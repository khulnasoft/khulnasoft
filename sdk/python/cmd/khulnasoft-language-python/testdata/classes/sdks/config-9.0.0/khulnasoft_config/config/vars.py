# coding=utf-8
# *** WARNING: this file was generated by khulnasoft-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import khulnasoft
import khulnasoft.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

import types

__config__ = khulnasoft.Config('config')


class _ExportableConfig(types.ModuleType):
    @property
    def name(self) -> Optional[str]:
        return __config__.get('name')

    @property
    def plugin_download_url(self) -> Optional[str]:
        return __config__.get('pluginDownloadURL')

