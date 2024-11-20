// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as utilities from "./utilities";

export class BasicResourceV2 extends khulnasoft.CustomResource {
    /**
     * Get an existing BasicResourceV2 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: khulnasoft.Input<khulnasoft.ID>, opts?: khulnasoft.CustomResourceOptions): BasicResourceV2 {
        return new BasicResourceV2(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __khulnasoftType = 'example:index:BasicResourceV2';

    /**
     * Returns true if the given object is an instance of BasicResourceV2.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BasicResourceV2 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__khulnasoftType'] === BasicResourceV2.__khulnasoftType;
    }

    public readonly bar!: khulnasoft.Output<string>;

    /**
     * Create a BasicResourceV2 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BasicResourceV2Args, opts?: khulnasoft.CustomResourceOptions) {
        let resourceInputs: khulnasoft.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.bar === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bar'");
            }
            resourceInputs["bar"] = args ? args.bar : undefined;
        } else {
            resourceInputs["bar"] = undefined /*out*/;
        }
        opts = khulnasoft.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "example:index:BasicResource" }] };
        opts = khulnasoft.mergeOptions(opts, aliasOpts);
        super(BasicResourceV2.__khulnasoftType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a BasicResourceV2 resource.
 */
export interface BasicResourceV2Args {
    bar: khulnasoft.Input<string>;
}
