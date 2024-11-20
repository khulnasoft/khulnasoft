// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as utilities from "./utilities";

// Export members:
export { ProviderArgs } from "./provider";
export type Provider = import("./provider").Provider;
export const Provider: typeof import("./provider").Provider = null as any;
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));

export { ResArgs } from "./res";
export type Res = import("./res").Res;
export const Res: typeof import("./res").Res = null as any;
utilities.lazyLoad(exports, ["Res"], () => require("./res"));

export { TestArgs, TestResult, TestOutputArgs } from "./test";
export const test: typeof import("./test").test = null as any;
export const testOutput: typeof import("./test").testOutput = null as any;
utilities.lazyLoad(exports, ["test","testOutput"], () => require("./test"));


// Export sub-modules:
import * as types from "./types";

export {
    types,
};

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): khulnasoft.Resource => {
        switch (type) {
            case "urnid:index:Res":
                return new Res(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
khulnasoft.runtime.registerResourceModule("urnid", "index", _module)
khulnasoft.runtime.registerResourcePackage("urnid", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): khulnasoft.ProviderResource => {
        if (type !== "khulnasoft:providers:urnid") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
