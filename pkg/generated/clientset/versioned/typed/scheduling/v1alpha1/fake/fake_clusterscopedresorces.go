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

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "sigs.k8s.io/scheduler-plugins/pkg/apis/scheduling/v1alpha1"
)

// FakeClusterScopedResorceses implements ClusterScopedResorcesInterface
type FakeClusterScopedResorceses struct {
	Fake *FakeSchedulingV1alpha1
	ns   string
}

var clusterscopedresorcesesResource = schema.GroupVersionResource{Group: "scheduling.sigs.k8s.io", Version: "v1alpha1", Resource: "clusterscopedresorceses"}

var clusterscopedresorcesesKind = schema.GroupVersionKind{Group: "scheduling.sigs.k8s.io", Version: "v1alpha1", Kind: "ClusterScopedResorces"}

// Get takes name of the clusterScopedResorces, and returns the corresponding clusterScopedResorces object, and an error if there is any.
func (c *FakeClusterScopedResorceses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterScopedResorces, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clusterscopedresorcesesResource, c.ns, name), &v1alpha1.ClusterScopedResorces{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterScopedResorces), err
}

// List takes label and field selectors, and returns the list of ClusterScopedResorceses that match those selectors.
func (c *FakeClusterScopedResorceses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterScopedResorcesList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clusterscopedresorcesesResource, clusterscopedresorcesesKind, c.ns, opts), &v1alpha1.ClusterScopedResorcesList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterScopedResorcesList{ListMeta: obj.(*v1alpha1.ClusterScopedResorcesList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterScopedResorcesList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterScopedResorceses.
func (c *FakeClusterScopedResorceses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clusterscopedresorcesesResource, c.ns, opts))

}

// Create takes the representation of a clusterScopedResorces and creates it.  Returns the server's representation of the clusterScopedResorces, and an error, if there is any.
func (c *FakeClusterScopedResorceses) Create(ctx context.Context, clusterScopedResorces *v1alpha1.ClusterScopedResorces, opts v1.CreateOptions) (result *v1alpha1.ClusterScopedResorces, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clusterscopedresorcesesResource, c.ns, clusterScopedResorces), &v1alpha1.ClusterScopedResorces{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterScopedResorces), err
}

// Update takes the representation of a clusterScopedResorces and updates it. Returns the server's representation of the clusterScopedResorces, and an error, if there is any.
func (c *FakeClusterScopedResorceses) Update(ctx context.Context, clusterScopedResorces *v1alpha1.ClusterScopedResorces, opts v1.UpdateOptions) (result *v1alpha1.ClusterScopedResorces, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clusterscopedresorcesesResource, c.ns, clusterScopedResorces), &v1alpha1.ClusterScopedResorces{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterScopedResorces), err
}

// Delete takes name of the clusterScopedResorces and deletes it. Returns an error if one occurs.
func (c *FakeClusterScopedResorceses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(clusterscopedresorcesesResource, c.ns, name), &v1alpha1.ClusterScopedResorces{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterScopedResorceses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clusterscopedresorcesesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterScopedResorcesList{})
	return err
}

// Patch applies the patch and returns the patched clusterScopedResorces.
func (c *FakeClusterScopedResorceses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterScopedResorces, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clusterscopedresorcesesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ClusterScopedResorces{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterScopedResorces), err
}