# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from .. import _utilities
import typing

# Make subpackages available:
if typing.TYPE_CHECKING:
    import khulnasoft_example.gcp.gke as __gke
    gke = __gke
else:
    gke = _utilities.lazy_import('khulnasoft_example.gcp.gke')

