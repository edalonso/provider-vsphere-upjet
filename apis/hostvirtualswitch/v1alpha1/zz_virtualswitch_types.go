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

type VirtualSwitchInitParameters struct {

	// The list of active network adapters used for load
	// balancing.
	// List of active network adapters used for load balancing.
	ActiveNics []*string `json:"activeNics,omitempty" tf:"active_nics,omitempty"`

	// Controls whether or not the virtual
	// network adapter is allowed to send network traffic with a different MAC
	// address than that of its own. Default: true.
	// Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than that of its own.
	AllowForgedTransmits *bool `json:"allowForgedTransmits,omitempty" tf:"allow_forged_transmits,omitempty"`

	// Controls whether or not the Media Access
	// Control (MAC) address can be changed. Default: true.
	// Controls whether or not the Media Access Control (MAC) address can be changed.
	AllowMacChanges *bool `json:"allowMacChanges,omitempty" tf:"allow_mac_changes,omitempty"`

	// Enable promiscuous mode on the network. This
	// flag indicates whether or not all traffic is seen on a given port. Default:
	// false.
	// Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
	AllowPromiscuous *bool `json:"allowPromiscuous,omitempty" tf:"allow_promiscuous,omitempty"`

	// The interval, in seconds, that a NIC beacon
	// packet is sent out. This can be used with check_beacon to
	// offer link failure capability beyond link status only. Default: 1.
	// Determines how often, in seconds, a beacon should be sent to probe for the validity of a link.
	BeaconInterval *float64 `json:"beaconInterval,omitempty" tf:"beacon_interval,omitempty"`

	// Enable beacon probing - this requires that the
	// beacon_interval option has been set in the bridge
	// options. If this is set to false, only link status is used to check for
	// failed NICs.  Default: false.
	// Enable beacon probing. Requires that the vSwitch has been configured to use a beacon. If disabled, link status is used only.
	CheckBeacon *bool `json:"checkBeacon,omitempty" tf:"check_beacon,omitempty"`

	// If set to true, the teaming policy will re-activate
	// failed interfaces higher in precedence when they come back up.  Default:
	// true.
	// If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
	Failback *bool `json:"failback,omitempty" tf:"failback,omitempty"`

	// The managed object ID of
	// the host to set the virtual switch up on. Forces a new resource if changed.
	// The managed object ID of the host to set the virtual switch up on.
	HostSystemID *string `json:"hostSystemId,omitempty" tf:"host_system_id,omitempty"`

	// Whether to advertise or listen
	// for link discovery traffic. Default: listen.
	// Whether to advertise or listen for link discovery. Valid values are advertise, both, listen, and none.
	LinkDiscoveryOperation *string `json:"linkDiscoveryOperation,omitempty" tf:"link_discovery_operation,omitempty"`

	// The discovery protocol type.  Valid
	// types are cpd and lldp. Default: cdp.
	// The discovery protocol type. Valid values are cdp and lldp.
	LinkDiscoveryProtocol *string `json:"linkDiscoveryProtocol,omitempty" tf:"link_discovery_protocol,omitempty"`

	// The maximum transmission unit (MTU) for the virtual
	// switch. Default: 1500.
	// The maximum transmission unit (MTU) of the virtual switch in bytes.
	Mtu *float64 `json:"mtu,omitempty" tf:"mtu,omitempty"`

	// The network interfaces to bind to the bridge.
	// The list of network adapters to bind to this virtual switch.
	NetworkAdapters []*string `json:"networkAdapters,omitempty" tf:"network_adapters,omitempty"`

	// If set to true, the teaming policy will
	// notify the broadcast network of a NIC failover, triggering cache updates.
	// Default: true.
	// If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
	NotifySwitches *bool `json:"notifySwitches,omitempty" tf:"notify_switches,omitempty"`

	// The number of ports to create with this
	// virtual switch. Default: 128.
	// The number of ports that this virtual switch is configured to use.
	NumberOfPorts *float64 `json:"numberOfPorts,omitempty" tf:"number_of_ports,omitempty"`

	// The average bandwidth in bits per
	// second if traffic shaping is enabled. Default: 0
	// The average bandwidth in bits per second if traffic shaping is enabled.
	ShapingAverageBandwidth *float64 `json:"shapingAverageBandwidth,omitempty" tf:"shaping_average_bandwidth,omitempty"`

	// The maximum burst size allowed in bytes if
	// shaping is enabled. Default: 0
	// The maximum burst size allowed in bytes if traffic shaping is enabled.
	ShapingBurstSize *float64 `json:"shapingBurstSize,omitempty" tf:"shaping_burst_size,omitempty"`

	// Set to true to enable the traffic shaper for
	// ports managed by this virtual switch. Default: false.
	// Enable traffic shaping on this virtual switch or port group.
	ShapingEnabled *bool `json:"shapingEnabled,omitempty" tf:"shaping_enabled,omitempty"`

	// The peak bandwidth during bursts in
	// bits per second if traffic shaping is enabled. Default: 0
	// The peak bandwidth during bursts in bits per second if traffic shaping is enabled.
	ShapingPeakBandwidth *float64 `json:"shapingPeakBandwidth,omitempty" tf:"shaping_peak_bandwidth,omitempty"`

	// The list of standby network adapters used for
	// failover.
	// List of standby network adapters used for failover.
	StandbyNics []*string `json:"standbyNics,omitempty" tf:"standby_nics,omitempty"`

	// The network adapter teaming policy. Can be one
	// of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid, or
	// failover_explicit. Default: loadbalance_srcid.
	// The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid, or failover_explicit.
	TeamingPolicy *string `json:"teamingPolicy,omitempty" tf:"teaming_policy,omitempty"`
}

