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
from ._enums import *

__all__ = ['ModuleResourceArgs', 'ModuleResource']

@khulnasoft.input_type
class ModuleResourceArgs:
    def __init__(__self__, *,
                 plain_required_bool: Optional[bool] = None,
                 plain_required_const: Optional[str] = None,
                 plain_required_number: Optional[float] = None,
                 plain_required_string: Optional[str] = None,
                 required_bool: Optional[khulnasoft.Input[bool]] = None,
                 required_enum: Optional[khulnasoft.Input['EnumThing']] = None,
                 required_number: Optional[khulnasoft.Input[float]] = None,
                 required_string: Optional[khulnasoft.Input[str]] = None,
                 optional_bool: Optional[khulnasoft.Input[bool]] = None,
                 optional_const: Optional[khulnasoft.Input[str]] = None,
                 optional_enum: Optional[khulnasoft.Input['EnumThing']] = None,
                 optional_number: Optional[khulnasoft.Input[float]] = None,
                 optional_string: Optional[khulnasoft.Input[str]] = None,
                 plain_optional_bool: Optional[bool] = None,
                 plain_optional_const: Optional[str] = None,
                 plain_optional_number: Optional[float] = None,
                 plain_optional_string: Optional[str] = None):
        """
        The set of arguments for constructing a ModuleResource resource.
        """
        if plain_required_bool is None:
            plain_required_bool = True
        khulnasoft.set(__self__, "plain_required_bool", plain_required_bool)
        if plain_required_const is None:
            plain_required_const = 'another'
        khulnasoft.set(__self__, "plain_required_const", 'val')
        if plain_required_number is None:
            plain_required_number = 42
        khulnasoft.set(__self__, "plain_required_number", plain_required_number)
        if plain_required_string is None:
            plain_required_string = 'buzzer'
        khulnasoft.set(__self__, "plain_required_string", plain_required_string)
        if required_bool is None:
            required_bool = True
        khulnasoft.set(__self__, "required_bool", required_bool)
        if required_enum is None:
            required_enum = 4
        khulnasoft.set(__self__, "required_enum", required_enum)
        if required_number is None:
            required_number = 42
        khulnasoft.set(__self__, "required_number", required_number)
        if required_string is None:
            required_string = 'buzzer'
        khulnasoft.set(__self__, "required_string", required_string)
        if optional_bool is None:
            optional_bool = True
        if optional_bool is not None:
            khulnasoft.set(__self__, "optional_bool", optional_bool)
        if optional_const is None:
            optional_const = 'another'
        if optional_const is not None:
            khulnasoft.set(__self__, "optional_const", 'val')
        if optional_enum is None:
            optional_enum = 8
        if optional_enum is not None:
            khulnasoft.set(__self__, "optional_enum", optional_enum)
        if optional_number is None:
            optional_number = 42
        if optional_number is not None:
            khulnasoft.set(__self__, "optional_number", optional_number)
        if optional_string is None:
            optional_string = 'buzzer'
        if optional_string is not None:
            khulnasoft.set(__self__, "optional_string", optional_string)
        if plain_optional_bool is None:
            plain_optional_bool = True
        if plain_optional_bool is not None:
            khulnasoft.set(__self__, "plain_optional_bool", plain_optional_bool)
        if plain_optional_const is None:
            plain_optional_const = 'another'
        if plain_optional_const is not None:
            khulnasoft.set(__self__, "plain_optional_const", 'val')
        if plain_optional_number is None:
            plain_optional_number = 42
        if plain_optional_number is not None:
            khulnasoft.set(__self__, "plain_optional_number", plain_optional_number)
        if plain_optional_string is None:
            plain_optional_string = 'buzzer'
        if plain_optional_string is not None:
            khulnasoft.set(__self__, "plain_optional_string", plain_optional_string)

    @property
    @khulnasoft.getter(name="plainRequiredBool")
    def plain_required_bool(self) -> bool:
        return khulnasoft.get(self, "plain_required_bool")

    @plain_required_bool.setter
    def plain_required_bool(self, value: bool):
        khulnasoft.set(self, "plain_required_bool", value)

    @property
    @khulnasoft.getter(name="plainRequiredConst")
    def plain_required_const(self) -> str:
        return khulnasoft.get(self, "plain_required_const")

    @plain_required_const.setter
    def plain_required_const(self, value: str):
        khulnasoft.set(self, "plain_required_const", value)

    @property
    @khulnasoft.getter(name="plainRequiredNumber")
    def plain_required_number(self) -> float:
        return khulnasoft.get(self, "plain_required_number")

    @plain_required_number.setter
    def plain_required_number(self, value: float):
        khulnasoft.set(self, "plain_required_number", value)

    @property
    @khulnasoft.getter(name="plainRequiredString")
    def plain_required_string(self) -> str:
        return khulnasoft.get(self, "plain_required_string")

    @plain_required_string.setter
    def plain_required_string(self, value: str):
        khulnasoft.set(self, "plain_required_string", value)

    @property
    @khulnasoft.getter(name="requiredBool")
    def required_bool(self) -> khulnasoft.Input[bool]:
        return khulnasoft.get(self, "required_bool")

    @required_bool.setter
    def required_bool(self, value: khulnasoft.Input[bool]):
        khulnasoft.set(self, "required_bool", value)

    @property
    @khulnasoft.getter(name="requiredEnum")
    def required_enum(self) -> khulnasoft.Input['EnumThing']:
        return khulnasoft.get(self, "required_enum")

    @required_enum.setter
    def required_enum(self, value: khulnasoft.Input['EnumThing']):
        khulnasoft.set(self, "required_enum", value)

    @property
    @khulnasoft.getter(name="requiredNumber")
    def required_number(self) -> khulnasoft.Input[float]:
        return khulnasoft.get(self, "required_number")

    @required_number.setter
    def required_number(self, value: khulnasoft.Input[float]):
        khulnasoft.set(self, "required_number", value)

    @property
    @khulnasoft.getter(name="requiredString")
    def required_string(self) -> khulnasoft.Input[str]:
        return khulnasoft.get(self, "required_string")

    @required_string.setter
    def required_string(self, value: khulnasoft.Input[str]):
        khulnasoft.set(self, "required_string", value)

    @property
    @khulnasoft.getter(name="optionalBool")
    def optional_bool(self) -> Optional[khulnasoft.Input[bool]]:
        return khulnasoft.get(self, "optional_bool")

    @optional_bool.setter
    def optional_bool(self, value: Optional[khulnasoft.Input[bool]]):
        khulnasoft.set(self, "optional_bool", value)

    @property
    @khulnasoft.getter(name="optionalConst")
    def optional_const(self) -> Optional[khulnasoft.Input[str]]:
        return khulnasoft.get(self, "optional_const")

    @optional_const.setter
    def optional_const(self, value: Optional[khulnasoft.Input[str]]):
        khulnasoft.set(self, "optional_const", value)

    @property
    @khulnasoft.getter(name="optionalEnum")
    def optional_enum(self) -> Optional[khulnasoft.Input['EnumThing']]:
        return khulnasoft.get(self, "optional_enum")

    @optional_enum.setter
    def optional_enum(self, value: Optional[khulnasoft.Input['EnumThing']]):
        khulnasoft.set(self, "optional_enum", value)

    @property
    @khulnasoft.getter(name="optionalNumber")
    def optional_number(self) -> Optional[khulnasoft.Input[float]]:
        return khulnasoft.get(self, "optional_number")

    @optional_number.setter
    def optional_number(self, value: Optional[khulnasoft.Input[float]]):
        khulnasoft.set(self, "optional_number", value)

    @property
    @khulnasoft.getter(name="optionalString")
    def optional_string(self) -> Optional[khulnasoft.Input[str]]:
        return khulnasoft.get(self, "optional_string")

    @optional_string.setter
    def optional_string(self, value: Optional[khulnasoft.Input[str]]):
        khulnasoft.set(self, "optional_string", value)

    @property
    @khulnasoft.getter(name="plainOptionalBool")
    def plain_optional_bool(self) -> Optional[bool]:
        return khulnasoft.get(self, "plain_optional_bool")

    @plain_optional_bool.setter
    def plain_optional_bool(self, value: Optional[bool]):
        khulnasoft.set(self, "plain_optional_bool", value)

    @property
    @khulnasoft.getter(name="plainOptionalConst")
    def plain_optional_const(self) -> Optional[str]:
        return khulnasoft.get(self, "plain_optional_const")

    @plain_optional_const.setter
    def plain_optional_const(self, value: Optional[str]):
        khulnasoft.set(self, "plain_optional_const", value)

    @property
    @khulnasoft.getter(name="plainOptionalNumber")
    def plain_optional_number(self) -> Optional[float]:
        return khulnasoft.get(self, "plain_optional_number")

    @plain_optional_number.setter
    def plain_optional_number(self, value: Optional[float]):
        khulnasoft.set(self, "plain_optional_number", value)

    @property
    @khulnasoft.getter(name="plainOptionalString")
    def plain_optional_string(self) -> Optional[str]:
        return khulnasoft.get(self, "plain_optional_string")

    @plain_optional_string.setter
    def plain_optional_string(self, value: Optional[str]):
        khulnasoft.set(self, "plain_optional_string", value)


