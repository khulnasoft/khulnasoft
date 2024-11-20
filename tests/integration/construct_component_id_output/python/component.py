# Copyright 2016-2024, Pulumi Corporation.  All rights reserved.

from typing import Any, Optional

import khulnasoft

class Component(khulnasoft.ComponentResource):
    def __init__(self, name: str, id: khulnasoft.Input[str], opts: Optional[khulnasoft.ResourceOptions] = None):
        props = dict()
        props["id"] = id
        super().__init__("testcomponent:index:Component", name, props, opts, True)

    @property
    @khulnasoft.getter
    def id(self) -> khulnasoft.Output[str]:
        return khulnasoft.get(self, "id")