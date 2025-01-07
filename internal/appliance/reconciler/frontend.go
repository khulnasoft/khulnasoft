package reconciler

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	netv1 "k8s.io/api/networking/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/khulnasoft/khulnasoft/internal/appliance/config"
	"github.com/khulnasoft/khulnasoft/internal/k8s/resource/container"
	"github.com/khulnasoft/khulnasoft/internal/k8s/resource/deployment"
	"github.com/khulnasoft/khulnasoft/internal/k8s/resource/ingress"
	"github.com/khulnasoft/khulnasoft/internal/k8s/resource/pod"
	"github.com/khulnasoft/khulnasoft/internal/k8s/resource/role"
	"github.com/khulnasoft/khulnasoft/internal/k8s/resource/rolebinding"
	"github.com/khulnasoft/khulnasoft/internal/k8s/resource/service"
	"github.com/khulnasoft/khulnasoft/internal/k8s/resource/serviceaccount"
	"github.com/khulnasoft/khulnasoft/lib/errors"
	"github.com/khulnasoft/khulnasoft/lib/pointers"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
)

func (r *Reconciler) reconcileFrontend(ctx context.Context, sg *config.Khulnasoft, owner client.Object) error {
	if err := r.reconcileFrontendDeployment(ctx, sg, owner); err != nil {
		return errors.Wrap(err, "reconciling Deployment")
	}
	if err := r.reconcileFrontendService(ctx, sg, owner); err != nil {
		return errors.Wrap(err, "reconciling Service")
	}
	if err := r.reconcileFrontendServiceInternal(ctx, sg, owner); err != nil {
		return errors.Wrap(err, "reconciling Service (internal)")
	}
	if err := r.reconcileFrontendServiceAccount(ctx, sg, owner); err != nil {
		return errors.Wrap(err, "reconciling ServiceAccount")
	}
	if err := r.reconcileFrontendRole(ctx, sg, owner); err != nil {
		return errors.Wrap(err, "reconciling Role")
	}
	if err := r.reconcileFrontendRoleBinding(ctx, sg, owner); err != nil {
		return errors.Wrap(err, "reconciling RoleBinding")
	}
	if err := r.reconcileFrontendIngress(ctx, sg, owner); err != nil {
		return errors.Wrap(err, "reconciling Ingress")
	}
	return nil
}

