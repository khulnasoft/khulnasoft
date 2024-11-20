import khulnasoft
import khulnasoft_simple_invoke as simple_invoke

explicit_provider = simple_invoke.Provider("explicitProvider")
data = simple_invoke.my_invoke_output(value="hello", opts=khulnasoft.InvokeOptions(provider=explicit_provider, parent=explicit_provider, version="10.0.0", plugin_download_url="https://example.com/github/example"))
khulnasoft.export("hello", data.result)
