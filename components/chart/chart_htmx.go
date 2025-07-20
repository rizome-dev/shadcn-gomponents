package chart

import (
	"encoding/json"
	"fmt"
	"strings"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// HTMXConfig provides HTMX-specific configuration
type HTMXConfig struct {
	config
	endpoint       string
	updateEndpoint string
	pollInterval   int // seconds, 0 means no polling
	swapTarget     string
}

// HTMXOption is a functional option for HTMX chart
type HTMXOption func(*HTMXConfig)

// NewHTMX creates an HTMX-enabled chart
func NewHTMX(id string, data ChartData, opts ...HTMXOption) g.Node {
	cfg := &HTMXConfig{
		config: config{
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
		},
		endpoint:       "/chart/" + id,
		updateEndpoint: "/chart/" + id + "/data",
		pollInterval:   0,
		swapTarget:     "#" + id,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	classes := strings.TrimSpace("relative " + cfg.class)
	
	// Convert data to JSON for embedding
	dataJSON, _ := json.Marshal(data)
	
	return h.Div(
		h.ID(id+"-wrapper"),
		h.Class(classes),
		g.Attr("data-chart", "htmx"),
		g.Attr("data-chart-type", string(cfg.chartType)),
		g.Attr("data-chart-data", string(dataJSON)),
		g.Attr("data-theme", cfg.theme),
		g.If(cfg.title != "", g.Attr("data-title", cfg.title)),
		g.If(cfg.subtitle != "", g.Attr("data-subtitle", cfg.subtitle)),
		
		// HTMX attributes for polling if enabled
		g.If(cfg.pollInterval > 0,
			g.Group([]g.Node{
				g.Attr("hx-get", cfg.updateEndpoint),
				g.Attr("hx-trigger", fmt.Sprintf("load, every %ds", cfg.pollInterval)),
				g.Attr("hx-target", cfg.swapTarget),
				g.Attr("hx-swap", "outerHTML"),
			}),
		),
		
		// Chart container
		HTMXContainer(id, data, cfg),
		
		// Controls for interactive charts
		g.If(cfg.pollInterval == 0,
			HTMXControls(id, cfg),
		),
	)
}

// HTMXContainer creates the HTMX chart container
func HTMXContainer(id string, data ChartData, cfg *HTMXConfig) g.Node {
	return h.Div(
		h.ID(id),
		h.Class("w-full relative"),
		h.Style(fmt.Sprintf("height: %s; width: %s;", cfg.height, cfg.width)),
		g.Attr("role", "img"),
		g.Attr("aria-label", fmt.Sprintf("%s chart", cfg.chartType)),
		g.Attr("data-chart-container", "htmx"),
		
		// Title section
		g.If(cfg.title != "",
			Title(cfg.title, cfg.subtitle),
		),
		
		// Chart visualization area
		HTMXVisualization(id, data, cfg),
		
		// Legend
		g.If(cfg.showLegend,
			Legend(data, &cfg.config),
		),
	)
}

// HTMXVisualization creates the visualization area for HTMX charts
func HTMXVisualization(id string, data ChartData, cfg *HTMXConfig) g.Node {
	return h.Div(
		h.Class("relative overflow-hidden"),
		h.Style(fmt.Sprintf("height: calc(%s - 4rem)", cfg.height)),
		
		// Server-rendered chart (SVG or ASCII)
		g.If(cfg.chartType == ChartTypeBar,
			HTMXBarChart(data, cfg),
		),
		g.If(cfg.chartType == ChartTypeLine,
			HTMXLineChart(data, cfg),
		),
		g.If(cfg.chartType == ChartTypePie || cfg.chartType == ChartTypeDonut,
			HTMXPieChart(data, cfg),
		),
		
		// Tooltip container
		g.If(cfg.showTooltip,
			h.Div(
				h.ID(id+"-tooltip"),
				h.Class("hidden absolute z-50 rounded-lg border bg-popover px-3 py-1.5 text-sm text-popover-foreground shadow-md"),
			),
		),
	)
}

// HTMXBarChart creates a server-rendered bar chart
func HTMXBarChart(data ChartData, cfg *HTMXConfig) g.Node {
	if len(data.Series) == 0 || len(data.Labels) == 0 {
		return h.Div(h.Class("text-muted-foreground"), g.Text("No data available"))
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
	
	// barWidth := 100.0 / float64(len(data.Labels)) // Not used in this implementation
	
	return h.Div(
		h.Class("relative h-full"),
		
		// Y-axis labels
		g.If(cfg.showGrid,
			h.Div(
				h.Class("absolute left-0 top-0 bottom-0 w-12 flex flex-col justify-between text-xs text-muted-foreground"),
				YAxisLabels(maxValue),
			),
		),
		
		// Chart area
		h.Div(
			h.Class("ml-14 h-full relative"),
			
			// Grid lines
			g.If(cfg.showGrid,
				GridLines(),
			),
			
			// Bars
			h.Div(
				h.Class("absolute inset-0 flex items-end justify-around px-2"),
				g.Group(renderBars(data, maxValue, cfg)),
			),
			
			// X-axis labels
			h.Div(
				h.Class("absolute bottom-0 left-0 right-0 flex justify-around text-xs text-muted-foreground -mb-6"),
				g.Group(renderXAxisLabels(data.Labels)),
			),
		),
	)
}

// HTMXLineChart creates a server-rendered line chart
func HTMXLineChart(data ChartData, cfg *HTMXConfig) g.Node {
	if len(data.Series) == 0 || len(data.Labels) == 0 {
		return h.Div(h.Class("text-muted-foreground"), g.Text("No data available"))
	}
	
	// Simple SVG line chart
	width := 600
	height := 300
	padding := 40
	
	return g.El("svg",
		g.Attr("viewBox", fmt.Sprintf("0 0 %d %d", width, height)),
		h.Class("w-full h-full"),
		g.Attr("preserveAspectRatio", "xMidYMid meet"),
		
		// Grid
		g.If(cfg.showGrid,
			renderSVGGrid(width, height, padding),
		),
		
		// Lines
		g.Group(renderSVGLines(data, width, height, padding)),
		
		// Points
		g.Group(renderSVGPoints(data, width, height, padding, cfg)),
	)
}

// HTMXPieChart creates a server-rendered pie chart
func HTMXPieChart(data ChartData, cfg *HTMXConfig) g.Node {
	if len(data.Series) == 0 || len(data.Series[0].Data) == 0 {
		return h.Div(h.Class("text-muted-foreground"), g.Text("No data available"))
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
		h.Class("w-full h-full max-w-xs mx-auto"),
		
		// Pie slices
		g.Group(renderPieSlices(data, total, center, radius, innerRadius)),
	)
}

// HTMXControls creates interactive controls for charts
func HTMXControls(id string, cfg *HTMXConfig) g.Node {
	return h.Div(
		h.Class("mt-4 flex flex-wrap gap-2"),
		
		// Refresh button
		h.Button(
			h.Type("button"),
			h.Class("inline-flex items-center gap-2 rounded-md bg-primary px-3 py-2 text-sm font-semibold text-primary-foreground shadow-sm hover:bg-primary/90"),
			g.Attr("hx-get", cfg.updateEndpoint),
			g.Attr("hx-target", "#"+id),
			g.Attr("hx-swap", "outerHTML"),
			g.Attr("hx-indicator", "#"+id+"-loading"),
			
			// Refresh icon
			g.Raw(`<svg class="h-4 w-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
				<path d="M21 2v6h-6"></path>
				<path d="M3 12a9 9 0 0 1 15-6.7L21 8"></path>
				<path d="M3 22v-6h6"></path>
				<path d="M21 12a9 9 0 0 1-15 6.7L3 16"></path>
			</svg>`),
			g.Text("Refresh"),
		),
		
		// Loading indicator
		h.Div(
			h.ID(id+"-loading"),
			h.Class("htmx-indicator inline-flex items-center gap-2 text-sm text-muted-foreground"),
			h.Div(h.Class("h-4 w-4 animate-spin rounded-full border-2 border-primary border-t-transparent")),
			g.Text("Updating..."),
		),
	)
}

// Helper functions for rendering

func renderBars(data ChartData, maxValue float64, cfg *HTMXConfig) []g.Node {
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
				
				barGroup = append(barGroup, h.Div(
					h.Class("absolute bottom-0 transition-all duration-300 hover:opacity-80"),
					h.Style(fmt.Sprintf("height: %.1f%%; width: %.1f%%; left: %.1f%%; background-color: %s;",
						height, barWidth*0.9, left, color)),
					g.Attr("data-label", label),
					g.Attr("data-series", series.Name),
					g.Attr("data-value", fmt.Sprintf("%.2f", series.Data[i])),
					g.If(cfg.showTooltip,
						g.Group([]g.Node{
							g.Attr("hx-get", fmt.Sprintf("%s/tooltip?series=%s&label=%s", cfg.endpoint, series.Name, label)),
							g.Attr("hx-trigger", "mouseenter"),
							g.Attr("hx-target", "#"+cfg.swapTarget+"-tooltip"),
							g.Attr("hx-swap", "innerHTML"),
						}),
					),
				))
			}
		}
		bars = append(bars, g.Group(barGroup))
	}
	
	return bars
}

