// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as dynamic from "@khulnasoft/khulnasoft/dynamic";

// NOTE: Dynamic provider is restarted every step, so unless we read this from some external state
// store, this would always be 0 anyway.
const id = 0;

export class Provider implements dynamic.ResourceProvider {
    public static readonly instance = new Provider();

    public async check(olds: any, news: any): Promise<dynamic.CheckResult> {
        return {
            inputs: news,
        };
    }

    public async create(inputs: any): Promise<dynamic.CreateResult> {
        if (inputs.state === 4) {
            return Promise.reject({
                message: "Resource failed to initialize", id: id.toString(), properties: inputs,
                reasons: ["state can't be 4"],
            });
        }

        return {
            id: id.toString(),
            outs: inputs,
        };
    }

    public async update(id: khulnasoft.ID, olds: any, news: any): Promise<dynamic.UpdateResult> {
        if (news.state === 4) {
            return Promise.reject({
                message: "Resource failed to initialize", id: id.toString(), properties: news,
                reasons: ["state can't be 4"],
            });
        }

        return {
            outs: news,
        };
    }
}

export class Resource extends dynamic.Resource {
    public readonly state: khulnasoft.Output<number>;

    constructor(name: string, num: khulnasoft.Input<number>, opts?: khulnasoft.ResourceOptions) {
        super(Provider.instance, name, { state: num }, opts);
    }
}
