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

type DiskInitParameters struct {

	// The adapter type for this virtual disk. Can be
	// one of ide, lsiLogic, or busLogic.  Default: lsiLogic.
	AdapterType *string `json:"adapterType,omitempty" tf:"adapter_type,omitempty"`

	// Tells the resource to create any
	// directories that are a part of the vmdk_path parameter if they are missing.
	// Default: false.
	CreateDirectories *bool `json:"createDirectories,omitempty" tf:"create_directories,omitempty"`

	// The name of the datacenter in which to create the
	// disk. Can be omitted when ESXi or if there is only one datacenter in
	// your infrastructure.
	Datacenter *string `json:"datacenter,omitempty" tf:"datacenter,omitempty"`

	// The name of the datastore in which to create the
	// disk.
	Datastore *string `json:"datastore,omitempty" tf:"datastore,omitempty"`

	// Size of the disk (in GB). Decreasing the size of a disk is not possible.
	// If a disk of a smaller size is required then the original has to be destroyed along with its data and a new one has to be
	// created.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The type of disk to create. Can be one of
	// eagerZeroedThick, lazy, or thin. Default: eagerZeroedThick. For
	// information on what each kind of disk provisioning policy means, click
	// here.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The path, including filename, of the virtual disk to
	// be created.  This needs to end in .vmdk.
	VmdkPath *string `json:"vmdkPath,omitempty" tf:"vmdk_path,omitempty"`
}

type DiskObservation struct {

	// The adapter type for this virtual disk. Can be
	// one of ide, lsiLogic, or busLogic.  Default: lsiLogic.
	AdapterType *string `json:"adapterType,omitempty" tf:"adapter_type,omitempty"`

	// Tells the resource to create any
	// directories that are a part of the vmdk_path parameter if they are missing.
	// Default: false.
	CreateDirectories *bool `json:"createDirectories,omitempty" tf:"create_directories,omitempty"`

	// The name of the datacenter in which to create the
	// disk. Can be omitted when ESXi or if there is only one datacenter in
	// your infrastructure.
	Datacenter *string `json:"datacenter,omitempty" tf:"datacenter,omitempty"`

	// The name of the datastore in which to create the
	// disk.
	Datastore *string `json:"datastore,omitempty" tf:"datastore,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Size of the disk (in GB). Decreasing the size of a disk is not possible.
	// If a disk of a smaller size is required then the original has to be destroyed along with its data and a new one has to be
	// created.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The type of disk to create. Can be one of
	// eagerZeroedThick, lazy, or thin. Default: eagerZeroedThick. For
	// information on what each kind of disk provisioning policy means, click
	// here.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The path, including filename, of the virtual disk to
	// be created.  This needs to end in .vmdk.
	VmdkPath *string `json:"vmdkPath,omitempty" tf:"vmdk_path,omitempty"`
}

type DiskParameters struct {

	// The adapter type for this virtual disk. Can be
	// one of ide, lsiLogic, or busLogic.  Default: lsiLogic.
	// +kubebuilder:validation:Optional
	AdapterType *string `json:"adapterType,omitempty" tf:"adapter_type,omitempty"`

	// Tells the resource to create any
	// directories that are a part of the vmdk_path parameter if they are missing.
	// Default: false.
	// +kubebuilder:validation:Optional
	CreateDirectories *bool `json:"createDirectories,omitempty" tf:"create_directories,omitempty"`

	// The name of the datacenter in which to create the
	// disk. Can be omitted when ESXi or if there is only one datacenter in
	// your infrastructure.
	// +kubebuilder:validation:Optional
	Datacenter *string `json:"datacenter,omitempty" tf:"datacenter,omitempty"`

	// The name of the datastore in which to create the
	// disk.
	// +kubebuilder:validation:Optional
	Datastore *string `json:"datastore,omitempty" tf:"datastore,omitempty"`

	// Size of the disk (in GB). Decreasing the size of a disk is not possible.
	// If a disk of a smaller size is required then the original has to be destroyed along with its data and a new one has to be
	// created.
	// +kubebuilder:validation:Optional
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The type of disk to create. Can be one of
	// eagerZeroedThick, lazy, or thin. Default: eagerZeroedThick. For
	// information on what each kind of disk provisioning policy means, click
	// here.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The path, including filename, of the virtual disk to
	// be created.  This needs to end in .vmdk.
	// +kubebuilder:validation:Optional
	VmdkPath *string `json:"vmdkPath,omitempty" tf:"vmdk_path,omitempty"`
}

// DiskSpec defines the desired state of Disk
type DiskSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DiskParameters `json:"forProvider"`
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
	InitProvider DiskInitParameters `json:"initProvider,omitempty"`
}

// DiskStatus defines the observed state of Disk.
type DiskStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DiskObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Disk is the Schema for the Disks API. Provides a vSphere virtual disk resource.  This can be used to create and delete virtual disks.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vsphere-upjet}
type Disk struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.datastore) || (has(self.initProvider) && has(self.initProvider.datastore))",message="spec.forProvider.datastore is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.size) || (has(self.initProvider) && has(self.initProvider.size))",message="spec.forProvider.size is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.vmdkPath) || (has(self.initProvider) && has(self.initProvider.vmdkPath))",message="spec.forProvider.vmdkPath is a required parameter"
	Spec   DiskSpec   `json:"spec"`
	Status DiskStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DiskList contains a list of Disks
type DiskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Disk `json:"items"`
}

// Repository type metadata.
var (
	Disk_Kind             = "Disk"
	Disk_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Disk_Kind}.String()
	Disk_KindAPIVersion   = Disk_Kind + "." + CRDGroupVersion.String()
	Disk_GroupVersionKind = CRDGroupVersion.WithKind(Disk_Kind)
)

func init() {
	SchemeBuilder.Register(&Disk{}, &DiskList{})
}
