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

__all__ = ['ConfigGroupArgs', 'ConfigGroup']

@khulnasoft.input_type
class ConfigGroupArgs:
    def __init__(__self__, *,
                 files: Optional[khulnasoft.Input[Union[str, Sequence[khulnasoft.Input[str]]]]] = None,
                 objs: Optional[khulnasoft.Input[Union[Any, Sequence[Any]]]] = None,
                 resource_prefix: Optional[khulnasoft.Input[str]] = None,
                 yaml: Optional[khulnasoft.Input[Union[str, Sequence[khulnasoft.Input[str]]]]] = None):
        """
        The set of arguments for constructing a ConfigGroup resource.
        :param khulnasoft.Input[Union[str, Sequence[khulnasoft.Input[str]]]] files: Set of paths or a URLs that uniquely identify files.
        :param khulnasoft.Input[Union[Any, Sequence[Any]]] objs: Objects representing Kubernetes resources.
        :param khulnasoft.Input[str] resource_prefix: An optional prefix for the auto-generated resource names. Example: A resource created with resourcePrefix="foo" would produce a resource named "foo-resourceName".
        :param khulnasoft.Input[Union[str, Sequence[khulnasoft.Input[str]]]] yaml: YAML text containing Kubernetes resource definitions.
        """
        if files is not None:
            khulnasoft.set(__self__, "files", files)
        if objs is not None:
            khulnasoft.set(__self__, "objs", objs)
        if resource_prefix is not None:
            khulnasoft.set(__self__, "resource_prefix", resource_prefix)
        if yaml is not None:
            khulnasoft.set(__self__, "yaml", yaml)

    @property
    @khulnasoft.getter
    def files(self) -> Optional[khulnasoft.Input[Union[str, Sequence[khulnasoft.Input[str]]]]]:
        """
        Set of paths or a URLs that uniquely identify files.
        """
        return khulnasoft.get(self, "files")

    @files.setter
    def files(self, value: Optional[khulnasoft.Input[Union[str, Sequence[khulnasoft.Input[str]]]]]):
        khulnasoft.set(self, "files", value)

    @property
    @khulnasoft.getter
    def objs(self) -> Optional[khulnasoft.Input[Union[Any, Sequence[Any]]]]:
        """
        Objects representing Kubernetes resources.
        """
        return khulnasoft.get(self, "objs")

    @objs.setter
    def objs(self, value: Optional[khulnasoft.Input[Union[Any, Sequence[Any]]]]):
        khulnasoft.set(self, "objs", value)

    @property
    @khulnasoft.getter(name="resourcePrefix")
    def resource_prefix(self) -> Optional[khulnasoft.Input[str]]:
        """
        An optional prefix for the auto-generated resource names. Example: A resource created with resourcePrefix="foo" would produce a resource named "foo-resourceName".
        """
        return khulnasoft.get(self, "resource_prefix")

    @resource_prefix.setter
    def resource_prefix(self, value: Optional[khulnasoft.Input[str]]):
        khulnasoft.set(self, "resource_prefix", value)

    @property
    @khulnasoft.getter
    def yaml(self) -> Optional[khulnasoft.Input[Union[str, Sequence[khulnasoft.Input[str]]]]]:
        """
        YAML text containing Kubernetes resource definitions.
        """
        return khulnasoft.get(self, "yaml")

    @yaml.setter
    def yaml(self, value: Optional[khulnasoft.Input[Union[str, Sequence[khulnasoft.Input[str]]]]]):
        khulnasoft.set(self, "yaml", value)


class ConfigGroup(khulnasoft.ComponentResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[khulnasoft.ResourceOptions] = None,
                 files: Optional[khulnasoft.Input[Union[str, Sequence[khulnasoft.Input[str]]]]] = None,
                 objs: Optional[khulnasoft.Input[Union[Any, Sequence[Any]]]] = None,
                 resource_prefix: Optional[khulnasoft.Input[str]] = None,
                 yaml: Optional[khulnasoft.Input[Union[str, Sequence[khulnasoft.Input[str]]]]] = None,
                 __props__=None):
        """
        A non-overlay component resource.

        :param str resource_name: The name of the resource.
        :param khulnasoft.ResourceOptions opts: Options for the resource.
        :param khulnasoft.Input[Union[str, Sequence[khulnasoft.Input[str]]]] files: Set of paths or a URLs that uniquely identify files.
        :param khulnasoft.Input[Union[Any, Sequence[Any]]] objs: Objects representing Kubernetes resources.
        :param khulnasoft.Input[str] resource_prefix: An optional prefix for the auto-generated resource names. Example: A resource created with resourcePrefix="foo" would produce a resource named "foo-resourceName".
        :param khulnasoft.Input[Union[str, Sequence[khulnasoft.Input[str]]]] yaml: YAML text containing Kubernetes resource definitions.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ConfigGroupArgs] = None,
                 opts: Optional[khulnasoft.ResourceOptions] = None):
        """
        A non-overlay component resource.

        :param str resource_name: The name of the resource.
        :param ConfigGroupArgs args: The arguments to use to populate this resource's properties.
        :param khulnasoft.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConfigGroupArgs, khulnasoft.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[khulnasoft.ResourceOptions] = None,
                 files: Optional[khulnasoft.Input[Union[str, Sequence[khulnasoft.Input[str]]]]] = None,
                 objs: Optional[khulnasoft.Input[Union[Any, Sequence[Any]]]] = None,
                 resource_prefix: Optional[khulnasoft.Input[str]] = None,
                 yaml: Optional[khulnasoft.Input[Union[str, Sequence[khulnasoft.Input[str]]]]] = None,
                 __props__=None):
        opts = khulnasoft.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, khulnasoft.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is not None:
            raise ValueError('ComponentResource classes do not support opts.id')
        else:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ConfigGroupArgs.__new__(ConfigGroupArgs)

            __props__.__dict__["files"] = files
            __props__.__dict__["objs"] = objs
            __props__.__dict__["resource_prefix"] = resource_prefix
            __props__.__dict__["yaml"] = yaml
            __props__.__dict__["resources"] = None
        super(ConfigGroup, __self__).__init__(
            'kubernetes:yaml/v2:ConfigGroup',
            resource_name,
            __props__,
            opts,
            remote=True)

    @property
    @khulnasoft.getter
    def resources(self) -> khulnasoft.Output[Optional[Sequence[Any]]]:
        """
        Resources created by the ConfigGroup.
        """
        return khulnasoft.get(self, "resources")

