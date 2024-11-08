package datastoreclustervmantiaffinityrule

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("vsphere_datastore_cluster_vm_anti_affinity_rule", func(r *config.Resource) {
        // We need to override the default group that upjet generated for
        // this resource, which would be "github"
        r.ShortGroup = "datastoreclustervmantiaffinityrule"
    })
}
