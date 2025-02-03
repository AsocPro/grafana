package dashboard

import (
	"github.com/grafana/grafana/sdkkinds/dashboard/v0alpha1"
	"github.com/grafana/grafana/sdkkinds/dashboard/v1alpha1"
	"github.com/grafana/grafana/sdkkinds/dashboard/v2alpha1"
)

manifest: {
	appName:       "dashboard"
	groupOverride: "dashboard.grafana.app"
	kinds: [
		{
			kind:       "Dashboard"
			pluralName: "Dashboards"
			current:    "v0alpha1"

			codegen: {
				frontend: false
				backend:  true
			}

			versions: {
				"v0alpha1": {
					schema: {
						spec: v0alpha1.DashboardSpec
					}
				}

				"v1alpha1": {
					schema: {
						spec: v1alpha1.DashboardSpec
					}
				}

				"v2alpha1": {
					schema: {
						spec: v2alpha1.DashboardSpec
					}
				}
			}
		},
	]
}
