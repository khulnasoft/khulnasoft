import khulnasoft
import khulnasoft_unknown as unknown

provider = khulnasoft.providers.Unknown("provider")
main = unknown.index.Main("main",
    first=hello,
    second={
        foo: bar,
    })
from_module = []
for range in [{"value": i} for i in range(0, 10)]:
    from_module.append(unknown.eks.Example(f"fromModule-{range['value']}", associated_main=main.id))
khulnasoft.export("mainId", main["id"])
khulnasoft.export("values", from_module["values"]["first"])
