// Copyright 2016-2022, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as dynamic from "@khulnasoft/khulnasoft/dynamic";

export class Provider implements dynamic.ResourceProvider {
    public static readonly instance = new Provider();

    private id: number = 0;

    public async check(olds: any, news: any): Promise<dynamic.CheckResult> {
        // When the engine re-creates a resource after it was deleted, it should
        // not pass the old (deleted) inputs to Check when re-creating.
        //
        // This Check implementation fails the test if this happens.
        if (olds.state === 99 && news.state === 22) {
            return {
                inputs: news,
                failures: [
                    {
                        property: "state",
                        reason: "engine did invalid comparison of old and new check inputs for recreated resource",
                    },
                ],
            };
        }

        return {
            inputs: news,
        };
    }

    public async diff(id: khulnasoft.ID, olds: any, news: any): Promise<dynamic.DiffResult> {
        if (olds.state !== news.state) {
            return {
                changes: true,
                replaces: ["state"],
                deleteBeforeReplace: true,
            };
        }

        return {
            changes: false,
        };
    }

    public async create(inputs: any): Promise<dynamic.CreateResult> {
        return {
            id: (this.id++).toString(),
            outs: inputs,
        };
    }
}

export class Resource extends khulnasoft.dynamic.Resource {
    public uniqueKey?: khulnasoft.Output<number>;
    public state: khulnasoft.Output<number>;

    constructor(name: string, props: ResourceProps, opts?: khulnasoft.ResourceOptions) {
        super(Provider.instance, name, props, opts);
    }
}

export interface ResourceProps {
    readonly uniqueKey?: khulnasoft.Input<number>;
    readonly state: khulnasoft.Input<number>;
}
