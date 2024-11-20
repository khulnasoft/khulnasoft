// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as utilities from "../utilities";

import {MainComponent} from "..";
import {Component2} from "./index";

export class Component extends khulnasoft.CustomResource {
    /**
     * Get an existing Component resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: khulnasoft.Input<khulnasoft.ID>, opts?: khulnasoft.CustomResourceOptions): Component {
        return new Component(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __khulnasoftType = 'example:mod:Component';

    /**
     * Returns true if the given object is an instance of Component.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Component {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__khulnasoftType'] === Component.__khulnasoftType;
    }


    /**
     * Create a Component resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ComponentArgs, opts?: khulnasoft.CustomResourceOptions) {
        let resourceInputs: khulnasoft.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["local"] = args ? args.local : undefined;
            resourceInputs["main"] = args ? args.main : undefined;
        } else {
        }
        opts = khulnasoft.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Component.__khulnasoftType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Component resource.
 */
export interface ComponentArgs {
    local?: khulnasoft.Input<Component2>;
    main?: khulnasoft.Input<MainComponent>;
}
