import khulnasoft
from khulnasoft import Input
from typing import Optional, Dict, TypedDict, Any
import khulnasoft_random as random

class SimpleComponent(khulnasoft.ComponentResource):
    def __init__(self, name: str, opts: Optional[khulnasoft.ResourceOptions] = None):
        super().__init__("components:index:SimpleComponent", name, {}, opts)

        first_password = random.RandomPassword(f"{name}-firstPassword",
            length=16,
            special=True,
            opts = khulnasoft.ResourceOptions(parent=self))

        second_password = random.RandomPassword(f"{name}-secondPassword",
            length=16,
            special=True,
            opts = khulnasoft.ResourceOptions(parent=self))

        self.register_outputs()
