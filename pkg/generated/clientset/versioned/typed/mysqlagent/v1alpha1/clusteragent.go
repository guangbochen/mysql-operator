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
	scheme "github.com/oracle/mysql-operator/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterAgentsGetter has a method to return a ClusterAgentInterface.
// A group's client should implement this interface.
type ClusterAgentsGetter interface {
	ClusterAgents(namespace string) ClusterAgentInterface
}

// ClusterAgentInterface has methods to work with ClusterAgent resources.
type ClusterAgentInterface interface {
	Create(*v1alpha1.ClusterAgent) (*v1alpha1.ClusterAgent, error)
	Update(*v1alpha1.ClusterAgent) (*v1alpha1.ClusterAgent, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ClusterAgent, error)
	List(opts v1.ListOptions) (*v1alpha1.ClusterAgentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ClusterAgent, err error)
	ClusterAgentExpansion
}

// clusterAgents implements ClusterAgentInterface
type clusterAgents struct {
	client rest.Interface
	ns     string
}

// newClusterAgents returns a ClusterAgents
func newClusterAgents(c *MySQLV1alpha1Client, namespace string) *clusterAgents {
	return &clusterAgents{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the clusterAgent, and returns the corresponding clusterAgent object, and an error if there is any.
func (c *clusterAgents) Get(name string, options v1.GetOptions) (result *v1alpha1.ClusterAgent, err error) {
	result = &v1alpha1.ClusterAgent{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusteragents").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterAgents that match those selectors.
func (c *clusterAgents) List(opts v1.ListOptions) (result *v1alpha1.ClusterAgentList, err error) {
	result = &v1alpha1.ClusterAgentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusteragents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterAgents.
func (c *clusterAgents) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("clusteragents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a clusterAgent and creates it.  Returns the server's representation of the clusterAgent, and an error, if there is any.
func (c *clusterAgents) Create(clusterAgent *v1alpha1.ClusterAgent) (result *v1alpha1.ClusterAgent, err error) {
	result = &v1alpha1.ClusterAgent{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("clusteragents").
		Body(clusterAgent).
		Do().
		Into(result)
	return
}

// Update takes the representation of a clusterAgent and updates it. Returns the server's representation of the clusterAgent, and an error, if there is any.
func (c *clusterAgents) Update(clusterAgent *v1alpha1.ClusterAgent) (result *v1alpha1.ClusterAgent, err error) {
	result = &v1alpha1.ClusterAgent{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clusteragents").
		Name(clusterAgent.Name).
		Body(clusterAgent).
		Do().
		Into(result)
	return
}

// Delete takes name of the clusterAgent and deletes it. Returns an error if one occurs.
func (c *clusterAgents) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clusteragents").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterAgents) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clusteragents").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched clusterAgent.
func (c *clusterAgents) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ClusterAgent, err error) {
	result = &v1alpha1.ClusterAgent{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("clusteragents").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
