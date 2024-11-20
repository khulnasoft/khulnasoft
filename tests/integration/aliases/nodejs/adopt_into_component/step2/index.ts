// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as khulnasoft from "@khulnasoft/khulnasoft";

class Resource extends khulnasoft.ComponentResource {
    constructor(name: string, opts?: khulnasoft.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #2 - adopt a resource into a component.  The component author is the same as the component user, and changes
// the component to be able to adopt the resource that was previously defined separately...
class Component extends khulnasoft.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: khulnasoft.ComponentResourceOptions) {
        super("my:module:Component", name, {}, opts);
        // The resource creation was moved from top level to inside the component.
        this.resource = new Resource(`${name}-child`, {
            // With a new parent
            parent: this,
            // But with an alias provided based on knowing where the resource existing before - in this case at top
            // level.  We use an absolute URN instead of a relative `Alias` because we are referencing a fixed resource
            // that was in some arbitrary other location in the hierarchy prior to being adopted into this component.
            aliases: [khulnasoft.createUrn("res2", "my:module:Resource")],
        });
    }
}
// The creation of the component is unchanged.
const comp2 = new Component("comp2");

// Scenario 3: adopt this resource into a new parent.
class Component2 extends khulnasoft.ComponentResource {
    constructor(name: string, opts?: khulnasoft.ComponentResourceOptions) {
        super("my:module:Component2", name, {}, opts);
    }
}

// validate that "parent: undefined" means "i didn't have a parent previously"
new Component2("unparented", {
    aliases: [{ parent: khulnasoft.rootStackResource }],
    parent: comp2,
});


// Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix
// in the next step to be parented by this.  Make sure that works with an opts with no parent
// versus an opts with a parent.

class Component3 extends khulnasoft.ComponentResource {
    constructor(name: string, opts: khulnasoft.ComponentResourceOptions = {}) {
        super("my:module:Component3", name, {}, opts);
        new Component2(name + "-child", { aliases: [{ parent: opts.parent }], parent: this });
    }
}

new Component3("parentedbystack");
new Component3("parentedbycomponent", { parent: comp2 });

// Scenario 5: Allow multiple aliases to the same resource.
class Component4 extends khulnasoft.ComponentResource {
    constructor(name: string, opts: khulnasoft.ComponentResourceOptions = {}) {
        super("my:module:Component4", name, {}, {
            aliases: [
                { parent: khulnasoft.rootStackResource },
                { parent: khulnasoft.rootStackResource },
            ],
            ...opts,
        });
    }
}

new Component4("duplicateAliases", { parent: comp2 });