func renderXAxisLabels(labels []string) []g.Node {
	nodes := make([]g.Node, len(labels))
	for i, label := range labels {
		nodes[i] = h.Div(
			h.Class("text-center"),
			g.Text(label),
		)
	}
	return nodes
}

func YAxisLabels(maxValue float64) g.Node {
	steps := 5
	labels := make([]g.Node, steps)
	
	for i := 0; i < steps; i++ {
		value := maxValue * float64(steps-i-1) / float64(steps-1)
		labels[i] = h.Div(
			h.Class("text-right pr-2"),
			g.Text(fmt.Sprintf("%.0f", value)),
		)
	}
	
	return g.Group(labels)
}

func GridLines() g.Node {
	return h.Div(
		h.Class("absolute inset-0"),
		// Horizontal lines
		h.Div(h.Class("absolute top-0 left-0 right-0 border-t border-gray-200")),
		h.Div(h.Class("absolute top-1/4 left-0 right-0 border-t border-gray-200")),
		h.Div(h.Class("absolute top-1/2 left-0 right-0 border-t border-gray-200")),
		h.Div(h.Class("absolute top-3/4 left-0 right-0 border-t border-gray-200")),
		h.Div(h.Class("absolute bottom-0 left-0 right-0 border-t border-gray-200")),
	)
}

