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
from . import http_module as _http_module
from ._enums import *

__all__ = ['Example_resourceArgs', 'Example_resource']

@khulnasoft.input_type
class Example_resourceArgs:
    def __init__(__self__, *,
                 map_enum: Optional[khulnasoft.Input[Sequence[khulnasoft.Input[Mapping[str, khulnasoft.Input['Enum_XYZ']]]]]] = None,
                 request__http: Optional[khulnasoft.Input['_http_module.RequestArgs']] = None):
        """
        The set of arguments for constructing a Example_resource resource.
        """
        if map_enum is not None:
            khulnasoft.set(__self__, "map_enum", map_enum)
        if request__http is not None:
            khulnasoft.set(__self__, "request__http", request__http)

    @property
    @khulnasoft.getter
    def map_enum(self) -> Optional[khulnasoft.Input[Sequence[khulnasoft.Input[Mapping[str, khulnasoft.Input['Enum_XYZ']]]]]]:
        return khulnasoft.get(self, "map_enum")

    @map_enum.setter
    def map_enum(self, value: Optional[khulnasoft.Input[Sequence[khulnasoft.Input[Mapping[str, khulnasoft.Input['Enum_XYZ']]]]]]):
        khulnasoft.set(self, "map_enum", value)

    @property
    @khulnasoft.getter(name="request_HTTP")
    def request__http(self) -> Optional[khulnasoft.Input['_http_module.RequestArgs']]:
        return khulnasoft.get(self, "request__http")

    @request__http.setter
    def request__http(self, value: Optional[khulnasoft.Input['_http_module.RequestArgs']]):
        khulnasoft.set(self, "request__http", value)


class Example_resource(khulnasoft.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[khulnasoft.ResourceOptions] = None,
                 map_enum: Optional[khulnasoft.Input[Sequence[khulnasoft.Input[Mapping[str, khulnasoft.Input['Enum_XYZ']]]]]] = None,
                 request__http: Optional[khulnasoft.Input[khulnasoft.InputType['_http_module.RequestArgs']]] = None,
                 __props__=None):
        """
        Create a Example_resource resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param khulnasoft.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[Example_resourceArgs] = None,
                 opts: Optional[khulnasoft.ResourceOptions] = None):
        """
        Create a Example_resource resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param Example_resourceArgs args: The arguments to use to populate this resource's properties.
        :param khulnasoft.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(Example_resourceArgs, khulnasoft.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[khulnasoft.ResourceOptions] = None,
                 map_enum: Optional[khulnasoft.Input[Sequence[khulnasoft.Input[Mapping[str, khulnasoft.Input['Enum_XYZ']]]]]] = None,
                 request__http: Optional[khulnasoft.Input[khulnasoft.InputType['_http_module.RequestArgs']]] = None,
                 __props__=None):
        opts = khulnasoft.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, khulnasoft.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = Example_resourceArgs.__new__(Example_resourceArgs)

            __props__.__dict__["map_enum"] = map_enum
            __props__.__dict__["request__http"] = request__http
            __props__.__dict__["url"] = None
            __props__.__dict__["good__urls"] = None
        super(Example_resource, __self__).__init__(
            'legacy_names:index:example_resource',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: khulnasoft.Input[str],
            opts: Optional[khulnasoft.ResourceOptions] = None) -> 'Example_resource':
        """
        Get an existing Example_resource resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param khulnasoft.Input[str] id: The unique provider ID of the resource to lookup.
        :param khulnasoft.ResourceOptions opts: Options for the resource.
        """
        opts = khulnasoft.ResourceOptions.merge(opts, khulnasoft.ResourceOptions(id=id))

        __props__ = Example_resourceArgs.__new__(Example_resourceArgs)

        __props__.__dict__["url"] = None
        __props__.__dict__["good__urls"] = None
        __props__.__dict__["map_enum"] = None
        return Example_resource(resource_name, opts=opts, __props__=__props__)

    @property
    @khulnasoft.getter(name="URL")
    def url(self) -> khulnasoft.Output[Optional[str]]:
        return khulnasoft.get(self, "url")

    @property
    @khulnasoft.getter(name="good_URLs")
    def good__urls(self) -> khulnasoft.Output[Optional[Sequence[str]]]:
        return khulnasoft.get(self, "good__urls")

    @property
    @khulnasoft.getter
    def map_enum(self) -> khulnasoft.Output[Optional[Sequence[Mapping[str, 'Enum_XYZ']]]]:
        return khulnasoft.get(self, "map_enum")
