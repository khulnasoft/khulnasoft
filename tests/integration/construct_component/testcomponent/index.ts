import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as dynamic from "@khulnasoft/khulnasoft/dynamic";
import * as provider from "@khulnasoft/khulnasoft/provider";

let currentID = 0;

class Resource extends dynamic.Resource {
    constructor(name: string, echo: khulnasoft.Input<any>, opts?: khulnasoft.CustomResourceOptions) {
        const provider = {
            create: async (inputs: any) => ({
                id: (currentID++).toString(),
                outs: undefined,
            }),
        };

        super(provider, name, {echo}, opts);
    }
}

class Component extends khulnasoft.ComponentResource {
    public readonly echo: khulnasoft.Output<any>;
    public readonly childId: khulnasoft.Output<khulnasoft.ID>;
    public readonly secret: khulnasoft.Output<string>;

    constructor(name: string, echo: khulnasoft.Input<any>, secret: khulnasoft.Output<string>, opts?: khulnasoft.ComponentResourceOptions) {
        super("testcomponent:index:Component", name, {}, opts);

        this.echo = khulnasoft.output(echo);
        this.childId = (new Resource(`child-${name}`, echo, {parent: this})).id;
        this.secret = secret;

        this.registerOutputs({
            echo: this.echo,
            childId: this.childId,
            secret: this.secret,
        })
    }
}

class Provider implements provider.Provider {
    public readonly version = "0.0.1";

    construct(name: string, type: string, inputs: khulnasoft.Inputs,
              options: khulnasoft.ComponentResourceOptions): Promise<provider.ConstructResult> {
        if (type != "testcomponent:index:Component") {
            throw new Error(`unknown resource type ${type}`);
        }

        const config = new khulnasoft.Config();
        const secretKey = "secret";
        const fullSecretKey = `${config.name}:${secretKey}`;
        // use internal khulnasoft prop to check secretness
        const isSecret = (khulnasoft.runtime as any).isConfigSecret(fullSecretKey); 
        if (!isSecret) {
            throw new Error(`expected config with key "${secretKey}" to be secret.`)
        }
        const secret = config.requireSecret(secretKey);


        const component = new Component(name, inputs["echo"], secret, options);
        return Promise.resolve({
            urn: component.urn,
            state: {
                echo: component.echo,
                childId: component.childId,
                secret: secret,
            },
        });
    }
}

export function main(args: string[]) {
    return provider.main(new Provider(), args);
}

main(process.argv.slice(2));
