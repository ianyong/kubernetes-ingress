// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"net/http"

	v1alpha1 "github.com/nginxinc/kubernetes-ingress/pkg/apis/configuration/v1alpha1"
	"github.com/nginxinc/kubernetes-ingress/pkg/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type K8sV1alpha1Interface interface {
	RESTClient() rest.Interface
	GlobalConfigurationsGetter
	PoliciesGetter
	TransportServersGetter
}

// K8sV1alpha1Client is used to interact with features provided by the k8s.nginx.org group.
type K8sV1alpha1Client struct {
	restClient rest.Interface
}

func (c *K8sV1alpha1Client) GlobalConfigurations(namespace string) GlobalConfigurationInterface {
	return newGlobalConfigurations(c, namespace)
}

func (c *K8sV1alpha1Client) Policies(namespace string) PolicyInterface {
	return newPolicies(c, namespace)
}

func (c *K8sV1alpha1Client) TransportServers(namespace string) TransportServerInterface {
	return newTransportServers(c, namespace)
}

// NewForConfig creates a new K8sV1alpha1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*K8sV1alpha1Client, error) {
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

// NewForConfigAndClient creates a new K8sV1alpha1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*K8sV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &K8sV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new K8sV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *K8sV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new K8sV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *K8sV1alpha1Client {
	return &K8sV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
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
func (c *K8sV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
