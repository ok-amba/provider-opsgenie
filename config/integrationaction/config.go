package integrationaction

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opsgenie_integration_action", func(r *config.Resource) {
		r.ShortGroup = ""
		r.Kind = "IntegrationAction"

		r.References["integration_id"] = config.Reference{
			Type: "ApiIntegration",
		}

		r.References["create.responders.id"] = config.Reference{
			Type: "Team",
		}
	})

}
