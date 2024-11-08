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
func (in *ClusterVMAntiAffinityRule) DeepCopyInto(out *ClusterVMAntiAffinityRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterVMAntiAffinityRule.
func (in *ClusterVMAntiAffinityRule) DeepCopy() *ClusterVMAntiAffinityRule {
	if in == nil {
		return nil
	}
	out := new(ClusterVMAntiAffinityRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterVMAntiAffinityRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterVMAntiAffinityRuleInitParameters) DeepCopyInto(out *ClusterVMAntiAffinityRuleInitParameters) {
	*out = *in
	if in.ComputeClusterID != nil {
		in, out := &in.ComputeClusterID, &out.ComputeClusterID
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Mandatory != nil {
		in, out := &in.Mandatory, &out.Mandatory
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.VirtualMachineIds != nil {
		in, out := &in.VirtualMachineIds, &out.VirtualMachineIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterVMAntiAffinityRuleInitParameters.
func (in *ClusterVMAntiAffinityRuleInitParameters) DeepCopy() *ClusterVMAntiAffinityRuleInitParameters {
	if in == nil {
		return nil
	}
	out := new(ClusterVMAntiAffinityRuleInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterVMAntiAffinityRuleList) DeepCopyInto(out *ClusterVMAntiAffinityRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterVMAntiAffinityRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterVMAntiAffinityRuleList.
func (in *ClusterVMAntiAffinityRuleList) DeepCopy() *ClusterVMAntiAffinityRuleList {
	if in == nil {
		return nil
	}
	out := new(ClusterVMAntiAffinityRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterVMAntiAffinityRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterVMAntiAffinityRuleObservation) DeepCopyInto(out *ClusterVMAntiAffinityRuleObservation) {
	*out = *in
	if in.ComputeClusterID != nil {
		in, out := &in.ComputeClusterID, &out.ComputeClusterID
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Mandatory != nil {
		in, out := &in.Mandatory, &out.Mandatory
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.VirtualMachineIds != nil {
		in, out := &in.VirtualMachineIds, &out.VirtualMachineIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterVMAntiAffinityRuleObservation.
func (in *ClusterVMAntiAffinityRuleObservation) DeepCopy() *ClusterVMAntiAffinityRuleObservation {
	if in == nil {
		return nil
	}
	out := new(ClusterVMAntiAffinityRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterVMAntiAffinityRuleParameters) DeepCopyInto(out *ClusterVMAntiAffinityRuleParameters) {
	*out = *in
	if in.ComputeClusterID != nil {
		in, out := &in.ComputeClusterID, &out.ComputeClusterID
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Mandatory != nil {
		in, out := &in.Mandatory, &out.Mandatory
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.VirtualMachineIds != nil {
		in, out := &in.VirtualMachineIds, &out.VirtualMachineIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterVMAntiAffinityRuleParameters.
func (in *ClusterVMAntiAffinityRuleParameters) DeepCopy() *ClusterVMAntiAffinityRuleParameters {
	if in == nil {
		return nil
	}
	out := new(ClusterVMAntiAffinityRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterVMAntiAffinityRuleSpec) DeepCopyInto(out *ClusterVMAntiAffinityRuleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterVMAntiAffinityRuleSpec.
func (in *ClusterVMAntiAffinityRuleSpec) DeepCopy() *ClusterVMAntiAffinityRuleSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterVMAntiAffinityRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterVMAntiAffinityRuleStatus) DeepCopyInto(out *ClusterVMAntiAffinityRuleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterVMAntiAffinityRuleStatus.
func (in *ClusterVMAntiAffinityRuleStatus) DeepCopy() *ClusterVMAntiAffinityRuleStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterVMAntiAffinityRuleStatus)
	in.DeepCopyInto(out)
	return out
}
