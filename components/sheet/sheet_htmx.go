package sheet

import (
	"net/http"
	
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	hx "maragu.dev/gomponents-htmx"
	"github.com/rizome-dev/shadcn-gomponents/lib"
	"github.com/rizome-dev/shadcn-gomponents/lib/icons"
)

// HTMXProps defines HTMX-specific properties for the Sheet
type HTMXProps struct {
	ID          string // Unique ID for the sheet
	TriggerPath string // Server path for trigger actions
	ClosePath   string // Server path for close actions
	ContentPath string // Server path for loading dynamic content
}

// NewHTMX creates an HTMX-enhanced Sheet component
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
		hx.Target("#" + htmxProps.ID),
		hx.Swap("outerHTML"),
		g.Group(children),
	)
}

// OverlayHTMX creates an HTMX-enhanced overlay with close functionality
func OverlayHTMX(props OverlayProps, htmxProps HTMXProps) g.Node {
	classes := lib.CN(
		"fixed inset-0 z-50 bg-black/80",
		"animate-in fade-in-0",
		props.Class,
	)
	
	return html.Div(
		html.Class(classes),
		g.Attr("data-sheet-overlay", ""),
		g.Attr("data-state", "open"),
		hx.Get(htmxProps.ClosePath),
		hx.Target("#" + htmxProps.ID),
		hx.Swap("outerHTML"),
		hx.Trigger("click"),
	)
}

// ContentHTMX creates HTMX-enhanced sheet content
func ContentHTMX(props ContentProps, htmxProps HTMXProps, children ...g.Node) g.Node {
	// Default side to right
	if props.Side == "" {
		props.Side = "right"
	}
	
	// Base classes
	baseClasses := "fixed z-50 gap-4 bg-background p-6 shadow-lg transition ease-in-out"
	animationClasses := "animate-in duration-500"
	
	// Side-specific classes
	var sideClasses string
	var slideClasses string
	
	switch props.Side {
	case "top":
		sideClasses = "inset-x-0 top-0 border-b"
		slideClasses = "slide-in-from-top"
	case "bottom":
		sideClasses = "inset-x-0 bottom-0 border-t"
		slideClasses = "slide-in-from-bottom"
	case "left":
		sideClasses = "inset-y-0 left-0 h-full w-3/4 border-r sm:max-w-sm"
		slideClasses = "slide-in-from-left"
	case "right":
		sideClasses = "inset-y-0 right-0 h-full w-3/4 border-l sm:max-w-sm"
		slideClasses = "slide-in-from-right"
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
			html.Class("absolute right-4 top-4 rounded-sm opacity-70 ring-offset-background transition-opacity hover:opacity-100 focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2"),
			g.Attr("aria-label", "Close"),
			hx.Get(htmxProps.ClosePath),
			hx.Target("#" + htmxProps.ID),
			hx.Swap("outerHTML"),
			icons.X(html.Class("h-4 w-4")),
			html.Span(html.Class("sr-only"), g.Text("Close")),
		)
		contentChildren = append([]g.Node{closeButton}, children...)
	}
	
	// Prevent clicks inside content from closing sheet
	return html.Div(
		html.Class(classes),
		html.Role("dialog"),
		g.Attr("aria-modal", "true"),
		g.Attr("data-state", "open"),
		g.Attr("data-sheet-content", ""),
		g.If(props.Side != "", g.Attr("data-side", props.Side)),
		hx.On("click", "event.stopPropagation()"),
		// Add escape key handler
		g.If(props.CloseOnEsc,
			hx.On("keyup", "if(event.key === 'Escape') htmx.ajax('GET', '"+htmxProps.ClosePath+"', {target: '#"+htmxProps.ID+"', swap: 'outerHTML'})"),
		),
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
		hx.Target("#" + htmxProps.ID),
		hx.Swap("outerHTML"),
		g.Group(children),
	)
}

