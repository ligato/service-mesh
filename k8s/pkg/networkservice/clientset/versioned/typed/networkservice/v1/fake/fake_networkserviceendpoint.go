// Copyright (c) 2019 Cisco and/or its affiliates.
// Copyright (c) 2019 Red Hat Inc. and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	networkservicev1 "github.com/networkservicemesh/networkservicemesh/k8s/pkg/apis/networkservice/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNetworkServiceEndpoints implements NetworkServiceEndpointInterface
type FakeNetworkServiceEndpoints struct {
	Fake *FakeNetworkservicemeshV1
	ns   string
}

var networkserviceendpointsResource = schema.GroupVersionResource{Group: "networkservicemesh.io", Version: "v1", Resource: "networkserviceendpoints"}

var networkserviceendpointsKind = schema.GroupVersionKind{Group: "networkservicemesh.io", Version: "v1", Kind: "NetworkServiceEndpoint"}

// Get takes name of the networkServiceEndpoint, and returns the corresponding networkServiceEndpoint object, and an error if there is any.
func (c *FakeNetworkServiceEndpoints) Get(name string, options v1.GetOptions) (result *networkservicev1.NetworkServiceEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networkserviceendpointsResource, c.ns, name), &networkservicev1.NetworkServiceEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*networkservicev1.NetworkServiceEndpoint), err
}

// List takes label and field selectors, and returns the list of NetworkServiceEndpoints that match those selectors.
func (c *FakeNetworkServiceEndpoints) List(opts v1.ListOptions) (result *networkservicev1.NetworkServiceEndpointList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networkserviceendpointsResource, networkserviceendpointsKind, c.ns, opts), &networkservicev1.NetworkServiceEndpointList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &networkservicev1.NetworkServiceEndpointList{ListMeta: obj.(*networkservicev1.NetworkServiceEndpointList).ListMeta}
	for _, item := range obj.(*networkservicev1.NetworkServiceEndpointList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkServiceEndpoints.
func (c *FakeNetworkServiceEndpoints) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networkserviceendpointsResource, c.ns, opts))

}

// Create takes the representation of a networkServiceEndpoint and creates it.  Returns the server's representation of the networkServiceEndpoint, and an error, if there is any.
func (c *FakeNetworkServiceEndpoints) Create(networkServiceEndpoint *networkservicev1.NetworkServiceEndpoint) (result *networkservicev1.NetworkServiceEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networkserviceendpointsResource, c.ns, networkServiceEndpoint), &networkservicev1.NetworkServiceEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*networkservicev1.NetworkServiceEndpoint), err
}

// Update takes the representation of a networkServiceEndpoint and updates it. Returns the server's representation of the networkServiceEndpoint, and an error, if there is any.
func (c *FakeNetworkServiceEndpoints) Update(networkServiceEndpoint *networkservicev1.NetworkServiceEndpoint) (result *networkservicev1.NetworkServiceEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networkserviceendpointsResource, c.ns, networkServiceEndpoint), &networkservicev1.NetworkServiceEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*networkservicev1.NetworkServiceEndpoint), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkServiceEndpoints) UpdateStatus(networkServiceEndpoint *networkservicev1.NetworkServiceEndpoint) (*networkservicev1.NetworkServiceEndpoint, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(networkserviceendpointsResource, "status", c.ns, networkServiceEndpoint), &networkservicev1.NetworkServiceEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*networkservicev1.NetworkServiceEndpoint), err
}

// Delete takes name of the networkServiceEndpoint and deletes it. Returns an error if one occurs.
func (c *FakeNetworkServiceEndpoints) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(networkserviceendpointsResource, c.ns, name), &networkservicev1.NetworkServiceEndpoint{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkServiceEndpoints) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networkserviceendpointsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &networkservicev1.NetworkServiceEndpointList{})
	return err
}

// Patch applies the patch and returns the patched networkServiceEndpoint.
func (c *FakeNetworkServiceEndpoints) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *networkservicev1.NetworkServiceEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networkserviceendpointsResource, c.ns, name, pt, data, subresources...), &networkservicev1.NetworkServiceEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*networkservicev1.NetworkServiceEndpoint), err
}
