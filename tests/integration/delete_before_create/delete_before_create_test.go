// Copyright 2018-2024, Pulumi Corporation.
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

//go:build (nodejs || all) && !xplatform_acceptance

package ints

import (
	"testing"

	"github.com/khulnasoft/khulnasoft/pkg/v3/testing/integration"
)

// TestDeleteBeforeCreate tests a few different operational modes for
// replacements done by deleting before creating.
//
//nolint:paralleltest // ProgramTest calls t.Parallel()
func TestDeleteBeforeCreate(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:          "step1",
		Dependencies: []string{"@khulnasoft/khulnasoft"},
		Quick:        true,
		EditDirs: []integration.EditDir{
			{
				Dir:      "step2",
				Additive: true,
			},
			{
				Dir:      "step3",
				Additive: true,
			},
			{
				Dir:      "step4",
				Additive: true,
			},
			{
				Dir:      "step5",
				Additive: true,
			},
			{
				Dir:      "step6",
				Additive: true,
			},
		},
	})
}