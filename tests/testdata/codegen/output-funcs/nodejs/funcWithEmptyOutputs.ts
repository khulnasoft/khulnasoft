// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as utilities from "./utilities";

/**
 * n/a
 */
export function funcWithEmptyOutputs(args: FuncWithEmptyOutputsArgs, opts?: khulnasoft.InvokeOptions): Promise<void> {
    opts = khulnasoft.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return khulnasoft.runtime.invoke("mypkg::funcWithEmptyOutputs", {
        "name": args.name,
    }, opts);
}

export interface FuncWithEmptyOutputsArgs {
    /**
     * The Name of the FeatureGroup.
     */
    name: string;
}