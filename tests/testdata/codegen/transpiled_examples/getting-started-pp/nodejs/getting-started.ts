import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as aws from "@khulnasoft/aws";

const mybucket = new aws.s3.Bucket("mybucket", {website: {
    indexDocument: "index.html",
}});
const indexhtml = new aws.s3.BucketObject("indexhtml", {
    bucket: mybucket.id,
    source: new khulnasoft.asset.StringAsset("<h1>Hello, world!</h1>"),
    acl: "public-read",
    contentType: "text/html",
});
export const bucketEndpoint = khulnasoft.interpolate`http://${mybucket.websiteEndpoint}`;
