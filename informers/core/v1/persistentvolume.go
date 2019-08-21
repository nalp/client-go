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
	time "time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "github.com/nalp/client-go/informers/internalinterfaces"
	kubernetes "github.com/nalp/client-go/kubernetes"
	v1 "github.com/nalp/client-go/listers/core/v1"
	cache "github.com/nalp/client-go/tools/cache"
)

// PersistentVolumeInformer provides access to a shared informer and lister for
// PersistentVolumes.
type PersistentVolumeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.PersistentVolumeLister
}

type persistentVolumeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewPersistentVolumeInformer constructs a new informer for PersistentVolume type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPersistentVolumeInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPersistentVolumeInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredPersistentVolumeInformer constructs a new informer for PersistentVolume type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPersistentVolumeInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().PersistentVolumes().List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().PersistentVolumes().Watch(options)
			},
		},
		&corev1.PersistentVolume{},
		resyncPeriod,
		indexers,
	)
}

func (f *persistentVolumeInformer) defaultInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPersistentVolumeInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *persistentVolumeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&corev1.PersistentVolume{}, f.defaultInformer)
}

func (f *persistentVolumeInformer) Lister() v1.PersistentVolumeLister {
	return v1.NewPersistentVolumeLister(f.Informer().GetIndexer())
}
