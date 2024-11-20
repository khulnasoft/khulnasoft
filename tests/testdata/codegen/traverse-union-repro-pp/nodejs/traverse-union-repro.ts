import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as aws from "@khulnasoft/aws";

const test = new aws.fsx.OpenZfsFileSystem("test", {
    storageCapacity: 64,
    subnetIds: [aws_subnet.test1.id],
    deploymentType: "SINGLE_AZ_1",
    throughputCapacity: 64,
});
