import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as random from "@khulnasoft/random";

class MyComponent extends khulnasoft.ComponentResource {
    constructor(name: string, opts: khulnasoft.ComponentResourceOptions) {
        super("pkg:index:MyComponent", name, {}, opts);

        new random.RandomPet("username", {}, { parent: this });
    }
}

const username = new random.RandomPet("username", {});

const component = new MyComponent("component", {
    // Add a dependency on the username resource to ensure it is created first. Depending on the order the
    // RandomPet resources are created the preview can generate different names for them. But our test expects
    // the first resource to be the renamed one.
    dependsOn: [username],
});