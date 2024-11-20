import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as fs from "fs";

const key = fs.readFileSync("key.pub", "utf8");
export const result = key;
