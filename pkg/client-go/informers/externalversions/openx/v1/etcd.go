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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	openxv1 "github.com/kzz45/neverdown/pkg/apis/openx/v1"
	versioned "github.com/kzz45/neverdown/pkg/client-go/clientset/versioned"
	internalinterfaces "github.com/kzz45/neverdown/pkg/client-go/informers/externalversions/internalinterfaces"
	v1 "github.com/kzz45/neverdown/pkg/client-go/listers/openx/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// EtcdInformer provides access to a shared informer and lister for
// Etcds.
type EtcdInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.EtcdLister
}

type etcdInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewEtcdInformer constructs a new informer for Etcd type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewEtcdInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredEtcdInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredEtcdInformer constructs a new informer for Etcd type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredEtcdInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OpenxV1().Etcds(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OpenxV1().Etcds(namespace).Watch(context.TODO(), options)
			},
		},
		&openxv1.Etcd{},
		resyncPeriod,
		indexers,
	)
}

func (f *etcdInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredEtcdInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *etcdInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&openxv1.Etcd{}, f.defaultInformer)
}

func (f *etcdInformer) Lister() v1.EtcdLister {
	return v1.NewEtcdLister(f.Informer().GetIndexer())
}
