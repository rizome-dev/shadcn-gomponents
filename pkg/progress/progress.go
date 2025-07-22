package progress

import (
	"fmt"
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// Props defines the properties for the Progress component
type Props struct {
	Value int    // Progress value (0-100)
	Max   int    // Maximum value (default: 100)
	Size  string // "sm" | "default" | "lg"
	Class string // Additional custom classes
}

// New creates a new Progress component
func New(props Props) g.Node {
	// Set defaults
	if props.Max == 0 {
		props.Max = 100
	}
	
	// Ensure value is within bounds
	if props.Value < 0 {
		props.Value = 0
	}
	if props.Value > props.Max {
		props.Value = props.Max
	}
	
	// Calculate percentage
	percentage := (props.Value * 100) / props.Max
	
	// Determine size classes
	sizeClass := "h-4" // default
	switch props.Size {
	case "sm":
		sizeClass = "h-2"
	case "lg":
		sizeClass = "h-6"
	}
	
	// Root classes
	rootClasses := lib.CN(
		"relative w-full overflow-hidden rounded-full bg-secondary",
		sizeClass,
		props.Class,
	)
	
	// Indicator classes
	indicatorClasses := "h-full bg-primary transition-all"
	
	return html.Div(
		html.Class(rootClasses),
		html.Role("progressbar"),
		g.Attr("aria-valuemin", "0"),
		g.Attr("aria-valuemax", fmt.Sprintf("%d", props.Max)),
		g.Attr("aria-valuenow", fmt.Sprintf("%d", props.Value)),
		g.Attr("aria-label", fmt.Sprintf("Progress: %d%%", percentage)),
		g.Attr("data-value", fmt.Sprintf("%d", props.Value)),
		g.Attr("data-max", fmt.Sprintf("%d", props.Max)),
		html.Div(
			html.Class(indicatorClasses),
			html.Style(fmt.Sprintf("width: %d%%", percentage)),
		),
	)
}

// Default creates a progress bar with default settings
func Default(value int) g.Node {
	return New(Props{Value: value})
}

// Small creates a small progress bar
func SmallComponent(value int) g.Node {
	return New(Props{
		Value: value,
		Size:  "sm",
	})
}

// Large creates a large progress bar
func Large(value int) g.Node {
	return New(Props{
		Value: value,
		Size:  "lg",
	})
}

// WithLabel creates a progress bar with a label
func WithLabel(value int, label string) g.Node {
	return html.Div(
		html.Class("space-y-2"),
		html.Div(html.Class("flex items-center justify-between text-sm"),
			html.Span(g.Text(label)),
			html.Span(html.Class("text-muted-foreground"), g.Textf("%d%%", value)),
		),
		Default(value),
	)
}

// Indeterminate creates an indeterminate progress bar (loading animation)
func Indeterminate() g.Node {
	// For indeterminate progress, we'll use a different animation approach
	rootClasses := lib.CN(
		"relative h-4 w-full overflow-hidden rounded-full bg-secondary",
	)
	
	indicatorClasses := lib.CN(
		"h-full bg-primary",
		"animate-pulse", // Simple pulse animation for indeterminate state
	)
	
	return html.Div(
		html.Class(rootClasses),
		html.Role("progressbar"),
		g.Attr("aria-label", "Loading..."),
		html.Div(
			html.Class(indicatorClasses),
			html.Style("width: 50%"), // Fixed width for visual effect
		),
	)
}

// Striped creates a striped progress bar
func Striped(value int) g.Node {
	// Add striped pattern using CSS
	stripedStyles := `
		background-image: linear-gradient(
			45deg,
			rgba(255, 255, 255, 0.15) 25%,
			transparent 25%,
			transparent 50%,
			rgba(255, 255, 255, 0.15) 50%,
			rgba(255, 255, 255, 0.15) 75%,
			transparent 75%,
			transparent
		);
		background-size: 1rem 1rem;
	`
	
	return html.Div(
		html.Class("relative h-4 w-full overflow-hidden rounded-full bg-secondary"),
		html.Role("progressbar"),
		g.Attr("aria-valuemin", "0"),
		g.Attr("aria-valuemax", "100"),
		g.Attr("aria-valuenow", fmt.Sprintf("%d", value)),
		html.Div(
			html.Class("h-full bg-primary transition-all"),
			html.Style(fmt.Sprintf("width: %d%%; %s", value, stripedStyles)),
		),
	)
}

// Multi creates a multi-segment progress bar
func Multi(segments []Segment) g.Node {
	var children []g.Node
	
	for i, segment := range segments {
		segmentClasses := lib.CN(
			"h-full transition-all",
			segment.Color,
		)
		
		children = append(children, html.Div(
			html.Class(segmentClasses),
			html.Style(fmt.Sprintf("width: %d%%", segment.Value)),
			g.If(i == 0, g.Attr("role", "presentation")), // Only first segment needs role
		))
	}
	
	// Calculate total
	total := 0
	for _, s := range segments {
		total += s.Value
	}
	
	return html.Div(
		html.Class("relative h-4 w-full overflow-hidden rounded-full bg-secondary flex"),
		html.Role("progressbar"),
		g.Attr("aria-valuemin", "0"),
		g.Attr("aria-valuemax", "100"),
		g.Attr("aria-valuenow", fmt.Sprintf("%d", total)),
		g.Attr("aria-label", fmt.Sprintf("Total progress: %d%%", total)),
		g.Group(children),
	)
}

// Segment defines a segment in a multi-segment progress bar
type Segment struct {
	Value int    // Percentage value
	Color string // Tailwind color class (e.g., "bg-blue-500")
}

// Circular creates a circular progress indicator
func Circular(value int, size string) g.Node {
	// Determine size
	var svgSize, strokeWidth int
	switch size {
	case "sm":
		svgSize = 32
		strokeWidth = 3
	case "lg":
		svgSize = 64
		strokeWidth = 5
	default:
		svgSize = 48
		strokeWidth = 4
	}
	
	radius := (svgSize - strokeWidth) / 2
	circumference := 2 * 3.14159 * float64(radius)
	strokeDashoffset := circumference * (1 - float64(value)/100)
	
	return html.Div(
		html.Class("relative inline-flex items-center justify-center"),
		g.El("svg",
			g.Attr("width", fmt.Sprintf("%d", svgSize)),
			g.Attr("height", fmt.Sprintf("%d", svgSize)),
			g.Attr("viewBox", fmt.Sprintf("0 0 %d %d", svgSize, svgSize)),
			g.Attr("role", "progressbar"),
			g.Attr("aria-valuemin", "0"),
			g.Attr("aria-valuemax", "100"),
			g.Attr("aria-valuenow", fmt.Sprintf("%d", value)),
			// Background circle
			g.El("circle",
				g.Attr("cx", fmt.Sprintf("%d", svgSize/2)),
				g.Attr("cy", fmt.Sprintf("%d", svgSize/2)),
				g.Attr("r", fmt.Sprintf("%d", radius)),
				g.Attr("fill", "none"),
				g.Attr("stroke", "currentColor"),
				g.Attr("stroke-width", fmt.Sprintf("%d", strokeWidth)),
				html.Class("text-secondary"),
			),
			// Progress circle
			g.El("circle",
				g.Attr("cx", fmt.Sprintf("%d", svgSize/2)),
				g.Attr("cy", fmt.Sprintf("%d", svgSize/2)),
				g.Attr("r", fmt.Sprintf("%d", radius)),
				g.Attr("fill", "none"),
				g.Attr("stroke", "currentColor"),
				g.Attr("stroke-width", fmt.Sprintf("%d", strokeWidth)),
				g.Attr("stroke-dasharray", fmt.Sprintf("%.2f", circumference)),
				g.Attr("stroke-dashoffset", fmt.Sprintf("%.2f", strokeDashoffset)),
				g.Attr("stroke-linecap", "round"),
				g.Attr("transform", fmt.Sprintf("rotate(-90 %d %d)", svgSize/2, svgSize/2)),
				html.Class("text-primary transition-all"),
			),
		),
		// Center text
		html.Div(
			html.Class("absolute inset-0 flex items-center justify-center"),
			html.Span(html.Class("text-xs font-semibold"), g.Textf("%d%%", value)),
		),
	)
}