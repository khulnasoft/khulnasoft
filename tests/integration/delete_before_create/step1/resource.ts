// Copyright 2016-2022, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as dynamic from "@khulnasoft/khulnasoft/dynamic";
import {v4 as uuidv4} from "uuid";

export class Provider implements dynamic.ResourceProvider {
    public static readonly instance = new Provider();

    public async check(olds: any, news: any): Promise<dynamic.CheckResult> {
        return {
            inputs: news,
        };
    }

    public async diff(id: khulnasoft.ID, olds: any, news: any): Promise<dynamic.DiffResult> {
        if (olds.state !== news.state) {
            return {
                changes: true,
                replaces: ["state"],
                deleteBeforeReplace: news.noDBR ? false : true,
            };
        }

        if (olds.noReplace !== news.noReplace) {
            return {
                changes: true,
            }
        }

        return {
            changes: false,
        };
    }

    public async create(inputs: any): Promise<dynamic.CreateResult> {
        return {
            id: uuidv4(),
            outs: inputs,
        };
    }
}

export class Resource extends khulnasoft.dynamic.Resource {
    public uniqueKey?: khulnasoft.Output<number>;
    public state: khulnasoft.Output<number>;
    public noReplace?: khulnasoft.Output<number>;

    constructor(name: string, props: ResourceProps, opts?: khulnasoft.CustomResourceOptions) {
        super(Provider.instance, name, props, opts);
    }
}

export interface ResourceProps {
    readonly uniqueKey?: khulnasoft.Input<number>;
    readonly state: khulnasoft.Input<number>;
    readonly noReplace?: khulnasoft.Input<number>;
    readonly noDBR?: khulnasoft.Input<boolean>;
}
