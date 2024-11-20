import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as subpackage from "@khulnasoft/subpackage";

// The resource name is based on the parameter value
const example = new subpackage.HelloWorld("example", {});
export const parameterValue = example.parameterValue;
