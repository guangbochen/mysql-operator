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

package fake

import (
	v1alpha1 "github.com/oracle/mysql-operator/pkg/apis/mysqlagent/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterAgents implements ClusterAgentInterface
type FakeClusterAgents struct {
	Fake *FakeMySQLV1alpha1
	ns   string
}

var clusteragentsResource = schema.GroupVersionResource{Group: "mysql.oracle.com", Version: "v1alpha1", Resource: "clusteragents"}

var clusteragentsKind = schema.GroupVersionKind{Group: "mysql.oracle.com", Version: "v1alpha1", Kind: "ClusterAgent"}

// Get takes name of the clusterAgent, and returns the corresponding clusterAgent object, and an error if there is any.
func (c *FakeClusterAgents) Get(name string, options v1.GetOptions) (result *v1alpha1.ClusterAgent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clusteragentsResource, c.ns, name), &v1alpha1.ClusterAgent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterAgent), err
}

// List takes label and field selectors, and returns the list of ClusterAgents that match those selectors.
func (c *FakeClusterAgents) List(opts v1.ListOptions) (result *v1alpha1.ClusterAgentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clusteragentsResource, clusteragentsKind, c.ns, opts), &v1alpha1.ClusterAgentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterAgentList{}
	for _, item := range obj.(*v1alpha1.ClusterAgentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterAgents.
func (c *FakeClusterAgents) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clusteragentsResource, c.ns, opts))

}

// Create takes the representation of a clusterAgent and creates it.  Returns the server's representation of the clusterAgent, and an error, if there is any.
func (c *FakeClusterAgents) Create(clusterAgent *v1alpha1.ClusterAgent) (result *v1alpha1.ClusterAgent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clusteragentsResource, c.ns, clusterAgent), &v1alpha1.ClusterAgent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterAgent), err
}

// Update takes the representation of a clusterAgent and updates it. Returns the server's representation of the clusterAgent, and an error, if there is any.
func (c *FakeClusterAgents) Update(clusterAgent *v1alpha1.ClusterAgent) (result *v1alpha1.ClusterAgent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clusteragentsResource, c.ns, clusterAgent), &v1alpha1.ClusterAgent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterAgent), err
}

// Delete takes name of the clusterAgent and deletes it. Returns an error if one occurs.
func (c *FakeClusterAgents) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(clusteragentsResource, c.ns, name), &v1alpha1.ClusterAgent{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterAgents) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clusteragentsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterAgentList{})
	return err
}

// Patch applies the patch and returns the patched clusterAgent.
func (c *FakeClusterAgents) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ClusterAgent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clusteragentsResource, c.ns, name, data, subresources...), &v1alpha1.ClusterAgent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterAgent), err
}
