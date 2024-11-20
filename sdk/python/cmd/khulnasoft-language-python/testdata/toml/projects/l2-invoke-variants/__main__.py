import khulnasoft
import khulnasoft_simple_invoke as simple_invoke

res = simple_invoke.StringResource("res")
khulnasoft.export("outputInput", simple_invoke.my_invoke_output(value=res.text).apply(lambda invoke: invoke.result))
khulnasoft.export("unit", simple_invoke.unit_output().apply(lambda invoke: invoke.result))
