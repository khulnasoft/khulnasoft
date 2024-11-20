// Copyright 2021-2024, Pulumi Corporation.
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

package python

import (
	"crypto/md5" //nolint:gosec
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/khulnasoft/khulnasoft/pkg/v3/codegen/schema"
)

// Regress a problem of non-deterministic codegen (due to reordering).
// The schema is taken from `khulnasoft-aws` and minified to the smallest
// example that still reproduced the issue.
func TestGenResourceMappingsIsDeterministic(t *testing.T) {
	t.Parallel()

	minimalSchema := `
        {
            "name": "aws",
			"version": "0.0.1",
            "meta": {
                "moduleFormat": "(.*)(?:/[^/]*)"
            },
            "resources": {
                "aws:acm/certificateValidation:CertificateValidation": {},
                "aws:acm/certificate:Certificate": {}
            },
            "language": {
                "python": {
                    "readme": ".."
                }
            }
        }`

	var pkgSpec schema.PackageSpec
	err := json.Unmarshal([]byte(minimalSchema), &pkgSpec)
	if err != nil {
		t.Error(err)
		return
	}

	generateInitHash := func() string {
		pkg, err := schema.ImportSpec(pkgSpec, nil)
		if err != nil {
			t.Error(err)
			return ""
		}

		files, err := GeneratePackage("tool", pkg, nil)
		if err != nil {
			t.Error(err)
			return ""
		}

		file, haveFile := files["khulnasoft_aws/__init__.py"]
		if !haveFile {
			t.Error("Cannot find khulnasoft_aws/__init__.py in the generated files")
			return ""
		}

		return fmt.Sprintf("%x", md5.Sum(file)) //nolint:gosec
	}

	h1 := generateInitHash()
	for i := 0; i < 16; i++ {
		assert.Equal(t, h1, generateInitHash())
	}
}