package popover

import (
	"fmt"
	"net/http"
	"strings"
	
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	hx "maragu.dev/gomponents-htmx"
	"github.com/rizome-dev/shadcn-gomponents/lib"
	"github.com/rizome-dev/shadcn-gomponents/lib/icons"
)

// HTMXProps defines HTMX-specific properties for the Popover
type HTMXProps struct {
	ID         string // Unique ID for the popover
	TogglePath string // Server path for toggle actions
	ClosePath  string // Server path for close actions
	LoadPath   string // Server path for loading dynamic content
}

// NewHTMX creates an HTMX-enhanced Popover component
func NewHTMX(props Props, htmxProps HTMXProps, children ...g.Node) g.Node {
	classes := lib.CN("relative inline-block", props.Class)
	
	return html.Div(
		html.ID(htmxProps.ID + "-container"),
		html.Class(classes),
		g.If(props.Open, g.Attr("data-state", "open")),
		g.If(!props.Open, g.Attr("data-state", "closed")),
		g.Group(children),
	)
}

// TriggerHTMX creates an HTMX-enhanced trigger
func TriggerHTMX(props TriggerProps, htmxProps HTMXProps, children ...g.Node) g.Node {
	classes := lib.CN(props.Class)
	
	triggerAttrs := []g.Node{
		html.Type("button"),
		g.If(classes != "", html.Class(classes)),
		g.Attr("aria-haspopup", "dialog"),
		g.Attr("aria-expanded", "false"),
		hx.Get(htmxProps.TogglePath),
		hx.Target("#" + htmxProps.ID + "-container"),
		hx.Swap("outerHTML"),
	}
	
	if props.AsChild && len(children) > 0 {
		// For AsChild, we need to add HTMX attributes to the first child
		return children[0]
	}
	
	return html.Button(append(triggerAttrs, children...)...)
}

// ContentHTMX creates HTMX-enhanced popover content
func ContentHTMX(props ContentProps, htmxProps HTMXProps, children ...g.Node) g.Node {
	// Default values
	if props.Side == "" {
		props.Side = "bottom"
	}
	if props.Align == "" {
		props.Align = "center"
	}
	
	// Build position classes based on side and align
	positionClasses := getPositionClasses(props.Side, props.Align)
	
	classes := lib.CN(
		"z-50 w-72 rounded-md border bg-popover p-4 text-popover-foreground shadow-md outline-none",
		"animate-in fade-in-0 zoom-in-95",
		getAnimationClasses(props.Side),
		positionClasses,
		props.Class,
	)
	
	// Add click outside handler
	return html.Div(
		html.ID(htmxProps.ID),
		html.Class(classes),
		html.Role("dialog"),
		g.Attr("data-state", "open"),
		g.If(props.Side != "", g.Attr("data-side", props.Side)),
		g.If(props.Align != "", g.Attr("data-align", props.Align)),
		// Click outside to close
		hx.On("click", fmt.Sprintf(`
			if (event.target.id === '%s' || event.target.closest('#%s')) {
				return;
			}
			htmx.ajax('GET', '%s', {target: '#%s-container', swap: 'outerHTML'});
		`, htmxProps.ID, htmxProps.ID, htmxProps.ClosePath, htmxProps.ID)),
		g.Group(children),
	)
}

// CloseHTMX creates an HTMX-enhanced close button
func CloseHTMX(htmxProps HTMXProps, class ...string) g.Node {
	classes := lib.CN(
		"absolute right-2 top-2 rounded-sm opacity-70 ring-offset-background transition-opacity hover:opacity-100 focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2",
		lib.CN(class...),
	)
	
	return html.Button(
		html.Type("button"),
		html.Class(classes),
		g.Attr("aria-label", "Close"),
		hx.Get(htmxProps.ClosePath),
		hx.Target("#" + htmxProps.ID + "-container"),
		hx.Swap("outerHTML"),
		icons.X(html.Class("h-4 w-4")),
	)
}

// DynamicContentHTMX creates popover content that loads dynamically
func DynamicContentHTMX(props ContentProps, htmxProps HTMXProps) g.Node {
	return ContentHTMX(
		props,
		htmxProps,
		html.Div(
			html.Class("relative"),
			hx.Get(htmxProps.LoadPath),
			hx.Trigger("load"),
			hx.Swap("innerHTML"),
			// Loading spinner
			html.Div(
				html.Class("flex items-center justify-center py-8"),
				icons.Loader(html.Class("h-6 w-6 animate-spin")),
			),
		),
	)
}

// RenderOpenPopover renders an open popover (for server response)
func RenderOpenPopover(props Props, contentProps ContentProps, htmxProps HTMXProps, trigger, content g.Node) g.Node {
	return NewHTMX(
		Props{Open: true},
		htmxProps,
		trigger,
		ContentHTMX(contentProps, htmxProps, content),
	)
}

