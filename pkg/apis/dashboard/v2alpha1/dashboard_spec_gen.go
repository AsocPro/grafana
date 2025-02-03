// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package v2alpha1

import (
	common "/common"
	json "encoding/json"
	errors "errors"
	fmt "fmt"
)

// Supported dashboard elements
// |* more element types in the future
// +k8s:openapi-gen=true
type DashboardspecElement = DashboardspecPanelKindOrDashboardspecLibraryPanelKind

// NewDashboardspecElement creates a new DashboardspecElement object.
func NewDashboardspecElement() *DashboardspecElement {
	return NewDashboardspecPanelKindOrDashboardspecLibraryPanelKind()
}

// +k8s:openapi-gen=true
type DashboardspecLibraryPanelKind struct {
	Kind string                        `json:"kind"`
	Spec DashboardspecLibraryPanelSpec `json:"spec"`
}

// NewDashboardspecLibraryPanelKind creates a new DashboardspecLibraryPanelKind object.
func NewDashboardspecLibraryPanelKind() *DashboardspecLibraryPanelKind {
	return &DashboardspecLibraryPanelKind{
		Kind: "LibraryPanel",
		Spec: *NewDashboardspecLibraryPanelSpec(),
	}
}

// +k8s:openapi-gen=true
type DashboardspecLibraryPanelSpec struct {
	// Panel ID for the library panel in the dashboard
	Id float64 `json:"id"`
	// Title for the library panel in the dashboard
	Title        string                       `json:"title"`
	LibraryPanel DashboardspecLibraryPanelRef `json:"libraryPanel"`
}

// NewDashboardspecLibraryPanelSpec creates a new DashboardspecLibraryPanelSpec object.
func NewDashboardspecLibraryPanelSpec() *DashboardspecLibraryPanelSpec {
	return &DashboardspecLibraryPanelSpec{
		LibraryPanel: *NewDashboardspecLibraryPanelRef(),
	}
}

// A library panel is a reusable panel that you can use in any dashboard.
// When you make a change to a library panel, that change propagates to all instances of where the panel is used.
// Library panels streamline reuse of panels across multiple dashboards.
// +k8s:openapi-gen=true
type DashboardspecLibraryPanelRef struct {
	// Library panel name
	Name string `json:"name"`
	// Library panel uid
	Uid string `json:"uid"`
}

// NewDashboardspecLibraryPanelRef creates a new DashboardspecLibraryPanelRef object.
func NewDashboardspecLibraryPanelRef() *DashboardspecLibraryPanelRef {
	return &DashboardspecLibraryPanelRef{}
}

// +k8s:openapi-gen=true
type DashboardspecAnnotationPanelFilter struct {
	// Should the specified panels be included or excluded
	Exclude *bool `json:"exclude,omitempty"`
	// Panel IDs that should be included or excluded
	Ids []uint8 `json:"ids"`
}

// NewDashboardspecAnnotationPanelFilter creates a new DashboardspecAnnotationPanelFilter object.
func NewDashboardspecAnnotationPanelFilter() *DashboardspecAnnotationPanelFilter {
	return &DashboardspecAnnotationPanelFilter{
		Exclude: (func(input bool) *bool { return &input })(false),
	}
}

// "Off" for no shared crosshair or tooltip (default).
// "Crosshair" for shared crosshair.
// "Tooltip" for shared crosshair AND shared tooltip.
// +k8s:openapi-gen=true
type DashboardspecDashboardCursorSync string

const (
	DashboardspecDashboardCursorSyncDashboardOff       DashboardspecDashboardCursorSync = "Off"
	DashboardspecDashboardCursorSyncDashboardCrosshair DashboardspecDashboardCursorSync = "Crosshair"
	DashboardspecDashboardCursorSyncDashboardTooltip   DashboardspecDashboardCursorSync = "Tooltip"
)

// Links with references to other dashboards or external resources
// +k8s:openapi-gen=true
type DashboardspecDashboardLink struct {
	// Title to display with the link
	Title string `json:"title"`
	// Link type. Accepted values are dashboards (to refer to another dashboard) and link (to refer to an external resource)
	// FIXME: The type is generated as `type: DashboardLinkType | dashboardLinkType.Link;` but it should be `type: DashboardLinkType`
	Type DashboardspecDashboardLinkType `json:"type"`
	// Icon name to be displayed with the link
	Icon string `json:"icon"`
	// Tooltip to display when the user hovers their mouse over it
	Tooltip string `json:"tooltip"`
	// Link URL. Only required/valid if the type is link
	Url *string `json:"url,omitempty"`
	// List of tags to limit the linked dashboards. If empty, all dashboards will be displayed. Only valid if the type is dashboards
	Tags []string `json:"tags"`
	// If true, all dashboards links will be displayed in a dropdown. If false, all dashboards links will be displayed side by side. Only valid if the type is dashboards
	AsDropdown bool `json:"asDropdown"`
	// If true, the link will be opened in a new tab
	TargetBlank bool `json:"targetBlank"`
	// If true, includes current template variables values in the link as query params
	IncludeVars bool `json:"includeVars"`
	// If true, includes current time range in the link as query params
	KeepTime bool `json:"keepTime"`
}

// NewDashboardspecDashboardLink creates a new DashboardspecDashboardLink object.
func NewDashboardspecDashboardLink() *DashboardspecDashboardLink {
	return &DashboardspecDashboardLink{
		AsDropdown:  false,
		TargetBlank: false,
		IncludeVars: false,
		KeepTime:    false,
	}
}

// +k8s:openapi-gen=true
type DashboardspecDataSourceRef struct {
	// The plugin type-id
	Type *string `json:"type,omitempty"`
	// Specific datasource instance
	Uid *string `json:"uid,omitempty"`
}

// NewDashboardspecDataSourceRef creates a new DashboardspecDataSourceRef object.
func NewDashboardspecDataSourceRef() *DashboardspecDataSourceRef {
	return &DashboardspecDataSourceRef{}
}

// Transformations allow to manipulate data returned by a query before the system applies a visualization.
// Using transformations you can: rename fields, join time series data, perform mathematical operations across queries,
// use the output of one transformation as the input to another transformation, etc.
// +k8s:openapi-gen=true
type DashboardspecDataTransformerConfig struct {
	// Unique identifier of transformer
	Id string `json:"id"`
	// Disabled transformations are skipped
	Disabled *bool `json:"disabled,omitempty"`
	// Optional frame matcher. When missing it will be applied to all results
	Filter *DashboardspecMatcherConfig `json:"filter,omitempty"`
	// Where to pull DataFrames from as input to transformation
	Topic *common.DashboardDataTopic `json:"topic,omitempty"`
	// Options to be passed to the transformer
	// Valid options depend on the transformer id
	Options interface{} `json:"options"`
}

// NewDashboardspecDataTransformerConfig creates a new DashboardspecDataTransformerConfig object.
func NewDashboardspecDataTransformerConfig() *DashboardspecDataTransformerConfig {
	return &DashboardspecDataTransformerConfig{}
}

// +k8s:openapi-gen=true
type DashboardspecDataLink struct {
	Title       string `json:"title"`
	Url         string `json:"url"`
	TargetBlank *bool  `json:"targetBlank,omitempty"`
}

// NewDashboardspecDataLink creates a new DashboardspecDataLink object.
func NewDashboardspecDataLink() *DashboardspecDataLink {
	return &DashboardspecDataLink{}
}

// The data model used in Grafana, namely the data frame, is a columnar-oriented table structure that unifies both time series and table query results.
// Each column within this structure is called a field. A field can represent a single time series or table column.
// Field options allow you to change how the data is displayed in your visualizations.
// +k8s:openapi-gen=true
type DashboardspecFieldConfigSource struct {
	// Defaults are the options applied to all fields.
	Defaults DashboardspecFieldConfig `json:"defaults"`
	// Overrides are the options applied to specific fields overriding the defaults.
	Overrides []V2alpha1DashboardspecFieldConfigSourceOverrides `json:"overrides"`
}

// NewDashboardspecFieldConfigSource creates a new DashboardspecFieldConfigSource object.
func NewDashboardspecFieldConfigSource() *DashboardspecFieldConfigSource {
	return &DashboardspecFieldConfigSource{
		Defaults: *NewDashboardspecFieldConfig(),
	}
}

// The data model used in Grafana, namely the data frame, is a columnar-oriented table structure that unifies both time series and table query results.
// Each column within this structure is called a field. A field can represent a single time series or table column.
// Field options allow you to change how the data is displayed in your visualizations.
// +k8s:openapi-gen=true
type DashboardspecFieldConfig struct {
	// The display value for this field.  This supports template variables blank is auto
	DisplayName *string `json:"displayName,omitempty"`
	// This can be used by data sources that return and explicit naming structure for values and labels
	// When this property is configured, this value is used rather than the default naming strategy.
	DisplayNameFromDS *string `json:"displayNameFromDS,omitempty"`
	// Human readable field metadata
	Description *string `json:"description,omitempty"`
	// An explicit path to the field in the datasource.  When the frame meta includes a path,
	// This will default to `${frame.meta.path}/${field.name}
	//
	// When defined, this value can be used as an identifier within the datasource scope, and
	// may be used to update the results
	Path *string `json:"path,omitempty"`
	// True if data source can write a value to the path. Auth/authz are supported separately
	Writeable *bool `json:"writeable,omitempty"`
	// True if data source field supports ad-hoc filters
	Filterable *bool `json:"filterable,omitempty"`
	// Unit a field should use. The unit you select is applied to all fields except time.
	// You can use the units ID availables in Grafana or a custom unit.
	// Available units in Grafana: https://github.com/grafana/grafana/blob/main/packages/grafana-data/src/valueFormats/categories.ts
	// As custom unit, you can use the following formats:
	// `suffix:<suffix>` for custom unit that should go after value.
	// `prefix:<prefix>` for custom unit that should go before value.
	// `time:<format>` For custom date time formats type for example `time:YYYY-MM-DD`.
	// `si:<base scale><unit characters>` for custom SI units. For example: `si: mF`. This one is a bit more advanced as you can specify both a unit and the source data scale. So if your source data is represented as milli (thousands of) something prefix the unit with that SI scale character.
	// `count:<unit>` for a custom count unit.
	// `currency:<unit>` for custom a currency unit.
	Unit *string `json:"unit,omitempty"`
	// Specify the number of decimals Grafana includes in the rendered value.
	// If you leave this field blank, Grafana automatically truncates the number of decimals based on the value.
	// For example 1.1234 will display as 1.12 and 100.456 will display as 100.
	// To display all decimals, set the unit to `String`.
	Decimals *float64 `json:"decimals,omitempty"`
	// The minimum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
	Min *float64 `json:"min,omitempty"`
	// The maximum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
	Max *float64 `json:"max,omitempty"`
	// Convert input values into a display string
	Mappings []DashboardspecValueMapping `json:"mappings,omitempty"`
	// Map numeric values to states
	Thresholds *DashboardspecThresholdsConfig `json:"thresholds,omitempty"`
	// Panel color configuration
	Color *DashboardspecFieldColor `json:"color,omitempty"`
	// The behavior when clicking on a result
	Links []interface{} `json:"links,omitempty"`
	// Alternative to empty string
	NoValue *string `json:"noValue,omitempty"`
	// custom is specified by the FieldConfig field
	// in panel plugin schemas.
	Custom map[string]interface{} `json:"custom,omitempty"`
}

// NewDashboardspecFieldConfig creates a new DashboardspecFieldConfig object.
func NewDashboardspecFieldConfig() *DashboardspecFieldConfig {
	return &DashboardspecFieldConfig{}
}

// +k8s:openapi-gen=true
type DashboardspecDynamicConfigValue struct {
	Id    string      `json:"id"`
	Value interface{} `json:"value,omitempty"`
}

// NewDashboardspecDynamicConfigValue creates a new DashboardspecDynamicConfigValue object.
func NewDashboardspecDynamicConfigValue() *DashboardspecDynamicConfigValue {
	return &DashboardspecDynamicConfigValue{
		Id: "",
	}
}

