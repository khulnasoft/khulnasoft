import logging
from os import getenv

from khulnasoft import Sandbox

KHULNASOFT_API_KEY = getenv("KHULNASOFT_API_KEY")

# Global logging configuration
logging.basicConfig(level=logging.INFO, format="GLOBAL - [%(asctime)s] - %(name)-32s - %(levelname)7s: %(message)s",
                    datefmt="%Y-%m-%d %H:%M:%S")  # $HighlightLine

# Or configure only khulnasoft logger

# Get khulnasoft logger
khulnasoft_logger = logging.getLogger("khulnasoft")  # $HighlightLine

# Set khulnasoft logger level to INFO
khulnasoft_logger.setLevel(logging.INFO)  # $HighlightLine

# Setup formatter
formatter = logging.Formatter("KhulnaSoft    - [%(asctime)s] - %(name)-32s - %(levelname)7s: %(message)s",
                              datefmt="%Y-%m-%d %H:%M:%S")

# Setup handler
handler = logging.StreamHandler()
handler.setFormatter(formatter)

# Add handler to khulnasoft logger
khulnasoft_logger.addHandler(handler)  # $HighlightLine


def main():
    sandbox = Sandbox(template="base", api_key=KHULNASOFT_API_KEY)
    sandbox.filesystem.write("test.txt", "Hello World")
    sandbox.close()


main()
