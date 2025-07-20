package dialog

import (
	"net/http"
	"strings"
	
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	hx "maragu.dev/gomponents-htmx"
	"github.com/rizome-dev/shadcn-gomponents/lib"
	"github.com/rizome-dev/shadcn-gomponents/lib/icons"
)

// HTMXProps defines HTMX-specific properties for the Dialog
type HTMXProps struct {
	ID          string // Unique ID for the dialog
	TriggerPath string // Server path for trigger actions
	ClosePath   string // Server path for close actions
	ContentPath string // Server path for loading dynamic content
}

// NewHTMX creates an HTMX-enhanced Dialog component
func NewHTMX(props Props, htmxProps HTMXProps, children ...g.Node) g.Node {
	if props.Open {
		classes := lib.CN("fixed inset-0 z-50", props.Class)
		return html.Div(
			html.ID(htmxProps.ID),
			html.Class(classes),
			g.Group(children),
		)
	}
	// Return empty div that can be replaced by HTMX
	return html.Div(html.ID(htmxProps.ID))
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
		lib.CN(class...),
	)

	return html.Div(
		html.Class(classes),
		hx.Get(htmxProps.ClosePath),
		html.Target("#" + htmxProps.ID),
		hx.Swap("outerHTML"),
		hx.Trigger("click"),
	)
}

