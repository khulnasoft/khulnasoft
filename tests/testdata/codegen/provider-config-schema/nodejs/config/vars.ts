// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

declare var exports: any;
const __config = new khulnasoft.Config("configstation");

export declare const favoritePlants: string[] | undefined;
Object.defineProperty(exports, "favoritePlants", {
    get() {
        return __config.getObject<string[]>("favoritePlants");
    },
    enumerable: true,
});

/**
 * omg my favorite sandwich
 */
export declare const favoriteSandwich: outputs.config.Sandwich | undefined;
Object.defineProperty(exports, "favoriteSandwich", {
    get() {
        return __config.getObject<outputs.config.Sandwich>("favoriteSandwich");
    },
    enumerable: true,
});

export declare const isMember: boolean;
Object.defineProperty(exports, "isMember", {
    get() {
        return __config.getObject<boolean>("isMember") ?? true;
    },
    enumerable: true,
});

export declare const kids: outputs.Child | undefined;
Object.defineProperty(exports, "kids", {
    get() {
        return __config.getObject<outputs.Child>("kids");
    },
    enumerable: true,
});

export declare const name: string | undefined;
Object.defineProperty(exports, "name", {
    get() {
        return __config.get("name");
    },
    enumerable: true,
});

export declare const numberOfSheep: number | undefined;
Object.defineProperty(exports, "numberOfSheep", {
    get() {
        return __config.getObject<number>("numberOfSheep");
    },
    enumerable: true,
});

/**
 * This is a huge secret
 */
export declare const secretCode: string | undefined;
Object.defineProperty(exports, "secretCode", {
    get() {
        return __config.get("secretCode") ?? utilities.getEnv("SECRET_CODE", "MY_SUPER_SECRET_CODE");
    },
    enumerable: true,
});
