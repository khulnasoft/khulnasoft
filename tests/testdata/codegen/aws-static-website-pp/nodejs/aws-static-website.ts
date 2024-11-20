import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as aws_static_website from "@khulnasoft/aws-static-website";

const websiteResource = new aws_static_website.Website("websiteResource", {
    sitePath: "string",
    indexHTML: "string",
    cacheTTL: 0,
    cdnArgs: {
        cloudfrontFunctionAssociations: [{
            eventType: "string",
            functionArn: "string",
        }],
        forwardedValues: {
            cookies: {
                forward: "string",
                whitelistedNames: ["string"],
            },
            queryString: false,
            headers: ["string"],
            queryStringCacheKeys: ["string"],
        },
        lambdaFunctionAssociations: [{
            eventType: "string",
            lambdaArn: "string",
            includeBody: false,
        }],
    },
    certificateARN: "string",
    error404: "string",
    addWebsiteVersionHeader: false,
    priceClass: "string",
    atomicDeployments: false,
    subdomain: "string",
    targetDomain: "string",
    withCDN: false,
    withLogs: false,
});
