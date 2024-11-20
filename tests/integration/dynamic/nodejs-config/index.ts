// Copyright 2016-2023, Pulumi Corporation.

import * as khulnasoft from "@khulnasoft/khulnasoft";

class CustomResource extends khulnasoft.dynamic.Resource {
    public readonly authenticated!: khulnasoft.Output<string>;
    public readonly color!: khulnasoft.Output<string>;

    constructor(name: string, opts?: khulnasoft.ResourceOptions) {
        super(
            new DummyResourceProvider(),
            name,
            {
                authenticated: undefined,
                color: undefined
            },
            opts,
            "custom-provider",
            "CustomResource",
        );
    }
}

class DummyResourceProvider implements khulnasoft.dynamic.ResourceProvider {
    private password: string;
    private color: string;

    async configure(req: khulnasoft.dynamic.ConfigureRequest): Promise<any> {
        this.password = req.config.require("password");
        this.color = req.config.get("colors:banana") ?? "blue";
    }

    async create(props: any): Promise<khulnasoft.dynamic.CreateResult> {
        return {
            id: "resource-id",
            outs: {
                authenticated: this.password === "s3cret" ? "200" : "401",
                color: this.color,
            },
        };
    }
}

const resource = new CustomResource("resource-name");

export const authenticated = resource.authenticated;
export const color = resource.color;
