# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import khulnasoft
import khulnasoft.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities
from . import outputs

__all__ = [
    'ListStorageAccountKeysResult',
    'AwaitableListStorageAccountKeysResult',
    'list_storage_account_keys',
    'list_storage_account_keys_output',
]

@khulnasoft.output_type
class ListStorageAccountKeysResult:
    """
    The response from the ListKeys operation.
    """
    def __init__(__self__, keys=None):
        if keys and not isinstance(keys, list):
            raise TypeError("Expected argument 'keys' to be a list")
        khulnasoft.set(__self__, "keys", keys)

    @property
    @khulnasoft.getter
    def keys(self) -> Sequence['outputs.StorageAccountKeyResponseResult']:
        """
        Gets the list of storage account keys and their properties for the specified storage account.
        """
        return khulnasoft.get(self, "keys")


class AwaitableListStorageAccountKeysResult(ListStorageAccountKeysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListStorageAccountKeysResult(
            keys=self.keys)


def list_storage_account_keys(account_name: Optional[str] = None,
                              expand: Optional[str] = None,
                              resource_group_name: Optional[str] = None,
                              opts: Optional[khulnasoft.InvokeOptions] = None) -> AwaitableListStorageAccountKeysResult:
    """
    The response from the ListKeys operation.
    API Version: 2021-02-01.


    :param str account_name: The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
    :param str expand: Specifies type of the key to be listed. Possible value is kerb.
    :param str resource_group_name: The name of the resource group within the user's subscription. The name is case insensitive.
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['expand'] = expand
    __args__['resourceGroupName'] = resource_group_name
    opts = khulnasoft.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = khulnasoft.runtime.invoke('mypkg::listStorageAccountKeys', __args__, opts=opts, typ=ListStorageAccountKeysResult).value

    return AwaitableListStorageAccountKeysResult(
        keys=khulnasoft.get(__ret__, 'keys'))
def list_storage_account_keys_output(account_name: Optional[khulnasoft.Input[str]] = None,
                                     expand: Optional[khulnasoft.Input[Optional[str]]] = None,
                                     resource_group_name: Optional[khulnasoft.Input[str]] = None,
                                     opts: Optional[khulnasoft.InvokeOptions] = None) -> khulnasoft.Output[ListStorageAccountKeysResult]:
    """
    The response from the ListKeys operation.
    API Version: 2021-02-01.


    :param str account_name: The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
    :param str expand: Specifies type of the key to be listed. Possible value is kerb.
    :param str resource_group_name: The name of the resource group within the user's subscription. The name is case insensitive.
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['expand'] = expand
    __args__['resourceGroupName'] = resource_group_name
    opts = khulnasoft.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = khulnasoft.runtime.invoke_output('mypkg::listStorageAccountKeys', __args__, opts=opts, typ=ListStorageAccountKeysResult)
    return __ret__.apply(lambda __response__: ListStorageAccountKeysResult(
        keys=khulnasoft.get(__response__, 'keys')))
