# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

from typing import Any, Optional

import khulnasoft

class Component(khulnasoft.ComponentResource):
    echo: khulnasoft.Output[Any]
    childId: khulnasoft.Output[str]

    def __init__(self, name: str, echo: khulnasoft.Input[Any], opts: Optional[khulnasoft.ResourceOptions] = None):
        props = dict()
        props["echo"] = echo
        props["childId"] = None
        props["secret"] = None
        super().__init__("testcomponent:index:Component", name, props, opts, True)
