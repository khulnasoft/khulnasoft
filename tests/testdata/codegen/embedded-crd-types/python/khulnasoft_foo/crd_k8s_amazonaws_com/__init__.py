# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from .. import _utilities
import typing

# Make subpackages available:
if typing.TYPE_CHECKING:
    import khulnasoft_foo.crd_k8s_amazonaws_com.v1alpha1 as __v1alpha1
    v1alpha1 = __v1alpha1
else:
    v1alpha1 = _utilities.lazy_import('khulnasoft_foo.crd_k8s_amazonaws_com.v1alpha1')

