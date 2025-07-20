package resizable

import (
	"fmt"
	
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// Props defines the properties for the Resizable component
type Props struct {
	Direction        string // "horizontal" | "vertical"
	OnResizeStart    string // JavaScript callback
	OnResizeEnd      string // JavaScript callback
	Class            string // Additional custom classes
	DefaultSize      int    // Default size percentage
	MinSize          int    // Minimum size percentage
	MaxSize          int    // Maximum size percentage
	CollapsedSize    int    // Collapsed size in pixels
	Collapsible      bool   // Whether panels can be collapsed
	Storage          bool   // Whether to persist sizes in localStorage
	StorageKey       string // Key for localStorage
}

// PanelProps defines properties for individual panels
type PanelProps struct {
	DefaultSize   int    // Default size percentage
	MinSize       int    // Minimum size percentage
	MaxSize       int    // Maximum size percentage
	Collapsible   bool   // Whether this panel can be collapsed
	CollapsedSize int    // Size when collapsed
	ID            string // Panel ID
	Order         int    // Panel order
	Class         string // Additional custom classes
}

// HandleProps defines properties for the resize handle
type HandleProps struct {
	WithHandle bool   // Whether to show the visual handle
	Disabled   bool   // Whether resizing is disabled
	Class      string // Additional custom classes
}

// PanelGroup creates a resizable panel group container
func PanelGroup(props Props, children ...g.Node) g.Node {
	// Set defaults
	if props.Direction == "" {
		props.Direction = "horizontal"
	}
	
	classes := lib.CN(
		"flex h-full w-full data-[panel-group-direction=vertical]:flex-col",
		props.Class,
	)
	
	attrs := []g.Node{
		html.Class(classes),
		g.Attr("data-panel-group", ""),
		g.Attr("data-panel-group-direction", props.Direction),
		g.If(props.Storage && props.StorageKey != "", g.Attr("data-panel-group-storage", props.StorageKey)),
	}
	
	// Add resize event handlers if provided
	if props.OnResizeStart != "" {
		attrs = append(attrs, g.Attr("data-onresizestart", props.OnResizeStart))
	}
	if props.OnResizeEnd != "" {
		attrs = append(attrs, g.Attr("data-onresizeend", props.OnResizeEnd))
	}
	
	return html.Div(append(attrs, children...)...)
}

// Panel creates a resizable panel
func Panel(props PanelProps, children ...g.Node) g.Node {
	// Calculate flex values
	defaultSize := props.DefaultSize
	if defaultSize == 0 {
		defaultSize = 50 // Default to 50%
	}
	
	minSize := props.MinSize
	if minSize == 0 {
		minSize = 10 // Default minimum 10%
	}
	
	maxSize := props.MaxSize
	if maxSize == 0 {
		maxSize = 90 // Default maximum 90%
	}
	
	classes := lib.CN(
		"relative",
		props.Class,
	)
	
	attrs := []g.Node{
		html.Class(classes),
		g.Attr("data-panel", ""),
		g.If(props.ID != "", html.ID(props.ID)),
		g.If(props.ID != "", g.Attr("data-panel-id", props.ID)),
		g.Attr("data-panel-size", fmt.Sprintf("%d", defaultSize)),
		g.Attr("data-panel-min-size", fmt.Sprintf("%d", minSize)),
		g.Attr("data-panel-max-size", fmt.Sprintf("%d", maxSize)),
		g.If(props.Collapsible, g.Attr("data-panel-collapsible", "true")),
		g.If(props.CollapsedSize > 0, g.Attr("data-panel-collapsed-size", fmt.Sprintf("%d", props.CollapsedSize))),
		g.If(props.Order > 0, g.Attr("data-panel-order", fmt.Sprintf("%d", props.Order))),
		html.Style(fmt.Sprintf("flex: %d %d 0%%", defaultSize, defaultSize)),
	}
	
	return html.Div(append(attrs, children...)...)
}

// Handle creates a resize handle between panels
func Handle(props HandleProps) g.Node {
	classes := lib.CN(
		"relative flex items-center justify-center",
		"bg-border",
		"after:absolute after:inset-0 hover:bg-accent",
		"data-[panel-group-direction=horizontal]:h-full data-[panel-group-direction=horizontal]:w-1 data-[panel-group-direction=horizontal]:cursor-col-resize",
		"data-[panel-group-direction=vertical]:h-1 data-[panel-group-direction=vertical]:w-full data-[panel-group-direction=vertical]:cursor-row-resize",
		lib.CNIf(props.Disabled, "cursor-not-allowed opacity-50", ""),
		"[&[data-panel-group-direction=vertical]>div]:rotate-90",
		props.Class,
	)
	
	handle := html.Div(
		html.Class(classes),
		html.Role("separator"),
		g.Attr("aria-valuenow", "50"),
		g.Attr("aria-valuemin", "0"),
		g.Attr("aria-valuemax", "100"),
		g.Attr("aria-orientation", "horizontal"),
		html.TabIndex("0"),
		g.Attr("data-panel-resize-handle", ""),
		g.If(props.Disabled, g.Attr("data-disabled", "")),
		
		// Visual handle indicator if requested
		g.If(props.WithHandle,
			html.Div(
				html.Class("z-10 flex h-4 w-3 items-center justify-center rounded-sm border bg-border"),
				// Handle dots
				g.El("svg",
					g.Attr("viewBox", "0 0 24 24"),
					g.Attr("width", "24"),
					g.Attr("height", "24"),
					g.Attr("fill", "currentColor"),
					html.Class("h-2.5 w-2.5"),
					g.El("path", g.Attr("fill-rule", "evenodd"), g.Attr("clip-rule", "evenodd"), 
						g.Attr("d", "M8 6a2 2 0 1 0 0-4 2 2 0 0 0 0 4zM8 14a2 2 0 1 0 0-4 2 2 0 0 0 0 4zM8 22a2 2 0 1 0 0-4 2 2 0 0 0 0 4zM16 6a2 2 0 1 0 0-4 2 2 0 0 0 0 4zM16 14a2 2 0 1 0 0-4 2 2 0 0 0 0 4zM16 22a2 2 0 1 0 0-4 2 2 0 0 0 0 4z"),
					),
				),
			),
		),
	)
	
	return handle
}

// HorizontalPanelGroup creates a horizontal panel group
func HorizontalPanelGroup(props Props, children ...g.Node) g.Node {
	props.Direction = "horizontal"
	return PanelGroup(props, children...)
}

// VerticalPanelGroup creates a vertical panel group
func VerticalPanelGroup(props Props, children ...g.Node) g.Node {
	props.Direction = "vertical"
	return PanelGroup(props, children...)
}

// CollapsiblePanel creates a panel that can be collapsed
func CollapsiblePanel(props PanelProps, children ...g.Node) g.Node {
	props.Collapsible = true
	if props.CollapsedSize == 0 {
		props.CollapsedSize = 4 // Default collapsed size
	}
	return Panel(props, children...)
}

// ResizableLayout creates a common resizable layout pattern
func ResizableLayout(sidebarContent, mainContent g.Node) g.Node {
	return PanelGroup(
		Props{Direction: "horizontal", Storage: true, StorageKey: "layout"},
		CollapsiblePanel(
			PanelProps{
				DefaultSize:   20,
				MinSize:       15,
				MaxSize:       30,
				CollapsedSize: 4,
				ID:            "sidebar",
			},
			sidebarContent,
		),
		Handle(HandleProps{WithHandle: true}),
		Panel(
			PanelProps{
				DefaultSize: 80,
				ID:          "main",
			},
			mainContent,
		),
	)
}

// TwoColumnLayout creates a simple two-column resizable layout
func TwoColumnLayout(leftContent, rightContent g.Node, leftSize, rightSize int) g.Node {
	if leftSize == 0 {
		leftSize = 50
	}
	if rightSize == 0 {
		rightSize = 50
	}
	
	return HorizontalPanelGroup(
		Props{},
		Panel(
			PanelProps{DefaultSize: leftSize},
			leftContent,
		),
		Handle(HandleProps{}),
		Panel(
			PanelProps{DefaultSize: rightSize},
			rightContent,
		),
	)
}

// ThreeColumnLayout creates a three-column resizable layout
func ThreeColumnLayout(leftContent, centerContent, rightContent g.Node) g.Node {
	return HorizontalPanelGroup(
		Props{},
		Panel(
			PanelProps{DefaultSize: 25, MinSize: 15, MaxSize: 35},
			leftContent,
		),
		Handle(HandleProps{}),
		Panel(
			PanelProps{DefaultSize: 50, MinSize: 30},
			centerContent,
		),
		Handle(HandleProps{}),
		Panel(
			PanelProps{DefaultSize: 25, MinSize: 15, MaxSize: 35},
			rightContent,
		),
	)
}

// IDELayout creates an IDE-like layout with collapsible sidebar and bottom panel
func IDELayout(sidebarContent, editorContent, terminalContent g.Node) g.Node {
	return VerticalPanelGroup(
		Props{Class: "h-screen"},
		// Top section with sidebar and editor
		Panel(
			PanelProps{DefaultSize: 70},
			HorizontalPanelGroup(
				Props{},
				CollapsiblePanel(
					PanelProps{
						DefaultSize:   20,
						MinSize:       15,
						MaxSize:       35,
						CollapsedSize: 4,
						ID:            "sidebar",
					},
					sidebarContent,
				),
				Handle(HandleProps{WithHandle: true}),
				Panel(
					PanelProps{DefaultSize: 80, ID: "editor"},
					editorContent,
				),
			),
		),
		Handle(HandleProps{WithHandle: true}),
		// Bottom terminal panel
		CollapsiblePanel(
			PanelProps{
				DefaultSize:   30,
				MinSize:       20,
				MaxSize:       50,
				CollapsedSize: 4,
				ID:            "terminal",
			},
			terminalContent,
		),
	)
}