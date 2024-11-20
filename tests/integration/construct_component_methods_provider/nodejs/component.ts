// Copyright 2016-2023, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";

interface ComponentArgs {
    first: khulnasoft.Input<string>;
    second: khulnasoft.Input<string>;
}

export class Component extends khulnasoft.ComponentResource {
    constructor(name: string, args: ComponentArgs, opts?: khulnasoft.ComponentResourceOptions) {
        super("testcomponent:index:Component", name, args, opts, true);
    }

    getMessage(args: Component.GetMessageArgs): khulnasoft.Output<Component.GetMessageResult> {
        return khulnasoft.runtime.call("testcomponent:index:Component/getMessage", {
            "__self__": this,
            "name": args.name,
        }, this);
    }
}

export namespace Component {
    export interface GetMessageArgs {
        name: khulnasoft.Input<string>;
    }

    export interface GetMessageResult {
        message: string;
    }
}
