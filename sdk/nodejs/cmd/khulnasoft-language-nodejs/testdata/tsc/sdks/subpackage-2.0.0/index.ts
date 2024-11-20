// *** WARNING: this file was generated by khulnasoft-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as utilities from "./utilities";

// Export members:
export { HelloWorldArgs } from "./helloWorld";
export type HelloWorld = import("./helloWorld").HelloWorld;
export const HelloWorld: typeof import("./helloWorld").HelloWorld = null as any;
utilities.lazyLoad(exports, ["HelloWorld"], () => require("./helloWorld"));

export { ProviderArgs } from "./provider";
export type Provider = import("./provider").Provider;
export const Provider: typeof import("./provider").Provider = null as any;
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): khulnasoft.Resource => {
        switch (type) {
            case "subpackage:index:HelloWorld":
                return new HelloWorld(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
khulnasoft.runtime.registerResourceModule("subpackage", "index", _module)
khulnasoft.runtime.registerResourcePackage("subpackage", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): khulnasoft.ProviderResource => {
        if (type !== "khulnasoft:providers:subpackage") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});