import khulnasoft
import khulnasoft_simple as simple

prov = simple.Provider("prov")
res = simple.Resource("res", value=True,
opts = khulnasoft.ResourceOptions(provider=prov))
