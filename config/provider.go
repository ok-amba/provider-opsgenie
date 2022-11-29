/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	"github.com/ok-amba/provider-opsgenie/config/alertpolicy"
	"github.com/ok-amba/provider-opsgenie/config/apiintegration"
	"github.com/ok-amba/provider-opsgenie/config/customrole"
	"github.com/ok-amba/provider-opsgenie/config/emailintegration"
	"github.com/ok-amba/provider-opsgenie/config/escalation"
	"github.com/ok-amba/provider-opsgenie/config/team"
	"github.com/ok-amba/provider-opsgenie/config/teamroutingrule"
)

const (
	resourcePrefix = "opsgenie"
	modulePath     = "github.com/ok-amba/provider-opsgenie"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		team.Configure,
		teamroutingrule.Configure,
		alertpolicy.Configure,
		escalation.Configure,
		apiintegration.Configure,
		customrole.Configure,
		emailintegration.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
