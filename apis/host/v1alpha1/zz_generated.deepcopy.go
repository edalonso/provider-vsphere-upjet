//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Host) DeepCopyInto(out *Host) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Host.
func (in *Host) DeepCopy() *Host {
	if in == nil {
		return nil
	}
	out := new(Host)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Host) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostInitParameters) DeepCopyInto(out *HostInitParameters) {
	*out = *in
	if in.Cluster != nil {
		in, out := &in.Cluster, &out.Cluster
		*out = new(string)
		**out = **in
	}
	if in.ClusterManaged != nil {
		in, out := &in.ClusterManaged, &out.ClusterManaged
		*out = new(bool)
		**out = **in
	}
	if in.Connected != nil {
		in, out := &in.Connected, &out.Connected
		*out = new(bool)
		**out = **in
	}
	if in.CustomAttributes != nil {
		in, out := &in.CustomAttributes, &out.CustomAttributes
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Datacenter != nil {
		in, out := &in.Datacenter, &out.Datacenter
		*out = new(string)
		**out = **in
	}
	if in.DatacenterRef != nil {
		in, out := &in.DatacenterRef, &out.DatacenterRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.DatacenterSelector != nil {
		in, out := &in.DatacenterSelector, &out.DatacenterSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Force != nil {
		in, out := &in.Force, &out.Force
		*out = new(bool)
		**out = **in
	}
	if in.Hostname != nil {
		in, out := &in.Hostname, &out.Hostname
		*out = new(string)
		**out = **in
	}
	if in.License != nil {
		in, out := &in.License, &out.License
		*out = new(string)
		**out = **in
	}
	if in.Lockdown != nil {
		in, out := &in.Lockdown, &out.Lockdown
		*out = new(string)
		**out = **in
	}
	if in.Maintenance != nil {
		in, out := &in.Maintenance, &out.Maintenance
		*out = new(bool)
		**out = **in
	}
	out.PasswordSecretRef = in.PasswordSecretRef
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]ServicesInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Thumbprint != nil {
		in, out := &in.Thumbprint, &out.Thumbprint
		*out = new(string)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostInitParameters.
func (in *HostInitParameters) DeepCopy() *HostInitParameters {
	if in == nil {
		return nil
	}
	out := new(HostInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostList) DeepCopyInto(out *HostList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Host, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostList.
func (in *HostList) DeepCopy() *HostList {
	if in == nil {
		return nil
	}
	out := new(HostList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HostList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostObservation) DeepCopyInto(out *HostObservation) {
	*out = *in
	if in.Cluster != nil {
		in, out := &in.Cluster, &out.Cluster
		*out = new(string)
		**out = **in
	}
	if in.ClusterManaged != nil {
		in, out := &in.ClusterManaged, &out.ClusterManaged
		*out = new(bool)
		**out = **in
	}
	if in.Connected != nil {
		in, out := &in.Connected, &out.Connected
		*out = new(bool)
		**out = **in
	}
	if in.CustomAttributes != nil {
		in, out := &in.CustomAttributes, &out.CustomAttributes
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Datacenter != nil {
		in, out := &in.Datacenter, &out.Datacenter
		*out = new(string)
		**out = **in
	}
	if in.Force != nil {
		in, out := &in.Force, &out.Force
		*out = new(bool)
		**out = **in
	}
	if in.Hostname != nil {
		in, out := &in.Hostname, &out.Hostname
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.License != nil {
		in, out := &in.License, &out.License
		*out = new(string)
		**out = **in
	}
	if in.Lockdown != nil {
		in, out := &in.Lockdown, &out.Lockdown
		*out = new(string)
		**out = **in
	}
	if in.Maintenance != nil {
		in, out := &in.Maintenance, &out.Maintenance
		*out = new(bool)
		**out = **in
	}
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]ServicesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Thumbprint != nil {
		in, out := &in.Thumbprint, &out.Thumbprint
		*out = new(string)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostObservation.
func (in *HostObservation) DeepCopy() *HostObservation {
	if in == nil {
		return nil
	}
	out := new(HostObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostParameters) DeepCopyInto(out *HostParameters) {
	*out = *in
	if in.Cluster != nil {
		in, out := &in.Cluster, &out.Cluster
		*out = new(string)
		**out = **in
	}
	if in.ClusterManaged != nil {
		in, out := &in.ClusterManaged, &out.ClusterManaged
		*out = new(bool)
		**out = **in
	}
	if in.Connected != nil {
		in, out := &in.Connected, &out.Connected
		*out = new(bool)
		**out = **in
	}
	if in.CustomAttributes != nil {
		in, out := &in.CustomAttributes, &out.CustomAttributes
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Datacenter != nil {
		in, out := &in.Datacenter, &out.Datacenter
		*out = new(string)
		**out = **in
	}
	if in.DatacenterRef != nil {
		in, out := &in.DatacenterRef, &out.DatacenterRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.DatacenterSelector != nil {
		in, out := &in.DatacenterSelector, &out.DatacenterSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Force != nil {
		in, out := &in.Force, &out.Force
		*out = new(bool)
		**out = **in
	}
	if in.Hostname != nil {
		in, out := &in.Hostname, &out.Hostname
		*out = new(string)
		**out = **in
	}
	if in.License != nil {
		in, out := &in.License, &out.License
		*out = new(string)
		**out = **in
	}
	if in.Lockdown != nil {
		in, out := &in.Lockdown, &out.Lockdown
		*out = new(string)
		**out = **in
	}
	if in.Maintenance != nil {
		in, out := &in.Maintenance, &out.Maintenance
		*out = new(bool)
		**out = **in
	}
	out.PasswordSecretRef = in.PasswordSecretRef
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]ServicesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Thumbprint != nil {
		in, out := &in.Thumbprint, &out.Thumbprint
		*out = new(string)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostParameters.
func (in *HostParameters) DeepCopy() *HostParameters {
	if in == nil {
		return nil
	}
	out := new(HostParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostSpec) DeepCopyInto(out *HostSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostSpec.
func (in *HostSpec) DeepCopy() *HostSpec {
	if in == nil {
		return nil
	}
	out := new(HostSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostStatus) DeepCopyInto(out *HostStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostStatus.
func (in *HostStatus) DeepCopy() *HostStatus {
	if in == nil {
		return nil
	}
	out := new(HostStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NtpdInitParameters) DeepCopyInto(out *NtpdInitParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.NtpServers != nil {
		in, out := &in.NtpServers, &out.NtpServers
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NtpdInitParameters.
func (in *NtpdInitParameters) DeepCopy() *NtpdInitParameters {
	if in == nil {
		return nil
	}
	out := new(NtpdInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NtpdObservation) DeepCopyInto(out *NtpdObservation) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.NtpServers != nil {
		in, out := &in.NtpServers, &out.NtpServers
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NtpdObservation.
func (in *NtpdObservation) DeepCopy() *NtpdObservation {
	if in == nil {
		return nil
	}
	out := new(NtpdObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NtpdParameters) DeepCopyInto(out *NtpdParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.NtpServers != nil {
		in, out := &in.NtpServers, &out.NtpServers
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NtpdParameters.
func (in *NtpdParameters) DeepCopy() *NtpdParameters {
	if in == nil {
		return nil
	}
	out := new(NtpdParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesInitParameters) DeepCopyInto(out *ServicesInitParameters) {
	*out = *in
	if in.Ntpd != nil {
		in, out := &in.Ntpd, &out.Ntpd
		*out = make([]NtpdInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesInitParameters.
func (in *ServicesInitParameters) DeepCopy() *ServicesInitParameters {
	if in == nil {
		return nil
	}
	out := new(ServicesInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesObservation) DeepCopyInto(out *ServicesObservation) {
	*out = *in
	if in.Ntpd != nil {
		in, out := &in.Ntpd, &out.Ntpd
		*out = make([]NtpdObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesObservation.
func (in *ServicesObservation) DeepCopy() *ServicesObservation {
	if in == nil {
		return nil
	}
	out := new(ServicesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesParameters) DeepCopyInto(out *ServicesParameters) {
	*out = *in
	if in.Ntpd != nil {
		in, out := &in.Ntpd, &out.Ntpd
		*out = make([]NtpdParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesParameters.
func (in *ServicesParameters) DeepCopy() *ServicesParameters {
	if in == nil {
		return nil
	}
	out := new(ServicesParameters)
	in.DeepCopyInto(out)
	return out
}