//
// Copyright 2019 HAProxy Technologies LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1alpha2

import (
	"net/http"

	v1alpha2 "github.com/haproxytech/kubernetes-ingress/crs/api/core/v1alpha2"
	"github.com/haproxytech/kubernetes-ingress/crs/generated/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type CoreV1alpha2Interface interface {
	RESTClient() rest.Interface
	BackendsGetter
	DefaultsGetter
	GlobalsGetter
}

// CoreV1alpha2Client is used to interact with features provided by the core.haproxy.org group.
type CoreV1alpha2Client struct {
	restClient rest.Interface
}

func (c *CoreV1alpha2Client) Backends(namespace string) BackendInterface {
	return newBackends(c, namespace)
}

func (c *CoreV1alpha2Client) Defaults(namespace string) DefaultsInterface {
	return newDefaults(c, namespace)
}

func (c *CoreV1alpha2Client) Globals(namespace string) GlobalInterface {
	return newGlobals(c, namespace)
}

// NewForConfig creates a new CoreV1alpha2Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*CoreV1alpha2Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new CoreV1alpha2Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*CoreV1alpha2Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &CoreV1alpha2Client{client}, nil
}

// NewForConfigOrDie creates a new CoreV1alpha2Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *CoreV1alpha2Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new CoreV1alpha2Client for the given RESTClient.
func New(c rest.Interface) *CoreV1alpha2Client {
	return &CoreV1alpha2Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha2.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *CoreV1alpha2Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
