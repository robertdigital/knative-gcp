/*
Copyright 2019 Google LLC

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
	v1alpha1 "github.com/google/knative-gcp/pkg/apis/events/v1alpha1"
	scheme "github.com/google/knative-gcp/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GCSsGetter has a method to return a GCSInterface.
// A group's client should implement this interface.
type GCSsGetter interface {
	GCSs(namespace string) GCSInterface
}

// GCSInterface has methods to work with GCS resources.
type GCSInterface interface {
	Create(*v1alpha1.GCS) (*v1alpha1.GCS, error)
	Update(*v1alpha1.GCS) (*v1alpha1.GCS, error)
	UpdateStatus(*v1alpha1.GCS) (*v1alpha1.GCS, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GCS, error)
	List(opts v1.ListOptions) (*v1alpha1.GCSList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GCS, err error)
	GCSExpansion
}

// gCSs implements GCSInterface
type gCSs struct {
	client rest.Interface
	ns     string
}

// newGCSs returns a GCSs
func newGCSs(c *EventsV1alpha1Client, namespace string) *gCSs {
	return &gCSs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the gCS, and returns the corresponding gCS object, and an error if there is any.
func (c *gCSs) Get(name string, options v1.GetOptions) (result *v1alpha1.GCS, err error) {
	result = &v1alpha1.GCS{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gcss").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GCSs that match those selectors.
func (c *gCSs) List(opts v1.ListOptions) (result *v1alpha1.GCSList, err error) {
	result = &v1alpha1.GCSList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gcss").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested gCSs.
func (c *gCSs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("gcss").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a gCS and creates it.  Returns the server's representation of the gCS, and an error, if there is any.
func (c *gCSs) Create(gCS *v1alpha1.GCS) (result *v1alpha1.GCS, err error) {
	result = &v1alpha1.GCS{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("gcss").
		Body(gCS).
		Do().
		Into(result)
	return
}

// Update takes the representation of a gCS and updates it. Returns the server's representation of the gCS, and an error, if there is any.
func (c *gCSs) Update(gCS *v1alpha1.GCS) (result *v1alpha1.GCS, err error) {
	result = &v1alpha1.GCS{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gcss").
		Name(gCS.Name).
		Body(gCS).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *gCSs) UpdateStatus(gCS *v1alpha1.GCS) (result *v1alpha1.GCS, err error) {
	result = &v1alpha1.GCS{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gcss").
		Name(gCS.Name).
		SubResource("status").
		Body(gCS).
		Do().
		Into(result)
	return
}

// Delete takes name of the gCS and deletes it. Returns an error if one occurs.
func (c *gCSs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gcss").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *gCSs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gcss").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched gCS.
func (c *gCSs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GCS, err error) {
	result = &v1alpha1.GCS{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("gcss").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
