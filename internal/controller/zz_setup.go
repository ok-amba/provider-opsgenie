/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	alertpolicy "github.com/ok-amba/provider-opsgenie/internal/controller/opsgenie/alertpolicy"
	escalation "github.com/ok-amba/provider-opsgenie/internal/controller/opsgenie/escalation"
	providerconfig "github.com/ok-amba/provider-opsgenie/internal/controller/providerconfig"
	team "github.com/ok-amba/provider-opsgenie/internal/controller/team/team"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alertpolicy.Setup,
		escalation.Setup,
		providerconfig.Setup,
		team.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
