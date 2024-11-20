// Copyright 2016-2023, Pulumi Corporation.  All rights reserved.
import * as khulnasoft from "@khulnasoft/khulnasoft";
import { Resource } from "./resource";


let config = new khulnasoft.Config();
export const a = new Resource("a", {
    state: {
        // Driven by table tests in steps_test.go.
        [config.require("propertyName")]: "foo",
    }
});
