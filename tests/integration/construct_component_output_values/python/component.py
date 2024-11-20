# Copyright 2016-2021, Pulumi Corporation.  All rights reserved.

import khulnasoft
from typing import Mapping, Optional


@khulnasoft.input_type
class BarArgs:
    def __init__(__self__, *,
                 tags: Optional[khulnasoft.Input[Mapping[str, khulnasoft.Input[str]]]] = None):
        if tags is not None:
            khulnasoft.set(__self__, "tags", tags)

    @property
    @khulnasoft.getter
    def tags(self) -> Optional[khulnasoft.Input[Mapping[str, khulnasoft.Input[str]]]]:
        return khulnasoft.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[khulnasoft.Input[Mapping[str, khulnasoft.Input[str]]]]):
        khulnasoft.set(self, "tags", value)


@khulnasoft.input_type
class FooArgs:
    def __init__(__self__, *,
                 something: Optional[khulnasoft.Input[str]] = None):
        if something is not None:
            khulnasoft.set(__self__, "something", something)

    @property
    @khulnasoft.getter
    def something(self) -> Optional[khulnasoft.Input[str]]:
        return khulnasoft.get(self, "something")

    @something.setter
    def something(self, value: Optional[khulnasoft.Input[str]]):
        khulnasoft.set(self, "something", value)


@khulnasoft.input_type
class ComponentArgs:
    def __init__(__self__, *,
                 bar: Optional[khulnasoft.Input['BarArgs']] = None,
                 foo: Optional['FooArgs'] = None):
        if bar is not None:
            khulnasoft.set(__self__, "bar", bar)
        if foo is not None:
            khulnasoft.set(__self__, "foo", foo)

    @property
    @khulnasoft.getter
    def bar(self) -> Optional[khulnasoft.Input['BarArgs']]:
        return khulnasoft.get(self, "bar")

    @bar.setter
    def bar(self, value: Optional[khulnasoft.Input['BarArgs']]):
        khulnasoft.set(self, "bar", value)

    @property
    @khulnasoft.getter
    def foo(self) -> Optional['FooArgs']:
        return khulnasoft.get(self, "foo")

    @foo.setter
    def foo(self, value: Optional['FooArgs']):
        khulnasoft.set(self, "foo", value)


class Component(khulnasoft.ComponentResource):
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ComponentArgs] = None,
                 opts: Optional[khulnasoft.ResourceOptions] = None):
        super().__init__('testcomponent:index:Component', resource_name, args, opts, remote=True)
