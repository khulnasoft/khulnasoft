# coding=utf-8
# *** WARNING: this file was generated by khulnasoft-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .provider import *
from .resource import *
_utilities.register(
    resource_modules="""
[
 {
  "pkg": "old",
  "mod": "index",
  "fqn": "khulnasoft_old",
  "classes": {
   "old:index:Resource": "Resource"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "old",
  "token": "khulnasoft:providers:old",
  "fqn": "khulnasoft_old",
  "class": "Provider"
 }
]
"""
)
