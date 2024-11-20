import khulnasoft

ref = khulnasoft.StackReference("ref")
ref.outputs["bucket"].apply(lambda bucket: khulnasoft.log.info(f"Bucket: {bucket}"))
