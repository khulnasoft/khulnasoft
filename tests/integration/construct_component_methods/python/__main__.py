# Copyright 2016-2021, Pulumi Corporation.  All rights reserved.

import khulnasoft

from component import Component

component = Component("component", first="Hello", second="World")
result = component.get_message("Alice")

khulnasoft.export("message", result.message)
