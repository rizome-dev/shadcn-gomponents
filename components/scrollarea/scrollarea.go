package scrollarea

import (
	"fmt"
	
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// Props defines the properties for the ScrollArea component
type Props struct {
	Orientation string // "vertical" | "horizontal" | "both"
	Type        string // "auto" | "always" | "scroll" | "hover"
	ScrollHideDelay int    // Delay in ms before hiding scrollbar
	Dir         string // "ltr" | "rtl"
	Class       string // Additional custom classes
}

// ViewportProps defines properties for the viewport
type ViewportProps struct {
	Class string // Additional custom classes
}

// ScrollbarProps defines properties for the scrollbar
type ScrollbarProps struct {
	Orientation string // "vertical" | "horizontal"
	ForceMount  bool   // Always render the scrollbar
	Class       string // Additional custom classes
}

// ThumbProps defines properties for the scrollbar thumb
type ThumbProps struct {
	Class string // Additional custom classes
}

// CornerProps defines properties for the corner where scrollbars meet
type CornerProps struct {
	Class string // Additional custom classes
}

// New creates a new ScrollArea container
func New(props Props, children ...g.Node) g.Node {
	// Set defaults
	if props.Orientation == "" {
		props.Orientation = "vertical"
	}
	if props.Type == "" {
		props.Type = "hover"
	}
	
	classes := lib.CN(
		"relative overflow-hidden",
		props.Class,
	)
	
	attrs := []g.Node{
		html.Class(classes),
		g.Attr("data-scroll-area", ""),
		g.If(props.Orientation != "", g.Attr("data-orientation", props.Orientation)),
		g.If(props.Type != "", g.Attr("data-type", props.Type)),
		g.If(props.Dir != "", g.Attr("dir", props.Dir)),
		g.If(props.ScrollHideDelay > 0, g.Attr("data-scroll-hide-delay", fmt.Sprintf("%d", props.ScrollHideDelay))),
	}
	
	return html.Div(append(attrs, children...)...)
}

// Viewport creates the scrollable viewport
func Viewport(props ViewportProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"h-full w-full rounded-[inherit]",
		props.Class,
	)
	
	return html.Div(
		html.Class(classes),
		g.Attr("data-scroll-area-viewport", ""),
		html.Style("overflow: scroll; -ms-overflow-style: none; scrollbar-width: none;"),
		// Hide webkit scrollbar
		g.Attr("style", "overflow: scroll; -ms-overflow-style: none; scrollbar-width: none; &::-webkit-scrollbar { display: none; }"),
		g.Group(children),
	)
}

// Scrollbar creates a scrollbar element
func Scrollbar(props ScrollbarProps, children ...g.Node) g.Node {
	// Set default orientation
	if props.Orientation == "" {
		props.Orientation = "vertical"
	}
	
	orientationClasses := ""
	if props.Orientation == "vertical" {
		orientationClasses = "h-full w-2.5 border-l border-l-transparent p-[1px]"
	} else {
		orientationClasses = "h-2.5 w-full border-t border-t-transparent p-[1px]"
	}
	
	classes := lib.CN(
		"flex touch-none select-none transition-colors",
		orientationClasses,
		props.Class,
	)
	
	attrs := []g.Node{
		html.Class(classes),
		g.Attr("data-scroll-area-scrollbar", ""),
		g.Attr("data-orientation", props.Orientation),
		g.If(props.ForceMount, g.Attr("data-state", "visible")),
	}
	
	return html.Div(append(attrs, children...)...)
}

// Thumb creates the scrollbar thumb
func Thumb(props ThumbProps) g.Node {
	classes := lib.CN(
		"relative flex-1 rounded-full bg-border",
		props.Class,
	)
	
	return html.Div(
		html.Class(classes),
		g.Attr("data-scroll-area-thumb", ""),
		html.Style("position: relative;"),
	)
}

// Corner creates the corner element where scrollbars meet
func Corner(props CornerProps) g.Node {
	classes := lib.CN(
		"absolute right-0 bottom-0",
		props.Class,
	)
	
	return html.Div(
		html.Class(classes),
		g.Attr("data-scroll-area-corner", ""),
	)
}

