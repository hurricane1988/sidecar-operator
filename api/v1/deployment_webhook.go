/*
Copyright 2023 hurricane1988.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var deploymentlog = logf.Log.WithName("deployment-resource")

func (r *Deployment) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-apps-hurricane-cn-v1-deployment,mutating=true,failurePolicy=fail,sideEffects=None,groups=apps.hurricane.cn,resources=deployments,verbs=create;update,versions=v1,name=mdeployment.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &Deployment{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Deployment) Default() {
	deploymentlog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
	switch r.Spec.Replicas {
	default:
		r.Spec.Replicas = new(int32)
		*r.Spec.Replicas = 1
	}
	// 判断注释是否有deployment.kubernetes.io/sidecar
	if r.Annotations["deployment.kubernetes.io/sidecar"] == "true" {
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:path=/validate-apps-hurricane-cn-v1-deployment,mutating=false,failurePolicy=fail,sideEffects=None,groups=apps.hurricane.cn,resources=deployments,verbs=create;update,versions=v1,name=vdeployment.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &Deployment{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Deployment) ValidateCreate() error {
	deploymentlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Deployment) ValidateUpdate(old runtime.Object) error {
	deploymentlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Deployment) ValidateDelete() error {
	deploymentlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
