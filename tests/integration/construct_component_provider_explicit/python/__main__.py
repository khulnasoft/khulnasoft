# Copyright 2016-2023, Pulumi Corporation.  All rights reserved.

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


class LocalComponent(khulnasoft.ComponentResource):
    message: khulnasoft.Output[str]

    def __init__(self, name: str, opts: Optional[khulnasoft.ResourceOptions] = None) -> None:
        super().__init__("my:index:LocalComponent", name, {}, opts)

        component = Component(f"{name}-mycomponent", khulnasoft.ResourceOptions(parent=self))
        self.message = component.message


provider = Provider("myprovider", "hello world")
component = Component("mycomponent", khulnasoft.ResourceOptions(
    provider=provider,
))
localComponent = LocalComponent("mylocalcomponent", khulnasoft.ResourceOptions(
    providers=[provider],
))

khulnasoft.export("message", component.message)
khulnasoft.export("nestedMessage", localComponent.message)
