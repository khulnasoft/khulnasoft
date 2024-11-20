// Copyright 2016-2024, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";

interface ComponentArgs {
    echo: khulnasoft.Input<any>;
}

export class Component extends khulnasoft.ComponentResource {
    public readonly echo!: khulnasoft.Output<any>;
    public readonly childId!: khulnasoft.Output<khulnasoft.ID>;

    constructor(name: string, args: ComponentArgs, opts?: khulnasoft.ComponentResourceOptions) {
        const inputs: any = {};
        inputs["echo"] = args.echo;
        inputs["childId"] = undefined /*out*/;
        inputs["secret"] = undefined /*out*/;

        super("testcomponent:index:Component", name, inputs, opts, true);
    }
}

