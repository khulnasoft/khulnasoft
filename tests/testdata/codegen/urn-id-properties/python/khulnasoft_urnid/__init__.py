# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .provider import *
from .res import *
from .test import *
from . import outputs
_utilities.register(
    resource_modules="""
[
 {
  "pkg": "urnid",
  "mod": "index",
  "fqn": "khulnasoft_urnid",
  "classes": {
   "urnid:index:Res": "Res"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "urnid",
  "token": "khulnasoft:providers:urnid",
  "fqn": "khulnasoft_urnid",
  "class": "Provider"
 }
]
"""
)