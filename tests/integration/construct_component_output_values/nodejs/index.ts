// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";

import { Component } from "./component";

new Component("component", {
    foo: {
        something: "hello",
    },
    bar: {
        tags: {
            "a": "world",
            "b": khulnasoft.secret("shh"),
        },
    },
});
