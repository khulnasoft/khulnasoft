package main

import (
	appsv1 "github.com/khulnasoft/khulnasoft-kubernetes/sdk/v3/go/kubernetes/apps/v1"
	corev1 "github.com/khulnasoft/khulnasoft-kubernetes/sdk/v3/go/kubernetes/core/v1"
	metav1 "github.com/khulnasoft/khulnasoft-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		_, err := appsv1.NewDeployment(ctx, "argocd_serverDeployment", &appsv1.DeploymentArgs{
			ApiVersion: khulnasoft.String("apps/v1"),
			Kind:       khulnasoft.String("Deployment"),
			Metadata: &metav1.ObjectMetaArgs{
				Name: khulnasoft.String("argocd-server"),
			},
			Spec: &appsv1.DeploymentSpecArgs{
				Selector: &metav1.LabelSelectorArgs{
					MatchLabels: khulnasoft.StringMap{
						"app": khulnasoft.String("server"),
					},
				},
				Replicas: khulnasoft.Int(1),
				Template: &corev1.PodTemplateSpecArgs{
					Metadata: &metav1.ObjectMetaArgs{
						Labels: khulnasoft.StringMap{
							"app": khulnasoft.String("server"),
						},
					},
					Spec: &corev1.PodSpecArgs{
						Containers: corev1.ContainerArray{
							&corev1.ContainerArgs{
								Name:  khulnasoft.String("nginx"),
								Image: khulnasoft.String("nginx"),
								ReadinessProbe: &corev1.ProbeArgs{
									HttpGet: &corev1.HTTPGetActionArgs{
										Port: khulnasoft.Any(8080),
									},
								},
							},
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
