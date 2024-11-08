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
func (in *Pool) DeepCopyInto(out *Pool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pool.
func (in *Pool) DeepCopy() *Pool {
	if in == nil {
		return nil
	}
	out := new(Pool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Pool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolInitParameters) DeepCopyInto(out *PoolInitParameters) {
	*out = *in
	if in.CPUExpandable != nil {
		in, out := &in.CPUExpandable, &out.CPUExpandable
		*out = new(bool)
		**out = **in
	}
	if in.CPULimit != nil {
		in, out := &in.CPULimit, &out.CPULimit
		*out = new(float64)
		**out = **in
	}
	if in.CPUReservation != nil {
		in, out := &in.CPUReservation, &out.CPUReservation
		*out = new(float64)
		**out = **in
	}
	if in.CPUShareLevel != nil {
		in, out := &in.CPUShareLevel, &out.CPUShareLevel
		*out = new(string)
		**out = **in
	}
	if in.CPUShares != nil {
		in, out := &in.CPUShares, &out.CPUShares
		*out = new(float64)
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
	if in.MemoryExpandable != nil {
		in, out := &in.MemoryExpandable, &out.MemoryExpandable
		*out = new(bool)
		**out = **in
	}
	if in.MemoryLimit != nil {
		in, out := &in.MemoryLimit, &out.MemoryLimit
		*out = new(float64)
		**out = **in
	}
	if in.MemoryReservation != nil {
		in, out := &in.MemoryReservation, &out.MemoryReservation
		*out = new(float64)
		**out = **in
	}
	if in.MemoryShareLevel != nil {
		in, out := &in.MemoryShareLevel, &out.MemoryShareLevel
		*out = new(string)
		**out = **in
	}
	if in.MemoryShares != nil {
		in, out := &in.MemoryShares, &out.MemoryShares
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ParentResourcePoolID != nil {
		in, out := &in.ParentResourcePoolID, &out.ParentResourcePoolID
		*out = new(string)
		**out = **in
	}
	if in.ScaleDescendantsShares != nil {
		in, out := &in.ScaleDescendantsShares, &out.ScaleDescendantsShares
		*out = new(string)
		**out = **in
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolInitParameters.
func (in *PoolInitParameters) DeepCopy() *PoolInitParameters {
	if in == nil {
		return nil
	}
	out := new(PoolInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolList) DeepCopyInto(out *PoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Pool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolList.
func (in *PoolList) DeepCopy() *PoolList {
	if in == nil {
		return nil
	}
	out := new(PoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolObservation) DeepCopyInto(out *PoolObservation) {
	*out = *in
	if in.CPUExpandable != nil {
		in, out := &in.CPUExpandable, &out.CPUExpandable
		*out = new(bool)
		**out = **in
	}
	if in.CPULimit != nil {
		in, out := &in.CPULimit, &out.CPULimit
		*out = new(float64)
		**out = **in
	}
	if in.CPUReservation != nil {
		in, out := &in.CPUReservation, &out.CPUReservation
		*out = new(float64)
		**out = **in
	}
	if in.CPUShareLevel != nil {
		in, out := &in.CPUShareLevel, &out.CPUShareLevel
		*out = new(string)
		**out = **in
	}
	if in.CPUShares != nil {
		in, out := &in.CPUShares, &out.CPUShares
		*out = new(float64)
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
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.MemoryExpandable != nil {
		in, out := &in.MemoryExpandable, &out.MemoryExpandable
		*out = new(bool)
		**out = **in
	}
	if in.MemoryLimit != nil {
		in, out := &in.MemoryLimit, &out.MemoryLimit
		*out = new(float64)
		**out = **in
	}
	if in.MemoryReservation != nil {
		in, out := &in.MemoryReservation, &out.MemoryReservation
		*out = new(float64)
		**out = **in
	}
	if in.MemoryShareLevel != nil {
		in, out := &in.MemoryShareLevel, &out.MemoryShareLevel
		*out = new(string)
		**out = **in
	}
	if in.MemoryShares != nil {
		in, out := &in.MemoryShares, &out.MemoryShares
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ParentResourcePoolID != nil {
		in, out := &in.ParentResourcePoolID, &out.ParentResourcePoolID
		*out = new(string)
		**out = **in
	}
	if in.ScaleDescendantsShares != nil {
		in, out := &in.ScaleDescendantsShares, &out.ScaleDescendantsShares
		*out = new(string)
		**out = **in
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolObservation.
func (in *PoolObservation) DeepCopy() *PoolObservation {
	if in == nil {
		return nil
	}
	out := new(PoolObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolParameters) DeepCopyInto(out *PoolParameters) {
	*out = *in
	if in.CPUExpandable != nil {
		in, out := &in.CPUExpandable, &out.CPUExpandable
		*out = new(bool)
		**out = **in
	}
	if in.CPULimit != nil {
		in, out := &in.CPULimit, &out.CPULimit
		*out = new(float64)
		**out = **in
	}
	if in.CPUReservation != nil {
		in, out := &in.CPUReservation, &out.CPUReservation
		*out = new(float64)
		**out = **in
	}
	if in.CPUShareLevel != nil {
		in, out := &in.CPUShareLevel, &out.CPUShareLevel
		*out = new(string)
		**out = **in
	}
	if in.CPUShares != nil {
		in, out := &in.CPUShares, &out.CPUShares
		*out = new(float64)
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
	if in.MemoryExpandable != nil {
		in, out := &in.MemoryExpandable, &out.MemoryExpandable
		*out = new(bool)
		**out = **in
	}
	if in.MemoryLimit != nil {
		in, out := &in.MemoryLimit, &out.MemoryLimit
		*out = new(float64)
		**out = **in
	}
	if in.MemoryReservation != nil {
		in, out := &in.MemoryReservation, &out.MemoryReservation
		*out = new(float64)
		**out = **in
	}
	if in.MemoryShareLevel != nil {
		in, out := &in.MemoryShareLevel, &out.MemoryShareLevel
		*out = new(string)
		**out = **in
	}
	if in.MemoryShares != nil {
		in, out := &in.MemoryShares, &out.MemoryShares
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ParentResourcePoolID != nil {
		in, out := &in.ParentResourcePoolID, &out.ParentResourcePoolID
		*out = new(string)
		**out = **in
	}
	if in.ScaleDescendantsShares != nil {
		in, out := &in.ScaleDescendantsShares, &out.ScaleDescendantsShares
		*out = new(string)
		**out = **in
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolParameters.
func (in *PoolParameters) DeepCopy() *PoolParameters {
	if in == nil {
		return nil
	}
	out := new(PoolParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolSpec) DeepCopyInto(out *PoolSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolSpec.
func (in *PoolSpec) DeepCopy() *PoolSpec {
	if in == nil {
		return nil
	}
	out := new(PoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolStatus) DeepCopyInto(out *PoolStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolStatus.
func (in *PoolStatus) DeepCopy() *PoolStatus {
	if in == nil {
		return nil
	}
	out := new(PoolStatus)
	in.DeepCopyInto(out)
	return out
}