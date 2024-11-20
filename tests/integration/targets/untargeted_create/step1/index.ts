// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";

let currentID = 0;

class Provider implements khulnasoft.dynamic.ResourceProvider {
    public static instance = new Provider();

    public create: (inputs: any) => Promise<khulnasoft.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: (currentID++) + "",
                outs: undefined,
            };
        };
    }
}

class Resource extends khulnasoft.dynamic.Resource {
    constructor(name: string, opts?: khulnasoft.ResourceOptions) {
        super(Provider.instance, name, {}, opts);
    }
}

// Create a resource using the default dynamic provider instance.
let a = new Resource("a");
let b = new Resource("b");

export const urn = a.urn;
