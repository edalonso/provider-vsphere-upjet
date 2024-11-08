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

type LicenseInitParameters struct {

	// A map of key/value pairs to be attached as labels (tags) to the license key.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The license key to add.
	LicenseKey *string `json:"licenseKey,omitempty" tf:"license_key,omitempty"`
}

type LicenseObservation struct {

	// The product edition of the license key.
	EditionKey *string `json:"editionKey,omitempty" tf:"edition_key,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A map of key/value pairs to be attached as labels (tags) to the license key.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The license key to add.
	LicenseKey *string `json:"licenseKey,omitempty" tf:"license_key,omitempty"`

	// The display name for the license.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Total number of units (example: CPUs) contained in the license.
	Total *float64 `json:"total,omitempty" tf:"total,omitempty"`

	// The number of units (example: CPUs) assigned to this license.
	Used *float64 `json:"used,omitempty" tf:"used,omitempty"`
}

type LicenseParameters struct {

	// A map of key/value pairs to be attached as labels (tags) to the license key.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The license key to add.
	// +kubebuilder:validation:Optional
	LicenseKey *string `json:"licenseKey,omitempty" tf:"license_key,omitempty"`
}

// LicenseSpec defines the desired state of License
type LicenseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LicenseParameters `json:"forProvider"`
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
	InitProvider LicenseInitParameters `json:"initProvider,omitempty"`
}

// LicenseStatus defines the observed state of License.
type LicenseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LicenseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// License is the Schema for the Licenses API. Provides a VMware vSphere license resource. This can be used to add and remove license keys.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vsphere-upjet}
type License struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.licenseKey) || (has(self.initProvider) && has(self.initProvider.licenseKey))",message="spec.forProvider.licenseKey is a required parameter"
	Spec   LicenseSpec   `json:"spec"`
	Status LicenseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LicenseList contains a list of Licenses
type LicenseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []License `json:"items"`
}

// Repository type metadata.
var (
	License_Kind             = "License"
	License_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: License_Kind}.String()
	License_KindAPIVersion   = License_Kind + "." + CRDGroupVersion.String()
	License_GroupVersionKind = CRDGroupVersion.WithKind(License_Kind)
)

func init() {
	SchemeBuilder.Register(&License{}, &LicenseList{})
}
