// Copyright 2016-2024, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as provider from "@khulnasoft/khulnasoft/provider";
import * as grpc from "@grpc/grpc-js";

class Component extends khulnasoft.ComponentResource {
    public readonly foo: khulnasoft.Output<string>;

    constructor(name: string, foo: khulnasoft.Input<string>, opts?: khulnasoft.ComponentResourceOptions) {
        super("testcomponent:index:Component", name, undefined, opts);

        this.foo = khulnasoft.output(foo);

        this.registerOutputs({
            foo: this.foo,
        })
    }
}

class Provider implements provider.Provider {
    public readonly version = "0.0.1";

    async construct(name: string, type: string, inputs: khulnasoft.Inputs,
              options: khulnasoft.ComponentResourceOptions): Promise<provider.ConstructResult> {
        if (type != "testcomponent:index:Component") {
            throw new Error(`unknown resource type ${type}`);
        }

        const component = new Component(name, inputs["foo"], options);

	throw new khulnasoft.InputPropertiesError({
	    message: "failing for a reason",
	    errors: [{propertyPath: "foo", reason: "the failure reason"}]});
    }
}

export function main(args: string[]) {
    return provider.main(new Provider(), args);
}

main(process.argv.slice(2));
