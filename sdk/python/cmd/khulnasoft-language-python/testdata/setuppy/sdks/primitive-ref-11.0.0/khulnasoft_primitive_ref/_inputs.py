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

__all__ = [
    'DataArgs',
    'DataArgsDict',
]

MYPY = False

if not MYPY:
    class DataArgsDict(TypedDict):
        bool_array: khulnasoft.Input[Sequence[khulnasoft.Input[bool]]]
        boolean: khulnasoft.Input[bool]
        float: khulnasoft.Input[float]
        integer: khulnasoft.Input[int]
        string: khulnasoft.Input[str]
        string_map: khulnasoft.Input[Mapping[str, khulnasoft.Input[str]]]
elif False:
    DataArgsDict: TypeAlias = Mapping[str, Any]

@khulnasoft.input_type
class DataArgs:
    def __init__(__self__, *,
                 bool_array: khulnasoft.Input[Sequence[khulnasoft.Input[bool]]],
                 boolean: khulnasoft.Input[bool],
                 float: khulnasoft.Input[float],
                 integer: khulnasoft.Input[int],
                 string: khulnasoft.Input[str],
                 string_map: khulnasoft.Input[Mapping[str, khulnasoft.Input[str]]]):
        khulnasoft.set(__self__, "bool_array", bool_array)
        khulnasoft.set(__self__, "boolean", boolean)
        khulnasoft.set(__self__, "float", float)
        khulnasoft.set(__self__, "integer", integer)
        khulnasoft.set(__self__, "string", string)
        khulnasoft.set(__self__, "string_map", string_map)

    @property
    @khulnasoft.getter(name="boolArray")
    def bool_array(self) -> khulnasoft.Input[Sequence[khulnasoft.Input[bool]]]:
        return khulnasoft.get(self, "bool_array")

    @bool_array.setter
    def bool_array(self, value: khulnasoft.Input[Sequence[khulnasoft.Input[bool]]]):
        khulnasoft.set(self, "bool_array", value)

    @property
    @khulnasoft.getter
    def boolean(self) -> khulnasoft.Input[bool]:
        return khulnasoft.get(self, "boolean")

    @boolean.setter
    def boolean(self, value: khulnasoft.Input[bool]):
        khulnasoft.set(self, "boolean", value)

    @property
    @khulnasoft.getter
    def float(self) -> khulnasoft.Input[float]:
        return khulnasoft.get(self, "float")

    @float.setter
    def float(self, value: khulnasoft.Input[float]):
        khulnasoft.set(self, "float", value)

    @property
    @khulnasoft.getter
    def integer(self) -> khulnasoft.Input[int]:
        return khulnasoft.get(self, "integer")

    @integer.setter
    def integer(self, value: khulnasoft.Input[int]):
        khulnasoft.set(self, "integer", value)

    @property
    @khulnasoft.getter
    def string(self) -> khulnasoft.Input[str]:
        return khulnasoft.get(self, "string")

    @string.setter
    def string(self, value: khulnasoft.Input[str]):
        khulnasoft.set(self, "string", value)

    @property
    @khulnasoft.getter(name="stringMap")
    def string_map(self) -> khulnasoft.Input[Mapping[str, khulnasoft.Input[str]]]:
        return khulnasoft.get(self, "string_map")

    @string_map.setter
    def string_map(self, value: khulnasoft.Input[Mapping[str, khulnasoft.Input[str]]]):
        khulnasoft.set(self, "string_map", value)


