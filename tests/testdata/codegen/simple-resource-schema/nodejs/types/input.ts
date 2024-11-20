// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

import {Resource} from "..";

export interface ConfigMapArgs {
    config?: khulnasoft.Input<string>;
}

export interface ObjectArgs {
    bar?: khulnasoft.Input<string>;
    configs?: khulnasoft.Input<khulnasoft.Input<inputs.ConfigMapArgs>[]>;
    foo?: khulnasoft.Input<Resource>;
    /**
     * List of lists of other objects
     */
    others?: khulnasoft.Input<khulnasoft.Input<khulnasoft.Input<inputs.SomeOtherObjectArgs>[]>[]>;
    /**
     * Mapping from string to list of some other object
     */
    stillOthers?: khulnasoft.Input<{[key: string]: khulnasoft.Input<khulnasoft.Input<inputs.SomeOtherObjectArgs>[]>}>;
}

export interface ObjectWithNodeOptionalInputsArgs {
    bar?: khulnasoft.Input<number>;
    foo?: khulnasoft.Input<string>;
}

export interface SomeOtherObjectArgs {
    baz?: khulnasoft.Input<string>;
}