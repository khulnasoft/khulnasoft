package main

import (
	appsv1 "github.com/khulnasoft/khulnasoft-kubernetes/sdk/v3/go/kubernetes/apps/v1"
	metav1 "github.com/khulnasoft/khulnasoft-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		// example
		_, err := appsv1.NewDeployment(ctx, "argocd_serverDeployment", &appsv1.DeploymentArgs{
			ApiVersion: khulnasoft.String("apps/v1"),
			Kind:       khulnasoft.String("Deployment"),
			Metadata: &metav1.ObjectMetaArgs{
				Labels: khulnasoft.StringMap{
					"app.kubernetes.io/component": khulnasoft.String("server"),
					"aws:region":                  khulnasoft.String("us-west-2"),
					"key%percent":                 khulnasoft.String("percent"),
					"key...ellipse":               khulnasoft.String("ellipse"),
					"key{bracket":                 khulnasoft.String("bracket"),
					"key}bracket":                 khulnasoft.String("bracket"),
					"key*asterix":                 khulnasoft.String("asterix"),
					"key?question":                khulnasoft.String("question"),
					"key,comma":                   khulnasoft.String("comma"),
					"key&&and":                    khulnasoft.String("and"),
					"key||or":                     khulnasoft.String("or"),
					"key!not":                     khulnasoft.String("not"),
					"key=>geq":                    khulnasoft.String("geq"),
					"key==eq":                     khulnasoft.String("equal"),
				},
				Name: khulnasoft.String("argocd-server"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
