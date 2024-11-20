// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";
import { Resource } from "./resource";

// resource "not-doomed" is updated, but the update partially fails.
const a = new Resource("not-doomed", 4);

// "a" should still be in the checkpoint with its new value.
