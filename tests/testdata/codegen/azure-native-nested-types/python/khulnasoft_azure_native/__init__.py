# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .provider import *

# Make subpackages available:
if typing.TYPE_CHECKING:
    import khulnasoft_azure_native.documentdb as __documentdb
    documentdb = __documentdb
else:
    documentdb = _utilities.lazy_import('khulnasoft_azure_native.documentdb')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "azure-native",
  "mod": "documentdb",
  "fqn": "khulnasoft_azure_native.documentdb",
  "classes": {
   "azure-native:documentdb:SqlResourceSqlContainer": "SqlResourceSqlContainer"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "azure-native",
  "token": "khulnasoft:providers:azure-native",
  "fqn": "khulnasoft_azure_native",
  "class": "Provider"
 }
]
"""
)