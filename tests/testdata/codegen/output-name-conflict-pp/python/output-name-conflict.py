import khulnasoft

config = khulnasoft.Config()
cidr_block = config.get("cidrBlock")
if cidr_block is None:
    cidr_block = "Test config variable"
khulnasoft.export("cidrBlock", cidr_block)
