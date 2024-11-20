// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";

let currentID = 0;

export class Provider implements khulnasoft.dynamic.ResourceProvider {
    public static readonly instance = new Provider();

    public readonly create: (inputs: any) => Promise<khulnasoft.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: (currentID++).toString(),
                outs: undefined,
            };
        };
    }
}

export class Resource extends khulnasoft.dynamic.Resource {
    public readonly state?: any;

    constructor(name: string, props: ResourceProps, opts?: khulnasoft.ResourceOptions) {
        super(Provider.instance, name, props, opts);
        this.state = props.state;
    }
}

export interface ResourceProps {
    state?: any; // arbitrary state bag that can be updated without replacing.
}
