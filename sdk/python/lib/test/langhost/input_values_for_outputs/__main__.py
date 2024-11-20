# Copyright 2016-2022, Pulumi Corporation.
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

import khulnasoft


class Instance(khulnasoft.CustomResource):
    public_ip: khulnasoft.Output[str]

    def __init__(
        self,
        resource_name,
        name: khulnasoft.Input[str] = None,
        value: khulnasoft.Input[str] = None,
        opts=None,
    ):
        if opts is None:
            opts = khulnasoft.ResourceOptions()
        if name is None and not opts.urn:
            raise TypeError("Missing required property 'name'")
        __props__ = {}
        __props__["public_ip"] = None
        __props__["name"] = name
        __props__["value"] = value
        super(Instance, self).__init__(
            "aws:ec2/instance:Instance", resource_name, __props__, opts
        )


@khulnasoft.input_type
class DefaultLogGroupArgs:
    def __init__(self, *, skip: Optional[bool] = None):
        if skip is not None:
            khulnasoft.set(self, "skip", skip)

    @property
    @khulnasoft.getter
    def skip(self) -> Optional[bool]:
        return khulnasoft.get(self, "skip")

    @skip.setter
    def skip(self, value: Optional[bool]):
        khulnasoft.set(self, "skip", value)


@khulnasoft.input_type
class FargateTaskDefinitionArgs:
    def __init__(self, *, log_group: Optional[DefaultLogGroupArgs] = None):
        if log_group is not None:
            khulnasoft.set(self, "log_group", log_group)

    @property
    @khulnasoft.getter(name="logGroup")
    def log_group(self) -> Optional[DefaultLogGroupArgs]:
        return khulnasoft.get(self, "log_group")

    @log_group.setter
    def log_group(self, value: Optional[DefaultLogGroupArgs]):
        khulnasoft.set(self, "log_group", value)


# This resource has an input named `logGroup` typed as `DefaultLogGroupArgs` and an output named `logGroup` typed
# as `Instance`. When the provider returns no value for `logGroup`, it should not try to set the output to the
# input value due to the type mismatch.
class FargateTaskDefinition(khulnasoft.ComponentResource):
    def __init__(
        self,
        resource_name: str,
        log_group: Optional[khulnasoft.InputType[DefaultLogGroupArgs]] = None,
    ):
        __props__ = FargateTaskDefinitionArgs.__new__(FargateTaskDefinitionArgs)
        __props__.__dict__["log_group"] = log_group
        super().__init__(
            "awsx:ecs:FargateTaskDefinition",
            resource_name,
            __props__,
            None,
            remote=True,
        )

    @property
    @khulnasoft.getter(name="logGroup")
    def log_group(self) -> khulnasoft.Output[Optional[Instance]]:
        return khulnasoft.get(self, "log_group")


task_def = FargateTaskDefinition("task_def", log_group=DefaultLogGroupArgs(skip=True))
