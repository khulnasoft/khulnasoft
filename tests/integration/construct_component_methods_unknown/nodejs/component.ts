// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";

export class Component extends khulnasoft.ComponentResource {
    constructor(name: string, opts?: khulnasoft.ComponentResourceOptions) {
        super("testcomponent:index:Component", name, undefined, opts, true);
    }

    getMessage(args: Component.GetMessageArgs): khulnasoft.Output<Component.GetMessageResult> {
        return khulnasoft.runtime.call("testcomponent:index:Component/getMessage", {
            "__self__": this,
            "echo": args.echo,
        }, this);
    }
}

export namespace Component {
    export interface GetMessageArgs {
        echo: khulnasoft.Input<string>;
    }

    export interface GetMessageResult {
        message: string;
    }
}
