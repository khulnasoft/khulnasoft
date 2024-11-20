import khulnasoft
import khulnasoft_aws as aws

site_bucket = aws.s3.Bucket("siteBucket")
test_file_asset = aws.s3.BucketObject("testFileAsset",
    bucket=site_bucket.id,
    source=khulnasoft.FileAsset("file.txt"))
test_string_asset = aws.s3.BucketObject("testStringAsset",
    bucket=site_bucket.id,
    source=khulnasoft.StringAsset("<h1>File contents</h1>"))
test_remote_asset = aws.s3.BucketObject("testRemoteAsset",
    bucket=site_bucket.id,
    source=khulnasoft.RemoteAsset("https://khulnasoft.test"))
test_file_archive = aws.lambda_.Function("testFileArchive",
    role=site_bucket.arn,
    code=khulnasoft.FileArchive("file.tar.gz"))
test_remote_archive = aws.lambda_.Function("testRemoteArchive",
    role=site_bucket.arn,
    code=khulnasoft.RemoteArchive("https://khulnasoft.test/foo.tar.gz"))
test_asset_archive = aws.lambda_.Function("testAssetArchive",
    role=site_bucket.arn,
    code=khulnasoft.AssetArchive({
        "file.txt": khulnasoft.FileAsset("file.txt"),
        "string.txt": khulnasoft.StringAsset("<h1>File contents</h1>"),
        "remote.txt": khulnasoft.RemoteAsset("https://khulnasoft.test"),
        "file.tar": khulnasoft.FileArchive("file.tar.gz"),
        "remote.tar": khulnasoft.RemoteArchive("https://khulnasoft.test/foo.tar.gz"),
        ".nestedDir": khulnasoft.AssetArchive({
            "file.txt": khulnasoft.FileAsset("file.txt"),
            "string.txt": khulnasoft.StringAsset("<h1>File contents</h1>"),
            "remote.txt": khulnasoft.RemoteAsset("https://khulnasoft.test"),
            "file.tar": khulnasoft.FileArchive("file.tar.gz"),
            "remote.tar": khulnasoft.RemoteArchive("https://khulnasoft.test/foo.tar.gz"),
        }),
    }))
