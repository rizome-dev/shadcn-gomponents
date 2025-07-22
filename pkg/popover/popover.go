package popover

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// Props defines the properties for the Popover component
type Props struct {
	Open      bool   // Whether the popover is open
	Side      string // "top" | "right" | "bottom" | "left"
	Align     string // "start" | "center" | "end"
	Class     string // Additional custom classes
	SideOffset int   // Offset from the side
	AlignOffset int  // Offset from the alignment
}

// TriggerProps defines properties for the Popover trigger
type TriggerProps struct {
	AsChild bool   // Whether to render as child element
	Class   string // Additional custom classes
}

// ContentProps defines properties for the Popover content
type ContentProps struct {
	Side        string // "top" | "right" | "bottom" | "left"
	Align       string // "start" | "center" | "end"
	SideOffset  int    // Offset from the side
	AlignOffset int    // Offset from the alignment
	Class       string // Additional custom classes
}

// New creates a new Popover container
func New(props Props, children ...g.Node) g.Node {
	classes := lib.CN("relative inline-block", props.Class)
	
	// Separate trigger from content
	var trigger g.Node
	var content g.Node
	
	for _, child := range children {
		// This is a simplified check - in a real implementation, 
		// we'd need a more robust way to identify components
		if trigger == nil {
			trigger = child
		} else if content == nil {
			content = child
		}
	}
	
	return html.Div(
		html.Class(classes),
		g.If(props.Open, g.Attr("data-state", "open")),
		g.If(!props.Open, g.Attr("data-state", "closed")),
		trigger,
		g.If(props.Open, content),
	)
}

// Trigger creates a Popover trigger element
func Trigger(props TriggerProps, children ...g.Node) g.Node {
	classes := lib.CN(props.Class)
	
	if props.AsChild && len(children) > 0 {
		// If AsChild is true, we modify the first child element
		return children[0]
	}
	
	return html.Button(
		html.Type("button"),
		g.If(classes != "", html.Class(classes)),
		g.Attr("aria-haspopup", "dialog"),
		g.Attr("aria-expanded", "false"),
		g.Group(children),
	)
}

// Content creates the Popover content
func ContentComponent(props ContentProps, children ...g.Node) g.Node {
	// Default values
	if props.Side == "" {
		props.Side = "bottom"
	}
	if props.Align == "" {
		props.Align = "center"
	}
	if props.SideOffset == 0 {
		props.SideOffset = 4
	}
	
	// Build position classes based on side and align
	positionClasses := getPositionClasses(props.Side, props.Align)
	
	classes := lib.CN(
		"z-50 w-72 rounded-md border bg-popover p-4 text-popover-foreground shadow-md outline-none",
		"data-[state=open]:animate-in data-[state=closed]:animate-out",
		"data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0",
		getAnimationClasses(props.Side),
		positionClasses,
		props.Class,
	)
	
	return html.Div(
		html.Class(classes),
		html.Role("dialog"),
		g.If(props.Side != "", g.Attr("data-side", props.Side)),
		g.If(props.Align != "", g.Attr("data-align", props.Align)),
		g.Group(children),
	)
}

// getPositionClasses returns position classes based on side and align
func getPositionClasses(side, align string) string {
	var classes []string
	
	// Position absolute
	classes = append(classes, "absolute")
	
	// Side positioning
	switch side {
	case "top":
		classes = append(classes, "bottom-full mb-2")
	case "right":
		classes = append(classes, "left-full ml-2")
	case "bottom":
		classes = append(classes, "top-full mt-2")
	case "left":
		classes = append(classes, "right-full mr-2")
	}
	
	// Alignment
	switch align {
	case "start":
		if side == "top" || side == "bottom" {
			classes = append(classes, "left-0")
		} else {
			classes = append(classes, "top-0")
		}
	case "center":
		if side == "top" || side == "bottom" {
			classes = append(classes, "left-1/2 -translate-x-1/2")
		} else {
			classes = append(classes, "top-1/2 -translate-y-1/2")
		}
	case "end":
		if side == "top" || side == "bottom" {
			classes = append(classes, "right-0")
		} else {
			classes = append(classes, "bottom-0")
		}
	}
	
	return lib.CN(classes...)
}

// getAnimationClasses returns animation classes based on side
func getAnimationClasses(side string) string {
	switch side {
	case "top":
		return "data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[state=closed]:slide-out-to-bottom-2 data-[state=open]:slide-in-from-bottom-2"
	case "right":
		return "data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[state=closed]:slide-out-to-left-2 data-[state=open]:slide-in-from-left-2"
	case "bottom":
		return "data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[state=closed]:slide-out-to-top-2 data-[state=open]:slide-in-from-top-2"
	case "left":
		return "data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[state=closed]:slide-out-to-right-2 data-[state=open]:slide-in-from-right-2"
	default:
		return "data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95"
	}
}

// WithArrow creates a popover content with an arrow
func WithArrow(props ContentProps, children ...g.Node) g.Node {
	arrowClasses := getArrowClasses(props.Side)
	
	return ContentComponent(
		props,
		append([]g.Node{
			// Arrow element
			html.Div(
				html.Class(arrowClasses),
			),
		}, children...)...,
	)
}

// getArrowClasses returns arrow classes based on side
func getArrowClasses(side string) string {
	base := "absolute w-2 h-2 bg-popover border rotate-45"
	
	switch side {
	case "top":
		return lib.CN(base, "bottom-[-5px] left-1/2 -translate-x-1/2 border-t-0 border-l-0")
	case "right":
		return lib.CN(base, "left-[-5px] top-1/2 -translate-y-1/2 border-r-0 border-t-0")
	case "bottom":
		return lib.CN(base, "top-[-5px] left-1/2 -translate-x-1/2 border-b-0 border-r-0")
	case "left":
		return lib.CN(base, "right-[-5px] top-1/2 -translate-y-1/2 border-l-0 border-b-0")
	default:
		return lib.CN(base, "top-[-5px] left-1/2 -translate-x-1/2 border-b-0 border-r-0")
	}
}

// Close creates a close button for the popover
func Close(class ...string) g.Node {
	classes := lib.CN(
		"absolute right-2 top-2 rounded-sm opacity-70 ring-offset-background transition-opacity hover:opacity-100 focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 disabled:pointer-events-none",
		lib.CN(class...),
	)
	
	return html.Button(
		html.Type("button"),
		html.Class(classes),
		g.Attr("aria-label", "Close"),
		// X icon
		g.El("svg",
			g.Attr("xmlns", "http://www.w3.org/2000/svg"),
			g.Attr("viewBox", "0 0 24 24"),
			g.Attr("fill", "none"),
			g.Attr("stroke", "currentColor"),
			g.Attr("stroke-width", "2"),
			g.Attr("stroke-linecap", "round"),
			g.Attr("stroke-linejoin", "round"),
			html.Class("h-4 w-4"),
			g.El("line", g.Attr("x1", "18"), g.Attr("y1", "6"), g.Attr("x2", "6"), g.Attr("y2", "18")),
			g.El("line", g.Attr("x1", "6"), g.Attr("y1", "6"), g.Attr("x2", "18"), g.Attr("y2", "18")),
		),
	)
}