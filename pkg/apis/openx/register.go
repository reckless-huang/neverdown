package openx

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// GroupName is the group name use in this package
const GroupName = "openx.neverdown.io"

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: runtime.APIVersionInternal}

func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

// Adds the list of known types to the given scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&Openx{},
		&OpenxList{},
		&Etcd{},
		&EtcdList{},
		&Mysql{},
		&MysqlList{},
		&Redis{},
		&RedisList{},
		&Affinity{},
		&AffinityList{},
		&Toleration{},
		&TolerationList{},
		&NodeSelector{},
		&NodeSelectorList{},
		&AliyunLoadBalancer{},
		&AliyunLoadBalancerList{},
		&AliyunAccessControl{},
		&AliyunAccessControlList{},
	)
	return nil
}
