package apiintegration

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opsgenie_api_integration", func(r *config.Resource) {
		r.ShortGroup = ""
		r.Kind = "ApiIntegration"

		r.References["owner_team_id"] = config.Reference{
			Type: "Team",
		}

		r.References["responders.id"] = config.Reference{
			Type: "Team",
		}
	})
}
