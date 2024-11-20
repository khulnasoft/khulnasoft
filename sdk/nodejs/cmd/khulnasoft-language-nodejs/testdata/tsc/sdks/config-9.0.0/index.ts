// *** WARNING: this file was generated by khulnasoft-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as utilities from "./utilities";

// Export members:
export { ProviderArgs } from "./provider";
export type Provider = import("./provider").Provider;
export const Provider: typeof import("./provider").Provider = null as any;
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));

export { ResourceArgs } from "./resource";
export type Resource = import("./resource").Resource;
export const Resource: typeof import("./resource").Resource = null as any;
utilities.lazyLoad(exports, ["Resource"], () => require("./resource"));


// Export sub-modules:
import * as config from "./config";

export {
    config,
};

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): khulnasoft.Resource => {
        switch (type) {
            case "config:index:Resource":
                return new Resource(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
khulnasoft.runtime.registerResourceModule("config", "index", _module)
khulnasoft.runtime.registerResourcePackage("config", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): khulnasoft.ProviderResource => {
        if (type !== "khulnasoft:providers:config") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
