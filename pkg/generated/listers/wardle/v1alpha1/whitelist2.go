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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "k8s.io/sample-apiserver/pkg/apis/wardle/v1alpha1"
)

// Whitelist2Lister helps list Whitelist2s.
type Whitelist2Lister interface {
	// List lists all Whitelist2s in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Whitelist2, err error)
	// Whitelist2s returns an object that can list and get Whitelist2s.
	Whitelist2s(namespace string) Whitelist2NamespaceLister
	Whitelist2ListerExpansion
}

// whitelist2Lister implements the Whitelist2Lister interface.
type whitelist2Lister struct {
	indexer cache.Indexer
}

// NewWhitelist2Lister returns a new Whitelist2Lister.
func NewWhitelist2Lister(indexer cache.Indexer) Whitelist2Lister {
	return &whitelist2Lister{indexer: indexer}
}

// List lists all Whitelist2s in the indexer.
func (s *whitelist2Lister) List(selector labels.Selector) (ret []*v1alpha1.Whitelist2, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Whitelist2))
	})
	return ret, err
}

// Whitelist2s returns an object that can list and get Whitelist2s.
func (s *whitelist2Lister) Whitelist2s(namespace string) Whitelist2NamespaceLister {
	return whitelist2NamespaceLister{indexer: s.indexer, namespace: namespace}
}

// Whitelist2NamespaceLister helps list and get Whitelist2s.
type Whitelist2NamespaceLister interface {
	// List lists all Whitelist2s in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Whitelist2, err error)
	// Get retrieves the Whitelist2 from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Whitelist2, error)
	Whitelist2NamespaceListerExpansion
}

// whitelist2NamespaceLister implements the Whitelist2NamespaceLister
// interface.
type whitelist2NamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Whitelist2s in the indexer for a given namespace.
func (s whitelist2NamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Whitelist2, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Whitelist2))
	})
	return ret, err
}

// Get retrieves the Whitelist2 from the indexer for a given namespace and name.
func (s whitelist2NamespaceLister) Get(name string) (*v1alpha1.Whitelist2, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("whitelist2"), name)
	}
	return obj.(*v1alpha1.Whitelist2), nil
}
