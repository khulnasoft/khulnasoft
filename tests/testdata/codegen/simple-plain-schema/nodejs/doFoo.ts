// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function doFoo(args: DoFooArgs, opts?: khulnasoft.InvokeOptions): Promise<void> {
    opts = khulnasoft.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return khulnasoft.runtime.invoke("example::doFoo", {
        "foo": args.foo,
    }, opts);
}

export interface DoFooArgs {
    foo: inputs.Foo;
}