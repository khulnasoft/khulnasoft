import khulnasoft
import khulnasoft_simple_invoke as simple_invoke

khulnasoft.export("hello", simple_invoke.my_invoke_output(value="hello").apply(lambda invoke: invoke.result))
khulnasoft.export("goodbye", simple_invoke.my_invoke_output(value="goodbye").apply(lambda invoke: invoke.result))
