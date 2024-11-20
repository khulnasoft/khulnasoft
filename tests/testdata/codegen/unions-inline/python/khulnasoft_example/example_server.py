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

__all__ = ['ExampleServerArgs', 'ExampleServer']

@khulnasoft.input_type
class ExampleServerArgs:
    def __init__(__self__, *,
                 properties: Optional[khulnasoft.Input[Union['ServerPropertiesForReplicaArgs', 'ServerPropertiesForRestoreArgs']]] = None):
        """
        The set of arguments for constructing a ExampleServer resource.
        """
        if properties is not None:
            khulnasoft.set(__self__, "properties", properties)

    @property
    @khulnasoft.getter
    def properties(self) -> Optional[khulnasoft.Input[Union['ServerPropertiesForReplicaArgs', 'ServerPropertiesForRestoreArgs']]]:
        return khulnasoft.get(self, "properties")

    @properties.setter
    def properties(self, value: Optional[khulnasoft.Input[Union['ServerPropertiesForReplicaArgs', 'ServerPropertiesForRestoreArgs']]]):
        khulnasoft.set(self, "properties", value)


class ExampleServer(khulnasoft.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[khulnasoft.ResourceOptions] = None,
                 properties: Optional[khulnasoft.Input[Union[khulnasoft.InputType['ServerPropertiesForReplicaArgs'], khulnasoft.InputType['ServerPropertiesForRestoreArgs']]]] = None,
                 __props__=None):
        """
        Create a ExampleServer resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param khulnasoft.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ExampleServerArgs] = None,
                 opts: Optional[khulnasoft.ResourceOptions] = None):
        """
        Create a ExampleServer resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ExampleServerArgs args: The arguments to use to populate this resource's properties.
        :param khulnasoft.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ExampleServerArgs, khulnasoft.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[khulnasoft.ResourceOptions] = None,
                 properties: Optional[khulnasoft.Input[Union[khulnasoft.InputType['ServerPropertiesForReplicaArgs'], khulnasoft.InputType['ServerPropertiesForRestoreArgs']]]] = None,
                 __props__=None):
        opts = khulnasoft.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, khulnasoft.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ExampleServerArgs.__new__(ExampleServerArgs)

            __props__.__dict__["properties"] = properties
            __props__.__dict__["name"] = None
        super(ExampleServer, __self__).__init__(
            'example:index:ExampleServer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: khulnasoft.Input[str],
            opts: Optional[khulnasoft.ResourceOptions] = None) -> 'ExampleServer':
        """
        Get an existing ExampleServer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param khulnasoft.Input[str] id: The unique provider ID of the resource to lookup.
        :param khulnasoft.ResourceOptions opts: Options for the resource.
        """
        opts = khulnasoft.ResourceOptions.merge(opts, khulnasoft.ResourceOptions(id=id))

        __props__ = ExampleServerArgs.__new__(ExampleServerArgs)

        __props__.__dict__["name"] = None
        return ExampleServer(resource_name, opts=opts, __props__=__props__)

    @property
    @khulnasoft.getter
    def name(self) -> khulnasoft.Output[Optional[str]]:
        return khulnasoft.get(self, "name")

