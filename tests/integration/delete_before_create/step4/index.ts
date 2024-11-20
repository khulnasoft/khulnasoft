// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";
import { Resource } from "./resource";

// Setup for the next test.
const a = new Resource("base", { uniqueKey: 1, state: 100 });
const b = new Resource("base-2", { uniqueKey: 2, state: 100 });
const c = new Resource("dependent", { state: khulnasoft.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate) });
