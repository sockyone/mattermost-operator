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
	v1alpha1 "github.com/oracle/mysql-operator/pkg/apis/mysql/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BackupLister helps list Backups.
type BackupLister interface {
	// List lists all Backups in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Backup, err error)
	// Backups returns an object that can list and get Backups.
	Backups(namespace string) BackupNamespaceLister
	BackupListerExpansion
}

// backupLister implements the BackupLister interface.
type backupLister struct {
	indexer cache.Indexer
}

// NewBackupLister returns a new BackupLister.
func NewBackupLister(indexer cache.Indexer) BackupLister {
	return &backupLister{indexer: indexer}
}

// List lists all Backups in the indexer.
func (s *backupLister) List(selector labels.Selector) (ret []*v1alpha1.Backup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Backup))
	})
	return ret, err
}

// Backups returns an object that can list and get Backups.
func (s *backupLister) Backups(namespace string) BackupNamespaceLister {
	return backupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BackupNamespaceLister helps list and get Backups.
type BackupNamespaceLister interface {
	// List lists all Backups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Backup, err error)
	// Get retrieves the Backup from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Backup, error)
	BackupNamespaceListerExpansion
}

// backupNamespaceLister implements the BackupNamespaceLister
// interface.
type backupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Backups in the indexer for a given namespace.
func (s backupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Backup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Backup))
	})
	return ret, err
}

// Get retrieves the Backup from the indexer for a given namespace and name.
func (s backupNamespaceLister) Get(name string) (*v1alpha1.Backup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("backup"), name)
	}
	return obj.(*v1alpha1.Backup), nil
}
