// Copyright 2016-2024, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";

class FailsOnCreate extends khulnasoft.CustomResource {
    public readonly value!: khulnasoft.Output<number>;
    constructor(name: string) {
        super("testprovider:index:FailsOnCreate", name, { value: undefined });
    }
}

export let xyz = "DEF";

export let foo = new FailsOnCreate("test").value;
