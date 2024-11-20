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
from . import outputs
from ._inputs import *

__all__ = ['ResourceWithAssetsArgs', 'ResourceWithAssets']

@khulnasoft.input_type
class ResourceWithAssetsArgs:
    def __init__(__self__, *,
                 source: khulnasoft.Input[Union[khulnasoft.Asset, khulnasoft.Archive]],
                 archive: Optional[khulnasoft.Input[khulnasoft.Archive]] = None,
                 nested: Optional[khulnasoft.Input['TypeWithAssetsArgs']] = None):
        """
        The set of arguments for constructing a ResourceWithAssets resource.
        """
        khulnasoft.set(__self__, "source", source)
        if archive is not None:
            khulnasoft.set(__self__, "archive", archive)
        if nested is not None:
            khulnasoft.set(__self__, "nested", nested)

    @property
    @khulnasoft.getter
    def source(self) -> khulnasoft.Input[Union[khulnasoft.Asset, khulnasoft.Archive]]:
        return khulnasoft.get(self, "source")

    @source.setter
    def source(self, value: khulnasoft.Input[Union[khulnasoft.Asset, khulnasoft.Archive]]):
        khulnasoft.set(self, "source", value)

    @property
    @khulnasoft.getter
    def archive(self) -> Optional[khulnasoft.Input[khulnasoft.Archive]]:
        return khulnasoft.get(self, "archive")

    @archive.setter
    def archive(self, value: Optional[khulnasoft.Input[khulnasoft.Archive]]):
        khulnasoft.set(self, "archive", value)

    @property
    @khulnasoft.getter
    def nested(self) -> Optional[khulnasoft.Input['TypeWithAssetsArgs']]:
        return khulnasoft.get(self, "nested")

    @nested.setter
    def nested(self, value: Optional[khulnasoft.Input['TypeWithAssetsArgs']]):
        khulnasoft.set(self, "nested", value)


class ResourceWithAssets(khulnasoft.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[khulnasoft.ResourceOptions] = None,
                 archive: Optional[khulnasoft.Input[khulnasoft.Archive]] = None,
                 nested: Optional[khulnasoft.Input[khulnasoft.InputType['TypeWithAssetsArgs']]] = None,
                 source: Optional[khulnasoft.Input[Union[khulnasoft.Asset, khulnasoft.Archive]]] = None,
                 __props__=None):
        """
        Create a ResourceWithAssets resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param khulnasoft.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ResourceWithAssetsArgs,
                 opts: Optional[khulnasoft.ResourceOptions] = None):
        """
        Create a ResourceWithAssets resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ResourceWithAssetsArgs args: The arguments to use to populate this resource's properties.
        :param khulnasoft.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ResourceWithAssetsArgs, khulnasoft.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[khulnasoft.ResourceOptions] = None,
                 archive: Optional[khulnasoft.Input[khulnasoft.Archive]] = None,
                 nested: Optional[khulnasoft.Input[khulnasoft.InputType['TypeWithAssetsArgs']]] = None,
                 source: Optional[khulnasoft.Input[Union[khulnasoft.Asset, khulnasoft.Archive]]] = None,
                 __props__=None):
        opts = khulnasoft.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, khulnasoft.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ResourceWithAssetsArgs.__new__(ResourceWithAssetsArgs)

            __props__.__dict__["archive"] = archive
            __props__.__dict__["nested"] = nested
            if source is None and not opts.urn:
                raise TypeError("Missing required property 'source'")
            __props__.__dict__["source"] = source
            __props__.__dict__["asset"] = None
        super(ResourceWithAssets, __self__).__init__(
            'example:index:ResourceWithAssets',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: khulnasoft.Input[str],
            opts: Optional[khulnasoft.ResourceOptions] = None) -> 'ResourceWithAssets':
        """
        Get an existing ResourceWithAssets resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param khulnasoft.Input[str] id: The unique provider ID of the resource to lookup.
        :param khulnasoft.ResourceOptions opts: Options for the resource.
        """
        opts = khulnasoft.ResourceOptions.merge(opts, khulnasoft.ResourceOptions(id=id))

        __props__ = ResourceWithAssetsArgs.__new__(ResourceWithAssetsArgs)

        __props__.__dict__["archive"] = None
        __props__.__dict__["asset"] = None
        __props__.__dict__["nested"] = None
        return ResourceWithAssets(resource_name, opts=opts, __props__=__props__)

    @property
    @khulnasoft.getter
    def archive(self) -> khulnasoft.Output[khulnasoft.Archive]:
        return khulnasoft.get(self, "archive")

    @property
    @khulnasoft.getter
    def asset(self) -> khulnasoft.Output[Union[khulnasoft.Asset, khulnasoft.Archive]]:
        return khulnasoft.get(self, "asset")

    @property
    @khulnasoft.getter
    def nested(self) -> khulnasoft.Output[Optional['outputs.TypeWithAssets']]:
        return khulnasoft.get(self, "nested")