func (r *Reconciler) reconcileFrontendDeployment(ctx context.Context, sg *config.Khulnasoft, owner client.Object) error {
	name := "frontend"
	cfg := sg.Spec.Frontend

	defaultImage := config.GetDefaultImage(sg, name)
	ctr := container.NewContainer(name, cfg, config.ContainerConfig{
		Image: defaultImage,
		Resources: &corev1.ResourceRequirements{
			Requests: corev1.ResourceList{
				corev1.ResourceCPU:              resource.MustParse("2"),
				corev1.ResourceMemory:           resource.MustParse("2G"),
				corev1.ResourceEphemeralStorage: resource.MustParse("4Gi"),
			},
			Limits: corev1.ResourceList{
				corev1.ResourceCPU:              resource.MustParse("2"),
				corev1.ResourceMemory:           resource.MustParse("4G"),
				corev1.ResourceEphemeralStorage: resource.MustParse("8Gi"),
			},
		},
	})

	ctr.Env = append(ctr.Env, frontendEnvVars(sg)...)
	ctr.Env = append(ctr.Env, dbAuthVars()...)
	ctr.Env = append(ctr.Env, container.EnvVarsRedis()...)
	ctr.Env = append(ctr.Env, container.EnvVarsOtel()...)

	ctr.Args = []string{"serve"}

	ctr.Ports = []corev1.ContainerPort{
		{Name: "http", ContainerPort: 3080},
		{Name: "http-internal", ContainerPort: 3090},
		{Name: "debug", ContainerPort: 6060},
	}

	ctr.LivenessProbe = &corev1.Probe{
		ProbeHandler: corev1.ProbeHandler{
			HTTPGet: &corev1.HTTPGetAction{
				Path: "/healthz",
				Port: intstr.FromString("debug"),
			},
		},
		InitialDelaySeconds: 300,
		TimeoutSeconds:      5,
	}
	ctr.ReadinessProbe = &corev1.Probe{
		ProbeHandler: corev1.ProbeHandler{
			HTTPGet: &corev1.HTTPGetAction{
				Path: "/ready",
				Port: intstr.FromString("debug"),
			},
		},
		PeriodSeconds:  5,
		TimeoutSeconds: 5,
	}
	ctr.VolumeMounts = []corev1.VolumeMount{
		{Name: "home-dir", MountPath: "/home/khulnasoft"},
	}

	template := pod.NewPodTemplate("sourcegraph-frontend", cfg)
	dbConnSpecs, err := r.getDBSecrets(ctx, sg)
	if err != nil {
		return err
	}
	dbConnHash, err := configHash(dbConnSpecs)
	if err != nil {
		return err
	}
	template.Template.ObjectMeta.Annotations["checksum/auth"] = dbConnHash

	redisConnSpecs, err := r.getRedisSecrets(ctx, sg)
	if err != nil {
		return err
	}
	redisConnHash, err := configHash(redisConnSpecs)
	if err != nil {
		return err
	}
	template.Template.ObjectMeta.Annotations["checksum/redis"] = redisConnHash

	template.Template.Spec.Containers = []corev1.Container{ctr}
	template.Template.Spec.Volumes = []corev1.Volume{pod.NewVolumeEmptyDir("home-dir")}
	template.Template.Spec.ServiceAccountName = "sourcegraph-frontend"

	if cfg.Migrator {
		migratorImage := config.GetDefaultImage(sg, "migrator")
		migratorCtr := container.NewContainer("migrator", cfg, config.ContainerConfig{
			Image: migratorImage,
			Resources: &corev1.ResourceRequirements{
				Requests: corev1.ResourceList{
					corev1.ResourceCPU:    resource.MustParse("100m"),
					corev1.ResourceMemory: resource.MustParse("50M"),
				},
				Limits: corev1.ResourceList{
					corev1.ResourceCPU:    resource.MustParse("500m"),
					corev1.ResourceMemory: resource.MustParse("100M"),
				},
			},
		})
		migratorCtr.Args = []string{"up"}
		migratorCtr.Env = append(migratorCtr.Env, frontendEnvVars(sg)...)
		migratorCtr.Env = append(migratorCtr.Env, dbAuthVars()...)
		template.Template.Spec.InitContainers = []corev1.Container{migratorCtr}
	}

	dep := deployment.NewDeployment("sourcegraph-frontend", sg.Namespace, sg.Spec.RequestedVersion)
	dep.Spec.Replicas = &cfg.Replicas
	dep.Spec.Strategy.RollingUpdate = &appsv1.RollingUpdateDeployment{
		MaxSurge:       pointers.Ptr(intstr.FromInt32(2)),
		MaxUnavailable: pointers.Ptr(intstr.FromInt32(0)),
	}
	dep.Spec.Template = template.Template

	ifChanged := struct {
		config.FrontendSpec
		DBConnSpecs
		RedisConnSpecs
	}{
		FrontendSpec:   cfg,
		DBConnSpecs:    dbConnSpecs,
		RedisConnSpecs: redisConnSpecs,
	}

	return reconcileObject(ctx, r, ifChanged, &dep, &appsv1.Deployment{}, sg, owner)
}

func (r *Reconciler) reconcileFrontendService(ctx context.Context, sg *config.Khulnasoft, owner client.Object) error {
	name := "sourcegraph-frontend"
	cfg := sg.Spec.Frontend
	logger := log.FromContext(ctx).WithValues("kind", "from ingress creation")
	namespacedName := types.NamespacedName{Namespace: sg.Namespace, Name: name}
	existingObj := &corev1.Service{}
	if err := r.Client.Get(ctx, namespacedName, existingObj); err != nil {
		// If we don't find an object, create one from the spec
		if kerrors.IsNotFound(err) {
			svc := service.NewService(name, sg.Namespace, cfg)
			svc.Spec.Ports = []corev1.ServicePort{
				{Name: "http", Port: 30080, TargetPort: intstr.FromString("http")},
			}
			svc.Spec.Selector = map[string]string{
				"app": "sourcegraph-appliance",
			}

			config.MarkObjectForAdoption(&svc)
			return reconcileObject(ctx, r, cfg, &svc, &corev1.Service{}, sg, owner)
		}
		logger.Error(err, "unexpected error getting object")
		return err
	}

	// If we found an object, we only want to change configmap-specified things,
	// and certain defaults such as the prometheus port.
	svcChanges := &corev1.Service{}
	svcChanges.SetAnnotations(map[string]string{
		"prometheus.io/port":            "6060",
		"sourcegraph.prometheus/scrape": "true",
	})
	config.MarkObjectForAdoption(svcChanges)
	newObj, err := MergeK8sObjects(existingObj, svcChanges)
	if err != nil {
		return errors.Wrap(err, "merging objects")
	}
	newSvc, ok := newObj.(*corev1.Service)
	if !ok {
		return errors.Wrap(err, "asserting type")
	}
	return reconcileObject(ctx, r, sg.Spec.Frontend, newSvc, &corev1.Service{}, sg, owner)
}

