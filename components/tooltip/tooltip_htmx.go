package tooltip

import (
	"fmt"
	"net/http"

	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	hx "maragu.dev/gomponents-htmx"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// HTMXProps defines properties for HTMX-enhanced tooltips
type HTMXProps struct {
	ID          string // Unique identifier for the tooltip
	ShowPath    string // Path to load tooltip content on show
	HidePath    string // Path to notify when tooltip is hidden
	LoadOnHover bool   // Whether to load content on hover
}

// HTMXTooltip creates an HTMX-enhanced tooltip
func HTMXTooltip(props Props, htmxProps HTMXProps, trigger, content g.Node) g.Node {
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

	// Generate unique IDs
	tooltipID := htmxProps.ID
	if tooltipID == "" {
		tooltipID = fmt.Sprintf("htmx-tooltip-%d", hashString(props.Content))
	}
	contentID := fmt.Sprintf("%s-content", tooltipID)

	triggerAttrs := []g.Node{
		html.ID(fmt.Sprintf("%s-trigger", tooltipID)),
		g.Attr("data-tooltip-trigger", "true"),
		g.Attr("aria-describedby", contentID),
	}

	// Add HTMX attributes for dynamic loading
	if htmxProps.LoadOnHover && htmxProps.ShowPath != "" {
		triggerAttrs = append(triggerAttrs,
			hx.Get(htmxProps.ShowPath),
			hx.Trigger(fmt.Sprintf("mouseenter delay:%dms", props.DelayMs)),
			hx.Target("#"+contentID),
			hx.Swap("innerHTML"),
		)
	}

	return html.Div(
		html.Class("relative inline-flex"),
		g.Attr("data-htmx-tooltip", "true"),
		
		// Enhanced trigger
		html.Span(
			append(triggerAttrs, trigger)...,
		),
		
		// Tooltip content container
		html.Div(
			html.ID(contentID),
			g.Attr("role", "tooltip"),
			g.Attr("data-state", lib.CNIf(props.Open, "open", "closed")),
			g.Attr("data-side", string(props.Side)),
			g.Attr("data-align", string(props.Align)),
			html.Class(lib.CN(
				"htmx-tooltip-content",
				"absolute z-50 w-max rounded-md bg-primary px-3 py-1.5 text-xs text-primary-foreground",
				"pointer-events-none opacity-0 invisible",
				"transition-all duration-100",
				getPositionClasses(props.Side, props.Align),
				lib.CNIf(props.Open, "opacity-100 visible", ""),
				props.ContentClass,
				props.Class,
			)),
			html.Style(getPositionStyles(props)),
			
			// Content (will be replaced by HTMX if LoadOnHover is true)
			g.If(htmxProps.LoadOnHover,
				html.Span(html.Class("text-muted-foreground"), g.Text("Loading...")),
			),
			g.If(!htmxProps.LoadOnHover,
				content,
			),
			
			// Arrow
			html.Div(
				html.Class(lib.CN(
					"absolute h-2 w-2 rotate-45 bg-primary",
					getArrowClasses(props.Side),
					props.ArrowClass,
				)),
			),
		),
		
		// Enhanced CSS for HTMX
		html.Style(fmt.Sprintf(`
			#%s-trigger:hover + #%s,
			#%s-trigger:focus + #%s,
			#%s-trigger[data-tooltip-open="true"] + #%s {
				opacity: 1;
				visibility: visible;
				transition-delay: %dms;
			}
			
			#%s-trigger:not(:hover):not(:focus) + #%s {
				transition-delay: 0ms;
			}
		`, tooltipID, contentID, tooltipID, contentID, tooltipID, contentID, 
		   props.DelayMs, tooltipID, contentID)),
	)
}

// HTMXTrigger creates an HTMX-enhanced tooltip trigger
func HTMXTrigger(htmxProps HTMXProps, children ...g.Node) g.Node {
	attrs := []g.Node{
		g.Attr("data-htmx-tooltip-trigger", "true"),
		html.Class("cursor-pointer"),
	}

	// Add show/hide HTMX handlers
	if htmxProps.ShowPath != "" {
		attrs = append(attrs,
			g.Attr("hx-get", htmxProps.ShowPath),
			g.Attr("hx-trigger", "mouseenter"),
			g.Attr("hx-target", "next .htmx-tooltip-content"),
		)
	}

	if htmxProps.HidePath != "" {
		attrs = append(attrs,
			g.Attr("hx-post", htmxProps.HidePath),
			g.Attr("hx-trigger", "mouseleave"),
		)
	}

	return html.Span(
		append(attrs, children...)...,
	)
}

// TooltipHandlers creates HTTP handlers for tooltip functionality
func TooltipHandlers(mux *http.ServeMux, baseProps Props, htmxProps HTMXProps) {
	// Handler to show tooltip content
	if htmxProps.ShowPath != "" {
		mux.HandleFunc(htmxProps.ShowPath, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}

			// You can customize this to load dynamic content
			// For now, return the static content
			content := baseProps.Content
			if content == "" {
				content = "Tooltip content"
			}

			// Write the content
			g.Text(content).Render(w)
		})
	}

	// Handler for hide events (if needed for tracking)
	if htmxProps.HidePath != "" {
		mux.HandleFunc(htmxProps.HidePath, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}

			// Could log tooltip hide events or update state
			w.WriteHeader(http.StatusOK)
		})
	}
}

