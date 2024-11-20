// Copyright 2016-2023, Pulumi Corporation.  All rights reserved.

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

class LocalComponent extends khulnasoft.ComponentResource {
    public readonly message: khulnasoft.Output<string>;

    constructor(name: string, opts?: khulnasoft.ComponentResourceOptions) {
        super("my:index:LocalComponent", name, {}, opts);

        const component = new Component(`${name}-mycomponent`, { parent: this });
        this.message = component.message;
    }
}

const provider = new Provider("myprovider", "hello world")
const component = new Component("mycomponent", { provider });
const localComponent = new LocalComponent("mylocalcomponent", { providers: [provider] });

export const message = component.message;
export const nestedMessage = localComponent.message;
