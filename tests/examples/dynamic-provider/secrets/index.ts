// Copyright 2023, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as dynamic from "@khulnasoft/khulnasoft/dynamic";

const config = new khulnasoft.Config();
const password = config.requireSecret("password");

class SimpleProvider implements khulnasoft.dynamic.ResourceProvider {
    async create(inputs: any) {
        // Need to use `password.get()` to get the underlying value of the secret from within the serialzied code.  
        // This simulates using this as a credential to talk to an external system.
        return { id: "0", outs: { authenticated: password.get() == "s3cret" ? "200" : "401" } };
    }
}

class SimpleResource extends dynamic.Resource {
    authenticated!: khulnasoft.Output<string>;
    constructor(name: string) {
        super(new SimpleProvider(), name, { authenticated: undefined }, undefined);
    }
}

let r = new SimpleResource("foo");
export const out = r.authenticated;

