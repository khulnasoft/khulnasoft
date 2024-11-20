import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as unknown from "@khulnasoft/unknown";

const data = unknown.index.getData({
    input: "hello",
});
const values = unknown.eks.moduleValues({});
export const content = data.content;
