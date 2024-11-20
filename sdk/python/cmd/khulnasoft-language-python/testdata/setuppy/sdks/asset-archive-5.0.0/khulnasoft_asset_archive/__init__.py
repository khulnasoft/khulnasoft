# coding=utf-8
# *** WARNING: this file was generated by khulnasoft-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .archive_resource import *
from .asset_resource import *
from .provider import *
_utilities.register(
    resource_modules="""
[
 {
  "pkg": "asset-archive",
  "mod": "index",
  "fqn": "khulnasoft_asset_archive",
  "classes": {
   "asset-archive:index:ArchiveResource": "ArchiveResource",
   "asset-archive:index:AssetResource": "AssetResource"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "asset-archive",
  "token": "khulnasoft:providers:asset-archive",
  "fqn": "khulnasoft_asset_archive",
  "class": "Provider"
 }
]
"""
)
