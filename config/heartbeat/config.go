package heartbeat

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opsgenie_heartbeat", func(r *config.Resource) {
		r.ShortGroup = ""
		r.Kind = "Heartbeat"

		r.References["owner_team_id"] = config.Reference{
			Type: "Team",
		}
	})
}
