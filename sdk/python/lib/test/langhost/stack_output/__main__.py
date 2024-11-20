# Copyright 2016-2018, Pulumi Corporation.
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
from typing import Any, Dict
import khulnasoft


class TestClass:
    def __init__(self):
        self.num = 1
        self._private = 2


my_test_class_instance = TestClass()

recursive: Dict[str, Any] = {"a": 1}
recursive["b"] = 2
recursive["c"] = recursive

khulnasoft.export("string", "khulnasoft")
khulnasoft.export("number", 1)
khulnasoft.export("boolean", True)
khulnasoft.export("list", [])
khulnasoft.export("list_with_none", [None])
khulnasoft.export("list_of_lists", [[], []])
khulnasoft.export(
    "list_of_outputs", [[khulnasoft.Output.from_input(1)], khulnasoft.Output.from_input([2])]
)
khulnasoft.export("set", set(["val"]))
khulnasoft.export("dict", {"a": 1})
khulnasoft.export("output", khulnasoft.Output.from_input(1))
khulnasoft.export("class", TestClass())
khulnasoft.export("recursive", recursive)
khulnasoft.export("duplicate_output_0", my_test_class_instance)
khulnasoft.export("duplicate_output_1", my_test_class_instance)
