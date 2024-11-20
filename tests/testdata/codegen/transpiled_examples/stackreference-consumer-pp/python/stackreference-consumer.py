import khulnasoft

stack_ref = khulnasoft.StackReference("stackRef", stack_name="PLACEHOLDER_ORG_NAME/stackreference-producer/PLACEHOLDER_STACK_NAME")
khulnasoft.export("referencedImageName", stack_ref.outputs["imageName"])
