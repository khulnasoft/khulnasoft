# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .component import *
from .my_function import *
from .provider import *
from ._inputs import *
from . import outputs
_utilities.register(
    resource_modules="""
[
 {
  "pkg": "typedDictExample",
  "mod": "index",
  "fqn": "khulnasoft_typedDictExample",
  "classes": {
   "typedDictExample:index:Component": "Component"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "typedDictExample",
  "token": "khulnasoft:providers:typedDictExample",
  "fqn": "khulnasoft_typedDictExample",
  "class": "Provider"
 }
]
"""
)