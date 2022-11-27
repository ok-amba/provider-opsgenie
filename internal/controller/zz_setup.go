/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	alertpolicy "github.com/ok-amba/provider-opsgenie/internal/controller/opsgenie/alertpolicy"
	escalation "github.com/ok-amba/provider-opsgenie/internal/controller/opsgenie/escalation"
	team "github.com/ok-amba/provider-opsgenie/internal/controller/opsgenie/team"
	teamroutingrule "github.com/ok-amba/provider-opsgenie/internal/controller/opsgenie/teamroutingrule"
	providerconfig "github.com/ok-amba/provider-opsgenie/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alertpolicy.Setup,
		escalation.Setup,
		team.Setup,
		teamroutingrule.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
