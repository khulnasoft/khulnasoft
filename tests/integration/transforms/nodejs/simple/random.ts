// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";

interface RandomArgs {
    length: khulnasoft.Input<number>;
    prefix?: khulnasoft.Input<string | undefined>;
}

export class Random extends khulnasoft.CustomResource {
    public readonly length!: khulnasoft.Output<number>;
    public readonly result!: khulnasoft.Output<string>;
    constructor(name: string, args: RandomArgs, opts?: khulnasoft.CustomResourceOptions) {
        super("testprovider:index:Random", name, args, opts);
    }

    randomInvoke(args) {
	return khulnasoft.runtime.invoke("testprovider:index:returnArgs", args);
    }
}


interface ComponentArgs {
    length: khulnasoft.Input<number>;
}

export class Component extends khulnasoft.ComponentResource {
    public readonly length!: khulnasoft.Output<number>;
    public readonly childId!: khulnasoft.Output<string>;
    constructor(name: string, args: ComponentArgs, opts?: khulnasoft.ComponentResourceOptions) {
        super("testprovider:index:Component", name, args, opts, true);
    }
}

export class TestProvider extends khulnasoft.ProviderResource {
    constructor(name: string, opts?: khulnasoft.ResourceOptions) {
        super("testprovider", name, {}, opts);
    }
}
