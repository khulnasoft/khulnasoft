// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as utilities from "./utilities";

// Export members:
export { MainComponentArgs } from "./mainComponent";
export type MainComponent = import("./mainComponent").MainComponent;
export const MainComponent: typeof import("./mainComponent").MainComponent = null as any;
utilities.lazyLoad(exports, ["MainComponent"], () => require("./mainComponent"));

export { ProviderArgs } from "./provider";
export type Provider = import("./provider").Provider;
export const Provider: typeof import("./provider").Provider = null as any;
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));

export { ResourceArgs } from "./resource";
export type Resource = import("./resource").Resource;
export const Resource: typeof import("./resource").Resource = null as any;
utilities.lazyLoad(exports, ["Resource"], () => require("./resource"));

export { ResourceInputArgs } from "./resourceInput";
export type ResourceInput = import("./resourceInput").ResourceInput;
export const ResourceInput: typeof import("./resourceInput").ResourceInput = null as any;
utilities.lazyLoad(exports, ["ResourceInput"], () => require("./resourceInput"));


// Export enums:
export * from "./types/enums";

// Export sub-modules:
import * as mod from "./mod";
import * as types from "./types";

export {
    mod,
    types,
};

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): khulnasoft.Resource => {
        switch (type) {
            case "example::MainComponent":
                return new MainComponent(name, <any>undefined, { urn })
            case "example::Resource":
                return new Resource(name, <any>undefined, { urn })
            case "example::ResourceInput":
                return new ResourceInput(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
khulnasoft.runtime.registerResourceModule("example", "", _module)
khulnasoft.runtime.registerResourcePackage("example", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): khulnasoft.ProviderResource => {
        if (type !== "khulnasoft:providers:example") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
