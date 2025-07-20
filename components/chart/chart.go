package chart

import (
	"encoding/json"
	"fmt"
	"strings"

	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

// ChartType represents the type of chart
type ChartType string

const (
	ChartTypeLine      ChartType = "line"
	ChartTypeBar       ChartType = "bar"
	ChartTypeArea      ChartType = "area"
	ChartTypePie       ChartType = "pie"
	ChartTypeDonut     ChartType = "donut"
	ChartTypeRadar     ChartType = "radar"
	ChartTypeRadialBar ChartType = "radialBar"
)

// ChartData represents the data structure for charts
type ChartData struct {
	Labels []string      `json:"labels"`
	Series []SeriesData  `json:"series"`
}

// SeriesData represents a data series
type SeriesData struct {
	Name   string    `json:"name"`
	Data   []float64 `json:"data"`
	Color  string    `json:"color,omitempty"`
}

// Option is a functional option for configuring a chart
type Option func(*config)

type config struct {
	class       string
	height      string
	width       string
	chartType   ChartType
	title       string
	subtitle    string
	showLegend  bool
	showGrid    bool
	showTooltip bool
	responsive  bool
	animations  bool
	theme       string // light or dark
}

// New creates a new chart component
func New(id string, data ChartData, opts ...Option) g.Node {
	cfg := &config{
		class:       "",
		height:      "400px",
		width:       "100%",
		chartType:   ChartTypeLine,
		showLegend:  true,
		showGrid:    true,
		showTooltip: true,
		responsive:  true,
		animations:  true,
		theme:       "light",
	}

	for _, opt := range opts {
		opt(cfg)
	}

	classes := strings.TrimSpace("relative " + cfg.class)
	
	// Convert data to JSON for embedding
	dataJSON, _ := json.Marshal(data)
	
	return html.Div(
		html.Class(classes),
		g.Attr("data-chart", "true"),
		g.Attr("data-chart-type", string(cfg.chartType)),
		g.Attr("data-chart-data", string(dataJSON)),
		g.Attr("data-theme", cfg.theme),
		g.If(cfg.title != "", g.Attr("data-title", cfg.title)),
		g.If(cfg.subtitle != "", g.Attr("data-subtitle", cfg.subtitle)),
		g.If(!cfg.showLegend, g.Attr("data-hide-legend", "true")),
		g.If(!cfg.showGrid, g.Attr("data-hide-grid", "true")),
		g.If(!cfg.showTooltip, g.Attr("data-hide-tooltip", "true")),
		g.If(!cfg.animations, g.Attr("data-no-animations", "true")),
		
		// Chart container
		Container(id, cfg),
		
		// Optional title
		g.If(cfg.title != "",
			Title(cfg.title, cfg.subtitle),
		),
	)
}

// Container creates the chart container
func Container(id string, cfg *config) g.Node {
	return html.Div(
		html.ID(id),
		html.Class("w-full"),
		html.Style(fmt.Sprintf("height: %s; width: %s;", cfg.height, cfg.width)),
		g.Attr("role", "img"),
		g.Attr("aria-label", fmt.Sprintf("%s chart", cfg.chartType)),
		
		// Canvas element for chart rendering
		Canvas(id+"-canvas", cfg),
		
		// Loading state
		LoadingState(),
	)
}

// Canvas creates the canvas element for chart rendering
func Canvas(id string, cfg *config) g.Node {
	return g.El("canvas",
		html.ID(id),
		html.Class("max-w-full"),
		g.Attr("role", "img"),
		g.If(cfg.responsive, g.Attr("data-responsive", "true")),
	)
}

// Title creates the chart title section
func Title(title, subtitle string) g.Node {
	return html.Div(
		html.Class("mb-4"),
		html.H3(
			html.Class("text-lg font-semibold"),
			g.Text(title),
		),
		g.If(subtitle != "",
			html.P(
				html.Class("text-sm text-muted-foreground"),
				g.Text(subtitle),
			),
		),
	)
}

// LoadingState creates a loading placeholder
func LoadingState() g.Node {
	return html.Div(
		html.Class("absolute inset-0 flex items-center justify-center bg-background/50"),
		g.Attr("data-chart-loading", "true"),
		html.Div(
			html.Class("flex flex-col items-center gap-2"),
			// Spinner
			html.Div(
				html.Class("h-8 w-8 animate-spin rounded-full border-4 border-primary border-t-transparent"),
			),
			html.Span(
				html.Class("text-sm text-muted-foreground"),
				g.Text("Loading chart..."),
			),
		),
	)
}

// Legend creates a custom legend component
func Legend(data ChartData, cfg *config) g.Node {
	if !cfg.showLegend || len(data.Series) == 0 {
		return nil
	}
	
	items := make([]g.Node, len(data.Series))
	for i, series := range data.Series {
		items[i] = LegendItem(series.Name, series.Color)
	}
	
	return html.Div(
		html.Class("flex flex-wrap gap-4 mt-4"),
		g.Attr("role", "list"),
		g.Attr("aria-label", "Chart legend"),
		g.Group(items),
	)
}

// LegendItem creates a single legend item
func LegendItem(name, color string) g.Node {
	if color == "" {
		color = "currentColor"
	}
	
	return html.Div(
		html.Class("flex items-center gap-2"),
		g.Attr("role", "listitem"),
		
		// Color indicator
		html.Span(
			html.Class("w-3 h-3 rounded-full"),
			html.Style(fmt.Sprintf("background-color: %s", color)),
			g.Attr("aria-hidden", "true"),
		),
		
		// Label
		html.Span(
			html.Class("text-sm"),
			g.Text(name),
		),
	)
}

// Tooltip creates a tooltip template
func Tooltip(cfg *config) g.Node {
	if !cfg.showTooltip {
		return nil
	}
	
	return html.Template(
		html.ID("chart-tooltip-template"),
		html.Div(
			html.Class("absolute z-50 rounded-lg border bg-popover px-3 py-1.5 text-sm text-popover-foreground shadow-md"),
			g.Attr("data-chart-tooltip", "true"),
			html.Div(
				html.Class("font-semibold"),
				g.Attr("data-tooltip-title", "true"),
			),
			html.Div(
				html.Class("mt-1 space-y-1"),
				g.Attr("data-tooltip-content", "true"),
			),
		),
	)
}

// Option functions

// WithClass adds custom CSS classes
func WithClass(class string) Option{
	return func(c *config) {
		c.class = class
	}
}

// WithHeight sets the chart height
func WithHeight(height string) Option{
	return func(c *config) {
		c.height = height
	}
}

// WithWidth sets the chart width
func WithWidth(width string) Option{
	return func(c *config) {
		c.width = width
	}
}

// WithType sets the chart type
func WithType(chartType ChartType) Option{
	return func(c *config) {
		c.chartType = chartType
	}
}

// WithTitle sets the chart title
func WithTitle(title string) Option{
	return func(c *config) {
		c.title = title
	}
}

// WithSubtitle sets the chart subtitle
func WithSubtitle(subtitle string) Option{
	return func(c *config) {
		c.subtitle = subtitle
	}
}

// WithoutLegend hides the legend
func WithoutLegend() Option{
	return func(c *config) {
		c.showLegend = false
	}
}

// WithoutGrid hides the grid
func WithoutGrid() Option{
	return func(c *config) {
		c.showGrid = false
	}
}

// WithoutTooltip hides tooltips
func WithoutTooltip() Option{
	return func(c *config) {
		c.showTooltip = false
	}
}

// WithoutAnimations disables animations
func WithoutAnimations() Option{
	return func(c *config) {
		c.animations = false
	}
}

// WithTheme sets the chart theme
func WithTheme(theme string) Option{
	return func(c *config) {
		c.theme = theme
	}
}

// WithoutResponsive disables responsive behavior
func WithoutResponsive() Option{
	return func(c *config) {
		c.responsive = false
	}
}