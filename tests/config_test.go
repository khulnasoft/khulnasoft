// Copyright 2016-2021, Pulumi Corporation.
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

package tests

import (
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/khulnasoft/khulnasoft/pkg/v3/testing/integration"
	ptesting "github.com/khulnasoft/khulnasoft/sdk/v3/go/common/testing"
)

const khulnasoftTestOrg = "moolumi"

func TestConfigCommands(t *testing.T) {
	t.Parallel()

	t.Run("SanityTest", func(t *testing.T) {
		t.Parallel()

		e := ptesting.NewEnvironment(t)
		defer e.DeleteIfNotFailed()

		integration.CreateBasicPulumiRepo(e)
		e.SetBackend(e.LocalURL())
		stackName := ptesting.RandomStackName()
		e.RunCommand("khulnasoft", "stack", "init", stackName)

		// check config is empty
		stdout, _ := e.RunCommand("khulnasoft", "config")
		assert.Equal(t, "KEY  VALUE", strings.Trim(stdout, "\r\n"))

		// set a bunch of config
		e.RunCommand("khulnasoft", "config", "set-all",
			"--plaintext", "key1=value1",
			"--plaintext", "outer.inner=value2",
			"--secret", "my_token=my_secret_token",
			"--plaintext", "myList[0]=foo")

		// check that it all exists
		stdout, _ = e.RunCommand("khulnasoft", "config", "get", "key1")
		assert.Equal(t, "value1", strings.Trim(stdout, "\r\n"))

		stdout, _ = e.RunCommand("khulnasoft", "config", "get", "outer.inner")
		assert.Equal(t, "value2", strings.Trim(stdout, "\r\n"))

		stdout, _ = e.RunCommand("khulnasoft", "config", "get", "my_token")
		assert.Equal(t, "my_secret_token", strings.Trim(stdout, "\r\n"))

		stdout, _ = e.RunCommand("khulnasoft", "config", "get", "myList[0]")
		assert.Equal(t, "foo", strings.Trim(stdout, "\r\n"))

		// check that the nested config does not exist because we didn't use path
		_, stderr := e.RunCommandExpectError("khulnasoft", "config", "get", "outer")
		assert.Equal(t,
			"error: configuration key 'outer' not found for stack '"+stackName+"'",
			strings.Trim(stderr, "\r\n"))

		_, stderr = e.RunCommandExpectError("khulnasoft", "config", "get", "myList")
		assert.Equal(t,
			"error: configuration key 'myList' not found for stack '"+stackName+"'",
			strings.Trim(stderr, "\r\n"))

		// set the nested config using --path
		e.RunCommand("khulnasoft", "config", "set-all", "--path",
			"--plaintext", "outer.inner=value2",
			"--plaintext", "myList[0]=foo")

		// check that the nested config now exists
		stdout, _ = e.RunCommand("khulnasoft", "config", "get", "outer")
		assert.Equal(t, "{\"inner\":\"value2\"}", strings.Trim(stdout, "\r\n"))

		stdout, _ = e.RunCommand("khulnasoft", "config", "get", "myList")
		assert.Equal(t, "[\"foo\"]", strings.Trim(stdout, "\r\n"))

		// remove the nested config values
		e.RunCommand("khulnasoft", "config", "rm-all", "--path", "outer.inner", "myList[0]")

		// check that it worked
		stdout, _ = e.RunCommand("khulnasoft", "config", "get", "outer")
		assert.Equal(t, "{}", strings.Trim(stdout, "\r\n"))

		stdout, _ = e.RunCommand("khulnasoft", "config", "get", "myList")
		assert.Equal(t, "[]", strings.Trim(stdout, "\r\n"))

		// remove other config values
		e.RunCommand("khulnasoft", "config", "rm-all",
			"outer.inner", "myList[0]", "outer", "myList", "key1", "my_token")

		// check that you can add keys with '=' in it
		e.RunCommand("khulnasoft", "config", "set-all",
			"--plaintext", "\"foo=\"=value2", "--plaintext", "'=some-weird=key='=value3")

		// check that they registered correctly
		stdout, _ = e.RunCommand("khulnasoft", "config", "get", "foo=")
		assert.Equal(t, "value2", strings.Trim(stdout, "\r\n"))

		stdout, _ = e.RunCommand("khulnasoft", "config", "get", "=some-weird=key=")
		assert.Equal(t, "value3", strings.Trim(stdout, "\r\n"))

		// remove the config key
		e.RunCommand("khulnasoft", "config", "rm-all", "foo=", "=some-weird=key=")

		// check config is empty again
		stdout, _ = e.RunCommand("khulnasoft", "config")
		assert.Equal(t, "KEY  VALUE", strings.Trim(stdout, "\r\n"))

		e.RunCommand("khulnasoft", "stack", "rm", "--yes")
	})

	t.Run("YAMLTest", func(t *testing.T) {
		t.Parallel()

		e := ptesting.NewEnvironment(t)
		defer e.DeleteIfNotFailed()

		integration.CreateBasicPulumiRepo(e)
		e.SetBackend(e.LocalURL())
		stackName := ptesting.RandomStackName()
		e.RunCommand("khulnasoft", "stack", "init", stackName)

		// check config is empty
		stdout, _ := e.RunCommand("khulnasoft", "config")
		assert.Equal(t, "KEY  VALUE", strings.Trim(stdout, "\r\n"))

		// set a config and verify that the generated config file matches the expected values
		e.RunCommand("khulnasoft", "config", "set",
			"a", "A",
		)
		expected := `^encryptionsalt: \S*
config:
  khulnasoft-test:a: A
$`
		b, err := os.ReadFile(filepath.Join(e.CWD, "Pulumi."+stackName+".yaml"))
		assert.NoError(t, err)
		assert.Regexp(t, regexp.MustCompile(expected), string(b))

		// set an additional secret config and verify that the generated config file matches the expected values
		e.RunCommand("khulnasoft", "config", "set",
			"b", "B", "--secret",
		)
		expected = `^encryptionsalt: \S*
config:
  khulnasoft-test:a: A
  khulnasoft-test:b:
    secure: \S*
$`
		b, err = os.ReadFile(filepath.Join(e.CWD, "Pulumi."+stackName+".yaml"))
		assert.NoError(t, err)
		assert.Regexp(t, regexp.MustCompile(expected), string(b))

		// update a config and verify that the generated config file matches the expected values
		e.RunCommand("khulnasoft", "config", "set",
			"a", "AA",
		)
		expected = `^encryptionsalt: \S*
config:
  khulnasoft-test:a: AA
  khulnasoft-test:b:
    secure: \S*
$`
		b, err = os.ReadFile(filepath.Join(e.CWD, "Pulumi."+stackName+".yaml"))
		assert.NoError(t, err)
		assert.Regexp(t, regexp.MustCompile(expected), string(b))

		// update the secret config and verify that the generated config file matches the expected values
		e.RunCommand("khulnasoft", "config", "set",
			"b", "BB", "--secret",
		)
		expected = `^encryptionsalt: \S*
config:
  khulnasoft-test:a: AA
  khulnasoft-test:b:
    secure: \S*
$`
		b, err = os.ReadFile(filepath.Join(e.CWD, "Pulumi."+stackName+".yaml"))
		assert.NoError(t, err)
		assert.Regexp(t, regexp.MustCompile(expected), string(b))

		// set a config with path=true and verify that the generated config file matches the expected values
		e.RunCommand("khulnasoft", "config", "set",
			"--path",
			"c", "C",
		)
		expected = `^encryptionsalt: \S*
config:
  khulnasoft-test:a: AA
  khulnasoft-test:b:
    secure: \S*
  khulnasoft-test:c: C
$`
		b, err = os.ReadFile(filepath.Join(e.CWD, "Pulumi."+stackName+".yaml"))
		assert.NoError(t, err)
		assert.Regexp(t, regexp.MustCompile(expected), string(b))

		// set a nested config and verify that the generated config file matches the expected values
		e.RunCommand("khulnasoft", "config", "set",
			"--path",
			"d.a", "D",
		)
		expected = `^encryptionsalt: \S*
config:
  khulnasoft-test:a: AA
  khulnasoft-test:b:
    secure: \S*
  khulnasoft-test:c: C
  khulnasoft-test:d:
    a: D
$`
		b, err = os.ReadFile(filepath.Join(e.CWD, "Pulumi."+stackName+".yaml"))
		assert.NoError(t, err)
		assert.Regexp(t, regexp.MustCompile(expected), string(b))

		// set an array config and verify that the generated config file matches the expected values
		e.RunCommand("khulnasoft", "config", "set",
			"--path",
			"e[0]", "E",
		)
		expected = `^encryptionsalt: \S*
config:
  khulnasoft-test:a: AA
  khulnasoft-test:b:
    secure: \S*
  khulnasoft-test:c: C
  khulnasoft-test:d:
    a: D
  khulnasoft-test:e:
    - E
$`
		b, err = os.ReadFile(filepath.Join(e.CWD, "Pulumi."+stackName+".yaml"))
		assert.NoError(t, err)
		assert.Regexp(t, regexp.MustCompile(expected), string(b))

		// set a nested array config and verify that the generated config file matches the expected values
		e.RunCommand("khulnasoft", "config", "set",
			"--path",
			"f.g[0]", "F",
		)
		expected = `^encryptionsalt: \S*
config:
  khulnasoft-test:a: AA
  khulnasoft-test:b:
    secure: \S*
  khulnasoft-test:c: C
  khulnasoft-test:d:
    a: D
  khulnasoft-test:e:
    - E
  khulnasoft-test:f:
    g:
      - F
$`
		b, err = os.ReadFile(filepath.Join(e.CWD, "Pulumi."+stackName+".yaml"))
		assert.NoError(t, err)
		assert.Regexp(t, regexp.MustCompile(expected), string(b))

		e.RunCommand("khulnasoft", "stack", "rm", "--yes")
	})
}

