// *** WARNING: this file was generated by khulnasoft-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as utilities from "./utilities";

// Export members:
export { MyInvokeArgs, MyInvokeResult, MyInvokeOutputArgs } from "./myInvoke";
export const myInvoke: typeof import("./myInvoke").myInvoke = null as any;
export const myInvokeOutput: typeof import("./myInvoke").myInvokeOutput = null as any;
utilities.lazyLoad(exports, ["myInvoke","myInvokeOutput"], () => require("./myInvoke"));

export { ProviderArgs } from "./provider";
export type Provider = import("./provider").Provider;
export const Provider: typeof import("./provider").Provider = null as any;
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));

export { SecretInvokeArgs, SecretInvokeResult, SecretInvokeOutputArgs } from "./secretInvoke";
export const secretInvoke: typeof import("./secretInvoke").secretInvoke = null as any;
export const secretInvokeOutput: typeof import("./secretInvoke").secretInvokeOutput = null as any;
utilities.lazyLoad(exports, ["secretInvoke","secretInvokeOutput"], () => require("./secretInvoke"));

export { StringResourceArgs } from "./stringResource";
export type StringResource = import("./stringResource").StringResource;
export const StringResource: typeof import("./stringResource").StringResource = null as any;
utilities.lazyLoad(exports, ["StringResource"], () => require("./stringResource"));

export { UnitArgs, UnitResult } from "./unit";
export const unit: typeof import("./unit").unit = null as any;
export const unitOutput: typeof import("./unit").unitOutput = null as any;
utilities.lazyLoad(exports, ["unit","unitOutput"], () => require("./unit"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): khulnasoft.Resource => {
        switch (type) {
            case "simple-invoke:index:StringResource":
                return new StringResource(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
khulnasoft.runtime.registerResourceModule("simple-invoke", "index", _module)
khulnasoft.runtime.registerResourcePackage("simple-invoke", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): khulnasoft.ProviderResource => {
        if (type !== "khulnasoft:providers:simple-invoke") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
