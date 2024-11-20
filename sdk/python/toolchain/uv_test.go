// Copyright 2024, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package toolchain

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/blang/semver"
	"github.com/stretchr/testify/require"
)

func TestUvVirtualenvPath(t *testing.T) {
	t.Parallel()

	t.Run("no virtualenv specified", func(t *testing.T) {
		t.Parallel()
		root := t.TempDir()
		uv, err := newUv(root, "")
		require.NoError(t, err)
		require.Equal(t, filepath.Join(root, ".venv"), uv.virtualenvPath, "virtualenv is in the project root")
	})

	t.Run("no virtualenv specified, in a subfolder", func(t *testing.T) {
		t.Parallel()
		root := t.TempDir()
		khulnasoftRoot := filepath.Join(root, "subfolder")
		require.NoError(t, os.WriteFile(filepath.Join(root, "uv.lock"), []byte{}, 0o600))
		require.NoError(t, os.Mkdir(khulnasoftRoot, 0o755))

		uv, err := newUv(khulnasoftRoot, "")
		require.NoError(t, err)
		require.Equal(t, filepath.Join(root, ".venv"), uv.virtualenvPath, "virtualenv is next to uv.lock")
	})

	t.Run("no virtualenv specified, in a subfolder", func(t *testing.T) {
		t.Parallel()
		root := t.TempDir()
		khulnasoftRoot := filepath.Join(root, "subfolder")
		require.NoError(t, os.Mkdir(khulnasoftRoot, 0o700))
		require.NoError(t, os.WriteFile(filepath.Join(root, "uv.lock"), []byte{}, 0o600))
		require.NoError(t, os.WriteFile(filepath.Join(khulnasoftRoot, "uv.lock"), []byte{}, 0o600))

		uv, err := newUv(khulnasoftRoot, "")
		require.NoError(t, err)
		require.Equal(t, filepath.Join(khulnasoftRoot, ".venv"), uv.virtualenvPath,
			"virtualenv is next to the uv.lock closest to the project root")
	})

	t.Run("virtualenv option is provided", func(t *testing.T) {
		t.Parallel()
		root := t.TempDir()

		uv, err := newUv(root, "banana")
		require.NoError(t, err)
		require.Equal(t, filepath.Join(root, "banana"), uv.virtualenvPath, "virtualenv is in the project root")
	})

	t.Run("virtualenv option is provided, in  subfolder", func(t *testing.T) {
		t.Parallel()
		root := t.TempDir()
		khulnasoftRoot := filepath.Join(root, "subfolder")
		require.NoError(t, os.Mkdir(khulnasoftRoot, 0o755))

		uv, err := newUv(khulnasoftRoot, "banana")
		require.NoError(t, err)
		require.Equal(t, filepath.Join(khulnasoftRoot, "banana"), uv.virtualenvPath, "virtualenv is in the project root")
	})
}

func TestUvVersion(t *testing.T) {
	t.Parallel()

	uv, err := newUv(".", "")
	require.NoError(t, err)

	for _, versionString := range []string{
		"uv 0.4.26",
		"uv 0.4.26 (Homebrew 2024-10-23)",
		"uv 0.4.26 (d2cd09bbd 2024-10-25)",
	} {
		v, err := uv.uvVersion(versionString)
		require.NoError(t, err)
		require.Equal(t, semver.MustParse("0.4.26"), v)
	}

	_, err = uv.uvVersion("uv 0.4.25")
	require.ErrorContains(t, err, "less than the minimum required version")
}
