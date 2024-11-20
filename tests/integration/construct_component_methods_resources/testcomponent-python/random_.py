# Copyright 2016-2021, Pulumi Corporation.  All rights reserved.

from typing import Optional

import khulnasoft

class Random(khulnasoft.CustomResource):
    def __init__(self,
                 resource_name: str,
                 length: khulnasoft.Input[int],
                 opts: Optional[khulnasoft.ResourceOptions] = None):
        props = {
            "length": length,
            "result": None,
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
