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
    'ENIConfigSpecArgs',
    'ENIConfigSpecArgsDict',
]

MYPY = False

if not MYPY:
    class ENIConfigSpecArgsDict(TypedDict):
        security_groups: NotRequired[khulnasoft.Input[Sequence[khulnasoft.Input[str]]]]
        subnet: NotRequired[khulnasoft.Input[str]]
elif False:
    ENIConfigSpecArgsDict: TypeAlias = Mapping[str, Any]

@khulnasoft.input_type
class ENIConfigSpecArgs:
    def __init__(__self__, *,
                 security_groups: Optional[khulnasoft.Input[Sequence[khulnasoft.Input[str]]]] = None,
                 subnet: Optional[khulnasoft.Input[str]] = None):
        if security_groups is not None:
            khulnasoft.set(__self__, "security_groups", security_groups)
        if subnet is not None:
            khulnasoft.set(__self__, "subnet", subnet)

    @property
    @khulnasoft.getter(name="securityGroups")
    def security_groups(self) -> Optional[khulnasoft.Input[Sequence[khulnasoft.Input[str]]]]:
        return khulnasoft.get(self, "security_groups")

    @security_groups.setter
    def security_groups(self, value: Optional[khulnasoft.Input[Sequence[khulnasoft.Input[str]]]]):
        khulnasoft.set(self, "security_groups", value)

    @property
    @khulnasoft.getter
    def subnet(self) -> Optional[khulnasoft.Input[str]]:
        return khulnasoft.get(self, "subnet")

    @subnet.setter
    def subnet(self, value: Optional[khulnasoft.Input[str]]):
        khulnasoft.set(self, "subnet", value)

