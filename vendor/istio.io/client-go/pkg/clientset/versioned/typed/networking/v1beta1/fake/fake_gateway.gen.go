// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1beta1 "istio.io/client-go/pkg/apis/networking/v1beta1"
	networkingv1beta1 "istio.io/client-go/pkg/applyconfiguration/networking/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGateways implements GatewayInterface
type FakeGateways struct {
	Fake *FakeNetworkingV1beta1
	ns   string
}

var gatewaysResource = v1beta1.SchemeGroupVersion.WithResource("gateways")

var gatewaysKind = v1beta1.SchemeGroupVersion.WithKind("Gateway")

// Get takes name of the gateway, and returns the corresponding gateway object, and an error if there is any.
func (c *FakeGateways) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Gateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(gatewaysResource, c.ns, name), &v1beta1.Gateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Gateway), err
}

// List takes label and field selectors, and returns the list of Gateways that match those selectors.
func (c *FakeGateways) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.GatewayList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(gatewaysResource, gatewaysKind, c.ns, opts), &v1beta1.GatewayList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.GatewayList{ListMeta: obj.(*v1beta1.GatewayList).ListMeta}
	for _, item := range obj.(*v1beta1.GatewayList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested gateways.
func (c *FakeGateways) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(gatewaysResource, c.ns, opts))

}

// Create takes the representation of a gateway and creates it.  Returns the server's representation of the gateway, and an error, if there is any.
func (c *FakeGateways) Create(ctx context.Context, gateway *v1beta1.Gateway, opts v1.CreateOptions) (result *v1beta1.Gateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(gatewaysResource, c.ns, gateway), &v1beta1.Gateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Gateway), err
}

// Update takes the representation of a gateway and updates it. Returns the server's representation of the gateway, and an error, if there is any.
func (c *FakeGateways) Update(ctx context.Context, gateway *v1beta1.Gateway, opts v1.UpdateOptions) (result *v1beta1.Gateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(gatewaysResource, c.ns, gateway), &v1beta1.Gateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Gateway), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGateways) UpdateStatus(ctx context.Context, gateway *v1beta1.Gateway, opts v1.UpdateOptions) (*v1beta1.Gateway, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(gatewaysResource, "status", c.ns, gateway), &v1beta1.Gateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Gateway), err
}

// Delete takes name of the gateway and deletes it. Returns an error if one occurs.
func (c *FakeGateways) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(gatewaysResource, c.ns, name, opts), &v1beta1.Gateway{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGateways) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(gatewaysResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.GatewayList{})
	return err
}

// Patch applies the patch and returns the patched gateway.
func (c *FakeGateways) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Gateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(gatewaysResource, c.ns, name, pt, data, subresources...), &v1beta1.Gateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Gateway), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied gateway.
func (c *FakeGateways) Apply(ctx context.Context, gateway *networkingv1beta1.GatewayApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.Gateway, err error) {
	if gateway == nil {
		return nil, fmt.Errorf("gateway provided to Apply must not be nil")
	}
	data, err := json.Marshal(gateway)
	if err != nil {
		return nil, err
	}
	name := gateway.Name
	if name == nil {
		return nil, fmt.Errorf("gateway.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(gatewaysResource, c.ns, *name, types.ApplyPatchType, data), &v1beta1.Gateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Gateway), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeGateways) ApplyStatus(ctx context.Context, gateway *networkingv1beta1.GatewayApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.Gateway, err error) {
	if gateway == nil {
		return nil, fmt.Errorf("gateway provided to Apply must not be nil")
	}
	data, err := json.Marshal(gateway)
	if err != nil {
		return nil, err
	}
	name := gateway.Name
	if name == nil {
		return nil, fmt.Errorf("gateway.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(gatewaysResource, c.ns, *name, types.ApplyPatchType, data, "status"), &v1beta1.Gateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Gateway), err
}
