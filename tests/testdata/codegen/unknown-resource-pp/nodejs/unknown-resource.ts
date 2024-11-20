import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as unknown from "@khulnasoft/unknown";

const provider = new khulnasoft.providers.Unknown("provider", {});
const main = new unknown.index.Main("main", {
    first: "hello",
    second: {
        foo: "bar",
    },
});
const fromModule: unknown.eks.Example[] = [];
for (const range = {value: 0}; range.value < 10; range.value++) {
    fromModule.push(new unknown.eks.Example(`fromModule-${range.value}`, {associatedMain: main.id}));
}
export const mainId = main.id;
export const values = fromModule.values.first;
