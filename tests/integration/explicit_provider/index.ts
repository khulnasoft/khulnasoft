// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";

class DynamicProvider extends khulnasoft.ProviderResource {
    constructor(name: string, opts?: khulnasoft.ResourceOptions) {
        super("khulnasoft-nodejs", name, {}, opts);
    }
}

class Provider implements khulnasoft.dynamic.ResourceProvider {
    public static instance = new Provider();

    public create: (inputs: any) => Promise<khulnasoft.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: "0",
                outs: undefined,
            };
        };
    }
}

class Resource extends khulnasoft.dynamic.Resource {
    constructor(name: string, provider?: khulnasoft.ProviderResource) {
        super(Provider.instance, name, {}, { provider: provider});
    }
}

// Create a resource using the default dynamic provider instance.
let a = new Resource("a");

// Create an explicit instance of the dynamic provider.
let p = new DynamicProvider("p");

// Create a resource using the explicit dynamic provider instance.
let b = new Resource("b", p);
