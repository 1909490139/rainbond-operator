// RAINBOND, Application Management Platform
// Copyright (C) 2014-2017 Goodrain Co., Ltd.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.
// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/GLYASAI/rainbond-operator/pkg/apis/rainbond/v1alpha1"
	scheme "github.com/GLYASAI/rainbond-operator/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RainbondsGetter has a method to return a RainbondInterface.
// A group's client should implement this interface.
type RainbondsGetter interface {
	Rainbonds(namespace string) RainbondInterface
}

// RainbondInterface has methods to work with Rainbond resources.
type RainbondInterface interface {
	Create(*v1alpha1.Rainbond) (*v1alpha1.Rainbond, error)
	Update(*v1alpha1.Rainbond) (*v1alpha1.Rainbond, error)
	UpdateStatus(*v1alpha1.Rainbond) (*v1alpha1.Rainbond, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Rainbond, error)
	List(opts v1.ListOptions) (*v1alpha1.RainbondList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Rainbond, err error)
	RainbondExpansion
}

// rainbonds implements RainbondInterface
type rainbonds struct {
	client rest.Interface
	ns     string
}

// newRainbonds returns a Rainbonds
func newRainbonds(c *RainbondV1alpha1Client, namespace string) *rainbonds {
	return &rainbonds{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the rainbond, and returns the corresponding rainbond object, and an error if there is any.
func (c *rainbonds) Get(name string, options v1.GetOptions) (result *v1alpha1.Rainbond, err error) {
	result = &v1alpha1.Rainbond{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("rainbonds").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Rainbonds that match those selectors.
func (c *rainbonds) List(opts v1.ListOptions) (result *v1alpha1.RainbondList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.RainbondList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("rainbonds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested rainbonds.
func (c *rainbonds) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("rainbonds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a rainbond and creates it.  Returns the server's representation of the rainbond, and an error, if there is any.
func (c *rainbonds) Create(rainbond *v1alpha1.Rainbond) (result *v1alpha1.Rainbond, err error) {
	result = &v1alpha1.Rainbond{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("rainbonds").
		Body(rainbond).
		Do().
		Into(result)
	return
}

// Update takes the representation of a rainbond and updates it. Returns the server's representation of the rainbond, and an error, if there is any.
func (c *rainbonds) Update(rainbond *v1alpha1.Rainbond) (result *v1alpha1.Rainbond, err error) {
	result = &v1alpha1.Rainbond{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("rainbonds").
		Name(rainbond.Name).
		Body(rainbond).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *rainbonds) UpdateStatus(rainbond *v1alpha1.Rainbond) (result *v1alpha1.Rainbond, err error) {
	result = &v1alpha1.Rainbond{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("rainbonds").
		Name(rainbond.Name).
		SubResource("status").
		Body(rainbond).
		Do().
		Into(result)
	return
}

// Delete takes name of the rainbond and deletes it. Returns an error if one occurs.
func (c *rainbonds) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("rainbonds").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *rainbonds) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("rainbonds").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched rainbond.
func (c *rainbonds) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Rainbond, err error) {
	result = &v1alpha1.Rainbond{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("rainbonds").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
