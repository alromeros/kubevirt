/*
Copyright 2021 The KubeVirt Authors.

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

package fake

import (
	monitoringv1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAlertmanagers implements AlertmanagerInterface
type FakeAlertmanagers struct {
	Fake *FakeMonitoringV1
	ns   string
}

var alertmanagersResource = schema.GroupVersionResource{Group: "monitoring.coreos.com", Version: "v1", Resource: "alertmanagers"}

var alertmanagersKind = schema.GroupVersionKind{Group: "monitoring.coreos.com", Version: "v1", Kind: "Alertmanager"}

// Get takes name of the alertmanager, and returns the corresponding alertmanager object, and an error if there is any.
func (c *FakeAlertmanagers) Get(name string, options v1.GetOptions) (result *monitoringv1.Alertmanager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(alertmanagersResource, c.ns, name), &monitoringv1.Alertmanager{})

	if obj == nil {
		return nil, err
	}
	return obj.(*monitoringv1.Alertmanager), err
}

// List takes label and field selectors, and returns the list of Alertmanagers that match those selectors.
func (c *FakeAlertmanagers) List(opts v1.ListOptions) (result *monitoringv1.AlertmanagerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(alertmanagersResource, alertmanagersKind, c.ns, opts), &monitoringv1.AlertmanagerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &monitoringv1.AlertmanagerList{ListMeta: obj.(*monitoringv1.AlertmanagerList).ListMeta}
	for _, item := range obj.(*monitoringv1.AlertmanagerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested alertmanagers.
func (c *FakeAlertmanagers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(alertmanagersResource, c.ns, opts))

}

// Create takes the representation of a alertmanager and creates it.  Returns the server's representation of the alertmanager, and an error, if there is any.
func (c *FakeAlertmanagers) Create(alertmanager *monitoringv1.Alertmanager) (result *monitoringv1.Alertmanager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(alertmanagersResource, c.ns, alertmanager), &monitoringv1.Alertmanager{})

	if obj == nil {
		return nil, err
	}
	return obj.(*monitoringv1.Alertmanager), err
}

// Update takes the representation of a alertmanager and updates it. Returns the server's representation of the alertmanager, and an error, if there is any.
func (c *FakeAlertmanagers) Update(alertmanager *monitoringv1.Alertmanager) (result *monitoringv1.Alertmanager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(alertmanagersResource, c.ns, alertmanager), &monitoringv1.Alertmanager{})

	if obj == nil {
		return nil, err
	}
	return obj.(*monitoringv1.Alertmanager), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAlertmanagers) UpdateStatus(alertmanager *monitoringv1.Alertmanager) (*monitoringv1.Alertmanager, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(alertmanagersResource, "status", c.ns, alertmanager), &monitoringv1.Alertmanager{})

	if obj == nil {
		return nil, err
	}
	return obj.(*monitoringv1.Alertmanager), err
}

// Delete takes name of the alertmanager and deletes it. Returns an error if one occurs.
func (c *FakeAlertmanagers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(alertmanagersResource, c.ns, name), &monitoringv1.Alertmanager{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAlertmanagers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(alertmanagersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &monitoringv1.AlertmanagerList{})
	return err
}

// Patch applies the patch and returns the patched alertmanager.
func (c *FakeAlertmanagers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *monitoringv1.Alertmanager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(alertmanagersResource, c.ns, name, pt, data, subresources...), &monitoringv1.Alertmanager{})

	if obj == nil {
		return nil, err
	}
	return obj.(*monitoringv1.Alertmanager), err
}
