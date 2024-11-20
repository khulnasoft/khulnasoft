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

__all__ = [
    'Pet',
]

@khulnasoft.output_type
class Pet(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None):
        if name is not None:
            khulnasoft.set(__self__, "name", name)

    @property
    @khulnasoft.getter
    def name(self) -> Optional[str]:
        return khulnasoft.get(self, "name")

