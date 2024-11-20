# Copyright 2016-2021, Pulumi Corporation.  All rights reserved.

from typing import Optional

import khulnasoft

class Component(khulnasoft.ComponentResource):
    def __init__(self,
                 name: str,
                 opts: Optional[khulnasoft.ResourceOptions] = None):
        super().__init__("testcomponent:index:Component", name, {}, opts, True)

    @khulnasoft.output_type
    class CreateRandomResult:
        def __init__(self, result: str):
            if result and not isinstance(result, str):
                raise TypeError("Expected argument 'result' to be a str")
            khulnasoft.set(self, "result", result)

        @property
        @khulnasoft.getter
        def result(self) -> str:
            return khulnasoft.get(self, "result")

    def create_random(__self__, length: khulnasoft.Input[int]) -> khulnasoft.Output['Component.CreateRandomResult']:
        __args__ = dict()
        __args__['__self__'] = __self__
        __args__['length'] = length
        return khulnasoft.runtime.call('testcomponent:index:Component/createRandom',
                                   __args__,
                                   res=__self__,
                                   typ=Component.CreateRandomResult)
