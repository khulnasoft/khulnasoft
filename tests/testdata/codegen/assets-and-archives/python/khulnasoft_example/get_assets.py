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
    'GetAssetsResult',
    'AwaitableGetAssetsResult',
    'get_assets',
    'get_assets_output',
]

@khulnasoft.output_type
class GetAssetsResult:
    def __init__(__self__, archive=None, source=None):
        if archive and not isinstance(archive, khulnasoft.Archive):
            raise TypeError("Expected argument 'archive' to be a khulnasoft.Archive")
        khulnasoft.set(__self__, "archive", archive)
        if source and not isinstance(source, Union[khulnasoft.Asset, khulnasoft.Archive]):
            raise TypeError("Expected argument 'source' to be a Union[khulnasoft.Asset, khulnasoft.Archive]")
        khulnasoft.set(__self__, "source", source)

    @property
    @khulnasoft.getter
    def archive(self) -> khulnasoft.Archive:
        return khulnasoft.get(self, "archive")

    @property
    @khulnasoft.getter
    def source(self) -> Union[khulnasoft.Asset, khulnasoft.Archive]:
        return khulnasoft.get(self, "source")


class AwaitableGetAssetsResult(GetAssetsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAssetsResult(
            archive=self.archive,
            source=self.source)


def get_assets(archive: Optional[khulnasoft.Archive] = None,
               source: Optional[Union[khulnasoft.Asset, khulnasoft.Archive]] = None,
               opts: Optional[khulnasoft.InvokeOptions] = None) -> AwaitableGetAssetsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['archive'] = archive
    __args__['source'] = source
    opts = khulnasoft.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = khulnasoft.runtime.invoke('example::GetAssets', __args__, opts=opts, typ=GetAssetsResult).value

    return AwaitableGetAssetsResult(
        archive=khulnasoft.get(__ret__, 'archive'),
        source=khulnasoft.get(__ret__, 'source'))
def get_assets_output(archive: Optional[khulnasoft.Input[khulnasoft.Archive]] = None,
                      source: Optional[khulnasoft.Input[Union[khulnasoft.Asset, khulnasoft.Archive]]] = None,
                      opts: Optional[khulnasoft.InvokeOptions] = None) -> khulnasoft.Output[GetAssetsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['archive'] = archive
    __args__['source'] = source
    opts = khulnasoft.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = khulnasoft.runtime.invoke_output('example::GetAssets', __args__, opts=opts, typ=GetAssetsResult)
    return __ret__.apply(lambda __response__: GetAssetsResult(
        archive=khulnasoft.get(__response__, 'archive'),
        source=khulnasoft.get(__response__, 'source')))
