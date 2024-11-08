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

type EntityInitParameters struct {

	// [Managed object ID|docs-about-morefs] of the vApp
	// container the entity is a member of.
	// Managed object ID of the vApp container the entity is a member of.
	ContainerID *string `json:"containerId,omitempty" tf:"container_id,omitempty"`

	// A list of custom attributes to set on this resource.
	// +mapType=granular
	CustomAttributes map[string]*string `json:"customAttributes,omitempty" tf:"custom_attributes,omitempty"`

	// How to start the entity. Valid settings are none
	// or powerOn. If set to none, then the entity does not participate in auto-start.
	// Default: powerOn
	// How to start the entity. Valid settings are none or powerOn. If set to none, then the entity does not participate in auto-start.
	StartAction *string `json:"startAction,omitempty" tf:"start_action,omitempty"`

	// Delay in seconds before continuing with the next
	// entity in the order of entities to be started. Default: 120
	// Delay in seconds before continuing with the next entity in the order of entities to be started.
	StartDelay *float64 `json:"startDelay,omitempty" tf:"start_delay,omitempty"`

	// Order to start and stop target in vApp. Default: 1
	// Order to start and stop target in vApp.
	StartOrder *float64 `json:"startOrder,omitempty" tf:"start_order,omitempty"`

	// Defines the stop action for the entity. Can be set
	// to none, powerOff, guestShutdown, or suspend. If set to none, then the entity
	// does not participate in auto-stop. Default: powerOff
	// Defines the stop action for the entity. Can be set to none, powerOff, guestShutdown, or suspend. If set to none, then the entity does not participate in auto-stop.
	StopAction *string `json:"stopAction,omitempty" tf:"stop_action,omitempty"`

	// Delay in seconds before continuing with the next
	// entity in the order sequence. This is only used if the stopAction is
	// guestShutdown. Default: 120
	// Delay in seconds before continuing with the next entity in the order of entities to be stopped.
	StopDelay *float64 `json:"stopDelay,omitempty" tf:"stop_delay,omitempty"`

	// A list of tag IDs to apply to this object.
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// [Managed object ID|docs-about-morefs] of the entity
	// to power on or power off. This can be a virtual machine or a vApp.
	// Managed object ID of the entity to power on or power off. This can be a virtual machine or a vApp.
	TargetID *string `json:"targetId,omitempty" tf:"target_id,omitempty"`

	// Determines if the VM should be marked as being
	// started when VMware Tools are ready instead of waiting for start_delay. This
	// property has no effect for vApps. Default: false
	// Determines if the VM should be marked as being started when VMware Tools are ready instead of waiting for start_delay. This property has no effect for vApps.
	WaitForGuest *bool `json:"waitForGuest,omitempty" tf:"wait_for_guest,omitempty"`
}

