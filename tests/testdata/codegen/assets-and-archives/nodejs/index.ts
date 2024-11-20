// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as utilities from "./utilities";

// Export members:
export { GetAssetsArgs, GetAssetsResult, GetAssetsOutputArgs } from "./getAssets";
export const getAssets: typeof import("./getAssets").getAssets = null as any;
export const getAssetsOutput: typeof import("./getAssets").getAssetsOutput = null as any;
utilities.lazyLoad(exports, ["getAssets","getAssetsOutput"], () => require("./getAssets"));

export { ProviderArgs } from "./provider";
export type Provider = import("./provider").Provider;
export const Provider: typeof import("./provider").Provider = null as any;
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));

export { ResourceWithAssetsArgs } from "./resourceWithAssets";
export type ResourceWithAssets = import("./resourceWithAssets").ResourceWithAssets;
export const ResourceWithAssets: typeof import("./resourceWithAssets").ResourceWithAssets = null as any;
utilities.lazyLoad(exports, ["ResourceWithAssets"], () => require("./resourceWithAssets"));


// Export sub-modules:
import * as types from "./types";

export {
    types,
};

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): khulnasoft.Resource => {
        switch (type) {
            case "example:index:ResourceWithAssets":
                return new ResourceWithAssets(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
khulnasoft.runtime.registerResourceModule("example", "index", _module)
khulnasoft.runtime.registerResourcePackage("example", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): khulnasoft.ProviderResource => {
        if (type !== "khulnasoft:providers:example") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});