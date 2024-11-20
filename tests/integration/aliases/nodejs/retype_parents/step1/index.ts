// Copyright 2016-2022, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";

class Resource extends khulnasoft.ComponentResource {
    constructor(name: string, opts?: khulnasoft.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #6 - Nested parents changing types
class ComponentSix extends khulnasoft.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: khulnasoft.ComponentResourceOptions) {
        super("my:module:ComponentSix-v0", name, {}, opts);
        this.resource = new Resource("otherchild", {parent: this});
    }
}

class ComponentSixParent extends khulnasoft.ComponentResource {
    child: ComponentSix;
    constructor(name: string, opts?: khulnasoft.ComponentResourceOptions) {
        super("my:module:ComponentSixParent-v0", name, {}, opts);
        this.child = new ComponentSix("child", {parent: this});
    }
}

const comp4 = new ComponentSixParent("comp6");
