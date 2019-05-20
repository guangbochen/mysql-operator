// Copyright 2019 Oracle and/or its affiliates. All rights reserved.
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
	v1alpha1 "github.com/oracle/mysql-operator/pkg/apis/mysqlagent/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterAgentLister helps list ClusterAgents.
type ClusterAgentLister interface {
	// List lists all ClusterAgents in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterAgent, err error)
	// ClusterAgents returns an object that can list and get ClusterAgents.
	ClusterAgents(namespace string) ClusterAgentNamespaceLister
	ClusterAgentListerExpansion
}

// clusterAgentLister implements the ClusterAgentLister interface.
type clusterAgentLister struct {
	indexer cache.Indexer
}

// NewClusterAgentLister returns a new ClusterAgentLister.
func NewClusterAgentLister(indexer cache.Indexer) ClusterAgentLister {
	return &clusterAgentLister{indexer: indexer}
}

// List lists all ClusterAgents in the indexer.
func (s *clusterAgentLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterAgent, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterAgent))
	})
	return ret, err
}

// ClusterAgents returns an object that can list and get ClusterAgents.
func (s *clusterAgentLister) ClusterAgents(namespace string) ClusterAgentNamespaceLister {
	return clusterAgentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ClusterAgentNamespaceLister helps list and get ClusterAgents.
type ClusterAgentNamespaceLister interface {
	// List lists all ClusterAgents in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterAgent, err error)
	// Get retrieves the ClusterAgent from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ClusterAgent, error)
	ClusterAgentNamespaceListerExpansion
}

// clusterAgentNamespaceLister implements the ClusterAgentNamespaceLister
// interface.
type clusterAgentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ClusterAgents in the indexer for a given namespace.
func (s clusterAgentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterAgent, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterAgent))
	})
	return ret, err
}

// Get retrieves the ClusterAgent from the indexer for a given namespace and name.
func (s clusterAgentNamespaceLister) Get(name string) (*v1alpha1.ClusterAgent, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("clusteragent"), name)
	}
	return obj.(*v1alpha1.ClusterAgent), nil
}
