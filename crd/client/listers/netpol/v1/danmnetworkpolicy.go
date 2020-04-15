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

package v1

import (
	v1 "github.com/nokia/danm-utils/crd/api/netpol/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DanmNetworkPolicyLister helps list DanmNetworkPolicies.
type DanmNetworkPolicyLister interface {
	// List lists all DanmNetworkPolicies in the indexer.
	List(selector labels.Selector) (ret []*v1.DanmNetworkPolicy, err error)
	// DanmNetworkPolicies returns an object that can list and get DanmNetworkPolicies.
	DanmNetworkPolicies(namespace string) DanmNetworkPolicyNamespaceLister
	DanmNetworkPolicyListerExpansion
}

// danmNetworkPolicyLister implements the DanmNetworkPolicyLister interface.
type danmNetworkPolicyLister struct {
	indexer cache.Indexer
}

// NewDanmNetworkPolicyLister returns a new DanmNetworkPolicyLister.
func NewDanmNetworkPolicyLister(indexer cache.Indexer) DanmNetworkPolicyLister {
	return &danmNetworkPolicyLister{indexer: indexer}
}

// List lists all DanmNetworkPolicies in the indexer.
func (s *danmNetworkPolicyLister) List(selector labels.Selector) (ret []*v1.DanmNetworkPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.DanmNetworkPolicy))
	})
	return ret, err
}

// DanmNetworkPolicies returns an object that can list and get DanmNetworkPolicies.
func (s *danmNetworkPolicyLister) DanmNetworkPolicies(namespace string) DanmNetworkPolicyNamespaceLister {
	return danmNetworkPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DanmNetworkPolicyNamespaceLister helps list and get DanmNetworkPolicies.
type DanmNetworkPolicyNamespaceLister interface {
	// List lists all DanmNetworkPolicies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.DanmNetworkPolicy, err error)
	// Get retrieves the DanmNetworkPolicy from the indexer for a given namespace and name.
	Get(name string) (*v1.DanmNetworkPolicy, error)
	DanmNetworkPolicyNamespaceListerExpansion
}

// danmNetworkPolicyNamespaceLister implements the DanmNetworkPolicyNamespaceLister
// interface.
type danmNetworkPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DanmNetworkPolicies in the indexer for a given namespace.
func (s danmNetworkPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1.DanmNetworkPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.DanmNetworkPolicy))
	})
	return ret, err
}

// Get retrieves the DanmNetworkPolicy from the indexer for a given namespace and name.
func (s danmNetworkPolicyNamespaceLister) Get(name string) (*v1.DanmNetworkPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("danmnetworkpolicy"), name)
	}
	return obj.(*v1.DanmNetworkPolicy), nil
}
