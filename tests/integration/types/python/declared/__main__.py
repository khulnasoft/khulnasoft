# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

from typing import Optional

import khulnasoft
from khulnasoft.dynamic import Resource, ResourceProvider, CreateResult


@khulnasoft.input_type
class AdditionalArgs:
    def __init__(self, first_value: khulnasoft.Input[str], second_value: Optional[khulnasoft.Input[float]] = None):
        khulnasoft.set(self, "first_value", first_value)
        khulnasoft.set(self, "second_value", second_value)

    # Property with empty getter/setter bodies.
    @property
    @khulnasoft.getter(name="firstValue")
    def first_value(self) -> khulnasoft.Input[str]:
        ...

    @first_value.setter
    def first_value(self, value: khulnasoft.Input[str]):
        ...

    # Property with explicitly specified getter/setter bodies.
    @property
    @khulnasoft.getter(name="secondValue")
    def second_value(self) -> Optional[khulnasoft.Input[float]]:
        return khulnasoft.get(self, "second_value")

    @second_value.setter
    def second_value(self, value: Optional[khulnasoft.Input[float]]):
        khulnasoft.set(self, "second_value", value)

@khulnasoft.output_type
class Additional(dict):
    def __init__(self, first_value: str, second_value: Optional[float]):
        khulnasoft.set(self, "first_value", first_value)
        khulnasoft.set(self, "second_value", second_value)

    # Property with empty getter body.
    @property
    @khulnasoft.getter(name="firstValue")
    def first_value(self) -> str:
        ...

    # Property with explicitly specified getter/setter bodies.
    @property
    @khulnasoft.getter(name="secondValue")
    def second_value(self) -> Optional[float]:
        return khulnasoft.get(self, "second_value")

current_id = 0

class MyResourceProvider(ResourceProvider):
    def create(self, inputs):
        global current_id
        current_id += 1
        return CreateResult(str(current_id), {"additional": inputs["additional"]})

class MyResource(Resource):
    additional: khulnasoft.Output[Additional]

    def __init__(self, name: str, additional: khulnasoft.InputType[AdditionalArgs]):
        super().__init__(MyResourceProvider(), name, {"additional": additional})


# Create a resource with input object.
res = MyResource("testres", additional=AdditionalArgs(first_value="hello", second_value=42))

# Create a resource using the output object of another resource.
res2 = MyResource("testres2", additional=AdditionalArgs(
    first_value=res.additional.first_value,
    second_value=res.additional.second_value))

# Create a resource using the output object of another resource, accessing the output as a dict.
res3 = MyResource("testres3", additional=AdditionalArgs(
    first_value=res.additional["first_value"],
    second_value=res.additional["second_value"]))

# Create a resource using a dict as the input.
# Note: These are camel case (not snake_case) since the resource does not do any translation of
# property names.
res4 = MyResource("testres4", additional={
    "firstValue": "hello",
    "secondValue": 42,
})

khulnasoft.export("res_first_value", res.additional.first_value)
khulnasoft.export("res_second_value", res.additional.second_value)
khulnasoft.export("res2_first_value", res2.additional.first_value)
khulnasoft.export("res2_second_value", res2.additional.second_value)
khulnasoft.export("res3_first_value", res3.additional.first_value)
khulnasoft.export("res3_second_value", res3.additional.second_value)
khulnasoft.export("res4_first_value", res4.additional.first_value)
khulnasoft.export("res4_second_value", res4.additional.second_value)
