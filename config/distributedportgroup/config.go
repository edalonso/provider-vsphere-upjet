package distributedportgroup

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("vsphere_distributed_port_group", func(r *config.Resource) {
        // We need to override the default group that upjet generated for
        // this resource, which would be "github"
        r.ShortGroup = "distributedportgroup"

        r.References["datacenter"] = config.Reference{
                Type: "github.com/edalonso/provider-vsphere-upjet/apis/datacenter/v1alpha1.moid",
       }
    })
}
