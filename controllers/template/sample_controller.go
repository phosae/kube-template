/*
Copyright 2022.

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

package template

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	templatev1 "zeng.dev/kube-template/apis/template/v1"
	v1 "zeng.dev/kube-template/apis/template/v1"
	"zeng.dev/kube-template/client/versioned"
)

// SampleReconciler reconciles a Sample object
type SampleReconciler struct {
	GenClient versioned.Clientset
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=template.zeng.dev,resources=samples,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=template.zeng.dev,resources=samples/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=template.zeng.dev,resources=samples/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Sample object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.1/pkg/reconcile
func (r *SampleReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	lg := log.FromContext(ctx)

	lg.Info(fmt.Sprintf("start handle %v", req))
	var s v1.Sample
	if err := r.Client.Get(ctx, client.ObjectKey{
		Namespace: req.Namespace,
		Name:      req.Name,
	}, &s); err != nil {
		return ctrl.Result{}, err
	}
	lg.Info("got", "foo", s.Spec.Foo, "hello", s.Spec.Hello)

	var ss, err = r.GenClient.TemplateV1().Samples(req.Namespace).Get(
		ctx, req.Name, metav1.GetOptions{})
	if err != nil {
		return ctrl.Result{}, err
	}
	lg.Info("got from gentclient", "foo", ss.Spec.Foo, "hello", ss.Spec.Hello)
	lg.Info(fmt.Sprintf("endof handle %v", req))
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *SampleReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&templatev1.Sample{}).
		Complete(r)
}
