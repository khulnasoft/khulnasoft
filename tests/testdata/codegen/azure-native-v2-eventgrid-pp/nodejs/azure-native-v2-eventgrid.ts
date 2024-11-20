import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as azure_native from "@khulnasoft/azure-native";

const example = new azure_native.eventgrid.EventSubscription("example", {
    destination: {
        endpointType: "EventHub",
        resourceId: "example",
    },
    expirationTimeUtc: "example",
    scope: "example",
});
