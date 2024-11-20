# coding=utf-8
# *** WARNING: this file was generated by khulnasoft-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import khulnasoft
import khulnasoft.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'DoEchoResult',
    'AwaitableDoEchoResult',
    'do_echo',
    'do_echo_output',
]

@khulnasoft.output_type
class DoEchoResult:
    def __init__(__self__, echo=None):
        if echo and not isinstance(echo, str):
            raise TypeError("Expected argument 'echo' to be a str")
        khulnasoft.set(__self__, "echo", echo)

    @property
    @khulnasoft.getter
    def echo(self) -> Optional[str]:
        return khulnasoft.get(self, "echo")


class AwaitableDoEchoResult(DoEchoResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return DoEchoResult(
            echo=self.echo)


def do_echo(echo: Optional[str] = None,
            opts: Optional[khulnasoft.InvokeOptions] = None) -> AwaitableDoEchoResult:
    """
    A test invoke that echoes its input.
    """
    __args__ = dict()
    __args__['echo'] = echo
    opts = khulnasoft.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = khulnasoft.runtime.invoke('pkg:index:doEcho', __args__, opts=opts, typ=DoEchoResult, package_ref=_utilities.get_package()).value

    return AwaitableDoEchoResult(
        echo=khulnasoft.get(__ret__, 'echo'))


@_utilities.lift_output_func(do_echo)
def do_echo_output(echo: Optional[khulnasoft.Input[Optional[str]]] = None,
                   opts: Optional[khulnasoft.InvokeOptions] = None) -> khulnasoft.Output[DoEchoResult]:
    """
    A test invoke that echoes its input.
    """
    ...