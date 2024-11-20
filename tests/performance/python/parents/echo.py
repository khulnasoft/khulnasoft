# Copyright 2024, Pulumi Corporation.

from typing import Any, Optional

import khulnasoft

class Echo(khulnasoft.CustomResource):
    def __init__(
        self,
        resource_name: str,
        echo: khulnasoft.Input[Any],
        opts: Optional[khulnasoft.ResourceOptions] = None,
    ):
        props = {
            "echo": echo,
        }
        super().__init__("testprovider:index:Echo", resource_name, props, opts)

    @property
    @khulnasoft.getter
    def echo(self) -> khulnasoft.Output[Any]:
        return khulnasoft.get(self, "echo")