type EntityObservation struct {

	// [Managed object ID|docs-about-morefs] of the vApp
	// container the entity is a member of.
	// Managed object ID of the vApp container the entity is a member of.
	ContainerID *string `json:"containerId,omitempty" tf:"container_id,omitempty"`

	// A list of custom attributes to set on this resource.
	// +mapType=granular
	CustomAttributes map[string]*string `json:"customAttributes,omitempty" tf:"custom_attributes,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// How to start the entity. Valid settings are none
	// or powerOn. If set to none, then the entity does not participate in auto-start.
	// Default: powerOn
	// How to start the entity. Valid settings are none or powerOn. If set to none, then the entity does not participate in auto-start.
	StartAction *string `json:"startAction,omitempty" tf:"start_action,omitempty"`

	// Delay in seconds before continuing with the next
	// entity in the order of entities to be started. Default: 120
	// Delay in seconds before continuing with the next entity in the order of entities to be started.
	StartDelay *float64 `json:"startDelay,omitempty" tf:"start_delay,omitempty"`

	// Order to start and stop target in vApp. Default: 1
	// Order to start and stop target in vApp.
	StartOrder *float64 `json:"startOrder,omitempty" tf:"start_order,omitempty"`

	// Defines the stop action for the entity. Can be set
	// to none, powerOff, guestShutdown, or suspend. If set to none, then the entity
	// does not participate in auto-stop. Default: powerOff
	// Defines the stop action for the entity. Can be set to none, powerOff, guestShutdown, or suspend. If set to none, then the entity does not participate in auto-stop.
	StopAction *string `json:"stopAction,omitempty" tf:"stop_action,omitempty"`

	// Delay in seconds before continuing with the next
	// entity in the order sequence. This is only used if the stopAction is
	// guestShutdown. Default: 120
	// Delay in seconds before continuing with the next entity in the order of entities to be stopped.
	StopDelay *float64 `json:"stopDelay,omitempty" tf:"stop_delay,omitempty"`

	// A list of tag IDs to apply to this object.
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// [Managed object ID|docs-about-morefs] of the entity
	// to power on or power off. This can be a virtual machine or a vApp.
	// Managed object ID of the entity to power on or power off. This can be a virtual machine or a vApp.
	TargetID *string `json:"targetId,omitempty" tf:"target_id,omitempty"`

	// Determines if the VM should be marked as being
	// started when VMware Tools are ready instead of waiting for start_delay. This
	// property has no effect for vApps. Default: false
	// Determines if the VM should be marked as being started when VMware Tools are ready instead of waiting for start_delay. This property has no effect for vApps.
	WaitForGuest *bool `json:"waitForGuest,omitempty" tf:"wait_for_guest,omitempty"`
}

type EntityParameters struct {

	// [Managed object ID|docs-about-morefs] of the vApp
	// container the entity is a member of.
	// Managed object ID of the vApp container the entity is a member of.
	// +kubebuilder:validation:Optional
	ContainerID *string `json:"containerId,omitempty" tf:"container_id,omitempty"`

	// A list of custom attributes to set on this resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	CustomAttributes map[string]*string `json:"customAttributes,omitempty" tf:"custom_attributes,omitempty"`

	// How to start the entity. Valid settings are none
	// or powerOn. If set to none, then the entity does not participate in auto-start.
	// Default: powerOn
	// How to start the entity. Valid settings are none or powerOn. If set to none, then the entity does not participate in auto-start.
	// +kubebuilder:validation:Optional
	StartAction *string `json:"startAction,omitempty" tf:"start_action,omitempty"`

	// Delay in seconds before continuing with the next
	// entity in the order of entities to be started. Default: 120
	// Delay in seconds before continuing with the next entity in the order of entities to be started.
	// +kubebuilder:validation:Optional
	StartDelay *float64 `json:"startDelay,omitempty" tf:"start_delay,omitempty"`

	// Order to start and stop target in vApp. Default: 1
	// Order to start and stop target in vApp.
	// +kubebuilder:validation:Optional
	StartOrder *float64 `json:"startOrder,omitempty" tf:"start_order,omitempty"`

	// Defines the stop action for the entity. Can be set
	// to none, powerOff, guestShutdown, or suspend. If set to none, then the entity
	// does not participate in auto-stop. Default: powerOff
	// Defines the stop action for the entity. Can be set to none, powerOff, guestShutdown, or suspend. If set to none, then the entity does not participate in auto-stop.
	// +kubebuilder:validation:Optional
	StopAction *string `json:"stopAction,omitempty" tf:"stop_action,omitempty"`

	// Delay in seconds before continuing with the next
	// entity in the order sequence. This is only used if the stopAction is
	// guestShutdown. Default: 120
	// Delay in seconds before continuing with the next entity in the order of entities to be stopped.
	// +kubebuilder:validation:Optional
	StopDelay *float64 `json:"stopDelay,omitempty" tf:"stop_delay,omitempty"`

	// A list of tag IDs to apply to this object.
	// +kubebuilder:validation:Optional
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// [Managed object ID|docs-about-morefs] of the entity
	// to power on or power off. This can be a virtual machine or a vApp.
	// Managed object ID of the entity to power on or power off. This can be a virtual machine or a vApp.
	// +kubebuilder:validation:Optional
	TargetID *string `json:"targetId,omitempty" tf:"target_id,omitempty"`

	// Determines if the VM should be marked as being
	// started when VMware Tools are ready instead of waiting for start_delay. This
	// property has no effect for vApps. Default: false
	// Determines if the VM should be marked as being started when VMware Tools are ready instead of waiting for start_delay. This property has no effect for vApps.
	// +kubebuilder:validation:Optional
	WaitForGuest *bool `json:"waitForGuest,omitempty" tf:"wait_for_guest,omitempty"`
}

// EntitySpec defines the desired state of Entity
type EntitySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EntityParameters `json:"forProvider"`
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
	InitProvider EntityInitParameters `json:"initProvider,omitempty"`
}

// EntityStatus defines the observed state of Entity.
type EntityStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EntityObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Entity is the Schema for the Entitys API. Provides a vSphere vApp entity resource. This can be used to describe the behavior of an entity (virtual machine or sub-vApp container) in a vApp container.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vsphere-upjet}
type Entity struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.containerId) || (has(self.initProvider) && has(self.initProvider.containerId))",message="spec.forProvider.containerId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.targetId) || (has(self.initProvider) && has(self.initProvider.targetId))",message="spec.forProvider.targetId is a required parameter"
	Spec   EntitySpec   `json:"spec"`
	Status EntityStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EntityList contains a list of Entitys
type EntityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Entity `json:"items"`
}

// Repository type metadata.
var (
	Entity_Kind             = "Entity"
	Entity_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Entity_Kind}.String()
	Entity_KindAPIVersion   = Entity_Kind + "." + CRDGroupVersion.String()
	Entity_GroupVersionKind = CRDGroupVersion.WithKind(Entity_Kind)
)

func init() {
	SchemeBuilder.Register(&Entity{}, &EntityList{})
}
