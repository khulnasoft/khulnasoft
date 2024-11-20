// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Config } from "@khulnasoft/khulnasoft";

const config = new Config("config_missing_js");
config.requireSecret("notFound")