// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as dynamic from "@khulnasoft/khulnasoft/dynamic";

class OperatorProvider implements dynamic.ResourceProvider {
    private op: (l: number, r: number) => any;

    constructor(op: (l: number, r: number) => any) {
        this.op = op;
    }

    public check(olds: any, news: any) { return Promise.resolve({ inputs: news }); }
    public diff(id: khulnasoft.ID, olds: any, news: any) { return Promise.resolve({}); }
    public delete(id: khulnasoft.ID, props: any) { return Promise.resolve(); }
    public create(inputs: any) { return Promise.resolve({ id: "0", outs: this.op(Number(inputs.left), Number(inputs.right)) }); }
    public update(id: string, olds: any, news: any) { return Promise.resolve({ outs: this.op(Number(news.left), Number(news.right)) }); }
}

class DivProvider extends OperatorProvider {
    constructor() {
        super((left: number, right: number) => <any>{ quotient: Math.floor(left / right), remainder: left % right });
    }

    public async check(olds: any, news: any) {
        return {
            inputs: news,
            failures: news.right == 0 ? [ { property: "right", reason: "divisor must be non-zero" } ] : [],
        }
    }
}

class Add extends dynamic.Resource {
    public readonly sum: khulnasoft.Output<number>;

    private static provider = new OperatorProvider((left: number, right: number) => <any>{ sum: left + right });

    constructor(name: string, left: khulnasoft.Input<number>, right: khulnasoft.Input<number>) {
        super(Add.provider, name, {left: left, right: right, sum: undefined}, undefined);
    }
}

class Mul extends dynamic.Resource {
    public readonly product: khulnasoft.Output<number>;

    private static provider = new OperatorProvider((left: number, right: number) => <any>{ product: left * right });

    constructor(name: string, left: khulnasoft.Input<number>, right: khulnasoft.Input<number>) {
        super(Mul.provider, name, {left: left, right: right, product: undefined}, undefined);
    }
}

class Sub extends dynamic.Resource {
    public readonly difference: khulnasoft.Output<number>;

    private static provider = new OperatorProvider((left: number, right: number) => <any>{ difference: left - right });

    constructor(name: string, left: khulnasoft.Input<number>, right: khulnasoft.Input<number>) {
        super(Sub.provider, name, {left: left, right: right, difference: undefined}, undefined);
    }
}

class Div extends dynamic.Resource {
    public readonly quotient: khulnasoft.Output<number>;
    public readonly remainder: khulnasoft.Output<number>;

    private static provider = new DivProvider();

    constructor(name: string, left: khulnasoft.Input<number>, right: khulnasoft.Input<number>) {
        super(Div.provider, name, {left: left, right: right, quotient: undefined, remainder: undefined}, undefined);
    }
}

let config = new khulnasoft.Config("simple");
let w = Number(config.require("w")), x = Number(config.require("x")), y = Number(config.require("y"));
let sum = new Add("sum", x, y);
let square = new Mul("square", sum.sum, sum.sum);
let diff = new Sub("diff", square.product, w);
let divrem = new Div("divrem", diff.difference, sum.sum);
let result = new Add("result", divrem.quotient, divrem.remainder);
export let outputSum: khulnasoft.Output<number> = result.sum;
result.sum.apply(result => {
    console.log(`((x + y)^2 - w) / (x + y) + ((x + y)^2 - w) %% (x + y) = ${result}`);
});
