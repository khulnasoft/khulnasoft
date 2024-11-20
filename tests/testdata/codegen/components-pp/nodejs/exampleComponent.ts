import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as random from "@khulnasoft/random";
import { SimpleComponent } from "./simpleComponent";

interface ExampleComponentArgs {
    /**
     * A simple input
     */
    input: khulnasoft.Input<string>,
    /**
     * The main CIDR blocks for the VPC
     * It is a map of strings
     */
    cidrBlocks: khulnasoft.Input<Record<string, khulnasoft.Input<string>>>,
    /**
     * GitHub app parameters, see your github app. Ensure the key is the base64-encoded `.pem` file (the output of `base64 app.private-key.pem`, not the content of `private-key.pem`).
     */
    githubApp: {
        id?: khulnasoft.Input<string>,
        keyBase64?: khulnasoft.Input<string>,
        webhookSecret?: khulnasoft.Input<string>,
    },
    /**
     * A list of servers
     */
    servers: {
        name?: khulnasoft.Input<string>,
    }[],
    /**
     * A map between for zones
     */
    deploymentZones: Record<string, {
        zone?: khulnasoft.Input<string>,
    }>,
    ipAddress: khulnasoft.Input<number[]>,
}

export class ExampleComponent extends khulnasoft.ComponentResource {
    public result: khulnasoft.Output<string>;
    constructor(name: string, args: ExampleComponentArgs, opts?: khulnasoft.ComponentResourceOptions) {
        super("components:index:ExampleComponent", name, args, opts);
        const password = new random.RandomPassword(`${name}-password`, {
            length: 16,
            special: true,
            overrideSpecial: args.input,
        }, {
            parent: this,
        });

        const githubPassword = new random.RandomPassword(`${name}-githubPassword`, {
            length: 16,
            special: true,
            overrideSpecial: args.githubApp.webhookSecret,
        }, {
            parent: this,
        });

        // Example of iterating a list of objects
        const serverPasswords: random.RandomPassword[] = [];
        for (const range = {value: 0}; range.value < args.servers.length; range.value++) {
            serverPasswords.push(new random.RandomPassword(`${name}-serverPasswords-${range.value}`, {
                length: 16,
                special: true,
                overrideSpecial: args.servers[range.value].name,
            }, {
            parent: this,
        }));
        }

        // Example of iterating a map of objects
        const zonePasswords: random.RandomPassword[] = [];
        for (const range of Object.entries(args.deploymentZones).map(([k, v]) => ({key: k, value: v}))) {
            zonePasswords.push(new random.RandomPassword(`${name}-zonePasswords-${range.key}`, {
                length: 16,
                special: true,
                overrideSpecial: range.value.zone,
            }, {
            parent: this,
        }));
        }

        const simpleComponent = new SimpleComponent(`${name}-simpleComponent`, {
            parent: this,
        });

        this.result = password.result;
        this.registerOutputs({
            result: password.result,
        });
    }
}
