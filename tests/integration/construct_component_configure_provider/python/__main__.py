# Copyright 2016-2023, Pulumi Corporation.  All rights reserved.

import khulnasoft
import khulnasoft_metaprovider
import khulnasoft_tls
import helpers

config = khulnasoft.Config()
proxy = config.require('proxy')

configurer = khulnasoft_metaprovider.Configurer("configurer", tls_proxy=helpers.unknownIfDryRun(proxy))
tls_provider = configurer.tls_provider()

key = khulnasoft_tls.PrivateKey("my-private-key",
                            algorithm="ECDSA",
                            ecdsa_curve="P384",
                            opts=khulnasoft.ResourceOptions(provider=tls_provider))

khulnasoft.export("keyAlgo", key.algorithm)
khulnasoft.export("meaningOfLife", configurer.meaning_of_life() + 1 - 1)

mix = configurer.object_mix()
key2 = khulnasoft_tls.PrivateKey("my-private-key-2",
                             algorithm="ECDSA",
                             ecdsa_curve="P384",
                             opts=khulnasoft.ResourceOptions(provider=mix.provider))

khulnasoft.export("keyAlgo2", key2.algorithm)
khulnasoft.export("meaningOfLife2", mix.meaning_of_life + 1 - 1)
