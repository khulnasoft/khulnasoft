// *** WARNING: this file was generated by khulnasoft-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function toSecret(args?: ToSecretArgs, opts?: khulnasoft.InvokeOptions): Promise<ToSecretResult> {
    args = args || {};
    opts = khulnasoft.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return khulnasoft.runtime.invoke("config-grpc:index:toSecret", {
        "bool1": args.bool1,
        "bool2": args.bool2,
        "bool3": args.bool3,
        "int1": args.int1,
        "int2": args.int2,
        "int3": args.int3,
        "listBool1": args.listBool1,
        "listBool2": args.listBool2,
        "listBool3": args.listBool3,
        "listInt1": args.listInt1,
        "listInt2": args.listInt2,
        "listInt3": args.listInt3,
        "listNum1": args.listNum1,
        "listNum2": args.listNum2,
        "listNum3": args.listNum3,
        "listSecretBool1": args.listSecretBool1,
        "listSecretBool2": args.listSecretBool2,
        "listSecretBool3": args.listSecretBool3,
        "listSecretInt1": args.listSecretInt1,
        "listSecretInt2": args.listSecretInt2,
        "listSecretInt3": args.listSecretInt3,
        "listSecretNum1": args.listSecretNum1,
        "listSecretNum2": args.listSecretNum2,
        "listSecretNum3": args.listSecretNum3,
        "listSecretString1": args.listSecretString1,
        "listSecretString2": args.listSecretString2,
        "listSecretString3": args.listSecretString3,
        "listString1": args.listString1,
        "listString2": args.listString2,
        "listString3": args.listString3,
        "mapBool1": args.mapBool1,
        "mapBool2": args.mapBool2,
        "mapBool3": args.mapBool3,
        "mapInt1": args.mapInt1,
        "mapInt2": args.mapInt2,
        "mapInt3": args.mapInt3,
        "mapNum1": args.mapNum1,
        "mapNum2": args.mapNum2,
        "mapNum3": args.mapNum3,
        "mapSecretBool1": args.mapSecretBool1,
        "mapSecretBool2": args.mapSecretBool2,
        "mapSecretBool3": args.mapSecretBool3,
        "mapSecretInt1": args.mapSecretInt1,
        "mapSecretInt2": args.mapSecretInt2,
        "mapSecretInt3": args.mapSecretInt3,
        "mapSecretNum1": args.mapSecretNum1,
        "mapSecretNum2": args.mapSecretNum2,
        "mapSecretNum3": args.mapSecretNum3,
        "mapSecretString1": args.mapSecretString1,
        "mapSecretString2": args.mapSecretString2,
        "mapSecretString3": args.mapSecretString3,
        "mapString1": args.mapString1,
        "mapString2": args.mapString2,
        "mapString3": args.mapString3,
        "num1": args.num1,
        "num2": args.num2,
        "num3": args.num3,
        "objBool1": args.objBool1,
        "objBool2": args.objBool2,
        "objBool3": args.objBool3,
        "objInt1": args.objInt1,
        "objInt2": args.objInt2,
        "objInt3": args.objInt3,
        "objNum1": args.objNum1,
        "objNum2": args.objNum2,
        "objNum3": args.objNum3,
        "objSecretBool1": args.objSecretBool1,
        "objSecretBool2": args.objSecretBool2,
        "objSecretBool3": args.objSecretBool3,
        "objSecretInt1": args.objSecretInt1,
        "objSecretInt2": args.objSecretInt2,
        "objSecretInt3": args.objSecretInt3,
        "objSecretNum1": args.objSecretNum1,
        "objSecretNum2": args.objSecretNum2,
        "objSecretNum3": args.objSecretNum3,
        "objSecretString1": args.objSecretString1,
        "objSecretString2": args.objSecretString2,
        "objSecretString3": args.objSecretString3,
        "objString1": args.objString1,
        "objString2": args.objString2,
        "objString3": args.objString3,
        "secretBool1": args.secretBool1,
        "secretBool2": args.secretBool2,
        "secretBool3": args.secretBool3,
        "secretInt1": args.secretInt1,
        "secretInt2": args.secretInt2,
        "secretInt3": args.secretInt3,
        "secretNum1": args.secretNum1,
        "secretNum2": args.secretNum2,
        "secretNum3": args.secretNum3,
        "secretString1": args.secretString1,
        "secretString2": args.secretString2,
        "secretString3": args.secretString3,
        "string1": args.string1,
        "string2": args.string2,
        "string3": args.string3,
    }, opts);
}

