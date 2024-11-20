import * as khulnasoft from "@khulnasoft/khulnasoft";
import { R } from "./res";

export const withoutSecret = new R("withoutSecret", {
    prefix: khulnasoft.output("it's a secret to everybody")
});

export const withSecret = new R("withSecret", {
    prefix: khulnasoft.secret("it's a secret to everybody")
});

export const withSecretAdditional = new R("withSecretAdditional", {
    prefix: khulnasoft.output("it's a secret to everybody")
}, {
    additionalSecretOutputs: ["prefix"]
});
