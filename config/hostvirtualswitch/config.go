package hostvirtualswitch

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("vsphere_host_virtual_switch", func(r *config.Resource) {
        // We need to override the default group that upjet generated for
        // this resource, which would be "github"
        r.ShortGroup = "hostvirtualswitch"

        r.References["host"] = config.Reference{
                Type: "github.com/edalonso/provider-vsphere-upjet/apis/host/v1alpha1.Host",
       }
    })
}
