// Copyright 2016-2024, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";

export class Random extends khulnasoft.Resource {
    result!: khulnasoft.Output<string | undefined>;
  
    constructor(name: string, length: number, opts?: khulnasoft.ResourceOptions) {
      const inputs: any = {};
      inputs["length"] = length;
      super("testprovider:index:Random", name, true, inputs, opts);
    }
  }