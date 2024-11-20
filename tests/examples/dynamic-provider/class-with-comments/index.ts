// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as dynamic from "@khulnasoft/khulnasoft/dynamic";

class SimpleProvider implements khulnasoft.dynamic.ResourceProvider {
    public create: (inputs: any) => Promise<khulnasoft.dynamic.CreateResult>;

    // Ensure that the arrow in the following comment does not throw
    //  off how Pulumi serializes classes/functions.
    // public update: (id: khulnasoft.ID, inputs: any) => Promise<khulnasoft.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: "0",
                outs: undefined,
            };
        };
    }
}

class SimpleResource extends dynamic.Resource {
    public value = 4;

    constructor(name: string) {
        super(new SimpleProvider(), name, {}, undefined);
    }
}

let r = new SimpleResource("foo");
export const val = r.value;
