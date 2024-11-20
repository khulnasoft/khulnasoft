# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import khulnasoft
import khulnasoft.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities
from ._inputs import *
import khulnasoft_kubernetes

__all__ = ['ComponentArgs', 'Component']

@khulnasoft.input_type
class ComponentArgs:
    def __init__(__self__, *,
                 my_type: Optional[khulnasoft.Input['MyTypeArgs']] = None):
        """
        The set of arguments for constructing a Component resource.
        """
        if my_type is not None:
            khulnasoft.set(__self__, "my_type", my_type)

    @property
    @khulnasoft.getter(name="myType")
    def my_type(self) -> Optional[khulnasoft.Input['MyTypeArgs']]:
        return khulnasoft.get(self, "my_type")

    @my_type.setter
    def my_type(self, value: Optional[khulnasoft.Input['MyTypeArgs']]):
        khulnasoft.set(self, "my_type", value)


class Component(khulnasoft.ComponentResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[khulnasoft.ResourceOptions] = None,
                 my_type: Optional[khulnasoft.Input[Union['MyTypeArgs', 'MyTypeArgsDict']]] = None,
                 __props__=None):
        """
        Create a Component resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param khulnasoft.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ComponentArgs] = None,
                 opts: Optional[khulnasoft.ResourceOptions] = None):
        """
        Create a Component resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ComponentArgs args: The arguments to use to populate this resource's properties.
        :param khulnasoft.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ComponentArgs, khulnasoft.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[khulnasoft.ResourceOptions] = None,
                 my_type: Optional[khulnasoft.Input[Union['MyTypeArgs', 'MyTypeArgsDict']]] = None,
                 __props__=None):
        opts = khulnasoft.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, khulnasoft.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is not None:
            raise ValueError('ComponentResource classes do not support opts.id')
        else:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ComponentArgs.__new__(ComponentArgs)

            __props__.__dict__["my_type"] = my_type
        super(Component, __self__).__init__(
            'typedDictExample:index:Component',
            resource_name,
            __props__,
            opts,
            remote=True)
