/*
Copyright 2016 The Kubernetes Authors.

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

package options

import (
	"fmt"

	"github.com/kzz45/neverdown/pkg/apiserver/server"

	"github.com/spf13/pflag"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/apiserver/pkg/features"
	"k8s.io/apiserver/pkg/storage/storagebackend"
	"k8s.io/apiserver/pkg/util/feature"
	utilflowcontrol "k8s.io/apiserver/pkg/util/flowcontrol"
	"k8s.io/client-go/kubernetes"
	"k8s.io/component-base/featuregate"
	"k8s.io/klog/v2"
)

// RecommendedOptions contains the recommended options for running an API server.
// If you add something to this list, it should be in a logical grouping.
// Each of them can be nil to leave the feature unconfigured on ApplyTo.
type RecommendedOptions struct {
	Etcd           *EtcdOptions
	SecureServing  *SecureServingOptionsWithLoopback
	Authentication *DelegatingAuthenticationOptions
	Authorization  *DelegatingAuthorizationOptions
	Audit          *AuditOptions
	Features       *FeatureOptions
	CoreAPI        *CoreAPIOptions

	// FeatureGate is a way to plumb feature gate through if you have them.
	FeatureGate featuregate.FeatureGate
	// ExtraAdmissionInitializers is called once after all ApplyTo from the options above, to pass the returned
	// admission plugin initializers to Admission.ApplyTo.
	ExtraAdmissionInitializers func(c *server.RecommendedConfig) ([]admission.PluginInitializer, error)
	Admission                  *AdmissionOptions
	// API Server Egress Selector is used to control outbound traffic from the API Server
	EgressSelector *EgressSelectorOptions
	// Traces contains options to control distributed request tracing.
	Traces *TracingOptions
}

func NewRecommendedOptions(prefix string, codec runtime.Codec) *RecommendedOptions {
	sso := NewSecureServingOptions()

	// We are composing recommended options for an aggregated api-server,
	// whose client is typically a proxy multiplexing many operations ---
	// notably including long-running ones --- into one HTTP/2 connection
	// into this server.  So allow many concurrent operations.
	sso.HTTP2MaxStreamsPerConnection = 1000

	return &RecommendedOptions{
		Etcd:           NewEtcdOptions(storagebackend.NewDefaultConfig(prefix, codec)),
		SecureServing:  sso.WithLoopback(),
		Authentication: NewDelegatingAuthenticationOptions(),
		Authorization:  NewDelegatingAuthorizationOptions(),
		Audit:          NewAuditOptions(),
		Features:       NewFeatureOptions(),
		CoreAPI:        NewCoreAPIOptions(),
		// Wired a global by default that sadly people will abuse to have different meanings in different repos.
		// Please consider creating your own FeatureGate so you can have a consistent meaning for what a variable contains
		// across different repos.  Future you will thank you.
		FeatureGate:                feature.DefaultFeatureGate,
		ExtraAdmissionInitializers: func(c *server.RecommendedConfig) ([]admission.PluginInitializer, error) { return nil, nil },
		Admission:                  NewAdmissionOptions(),
		EgressSelector:             NewEgressSelectorOptions(),
		Traces:                     NewTracingOptions(),
	}
}

func (o *RecommendedOptions) AddFlags(fs *pflag.FlagSet) {
	o.Etcd.AddFlags(fs)
	o.SecureServing.AddFlags(fs)
	o.Authentication.AddFlags(fs)
	o.Authorization.AddFlags(fs)
	o.Audit.AddFlags(fs)
	o.Features.AddFlags(fs)
	o.CoreAPI.AddFlags(fs)
	o.Admission.AddFlags(fs)
	o.EgressSelector.AddFlags(fs)
	o.Traces.AddFlags(fs)
}

// ApplyTo adds RecommendedOptions to the server configuration.
// pluginInitializers can be empty, it is only need for additional initializers.
func (o *RecommendedOptions) ApplyTo(config *server.RecommendedConfig) error {
	// if err := o.Etcd.Complete(config.Config.StorageObjectCountTracker, config.Config.DrainedNotify(), config.Config.AddPostStartHook); err != nil {
	// 	return err
	// }
	// if err := o.Etcd.ApplyTo(&config.Config); err != nil {
	// 	return err
	// }
	// if err := o.EgressSelector.ApplyTo(&config.Config); err != nil {
	// 	return err
	// }
	// if err := o.Traces.ApplyTo(config.Config.EgressSelector, &config.Config); err != nil {
	// 	return err
	// }
	// if err := o.SecureServing.ApplyTo(&config.Config.SecureServing, &config.Config.LoopbackClientConfig); err != nil {
	// 	return err
	// }
	// if err := o.Authentication.ApplyTo(&config.Config.Authentication, config.SecureServing, config.OpenAPIConfig); err != nil {
	// 	return err
	// }
	// if err := o.Authorization.ApplyTo(&config.Config.Authorization); err != nil {
	// 	return err
	// }
	// if err := o.Audit.ApplyTo(&config.Config); err != nil {
	// 	return err
	// }
	// if err := o.Features.ApplyTo(&config.Config); err != nil {
	// 	return err
	// }
	// if err := o.CoreAPI.ApplyTo(config); err != nil {
	// 	return err
	// }
	// initializers, err := o.ExtraAdmissionInitializers(config)
	// if err != nil {
	// 	return err
	// }
	// kubeClient, err := kubernetes.NewForConfig(config.ClientConfig)
	// if err != nil {
	// 	return err
	// }
	// dynamicClient, err := dynamic.NewForConfig(config.ClientConfig)
	// if err != nil {
	// 	return err
	// }
	// if err := o.Admission.ApplyTo(&config.Config, config.SharedInformerFactory, kubeClient, dynamicClient, o.FeatureGate,
	// 	initializers...); err != nil {
	// 	return err
	// }
	if feature.DefaultFeatureGate.Enabled(features.APIPriorityAndFairness) {
		if config.ClientConfig != nil {
			if config.MaxRequestsInFlight+config.MaxMutatingRequestsInFlight <= 0 {
				return fmt.Errorf("invalid configuration: MaxRequestsInFlight=%d and MaxMutatingRequestsInFlight=%d; they must add up to something positive", config.MaxRequestsInFlight, config.MaxMutatingRequestsInFlight)

			}
			config.FlowControl = utilflowcontrol.New(
				config.SharedInformerFactory,
				kubernetes.NewForConfigOrDie(config.ClientConfig).FlowcontrolV1beta3(),
				config.MaxRequestsInFlight+config.MaxMutatingRequestsInFlight,
				config.RequestTimeout/4,
			)
		} else {
			klog.Warningf("Neither kubeconfig is provided nor service-account is mounted, so APIPriorityAndFairness will be disabled")
		}
	}
	return nil
}

func (o *RecommendedOptions) Validate() []error {
	errors := []error{}
	errors = append(errors, o.Etcd.Validate()...)
	errors = append(errors, o.SecureServing.Validate()...)
	errors = append(errors, o.Authentication.Validate()...)
	errors = append(errors, o.Authorization.Validate()...)
	errors = append(errors, o.Audit.Validate()...)
	errors = append(errors, o.Features.Validate()...)
	errors = append(errors, o.CoreAPI.Validate()...)
	errors = append(errors, o.Admission.Validate()...)
	errors = append(errors, o.EgressSelector.Validate()...)
	errors = append(errors, o.Traces.Validate()...)

	return errors
}
