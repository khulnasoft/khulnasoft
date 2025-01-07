package autoindexing

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/khulnasoft/khulnasoft/internal/api"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/autoindexing/internal/jobselector"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/autoindexing/shared"
	"github.com/khulnasoft/khulnasoft/internal/codeintel/dependencies"
	uploadsshared "github.com/khulnasoft/khulnasoft/internal/codeintel/uploads/shared"
	"github.com/khulnasoft/khulnasoft/internal/database/dbmocks"
	"github.com/khulnasoft/khulnasoft/internal/gitserver"
	"github.com/khulnasoft/khulnasoft/internal/observation"
	internaltypes "github.com/khulnasoft/khulnasoft/internal/types"
	"github.com/khulnasoft/khulnasoft/lib/codeintel/autoindex/config"
)

func init() {
	jobselector.MaximumIndexJobsPerInferredConfiguration = 50
}

func TestQueueAutoIndexJobsExplicit(t *testing.T) {
	conf := `{
		"index_jobs": [
			{
				"steps": [
					{
						// Comments are the future
						"image": "go:latest",
						"commands": ["go mod vendor"],
					}
				],
				"indexer": "lsif-go",
				"indexer_args": ["--no-animation"],
			},
			{
				"root": "web/",
				"indexer": "scip-typescript",
				"indexer_args": ["index", "--no-progress-bar"],
				"outfile": "lsif.dump",
			},
		]
	}`

	mockDBStore := NewMockStore()
	mockDBStore.InsertJobsFunc.SetDefaultHook(func(ctx context.Context, jobs []uploadsshared.AutoIndexJob) ([]uploadsshared.AutoIndexJob, error) {
		return jobs, nil
	})
	mockDBStore.RepositoryExceptionsFunc.SetDefaultReturn(true, true, nil)

	mockGitserverClient := gitserver.NewMockClient()
	mockGitserverClient.ResolveRevisionFunc.SetDefaultHook(func(ctx context.Context, repo api.RepoName, rev string, opts gitserver.ResolveRevisionOptions) (api.CommitID, error) {
		return api.CommitID(fmt.Sprintf("c%s", repo)), nil
	})

	inferenceService := NewMockInferenceService()

	service := newService(
		observation.TestContextTB(t),
		mockDBStore,
		inferenceService,
		defaultMockRepoStore(), // repoStore
		mockGitserverClient,
	)
	_, _ = service.QueueAutoIndexJobs(context.Background(), 42, "HEAD", conf, false, false)

	if len(mockDBStore.IsQueuedFunc.History()) != 1 {
		t.Errorf("unexpected number of calls to IsQueued. want=%d have=%d", 1, len(mockDBStore.IsQueuedFunc.History()))
	} else {
		var commits []string
		for _, call := range mockDBStore.IsQueuedFunc.History() {
			commits = append(commits, call.Arg2)
		}
		sort.Strings(commits)

		if diff := cmp.Diff([]string{"cr42"}, commits); diff != "" {
			t.Errorf("unexpected commits (-want +got):\n%s", diff)
		}
	}

	var jobs []uploadsshared.AutoIndexJob
	for _, call := range mockDBStore.InsertJobsFunc.History() {
		jobs = append(jobs, call.Result0...)
	}

	expectedIndexes := []uploadsshared.AutoIndexJob{
		{
			RepositoryID: 42,
			Commit:       "cr42",
			State:        "queued",
			DockerSteps: []uploadsshared.DockerStep{
				{
					Image:    "go:latest",
					Commands: []string{"go mod vendor"},
				},
			},
			Indexer:     "lsif-go",
			IndexerArgs: []string{"--no-animation"},
		},
		{
			RepositoryID: 42,
			Commit:       "cr42",
			State:        "queued",
			DockerSteps:  nil,
			Root:         "web/",
			Indexer:      "scip-typescript",
			IndexerArgs:  []string{"index", "--no-progress-bar"},
			Outfile:      "lsif.dump",
		},
	}
	if diff := cmp.Diff(expectedIndexes, jobs, cmpopts.EquateEmpty()); diff != "" {
		t.Errorf("unexpected indexes (-want +got):\n%s", diff)
	}
}

