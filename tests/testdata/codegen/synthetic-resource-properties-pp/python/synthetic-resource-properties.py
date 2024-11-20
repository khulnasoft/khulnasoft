import khulnasoft
import khulnasoft_synthetic as synthetic

rt = synthetic.resource_properties.Root("rt")
khulnasoft.export("trivial", rt)
khulnasoft.export("simple", rt.res1)
khulnasoft.export("foo", rt.res1.obj1.res2.obj2)
khulnasoft.export("complex", rt.res1.obj1.res2.obj2.answer)
