/*
Copyright 2017 The Kubernetes Authors.

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

package flunder

import (
	"context"

	metainternalversion "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/sample-apiserver/pkg/apis/wardle/v1beta1"
)

type FlunderStorage struct {
}

func NewStorage() *FlunderStorage {
	return &FlunderStorage{}
}

func (s *FlunderStorage) Kind() string {
	return "Flunder"
}

func (s *FlunderStorage) New() runtime.Object {
	return &v1beta1.Flunder{}
}

func (s *FlunderStorage) NewList() runtime.Object {
	return &v1beta1.FlunderList{}
}

func (m *FlunderStorage) Get(ctx context.Context, name string, opts *metav1.GetOptions) (runtime.Object, error) {
	return &v1beta1.Flunder{}, nil
}

func (m *FlunderStorage) List(ctx context.Context, options *metainternalversion.ListOptions) (runtime.Object, error) {
	return &v1beta1.FlunderList{}, nil
}

func (m *FlunderStorage) NamespaceScoped() bool {
	return false
}
