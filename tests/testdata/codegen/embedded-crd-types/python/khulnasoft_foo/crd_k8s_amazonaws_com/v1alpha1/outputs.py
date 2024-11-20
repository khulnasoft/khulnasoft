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

__all__ = [
    'ENIConfigSpec',
]

@khulnasoft.output_type
class ENIConfigSpec(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "securityGroups":
            suggest = "security_groups"

        if suggest:
            khulnasoft.log.warn(f"Key '{key}' not found in ENIConfigSpec. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ENIConfigSpec.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ENIConfigSpec.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 security_groups: Optional[Sequence[str]] = None,
                 subnet: Optional[str] = None):
        if security_groups is not None:
            khulnasoft.set(__self__, "security_groups", security_groups)
        if subnet is not None:
            khulnasoft.set(__self__, "subnet", subnet)

    @property
    @khulnasoft.getter(name="securityGroups")
    def security_groups(self) -> Optional[Sequence[str]]:
        return khulnasoft.get(self, "security_groups")

    @property
    @khulnasoft.getter
    def subnet(self) -> Optional[str]:
        return khulnasoft.get(self, "subnet")