func TestBasicConfigGetRetrievedValueFromProject(t *testing.T) {
	t.Parallel()

	e := ptesting.NewEnvironment(t)
	defer e.DeleteIfNotFailed()

	khulnasoftProject := `
name: khulnasoft-test
runtime: go
config:
  first-value:
    type: string
    default: first`

	integration.CreatePulumiRepo(e, khulnasoftProject)
	e.SetBackend(e.LocalURL())
	stackName := ptesting.RandomStackName()
	e.RunCommand("khulnasoft", "stack", "init", stackName)
	stdout, _ := e.RunCommand("khulnasoft", "config", "get", "first-value")
	assert.Equal(t, "first", strings.Trim(stdout, "\r\n"))
}

func TestConfigSetAppendsValuesToEnd(t *testing.T) {
	t.Parallel()

	e := ptesting.NewEnvironment(t)
	defer e.DeleteIfNotFailed()

	khulnasoftProject := `
name: khulnasoft-test
runtime: go`

	integration.CreatePulumiRepo(e, khulnasoftProject)
	e.SetBackend(e.LocalURL())
	stackName := ptesting.RandomStackName()

	e.RunCommand("khulnasoft", "stack", "init", stackName)
	e.RunCommand("khulnasoft", "config", "set", "Bconfig", "one")
	e.RunCommand("khulnasoft", "config", "set", "Aconfig", "shouldBeAtEnd")
	e.RunCommand("khulnasoft", "config", "set", "Cconfig", "shouldAlsoBeAtEnd")
	e.RunCommand("khulnasoft", "config", "set", "Bconfig", "shouldOverWrite")

	b, err := os.ReadFile(filepath.Join(e.CWD, "Pulumi."+stackName+".yaml"))
	assert.NoError(t, err)

	expectedRegex := `encryptionsalt: .*
config:
  khulnasoft-test:Bconfig: shouldOverWrite
  khulnasoft-test:Aconfig: shouldBeAtEnd
  khulnasoft-test:Cconfig: shouldAlsoBeAtEnd`

	assert.Regexp(t, regexp.MustCompile(expectedRegex), strings.TrimSpace(string(b)))
}

