import khulnasoft
import khulnasoft_large as large

res = large.String("res", value="hello world")
khulnasoft.export("output", res.value)
