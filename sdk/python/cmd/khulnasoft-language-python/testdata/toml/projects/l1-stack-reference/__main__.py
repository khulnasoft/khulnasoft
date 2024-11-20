import khulnasoft

ref = khulnasoft.StackReference("ref", stack_name="organization/other/dev")
khulnasoft.export("plain", ref.get_output("plain"))
khulnasoft.export("secret", ref.get_output("secret"))
