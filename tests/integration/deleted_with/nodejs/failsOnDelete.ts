// Copyright 2016-2023, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";

export class FailsOnDelete extends khulnasoft.CustomResource {
    constructor(name: string, opts?: khulnasoft.CustomResourceOptions) {
        super("testprovider:index:FailsOnDelete", name, {}, opts);
    }
}
