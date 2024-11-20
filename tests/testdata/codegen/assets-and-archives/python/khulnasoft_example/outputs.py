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
    'TypeWithAssets',
]

@khulnasoft.output_type
class TypeWithAssets(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "plainArchive":
            suggest = "plain_archive"
        elif key == "plainAsset":
            suggest = "plain_asset"

        if suggest:
            khulnasoft.log.warn(f"Key '{key}' not found in TypeWithAssets. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TypeWithAssets.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TypeWithAssets.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 asset: Union[khulnasoft.Asset, khulnasoft.Archive],
                 plain_archive: khulnasoft.Archive,
                 archive: Optional[khulnasoft.Archive] = None,
                 plain_asset: Optional[Union[khulnasoft.Asset, khulnasoft.Archive]] = None):
        khulnasoft.set(__self__, "asset", asset)
        khulnasoft.set(__self__, "plain_archive", plain_archive)
        if archive is not None:
            khulnasoft.set(__self__, "archive", archive)
        if plain_asset is not None:
            khulnasoft.set(__self__, "plain_asset", plain_asset)

    @property
    @khulnasoft.getter
    def asset(self) -> Union[khulnasoft.Asset, khulnasoft.Archive]:
        return khulnasoft.get(self, "asset")

    @property
    @khulnasoft.getter(name="plainArchive")
    def plain_archive(self) -> khulnasoft.Archive:
        return khulnasoft.get(self, "plain_archive")

    @property
    @khulnasoft.getter
    def archive(self) -> Optional[khulnasoft.Archive]:
        return khulnasoft.get(self, "archive")

    @property
    @khulnasoft.getter(name="plainAsset")
    def plain_asset(self) -> Optional[Union[khulnasoft.Asset, khulnasoft.Archive]]:
        return khulnasoft.get(self, "plain_asset")


