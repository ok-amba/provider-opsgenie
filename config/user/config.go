package user

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opsgenie_user", func(r *config.Resource) {
		r.ShortGroup = ""
		r.Kind = "User"
	})
}
