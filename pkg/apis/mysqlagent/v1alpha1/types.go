// Copyright 2018 Oracle and/or its affiliates. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MinimumMySQLVersion is the minimum version of MySQL server supported by the
// MySQL Operator.
const MinimumMySQLVersion = "5.7.23"

// ClusterAgentSpec defines the attributes a user can specify when creating a cluster
type ClusterAgentSpec struct {
	// Version defines the MySQL Docker image version.
	Version string `json:"version"`
}

// ClusterAgentConditionType represents a valid condition of a Cluster.
type ClusterAgentConditionType string

const (
	// ClusterReady means the Cluster is able to service requests.
	ClusterReady ClusterAgentConditionType = "Ready"
)

// ClusterCondition describes the observed state of a Cluster at a certain point.
type ClusterAgentCondition struct {
	Type   ClusterAgentConditionType
	Status corev1.ConditionStatus
	// +optional
	LastTransitionTime metav1.Time
	// +optional
	Reason string
	// +optional
	Message string
}

// ClusterAgentStatus defines the current status of a MySQL cluster
// propagating useful information back to the cluster admin
type ClusterAgentStatus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	// +optional
	Conditions []ClusterAgentCondition
}

// +genclient
// +genclient:noStatus
// +resourceName=clusteragents
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Cluster represents a cluster spec and associated metadata
type ClusterAgent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   ClusterAgentSpec   `json:"spec"`
	Status ClusterAgentStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterAgentList is a placeholder type for a list of MySQL cluster agents
type ClusterAgentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ClusterAgent `json:"items"`
}
