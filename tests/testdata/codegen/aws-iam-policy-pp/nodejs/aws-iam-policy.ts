import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as aws from "@khulnasoft/aws";

// Create a policy with multiple Condition keys
const policy = new aws.iam.Policy("policy", {
    path: "/",
    description: "My test policy",
    policy: JSON.stringify({
        Version: "2012-10-17",
        Statement: [{
            Effect: "Allow",
            Action: "lambda:*",
            Resource: "arn:aws:lambda:*:*:function:*",
            Condition: {
                StringEquals: {
                    "aws:RequestTag/Team": [
                        "iamuser-admin",
                        "iamuser2-admin",
                    ],
                },
                "ForAllValues:StringEquals": {
                    "aws:TagKeys": ["Team"],
                },
            },
        }],
    }),
});
export const policyName = policy.name;
