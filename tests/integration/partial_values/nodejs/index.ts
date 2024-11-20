// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as assert from "assert";
import * as khulnasoft from "@khulnasoft/khulnasoft";
import { Resource } from "./resource";

const unknown = <any>khulnasoft.output(khulnasoft.runtime.isDryRun() ? { __khulnasoftUnknown: true } : "foo");

let a = new Resource("res", {
    foo: "foo",
    bar: { value: "foo", unknown },
    baz: [ "foo", unknown ],
});

export let o = Promise.all([
    (<any>a.foo).isKnown,
    (<any>a.bar.value).isKnown,
    (<any>a.bar.unknown).isKnown,
    (<any>a.baz[0]).isKnown,
    (<any>a.baz[1]).isKnown,
]).then(([r1, r2, r3, r4, r5]) => {
    assert.equal(r1, true);
    assert.equal(r2, true);
    assert.equal(r3, !khulnasoft.runtime.isDryRun());
    assert.equal(r4, true);
    assert.equal(r5, !khulnasoft.runtime.isDryRun());

    console.log("ok");
    return "checked";
});
