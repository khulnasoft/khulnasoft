import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as synthetic from "@khulnasoft/synthetic";

const rt = new synthetic.resourceproperties.Root("rt", {});
export const trivial = rt;
export const simple = rt.res1;
export const foo = rt.res1.apply(res1 => res1.obj1?.res2?.obj2);
export const complex = rt.res1.apply(res1 => res1.obj1?.res2?.obj2?.answer);
