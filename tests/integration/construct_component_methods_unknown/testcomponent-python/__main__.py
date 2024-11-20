# Copyright 2016-2021, Pulumi Corporation.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

from typing import Optional
import sys

import khulnasoft
import khulnasoft.provider as provider


def panic(text: str):
    print(text)
    sys.exit(1)


class Component(khulnasoft.ComponentResource):
    def __init__(self,
                 resource_name: str,
                 opts: Optional[khulnasoft.ResourceOptions] = None) -> None:
        super().__init__("testcomponent:index:Component", resource_name, {}, opts)

    def get_message(self, echo: khulnasoft.Input[str]) -> khulnasoft.Output[str]:
        return khulnasoft.Output.from_input(echo).apply(lambda v: panic("should not run (echo)"))

class Provider(provider.Provider):
    VERSION = "0.0.1"

    class Module(khulnasoft.runtime.ResourceModule):
        def version(self):
            return Provider.VERSION

        def construct(self, name: str, typ: str, urn: str) -> khulnasoft.Resource:
            if typ == "testcomponent:index:Component":
                return Component(name, khulnasoft.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")

    def __init__(self):
        super().__init__(Provider.VERSION)
        khulnasoft.runtime.register_resource_module("testcomponent", "index", Provider.Module())

    def construct(self, name: str, resource_type: str, inputs: khulnasoft.Inputs,
                  options: Optional[khulnasoft.ResourceOptions] = None) -> provider.ConstructResult:

        if resource_type != "testcomponent:index:Component":
            raise Exception(f"unknown resource type {resource_type}")

        component = Component(name, options)

        return provider.ConstructResult(
            urn=component.urn,
            state=inputs)

    def call(self, token: str, args: khulnasoft.Inputs) -> provider.CallResult:
        if token != "testcomponent:index:Component/getMessage":
            raise Exception(f'unknown method {token}')

        comp: Component = args["__self__"]
        outputs = {
            "message": comp.get_message(args["echo"])
        }
        return provider.CallResult(outputs=outputs)


if __name__ == "__main__":
    provider.main(Provider(), sys.argv[1:])
