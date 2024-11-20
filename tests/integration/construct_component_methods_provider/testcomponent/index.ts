// Copyright 2016-2023, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as provider from "@khulnasoft/khulnasoft/provider";

interface ComponentArgs {
    first: khulnasoft.Input<string>;
    second: khulnasoft.Input<string>;
}

class Component extends khulnasoft.ComponentResource {
    public readonly first!: khulnasoft.Output<string>;
    public readonly second!: khulnasoft.Output<string>;

    constructor(name: string, args: ComponentArgs, opts?: khulnasoft.ComponentResourceOptions) {
        if (opts?.urn) {
            args = <any>{
                first: undefined,
                second: undefined,
            };
        }

        super("testcomponent:index:Component", name, args, opts);

        // Skip further initialization when being initialized from a resource reference.
        if (opts?.urn) {
            return;
        }

        this.registerOutputs(args);
    }

    getMessage(name: khulnasoft.Input<string>): khulnasoft.Output<string> {
        return khulnasoft.all([this.first, this.second, name]).apply(([frst, scnd, nm]) => `${frst} ${scnd}, ${nm}!`);
    }
}

class Provider implements provider.Provider {
    public readonly version = "0.0.1";

    constructor() {
        // Register any resources that can come back as resource references that need to be rehydrated.
        khulnasoft.runtime.registerResourceModule("testcomponent", "index", {
            version: this.version,
            construct: (name, type, urn) => {
                switch (type) {
                    case "testcomponent:index:Component":
                        return new Component(name, <any>undefined, { urn });
                    default:
                        throw new Error(`unknown resource type ${type}`);
                }
            },
        });
    }

    async construct(name: string, type: string, inputs: khulnasoft.Inputs,
              options: khulnasoft.ComponentResourceOptions): Promise<provider.ConstructResult> {
        if (type != "testcomponent:index:Component") {
            throw new Error(`unknown resource type ${type}`);
        }

        const component = new Component(name, <ComponentArgs>inputs, options);
        return {
            urn: component.urn,
            state: inputs,
        };
    }

    async call(token: string, inputs: khulnasoft.Inputs): Promise<provider.InvokeResult> {
        switch (token) {
            case "testcomponent:index:Component/getMessage":
                const self: Component = inputs.__self__;
                return {
                    outputs: {
                        message: self.getMessage(inputs.name),
                    },
                };

            default:
                throw new Error(`unknown method ${token}`);
        }
    }
}

export function main(args: string[]) {
    return provider.main(new Provider(), args);
}

main(process.argv.slice(2));