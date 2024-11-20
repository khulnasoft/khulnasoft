import * as khulnasoft from "@khulnasoft/khulnasoft";

const config = new khulnasoft.Config();
const value = config.require("value");
const tags = config.getObject<Record<string, string>>("tags") || {
    [`interpolated/${value}`]: "value",
};
