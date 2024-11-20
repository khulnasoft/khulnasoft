// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Check codegen of functions with all optional inputs.
 */
export function funcWithAllOptionalInputs(args?: FuncWithAllOptionalInputsArgs, opts?: khulnasoft.InvokeOptions): Promise<FuncWithAllOptionalInputsResult> {
    args = args || {};
    opts = khulnasoft.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return khulnasoft.runtime.invoke("mypkg::funcWithAllOptionalInputs", {
        "a": args.a ? inputs.helmReleaseSettingsProvideDefaults(args.a) : undefined,
        "b": args.b,
    }, opts);
}

export interface FuncWithAllOptionalInputsArgs {
    /**
     * Property A
     */
    a?: inputs.HelmReleaseSettings;
    /**
     * Property B
     */
    b?: string;
}

export interface FuncWithAllOptionalInputsResult {
    readonly r: string;
}
/**
 * Check codegen of functions with all optional inputs.
 */
export function funcWithAllOptionalInputsOutput(args?: FuncWithAllOptionalInputsOutputArgs, opts?: khulnasoft.InvokeOptions): khulnasoft.Output<FuncWithAllOptionalInputsResult> {
    args = args || {};
    opts = khulnasoft.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return khulnasoft.runtime.invokeOutput("mypkg::funcWithAllOptionalInputs", {
        "a": args.a ? khulnasoft.output(args.a).apply(inputs.helmReleaseSettingsProvideDefaults) : undefined,
        "b": args.b,
    }, opts);
}

export interface FuncWithAllOptionalInputsOutputArgs {
    /**
     * Property A
     */
    a?: khulnasoft.Input<inputs.HelmReleaseSettingsArgs>;
    /**
     * Property B
     */
    b?: khulnasoft.Input<string>;
}