class ModuleResource(khulnasoft.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[khulnasoft.ResourceOptions] = None,
                 optional_bool: Optional[khulnasoft.Input[bool]] = None,
                 optional_const: Optional[khulnasoft.Input[str]] = None,
                 optional_enum: Optional[khulnasoft.Input['EnumThing']] = None,
                 optional_number: Optional[khulnasoft.Input[float]] = None,
                 optional_string: Optional[khulnasoft.Input[str]] = None,
                 plain_optional_bool: Optional[bool] = None,
                 plain_optional_const: Optional[str] = None,
                 plain_optional_number: Optional[float] = None,
                 plain_optional_string: Optional[str] = None,
                 plain_required_bool: Optional[bool] = None,
                 plain_required_const: Optional[str] = None,
                 plain_required_number: Optional[float] = None,
                 plain_required_string: Optional[str] = None,
                 required_bool: Optional[khulnasoft.Input[bool]] = None,
                 required_enum: Optional[khulnasoft.Input['EnumThing']] = None,
                 required_number: Optional[khulnasoft.Input[float]] = None,
                 required_string: Optional[khulnasoft.Input[str]] = None,
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
                 args: ModuleResourceArgs,
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
                 optional_bool: Optional[khulnasoft.Input[bool]] = None,
                 optional_const: Optional[khulnasoft.Input[str]] = None,
                 optional_enum: Optional[khulnasoft.Input['EnumThing']] = None,
                 optional_number: Optional[khulnasoft.Input[float]] = None,
                 optional_string: Optional[khulnasoft.Input[str]] = None,
                 plain_optional_bool: Optional[bool] = None,
                 plain_optional_const: Optional[str] = None,
                 plain_optional_number: Optional[float] = None,
                 plain_optional_string: Optional[str] = None,
                 plain_required_bool: Optional[bool] = None,
                 plain_required_const: Optional[str] = None,
                 plain_required_number: Optional[float] = None,
                 plain_required_string: Optional[str] = None,
                 required_bool: Optional[khulnasoft.Input[bool]] = None,
                 required_enum: Optional[khulnasoft.Input['EnumThing']] = None,
                 required_number: Optional[khulnasoft.Input[float]] = None,
                 required_string: Optional[khulnasoft.Input[str]] = None,
                 __props__=None):
        opts = khulnasoft.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, khulnasoft.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ModuleResourceArgs.__new__(ModuleResourceArgs)

            if optional_bool is None:
                optional_bool = True
            __props__.__dict__["optional_bool"] = optional_bool
            if optional_const is None:
                optional_const = 'another'
            __props__.__dict__["optional_const"] = 'val'
            if optional_enum is None:
                optional_enum = 8
            __props__.__dict__["optional_enum"] = optional_enum
            if optional_number is None:
                optional_number = 42
            __props__.__dict__["optional_number"] = optional_number
            if optional_string is None:
                optional_string = 'buzzer'
            __props__.__dict__["optional_string"] = optional_string
            if plain_optional_bool is None:
                plain_optional_bool = True
            __props__.__dict__["plain_optional_bool"] = plain_optional_bool
            if plain_optional_const is None:
                plain_optional_const = 'another'
            __props__.__dict__["plain_optional_const"] = 'val'
            if plain_optional_number is None:
                plain_optional_number = 42
            __props__.__dict__["plain_optional_number"] = plain_optional_number
            if plain_optional_string is None:
                plain_optional_string = 'buzzer'
            __props__.__dict__["plain_optional_string"] = plain_optional_string
            if plain_required_bool is None:
                plain_required_bool = True
            if plain_required_bool is None and not opts.urn:
                raise TypeError("Missing required property 'plain_required_bool'")
            __props__.__dict__["plain_required_bool"] = plain_required_bool
            if plain_required_const is None:
                plain_required_const = 'another'
            if plain_required_const is None and not opts.urn:
                raise TypeError("Missing required property 'plain_required_const'")
            __props__.__dict__["plain_required_const"] = 'val'
            if plain_required_number is None:
                plain_required_number = 42
            if plain_required_number is None and not opts.urn:
                raise TypeError("Missing required property 'plain_required_number'")
            __props__.__dict__["plain_required_number"] = plain_required_number
            if plain_required_string is None:
                plain_required_string = 'buzzer'
            if plain_required_string is None and not opts.urn:
                raise TypeError("Missing required property 'plain_required_string'")
            __props__.__dict__["plain_required_string"] = plain_required_string
            if required_bool is None:
                required_bool = True
            if required_bool is None and not opts.urn:
                raise TypeError("Missing required property 'required_bool'")
            __props__.__dict__["required_bool"] = required_bool
            if required_enum is None:
                required_enum = 4
            if required_enum is None and not opts.urn:
                raise TypeError("Missing required property 'required_enum'")
            __props__.__dict__["required_enum"] = required_enum
            if required_number is None:
                required_number = 42
            if required_number is None and not opts.urn:
                raise TypeError("Missing required property 'required_number'")
            __props__.__dict__["required_number"] = required_number
            if required_string is None:
                required_string = 'buzzer'
            if required_string is None and not opts.urn:
                raise TypeError("Missing required property 'required_string'")
            __props__.__dict__["required_string"] = required_string
        super(ModuleResource, __self__).__init__(
            'foobar::ModuleResource',
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

        __props__.__dict__["optional_bool"] = None
        return ModuleResource(resource_name, opts=opts, __props__=__props__)

    @property
    @khulnasoft.getter(name="optionalBool")
    def optional_bool(self) -> khulnasoft.Output[Optional[bool]]:
        return khulnasoft.get(self, "optional_bool")
