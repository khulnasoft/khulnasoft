// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as utilities from "./utilities";

/**
 * Codegen demo with const inputs
 */
export function funcWithConstInput(args?: FuncWithConstInputArgs, opts?: khulnasoft.InvokeOptions): Promise<void> {
    args = args || {};
    opts = khulnasoft.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return khulnasoft.runtime.invoke("mypkg::funcWithConstInput", {
        "plainInput": args.plainInput,
    }, opts);
}

export interface FuncWithConstInputArgs {
    plainInput?: "fixed";
}