// ScrollAreaWithBar creates a scroll area with visible scrollbars
func ScrollAreaWithBar(props Props, children ...g.Node) g.Node {
	props.Type = "always"
	
	return New(
		props,
		Viewport(
			ViewportProps{},
			children...,
		),
		Scrollbar(
			ScrollbarProps{Orientation: "vertical"},
			Thumb(ThumbProps{}),
		),
		g.If(props.Orientation == "both",
			Scrollbar(
				ScrollbarProps{Orientation: "horizontal"},
				Thumb(ThumbProps{}),
			),
		),
		g.If(props.Orientation == "both",
			Corner(CornerProps{}),
		),
	)
}

// HorizontalScrollArea creates a horizontal scroll area
func HorizontalScrollArea(props Props, children ...g.Node) g.Node {
	props.Orientation = "horizontal"
	
	return New(
		props,
		Viewport(
			ViewportProps{},
			children...,
		),
		Scrollbar(
			ScrollbarProps{Orientation: "horizontal"},
			Thumb(ThumbProps{}),
		),
	)
}

// ScrollAreaAuto creates a scroll area with auto-hiding scrollbars
func ScrollAreaAuto(props Props, children ...g.Node) g.Node {
	props.Type = "auto"
	
	return New(
		props,
		Viewport(
			ViewportProps{},
			children...,
		),
		Scrollbar(
			ScrollbarProps{Orientation: "vertical"},
			Thumb(ThumbProps{}),
		),
	)
}

// ScrollAreaHover creates a scroll area with hover-activated scrollbars
func ScrollAreaHover(props Props, children ...g.Node) g.Node {
	props.Type = "hover"
	props.ScrollHideDelay = 600 // Default hide delay
	
	return New(
		props,
		Viewport(
			ViewportProps{},
			children...,
		),
		Scrollbar(
			ScrollbarProps{Orientation: "vertical"},
			Thumb(ThumbProps{}),
		),
	)
}

// CodeScrollArea creates a scroll area optimized for code blocks
func CodeScrollArea(class string, code ...g.Node) g.Node {
	return ScrollAreaWithBar(
		Props{
			Orientation: "both",
			Class:       lib.CN("rounded-md border", class),
		},
		Viewport(
			ViewportProps{},
			html.Pre(
				html.Class("p-4"),
				html.Code(
					html.Class("text-sm"),
					g.Group(code),
				),
			),
		),
	)
}

// ListScrollArea creates a scroll area optimized for lists
func ListScrollArea(maxHeight string, items ...g.Node) g.Node {
	return ScrollAreaAuto(
		Props{
			Class: lib.CN("rounded-md border", fmt.Sprintf("max-h-[%s]", maxHeight)),
		},
		Viewport(
			ViewportProps{},
			html.Div(
				html.Class("p-4"),
				g.Group(items),
			),
		),
	)
}

// ChatScrollArea creates a scroll area optimized for chat interfaces
func ChatScrollArea(messages ...g.Node) g.Node {
	return New(
		Props{
			Type:  "auto",
			Class: "h-full",
		},
		Viewport(
			ViewportProps{Class: "pb-4"},
			html.Div(
				html.Class("flex flex-col gap-4 p-4"),
				g.Group(messages),
			),
		),
		Scrollbar(
			ScrollbarProps{Orientation: "vertical"},
			Thumb(ThumbProps{}),
		),
	)
}

// ImageGalleryScrollArea creates a horizontal scroll area for image galleries
func ImageGalleryScrollArea(images ...g.Node) g.Node {
	return HorizontalScrollArea(
		Props{
			Class: "w-full whitespace-nowrap rounded-md border",
		},
		html.Div(
			html.Class("flex w-max space-x-4 p-4"),
			g.Group(images),
		),
	)
}

// TableScrollArea creates a scroll area for tables with both scrollbars
func TableScrollArea(table g.Node) g.Node {
	return ScrollAreaWithBar(
		Props{
			Orientation: "both",
			Class:       "rounded-md border",
		},
		Viewport(
			ViewportProps{},
			table,
		),
	)
}