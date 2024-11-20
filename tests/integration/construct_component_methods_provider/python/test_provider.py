# Copyright 2016-2023, Pulumi Corporation.  All rights reserved.

import khulnasoft

class TestProvider(khulnasoft.ProviderResource):
    def __init__(__self__, resource_name: str):
        super().__init__("testprovider", resource_name)
