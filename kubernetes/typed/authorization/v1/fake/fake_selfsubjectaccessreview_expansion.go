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

package fake

import (
	authorizationapi "k8s.io/api/authorization/v1"
	core "github.com/nalp/client-go/testing"
)

func (c *FakeSelfSubjectAccessReviews) Create(sar *authorizationapi.SelfSubjectAccessReview) (result *authorizationapi.SelfSubjectAccessReview, err error) {
	obj, err := c.Fake.Invokes(core.NewRootCreateAction(authorizationapi.SchemeGroupVersion.WithResource("selfsubjectaccessreviews"), sar), &authorizationapi.SelfSubjectAccessReview{})
	return obj.(*authorizationapi.SelfSubjectAccessReview), err
}
