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
                outs: inputs,
            };
        };
    }
}

export class Resource extends khulnasoft.dynamic.Resource {
    public readonly foo: khulnasoft.Output<string>;
    public readonly bar: khulnasoft.Output<{ value: string, unknown: string }>;
    public readonly baz: khulnasoft.Output<any[]>;

    constructor(name: string, props: ResourceProps, opts?: khulnasoft.ResourceOptions) {
        super(Provider.instance, name, props, opts);
    }
}

export interface ResourceProps {
    foo: khulnasoft.Input<string>;
    bar: khulnasoft.Input<{ value: khulnasoft.Input<string>, unknown: khulnasoft.Input<string> }>;
    baz: khulnasoft.Input<khulnasoft.Input<any>[]>;
}
