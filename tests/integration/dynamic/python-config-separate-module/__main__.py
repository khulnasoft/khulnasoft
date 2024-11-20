# Copyright 2024, Pulumi Corporation.  All rights reserved.

import khulnasoft
from khulnasoft.dynamic import Resource

from provider import SimpleProvider


class SimpleResource(Resource):
    authenticated: khulnasoft.Output[str]
    color: khulnasoft.Output[str]

    def __init__(self, name):
        super().__init__(SimpleProvider(), name, {"authenticated": None, "color": None})


r = SimpleResource("foo")
khulnasoft.export("authenticated", r.authenticated)
khulnasoft.export("color", r.color)
