# coding=utf-8
# *** WARNING: this file was generated by khulnasoft-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import khulnasoft
import khulnasoft.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['FailsOnCreateArgs', 'FailsOnCreate']

@khulnasoft.input_type
class FailsOnCreateArgs:
    def __init__(__self__):
        """
        The set of arguments for constructing a FailsOnCreate resource.
        """
        pass


class FailsOnCreate(khulnasoft.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[khulnasoft.ResourceOptions] = None,
                 __props__=None):
        """
        A test resource fails on create.

        :param str resource_name: The name of the resource.
        :param khulnasoft.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[FailsOnCreateArgs] = None,
                 opts: Optional[khulnasoft.ResourceOptions] = None):
        """
        A test resource fails on create.

        :param str resource_name: The name of the resource.
        :param FailsOnCreateArgs args: The arguments to use to populate this resource's properties.
        :param khulnasoft.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FailsOnCreateArgs, khulnasoft.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[khulnasoft.ResourceOptions] = None,
                 __props__=None):
        opts = khulnasoft.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, khulnasoft.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FailsOnCreateArgs.__new__(FailsOnCreateArgs)

        super(FailsOnCreate, __self__).__init__(
            'pkg:index:FailsOnCreate',
            resource_name,
            __props__,
            opts,
            package_ref=_utilities.get_package())

    @staticmethod
    def get(resource_name: str,
            id: khulnasoft.Input[str],
            opts: Optional[khulnasoft.ResourceOptions] = None) -> 'FailsOnCreate':
        """
        Get an existing FailsOnCreate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param khulnasoft.Input[str] id: The unique provider ID of the resource to lookup.
        :param khulnasoft.ResourceOptions opts: Options for the resource.
        """
        opts = khulnasoft.ResourceOptions.merge(opts, khulnasoft.ResourceOptions(id=id))

        __props__ = FailsOnCreateArgs.__new__(FailsOnCreateArgs)

        return FailsOnCreate(resource_name, opts=opts, __props__=__props__)

