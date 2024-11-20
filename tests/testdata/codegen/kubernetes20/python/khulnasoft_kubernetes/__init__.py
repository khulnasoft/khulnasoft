# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .provider import *
from ._inputs import *

# Make subpackages available:
if typing.TYPE_CHECKING:
    import khulnasoft_kubernetes.core as __core
    core = __core
    import khulnasoft_kubernetes.helm as __helm
    helm = __helm
    import khulnasoft_kubernetes.meta as __meta
    meta = __meta
    import khulnasoft_kubernetes.yaml as __yaml
    yaml = __yaml
else:
    core = _utilities.lazy_import('khulnasoft_kubernetes.core')
    helm = _utilities.lazy_import('khulnasoft_kubernetes.helm')
    meta = _utilities.lazy_import('khulnasoft_kubernetes.meta')
    yaml = _utilities.lazy_import('khulnasoft_kubernetes.yaml')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "kubernetes",
  "mod": "core/v1",
  "fqn": "khulnasoft_kubernetes.core.v1",
  "classes": {
   "kubernetes:core/v1:ConfigMap": "ConfigMap",
   "kubernetes:core/v1:ConfigMapList": "ConfigMapList"
  }
 },
 {
  "pkg": "kubernetes",
  "mod": "helm.sh/v3",
  "fqn": "khulnasoft_kubernetes.helm.v3",
  "classes": {
   "kubernetes:helm.sh/v3:Release": "Release"
  }
 },
 {
  "pkg": "kubernetes",
  "mod": "yaml/v2",
  "fqn": "khulnasoft_kubernetes.yaml.v2",
  "classes": {
   "kubernetes:yaml/v2:ConfigGroup": "ConfigGroup"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "kubernetes",
  "token": "khulnasoft:providers:kubernetes",
  "fqn": "khulnasoft_kubernetes",
  "class": "Provider"
 }
]
"""
)
