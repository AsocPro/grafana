package dashboard

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

			// TODO: enable v2, currently crashing with an error
			// panic: runtime error: index out of range [0] with length 0
			// goroutine 1 [running]:
			// github.com/grafana/cog/internal/ast/compiler.(*DisjunctionInferMapping).inferDiscriminatorField
			// github.com/grafana/cog@v0.0.18/internal/ast/compiler/disjunctions_infer_mapping.go:114 +0x4d0
			// versions: v0 & v1 & v2

			versions: v0 & v1
		},
	]
}
