// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as provider from "@khulnasoft/khulnasoft/provider";

class Component extends khulnasoft.ComponentResource {
    public readonly foo: khulnasoft.Output<string>;

    constructor(name: string, foo: khulnasoft.Input<string>, opts?: khulnasoft.ComponentResourceOptions) {
        super("testcomponent:index:Component", name, {}, opts);

        this.foo = khulnasoft.output(foo);

        this.registerOutputs({
            foo: this.foo,
        })
    }
}

class Provider implements provider.Provider {
    public readonly version = "0.0.1";

    construct(name: string, type: string, inputs: khulnasoft.Inputs,
              options: khulnasoft.ComponentResourceOptions): Promise<provider.ConstructResult> {
        if (type != "testcomponent:index:Component") {
            throw new Error(`unknown resource type ${type}`);
        }

       const foo = khulnasoft.output("").apply(a => {
           throw new Error("intentional error from within an apply");
           return a;
       });


        const component = new Component(name, foo);
        return Promise.resolve({
            urn: component.urn,
            state: {
                foo: component.foo
            },
        });
    }
}

export function main(args: string[]) {
    return provider.main(new Provider(), args);
}

main(process.argv.slice(2));
