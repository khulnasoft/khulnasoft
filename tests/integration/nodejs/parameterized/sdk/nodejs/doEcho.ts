// *** WARNING: this file was generated by khulnasoft-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as utilities from "./utilities";

/**
 * A test invoke that echoes its input.
 */
export function doEcho(args?: DoEchoArgs, opts?: khulnasoft.InvokeOptions): Promise<DoEchoResult> {
    args = args || {};

    opts = khulnasoft.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return khulnasoft.runtime.invoke("pkg:index:doEcho", {
        "echo": args.echo,
    }, opts, utilities.getPackage());
}

export interface DoEchoArgs {
    echo?: string;
}

export interface DoEchoResult {
    readonly echo?: string;
}
/**
 * A test invoke that echoes its input.
 */
export function doEchoOutput(args?: DoEchoOutputArgs, opts?: khulnasoft.InvokeOptions): khulnasoft.Output<DoEchoResult> {
    return khulnasoft.output(args).apply((a: any) => doEcho(a, opts))
}

export interface DoEchoOutputArgs {
    echo?: khulnasoft.Input<string>;
}
