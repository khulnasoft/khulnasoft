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
from .. import _utilities
from .. import _inputs as _root_inputs
from .. import outputs as _root_outputs

__all__ = ['ModuleResourceArgs', 'ModuleResource']

@khulnasoft.input_type
class ModuleResourceArgs:
    def __init__(__self__, *,
                 thing: Optional[khulnasoft.Input['_root_inputs.TopLevelArgs']] = None):
        """
        The set of arguments for constructing a ModuleResource resource.
        """
        if thing is not None:
            khulnasoft.set(__self__, "thing", thing)

    @property
    @khulnasoft.getter
    def thing(self) -> Optional[khulnasoft.Input['_root_inputs.TopLevelArgs']]:
        return khulnasoft.get(self, "thing")

    @thing.setter
    def thing(self, value: Optional[khulnasoft.Input['_root_inputs.TopLevelArgs']]):
        khulnasoft.set(self, "thing", value)


class ModuleResource(khulnasoft.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[khulnasoft.ResourceOptions] = None,
                 thing: Optional[khulnasoft.Input[Union['_root_inputs.TopLevelArgs', '_root_inputs.TopLevelArgsDict']]] = None,
                 __props__=None):
        """
        Create a ModuleResource resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param khulnasoft.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ModuleResourceArgs] = None,
                 opts: Optional[khulnasoft.ResourceOptions] = None):
        """
        Create a ModuleResource resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ModuleResourceArgs args: The arguments to use to populate this resource's properties.
        :param khulnasoft.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ModuleResourceArgs, khulnasoft.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[khulnasoft.ResourceOptions] = None,
                 thing: Optional[khulnasoft.Input[Union['_root_inputs.TopLevelArgs', '_root_inputs.TopLevelArgsDict']]] = None,
                 __props__=None):
        opts = khulnasoft.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, khulnasoft.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ModuleResourceArgs.__new__(ModuleResourceArgs)

            __props__.__dict__["thing"] = thing
        super(ModuleResource, __self__).__init__(
            'foo-bar:submodule1:ModuleResource',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: khulnasoft.Input[str],
            opts: Optional[khulnasoft.ResourceOptions] = None) -> 'ModuleResource':
        """
        Get an existing ModuleResource resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param khulnasoft.Input[str] id: The unique provider ID of the resource to lookup.
        :param khulnasoft.ResourceOptions opts: Options for the resource.
        """
        opts = khulnasoft.ResourceOptions.merge(opts, khulnasoft.ResourceOptions(id=id))

        __props__ = ModuleResourceArgs.__new__(ModuleResourceArgs)

        __props__.__dict__["thing"] = None
        return ModuleResource(resource_name, opts=opts, __props__=__props__)

    @property
    @khulnasoft.getter
    def thing(self) -> khulnasoft.Output[Optional['_root_outputs.TopLevel']]:
        return khulnasoft.get(self, "thing")

