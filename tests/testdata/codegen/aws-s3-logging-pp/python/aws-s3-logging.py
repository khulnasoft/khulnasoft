import khulnasoft
import khulnasoft_aws as aws

logs = aws.s3.Bucket("logs")
bucket = aws.s3.Bucket("bucket", loggings=[{
    "target_bucket": logs.bucket,
}])
khulnasoft.export("targetBucket", bucket.loggings[0].target_bucket)
