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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/kzz45/neverdown/pkg/client-go/clientset/versioned/typed/jingx/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeJingxV1 struct {
	*testing.Fake
}

func (c *FakeJingxV1) Events(namespace string) v1.EventInterface {
	return &FakeEvents{c, namespace}
}

func (c *FakeJingxV1) Projects(namespace string) v1.ProjectInterface {
	return &FakeProjects{c, namespace}
}

func (c *FakeJingxV1) Repositories(namespace string) v1.RepositoryInterface {
	return &FakeRepositories{c, namespace}
}

func (c *FakeJingxV1) Tags(namespace string) v1.TagInterface {
	return &FakeTags{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeJingxV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