// RenderOpenSheet renders an open sheet (for server response)
func RenderOpenSheet(props Props, contentProps ContentProps, htmxProps HTMXProps, content g.Node) g.Node {
	return NewHTMX(
		Props{Open: true},
		htmxProps,
		OverlayHTMX(OverlayProps{}, htmxProps),
		ContentHTMX(contentProps, htmxProps, content),
	)
}

// RenderClosedSheet renders a closed sheet (for server response)
func RenderClosedSheet(htmxProps HTMXProps) g.Node {
	return html.Div(html.ID(htmxProps.ID), g.Attr("data-state", "closed"))
}

// FormSheetHTMX creates a sheet with an HTMX-enhanced form
func FormSheetHTMX(htmxProps HTMXProps, formAction string, title string, children ...g.Node) g.Node {
	return NewHTMX(
		Props{Open: true},
		htmxProps,
		OverlayHTMX(OverlayProps{}, htmxProps),
		ContentHTMX(
			ContentProps{Side: "right", ShowCloseButton: true},
			htmxProps,
			html.Form(
				hx.Post(formAction),
				hx.Target("#" + htmxProps.ID),
				hx.Swap("outerHTML"),
				HeaderComponent(
					HeaderProps{},
					Title(TitleProps{}, g.Text(title)),
				),
				g.Group(children),
			),
		),
	)
}

