// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as utilities from "./utilities";

// Export members:
export { Example_resourceArgs } from "./example_resource";
export type Example_resource = import("./example_resource").Example_resource;
export const Example_resource: typeof import("./example_resource").Example_resource = null as any;
utilities.lazyLoad(exports, ["Example_resource"], () => require("./example_resource"));

export { ProviderArgs } from "./provider";
export type Provider = import("./provider").Provider;
export const Provider: typeof import("./provider").Provider = null as any;
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));


// Export enums:
export * from "./types/enums";

// Export sub-modules:
import * as types from "./types";

export {
    types,
};

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): khulnasoft.Resource => {
        switch (type) {
            case "legacy_names:index:example_resource":
                return new Example_resource(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
khulnasoft.runtime.registerResourceModule("legacy_names", "index", _module)
khulnasoft.runtime.registerResourcePackage("legacy_names", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): khulnasoft.ProviderResource => {
        if (type !== "khulnasoft:providers:legacy_names") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});