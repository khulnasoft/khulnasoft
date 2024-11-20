// Copyright 2022-2024, Pulumi Corporation.
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

package batchyaml

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/khulnasoft/khulnasoft/pkg/v3/codegen/dotnet"
	"github.com/khulnasoft/khulnasoft/pkg/v3/codegen/testing/test"
)

// This specifically tests the synced examples from khulnasoft/yaml with
// testing/test/testdata/transpiled_examples, as it requires a different SDK path in Check
func TestGenerateProgram(t *testing.T) {
	t.Parallel()

	rootDir, err := filepath.Abs(filepath.Join("..", "..", "..", "..", ".."))
	assert.NoError(t, err)

	test.GenerateDotnetYAMLBatchTest(t, rootDir, dotnet.GenerateProgram)
}