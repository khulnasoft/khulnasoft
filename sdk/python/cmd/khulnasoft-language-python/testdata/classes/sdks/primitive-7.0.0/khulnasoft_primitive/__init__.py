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
  "pkg": "primitive",
  "mod": "index",
  "fqn": "khulnasoft_primitive",
  "classes": {
   "primitive:index:Resource": "Resource"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "primitive",
  "token": "khulnasoft:providers:primitive",
  "fqn": "khulnasoft_primitive",
  "class": "Provider"
 }
]
"""
)
