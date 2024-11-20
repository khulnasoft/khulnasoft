// Copyright 2016-2022, Pulumi Corporation.

import * as khulnasoft from '@khulnasoft/khulnasoft'

// A ResourceProvider using the default generic type, with explicit return type defined.
class DefaultGenericProvider implements khulnasoft.dynamic.ResourceProvider {
  async create (props: any): Promise<khulnasoft.dynamic.CreateResult> {
    return { id: 'resource-id', outs: {} }
  }

  async check (olds: any, news: any): Promise<khulnasoft.dynamic.CheckResult> {
    return Promise.resolve({ inputs: news })
  }

  async diff (id: khulnasoft.ID, olds: any, news: any): Promise<khulnasoft.dynamic.DiffResult> {
    return Promise.resolve({})
  }

  async delete (id: khulnasoft.ID, props: any): Promise<void> { return Promise.resolve() }

  async update (id: khulnasoft.ID, olds: any, news: any): Promise<khulnasoft.dynamic.UpdateResult> {
    return Promise.resolve({ outs: {} })
  }

  async read(id: khulnasoft.ID, props: any): Promise<khulnasoft.dynamic.ReadResult> {
    return Promise.resolve({ props: {} })
  }
}

type InputArgs = {
  names: string
}

type OutputArgs = {
  resourceId: string
  name: string
}

// All parameters and returns typed are inferred through the generic types provided.
class TypedGenericProvider implements khulnasoft.dynamic.ResourceProvider<InputArgs, OutputArgs> {
  async create (props) {
    return { id: 'resource-id', outs: { resourceId: "id", name: "test" } }
  }

  async check (olds, news) {
    return Promise.resolve({ inputs: news })
  }

  async diff (id, olds, news) {
    return Promise.resolve({})
  }

  async delete (id, props) { return Promise.resolve() }

  async update (id, olds, news) {
    return Promise.resolve({ outs: { resourceId: olds.resourceId, ...news } })
  }

  async read(id, props) {
    return Promise.resolve({ props: props })
  }
}
