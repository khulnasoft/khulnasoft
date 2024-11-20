import * as khulnasoft from "@khulnasoft/khulnasoft";

class Random extends khulnasoft.Resource {
  result!: khulnasoft.Output<string | undefined>;

  constructor(name: string, length: number, opts?: khulnasoft.ResourceOptions) {
    const inputs: any = {};
    inputs["length"] = length;
    super("testprovider:index:Random", name, true, inputs, opts);
  }
}

class RandomProvider extends khulnasoft.ProviderResource {
  constructor(name: string, opts?: khulnasoft.ResourceOptions) {
    super("testprovider", name, {}, opts);
  }
}

const r = new Random("default", 10, {
  pluginDownloadURL: "get.example.test",
});
export const defaultProvider = r.result;

const provider = new RandomProvider("explicit", {
  pluginDownloadURL: "get.khulnasoft.test/providers",
});

new Random("explicit", 8, { provider: provider });
