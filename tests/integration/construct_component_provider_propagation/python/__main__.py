# Copyright 2016-2023, Pulumi Corporation.  All rights reserved.

from typing import Optional

import khulnasoft

class Component(khulnasoft.ComponentResource):
    """
    Python-level remote component for the component resource
    defined in sibling testcomponent-go directory.
    """

    result: khulnasoft.Output[str]

    def __init__(self,
                 name: str,
                 opts: Optional[khulnasoft.ResourceOptions] = None):
        props = {"result": None}
        super().__init__("testcomponent:index:Component", name, props, opts, remote=True)


class RandomProvider(khulnasoft.ProviderResource):
    """
    Provider for the testprovider:index:Random resource.

    Implemented in tests/testprovider.
    """

    def __init__(self, name, opts: Optional[khulnasoft.ResourceOptions]=None):
        super().__init__("testprovider", name, {}, opts)


explicit_provider = RandomProvider("explicit")

# Should pick up the default provider.
Component("uses_default")

# Should use the provider passed in as an argument.
Component("uses_provider", opts=khulnasoft.ResourceOptions(
    provider=explicit_provider,
))

# Should use the provider passed in as an argument
Component("uses_providers", opts=khulnasoft.ResourceOptions(
    providers=[explicit_provider],
))

# Should use the provider passed in as an argument
Component("uses_providers_map", opts=khulnasoft.ResourceOptions(
    providers={"testprovider": explicit_provider},
))