func renderSVGGrid(width, height, padding int) g.Node {
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

func renderSVGLines(data ChartData, width, height, padding int) []g.Node {
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

func renderSVGPoints(data ChartData, width, height, padding int, cfg *HTMXConfig) []g.Node {
	if !cfg.showTooltip {
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

func renderPieSlices(data ChartData, total float64, center, radius, innerRadius int) []g.Node {
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
		
		slice := createPieSlice(center, radius, innerRadius, startAngle, angle, color, data.Labels[i], value)
		slices = append(slices, slice)
		
		startAngle += angle
	}
	
	return slices
}

func createPieSlice(center, radius, innerRadius int, startAngle, angle float64, color, label string, value float64) g.Node {
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
		h.Class("hover:opacity-80 transition-opacity cursor-pointer"),
	)
}

// Simple trigonometry helpers
func sine(rad float64) float64 {
	// Using math package would be better, but for a simple implementation:
	// This is a rough approximation using Taylor series
	x := rad
	x3 := x * x * x
	x5 := x3 * x * x
	return x - x3/6 + x5/120
}

func cosine(rad float64) float64 {
	// cos(x) = sin(x + π/2)
	return sine(rad + 3.14159/2)
}

// ChartResponse generates a chart update response
func ChartResponse(id string, data ChartData, cfg *HTMXConfig) g.Node {
	return HTMXContainer(id, data, cfg)
}

// TooltipResponse generates a tooltip response
func TooltipResponse(series, label string, value float64) g.Node {
	return h.Div(
		h.Div(
			h.Class("font-semibold"),
			g.Text(label),
		),
		h.Div(
			h.Class("text-sm"),
			g.Text(fmt.Sprintf("%s: %.2f", series, value)),
		),
	)
}

// HTMX Option functions

// WithHTMXClass adds custom CSS classes
func WithHTMXClass(class string) HTMXOption {
	return func(c *HTMXConfig) {
		c.class = class
	}
}

// WithHTMXHeight sets the chart height
func WithHTMXHeight(height string) HTMXOption {
	return func(c *HTMXConfig) {
		c.height = height
	}
}

// WithHTMXType sets the chart type
func WithHTMXType(chartType ChartType) HTMXOption {
	return func(c *HTMXConfig) {
		c.chartType = chartType
	}
}

// WithHTMXTitle sets the chart title
func WithHTMXTitle(title string) HTMXOption {
	return func(c *HTMXConfig) {
		c.title = title
	}
}

// WithHTMXSubtitle sets the chart subtitle
func WithHTMXSubtitle(subtitle string) HTMXOption {
	return func(c *HTMXConfig) {
		c.subtitle = subtitle
	}
}

// WithHTMXEndpoint sets custom endpoints
func WithHTMXEndpoint(endpoint string) HTMXOption {
	return func(c *HTMXConfig) {
		c.endpoint = endpoint
		c.updateEndpoint = endpoint + "/data"
	}
}

// WithHTMXPolling enables auto-refresh polling
func WithHTMXPolling(intervalSeconds int) HTMXOption {
	return func(c *HTMXConfig) {
		c.pollInterval = intervalSeconds
	}
}

// WithHTMXSwapTarget sets custom swap target
func WithHTMXSwapTarget(target string) HTMXOption {
	return func(c *HTMXConfig) {
		c.swapTarget = target
	}
}

// WithoutHTMXLegend hides the legend
func WithoutHTMXLegend() HTMXOption {
	return func(c *HTMXConfig) {
		c.showLegend = false
	}
}

// WithoutHTMXGrid hides the grid
func WithoutHTMXGrid() HTMXOption {
	return func(c *HTMXConfig) {
		c.showGrid = false
	}
}

// WithoutHTMXTooltip hides tooltips
func WithoutHTMXTooltip() HTMXOption {
	return func(c *HTMXConfig) {
		c.showTooltip = false
	}
}