import * as khulnasoft from "@khulnasoft/khulnasoft";

function singleOrNone<T>(elements: khulnasoft.Input<T>[]): khulnasoft.Input<T> {
    if (elements.length != 1) {
        throw new Error("singleOrNone expected input list to have a single element");
    }
    return elements[0];
}

export const result = singleOrNone([1]);
