package slider

import (
	"fmt"

	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// Props defines the properties for the slider component
type Props struct {
	Min         int
	Max         int
	Step        int
	Value       []int // Can have multiple values for range sliders
	Disabled    bool
	Orientation string // "horizontal" | "vertical"
	Class       string
	ID          string
	Name        string
}

// New creates a new slider component
func New(props Props, attrs ...g.Node) g.Node {
	// Set defaults
	if props.Max == 0 {
		props.Max = 100
	}
	if props.Step == 0 {
		props.Step = 1
	}
	if props.Orientation == "" {
		props.Orientation = "horizontal"
	}
	if len(props.Value) == 0 {
		props.Value = []int{props.Min}
	}

	// Calculate range percentage
	rangeStart := float64(props.Value[0]-props.Min) / float64(props.Max-props.Min) * 100
	rangeEnd := float64(100)
	if len(props.Value) > 1 {
		rangeEnd = float64(props.Value[len(props.Value)-1]-props.Min) / float64(props.Max-props.Min) * 100
	}

	attributes := []g.Node{
		g.Attr("data-slider", "true"),
		g.Attr("data-orientation", props.Orientation),
		html.Class(lib.CN(
			"relative flex w-full touch-none items-center select-none",
			"data-[disabled]:opacity-50",
			"data-[orientation=vertical]:h-full data-[orientation=vertical]:min-h-44",
			"data-[orientation=vertical]:w-auto data-[orientation=vertical]:flex-col",
			props.Class,
		)),
	}

	if props.Disabled {
		attributes = append(attributes, g.Attr("data-disabled", "true"))
	}

	if props.ID != "" {
		attributes = append(attributes, g.Attr("id", props.ID))
	}

	// Build slider structure
	sliderContent := []g.Node{
		// Track
		html.Div(
			g.Attr("data-slider-track", "true"),
			g.Attr("data-orientation", props.Orientation),
			html.Class(lib.CN(
				"bg-muted relative grow overflow-hidden rounded-full",
				"data-[orientation=horizontal]:h-1.5 data-[orientation=horizontal]:w-full",
				"data-[orientation=vertical]:h-full data-[orientation=vertical]:w-1.5",
			)),
			// Range (filled portion)
			html.Div(
				g.Attr("data-slider-range", "true"),
				html.Class(lib.CN(
					"bg-primary absolute",
					"data-[orientation=horizontal]:h-full",
					"data-[orientation=vertical]:w-full",
				)),
				g.If(props.Orientation == "horizontal",
					g.Attr("style", fmt.Sprintf("left: %.2f%%; width: %.2f%%", rangeStart, rangeEnd-rangeStart)),
				),
				g.If(props.Orientation == "vertical",
					g.Attr("style", fmt.Sprintf("bottom: %.2f%%; height: %.2f%%", rangeStart, rangeEnd-rangeStart)),
				),
			),
		),
	}

	// Add thumbs
	for i, value := range props.Value {
		position := float64(value-props.Min) / float64(props.Max-props.Min) * 100
		
		thumb := html.Div(
			g.Attr("data-slider-thumb", "true"),
			g.Attr("data-index", fmt.Sprintf("%d", i)),
			g.Attr("tabindex", "0"),
			g.Attr("role", "slider"),
			g.Attr("aria-valuemin", fmt.Sprintf("%d", props.Min)),
			g.Attr("aria-valuemax", fmt.Sprintf("%d", props.Max)),
			g.Attr("aria-valuenow", fmt.Sprintf("%d", value)),
			g.Attr("aria-orientation", props.Orientation),
			html.Class(lib.CN(
				"border-primary bg-background ring-ring/50",
				"block size-4 shrink-0 rounded-full border shadow-sm",
				"transition-[color,box-shadow] hover:ring-4 focus-visible:ring-4",
				"focus-visible:outline-none disabled:pointer-events-none disabled:opacity-50",
				"absolute",
			)),
			g.If(props.Orientation == "horizontal",
				g.Attr("style", fmt.Sprintf("left: calc(%.2f%% - 8px); top: 50%%; transform: translateY(-50%%)", position)),
			),
			g.If(props.Orientation == "vertical",
				g.Attr("style", fmt.Sprintf("bottom: calc(%.2f%% - 8px); left: 50%%; transform: translateX(-50%%)", position)),
			),
			g.If(props.Disabled,
				g.Attr("disabled", "true"),
			),
		)

		sliderContent = append(sliderContent, thumb)
	}

	// Hidden inputs for form submission
	if props.Name != "" {
		for i, value := range props.Value {
			sliderContent = append(sliderContent,
				html.Input(
					html.Type("hidden"),
					html.Name(fmt.Sprintf("%s[%d]", props.Name, i)),
					html.Value(fmt.Sprintf("%d", value)),
					g.If(props.ID != "", g.Attr("id", fmt.Sprintf("%s-input-%d", props.ID, i))),
				),
			)
		}
	}

	// Add any additional attributes
	sliderContent = g.Group(append(sliderContent, attrs...))

	return html.Div(
		g.Group(append(attributes, sliderContent...)),
	)
}

// Single creates a single-value slider
func Single(props Props) g.Node {
	if len(props.Value) > 1 {
		props.Value = props.Value[:1]
	}
	return New(props)
}

// Range creates a range slider with two thumbs
func Range(props Props) g.Node {
	// Set default max if not provided
	if props.Max == 0 {
		props.Max = 100
	}
	
	// Ensure we have exactly 2 values for range
	if len(props.Value) == 0 {
		props.Value = []int{props.Min, props.Max}
	} else if len(props.Value) == 1 {
		props.Value = append(props.Value, props.Max)
	} else if len(props.Value) > 2 {
		props.Value = props.Value[:2]
	}
	
	// Ensure first value is less than second
	if props.Value[0] > props.Value[1] {
		props.Value[0], props.Value[1] = props.Value[1], props.Value[0]
	}
	
	return New(props)
}

// Vertical creates a vertical slider
func Vertical(props Props) g.Node {
	props.Orientation = "vertical"
	return New(props)
}

// WithLabels creates a slider with min/max labels
func WithLabels(props Props) g.Node {
	labelClass := "text-xs text-muted-foreground"
	
	return html.Div(
		html.Class("space-y-2"),
		// Labels
		html.Div(
			html.Class("flex justify-between"),
			html.Span(html.Class(labelClass), g.Text(fmt.Sprintf("%d", props.Min))),
			html.Span(html.Class(labelClass), g.Text(fmt.Sprintf("%d", props.Max))),
		),
		// Slider
		New(props),
	)
}

// WithValue creates a slider with current value display
func WithValue(props Props) g.Node {
	valueText := fmt.Sprintf("%d", props.Value[0])
	if len(props.Value) > 1 {
		valueText = fmt.Sprintf("%d - %d", props.Value[0], props.Value[1])
	}
	
	return html.Div(
		html.Class("space-y-2"),
		// Value display
		html.Div(
			html.Class("flex justify-between items-center"),
			html.Label(g.Text("Value")),
			html.Span(
				html.Class("text-sm font-medium"),
				g.If(props.ID != "", g.Attr("id", props.ID+"-value")),
				g.Text(valueText),
			),
		),
		// Slider
		New(props),
	)
}

// WithTicks creates a slider with tick marks
func WithTicks(props Props, tickCount int) g.Node {
	if tickCount < 2 {
		tickCount = 2
	}
	
	// Set default orientation if not specified
	if props.Orientation == "" {
		props.Orientation = "horizontal"
	}
	
	// Generate tick positions
	ticks := []g.Node{}
	for i := 0; i < tickCount; i++ {
		position := float64(i) / float64(tickCount-1) * 100
		tick := html.Div(
			html.Class(lib.CN(
				"absolute bg-border",
				lib.CNIf(props.Orientation == "horizontal",
					"w-px h-2 -bottom-3",
					"h-px w-2 -right-3",
				),
			)),
			g.If(props.Orientation == "horizontal",
				g.Attr("style", fmt.Sprintf("left: %.2f%%", position)),
			),
			g.If(props.Orientation == "vertical",
				g.Attr("style", fmt.Sprintf("bottom: %.2f%%", position)),
			),
		)
		ticks = append(ticks, tick)
	}
	
	return html.Div(
		html.Class("relative"),
		New(props),
		// Tick marks container
		html.Div(
			html.Class("absolute inset-0 pointer-events-none"),
			g.Group(ticks),
		),
	)
}