import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as random from "@khulnasoft/random";

const randomPassword = new random.RandomPassword("randomPassword", {
    length: 16,
    special: true,
    overrideSpecial: "_%@",
});
export const password = randomPassword.result;
