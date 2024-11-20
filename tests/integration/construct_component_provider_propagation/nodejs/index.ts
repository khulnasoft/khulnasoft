// Copyright 2016-2023, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";

class Component extends khulnasoft.ComponentResource {
    public readonly result!: khulnasoft.Output<string>;

    constructor(name: string, opts?: khulnasoft.ComponentResourceOptions) {
        const inputs = {
            result: undefined /*out*/,
        };
        super("testcomponent:index:Component", name, inputs, opts, true);
    }
}

class RandomProvider extends khulnasoft.ProviderResource {
    constructor(name: string, opts?: khulnasoft.ResourceOptions) {
        super("testprovider", name, {}, opts);
    }
}

const explicitProvider = new RandomProvider("explicit");

new Component("uses_default");
new Component("uses_provider", {provider: explicitProvider});
new Component("uses_providers", {providers: [explicitProvider]});
new Component("uses_providers_map", {providers: {testprovider: explicitProvider}});
