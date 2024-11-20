// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";

class MyResource extends khulnasoft.dynamic.Resource {
    constructor(name: string, props: khulnasoft.Inputs, opts?: khulnasoft.CustomResourceOptions) {
        super({
            create: async (inputs: any) => {
                return {
                    id: "0",
                    outs: inputs,
                }
            },
        }, name, props, opts);
    }
}

class GetResource extends khulnasoft.Resource {
    foo: khulnasoft.Output<string>;
    bar: khulnasoft.Output<string>;

    constructor(urn: khulnasoft.URN) {
        const props = {
            foo: undefined,
            bar: undefined,
        };
        super("unused:unused:unused", "unused", true, props, { urn });
    }
}

const a = new MyResource("a", {
    foo: "foo",
    bar: khulnasoft.secret("my-$ecret"),
});

const getFoo = a.urn.apply(urn => {
    const r = new GetResource(urn);
    return r.foo
});

const getBar = a.urn.apply(urn => {
    const r = new GetResource(urn);
    return r.bar
});


export const foo = getFoo;
export const secret = getBar;