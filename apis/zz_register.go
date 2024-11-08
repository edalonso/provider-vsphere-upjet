// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/edalonso/provider-vsphere-upjet/apis/computecluster/v1alpha1"
	v1alpha1computeclusterhostgroup "github.com/edalonso/provider-vsphere-upjet/apis/computeclusterhostgroup/v1alpha1"
	v1alpha1computeclustervmaffinityrule "github.com/edalonso/provider-vsphere-upjet/apis/computeclustervmaffinityrule/v1alpha1"
	v1alpha1computeclustervmantiaffinityrule "github.com/edalonso/provider-vsphere-upjet/apis/computeclustervmantiaffinityrule/v1alpha1"
	v1alpha1computeclustervmdependencyrule "github.com/edalonso/provider-vsphere-upjet/apis/computeclustervmdependencyrule/v1alpha1"
	v1alpha1computeclustervmgroup "github.com/edalonso/provider-vsphere-upjet/apis/computeclustervmgroup/v1alpha1"
	v1alpha1computeclustervmhostrule "github.com/edalonso/provider-vsphere-upjet/apis/computeclustervmhostrule/v1alpha1"
	v1alpha1contentlibrary "github.com/edalonso/provider-vsphere-upjet/apis/contentlibrary/v1alpha1"
	v1alpha1contentlibraryitem "github.com/edalonso/provider-vsphere-upjet/apis/contentlibraryitem/v1alpha1"
	v1alpha1customattribute "github.com/edalonso/provider-vsphere-upjet/apis/customattribute/v1alpha1"
	v1alpha1datacenter "github.com/edalonso/provider-vsphere-upjet/apis/datacenter/v1alpha1"
	v1alpha1datastorecluster "github.com/edalonso/provider-vsphere-upjet/apis/datastorecluster/v1alpha1"
	v1alpha1datastoreclustervmantiaffinityrule "github.com/edalonso/provider-vsphere-upjet/apis/datastoreclustervmantiaffinityrule/v1alpha1"
	v1alpha1distributedportgroup "github.com/edalonso/provider-vsphere-upjet/apis/distributedportgroup/v1alpha1"
	v1alpha1distributedvirtualswitch "github.com/edalonso/provider-vsphere-upjet/apis/distributedvirtualswitch/v1alpha1"
	v1alpha1dpmhostoverride "github.com/edalonso/provider-vsphere-upjet/apis/dpmhostoverride/v1alpha1"
	v1alpha1drsvmoverride "github.com/edalonso/provider-vsphere-upjet/apis/drsvmoverride/v1alpha1"
	v1alpha1entitypermissions "github.com/edalonso/provider-vsphere-upjet/apis/entitypermissions/v1alpha1"
	v1alpha1file "github.com/edalonso/provider-vsphere-upjet/apis/file/v1alpha1"
	v1alpha1folder "github.com/edalonso/provider-vsphere-upjet/apis/folder/v1alpha1"
	v1alpha1havmoverride "github.com/edalonso/provider-vsphere-upjet/apis/havmoverride/v1alpha1"
	v1alpha1host "github.com/edalonso/provider-vsphere-upjet/apis/host/v1alpha1"
	v1alpha1hostportgroup "github.com/edalonso/provider-vsphere-upjet/apis/hostportgroup/v1alpha1"
	v1alpha1hostvirtualswitch "github.com/edalonso/provider-vsphere-upjet/apis/hostvirtualswitch/v1alpha1"
	v1alpha1license "github.com/edalonso/provider-vsphere-upjet/apis/license/v1alpha1"
	v1alpha1nasdatastore "github.com/edalonso/provider-vsphere-upjet/apis/nasdatastore/v1alpha1"
	v1alpha1resourcepool "github.com/edalonso/provider-vsphere-upjet/apis/resourcepool/v1alpha1"
	v1alpha1role "github.com/edalonso/provider-vsphere-upjet/apis/role/v1alpha1"
	v1alpha1storagedrsvmoverride "github.com/edalonso/provider-vsphere-upjet/apis/storagedrsvmoverride/v1alpha1"
	v1alpha1tag "github.com/edalonso/provider-vsphere-upjet/apis/tag/v1alpha1"
	v1alpha1tagcategory "github.com/edalonso/provider-vsphere-upjet/apis/tagcategory/v1alpha1"
	v1alpha1apis "github.com/edalonso/provider-vsphere-upjet/apis/v1alpha1"
	v1beta1 "github.com/edalonso/provider-vsphere-upjet/apis/v1beta1"
	v1alpha1vappcontainer "github.com/edalonso/provider-vsphere-upjet/apis/vappcontainer/v1alpha1"
	v1alpha1vappentity "github.com/edalonso/provider-vsphere-upjet/apis/vappentity/v1alpha1"
	v1alpha1virtualdisk "github.com/edalonso/provider-vsphere-upjet/apis/virtualdisk/v1alpha1"
	v1alpha1virtualmachine "github.com/edalonso/provider-vsphere-upjet/apis/virtualmachine/v1alpha1"
	v1alpha1virtualmachinesnapshot "github.com/edalonso/provider-vsphere-upjet/apis/virtualmachinesnapshot/v1alpha1"
	v1alpha1vmfsdatastore "github.com/edalonso/provider-vsphere-upjet/apis/vmfsdatastore/v1alpha1"
	v1alpha1vmstoragepolicy "github.com/edalonso/provider-vsphere-upjet/apis/vmstoragepolicy/v1alpha1"
	v1alpha1vnic "github.com/edalonso/provider-vsphere-upjet/apis/vnic/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1computeclusterhostgroup.SchemeBuilder.AddToScheme,
		v1alpha1computeclustervmaffinityrule.SchemeBuilder.AddToScheme,
		v1alpha1computeclustervmantiaffinityrule.SchemeBuilder.AddToScheme,
		v1alpha1computeclustervmdependencyrule.SchemeBuilder.AddToScheme,
		v1alpha1computeclustervmgroup.SchemeBuilder.AddToScheme,
		v1alpha1computeclustervmhostrule.SchemeBuilder.AddToScheme,
		v1alpha1contentlibrary.SchemeBuilder.AddToScheme,
		v1alpha1contentlibraryitem.SchemeBuilder.AddToScheme,
		v1alpha1customattribute.SchemeBuilder.AddToScheme,
		v1alpha1datacenter.SchemeBuilder.AddToScheme,
		v1alpha1datastorecluster.SchemeBuilder.AddToScheme,
		v1alpha1datastoreclustervmantiaffinityrule.SchemeBuilder.AddToScheme,
		v1alpha1distributedportgroup.SchemeBuilder.AddToScheme,
		v1alpha1distributedvirtualswitch.SchemeBuilder.AddToScheme,
		v1alpha1dpmhostoverride.SchemeBuilder.AddToScheme,
		v1alpha1drsvmoverride.SchemeBuilder.AddToScheme,
		v1alpha1entitypermissions.SchemeBuilder.AddToScheme,
		v1alpha1file.SchemeBuilder.AddToScheme,
		v1alpha1folder.SchemeBuilder.AddToScheme,
		v1alpha1havmoverride.SchemeBuilder.AddToScheme,
		v1alpha1host.SchemeBuilder.AddToScheme,
		v1alpha1hostportgroup.SchemeBuilder.AddToScheme,
		v1alpha1hostvirtualswitch.SchemeBuilder.AddToScheme,
		v1alpha1license.SchemeBuilder.AddToScheme,
		v1alpha1nasdatastore.SchemeBuilder.AddToScheme,
		v1alpha1resourcepool.SchemeBuilder.AddToScheme,
		v1alpha1role.SchemeBuilder.AddToScheme,
		v1alpha1storagedrsvmoverride.SchemeBuilder.AddToScheme,
		v1alpha1tag.SchemeBuilder.AddToScheme,
		v1alpha1tagcategory.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1beta1.SchemeBuilder.AddToScheme,
		v1alpha1vappcontainer.SchemeBuilder.AddToScheme,
		v1alpha1vappentity.SchemeBuilder.AddToScheme,
		v1alpha1virtualdisk.SchemeBuilder.AddToScheme,
		v1alpha1virtualmachine.SchemeBuilder.AddToScheme,
		v1alpha1virtualmachinesnapshot.SchemeBuilder.AddToScheme,
		v1alpha1vmfsdatastore.SchemeBuilder.AddToScheme,
		v1alpha1vmstoragepolicy.SchemeBuilder.AddToScheme,
		v1alpha1vnic.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
