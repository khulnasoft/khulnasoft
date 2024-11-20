package main

import (
	appsv1 "github.com/khulnasoft/khulnasoft-kubernetes/sdk/v3/go/kubernetes/apps/v1"
	corev1 "github.com/khulnasoft/khulnasoft-kubernetes/sdk/v3/go/kubernetes/core/v1"
	metav1 "github.com/khulnasoft/khulnasoft-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	rbacv1 "github.com/khulnasoft/khulnasoft-kubernetes/sdk/v3/go/kubernetes/rbac/v1"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"
)

func main() {
	khulnasoft.Run(func(ctx *khulnasoft.Context) error {
		_, err := appsv1.NewDeployment(ctx, "khulnasoft_kubernetes_operatorDeployment", &appsv1.DeploymentArgs{
			ApiVersion: khulnasoft.String("apps/v1"),
			Kind:       khulnasoft.String("Deployment"),
			Metadata: &metav1.ObjectMetaArgs{
				Name: khulnasoft.String("khulnasoft-kubernetes-operator"),
			},
			Spec: &appsv1.DeploymentSpecArgs{
				Replicas: khulnasoft.Int(1),
				Selector: &metav1.LabelSelectorArgs{
					MatchLabels: khulnasoft.StringMap{
						"name": khulnasoft.String("khulnasoft-kubernetes-operator"),
					},
				},
				Template: &corev1.PodTemplateSpecArgs{
					Metadata: &metav1.ObjectMetaArgs{
						Labels: khulnasoft.StringMap{
							"name": khulnasoft.String("khulnasoft-kubernetes-operator"),
						},
					},
					Spec: &corev1.PodSpecArgs{
						ServiceAccountName: khulnasoft.String("khulnasoft-kubernetes-operator"),
						ImagePullSecrets: corev1.LocalObjectReferenceArray{
							&corev1.LocalObjectReferenceArgs{
								Name: khulnasoft.String("khulnasoft-kubernetes-operator"),
							},
						},
						Containers: corev1.ContainerArray{
							&corev1.ContainerArgs{
								Name:  khulnasoft.String("khulnasoft-kubernetes-operator"),
								Image: khulnasoft.String("khulnasoft/khulnasoft-kubernetes-operator:v0.0.2"),
								Command: khulnasoft.StringArray{
									khulnasoft.String("khulnasoft-kubernetes-operator"),
								},
								Args: khulnasoft.StringArray{
									khulnasoft.String("--zap-level=debug"),
								},
								ImagePullPolicy: khulnasoft.String("Always"),
								Env: corev1.EnvVarArray{
									&corev1.EnvVarArgs{
										Name: khulnasoft.String("WATCH_NAMESPACE"),
										ValueFrom: &corev1.EnvVarSourceArgs{
											FieldRef: &corev1.ObjectFieldSelectorArgs{
												FieldPath: khulnasoft.String("metadata.namespace"),
											},
										},
									},
									&corev1.EnvVarArgs{
										Name: khulnasoft.String("POD_NAME"),
										ValueFrom: &corev1.EnvVarSourceArgs{
											FieldRef: &corev1.ObjectFieldSelectorArgs{
												FieldPath: khulnasoft.String("metadata.name"),
											},
										},
									},
									&corev1.EnvVarArgs{
										Name:  khulnasoft.String("OPERATOR_NAME"),
										Value: khulnasoft.String("khulnasoft-kubernetes-operator"),
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
		_, err = rbacv1.NewRole(ctx, "khulnasoft_kubernetes_operatorRole", &rbacv1.RoleArgs{
			ApiVersion: khulnasoft.String("rbac.authorization.k8s.io/v1"),
			Kind:       khulnasoft.String("Role"),
			Metadata: &metav1.ObjectMetaArgs{
				CreationTimestamp: nil,
				Name:              khulnasoft.String("khulnasoft-kubernetes-operator"),
			},
			Rules: rbacv1.PolicyRuleArray{
				&rbacv1.PolicyRuleArgs{
					ApiGroups: khulnasoft.StringArray{
						khulnasoft.String(""),
					},
					Resources: khulnasoft.StringArray{
						khulnasoft.String("pods"),
						khulnasoft.String("services"),
						khulnasoft.String("services/finalizers"),
						khulnasoft.String("endpoints"),
						khulnasoft.String("persistentvolumeclaims"),
						khulnasoft.String("events"),
						khulnasoft.String("configmaps"),
						khulnasoft.String("secrets"),
					},
					Verbs: khulnasoft.StringArray{
						khulnasoft.String("create"),
						khulnasoft.String("delete"),
						khulnasoft.String("get"),
						khulnasoft.String("list"),
						khulnasoft.String("patch"),
						khulnasoft.String("update"),
						khulnasoft.String("watch"),
					},
				},
				&rbacv1.PolicyRuleArgs{
					ApiGroups: khulnasoft.StringArray{
						khulnasoft.String("apps"),
					},
					Resources: khulnasoft.StringArray{
						khulnasoft.String("deployments"),
						khulnasoft.String("daemonsets"),
						khulnasoft.String("replicasets"),
						khulnasoft.String("statefulsets"),
					},
					Verbs: khulnasoft.StringArray{
						khulnasoft.String("create"),
						khulnasoft.String("delete"),
						khulnasoft.String("get"),
						khulnasoft.String("list"),
						khulnasoft.String("patch"),
						khulnasoft.String("update"),
						khulnasoft.String("watch"),
					},
				},
				&rbacv1.PolicyRuleArgs{
					ApiGroups: khulnasoft.StringArray{
						khulnasoft.String("monitoring.coreos.com"),
					},
					Resources: khulnasoft.StringArray{
						khulnasoft.String("servicemonitors"),
					},
					Verbs: khulnasoft.StringArray{
						khulnasoft.String("get"),
						khulnasoft.String("create"),
					},
				},
				&rbacv1.PolicyRuleArgs{
					ApiGroups: khulnasoft.StringArray{
						khulnasoft.String("apps"),
					},
					ResourceNames: khulnasoft.StringArray{
						khulnasoft.String("khulnasoft-kubernetes-operator"),
					},
					Resources: khulnasoft.StringArray{
						khulnasoft.String("deployments/finalizers"),
					},
					Verbs: khulnasoft.StringArray{
						khulnasoft.String("update"),
					},
				},
				&rbacv1.PolicyRuleArgs{
					ApiGroups: khulnasoft.StringArray{
						khulnasoft.String(""),
					},
					Resources: khulnasoft.StringArray{
						khulnasoft.String("pods"),
					},
					Verbs: khulnasoft.StringArray{
						khulnasoft.String("get"),
					},
				},
				&rbacv1.PolicyRuleArgs{
					ApiGroups: khulnasoft.StringArray{
						khulnasoft.String("apps"),
					},
					Resources: khulnasoft.StringArray{
						khulnasoft.String("replicasets"),
						khulnasoft.String("deployments"),
					},
					Verbs: khulnasoft.StringArray{
						khulnasoft.String("get"),
					},
				},
				&rbacv1.PolicyRuleArgs{
					ApiGroups: khulnasoft.StringArray{
						khulnasoft.String("khulnasoft.com"),
					},
					Resources: khulnasoft.StringArray{
						khulnasoft.String("*"),
					},
					Verbs: khulnasoft.StringArray{
						khulnasoft.String("create"),
						khulnasoft.String("delete"),
						khulnasoft.String("get"),
						khulnasoft.String("list"),
						khulnasoft.String("patch"),
						khulnasoft.String("update"),
						khulnasoft.String("watch"),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = rbacv1.NewRoleBinding(ctx, "khulnasoft_kubernetes_operatorRoleBinding", &rbacv1.RoleBindingArgs{
			Kind:       khulnasoft.String("RoleBinding"),
			ApiVersion: khulnasoft.String("rbac.authorization.k8s.io/v1"),
			Metadata: &metav1.ObjectMetaArgs{
				Name: khulnasoft.String("khulnasoft-kubernetes-operator"),
			},
			Subjects: rbacv1.SubjectArray{
				&rbacv1.SubjectArgs{
					Kind: khulnasoft.String("ServiceAccount"),
					Name: khulnasoft.String("khulnasoft-kubernetes-operator"),
				},
			},
			RoleRef: &rbacv1.RoleRefArgs{
				Kind:     khulnasoft.String("Role"),
				Name:     khulnasoft.String("khulnasoft-kubernetes-operator"),
				ApiGroup: khulnasoft.String("rbac.authorization.k8s.io"),
			},
		})
		if err != nil {
			return err
		}
		_, err = corev1.NewServiceAccount(ctx, "khulnasoft_kubernetes_operatorServiceAccount", &corev1.ServiceAccountArgs{
			ApiVersion: khulnasoft.String("v1"),
			Kind:       khulnasoft.String("ServiceAccount"),
			Metadata: &metav1.ObjectMetaArgs{
				Name: khulnasoft.String("khulnasoft-kubernetes-operator"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
