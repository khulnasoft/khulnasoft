# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from ._enums import *
from .list_configurations import *
from .list_product_families import *
from .provider import *
from ._inputs import *
from . import outputs
_utilities.register(
    resource_modules="""
[]
""",
    resource_packages="""
[
 {
  "pkg": "myedgeorder",
  "token": "khulnasoft:providers:myedgeorder",
  "fqn": "khulnasoft_myedgeorder",
  "class": "Provider"
 }
]
"""
)
