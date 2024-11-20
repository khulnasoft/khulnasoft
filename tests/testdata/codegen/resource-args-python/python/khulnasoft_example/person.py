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

__all__ = ['PersonArgs', 'Person']

@khulnasoft.input_type
class PersonArgs:
    def __init__(__self__, *,
                 name: Optional[khulnasoft.Input[str]] = None,
                 pets: Optional[khulnasoft.Input[Sequence[khulnasoft.Input['PetArgs']]]] = None):
        """
        The set of arguments for constructing a Person resource.
        """
        if name is not None:
            khulnasoft.set(__self__, "name", name)
        if pets is not None:
            khulnasoft.set(__self__, "pets", pets)

    @property
    @khulnasoft.getter
    def name(self) -> Optional[khulnasoft.Input[str]]:
        return khulnasoft.get(self, "name")

    @name.setter
    def name(self, value: Optional[khulnasoft.Input[str]]):
        khulnasoft.set(self, "name", value)

    @property
    @khulnasoft.getter
    def pets(self) -> Optional[khulnasoft.Input[Sequence[khulnasoft.Input['PetArgs']]]]:
        return khulnasoft.get(self, "pets")

    @pets.setter
    def pets(self, value: Optional[khulnasoft.Input[Sequence[khulnasoft.Input['PetArgs']]]]):
        khulnasoft.set(self, "pets", value)


class Person(khulnasoft.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[khulnasoft.ResourceOptions] = None,
                 name: Optional[khulnasoft.Input[str]] = None,
                 pets: Optional[khulnasoft.Input[Sequence[khulnasoft.Input[Union['PetArgs', 'PetArgsDict']]]]] = None,
                 __props__=None):
        """
        Create a Person resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param khulnasoft.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[PersonArgs] = None,
                 opts: Optional[khulnasoft.ResourceOptions] = None):
        """
        Create a Person resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param PersonArgs args: The arguments to use to populate this resource's properties.
        :param khulnasoft.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PersonArgs, khulnasoft.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[khulnasoft.ResourceOptions] = None,
                 name: Optional[khulnasoft.Input[str]] = None,
                 pets: Optional[khulnasoft.Input[Sequence[khulnasoft.Input[Union['PetArgs', 'PetArgsDict']]]]] = None,
                 __props__=None):
        opts = khulnasoft.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, khulnasoft.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PersonArgs.__new__(PersonArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["pets"] = pets
        super(Person, __self__).__init__(
            'example::Person',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: khulnasoft.Input[str],
            opts: Optional[khulnasoft.ResourceOptions] = None) -> 'Person':
        """
        Get an existing Person resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param khulnasoft.Input[str] id: The unique provider ID of the resource to lookup.
        :param khulnasoft.ResourceOptions opts: Options for the resource.
        """
        opts = khulnasoft.ResourceOptions.merge(opts, khulnasoft.ResourceOptions(id=id))

        __props__ = PersonArgs.__new__(PersonArgs)

        __props__.__dict__["name"] = None
        __props__.__dict__["pets"] = None
        return Person(resource_name, opts=opts, __props__=__props__)

    @property
    @khulnasoft.getter
    def name(self) -> khulnasoft.Output[Optional[str]]:
        return khulnasoft.get(self, "name")

    @property
    @khulnasoft.getter
    def pets(self) -> khulnasoft.Output[Optional[Sequence['outputs.Pet']]]:
        return khulnasoft.get(self, "pets")

