// *** WARNING: this file was generated by khulnasoft-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as utilities from "./utilities";

export function myInvoke(args: MyInvokeArgs, opts?: khulnasoft.InvokeOptions): Promise<MyInvokeResult> {
    opts = khulnasoft.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return khulnasoft.runtime.invoke("simple-invoke:index:myInvoke", {
        "value": args.value,
    }, opts);
}

export interface MyInvokeArgs {
    value: string;
}

export interface MyInvokeResult {
    readonly result: string;
}
export function myInvokeOutput(args: MyInvokeOutputArgs, opts?: khulnasoft.InvokeOptions): khulnasoft.Output<MyInvokeResult> {
    opts = khulnasoft.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return khulnasoft.runtime.invokeOutput("simple-invoke:index:myInvoke", {
        "value": args.value,
    }, opts);
}

export interface MyInvokeOutputArgs {
    value: khulnasoft.Input<string>;
}