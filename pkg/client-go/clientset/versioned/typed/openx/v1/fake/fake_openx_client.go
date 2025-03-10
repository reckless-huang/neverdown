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
	v1 "github.com/kzz45/neverdown/pkg/client-go/clientset/versioned/typed/openx/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeOpenxV1 struct {
	*testing.Fake
}

func (c *FakeOpenxV1) Affinities(namespace string) v1.AffinityInterface {
	return &FakeAffinities{c, namespace}
}

func (c *FakeOpenxV1) AliyunAccessControls(namespace string) v1.AliyunAccessControlInterface {
	return &FakeAliyunAccessControls{c, namespace}
}

func (c *FakeOpenxV1) AliyunLoadBalancers(namespace string) v1.AliyunLoadBalancerInterface {
	return &FakeAliyunLoadBalancers{c, namespace}
}

func (c *FakeOpenxV1) Etcds(namespace string) v1.EtcdInterface {
	return &FakeEtcds{c, namespace}
}

func (c *FakeOpenxV1) Mysqls(namespace string) v1.MysqlInterface {
	return &FakeMysqls{c, namespace}
}

func (c *FakeOpenxV1) NodeSelectors(namespace string) v1.NodeSelectorInterface {
	return &FakeNodeSelectors{c, namespace}
}

func (c *FakeOpenxV1) Openxes(namespace string) v1.OpenxInterface {
	return &FakeOpenxes{c, namespace}
}

func (c *FakeOpenxV1) Redises(namespace string) v1.RedisInterface {
	return &FakeRedises{c, namespace}
}

func (c *FakeOpenxV1) Tolerations(namespace string) v1.TolerationInterface {
	return &FakeTolerations{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeOpenxV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