// Matcher is a predicate configuration. Based on the config a set of field(s) or values is filtered in order to apply override / transformation.
// It comes with in id ( to resolve implementation from registry) and a configuration that’s specific to a particular matcher type.
// +k8s:openapi-gen=true
type DashboardspecMatcherConfig struct {
	// The matcher id. This is used to find the matcher implementation from registry.
	Id string `json:"id"`
	// The matcher options. This is specific to the matcher implementation.
	Options interface{} `json:"options,omitempty"`
}

// NewDashboardspecMatcherConfig creates a new DashboardspecMatcherConfig object.
func NewDashboardspecMatcherConfig() *DashboardspecMatcherConfig {
	return &DashboardspecMatcherConfig{
		Id: "",
	}
}

// +k8s:openapi-gen=true
type DashboardspecThreshold struct {
	Value float64 `json:"value"`
	Color string  `json:"color"`
}

// NewDashboardspecThreshold creates a new DashboardspecThreshold object.
func NewDashboardspecThreshold() *DashboardspecThreshold {
	return &DashboardspecThreshold{}
}

// +k8s:openapi-gen=true
type DashboardspecThresholdsMode string

const (
	DashboardspecThresholdsModeDashboardAbsolute   DashboardspecThresholdsMode = "absolute"
	DashboardspecThresholdsModeDashboardPercentage DashboardspecThresholdsMode = "percentage"
)

// +k8s:openapi-gen=true
type DashboardspecThresholdsConfig struct {
	Mode  DashboardspecThresholdsMode `json:"mode"`
	Steps []DashboardspecThreshold    `json:"steps"`
}

// NewDashboardspecThresholdsConfig creates a new DashboardspecThresholdsConfig object.
func NewDashboardspecThresholdsConfig() *DashboardspecThresholdsConfig {
	return &DashboardspecThresholdsConfig{}
}

// +k8s:openapi-gen=true
type DashboardspecValueMapping = DashboardspecValueMapOrDashboardspecRangeMapOrDashboardspecRegexMapOrDashboardspecSpecialValueMap

// NewDashboardspecValueMapping creates a new DashboardspecValueMapping object.
func NewDashboardspecValueMapping() *DashboardspecValueMapping {
	return NewDashboardspecValueMapOrDashboardspecRangeMapOrDashboardspecRegexMapOrDashboardspecSpecialValueMap()
}

// Supported value mapping types
// `value`: Maps text values to a color or different display text and color. For example, you can configure a value mapping so that all instances of the value 10 appear as Perfection! rather than the number.
// `range`: Maps numerical ranges to a display text and color. For example, if a value is within a certain range, you can configure a range value mapping to display Low or High rather than the number.
// `regex`: Maps regular expressions to replacement text and a color. For example, if a value is www.example.com, you can configure a regex value mapping so that Grafana displays www and truncates the domain.
// `special`: Maps special values like Null, NaN (not a number), and boolean values like true and false to a display text and color. See SpecialValueMatch to see the list of special values. For example, you can configure a special value mapping so that null values appear as N/A.
// +k8s:openapi-gen=true
type DashboardspecMappingType string

const (
	DashboardspecMappingTypeDashboardValueToText  DashboardspecMappingType = "value"
	DashboardspecMappingTypeDashboardRangeToText  DashboardspecMappingType = "range"
	DashboardspecMappingTypeDashboardRegexToText  DashboardspecMappingType = "regex"
	DashboardspecMappingTypeDashboardSpecialValue DashboardspecMappingType = "special"
)

// Maps text values to a color or different display text and color.
// For example, you can configure a value mapping so that all instances of the value 10 appear as Perfection! rather than the number.
// +k8s:openapi-gen=true
type DashboardspecValueMap struct {
	Type string `json:"type"`
	// Map with <value_to_match>: ValueMappingResult. For example: { "10": { text: "Perfection!", color: "green" } }
	Options map[string]DashboardspecValueMappingResult `json:"options"`
}

// NewDashboardspecValueMap creates a new DashboardspecValueMap object.
func NewDashboardspecValueMap() *DashboardspecValueMap {
	return &DashboardspecValueMap{
		Type: "value",
	}
}

// Maps numerical ranges to a display text and color.
// For example, if a value is within a certain range, you can configure a range value mapping to display Low or High rather than the number.
// +k8s:openapi-gen=true
type DashboardspecRangeMap struct {
	Type string `json:"type"`
	// Range to match against and the result to apply when the value is within the range
	Options V2alpha1DashboardspecRangeMapOptions `json:"options"`
}

// NewDashboardspecRangeMap creates a new DashboardspecRangeMap object.
func NewDashboardspecRangeMap() *DashboardspecRangeMap {
	return &DashboardspecRangeMap{
		Type:    "range",
		Options: *NewV2alpha1DashboardspecRangeMapOptions(),
	}
}

// Maps regular expressions to replacement text and a color.
// For example, if a value is www.example.com, you can configure a regex value mapping so that Grafana displays www and truncates the domain.
// +k8s:openapi-gen=true
type DashboardspecRegexMap struct {
	Type string `json:"type"`
	// Regular expression to match against and the result to apply when the value matches the regex
	Options V2alpha1DashboardspecRegexMapOptions `json:"options"`
}

// NewDashboardspecRegexMap creates a new DashboardspecRegexMap object.
func NewDashboardspecRegexMap() *DashboardspecRegexMap {
	return &DashboardspecRegexMap{
		Type:    "regex",
		Options: *NewV2alpha1DashboardspecRegexMapOptions(),
	}
}

// Maps special values like Null, NaN (not a number), and boolean values like true and false to a display text and color.
// See SpecialValueMatch to see the list of special values.
// For example, you can configure a special value mapping so that null values appear as N/A.
// +k8s:openapi-gen=true
type DashboardspecSpecialValueMap struct {
	Type    string                                      `json:"type"`
	Options V2alpha1DashboardspecSpecialValueMapOptions `json:"options"`
}

// NewDashboardspecSpecialValueMap creates a new DashboardspecSpecialValueMap object.
func NewDashboardspecSpecialValueMap() *DashboardspecSpecialValueMap {
	return &DashboardspecSpecialValueMap{
		Type:    "special",
		Options: *NewV2alpha1DashboardspecSpecialValueMapOptions(),
	}
}

// Special value types supported by the `SpecialValueMap`
// +k8s:openapi-gen=true
type DashboardspecSpecialValueMatch string

const (
	DashboardspecSpecialValueMatchDashboardTrue       DashboardspecSpecialValueMatch = "true"
	DashboardspecSpecialValueMatchDashboardFalse      DashboardspecSpecialValueMatch = "false"
	DashboardspecSpecialValueMatchDashboardNull       DashboardspecSpecialValueMatch = "null"
	DashboardspecSpecialValueMatchDashboardNaN        DashboardspecSpecialValueMatch = "nan"
	DashboardspecSpecialValueMatchDashboardNullAndNaN DashboardspecSpecialValueMatch = "null+nan"
	DashboardspecSpecialValueMatchDashboardEmpty      DashboardspecSpecialValueMatch = "empty"
)

// Result used as replacement with text and color when the value matches
// +k8s:openapi-gen=true
type DashboardspecValueMappingResult struct {
	// Text to display when the value matches
	Text *string `json:"text,omitempty"`
	// Text to use when the value matches
	Color *string `json:"color,omitempty"`
	// Icon to display when the value matches. Only specific visualizations.
	Icon *string `json:"icon,omitempty"`
	// Position in the mapping array. Only used internally.
	Index *int32 `json:"index,omitempty"`
}

// NewDashboardspecValueMappingResult creates a new DashboardspecValueMappingResult object.
func NewDashboardspecValueMappingResult() *DashboardspecValueMappingResult {
	return &DashboardspecValueMappingResult{}
}

// Color mode for a field. You can specify a single color, or select a continuous (gradient) color schemes, based on a value.
// Continuous color interpolates a color using the percentage of a value relative to min and max.
// Accepted values are:
// `thresholds`: From thresholds. Informs Grafana to take the color from the matching threshold
// `palette-classic`: Classic palette. Grafana will assign color by looking up a color in a palette by series index. Useful for Graphs and pie charts and other categorical data visualizations
// `palette-classic-by-name`: Classic palette (by name). Grafana will assign color by looking up a color in a palette by series name. Useful for Graphs and pie charts and other categorical data visualizations
// `continuous-GrYlRd`: ontinuous Green-Yellow-Red palette mode
// `continuous-RdYlGr`: Continuous Red-Yellow-Green palette mode
// `continuous-BlYlRd`: Continuous Blue-Yellow-Red palette mode
// `continuous-YlRd`: Continuous Yellow-Red palette mode
// `continuous-BlPu`: Continuous Blue-Purple palette mode
// `continuous-YlBl`: Continuous Yellow-Blue palette mode
// `continuous-blues`: Continuous Blue palette mode
// `continuous-reds`: Continuous Red palette mode
// `continuous-greens`: Continuous Green palette mode
// `continuous-purples`: Continuous Purple palette mode
// `shades`: Shades of a single color. Specify a single color, useful in an override rule.
// `fixed`: Fixed color mode. Specify a single color, useful in an override rule.
// +k8s:openapi-gen=true
type DashboardspecFieldColorModeId string

const (
	DashboardspecFieldColorModeIdDashboardThresholds           DashboardspecFieldColorModeId = "thresholds"
	DashboardspecFieldColorModeIdDashboardPaletteClassic       DashboardspecFieldColorModeId = "palette-classic"
	DashboardspecFieldColorModeIdDashboardPaletteClassicByName DashboardspecFieldColorModeId = "palette-classic-by-name"
	DashboardspecFieldColorModeIdDashboardContinuousGrYlRd     DashboardspecFieldColorModeId = "continuous-GrYlRd"
	DashboardspecFieldColorModeIdDashboardContinuousRdYlGr     DashboardspecFieldColorModeId = "continuous-RdYlGr"
	DashboardspecFieldColorModeIdDashboardContinuousBlYlRd     DashboardspecFieldColorModeId = "continuous-BlYlRd"
	DashboardspecFieldColorModeIdDashboardContinuousYlRd       DashboardspecFieldColorModeId = "continuous-YlRd"
	DashboardspecFieldColorModeIdDashboardContinuousBlPu       DashboardspecFieldColorModeId = "continuous-BlPu"
	DashboardspecFieldColorModeIdDashboardContinuousYlBl       DashboardspecFieldColorModeId = "continuous-YlBl"
	DashboardspecFieldColorModeIdDashboardContinuousBlues      DashboardspecFieldColorModeId = "continuous-blues"
	DashboardspecFieldColorModeIdDashboardContinuousReds       DashboardspecFieldColorModeId = "continuous-reds"
	DashboardspecFieldColorModeIdDashboardContinuousGreens     DashboardspecFieldColorModeId = "continuous-greens"
	DashboardspecFieldColorModeIdDashboardContinuousPurples    DashboardspecFieldColorModeId = "continuous-purples"
	DashboardspecFieldColorModeIdDashboardFixed                DashboardspecFieldColorModeId = "fixed"
	DashboardspecFieldColorModeIdDashboardShades               DashboardspecFieldColorModeId = "shades"
)

// Defines how to assign a series color from "by value" color schemes. For example for an aggregated data points like a timeseries, the color can be assigned by the min, max or last value.
// +k8s:openapi-gen=true
type DashboardspecFieldColorSeriesByMode string

const (
	DashboardspecFieldColorSeriesByModeDashboardMin  DashboardspecFieldColorSeriesByMode = "min"
	DashboardspecFieldColorSeriesByModeDashboardMax  DashboardspecFieldColorSeriesByMode = "max"
	DashboardspecFieldColorSeriesByModeDashboardLast DashboardspecFieldColorSeriesByMode = "last"
)

// Map a field to a color.
// +k8s:openapi-gen=true
type DashboardspecFieldColor struct {
	// The main color scheme mode.
	Mode DashboardspecFieldColorModeId `json:"mode"`
	// The fixed color value for fixed or shades color modes.
	FixedColor *string `json:"fixedColor,omitempty"`
	// Some visualizations need to know how to assign a series color from by value color schemes.
	SeriesBy *DashboardspecFieldColorSeriesByMode `json:"seriesBy,omitempty"`
}

