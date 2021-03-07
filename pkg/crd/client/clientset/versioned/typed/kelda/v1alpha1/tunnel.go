/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/kelda-inc/kelda/pkg/crd/apis/kelda/v1alpha1"
	scheme "github.com/kelda-inc/kelda/pkg/crd/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TunnelsGetter has a method to return a TunnelInterface.
// A group's client should implement this interface.
type TunnelsGetter interface {
	Tunnels(namespace string) TunnelInterface
}

// TunnelInterface has methods to work with Tunnel resources.
type TunnelInterface interface {
	Create(*v1alpha1.Tunnel) (*v1alpha1.Tunnel, error)
	Update(*v1alpha1.Tunnel) (*v1alpha1.Tunnel, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Tunnel, error)
	List(opts v1.ListOptions) (*v1alpha1.TunnelList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Tunnel, err error)
	TunnelExpansion
}

// tunnels implements TunnelInterface
type tunnels struct {
	client rest.Interface
	ns     string
}

// newTunnels returns a Tunnels
func newTunnels(c *KeldaV1alpha1Client, namespace string) *tunnels {
	return &tunnels{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the tunnel, and returns the corresponding tunnel object, and an error if there is any.
func (c *tunnels) Get(name string, options v1.GetOptions) (result *v1alpha1.Tunnel, err error) {
	result = &v1alpha1.Tunnel{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tunnels").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Tunnels that match those selectors.
func (c *tunnels) List(opts v1.ListOptions) (result *v1alpha1.TunnelList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.TunnelList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tunnels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested tunnels.
func (c *tunnels) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("tunnels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a tunnel and creates it.  Returns the server's representation of the tunnel, and an error, if there is any.
func (c *tunnels) Create(tunnel *v1alpha1.Tunnel) (result *v1alpha1.Tunnel, err error) {
	result = &v1alpha1.Tunnel{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("tunnels").
		Body(tunnel).
		Do().
		Into(result)
	return
}

// Update takes the representation of a tunnel and updates it. Returns the server's representation of the tunnel, and an error, if there is any.
func (c *tunnels) Update(tunnel *v1alpha1.Tunnel) (result *v1alpha1.Tunnel, err error) {
	result = &v1alpha1.Tunnel{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tunnels").
		Name(tunnel.Name).
		Body(tunnel).
		Do().
		Into(result)
	return
}

// Delete takes name of the tunnel and deletes it. Returns an error if one occurs.
func (c *tunnels) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tunnels").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *tunnels) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tunnels").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched tunnel.
func (c *tunnels) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Tunnel, err error) {
	result = &v1alpha1.Tunnel{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("tunnels").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
