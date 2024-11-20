// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as utilities from "./utilities";

// Export members:
export { ProviderArgs } from "./provider";
export type Provider = import("./provider").Provider;
export const Provider: typeof import("./provider").Provider = null as any;
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));

export { RegistryGeoReplicationArgs } from "./registryGeoReplication";
export type RegistryGeoReplication = import("./registryGeoReplication").RegistryGeoReplication;
export const RegistryGeoReplication: typeof import("./registryGeoReplication").RegistryGeoReplication = null as any;
utilities.lazyLoad(exports, ["RegistryGeoReplication"], () => require("./registryGeoReplication"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): khulnasoft.Resource => {
        switch (type) {
            case "registrygeoreplication:index:RegistryGeoReplication":
                return new RegistryGeoReplication(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
khulnasoft.runtime.registerResourceModule("registrygeoreplication", "index", _module)
khulnasoft.runtime.registerResourcePackage("registrygeoreplication", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): khulnasoft.ProviderResource => {
        if (type !== "khulnasoft:providers:registrygeoreplication") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});