// Copyright 2016-2023, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";

export class TestProvider extends khulnasoft.ProviderResource {
    constructor(name: string) {
        super("testprovider", name);
    }
}
