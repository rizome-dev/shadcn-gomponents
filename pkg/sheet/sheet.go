package sheet

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// Props defines the properties for the Sheet component
type Props struct {
	Open      bool   // Whether the sheet is open
	Side      string // "top" | "right" | "bottom" | "left"
	Class     string // Additional custom classes
	OnOpenChange string // JavaScript callback for open state changes
}

// TriggerProps defines properties for the Sheet trigger
type TriggerProps struct {
	AsChild bool   // Whether to render as child element
	Class   string // Additional custom classes
}

// ContentProps defines properties for the Sheet content
type ContentProps struct {
	Side             string // "top" | "right" | "bottom" | "left"
	Class            string // Additional custom classes
	ShowCloseButton  bool   // Whether to show the close button
	CloseOnOverlay   bool   // Whether clicking overlay closes the sheet
	CloseOnEsc       bool   // Whether ESC key closes the sheet
}

// OverlayProps defines properties for the Sheet overlay
type OverlayProps struct {
	Class string // Additional custom classes
}

// HeaderProps defines properties for the Sheet header
type HeaderProps struct {
	Class string // Additional custom classes
}

// TitleProps defines properties for the Sheet title
type TitleProps struct {
	Class string // Additional custom classes
}

// DescriptionProps defines properties for the Sheet description
type DescriptionProps struct {
	Class string // Additional custom classes
}

// FooterProps defines properties for the Sheet footer
type FooterProps struct {
	Class string // Additional custom classes
}

// CloseProps defines properties for the close button
type CloseProps struct {
	Class string // Additional custom classes
}

// New creates a new Sheet container
func New(props Props, children ...g.Node) g.Node {
	if props.Open {
		classes := lib.CN("fixed inset-0 z-50", props.Class)
		return html.Div(
			html.Class(classes),
			g.Attr("data-state", "open"),
			g.Group(children),
		)
	}
	// Return empty div when closed
	return html.Div(g.Attr("data-state", "closed"))
}

// Trigger creates a Sheet trigger element
func Trigger(props TriggerProps, children ...g.Node) g.Node {
	classes := lib.CN(props.Class)
	
	if props.AsChild && len(children) > 0 {
		// If AsChild is true, we would modify the first child element
		// In a real implementation, this would clone and modify the child
		return children[0]
	}
	
	return html.Button(
		html.Type("button"),
		g.If(classes != "", html.Class(classes)),
		g.Attr("aria-haspopup", "dialog"),
		g.Group(children),
	)
}

// Overlay creates the Sheet overlay/backdrop
func Overlay(props OverlayProps) g.Node {
	classes := lib.CN(
		"fixed inset-0 z-50 bg-black/80",
		"data-[state=open]:animate-in data-[state=closed]:animate-out",
		"data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0",
		props.Class,
	)
	
	return html.Div(
		html.Class(classes),
		g.Attr("data-sheet-overlay", ""),
		g.Attr("data-state", "open"),
	)
}

// Content creates the Sheet content
func ContentComponent(props ContentProps, children ...g.Node) g.Node {
	// Default side to right
	if props.Side == "" {
		props.Side = "right"
	}
	
	// Base classes
	baseClasses := "fixed z-50 gap-4 bg-background p-6 shadow-lg transition ease-in-out"
	animationClasses := "data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:duration-300 data-[state=open]:duration-500"
	
	// Side-specific classes
	var sideClasses string
	var slideClasses string
	
	switch props.Side {
	case "top":
		sideClasses = "inset-x-0 top-0 border-b"
		slideClasses = "data-[state=closed]:slide-out-to-top data-[state=open]:slide-in-from-top"
	case "bottom":
		sideClasses = "inset-x-0 bottom-0 border-t"
		slideClasses = "data-[state=closed]:slide-out-to-bottom data-[state=open]:slide-in-from-bottom"
	case "left":
		sideClasses = "inset-y-0 left-0 h-full w-3/4 border-r sm:max-w-sm"
		slideClasses = "data-[state=closed]:slide-out-to-left data-[state=open]:slide-in-from-left"
	case "right":
		sideClasses = "inset-y-0 right-0 h-full w-3/4 border-l sm:max-w-sm"
		slideClasses = "data-[state=closed]:slide-out-to-right data-[state=open]:slide-in-from-right"
	}
	
	classes := lib.CN(
		baseClasses,
		animationClasses,
		sideClasses,
		slideClasses,
		props.Class,
	)
	
	contentChildren := children
	if props.ShowCloseButton {
		closeButton := html.Button(
			html.Type("button"),
			html.Class("absolute right-4 top-4 rounded-sm opacity-70 ring-offset-background transition-opacity hover:opacity-100 focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 disabled:pointer-events-none"),
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
			html.Span(html.Class("sr-only"), g.Text("Close")),
		)
		contentChildren = append([]g.Node{closeButton}, children...)
	}
	
	return html.Div(
		html.Class(classes),
		html.Role("dialog"),
		g.Attr("aria-modal", "true"),
		g.Attr("data-state", "open"),
		g.Attr("data-sheet-content", ""),
		g.If(props.Side != "", g.Attr("data-side", props.Side)),
		g.Group(contentChildren),
	)
}

