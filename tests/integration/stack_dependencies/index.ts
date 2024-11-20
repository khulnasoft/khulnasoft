// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";

class Provider implements khulnasoft.dynamic.ResourceProvider {
    public readonly create: (inputs: any) => Promise<khulnasoft.dynamic.CreateResult>;

    constructor(num: number) {
        this.create = async (inputs: any) => {
            return {
                id: "0",
                outs: { value: num }
            }
        }
    }
}


class FirstResource extends khulnasoft.dynamic.Resource {
    public readonly value: khulnasoft.Output<number>;

    private static provider: Provider = new Provider(42);
    constructor(name: string) {
        super(FirstResource.provider, name, { value: undefined }, undefined);
    }
}

class SecondResource extends khulnasoft.dynamic.Resource {
    public readonly dep: khulnasoft.Output<number>;

    private static provider: Provider = new Provider(99);

    constructor(name: string, prop: khulnasoft.Input<number>) {
        super(SecondResource.provider, name, {dep: prop}, undefined);
    }
}

const first = new FirstResource("first");
first.value.apply(v => {
    console.log(`first.value: ${v}`);
});


const second = new SecondResource("second", first.value);
second.dep.apply(d => {
    console.log(`second.dep: ${d}`);
});