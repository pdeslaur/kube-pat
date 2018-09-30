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

package v1beta1

import (
	time "time"

	portaddresstranslationv1beta1 "github.com/pdeslaur/kube-pat/pkg/apis/portaddresstranslation/v1beta1"
	versioned "github.com/pdeslaur/kube-pat/pkg/client/clientset/versioned"
	internalinterfaces "github.com/pdeslaur/kube-pat/pkg/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/pdeslaur/kube-pat/pkg/client/listers/portaddresstranslation/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PortAddressTranslationInformer provides access to a shared informer and lister for
// PortAddressTranslations.
type PortAddressTranslationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.PortAddressTranslationLister
}

type portAddressTranslationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPortAddressTranslationInformer constructs a new informer for PortAddressTranslation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPortAddressTranslationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPortAddressTranslationInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPortAddressTranslationInformer constructs a new informer for PortAddressTranslation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPortAddressTranslationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.K8sV1beta1().PortAddressTranslations(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.K8sV1beta1().PortAddressTranslations(namespace).Watch(options)
			},
		},
		&portaddresstranslationv1beta1.PortAddressTranslation{},
		resyncPeriod,
		indexers,
	)
}

func (f *portAddressTranslationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPortAddressTranslationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *portAddressTranslationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&portaddresstranslationv1beta1.PortAddressTranslation{}, f.defaultInformer)
}

func (f *portAddressTranslationInformer) Lister() v1beta1.PortAddressTranslationLister {
	return v1beta1.NewPortAddressTranslationLister(f.Informer().GetIndexer())
}
