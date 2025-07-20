package drawer

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

// HTMXProps defines HTMX-specific properties for the Drawer
type HTMXProps struct {
	ID          string // Unique ID for the drawer
	TriggerPath string // Server path for trigger actions
	ClosePath   string // Server path for close actions
	ContentPath string // Server path for loading dynamic content
}

// NewHTMX creates an HTMX-enhanced Drawer component
func NewHTMX(props Props, htmxProps HTMXProps, children ...g.Node) g.Node {
	if props.Open {
		classes := lib.CN("fixed inset-0 z-50", props.Class)
		return html.Div(
			html.ID(htmxProps.ID),
			html.Class(classes),
			g.Attr("data-state", "open"),
			g.Group(children),
		)
	}
	// Return empty div that can be replaced by HTMX
	return html.Div(html.ID(htmxProps.ID), g.Attr("data-state", "closed"))
}

// TriggerHTMX creates an HTMX-enhanced trigger
func TriggerHTMX(props TriggerProps, htmxProps HTMXProps, children ...g.Node) g.Node {
	classes := lib.CN(props.Class)
	
	return html.Button(
		html.Type("button"),
		g.If(classes != "", html.Class(classes)),
		hx.Get(htmxProps.TriggerPath),
		html.Target("#" + htmxProps.ID),
		hx.Swap("outerHTML"),
		g.Group(children),
	)
}

// OverlayHTMX creates an HTMX-enhanced overlay with close functionality
func OverlayHTMX(htmxProps HTMXProps, class ...string) g.Node {
	classes := lib.CN(
		"fixed inset-0 z-50 bg-black/80",
		"animate-in fade-in-0",
		lib.CN(class...),
	)

	return html.Div(
		html.Class(classes),
		g.Attr("data-state", "open"),
		hx.Get(htmxProps.ClosePath),
		html.Target("#" + htmxProps.ID),
		hx.Swap("outerHTML"),
		hx.Trigger("click"),
	)
}

// ContentHTMX creates HTMX-enhanced drawer content
func ContentHTMX(props ContentProps, htmxProps HTMXProps, side string, children ...g.Node) g.Node {
	// Get variant classes
	classes := drawerVariants.GetClasses(lib.VariantProps{
		Variant: side,
		Class:   props.Class,
	})

	// Add animation classes
	classes = lib.CN(classes, "animate-in")

	// Prevent clicks inside content from closing drawer
	return html.Div(
		html.Class(classes),
		g.Attr("data-state", "open"),
		g.Attr("data-side", side),
		hx.On("click", "event.stopPropagation()"),
		g.Group(children),
	)
}

// CloseHTMX creates an HTMX-enhanced close button
func CloseHTMX(props CloseProps, htmxProps HTMXProps, children ...g.Node) g.Node {
	classes := lib.CN(props.Class)

	return html.Button(
		html.Type("button"),
		g.If(classes != "", html.Class(classes)),
		hx.Get(htmxProps.ClosePath),
		html.Target("#" + htmxProps.ID),
		hx.Swap("outerHTML"),
		g.Group(children),
	)
}

// CloseButtonHTMX creates a standard close button with X icon
func CloseButtonHTMX(htmxProps HTMXProps) g.Node {
	return html.Button(
		html.Type("button"),
		html.Class("absolute right-4 top-4 rounded-sm opacity-70 ring-offset-background transition-opacity hover:opacity-100 focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 disabled:pointer-events-none"),
		hx.Get(htmxProps.ClosePath),
		html.Target("#" + htmxProps.ID),
		hx.Swap("outerHTML"),
		icons.X(html.Class("h-4 w-4")),
		html.Span(html.Class("sr-only"), g.Text("Close")),
	)
}

// FormDrawerHTMX creates a drawer with an HTMX-enhanced form
func FormDrawerHTMX(htmxProps HTMXProps, side string, formAction string, title string, children ...g.Node) g.Node {
	return NewHTMX(
		Props{Open: true},
		htmxProps,
		OverlayHTMX(htmxProps),
		ContentHTMX(
			ContentProps{},
			htmxProps,
			side,
			CloseButtonHTMX(htmxProps),
			html.Form(
				hx.Post(formAction),
				html.Target("#" + htmxProps.ID),
				hx.Swap("outerHTML"),
				DrawerHeader(
					HeaderProps{},
					DrawerTitle(TitleProps{}, g.Text(title)),
				),
				g.Group(children),
			),
		),
	)
}