// RenderClosedPopover renders a closed popover (for server response)
func RenderClosedPopover(props Props, htmxProps HTMXProps, trigger g.Node) g.Node {
	return NewHTMX(
		Props{Open: false},
		htmxProps,
		trigger,
	)
}

// MenuPopoverHTMX creates a popover with menu items
func MenuPopoverHTMX(htmxProps HTMXProps, items []MenuItem) g.Node {
	return ContentHTMX(
		ContentProps{Class: "w-56 p-1"},
		htmxProps,
		g.Group(g.Map(items, func(item MenuItem) g.Node {
			if item.Separator {
				return html.Div(html.Class("h-px bg-border my-1"))
			}
			
			btnClasses := lib.CN(
				"flex w-full items-center rounded-sm px-2 py-1.5 text-sm hover:bg-accent hover:text-accent-foreground",
				lib.CNIf(item.Disabled, "opacity-50 cursor-not-allowed", "cursor-pointer"),
			)
			
			btn := html.Button(
				html.Type("button"),
				html.Class(btnClasses),
				g.If(item.Disabled, html.Disabled()),
			)
			
			if item.Action != "" && !item.Disabled {
				btn = html.Button(
					html.Type("button"),
					html.Class(btnClasses),
					hx.Post(item.Action),
					hx.Target("#" + htmxProps.ID + "-container"),
					hx.Swap("outerHTML"),
				)
			}
			
			return g.El("div", btn,
				g.If(item.Icon != nil, item.Icon),
				html.Span(g.If(item.Icon != nil, html.Class("ml-2")), g.Text(item.Label)),
				g.If(item.Shortcut != "", 
					html.Span(html.Class("ml-auto text-xs text-muted-foreground"), g.Text(item.Shortcut)),
				),
			)
		})),
	)
}

// MenuItem represents a menu item in the popover
type MenuItem struct {
	Label     string
	Icon      g.Node
	Action    string // HTMX action endpoint
	Shortcut  string
	Disabled  bool
	Separator bool
}

// ProfilePopoverHTMX creates a profile popover with user info
func ProfilePopoverHTMX(htmxProps HTMXProps, userName, userEmail string, avatarUrl string) g.Node {
	return ContentHTMX(
		ContentProps{Class: "w-80"},
		htmxProps,
		html.Div(html.Class("flex gap-4"),
			// Avatar
			html.Div(html.Class("h-12 w-12 rounded-full bg-slate-200 overflow-hidden"),
				g.If(avatarUrl != "",
					html.Img(html.Src(avatarUrl), html.Alt(userName), html.Class("h-full w-full object-cover")),
				),
				g.If(avatarUrl == "",
					html.Div(html.Class("h-full w-full flex items-center justify-center text-lg font-medium"),
						g.Text(strings.ToUpper(string(userName[0]))),
					),
				),
			),
			// User info
			html.Div(html.Class("flex-1 space-y-1"),
				html.P(html.Class("text-sm font-medium leading-none"), g.Text(userName)),
				html.P(html.Class("text-sm text-muted-foreground"), g.Text(userEmail)),
			),
		),
		html.Div(html.Class("mt-4 pt-4 border-t"),
			MenuPopoverHTMX(htmxProps, []MenuItem{
				{Label: "Profile", Icon: icons.User(html.Class("mr-2 h-4 w-4")), Action: "/profile"},
				{Label: "Settings", Icon: icons.Settings(html.Class("mr-2 h-4 w-4")), Action: "/settings"},
				{Separator: true},
				{Label: "Logout", Icon: icons.LogOut(html.Class("mr-2 h-4 w-4")), Action: "/logout"},
			}),
		),
	)
}

