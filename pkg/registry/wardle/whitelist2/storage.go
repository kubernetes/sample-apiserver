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

package whitelist2

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/sample-apiserver/pkg/apis/wardle/v1"
	"k8s.io/sample-apiserver/pkg/list"

	metainternalversion "k8s.io/apimachinery/pkg/apis/meta/internalversion"
)

type Whitelist2Storage struct {
}

func NewStorage() *Whitelist2Storage {
	return &Whitelist2Storage{}
}

func (s *Whitelist2Storage) Kind() string {
	return "Whitelist2"
}

func (s *Whitelist2Storage) New() runtime.Object {
	return &v1.Whitelist2{}
}

//func (m *Whitelist2Storage) Get(ctx context.Context, name string, opts *metav1.GetOptions) (runtime.Object, error) {
//	return &v1.Whitelist2{}, nil
//}

func (m *Whitelist2Storage) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	return &v1.Whitelist2{}, nil
}

//func (m *Whitelist2Storage) NewGetOptions() (runtime.Object, bool, string) {
//	return &v1.Whitelist2{}, true, ""
//}

func (m *Whitelist2Storage) List(ctx context.Context, options *metainternalversion.ListOptions, extraOptions *list.ListOptions) (runtime.Object, error) {
	return &v1.Whitelist2List{}, nil
}

func (s *Whitelist2Storage) NewList() runtime.Object {
	return &v1.Whitelist2List{}
}

func (m *Whitelist2Storage) NamespaceScoped() bool {
	return false
}
