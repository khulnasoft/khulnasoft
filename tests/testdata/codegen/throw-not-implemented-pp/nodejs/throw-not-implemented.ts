import * as khulnasoft from "@khulnasoft/khulnasoft";

function notImplemented(message: string) {
    throw new Error(message);
}

export const result = notImplemented("expression here is not implemented yet");
