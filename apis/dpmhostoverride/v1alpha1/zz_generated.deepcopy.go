//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostOverride) DeepCopyInto(out *HostOverride) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostOverride.
func (in *HostOverride) DeepCopy() *HostOverride {
	if in == nil {
		return nil
	}
	out := new(HostOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HostOverride) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostOverrideInitParameters) DeepCopyInto(out *HostOverrideInitParameters) {
	*out = *in
	if in.ComputeClusterID != nil {
		in, out := &in.ComputeClusterID, &out.ComputeClusterID
		*out = new(string)
		**out = **in
	}
	if in.DpmAutomationLevel != nil {
		in, out := &in.DpmAutomationLevel, &out.DpmAutomationLevel
		*out = new(string)
		**out = **in
	}
	if in.DpmEnabled != nil {
		in, out := &in.DpmEnabled, &out.DpmEnabled
		*out = new(bool)
		**out = **in
	}
	if in.HostSystemID != nil {
		in, out := &in.HostSystemID, &out.HostSystemID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostOverrideInitParameters.
func (in *HostOverrideInitParameters) DeepCopy() *HostOverrideInitParameters {
	if in == nil {
		return nil
	}
	out := new(HostOverrideInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostOverrideList) DeepCopyInto(out *HostOverrideList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HostOverride, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostOverrideList.
func (in *HostOverrideList) DeepCopy() *HostOverrideList {
	if in == nil {
		return nil
	}
	out := new(HostOverrideList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HostOverrideList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostOverrideObservation) DeepCopyInto(out *HostOverrideObservation) {
	*out = *in
	if in.ComputeClusterID != nil {
		in, out := &in.ComputeClusterID, &out.ComputeClusterID
		*out = new(string)
		**out = **in
	}
	if in.DpmAutomationLevel != nil {
		in, out := &in.DpmAutomationLevel, &out.DpmAutomationLevel
		*out = new(string)
		**out = **in
	}
	if in.DpmEnabled != nil {
		in, out := &in.DpmEnabled, &out.DpmEnabled
		*out = new(bool)
		**out = **in
	}
	if in.HostSystemID != nil {
		in, out := &in.HostSystemID, &out.HostSystemID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostOverrideObservation.
func (in *HostOverrideObservation) DeepCopy() *HostOverrideObservation {
	if in == nil {
		return nil
	}
	out := new(HostOverrideObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostOverrideParameters) DeepCopyInto(out *HostOverrideParameters) {
	*out = *in
	if in.ComputeClusterID != nil {
		in, out := &in.ComputeClusterID, &out.ComputeClusterID
		*out = new(string)
		**out = **in
	}
	if in.DpmAutomationLevel != nil {
		in, out := &in.DpmAutomationLevel, &out.DpmAutomationLevel
		*out = new(string)
		**out = **in
	}
	if in.DpmEnabled != nil {
		in, out := &in.DpmEnabled, &out.DpmEnabled
		*out = new(bool)
		**out = **in
	}
	if in.HostSystemID != nil {
		in, out := &in.HostSystemID, &out.HostSystemID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostOverrideParameters.
func (in *HostOverrideParameters) DeepCopy() *HostOverrideParameters {
	if in == nil {
		return nil
	}
	out := new(HostOverrideParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostOverrideSpec) DeepCopyInto(out *HostOverrideSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostOverrideSpec.
func (in *HostOverrideSpec) DeepCopy() *HostOverrideSpec {
	if in == nil {
		return nil
	}
	out := new(HostOverrideSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostOverrideStatus) DeepCopyInto(out *HostOverrideStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostOverrideStatus.
func (in *HostOverrideStatus) DeepCopy() *HostOverrideStatus {
	if in == nil {
		return nil
	}
	out := new(HostOverrideStatus)
	in.DeepCopyInto(out)
	return out
}
