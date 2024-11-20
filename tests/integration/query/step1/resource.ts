// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";

let currentID = 0;

export class Provider implements khulnasoft.dynamic.ResourceProvider {
    public static readonly instance = new Provider();

    public async create(inputs: any) {
        return {
            id: (currentID++).toString(),
            outs: undefined,
        };
    }
}

export class Resource extends khulnasoft.dynamic.Resource {
    public isInstance(o: any): o is Resource {
        return o.__khulnasoftType === "khulnasoft-nodejs:dynamic:Resource";
    }

    constructor(name: string, props: khulnasoft.Inputs, opts?: khulnasoft.ResourceOptions) {
        super(Provider.instance, name, props, opts);
    }
}
