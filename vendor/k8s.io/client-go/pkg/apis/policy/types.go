/*
Copyright 2016 The Kubernetes Authors.

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

package policy

import (
	"k8s.io/client-go/pkg/api"
	"k8s.io/client-go/pkg/api/unversioned"
	"k8s.io/client-go/pkg/util/intstr"
)

// PodDisruptionBudgetSpec is a description of a PodDisruptionBudget.
type PodDisruptionBudgetSpec struct {
	// An eviction is allowed if at least "minAvailable" pods selected by
	// "selector" will still be available after the eviction, i.e. even in the
	// absence of the evicted pod.  So for example you can prevent all voluntary
	// evictions by specifying "100%".
	// +optional
	MinAvailable intstr.IntOrString `json:"minAvailable,omitempty"`

	// Label query over pods whose evictions are managed by the disruption
	// budget.
	// +optional
	Selector *unversioned.LabelSelector `json:"selector,omitempty"`
}

// PodDisruptionBudgetStatus represents information about the status of a
// PodDisruptionBudget. Status may trail the actual state of a system.
type PodDisruptionBudgetStatus struct {
	// Whether or not a disruption is currently allowed.
	PodDisruptionAllowed bool `json:"disruptionAllowed"`

	// current number of healthy pods
	CurrentHealthy int32 `json:"currentHealthy"`

	// minimum desired number of healthy pods
	DesiredHealthy int32 `json:"desiredHealthy"`

	// total number of pods counted by this disruption budget
	ExpectedPods int32 `json:"expectedPods"`
}

// +genclient=true

// PodDisruptionBudget is an object to define the max disruption that can be caused to a collection of pods
type PodDisruptionBudget struct {
	unversioned.TypeMeta `json:",inline"`
	// +optional
	api.ObjectMeta `json:"metadata,omitempty"`

	// Specification of the desired behavior of the PodDisruptionBudget.
	// +optional
	Spec PodDisruptionBudgetSpec `json:"spec,omitempty"`
	// Most recently observed status of the PodDisruptionBudget.
	// +optional
	Status PodDisruptionBudgetStatus `json:"status,omitempty"`
}

// PodDisruptionBudgetList is a collection of PodDisruptionBudgets.
type PodDisruptionBudgetList struct {
	unversioned.TypeMeta `json:",inline"`
	// +optional
	unversioned.ListMeta `json:"metadata,omitempty"`
	Items                []PodDisruptionBudget `json:"items"`
}

// Eviction evicts a pod from its node subject to certain policies and safety constraints.
// This is a subresource of Pod.  A request to cause such an eviction is
// created by POSTing to .../pods/<pod name>/evictions.
type Eviction struct {
	unversioned.TypeMeta `json:",inline"`

	// ObjectMeta describes the pod that is being evicted.
	// +optional
	api.ObjectMeta `json:"metadata,omitempty"`

	// DeleteOptions may be provided
	// +optional
	DeleteOptions *api.DeleteOptions `json:"deleteOptions,omitempty"`
}
