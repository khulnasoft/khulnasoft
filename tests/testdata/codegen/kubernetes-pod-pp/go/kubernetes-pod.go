package main

import (
	corev1 "github.com/khulnasoft/khulnasoft-kubernetes/sdk/v3/go/kubernetes/core/v1"
	metav1 "github.com/khulnasoft/khulnasoft-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		bar, err := corev1.NewPod(ctx, "bar", &corev1.PodArgs{
			ApiVersion: khulnasoft.String("v1"),
			Metadata: &metav1.ObjectMetaArgs{
				Namespace: khulnasoft.String("foo"),
				Name:      khulnasoft.String("bar"),
				Labels: khulnasoft.StringMap{
					"app.kubernetes.io/name":    khulnasoft.String("cilium-agent"),
					"app.kubernetes.io/part-of": khulnasoft.String("cilium"),
					"k8s-app":                   khulnasoft.String("cilium"),
				},
			},
			Spec: &corev1.PodSpecArgs{
				Containers: corev1.ContainerArray{
					&corev1.ContainerArgs{
						Name:  khulnasoft.String("nginx"),
						Image: khulnasoft.String("nginx:1.14-alpine"),
						Ports: corev1.ContainerPortArray{
							&corev1.ContainerPortArgs{
								ContainerPort: khulnasoft.Int(80),
							},
						},
						Resources: &corev1.ResourceRequirementsArgs{
							Limits: khulnasoft.StringMap{
								"memory": khulnasoft.String("20Mi"),
								"cpu":    khulnasoft.String("0.2"),
							},
						},
					},
					&corev1.ContainerArgs{
						Name:  khulnasoft.String("nginx2"),
						Image: khulnasoft.String("nginx:1.14-alpine"),
						Ports: corev1.ContainerPortArray{
							&corev1.ContainerPortArgs{
								ContainerPort: khulnasoft.Int(80),
							},
						},
						Resources: &corev1.ResourceRequirementsArgs{
							Limits: khulnasoft.StringMap{
								"memory": khulnasoft.String("20Mi"),
								"cpu":    khulnasoft.String("0.2"),
							},
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		// Test that we can assign from a constant without type errors
		_ := bar.Kind
		return nil
	})
}
