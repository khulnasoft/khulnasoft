# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import khulnasoft
from component import Component

component_a = Component("a", id="hello")

khulnasoft.export("id", component_a.id)