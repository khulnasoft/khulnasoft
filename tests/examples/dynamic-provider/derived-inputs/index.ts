// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as dynamic from "@khulnasoft/khulnasoft/dynamic";

const sleep = require("sleep-promise");

class InputProvider implements dynamic.ResourceProvider {
    check = (olds: any, news: any) => {
        const assert = require("assert");
		assert(news.input);
		return Promise.resolve({ inputs: news });
	};
    diff = (id: khulnasoft.ID, olds: any, news: any) => Promise.resolve({});
    create = (inputs: any) => Promise.resolve({ id: "0" });
    update = (id: string, olds: any, news: any) => Promise.resolve({});
    delete = (id: khulnasoft.ID, props: any) => Promise.resolve();
}

class InputResource extends dynamic.Resource {
    constructor(name: string, input: khulnasoft.Input<string>) {
        super(new InputProvider(), name, { input: input }, undefined);
    }
}

(async () => {
    try {
        const a = new InputResource("a", "string");
		const b = new InputResource("b", a.urn);
    } catch (err) {
        console.error(err);
        process.exit(-1);
    }
})();
