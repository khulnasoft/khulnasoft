import khulnasoft
import khulnasoft_unknown as unknown

data = unknown.index.get_data(input="hello")
values = unknown.eks.module_values()
khulnasoft.export("content", data["content"])
