# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .configurer import *
from .provider import *
_utilities.register(
    resource_modules="""
[
 {
  "pkg": "metaprovider",
  "mod": "index",
  "fqn": "khulnasoft_metaprovider",
  "classes": {
   "metaprovider:index:Configurer": "Configurer"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "metaprovider",
  "token": "khulnasoft:providers:metaprovider",
  "fqn": "khulnasoft_metaprovider",
  "class": "Provider"
 }
]
"""
)
