// Copyright 2021 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	testjobv1 "github.com/kubeflow/common/test_job/apis/test_job/v1"
	versioned "github.com/kubeflow/common/test_job/client/clientset/versioned"
	internalinterfaces "github.com/kubeflow/common/test_job/client/informers/externalversions/internalinterfaces"
	v1 "github.com/kubeflow/common/test_job/client/listers/test_job/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TestJobInformer provides access to a shared informer and lister for
// TestJobs.
type TestJobInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.TestJobLister
}

type testJobInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTestJobInformer constructs a new informer for TestJob type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTestJobInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTestJobInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTestJobInformer constructs a new informer for TestJob type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTestJobInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeflowV1().TestJobs(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeflowV1().TestJobs(namespace).Watch(options)
			},
		},
		&testjobv1.TestJob{},
		resyncPeriod,
		indexers,
	)
}

func (f *testJobInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTestJobInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *testJobInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&testjobv1.TestJob{}, f.defaultInformer)
}

func (f *testJobInformer) Lister() v1.TestJobLister {
	return v1.NewTestJobLister(f.Informer().GetIndexer())
}
