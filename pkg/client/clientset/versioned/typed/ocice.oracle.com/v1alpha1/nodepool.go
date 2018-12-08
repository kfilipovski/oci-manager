/*
Copyright 2018 Oracle and/or its affiliates. All rights reserved.

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
	v1alpha1 "github.com/oracle/oci-manager/pkg/apis/ocice.oracle.com/v1alpha1"
	scheme "github.com/oracle/oci-manager/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NodePoolsGetter has a method to return a NodePoolInterface.
// A group's client should implement this interface.
type NodePoolsGetter interface {
	NodePools(namespace string) NodePoolInterface
}

// NodePoolInterface has methods to work with NodePool resources.
type NodePoolInterface interface {
	Create(*v1alpha1.NodePool) (*v1alpha1.NodePool, error)
	Update(*v1alpha1.NodePool) (*v1alpha1.NodePool, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.NodePool, error)
	List(opts v1.ListOptions) (*v1alpha1.NodePoolList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NodePool, err error)
	NodePoolExpansion
}

// nodePools implements NodePoolInterface
type nodePools struct {
	client rest.Interface
	ns     string
}

// newNodePools returns a NodePools
func newNodePools(c *OciceV1alpha1Client, namespace string) *nodePools {
	return &nodePools{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the nodePool, and returns the corresponding nodePool object, and an error if there is any.
func (c *nodePools) Get(name string, options v1.GetOptions) (result *v1alpha1.NodePool, err error) {
	result = &v1alpha1.NodePool{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("nodepools").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NodePools that match those selectors.
func (c *nodePools) List(opts v1.ListOptions) (result *v1alpha1.NodePoolList, err error) {
	result = &v1alpha1.NodePoolList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("nodepools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested nodePools.
func (c *nodePools) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("nodepools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a nodePool and creates it.  Returns the server's representation of the nodePool, and an error, if there is any.
func (c *nodePools) Create(nodePool *v1alpha1.NodePool) (result *v1alpha1.NodePool, err error) {
	result = &v1alpha1.NodePool{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("nodepools").
		Body(nodePool).
		Do().
		Into(result)
	return
}

// Update takes the representation of a nodePool and updates it. Returns the server's representation of the nodePool, and an error, if there is any.
func (c *nodePools) Update(nodePool *v1alpha1.NodePool) (result *v1alpha1.NodePool, err error) {
	result = &v1alpha1.NodePool{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("nodepools").
		Name(nodePool.Name).
		Body(nodePool).
		Do().
		Into(result)
	return
}

// Delete takes name of the nodePool and deletes it. Returns an error if one occurs.
func (c *nodePools) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("nodepools").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *nodePools) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("nodepools").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched nodePool.
func (c *nodePools) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NodePool, err error) {
	result = &v1alpha1.NodePool{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("nodepools").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}