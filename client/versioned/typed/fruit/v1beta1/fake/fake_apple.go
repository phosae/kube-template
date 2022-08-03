/*
Copyright 2022.

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
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1beta1 "zeng.dev/kube-template/apis/fruit/v1beta1"
)

// FakeApples implements AppleInterface
type FakeApples struct {
	Fake *FakeFruitV1beta1
	ns   string
}

var applesResource = schema.GroupVersionResource{Group: "fruit", Version: "v1beta1", Resource: "apples"}

var applesKind = schema.GroupVersionKind{Group: "fruit", Version: "v1beta1", Kind: "Apple"}

// Get takes name of the apple, and returns the corresponding apple object, and an error if there is any.
func (c *FakeApples) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Apple, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(applesResource, c.ns, name), &v1beta1.Apple{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Apple), err
}

// List takes label and field selectors, and returns the list of Apples that match those selectors.
func (c *FakeApples) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.AppleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(applesResource, applesKind, c.ns, opts), &v1beta1.AppleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.AppleList{ListMeta: obj.(*v1beta1.AppleList).ListMeta}
	for _, item := range obj.(*v1beta1.AppleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apples.
func (c *FakeApples) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(applesResource, c.ns, opts))

}

// Create takes the representation of a apple and creates it.  Returns the server's representation of the apple, and an error, if there is any.
func (c *FakeApples) Create(ctx context.Context, apple *v1beta1.Apple, opts v1.CreateOptions) (result *v1beta1.Apple, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(applesResource, c.ns, apple), &v1beta1.Apple{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Apple), err
}

// Update takes the representation of a apple and updates it. Returns the server's representation of the apple, and an error, if there is any.
func (c *FakeApples) Update(ctx context.Context, apple *v1beta1.Apple, opts v1.UpdateOptions) (result *v1beta1.Apple, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(applesResource, c.ns, apple), &v1beta1.Apple{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Apple), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApples) UpdateStatus(ctx context.Context, apple *v1beta1.Apple, opts v1.UpdateOptions) (*v1beta1.Apple, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(applesResource, "status", c.ns, apple), &v1beta1.Apple{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Apple), err
}

// Delete takes name of the apple and deletes it. Returns an error if one occurs.
func (c *FakeApples) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(applesResource, c.ns, name), &v1beta1.Apple{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApples) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(applesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.AppleList{})
	return err
}

// Patch applies the patch and returns the patched apple.
func (c *FakeApples) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Apple, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(applesResource, c.ns, name, pt, data, subresources...), &v1beta1.Apple{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Apple), err
}