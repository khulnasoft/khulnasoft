import * as khulnasoft from "@khulnasoft/khulnasoft";

// Regression test for [khulnasoft/khulnasoft#2741], you should be able to create an instance of a first class provider
// with secret configuration values, so long as these values are themselves strings.
class DynamicProvider extends khulnasoft.ProviderResource {
    constructor(name: string, opts?: khulnasoft.ResourceOptions) {
        super("khulnasoft-nodejs", name,  { secretProperty: khulnasoft.secret("it's a secret to everybody") }, opts);
    }
}

const p = new DynamicProvider("p");