// NewDashboardspecFieldColor creates a new DashboardspecFieldColor object.
func NewDashboardspecFieldColor() *DashboardspecFieldColor {
	return &DashboardspecFieldColor{}
}

// Dashboard Link type. Accepted values are dashboards (to refer to another dashboard) and link (to refer to an external resource)
// +k8s:openapi-gen=true
type DashboardspecDashboardLinkType string

const (
	DashboardspecDashboardLinkTypeDashboardLink       DashboardspecDashboardLinkType = "link"
	DashboardspecDashboardLinkTypeDashboardDashboards DashboardspecDashboardLinkType = "dashboards"
)

// --- Common types ---
// +k8s:openapi-gen=true
type DashboardspecKind struct {
	Kind     string      `json:"kind"`
	Spec     interface{} `json:"spec"`
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewDashboardspecKind creates a new DashboardspecKind object.
func NewDashboardspecKind() *DashboardspecKind {
	return &DashboardspecKind{}
}

// --- Kinds ---
// +k8s:openapi-gen=true
type DashboardspecVizConfigSpec struct {
	PluginVersion string                         `json:"pluginVersion"`
	Options       map[string]interface{}         `json:"options"`
	FieldConfig   DashboardspecFieldConfigSource `json:"fieldConfig"`
}

// NewDashboardspecVizConfigSpec creates a new DashboardspecVizConfigSpec object.
func NewDashboardspecVizConfigSpec() *DashboardspecVizConfigSpec {
	return &DashboardspecVizConfigSpec{
		FieldConfig: *NewDashboardspecFieldConfigSource(),
	}
}

// +k8s:openapi-gen=true
type DashboardspecVizConfigKind struct {
	// The kind of a VizConfigKind is the plugin ID
	Kind string                     `json:"kind"`
	Spec DashboardspecVizConfigSpec `json:"spec"`
}

// NewDashboardspecVizConfigKind creates a new DashboardspecVizConfigKind object.
func NewDashboardspecVizConfigKind() *DashboardspecVizConfigKind {
	return &DashboardspecVizConfigKind{
		Spec: *NewDashboardspecVizConfigSpec(),
	}
}

// +k8s:openapi-gen=true
type DashboardspecAnnotationQuerySpec struct {
	Datasource *DashboardspecDataSourceRef         `json:"datasource,omitempty"`
	Query      *DashboardspecDataQueryKind         `json:"query,omitempty"`
	Enable     bool                                `json:"enable"`
	Hide       bool                                `json:"hide"`
	IconColor  string                              `json:"iconColor"`
	Name       string                              `json:"name"`
	BuiltIn    *bool                               `json:"builtIn,omitempty"`
	Filter     *DashboardspecAnnotationPanelFilter `json:"filter,omitempty"`
}

// NewDashboardspecAnnotationQuerySpec creates a new DashboardspecAnnotationQuerySpec object.
func NewDashboardspecAnnotationQuerySpec() *DashboardspecAnnotationQuerySpec {
	return &DashboardspecAnnotationQuerySpec{
		BuiltIn: (func(input bool) *bool { return &input })(false),
	}
}

// +k8s:openapi-gen=true
type DashboardspecAnnotationQueryKind struct {
	Kind string                           `json:"kind"`
	Spec DashboardspecAnnotationQuerySpec `json:"spec"`
}

// NewDashboardspecAnnotationQueryKind creates a new DashboardspecAnnotationQueryKind object.
func NewDashboardspecAnnotationQueryKind() *DashboardspecAnnotationQueryKind {
	return &DashboardspecAnnotationQueryKind{
		Kind: "AnnotationQuery",
		Spec: *NewDashboardspecAnnotationQuerySpec(),
	}
}

// +k8s:openapi-gen=true
type DashboardspecQueryOptionsSpec struct {
	TimeFrom         *string `json:"timeFrom,omitempty"`
	MaxDataPoints    *int64  `json:"maxDataPoints,omitempty"`
	TimeShift        *string `json:"timeShift,omitempty"`
	QueryCachingTTL  *int64  `json:"queryCachingTTL,omitempty"`
	Interval         *string `json:"interval,omitempty"`
	CacheTimeout     *string `json:"cacheTimeout,omitempty"`
	HideTimeOverride *bool   `json:"hideTimeOverride,omitempty"`
}

// NewDashboardspecQueryOptionsSpec creates a new DashboardspecQueryOptionsSpec object.
func NewDashboardspecQueryOptionsSpec() *DashboardspecQueryOptionsSpec {
	return &DashboardspecQueryOptionsSpec{}
}

// +k8s:openapi-gen=true
type DashboardspecDataQueryKind struct {
	// The kind of a DataQueryKind is the datasource type
	Kind string                 `json:"kind"`
	Spec map[string]interface{} `json:"spec"`
}

// NewDashboardspecDataQueryKind creates a new DashboardspecDataQueryKind object.
func NewDashboardspecDataQueryKind() *DashboardspecDataQueryKind {
	return &DashboardspecDataQueryKind{}
}

// +k8s:openapi-gen=true
type DashboardspecPanelQuerySpec struct {
	Query      DashboardspecDataQueryKind  `json:"query"`
	Datasource *DashboardspecDataSourceRef `json:"datasource,omitempty"`
	RefId      string                      `json:"refId"`
	Hidden     bool                        `json:"hidden"`
}

// NewDashboardspecPanelQuerySpec creates a new DashboardspecPanelQuerySpec object.
func NewDashboardspecPanelQuerySpec() *DashboardspecPanelQuerySpec {
	return &DashboardspecPanelQuerySpec{
		Query: *NewDashboardspecDataQueryKind(),
	}
}

// +k8s:openapi-gen=true
type DashboardspecPanelQueryKind struct {
	Kind string                      `json:"kind"`
	Spec DashboardspecPanelQuerySpec `json:"spec"`
}

// NewDashboardspecPanelQueryKind creates a new DashboardspecPanelQueryKind object.
func NewDashboardspecPanelQueryKind() *DashboardspecPanelQueryKind {
	return &DashboardspecPanelQueryKind{
		Kind: "PanelQuery",
		Spec: *NewDashboardspecPanelQuerySpec(),
	}
}

// +k8s:openapi-gen=true
type DashboardspecTransformationKind struct {
	// The kind of a TransformationKind is the transformation ID
	Kind string                             `json:"kind"`
	Spec DashboardspecDataTransformerConfig `json:"spec"`
}

// NewDashboardspecTransformationKind creates a new DashboardspecTransformationKind object.
func NewDashboardspecTransformationKind() *DashboardspecTransformationKind {
	return &DashboardspecTransformationKind{
		Spec: *NewDashboardspecDataTransformerConfig(),
	}
}

// +k8s:openapi-gen=true
type DashboardspecQueryGroupSpec struct {
	Queries         []DashboardspecPanelQueryKind     `json:"queries"`
	Transformations []DashboardspecTransformationKind `json:"transformations"`
	QueryOptions    DashboardspecQueryOptionsSpec     `json:"queryOptions"`
}

// NewDashboardspecQueryGroupSpec creates a new DashboardspecQueryGroupSpec object.
func NewDashboardspecQueryGroupSpec() *DashboardspecQueryGroupSpec {
	return &DashboardspecQueryGroupSpec{
		QueryOptions: *NewDashboardspecQueryOptionsSpec(),
	}
}

// +k8s:openapi-gen=true
type DashboardspecQueryGroupKind struct {
	Kind string                      `json:"kind"`
	Spec DashboardspecQueryGroupSpec `json:"spec"`
}

// NewDashboardspecQueryGroupKind creates a new DashboardspecQueryGroupKind object.
func NewDashboardspecQueryGroupKind() *DashboardspecQueryGroupKind {
	return &DashboardspecQueryGroupKind{
		Kind: "QueryGroup",
		Spec: *NewDashboardspecQueryGroupSpec(),
	}
}

// Time configuration
// It defines the default time config for the time picker, the refresh picker for the specific dashboard.
// +k8s:openapi-gen=true
type DashboardspecTimeSettingsSpec struct {
	// Timezone of dashboard. Accepted values are IANA TZDB zone ID or "browser" or "utc".
	Timezone *string `json:"timezone,omitempty"`
	// Start time range for dashboard.
	// Accepted values are relative time strings like "now-6h" or absolute time strings like "2020-07-10T08:00:00.000Z".
	From string `json:"from"`
	// End time range for dashboard.
	// Accepted values are relative time strings like "now-6h" or absolute time strings like "2020-07-10T08:00:00.000Z".
	To string `json:"to"`
	// Refresh rate of dashboard. Represented via interval string, e.g. "5s", "1m", "1h", "1d".
	// v1: refresh
	AutoRefresh string `json:"autoRefresh"`
	// Interval options available in the refresh picker dropdown.
	// v1: timepicker.refresh_intervals
	AutoRefreshIntervals []string `json:"autoRefreshIntervals"`
	// Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.
	// v1: timepicker.time_options , not exposed in the UI
	QuickRanges []string `json:"quickRanges"`
	// Whether timepicker is visible or not.
	// v1: timepicker.hidden
	HideTimepicker bool `json:"hideTimepicker"`
	// Day when the week starts. Expressed by the name of the day in lowercase, e.g. "monday".
	WeekStart string `json:"weekStart"`
	// The month that the fiscal year starts on. 0 = January, 11 = December
	FiscalYearStartMonth int64 `json:"fiscalYearStartMonth"`
	// Override the now time by entering a time delay. Use this option to accommodate known delays in data aggregation to avoid null values.
	// v1: timepicker.nowDelay
	NowDelay *string `json:"nowDelay,omitempty"`
}

// NewDashboardspecTimeSettingsSpec creates a new DashboardspecTimeSettingsSpec object.
func NewDashboardspecTimeSettingsSpec() *DashboardspecTimeSettingsSpec {
	return &DashboardspecTimeSettingsSpec{
		Timezone:             (func(input string) *string { return &input })("browser"),
		From:                 "now-6h",
		To:                   "now",
		AutoRefreshIntervals: []string{"5s", "10s", "30s", "1m", "5m", "15m", "30m", "1h", "2h", "1d"},
		QuickRanges:          []string{"5m", "15m", "1h", "6h", "12h", "24h", "2d", "7d", "30d"},
	}
}

// other repeat modes will be added in the future: label, frame
// +k8s:openapi-gen=true
const DashboardspecRepeatMode = "variable"

// +k8s:openapi-gen=true
type DashboardspecRepeatOptions struct {
	Mode      string                               `json:"mode"`
	Value     string                               `json:"value"`
	Direction *DashboardspecRepeatOptionsDirection `json:"direction,omitempty"`
	MaxPerRow *int64                               `json:"maxPerRow,omitempty"`
}

// NewDashboardspecRepeatOptions creates a new DashboardspecRepeatOptions object.
func NewDashboardspecRepeatOptions() *DashboardspecRepeatOptions {
	return &DashboardspecRepeatOptions{}
}

// +k8s:openapi-gen=true
type DashboardspecRowRepeatOptions struct {
	Mode  string `json:"mode"`
	Value string `json:"value"`
}

// NewDashboardspecRowRepeatOptions creates a new DashboardspecRowRepeatOptions object.
func NewDashboardspecRowRepeatOptions() *DashboardspecRowRepeatOptions {
	return &DashboardspecRowRepeatOptions{}
}

// +k8s:openapi-gen=true
type DashboardspecGridLayoutItemSpec struct {
	X      int64 `json:"x"`
	Y      int64 `json:"y"`
	Width  int64 `json:"width"`
	Height int64 `json:"height"`
	// reference to a PanelKind from dashboard.spec.elements Expressed as JSON Schema reference
	Element DashboardspecElementReference `json:"element"`
	Repeat  *DashboardspecRepeatOptions   `json:"repeat,omitempty"`
}

// NewDashboardspecGridLayoutItemSpec creates a new DashboardspecGridLayoutItemSpec object.
func NewDashboardspecGridLayoutItemSpec() *DashboardspecGridLayoutItemSpec {
	return &DashboardspecGridLayoutItemSpec{
		Element: *NewDashboardspecElementReference(),
	}
}

// +k8s:openapi-gen=true
type DashboardspecGridLayoutItemKind struct {
	Kind string                          `json:"kind"`
	Spec DashboardspecGridLayoutItemSpec `json:"spec"`
}

// NewDashboardspecGridLayoutItemKind creates a new DashboardspecGridLayoutItemKind object.
func NewDashboardspecGridLayoutItemKind() *DashboardspecGridLayoutItemKind {
	return &DashboardspecGridLayoutItemKind{
		Kind: "GridLayoutItem",
		Spec: *NewDashboardspecGridLayoutItemSpec(),
	}
}

// +k8s:openapi-gen=true
type DashboardspecGridLayoutRowKind struct {
	Kind string                         `json:"kind"`
	Spec DashboardspecGridLayoutRowSpec `json:"spec"`
}

// NewDashboardspecGridLayoutRowKind creates a new DashboardspecGridLayoutRowKind object.
func NewDashboardspecGridLayoutRowKind() *DashboardspecGridLayoutRowKind {
	return &DashboardspecGridLayoutRowKind{
		Kind: "GridLayoutRow",
		Spec: *NewDashboardspecGridLayoutRowSpec(),
	}
}

// +k8s:openapi-gen=true
type DashboardspecGridLayoutRowSpec struct {
	Y         int64  `json:"y"`
	Collapsed bool   `json:"collapsed"`
	Title     string `json:"title"`
	// Grid items in the row will have their Y value be relative to the rows Y value. This means a panel positioned at Y: 0 in a row with Y: 10 will be positioned at Y: 11 (row header has a heigh of 1) in the dashboard.
	Elements []DashboardspecGridLayoutItemKind `json:"elements"`
	Repeat   *DashboardspecRowRepeatOptions    `json:"repeat,omitempty"`
}

// NewDashboardspecGridLayoutRowSpec creates a new DashboardspecGridLayoutRowSpec object.
func NewDashboardspecGridLayoutRowSpec() *DashboardspecGridLayoutRowSpec {
	return &DashboardspecGridLayoutRowSpec{}
}

// +k8s:openapi-gen=true
type DashboardspecGridLayoutSpec struct {
	Items []DashboardspecGridLayoutItemKindOrDashboardspecGridLayoutRowKind `json:"items"`
}

// NewDashboardspecGridLayoutSpec creates a new DashboardspecGridLayoutSpec object.
func NewDashboardspecGridLayoutSpec() *DashboardspecGridLayoutSpec {
	return &DashboardspecGridLayoutSpec{}
}

// +k8s:openapi-gen=true
type DashboardspecGridLayoutKind struct {
	Kind string                      `json:"kind"`
	Spec DashboardspecGridLayoutSpec `json:"spec"`
}

// NewDashboardspecGridLayoutKind creates a new DashboardspecGridLayoutKind object.
func NewDashboardspecGridLayoutKind() *DashboardspecGridLayoutKind {
	return &DashboardspecGridLayoutKind{
		Kind: "GridLayout",
		Spec: *NewDashboardspecGridLayoutSpec(),
	}
}

// +k8s:openapi-gen=true
type DashboardspecPanelSpec struct {
	Id          float64                     `json:"id"`
	Title       string                      `json:"title"`
	Description string                      `json:"description"`
	Links       []DashboardspecDataLink     `json:"links"`
	Data        DashboardspecQueryGroupKind `json:"data"`
	VizConfig   DashboardspecVizConfigKind  `json:"vizConfig"`
	Transparent *bool                       `json:"transparent,omitempty"`
}

// NewDashboardspecPanelSpec creates a new DashboardspecPanelSpec object.
func NewDashboardspecPanelSpec() *DashboardspecPanelSpec {
	return &DashboardspecPanelSpec{
		Data:      *NewDashboardspecQueryGroupKind(),
		VizConfig: *NewDashboardspecVizConfigKind(),
	}
}

// +k8s:openapi-gen=true
type DashboardspecPanelKind struct {
	Kind string                 `json:"kind"`
	Spec DashboardspecPanelSpec `json:"spec"`
}

// NewDashboardspecPanelKind creates a new DashboardspecPanelKind object.
func NewDashboardspecPanelKind() *DashboardspecPanelKind {
	return &DashboardspecPanelKind{
		Kind: "Panel",
		Spec: *NewDashboardspecPanelSpec(),
	}
}

// +k8s:openapi-gen=true
type DashboardspecElementReference struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
}

