/*
Copyright 2025.

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

package controller

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// EventReconciler reconciles a Event object
type EventReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=core,resources=events,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=events/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=core,resources=events/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.21.0/pkg/reconcile
func (r *EventReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := logf.FromContext(ctx).WithName("k8s_events_logger")

	e := &corev1.Event{}
	if err := r.Get(ctx, req.NamespacedName, e); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	log.WithValues(
		"namespace", e.Namespace,
		"reason", e.Reason,
		"source", e.Source,
		"type", e.Type,
		"action", e.Action,
		"reporting_controller", e.ReportingController,
		"reporting_instance", e.ReportingInstance,
		"event_time", e.LastTimestamp.Time,
		"involved_object", e.InvolvedObject,
		"related", e.Related,
	).Info(e.Message)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *EventReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(
			&corev1.Event{},
			builder.WithPredicates(predicate.Funcs{
				CreateFunc: func(e event.TypedCreateEvent[client.Object]) bool {
					return !e.IsInInitialList
				},
				DeleteFunc: func(_ event.TypedDeleteEvent[client.Object]) bool {
					return false
				},
				UpdateFunc: func(e event.TypedUpdateEvent[client.Object]) bool {
					oldEvent, okOld := e.ObjectOld.(*corev1.Event)
					newEvent, okNew := e.ObjectNew.(*corev1.Event)
					if !okOld || !okNew {
						mgr.GetLogger().
							WithValues(
								"old_object", fmt.Sprintf("%T", e.ObjectOld),
								"new_object", fmt.Sprintf("%T", e.ObjectNew),
							).
							Error(
								fmt.Errorf("reconcile watcher: cant cast client object to corev1.Event"),
								"cant ast client object to corev1.Event",
							)
						return false
					}
					return !oldEvent.LastTimestamp.Equal(&newEvent.LastTimestamp)
				},
				GenericFunc: func(e event.TypedGenericEvent[client.Object]) bool {
					return false
				},
			}),
		).
		Named("event").
		Complete(r)
}
