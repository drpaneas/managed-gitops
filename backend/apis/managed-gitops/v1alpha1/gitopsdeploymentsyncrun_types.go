/*
Copyright 2021.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GitOpsDeploymentSyncRunSpec defines the desired state of GitOpsDeploymentSyncRun
type GitOpsDeploymentSyncRunSpec struct {
	GitopsDeploymentName string `json:"gitopsDeploymentName"`
	RevisionID           string `json:"revisionID,omitempty"`
}

// GitOpsDeploymentSyncRunStatus defines the observed state of GitOpsDeploymentSyncRun
type GitOpsDeploymentSyncRunStatus struct {
	SyncStatus   string `json:"syncStatus,omitempty"`
	HealthStatus string `json:"health,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// GitOpsDeploymentSyncRun is the Schema for the gitopsdeploymentsyncruns API
type GitOpsDeploymentSyncRun struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GitOpsDeploymentSyncRunSpec   `json:"spec,omitempty"`
	Status GitOpsDeploymentSyncRunStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// GitOpsDeploymentSyncRunList contains a list of GitOpsDeploymentSyncRun
type GitOpsDeploymentSyncRunList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GitOpsDeploymentSyncRun `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GitOpsDeploymentSyncRun{}, &GitOpsDeploymentSyncRunList{})
}
