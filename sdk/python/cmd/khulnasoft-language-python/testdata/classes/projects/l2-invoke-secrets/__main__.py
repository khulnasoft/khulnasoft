import khulnasoft
import khulnasoft_simple as simple
import khulnasoft_simple_invoke as simple_invoke

res = simple.Resource("res", value=True)
khulnasoft.export("nonSecret", simple_invoke.secret_invoke_output(value="hello",
    secret_response=False).apply(lambda invoke: invoke.response))
khulnasoft.export("firstSecret", simple_invoke.secret_invoke_output(value="hello",
    secret_response=res.value).apply(lambda invoke: invoke.response))
khulnasoft.export("secondSecret", simple_invoke.secret_invoke_output(value=khulnasoft.Output.secret("goodbye"),
    secret_response=False).apply(lambda invoke: invoke.response))