// ExampleHTMX creates an HTMX-enhanced drawer example
func ExampleHTMX() g.Node {
	htmxProps := HTMXProps{
		ID:          "drawer-example",
		TriggerPath: "/api/drawer/open",
		ClosePath:   "/api/drawer/close",
	}
	
	return html.Div(
		TriggerHTMX(
			TriggerProps{Class: "bg-primary text-primary-foreground hover:bg-primary/90 px-4 py-2 rounded-md"},
			htmxProps,
			g.Text("Open Drawer"),
		),
		html.Div(html.ID(htmxProps.ID)),
	)
}

// RenderOpenDrawer renders an open drawer (for server response)
func RenderOpenDrawer(htmxProps HTMXProps, side string) g.Node {
	return NewHTMX(
		Props{Open: true},
		htmxProps,
		OverlayHTMX(htmxProps),
		ContentHTMX(
			ContentProps{},
			htmxProps,
			side,
			CloseButtonHTMX(htmxProps),
			DrawerHeader(
				HeaderProps{},
				DrawerTitle(TitleProps{}, g.Text("Edit Profile")),
				DrawerDescription(
					DescriptionProps{},
					g.Text("Make changes to your profile here. Click save when you're done."),
				),
			),
			html.Div(html.Class("grid gap-4 py-4"),
				html.Div(html.Class("grid grid-cols-4 items-center gap-4"),
					html.Label(html.For("name"), html.Class("text-right"), g.Text("Name")),
					html.Input(
						html.ID("name"),
						html.Name("name"),
						html.Value("Pedro Duarte"),
						html.Class("col-span-3"),
					),
				),
				html.Div(html.Class("grid grid-cols-4 items-center gap-4"),
					html.Label(html.For("username"), html.Class("text-right"), g.Text("Username")),
					html.Input(
						html.ID("username"),
						html.Name("username"),
						html.Value("@peduarte"),
						html.Class("col-span-3"),
					),
				),
			),
			DrawerFooter(
				FooterProps{},
				CloseHTMX(
					CloseProps{Class: "border hover:bg-accent px-4 py-2 rounded-md"},
					htmxProps,
					g.Text("Cancel"),
				),
				html.Button(
					html.Type("submit"),
					html.Class("bg-primary text-primary-foreground hover:bg-primary/90 px-4 py-2 rounded-md"),
					hx.Post("/api/drawer/save-profile"),
					html.Target("#" + htmxProps.ID),
					hx.Swap("outerHTML"),
					g.Text("Save changes"),
				),
			),
		),
	)
}

// RenderClosedDrawer renders a closed drawer (for server response)
func RenderClosedDrawer(htmxProps HTMXProps) g.Node {
	return html.Div(html.ID(htmxProps.ID), g.Attr("data-state", "closed"))
}

// NavigationDrawerExampleHTMX creates a navigation drawer with HTMX
func NavigationDrawerExampleHTMX() g.Node {
	htmxProps := HTMXProps{
		ID:          "nav-drawer",
		TriggerPath: "/api/drawer/nav/open",
		ClosePath:   "/api/drawer/nav/close",
	}
	
	return html.Div(
		html.Button(
			html.Type("button"),
			html.Class("fixed left-4 top-4 z-40 rounded-md border p-2"),
			hx.Get(htmxProps.TriggerPath),
			html.Target("#" + htmxProps.ID),
			hx.Swap("outerHTML"),
			icons.Menu(html.Class("h-4 w-4")),
			html.Span(html.Class("sr-only"), g.Text("Open navigation")),
		),
		html.Div(html.ID(htmxProps.ID)),
	)
}

// RenderNavigationDrawer renders a navigation drawer (for server response)
func RenderNavigationDrawer(htmxProps HTMXProps) g.Node {
	navItems := []struct {
		Title string
		Href  string
		Icon  g.Node
	}{
		{Title: "Dashboard", Href: "/dashboard", Icon: icons.Home(html.Class("h-4 w-4"))},
		{Title: "Users", Href: "/users", Icon: icons.Users(html.Class("h-4 w-4"))},
		{Title: "Products", Href: "/products", Icon: icons.Package(html.Class("h-4 w-4"))},
		{Title: "Settings", Href: "/settings", Icon: icons.Settings(html.Class("h-4 w-4"))},
	}

	return NewHTMX(
		Props{Open: true},
		htmxProps,
		OverlayHTMX(htmxProps),
		ContentHTMX(
			ContentProps{},
			htmxProps,
			"left",
			html.Div(html.Class("flex h-full flex-col"),
				html.Div(html.Class("flex items-center justify-between p-4 pb-2"),
					html.H2(html.Class("text-lg font-semibold"), g.Text("Navigation")),
					html.Button(
						html.Type("button"),
						html.Class("rounded-sm opacity-70 ring-offset-background transition-opacity hover:opacity-100"),
						hx.Get(htmxProps.ClosePath),
						html.Target("#" + htmxProps.ID),
						hx.Swap("outerHTML"),
						icons.X(html.Class("h-4 w-4")),
						html.Span(html.Class("sr-only"), g.Text("Close")),
					),
				),
				html.Nav(html.Class("flex-1 space-y-1 p-4"),
					g.Group(g.Map(navItems, func(item struct {
						Title string
						Href  string
						Icon  g.Node
					}) g.Node {
						return html.A(
							html.Href(item.Href),
							html.Class("flex items-center gap-3 rounded-lg px-3 py-2 text-sm transition-colors hover:bg-accent hover:text-accent-foreground"),
							item.Icon,
							g.Text(item.Title),
						)
					})),
				),
			),
		),
	)
}