// NewDashboardspecElementReference creates a new DashboardspecElementReference object.
func NewDashboardspecElementReference() *DashboardspecElementReference {
	return &DashboardspecElementReference{
		Kind: "ElementReference",
	}
}

// Variable types
// +k8s:openapi-gen=true
type DashboardspecVariableValue = StringOrBoolOrFloat64OrDashboardspecCustomVariableValueOrArrayOfDashboardspecVariableValueSingle

// NewDashboardspecVariableValue creates a new DashboardspecVariableValue object.
func NewDashboardspecVariableValue() *DashboardspecVariableValue {
	return NewStringOrBoolOrFloat64OrDashboardspecCustomVariableValueOrArrayOfDashboardspecVariableValueSingle()
}

// +k8s:openapi-gen=true
type DashboardspecVariableValueSingle = StringOrBoolOrFloat64OrDashboardspecCustomVariableValue

// NewDashboardspecVariableValueSingle creates a new DashboardspecVariableValueSingle object.
func NewDashboardspecVariableValueSingle() *DashboardspecVariableValueSingle {
	return NewStringOrBoolOrFloat64OrDashboardspecCustomVariableValue()
}

// Custom formatter variable
// +k8s:openapi-gen=true
type DashboardspecCustomFormatterVariable struct {
	Name       string                    `json:"name"`
	Type       DashboardspecVariableType `json:"type"`
	Multi      bool                      `json:"multi"`
	IncludeAll bool                      `json:"includeAll"`
}

// NewDashboardspecCustomFormatterVariable creates a new DashboardspecCustomFormatterVariable object.
func NewDashboardspecCustomFormatterVariable() *DashboardspecCustomFormatterVariable {
	return &DashboardspecCustomFormatterVariable{}
}

// Custom variable value
// +k8s:openapi-gen=true
type DashboardspecCustomVariableValue struct {
	// The format name or function used in the expression
	Formatter StringOrDashboardspecVariableCustomFormatterFn `json:"formatter"`
}

// NewDashboardspecCustomVariableValue creates a new DashboardspecCustomVariableValue object.
func NewDashboardspecCustomVariableValue() *DashboardspecCustomVariableValue {
	return &DashboardspecCustomVariableValue{
		Formatter: *NewStringOrDashboardspecVariableCustomFormatterFn(),
	}
}

// Custom formatter function
// +k8s:openapi-gen=true
type DashboardspecVariableCustomFormatterFn struct {
	Value                  interface{}                                                       `json:"value"`
	LegacyVariableModel    V2alpha1DashboardspecVariableCustomFormatterFnLegacyVariableModel `json:"legacyVariableModel"`
	LegacyDefaultFormatter *DashboardspecVariableCustomFormatterFn                           `json:"legacyDefaultFormatter,omitempty"`
}

// NewDashboardspecVariableCustomFormatterFn creates a new DashboardspecVariableCustomFormatterFn object.
func NewDashboardspecVariableCustomFormatterFn() *DashboardspecVariableCustomFormatterFn {
	return &DashboardspecVariableCustomFormatterFn{
		LegacyVariableModel: *NewV2alpha1DashboardspecVariableCustomFormatterFnLegacyVariableModel(),
	}
}

// Dashboard variable type
// `query`: Query-generated list of values such as metric names, server names, sensor IDs, data centers, and so on.
// `adhoc`: Key/value filters that are automatically added to all metric queries for a data source (Prometheus, Loki, InfluxDB, and Elasticsearch only).
// `constant`: 	Define a hidden constant.
// `datasource`: Quickly change the data source for an entire dashboard.
// `interval`: Interval variables represent time spans.
// `textbox`: Display a free text input field with an optional default value.
// `custom`: Define the variable options manually using a comma-separated list.
// `system`: Variables defined by Grafana. See: https://grafana.com/docs/grafana/latest/dashboards/variables/add-template-variables/#global-variables
// +k8s:openapi-gen=true
type DashboardspecVariableType string

const (
	DashboardspecVariableTypeDashboardQuery      DashboardspecVariableType = "query"
	DashboardspecVariableTypeDashboardAdhoc      DashboardspecVariableType = "adhoc"
	DashboardspecVariableTypeDashboardGroupby    DashboardspecVariableType = "groupby"
	DashboardspecVariableTypeDashboardConstant   DashboardspecVariableType = "constant"
	DashboardspecVariableTypeDashboardDatasource DashboardspecVariableType = "datasource"
	DashboardspecVariableTypeDashboardInterval   DashboardspecVariableType = "interval"
	DashboardspecVariableTypeDashboardTextbox    DashboardspecVariableType = "textbox"
	DashboardspecVariableTypeDashboardCustom     DashboardspecVariableType = "custom"
	DashboardspecVariableTypeDashboardSystem     DashboardspecVariableType = "system"
	DashboardspecVariableTypeDashboardSnapshot   DashboardspecVariableType = "snapshot"
)

// +k8s:openapi-gen=true
type DashboardspecVariableKind = DashboardspecQueryVariableKindOrDashboardspecTextVariableKindOrDashboardspecConstantVariableKindOrDashboardspecDatasourceVariableKindOrDashboardspecIntervalVariableKindOrDashboardspecCustomVariableKindOrDashboardspecGroupByVariableKindOrDashboardspecAdhocVariableKind

// NewDashboardspecVariableKind creates a new DashboardspecVariableKind object.
func NewDashboardspecVariableKind() *DashboardspecVariableKind {
	return NewDashboardspecQueryVariableKindOrDashboardspecTextVariableKindOrDashboardspecConstantVariableKindOrDashboardspecDatasourceVariableKindOrDashboardspecIntervalVariableKindOrDashboardspecCustomVariableKindOrDashboardspecGroupByVariableKindOrDashboardspecAdhocVariableKind()
}

// Sort variable options
// Accepted values are:
// `disabled`: No sorting
// `alphabeticalAsc`: Alphabetical ASC
// `alphabeticalDesc`: Alphabetical DESC
// `numericalAsc`: Numerical ASC
// `numericalDesc`: Numerical DESC
// `alphabeticalCaseInsensitiveAsc`: Alphabetical Case Insensitive ASC
// `alphabeticalCaseInsensitiveDesc`: Alphabetical Case Insensitive DESC
// `naturalAsc`: Natural ASC
// `naturalDesc`: Natural DESC
// VariableSort enum with default value
// +k8s:openapi-gen=true
type DashboardspecVariableSort string

const (
	DashboardspecVariableSortDashboardDisabled                        DashboardspecVariableSort = "disabled"
	DashboardspecVariableSortDashboardAlphabeticalAsc                 DashboardspecVariableSort = "alphabeticalAsc"
	DashboardspecVariableSortDashboardAlphabeticalDesc                DashboardspecVariableSort = "alphabeticalDesc"
	DashboardspecVariableSortDashboardNumericalAsc                    DashboardspecVariableSort = "numericalAsc"
	DashboardspecVariableSortDashboardNumericalDesc                   DashboardspecVariableSort = "numericalDesc"
	DashboardspecVariableSortDashboardAlphabeticalCaseInsensitiveAsc  DashboardspecVariableSort = "alphabeticalCaseInsensitiveAsc"
	DashboardspecVariableSortDashboardAlphabeticalCaseInsensitiveDesc DashboardspecVariableSort = "alphabeticalCaseInsensitiveDesc"
	DashboardspecVariableSortDashboardNaturalAsc                      DashboardspecVariableSort = "naturalAsc"
	DashboardspecVariableSortDashboardNaturalDesc                     DashboardspecVariableSort = "naturalDesc"
)

