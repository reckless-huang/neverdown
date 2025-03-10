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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/kzz45/neverdown/pkg/apis/openx/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AffinityLister helps list Affinities.
// All objects returned here must be treated as read-only.
type AffinityLister interface {
	// List lists all Affinities in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Affinity, err error)
	// Affinities returns an object that can list and get Affinities.
	Affinities(namespace string) AffinityNamespaceLister
	AffinityListerExpansion
}

// affinityLister implements the AffinityLister interface.
type affinityLister struct {
	indexer cache.Indexer
}

// NewAffinityLister returns a new AffinityLister.
func NewAffinityLister(indexer cache.Indexer) AffinityLister {
	return &affinityLister{indexer: indexer}
}

// List lists all Affinities in the indexer.
func (s *affinityLister) List(selector labels.Selector) (ret []*v1.Affinity, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Affinity))
	})
	return ret, err
}

// Affinities returns an object that can list and get Affinities.
func (s *affinityLister) Affinities(namespace string) AffinityNamespaceLister {
	return affinityNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AffinityNamespaceLister helps list and get Affinities.
// All objects returned here must be treated as read-only.
type AffinityNamespaceLister interface {
	// List lists all Affinities in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Affinity, err error)
	// Get retrieves the Affinity from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Affinity, error)
	AffinityNamespaceListerExpansion
}

// affinityNamespaceLister implements the AffinityNamespaceLister
// interface.
type affinityNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Affinities in the indexer for a given namespace.
func (s affinityNamespaceLister) List(selector labels.Selector) (ret []*v1.Affinity, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Affinity))
	})
	return ret, err
}

// Get retrieves the Affinity from the indexer for a given namespace and name.
func (s affinityNamespaceLister) Get(name string) (*v1.Affinity, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("affinity"), name)
	}
	return obj.(*v1.Affinity), nil
}
