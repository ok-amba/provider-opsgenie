package alertpolicy

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	/* p.AddResourceConfigurator("opsgenie_team", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.ShortGroup = "team"
	}) */
	p.AddResourceConfigurator("opsgenie_alert_policy", func(r *config.Resource) {
		r.ShortGroup = ""
		r.Kind = "AlertPolicy"

		r.References["team_id"] = config.Reference{
			Type: "github.com/ok-amba/provider-opsgenie/apis/team/v1alpha1.Team",
		}

		r.References["responders.id"] = config.Reference{
			Type: "github.com/ok-amba/provider-opsgenie/apis/team/v1alpha1.Team",
		}
	})

}
