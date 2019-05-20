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

package cluster

import (
	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	corelistersv1 "k8s.io/client-go/listers/core/v1"

	"github.com/oracle/mysql-operator/pkg/controllers/util"
)

// PodControlInterface defines the interface that the
// ClusterController uses to create, update, and delete mysql pods. It
// is implemented as an interface to enable testing.
type PodControlInterface interface {
	PatchPod(old *v1.Pod, new *v1.Pod) error
}

type realPodControl struct {
	client    kubernetes.Interface
	podLister corelistersv1.PodLister
}

// NewRealPodControl creates a concrete implementation of the
// PodControlInterface.
func NewRealPodControl(client kubernetes.Interface, podLister corelistersv1.PodLister) PodControlInterface {
	return &realPodControl{client: client, podLister: podLister}
}

func (rpc *realPodControl) PatchPod(old *v1.Pod, new *v1.Pod) error {
	_, err := util.PatchPod(rpc.client, old, new)
	return err
}

