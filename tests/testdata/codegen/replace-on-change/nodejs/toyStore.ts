// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

import {Cat, Dog} from "./index";

export class ToyStore extends khulnasoft.CustomResource {
    /**
     * Get an existing ToyStore resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: khulnasoft.Input<khulnasoft.ID>, opts?: khulnasoft.CustomResourceOptions): ToyStore {
        return new ToyStore(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __khulnasoftType = 'example::ToyStore';

    /**
     * Returns true if the given object is an instance of ToyStore.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ToyStore {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__khulnasoftType'] === ToyStore.__khulnasoftType;
    }

    public /*out*/ readonly chew!: khulnasoft.Output<outputs.Chew | undefined>;
    public /*out*/ readonly laser!: khulnasoft.Output<outputs.Laser | undefined>;
    public /*out*/ readonly stuff!: khulnasoft.Output<outputs.Toy[] | undefined>;
    public /*out*/ readonly wanted!: khulnasoft.Output<outputs.Toy[] | undefined>;

    /**
     * Create a ToyStore resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ToyStoreArgs, opts?: khulnasoft.CustomResourceOptions) {
        let resourceInputs: khulnasoft.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["chew"] = undefined /*out*/;
            resourceInputs["laser"] = undefined /*out*/;
            resourceInputs["stuff"] = undefined /*out*/;
            resourceInputs["wanted"] = undefined /*out*/;
        } else {
            resourceInputs["chew"] = undefined /*out*/;
            resourceInputs["laser"] = undefined /*out*/;
            resourceInputs["stuff"] = undefined /*out*/;
            resourceInputs["wanted"] = undefined /*out*/;
        }
        opts = khulnasoft.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["chew.owner", "laser.batteries", "stuff[*].associated.color", "stuff[*].color", "wanted[*]"] };
        opts = khulnasoft.mergeOptions(opts, replaceOnChanges);
        super(ToyStore.__khulnasoftType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ToyStore resource.
 */
export interface ToyStoreArgs {
}