export interface ToSecretArgs {
    bool1?: boolean;
    bool2?: boolean;
    bool3?: boolean;
    int1?: number;
    int2?: number;
    int3?: number;
    listBool1?: boolean[];
    listBool2?: boolean[];
    listBool3?: boolean[];
    listInt1?: number[];
    listInt2?: number[];
    listInt3?: number[];
    listNum1?: number[];
    listNum2?: number[];
    listNum3?: number[];
    listSecretBool1?: boolean[];
    listSecretBool2?: boolean[];
    listSecretBool3?: boolean[];
    listSecretInt1?: number[];
    listSecretInt2?: number[];
    listSecretInt3?: number[];
    listSecretNum1?: number[];
    listSecretNum2?: number[];
    listSecretNum3?: number[];
    listSecretString1?: string[];
    listSecretString2?: string[];
    listSecretString3?: string[];
    listString1?: string[];
    listString2?: string[];
    listString3?: string[];
    mapBool1?: {[key: string]: boolean};
    mapBool2?: {[key: string]: boolean};
    mapBool3?: {[key: string]: boolean};
    mapInt1?: {[key: string]: number};
    mapInt2?: {[key: string]: number};
    mapInt3?: {[key: string]: number};
    mapNum1?: {[key: string]: number};
    mapNum2?: {[key: string]: number};
    mapNum3?: {[key: string]: number};
    mapSecretBool1?: {[key: string]: boolean};
    mapSecretBool2?: {[key: string]: boolean};
    mapSecretBool3?: {[key: string]: boolean};
    mapSecretInt1?: {[key: string]: number};
    mapSecretInt2?: {[key: string]: number};
    mapSecretInt3?: {[key: string]: number};
    mapSecretNum1?: {[key: string]: number};
    mapSecretNum2?: {[key: string]: number};
    mapSecretNum3?: {[key: string]: number};
    mapSecretString1?: {[key: string]: string};
    mapSecretString2?: {[key: string]: string};
    mapSecretString3?: {[key: string]: string};
    mapString1?: {[key: string]: string};
    mapString2?: {[key: string]: string};
    mapString3?: {[key: string]: string};
    num1?: number;
    num2?: number;
    num3?: number;
    objBool1?: inputs.Tbool1;
    objBool2?: inputs.Tbool2;
    objBool3?: inputs.Tbool3;
    objInt1?: inputs.Tint1;
    objInt2?: inputs.Tint2;
    objInt3?: inputs.Tint3;
    objNum1?: inputs.Tnum1;
    objNum2?: inputs.Tnum2;
    objNum3?: inputs.Tnum3;
    objSecretBool1?: inputs.TsecretBool1;
    objSecretBool2?: inputs.TsecretBool2;
    objSecretBool3?: inputs.TsecretBool3;
    objSecretInt1?: inputs.TsecretInt1;
    objSecretInt2?: inputs.TsecretInt2;
    objSecretInt3?: inputs.TsecretInt3;
    objSecretNum1?: inputs.TsecretNum1;
    objSecretNum2?: inputs.TsecretNum2;
    objSecretNum3?: inputs.TsecretNum3;
    objSecretString1?: inputs.TsecretString1;
    objSecretString2?: inputs.TsecretString2;
    objSecretString3?: inputs.TsecretString3;
    objString1?: inputs.Tstring1;
    objString2?: inputs.Tstring2;
    objString3?: inputs.Tstring3;
    secretBool1?: boolean;
    secretBool2?: boolean;
    secretBool3?: boolean;
    secretInt1?: number;
    secretInt2?: number;
    secretInt3?: number;
    secretNum1?: number;
    secretNum2?: number;
    secretNum3?: number;
    secretString1?: string;
    secretString2?: string;
    secretString3?: string;
    string1?: string;
    string2?: string;
    string3?: string;
}

