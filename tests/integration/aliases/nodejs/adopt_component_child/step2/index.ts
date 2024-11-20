// Copyright 2016-2023, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";

class FooResource extends khulnasoft.ComponentResource {
    constructor(name: string, opts?: khulnasoft.ComponentResourceOptions) {
        super("my:module:FooResource", name, {}, opts);
    }
}

class ComponentResource extends khulnasoft.ComponentResource {
    constructor(name: string, opts?: khulnasoft.ComponentResourceOptions) {
        super("my:module:ComponentResource", name, {}, opts);
        new FooResource("child", {
            parent: this,
            aliases: [{ parent: khulnasoft.rootStackResource }],
        });
    }
}

new ComponentResource("comp");
