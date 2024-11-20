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

__all__ = [
    'FuncWithSecretsResult',
    'AwaitableFuncWithSecretsResult',
    'func_with_secrets',
    'func_with_secrets_output',
]

@khulnasoft.output_type
class FuncWithSecretsResult:
    def __init__(__self__, ciphertext=None, crypto_key=None, id=None, plaintext=None):
        if ciphertext and not isinstance(ciphertext, str):
            raise TypeError("Expected argument 'ciphertext' to be a str")
        khulnasoft.set(__self__, "ciphertext", ciphertext)
        if crypto_key and not isinstance(crypto_key, str):
            raise TypeError("Expected argument 'crypto_key' to be a str")
        khulnasoft.set(__self__, "crypto_key", crypto_key)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        khulnasoft.set(__self__, "id", id)
        if plaintext and not isinstance(plaintext, str):
            raise TypeError("Expected argument 'plaintext' to be a str")
        khulnasoft.set(__self__, "plaintext", plaintext)

    @property
    @khulnasoft.getter
    def ciphertext(self) -> str:
        return khulnasoft.get(self, "ciphertext")

    @property
    @khulnasoft.getter(name="cryptoKey")
    def crypto_key(self) -> str:
        return khulnasoft.get(self, "crypto_key")

    @property
    @khulnasoft.getter
    def id(self) -> str:
        return khulnasoft.get(self, "id")

    @property
    @khulnasoft.getter
    def plaintext(self) -> str:
        return khulnasoft.get(self, "plaintext")


class AwaitableFuncWithSecretsResult(FuncWithSecretsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return FuncWithSecretsResult(
            ciphertext=self.ciphertext,
            crypto_key=self.crypto_key,
            id=self.id,
            plaintext=self.plaintext)


def func_with_secrets(crypto_key: Optional[str] = None,
                      plaintext: Optional[str] = None,
                      opts: Optional[khulnasoft.InvokeOptions] = None) -> AwaitableFuncWithSecretsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['cryptoKey'] = crypto_key
    __args__['plaintext'] = plaintext
    opts = khulnasoft.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = khulnasoft.runtime.invoke('mypkg::funcWithSecrets', __args__, opts=opts, typ=FuncWithSecretsResult).value

    return AwaitableFuncWithSecretsResult(
        ciphertext=khulnasoft.get(__ret__, 'ciphertext'),
        crypto_key=khulnasoft.get(__ret__, 'crypto_key'),
        id=khulnasoft.get(__ret__, 'id'),
        plaintext=khulnasoft.get(__ret__, 'plaintext'))
def func_with_secrets_output(crypto_key: Optional[khulnasoft.Input[str]] = None,
                             plaintext: Optional[khulnasoft.Input[str]] = None,
                             opts: Optional[khulnasoft.InvokeOptions] = None) -> khulnasoft.Output[FuncWithSecretsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['cryptoKey'] = crypto_key
    __args__['plaintext'] = plaintext
    opts = khulnasoft.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = khulnasoft.runtime.invoke_output('mypkg::funcWithSecrets', __args__, opts=opts, typ=FuncWithSecretsResult)
    return __ret__.apply(lambda __response__: FuncWithSecretsResult(
        ciphertext=khulnasoft.get(__response__, 'ciphertext'),
        crypto_key=khulnasoft.get(__response__, 'crypto_key'),
        id=khulnasoft.get(__response__, 'id'),
        plaintext=khulnasoft.get(__response__, 'plaintext')))