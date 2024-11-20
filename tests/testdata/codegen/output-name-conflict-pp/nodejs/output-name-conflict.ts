import * as khulnasoft from "@khulnasoft/khulnasoft";

export = async () => {
    const config = new khulnasoft.Config();
    const cidrBlock = config.get("cidrBlock") || "Test config variable";
    return {
        cidrBlock: cidrBlock,
    };
}