func (r *Reconciler) reconcileFrontendServiceInternal(ctx context.Context, sg *config.Khulnasoft, owner client.Object) error {
	cfg := sg.Spec.Frontend

	svc := service.NewService("sourcegraph-frontend-internal", sg.Namespace, nil)
	svc.Spec.Ports = []corev1.ServicePort{
		{Name: "http-internal", Port: 80, TargetPort: intstr.FromString("http-internal")},
	}
	svc.Spec.Selector = map[string]string{
		"app": "sourcegraph-frontend",
	}

	return reconcileObject(ctx, r, cfg, &svc, &corev1.Service{}, sg, owner)
}

func (r *Reconciler) reconcileFrontendRole(ctx context.Context, sg *config.Khulnasoft, owner client.Object) error {
	name := "sourcegraph-frontend"
	cfg := sg.Spec.Frontend

	role := role.NewRole(name, sg.Namespace)

	readVerbs := []string{"get", "list", "watch"}
	role.Rules = []rbacv1.PolicyRule{
		{
			APIGroups: []string{""},
			Resources: []string{"endpoints", "services"},
			Verbs:     readVerbs,
		},
		{
			APIGroups: []string{"apps"},
			Resources: []string{"statefulsets"},
			Verbs:     readVerbs,
		},
	}

	return reconcileObject(ctx, r, cfg, &role, &rbacv1.Role{}, sg, owner)
}

func (r *Reconciler) reconcileFrontendServiceAccount(ctx context.Context, sg *config.Khulnasoft, owner client.Object) error {
	cfg := sg.Spec.Frontend
	sa := serviceaccount.NewServiceAccount("sourcegraph-frontend", sg.Namespace, cfg)
	return reconcileObject(ctx, r, cfg, &sa, &corev1.ServiceAccount{}, sg, owner)
}

func (r *Reconciler) reconcileFrontendRoleBinding(ctx context.Context, sg *config.Khulnasoft, owner client.Object) error {
	name := "sourcegraph-frontend"
	binding := rolebinding.NewRoleBinding(name, sg.Namespace)
	binding.RoleRef = rbacv1.RoleRef{
		Kind: "Role",
		Name: name,
	}
	binding.Subjects = []rbacv1.Subject{
		{
			Kind:      "ServiceAccount",
			Name:      name,
			Namespace: sg.Namespace,
		},
	}
	return reconcileObject(ctx, r, sg.Spec.Frontend, &binding, &rbacv1.RoleBinding{}, sg, owner)
}

func (r *Reconciler) reconcileFrontendIngress(ctx context.Context, sg *config.Khulnasoft, owner client.Object) error {
	name := "sourcegraph-frontend"
	cfg := sg.Spec.Frontend
	logger := log.FromContext(ctx).WithValues("kind", "from ingress creation")
	namespacedName := types.NamespacedName{Namespace: sg.Namespace, Name: name}
	existingObj := &netv1.Ingress{}
	if err := r.Client.Get(ctx, namespacedName, existingObj); err != nil {
		// If we don't find an object, create one from the spec
		if kerrors.IsNotFound(err) {
			ingress := ingress.NewIngress(name, sg.Namespace)
			if cfg.Ingress == nil {
				return ensureObjectDeleted(ctx, r, owner, &ingress)
			}

			ingress.SetAnnotations(cfg.Ingress.Annotations)

			if cfg.Ingress.TLSSecret != "" {
				ingress.Spec.TLS = []netv1.IngressTLS{{
					Hosts:      []string{cfg.Ingress.Host},
					SecretName: cfg.Ingress.TLSSecret,
				}}
			}

			ingress.Spec.Rules = []netv1.IngressRule{{
				Host: cfg.Ingress.Host,
				IngressRuleValue: netv1.IngressRuleValue{
					HTTP: &netv1.HTTPIngressRuleValue{
						Paths: []netv1.HTTPIngressPath{{
							Path:     "/",
							PathType: pointers.Ptr(netv1.PathTypePrefix),
							Backend: netv1.IngressBackend{
								Service: &netv1.IngressServiceBackend{
									Name: name,
									Port: netv1.ServiceBackendPort{
										Number: 30080,
									},
								},
							},
						}},
					},
				},
			}}

			ingress.Spec.IngressClassName = cfg.Ingress.IngressClassName

			config.MarkObjectForAdoption(&ingress)
			return reconcileObject(ctx, r, sg.Spec.Frontend, &ingress, &netv1.Ingress{}, sg, owner)
		}
		logger.Error(err, "unexpected error getting object")
		return err
	}

	//Otherwise we found an object and only want to craft changes
	cfgIngress := ingress.NewIngress(name, sg.Namespace)
	cfgIngress.SetAnnotations(cfg.Ingress.Annotations)
	if cfg.Ingress.TLSSecret != "" {
		cfgIngress.Spec.TLS = []netv1.IngressTLS{{
			Hosts:      []string{cfg.Ingress.Host},
			SecretName: cfg.Ingress.TLSSecret,
		}}
	}
	cfgIngress.Spec.IngressClassName = cfg.Ingress.IngressClassName
	config.MarkObjectForAdoption(&cfgIngress)
	newObj, err := MergeK8sObjects(existingObj, &cfgIngress)
	if err != nil {
		return errors.Wrap(err, "merging objects")
	}
	newObjAsIngress, ok := newObj.(*netv1.Ingress)
	if !ok {
		return errors.Wrap(err, "asserting type")
	}
	return reconcileObject(ctx, r, sg.Spec.Frontend, newObjAsIngress, &netv1.Ingress{}, sg, owner)
}

