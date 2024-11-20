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

import "mocha";
import * as assert from "assert";
import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as sut from "..";

khulnasoft.runtime.setMocks({
    newResource: function(_: khulnasoft.runtime.MockResourceArgs): {id: string, state: any} {
        throw new Error("newResource not implemented");
    },
    call: function(args: khulnasoft.runtime.MockCallArgs) {
        if (args.token == "mypkg::listStorageAccountKeys") {
            return {
                "keys": [{
                    "creationTime": "my-creation-time",
                    "keyName": "my-key-name",
                    "permissions": "my-permissions",
                    "value": JSON.stringify(args.inputs),
                }]
            };
        }
        if (args.token == "mypkg::getIntegrationRuntimeObjectMetadatum") {
            return {nextLink: JSON.stringify(args.inputs)};
        }
        if (args.token == "mypkg::funcWithAllOptionalInputs" ||
            args.token == "mypkg::funcWithDefaultValue" ||
            args.token == "mypkg::funcWithListParam" ||
            args.token == "mypkg::funcWithDictParam")
        {
            return {r: JSON.stringify(args.inputs)};
        }
        throw new Error("call not implemented for " + args.token);
    },
});

function checkTable(done: any, transform: (res: any) => any, table: {given: khulnasoft.Output<any>, expect: any}[]) {
    checkOutput(done, khulnasoft.all(table.map(x => x.given)), res => {
        res.map((actual, i) => {
            assert.deepStrictEqual(transform(actual), table[i].expect);
        });
    });
}

describe("output-funcs", () => {
    it("funcWithAllOptionalInputsOutput", (done) => {
        checkTable(done, res => JSON.parse(res.r), [
            {given: sut.funcWithAllOptionalInputsOutput({}),
             expect: {}},
            {given: sut.funcWithAllOptionalInputsOutput({a: khulnasoft.output("my-a")}),
             expect: {"a": "my-a"}},
            {given: sut.funcWithAllOptionalInputsOutput({a: khulnasoft.output("my-a"),
                                                         b: khulnasoft.output("my-b")}),
             expect: {"a": "my-a", "b": "my-b"}}
        ]);
    });

    // TODO[khulnasoft/khulnasoft#7973] Node codegen does not respect default
    // values at the moment, otherwise "b" parameter would receive the
    // default value from the schema.
    it("funcWithDefaultValueOutput", (done) => {
        checkTable(done, res => JSON.parse(res.r), [
            {given: sut.funcWithDefaultValueOutput({"a": khulnasoft.output("my-a")}),
             expect: {"a": "my-a"}},
            {given: sut.funcWithDefaultValueOutput({"a": khulnasoft.output("my-a"),
                                                    "b": khulnasoft.output("my-b")}),
             expect: {"a": "my-a", "b": "my-b"}}
        ]);
    });

    it("funcWithListParamOutput", (done) => {
        const l = ["a", "b", "c"];
        checkTable(done, res => JSON.parse(res.r), [
            {given: sut.funcWithListParamOutput({}),
             expect: {}},
            {given: sut.funcWithListParamOutput({"a": khulnasoft.output(l)}),
             expect: {"a": l}},
            {given: sut.funcWithListParamOutput({"a": khulnasoft.output(l),
                                                 "b": khulnasoft.output("my-b")}),
             expect: {"a": l, "b": "my-b"}},
        ]);
    });

    it("funcWithDictParamOutput", (done) => {
        const d = {"key-a": "value-a", "key-b": "value-b"};
        checkTable(done, res => JSON.parse(res.r), [
            {given: sut.funcWithDictParamOutput({}),
             expect: {}},
            {given: sut.funcWithDictParamOutput({"a": khulnasoft.output(d)}),
             expect: {"a": d}},
            {given: sut.funcWithDictParamOutput({"a": khulnasoft.output(d),
                                                 "b": khulnasoft.output("my-b")}),
             expect: {"a": d, "b": "my-b"}},
        ]);
    });

    it("listStorageAccountKeysOutput", (done) => {
        const output = sut.listStorageAccountKeysOutput({
            accountName: khulnasoft.output("my-account-name"),
            resourceGroupName: khulnasoft.output("my-resource-group-name"),
        });
        checkOutput(done, output, (res: sut.ListStorageAccountKeysResult) => {
            assert.equal(res.keys.length, 1);
            const k = res.keys[0];
            assert.equal(k.creationTime, "my-creation-time");
            assert.equal(k.keyName, "my-key-name");
            assert.equal(k.permissions, "my-permissions");
            assert.deepStrictEqual(JSON.parse(k.value), {
                "accountName": "my-account-name",
                "resourceGroupName": "my-resource-group-name"
            });
        });
    });

    it("listStorageAccountKeysOutput with optional arg set", (done) => {
        const output = sut.listStorageAccountKeysOutput({
            accountName: khulnasoft.output("my-account-name"),
            resourceGroupName: khulnasoft.output("my-resource-group-name"),
            expand: khulnasoft.output("my-expand"),
        });
        checkOutput(done, output, (res: sut.ListStorageAccountKeysResult) => {
            assert.equal(res.keys.length, 1);
            const k = res.keys[0];
            assert.equal(k.creationTime, "my-creation-time");
            assert.equal(k.keyName, "my-key-name");
            assert.equal(k.permissions, "my-permissions");
            assert.deepStrictEqual(JSON.parse(k.value), {
                "accountName": "my-account-name",
                "resourceGroupName": "my-resource-group-name",
                "expand": "my-expand"
            });
        });
    });

    it("listStorageAccountKeysOutput with unknown inputs returns unknown", (done) => {
        const output = sut.listStorageAccountKeysOutput({
            accountName: khulnasoft.output("my-account-name"),
            resourceGroupName: khulnasoft.output("my-resource-group-name"),
            expand: khulnasoft.unknown as any,
        });

        output.apply((res) => {
            done(new Error("apply should not be called when the result is unknown"));
        });

        setTimeout(done, 1000);
    });

    it("getIntegrationRuntimeObjectMetadatumOutput", (done) => {
        checkTable(
            done,
            (res: sut.GetIntegrationRuntimeObjectMetadatumResult) =>
                JSON.parse(res.nextLink || "{}"),
            [{given: sut.getIntegrationRuntimeObjectMetadatumOutput({
                factoryName: khulnasoft.output("my-factory-name"),
                integrationRuntimeName: khulnasoft.output("my-integration-runtime-name"),
                metadataPath: khulnasoft.output("my-metadata-path"),
                resourceGroupName: khulnasoft.output("my-resource-group-name")}),
              expect: {"factoryName": "my-factory-name",
                       "integrationRuntimeName": "my-integration-runtime-name",
                       "metadataPath": "my-metadata-path",
                       "resourceGroupName": "my-resource-group-name"}}]
        );
    });
 });


function checkOutput<T>(done: any, output: khulnasoft.Output<T>, check: (value: T) => void) {
    output.apply(value => {
        try {
            check(value);
            done();
        } catch (error) {
            done(error);
        }
    });
}