// DialogContentHTMX creates HTMX-enhanced dialog content with close button
func DialogContentHTMX(props ContentProps, htmxProps HTMXProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"fixed left-[50%] top-[50%] z-50 grid w-full max-w-lg translate-x-[-50%] translate-y-[-50%] gap-4 border bg-background p-6 shadow-lg sm:rounded-lg",
		props.Class,
	)

	contentChildren := children
	if props.ShowCloseButton {
		closeButton := html.Button(
			html.Type("button"),
			html.Class("absolute right-4 top-4 rounded-sm opacity-70 ring-offset-background transition-opacity hover:opacity-100 focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 disabled:pointer-events-none"),
			hx.Get(htmxProps.ClosePath),
			html.Target("#" + htmxProps.ID),
			hx.Swap("outerHTML"),
			icons.X(html.Class("h-4 w-4")),
			html.Span(html.Class("sr-only"), g.Text("Close")),
		)
		contentChildren = append([]g.Node{closeButton}, children...)
	}

	// Prevent clicks inside content from closing dialog
	return html.Div(
		html.Class(classes),
		hx.On("click", "event.stopPropagation()"),
		g.Group(contentChildren),
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

// FormDialogHTMX creates a dialog with an HTMX-enhanced form
func FormDialogHTMX(htmxProps HTMXProps, formAction string, title string, children ...g.Node) g.Node {
	return NewHTMX(
		Props{Open: true},
		htmxProps,
		OverlayHTMX(htmxProps),
		DialogContentHTMX(
			ContentProps{ShowCloseButton: true},
			htmxProps,
			html.Form(
				hx.Post(formAction),
				html.Target("#" + htmxProps.ID),
				hx.Swap("outerHTML"),
				DialogHeader(
					HeaderProps{},
					DialogTitle(TitleProps{}, title),
				),
				g.Group(children),
			),
		),
	)
}

// ExampleHTMX creates an HTMX-enhanced dialog example
func ExampleHTMX() g.Node {
	htmxProps := HTMXProps{
		ID:          "dialog-example",
		TriggerPath: "/api/dialog/open",
		ClosePath:   "/api/dialog/close",
	}
	
	return html.Div(
		TriggerHTMX(
			TriggerProps{},
			htmxProps,
			g.Text("Open Dialog"),
		),
		html.Div(html.ID(htmxProps.ID)),
	)
}

// RenderOpenDialog renders an open dialog (for server response)
func RenderOpenDialog(htmxProps HTMXProps) g.Node {
	return NewHTMX(
		Props{Open: true},
		htmxProps,
		OverlayHTMX(htmxProps),
		DialogContentHTMX(
			ContentProps{ShowCloseButton: true},
			htmxProps,
			DialogHeader(
				HeaderProps{},
				DialogTitle(TitleProps{}, "Edit Profile"),
				Description(DescriptionProps{}, "Make changes to your profile here. Click save when you're done."),
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
			DialogFooter(
				FooterProps{},
				CloseHTMX(CloseProps{Class: "border hover:bg-accent"}, htmxProps, g.Text("Cancel")),
				html.Button(
					html.Type("submit"),
					html.Class("bg-primary text-primary-foreground hover:bg-primary/90"),
					hx.Post("/api/dialog/save-profile"),
					html.Target("#" + htmxProps.ID),
					hx.Swap("outerHTML"),
					g.Text("Save changes"),
				),
			),
		),
	)
}

// RenderClosedDialog renders a closed dialog (for server response)
func RenderClosedDialog(htmxProps HTMXProps) g.Node {
	return html.Div(html.ID(htmxProps.ID))
}

// LoginFormExampleHTMX creates a login form dialog with HTMX
func LoginFormExampleHTMX() g.Node {
	htmxProps := HTMXProps{
		ID:          "login-dialog",
		TriggerPath: "/api/dialog/login/open",
		ClosePath:   "/api/dialog/login/close",
	}
	
	return html.Div(
		TriggerHTMX(
			TriggerProps{Class: "bg-primary text-primary-foreground hover:bg-primary/90"},
			htmxProps,
			g.Text("Login"),
		),
		html.Div(html.ID(htmxProps.ID)),
	)
}

// RenderLoginDialog renders the login dialog (for server response)
func RenderLoginDialog(htmxProps HTMXProps) g.Node {
	return FormDialogHTMX(
		htmxProps,
		"/api/auth/login",
		"Login to your account",
		Description(DescriptionProps{}, "Enter your credentials to access your account."),
		html.Div(html.Class("grid gap-4 py-4"),
			html.Div(html.Class("grid gap-2"),
				html.Label(html.For("email"), g.Text("Email")),
				html.Input(
					html.Type("email"),
					html.ID("email"),
					html.Name("email"),
					html.Placeholder("m@example.com"),
					html.Required(),
					hx.Trigger("keyup changed delay:500ms"),
					hx.Post("/api/validate/email"),
					html.Target("#email-error"),
					hx.Swap("innerHTML"),
				),
				html.Div(html.ID("email-error"), html.Class("text-sm text-destructive")),
			),
			html.Div(html.Class("grid gap-2"),
				html.Label(html.For("password"), g.Text("Password")),
				html.Input(
					html.Type("password"),
					html.ID("password"),
					html.Name("password"),
					html.Required(),
				),
			),
			html.Div(html.Class("flex items-center space-x-2"),
				html.Input(
					html.Type("checkbox"),
					html.ID("remember"),
					html.Name("remember"),
					html.Class("h-4 w-4"),
				),
				html.Label(
					html.For("remember"),
					html.Class("text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"),
					g.Text("Remember me"),
				),
			),
		),
		DialogFooter(
			FooterProps{},
			CloseHTMX(CloseProps{Class: "border hover:bg-accent"}, htmxProps, g.Text("Cancel")),
			html.Button(
				html.Type("submit"),
				html.Class("bg-primary text-primary-foreground hover:bg-primary/90"),
				g.Text("Login"),
			),
		),
	)
}

// SearchDialogExampleHTMX creates a search dialog with HTMX
func SearchDialogExampleHTMX() g.Node {
	htmxProps := HTMXProps{
		ID:          "search-dialog",
		TriggerPath: "/api/dialog/search/open",
		ClosePath:   "/api/dialog/search/close",
	}
	
	return html.Div(
		html.Button(
			html.Type("button"),
			html.Class("inline-flex items-center gap-2 border rounded-md px-3 py-2 text-sm"),
			hx.Get(htmxProps.TriggerPath),
			html.Target("#" + htmxProps.ID),
			hx.Swap("outerHTML"),
			icons.Search(html.Class("h-4 w-4")),
			g.Text("Searchtml..."),
			html.Kbd(html.Class("ml-auto text-xs"), g.Text("⌘K")),
		),
		html.Div(html.ID(htmxProps.ID)),
	)
}

// RenderSearchDialog renders the search dialog (for server response)
func RenderSearchDialog(htmxProps HTMXProps) g.Node {
	return NewHTMX(
		Props{Open: true},
		htmxProps,
		OverlayHTMX(htmxProps),
		DialogContentHTMX(
			ContentProps{Class: "max-w-2xl", ShowCloseButton: false},
			htmxProps,
			html.Div(html.Class("flex items-center border-b px-3"),
				icons.Search(html.Class("mr-2 h-4 w-4 shrink-0 opacity-50")),
				html.Input(
					html.Type("text"),
					html.Class("flex h-11 w-full rounded-md bg-transparent py-3 text-sm outline-none placeholder:text-muted-foreground disabled:cursor-not-allowed disabled:opacity-50"),
					html.Placeholder("Type a command or searchtml..."),
					hx.Post("/api/search"),
					hx.Trigger("keyup changed delay:300ms"),
					html.Target("#search-results"),
					hx.Swap("innerHTML"),
					hx.Indicator("#search-spinner"),
				),
				html.Div(html.ID("search-spinner"), html.Class("htmx-indicator"),
					icons.Loader(html.Class("h-4 w-4")),
				),
			),
			html.Div(html.ID("search-results"), html.Class("max-h-[300px] overflow-y-auto p-4"),
				// Results will be loaded here
				html.P(html.Class("text-sm text-muted-foreground text-center py-6"),
					g.Text("Start typing to searchtml..."),
				),
			),
		),
	)
}

// DialogHandlers creates HTTP handlers for dialog components
func DialogHandlers(mux *http.ServeMux) {
	// Basic dialog handlers
	htmxProps := HTMXProps{
		ID:          "dialog-example",
		TriggerPath: "/api/dialog/open",
		ClosePath:   "/api/dialog/close",
	}
	
	mux.HandleFunc("/api/dialog/open", func(w http.ResponseWriter, r *http.Request) {
		node := RenderOpenDialog(htmxProps)
		node.Render(w)
	})
	
	mux.HandleFunc("/api/dialog/close", func(w http.ResponseWriter, r *http.Request) {
		node := RenderClosedDialog(htmxProps)
		node.Render(w)
	})
	
	mux.HandleFunc("/api/dialog/save-profile", func(w http.ResponseWriter, r *http.Request) {
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
	
	// Login dialog handlers
	loginProps := HTMXProps{
		ID:          "login-dialog",
		TriggerPath: "/api/dialog/login/open",
		ClosePath:   "/api/dialog/login/close",
	}
	
	mux.HandleFunc("/api/dialog/login/open", func(w http.ResponseWriter, r *http.Request) {
		node := RenderLoginDialog(loginProps)
		node.Render(w)
	})
	
	mux.HandleFunc("/api/dialog/login/close", func(w http.ResponseWriter, r *http.Request) {
		node := RenderClosedDialog(loginProps)
		node.Render(w)
	})
	
	mux.HandleFunc("/api/validate/email", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		email := r.FormValue("email")
		
		// Simple email validation
		if email == "" {
			return
		}
		
		if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
			node := html.Span(html.Class("text-destructive text-sm"), g.Text("Please enter a valid email address"))
			node.Render(w)
			return
		}
		
		// Check if email exists (in a real app)
		if email == "taken@example.com" {
			node := html.Span(html.Class("text-destructive text-sm"), g.Text("This email is already registered"))
			node.Render(w)
			return
		}
	})
	
	mux.HandleFunc("/api/auth/login", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		email := r.FormValue("email")
		password := r.FormValue("password")
		remember := r.FormValue("remember") == "on"
		
		// In a real app, authenticate the user
		if email == "test@example.com" && password == "password" {
			// Success
			node := html.Div(
				html.ID(loginProps.ID),
				html.Class("fixed bottom-4 right-4 bg-green-500 text-white px-4 py-2 rounded-md shadow-lg"),
				g.Textf("Welcome back! Remember me: %v", remember),
				// In a real app, you would set cookies and redirect
			)
			node.Render(w)
		} else {
			// Error
			node := html.Div(
				html.ID(loginProps.ID),
				html.Class("fixed bottom-4 right-4 bg-red-500 text-white px-4 py-2 rounded-md shadow-lg"),
				g.Text("Invalid email or password"),
			)
			node.Render(w)
		}
	})
	
	// Search dialog handlers
	searchProps := HTMXProps{
		ID:          "search-dialog",
		TriggerPath: "/api/dialog/search/open",
		ClosePath:   "/api/dialog/search/close",
	}
	
	mux.HandleFunc("/api/dialog/search/open", func(w http.ResponseWriter, r *http.Request) {
		node := RenderSearchDialog(searchProps)
		node.Render(w)
	})
	
	mux.HandleFunc("/api/dialog/search/close", func(w http.ResponseWriter, r *http.Request) {
		node := RenderClosedDialog(searchProps)
		node.Render(w)
	})
	
	mux.HandleFunc("/api/search", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		query := r.FormValue("value")
		
		if query == "" {
			node := html.P(html.Class("text-sm text-muted-foreground text-center py-6"),
				g.Text("Start typing to searchtml..."),
			)
			node.Render(w)
			return
		}
		
		// Mock search results
		results := []struct {
			Title       string
			Description string
			Category    string
		}{
			{"Components", "Browse all UI components", "Navigation"},
			{"Documentation", "Read the documentation", "Help"},
			{"Settings", "Manage your account settings", "Account"},
			{"Keyboard Shortcuts", "View all keyboard shortcuts", "Help"},
		}
		
		// Filter results
		var filtered []struct {
			Title       string
			Description string
			Category    string
		}
		
		for _, result := range results {
			if strings.Contains(strings.ToLower(result.Title), strings.ToLower(query)) ||
				strings.Contains(strings.ToLower(result.Description), strings.ToLower(query)) {
				filtered = append(filtered, result)
			}
		}
		
		if len(filtered) == 0 {
			node := html.P(html.Class("text-sm text-muted-foreground text-center py-6"),
				g.Textf("No results found for '%s'", query),
			)
			node.Render(w)
			return
		}
		
		// Render results
		node := html.Div(
			g.Group(g.Map(filtered, func(result struct {
				Title       string
				Description string
				Category    string
			}) g.Node {
				return html.Button(
					html.Type("button"),
					html.Class("w-full flex items-start gap-3 rounded-lg px-3 py-2 text-left hover:bg-accent"),
					html.Div(html.Class("flex flex-col gap-1"),
						html.Div(html.Class("text-sm font-medium"), g.Text(result.Title)),
						html.Div(html.Class("text-xs text-muted-foreground"), g.Text(result.Description)),
					),
					html.Span(html.Class("ml-auto text-xs text-muted-foreground"), g.Text(result.Category)),
				)
			})),
		)
		node.Render(w)
	})
}