# Copyright 2016-2024, Pulumi Corporation.  All rights reserved.

from typing import Optional

import khulnasoft

class Component(khulnasoft.ComponentResource):
    foo: khulnasoft.Output[str]

    def __init__(self, name: str, foo: khulnasoft.Input[str], opts: Optional[khulnasoft.ResourceOptions] = None):
        props = dict()
        props["foo"] = foo
        super().__init__("testcomponent:index:Component", name, props, opts, True)
