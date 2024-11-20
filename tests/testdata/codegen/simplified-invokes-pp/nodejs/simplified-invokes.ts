import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as std from "@khulnasoft/std";

const everyArg = std.AbsMultiArgs(10, 20, 30);
const onlyRequiredArgs = std.AbsMultiArgs(10);
const optionalArgs = std.AbsMultiArgs(10, undefined, 30);
const nestedUse = Promise.all([everyArg, std.AbsMultiArgs(42)]).then(([everyArg, invoke]) => std.AbsMultiArgs(everyArg, invoke));
export const result = nestedUse;
