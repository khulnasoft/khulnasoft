// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as utilities from "./utilities";

export function getAssets(args: GetAssetsArgs, opts?: khulnasoft.InvokeOptions): Promise<GetAssetsResult> {
    opts = khulnasoft.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return khulnasoft.runtime.invoke("example::GetAssets", {
        "archive": args.archive,
        "source": args.source,
    }, opts);
}

export interface GetAssetsArgs {
    archive: khulnasoft.asset.Archive;
    source: khulnasoft.asset.Asset | khulnasoft.asset.Archive;
}

export interface GetAssetsResult {
    readonly archive: khulnasoft.asset.Archive;
    readonly source: khulnasoft.asset.Asset | khulnasoft.asset.Archive;
}
export function getAssetsOutput(args: GetAssetsOutputArgs, opts?: khulnasoft.InvokeOptions): khulnasoft.Output<GetAssetsResult> {
    opts = khulnasoft.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return khulnasoft.runtime.invokeOutput("example::GetAssets", {
        "archive": args.archive,
        "source": args.source,
    }, opts);
}

export interface GetAssetsOutputArgs {
    archive: khulnasoft.Input<khulnasoft.asset.Archive>;
    source: khulnasoft.Input<khulnasoft.asset.Asset | khulnasoft.asset.Archive>;
}
