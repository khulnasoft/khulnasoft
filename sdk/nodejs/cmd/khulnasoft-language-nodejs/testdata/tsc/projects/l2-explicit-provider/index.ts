import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as simple from "@khulnasoft/simple";

const prov = new simple.Provider("prov", {});
const res = new simple.Resource("res", {value: true}, {
    provider: prov,
});
