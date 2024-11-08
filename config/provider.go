/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/edalonso/provider-vsphere-upjet/config/license"
        "github.com/edalonso/provider-vsphere-upjet/config/datacenter"
        "github.com/edalonso/provider-vsphere-upjet/config/host"
        "github.com/edalonso/provider-vsphere-upjet/config/customattribute"
        "github.com/edalonso/provider-vsphere-upjet/config/contentlibrary"
        "github.com/edalonso/provider-vsphere-upjet/config/folder"
        "github.com/edalonso/provider-vsphere-upjet/config/tagcategory"
        "github.com/edalonso/provider-vsphere-upjet/config/tag"
        "github.com/edalonso/provider-vsphere-upjet/config/hostvirtualswitch"
        "github.com/edalonso/provider-vsphere-upjet/config/hostportgroup"
        "github.com/edalonso/provider-vsphere-upjet/config/distributedvirtualswitch"
        "github.com/edalonso/provider-vsphere-upjet/config/distributedportgroup"
        "github.com/edalonso/provider-vsphere-upjet/config/vnic"
        "github.com/edalonso/provider-vsphere-upjet/config/role"
        "github.com/edalonso/provider-vsphere-upjet/config/entitypermissions"
        "github.com/edalonso/provider-vsphere-upjet/config/vmfsdatastore"
        "github.com/edalonso/provider-vsphere-upjet/config/nasdatastore"
        "github.com/edalonso/provider-vsphere-upjet/config/file"
        "github.com/edalonso/provider-vsphere-upjet/config/vmstoragepolicy"
        "github.com/edalonso/provider-vsphere-upjet/config/datastorecluster"
        "github.com/edalonso/provider-vsphere-upjet/config/computecluster"
        "github.com/edalonso/provider-vsphere-upjet/config/resourcepool"
        "github.com/edalonso/provider-vsphere-upjet/config/virtualdisk"
        "github.com/edalonso/provider-vsphere-upjet/config/vappentity"
        "github.com/edalonso/provider-vsphere-upjet/config/virtualmachine"
        "github.com/edalonso/provider-vsphere-upjet/config/virtualmachinesnapshot"
        "github.com/edalonso/provider-vsphere-upjet/config/vappcontainer"
        "github.com/edalonso/provider-vsphere-upjet/config/contentlibraryitem"
        "github.com/edalonso/provider-vsphere-upjet/config/computeclusterhostgroup"
        "github.com/edalonso/provider-vsphere-upjet/config/computeclustervmaffinityrule"
        "github.com/edalonso/provider-vsphere-upjet/config/computeclustervmantiaffinityrule"
        "github.com/edalonso/provider-vsphere-upjet/config/computeclustervmdependencyrule"
        "github.com/edalonso/provider-vsphere-upjet/config/computeclustervmgroup"
        "github.com/edalonso/provider-vsphere-upjet/config/computeclustervmhostrule"
        "github.com/edalonso/provider-vsphere-upjet/config/drsvmoverride"
        "github.com/edalonso/provider-vsphere-upjet/config/dpmhostoverride"
        "github.com/edalonso/provider-vsphere-upjet/config/havmoverride"
        "github.com/edalonso/provider-vsphere-upjet/config/storagedrsvmoverride"
        "github.com/edalonso/provider-vsphere-upjet/config/datastoreclustervmantiaffinityrule"
)

const (
	resourcePrefix = "vsphere-upjet"
	modulePath     = "github.com/edalonso/provider-vsphere-upjet"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("disashop.com"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
                license.Configure,
                datacenter.Configure,
		host.Configure,
                customattribute.Configure,
                folder.Configure,
                contentlibrary.Configure,
		tagcategory.Configure,
                tag.Configure,
                hostvirtualswitch.Configure,
                hostportgroup.Configure,
                distributedvirtualswitch.Configure,
                distributedportgroup.Configure,
                vnic.Configure,
                role.Configure,
                entitypermissions.Configure,
                vmfsdatastore.Configure,
                nasdatastore.Configure,
                file.Configure,
                vmstoragepolicy.Configure,
                datastorecluster.Configure,
                computecluster.Configure,
                resourcepool.Configure,
                virtualdisk.Configure,
                vappentity.Configure,
                virtualmachine.Configure,
                virtualmachinesnapshot.Configure,
                vappcontainer.Configure,
                contentlibraryitem.Configure,
		computeclusterhostgroup.Configure,
                computeclustervmaffinityrule.Configure,
                computeclustervmantiaffinityrule.Configure,
		computeclustervmdependencyrule.Configure,
                computeclustervmgroup.Configure,
                computeclustervmhostrule.Configure,
                drsvmoverride.Configure,
                dpmhostoverride.Configure,
		havmoverride.Configure,
                storagedrsvmoverride.Configure,
		datastoreclustervmantiaffinityrule.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