// HTMXDynamicTooltip creates a tooltip that loads content dynamically
func HTMXDynamicTooltip(trigger g.Node, loadPath string) g.Node {
	return HTMXTooltip(
		Props{
			DelayMs: 200,
		},
		HTMXProps{
			ShowPath:    loadPath,
			LoadOnHover: true,
		},
		trigger,
		nil, // Content will be loaded dynamically
	)
}

// HTMXFormTooltip creates a tooltip for form field help
func HTMXFormTooltip(fieldName, helpPath string) g.Node {
	trigger := g.Raw(`<svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 text-muted-foreground hover:text-foreground transition-colors" viewBox="0 0 20 20" fill="currentColor">
		<path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-8-3a1 1 0 00-.867.5 1 1 0 11-1.731-1A3 3 0 0113 8a3.001 3.001 0 01-2 2.83V11a1 1 0 11-2 0v-1a1 1 0 011-1 1 1 0 100-2zm0 8a1 1 0 100-2 1 1 0 000 2z" clip-rule="evenodd" />
	</svg>`)

	return HTMXTooltip(
		Props{
			Side:    SideRight,
			DelayMs: 100,
		},
		HTMXProps{
			ID:          fmt.Sprintf("help-%s", fieldName),
			ShowPath:    fmt.Sprintf("%s?field=%s", helpPath, fieldName),
			LoadOnHover: true,
		},
		trigger,
		nil,
	)
}

// HTMXUserTooltip creates a tooltip that shows user information
func HTMXUserTooltip(userID string, trigger g.Node) g.Node {
	return HTMXTooltip(
		Props{
			Side:         SideBottom,
			DelayMs:      300,
			ContentClass: "p-4 max-w-xs",
		},
		HTMXProps{
			ID:          fmt.Sprintf("user-%s", userID),
			ShowPath:    fmt.Sprintf("/api/users/%s/tooltip", userID),
			LoadOnHover: true,
		},
		trigger,
		nil,
	)
}

// HTMXStatusTooltip creates a tooltip that shows live status
func HTMXStatusTooltip(statusID string, trigger g.Node) g.Node {
	return HTMXTooltip(
		Props{
			Side:    SideTop,
			DelayMs: 0,
		},
		HTMXProps{
			ID:          fmt.Sprintf("status-%s", statusID),
			ShowPath:    fmt.Sprintf("/api/status/%s", statusID),
			LoadOnHover: true,
		},
		trigger,
		html.Div(
			html.Class("flex items-center gap-2"),
			html.Span(html.Class("inline-block w-2 h-2 bg-green-500 rounded-full animate-pulse")),
			html.Span(g.Text("Loading status...")),
		),
	)
}

// HTMXProgressTooltip creates a tooltip showing progress
func HTMXProgressTooltip(taskID string, trigger g.Node) g.Node {
	return HTMXTooltip(
		Props{
			Side:         SideBottom,
			DelayMs:      0,
			ContentClass: "p-2",
		},
		HTMXProps{
			ID:          fmt.Sprintf("progress-%s", taskID),
			ShowPath:    fmt.Sprintf("/api/tasks/%s/progress", taskID),
			LoadOnHover: true,
		},
		trigger,
		html.Div(
			html.Class("w-32"),
			html.Div(html.Class("text-xs mb-1"), g.Text("Loading...")),
			html.Div(
				html.Class("w-full bg-secondary rounded-full h-2"),
				html.Div(html.Class("bg-primary h-2 rounded-full transition-all"), html.Style("width: 0%")),
			),
		),
	)
}

// BatchTooltipHandler creates a handler for multiple tooltips with shared logic
func BatchTooltipHandler(mux *http.ServeMux, basePath string, contentFunc func(string) g.Node) {
	mux.HandleFunc(basePath, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Extract tooltip ID from query params or path
		tooltipID := r.URL.Query().Get("id")
		if tooltipID == "" {
			http.Error(w, "Missing tooltip ID", http.StatusBadRequest)
			return
		}

		// Generate content based on ID
		content := contentFunc(tooltipID)
		if content == nil {
			http.Error(w, "Tooltip not found", http.StatusNotFound)
			return
		}

		// Render the content
		content.Render(w)
	})
}