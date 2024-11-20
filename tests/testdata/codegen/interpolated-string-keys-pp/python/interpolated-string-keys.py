import khulnasoft

config = khulnasoft.Config()
value = config.require("value")
tags = config.get_object("tags")
if tags is None:
    tags = {
        f"interpolated/{value}": "value",
    }
