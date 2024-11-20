// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";
import { Resource } from "./resource";

// resource "not-doomed" is created successfully.
const a = new Resource("not-doomed", 5);

// "a" should be in the checkpoint.
