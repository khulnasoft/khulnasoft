# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .func_with_all_optional_inputs import *
from .func_with_const_input import *
from .func_with_default_value import *
from .func_with_dict_param import *
from .func_with_empty_outputs import *
from .func_with_list_param import *
from .get_bastion_shareable_link import *
from .get_client_config import *
from .get_integration_runtime_object_metadatum import *
from .list_storage_account_keys import *
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
  "pkg": "mypkg",
  "token": "khulnasoft:providers:mypkg",
  "fqn": "khulnasoft_mypkg",
  "class": "Provider"
 }
]
"""
)
