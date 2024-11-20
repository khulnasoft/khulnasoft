// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";

class Resource extends khulnasoft.ComponentResource {
    constructor(name: string, opts?: khulnasoft.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive extends khulnasoft.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: khulnasoft.ComponentResourceOptions) {
        super("my:module:ComponentFive", name, {}, opts);
        this.resource = new Resource("otherchild", {parent: this});
    }
}
const comp5 = new ComponentFive("comp5");
