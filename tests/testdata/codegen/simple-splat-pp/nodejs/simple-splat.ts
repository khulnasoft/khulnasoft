import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as splat from "@khulnasoft/splat";

const allKeys = splat.getSshKeys({});
const main = new splat.Server("main", {sshKeys: allKeys.then(allKeys => allKeys.sshKeys.map(__item => __item.name))});
