// Copyright 2016-2023, Pulumi Corporation.
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

import * as assert from "assert";
import * as khulnasoft from "..";

class TestMocks implements khulnasoft.runtime.Mocks {
    outputs: any;

    constructor(outputs: any) {
        this.outputs = outputs;
    }

    call(args: khulnasoft.runtime.MockCallArgs): Record<string, any> {
        throw new Error(`unknown function ${args.token}`);
    }

    newResource(args: khulnasoft.runtime.MockResourceArgs): { id: string | undefined; state: Record<string, any> } {
        switch (args.type) {
            case "khulnasoft:khulnasoft:StackReference":
                return {
                    id: `${args.name}_id`,
                    state: {
                        outputs: this.outputs,
                    },
                };
            default:
                throw new Error(`unknown type ${args.type}`);
        }
    }
}

describe("StackReference.getOutputDetails", () => {
    // The two tests don't share a mock because in the JS SDK,
    // if a map item is a secret, the entire map gets promoted to secret.

    it("supports plain text", async () => {
        khulnasoft.runtime.setMocks(
            new TestMocks({
                bucket: "mybucket-1234",
            }),
        );
        const ref = new khulnasoft.StackReference("foo");

        assert.deepStrictEqual(await ref.getOutputDetails("bucket"), {
            value: "mybucket-1234",
        });
    });

    it("supports secrets", async () => {
        khulnasoft.runtime.setMocks(
            new TestMocks({
                password: khulnasoft.secret("supersecretpassword"),
            }),
        );
        const ref = new khulnasoft.StackReference("foo");

        assert.deepStrictEqual(await ref.getOutputDetails("password"), {
            secretValue: "supersecretpassword",
        });
    });

    it("types applies correctly", async () => {
        const passwordValue = "supersecretpassword";
        khulnasoft.runtime.setMocks(
            new TestMocks({
                password: khulnasoft.secret(passwordValue),
            }),
        );
        const ref = new khulnasoft.StackReference("foo");
        const password: khulnasoft.Output<number> = ref.outputs["password"].apply((x) => x.length);

        assert.deepStrictEqual(await password.promise(), passwordValue.length);
    });
});
