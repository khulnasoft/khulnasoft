// Copyright 2016-2024, Pulumi Corporation.
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

import * as runtime from "@khulnasoft/khulnasoft/runtime";

(async function() {
    const deps = await runtime.computeCodePaths();

    const actual = JSON.stringify([...deps.keys()]);
    const expected = "[]";

    if (actual !== expected) {
        throw new Error(`Got '${actual}' expected '${expected}'`);
    }
})()