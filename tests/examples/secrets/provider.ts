import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as dynamic from "@khulnasoft/khulnasoft/dynamic";

class ReflectProvider implements dynamic.ResourceProvider {
    public check(olds: any, news: any) { return Promise.resolve({ inputs: news }); }
    public diff(id: khulnasoft.ID, olds: any, news: any) { return Promise.resolve({}); }
    public delete(id: khulnasoft.ID, props: any) { return Promise.resolve(); }
    public create(inputs: any) { return Promise.resolve({ id: "0", outs: inputs }); }
    public update(id: string, olds: any, news: any) { return Promise.resolve({ outs: news }); }
}

export class ReflectResource<T> extends dynamic.Resource {
    public readonly value!: khulnasoft.Output<T>;

    constructor(name: string, value: khulnasoft.Input<T>, opts?: khulnasoft.CustomResourceOptions) {
        super(new ReflectProvider(), name, {value: value}, opts);
    }
}

class DummyProvider implements dynamic.ResourceProvider {
    public check(olds: any, news: any) { return Promise.resolve({ inputs: news }); }
    public diff(id: khulnasoft.ID, olds: any, news: any) { return Promise.resolve({}); }
    public delete(id: khulnasoft.ID, props: any) { return Promise.resolve(); }
    public create(inputs: any) { return Promise.resolve({ id: "0", outs: {"value": "hello"} }); }
    public update(id: string, olds: any, news: any) { return Promise.resolve({ outs: news }); }
}

export class DummyResource extends dynamic.Resource {
    public readonly value!: khulnasoft.Output<string>;

    constructor(name: string, opts?: khulnasoft.CustomResourceOptions) {
        super(new DummyProvider(), name, {}, opts);
    }
}
