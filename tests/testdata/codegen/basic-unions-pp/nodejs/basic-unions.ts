import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as basic_unions from "@khulnasoft/basic-unions";

// properties field is bound to union case ServerPropertiesForReplica
const replica = new basic_unions.ExampleServer("replica", {properties: {
    createMode: "Replica",
    version: "0.1.0-dev",
}});
// properties field is bound to union case ServerPropertiesForRestore
const restore = new basic_unions.ExampleServer("restore", {properties: {
    createMode: "PointInTimeRestore",
    restorePointInTime: "example",
}});
