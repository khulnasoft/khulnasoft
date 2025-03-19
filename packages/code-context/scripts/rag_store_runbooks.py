from __future__ import annotations

import os
import sqlite3
import struct
import sys
from sqlite3 import Connection
from typing import List, Tuple

import sqlite_vec
import vertexai
from git import GitCommandError, Repo
from langchain_text_splitters import MarkdownTextSplitter
from vertexai.language_models import (TextEmbedding, TextEmbeddingInput,
                                      TextEmbeddingModel)

PROJECT_ID = "tkhandelwal-7239f9f8"
REGION = "us-central1"
MODEL_ID = "text-embedding-005"
TASK_TYPE = "RETRIEVAL_DOCUMENT"

DEFAULT_BRANCH = "master"
REPOSITORY_URL = "https://gitlab.com/gitlab-com/runbooks"
CLONE_PATH = "/var/folders/_t/jzjnv0ln08g0gt5cf61f9mj00000gn/T/tmp.qQ7pw5CAcJ/runbooks"


def serialize(vector: List[float]) -> bytes:
    """serializes a list of floats into a compact "raw bytes" format"""
    return struct.pack("%sf" % len(vector), *vector)

def init_db(db_path) -> Connection:
    # ref: https://github.com/asg017/sqlite-vec/blob/2abca78d359b7059f2967afd6894adb46d6dba30/examples/python-recipes/openai-sample.py#L30
    db: Connection = sqlite3.connect(db_path)
    db.enable_load_extension(True)
    sqlite_vec.load(db)
    db.enable_load_extension(False)
    return db

def create_embeddings_table(db: Connection):
    db.execute(
        """
        CREATE VIRTUAL TABLE IF NOT EXISTS vec_runbooks_docs USING vec0(
          id INTEGER PRIMARY KEY AUTO_INCREMENT,
          embedding FLOAT[768],
          source_link TEXT
         );
        """
    )

def setup_db(db_path):
    db = init_db(db_path)
    create_embeddings_table(db)
    return db


def find_md_files(repo_path) -> List[str]:
    """
    Crawls through the specified repository directory and finds all .md files.

    Args:
        repo_path (str): The path to the cloned repository.

    Returns:
        list: A list of paths to .md files.
    """
    md_files: List[str] = []

    for root, _, files in os.walk(repo_path):
        for file in files:
            if file.endswith(".md"):
                md_files.append(os.path.join(root, file))
    
    return md_files

def clone_or_update_repo(repo_url, clone_path, branch=DEFAULT_BRANCH):
    """
    Clones a Git repository to the specified path or updates it if it already exists.

    Args:
        repo_url (str): The URL of the Git repository.
        clone_path (str): The path where the repository will be cloned.

    Returns:
        str: The path to the repository.
    """
    if os.path.exists(clone_path):
        print(f"Repository already exists at {clone_path}. Pulling latest changes...")
        try:
            repo = Repo(clone_path)
            repo.remotes.origin.pull(branch)
            print("Repository updated successfully.")
        except GitCommandError as e:
            print(f"Error during git pull: {e}")
            return None
    else:
        print(f"Cloning repository from {repo_url} to {clone_path}...")
        try:
            repo = Repo.clone_from(repo_url, clone_path)
            print("Repository cloned successfully.")
        except GitCommandError as e:
            print(f"Error during git clone: {e}")
            return None
    return clone_path

def chunk_md_file(file_path, **kwargs) -> List[str]:
    """
    Reads and chunks a file using LangChain's text splitter.

    Args:
        file_path (str): Path to the Markdown file.
        chunk_size (int): Maximum size of each chunk.
        chunk_overlap (int): Overlap between chunks.

    Returns:
        list: List of text chunks.
    """
    with open(file_path, "r", encoding="utf-8") as file:
        content = file.read()

    markdown_splitter = MarkdownTextSplitter(**kwargs)
    chunks = markdown_splitter.split_text(content)
    return chunks

def setup_model() -> Tuple[TextEmbeddingModel, dict]:
    vertexai.init(project=PROJECT_ID, location=REGION)
    model = TextEmbeddingModel.from_pretrained(MODEL_ID)

    dimensionality = 768
    model_kwargs = dict(output_dimensionality=dimensionality)

    return model, model_kwargs

def get_title(file_path: str, repo_path: str) -> str:
    relpath = os.path.relpath(file_path, repo_path)
    title = relpath.replace("/", " ").rstrip(".md")
    return title

def get_file_raw_web_link(file_path: str, repo_path: str) -> str:
    relpath = os.path.relpath(file_path, repo_path)
    return REPOSITORY_URL.rstrip(".git") + f"/raw/{DEFAULT_BRANCH}/{relpath}"

def process_chunks(chunks: List[str], title: str, model: TextEmbeddingModel, kwargs: dict) -> List[TextEmbedding]:
    embeddings: List[TextEmbedding] = [] 
    if len(chunks) > 250:
        batch_size = 250
        for i in range(0, len(chunks), batch_size):
            batch = chunks[i:i+batch_size]
            batch_inputs = [TextEmbeddingInput(text=text, task_type=TASK_TYPE, title=title) for text in batch]
            embeddings += model.get_embeddings(batch_inputs, **kwargs)
    else:
        inputs = [TextEmbeddingInput(text=text, task_type=TASK_TYPE, title=title) for text in chunks]
        embeddings = model.get_embeddings(inputs, **kwargs)

    return embeddings

def persist_embeddings(db: Connection, embeddings: List[TextEmbedding], chunks: List[str], file_web_link: str):
    with db:
        for i in range(len(embeddings)):
            embedding_vector = serialize(embeddings[i].values)
            db.execute(
                """
                INSERT INTO vec_runbooks_docs (embedding, source_link)
                VALUES (?, ?)
                """,
                [embedding_vector, file_web_link],
            )

if __name__ == "__main__":
    model, kwargs = setup_model()
    repo_path = clone_or_update_repo(REPOSITORY_URL, CLONE_PATH)
    if repo_path is None:
        print("Error cloning or updating repository.")
        sys.exit(1)
    docs_dir = os.path.join(repo_path, "docs")
    markdown_files = find_md_files(docs_dir)
    db = setup_db(db_path="runbooks.db")
    for file in markdown_files: 
        title: str = get_title(file, repo_path)
        file_web_link: str = get_file_raw_web_link(file, repo_path)
        chunks: List[str] = chunk_md_file(file, chunk_size=200, chunk_overlap=20)
        embeddings: List[TextEmbedding] = process_chunks(chunks, title, model, kwargs)
        persist_embeddings(db, embeddings, chunks, file_web_link)
