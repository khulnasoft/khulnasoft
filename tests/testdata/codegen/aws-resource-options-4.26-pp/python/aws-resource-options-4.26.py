import khulnasoft
import khulnasoft_aws as aws

provider = aws.Provider("provider", region="us-west-2")
bucket1 = aws.s3.Bucket("bucket1", opts = khulnasoft.ResourceOptions(provider=provider,
    depends_on=[provider],
    protect=True,
    ignore_changes=[
        "bucket",
        "lifecycleRules[0]",
    ]))