// DrawerHandlers creates HTTP handlers for drawer components
func DrawerHandlers(mux *http.ServeMux) {
	// Basic drawer handlers
	htmxProps := HTMXProps{
		ID:          "drawer-example",
		TriggerPath: "/api/drawer/open",
		ClosePath:   "/api/drawer/close",
	}
	
	mux.HandleFunc("/api/drawer/open", func(w http.ResponseWriter, r *http.Request) {
		side := r.URL.Query().Get("side")
		if side == "" {
			side = "right"
		}
		node := RenderOpenDrawer(htmxProps, side)
		node.Render(w)
	})
	
	mux.HandleFunc("/api/drawer/close", func(w http.ResponseWriter, r *http.Request) {
		node := RenderClosedDrawer(htmxProps)
		node.Render(w)
	})
	
	mux.HandleFunc("/api/drawer/save-profile", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		name := r.FormValue("name")
		username := r.FormValue("username")
		
		// In a real app, save the profile here
		
		// Return success message
		node := html.Div(
			html.ID(htmxProps.ID),
			html.Class("fixed bottom-4 right-4 bg-green-500 text-white px-4 py-2 rounded-md shadow-lg"),
			g.Textf("Profile updated! Name: %s, Username: %s", name, username),
			// Auto-hide after 3 seconds
			g.Attr("x-data", "{}"),
			g.Attr("x-init", "setTimeout(() => $el.remove(), 3000)"),
		)
		node.Render(w)
	})
	
	// Navigation drawer handlers
	navProps := HTMXProps{
		ID:          "nav-drawer",
		TriggerPath: "/api/drawer/nav/open",
		ClosePath:   "/api/drawer/nav/close",
	}
	
	mux.HandleFunc("/api/drawer/nav/open", func(w http.ResponseWriter, r *http.Request) {
		node := RenderNavigationDrawer(navProps)
		node.Render(w)
	})
	
	mux.HandleFunc("/api/drawer/nav/close", func(w http.ResponseWriter, r *http.Request) {
		node := RenderClosedDrawer(navProps)
		node.Render(w)
	})
	
	// Multiple drawers example
	for _, side := range []string{"left", "right", "top", "bottom"} {
		s := side // capture loop variable
		mux.HandleFunc(fmt.Sprintf("/api/drawer/%s/open", s), func(w http.ResponseWriter, r *http.Request) {
			props := HTMXProps{
				ID:          fmt.Sprintf("%s-drawer", s),
				TriggerPath: fmt.Sprintf("/api/drawer/%s/open", s),
				ClosePath:   fmt.Sprintf("/api/drawer/%s/close", s),
			}
			
			node := NewHTMX(
				Props{Open: true},
				props,
				OverlayHTMX(props),
				ContentHTMX(
					ContentProps{},
					props,
					s,
					CloseButtonHTMX(props),
					DrawerHeader(
						HeaderProps{},
						DrawerTitle(TitleProps{}, g.Textf("%s Drawer", strings.Title(s))),
						DrawerDescription(
							DescriptionProps{},
							g.Textf("This drawer slides in from the %s.", s),
						),
					),
					html.Div(html.Class("flex-1 py-4"),
						html.P(html.Class("text-muted-foreground"),
							g.Text("Drawer content goes here..."),
						),
					),
					DrawerFooter(
						FooterProps{},
						CloseHTMX(
							CloseProps{Class: "border hover:bg-accent px-4 py-2 rounded-md"},
							props,
							g.Text("Close"),
						),
					),
				),
			)
			node.Render(w)
		})
		
		mux.HandleFunc(fmt.Sprintf("/api/drawer/%s/close", s), func(w http.ResponseWriter, r *http.Request) {
			props := HTMXProps{
				ID: fmt.Sprintf("%s-drawer", s),
			}
			node := RenderClosedDrawer(props)
			node.Render(w)
		})
	}
}