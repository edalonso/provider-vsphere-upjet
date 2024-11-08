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
func (in *DrsVMOverride) DeepCopyInto(out *DrsVMOverride) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DrsVMOverride.
func (in *DrsVMOverride) DeepCopy() *DrsVMOverride {
	if in == nil {
		return nil
	}
	out := new(DrsVMOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DrsVMOverride) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DrsVMOverrideInitParameters) DeepCopyInto(out *DrsVMOverrideInitParameters) {
	*out = *in
	if in.DatastoreClusterID != nil {
		in, out := &in.DatastoreClusterID, &out.DatastoreClusterID
		*out = new(string)
		**out = **in
	}
	if in.SdrsAutomationLevel != nil {
		in, out := &in.SdrsAutomationLevel, &out.SdrsAutomationLevel
		*out = new(string)
		**out = **in
	}
	if in.SdrsEnabled != nil {
		in, out := &in.SdrsEnabled, &out.SdrsEnabled
		*out = new(string)
		**out = **in
	}
	if in.SdrsIntraVMAffinity != nil {
		in, out := &in.SdrsIntraVMAffinity, &out.SdrsIntraVMAffinity
		*out = new(string)
		**out = **in
	}
	if in.VirtualMachineID != nil {
		in, out := &in.VirtualMachineID, &out.VirtualMachineID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DrsVMOverrideInitParameters.
func (in *DrsVMOverrideInitParameters) DeepCopy() *DrsVMOverrideInitParameters {
	if in == nil {
		return nil
	}
	out := new(DrsVMOverrideInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DrsVMOverrideList) DeepCopyInto(out *DrsVMOverrideList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DrsVMOverride, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DrsVMOverrideList.
func (in *DrsVMOverrideList) DeepCopy() *DrsVMOverrideList {
	if in == nil {
		return nil
	}
	out := new(DrsVMOverrideList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DrsVMOverrideList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DrsVMOverrideObservation) DeepCopyInto(out *DrsVMOverrideObservation) {
	*out = *in
	if in.DatastoreClusterID != nil {
		in, out := &in.DatastoreClusterID, &out.DatastoreClusterID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.SdrsAutomationLevel != nil {
		in, out := &in.SdrsAutomationLevel, &out.SdrsAutomationLevel
		*out = new(string)
		**out = **in
	}
	if in.SdrsEnabled != nil {
		in, out := &in.SdrsEnabled, &out.SdrsEnabled
		*out = new(string)
		**out = **in
	}
	if in.SdrsIntraVMAffinity != nil {
		in, out := &in.SdrsIntraVMAffinity, &out.SdrsIntraVMAffinity
		*out = new(string)
		**out = **in
	}
	if in.VirtualMachineID != nil {
		in, out := &in.VirtualMachineID, &out.VirtualMachineID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DrsVMOverrideObservation.
func (in *DrsVMOverrideObservation) DeepCopy() *DrsVMOverrideObservation {
	if in == nil {
		return nil
	}
	out := new(DrsVMOverrideObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DrsVMOverrideParameters) DeepCopyInto(out *DrsVMOverrideParameters) {
	*out = *in
	if in.DatastoreClusterID != nil {
		in, out := &in.DatastoreClusterID, &out.DatastoreClusterID
		*out = new(string)
		**out = **in
	}
	if in.SdrsAutomationLevel != nil {
		in, out := &in.SdrsAutomationLevel, &out.SdrsAutomationLevel
		*out = new(string)
		**out = **in
	}
	if in.SdrsEnabled != nil {
		in, out := &in.SdrsEnabled, &out.SdrsEnabled
		*out = new(string)
		**out = **in
	}
	if in.SdrsIntraVMAffinity != nil {
		in, out := &in.SdrsIntraVMAffinity, &out.SdrsIntraVMAffinity
		*out = new(string)
		**out = **in
	}
	if in.VirtualMachineID != nil {
		in, out := &in.VirtualMachineID, &out.VirtualMachineID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DrsVMOverrideParameters.
func (in *DrsVMOverrideParameters) DeepCopy() *DrsVMOverrideParameters {
	if in == nil {
		return nil
	}
	out := new(DrsVMOverrideParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DrsVMOverrideSpec) DeepCopyInto(out *DrsVMOverrideSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DrsVMOverrideSpec.
func (in *DrsVMOverrideSpec) DeepCopy() *DrsVMOverrideSpec {
	if in == nil {
		return nil
	}
	out := new(DrsVMOverrideSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DrsVMOverrideStatus) DeepCopyInto(out *DrsVMOverrideStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DrsVMOverrideStatus.
func (in *DrsVMOverrideStatus) DeepCopy() *DrsVMOverrideStatus {
	if in == nil {
		return nil
	}
	out := new(DrsVMOverrideStatus)
	in.DeepCopyInto(out)
	return out
}