type VirtualSwitchObservation struct {

	// The list of active network adapters used for load
	// balancing.
	// List of active network adapters used for load balancing.
	ActiveNics []*string `json:"activeNics,omitempty" tf:"active_nics,omitempty"`

	// Controls whether or not the virtual
	// network adapter is allowed to send network traffic with a different MAC
	// address than that of its own. Default: true.
	// Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than that of its own.
	AllowForgedTransmits *bool `json:"allowForgedTransmits,omitempty" tf:"allow_forged_transmits,omitempty"`

	// Controls whether or not the Media Access
	// Control (MAC) address can be changed. Default: true.
	// Controls whether or not the Media Access Control (MAC) address can be changed.
	AllowMacChanges *bool `json:"allowMacChanges,omitempty" tf:"allow_mac_changes,omitempty"`

	// Enable promiscuous mode on the network. This
	// flag indicates whether or not all traffic is seen on a given port. Default:
	// false.
	// Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
	AllowPromiscuous *bool `json:"allowPromiscuous,omitempty" tf:"allow_promiscuous,omitempty"`

	// The interval, in seconds, that a NIC beacon
	// packet is sent out. This can be used with check_beacon to
	// offer link failure capability beyond link status only. Default: 1.
	// Determines how often, in seconds, a beacon should be sent to probe for the validity of a link.
	BeaconInterval *float64 `json:"beaconInterval,omitempty" tf:"beacon_interval,omitempty"`

	// Enable beacon probing - this requires that the
	// beacon_interval option has been set in the bridge
	// options. If this is set to false, only link status is used to check for
	// failed NICs.  Default: false.
	// Enable beacon probing. Requires that the vSwitch has been configured to use a beacon. If disabled, link status is used only.
	CheckBeacon *bool `json:"checkBeacon,omitempty" tf:"check_beacon,omitempty"`

	// If set to true, the teaming policy will re-activate
	// failed interfaces higher in precedence when they come back up.  Default:
	// true.
	// If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
	Failback *bool `json:"failback,omitempty" tf:"failback,omitempty"`

	// The managed object ID of
	// the host to set the virtual switch up on. Forces a new resource if changed.
	// The managed object ID of the host to set the virtual switch up on.
	HostSystemID *string `json:"hostSystemId,omitempty" tf:"host_system_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether to advertise or listen
	// for link discovery traffic. Default: listen.
	// Whether to advertise or listen for link discovery. Valid values are advertise, both, listen, and none.
	LinkDiscoveryOperation *string `json:"linkDiscoveryOperation,omitempty" tf:"link_discovery_operation,omitempty"`

	// The discovery protocol type.  Valid
	// types are cpd and lldp. Default: cdp.
	// The discovery protocol type. Valid values are cdp and lldp.
	LinkDiscoveryProtocol *string `json:"linkDiscoveryProtocol,omitempty" tf:"link_discovery_protocol,omitempty"`

	// The maximum transmission unit (MTU) for the virtual
	// switch. Default: 1500.
	// The maximum transmission unit (MTU) of the virtual switch in bytes.
	Mtu *float64 `json:"mtu,omitempty" tf:"mtu,omitempty"`

	// The network interfaces to bind to the bridge.
	// The list of network adapters to bind to this virtual switch.
	NetworkAdapters []*string `json:"networkAdapters,omitempty" tf:"network_adapters,omitempty"`

	// If set to true, the teaming policy will
	// notify the broadcast network of a NIC failover, triggering cache updates.
	// Default: true.
	// If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
	NotifySwitches *bool `json:"notifySwitches,omitempty" tf:"notify_switches,omitempty"`

	// The number of ports to create with this
	// virtual switch. Default: 128.
	// The number of ports that this virtual switch is configured to use.
	NumberOfPorts *float64 `json:"numberOfPorts,omitempty" tf:"number_of_ports,omitempty"`

	// The average bandwidth in bits per
	// second if traffic shaping is enabled. Default: 0
	// The average bandwidth in bits per second if traffic shaping is enabled.
	ShapingAverageBandwidth *float64 `json:"shapingAverageBandwidth,omitempty" tf:"shaping_average_bandwidth,omitempty"`

	// The maximum burst size allowed in bytes if
	// shaping is enabled. Default: 0
	// The maximum burst size allowed in bytes if traffic shaping is enabled.
	ShapingBurstSize *float64 `json:"shapingBurstSize,omitempty" tf:"shaping_burst_size,omitempty"`

	// Set to true to enable the traffic shaper for
	// ports managed by this virtual switch. Default: false.
	// Enable traffic shaping on this virtual switch or port group.
	ShapingEnabled *bool `json:"shapingEnabled,omitempty" tf:"shaping_enabled,omitempty"`

	// The peak bandwidth during bursts in
	// bits per second if traffic shaping is enabled. Default: 0
	// The peak bandwidth during bursts in bits per second if traffic shaping is enabled.
	ShapingPeakBandwidth *float64 `json:"shapingPeakBandwidth,omitempty" tf:"shaping_peak_bandwidth,omitempty"`

	// The list of standby network adapters used for
	// failover.
	// List of standby network adapters used for failover.
	StandbyNics []*string `json:"standbyNics,omitempty" tf:"standby_nics,omitempty"`

	// The network adapter teaming policy. Can be one
	// of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid, or
	// failover_explicit. Default: loadbalance_srcid.
	// The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid, or failover_explicit.
	TeamingPolicy *string `json:"teamingPolicy,omitempty" tf:"teaming_policy,omitempty"`
}

type VirtualSwitchParameters struct {

	// The list of active network adapters used for load
	// balancing.
	// List of active network adapters used for load balancing.
	// +kubebuilder:validation:Optional
	ActiveNics []*string `json:"activeNics,omitempty" tf:"active_nics,omitempty"`

	// Controls whether or not the virtual
	// network adapter is allowed to send network traffic with a different MAC
	// address than that of its own. Default: true.
	// Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than that of its own.
	// +kubebuilder:validation:Optional
	AllowForgedTransmits *bool `json:"allowForgedTransmits,omitempty" tf:"allow_forged_transmits,omitempty"`

	// Controls whether or not the Media Access
	// Control (MAC) address can be changed. Default: true.
	// Controls whether or not the Media Access Control (MAC) address can be changed.
	// +kubebuilder:validation:Optional
	AllowMacChanges *bool `json:"allowMacChanges,omitempty" tf:"allow_mac_changes,omitempty"`

	// Enable promiscuous mode on the network. This
	// flag indicates whether or not all traffic is seen on a given port. Default:
	// false.
	// Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
	// +kubebuilder:validation:Optional
	AllowPromiscuous *bool `json:"allowPromiscuous,omitempty" tf:"allow_promiscuous,omitempty"`

	// The interval, in seconds, that a NIC beacon
	// packet is sent out. This can be used with check_beacon to
	// offer link failure capability beyond link status only. Default: 1.
	// Determines how often, in seconds, a beacon should be sent to probe for the validity of a link.
	// +kubebuilder:validation:Optional
	BeaconInterval *float64 `json:"beaconInterval,omitempty" tf:"beacon_interval,omitempty"`

	// Enable beacon probing - this requires that the
	// beacon_interval option has been set in the bridge
	// options. If this is set to false, only link status is used to check for
	// failed NICs.  Default: false.
	// Enable beacon probing. Requires that the vSwitch has been configured to use a beacon. If disabled, link status is used only.
	// +kubebuilder:validation:Optional
	CheckBeacon *bool `json:"checkBeacon,omitempty" tf:"check_beacon,omitempty"`

	// If set to true, the teaming policy will re-activate
	// failed interfaces higher in precedence when they come back up.  Default:
	// true.
	// If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
	// +kubebuilder:validation:Optional
	Failback *bool `json:"failback,omitempty" tf:"failback,omitempty"`

	// The managed object ID of
	// the host to set the virtual switch up on. Forces a new resource if changed.
	// The managed object ID of the host to set the virtual switch up on.
	// +kubebuilder:validation:Optional
	HostSystemID *string `json:"hostSystemId,omitempty" tf:"host_system_id,omitempty"`

	// Whether to advertise or listen
	// for link discovery traffic. Default: listen.
	// Whether to advertise or listen for link discovery. Valid values are advertise, both, listen, and none.
	// +kubebuilder:validation:Optional
	LinkDiscoveryOperation *string `json:"linkDiscoveryOperation,omitempty" tf:"link_discovery_operation,omitempty"`

	// The discovery protocol type.  Valid
	// types are cpd and lldp. Default: cdp.
	// The discovery protocol type. Valid values are cdp and lldp.
	// +kubebuilder:validation:Optional
	LinkDiscoveryProtocol *string `json:"linkDiscoveryProtocol,omitempty" tf:"link_discovery_protocol,omitempty"`

	// The maximum transmission unit (MTU) for the virtual
	// switch. Default: 1500.
	// The maximum transmission unit (MTU) of the virtual switch in bytes.
	// +kubebuilder:validation:Optional
	Mtu *float64 `json:"mtu,omitempty" tf:"mtu,omitempty"`

	// The network interfaces to bind to the bridge.
	// The list of network adapters to bind to this virtual switch.
	// +kubebuilder:validation:Optional
	NetworkAdapters []*string `json:"networkAdapters,omitempty" tf:"network_adapters,omitempty"`

	// If set to true, the teaming policy will
	// notify the broadcast network of a NIC failover, triggering cache updates.
	// Default: true.
	// If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
	// +kubebuilder:validation:Optional
	NotifySwitches *bool `json:"notifySwitches,omitempty" tf:"notify_switches,omitempty"`

	// The number of ports to create with this
	// virtual switch. Default: 128.
	// The number of ports that this virtual switch is configured to use.
	// +kubebuilder:validation:Optional
	NumberOfPorts *float64 `json:"numberOfPorts,omitempty" tf:"number_of_ports,omitempty"`

	// The average bandwidth in bits per
	// second if traffic shaping is enabled. Default: 0
	// The average bandwidth in bits per second if traffic shaping is enabled.
	// +kubebuilder:validation:Optional
	ShapingAverageBandwidth *float64 `json:"shapingAverageBandwidth,omitempty" tf:"shaping_average_bandwidth,omitempty"`

	// The maximum burst size allowed in bytes if
	// shaping is enabled. Default: 0
	// The maximum burst size allowed in bytes if traffic shaping is enabled.
	// +kubebuilder:validation:Optional
	ShapingBurstSize *float64 `json:"shapingBurstSize,omitempty" tf:"shaping_burst_size,omitempty"`

	// Set to true to enable the traffic shaper for
	// ports managed by this virtual switch. Default: false.
	// Enable traffic shaping on this virtual switch or port group.
	// +kubebuilder:validation:Optional
	ShapingEnabled *bool `json:"shapingEnabled,omitempty" tf:"shaping_enabled,omitempty"`

	// The peak bandwidth during bursts in
	// bits per second if traffic shaping is enabled. Default: 0
	// The peak bandwidth during bursts in bits per second if traffic shaping is enabled.
	// +kubebuilder:validation:Optional
	ShapingPeakBandwidth *float64 `json:"shapingPeakBandwidth,omitempty" tf:"shaping_peak_bandwidth,omitempty"`

	// The list of standby network adapters used for
	// failover.
	// List of standby network adapters used for failover.
	// +kubebuilder:validation:Optional
	StandbyNics []*string `json:"standbyNics,omitempty" tf:"standby_nics,omitempty"`

	// The network adapter teaming policy. Can be one
	// of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid, or
	// failover_explicit. Default: loadbalance_srcid.
	// The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid, or failover_explicit.
	// +kubebuilder:validation:Optional
	TeamingPolicy *string `json:"teamingPolicy,omitempty" tf:"teaming_policy,omitempty"`
}

// VirtualSwitchSpec defines the desired state of VirtualSwitch
type VirtualSwitchSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualSwitchParameters `json:"forProvider"`
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
	InitProvider VirtualSwitchInitParameters `json:"initProvider,omitempty"`
}

// VirtualSwitchStatus defines the observed state of VirtualSwitch.
type VirtualSwitchStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualSwitchObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// VirtualSwitch is the Schema for the VirtualSwitchs API. Provides a vSphere Host Virtual Switch Resource. This can be used to configure vSwitches direct on an ESXi host.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vsphere-upjet}
type VirtualSwitch struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.activeNics) || (has(self.initProvider) && has(self.initProvider.activeNics))",message="spec.forProvider.activeNics is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.hostSystemId) || (has(self.initProvider) && has(self.initProvider.hostSystemId))",message="spec.forProvider.hostSystemId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.networkAdapters) || (has(self.initProvider) && has(self.initProvider.networkAdapters))",message="spec.forProvider.networkAdapters is a required parameter"
	Spec   VirtualSwitchSpec   `json:"spec"`
	Status VirtualSwitchStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualSwitchList contains a list of VirtualSwitchs
type VirtualSwitchList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualSwitch `json:"items"`
}

// Repository type metadata.
var (
	VirtualSwitch_Kind             = "VirtualSwitch"
	VirtualSwitch_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VirtualSwitch_Kind}.String()
	VirtualSwitch_KindAPIVersion   = VirtualSwitch_Kind + "." + CRDGroupVersion.String()
	VirtualSwitch_GroupVersionKind = CRDGroupVersion.WithKind(VirtualSwitch_Kind)
)

func init() {
	SchemeBuilder.Register(&VirtualSwitch{}, &VirtualSwitchList{})
}
