import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as simple_invoke from "@khulnasoft/simple-invoke";

const res = new simple_invoke.StringResource("res", {});
export const outputInput = simple_invoke.myInvokeOutput({
    value: res.text,
}).apply(invoke => invoke.result);
export const unit = simple_invoke.unitOutput({}).apply(invoke => invoke.result);
