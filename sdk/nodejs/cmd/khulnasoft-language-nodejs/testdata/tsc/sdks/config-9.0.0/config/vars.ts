// *** WARNING: this file was generated by khulnasoft-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as utilities from "../utilities";

declare var exports: any;
const __config = new khulnasoft.Config("config");

export declare const name: string | undefined;
Object.defineProperty(exports, "name", {
    get() {
        return __config.get("name");
    },
    enumerable: true,
});

export declare const pluginDownloadURL: string | undefined;
Object.defineProperty(exports, "pluginDownloadURL", {
    get() {
        return __config.get("pluginDownloadURL");
    },
    enumerable: true,
});

