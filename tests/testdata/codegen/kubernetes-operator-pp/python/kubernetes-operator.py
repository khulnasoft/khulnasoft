import khulnasoft
import khulnasoft_kubernetes as kubernetes

khulnasoft_kubernetes_operator_deployment = kubernetes.apps.v1.Deployment("khulnasoft_kubernetes_operatorDeployment",
    api_version="apps/v1",
    kind="Deployment",
    metadata={
        "name": "khulnasoft-kubernetes-operator",
    },
    spec={
        "replicas": 1,
        "selector": {
            "match_labels": {
                "name": "khulnasoft-kubernetes-operator",
            },
        },
        "template": {
            "metadata": {
                "labels": {
                    "name": "khulnasoft-kubernetes-operator",
                },
            },
            "spec": {
                "service_account_name": "khulnasoft-kubernetes-operator",
                "image_pull_secrets": [{
                    "name": "khulnasoft-kubernetes-operator",
                }],
                "containers": [{
                    "name": "khulnasoft-kubernetes-operator",
                    "image": "khulnasoft/khulnasoft-kubernetes-operator:v0.0.2",
                    "command": ["khulnasoft-kubernetes-operator"],
                    "args": ["--zap-level=debug"],
                    "image_pull_policy": "Always",
                    "env": [
                        {
                            "name": "WATCH_NAMESPACE",
                            "value_from": {
                                "field_ref": {
                                    "field_path": "metadata.namespace",
                                },
                            },
                        },
                        {
                            "name": "POD_NAME",
                            "value_from": {
                                "field_ref": {
                                    "field_path": "metadata.name",
                                },
                            },
                        },
                        {
                            "name": "OPERATOR_NAME",
                            "value": "khulnasoft-kubernetes-operator",
                        },
                    ],
                }],
            },
        },
    })
khulnasoft_kubernetes_operator_role = kubernetes.rbac.v1.Role("khulnasoft_kubernetes_operatorRole",
    api_version="rbac.authorization.k8s.io/v1",
    kind="Role",
    metadata={
        "creation_timestamp": None,
        "name": "khulnasoft-kubernetes-operator",
    },
    rules=[
        {
            "api_groups": [""],
            "resources": [
                "pods",
                "services",
                "services/finalizers",
                "endpoints",
                "persistentvolumeclaims",
                "events",
                "configmaps",
                "secrets",
            ],
            "verbs": [
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
            "api_groups": ["apps"],
            "resources": [
                "deployments",
                "daemonsets",
                "replicasets",
                "statefulsets",
            ],
            "verbs": [
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
            "api_groups": ["monitoring.coreos.com"],
            "resources": ["servicemonitors"],
            "verbs": [
                "get",
                "create",
            ],
        },
        {
            "api_groups": ["apps"],
            "resource_names": ["khulnasoft-kubernetes-operator"],
            "resources": ["deployments/finalizers"],
            "verbs": ["update"],
        },
        {
            "api_groups": [""],
            "resources": ["pods"],
            "verbs": ["get"],
        },
        {
            "api_groups": ["apps"],
            "resources": [
                "replicasets",
                "deployments",
            ],
            "verbs": ["get"],
        },
        {
            "api_groups": ["khulnasoft.com"],
            "resources": ["*"],
            "verbs": [
                "create",
                "delete",
                "get",
                "list",
                "patch",
                "update",
                "watch",
            ],
        },
    ])
khulnasoft_kubernetes_operator_role_binding = kubernetes.rbac.v1.RoleBinding("khulnasoft_kubernetes_operatorRoleBinding",
    kind="RoleBinding",
    api_version="rbac.authorization.k8s.io/v1",
    metadata={
        "name": "khulnasoft-kubernetes-operator",
    },
    subjects=[{
        "kind": "ServiceAccount",
        "name": "khulnasoft-kubernetes-operator",
    }],
    role_ref={
        "kind": "Role",
        "name": "khulnasoft-kubernetes-operator",
        "api_group": "rbac.authorization.k8s.io",
    })
khulnasoft_kubernetes_operator_service_account = kubernetes.core.v1.ServiceAccount("khulnasoft_kubernetes_operatorServiceAccount",
    api_version="v1",
    kind="ServiceAccount",
    metadata={
        "name": "khulnasoft-kubernetes-operator",
    })
