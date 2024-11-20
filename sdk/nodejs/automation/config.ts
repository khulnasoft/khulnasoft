// Copyright 2016-2020, Pulumi Corporation.
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

/**
 * An input to/output from a `khulnasoft config` command.
 */
export interface ConfigValue {
    /**
     * The underlying configuration value.
     */
    value: string;

    /**
     * True if and only if this configuration value is a secret.
     */
    secret?: boolean;
}

/**
 * A map of configuration values.
 */
export type ConfigMap = { [key: string]: ConfigValue };
