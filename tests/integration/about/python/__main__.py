"""A Kubernetes Python Pulumi program"""

import khulnasoft
from khulnasoft_kubernetes.apps.v1 import Deployment, DeploymentSpecArgs
from khulnasoft_kubernetes.meta.v1 import LabelSelectorArgs, ObjectMetaArgs
from khulnasoft_kubernetes.core.v1 import ContainerArgs, PodSpecArgs, PodTemplateSpecArgs

app_labels = { "app": "nginx" }

deployment = Deployment(
    "nginx",
    spec=DeploymentSpecArgs(
        selector=LabelSelectorArgs(match_labels=app_labels),
        replicas=1,
        template=PodTemplateSpecArgs(
            metadata=ObjectMetaArgs(labels=app_labels),
            spec=PodSpecArgs(containers=[ContainerArgs(name="nginx", image="nginx")])
        ),
    ))

khulnasoft.export("name", deployment.metadata["name"])