// PopoverHandlers creates HTTP handlers for popover components
func PopoverHandlers(mux *http.ServeMux) {
	// Basic popover example
	basicProps := HTMXProps{
		ID:         "popover-basic",
		TogglePath: "/api/popover/basic/toggle",
		ClosePath:  "/api/popover/basic/close",
	}
	
	mux.HandleFunc("/api/popover/basic/toggle", func(w http.ResponseWriter, r *http.Request) {
		// Toggle popover state (in real app, track state server-side)
		trigger := TriggerHTMX(
			TriggerProps{},
			basicProps,
			g.Text("Open popover"),
		)
		
		content := html.Div(
			html.H3(html.Class("font-medium"), g.Text("Popover Content")),
			html.P(html.Class("text-sm text-muted-foreground mt-2"), 
				g.Text("This is a basic popover with some content."),
			),
		)
		
		node := RenderOpenPopover(
			Props{},
			ContentProps{},
			basicProps,
			trigger,
			content,
		)
		node.Render(w)
	})
	
	mux.HandleFunc("/api/popover/basic/close", func(w http.ResponseWriter, r *http.Request) {
		trigger := TriggerHTMX(
			TriggerProps{},
			basicProps,
			g.Text("Open popover"),
		)
		
		node := RenderClosedPopover(Props{}, basicProps, trigger)
		node.Render(w)
	})
	
	// Menu popover example
	menuProps := HTMXProps{
		ID:         "popover-menu",
		TogglePath: "/api/popover/menu/toggle",
		ClosePath:  "/api/popover/menu/close",
	}
	
	mux.HandleFunc("/api/popover/menu/toggle", func(w http.ResponseWriter, r *http.Request) {
		trigger := TriggerHTMX(
			TriggerProps{Class: "inline-flex items-center gap-2"},
			menuProps,
			icons.MoreVertical(html.Class("h-4 w-4")),
			g.Text("Options"),
		)
		
		node := RenderOpenPopover(
			Props{},
			ContentProps{Side: "bottom", Align: "end"},
			menuProps,
			trigger,
			MenuPopoverHTMX(menuProps, []MenuItem{
				{Label: "Edit", Icon: icons.Edit(html.Class("mr-2 h-4 w-4")), Action: "/api/item/edit"},
				{Label: "Duplicate", Icon: icons.Copy(html.Class("mr-2 h-4 w-4")), Action: "/api/item/duplicate"},
				{Separator: true},
				{Label: "Archive", Icon: icons.Archive(html.Class("mr-2 h-4 w-4")), Action: "/api/item/archive"},
				{Label: "Delete", Icon: icons.Trash(html.Class("mr-2 h-4 w-4")), Action: "/api/item/delete"},
			}),
		)
		node.Render(w)
	})
	
	mux.HandleFunc("/api/popover/menu/close", func(w http.ResponseWriter, r *http.Request) {
		trigger := TriggerHTMX(
			TriggerProps{Class: "inline-flex items-center gap-2"},
			menuProps,
			icons.MoreVertical(html.Class("h-4 w-4")),
			g.Text("Options"),
		)
		
		node := RenderClosedPopover(Props{}, menuProps, trigger)
		node.Render(w)
	})
	
	// Handle menu actions
	mux.HandleFunc("/api/item/", func(w http.ResponseWriter, r *http.Request) {
		action := strings.TrimPrefix(r.URL.Path, "/api/item/")
		
		// Return a success message
		trigger := TriggerHTMX(
			TriggerProps{Class: "inline-flex items-center gap-2"},
			menuProps,
			icons.MoreVertical(html.Class("h-4 w-4")),
			g.Text("Options"),
		)
		
		// Close popover with success message
		node := html.Div(
			html.ID(menuProps.ID + "-container"),
			html.Class("relative inline-block"),
			trigger,
			html.Div(
				html.Class("absolute top-full mt-2 left-0 bg-green-500 text-white px-3 py-1 rounded-md text-sm"),
				g.Textf("Action '%s' completed!", action),
				// Auto-hide after 2 seconds
				g.Attr("x-data", "{}"),
				g.Attr("x-init", "setTimeout(() => $el.remove(), 2000)"),
			),
		)
		node.Render(w)
	})
	
	// Dynamic content popover
	dynamicProps := HTMXProps{
		ID:         "popover-dynamic",
		TogglePath: "/api/popover/dynamic/toggle",
		ClosePath:  "/api/popover/dynamic/close",
		LoadPath:   "/api/popover/dynamic/content",
	}
	
	mux.HandleFunc("/api/popover/dynamic/toggle", func(w http.ResponseWriter, r *http.Request) {
		trigger := TriggerHTMX(
			TriggerProps{},
			dynamicProps,
			g.Text("Load dynamic content"),
		)
		
		node := RenderOpenPopover(
			Props{},
			ContentProps{Class: "w-80"},
			dynamicProps,
			trigger,
			DynamicContentHTMX(ContentProps{}, dynamicProps),
		)
		node.Render(w)
	})
	
	mux.HandleFunc("/api/popover/dynamic/content", func(w http.ResponseWriter, r *http.Request) {
		// Simulate loading delay
		// time.Sleep(500 * time.Millisecond)
		
		node := html.Div(
			html.H3(html.Class("font-medium mb-2"), g.Text("Dynamic Content")),
			html.P(html.Class("text-sm text-muted-foreground"), 
				g.Text("This content was loaded dynamically from the server."),
			),
			html.Div(html.Class("mt-4 space-y-2"),
				html.Div(html.Class("flex justify-between text-sm"),
					html.Span(g.Text("Status:")),
					html.Span(html.Class("font-medium text-green-600"), g.Text("Active")),
				),
				html.Div(html.Class("flex justify-between text-sm"),
					html.Span(g.Text("Last updated:")),
					html.Span(html.Class("font-medium"), g.Text("2 minutes ago")),
				),
			),
		)
		node.Render(w)
	})
}