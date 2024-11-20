# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import khulnasoft

config = khulnasoft.Config('config_missing_py')
config.require_secret('notFound')