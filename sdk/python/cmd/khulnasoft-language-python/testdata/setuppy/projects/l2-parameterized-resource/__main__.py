import khulnasoft
import khulnasoft_subpackage as subpackage

# The resource name is based on the parameter value
example = subpackage.HelloWorld("example")
khulnasoft.export("parameterValue", example.parameter_value)
