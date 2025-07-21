package tooltip

import (
	"fmt"

	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// Side defines the preferred side for the tooltip
type Side string

const (
	SideTop    Side = "top"
	SideRight  Side = "right"
	SideBottom Side = "bottom"
	SideLeft   Side = "left"
)

// Align defines the alignment of the tooltip
type Align string

const (
	AlignStart  Align = "start"
	AlignCenter Align = "center"
	AlignEnd    Align = "end"
)

// Props defines the properties for the Tooltip component
type Props struct {
	ID           string // HTML id attribute
	Content      string // The tooltip content
	Side         Side   // Preferred side: "top" | "right" | "bottom" | "left"
	Align        Align  // Alignment: "start" | "center" | "end"
	DelayMs      int    // Delay in milliseconds before showing
	SideOffset   int    // Distance from trigger in pixels
	AlignOffset  int    // Offset along the side in pixels
	Class        string // Additional custom classes for content
	ContentClass string // Classes specifically for content wrapper
	ArrowClass   string // Classes for the arrow
	Open         bool   // Whether tooltip is open (for controlled mode)
	AsChild      bool   // Whether to render trigger as child
}

// TriggerProps defines properties for the tooltip trigger
type TriggerProps struct {
	ID        string // HTML id for the trigger
	Class     string // Additional custom classes
	AsChild   bool   // Whether to render as child
	OnHover   string // JavaScript onMouseEnter handler
	OnLeave   string // JavaScript onMouseLeave handler
	OnFocus   string // JavaScript onFocus handler
	OnBlur    string // JavaScript onBlur handler
}

// New creates a tooltip wrapper with provider
func New(props Props, trigger, content g.Node) g.Node {
	// Set defaults
	if props.Side == "" {
		props.Side = SideTop
	}
	if props.Align == "" {
		props.Align = AlignCenter
	}
	if props.DelayMs == 0 {
		props.DelayMs = 0
	}
	if props.SideOffset == 0 {
		props.SideOffset = 4
	}

	// Generate unique IDs if not provided
	tooltipID := props.ID
	if tooltipID == "" {
		tooltipID = fmt.Sprintf("tooltip-%d", hashString(props.Content))
	}
	triggerID := fmt.Sprintf("%s-trigger", tooltipID)
	contentID := fmt.Sprintf("%s-content", tooltipID)

	return html.Div(
		html.Class("relative inline-flex"),
		g.Attr("data-tooltip-container", "true"),
		
		// Trigger element
		trigger,
		
		// Tooltip content (hidden by default)
		html.Div(
			html.ID(contentID),
			g.Attr("role", "tooltip"),
			g.Attr("data-state", lib.CNIf(props.Open, "open", "closed")),
			g.Attr("data-side", string(props.Side)),
			g.Attr("data-align", string(props.Align)),
			html.Class(lib.CN(
				"tooltip-content",
				"absolute z-50 w-max rounded-md bg-primary px-3 py-1.5 text-xs text-primary-foreground",
				"pointer-events-none opacity-0 invisible",
				"transition-all duration-100",
				getPositionClasses(props.Side, props.Align),
				lib.CNIf(props.Open,
					"opacity-100 visible",
					"",
				),
				props.ContentClass,
				props.Class,
			)),
			html.Style(getPositionStyles(props)),
			
			// Content
			content,
			
			// Arrow
			html.Div(
				html.Class(lib.CN(
					"tooltip-arrow",
					"absolute h-2 w-2 rotate-45 bg-primary",
					getArrowClasses(props.Side),
					props.ArrowClass,
				)),
			),
		),
		
		// CSS for hover behavior
		html.Style(fmt.Sprintf(`
			#%s:hover + #%s,
			#%s:focus + #%s,
			#%s.tooltip-open + #%s {
				opacity: 1;
				visibility: visible;
				transition-delay: %dms;
			}
		`, triggerID, contentID, triggerID, contentID, triggerID, contentID, props.DelayMs)),
	)
}

// Trigger creates a tooltip trigger element
func Trigger(props TriggerProps, children ...g.Node) g.Node {
	attrs := []g.Node{
		g.Attr("data-tooltip-trigger", "true"),
		g.Attr("aria-describedby", props.ID + "-content"),
		html.Class(lib.CN(
			"tooltip-trigger",
			props.Class,
		)),
	}

	if props.ID != "" {
		attrs = append(attrs, html.ID(props.ID))
	}

	// Add event handlers
	if props.OnHover != "" {
		attrs = append(attrs, g.Attr("onmouseenter", props.OnHover))
	}
	if props.OnLeave != "" {
		attrs = append(attrs, g.Attr("onmouseleave", props.OnLeave))
	}
	if props.OnFocus != "" {
		attrs = append(attrs, g.Attr("onfocus", props.OnFocus))
	}
	if props.OnBlur != "" {
		attrs = append(attrs, g.Attr("onblur", props.OnBlur))
	}

	if props.AsChild && len(children) > 0 {
		// TODO: Implement AsChild pattern properly
		// For now, just return the children with attributes
		return g.Group(children)
	}

	return html.Span(
		append(attrs, children...)...,
	)
}

// TooltipContent creates tooltip content
func TooltipContent(props Props, children ...g.Node) g.Node {
	return g.Group(children)
}

// Simple creates a simple tooltip with text content
func Simple(content string, trigger g.Node) g.Node {
	return New(Props{
		Content: content,
	}, trigger, g.Text(content))
}

// WithSide creates a tooltip with specific side
func WithSide(content string, side Side, trigger g.Node) g.Node {
	return New(Props{
		Content: content,
		Side:    side,
	}, trigger, g.Text(content))
}

// WithDelay creates a tooltip with delay
func WithDelay(content string, delayMs int, trigger g.Node) g.Node {
	return New(Props{
		Content: content,
		DelayMs: delayMs,
	}, trigger, g.Text(content))
}

// Provider creates a tooltip provider wrapper (for consistency with React version)
func Provider(delayMs int, children ...g.Node) g.Node {
	return html.Div(
		g.Attr("data-tooltip-provider", "true"),
		g.Attr("data-delay", fmt.Sprintf("%d", delayMs)),
		g.Group(children),
	)
}

// Helper functions

func getPositionClasses(side Side, align Align) string {
	// Base positioning classes
	switch side {
	case SideTop:
		return "bottom-full mb-2"
	case SideBottom:
		return "top-full mt-2"
	case SideLeft:
		return "right-full mr-2"
	case SideRight:
		return "left-full ml-2"
	default:
		return "bottom-full mb-2"
	}
}

func getPositionStyles(props Props) string {
	alignOffset := props.AlignOffset
	
	var styles []string
	
	// Apply align offset
	switch props.Side {
	case SideTop, SideBottom:
		switch props.Align {
		case AlignStart:
			styles = append(styles, fmt.Sprintf("left: %dpx", alignOffset))
		case AlignEnd:
			styles = append(styles, fmt.Sprintf("right: %dpx", alignOffset))
		case AlignCenter:
			styles = append(styles, "left: 50%; transform: translateX(-50%)")
		}
	case SideLeft, SideRight:
		switch props.Align {
		case AlignStart:
			styles = append(styles, fmt.Sprintf("top: %dpx", alignOffset))
		case AlignEnd:
			styles = append(styles, fmt.Sprintf("bottom: %dpx", alignOffset))
		case AlignCenter:
			styles = append(styles, "top: 50%; transform: translateY(-50%)")
		}
	}
	
	return lib.CN(styles...)
}

func getArrowClasses(side Side) string {
	switch side {
	case SideTop:
		return "bottom-0 left-1/2 -translate-x-1/2 translate-y-1/2"
	case SideBottom:
		return "top-0 left-1/2 -translate-x-1/2 -translate-y-1/2"
	case SideLeft:
		return "right-0 top-1/2 translate-x-1/2 -translate-y-1/2"
	case SideRight:
		return "left-0 top-1/2 -translate-x-1/2 -translate-y-1/2"
	default:
		return "bottom-0 left-1/2 -translate-x-1/2 translate-y-1/2"
	}
}

// Simple hash function for generating IDs
func hashString(s string) uint32 {
	var h uint32
	for _, c := range s {
		h = h*31 + uint32(c)
	}
	return h
}

// Preset tooltip configurations

// InfoTooltip creates an info tooltip
func InfoTooltip(content string, trigger g.Node) g.Node {
	return New(Props{
		Content: content,
		ContentClass: "bg-blue-600 text-white",
		ArrowClass: "bg-blue-600",
	}, trigger, g.Group([]g.Node{
		g.Raw(`<svg xmlns="http://www.w3.org/2000/svg" class="inline-block w-3 h-3 mr-1" viewBox="0 0 20 20" fill="currentColor"><path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z" clip-rule="evenodd" /></svg>`),
		g.Text(content),
	}))
}

// WarningTooltip creates a warning tooltip
func WarningTooltip(content string, trigger g.Node) g.Node {
	return New(Props{
		Content: content,
		ContentClass: "bg-yellow-600 text-white",
		ArrowClass: "bg-yellow-600",
	}, trigger, g.Group([]g.Node{
		g.Raw(`<svg xmlns="http://www.w3.org/2000/svg" class="inline-block w-3 h-3 mr-1" viewBox="0 0 20 20" fill="currentColor"><path fill-rule="evenodd" d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z" clip-rule="evenodd" /></svg>`),
		g.Text(content),
	}))
}

// ErrorTooltip creates an error tooltip
func ErrorTooltip(content string, trigger g.Node) g.Node {
	return New(Props{
		Content: content,
		ContentClass: "bg-red-600 text-white",
		ArrowClass: "bg-red-600",
	}, trigger, g.Group([]g.Node{
		g.Raw(`<svg xmlns="http://www.w3.org/2000/svg" class="inline-block w-3 h-3 mr-1" viewBox="0 0 20 20" fill="currentColor"><path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clip-rule="evenodd" /></svg>`),
		g.Text(content),
	}))
}

// IconTooltip creates a tooltip with custom icon
func IconTooltip(content string, icon g.Node, trigger g.Node) g.Node {
	return New(Props{
		Content: content,
	}, trigger, g.Group([]g.Node{
		icon,
		g.Text(" "),
		g.Text(content),
	}))
}