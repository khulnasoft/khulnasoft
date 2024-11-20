# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .foo import *
from .provider import *

# Make subpackages available:
if typing.TYPE_CHECKING:
    import khulnasoft_example.nested as __nested
    nested = __nested
else:
    nested = _utilities.lazy_import('khulnasoft_example.nested')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "example",
  "mod": "",
  "fqn": "khulnasoft_example",
  "classes": {
   "example::Foo": "Foo"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "example",
  "token": "khulnasoft:providers:example",
  "fqn": "khulnasoft_example",
  "class": "Provider"
 }
]
"""
)