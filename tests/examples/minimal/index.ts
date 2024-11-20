// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Config } from "@khulnasoft/khulnasoft";

let config = new Config();
console.log(`Hello, ${config.require("name")}!`);
console.log(`Psst, ${config.require("secret")}`);

