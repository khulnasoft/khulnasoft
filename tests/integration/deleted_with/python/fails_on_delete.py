# Copyright 2016-2023, Pulumi Corporation.  All rights reserved.

from typing import Optional

import khulnasoft

class FailsOnDelete(khulnasoft.CustomResource):
    def __init__(self,
                 resource_name: str,
                 opts: Optional[khulnasoft.ResourceOptions] = None):
        super().__init__("testprovider:index:FailsOnDelete", resource_name, {}, opts)
