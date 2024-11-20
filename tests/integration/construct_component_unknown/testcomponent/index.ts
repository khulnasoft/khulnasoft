import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as provider from "@khulnasoft/khulnasoft/provider";

interface ComponentArgs {
    message: khulnasoft.Input<string>;
    nested: khulnasoft.Input<{
        value: khulnasoft.Input<string>;
    }>;
}

class Component extends khulnasoft.ComponentResource {
    constructor(name: string, args: ComponentArgs, opts?: khulnasoft.ComponentResourceOptions) {
        super("testcomponent:index:Component", name, {}, opts);

        // These `apply`s should not run.
        khulnasoft.output(args.message).apply(v => { console.log("should not run (message)"); process.exit(1); });
        khulnasoft.output(args.nested).apply(v => { console.log("should not run (nested)"); process.exit(1); });
    }
}

class Provider implements provider.Provider {
    public readonly version = "0.0.1";

    construct(name: string, type: string, inputs: khulnasoft.Inputs,
              options: khulnasoft.ComponentResourceOptions): Promise<provider.ConstructResult> {
        if (type != "testcomponent:index:Component") {
            throw new Error(`unknown resource type ${type}`);
        }

        const component = new Component(name, <ComponentArgs>inputs, options);
        return Promise.resolve({
            urn: component.urn,
            state: {},
        });
    }
}

export function main(args: string[]) {
    return provider.main(new Provider(), args);
}

main(process.argv.slice(2));
