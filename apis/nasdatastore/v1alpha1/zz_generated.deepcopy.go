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
func (in *Datastore) DeepCopyInto(out *Datastore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Datastore.
func (in *Datastore) DeepCopy() *Datastore {
	if in == nil {
		return nil
	}
	out := new(Datastore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Datastore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatastoreInitParameters) DeepCopyInto(out *DatastoreInitParameters) {
	*out = *in
	if in.AccessMode != nil {
		in, out := &in.AccessMode, &out.AccessMode
		*out = new(string)
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
	if in.DatastoreClusterID != nil {
		in, out := &in.DatastoreClusterID, &out.DatastoreClusterID
		*out = new(string)
		**out = **in
	}
	if in.Folder != nil {
		in, out := &in.Folder, &out.Folder
		*out = new(string)
		**out = **in
	}
	if in.HostSystemIds != nil {
		in, out := &in.HostSystemIds, &out.HostSystemIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RemoteHosts != nil {
		in, out := &in.RemoteHosts, &out.RemoteHosts
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RemotePath != nil {
		in, out := &in.RemotePath, &out.RemotePath
		*out = new(string)
		**out = **in
	}
	if in.SecurityType != nil {
		in, out := &in.SecurityType, &out.SecurityType
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
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatastoreInitParameters.
func (in *DatastoreInitParameters) DeepCopy() *DatastoreInitParameters {
	if in == nil {
		return nil
	}
	out := new(DatastoreInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatastoreList) DeepCopyInto(out *DatastoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Datastore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatastoreList.
func (in *DatastoreList) DeepCopy() *DatastoreList {
	if in == nil {
		return nil
	}
	out := new(DatastoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatastoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatastoreObservation) DeepCopyInto(out *DatastoreObservation) {
	*out = *in
	if in.AccessMode != nil {
		in, out := &in.AccessMode, &out.AccessMode
		*out = new(string)
		**out = **in
	}
	if in.Accessible != nil {
		in, out := &in.Accessible, &out.Accessible
		*out = new(bool)
		**out = **in
	}
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
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
	if in.DatastoreClusterID != nil {
		in, out := &in.DatastoreClusterID, &out.DatastoreClusterID
		*out = new(string)
		**out = **in
	}
	if in.Folder != nil {
		in, out := &in.Folder, &out.Folder
		*out = new(string)
		**out = **in
	}
	if in.FreeSpace != nil {
		in, out := &in.FreeSpace, &out.FreeSpace
		*out = new(float64)
		**out = **in
	}
	if in.HostSystemIds != nil {
		in, out := &in.HostSystemIds, &out.HostSystemIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.MaintenanceMode != nil {
		in, out := &in.MaintenanceMode, &out.MaintenanceMode
		*out = new(string)
		**out = **in
	}
	if in.MultipleHostAccess != nil {
		in, out := &in.MultipleHostAccess, &out.MultipleHostAccess
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ProtocolEndpoint != nil {
		in, out := &in.ProtocolEndpoint, &out.ProtocolEndpoint
		*out = new(bool)
		**out = **in
	}
	if in.RemoteHosts != nil {
		in, out := &in.RemoteHosts, &out.RemoteHosts
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RemotePath != nil {
		in, out := &in.RemotePath, &out.RemotePath
		*out = new(string)
		**out = **in
	}
	if in.SecurityType != nil {
		in, out := &in.SecurityType, &out.SecurityType
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
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
	if in.UncommittedSpace != nil {
		in, out := &in.UncommittedSpace, &out.UncommittedSpace
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatastoreObservation.
func (in *DatastoreObservation) DeepCopy() *DatastoreObservation {
	if in == nil {
		return nil
	}
	out := new(DatastoreObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatastoreParameters) DeepCopyInto(out *DatastoreParameters) {
	*out = *in
	if in.AccessMode != nil {
		in, out := &in.AccessMode, &out.AccessMode
		*out = new(string)
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
	if in.DatastoreClusterID != nil {
		in, out := &in.DatastoreClusterID, &out.DatastoreClusterID
		*out = new(string)
		**out = **in
	}
	if in.Folder != nil {
		in, out := &in.Folder, &out.Folder
		*out = new(string)
		**out = **in
	}
	if in.HostSystemIds != nil {
		in, out := &in.HostSystemIds, &out.HostSystemIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RemoteHosts != nil {
		in, out := &in.RemoteHosts, &out.RemoteHosts
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RemotePath != nil {
		in, out := &in.RemotePath, &out.RemotePath
		*out = new(string)
		**out = **in
	}
	if in.SecurityType != nil {
		in, out := &in.SecurityType, &out.SecurityType
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
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatastoreParameters.
func (in *DatastoreParameters) DeepCopy() *DatastoreParameters {
	if in == nil {
		return nil
	}
	out := new(DatastoreParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatastoreSpec) DeepCopyInto(out *DatastoreSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatastoreSpec.
func (in *DatastoreSpec) DeepCopy() *DatastoreSpec {
	if in == nil {
		return nil
	}
	out := new(DatastoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatastoreStatus) DeepCopyInto(out *DatastoreStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatastoreStatus.
func (in *DatastoreStatus) DeepCopy() *DatastoreStatus {
	if in == nil {
		return nil
	}
	out := new(DatastoreStatus)
	in.DeepCopyInto(out)
	return out
}
