package collapsible

import (
	"io"

	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
	"github.com/rizome-dev/shadcn-gomponents/lib/icons"
)

// Props defines the properties for the Collapsible component
type Props struct {
	Open  bool   // Whether the collapsible is open
	Class string // Additional custom classes
}

// TriggerProps defines the properties for the CollapsibleTrigger
type TriggerProps struct {
	Class string
}

// ContentProps defines the properties for the CollapsibleContent
type ContentProps struct {
	Class string
}

// New creates a new Collapsible component using HTML details/summary elements
// This provides native collapsible functionality without JavaScript
func New(props Props, children ...g.Node) g.Node {
	classes := lib.CN(
		props.Class,
	)

	// Split children into summary (trigger) and content
	var summaryNodes []g.Node
	var contentNodes []g.Node
	inContent := false

	for _, child := range children {
		if _, ok := child.(contentMarker); ok {
			inContent = true
			continue
		}
		if inContent {
			contentNodes = append(contentNodes, child)
		} else {
			summaryNodes = append(summaryNodes, child)
		}
	}

	attrs := []g.Node{
		g.If(classes != "", html.Class(classes)),
	}
	if props.Open {
		attrs = append(attrs, g.Attr("open", ""))
	}

	return g.El("details",
		append(attrs,
			g.El("summary", summaryNodes...),
			g.Group(contentNodes),
		)...,
	)
}

// contentMarker is a zero-size type used to mark where content begins
type contentMarker struct{}

func (contentMarker) Render(io.Writer) error { return nil }

// ContentMarker marks the boundary between trigger and content in a collapsible
var ContentMarker = contentMarker{}

// Alternative implementation using div-based structure for more control
// This requires JavaScript for interactivity but provides more styling flexibility

// DivCollapsible creates a div-based collapsible (requires JavaScript for interactivity)
func DivCollapsible(props Props, children ...g.Node) g.Node {
	classes := lib.CN(
		"space-y-2",
		props.Class,
	)

	dataState := "closed"
	if props.Open {
		dataState = "open"
	}

	return html.Div(
		html.Class(classes),
		g.Attr("data-state", dataState),
		g.Group(children),
	)
}

// Trigger creates a CollapsibleTrigger component
func Trigger(props TriggerProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"flex items-center justify-between cursor-pointer",
		props.Class,
	)

	return html.Div(
		html.Class(classes),
		g.Attr("role", "button"),
		g.Attr("tabindex", "0"),
		g.Group(children),
	)
}

// CollapsibleContent creates a CollapsibleContent component
func CollapsibleContent(props ContentProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"overflow-hidden transition-all data-[state=closed]:animate-accordion-up data-[state=open]:animate-accordion-down",
		props.Class,
	)

	return html.Div(
		html.Class(classes),
		g.Attr("data-state", "closed"), // Default state, should be controlled by JavaScript
		g.Group(children),
	)
}

// TriggerButton creates a trigger button with chevron icon
func TriggerButton(props TriggerProps, isOpen bool) g.Node {
	classes := lib.CN(
		"inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium transition-all disabled:pointer-events-none disabled:opacity-50 outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px]",
		"hover:bg-accent hover:text-accent-foreground",
		"h-9 w-9 p-0",
		props.Class,
	)

	iconClasses := lib.CN(
		"h-4 w-4 transition-transform duration-200",
		lib.CNIf(isOpen, "transform rotate-180", ""),
	)

	return html.Button(
		html.Type("button"),
		html.Class(classes),
		icons.ChevronsUpDown(html.Class(iconClasses)),
		html.Span(html.Class("sr-only"), g.Text("Toggle")),
	)
}

// Example creates a basic collapsible example using details/summary
func Example() g.Node {
	return g.El("details",
		html.Class("w-[350px] space-y-2"),
		g.Attr("open", ""),
		g.El("summary",
			html.Class("flex items-center justify-between space-x-4 px-4 cursor-pointer list-none"),
			html.Div(
				html.H4(html.Class("text-sm font-semibold"), g.Text("@peduarte starred 3 repositories")),
			),
			html.Div(
				html.Class("inline-flex items-center justify-center rounded-md text-sm font-medium transition-colors hover:bg-accent hover:text-accent-foreground h-9 w-9 p-0"),
				icons.ChevronsUpDown(html.Class("h-4 w-4")),
			),
		),
		html.Div(html.Class("space-y-2 px-4"),
			html.Div(html.Class("rounded-md border px-4 py-3 font-mono text-sm"),
				g.Text("@radix-ui/primitives"),
			),
			html.Div(html.Class("rounded-md border px-4 py-3 font-mono text-sm"),
				g.Text("@radix-ui/colors"),
			),
			html.Div(html.Class("rounded-md border px-4 py-3 font-mono text-sm"),
				g.Text("@stitches/react"),
			),
		),
	)
}

// ExampleClosed creates a closed collapsible example
func ExampleClosed() g.Node {
	return g.El("details",
		html.Class("w-[350px] space-y-2"),
		// No "open" attribute means it starts closed
		g.El("summary",
			html.Class("flex items-center justify-between space-x-4 px-4 cursor-pointer list-none"),
			html.Div(
				html.H4(html.Class("text-sm font-semibold"), g.Text("Click to expand")),
			),
			html.Div(
				html.Class("inline-flex items-center justify-center rounded-md text-sm font-medium transition-colors hover:bg-accent hover:text-accent-foreground h-9 w-9 p-0"),
				icons.ChevronsUpDown(html.Class("h-4 w-4")),
			),
		),
		html.Div(html.Class("space-y-2 px-4"),
			html.Div(html.Class("rounded-md border px-4 py-3 font-mono text-sm"),
				g.Text("Hidden content 1"),
			),
			html.Div(html.Class("rounded-md border px-4 py-3 font-mono text-sm"),
				g.Text("Hidden content 2"),
			),
		),
	)
}

// ExampleStyled creates a styled collapsible with custom appearance
func ExampleStyled() g.Node {
	return g.El("details",
		html.Class("w-full max-w-md mx-auto"),
		g.El("summary",
			html.Class("flex items-center justify-between p-4 bg-muted rounded-lg cursor-pointer list-none hover:bg-muted/80 transition-colors"),
			html.Div(html.Class("flex items-center gap-3"),
				// Icon placeholder
				html.Div(html.Class("w-10 h-10 bg-primary/10 rounded-full flex items-center justify-center"),
					g.Text("📦"),
				),
				html.Div(
					html.H3(html.Class("font-semibold"), g.Text("Package Details")),
					html.P(html.Class("text-sm text-muted-foreground"), g.Text("View package information")),
				),
			),
			icons.ChevronRight(html.Class("h-5 w-5 transition-transform [details[open]>&]:rotate-90")),
		),
		html.Div(html.Class("p-4 space-y-4"),
			html.Div(html.Class("grid gap-2"),
				html.Div(html.Class("flex justify-between"),
					html.Span(html.Class("text-sm font-medium"), g.Text("Version")),
					html.Span(html.Class("text-sm text-muted-foreground"), g.Text("1.0.0")),
				),
				html.Div(html.Class("flex justify-between"),
					html.Span(html.Class("text-sm font-medium"), g.Text("License")),
					html.Span(html.Class("text-sm text-muted-foreground"), g.Text("MIT")),
				),
				html.Div(html.Class("flex justify-between"),
					html.Span(html.Class("text-sm font-medium"), g.Text("Downloads")),
					html.Span(html.Class("text-sm text-muted-foreground"), g.Text("1.2M")),
				),
			),
		),
	)
}