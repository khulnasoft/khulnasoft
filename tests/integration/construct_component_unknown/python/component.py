# Copyright 2016-2021, Pulumi Corporation.  All rights reserved.

from typing import Optional

import khulnasoft

class Component(khulnasoft.ComponentResource):
    def __init__(self, name: str, args: khulnasoft.Inputs, opts: Optional[khulnasoft.ResourceOptions] = None):
        super().__init__("testcomponent:index:Component", name, args, opts, True)
