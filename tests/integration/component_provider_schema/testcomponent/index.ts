// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";

class Provider implements khulnasoft.provider.Provider {
    public readonly version = "0.0.1";
    constructor(public readonly schema?: string) {
    }
}

export function main(args: string[]) {
    const schema = process.env.INCLUDE_SCHEMA ? `{"hello": "world"}` : undefined;
    return khulnasoft.provider.main(new Provider(schema), args);
}

main(process.argv.slice(2));
