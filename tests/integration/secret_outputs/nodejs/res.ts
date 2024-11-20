import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as dynamic from "@khulnasoft/khulnasoft/dynamic";

export interface RArgs {
    prefix: khulnasoft.Input<string>
}

const provider: khulnasoft.dynamic.ResourceProvider = {
    async create(inputs) {
        return { id: "1", outs: {
            prefix: inputs["prefix"]
        }};
    }
}

export class R extends dynamic.Resource {
    public prefix!: khulnasoft.Output<string>;

    constructor(name: string, props: RArgs, opts?: khulnasoft.CustomResourceOptions) {
        super(provider, name, props, opts)
    }
}