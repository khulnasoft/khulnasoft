# Copyright 2016-2023, Pulumi Corporation.  All rights reserved.

import khulnasoft

from fails_on_delete import FailsOnDelete
from random_ import Random

rand = Random("random", length=10)
FailsOnDelete("failsondelete", opts=khulnasoft.ResourceOptions(deleted_with=rand))
