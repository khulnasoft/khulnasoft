// Copyright 2016-2024, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";
import { Random } from "./random";

khulnasoft.runtime.registerStackTransform(async ({ type, props, opts }) => {
    console.log("stack transform");
    return undefined;
});

new Random("res1", { length: khulnasoft.secret(5) });