import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as aws from "@khulnasoft/aws";

const logs = new aws.s3.Bucket("logs", {});
const bucket = new aws.s3.Bucket("bucket", {loggings: [{
    targetBucket: logs.bucket,
}]});
export const targetBucket = bucket.loggings.apply(loggings => loggings?.[0]?.targetBucket);
