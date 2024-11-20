# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .provider import *

# Make subpackages available:
if typing.TYPE_CHECKING:
    import khulnasoft_example.mymodule as __mymodule
    mymodule = __mymodule
else:
    mymodule = _utilities.lazy_import('khulnasoft_example.mymodule')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "example",
  "mod": "myModule",
  "fqn": "khulnasoft_example.mymodule",
  "classes": {
   "example:myModule:IamResource": "IamResource"
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
