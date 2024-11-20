# Copyright 2024, Pulumi Corporation.

from khulnasoft import ComponentResource, ResourceOptions

from echo import Echo

for i in range(0, 100):
    Echo(f"echo-{i}", echo=f"hello-{i}")
