# coding=utf-8
# *** WARNING: this file was generated by khulnasoft-language-python. ***
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
from . import outputs
from ._inputs import *

__all__ = ['ResourceArgs', 'Resource']

@khulnasoft.input_type
class ResourceArgs:
    def __init__(__self__, *,
                 data: 'DataArgs',
                 non_plain_data: Optional[khulnasoft.Input['DataArgs']] = None):
        """
        The set of arguments for constructing a Resource resource.
        :param khulnasoft.Input['DataArgs'] non_plain_data: A non plain input to compare against the plain inputs, as well as testing plain/non-plain nesting.
        """
        khulnasoft.set(__self__, "data", data)
        if non_plain_data is not None:
            khulnasoft.set(__self__, "non_plain_data", non_plain_data)

    @property
    @khulnasoft.getter
    def data(self) -> 'DataArgs':
        return khulnasoft.get(self, "data")

    @data.setter
    def data(self, value: 'DataArgs'):
        khulnasoft.set(self, "data", value)

    @property
    @khulnasoft.getter(name="nonPlainData")
    def non_plain_data(self) -> Optional[khulnasoft.Input['DataArgs']]:
        """
        A non plain input to compare against the plain inputs, as well as testing plain/non-plain nesting.
        """
        return khulnasoft.get(self, "non_plain_data")

    @non_plain_data.setter
    def non_plain_data(self, value: Optional[khulnasoft.Input['DataArgs']]):
        khulnasoft.set(self, "non_plain_data", value)


class Resource(khulnasoft.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[khulnasoft.ResourceOptions] = None,
                 data: Optional[Union['DataArgs', 'DataArgsDict']] = None,
                 non_plain_data: Optional[khulnasoft.Input[Union['DataArgs', 'DataArgsDict']]] = None,
                 __props__=None):
        """
        Create a Resource resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param khulnasoft.ResourceOptions opts: Options for the resource.
        :param khulnasoft.Input[Union['DataArgs', 'DataArgsDict']] non_plain_data: A non plain input to compare against the plain inputs, as well as testing plain/non-plain nesting.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ResourceArgs,
                 opts: Optional[khulnasoft.ResourceOptions] = None):
        """
        Create a Resource resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ResourceArgs args: The arguments to use to populate this resource's properties.
        :param khulnasoft.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ResourceArgs, khulnasoft.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[khulnasoft.ResourceOptions] = None,
                 data: Optional[Union['DataArgs', 'DataArgsDict']] = None,
                 non_plain_data: Optional[khulnasoft.Input[Union['DataArgs', 'DataArgsDict']]] = None,
                 __props__=None):
        opts = khulnasoft.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, khulnasoft.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ResourceArgs.__new__(ResourceArgs)

            if data is None and not opts.urn:
                raise TypeError("Missing required property 'data'")
            __props__.__dict__["data"] = data
            __props__.__dict__["non_plain_data"] = non_plain_data
        super(Resource, __self__).__init__(
            'plain:index:Resource',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: khulnasoft.Input[str],
            opts: Optional[khulnasoft.ResourceOptions] = None) -> 'Resource':
        """
        Get an existing Resource resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param khulnasoft.Input[str] id: The unique provider ID of the resource to lookup.
        :param khulnasoft.ResourceOptions opts: Options for the resource.
        """
        opts = khulnasoft.ResourceOptions.merge(opts, khulnasoft.ResourceOptions(id=id))

        __props__ = ResourceArgs.__new__(ResourceArgs)

        __props__.__dict__["data"] = None
        __props__.__dict__["non_plain_data"] = None
        return Resource(resource_name, opts=opts, __props__=__props__)

    @property
    @khulnasoft.getter
    def data(self) -> khulnasoft.Output['outputs.Data']:
        return khulnasoft.get(self, "data")

    @property
    @khulnasoft.getter(name="nonPlainData")
    def non_plain_data(self) -> khulnasoft.Output[Optional['outputs.Data']]:
        """
        A non plain input to compare against the plain inputs, as well as testing plain/non-plain nesting.
        """
        return khulnasoft.get(self, "non_plain_data")

