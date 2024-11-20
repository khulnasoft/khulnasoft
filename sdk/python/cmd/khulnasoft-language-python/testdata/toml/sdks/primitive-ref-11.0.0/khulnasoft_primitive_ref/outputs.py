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
    'Data',
]

@khulnasoft.output_type
class Data(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "boolArray":
            suggest = "bool_array"
        elif key == "stringMap":
            suggest = "string_map"

        if suggest:
            khulnasoft.log.warn(f"Key '{key}' not found in Data. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        Data.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        Data.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 bool_array: Sequence[bool],
                 boolean: bool,
                 float: float,
                 integer: int,
                 string: str,
                 string_map: Mapping[str, str]):
        khulnasoft.set(__self__, "bool_array", bool_array)
        khulnasoft.set(__self__, "boolean", boolean)
        khulnasoft.set(__self__, "float", float)
        khulnasoft.set(__self__, "integer", integer)
        khulnasoft.set(__self__, "string", string)
        khulnasoft.set(__self__, "string_map", string_map)

    @property
    @khulnasoft.getter(name="boolArray")
    def bool_array(self) -> Sequence[bool]:
        return khulnasoft.get(self, "bool_array")

    @property
    @khulnasoft.getter
    def boolean(self) -> bool:
        return khulnasoft.get(self, "boolean")

    @property
    @khulnasoft.getter
    def float(self) -> float:
        return khulnasoft.get(self, "float")

    @property
    @khulnasoft.getter
    def integer(self) -> int:
        return khulnasoft.get(self, "integer")

    @property
    @khulnasoft.getter
    def string(self) -> str:
        return khulnasoft.get(self, "string")

    @property
    @khulnasoft.getter(name="stringMap")
    def string_map(self) -> Mapping[str, str]:
        return khulnasoft.get(self, "string_map")


