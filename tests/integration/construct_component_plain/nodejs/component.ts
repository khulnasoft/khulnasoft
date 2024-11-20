// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";

export interface ComponentArgs {
    children?: number;
}

export class Component extends khulnasoft.ComponentResource {
    constructor(name: string, args: ComponentArgs, opts?: khulnasoft.ComponentResourceOptions) {
        super("testcomponent:index:Component", name, args, opts, true);
    }
}
