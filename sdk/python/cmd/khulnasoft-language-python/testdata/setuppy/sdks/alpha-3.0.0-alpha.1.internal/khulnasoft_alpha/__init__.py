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
  "pkg": "alpha",
  "mod": "index",
  "fqn": "khulnasoft_alpha",
  "classes": {
   "alpha:index:Resource": "Resource"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "alpha",
  "token": "khulnasoft:providers:alpha",
  "fqn": "khulnasoft_alpha",
  "class": "Provider"
 }
]
"""
)