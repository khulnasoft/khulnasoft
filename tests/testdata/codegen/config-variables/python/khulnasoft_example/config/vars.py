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

import types

__config__ = khulnasoft.Config('example')


class _ExportableConfig(types.ModuleType):
    @property
    def http_retry_wait_max(self) -> Optional[float]:
        return __config__.get_float('httpRetryWaitMax')
