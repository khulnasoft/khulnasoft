// Copyright 2016-2023, Pulumi Corporation.

import * as khulnasoft from '@khulnasoft/khulnasoft'

class CustomResource extends khulnasoft.dynamic.Resource {
  constructor (name: string, opts?: khulnasoft.ResourceOptions) {
    super(new DummyResourceProvider(), name, {}, opts, "custom-provider", "CustomResource")
  }
}

class DummyResourceProvider implements khulnasoft.dynamic.ResourceProvider {
  async create (props: any): Promise<khulnasoft.dynamic.CreateResult> {
    var config = JSON.parse(process.env.PULUMI_CONFIG)

    return { id: config["id"], outs: {} }
  }
}

const resource = new CustomResource('resource-name')

export const rid = resource.id
