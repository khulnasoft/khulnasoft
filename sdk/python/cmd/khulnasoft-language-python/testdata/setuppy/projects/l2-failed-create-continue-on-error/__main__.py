import khulnasoft
import khulnasoft_fail_on_create as fail_on_create
import khulnasoft_simple as simple

failing = fail_on_create.Resource("failing", value=False)
dependent = simple.Resource("dependent", value=True,
opts = khulnasoft.ResourceOptions(depends_on=[failing]))
dependent_on_output = simple.Resource("dependent_on_output", value=failing.value)
independent = simple.Resource("independent", value=True)
double_dependency = simple.Resource("double_dependency", value=True,
opts = khulnasoft.ResourceOptions(depends_on=[
        independent,
        dependent_on_output,
    ]))