func TestQueueAutoIndexJobsInDatabase(t *testing.T) {
	indexConfiguration := shared.IndexConfiguration{
		ID:           1,
		RepositoryID: 42,
		Data: []byte(`{
			"index_jobs": [
				{
					"steps": [
						{
							// Comments are the future
							"image": "go:latest",
							"commands": ["go mod vendor"],
						}
					],
					"indexer": "lsif-go",
					"indexer_args": ["--no-animation"],
				},
				{
					"root": "web/",
					"indexer": "scip-typescript",
					"indexer_args": ["index", "--no-progress-bar"],
					"outfile": "lsif.dump",
				},
			]
		}`),
	}

	mockDBStore := NewMockStore()
	mockDBStore.InsertJobsFunc.SetDefaultHook(func(ctx context.Context, jobs []uploadsshared.AutoIndexJob) ([]uploadsshared.AutoIndexJob, error) {
		return jobs, nil
	})
	mockDBStore.GetIndexConfigurationByRepositoryIDFunc.SetDefaultReturn(indexConfiguration, true, nil)
	mockDBStore.RepositoryExceptionsFunc.SetDefaultReturn(true, true, nil)

	mockGitserverClient := gitserver.NewMockClient()
	mockGitserverClient.ResolveRevisionFunc.SetDefaultHook(func(ctx context.Context, repo api.RepoName, rev string, opts gitserver.ResolveRevisionOptions) (api.CommitID, error) {
		return api.CommitID(fmt.Sprintf("c%s", repo)), nil
	})
	inferenceService := NewMockInferenceService()

	service := newService(
		observation.TestContextTB(t),
		mockDBStore,
		inferenceService,
		defaultMockRepoStore(), // repoStore
		mockGitserverClient,
	)
	_, _ = service.QueueAutoIndexJobs(context.Background(), 42, "HEAD", "", false, false)

	if len(mockDBStore.GetIndexConfigurationByRepositoryIDFunc.History()) != 1 {
		t.Errorf("unexpected number of calls to GetIndexConfigurationByRepositoryID. want=%d have=%d", 1, len(mockDBStore.GetIndexConfigurationByRepositoryIDFunc.History()))
	} else {
		var repositoryIDs []int
		for _, call := range mockDBStore.GetIndexConfigurationByRepositoryIDFunc.History() {
			repositoryIDs = append(repositoryIDs, call.Arg1)
		}
		sort.Ints(repositoryIDs)

		if diff := cmp.Diff([]int{42}, repositoryIDs); diff != "" {
			t.Errorf("unexpected repository identifiers (-want +got):\n%s", diff)
		}
	}

	if len(mockDBStore.IsQueuedFunc.History()) != 1 {
		t.Errorf("unexpected number of calls to IsQueued. want=%d have=%d", 1, len(mockDBStore.IsQueuedFunc.History()))
	} else {
		var commits []string
		for _, call := range mockDBStore.IsQueuedFunc.History() {
			commits = append(commits, call.Arg2)
		}
		sort.Strings(commits)

		if diff := cmp.Diff([]string{"cr42"}, commits); diff != "" {
			t.Errorf("unexpected commits (-want +got):\n%s", diff)
		}
	}

	var jobs []uploadsshared.AutoIndexJob
	for _, call := range mockDBStore.InsertJobsFunc.History() {
		jobs = append(jobs, call.Result0...)
	}

	expectedIndexes := []uploadsshared.AutoIndexJob{
		{
			RepositoryID: 42,
			Commit:       "cr42",
			State:        "queued",
			DockerSteps: []uploadsshared.DockerStep{
				{
					Image:    "go:latest",
					Commands: []string{"go mod vendor"},
				},
			},
			Indexer:     "lsif-go",
			IndexerArgs: []string{"--no-animation"},
		},
		{
			RepositoryID: 42,
			Commit:       "cr42",
			State:        "queued",
			DockerSteps:  nil,
			Root:         "web/",
			Indexer:      "scip-typescript",
			IndexerArgs:  []string{"index", "--no-progress-bar"},
			Outfile:      "lsif.dump",
		},
	}
	if diff := cmp.Diff(expectedIndexes, jobs, cmpopts.EquateEmpty()); diff != "" {
		t.Errorf("unexpected indexes (-want +got):\n%s", diff)
	}
}