// Options to config when to refresh a variable
// `never`: Never refresh the variable
// `onDashboardLoad`: Queries the data source every time the dashboard loads.
// `onTimeRangeChanged`: Queries the data source when the dashboard time range changes.
// +k8s:openapi-gen=true
type DashboardspecVariableRefresh string

const (
	DashboardspecVariableRefreshDashboardNever              DashboardspecVariableRefresh = "never"
	DashboardspecVariableRefreshDashboardOnDashboardLoad    DashboardspecVariableRefresh = "onDashboardLoad"
	DashboardspecVariableRefreshDashboardOnTimeRangeChanged DashboardspecVariableRefresh = "onTimeRangeChanged"
)

// Determine if the variable shows on dashboard
// Accepted values are `dontHide` (show label and value), `hideLabel` (show value only), `hideVariable` (show nothing).
// +k8s:openapi-gen=true
type DashboardspecVariableHide string

const (
	DashboardspecVariableHideDashboardDontHide     DashboardspecVariableHide = "dontHide"
	DashboardspecVariableHideDashboardHideLabel    DashboardspecVariableHide = "hideLabel"
	DashboardspecVariableHideDashboardHideVariable DashboardspecVariableHide = "hideVariable"
)

// FIXME: should we introduce this? --- Variable value option
// +k8s:openapi-gen=true
type DashboardspecVariableValueOption struct {
	Label string                           `json:"label"`
	Value DashboardspecVariableValueSingle `json:"value"`
	Group *string                          `json:"group,omitempty"`
}

// NewDashboardspecVariableValueOption creates a new DashboardspecVariableValueOption object.
func NewDashboardspecVariableValueOption() *DashboardspecVariableValueOption {
	return &DashboardspecVariableValueOption{
		Value: *NewDashboardspecVariableValueSingle(),
	}
}

// Variable option specification
// +k8s:openapi-gen=true
type DashboardspecVariableOption struct {
	// Whether the option is selected or not
	Selected *bool `json:"selected,omitempty"`
	// Text to be displayed for the option
	Text StringOrArrayOfString `json:"text"`
	// Value of the option
	Value StringOrArrayOfString `json:"value"`
}

// NewDashboardspecVariableOption creates a new DashboardspecVariableOption object.
func NewDashboardspecVariableOption() *DashboardspecVariableOption {
	return &DashboardspecVariableOption{
		Text:  *NewStringOrArrayOfString(),
		Value: *NewStringOrArrayOfString(),
	}
}

// Query variable specification
// +k8s:openapi-gen=true
type DashboardspecQueryVariableSpec struct {
	Name        string                             `json:"name"`
	Current     DashboardspecVariableOption        `json:"current"`
	Label       *string                            `json:"label,omitempty"`
	Hide        DashboardspecVariableHide          `json:"hide"`
	Refresh     DashboardspecVariableRefresh       `json:"refresh"`
	SkipUrlSync bool                               `json:"skipUrlSync"`
	Description *string                            `json:"description,omitempty"`
	Datasource  *DashboardspecDataSourceRef        `json:"datasource,omitempty"`
	Query       StringOrDashboardspecDataQueryKind `json:"query"`
	Regex       string                             `json:"regex"`
	Sort        DashboardspecVariableSort          `json:"sort"`
	Definition  *string                            `json:"definition,omitempty"`
	Options     []DashboardspecVariableOption      `json:"options"`
	Multi       bool                               `json:"multi"`
	IncludeAll  bool                               `json:"includeAll"`
	AllValue    *string                            `json:"allValue,omitempty"`
	Placeholder *string                            `json:"placeholder,omitempty"`
}

// NewDashboardspecQueryVariableSpec creates a new DashboardspecQueryVariableSpec object.
func NewDashboardspecQueryVariableSpec() *DashboardspecQueryVariableSpec {
	return &DashboardspecQueryVariableSpec{
		Name: "",
		Current: DashboardspecVariableOption{
			Text: StringOrArrayOfString{
				String: (func(input string) *string { return &input })(""),
			},
			Value: StringOrArrayOfString{
				String: (func(input string) *string { return &input })(""),
			},
		},
		Hide:        DashboardspecVariableHideDashboardDontHide,
		Refresh:     DashboardspecVariableRefreshDashboardNever,
		SkipUrlSync: false,
		Query:       *NewStringOrDashboardspecDataQueryKind(),
		Regex:       "",
		Multi:       false,
		IncludeAll:  false,
	}
}

// Query variable kind
// +k8s:openapi-gen=true
type DashboardspecQueryVariableKind struct {
	Kind string                         `json:"kind"`
	Spec DashboardspecQueryVariableSpec `json:"spec"`
}

// NewDashboardspecQueryVariableKind creates a new DashboardspecQueryVariableKind object.
func NewDashboardspecQueryVariableKind() *DashboardspecQueryVariableKind {
	return &DashboardspecQueryVariableKind{
		Kind: "QueryVariable",
		Spec: *NewDashboardspecQueryVariableSpec(),
	}
}

// Text variable specification
// +k8s:openapi-gen=true
type DashboardspecTextVariableSpec struct {
	Name        string                      `json:"name"`
	Current     DashboardspecVariableOption `json:"current"`
	Query       string                      `json:"query"`
	Label       *string                     `json:"label,omitempty"`
	Hide        DashboardspecVariableHide   `json:"hide"`
	SkipUrlSync bool                        `json:"skipUrlSync"`
	Description *string                     `json:"description,omitempty"`
}

// NewDashboardspecTextVariableSpec creates a new DashboardspecTextVariableSpec object.
func NewDashboardspecTextVariableSpec() *DashboardspecTextVariableSpec {
	return &DashboardspecTextVariableSpec{
		Name: "",
		Current: DashboardspecVariableOption{
			Text: StringOrArrayOfString{
				String: (func(input string) *string { return &input })(""),
			},
			Value: StringOrArrayOfString{
				String: (func(input string) *string { return &input })(""),
			},
		},
		Query:       "",
		Hide:        DashboardspecVariableHideDashboardDontHide,
		SkipUrlSync: false,
	}
}

// Text variable kind
// +k8s:openapi-gen=true
type DashboardspecTextVariableKind struct {
	Kind string                        `json:"kind"`
	Spec DashboardspecTextVariableSpec `json:"spec"`
}

// NewDashboardspecTextVariableKind creates a new DashboardspecTextVariableKind object.
func NewDashboardspecTextVariableKind() *DashboardspecTextVariableKind {
	return &DashboardspecTextVariableKind{
		Kind: "TextVariable",
		Spec: *NewDashboardspecTextVariableSpec(),
	}
}

// Constant variable specification
// +k8s:openapi-gen=true
type DashboardspecConstantVariableSpec struct {
	Name        string                      `json:"name"`
	Query       string                      `json:"query"`
	Current     DashboardspecVariableOption `json:"current"`
	Label       *string                     `json:"label,omitempty"`
	Hide        DashboardspecVariableHide   `json:"hide"`
	SkipUrlSync bool                        `json:"skipUrlSync"`
	Description *string                     `json:"description,omitempty"`
}

// NewDashboardspecConstantVariableSpec creates a new DashboardspecConstantVariableSpec object.
func NewDashboardspecConstantVariableSpec() *DashboardspecConstantVariableSpec {
	return &DashboardspecConstantVariableSpec{
		Name:  "",
		Query: "",
		Current: DashboardspecVariableOption{
			Text: StringOrArrayOfString{
				String: (func(input string) *string { return &input })(""),
			},
			Value: StringOrArrayOfString{
				String: (func(input string) *string { return &input })(""),
			},
		},
		Hide:        DashboardspecVariableHideDashboardDontHide,
		SkipUrlSync: false,
	}
}

// Constant variable kind
// +k8s:openapi-gen=true
type DashboardspecConstantVariableKind struct {
	Kind string                            `json:"kind"`
	Spec DashboardspecConstantVariableSpec `json:"spec"`
}

// NewDashboardspecConstantVariableKind creates a new DashboardspecConstantVariableKind object.
func NewDashboardspecConstantVariableKind() *DashboardspecConstantVariableKind {
	return &DashboardspecConstantVariableKind{
		Kind: "ConstantVariable",
		Spec: *NewDashboardspecConstantVariableSpec(),
	}
}

// Datasource variable specification
// +k8s:openapi-gen=true
type DashboardspecDatasourceVariableSpec struct {
	Name        string                        `json:"name"`
	PluginId    string                        `json:"pluginId"`
	Refresh     DashboardspecVariableRefresh  `json:"refresh"`
	Regex       string                        `json:"regex"`
	Current     DashboardspecVariableOption   `json:"current"`
	Options     []DashboardspecVariableOption `json:"options"`
	Multi       bool                          `json:"multi"`
	IncludeAll  bool                          `json:"includeAll"`
	AllValue    *string                       `json:"allValue,omitempty"`
	Label       *string                       `json:"label,omitempty"`
	Hide        DashboardspecVariableHide     `json:"hide"`
	SkipUrlSync bool                          `json:"skipUrlSync"`
	Description *string                       `json:"description,omitempty"`
}

// NewDashboardspecDatasourceVariableSpec creates a new DashboardspecDatasourceVariableSpec object.
func NewDashboardspecDatasourceVariableSpec() *DashboardspecDatasourceVariableSpec {
	return &DashboardspecDatasourceVariableSpec{
		Name:     "",
		PluginId: "",
		Refresh:  DashboardspecVariableRefreshDashboardNever,
		Regex:    "",
		Current: DashboardspecVariableOption{
			Text: StringOrArrayOfString{
				String: (func(input string) *string { return &input })(""),
			},
			Value: StringOrArrayOfString{
				String: (func(input string) *string { return &input })(""),
			},
		},
		Multi:       false,
		IncludeAll:  false,
		Hide:        DashboardspecVariableHideDashboardDontHide,
		SkipUrlSync: false,
	}
}

// Datasource variable kind
// +k8s:openapi-gen=true
type DashboardspecDatasourceVariableKind struct {
	Kind string                              `json:"kind"`
	Spec DashboardspecDatasourceVariableSpec `json:"spec"`
}

// NewDashboardspecDatasourceVariableKind creates a new DashboardspecDatasourceVariableKind object.
func NewDashboardspecDatasourceVariableKind() *DashboardspecDatasourceVariableKind {
	return &DashboardspecDatasourceVariableKind{
		Kind: "DatasourceVariable",
		Spec: *NewDashboardspecDatasourceVariableSpec(),
	}
}

// Interval variable specification
// +k8s:openapi-gen=true
type DashboardspecIntervalVariableSpec struct {
	Name        string                        `json:"name"`
	Query       string                        `json:"query"`
	Current     DashboardspecVariableOption   `json:"current"`
	Options     []DashboardspecVariableOption `json:"options"`
	Auto        bool                          `json:"auto"`
	AutoMin     string                        `json:"auto_min"`
	AutoCount   int64                         `json:"auto_count"`
	Refresh     DashboardspecVariableRefresh  `json:"refresh"`
	Label       *string                       `json:"label,omitempty"`
	Hide        DashboardspecVariableHide     `json:"hide"`
	SkipUrlSync bool                          `json:"skipUrlSync"`
	Description *string                       `json:"description,omitempty"`
}

// NewDashboardspecIntervalVariableSpec creates a new DashboardspecIntervalVariableSpec object.
func NewDashboardspecIntervalVariableSpec() *DashboardspecIntervalVariableSpec {
	return &DashboardspecIntervalVariableSpec{
		Name:  "",
		Query: "",
		Current: DashboardspecVariableOption{
			Text: StringOrArrayOfString{
				String: (func(input string) *string { return &input })(""),
			},
			Value: StringOrArrayOfString{
				String: (func(input string) *string { return &input })(""),
			},
		},
		Auto:        false,
		AutoMin:     "",
		AutoCount:   0,
		Refresh:     DashboardspecVariableRefreshDashboardNever,
		Hide:        DashboardspecVariableHideDashboardDontHide,
		SkipUrlSync: false,
	}
}

