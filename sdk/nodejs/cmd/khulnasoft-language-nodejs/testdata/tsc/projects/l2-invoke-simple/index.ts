import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as simple_invoke from "@khulnasoft/simple-invoke";

export const hello = simple_invoke.myInvokeOutput({
    value: "hello",
}).apply(invoke => invoke.result);
export const goodbye = simple_invoke.myInvokeOutput({
    value: "goodbye",
}).apply(invoke => invoke.result);