var yamlIndexConfiguration = []byte(`
index_jobs:
  -
    steps:
      - image: go:latest
        commands:
          - go mod vendor
    indexer: lsif-go
    indexer_args:
      - --no-animation
  -
    root: web/
    indexer: scip-typescript
    indexer_args: ['index', '--no-progress-bar']
    outfile: lsif.dump
`)

func TestQueueAutoIndexJobsInRepository(t *testing.T) {
	mockDBStore := NewMockStore()
	mockDBStore.InsertJobsFunc.SetDefaultHook(func(ctx context.Context, jobs []uploadsshared.AutoIndexJob) ([]uploadsshared.AutoIndexJob, error) {
		return jobs, nil
	})
	mockDBStore.RepositoryExceptionsFunc.SetDefaultReturn(true, true, nil)

	gitserverClient := gitserver.NewMockClient()
	gitserverClient.ResolveRevisionFunc.SetDefaultHook(func(ctx context.Context, repo api.RepoName, rev string, opts gitserver.ResolveRevisionOptions) (api.CommitID, error) {
		return api.CommitID(fmt.Sprintf("c%s", repo)), nil
	})
	gitserverClient.NewFileReaderFunc.SetDefaultReturn(io.NopCloser(bytes.NewReader(yamlIndexConfiguration)), nil)
	inferenceService := NewMockInferenceService()

	service := newService(
		observation.TestContextTB(t),
		mockDBStore,
		inferenceService,
		defaultMockRepoStore(), // repoStore
		gitserverClient,
	)

	if _, err := service.QueueAutoIndexJobs(context.Background(), 42, "HEAD", "", false, false); err != nil {
		t.Fatalf("unexpected error performing update: %s", err)
	}

	if len(mockDBStore.IsQueuedFunc.History()) != 1 {
		t.Errorf("unexpected number of calls to IsQueued. want=%d have=%d", 1, len(mockDBStore.IsQueuedFunc.History()))
	} else {
		var commits []string
		for _, call := range mockDBStore.IsQueuedFunc.History() {
			commits = append(commits, call.Arg2)
		}
		sort.Strings(commits)

		if diff := cmp.Diff([]string{"cr42"}, commits); diff != "" {
			t.Errorf("unexpected commits (-want +got):\n%s", diff)
		}
	}

	var jobs []uploadsshared.AutoIndexJob
	for _, call := range mockDBStore.InsertJobsFunc.History() {
		jobs = append(jobs, call.Result0...)
	}

	expectedIndexes := []uploadsshared.AutoIndexJob{
		{
			RepositoryID: 42,
			Commit:       "cr42",
			State:        "queued",
			DockerSteps: []uploadsshared.DockerStep{
				{
					Image:    "go:latest",
					Commands: []string{"go mod vendor"},
				},
			},
			Indexer:     "lsif-go",
			IndexerArgs: []string{"--no-animation"},
		},
		{
			RepositoryID: 42,
			Commit:       "cr42",
			State:        "queued",
			DockerSteps:  nil,
			Root:         "web/",
			Indexer:      "scip-typescript",
			IndexerArgs:  []string{"index", "--no-progress-bar"},
			Outfile:      "lsif.dump",
		},
	}
	if diff := cmp.Diff(expectedIndexes, jobs, cmpopts.EquateEmpty()); diff != "" {
		t.Errorf("unexpected indexes (-want +got):\n%s", diff)
	}
}

