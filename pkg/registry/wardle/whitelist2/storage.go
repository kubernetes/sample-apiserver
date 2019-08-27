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
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/sample-apiserver/pkg/apis/wardle"
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

func (m *Whitelist2Storage) New() runtime.Object {
	obj := &wardle.Whitelist2{}
	obj.Name = "ccc"
	return obj
}

func (m *Whitelist2Storage) Create(ctx context.Context, obj runtime.Object, createValidation rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	return obj, nil
}

//func (m *Whitelist2Storage) Get(ctx context.Context, name string, opts *metav1.GetOptions) (runtime.Object, error) {
//	return &v1.Whitelist2{}, nil
//}

func (m *Whitelist2Storage) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	return &wardle.Whitelist2{
		Name: name,
		Ips:  []string{"127.0.0.3", "127.0.0.4"},
		ID:   1,
	}, nil
}

//func (m *Whitelist2Storage) NewGetOptions() (runtime.Object, bool, string) {
//	return &wardle.Whitelist2{}, true, ""
//}

func (m *Whitelist2Storage) List(ctx context.Context, options *metainternalversion.ListOptions, extraOptions *list.ListOptions) (runtime.Object, error) {
	return &wardle.Whitelist2List{
		Items: []wardle.Whitelist2{
			{
				Name: "default",
				Ips:  []string{"127.0.0.1", "127.0.0.2"},
				ID:   1,
			},
			{
				Name: "localhost",
				Ips:  []string{"127.0.0.3", "127.0.0.4"},
				ID:   2,
			},
		},
	}, nil
}

func (s *Whitelist2Storage) NewList() runtime.Object {
	return &wardle.Whitelist2List{}
}

func (m *Whitelist2Storage) NamespaceScoped() bool {
	return false
}