func TestConfigRmRemovesValuesFromConfig(t *testing.T) {
	t.Parallel()

	e := ptesting.NewEnvironment(t)
	defer e.DeleteIfNotFailed()

	khulnasoftProject := `
name: khulnasoft-test
runtime: go`

	integration.CreatePulumiRepo(e, khulnasoftProject)
	e.SetBackend(e.LocalURL())
	stackName := ptesting.RandomStackName()

	e.RunCommand("khulnasoft", "stack", "init", stackName)
	e.RunCommand("khulnasoft", "config", "set", "Aconfig", "shouldBeRemoved")
	e.RunCommand("khulnasoft", "config", "set", "Bconfig", "shouldRemain")
	e.RunCommand("khulnasoft", "config", "rm", "Aconfig")

	b, err := os.ReadFile(filepath.Join(e.CWD, "Pulumi."+stackName+".yaml"))
	assert.NoError(t, err)

	expectedRegex := `encryptionsalt: .*
config:
  khulnasoft-test:Bconfig: shouldRemain`

	assert.Regexp(t, regexp.MustCompile(expectedRegex), strings.TrimSpace(string(b)))

	e.RunCommand("khulnasoft", "config", "rm", "Bconfig")

	b, err = os.ReadFile(filepath.Join(e.CWD, "Pulumi."+stackName+".yaml"))
	assert.NoError(t, err)

	expectedRegex = `encryptionsalt: .*`
	assert.Regexp(t, regexp.MustCompile(expectedRegex), strings.TrimSpace(string(b)))
}

