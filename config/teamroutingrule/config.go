package teamroutingrule

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opsgenie_team_routing_rule", func(r *config.Resource) {
		r.ShortGroup = ""
		r.Kind = "TeamRoutingRule"

		r.References["team_id"] = config.Reference{
			Type: "Team",
		}

		r.References["notify.id"] = config.Reference{
			Type: "Escalation",
		}
	})
}