// Header creates a Sheet header
func HeaderComponent(props HeaderProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"flex flex-col space-y-2 text-center sm:text-left",
		props.Class,
	)
	
	return html.Div(
		html.Class(classes),
		g.Group(children),
	)
}

// Footer creates a Sheet footer
func FooterComponent(props FooterProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"flex flex-col-reverse sm:flex-row sm:justify-end sm:space-x-2",
		props.Class,
	)
	
	return html.Div(
		html.Class(classes),
		g.Group(children),
	)
}

// Title creates a Sheet title
func Title(props TitleProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"text-lg font-semibold text-foreground",
		props.Class,
	)
	
	return html.H2(
		html.Class(classes),
		g.Group(children),
	)
}

// Description creates a Sheet description
func Description(props DescriptionProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"text-sm text-muted-foreground",
		props.Class,
	)
	
	return html.P(
		html.Class(classes),
		g.Group(children),
	)
}

// Close creates a close button for the sheet
func Close(props CloseProps, children ...g.Node) g.Node {
	classes := lib.CN(props.Class)
	
	return html.Button(
		html.Type("button"),
		g.If(classes != "", html.Class(classes)),
		g.Group(children),
	)
}

// WithForm creates a sheet with a form
func WithForm(props Props, contentProps ContentProps, formAction string, children ...g.Node) g.Node {
	return New(
		props,
		Overlay(OverlayProps{}),
		ContentComponent(
			contentProps,
			html.Form(
				html.Action(formAction),
				html.Method("POST"),
				g.Group(children),
			),
		),
	)
}

// RightSheet creates a right-side sheet
func RightSheet(props Props, children ...g.Node) g.Node {
	return New(
		props,
		Overlay(OverlayProps{}),
		ContentComponent(
			ContentProps{Side: "right", ShowCloseButton: true},
			children...,
		),
	)
}

// LeftSheet creates a left-side sheet
func LeftSheet(props Props, children ...g.Node) g.Node {
	return New(
		props,
		Overlay(OverlayProps{}),
		ContentComponent(
			ContentProps{Side: "left", ShowCloseButton: true},
			children...,
		),
	)
}

// TopSheet creates a top sheet
func TopSheet(props Props, children ...g.Node) g.Node {
	return New(
		props,
		Overlay(OverlayProps{}),
		ContentComponent(
			ContentProps{Side: "top", ShowCloseButton: true},
			children...,
		),
	)
}

// BottomSheet creates a bottom sheet
func BottomSheet(props Props, children ...g.Node) g.Node {
	return New(
		props,
		Overlay(OverlayProps{}),
		ContentComponent(
			ContentProps{Side: "bottom", ShowCloseButton: true},
			children...,
		),
	)
}

// MobileSheet creates a sheet optimized for mobile (bottom sheet on mobile, right sheet on desktop)
func MobileSheet(props Props, children ...g.Node) g.Node {
	return New(
		props,
		Overlay(OverlayProps{}),
		ContentComponent(
			ContentProps{
				Side:            "bottom",
				ShowCloseButton: true,
				Class:           "sm:max-w-lg h-[90vh] sm:h-full sm:inset-y-0 sm:right-0 sm:left-auto sm:bottom-auto sm:border-l sm:border-t-0",
			},
			children...,
		),
	)
}