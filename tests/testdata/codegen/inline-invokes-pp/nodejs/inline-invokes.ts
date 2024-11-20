import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as aws from "@khulnasoft/aws";

const webSecurityGroup = new aws.ec2.SecurityGroup("webSecurityGroup", {vpcId: aws.ec2.getVpc({
    "default": true,
}).then(invoke => invoke.id)});
