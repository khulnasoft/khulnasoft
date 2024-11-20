// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

/**
 * The list of product families.
 * API Version: 2020-12-01-preview.
 */
export function listProductFamilies(args: ListProductFamiliesArgs, opts?: khulnasoft.InvokeOptions): Promise<ListProductFamiliesResult> {
    opts = khulnasoft.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return khulnasoft.runtime.invoke("myedgeorder::listProductFamilies", {
        "customerSubscriptionDetails": args.customerSubscriptionDetails,
        "expand": args.expand,
        "filterableProperties": args.filterableProperties,
        "skipToken": args.skipToken,
    }, opts);
}

export interface ListProductFamiliesArgs {
    /**
     * Customer subscription properties. Clients can display available products to unregistered customers by explicitly passing subscription details
     */
    customerSubscriptionDetails?: inputs.CustomerSubscriptionDetails;
    /**
     * $expand is supported on configurations parameter for product, which provides details on the configurations for the product.
     */
    expand?: string;
    /**
     * Dictionary of filterable properties on product family.
     */
    filterableProperties: {[key: string]: inputs.FilterableProperty[]};
    /**
     * $skipToken is supported on list of product families, which provides the next page in the list of product families.
     */
    skipToken?: string;
}

/**
 * The list of product families.
 */
export interface ListProductFamiliesResult {
    /**
     * Link for the next set of product families.
     */
    readonly nextLink?: string;
    /**
     * List of product families.
     */
    readonly value: outputs.ProductFamilyResponse[];
}
/**
 * The list of product families.
 * API Version: 2020-12-01-preview.
 */
export function listProductFamiliesOutput(args: ListProductFamiliesOutputArgs, opts?: khulnasoft.InvokeOptions): khulnasoft.Output<ListProductFamiliesResult> {
    opts = khulnasoft.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return khulnasoft.runtime.invokeOutput("myedgeorder::listProductFamilies", {
        "customerSubscriptionDetails": args.customerSubscriptionDetails,
        "expand": args.expand,
        "filterableProperties": args.filterableProperties,
        "skipToken": args.skipToken,
    }, opts);
}

export interface ListProductFamiliesOutputArgs {
    /**
     * Customer subscription properties. Clients can display available products to unregistered customers by explicitly passing subscription details
     */
    customerSubscriptionDetails?: khulnasoft.Input<inputs.CustomerSubscriptionDetailsArgs>;
    /**
     * $expand is supported on configurations parameter for product, which provides details on the configurations for the product.
     */
    expand?: khulnasoft.Input<string>;
    /**
     * Dictionary of filterable properties on product family.
     */
    filterableProperties: khulnasoft.Input<{[key: string]: khulnasoft.Input<khulnasoft.Input<inputs.FilterablePropertyArgs>[]>}>;
    /**
     * $skipToken is supported on list of product families, which provides the next page in the list of product families.
     */
    skipToken?: khulnasoft.Input<string>;
}
