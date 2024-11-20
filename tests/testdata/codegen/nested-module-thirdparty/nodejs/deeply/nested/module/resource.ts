// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as utilities from "../../../utilities";

export class Resource extends khulnasoft.CustomResource {
    /**
     * Get an existing Resource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: khulnasoft.Input<khulnasoft.ID>, opts?: khulnasoft.CustomResourceOptions): Resource {
        return new Resource(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __khulnasoftType = 'foo-bar:deeply/nested/module:Resource';

    /**
     * Returns true if the given object is an instance of Resource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Resource {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__khulnasoftType'] === Resource.__khulnasoftType;
    }

    public readonly baz!: khulnasoft.Output<string | undefined>;

    /**
     * Create a Resource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ResourceArgs, opts?: khulnasoft.CustomResourceOptions) {
        let resourceInputs: khulnasoft.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["baz"] = args?.baz ? khulnasoft.secret(args.baz) : undefined;
        } else {
            resourceInputs["baz"] = undefined /*out*/;
        }
        opts = khulnasoft.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["baz"] };
        opts = khulnasoft.mergeOptions(opts, secretOpts);
        super(Resource.__khulnasoftType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Resource resource.
 */
export interface ResourceArgs {
    baz?: khulnasoft.Input<string>;
}
