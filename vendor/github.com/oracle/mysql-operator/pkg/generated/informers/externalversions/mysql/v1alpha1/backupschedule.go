// Copyright 2018 Oracle and/or its affiliates. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	time "time"

	mysql_v1alpha1 "github.com/oracle/mysql-operator/pkg/apis/mysql/v1alpha1"
	versioned "github.com/oracle/mysql-operator/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/oracle/mysql-operator/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/oracle/mysql-operator/pkg/generated/listers/mysql/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// BackupScheduleInformer provides access to a shared informer and lister for
// BackupSchedules.
type BackupScheduleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.BackupScheduleLister
}

type backupScheduleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewBackupScheduleInformer constructs a new informer for BackupSchedule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBackupScheduleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBackupScheduleInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredBackupScheduleInformer constructs a new informer for BackupSchedule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBackupScheduleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MySQLV1alpha1().BackupSchedules(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MySQLV1alpha1().BackupSchedules(namespace).Watch(options)
			},
		},
		&mysql_v1alpha1.BackupSchedule{},
		resyncPeriod,
		indexers,
	)
}

func (f *backupScheduleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBackupScheduleInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *backupScheduleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&mysql_v1alpha1.BackupSchedule{}, f.defaultInformer)
}

func (f *backupScheduleInformer) Lister() v1alpha1.BackupScheduleLister {
	return v1alpha1.NewBackupScheduleLister(f.Informer().GetIndexer())
}