func TestConfigGetRetrievedValueFromBothStackAndProjectUsingJson(t *testing.T) {
	t.Parallel()

	e := ptesting.NewEnvironment(t)
	defer e.DeleteIfNotFailed()

	khulnasoftProject := `
name: khulnasoft-test
runtime: go
config:
  first-value:
    type: string
    default: first
  second-value:
    type: string
  third-value:
    type: array
    items:
      type: string
    default: [third]`

	integration.CreatePulumiRepo(e, khulnasoftProject)
	e.SetBackend(e.LocalURL())
	stackName := ptesting.RandomStackName()

	e.RunCommand("khulnasoft", "stack", "init", stackName)
	e.RunCommand("khulnasoft", "config", "set", "second-value", "second")
	stdout, _ := e.RunCommand("khulnasoft", "config", "--json")
	// check that stdout is an array containing 2 objects
	var config map[string]interface{}
	jsonError := json.Unmarshal([]byte(stdout), &config)
	assert.Nil(t, jsonError)
	assert.Equal(t, 3, len(config))
	assert.Equal(t, "first", config["khulnasoft-test:first-value"].(map[string]interface{})["value"])
	assert.Equal(t, "second", config["khulnasoft-test:second-value"].(map[string]interface{})["value"])
	thirdValue := config["khulnasoft-test:third-value"].(map[string]interface{})
	assert.Equal(t, "[\"third\"]", thirdValue["value"])
	assert.Equal(t, []interface{}{"third"}, thirdValue["objectValue"])
}

func TestConfigCommandsUsingEnvironments(t *testing.T) {
	if getTestOrg() != khulnasoftTestOrg {
		t.Skip("Skipping test because the required environment is in the moolumi org.")
	}
	t.Parallel()

	e := ptesting.NewEnvironment(t)
	defer deleteIfNotFailed(e)

	integration.CreateBasicPulumiRepo(e)
	e.RunCommand("khulnasoft", "org", "set-default", getTestOrg())
	stackName := ptesting.RandomStackName()
	e.RunCommand("khulnasoft", "stack", "init", stackName)

	// check config is empty
	stdout, _ := e.RunCommand("khulnasoft", "config")
	assert.Equal(t, "KEY  VALUE", strings.Trim(stdout, "\r\n"))

	// set an esc environment
	e.RunCommand("khulnasoft", "config", "env", "add", "secrets-test-env-DO-NOT-DELETE", "--yes")

	// just `khulnasoft config`
	stdout, _ = e.RunCommand("khulnasoft", "config")
	assert.Equal(t, `KEY          VALUE
test_secret  [secret]`, strings.Trim(stdout, "\r\n"))

	// `khulnasoft config --show-secrets`
	stdout, _ = e.RunCommand("khulnasoft", "config", "--show-secrets")
	assert.Equal(t, `KEY          VALUE
test_secret  this_is_my_secret`, strings.Trim(stdout, "\r\n"))

	// `khulnasoft config --open`
	stdout, _ = e.RunCommand("khulnasoft", "config", "--open")
	assert.Equal(t, `KEY          VALUE
test_secret  [secret]`, strings.Trim(stdout, "\r\n"))

	// `khulnasoft config --show-secrets --open`
	stdout, _ = e.RunCommand("khulnasoft", "config", "--show-secrets", "--open")
	assert.Equal(t, `KEY          VALUE
test_secret  this_is_my_secret`, strings.Trim(stdout, "\r\n"))

	// `khulnasoft config --show-secrets --open=false`
	stdout, _ = e.RunCommand("khulnasoft", "config", "--show-secrets", "--open=false")
	assert.Equal(t, `KEY          VALUE
test_secret  [unknown]`, strings.Trim(stdout, "\r\n"))

	// `khulnasoft config get`
	stdout, _ = e.RunCommand("khulnasoft", "config", "get", "test_secret")
	assert.Equal(t, "this_is_my_secret", strings.Trim(stdout, "\r\n"))

	// `khulnasoft config get --open=false`
	stdout, _ = e.RunCommand("khulnasoft", "config", "get", "test_secret", "--open=false")
	assert.Equal(t, "[unknown]", strings.Trim(stdout, "\r\n"))

	// delete the stack
	e.RunCommand("khulnasoft", "stack", "rm", "-s", stackName, "--yes")
}

func getTestOrg() string {
	testOrg := khulnasoftTestOrg
	if _, set := os.LookupEnv("PULUMI_TEST_ORG"); set {
		testOrg = os.Getenv("PULUMI_TEST_ORG")
	}
	return testOrg
}
