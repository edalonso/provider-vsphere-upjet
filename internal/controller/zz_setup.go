// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	cluster "github.com/edalonso/provider-vsphere-upjet/internal/controller/computecluster/cluster"
	clusterhostgroup "github.com/edalonso/provider-vsphere-upjet/internal/controller/computeclusterhostgroup/clusterhostgroup"
	clustervmaffinityrule "github.com/edalonso/provider-vsphere-upjet/internal/controller/computeclustervmaffinityrule/clustervmaffinityrule"
	clustervmantiaffinityrule "github.com/edalonso/provider-vsphere-upjet/internal/controller/computeclustervmantiaffinityrule/clustervmantiaffinityrule"
	clustervmdependencyrule "github.com/edalonso/provider-vsphere-upjet/internal/controller/computeclustervmdependencyrule/clustervmdependencyrule"
	clustervmgroup "github.com/edalonso/provider-vsphere-upjet/internal/controller/computeclustervmgroup/clustervmgroup"
	clustervmhostrule "github.com/edalonso/provider-vsphere-upjet/internal/controller/computeclustervmhostrule/clustervmhostrule"
	library "github.com/edalonso/provider-vsphere-upjet/internal/controller/contentlibrary/library"
	libraryitem "github.com/edalonso/provider-vsphere-upjet/internal/controller/contentlibraryitem/libraryitem"
	attribute "github.com/edalonso/provider-vsphere-upjet/internal/controller/customattribute/attribute"
	datacenter "github.com/edalonso/provider-vsphere-upjet/internal/controller/datacenter/datacenter"
	clusterdatastorecluster "github.com/edalonso/provider-vsphere-upjet/internal/controller/datastorecluster/cluster"
	clustervmantiaffinityruledatastoreclustervmantiaffinityrule "github.com/edalonso/provider-vsphere-upjet/internal/controller/datastoreclustervmantiaffinityrule/clustervmantiaffinityrule"
	portgroup "github.com/edalonso/provider-vsphere-upjet/internal/controller/distributedportgroup/portgroup"
	virtualswitch "github.com/edalonso/provider-vsphere-upjet/internal/controller/distributedvirtualswitch/virtualswitch"
	hostoverride "github.com/edalonso/provider-vsphere-upjet/internal/controller/dpmhostoverride/hostoverride"
	vmoverride "github.com/edalonso/provider-vsphere-upjet/internal/controller/drsvmoverride/vmoverride"
	permissions "github.com/edalonso/provider-vsphere-upjet/internal/controller/entitypermissions/permissions"
	file "github.com/edalonso/provider-vsphere-upjet/internal/controller/file/file"
	folder "github.com/edalonso/provider-vsphere-upjet/internal/controller/folder/folder"
	vmoverridehavmoverride "github.com/edalonso/provider-vsphere-upjet/internal/controller/havmoverride/vmoverride"
	host "github.com/edalonso/provider-vsphere-upjet/internal/controller/host/host"
	portgrouphostportgroup "github.com/edalonso/provider-vsphere-upjet/internal/controller/hostportgroup/portgroup"
	virtualswitchhostvirtualswitch "github.com/edalonso/provider-vsphere-upjet/internal/controller/hostvirtualswitch/virtualswitch"
	license "github.com/edalonso/provider-vsphere-upjet/internal/controller/license/license"
	datastore "github.com/edalonso/provider-vsphere-upjet/internal/controller/nasdatastore/datastore"
	providerconfig "github.com/edalonso/provider-vsphere-upjet/internal/controller/providerconfig"
	pool "github.com/edalonso/provider-vsphere-upjet/internal/controller/resourcepool/pool"
	role "github.com/edalonso/provider-vsphere-upjet/internal/controller/role/role"
	drsvmoverride "github.com/edalonso/provider-vsphere-upjet/internal/controller/storagedrsvmoverride/drsvmoverride"
	tag "github.com/edalonso/provider-vsphere-upjet/internal/controller/tag/tag"
	category "github.com/edalonso/provider-vsphere-upjet/internal/controller/tagcategory/category"
	container "github.com/edalonso/provider-vsphere-upjet/internal/controller/vappcontainer/container"
	entity "github.com/edalonso/provider-vsphere-upjet/internal/controller/vappentity/entity"
	disk "github.com/edalonso/provider-vsphere-upjet/internal/controller/virtualdisk/disk"
	machine "github.com/edalonso/provider-vsphere-upjet/internal/controller/virtualmachine/machine"
	machinesnapshot "github.com/edalonso/provider-vsphere-upjet/internal/controller/virtualmachinesnapshot/machinesnapshot"
	datastorevmfsdatastore "github.com/edalonso/provider-vsphere-upjet/internal/controller/vmfsdatastore/datastore"
	storagepolicy "github.com/edalonso/provider-vsphere-upjet/internal/controller/vmstoragepolicy/storagepolicy"
	vnic "github.com/edalonso/provider-vsphere-upjet/internal/controller/vnic/vnic"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.Setup,
		clusterhostgroup.Setup,
		clustervmaffinityrule.Setup,
		clustervmantiaffinityrule.Setup,
		clustervmdependencyrule.Setup,
		clustervmgroup.Setup,
		clustervmhostrule.Setup,
		library.Setup,
		libraryitem.Setup,
		attribute.Setup,
		datacenter.Setup,
		clusterdatastorecluster.Setup,
		clustervmantiaffinityruledatastoreclustervmantiaffinityrule.Setup,
		portgroup.Setup,
		virtualswitch.Setup,
		hostoverride.Setup,
		vmoverride.Setup,
		permissions.Setup,
		file.Setup,
		folder.Setup,
		vmoverridehavmoverride.Setup,
		host.Setup,
		portgrouphostportgroup.Setup,
		virtualswitchhostvirtualswitch.Setup,
		license.Setup,
		datastore.Setup,
		providerconfig.Setup,
		pool.Setup,
		role.Setup,
		drsvmoverride.Setup,
		tag.Setup,
		category.Setup,
		container.Setup,
		entity.Setup,
		disk.Setup,
		machine.Setup,
		machinesnapshot.Setup,
		datastorevmfsdatastore.Setup,
		storagepolicy.Setup,
		vnic.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
