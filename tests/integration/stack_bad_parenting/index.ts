// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";

let currentID = 0;

class Provider implements khulnasoft.dynamic.ResourceProvider {
    public static instance = new Provider();

    public create: (inputs: any) => Promise<khulnasoft.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: (currentID++).toString(),
                outs: undefined,
            };
        };
    }
}

class Resource extends khulnasoft.dynamic.Resource {
    constructor(name: string, parent?: khulnasoft.Resource) {
        super(Provider.instance, name, {}, { parent: parent });
    }
}

// Ensure we throw if pass an non-resource as a parent.
let a = new Resource("a", <any>this);
