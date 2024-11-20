# Copyright 2016-2021, Pulumi Corporation.  All rights reserved.

from typing import Optional

import khulnasoft

class Random(khulnasoft.CustomResource):
    def __init__(self,
                 resource_name: str,
                 length: khulnasoft.Input[int],
                 prefix: Optional[khulnasoft.Input[str]] = None,
                 opts: Optional[khulnasoft.ResourceOptions] = None):
        props = {
            "length": length,
            "result": None,
            "prefix": prefix,
        }
        super().__init__("testprovider:index:Random", resource_name, props, opts)

    @property
    @khulnasoft.getter
    def length(self) -> khulnasoft.Output[int]:
        return khulnasoft.get(self, "length")

    @property
    @khulnasoft.getter
    def result(self) -> khulnasoft.Output[str]:
        return khulnasoft.get(self, "result")

    def invoke(self, args):
        return khulnasoft.runtime.invoke("testprovider:index:returnArgs", args)


class Component(khulnasoft.ComponentResource):
    def __init__(self,
                 resource_name: str,
                 length: khulnasoft.Input[int],
                 opts: Optional[khulnasoft.ResourceOptions] = None):
        props = {
            "length": length,
            "childId": None,
        }
        super().__init__("testprovider:index:Component", resource_name, props, opts, True)

    @property
    @khulnasoft.getter
    def length(self) -> khulnasoft.Output[int]:
        return khulnasoft.get(self, "length")

    @property
    @khulnasoft.getter
    def child_id(self) -> khulnasoft.Output[str]:
        return khulnasoft.get(self, "childId")

class Provider(khulnasoft.ProviderResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[khulnasoft.ResourceOptions] = None):
        super(Provider, __self__).__init__(
            'testprovider',
            resource_name,
            None,
            opts)
