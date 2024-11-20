# Copyright 2016-2023, Pulumi Corporation.  All rights reserved.

import khulnasoft


class FooResource(khulnasoft.ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:FooResource", name, None, opts)


class ComponentResource(khulnasoft.ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:ComponentResource", name, None, opts)


comp = ComponentResource("comp")


FooResource("child", khulnasoft.ResourceOptions(
    aliases=[khulnasoft.Alias(parent=comp)]
))
