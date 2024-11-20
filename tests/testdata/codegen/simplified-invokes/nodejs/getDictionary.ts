// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getDictionary(a?: number, opts?: khulnasoft.InvokeOptions): Promise<{[key: string]: outputs.AnotherCustomResult}> {
    opts = khulnasoft.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return khulnasoft.runtime.invoke("std:index:GetDictionary", {
        "a": a,
    }, opts);
}
export function getDictionaryOutput(a?: khulnasoft.Input<number | undefined>, opts?: khulnasoft.InvokeOptions): khulnasoft.Output<{[key: string]: outputs.AnotherCustomResult}> {
    opts = khulnasoft.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return khulnasoft.runtime.invokeOutput("std:index:GetDictionary", {
        "a": a,
    }, opts);
}