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
        new FooResource("childrenamed", {
            parent: this,
            aliases: [{ name: "child" }],
        });
    }
}

new ComponentResource("comp");
