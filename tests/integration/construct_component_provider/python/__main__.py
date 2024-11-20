# Copyright 2016-2021, Pulumi Corporation.  All rights reserved.

from typing import Optional

import khulnasoft


class Provider(khulnasoft.ProviderResource):
    message: khulnasoft.Output[str]

    def __init__(self, name: str, message: khulnasoft.Input[str], opts: Optional[khulnasoft.ResourceOptions] = None) -> None:
        super().__init__("testcomponent", name, {"message": message}, opts)


class Component(khulnasoft.ComponentResource):
    message: khulnasoft.Output[str]

    def __init__(self, name: str, opts: Optional[khulnasoft.ResourceOptions] = None) -> None:
        props = {
            "message": None
        }
        super().__init__("testcomponent:index:Component", name, props, opts, True)


component = Component("mycomponent", khulnasoft.ResourceOptions(
    providers={
        "testcomponent": Provider("myprovider", "hello world"),
    })
)


khulnasoft.export("message", component.message)
