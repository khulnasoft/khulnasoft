// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

export class ModuleResource extends khulnasoft.CustomResource {
    /**
     * Get an existing ModuleResource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: khulnasoft.Input<khulnasoft.ID>, opts?: khulnasoft.CustomResourceOptions): ModuleResource {
        return new ModuleResource(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __khulnasoftType = 'foobar::ModuleResource';

    /**
     * Returns true if the given object is an instance of ModuleResource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ModuleResource {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__khulnasoftType'] === ModuleResource.__khulnasoftType;
    }

    public readonly optionalBool!: khulnasoft.Output<boolean | undefined>;

    /**
     * Create a ModuleResource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ModuleResourceArgs, opts?: khulnasoft.CustomResourceOptions) {
        let resourceInputs: khulnasoft.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.plainRequiredBool === undefined) && !opts.urn) {
                throw new Error("Missing required property 'plainRequiredBool'");
            }
            if ((!args || args.plainRequiredConst === undefined) && !opts.urn) {
                throw new Error("Missing required property 'plainRequiredConst'");
            }
            if ((!args || args.plainRequiredNumber === undefined) && !opts.urn) {
                throw new Error("Missing required property 'plainRequiredNumber'");
            }
            if ((!args || args.plainRequiredString === undefined) && !opts.urn) {
                throw new Error("Missing required property 'plainRequiredString'");
            }
            if ((!args || args.requiredBool === undefined) && !opts.urn) {
                throw new Error("Missing required property 'requiredBool'");
            }
            if ((!args || args.requiredEnum === undefined) && !opts.urn) {
                throw new Error("Missing required property 'requiredEnum'");
            }
            if ((!args || args.requiredNumber === undefined) && !opts.urn) {
                throw new Error("Missing required property 'requiredNumber'");
            }
            if ((!args || args.requiredString === undefined) && !opts.urn) {
                throw new Error("Missing required property 'requiredString'");
            }
            resourceInputs["optionalBool"] = (args ? args.optionalBool : undefined) ?? true;
            resourceInputs["optionalConst"] = "val";
            resourceInputs["optionalEnum"] = (args ? args.optionalEnum : undefined) ?? 8;
            resourceInputs["optionalNumber"] = (args ? args.optionalNumber : undefined) ?? 42;
            resourceInputs["optionalString"] = (args ? args.optionalString : undefined) ?? "buzzer";
            resourceInputs["plainOptionalBool"] = (args ? args.plainOptionalBool : undefined) ?? true;
            resourceInputs["plainOptionalConst"] = "val";
            resourceInputs["plainOptionalNumber"] = (args ? args.plainOptionalNumber : undefined) ?? 42;
            resourceInputs["plainOptionalString"] = (args ? args.plainOptionalString : undefined) ?? "buzzer";
            resourceInputs["plainRequiredBool"] = (args ? args.plainRequiredBool : undefined) ?? true;
            resourceInputs["plainRequiredConst"] = "val";
            resourceInputs["plainRequiredNumber"] = (args ? args.plainRequiredNumber : undefined) ?? 42;
            resourceInputs["plainRequiredString"] = (args ? args.plainRequiredString : undefined) ?? "buzzer";
            resourceInputs["requiredBool"] = (args ? args.requiredBool : undefined) ?? true;
            resourceInputs["requiredEnum"] = (args ? args.requiredEnum : undefined) ?? 4;
            resourceInputs["requiredNumber"] = (args ? args.requiredNumber : undefined) ?? 42;
            resourceInputs["requiredString"] = (args ? args.requiredString : undefined) ?? "buzzer";
        } else {
            resourceInputs["optionalBool"] = undefined /*out*/;
        }
        opts = khulnasoft.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ModuleResource.__khulnasoftType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ModuleResource resource.
 */
export interface ModuleResourceArgs {
    optionalBool?: khulnasoft.Input<boolean>;
    optionalConst?: khulnasoft.Input<"val">;
    optionalEnum?: khulnasoft.Input<enums.EnumThing>;
    optionalNumber?: khulnasoft.Input<number>;
    optionalString?: khulnasoft.Input<string>;
    plainOptionalBool?: boolean;
    plainOptionalConst?: "val";
    plainOptionalNumber?: number;
    plainOptionalString?: string;
    plainRequiredBool: boolean;
    plainRequiredConst: "val";
    plainRequiredNumber: number;
    plainRequiredString: string;
    requiredBool: khulnasoft.Input<boolean>;
    requiredEnum: khulnasoft.Input<enums.EnumThing>;
    requiredNumber: khulnasoft.Input<number>;
    requiredString: khulnasoft.Input<string>;
}