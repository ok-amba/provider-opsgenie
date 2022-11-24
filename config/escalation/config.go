package escalation

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opsgenie_escalation", func(r *config.Resource) {
		r.ShortGroup = ""
		r.Kind = "Escalation"

		r.References["owner_team_id"] = config.Reference{
			Type: "github.com/ok-amba/provider-opsgenie/apis/team/v1alpha1.Team",
		}

		r.References["rules.recipient.id"] = config.Reference{
			Type: "github.com/ok-amba/provider-opsgenie/apis/team/v1alpha1.Team",
		}
	})

}
