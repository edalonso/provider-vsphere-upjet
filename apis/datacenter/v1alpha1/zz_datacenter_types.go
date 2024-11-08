// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DatacenterInitParameters struct {

	// Map of custom attribute ids to value
	// strings to set for datacenter resource. See
	// here for a reference on how to set values
	// for custom attributes.
	// A list of custom attributes to set on this resource.
	// +mapType=granular
	CustomAttributes map[string]*string `json:"customAttributes,omitempty" tf:"custom_attributes,omitempty"`

	// The folder where the datacenter should be created.
	// Forces a new resource if changed.
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// The IDs of any tags to attach to this resource. See
	// here for a reference on how to apply tags.
	// A list of tag IDs to apply to this object.
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DatacenterObservation struct {

	// Map of custom attribute ids to value
	// strings to set for datacenter resource. See
	// here for a reference on how to set values
	// for custom attributes.
	// A list of custom attributes to set on this resource.
	// +mapType=granular
	CustomAttributes map[string]*string `json:"customAttributes,omitempty" tf:"custom_attributes,omitempty"`

	// The folder where the datacenter should be created.
	// Forces a new resource if changed.
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// The name of this datacenter. This will be changed to the managed
	// object ID in v2.0.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Managed object ID of this datacenter.
	// Managed object ID of the datacenter.
	Moid *string `json:"moid,omitempty" tf:"moid,omitempty"`

	// The IDs of any tags to attach to this resource. See
	// here for a reference on how to apply tags.
	// A list of tag IDs to apply to this object.
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DatacenterParameters struct {

	// Map of custom attribute ids to value
	// strings to set for datacenter resource. See
	// here for a reference on how to set values
	// for custom attributes.
	// A list of custom attributes to set on this resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	CustomAttributes map[string]*string `json:"customAttributes,omitempty" tf:"custom_attributes,omitempty"`

	// The folder where the datacenter should be created.
	// Forces a new resource if changed.
	// +kubebuilder:validation:Optional
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// The IDs of any tags to attach to this resource. See
	// here for a reference on how to apply tags.
	// A list of tag IDs to apply to this object.
	// +kubebuilder:validation:Optional
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// DatacenterSpec defines the desired state of Datacenter
type DatacenterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatacenterParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider DatacenterInitParameters `json:"initProvider,omitempty"`
}

// DatacenterStatus defines the observed state of Datacenter.
type DatacenterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatacenterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Datacenter is the Schema for the Datacenters API. Provides a VMware vSphere datacenter resource. This can be used as the primary container of inventory objects such as hosts and virtual machines.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vsphere-upjet}
type Datacenter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatacenterSpec   `json:"spec"`
	Status            DatacenterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatacenterList contains a list of Datacenters
type DatacenterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Datacenter `json:"items"`
}

// Repository type metadata.
var (
	Datacenter_Kind             = "Datacenter"
	Datacenter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Datacenter_Kind}.String()
	Datacenter_KindAPIVersion   = Datacenter_Kind + "." + CRDGroupVersion.String()
	Datacenter_GroupVersionKind = CRDGroupVersion.WithKind(Datacenter_Kind)
)

func init() {
	SchemeBuilder.Register(&Datacenter{}, &DatacenterList{})
}