// Interval variable kind
// +k8s:openapi-gen=true
type DashboardspecIntervalVariableKind struct {
	Kind string                            `json:"kind"`
	Spec DashboardspecIntervalVariableSpec `json:"spec"`
}

// NewDashboardspecIntervalVariableKind creates a new DashboardspecIntervalVariableKind object.
func NewDashboardspecIntervalVariableKind() *DashboardspecIntervalVariableKind {
	return &DashboardspecIntervalVariableKind{
		Kind: "IntervalVariable",
		Spec: *NewDashboardspecIntervalVariableSpec(),
	}
}

// Custom variable specification
// +k8s:openapi-gen=true
type DashboardspecCustomVariableSpec struct {
	Name        string                        `json:"name"`
	Query       string                        `json:"query"`
	Current     DashboardspecVariableOption   `json:"current"`
	Options     []DashboardspecVariableOption `json:"options"`
	Multi       bool                          `json:"multi"`
	IncludeAll  bool                          `json:"includeAll"`
	AllValue    *string                       `json:"allValue,omitempty"`
	Label       *string                       `json:"label,omitempty"`
	Hide        DashboardspecVariableHide     `json:"hide"`
	SkipUrlSync bool                          `json:"skipUrlSync"`
	Description *string                       `json:"description,omitempty"`
}

// NewDashboardspecCustomVariableSpec creates a new DashboardspecCustomVariableSpec object.
func NewDashboardspecCustomVariableSpec() *DashboardspecCustomVariableSpec {
	return &DashboardspecCustomVariableSpec{
		Name:        "",
		Query:       "",
		Current:     *NewDashboardspecVariableOption(),
		Multi:       false,
		IncludeAll:  false,
		Hide:        DashboardspecVariableHideDashboardDontHide,
		SkipUrlSync: false,
	}
}

// Custom variable kind
// +k8s:openapi-gen=true
type DashboardspecCustomVariableKind struct {
	Kind string                          `json:"kind"`
	Spec DashboardspecCustomVariableSpec `json:"spec"`
}

// NewDashboardspecCustomVariableKind creates a new DashboardspecCustomVariableKind object.
func NewDashboardspecCustomVariableKind() *DashboardspecCustomVariableKind {
	return &DashboardspecCustomVariableKind{
		Kind: "CustomVariable",
		Spec: *NewDashboardspecCustomVariableSpec(),
	}
}

// GroupBy variable specification
// +k8s:openapi-gen=true
type DashboardspecGroupByVariableSpec struct {
	Name        string                        `json:"name"`
	Datasource  *DashboardspecDataSourceRef   `json:"datasource,omitempty"`
	Current     DashboardspecVariableOption   `json:"current"`
	Options     []DashboardspecVariableOption `json:"options"`
	Multi       bool                          `json:"multi"`
	Label       *string                       `json:"label,omitempty"`
	Hide        DashboardspecVariableHide     `json:"hide"`
	SkipUrlSync bool                          `json:"skipUrlSync"`
	Description *string                       `json:"description,omitempty"`
}

// NewDashboardspecGroupByVariableSpec creates a new DashboardspecGroupByVariableSpec object.
func NewDashboardspecGroupByVariableSpec() *DashboardspecGroupByVariableSpec {
	return &DashboardspecGroupByVariableSpec{
		Name: "",
		Current: DashboardspecVariableOption{
			Text: StringOrArrayOfString{
				String: (func(input string) *string { return &input })(""),
			},
			Value: StringOrArrayOfString{
				String: (func(input string) *string { return &input })(""),
			},
		},
		Multi:       false,
		Hide:        DashboardspecVariableHideDashboardDontHide,
		SkipUrlSync: false,
	}
}

// Group variable kind
// +k8s:openapi-gen=true
type DashboardspecGroupByVariableKind struct {
	Kind string                           `json:"kind"`
	Spec DashboardspecGroupByVariableSpec `json:"spec"`
}

// NewDashboardspecGroupByVariableKind creates a new DashboardspecGroupByVariableKind object.
func NewDashboardspecGroupByVariableKind() *DashboardspecGroupByVariableKind {
	return &DashboardspecGroupByVariableKind{
		Kind: "GroupByVariable",
		Spec: *NewDashboardspecGroupByVariableSpec(),
	}
}

// Adhoc variable specification
// +k8s:openapi-gen=true
type DashboardspecAdhocVariableSpec struct {
	Name        string                               `json:"name"`
	Datasource  *DashboardspecDataSourceRef          `json:"datasource,omitempty"`
	BaseFilters []DashboardspecAdHocFilterWithLabels `json:"baseFilters"`
	Filters     []DashboardspecAdHocFilterWithLabels `json:"filters"`
	DefaultKeys []DashboardspecMetricFindValue       `json:"defaultKeys"`
	Label       *string                              `json:"label,omitempty"`
	Hide        DashboardspecVariableHide            `json:"hide"`
	SkipUrlSync bool                                 `json:"skipUrlSync"`
	Description *string                              `json:"description,omitempty"`
}

// NewDashboardspecAdhocVariableSpec creates a new DashboardspecAdhocVariableSpec object.
func NewDashboardspecAdhocVariableSpec() *DashboardspecAdhocVariableSpec {
	return &DashboardspecAdhocVariableSpec{
		Name:        "",
		Hide:        DashboardspecVariableHideDashboardDontHide,
		SkipUrlSync: false,
	}
}

// Define the MetricFindValue type
// +k8s:openapi-gen=true
type DashboardspecMetricFindValue struct {
	Text       string           `json:"text"`
	Value      *StringOrFloat64 `json:"value,omitempty"`
	Group      *string          `json:"group,omitempty"`
	Expandable *bool            `json:"expandable,omitempty"`
}

// NewDashboardspecMetricFindValue creates a new DashboardspecMetricFindValue object.
func NewDashboardspecMetricFindValue() *DashboardspecMetricFindValue {
	return &DashboardspecMetricFindValue{}
}

// Define the AdHocFilterWithLabels type
// +k8s:openapi-gen=true
type DashboardspecAdHocFilterWithLabels struct {
	Key         string   `json:"key"`
	Operator    string   `json:"operator"`
	Value       string   `json:"value"`
	Values      []string `json:"values,omitempty"`
	KeyLabel    *string  `json:"keyLabel,omitempty"`
	ValueLabels []string `json:"valueLabels,omitempty"`
	ForceEdit   *bool    `json:"forceEdit,omitempty"`
	// @deprecated
	Condition *string `json:"condition,omitempty"`
}

// NewDashboardspecAdHocFilterWithLabels creates a new DashboardspecAdHocFilterWithLabels object.
func NewDashboardspecAdHocFilterWithLabels() *DashboardspecAdHocFilterWithLabels {
	return &DashboardspecAdHocFilterWithLabels{}
}

// Adhoc variable kind
// +k8s:openapi-gen=true
type DashboardspecAdhocVariableKind struct {
	Kind string                         `json:"kind"`
	Spec DashboardspecAdhocVariableSpec `json:"spec"`
}

// NewDashboardspecAdhocVariableKind creates a new DashboardspecAdhocVariableKind object.
func NewDashboardspecAdhocVariableKind() *DashboardspecAdhocVariableKind {
	return &DashboardspecAdhocVariableKind{
		Kind: "AdhocVariable",
		Spec: *NewDashboardspecAdhocVariableSpec(),
	}
}

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
	CursorSync DashboardspecDashboardCursorSync `json:"cursorSync"`
	// When set to true, the dashboard will redraw panels at an interval matching the pixel width.
	// This will keep data "moving left" regardless of the query refresh rate. This setting helps
	// avoid dashboards presenting stale live data.
	LiveNow *bool `json:"liveNow,omitempty"`
	// When set to true, the dashboard will load all panels in the dashboard when it's loaded.
	Preload bool `json:"preload"`
	// Whether a dashboard is editable or not.
	Editable *bool `json:"editable,omitempty"`
	// Links with references to other dashboards or external websites.
	Links []DashboardspecDashboardLink `json:"links"`
	// Tags associated with dashboard.
	Tags         []string                      `json:"tags"`
	TimeSettings DashboardspecTimeSettingsSpec `json:"timeSettings"`
	// Configured template variables.
	Variables   []DashboardspecVariableKind        `json:"variables"`
	Elements    map[string]DashboardspecElement    `json:"elements"`
	Annotations []DashboardspecAnnotationQueryKind `json:"annotations"`
	Layout      DashboardspecGridLayoutKind        `json:"layout"`
	// Plugins only. The version of the dashboard installed together with the plugin.
	// This is used to determine if the dashboard should be updated when the plugin is updated.
	Revision *uint16 `json:"revision,omitempty"`
}

// NewDashboardSpec creates a new DashboardSpec object.
func NewDashboardSpec() *DashboardSpec {
	return &DashboardSpec{
		Editable:     (func(input bool) *bool { return &input })(true),
		TimeSettings: *NewDashboardspecTimeSettingsSpec(),
		Layout:       *NewDashboardspecGridLayoutKind(),
	}
}

type DashboardspecRepeatOptionsDirection string

const (
	DashboardspecRepeatOptionsDirectionDashboardH DashboardspecRepeatOptionsDirection = "h"
	DashboardspecRepeatOptionsDirectionDashboardV DashboardspecRepeatOptionsDirection = "v"
)

type V2alpha1DashboardspecFieldConfigSourceOverrides struct {
	Matcher    DashboardspecMatcherConfig        `json:"matcher"`
	Properties []DashboardspecDynamicConfigValue `json:"properties"`
}

// NewV2alpha1DashboardspecFieldConfigSourceOverrides creates a new V2alpha1DashboardspecFieldConfigSourceOverrides object.
func NewV2alpha1DashboardspecFieldConfigSourceOverrides() *V2alpha1DashboardspecFieldConfigSourceOverrides {
	return &V2alpha1DashboardspecFieldConfigSourceOverrides{
		Matcher: *NewDashboardspecMatcherConfig(),
	}
}

type V2alpha1DashboardspecRangeMapOptions struct {
	// Min value of the range. It can be null which means -Infinity
	From *float64 `json:"from"`
	// Max value of the range. It can be null which means +Infinity
	To *float64 `json:"to"`
	// Config to apply when the value is within the range
	Result DashboardspecValueMappingResult `json:"result"`
}

// NewV2alpha1DashboardspecRangeMapOptions creates a new V2alpha1DashboardspecRangeMapOptions object.
func NewV2alpha1DashboardspecRangeMapOptions() *V2alpha1DashboardspecRangeMapOptions {
	return &V2alpha1DashboardspecRangeMapOptions{
		Result: *NewDashboardspecValueMappingResult(),
	}
}

type V2alpha1DashboardspecRegexMapOptions struct {
	// Regular expression to match against
	Pattern string `json:"pattern"`
	// Config to apply when the value matches the regex
	Result DashboardspecValueMappingResult `json:"result"`
}

// NewV2alpha1DashboardspecRegexMapOptions creates a new V2alpha1DashboardspecRegexMapOptions object.
func NewV2alpha1DashboardspecRegexMapOptions() *V2alpha1DashboardspecRegexMapOptions {
	return &V2alpha1DashboardspecRegexMapOptions{
		Result: *NewDashboardspecValueMappingResult(),
	}
}

type V2alpha1DashboardspecSpecialValueMapOptions struct {
	// Special value to match against
	Match DashboardspecSpecialValueMatch `json:"match"`
	// Config to apply when the value matches the special value
	Result DashboardspecValueMappingResult `json:"result"`
}

// NewV2alpha1DashboardspecSpecialValueMapOptions creates a new V2alpha1DashboardspecSpecialValueMapOptions object.
func NewV2alpha1DashboardspecSpecialValueMapOptions() *V2alpha1DashboardspecSpecialValueMapOptions {
	return &V2alpha1DashboardspecSpecialValueMapOptions{
		Result: *NewDashboardspecValueMappingResult(),
	}
}

type V2alpha1DashboardspecVariableCustomFormatterFnLegacyVariableModel struct {
	Name       string                    `json:"name"`
	Type       DashboardspecVariableType `json:"type"`
	Multi      bool                      `json:"multi"`
	IncludeAll bool                      `json:"includeAll"`
}

