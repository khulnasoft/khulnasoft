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
        const component = new Component(name, inputs["foo"], options);

	return {
	    urn: component.urn,
	    state: inputs,
	};
    }

    async call(token: string, inputs: khulnasoft.Inputs): Promise<provider.InvokeResult> {
	switch (token) {
	    case "testcomponent:index:Component/getMessage":
		throw new khulnasoft.InputPropertyError({propertyPath: "foo", reason: "the failure reason"});
	    default:
		throw new Error(`unknown method ${token}`);
	};
    }
}

export function main(args: string[]) {
    return provider.main(new Provider(), args);
}

main(process.argv.slice(2));
