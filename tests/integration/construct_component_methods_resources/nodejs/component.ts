// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";

export class Component extends khulnasoft.ComponentResource {
    constructor(name: string, opts?: khulnasoft.ComponentResourceOptions) {
        super("testcomponent:index:Component", name, undefined, opts, true);
    }

    createRandom(args: Component.CreateRandomArgs): khulnasoft.Output<Component.CreateRandomResult> {
        return khulnasoft.runtime.call("testcomponent:index:Component/createRandom", {
            "__self__": this,
            "length": args.length,
        }, this);
    }
}

export namespace Component {
    export interface CreateRandomArgs {
        length: khulnasoft.Input<number>;
    }

    export interface CreateRandomResult {
        result: string;
    }
}
