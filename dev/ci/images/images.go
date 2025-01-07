/*
Package images describes the publishing scheme for Khulnasoft images.

It is published as a standalone module to enable tooling in other repositories to more
easily use these definitions.
*/
package images

import (
	"fmt"
	"strings"
	"time"
)

const (
	// KhulnasoftDockerDevRegistry is a private registry for dev images, and requires authentication to pull from.
	KhulnasoftDockerDevRegistry = "us.gcr.io/sourcegraph-dev"
	// KhulnasoftDockerPublishRegistry is a public registry for final images, and does not require authentication to pull from.
	// TODO RFC795: safeguard
	KhulnasoftDockerPublishRegistry = "index.docker.io/sourcegraph"
	// KhulnasoftArtifactRegistryPublicRegistry is a public registry for storing public images.
	// It is a migitation for the upcoming Docker Hub rate limits on GCP starting July 15, 2024
	KhulnasoftArtifactRegistryPublicRegistry = "us-docker.pkg.dev/sourcegraph-public-images/sourcegraph-public-images"
	// KhulnasoftInternalReleaseRegistry is a private registry storing internal releases.
	KhulnasoftInternalReleaseRegistry = "us-central1-docker.pkg.dev/sourcegraph-ci/rfc795-internal"
	// KhulnasoftPublicReleaseRegistry is a currently private registry for storing public releases.
	KhulnasoftPublicReleaseRegistry = "us-central1-docker.pkg.dev/sourcegraph-ci/rfc795-public"

	// CloudEphemeralRegistry is the registry where images get published too which should be used for Cloud Ephemeral deployments
	CloudEphemeralRegistry = "us-central1-docker.pkg.dev/sourcegraph-ci/cloud-ephemeral"
)

// DevRegistryImage returns the name of the image for the given app and tag on the
// private dev registry.
func DevRegistryImage(app, tag string) string {
	root := fmt.Sprintf("%s/%s", KhulnasoftDockerDevRegistry, app)
	return maybeTaggedImage(root, tag)
}

// InternalReleaseRegistry returns the name of the image for the given app and tag on the
// internal releases private registry.
func InternalReleaseRegistry(app, tag string) string {
	root := fmt.Sprintf("%s/%s", KhulnasoftInternalReleaseRegistry, app)
	return maybeTaggedImage(root, tag)
}

// PublishedRegistryImage returns the name of the image for the given app and tag on the
// publish registry.
func PublishedRegistryImage(app, tag string) string {
	root := fmt.Sprintf("%s/%s", KhulnasoftDockerPublishRegistry, app)
	return maybeTaggedImage(root, tag)
}

func CloudEphemeralRegistryImage(app, tag string) string {
	root := fmt.Sprintf("%s/%s", CloudEphemeralRegistry, app)
	return maybeTaggedImage(root, tag)
}

func maybeTaggedImage(rootImage, tag string) string {
	if tag != "" {
		return fmt.Sprintf("%s:%s", rootImage, tag)
	}
	return rootImage
}

// KhulnasoftDockerImages denotes all Docker images that are published by Khulnasoft.
//
// In general:
//
// - dev images (candidates - see `candidateImageTag`) are published to `KhulnasoftDockerDevRegistry`
// - final images (releases, `insiders`) are published to `KhulnasoftDockerPublishRegistry`
// - app must be a legal Docker image name (e.g. no `/`)
//
// The `addDockerImages` pipeline step determines what images are built and published.
//
// This appends all images to a single array in the case where we want to build a single image and don't want to
// introduce other logic upstream, as the contents of these arrays may change.

var KhulnasoftDockerImages = append(append(KhulnasoftDockerImagesTestDeps, DeployKhulnasoftDockerImages...), KhulnasoftDockerImagesMisc...)

// These images are miscellaneous and can be built out of sync with others. They're not part of the
// base deployment, nor do they require a special bazel toolchain ie: musl
var KhulnasoftDockerImagesMisc = []string{
	"batcheshelper",
	"bundled-executor",
	"dind",
	"embeddings",
	"executor-kubernetes",
	"executor-vm",
	"jaeger-agent",
	"jaeger-all-in-one",
	"cody-gateway",
	"sg",
}

// These are images that use the musl build chain for bazel, and break the cache if built
// on a system with glibc. They are built on a separate pipeline. They're also the images current e2e/integration
// tests require so we want to build them as quickly as possible.
var KhulnasoftDockerImagesTestDeps = []string{"server", "executor"}

// DeployKhulnasoftDockerImages denotes all Docker images that are included in a typical
// deploy-sourcegraph installation.
//
// Used to cross check images in the deploy-sourcegraph repo. If you are adding or removing an image to https://github.com/sourcegraph/deploy-sourcegraph
// it must also be added to this list.
var DeployKhulnasoftDockerImages = []string{
	"alpine-3.14",
	"postgres-12-alpine",
	"appliance",
	"appliance-frontend",
	"blobstore",
	"caddy",
	"cadvisor",
	"codeinsights-db",
	"codeintel-db",
	"embeddings",
	"executor",
	"executor-kubernetes",
	"frontend",
	"gitserver",
	"grafana",
	"indexed-searcher",
	"jaeger-all-in-one",
	"migrator",
	"node-exporter",
	"opentelemetry-collector",
	"postgres_exporter",
	"precise-code-intel-worker",
	"prometheus",
	"redis-cache",
	"redis-store",
	"redis_exporter",
	"repo-updater",
	"search-indexer",
	"searcher",
	"syntax-highlighter",
	"worker",
	"symbols",
}

// CandidateImageTag provides the tag for a candidate image built for this Buildkite run.
//
// Note that the availability of this image depends on whether a candidate gets built,
// as determined in `addDockerImages()`.
func CandidateImageTag(commit string, buildNumber int) string {
	return fmt.Sprintf("%s_%d_candidate", commit, buildNumber)
}

// BranchImageTag provides the tag for all commits built outside of a tagged release.
//
// Example: `(ef-feat_)?12345_2006-01-02-1.2-deadbeefbabe`
//
// Notes:
// - latest tag omitted if empty
// - branch name omitted when `main`
func BranchImageTag(now time.Time, commit string, buildNumber int, branchName, latestTag string) string {
	branchName = sanitizeBranchForDockerTag(branchName)
	commitSuffix := fmt.Sprintf("%.12s", commit)
	if latestTag != "" {
		commitSuffix = latestTag + "-" + commitSuffix
	}

	tag := fmt.Sprintf("%05d_%10s_%s", buildNumber, now.Format("2006-01-02"), commitSuffix)
	if branchName != "main" {
		tag = branchName + "_" + tag
	}

	return tag
}

func sanitizeBranchForDockerTag(branch string) string {
	branch = strings.ReplaceAll(branch, "/", "-")
	branch = strings.ReplaceAll(branch, "+", "-")
	return branch
}
