// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as provider from "@khulnasoft/khulnasoft/provider";

const version = "0.0.1";

class Provider extends khulnasoft.ProviderResource {
    public readonly message!: khulnasoft.Output<string>;

    constructor(name: string, opts?: khulnasoft.ResourceOptions) {
        super("testcomponent", name, { "message": undefined }, opts);
    }
}

class Component extends khulnasoft.ComponentResource {
    public message!: khulnasoft.Output<string>;

    constructor(name: string, opts?: khulnasoft.ComponentResourceOptions) {
        super("testcomponent:index:Component", name, {}, opts);
    }

    protected async initialize(args: khulnasoft.Inputs) {
        const provider = this.getProvider("testcomponent::");
        if (!(provider instanceof Provider)) {
            throw new Error("provider is not an instance of Provider");
        }
        this.message = provider.message;
        this.registerOutputs({
            message: provider.message,
        });
        return undefined;
    }
}

class ProviderServer implements provider.Provider {
    public readonly version = version;

    constructor() {
        khulnasoft.runtime.registerResourcePackage("testcomponent", {
            version,
            constructProvider: (name: string, type: string, urn: string): khulnasoft.ProviderResource => {
                if (type !== "khulnasoft:providers:testcomponent") {
                    throw new Error(`unknown provider type ${type}`);
                }
                return new Provider(name, { urn });
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
            state: {
                message: component.message,
            },
        };
    }
}

export function main(args: string[]) {
    return provider.main(new ProviderServer(), args);
}

main(process.argv.slice(2));
