import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as kubernetes from "@khulnasoft/kubernetes";

const khulnasoft_kubernetes_operatorDeployment = new kubernetes.apps.v1.Deployment("khulnasoft_kubernetes_operatorDeployment", {
    apiVersion: "apps/v1",
    kind: "Deployment",
    metadata: {
        name: "khulnasoft-kubernetes-operator",
    },
    spec: {
        replicas: 1,
        selector: {
            matchLabels: {
                name: "khulnasoft-kubernetes-operator",
            },
        },
        template: {
            metadata: {
                labels: {
                    name: "khulnasoft-kubernetes-operator",
                },
            },
            spec: {
                serviceAccountName: "khulnasoft-kubernetes-operator",
                imagePullSecrets: [{
                    name: "khulnasoft-kubernetes-operator",
                }],
                containers: [{
                    name: "khulnasoft-kubernetes-operator",
                    image: "khulnasoft/khulnasoft-kubernetes-operator:v0.0.2",
                    command: ["khulnasoft-kubernetes-operator"],
                    args: ["--zap-level=debug"],
                    imagePullPolicy: "Always",
                    env: [
                        {
                            name: "WATCH_NAMESPACE",
                            valueFrom: {
                                fieldRef: {
                                    fieldPath: "metadata.namespace",
                                },
                            },
                        },
                        {
                            name: "POD_NAME",
                            valueFrom: {
                                fieldRef: {
                                    fieldPath: "metadata.name",
                                },
                            },
                        },
                        {
                            name: "OPERATOR_NAME",
                            value: "khulnasoft-kubernetes-operator",
                        },
                    ],
                }],
            },
        },
    },
});
const khulnasoft_kubernetes_operatorRole = new kubernetes.rbac.v1.Role("khulnasoft_kubernetes_operatorRole", {
    apiVersion: "rbac.authorization.k8s.io/v1",
    kind: "Role",
    metadata: {
        creationTimestamp: undefined,
        name: "khulnasoft-kubernetes-operator",
    },
    rules: [
        {
            apiGroups: [""],
            resources: [
                "pods",
                "services",
                "services/finalizers",
                "endpoints",
                "persistentvolumeclaims",
                "events",
                "configmaps",
                "secrets",
            ],
            verbs: [
                "create",
                "delete",
                "get",
                "list",
                "patch",
                "update",
                "watch",
            ],
        },
        {
            apiGroups: ["apps"],
            resources: [
                "deployments",
                "daemonsets",
                "replicasets",
                "statefulsets",
            ],
            verbs: [
                "create",
                "delete",
                "get",
                "list",
                "patch",
                "update",
                "watch",
            ],
        },
        {
            apiGroups: ["monitoring.coreos.com"],
            resources: ["servicemonitors"],
            verbs: [
                "get",
                "create",
            ],
        },
        {
            apiGroups: ["apps"],
            resourceNames: ["khulnasoft-kubernetes-operator"],
            resources: ["deployments/finalizers"],
            verbs: ["update"],
        },
        {
            apiGroups: [""],
            resources: ["pods"],
            verbs: ["get"],
        },
        {
            apiGroups: ["apps"],
            resources: [
                "replicasets",
                "deployments",
            ],
            verbs: ["get"],
        },
        {
            apiGroups: ["khulnasoft.com"],
            resources: ["*"],
            verbs: [
                "create",
                "delete",
                "get",
                "list",
                "patch",
                "update",
                "watch",
            ],
        },
    ],
});
const khulnasoft_kubernetes_operatorRoleBinding = new kubernetes.rbac.v1.RoleBinding("khulnasoft_kubernetes_operatorRoleBinding", {
    kind: "RoleBinding",
    apiVersion: "rbac.authorization.k8s.io/v1",
    metadata: {
        name: "khulnasoft-kubernetes-operator",
    },
    subjects: [{
        kind: "ServiceAccount",
        name: "khulnasoft-kubernetes-operator",
    }],
    roleRef: {
        kind: "Role",
        name: "khulnasoft-kubernetes-operator",
        apiGroup: "rbac.authorization.k8s.io",
    },
});
const khulnasoft_kubernetes_operatorServiceAccount = new kubernetes.core.v1.ServiceAccount("khulnasoft_kubernetes_operatorServiceAccount", {
    apiVersion: "v1",
    kind: "ServiceAccount",
    metadata: {
        name: "khulnasoft-kubernetes-operator",
    },
});
