import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as kubernetes from "@khulnasoft/kubernetes";

const argocd_serverDeployment = new kubernetes.apps.v1.Deployment("argocd_serverDeployment", {
    apiVersion: "apps/v1",
    kind: "Deployment",
    metadata: {
        name: "argocd-server",
    },
    spec: {
        selector: {
            matchLabels: {
                app: "server",
            },
        },
        replicas: 1,
        template: {
            metadata: {
                labels: {
                    app: "server",
                },
            },
            spec: {
                containers: [{
                    name: "nginx",
                    image: "nginx",
                    ports: [{
                        containerPort: 80,
                    }],
                }],
            },
        },
    },
});