// DynamicContentSheetHTMX creates a sheet that loads content dynamically
func DynamicContentSheetHTMX(props ContentProps, htmxProps HTMXProps) g.Node {
	return ContentHTMX(
		props,
		htmxProps,
		html.Div(
			html.Class("relative min-h-[200px]"),
			hx.Get(htmxProps.ContentPath),
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

// NavigationSheetHTMX creates a navigation sheet
func NavigationSheetHTMX(htmxProps HTMXProps, items []NavItem) g.Node {
	return ContentHTMX(
		ContentProps{Side: "left", ShowCloseButton: true},
		htmxProps,
		html.Nav(
			html.Class("flex flex-col space-y-4"),
			HeaderComponent(
				HeaderProps{},
				Title(TitleProps{}, g.Text("Navigation")),
			),
			html.Div(
				html.Class("flex-1 overflow-y-auto"),
				g.Group(g.Map(items, func(item NavItem) g.Node {
					itemClass := "flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground transition-all hover:text-primary"
					if item.Active {
						itemClass = "flex items-center gap-3 rounded-lg bg-muted px-3 py-2 text-primary transition-all"
					}
					
					if item.Action != "" {
						return html.A(
							html.Href("#"),
							html.Class(itemClass),
							hx.Get(item.Action),
							hx.Target("#main-content"),
							hx.Swap("innerHTML"),
							hx.On("click", "htmx.ajax('GET', '"+htmxProps.ClosePath+"', {target: '#"+htmxProps.ID+"', swap: 'outerHTML'})"),
							g.If(item.Icon != nil, item.Icon),
							g.Text(item.Label),
							g.If(item.Badge != "",
								html.Span(
									html.Class("ml-auto flex h-6 w-6 shrink-0 items-center justify-center rounded-full bg-primary text-primary-foreground text-xs"),
									g.Text(item.Badge),
								),
							),
						)
					}
					
					return html.A(
						html.Href(item.Href),
						html.Class(itemClass),
						g.If(item.Icon != nil, item.Icon),
						g.Text(item.Label),
						g.If(item.Badge != "",
							html.Span(
								html.Class("ml-auto flex h-6 w-6 shrink-0 items-center justify-center rounded-full bg-primary text-primary-foreground text-xs"),
								g.Text(item.Badge),
							),
						),
					)
				})),
			),
		),
	)
}

// NavItem represents a navigation item
type NavItem struct {
	Label  string
	Href   string
	Icon   g.Node
	Active bool
	Badge  string
	Action string // HTMX action endpoint
}

// SheetHandlers creates HTTP handlers for sheet components
func SheetHandlers(mux *http.ServeMux) {
	// Basic sheet example
	basicProps := HTMXProps{
		ID:          "sheet-basic",
		TriggerPath: "/api/sheet/basic/open",
		ClosePath:   "/api/sheet/basic/close",
	}
	
	mux.HandleFunc("/api/sheet/basic/open", func(w http.ResponseWriter, r *http.Request) {
		content := html.Div(
			HeaderComponent(
				HeaderProps{},
				Title(TitleProps{}, g.Text("Edit Profile")),
				Description(DescriptionProps{}, g.Text("Make changes to your profile here. Click save when you're done.")),
			),
			html.Div(html.Class("grid gap-4 py-4"),
				html.Div(html.Class("grid grid-cols-4 items-center gap-4"),
					html.Label(html.For("name"), html.Class("text-right"), g.Text("Name")),
					html.Input(
						html.ID("name"),
						html.Value("Pedro Duarte"),
						html.Class("col-span-3"),
					),
				),
				html.Div(html.Class("grid grid-cols-4 items-center gap-4"),
					html.Label(html.For("username"), html.Class("text-right"), g.Text("Username")),
					html.Input(
						html.ID("username"),
						html.Value("@peduarte"),
						html.Class("col-span-3"),
					),
				),
			),
			FooterComponent(
				FooterProps{},
				CloseHTMX(CloseProps{Class: "border"}, basicProps, g.Text("Cancel")),
				html.Button(html.Type("submit"), g.Text("Save changes")),
			),
		)
		
		node := RenderOpenSheet(
			Props{},
			ContentProps{Side: "right", ShowCloseButton: true},
			basicProps,
			content,
		)
		node.Render(w)
	})
	
	mux.HandleFunc("/api/sheet/basic/close", func(w http.ResponseWriter, r *http.Request) {
		node := RenderClosedSheet(basicProps)
		node.Render(w)
	})
	
	// Navigation sheet example
	navProps := HTMXProps{
		ID:          "sheet-nav",
		TriggerPath: "/api/sheet/nav/open",
		ClosePath:   "/api/sheet/nav/close",
	}
	
	mux.HandleFunc("/api/sheet/nav/open", func(w http.ResponseWriter, r *http.Request) {
		items := []NavItem{
			{Label: "Dashboard", Href: "/", Icon: icons.Home(html.Class("h-4 w-4")), Active: true},
			{Label: "Orders", Href: "/orders", Icon: icons.Package(html.Class("h-4 w-4")), Badge: "3"},
			{Label: "Products", Href: "/products", Icon: icons.Package(html.Class("h-4 w-4"))},
			{Label: "Customers", Href: "/customers", Icon: icons.Users(html.Class("h-4 w-4"))},
			{Label: "Settings", Href: "/settings", Icon: icons.Settings(html.Class("h-4 w-4"))},
		}
		
		node := RenderOpenSheet(
			Props{},
			ContentProps{},
			navProps,
			NavigationSheetHTMX(navProps, items),
		)
		node.Render(w)
	})
	
	mux.HandleFunc("/api/sheet/nav/close", func(w http.ResponseWriter, r *http.Request) {
		node := RenderClosedSheet(navProps)
		node.Render(w)
	})
	
	// Form sheet example
	formProps := HTMXProps{
		ID:          "sheet-form",
		TriggerPath: "/api/sheet/form/open",
		ClosePath:   "/api/sheet/form/close",
	}
	
	mux.HandleFunc("/api/sheet/form/open", func(w http.ResponseWriter, r *http.Request) {
		node := FormSheetHTMX(
			formProps,
			"/api/sheet/form/submit",
			"Create New Product",
			html.Div(html.Class("grid gap-4 py-4"),
				html.Div(html.Class("grid gap-2"),
					html.Label(html.For("product-name"), g.Text("Product Name")),
					html.Input(
						html.ID("product-name"),
						html.Name("name"),
						html.Placeholder("Enter product name"),
						html.Required(),
					),
				),
				html.Div(html.Class("grid gap-2"),
					html.Label(html.For("product-price"), g.Text("Price")),
					html.Input(
						html.Type("number"),
						html.ID("product-price"),
						html.Name("price"),
						html.Placeholder("0.00"),
						html.Step("0.01"),
						html.Required(),
					),
				),
				html.Div(html.Class("grid gap-2"),
					html.Label(html.For("product-description"), g.Text("Description")),
					html.Textarea(
						html.ID("product-description"),
						html.Name("description"),
						html.Placeholder("Enter product description"),
						html.Class("min-h-[100px]"),
					),
				),
			),
			FooterComponent(
				FooterProps{},
				CloseHTMX(CloseProps{Class: "border"}, formProps, g.Text("Cancel")),
				html.Button(html.Type("submit"), g.Text("Create Product")),
			),
		)
		node.Render(w)
	})
	
	mux.HandleFunc("/api/sheet/form/submit", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		name := r.FormValue("name")
		price := r.FormValue("price")
		description := r.FormValue("description")
		
		// In a real app, save the product here
		
		// Return success message and close sheet
		node := html.Div(
			html.ID(formProps.ID),
			g.Attr("data-state", "closed"),
			html.Div(
				html.Class("fixed bottom-4 right-4 bg-green-500 text-white px-4 py-2 rounded-md shadow-lg"),
				g.Textf("Product '%s' created successfully! Price: $%s", name, price),
				g.If(description != "", html.P(html.Class("text-sm mt-1"), g.Textf("Description: %s", description))),
				// Auto-hide after 3 seconds
				g.Attr("x-data", "{}"),
				g.Attr("x-init", "setTimeout(() => $el.remove(), 3000)"),
			),
		)
		node.Render(w)
	})
	
	// Dynamic content sheet
	dynamicProps := HTMXProps{
		ID:          "sheet-dynamic",
		TriggerPath: "/api/sheet/dynamic/open",
		ClosePath:   "/api/sheet/dynamic/close",
		ContentPath: "/api/sheet/dynamic/content",
	}
	
	mux.HandleFunc("/api/sheet/dynamic/open", func(w http.ResponseWriter, r *http.Request) {
		node := RenderOpenSheet(
			Props{},
			ContentProps{Side: "right", ShowCloseButton: true},
			dynamicProps,
			DynamicContentSheetHTMX(ContentProps{}, dynamicProps),
		)
		node.Render(w)
	})
	
	mux.HandleFunc("/api/sheet/dynamic/content", func(w http.ResponseWriter, r *http.Request) {
		// Simulate loading delay
		// time.Sleep(500 * time.Millisecond)
		
		node := html.Div(
			HeaderComponent(
				HeaderProps{},
				Title(TitleProps{}, g.Text("Account Details")),
				Description(DescriptionProps{}, g.Text("View and manage your account information.")),
			),
			html.Div(html.Class("space-y-4 py-4"),
				html.Div(
					html.H3(html.Class("text-sm font-medium"), g.Text("Personal Information")),
					html.Dl(html.Class("mt-2 divide-y"),
						html.Div(html.Class("py-2"),
							html.Dt(html.Class("text-sm text-muted-foreground"), g.Text("Full Name")),
							html.Dd(html.Class("mt-1 text-sm"), g.Text("John Doe")),
						),
						html.Div(html.Class("py-2"),
							html.Dt(html.Class("text-sm text-muted-foreground"), g.Text("Email")),
							html.Dd(html.Class("mt-1 text-sm"), g.Text("john.doe@example.com")),
						),
						html.Div(html.Class("py-2"),
							html.Dt(html.Class("text-sm text-muted-foreground"), g.Text("Phone")),
							html.Dd(html.Class("mt-1 text-sm"), g.Text("+1 (555) 123-4567")),
						),
					),
				),
				html.Div(
					html.H3(html.Class("text-sm font-medium"), g.Text("Subscription")),
					html.Div(html.Class("mt-2 rounded-lg border p-4"),
						html.Div(html.Class("flex items-center justify-between"),
							html.Div(
								html.P(html.Class("font-medium"), g.Text("Pro Plan")),
								html.P(html.Class("text-sm text-muted-foreground"), g.Text("$29/month")),
							),
							html.Span(html.Class("rounded-full bg-green-100 px-2 py-1 text-xs text-green-700"), g.Text("Active")),
						),
					),
				),
			),
		)
		node.Render(w)
	})
	
	// Multi-step form sheet
	multiStepProps := HTMXProps{
		ID:          "sheet-multistep",
		TriggerPath: "/api/sheet/multistep/open",
		ClosePath:   "/api/sheet/multistep/close",
	}
	
	mux.HandleFunc("/api/sheet/multistep/open", func(w http.ResponseWriter, r *http.Request) {
		step := r.URL.Query().Get("step")
		if step == "" {
			step = "1"
		}
		
		var content g.Node
		switch step {
		case "1":
			content = html.Div(
				HeaderComponent(
					HeaderProps{},
					Title(TitleProps{}, g.Text("Step 1: Basic Information")),
					Description(DescriptionProps{}, g.Text("Let's start with your basic details.")),
				),
				html.Form(
					hx.Post("/api/sheet/multistep/next?step=2"),
					hx.Target("#" + multiStepProps.ID),
					hx.Swap("outerHTML"),
					html.Div(html.Class("grid gap-4 py-4"),
						html.Div(html.Class("grid gap-2"),
							html.Label(html.For("first-name"), g.Text("First Name")),
							html.Input(html.ID("first-name"), html.Name("firstName"), html.Required()),
						),
						html.Div(html.Class("grid gap-2"),
							html.Label(html.For("last-name"), g.Text("Last Name")),
							html.Input(html.ID("last-name"), html.Name("lastName"), html.Required()),
						),
					),
					FooterComponent(
						FooterProps{},
						CloseHTMX(CloseProps{Class: "border"}, multiStepProps, g.Text("Cancel")),
						html.Button(html.Type("submit"), g.Text("Next")),
					),
				),
			)
		case "2":
			content = html.Div(
				HeaderComponent(
					HeaderProps{},
					Title(TitleProps{}, g.Text("Step 2: Contact Information")),
					Description(DescriptionProps{}, g.Text("How can we reach you?")),
				),
				html.Form(
					hx.Post("/api/sheet/multistep/next?step=3"),
					hx.Target("#" + multiStepProps.ID),
					hx.Swap("outerHTML"),
					html.Div(html.Class("grid gap-4 py-4"),
						html.Div(html.Class("grid gap-2"),
							html.Label(html.For("email"), g.Text("Email")),
							html.Input(html.Type("email"), html.ID("email"), html.Name("email"), html.Required()),
						),
						html.Div(html.Class("grid gap-2"),
							html.Label(html.For("phone"), g.Text("Phone")),
							html.Input(html.Type("tel"), html.ID("phone"), html.Name("phone")),
						),
					),
					FooterComponent(
						FooterProps{},
						html.Button(
							html.Type("button"),
							html.Class("border mr-auto"),
							hx.Get("/api/sheet/multistep/open?step=1"),
							hx.Target("#" + multiStepProps.ID),
							hx.Swap("outerHTML"),
							g.Text("Back"),
						),
						CloseHTMX(CloseProps{Class: "border"}, multiStepProps, g.Text("Cancel")),
						html.Button(html.Type("submit"), g.Text("Finish")),
					),
				),
			)
		case "3":
			content = html.Div(
				HeaderComponent(
					HeaderProps{},
					Title(TitleProps{}, g.Text("All Done!")),
					Description(DescriptionProps{}, g.Text("Your information has been saved successfully.")),
				),
				html.Div(html.Class("py-8 text-center"),
					html.Div(html.Class("mx-auto w-12 h-12 rounded-full bg-green-100 flex items-center justify-center mb-4"),
						icons.Check(html.Class("h-6 w-6 text-green-600")),
					),
					html.P(html.Class("text-sm text-muted-foreground"), g.Text("Thank you for completing the form.")),
				),
				FooterComponent(
					FooterProps{},
					CloseHTMX(CloseProps{Class: "w-full"}, multiStepProps, g.Text("Close")),
				),
			)
		}
		
		node := RenderOpenSheet(
			Props{},
			ContentProps{Side: "right", ShowCloseButton: true},
			multiStepProps,
			content,
		)
		node.Render(w)
	})
	
	mux.HandleFunc("/api/sheet/multistep/next", func(w http.ResponseWriter, r *http.Request) {
		step := r.URL.Query().Get("step")
		if step == "" {
			step = "1"
		}
		
		// In a real app, you would save the form data here
		r.ParseForm()
		
		// Redirect to the next step
		http.Redirect(w, r, "/api/sheet/multistep/open?step="+step, http.StatusSeeOther)
	})
}

// Example creates an example sheet
func ExampleHTMX() g.Node {
	htmxProps := HTMXProps{
		ID:          "sheet-example",
		TriggerPath: "/api/sheet/basic/open",
		ClosePath:   "/api/sheet/basic/close",
	}
	
	return html.Div(
		TriggerHTMX(
			TriggerProps{},
			htmxProps,
			g.Text("Open Sheet"),
		),
		html.Div(html.ID(htmxProps.ID)),
	)
}

// MobileMenuExampleHTMX creates a mobile menu sheet
func MobileMenuExampleHTMX() g.Node {
	htmxProps := HTMXProps{
		ID:          "mobile-menu",
		TriggerPath: "/api/sheet/nav/open",
		ClosePath:   "/api/sheet/nav/close",
	}
	
	return html.Div(
		TriggerHTMX(
			TriggerProps{Class: "lg:hidden"},
			htmxProps,
			icons.Menu(html.Class("h-6 w-6")),
		),
		html.Div(html.ID(htmxProps.ID)),
	)
}

// FilterSheetExampleHTMX creates a filter sheet
func FilterSheetExampleHTMX() g.Node {
	htmxProps := HTMXProps{
		ID:          "filter-sheet",
		TriggerPath: "/api/sheet/filter/open",
		ClosePath:   "/api/sheet/filter/close",
	}
	
	return html.Div(
		TriggerHTMX(
			TriggerProps{Class: "inline-flex items-center gap-2"},
			htmxProps,
			g.El("svg",
				g.Attr("xmlns", "http://www.w3.org/2000/svg"),
				g.Attr("viewBox", "0 0 24 24"),
				g.Attr("fill", "none"),
				g.Attr("stroke", "currentColor"),
				g.Attr("stroke-width", "2"),
				g.Attr("stroke-linecap", "round"),
				g.Attr("stroke-linejoin", "round"),
				html.Class("h-4 w-4"),
				g.El("polygon", g.Attr("points", "22 3 2 3 10 12.46 10 19 14 21 14 12.46 22 3")),
			),
			g.Text("Filters"),
		),
		html.Div(html.ID(htmxProps.ID)),
	)
}

// SettingsSheetExampleHTMX creates a settings sheet
func SettingsSheetExampleHTMX() g.Node {
	htmxProps := HTMXProps{
		ID:          "settings-sheet",
		TriggerPath: "/api/sheet/dynamic/open",
		ClosePath:   "/api/sheet/dynamic/close",
		ContentPath: "/api/sheet/dynamic/content",
	}
	
	return html.Div(
		TriggerHTMX(
			TriggerProps{Class: "inline-flex items-center gap-2"},
			htmxProps,
			icons.Settings(html.Class("h-4 w-4")),
			g.Text("Settings"),
		),
		html.Div(html.ID(htmxProps.ID)),
	)
}