export interface ToSecretResult {
    readonly bool1: boolean;
    readonly bool2: boolean;
    readonly bool3: boolean;
    readonly int1: number;
    readonly int2: number;
    readonly int3: number;
    readonly listBool1: boolean[];
    readonly listBool2: boolean[];
    readonly listBool3: boolean[];
    readonly listInt1: number[];
    readonly listInt2: number[];
    readonly listInt3: number[];
    readonly listNum1: number[];
    readonly listNum2: number[];
    readonly listNum3: number[];
    readonly listSecretBool1: boolean[];
    readonly listSecretBool2: boolean[];
    readonly listSecretBool3: boolean[];
    readonly listSecretInt1: number[];
    readonly listSecretInt2: number[];
    readonly listSecretInt3: number[];
    readonly listSecretNum1: number[];
    readonly listSecretNum2: number[];
    readonly listSecretNum3: number[];
    readonly listSecretString1: string[];
    readonly listSecretString2: string[];
    readonly listSecretString3: string[];
    readonly listString1: string[];
    readonly listString2: string[];
    readonly listString3: string[];
    readonly mapBool1: {[key: string]: boolean};
    readonly mapBool2: {[key: string]: boolean};
    readonly mapBool3: {[key: string]: boolean};
    readonly mapInt1: {[key: string]: number};
    readonly mapInt2: {[key: string]: number};
    readonly mapInt3: {[key: string]: number};
    readonly mapNum1: {[key: string]: number};
    readonly mapNum2: {[key: string]: number};
    readonly mapNum3: {[key: string]: number};
    readonly mapSecretBool1: {[key: string]: boolean};
    readonly mapSecretBool2: {[key: string]: boolean};
    readonly mapSecretBool3: {[key: string]: boolean};
    readonly mapSecretInt1: {[key: string]: number};
    readonly mapSecretInt2: {[key: string]: number};
    readonly mapSecretInt3: {[key: string]: number};
    readonly mapSecretNum1: {[key: string]: number};
    readonly mapSecretNum2: {[key: string]: number};
    readonly mapSecretNum3: {[key: string]: number};
    readonly mapSecretString1: {[key: string]: string};
    readonly mapSecretString2: {[key: string]: string};
    readonly mapSecretString3: {[key: string]: string};
    readonly mapString1: {[key: string]: string};
    readonly mapString2: {[key: string]: string};
    readonly mapString3: {[key: string]: string};
    readonly num1: number;
    readonly num2: number;
    readonly num3: number;
    readonly objBool1: outputs.Tbool1;
    readonly objBool2: outputs.Tbool2;
    readonly objBool3: outputs.Tbool3;
    readonly objInt1: outputs.Tint1;
    readonly objInt2: outputs.Tint2;
    readonly objInt3: outputs.Tint3;
    readonly objNum1: outputs.Tnum1;
    readonly objNum2: outputs.Tnum2;
    readonly objNum3: outputs.Tnum3;
    readonly objSecretBool1: outputs.TsecretBool1;
    readonly objSecretBool2: outputs.TsecretBool2;
    readonly objSecretBool3: outputs.TsecretBool3;
    readonly objSecretInt1: outputs.TsecretInt1;
    readonly objSecretInt2: outputs.TsecretInt2;
    readonly objSecretInt3: outputs.TsecretInt3;
    readonly objSecretNum1: outputs.TsecretNum1;
    readonly objSecretNum2: outputs.TsecretNum2;
    readonly objSecretNum3: outputs.TsecretNum3;
    readonly objSecretString1: outputs.TsecretString1;
    readonly objSecretString2: outputs.TsecretString2;
    readonly objSecretString3: outputs.TsecretString3;
    readonly objString1: outputs.Tstring1;
    readonly objString2: outputs.Tstring2;
    readonly objString3: outputs.Tstring3;
    readonly secretBool1: boolean;
    readonly secretBool2: boolean;
    readonly secretBool3: boolean;
    readonly secretInt1: number;
    readonly secretInt2: number;
    readonly secretInt3: number;
    readonly secretNum1: number;
    readonly secretNum2: number;
    readonly secretNum3: number;
    readonly secretString1: string;
    readonly secretString2: string;
    readonly secretString3: string;
    readonly string1: string;
    readonly string2: string;
    readonly string3: string;
}
export function toSecretOutput(args?: ToSecretOutputArgs, opts?: khulnasoft.InvokeOptions): khulnasoft.Output<ToSecretResult> {
    args = args || {};
    opts = khulnasoft.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return khulnasoft.runtime.invokeOutput("config-grpc:index:toSecret", {
        "bool1": args.bool1,
        "bool2": args.bool2,
        "bool3": args.bool3,
        "int1": args.int1,
        "int2": args.int2,
        "int3": args.int3,
        "listBool1": args.listBool1,
        "listBool2": args.listBool2,
        "listBool3": args.listBool3,
        "listInt1": args.listInt1,
        "listInt2": args.listInt2,
        "listInt3": args.listInt3,
        "listNum1": args.listNum1,
        "listNum2": args.listNum2,
        "listNum3": args.listNum3,
        "listSecretBool1": args.listSecretBool1,
        "listSecretBool2": args.listSecretBool2,
        "listSecretBool3": args.listSecretBool3,
        "listSecretInt1": args.listSecretInt1,
        "listSecretInt2": args.listSecretInt2,
        "listSecretInt3": args.listSecretInt3,
        "listSecretNum1": args.listSecretNum1,
        "listSecretNum2": args.listSecretNum2,
        "listSecretNum3": args.listSecretNum3,
        "listSecretString1": args.listSecretString1,
        "listSecretString2": args.listSecretString2,
        "listSecretString3": args.listSecretString3,
        "listString1": args.listString1,
        "listString2": args.listString2,
        "listString3": args.listString3,
        "mapBool1": args.mapBool1,
        "mapBool2": args.mapBool2,
        "mapBool3": args.mapBool3,
        "mapInt1": args.mapInt1,
        "mapInt2": args.mapInt2,
        "mapInt3": args.mapInt3,
        "mapNum1": args.mapNum1,
        "mapNum2": args.mapNum2,
        "mapNum3": args.mapNum3,
        "mapSecretBool1": args.mapSecretBool1,
        "mapSecretBool2": args.mapSecretBool2,
        "mapSecretBool3": args.mapSecretBool3,
        "mapSecretInt1": args.mapSecretInt1,
        "mapSecretInt2": args.mapSecretInt2,
        "mapSecretInt3": args.mapSecretInt3,
        "mapSecretNum1": args.mapSecretNum1,
        "mapSecretNum2": args.mapSecretNum2,
        "mapSecretNum3": args.mapSecretNum3,
        "mapSecretString1": args.mapSecretString1,
        "mapSecretString2": args.mapSecretString2,
        "mapSecretString3": args.mapSecretString3,
        "mapString1": args.mapString1,
        "mapString2": args.mapString2,
        "mapString3": args.mapString3,
        "num1": args.num1,
        "num2": args.num2,
        "num3": args.num3,
        "objBool1": args.objBool1,
        "objBool2": args.objBool2,
        "objBool3": args.objBool3,
        "objInt1": args.objInt1,
        "objInt2": args.objInt2,
        "objInt3": args.objInt3,
        "objNum1": args.objNum1,
        "objNum2": args.objNum2,
        "objNum3": args.objNum3,
        "objSecretBool1": args.objSecretBool1,
        "objSecretBool2": args.objSecretBool2,
        "objSecretBool3": args.objSecretBool3,
        "objSecretInt1": args.objSecretInt1,
        "objSecretInt2": args.objSecretInt2,
        "objSecretInt3": args.objSecretInt3,
        "objSecretNum1": args.objSecretNum1,
        "objSecretNum2": args.objSecretNum2,
        "objSecretNum3": args.objSecretNum3,
        "objSecretString1": args.objSecretString1,
        "objSecretString2": args.objSecretString2,
        "objSecretString3": args.objSecretString3,
        "objString1": args.objString1,
        "objString2": args.objString2,
        "objString3": args.objString3,
        "secretBool1": args.secretBool1,
        "secretBool2": args.secretBool2,
        "secretBool3": args.secretBool3,
        "secretInt1": args.secretInt1,
        "secretInt2": args.secretInt2,
        "secretInt3": args.secretInt3,
        "secretNum1": args.secretNum1,
        "secretNum2": args.secretNum2,
        "secretNum3": args.secretNum3,
        "secretString1": args.secretString1,
        "secretString2": args.secretString2,
        "secretString3": args.secretString3,
        "string1": args.string1,
        "string2": args.string2,
        "string3": args.string3,
    }, opts);
}

