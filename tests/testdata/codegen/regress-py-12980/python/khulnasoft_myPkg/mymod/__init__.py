# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from .. import _utilities
import typing

# Make subpackages available:
if typing.TYPE_CHECKING:
    import khulnasoft_myPkg.mymod.childa as __childa
    childa = __childa
    import khulnasoft_myPkg.mymod.childb as __childb
    childb = __childb
else:
    childa = _utilities.lazy_import('khulnasoft_myPkg.mymod.childa')
    childb = _utilities.lazy_import('khulnasoft_myPkg.mymod.childb')
