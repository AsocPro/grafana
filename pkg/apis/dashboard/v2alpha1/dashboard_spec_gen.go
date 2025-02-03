// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package v2alpha1

// +k8s:openapi-gen=true
type DashboardSpec struct {
	// Title of dashboard.
	Title string `json:"title"`
	// Description of dashboard.
	Description *string `json:"description,omitempty"`
	// Configuration of dashboard cursor sync behavior.
	// "Off" for no shared crosshair or tooltip (default).
	// "Crosshair" for shared crosshair.
	// "Tooltip" for shared crosshair AND shared tooltip.
	CursorSync DashboardDashboardCursorSync `json:"cursorSync"`
	// When set to true, the dashboard will redraw panels at an interval matching the pixel width.
	// This will keep data "moving left" regardless of the query refresh rate. This setting helps
	// avoid dashboards presenting stale live data.
	LiveNow *bool `json:"liveNow,omitempty"`
	// When set to true, the dashboard will load all panels in the dashboard when it's loaded.
	Preload bool `json:"preload"`
	// Whether a dashboard is editable or not.
	Editable *bool `json:"editable,omitempty"`
	// Links with references to other dashboards or external websites.
	Links []DashboardDashboardLink `json:"links"`
	// Tags associated with dashboard.
	Tags         []string                  `json:"tags"`
	TimeSettings DashboardTimeSettingsSpec `json:"timeSettings"`
	// Configured template variables.
	Variables   []DashboardVariableKind        `json:"variables"`
	Elements    map[string]DashboardElement    `json:"elements"`
	Annotations []DashboardAnnotationQueryKind `json:"annotations"`
	Layout      DashboardGridLayoutKind        `json:"layout"`
	// Plugins only. The version of the dashboard installed together with the plugin.
	// This is used to determine if the dashboard should be updated when the plugin is updated.
	Revision *uint16 `json:"revision,omitempty"`
}

// NewDashboardSpec creates a new DashboardSpec object.
func NewDashboardSpec() *DashboardSpec {
	return &DashboardSpec{
		Editable: (func(input bool) *bool { return &input })(true),
	}
}
