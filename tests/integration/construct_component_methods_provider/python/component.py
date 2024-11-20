# Copyright 2016-2023, Pulumi Corporation.  All rights reserved.

from typing import Optional

import khulnasoft

class Component(khulnasoft.ComponentResource):
    def __init__(self,
                 name: str,
                 first: khulnasoft.Input[str],
                 second: khulnasoft.Input[str],
                 opts: Optional[khulnasoft.ResourceOptions] = None):
        props = {
            "first": first,
            "second": second,
        }
        super().__init__("testcomponent:index:Component", name, props, opts, True)

    @khulnasoft.output_type
    class GetMessageResult:
        def __init__(self, message: str):
            if message and not isinstance(message, str):
                 raise TypeError("Expected argument 'message' to be a str")
            khulnasoft.set(self, "message", message)

        @property
        @khulnasoft.getter
        def message(self) -> str:
            return khulnasoft.get(self, "message")

    def get_message(__self__, name: khulnasoft.Input[str]) -> khulnasoft.Output['Component.GetMessageResult']:
        __args__ = dict()
        __args__['__self__'] = __self__
        __args__['name'] = name
        return khulnasoft.runtime.call('testcomponent:index:Component/getMessage',
                                   __args__,
                                   res=__self__,
                                   typ=Component.GetMessageResult)