func TestQueueAutoIndexJobsInferred(t *testing.T) {
	mockDBStore := NewMockStore()
	mockDBStore.InsertJobsFunc.SetDefaultHook(func(ctx context.Context, jobs []uploadsshared.AutoIndexJob) ([]uploadsshared.AutoIndexJob, error) {
		return jobs, nil
	})
	mockDBStore.RepositoryExceptionsFunc.SetDefaultReturn(true, true, nil)

	gitserverClient := gitserver.NewMockClient()
	gitserverClient.ResolveRevisionFunc.SetDefaultHook(func(ctx context.Context, repo api.RepoName, rev string, opts gitserver.ResolveRevisionOptions) (api.CommitID, error) {
		return api.CommitID(fmt.Sprintf("c%s", repo)), nil
	})
	gitserverClient.NewFileReaderFunc.SetDefaultReturn(nil, os.ErrNotExist)

	inferenceService := NewMockInferenceService()
	inferenceService.InferIndexJobsFunc.SetDefaultHook(func(ctx context.Context, rn api.RepoName, s1, s2 string) (*shared.InferenceResult, error) {
		switch string(rn) {
		case "r42":
			return &shared.InferenceResult{IndexJobs: []config.AutoIndexJobSpec{{Root: ""}}}, nil
		case "r44":
			return &shared.InferenceResult{IndexJobs: []config.AutoIndexJobSpec{{Root: "a"}, {Root: "b"}}}, nil
		default:
			return &shared.InferenceResult{IndexJobs: nil}, nil
		}
	})

	service := newService(
		observation.TestContextTB(t),
		mockDBStore,
		inferenceService,
		defaultMockRepoStore(), // repoStore
		gitserverClient,
	)

	for _, id := range []int{41, 42, 43, 44} {
		if _, err := service.QueueAutoIndexJobs(context.Background(), id, "HEAD", "", false, false); err != nil {
			t.Fatalf("unexpected error performing update: %s", err)
		}
	}

	indexRoots := map[int][]string{}
	for _, call := range mockDBStore.InsertJobsFunc.History() {
		for _, index := range call.Result0 {
			indexRoots[index.RepositoryID] = append(indexRoots[index.RepositoryID], index.Root)
		}
	}

	expectedIndexRoots := map[int][]string{
		42: {""},
		44: {"a", "b"},
	}
	if diff := cmp.Diff(expectedIndexRoots, indexRoots); diff != "" {
		t.Errorf("unexpected indexes (-want +got):\n%s", diff)
	}

	if len(mockDBStore.IsQueuedFunc.History()) != 4 {
		t.Errorf("unexpected number of calls to IsQueued. want=%d have=%d", 4, len(mockDBStore.IsQueuedFunc.History()))
	} else {
		var commits []string
		for _, call := range mockDBStore.IsQueuedFunc.History() {
			commits = append(commits, call.Arg2)
		}
		sort.Strings(commits)

		if diff := cmp.Diff([]string{"cr41", "cr42", "cr43", "cr44"}, commits); diff != "" {
			t.Errorf("unexpected commits (-want +got):\n%s", diff)
		}
	}
}

