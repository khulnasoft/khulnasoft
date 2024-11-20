// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as dynamic from "@khulnasoft/khulnasoft/dynamic";

export class Provider implements dynamic.ResourceProvider {
    public static readonly instance = new Provider();

    private id: number = 0;

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
