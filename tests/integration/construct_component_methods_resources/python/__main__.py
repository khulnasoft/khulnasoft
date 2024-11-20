# Copyright 2016-2021, Pulumi Corporation.  All rights reserved.

import khulnasoft

from component import Component

component = Component("component")
result = component.create_random(length=10).result

khulnasoft.export("result", result)
