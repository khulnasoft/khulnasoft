import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as aws from "@khulnasoft/aws";

const siteBucket = new aws.s3.Bucket("siteBucket", {});
const testFileAsset = new aws.s3.BucketObject("testFileAsset", {
    bucket: siteBucket.id,
    source: new khulnasoft.asset.FileAsset("file.txt"),
});
const testStringAsset = new aws.s3.BucketObject("testStringAsset", {
    bucket: siteBucket.id,
    source: new khulnasoft.asset.StringAsset("<h1>File contents</h1>"),
});
const testRemoteAsset = new aws.s3.BucketObject("testRemoteAsset", {
    bucket: siteBucket.id,
    source: new khulnasoft.asset.RemoteAsset("https://khulnasoft.test"),
});
const testFileArchive = new aws.lambda.Function("testFileArchive", {
    role: siteBucket.arn,
    code: new khulnasoft.asset.FileArchive("file.tar.gz"),
});
const testRemoteArchive = new aws.lambda.Function("testRemoteArchive", {
    role: siteBucket.arn,
    code: new khulnasoft.asset.RemoteArchive("https://khulnasoft.test/foo.tar.gz"),
});
const testAssetArchive = new aws.lambda.Function("testAssetArchive", {
    role: siteBucket.arn,
    code: new khulnasoft.asset.AssetArchive({
        "file.txt": new khulnasoft.asset.FileAsset("file.txt"),
        "string.txt": new khulnasoft.asset.StringAsset("<h1>File contents</h1>"),
        "remote.txt": new khulnasoft.asset.RemoteAsset("https://khulnasoft.test"),
        "file.tar": new khulnasoft.asset.FileArchive("file.tar.gz"),
        "remote.tar": new khulnasoft.asset.RemoteArchive("https://khulnasoft.test/foo.tar.gz"),
        ".nestedDir": new khulnasoft.asset.AssetArchive({
            "file.txt": new khulnasoft.asset.FileAsset("file.txt"),
            "string.txt": new khulnasoft.asset.StringAsset("<h1>File contents</h1>"),
            "remote.txt": new khulnasoft.asset.RemoteAsset("https://khulnasoft.test"),
            "file.tar": new khulnasoft.asset.FileArchive("file.tar.gz"),
            "remote.tar": new khulnasoft.asset.RemoteArchive("https://khulnasoft.test/foo.tar.gz"),
        }),
    }),
});
