# Copyright 2016-2023, Pulumi Corporation.  All rights reserved.

import khulnasoft


class FooResource(khulnasoft.ComponentResource):
    def __init__(self, name, opts=None):
        alias_opts = khulnasoft.ResourceOptions(aliases=[khulnasoft.Alias(type_="my:module:FooResource")])
        opts = khulnasoft.ResourceOptions.merge(opts, alias_opts)
        super().__init__("my:module:FooResourceNew", name, None, opts)


class ComponentResource(khulnasoft.ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:ComponentResource", name, None, opts)
        FooResource("child", khulnasoft.ResourceOptions(parent=self))


ComponentResource("comp")