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

package v1beta1

import (
	v1beta1 "k8s.io/api/admissionregistration/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"github.com/nalp/client-go/tools/cache"
)

// MutatingWebhookConfigurationLister helps list MutatingWebhookConfigurations.
type MutatingWebhookConfigurationLister interface {
	// List lists all MutatingWebhookConfigurations in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.MutatingWebhookConfiguration, err error)
	// Get retrieves the MutatingWebhookConfiguration from the index for a given name.
	Get(name string) (*v1beta1.MutatingWebhookConfiguration, error)
	MutatingWebhookConfigurationListerExpansion
}

// mutatingWebhookConfigurationLister implements the MutatingWebhookConfigurationLister interface.
type mutatingWebhookConfigurationLister struct {
	indexer cache.Indexer
}

// NewMutatingWebhookConfigurationLister returns a new MutatingWebhookConfigurationLister.
func NewMutatingWebhookConfigurationLister(indexer cache.Indexer) MutatingWebhookConfigurationLister {
	return &mutatingWebhookConfigurationLister{indexer: indexer}
}

// List lists all MutatingWebhookConfigurations in the indexer.
func (s *mutatingWebhookConfigurationLister) List(selector labels.Selector) (ret []*v1beta1.MutatingWebhookConfiguration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.MutatingWebhookConfiguration))
	})
	return ret, err
}

// Get retrieves the MutatingWebhookConfiguration from the index for a given name.
func (s *mutatingWebhookConfigurationLister) Get(name string) (*v1beta1.MutatingWebhookConfiguration, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("mutatingwebhookconfiguration"), name)
	}
	return obj.(*v1beta1.MutatingWebhookConfiguration), nil
}
