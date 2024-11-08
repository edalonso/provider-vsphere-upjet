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
func (in *File) DeepCopyInto(out *File) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new File.
func (in *File) DeepCopy() *File {
	if in == nil {
		return nil
	}
	out := new(File)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *File) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileInitParameters) DeepCopyInto(out *FileInitParameters) {
	*out = *in
	if in.CreateDirectories != nil {
		in, out := &in.CreateDirectories, &out.CreateDirectories
		*out = new(bool)
		**out = **in
	}
	if in.Datacenter != nil {
		in, out := &in.Datacenter, &out.Datacenter
		*out = new(string)
		**out = **in
	}
	if in.Datastore != nil {
		in, out := &in.Datastore, &out.Datastore
		*out = new(string)
		**out = **in
	}
	if in.DestinationFile != nil {
		in, out := &in.DestinationFile, &out.DestinationFile
		*out = new(string)
		**out = **in
	}
	if in.SourceDatacenter != nil {
		in, out := &in.SourceDatacenter, &out.SourceDatacenter
		*out = new(string)
		**out = **in
	}
	if in.SourceDatastore != nil {
		in, out := &in.SourceDatastore, &out.SourceDatastore
		*out = new(string)
		**out = **in
	}
	if in.SourceFile != nil {
		in, out := &in.SourceFile, &out.SourceFile
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileInitParameters.
func (in *FileInitParameters) DeepCopy() *FileInitParameters {
	if in == nil {
		return nil
	}
	out := new(FileInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileList) DeepCopyInto(out *FileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]File, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileList.
func (in *FileList) DeepCopy() *FileList {
	if in == nil {
		return nil
	}
	out := new(FileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileObservation) DeepCopyInto(out *FileObservation) {
	*out = *in
	if in.CreateDirectories != nil {
		in, out := &in.CreateDirectories, &out.CreateDirectories
		*out = new(bool)
		**out = **in
	}
	if in.Datacenter != nil {
		in, out := &in.Datacenter, &out.Datacenter
		*out = new(string)
		**out = **in
	}
	if in.Datastore != nil {
		in, out := &in.Datastore, &out.Datastore
		*out = new(string)
		**out = **in
	}
	if in.DestinationFile != nil {
		in, out := &in.DestinationFile, &out.DestinationFile
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.SourceDatacenter != nil {
		in, out := &in.SourceDatacenter, &out.SourceDatacenter
		*out = new(string)
		**out = **in
	}
	if in.SourceDatastore != nil {
		in, out := &in.SourceDatastore, &out.SourceDatastore
		*out = new(string)
		**out = **in
	}
	if in.SourceFile != nil {
		in, out := &in.SourceFile, &out.SourceFile
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileObservation.
func (in *FileObservation) DeepCopy() *FileObservation {
	if in == nil {
		return nil
	}
	out := new(FileObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileParameters) DeepCopyInto(out *FileParameters) {
	*out = *in
	if in.CreateDirectories != nil {
		in, out := &in.CreateDirectories, &out.CreateDirectories
		*out = new(bool)
		**out = **in
	}
	if in.Datacenter != nil {
		in, out := &in.Datacenter, &out.Datacenter
		*out = new(string)
		**out = **in
	}
	if in.Datastore != nil {
		in, out := &in.Datastore, &out.Datastore
		*out = new(string)
		**out = **in
	}
	if in.DestinationFile != nil {
		in, out := &in.DestinationFile, &out.DestinationFile
		*out = new(string)
		**out = **in
	}
	if in.SourceDatacenter != nil {
		in, out := &in.SourceDatacenter, &out.SourceDatacenter
		*out = new(string)
		**out = **in
	}
	if in.SourceDatastore != nil {
		in, out := &in.SourceDatastore, &out.SourceDatastore
		*out = new(string)
		**out = **in
	}
	if in.SourceFile != nil {
		in, out := &in.SourceFile, &out.SourceFile
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileParameters.
func (in *FileParameters) DeepCopy() *FileParameters {
	if in == nil {
		return nil
	}
	out := new(FileParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSpec) DeepCopyInto(out *FileSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSpec.
func (in *FileSpec) DeepCopy() *FileSpec {
	if in == nil {
		return nil
	}
	out := new(FileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileStatus) DeepCopyInto(out *FileStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileStatus.
func (in *FileStatus) DeepCopy() *FileStatus {
	if in == nil {
		return nil
	}
	out := new(FileStatus)
	in.DeepCopyInto(out)
	return out
}