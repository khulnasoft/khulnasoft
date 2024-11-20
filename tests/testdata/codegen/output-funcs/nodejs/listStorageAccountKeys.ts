// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The response from the ListKeys operation.
 * API Version: 2021-02-01.
 */
export function listStorageAccountKeys(args: ListStorageAccountKeysArgs, opts?: khulnasoft.InvokeOptions): Promise<ListStorageAccountKeysResult> {
    opts = khulnasoft.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return khulnasoft.runtime.invoke("mypkg::listStorageAccountKeys", {
        "accountName": args.accountName,
        "expand": args.expand,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListStorageAccountKeysArgs {
    /**
     * The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
     */
    accountName: string;
    /**
     * Specifies type of the key to be listed. Possible value is kerb.
     */
    expand?: string;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * The response from the ListKeys operation.
 */
export interface ListStorageAccountKeysResult {
    /**
     * Gets the list of storage account keys and their properties for the specified storage account.
     */
    readonly keys: outputs.StorageAccountKeyResponse[];
}
/**
 * The response from the ListKeys operation.
 * API Version: 2021-02-01.
 */
export function listStorageAccountKeysOutput(args: ListStorageAccountKeysOutputArgs, opts?: khulnasoft.InvokeOptions): khulnasoft.Output<ListStorageAccountKeysResult> {
    opts = khulnasoft.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return khulnasoft.runtime.invokeOutput("mypkg::listStorageAccountKeys", {
        "accountName": args.accountName,
        "expand": args.expand,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListStorageAccountKeysOutputArgs {
    /**
     * The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
     */
    accountName: khulnasoft.Input<string>;
    /**
     * Specifies type of the key to be listed. Possible value is kerb.
     */
    expand?: khulnasoft.Input<string>;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    resourceGroupName: khulnasoft.Input<string>;
}
