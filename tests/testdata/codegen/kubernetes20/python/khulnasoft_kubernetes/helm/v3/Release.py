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
from ... import _utilities

__all__ = ['ReleaseArgs', 'Release']

@khulnasoft.input_type
class ReleaseArgs:
    def __init__(__self__, *,
                 chart: khulnasoft.Input[str],
                 value_yaml_files: Optional[khulnasoft.Input[Sequence[khulnasoft.Input[Union[khulnasoft.Asset, khulnasoft.Archive]]]]] = None,
                 values: Optional[khulnasoft.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a Release resource.
        :param khulnasoft.Input[str] chart: Chart name to be installed. A path may be used.
        :param khulnasoft.Input[Sequence[khulnasoft.Input[Union[khulnasoft.Asset, khulnasoft.Archive]]]] value_yaml_files: List of assets (raw yaml files). Content is read and merged with values.
        :param khulnasoft.Input[Mapping[str, Any]] values: Custom values set for the release.
        """
        khulnasoft.set(__self__, "chart", chart)
        if value_yaml_files is not None:
            khulnasoft.set(__self__, "value_yaml_files", value_yaml_files)
        if values is not None:
            khulnasoft.set(__self__, "values", values)

    @property
    @khulnasoft.getter
    def chart(self) -> khulnasoft.Input[str]:
        """
        Chart name to be installed. A path may be used.
        """
        return khulnasoft.get(self, "chart")

    @chart.setter
    def chart(self, value: khulnasoft.Input[str]):
        khulnasoft.set(self, "chart", value)

    @property
    @khulnasoft.getter(name="valueYamlFiles")
    def value_yaml_files(self) -> Optional[khulnasoft.Input[Sequence[khulnasoft.Input[Union[khulnasoft.Asset, khulnasoft.Archive]]]]]:
        """
        List of assets (raw yaml files). Content is read and merged with values.
        """
        return khulnasoft.get(self, "value_yaml_files")

    @value_yaml_files.setter
    def value_yaml_files(self, value: Optional[khulnasoft.Input[Sequence[khulnasoft.Input[Union[khulnasoft.Asset, khulnasoft.Archive]]]]]):
        khulnasoft.set(self, "value_yaml_files", value)

    @property
    @khulnasoft.getter
    def values(self) -> Optional[khulnasoft.Input[Mapping[str, Any]]]:
        """
        Custom values set for the release.
        """
        return khulnasoft.get(self, "values")

    @values.setter
    def values(self, value: Optional[khulnasoft.Input[Mapping[str, Any]]]):
        khulnasoft.set(self, "values", value)


class Release(khulnasoft.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[khulnasoft.ResourceOptions] = None,
                 chart: Optional[khulnasoft.Input[str]] = None,
                 value_yaml_files: Optional[khulnasoft.Input[Sequence[khulnasoft.Input[Union[khulnasoft.Asset, khulnasoft.Archive]]]]] = None,
                 values: Optional[khulnasoft.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        """
        A non-overlay, non-component, non-Kubernetes resource.

        :param str resource_name: The name of the resource.
        :param khulnasoft.ResourceOptions opts: Options for the resource.
        :param khulnasoft.Input[str] chart: Chart name to be installed. A path may be used.
        :param khulnasoft.Input[Sequence[khulnasoft.Input[Union[khulnasoft.Asset, khulnasoft.Archive]]]] value_yaml_files: List of assets (raw yaml files). Content is read and merged with values.
        :param khulnasoft.Input[Mapping[str, Any]] values: Custom values set for the release.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ReleaseArgs,
                 opts: Optional[khulnasoft.ResourceOptions] = None):
        """
        A non-overlay, non-component, non-Kubernetes resource.

        :param str resource_name: The name of the resource.
        :param ReleaseArgs args: The arguments to use to populate this resource's properties.
        :param khulnasoft.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ReleaseArgs, khulnasoft.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[khulnasoft.ResourceOptions] = None,
                 chart: Optional[khulnasoft.Input[str]] = None,
                 value_yaml_files: Optional[khulnasoft.Input[Sequence[khulnasoft.Input[Union[khulnasoft.Asset, khulnasoft.Archive]]]]] = None,
                 values: Optional[khulnasoft.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        opts = khulnasoft.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, khulnasoft.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ReleaseArgs.__new__(ReleaseArgs)

            if chart is None and not opts.urn:
                raise TypeError("Missing required property 'chart'")
            __props__.__dict__["chart"] = chart
            __props__.__dict__["value_yaml_files"] = value_yaml_files
            __props__.__dict__["values"] = values
        super(Release, __self__).__init__(
            'kubernetes:helm.sh/v3:Release',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: khulnasoft.Input[str],
            opts: Optional[khulnasoft.ResourceOptions] = None) -> 'Release':
        """
        Get an existing Release resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param khulnasoft.Input[str] id: The unique provider ID of the resource to lookup.
        :param khulnasoft.ResourceOptions opts: Options for the resource.
        """
        opts = khulnasoft.ResourceOptions.merge(opts, khulnasoft.ResourceOptions(id=id))

        __props__ = ReleaseArgs.__new__(ReleaseArgs)

        __props__.__dict__["chart"] = None
        __props__.__dict__["value_yaml_files"] = None
        __props__.__dict__["values"] = None
        return Release(resource_name, opts=opts, __props__=__props__)

    @property
    @khulnasoft.getter
    def chart(self) -> khulnasoft.Output[str]:
        """
        Chart name to be installed. A path may be used.
        """
        return khulnasoft.get(self, "chart")

    @property
    @khulnasoft.getter(name="valueYamlFiles")
    def value_yaml_files(self) -> khulnasoft.Output[Optional[Sequence[Union[khulnasoft.Asset, khulnasoft.Archive]]]]:
        """
        List of assets (raw yaml files). Content is read and merged with values (with values taking precedence).
        """
        return khulnasoft.get(self, "value_yaml_files")

    @property
    @khulnasoft.getter
    def values(self) -> khulnasoft.Output[Optional[Mapping[str, Any]]]:
        """
        Custom values set for the release.
        """
        return khulnasoft.get(self, "values")

