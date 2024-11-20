import * as khulnasoft from "@khulnasoft/khulnasoft";

const config = new khulnasoft.Config();

export const out = config.requireSecret("mysecret");
