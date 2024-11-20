# Copyright 2016-2023, Pulumi Corporation.  All rights reserved.

import khulnasoft

a = khulnasoft.Output.from_input([1, 2])

khulnasoft.export("export1", a)
khulnasoft.export("export2", a)