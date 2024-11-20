// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package httpstate

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

//nolint:paralleltest // sets env var, must be run in isolation
func TestConsoleURL(t *testing.T) {
	//nolint:paralleltest // sets env var, must be run in isolation
	t.Run("HonorEnvVar", func(t *testing.T) {
		// Honor the PULUMI_CONSOLE_DOMAIN environment variable.
		t.Setenv("PULUMI_CONSOLE_DOMAIN", "khulnasoft-console.contoso.com")
		assert.Equal(t,
			"https://khulnasoft-console.contoso.com/1/2",
			cloudConsoleURL("https://api.khulnasoft.contoso.com", "1", "2"))

		// Unset the variable, confirm the "standard behavior" where we
		// replace "api." with "app.".
		os.Unsetenv("PULUMI_CONSOLE_DOMAIN")
		assert.Equal(t,
			"https://app.khulnasoft.contoso.com/1/2",
			cloudConsoleURL("https://api.khulnasoft.contoso.com", "1", "2"))
	})

	t.Run("CloudURLUsingStandardPattern", func(t *testing.T) {
		assert.Equal(t,
			"https://app.khulnasoft.com/khulnasoft-bot/my-stack",
			cloudConsoleURL("https://api.khulnasoft.com", "khulnasoft-bot", "my-stack"))

		assert.Equal(t,
			"http://app.khulnasoft.example.com/khulnasoft-bot/my-stack",
			cloudConsoleURL("http://api.khulnasoft.example.com", "khulnasoft-bot", "my-stack"))
	})

	t.Run("LocalDevelopment", func(t *testing.T) {
		assert.Equal(t,
			"http://localhost:3000/khulnasoft-bot/my-stack",
			cloudConsoleURL("http://localhost:8080", "khulnasoft-bot", "my-stack"))
	})

	t.Run("ConsoleDomainUnknown", func(t *testing.T) {
		assert.Equal(t, "", cloudConsoleURL("https://khulnasoft.example.com", "khulnasoft-bot", "my-stack"))
		assert.Equal(t, "", cloudConsoleURL("not-even-a-real-url", "khulnasoft-bot", "my-stack"))
	})
}
