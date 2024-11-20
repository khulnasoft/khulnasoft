# Copyright 2016-2023, Pulumi Corporation.  All rights reserved.

import khulnasoft

from component import Component
from test_provider import TestProvider

test_provider = TestProvider("testProvider")

component1 = Component("component1", first="Hello", second="World",
    opts=khulnasoft.ResourceOptions(provider=test_provider))
result1 = component1.get_message("Alice")

component2 = Component("component2", first="Hi", second="There",
    opts=khulnasoft.ResourceOptions(providers=[test_provider]))
result2 = component2.get_message("Bob")


khulnasoft.export("message1", result1.message)
khulnasoft.export("message2", result2.message)
