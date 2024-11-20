import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as simple from "@khulnasoft/simple";
import * as simple_invoke from "@khulnasoft/simple-invoke";

const res = new simple.Resource("res", {value: true});
export const nonSecret = simple_invoke.secretInvokeOutput({
    value: "hello",
    secretResponse: false,
}).apply(invoke => invoke.response);
export const firstSecret = simple_invoke.secretInvokeOutput({
    value: "hello",
    secretResponse: res.value,
}).apply(invoke => invoke.response);
export const secondSecret = simple_invoke.secretInvokeOutput({
    value: khulnasoft.secret("goodbye"),
    secretResponse: false,
}).apply(invoke => invoke.response);
