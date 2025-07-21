package chart

import (
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

// New creates a new chart component with server-side rendering
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
	
	return html.Div(
		html.Class(classes),
		g.Attr("data-chart", "static"),
		g.Attr("data-chart-type", string(cfg.chartType)),
		g.Attr("data-theme", cfg.theme),
		
		// Chart container with server-rendered content
		ServerRenderedContainer(id, data, cfg),
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

// ServerRenderedContainer creates a server-rendered chart container
func ServerRenderedContainer(id string, data ChartData, cfg *config) g.Node {
	return html.Div(
		html.ID(id),
		html.Class("w-full"),
		html.Style(fmt.Sprintf("height: %s; width: %s;", cfg.height, cfg.width)),
		g.Attr("role", "img"),
		g.Attr("aria-label", fmt.Sprintf("%s chart", cfg.chartType)),
		
		// Title section
		g.If(cfg.title != "",
			Title(cfg.title, cfg.subtitle),
		),
		
		// Chart visualization area
		ServerRenderedVisualization(id, data, cfg),
		
		// Legend
		g.If(cfg.showLegend,
			Legend(data, cfg),
		),
	)
}

// ServerRenderedVisualization creates the visualization area for static charts
func ServerRenderedVisualization(id string, data ChartData, cfg *config) g.Node {
	return html.Div(
		html.Class("relative overflow-hidden"),
		html.Style(fmt.Sprintf("height: calc(%s - 4rem)", cfg.height)),
		
		// Server-rendered chart based on type
		g.If(cfg.chartType == ChartTypeBar,
			RenderBarChart(data, cfg),
		),
		g.If(cfg.chartType == ChartTypeLine,
			RenderLineChart(data, cfg),
		),
		g.If(cfg.chartType == ChartTypePie || cfg.chartType == ChartTypeDonut,
			RenderPieChart(data, cfg),
		),
		g.If(cfg.chartType == ChartTypeArea,
			RenderAreaChart(data, cfg),
		),
	)
}

// RenderBarChart creates a static bar chart
func RenderBarChart(data ChartData, cfg *config) g.Node {
	if len(data.Series) == 0 || len(data.Labels) == 0 {
		return html.Div(html.Class("text-muted-foreground"), g.Text("No data available"))
	}
	
	// Find max value for scaling
	maxValue := 0.0
	for _, series := range data.Series {
		for _, value := range series.Data {
			if value > maxValue {
				maxValue = value
			}
		}
	}
	
	return html.Div(
		html.Class("relative h-full"),
		
		// Y-axis labels
		g.If(cfg.showGrid,
			html.Div(
				html.Class("absolute left-0 top-0 bottom-0 w-12 flex flex-col justify-between text-xs text-muted-foreground"),
				renderYAxisLabels(maxValue),
			),
		),
		
		// Chart area
		html.Div(
			html.Class("ml-14 h-full relative"),
			
			// Grid lines
			g.If(cfg.showGrid,
				renderGridLines(),
			),
			
			// Bars
			html.Div(
				html.Class("absolute inset-0 flex items-end justify-around px-2"),
				g.Group(renderStaticBars(data, maxValue, cfg)),
			),
			
			// X-axis labels
			html.Div(
				html.Class("absolute bottom-0 left-0 right-0 flex justify-around text-xs text-muted-foreground -mb-6"),
				g.Group(renderStaticXAxisLabels(data.Labels)),
			),
		),
	)
}

// RenderLineChart creates a static line chart
func RenderLineChart(data ChartData, cfg *config) g.Node {
	if len(data.Series) == 0 || len(data.Labels) == 0 {
		return html.Div(html.Class("text-muted-foreground"), g.Text("No data available"))
	}
	
	// Simple SVG line chart
	width := 600
	height := 300
	padding := 40
	
	return g.El("svg",
		g.Attr("viewBox", fmt.Sprintf("0 0 %d %d", width, height)),
		html.Class("w-full h-full"),
		g.Attr("preserveAspectRatio", "xMidYMid meet"),
		
		// Grid
		g.If(cfg.showGrid,
			renderStaticSVGGrid(width, height, padding),
		),
		
		// Lines
		g.Group(renderStaticSVGLines(data, width, height, padding)),
		
		// Points
		g.Group(renderStaticSVGPoints(data, width, height, padding)),
	)
}

// RenderPieChart creates a static pie chart
func RenderPieChart(data ChartData, cfg *config) g.Node {
	if len(data.Series) == 0 || len(data.Series[0].Data) == 0 {
		return html.Div(html.Class("text-muted-foreground"), g.Text("No data available"))
	}
	
	// Calculate total
	total := 0.0
	for _, value := range data.Series[0].Data {
		total += value
	}
	
	// SVG pie chart
	size := 300
	center := size / 2
	radius := size/2 - 20
	innerRadius := 0
	if cfg.chartType == ChartTypeDonut {
		innerRadius = radius / 2
	}
	
	return g.El("svg",
		g.Attr("viewBox", fmt.Sprintf("0 0 %d %d", size, size)),
		html.Class("w-full h-full max-w-xs mx-auto"),
		
		// Pie slices
		g.Group(renderStaticPieSlices(data, total, center, radius, innerRadius)),
	)
}

// RenderAreaChart creates a static area chart
func RenderAreaChart(data ChartData, cfg *config) g.Node {
	if len(data.Series) == 0 || len(data.Labels) == 0 {
		return html.Div(html.Class("text-muted-foreground"), g.Text("No data available"))
	}
	
	// SVG area chart
	width := 600
	height := 300
	padding := 40
	
	return g.El("svg",
		g.Attr("viewBox", fmt.Sprintf("0 0 %d %d", width, height)),
		html.Class("w-full h-full"),
		g.Attr("preserveAspectRatio", "xMidYMid meet"),
		
		// Grid
		g.If(cfg.showGrid,
			renderStaticSVGGrid(width, height, padding),
		),
		
		// Area fills
		g.Group(renderStaticSVGAreas(data, width, height, padding)),
		
		// Lines on top
		g.Group(renderStaticSVGLines(data, width, height, padding)),
	)
}

// Helper functions for static rendering

func renderYAxisLabels(maxValue float64) g.Node {
	steps := 5
	labels := make([]g.Node, steps)
	
	for i := 0; i < steps; i++ {
		value := maxValue * float64(steps-i-1) / float64(steps-1)
		labels[i] = html.Div(
			html.Class("text-right pr-2"),
			g.Text(fmt.Sprintf("%.0f", value)),
		)
	}
	
	return g.Group(labels)
}

func renderGridLines() g.Node {
	return html.Div(
		html.Class("absolute inset-0"),
		// Horizontal lines
		html.Div(html.Class("absolute top-0 left-0 right-0 border-t border-gray-200")),
		html.Div(html.Class("absolute top-1/4 left-0 right-0 border-t border-gray-200")),
		html.Div(html.Class("absolute top-1/2 left-0 right-0 border-t border-gray-200")),
		html.Div(html.Class("absolute top-3/4 left-0 right-0 border-t border-gray-200")),
		html.Div(html.Class("absolute bottom-0 left-0 right-0 border-t border-gray-200")),
	)
}

func renderStaticBars(data ChartData, maxValue float64, cfg *config) []g.Node {
	bars := []g.Node{}
	barGroupWidth := 100.0 / float64(len(data.Labels))
	
	for i, label := range data.Labels {
		barGroup := []g.Node{}
		seriesCount := len(data.Series)
		barWidth := barGroupWidth / float64(seriesCount) * 0.8
		
		for j, series := range data.Series {
			if i < len(series.Data) {
				height := (series.Data[i] / maxValue) * 100
				left := (float64(i) * barGroupWidth) + (float64(j) * barWidth)
				
				color := series.Color
				if color == "" {
					colors := []string{"#3b82f6", "#10b981", "#f59e0b", "#ef4444", "#8b5cf6"}
					color = colors[j%len(colors)]
				}
				
				barGroup = append(barGroup, html.Div(
					html.Class("absolute bottom-0 transition-all duration-300 hover:opacity-80"),
					html.Style(fmt.Sprintf("height: %.1f%%; width: %.1f%%; left: %.1f%%; background-color: %s;",
						height, barWidth*0.9, left, color)),
					g.Attr("data-label", label),
					g.Attr("data-series", series.Name),
					g.Attr("data-value", fmt.Sprintf("%.2f", series.Data[i])),
				))
			}
		}
		bars = append(bars, g.Group(barGroup))
	}
	
	return bars
}

func renderStaticXAxisLabels(labels []string) []g.Node {
	nodes := make([]g.Node, len(labels))
	for i, label := range labels {
		nodes[i] = html.Div(
			html.Class("text-center"),
			g.Text(label),
		)
	}
	return nodes
}

func renderStaticSVGGrid(width, height, padding int) g.Node {
	lines := []g.Node{}
	
	// Horizontal lines
	for i := 0; i <= 5; i++ {
		y := padding + (height-2*padding)*i/5
		lines = append(lines, g.El("line",
			g.Attr("x1", fmt.Sprintf("%d", padding)),
			g.Attr("y1", fmt.Sprintf("%d", y)),
			g.Attr("x2", fmt.Sprintf("%d", width-padding)),
			g.Attr("y2", fmt.Sprintf("%d", y)),
			g.Attr("stroke", "#e5e7eb"),
			g.Attr("stroke-width", "1"),
		))
	}
	
	return g.El("g", g.Group(lines))
}

func renderStaticSVGLines(data ChartData, width, height, padding int) []g.Node {
	if len(data.Series) == 0 {
		return nil
	}
	
	// Find max value
	maxValue := 0.0
	for _, series := range data.Series {
		for _, value := range series.Data {
			if value > maxValue {
				maxValue = value
			}
		}
	}
	
	lines := []g.Node{}
	colors := []string{"#3b82f6", "#10b981", "#f59e0b", "#ef4444", "#8b5cf6"}
	
	for i, series := range data.Series {
		if len(series.Data) < 2 {
			continue
		}
		
		points := []string{}
		xStep := float64(width-2*padding) / float64(len(series.Data)-1)
		
		for j, value := range series.Data {
			x := float64(padding) + float64(j)*xStep
			y := float64(height-padding) - (value/maxValue)*float64(height-2*padding)
			points = append(points, fmt.Sprintf("%.1f,%.1f", x, y))
		}
		
		color := series.Color
		if color == "" {
			color = colors[i%len(colors)]
		}
		
		lines = append(lines, g.El("polyline",
			g.Attr("points", strings.Join(points, " ")),
			g.Attr("fill", "none"),
			g.Attr("stroke", color),
			g.Attr("stroke-width", "2"),
		))
	}
	
	return lines
}

func renderStaticSVGPoints(data ChartData, width, height, padding int) []g.Node {
	// Find max value
	maxValue := 0.0
	for _, series := range data.Series {
		for _, value := range series.Data {
			if value > maxValue {
				maxValue = value
			}
		}
	}
	
	points := []g.Node{}
	colors := []string{"#3b82f6", "#10b981", "#f59e0b", "#ef4444", "#8b5cf6"}
	
	for i, series := range data.Series {
		xStep := float64(width-2*padding) / float64(len(series.Data)-1)
		
		color := series.Color
		if color == "" {
			color = colors[i%len(colors)]
		}
		
		for j, value := range series.Data {
			x := float64(padding) + float64(j)*xStep
			y := float64(height-padding) - (value/maxValue)*float64(height-2*padding)
			
			points = append(points, g.El("circle",
				g.Attr("cx", fmt.Sprintf("%.1f", x)),
				g.Attr("cy", fmt.Sprintf("%.1f", y)),
				g.Attr("r", "4"),
				g.Attr("fill", color),
				g.Attr("data-series", series.Name),
				g.Attr("data-value", fmt.Sprintf("%.2f", value)),
			))
		}
	}
	
	return points
}

func renderStaticSVGAreas(data ChartData, width, height, padding int) []g.Node {
	if len(data.Series) == 0 {
		return nil
	}
	
	// Find max value
	maxValue := 0.0
	for _, series := range data.Series {
		for _, value := range series.Data {
			if value > maxValue {
				maxValue = value
			}
		}
	}
	
	areas := []g.Node{}
	colors := []string{"#3b82f6", "#10b981", "#f59e0b", "#ef4444", "#8b5cf6"}
	
	for i, series := range data.Series {
		if len(series.Data) < 2 {
			continue
		}
		
		points := []string{}
		xStep := float64(width-2*padding) / float64(len(series.Data)-1)
		
		// Start from bottom left
		points = append(points, fmt.Sprintf("%d,%d", padding, height-padding))
		
		// Add data points
		for j, value := range series.Data {
			x := float64(padding) + float64(j)*xStep
			y := float64(height-padding) - (value/maxValue)*float64(height-2*padding)
			points = append(points, fmt.Sprintf("%.1f,%.1f", x, y))
		}
		
		// Close at bottom right
		points = append(points, fmt.Sprintf("%d,%d", width-padding, height-padding))
		
		color := series.Color
		if color == "" {
			color = colors[i%len(colors)]
		}
		
		areas = append(areas, g.El("polygon",
			g.Attr("points", strings.Join(points, " ")),
			g.Attr("fill", color),
			g.Attr("fill-opacity", "0.3"),
		))
	}
	
	return areas
}

func renderStaticPieSlices(data ChartData, total float64, center, radius, innerRadius int) []g.Node {
	slices := []g.Node{}
	colors := []string{"#3b82f6", "#10b981", "#f59e0b", "#ef4444", "#8b5cf6", "#ec4899", "#14b8a6"}
	
	startAngle := -90.0 // Start from top
	
	for i, value := range data.Series[0].Data {
		if i >= len(data.Labels) {
			break
		}
		
		percentage := value / total
		angle := percentage * 360
		
		color := colors[i%len(colors)]
		if len(data.Series[0].Color) > 0 {
			color = data.Series[0].Color
		}
		
		slice := createStaticPieSlice(center, radius, innerRadius, startAngle, angle, color, data.Labels[i], value)
		slices = append(slices, slice)
		
		startAngle += angle
	}
	
	return slices
}

func createStaticPieSlice(center, radius, innerRadius int, startAngle, angle float64, color, label string, value float64) g.Node {
	// Convert to radians
	startRad := startAngle * 3.14159 / 180
	endRad := (startAngle + angle) * 3.14159 / 180
	
	// Calculate points
	x1 := float64(center) + float64(radius)*cosine(startRad)
	y1 := float64(center) + float64(radius)*sine(startRad)
	x2 := float64(center) + float64(radius)*cosine(endRad)
	y2 := float64(center) + float64(radius)*sine(endRad)
	
	largeArc := 0
	if angle > 180 {
		largeArc = 1
	}
	
	var path string
	if innerRadius > 0 {
		// Donut chart
		ix1 := float64(center) + float64(innerRadius)*cosine(startRad)
		iy1 := float64(center) + float64(innerRadius)*sine(startRad)
		ix2 := float64(center) + float64(innerRadius)*cosine(endRad)
		iy2 := float64(center) + float64(innerRadius)*sine(endRad)
		
		path = fmt.Sprintf("M %.1f %.1f A %d %d 0 %d 1 %.1f %.1f L %.1f %.1f A %d %d 0 %d 0 %.1f %.1f Z",
			x1, y1, radius, radius, largeArc, x2, y2,
			ix2, iy2, innerRadius, innerRadius, largeArc, ix1, iy1)
	} else {
		// Pie chart
		path = fmt.Sprintf("M %d %d L %.1f %.1f A %d %d 0 %d 1 %.1f %.1f Z",
			center, center, x1, y1, radius, radius, largeArc, x2, y2)
	}
	
	return g.El("path",
		g.Attr("d", path),
		g.Attr("fill", color),
		g.Attr("stroke", "white"),
		g.Attr("stroke-width", "2"),
		g.Attr("data-label", label),
		g.Attr("data-value", fmt.Sprintf("%.2f", value)),
		html.Class("hover:opacity-80 transition-opacity cursor-pointer"),
	)
}

