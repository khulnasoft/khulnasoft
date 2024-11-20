// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";

interface RandomArgs {
    length: khulnasoft.Input<number>;
}

export class Random extends khulnasoft.CustomResource {
    public readonly length!: khulnasoft.Output<number>;
    public readonly result!: khulnasoft.Output<string>;
    constructor(name: string, args: RandomArgs, opts?: khulnasoft.CustomResourceOptions) {
        const props = {
            length: args.length,
            result: undefined,
        }
        super("testprovider:index:Random", name, props, opts);
    }
}
