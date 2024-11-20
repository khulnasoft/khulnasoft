// Copyright 2016-2023, Pulumi Corporation.  All rights reserved.
import * as khulnasoft from "@khulnasoft/khulnasoft";

let currentID = 0;

export class Provider implements khulnasoft.dynamic.ResourceProvider {
    public static readonly instance = new Provider();

    constructor() {}

    public async create(inputs: any) {
        return {
            id: (currentID++).toString(),
            outs: inputs,
        };
    }

    public async delete(id: khulnasoft.ID, props: any) {}

    public async diff(id: khulnasoft.ID, olds: any, news: any) { return {}; }

    public async update(id: khulnasoft.ID, olds: any, news: any) {
        return news;
    }
}

export class Resource extends khulnasoft.dynamic.Resource {
    constructor(name: string, props: ResourceProps, opts?: khulnasoft.ResourceOptions) {
        super(Provider.instance, name, props, opts);
    }
}

export interface ResourceProps {
    state?: any; // arbitrary state bag that can be updated without replacing.
}
