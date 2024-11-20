// *** WARNING: this file was generated by khulnasoft-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as utilities from "./utilities";

export class Provider extends khulnasoft.ProviderResource {
    /** @internal */
    public static readonly __khulnasoftType = 'config';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__khulnasoftType'] === "khulnasoft:providers:" + Provider.__khulnasoftType;
    }

    public readonly name!: khulnasoft.Output<string>;
    public readonly pluginDownloadURL!: khulnasoft.Output<string | undefined>;
    public /*out*/ readonly version!: khulnasoft.Output<string>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProviderArgs, opts?: khulnasoft.ResourceOptions) {
        let resourceInputs: khulnasoft.Inputs = {};
        opts = opts || {};
        {
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["pluginDownloadURL"] = args ? args.pluginDownloadURL : undefined;
            resourceInputs["version"] = undefined /*out*/;
        }
        opts = khulnasoft.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Provider.__khulnasoftType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    name: khulnasoft.Input<string>;
    pluginDownloadURL?: khulnasoft.Input<string>;
}