// NewV2alpha1DashboardspecVariableCustomFormatterFnLegacyVariableModel creates a new V2alpha1DashboardspecVariableCustomFormatterFnLegacyVariableModel object.
func NewV2alpha1DashboardspecVariableCustomFormatterFnLegacyVariableModel() *V2alpha1DashboardspecVariableCustomFormatterFnLegacyVariableModel {
	return &V2alpha1DashboardspecVariableCustomFormatterFnLegacyVariableModel{}
}

type DashboardspecPanelKindOrDashboardspecLibraryPanelKind struct {
	DashboardspecPanelKind        *DashboardspecPanelKind        `json:"DashboardspecPanelKind,omitempty"`
	DashboardspecLibraryPanelKind *DashboardspecLibraryPanelKind `json:"DashboardspecLibraryPanelKind,omitempty"`
}

// NewDashboardspecPanelKindOrDashboardspecLibraryPanelKind creates a new DashboardspecPanelKindOrDashboardspecLibraryPanelKind object.
func NewDashboardspecPanelKindOrDashboardspecLibraryPanelKind() *DashboardspecPanelKindOrDashboardspecLibraryPanelKind {
	return &DashboardspecPanelKindOrDashboardspecLibraryPanelKind{}
}

// MarshalJSON implements a custom JSON marshalling logic to encode `DashboardspecPanelKindOrDashboardspecLibraryPanelKind` as JSON.
func (resource DashboardspecPanelKindOrDashboardspecLibraryPanelKind) MarshalJSON() ([]byte, error) {
	if resource.DashboardspecPanelKind != nil {
		return json.Marshal(resource.DashboardspecPanelKind)
	}
	if resource.DashboardspecLibraryPanelKind != nil {
		return json.Marshal(resource.DashboardspecLibraryPanelKind)
	}

	return nil, fmt.Errorf("no value for disjunction of refs")
}

// UnmarshalJSON implements a custom JSON unmarshalling logic to decode `DashboardspecPanelKindOrDashboardspecLibraryPanelKind` from JSON.
func (resource *DashboardspecPanelKindOrDashboardspecLibraryPanelKind) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]interface{})
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["kind"]
	if !found {
		return errors.New("discriminator field 'kind' not found in payload")
	}

	switch discriminator {
	case "LibraryPanel":
		var dashboardspecLibraryPanelKind DashboardspecLibraryPanelKind
		if err := json.Unmarshal(raw, &dashboardspecLibraryPanelKind); err != nil {
			return err
		}

		resource.DashboardspecLibraryPanelKind = &dashboardspecLibraryPanelKind
		return nil
	case "Panel":
		var dashboardspecPanelKind DashboardspecPanelKind
		if err := json.Unmarshal(raw, &dashboardspecPanelKind); err != nil {
			return err
		}

		resource.DashboardspecPanelKind = &dashboardspecPanelKind
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `kind = %v`", discriminator)
}

type DashboardspecValueMapOrDashboardspecRangeMapOrDashboardspecRegexMapOrDashboardspecSpecialValueMap struct {
	DashboardspecValueMap        *DashboardspecValueMap        `json:"DashboardspecValueMap,omitempty"`
	DashboardspecRangeMap        *DashboardspecRangeMap        `json:"DashboardspecRangeMap,omitempty"`
	DashboardspecRegexMap        *DashboardspecRegexMap        `json:"DashboardspecRegexMap,omitempty"`
	DashboardspecSpecialValueMap *DashboardspecSpecialValueMap `json:"DashboardspecSpecialValueMap,omitempty"`
}

// NewDashboardspecValueMapOrDashboardspecRangeMapOrDashboardspecRegexMapOrDashboardspecSpecialValueMap creates a new DashboardspecValueMapOrDashboardspecRangeMapOrDashboardspecRegexMapOrDashboardspecSpecialValueMap object.
func NewDashboardspecValueMapOrDashboardspecRangeMapOrDashboardspecRegexMapOrDashboardspecSpecialValueMap() *DashboardspecValueMapOrDashboardspecRangeMapOrDashboardspecRegexMapOrDashboardspecSpecialValueMap {
	return &DashboardspecValueMapOrDashboardspecRangeMapOrDashboardspecRegexMapOrDashboardspecSpecialValueMap{}
}

// MarshalJSON implements a custom JSON marshalling logic to encode `DashboardspecValueMapOrDashboardspecRangeMapOrDashboardspecRegexMapOrDashboardspecSpecialValueMap` as JSON.
func (resource DashboardspecValueMapOrDashboardspecRangeMapOrDashboardspecRegexMapOrDashboardspecSpecialValueMap) MarshalJSON() ([]byte, error) {
	if resource.DashboardspecValueMap != nil {
		return json.Marshal(resource.DashboardspecValueMap)
	}
	if resource.DashboardspecRangeMap != nil {
		return json.Marshal(resource.DashboardspecRangeMap)
	}
	if resource.DashboardspecRegexMap != nil {
		return json.Marshal(resource.DashboardspecRegexMap)
	}
	if resource.DashboardspecSpecialValueMap != nil {
		return json.Marshal(resource.DashboardspecSpecialValueMap)
	}

	return nil, fmt.Errorf("no value for disjunction of refs")
}

// UnmarshalJSON implements a custom JSON unmarshalling logic to decode `DashboardspecValueMapOrDashboardspecRangeMapOrDashboardspecRegexMapOrDashboardspecSpecialValueMap` from JSON.
func (resource *DashboardspecValueMapOrDashboardspecRangeMapOrDashboardspecRegexMapOrDashboardspecSpecialValueMap) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]interface{})
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["type"]
	if !found {
		return errors.New("discriminator field 'type' not found in payload")
	}

	switch discriminator {
	case "range":
		var dashboardspecRangeMap DashboardspecRangeMap
		if err := json.Unmarshal(raw, &dashboardspecRangeMap); err != nil {
			return err
		}

		resource.DashboardspecRangeMap = &dashboardspecRangeMap
		return nil
	case "regex":
		var dashboardspecRegexMap DashboardspecRegexMap
		if err := json.Unmarshal(raw, &dashboardspecRegexMap); err != nil {
			return err
		}

		resource.DashboardspecRegexMap = &dashboardspecRegexMap
		return nil
	case "special":
		var dashboardspecSpecialValueMap DashboardspecSpecialValueMap
		if err := json.Unmarshal(raw, &dashboardspecSpecialValueMap); err != nil {
			return err
		}

		resource.DashboardspecSpecialValueMap = &dashboardspecSpecialValueMap
		return nil
	case "value":
		var dashboardspecValueMap DashboardspecValueMap
		if err := json.Unmarshal(raw, &dashboardspecValueMap); err != nil {
			return err
		}

		resource.DashboardspecValueMap = &dashboardspecValueMap
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `type = %v`", discriminator)
}

type DashboardspecGridLayoutItemKindOrDashboardspecGridLayoutRowKind struct {
	DashboardspecGridLayoutItemKind *DashboardspecGridLayoutItemKind `json:"DashboardspecGridLayoutItemKind,omitempty"`
	DashboardspecGridLayoutRowKind  *DashboardspecGridLayoutRowKind  `json:"DashboardspecGridLayoutRowKind,omitempty"`
}

// NewDashboardspecGridLayoutItemKindOrDashboardspecGridLayoutRowKind creates a new DashboardspecGridLayoutItemKindOrDashboardspecGridLayoutRowKind object.
func NewDashboardspecGridLayoutItemKindOrDashboardspecGridLayoutRowKind() *DashboardspecGridLayoutItemKindOrDashboardspecGridLayoutRowKind {
	return &DashboardspecGridLayoutItemKindOrDashboardspecGridLayoutRowKind{}
}

// MarshalJSON implements a custom JSON marshalling logic to encode `DashboardspecGridLayoutItemKindOrDashboardspecGridLayoutRowKind` as JSON.
func (resource DashboardspecGridLayoutItemKindOrDashboardspecGridLayoutRowKind) MarshalJSON() ([]byte, error) {
	if resource.DashboardspecGridLayoutItemKind != nil {
		return json.Marshal(resource.DashboardspecGridLayoutItemKind)
	}
	if resource.DashboardspecGridLayoutRowKind != nil {
		return json.Marshal(resource.DashboardspecGridLayoutRowKind)
	}

	return nil, fmt.Errorf("no value for disjunction of refs")
}

// UnmarshalJSON implements a custom JSON unmarshalling logic to decode `DashboardspecGridLayoutItemKindOrDashboardspecGridLayoutRowKind` from JSON.
func (resource *DashboardspecGridLayoutItemKindOrDashboardspecGridLayoutRowKind) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]interface{})
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["kind"]
	if !found {
		return errors.New("discriminator field 'kind' not found in payload")
	}

	switch discriminator {
	case "GridLayoutItem":
		var dashboardspecGridLayoutItemKind DashboardspecGridLayoutItemKind
		if err := json.Unmarshal(raw, &dashboardspecGridLayoutItemKind); err != nil {
			return err
		}

		resource.DashboardspecGridLayoutItemKind = &dashboardspecGridLayoutItemKind
		return nil
	case "GridLayoutRow":
		var dashboardspecGridLayoutRowKind DashboardspecGridLayoutRowKind
		if err := json.Unmarshal(raw, &dashboardspecGridLayoutRowKind); err != nil {
			return err
		}

		resource.DashboardspecGridLayoutRowKind = &dashboardspecGridLayoutRowKind
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `kind = %v`", discriminator)
}

type StringOrBoolOrFloat64OrDashboardspecCustomVariableValueOrArrayOfDashboardspecVariableValueSingle struct {
	String                                  *string                            `json:"String,omitempty"`
	Bool                                    *bool                              `json:"Bool,omitempty"`
	Float64                                 *float64                           `json:"Float64,omitempty"`
	DashboardspecCustomVariableValue        *DashboardspecCustomVariableValue  `json:"DashboardspecCustomVariableValue,omitempty"`
	ArrayOfDashboardspecVariableValueSingle []DashboardspecVariableValueSingle `json:"ArrayOfDashboardspecVariableValueSingle,omitempty"`
}

// NewStringOrBoolOrFloat64OrDashboardspecCustomVariableValueOrArrayOfDashboardspecVariableValueSingle creates a new StringOrBoolOrFloat64OrDashboardspecCustomVariableValueOrArrayOfDashboardspecVariableValueSingle object.
func NewStringOrBoolOrFloat64OrDashboardspecCustomVariableValueOrArrayOfDashboardspecVariableValueSingle() *StringOrBoolOrFloat64OrDashboardspecCustomVariableValueOrArrayOfDashboardspecVariableValueSingle {
	return &StringOrBoolOrFloat64OrDashboardspecCustomVariableValueOrArrayOfDashboardspecVariableValueSingle{}
}

type StringOrBoolOrFloat64OrDashboardspecCustomVariableValue struct {
	String                           *string                           `json:"String,omitempty"`
	Bool                             *bool                             `json:"Bool,omitempty"`
	Float64                          *float64                          `json:"Float64,omitempty"`
	DashboardspecCustomVariableValue *DashboardspecCustomVariableValue `json:"DashboardspecCustomVariableValue,omitempty"`
}

// NewStringOrBoolOrFloat64OrDashboardspecCustomVariableValue creates a new StringOrBoolOrFloat64OrDashboardspecCustomVariableValue object.
func NewStringOrBoolOrFloat64OrDashboardspecCustomVariableValue() *StringOrBoolOrFloat64OrDashboardspecCustomVariableValue {
	return &StringOrBoolOrFloat64OrDashboardspecCustomVariableValue{}
}

type StringOrDashboardspecVariableCustomFormatterFn struct {
	String                                 *string                                 `json:"String,omitempty"`
	DashboardspecVariableCustomFormatterFn *DashboardspecVariableCustomFormatterFn `json:"DashboardspecVariableCustomFormatterFn,omitempty"`
}

// NewStringOrDashboardspecVariableCustomFormatterFn creates a new StringOrDashboardspecVariableCustomFormatterFn object.
func NewStringOrDashboardspecVariableCustomFormatterFn() *StringOrDashboardspecVariableCustomFormatterFn {
	return &StringOrDashboardspecVariableCustomFormatterFn{}
}

