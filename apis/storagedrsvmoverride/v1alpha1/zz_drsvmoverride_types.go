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

type DrsVMOverrideInitParameters struct {

	// The managed object reference
	// ID of the datastore cluster to put the override in.
	// Forces a new resource if changed.
	// The managed object ID of the datastore cluster.
	DatastoreClusterID *string `json:"datastoreClusterId,omitempty" tf:"datastore_cluster_id,omitempty"`

	// Overrides any Storage DRS automation
	// levels for this virtual machine. Can be one of automated or manual. When
	// not specified, the datastore cluster's settings are used according to the
	// specific SDRS subsystem.
	// Overrides any Storage DRS automation levels for this virtual machine.
	SdrsAutomationLevel *string `json:"sdrsAutomationLevel,omitempty" tf:"sdrs_automation_level,omitempty"`

	// Overrides the default Storage DRS setting for
	// this virtual machine. When not specified, the datastore cluster setting is
	// used.
	// Overrides the default Storage DRS setting for this virtual machine.
	SdrsEnabled *string `json:"sdrsEnabled,omitempty" tf:"sdrs_enabled,omitempty"`

	// Overrides the intra-VM affinity setting
	// for this virtual machine. When true, all disks for this virtual machine
	// will be kept on the same datastore. When false, Storage DRS may locate
	// individual disks on different datastores if it helps satisfy cluster
	// requirements. When not specified, the datastore cluster's settings are used.
	// Overrides the intra-VM affinity setting for this virtual machine.
	SdrsIntraVMAffinity *string `json:"sdrsIntraVmAffinity,omitempty" tf:"sdrs_intra_vm_affinity,omitempty"`

	// The UUID of the virtual machine to create
	// the override for.  Forces a new resource if changed.
	// The managed object ID of the virtual machine.
	VirtualMachineID *string `json:"virtualMachineId,omitempty" tf:"virtual_machine_id,omitempty"`
}

type DrsVMOverrideObservation struct {

	// The managed object reference
	// ID of the datastore cluster to put the override in.
	// Forces a new resource if changed.
	// The managed object ID of the datastore cluster.
	DatastoreClusterID *string `json:"datastoreClusterId,omitempty" tf:"datastore_cluster_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Overrides any Storage DRS automation
	// levels for this virtual machine. Can be one of automated or manual. When
	// not specified, the datastore cluster's settings are used according to the
	// specific SDRS subsystem.
	// Overrides any Storage DRS automation levels for this virtual machine.
	SdrsAutomationLevel *string `json:"sdrsAutomationLevel,omitempty" tf:"sdrs_automation_level,omitempty"`

	// Overrides the default Storage DRS setting for
	// this virtual machine. When not specified, the datastore cluster setting is
	// used.
	// Overrides the default Storage DRS setting for this virtual machine.
	SdrsEnabled *string `json:"sdrsEnabled,omitempty" tf:"sdrs_enabled,omitempty"`

	// Overrides the intra-VM affinity setting
	// for this virtual machine. When true, all disks for this virtual machine
	// will be kept on the same datastore. When false, Storage DRS may locate
	// individual disks on different datastores if it helps satisfy cluster
	// requirements. When not specified, the datastore cluster's settings are used.
	// Overrides the intra-VM affinity setting for this virtual machine.
	SdrsIntraVMAffinity *string `json:"sdrsIntraVmAffinity,omitempty" tf:"sdrs_intra_vm_affinity,omitempty"`

	// The UUID of the virtual machine to create
	// the override for.  Forces a new resource if changed.
	// The managed object ID of the virtual machine.
	VirtualMachineID *string `json:"virtualMachineId,omitempty" tf:"virtual_machine_id,omitempty"`
}

type DrsVMOverrideParameters struct {

	// The managed object reference
	// ID of the datastore cluster to put the override in.
	// Forces a new resource if changed.
	// The managed object ID of the datastore cluster.
	// +kubebuilder:validation:Optional
	DatastoreClusterID *string `json:"datastoreClusterId,omitempty" tf:"datastore_cluster_id,omitempty"`

	// Overrides any Storage DRS automation
	// levels for this virtual machine. Can be one of automated or manual. When
	// not specified, the datastore cluster's settings are used according to the
	// specific SDRS subsystem.
	// Overrides any Storage DRS automation levels for this virtual machine.
	// +kubebuilder:validation:Optional
	SdrsAutomationLevel *string `json:"sdrsAutomationLevel,omitempty" tf:"sdrs_automation_level,omitempty"`

	// Overrides the default Storage DRS setting for
	// this virtual machine. When not specified, the datastore cluster setting is
	// used.
	// Overrides the default Storage DRS setting for this virtual machine.
	// +kubebuilder:validation:Optional
	SdrsEnabled *string `json:"sdrsEnabled,omitempty" tf:"sdrs_enabled,omitempty"`

	// Overrides the intra-VM affinity setting
	// for this virtual machine. When true, all disks for this virtual machine
	// will be kept on the same datastore. When false, Storage DRS may locate
	// individual disks on different datastores if it helps satisfy cluster
	// requirements. When not specified, the datastore cluster's settings are used.
	// Overrides the intra-VM affinity setting for this virtual machine.
	// +kubebuilder:validation:Optional
	SdrsIntraVMAffinity *string `json:"sdrsIntraVmAffinity,omitempty" tf:"sdrs_intra_vm_affinity,omitempty"`

	// The UUID of the virtual machine to create
	// the override for.  Forces a new resource if changed.
	// The managed object ID of the virtual machine.
	// +kubebuilder:validation:Optional
	VirtualMachineID *string `json:"virtualMachineId,omitempty" tf:"virtual_machine_id,omitempty"`
}

// DrsVMOverrideSpec defines the desired state of DrsVMOverride
type DrsVMOverrideSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DrsVMOverrideParameters `json:"forProvider"`
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
	InitProvider DrsVMOverrideInitParameters `json:"initProvider,omitempty"`
}

// DrsVMOverrideStatus defines the observed state of DrsVMOverride.
type DrsVMOverrideStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DrsVMOverrideObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DrsVMOverride is the Schema for the DrsVMOverrides API. Provides a VMware vSphere Storage DRS virtual machine override resource. This can be used to override Storage DRS settings in a datastore cluster.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vsphere-upjet}
type DrsVMOverride struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.datastoreClusterId) || (has(self.initProvider) && has(self.initProvider.datastoreClusterId))",message="spec.forProvider.datastoreClusterId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.virtualMachineId) || (has(self.initProvider) && has(self.initProvider.virtualMachineId))",message="spec.forProvider.virtualMachineId is a required parameter"
	Spec   DrsVMOverrideSpec   `json:"spec"`
	Status DrsVMOverrideStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DrsVMOverrideList contains a list of DrsVMOverrides
type DrsVMOverrideList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DrsVMOverride `json:"items"`
}

// Repository type metadata.
var (
	DrsVMOverride_Kind             = "DrsVMOverride"
	DrsVMOverride_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DrsVMOverride_Kind}.String()
	DrsVMOverride_KindAPIVersion   = DrsVMOverride_Kind + "." + CRDGroupVersion.String()
	DrsVMOverride_GroupVersionKind = CRDGroupVersion.WithKind(DrsVMOverride_Kind)
)

func init() {
	SchemeBuilder.Register(&DrsVMOverride{}, &DrsVMOverrideList{})
}