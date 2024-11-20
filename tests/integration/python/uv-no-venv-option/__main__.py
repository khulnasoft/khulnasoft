# Copyright 2024, Pulumi Corporation.  All rights reserved.

"""A Python project that uses the uv toolchain, without specifing the location for the virtualenv."""

import khulnasoft

khulnasoft.export("foo", "bar")
