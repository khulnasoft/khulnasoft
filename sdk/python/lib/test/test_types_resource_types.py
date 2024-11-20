# Copyright 2016-2020, Pulumi Corporation.
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

from khulnasoft._types import resource_types
import khulnasoft


class Resource1(khulnasoft.Resource):
    pass


class Resource2(khulnasoft.Resource):
    foo: khulnasoft.Output[str]


class Resource3(khulnasoft.Resource):
    nested: khulnasoft.Output["Nested"]


class Resource4(khulnasoft.Resource):
    nested_value: khulnasoft.Output["Nested"] = khulnasoft.property("nestedValue")


class Resource5(khulnasoft.Resource):
    @property
    @khulnasoft.getter
    def foo(self) -> khulnasoft.Output[str]: ...  # type: ignore


class Resource6(khulnasoft.Resource):
    @property
    @khulnasoft.getter
    def nested(self) -> khulnasoft.Output["Nested"]: ...  # type: ignore


class Resource7(khulnasoft.Resource):
    @property
    @khulnasoft.getter(name="nestedValue")
    def nested_value(self) -> khulnasoft.Output["Nested"]: ...  # type: ignore


class Resource8(khulnasoft.Resource):
    foo: khulnasoft.Output


class Resource9(khulnasoft.Resource):
    @property
    @khulnasoft.getter
    def foo(self) -> khulnasoft.Output: ...  # type: ignore


class Resource10(khulnasoft.Resource):
    foo: str


class Resource11(khulnasoft.Resource):
    @property
    @khulnasoft.getter
    def foo(self) -> str: ...  # type: ignore


class Resource12(khulnasoft.Resource):
    @property
    @khulnasoft.getter
    def foo(self): ...  # type: ignore


@khulnasoft.output_type
class Nested:
    first: str
    second: str


class ResourceTypesTests(unittest.TestCase):
    def test_resource_types(self):
        self.assertEqual({}, resource_types(Resource1))

        self.assertEqual({"foo": str}, resource_types(Resource2))
        self.assertEqual({"nested": Nested}, resource_types(Resource3))
        self.assertEqual({"nestedValue": Nested}, resource_types(Resource4))

        self.assertEqual({"foo": str}, resource_types(Resource5))
        self.assertEqual({"nested": Nested}, resource_types(Resource6))
        self.assertEqual({"nestedValue": Nested}, resource_types(Resource7))

        # Non-generic Output excluded from types.
        self.assertEqual({}, resource_types(Resource8))
        self.assertEqual({}, resource_types(Resource9))

        # Type annotations not using Output.
        self.assertEqual({"foo": str}, resource_types(Resource10))
        self.assertEqual({"foo": str}, resource_types(Resource11))

        # No return type annotation from the property getter.
        self.assertEqual({}, resource_types(Resource12))
