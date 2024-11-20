// Copyright 2016-2024, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";

class MyComponent extends khulnasoft.ComponentResource {
    constructor(name: string) {
        super("test:index:MyComponent", name);
    }
}

khulnasoft.log.debug("A debug message");

new MyComponent("mycomponent");
