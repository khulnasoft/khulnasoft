// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";

interface ComponentArgs {
    foo: khulnasoft.Input<string>;
}

export class Component extends khulnasoft.ComponentResource {
    public readonly foo!: khulnasoft.Output<string>;

    constructor(name: string, args: ComponentArgs, opts?: khulnasoft.ComponentResourceOptions) {
        const inputs: any = {};
        inputs["foo"] = args.foo;

        super("testcomponent:index:Component", name, inputs, opts, true);
    }
}