export interface ToSecretOutputArgs {
    bool1?: khulnasoft.Input<boolean>;
    bool2?: khulnasoft.Input<boolean>;
    bool3?: khulnasoft.Input<boolean>;
    int1?: khulnasoft.Input<number>;
    int2?: khulnasoft.Input<number>;
    int3?: khulnasoft.Input<number>;
    listBool1?: khulnasoft.Input<khulnasoft.Input<boolean>[]>;
    listBool2?: khulnasoft.Input<khulnasoft.Input<boolean>[]>;
    listBool3?: khulnasoft.Input<khulnasoft.Input<boolean>[]>;
    listInt1?: khulnasoft.Input<khulnasoft.Input<number>[]>;
    listInt2?: khulnasoft.Input<khulnasoft.Input<number>[]>;
    listInt3?: khulnasoft.Input<khulnasoft.Input<number>[]>;
    listNum1?: khulnasoft.Input<khulnasoft.Input<number>[]>;
    listNum2?: khulnasoft.Input<khulnasoft.Input<number>[]>;
    listNum3?: khulnasoft.Input<khulnasoft.Input<number>[]>;
    listSecretBool1?: khulnasoft.Input<khulnasoft.Input<boolean>[]>;
    listSecretBool2?: khulnasoft.Input<khulnasoft.Input<boolean>[]>;
    listSecretBool3?: khulnasoft.Input<khulnasoft.Input<boolean>[]>;
    listSecretInt1?: khulnasoft.Input<khulnasoft.Input<number>[]>;
    listSecretInt2?: khulnasoft.Input<khulnasoft.Input<number>[]>;
    listSecretInt3?: khulnasoft.Input<khulnasoft.Input<number>[]>;
    listSecretNum1?: khulnasoft.Input<khulnasoft.Input<number>[]>;
    listSecretNum2?: khulnasoft.Input<khulnasoft.Input<number>[]>;
    listSecretNum3?: khulnasoft.Input<khulnasoft.Input<number>[]>;
    listSecretString1?: khulnasoft.Input<khulnasoft.Input<string>[]>;
    listSecretString2?: khulnasoft.Input<khulnasoft.Input<string>[]>;
    listSecretString3?: khulnasoft.Input<khulnasoft.Input<string>[]>;
    listString1?: khulnasoft.Input<khulnasoft.Input<string>[]>;
    listString2?: khulnasoft.Input<khulnasoft.Input<string>[]>;
    listString3?: khulnasoft.Input<khulnasoft.Input<string>[]>;
    mapBool1?: khulnasoft.Input<{[key: string]: khulnasoft.Input<boolean>}>;
    mapBool2?: khulnasoft.Input<{[key: string]: khulnasoft.Input<boolean>}>;
    mapBool3?: khulnasoft.Input<{[key: string]: khulnasoft.Input<boolean>}>;
    mapInt1?: khulnasoft.Input<{[key: string]: khulnasoft.Input<number>}>;
    mapInt2?: khulnasoft.Input<{[key: string]: khulnasoft.Input<number>}>;
    mapInt3?: khulnasoft.Input<{[key: string]: khulnasoft.Input<number>}>;
    mapNum1?: khulnasoft.Input<{[key: string]: khulnasoft.Input<number>}>;
    mapNum2?: khulnasoft.Input<{[key: string]: khulnasoft.Input<number>}>;
    mapNum3?: khulnasoft.Input<{[key: string]: khulnasoft.Input<number>}>;
    mapSecretBool1?: khulnasoft.Input<{[key: string]: khulnasoft.Input<boolean>}>;
    mapSecretBool2?: khulnasoft.Input<{[key: string]: khulnasoft.Input<boolean>}>;
    mapSecretBool3?: khulnasoft.Input<{[key: string]: khulnasoft.Input<boolean>}>;
    mapSecretInt1?: khulnasoft.Input<{[key: string]: khulnasoft.Input<number>}>;
    mapSecretInt2?: khulnasoft.Input<{[key: string]: khulnasoft.Input<number>}>;
    mapSecretInt3?: khulnasoft.Input<{[key: string]: khulnasoft.Input<number>}>;
    mapSecretNum1?: khulnasoft.Input<{[key: string]: khulnasoft.Input<number>}>;
    mapSecretNum2?: khulnasoft.Input<{[key: string]: khulnasoft.Input<number>}>;
    mapSecretNum3?: khulnasoft.Input<{[key: string]: khulnasoft.Input<number>}>;
    mapSecretString1?: khulnasoft.Input<{[key: string]: khulnasoft.Input<string>}>;
    mapSecretString2?: khulnasoft.Input<{[key: string]: khulnasoft.Input<string>}>;
    mapSecretString3?: khulnasoft.Input<{[key: string]: khulnasoft.Input<string>}>;
    mapString1?: khulnasoft.Input<{[key: string]: khulnasoft.Input<string>}>;
    mapString2?: khulnasoft.Input<{[key: string]: khulnasoft.Input<string>}>;
    mapString3?: khulnasoft.Input<{[key: string]: khulnasoft.Input<string>}>;
    num1?: khulnasoft.Input<number>;
    num2?: khulnasoft.Input<number>;
    num3?: khulnasoft.Input<number>;
    objBool1?: khulnasoft.Input<inputs.Tbool1Args>;
    objBool2?: khulnasoft.Input<inputs.Tbool2Args>;
    objBool3?: khulnasoft.Input<inputs.Tbool3Args>;
    objInt1?: khulnasoft.Input<inputs.Tint1Args>;
    objInt2?: khulnasoft.Input<inputs.Tint2Args>;
    objInt3?: khulnasoft.Input<inputs.Tint3Args>;
    objNum1?: khulnasoft.Input<inputs.Tnum1Args>;
    objNum2?: khulnasoft.Input<inputs.Tnum2Args>;
    objNum3?: khulnasoft.Input<inputs.Tnum3Args>;
    objSecretBool1?: khulnasoft.Input<inputs.TsecretBool1Args>;
    objSecretBool2?: khulnasoft.Input<inputs.TsecretBool2Args>;
    objSecretBool3?: khulnasoft.Input<inputs.TsecretBool3Args>;
    objSecretInt1?: khulnasoft.Input<inputs.TsecretInt1Args>;
    objSecretInt2?: khulnasoft.Input<inputs.TsecretInt2Args>;
    objSecretInt3?: khulnasoft.Input<inputs.TsecretInt3Args>;
    objSecretNum1?: khulnasoft.Input<inputs.TsecretNum1Args>;
    objSecretNum2?: khulnasoft.Input<inputs.TsecretNum2Args>;
    objSecretNum3?: khulnasoft.Input<inputs.TsecretNum3Args>;
    objSecretString1?: khulnasoft.Input<inputs.TsecretString1Args>;
    objSecretString2?: khulnasoft.Input<inputs.TsecretString2Args>;
    objSecretString3?: khulnasoft.Input<inputs.TsecretString3Args>;
    objString1?: khulnasoft.Input<inputs.Tstring1Args>;
    objString2?: khulnasoft.Input<inputs.Tstring2Args>;
    objString3?: khulnasoft.Input<inputs.Tstring3Args>;
    secretBool1?: khulnasoft.Input<boolean>;
    secretBool2?: khulnasoft.Input<boolean>;
    secretBool3?: khulnasoft.Input<boolean>;
    secretInt1?: khulnasoft.Input<number>;
    secretInt2?: khulnasoft.Input<number>;
    secretInt3?: khulnasoft.Input<number>;
    secretNum1?: khulnasoft.Input<number>;
    secretNum2?: khulnasoft.Input<number>;
    secretNum3?: khulnasoft.Input<number>;
    secretString1?: khulnasoft.Input<string>;
    secretString2?: khulnasoft.Input<string>;
    secretString3?: khulnasoft.Input<string>;
    string1?: khulnasoft.Input<string>;
    string2?: khulnasoft.Input<string>;
    string3?: khulnasoft.Input<string>;
}
