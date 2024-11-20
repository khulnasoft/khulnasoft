import * as khulnasoft from "@khulnasoft/khulnasoft";

const ref = new khulnasoft.StackReference("ref", {name: "organization/other/dev"});
export const plain = ref.getOutput("plain");
export const secret = ref.getOutput("secret");
