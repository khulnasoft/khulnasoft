#!/usr/bin/env python3

from typing import Any, Optional

from khulnasoft import Resource, export
from khulnasoft.resource import ProviderResource as Provider
from khulnasoft.resource import ResourceOptions


class Random(Resource):
    def __init__(self, name: str, length=int, opts: Optional[ResourceOptions]=None):
        props = {"length": length, "result": None}
        self.length = length
        Resource.__init__(self, "testprovider:index:Random", name, True, props, opts)
        print(props)

class RandomProvider(Provider):
    def __init__(self, opts: Optional[ResourceOptions]=None):
        Provider.__init__(self, "testprovider", "provider", None, opts)

example_url = ResourceOptions(plugin_download_url="get.example.test")
provider_url = ResourceOptions(plugin_download_url="get.khulnasoft.test/providers")

# Create resource with specified PluginDownloadURL
r = Random("default", length=10, opts=example_url)
export("default provider", r.result)

# Create provider with specified PluginDownloadURL
provider = RandomProvider(provider_url)
# Create resource that inherits the providers PluginDownloadURL
e = Random("provided", length=8, opts=ResourceOptions(provider=provider))

export("explicit provider", e.result)
