// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";

// Step 3: Run a query during `khulnasoft query`.
khulnasoft.runtime
    .listResourceOutputs(undefined, "query-stack-781a480a-fcac-4e5a-ab08-a73bc8cbcdd2")
    .groupBy<string, khulnasoft.Resource>(r => (<any>r).__khulnasoftType)
    .all(async function(group) {
        const count = await group.count();
        if (group.key === "khulnasoft-nodejs:dynamic:Resource" && count !== 2) {
            throw Error(`Expected 2 registered resources, got ${count}`);
        }
        console.log(group.key);
        return (
            group.key === "khulnasoft-nodejs:dynamic:Resource" ||
            group.key === "khulnasoft:providers:khulnasoft-nodejs" ||
            group.key === "khulnasoft:khulnasoft:Stack"
        );
    })
    .then(res => {
        if (res !== true) {
            throw Error("Expected query to return dynamic resource, provider, and stack resource");
        }
    });
