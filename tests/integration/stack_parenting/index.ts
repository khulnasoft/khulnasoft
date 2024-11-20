// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";

let currentID = 0;

class Provider implements khulnasoft.dynamic.ResourceProvider {
    public static instance = new Provider();

    public create: (inputs: any) => Promise<khulnasoft.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: (currentID++).toString(),
                outs: undefined,
            };
        };
    }
}

class Component extends khulnasoft.ComponentResource {
    constructor(name: string, parent?: khulnasoft.ComponentResource) {
        super("component", name, {}, { parent: parent });
    }
}

class Resource extends khulnasoft.dynamic.Resource {
    constructor(name: string, parent?: khulnasoft.ComponentResource) {
        super(Provider.instance, name, {}, { parent: parent });
    }
}

// Just allocate a few resources and make sure their URNs are correct with respect to parents, etc.  This
// should form a tree of roughly the following structure:
//
//     A      F
//    / \      \
//   B   C      G
//      / \
//     D   E
//
// with the caveat, of course, that A and F will share a common parent, the implicit stack.
let a = new Component("a");

let b = new Resource("b", a);
let c = new Component("c", a);

let d = new Resource("d", c);
let e = new Resource("e", c);

let f = new Component("f");

let g = new Resource("g", f);
