# Copyright 2016-2021, Pulumi Corporation.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import unittest

from typing import Optional

from khulnasoft._types import input_type_types
import khulnasoft


@khulnasoft.input_type
class Foo:
    @property
    @khulnasoft.getter()
    def bar(self) -> khulnasoft.Input[str]: ...  # type: ignore


@khulnasoft.input_type
class MySimpleInputType:
    a: str
    b: Optional[str]
    c: khulnasoft.Input[str]
    d: Optional[khulnasoft.Input[str]]
    e: Foo
    f: Optional[Foo]
    g: khulnasoft.Input[Foo]
    h: Optional[khulnasoft.Input[Foo]]
    i: khulnasoft.InputType[Foo]
    j: Optional[khulnasoft.InputType[Foo]]
    k: khulnasoft.Input[khulnasoft.InputType[Foo]]
    l: Optional[khulnasoft.Input[khulnasoft.InputType[Foo]]]


@khulnasoft.input_type
class MyPropertiesInputType:
    @property
    @khulnasoft.getter()
    def a(self) -> str: ...  # type: ignore

    @property
    @khulnasoft.getter()
    def b(self) -> Optional[str]: ...  # type: ignore

    @property
    @khulnasoft.getter()
    def c(self) -> khulnasoft.Input[str]: ...  # type: ignore

    @property
    @khulnasoft.getter()
    def d(self) -> Optional[khulnasoft.Input[str]]: ...  # type: ignore

    @property
    @khulnasoft.getter()
    def e(self) -> Foo: ...  # type: ignore

    @property
    @khulnasoft.getter()
    def f(self) -> Optional[Foo]: ...  # type: ignore

    @property
    @khulnasoft.getter()
    def g(self) -> khulnasoft.Input[Foo]: ...  # type: ignore

    @property
    @khulnasoft.getter()
    def h(self) -> Optional[khulnasoft.Input[Foo]]: ...  # type: ignore

    @property
    @khulnasoft.getter()
    def i(self) -> khulnasoft.InputType[Foo]: ...  # type: ignore

    @property
    @khulnasoft.getter()
    def j(self) -> Optional[khulnasoft.InputType[Foo]]: ...  # type: ignore

    @property
    @khulnasoft.getter()
    def k(self) -> khulnasoft.Input[khulnasoft.InputType[Foo]]: ...  # type: ignore

    @property
    @khulnasoft.getter()
    def l(self) -> Optional[khulnasoft.Input[khulnasoft.InputType[Foo]]]: ...  # type: ignore


class InputTypeTypesTests(unittest.TestCase):
    def test_input_type_types(self):
        expected = {
            "a": str,
            "b": str,
            "c": str,
            "d": str,
            "e": Foo,
            "f": Foo,
            "g": Foo,
            "h": Foo,
            "i": Foo,
            "j": Foo,
            "k": Foo,
            "l": Foo,
        }
        self.assertEqual(expected, input_type_types(MySimpleInputType))
        self.assertEqual(expected, input_type_types(MyPropertiesInputType))
