// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	samplesv1 "github.com/openshift/cluster-samples-operator/pkg/apis/samples/v1"
	versioned "github.com/openshift/cluster-samples-operator/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/openshift/cluster-samples-operator/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/cluster-samples-operator/pkg/generated/listers/samples/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ConfigInformer provides access to a shared informer and lister for
// Configs.
type ConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ConfigLister
}

type configInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewConfigInformer constructs a new informer for Config type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewConfigInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredConfigInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredConfigInformer constructs a new informer for Config type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredConfigInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SamplesV1().Configs().List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SamplesV1().Configs().Watch(options)
			},
		},
		&samplesv1.Config{},
		resyncPeriod,
		indexers,
	)
}

func (f *configInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredConfigInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *configInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&samplesv1.Config{}, f.defaultInformer)
}

func (f *configInformer) Lister() v1.ConfigLister {
	return v1.NewConfigLister(f.Informer().GetIndexer())
}
