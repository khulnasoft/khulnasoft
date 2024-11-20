// Copyright 2020-2024, Pulumi Corporation.
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

package nodejs

import (
	"testing"

	"github.com/khulnasoft/khulnasoft/pkg/v3/codegen/pcl"
	"github.com/khulnasoft/khulnasoft/pkg/v3/codegen/schema"
	"github.com/khulnasoft/khulnasoft/pkg/v3/codegen/testing/test"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/common/workspace"
	"github.com/stretchr/testify/assert"
)

func TestGenerateProgramVersionSelection(t *testing.T) {
	t.Parallel()

	test.GenerateNodeJSProgramTest(
		t,
		GenerateProgram,
		func(
			directory string, project workspace.Project, program *pcl.Program, localDependencies map[string]string,
		) error {
			return GenerateProject(directory, project, program, localDependencies, false)
		},
	)
}

func TestEnumReferencesCorrectIdentifier(t *testing.T) {
	t.Parallel()
	s := &schema.Package{
		Name: "khulnasoftservice",
		Language: map[string]interface{}{
			"nodejs": NodePackageInfo{
				PackageName: "@khulnasoft/bar",
			},
		},
	}
	result, err := enumNameWithPackage("khulnasoftservice:index:WebhookFilters", s.Reference())
	assert.NoError(t, err)
	assert.Equal(t, "khulnasoftservice.WebhookFilters", result)

	// These are redundant, but serve to clarify our expectations around package alias names.
	assert.NotEqual(t, "bar.WebhookFilters", result)
	assert.NotEqual(t, "@khulnasoft/bar.WebhookFilters", result)
}
