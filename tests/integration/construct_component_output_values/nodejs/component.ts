// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";

export interface BarArgs {
    tags?: khulnasoft.Input<{[key: string]: khulnasoft.Input<string>}>;
}

export interface FooArgs {
    something?: khulnasoft.Input<string>;
}

export interface ComponentArgs {
    bar?: khulnasoft.Input<BarArgs>;
    foo?: FooArgs;
}

export class Component extends khulnasoft.ComponentResource {
    constructor(name: string, args?: ComponentArgs, opts?: khulnasoft.ComponentResourceOptions) {
        super("testcomponent:index:Component", name, args, opts, true /*remote*/);
    }
}