type DashboardspecQueryVariableKindOrDashboardspecTextVariableKindOrDashboardspecConstantVariableKindOrDashboardspecDatasourceVariableKindOrDashboardspecIntervalVariableKindOrDashboardspecCustomVariableKindOrDashboardspecGroupByVariableKindOrDashboardspecAdhocVariableKind struct {
	DashboardspecQueryVariableKind      *DashboardspecQueryVariableKind      `json:"DashboardspecQueryVariableKind,omitempty"`
	DashboardspecTextVariableKind       *DashboardspecTextVariableKind       `json:"DashboardspecTextVariableKind,omitempty"`
	DashboardspecConstantVariableKind   *DashboardspecConstantVariableKind   `json:"DashboardspecConstantVariableKind,omitempty"`
	DashboardspecDatasourceVariableKind *DashboardspecDatasourceVariableKind `json:"DashboardspecDatasourceVariableKind,omitempty"`
	DashboardspecIntervalVariableKind   *DashboardspecIntervalVariableKind   `json:"DashboardspecIntervalVariableKind,omitempty"`
	DashboardspecCustomVariableKind     *DashboardspecCustomVariableKind     `json:"DashboardspecCustomVariableKind,omitempty"`
	DashboardspecGroupByVariableKind    *DashboardspecGroupByVariableKind    `json:"DashboardspecGroupByVariableKind,omitempty"`
	DashboardspecAdhocVariableKind      *DashboardspecAdhocVariableKind      `json:"DashboardspecAdhocVariableKind,omitempty"`
}

// NewDashboardspecQueryVariableKindOrDashboardspecTextVariableKindOrDashboardspecConstantVariableKindOrDashboardspecDatasourceVariableKindOrDashboardspecIntervalVariableKindOrDashboardspecCustomVariableKindOrDashboardspecGroupByVariableKindOrDashboardspecAdhocVariableKind creates a new DashboardspecQueryVariableKindOrDashboardspecTextVariableKindOrDashboardspecConstantVariableKindOrDashboardspecDatasourceVariableKindOrDashboardspecIntervalVariableKindOrDashboardspecCustomVariableKindOrDashboardspecGroupByVariableKindOrDashboardspecAdhocVariableKind object.
func NewDashboardspecQueryVariableKindOrDashboardspecTextVariableKindOrDashboardspecConstantVariableKindOrDashboardspecDatasourceVariableKindOrDashboardspecIntervalVariableKindOrDashboardspecCustomVariableKindOrDashboardspecGroupByVariableKindOrDashboardspecAdhocVariableKind() *DashboardspecQueryVariableKindOrDashboardspecTextVariableKindOrDashboardspecConstantVariableKindOrDashboardspecDatasourceVariableKindOrDashboardspecIntervalVariableKindOrDashboardspecCustomVariableKindOrDashboardspecGroupByVariableKindOrDashboardspecAdhocVariableKind {
	return &DashboardspecQueryVariableKindOrDashboardspecTextVariableKindOrDashboardspecConstantVariableKindOrDashboardspecDatasourceVariableKindOrDashboardspecIntervalVariableKindOrDashboardspecCustomVariableKindOrDashboardspecGroupByVariableKindOrDashboardspecAdhocVariableKind{}
}

// MarshalJSON implements a custom JSON marshalling logic to encode `DashboardspecQueryVariableKindOrDashboardspecTextVariableKindOrDashboardspecConstantVariableKindOrDashboardspecDatasourceVariableKindOrDashboardspecIntervalVariableKindOrDashboardspecCustomVariableKindOrDashboardspecGroupByVariableKindOrDashboardspecAdhocVariableKind` as JSON.
func (resource DashboardspecQueryVariableKindOrDashboardspecTextVariableKindOrDashboardspecConstantVariableKindOrDashboardspecDatasourceVariableKindOrDashboardspecIntervalVariableKindOrDashboardspecCustomVariableKindOrDashboardspecGroupByVariableKindOrDashboardspecAdhocVariableKind) MarshalJSON() ([]byte, error) {
	if resource.DashboardspecQueryVariableKind != nil {
		return json.Marshal(resource.DashboardspecQueryVariableKind)
	}
	if resource.DashboardspecTextVariableKind != nil {
		return json.Marshal(resource.DashboardspecTextVariableKind)
	}
	if resource.DashboardspecConstantVariableKind != nil {
		return json.Marshal(resource.DashboardspecConstantVariableKind)
	}
	if resource.DashboardspecDatasourceVariableKind != nil {
		return json.Marshal(resource.DashboardspecDatasourceVariableKind)
	}
	if resource.DashboardspecIntervalVariableKind != nil {
		return json.Marshal(resource.DashboardspecIntervalVariableKind)
	}
	if resource.DashboardspecCustomVariableKind != nil {
		return json.Marshal(resource.DashboardspecCustomVariableKind)
	}
	if resource.DashboardspecGroupByVariableKind != nil {
		return json.Marshal(resource.DashboardspecGroupByVariableKind)
	}
	if resource.DashboardspecAdhocVariableKind != nil {
		return json.Marshal(resource.DashboardspecAdhocVariableKind)
	}

	return nil, fmt.Errorf("no value for disjunction of refs")
}

// UnmarshalJSON implements a custom JSON unmarshalling logic to decode `DashboardspecQueryVariableKindOrDashboardspecTextVariableKindOrDashboardspecConstantVariableKindOrDashboardspecDatasourceVariableKindOrDashboardspecIntervalVariableKindOrDashboardspecCustomVariableKindOrDashboardspecGroupByVariableKindOrDashboardspecAdhocVariableKind` from JSON.
func (resource *DashboardspecQueryVariableKindOrDashboardspecTextVariableKindOrDashboardspecConstantVariableKindOrDashboardspecDatasourceVariableKindOrDashboardspecIntervalVariableKindOrDashboardspecCustomVariableKindOrDashboardspecGroupByVariableKindOrDashboardspecAdhocVariableKind) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]interface{})
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["kind"]
	if !found {
		return errors.New("discriminator field 'kind' not found in payload")
	}

	switch discriminator {
	case "AdhocVariable":
		var dashboardspecAdhocVariableKind DashboardspecAdhocVariableKind
		if err := json.Unmarshal(raw, &dashboardspecAdhocVariableKind); err != nil {
			return err
		}

		resource.DashboardspecAdhocVariableKind = &dashboardspecAdhocVariableKind
		return nil
	case "ConstantVariable":
		var dashboardspecConstantVariableKind DashboardspecConstantVariableKind
		if err := json.Unmarshal(raw, &dashboardspecConstantVariableKind); err != nil {
			return err
		}

		resource.DashboardspecConstantVariableKind = &dashboardspecConstantVariableKind
		return nil
	case "CustomVariable":
		var dashboardspecCustomVariableKind DashboardspecCustomVariableKind
		if err := json.Unmarshal(raw, &dashboardspecCustomVariableKind); err != nil {
			return err
		}

		resource.DashboardspecCustomVariableKind = &dashboardspecCustomVariableKind
		return nil
	case "DatasourceVariable":
		var dashboardspecDatasourceVariableKind DashboardspecDatasourceVariableKind
		if err := json.Unmarshal(raw, &dashboardspecDatasourceVariableKind); err != nil {
			return err
		}

		resource.DashboardspecDatasourceVariableKind = &dashboardspecDatasourceVariableKind
		return nil
	case "GroupByVariable":
		var dashboardspecGroupByVariableKind DashboardspecGroupByVariableKind
		if err := json.Unmarshal(raw, &dashboardspecGroupByVariableKind); err != nil {
			return err
		}

		resource.DashboardspecGroupByVariableKind = &dashboardspecGroupByVariableKind
		return nil
	case "IntervalVariable":
		var dashboardspecIntervalVariableKind DashboardspecIntervalVariableKind
		if err := json.Unmarshal(raw, &dashboardspecIntervalVariableKind); err != nil {
			return err
		}

		resource.DashboardspecIntervalVariableKind = &dashboardspecIntervalVariableKind
		return nil
	case "QueryVariable":
		var dashboardspecQueryVariableKind DashboardspecQueryVariableKind
		if err := json.Unmarshal(raw, &dashboardspecQueryVariableKind); err != nil {
			return err
		}

		resource.DashboardspecQueryVariableKind = &dashboardspecQueryVariableKind
		return nil
	case "TextVariable":
		var dashboardspecTextVariableKind DashboardspecTextVariableKind
		if err := json.Unmarshal(raw, &dashboardspecTextVariableKind); err != nil {
			return err
		}

		resource.DashboardspecTextVariableKind = &dashboardspecTextVariableKind
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `kind = %v`", discriminator)
}

type StringOrArrayOfString struct {
	String        *string  `json:"String,omitempty"`
	ArrayOfString []string `json:"ArrayOfString,omitempty"`
}

// NewStringOrArrayOfString creates a new StringOrArrayOfString object.
func NewStringOrArrayOfString() *StringOrArrayOfString {
	return &StringOrArrayOfString{}
}

// MarshalJSON implements a custom JSON marshalling logic to encode `StringOrArrayOfString` as JSON.
func (resource StringOrArrayOfString) MarshalJSON() ([]byte, error) {
	if resource.String != nil {
		return json.Marshal(resource.String)
	}

	if resource.ArrayOfString != nil {
		return json.Marshal(resource.ArrayOfString)
	}

	return nil, fmt.Errorf("no value for disjunction of scalars")
}

// UnmarshalJSON implements a custom JSON unmarshalling logic to decode `StringOrArrayOfString` from JSON.
func (resource *StringOrArrayOfString) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	var errList []error

	// String
	var String string
	if err := json.Unmarshal(raw, &String); err != nil {
		errList = append(errList, err)
		resource.String = nil
	} else {
		resource.String = &String
		return nil
	}

	// ArrayOfString
	var ArrayOfString []string
	if err := json.Unmarshal(raw, &ArrayOfString); err != nil {
		errList = append(errList, err)
		resource.ArrayOfString = nil
	} else {
		resource.ArrayOfString = ArrayOfString
		return nil
	}

	return errors.Join(errList...)
}

type StringOrDashboardspecDataQueryKind struct {
	String                     *string                     `json:"String,omitempty"`
	DashboardspecDataQueryKind *DashboardspecDataQueryKind `json:"DashboardspecDataQueryKind,omitempty"`
}

// NewStringOrDashboardspecDataQueryKind creates a new StringOrDashboardspecDataQueryKind object.
func NewStringOrDashboardspecDataQueryKind() *StringOrDashboardspecDataQueryKind {
	return &StringOrDashboardspecDataQueryKind{}
}

type StringOrFloat64 struct {
	String  *string  `json:"String,omitempty"`
	Float64 *float64 `json:"Float64,omitempty"`
}

// NewStringOrFloat64 creates a new StringOrFloat64 object.
func NewStringOrFloat64() *StringOrFloat64 {
	return &StringOrFloat64{}
}

// MarshalJSON implements a custom JSON marshalling logic to encode `StringOrFloat64` as JSON.
func (resource StringOrFloat64) MarshalJSON() ([]byte, error) {
	if resource.String != nil {
		return json.Marshal(resource.String)
	}

	if resource.Float64 != nil {
		return json.Marshal(resource.Float64)
	}

	return nil, fmt.Errorf("no value for disjunction of scalars")
}

// UnmarshalJSON implements a custom JSON unmarshalling logic to decode `StringOrFloat64` from JSON.
func (resource *StringOrFloat64) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	var errList []error

	// String
	var String string
	if err := json.Unmarshal(raw, &String); err != nil {
		errList = append(errList, err)
		resource.String = nil
	} else {
		resource.String = &String
		return nil
	}

	// Float64
	var Float64 float64
	if err := json.Unmarshal(raw, &Float64); err != nil {
		errList = append(errList, err)
		resource.Float64 = nil
	} else {
		resource.Float64 = &Float64
		return nil
	}

	return errors.Join(errList...)
}
