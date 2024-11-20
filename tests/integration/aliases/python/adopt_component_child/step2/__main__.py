# Copyright 2016-2023, Pulumi Corporation.  All rights reserved.

import khulnasoft


class FooResource(khulnasoft.ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:FooResource", name, None, opts)


class ComponentResource(khulnasoft.ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:ComponentResource", name, None, opts)
        FooResource("child", khulnasoft.ResourceOptions(
            parent=self,
            aliases=[khulnasoft.Alias(parent=khulnasoft.ROOT_STACK_RESOURCE)]
        ))


ComponentResource("comp")
