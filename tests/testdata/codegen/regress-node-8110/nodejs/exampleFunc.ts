// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

export function exampleFunc(args?: ExampleFuncArgs, opts?: khulnasoft.InvokeOptions): Promise<void> {
    args = args || {};
    opts = khulnasoft.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return khulnasoft.runtime.invoke("my8110::exampleFunc", {
        "enums": args.enums,
    }, opts);
}

export interface ExampleFuncArgs {
    enums?: (string | enums.MyEnum)[];
}
