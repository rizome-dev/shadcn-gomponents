package sidebar

import (
	"fmt"
	"net/http"

	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	hx "maragu.dev/gomponents-htmx"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// HTMXProps defines properties for HTMX-enhanced sidebar
type HTMXProps struct {
	ID              string
	TogglePath      string
	StatePath       string
	MobileTogglePath string
	DefaultOpen     bool
}

// HTMXProvider creates an HTMX-enhanced sidebar provider
func HTMXProvider(props ProviderProps, htmxProps HTMXProps, children ...g.Node) g.Node {
	style := fmt.Sprintf("--sidebar-width: %s; --sidebar-width-icon: %s;", SidebarWidth, SidebarWidthIcon)
	if props.Style != "" {
		style = style + " " + props.Style
	}

	classes := lib.CN(
		"group/sidebar-wrapper has-[[data-variant=inset]]:bg-sidebar flex min-h-svh w-full",
		props.Class,
	)

	return html.Div(
		g.Attr("id", htmxProps.ID),
		g.Attr("data-sidebar-wrapper", "true"),
		g.Attr("style", style),
		html.Class(classes),
		hx.Get(htmxProps.StatePath),
		hx.Trigger("load"),
		hx.Target("this"),
		hx.Swap("outerHTML"),
		g.Group(children),
	)
}

// HTMXSidebar creates an HTMX-enhanced sidebar
func HTMXSidebar(props Props, htmxProps HTMXProps, isOpen bool, children ...g.Node) g.Node {
	// Set defaults
	if props.Side == "" {
		props.Side = "left"
	}
	if props.Variant == "" {
		props.Variant = "sidebar"
	}
	if props.Collapsible == "" {
		props.Collapsible = "offcanvas"
	}

	state := "collapsed"
	if isOpen {
		state = "expanded"
	}

	// For non-collapsible sidebar
	if props.Collapsible == "none" {
		return html.Div(
			g.Attr("data-sidebar", "true"),
			html.Class(lib.CN(
				"bg-sidebar text-sidebar-foreground flex h-full w-[var(--sidebar-width)] flex-col",
				props.Class,
			)),
			g.Group(children),
		)
	}

	// Main sidebar container
	return html.Div(
		html.Class("group peer text-sidebar-foreground hidden md:block"),
		g.Attr("data-sidebar", "true"),
		g.Attr("data-state", state),
		g.Attr("data-variant", props.Variant),
		g.Attr("data-side", props.Side),
		g.Attr("data-collapsible", lib.CNIf(state == "collapsed", props.Collapsible, "")),
		g.If(htmxProps.ID != "", g.Attr("id", htmxProps.ID)),
		
		// Sidebar gap
		html.Div(
			g.Attr("data-sidebar-gap", "true"),
			html.Class(lib.CN(
				"relative w-[var(--sidebar-width)] bg-transparent transition-[width] duration-200 ease-linear",
				"group-data-[collapsible=offcanvas]:w-0",
				"group-data-[side=right]:rotate-180",
				lib.CNIf(props.Variant == "floating" || props.Variant == "inset",
					"group-data-[collapsible=icon]:w-[calc(var(--sidebar-width-icon)+theme(spacing.4))]",
					"group-data-[collapsible=icon]:w-[var(--sidebar-width-icon)]",
				),
			)),
		),
		
		// Sidebar container
		html.Div(
			g.Attr("data-sidebar-container", "true"),
			html.Class(lib.CN(
				"fixed inset-y-0 z-10 hidden h-svh w-[var(--sidebar-width)] transition-[left,right,width] duration-200 ease-linear md:flex",
				lib.CNIf(props.Side == "left",
					"left-0 group-data-[collapsible=offcanvas]:left-[calc(var(--sidebar-width)*-1)]",
					"right-0 group-data-[collapsible=offcanvas]:right-[calc(var(--sidebar-width)*-1)]",
				),
				lib.CNIf(props.Variant == "floating" || props.Variant == "inset",
					"p-2 group-data-[collapsible=icon]:w-[calc(var(--sidebar-width-icon)+theme(spacing.4)+2px)]",
					"group-data-[collapsible=icon]:w-[var(--sidebar-width-icon)] group-data-[side=left]:border-r group-data-[side=right]:border-l",
				),
				props.Class,
			)),
			
			// Inner sidebar
			html.Div(
				g.Attr("data-sidebar-inner", "true"),
				html.Class("bg-sidebar group-data-[variant=floating]:border-sidebar-border flex h-full w-full flex-col group-data-[variant=floating]:rounded-lg group-data-[variant=floating]:border group-data-[variant=floating]:shadow-sm"),
				g.Group(children),
			),
		),
	)
}

// HTMXTrigger creates an HTMX-enhanced sidebar trigger
func HTMXTrigger(props Props, htmxProps HTMXProps, children ...g.Node) g.Node {
	buttonChildren := []g.Node{
		// Default icon if no children provided
		g.Raw(`<svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg" class="size-4"><path d="M2 4.5C2 4.22386 2.22386 4 2.5 4H12.5C12.7761 4 13 4.22386 13 4.5C13 4.77614 12.7761 5 12.5 5H2.5C2.22386 5 2 4.77614 2 4.5ZM2 7.5C2 7.22386 2.22386 7 2.5 7H12.5C12.7761 7 13 7.22386 13 7.5C13 7.77614 12.7761 8 12.5 8H2.5C2.22386 8 2 7.77614 2 7.5ZM2 10.5C2 10.2239 2.22386 10 2.5 10H12.5C12.7761 10 13 10.2239 13 10.5C13 10.7761 12.7761 11 12.5 11H2.5C2.22386 11 2 10.7761 2 10.5Z" fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"></path></svg>`),
		html.Span(html.Class("sr-only"), g.Text("Toggle Sidebar")),
	}
	
	if len(children) > 0 {
		buttonChildren = children
	}

	return html.Button(
		g.Attr("data-sidebar-trigger", "true"),
		g.Attr("type", "button"),
		html.Class(lib.CN("size-7", props.Class)),
		hx.Post(htmxProps.TogglePath),
		hx.Target("#" + htmxProps.ID),
		hx.Swap("outerHTML"),
		g.Group(buttonChildren),
	)
}

// HTMXRail creates an HTMX-enhanced sidebar rail
func HTMXRail(props Props, htmxProps HTMXProps) g.Node {
	return html.Button(
		g.Attr("data-sidebar-rail", "true"),
		g.Attr("aria-label", "Toggle Sidebar"),
		g.Attr("tabindex", "-1"),
		g.Attr("title", "Toggle Sidebar"),
		html.Class(lib.CN(
			"hover:after:bg-sidebar-border absolute inset-y-0 z-20 hidden w-4 -translate-x-1/2 transition-all ease-linear group-data-[side=left]:-right-4 group-data-[side=right]:left-0 after:absolute after:inset-y-0 after:left-1/2 after:w-[2px] sm:flex",
			"[[data-side=left]_&]:cursor-w-resize [[data-side=right]_&]:cursor-e-resize",
			"[[data-side=left][data-state=collapsed]_&]:cursor-e-resize [[data-side=right][data-state=collapsed]_&]:cursor-w-resize",
			"hover:group-data-[collapsible=offcanvas]:bg-sidebar group-data-[collapsible=offcanvas]:translate-x-0 group-data-[collapsible=offcanvas]:after:left-full",
			"[[data-side=left][data-collapsible=offcanvas]_&]:-right-2",
			"[[data-side=right][data-collapsible=offcanvas]_&]:-left-2",
			props.Class,
		)),
		hx.Post(htmxProps.TogglePath),
		hx.Target("#" + htmxProps.ID),
		hx.Swap("outerHTML"),
	)
}

// MobileSheet creates a mobile sidebar sheet overlay
func MobileSheet(props Props, htmxProps HTMXProps, isOpen bool, children ...g.Node) g.Node {
	if !isOpen {
		return nil
	}

	return html.Div(
		html.Class("fixed inset-0 z-50 md:hidden"),
		// Backdrop
		html.Div(
			html.Class("fixed inset-0 bg-background/80 backdrop-blur-sm"),
			hx.Post(htmxProps.MobileTogglePath + "?close=true"),
			hx.Target("#" + htmxProps.ID + "-mobile"),
			hx.Swap("outerHTML"),
		),
		// Sheet content
		html.Div(
			g.Attr("data-sidebar", "true"),
			g.Attr("data-mobile", "true"),
			html.Class(lib.CN(
				"bg-sidebar text-sidebar-foreground fixed inset-y-0 w-[var(--sidebar-width)] p-0",
				lib.CNIf(props.Side == "left", "left-0", "right-0"),
			)),
			g.Attr("style", fmt.Sprintf("--sidebar-width: %s;", SidebarWidthMobile)),
			// Close button
			html.Button(
				html.Type("button"),
				html.Class("absolute right-4 top-4 rounded-sm opacity-70 ring-offset-background transition-opacity hover:opacity-100 focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 disabled:pointer-events-none data-[state=open]:bg-secondary"),
				hx.Post(htmxProps.MobileTogglePath + "?close=true"),
				hx.Target("#" + htmxProps.ID + "-mobile"),
				hx.Swap("outerHTML"),
				g.Raw(`<svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg" class="h-4 w-4"><path d="M11.7816 4.03157C12.0062 3.80702 12.0062 3.44295 11.7816 3.2184C11.5571 2.99385 11.193 2.99385 10.9685 3.2184L7.50005 6.68682L4.03164 3.2184C3.80708 2.99385 3.44301 2.99385 3.21846 3.2184C2.99391 3.44295 2.99391 3.80702 3.21846 4.03157L6.68688 7.49999L3.21846 10.9684C2.99391 11.193 2.99391 11.557 3.21846 11.7816C3.44301 12.0061 3.80708 12.0061 4.03164 11.7816L7.50005 8.31316L10.9685 11.7816C11.193 12.0061 11.5571 12.0061 11.7816 11.7816C12.0062 11.557 12.0062 11.193 11.7816 10.9684L8.31322 7.49999L11.7816 4.03157Z" fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"></path></svg>`),
				html.Span(html.Class("sr-only"), g.Text("Close")),
			),
			// Content
			html.Div(
				html.Class("flex h-full w-full flex-col"),
				g.Group(children),
			),
		),
	)
}

// Server-side state management
var sidebarStates = make(map[string]bool)

// SidebarHandlers creates HTTP handlers for sidebar functionality
func SidebarHandlers(mux *http.ServeMux, baseProps Props, htmxProps HTMXProps) {
	// Initialize state
	if _, exists := sidebarStates[htmxProps.ID]; !exists {
		sidebarStates[htmxProps.ID] = htmxProps.DefaultOpen
	}

	// Toggle handler
	mux.HandleFunc(htmxProps.TogglePath, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Toggle state
		sidebarStates[htmxProps.ID] = !sidebarStates[htmxProps.ID]

		// Return updated sidebar
		sidebar := HTMXSidebar(baseProps, htmxProps, sidebarStates[htmxProps.ID], 
			// You would pass the actual content here
			HeaderComponent(Props{}, g.Text("Sidebar Header")),
			ContentComponent(Props{}, g.Text("Sidebar Content")),
			FooterComponent(Props{}, g.Text("Sidebar Footer")),
		)
		
		sidebar.Render(w)
	})

	// State handler (for initial load)
	mux.HandleFunc(htmxProps.StatePath, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Return current state sidebar
		sidebar := HTMXSidebar(baseProps, htmxProps, sidebarStates[htmxProps.ID],
			// You would pass the actual content here
			HeaderComponent(Props{}, g.Text("Sidebar Header")),
			ContentComponent(Props{}, g.Text("Sidebar Content")),
			FooterComponent(Props{}, g.Text("Sidebar Footer")),
		)
		
		sidebar.Render(w)
	})

	// Mobile toggle handler
	if htmxProps.MobileTogglePath != "" {
		mobileID := htmxProps.ID + "-mobile"
		mux.HandleFunc(htmxProps.MobileTogglePath, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}

			close := r.URL.Query().Get("close") == "true"
			
			if close {
				// Return empty div to close
				html.Div(g.Attr("id", mobileID)).Render(w)
			} else {
				// Toggle mobile state
				isOpen := !close
				if _, exists := sidebarStates[mobileID]; exists {
					isOpen = !sidebarStates[mobileID]
				}
				sidebarStates[mobileID] = isOpen

				// Return mobile sheet
				sheet := html.Div(
					g.Attr("id", mobileID),
					g.If(isOpen, MobileSheet(baseProps, htmxProps, isOpen,
						// You would pass the actual content here
						HeaderComponent(Props{}, g.Text("Sidebar Header")),
						ContentComponent(Props{}, g.Text("Sidebar Content")),
						FooterComponent(Props{}, g.Text("Sidebar Footer")),
					)),
				)
				
				sheet.Render(w)
			}
		})
	}
}