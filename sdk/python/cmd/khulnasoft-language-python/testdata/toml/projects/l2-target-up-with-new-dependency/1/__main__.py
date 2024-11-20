import khulnasoft
import khulnasoft_simple as simple

target_only = simple.Resource("targetOnly", value=True)
dep = simple.Resource("dep", value=True)
unrelated = simple.Resource("unrelated", value=True,
opts = khulnasoft.ResourceOptions(depends_on=[dep]))