func TestQueueAutoIndexJobsForPackage(t *testing.T) {
	mockDBStore := NewMockStore()
	mockDBStore.InsertJobsFunc.SetDefaultHook(func(ctx context.Context, jobs []uploadsshared.AutoIndexJob) ([]uploadsshared.AutoIndexJob, error) {
		return jobs, nil
	})
	mockDBStore.IsQueuedFunc.SetDefaultReturn(false, nil)
	mockDBStore.RepositoryExceptionsFunc.SetDefaultReturn(true, true, nil)

	gitserverClient := gitserver.NewMockClient()
	gitserverClient.ResolveRevisionFunc.SetDefaultHook(func(ctx context.Context, repo api.RepoName, versionString string, opts gitserver.ResolveRevisionOptions) (api.CommitID, error) {
		if repo != "r42" && versionString != "4e7eeb0f8a96" {
			t.Errorf("unexpected (repoID, versionString) (%v, %v) supplied to EnqueueRepoUpdate", repo, versionString)
		}
		return "c42", nil
	})
	gitserverClient.NewFileReaderFunc.SetDefaultReturn(nil, os.ErrNotExist)

	inferenceService := NewMockInferenceService()
	inferenceService.InferIndexJobsFunc.SetDefaultHook(func(ctx context.Context, rn api.RepoName, s1, s2 string) (*shared.InferenceResult, error) {
		return &shared.InferenceResult{
			IndexJobs: []config.AutoIndexJobSpec{
				{
					Root: "",
					Steps: []config.DockerStep{
						{
							Image:    "sourcegraph/lsif-go:latest",
							Commands: []string{"go mod download"},
						},
					},
					Indexer:     "sourcegraph/lsif-go:latest",
					IndexerArgs: []string{"lsif-go", "--no-animation"},
				},
			},
		}, nil
	})

	mockRepoStore := defaultMockRepoStore()
	mockRepoStore.GetByNameFunc.SetDefaultHook(func(ctx context.Context, repoName api.RepoName) (*internaltypes.Repo, error) {
		if repoName != "github.com/khulnasoft/khulnasoft" {
			t.Errorf("unexpected repo %v supplied to EnqueueRepoUpdate", repoName)
		}
		return &internaltypes.Repo{ID: 42, Name: "github.com/khulnasoft/khulnasoft"}, nil
	})

	service := newService(
		observation.TestContextTB(t),
		mockDBStore,
		inferenceService,
		mockRepoStore, // repoStore
		gitserverClient,
	)

	_ = service.QueueAutoIndexJobsForPackage(context.Background(), dependencies.MinimialVersionedPackageRepo{
		Scheme:  "gomod",
		Name:    "https://github.com/khulnasoft/khulnasoft",
		Version: "v3.26.0-4e7eeb0f8a96",
	})

	if len(mockDBStore.IsQueuedFunc.History()) != 1 {
		t.Errorf("unexpected number of calls to IsQueued. want=%d have=%d", 1, len(mockDBStore.IsQueuedFunc.History()))
	} else {
		var commits []string
		for _, call := range mockDBStore.IsQueuedFunc.History() {
			commits = append(commits, call.Arg2)
		}
		sort.Strings(commits)

		if diff := cmp.Diff([]string{"c42"}, commits); diff != "" {
			t.Errorf("unexpected commits (-want +got):\n%s", diff)
		}
	}

	if len(mockDBStore.InsertJobsFunc.History()) != 1 {
		t.Errorf("unexpected number of calls to InsertJobs. want=%d have=%d", 1, len(mockDBStore.InsertJobsFunc.History()))
	} else {
		var jobs []uploadsshared.AutoIndexJob
		for _, call := range mockDBStore.InsertJobsFunc.History() {
			jobs = append(jobs, call.Result0...)
		}

		expectedIndexes := []uploadsshared.AutoIndexJob{
			{
				RepositoryID: 42,
				Commit:       "c42",
				State:        "queued",
				DockerSteps: []uploadsshared.DockerStep{
					{
						Image:    "sourcegraph/lsif-go:latest",
						Commands: []string{"go mod download"},
					},
				},
				Indexer:     "sourcegraph/lsif-go:latest",
				IndexerArgs: []string{"lsif-go", "--no-animation"},
			},
		}
		if diff := cmp.Diff(expectedIndexes, jobs, cmpopts.EquateEmpty()); diff != "" {
			t.Errorf("unexpected indexes (-want +got):\n%s", diff)
		}
	}
}

func defaultMockRepoStore() *dbmocks.MockRepoStore {
	repoStore := dbmocks.NewMockRepoStore()
	repoStore.GetFunc.SetDefaultHook(func(ctx context.Context, id api.RepoID) (*internaltypes.Repo, error) {
		return &internaltypes.Repo{
			ID:   id,
			Name: api.RepoName(fmt.Sprintf("r%d", id)),
		}, nil
	})
	return repoStore
}
