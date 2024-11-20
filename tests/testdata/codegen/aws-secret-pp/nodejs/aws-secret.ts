import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as aws from "@khulnasoft/aws";

const dbCluster = new aws.rds.Cluster("dbCluster", {masterPassword: khulnasoft.secret("foobar")});
