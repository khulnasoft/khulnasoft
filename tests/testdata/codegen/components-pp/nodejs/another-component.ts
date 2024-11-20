import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as random from "@khulnasoft/random";

export class AnotherComponent extends khulnasoft.ComponentResource {
    constructor(name: string, opts?: khulnasoft.ComponentResourceOptions) {
        super("components:index:AnotherComponent", name, {}, opts);
        const firstPassword = new random.RandomPassword(`${name}-firstPassword`, {
            length: 16,
            special: true,
        }, {
            parent: this,
        });

        this.registerOutputs();
    }
}
