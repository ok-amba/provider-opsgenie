package alertpolicy

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opsgenie_alert_policy", func(r *config.Resource) {
		r.ShortGroup = ""
		r.Kind = "AlertPolicy"

		r.References["team_id"] = config.Reference{
			Type: "Team",
		}

		r.References["responders.id"] = config.Reference{
			Type: "Team",
		}
	})

}
