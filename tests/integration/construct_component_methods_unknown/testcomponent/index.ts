// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as provider from "@khulnasoft/khulnasoft/provider";

class Component extends khulnasoft.ComponentResource {
    constructor(name: string, opts?: khulnasoft.ComponentResourceOptions) {
        super("testcomponent:index:Component", name, undefined, opts);
    }

    getMessage(echo: khulnasoft.Input<string>): khulnasoft.Output<string> {
        return khulnasoft.output(echo).apply(v => { console.log("should not run (echo)"); process.exit(1); });
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
                        return new Component(name, { urn });
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

        const component = new Component(name, options);
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
                        message: self.getMessage(inputs.echo),
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