func frontendEnvVars(sg *config.Khulnasoft) []corev1.EnvVar {
	vars := []corev1.EnvVar{
		{Name: "DEPLOY_TYPE", Value: "appliance"},
	}
	if !sg.Spec.Grafana.IsDisabled() {
		vars = append(vars, corev1.EnvVar{Name: "GRAFANA_SERVER_URL", Value: "http://grafana:30070"})
	}
	if !sg.Spec.Jaeger.IsDisabled() {
		vars = append(vars, corev1.EnvVar{Name: "JAEGER_SERVER_URL", Value: "http://jaeger-query:16686"})
	}
	if !sg.Spec.Prometheus.IsDisabled() {
		vars = append(vars, corev1.EnvVar{Name: "PROMETHEUS_URL", Value: "http://prometheus:30090"})
	}
	return vars
}

func dbAuthVars() []corev1.EnvVar {
	return []corev1.EnvVar{
		container.NewEnvVarSecretKeyRef("PGDATABASE", pgsqlSecretName, "database"),
		container.NewEnvVarSecretKeyRef("PGHOST", pgsqlSecretName, "host"),
		container.NewEnvVarSecretKeyRef("PGPASSWORD", pgsqlSecretName, "password"),
		container.NewEnvVarSecretKeyRef("PGPORT", pgsqlSecretName, "port"),
		container.NewEnvVarSecretKeyRef("PGUSER", pgsqlSecretName, "user"),
		container.NewEnvVarSecretKeyRef("CODEINTEL_PGDATABASE", codeIntelDBSecretName, "database"),
		container.NewEnvVarSecretKeyRef("CODEINTEL_PGHOST", codeIntelDBSecretName, "host"),
		container.NewEnvVarSecretKeyRef("CODEINTEL_PGPASSWORD", codeIntelDBSecretName, "password"),
		container.NewEnvVarSecretKeyRef("CODEINTEL_PGPORT", codeIntelDBSecretName, "port"),
		container.NewEnvVarSecretKeyRef("CODEINTEL_PGUSER", codeIntelDBSecretName, "user"),
		container.NewEnvVarSecretKeyRef("CODEINSIGHTS_PGDATABASE", codeInsightsDBSecretName, "database"),
		container.NewEnvVarSecretKeyRef("CODEINSIGHTS_PGHOST", codeInsightsDBSecretName, "host"),
		container.NewEnvVarSecretKeyRef("CODEINSIGHTS_PGPASSWORD", codeInsightsDBSecretName, "password"),
		container.NewEnvVarSecretKeyRef("CODEINSIGHTS_PGPORT", codeInsightsDBSecretName, "port"),
		container.NewEnvVarSecretKeyRef("CODEINSIGHTS_PGUSER", codeInsightsDBSecretName, "user"),
	}
}

// MergeK8sObjects merges a Kubernetes object that already exists within the cluster
// with an existing Kubernetes object definition.
func MergeK8sObjects(existingObj client.Object, newObject client.Object) (client.Object, error) {
	// Convert existing object to unstructured
	existingUnstructured, err := runtime.DefaultUnstructuredConverter.ToUnstructured(existingObj)
	if err != nil {
		return nil, errors.Newf("failed to convert existing object to unstructured: %w", err)
	}

	newUnstructured, err := runtime.DefaultUnstructuredConverter.ToUnstructured(newObject)
	if err != nil {
		return nil, errors.Newf("failed to convert new object to unstructured: %w", err)
	}

	// Merge the objects using strategic merge patch
	mergedUnstructured, err := strategicpatch.StrategicMergeMapPatch(existingUnstructured, newUnstructured, existingObj)
	if err != nil {
		return nil, errors.Newf("failed to merge objects: %w", err)
	}

	// Convert the merged unstructured object back to the original type
	mergedObj := existingObj.DeepCopyObject().(client.Object)
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(mergedUnstructured, mergedObj)
	if err != nil {
		return nil, errors.Newf("failed to convert merged object from unstructured: %w", err)
	}

	return mergedObj, nil
}
