# Copyright 2016-2021, Pulumi Corporation.  All rights reserved.

from khulnasoft import export
from khulnasoft.dynamic import CreateResult, Resource, ResourceProvider


class CustomResource(
    Resource, module="custom-provider", name="CustomResource"
):
    def __init__(self, name, opts=None):
        super().__init__(DummyResourceProvider(), name, {}, opts)


class DummyResourceProvider(ResourceProvider):
    def create(self, props):
        return CreateResult("resource-id", {})


resource = CustomResource("resource-name")

export("urn", resource.urn)
