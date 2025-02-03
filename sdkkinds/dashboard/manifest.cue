package dashboard

import (
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
						// For now, we use unstructured for the spec,
						// and it cannot be produced by the SDK codegen.
						spec: {}
					}
				}

				"v1alpha1": {
					schema: {
						// For now, we use unstructured for the spec,
						// and it cannot be produced by the SDK codegen.
						spec: {}
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
