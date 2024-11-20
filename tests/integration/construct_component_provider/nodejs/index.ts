// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";

class Provider extends khulnasoft.ProviderResource {
    public readonly message!: khulnasoft.Output<string>;

    constructor(name: string, message: string, opts?: khulnasoft.ResourceOptions) {
        super("testcomponent", name, { message }, opts);
    }
}

class Component extends khulnasoft.ComponentResource {
    public readonly message!: khulnasoft.Output<string>;

    constructor(name: string, opts?: khulnasoft.ComponentResourceOptions) {
        const inputs = {
            message: undefined /*out*/,
        };
        super("testcomponent:index:Component", name, inputs, opts, true);
    }
}

const component = new Component("mycomponent", {
    providers: {
        "testcomponent": new Provider("myprovider", "hello world"),
    },
});

export const message = component.message;
