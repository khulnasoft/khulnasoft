import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as primitive from "@khulnasoft/primitive";

const res = new primitive.Resource("res", {
    boolean: true,
    float: 3.14,
    integer: 42,
    string: "hello",
    numberArray: [
        -1,
        0,
        1,
    ],
    booleanMap: {
        t: true,
        f: false,
    },
});
