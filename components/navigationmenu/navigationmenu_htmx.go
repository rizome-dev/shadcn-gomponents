package navigationmenu

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	hx "maragu.dev/gomponents-htmx"
	"github.com/rizome-dev/shadcn-gomponents/lib"
	"fmt"
)

// HTMXProps defines HTMX-specific properties for the NavigationMenu
type HTMXProps struct {
	ID            string // Unique ID for the navigation menu
	ContentPath   string // Server path for menu content
	ViewportID    string // ID for the viewport element
	IndicatorID   string // ID for the indicator element
	PushURL       bool   // Whether to push URL to history
	PreloadOnHover bool   // Whether to preload content on hover
}

// ItemHTMXProps defines HTMX properties for individual items
type ItemHTMXProps struct {
	ContentPath string // Server path to fetch item content
	Target      string // HTMX target for content
	Indicator   string // Loading indicator ID
}

// NewHTMX creates an HTMX-enhanced NavigationMenu component
func NewHTMX(props Props, htmxProps HTMXProps, children ...g.Node) g.Node {
	// Set defaults
	if props.Orientation == "" {
		props.Orientation = "horizontal"
	}

	classes := lib.CN(
		"relative z-10 flex max-w-max flex-1 items-center justify-center",
		props.Class,
	)

	attrs := []g.Node{
		html.Class(classes),
		html.Role("navigation"),
		g.Attr("data-orientation", props.Orientation),
		g.If(htmxProps.ID != "", html.ID(htmxProps.ID)),
		g.Attr("data-navigation-menu", "true"),
	}

	// Add global event handlers
	attrs = append(attrs,
		// Close other menus when opening a new one
		g.Attr("hx-on:htmx:after-request", `
			if (event.detail.target.matches('[data-navigation-content]')) {
				// Close all other open menus
				const allContents = this.querySelectorAll('[data-navigation-content]');
				allContents.forEach(content => {
					if (content !== event.detail.target) {
						content.style.display = 'none';
						content.setAttribute('data-state', 'closed');
						const trigger = document.querySelector('[data-navigation-trigger="' + content.dataset.navigationContent + '"]');
						if (trigger) {
							trigger.setAttribute('data-state', 'closed');
							trigger.setAttribute('aria-expanded', 'false');
						}
					}
				});
			}
		`),
	)

	return html.Nav(
		append(attrs, children...)...,
	)
}

// TriggerHTMX creates an HTMX-enhanced navigation trigger
func TriggerHTMX(props TriggerProps, itemValue string, htmxProps ItemHTMXProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"group inline-flex h-10 w-max items-center justify-center rounded-md bg-background px-4 py-2 text-sm font-medium transition-colors",
		"hover:bg-accent hover:text-accent-foreground",
		"focus:bg-accent focus:text-accent-foreground focus:outline-none",
		"disabled:pointer-events-none disabled:opacity-50",
		"data-[active]:bg-accent/50 data-[state=open]:bg-accent/50",
		props.Class,
	)

	target := htmxProps.Target
	if target == "" {
		target = fmt.Sprintf("#nav-content-%s", itemValue)
	}

	attrs := []g.Node{
		html.Type("button"),
		html.Class(classes),
		html.Role("menuitem"),
		g.Attr("aria-haspopup", "menu"),
		g.Attr("aria-expanded", "false"),
		g.Attr("data-state", "closed"),
		g.Attr("data-navigation-trigger", itemValue),
		
		// HTMX attributes
		hx.Get(htmxProps.ContentPath),
		hx.Target(target),
		hx.Swap("innerHTML"),
		hx.Trigger("click"),
		g.If(htmxProps.Indicator != "", hx.Indicator("#"+htmxProps.Indicator)),
		
		// Preload on hover if enabled
		g.If(htmxProps.ContentPath != "",
			g.Attr("hx-trigger", "click, mouseenter[ctrlKey||metaKey||shiftKey] once"),
		),
		
		// Toggle menu state
		g.Attr("onclick", fmt.Sprintf(`
			const trigger = this;
			const isOpen = trigger.getAttribute('data-state') === 'open';
			const nav = trigger.closest('[data-navigation-menu]');
			const content = document.querySelector('%s');
			
			if (!isOpen) {
				// Close other triggers
				if (nav) {
					nav.querySelectorAll('[data-navigation-trigger]').forEach(t => {
						if (t !== trigger) {
							t.setAttribute('data-state', 'closed');
							t.setAttribute('aria-expanded', 'false');
						}
					});
				}
				
				trigger.setAttribute('data-state', 'open');
				trigger.setAttribute('aria-expanded', 'true');
				if (content) {
					content.style.display = 'block';
					content.setAttribute('data-state', 'open');
				}
			} else {
				trigger.setAttribute('data-state', 'closed');
				trigger.setAttribute('aria-expanded', 'false');
				if (content) {
					content.style.display = 'none';
					content.setAttribute('data-state', 'closed');
				}
			}
		`, target)),
	}

	if props.Disabled {
		attrs = append(attrs, html.Disabled())
	}

	// Add chevron down icon
	chevron := g.Raw(`<svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg" class="relative top-[1px] ml-1 h-3 w-3 transition duration-200 group-data-[state=open]:rotate-180" aria-hidden="true">
		<path d="M3.13523 6.15803C3.3241 5.95657 3.64052 5.94637 3.84197 6.13523L7.5 9.56464L11.158 6.13523C11.3595 5.94637 11.6759 5.95657 11.8648 6.15803C12.0536 6.35949 12.0434 6.67591 11.842 6.86477L7.84197 10.6148C7.64964 10.7951 7.35036 10.7951 7.15803 10.6148L3.15803 6.86477C2.95657 6.67591 2.94637 6.35949 3.13523 6.15803Z" fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"></path>
	</svg>`)

	return html.Button(
		append(attrs, g.Group(children), chevron)...,
	)
}

// ContentHTMX creates HTMX-aware navigation content
func ContentHTMX(props ContentProps, itemValue string, children ...g.Node) g.Node {
	classes := lib.CN(
		"left-0 top-0 w-full animate-in fade-in zoom-in-90",
		"md:absolute md:w-auto",
		props.Class,
	)

	contentID := fmt.Sprintf("nav-content-%s", itemValue)

	return html.Div(
		html.ID(contentID),
		html.Class(classes),
		g.Attr("data-state", "open"),
		g.Attr("data-navigation-content", itemValue),
		
		// Click outside to close
		g.Attr("hx-on:click.outside", fmt.Sprintf(`
			const trigger = document.querySelector('[data-navigation-trigger="%s"]');
			if (trigger) {
				trigger.setAttribute('data-state', 'closed');
				trigger.setAttribute('aria-expanded', 'false');
			}
			this.style.display = 'none';
			this.setAttribute('data-state', 'closed');
		`, itemValue)),
		
		// Keyboard navigation
		g.Attr("onkeydown", fmt.Sprintf(`
			if (event.key === 'Escape') {
				const trigger = document.querySelector('[data-navigation-trigger="%s"]');
				if (trigger) {
					trigger.setAttribute('data-state', 'closed');
					trigger.setAttribute('aria-expanded', 'false');
					trigger.focus();
				}
				this.style.display = 'none';
				this.setAttribute('data-state', 'closed');
			}
		`, itemValue)),
		
		g.Group(children),
	)
}

// LinkHTMX creates an HTMX-enhanced navigation link
func LinkHTMX(props LinkProps, htmxProps HTMXProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"block select-none space-y-1 rounded-md p-3 leading-none no-underline outline-none transition-colors",
		"hover:bg-accent hover:text-accent-foreground",
		"focus:bg-accent focus:text-accent-foreground",
		lib.CNIf(props.Active,
			"bg-accent text-accent-foreground",
			"",
		),
		lib.CNIf(props.Disabled,
			"pointer-events-none opacity-50",
			"",
		),
		props.Class,
	)

	attrs := []g.Node{
		html.Class(classes),
		html.Href(props.Href),
		html.Role("menuitem"),
	}

	if props.Active {
		attrs = append(attrs, g.Attr("aria-current", "page"))
	}

	if props.Disabled {
		attrs = append(attrs, g.Attr("aria-disabled", "true"))
	}

	// Add HTMX navigation if push URL is enabled
	if htmxProps.PushURL && props.Href != "" && !props.Disabled {
		attrs = append(attrs,
			hx.Get(props.Href),
			hx.PushURL("true"),
			hx.Target("body"),
			hx.Swap("innerHTML transition:true"),
		)
	}

	// Close menu on click
	attrs = append(attrs,
		g.Attr("onclick", `
			const content = this.closest('[data-navigation-content]');
			if (content) {
				const itemValue = content.dataset.navigationContent;
				const trigger = document.querySelector('[data-navigation-trigger="' + itemValue + '"]');
				if (trigger) {
					trigger.setAttribute('data-state', 'closed');
					trigger.setAttribute('aria-expanded', 'false');
				}
				content.style.display = 'none';
				content.setAttribute('data-state', 'closed');
			}
		`),
	)

	return html.A(
		append(attrs, children...)...,
	)
}

// ViewportHTMX creates an HTMX-aware viewport
func ViewportHTMX(props ViewportProps, htmxProps HTMXProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"origin-top-center relative mt-1.5 overflow-hidden rounded-md border bg-popover text-popover-foreground shadow-lg",
		"animate-in zoom-in-90",
		"md:w-auto",
		props.Class,
	)

	viewportID := htmxProps.ViewportID
	if viewportID == "" {
		viewportID = "navigation-viewport"
	}

	return html.Div(
		html.ID(viewportID),
		html.Class(classes),
		g.Attr("data-state", "open"),
		g.Group(children),
	)
}

// ExampleHTMX creates an HTMX-enhanced navigation menu example
func ExampleHTMX() g.Node {
	htmxProps := HTMXProps{
		ID:             "main-nav",
		PreloadOnHover: true,
		PushURL:        false,
	}

	return NewHTMX(
		Props{},
		htmxProps,
		ListComponent(
			ListProps{},
			Item(
				ItemProps{Value: "getting-started"},
				TriggerHTMX(
					TriggerProps{},
					"getting-started",
					ItemHTMXProps{
						ContentPath: "/api/nav/getting-started",
					},
					g.Text("Getting started"),
				),
				html.Div(html.ID("nav-content-getting-started")), // Content placeholder
			),
			Item(
				ItemProps{Value: "components"},
				TriggerHTMX(
					TriggerProps{},
					"components",
					ItemHTMXProps{
						ContentPath: "/api/nav/components",
					},
					g.Text("Components"),
				),
				html.Div(html.ID("nav-content-components")), // Content placeholder
			),
			Item(
				ItemProps{},
				LinkHTMX(
					LinkProps{Href: "/docs"},
					htmxProps,
					g.Text("Documentation"),
				),
			),
		),
	)
}

// RenderGettingStartedContent renders the getting started menu content (for server response)
func RenderGettingStartedContent() g.Node {
	return ContentHTMX(
		ContentProps{},
		"getting-started",
		html.Ul(html.Class("grid gap-3 p-4 md:w-[400px] lg:w-[500px] lg:grid-cols-[.75fr_1fr]"),
			html.Li(html.Class("row-span-3"),
				LinkHTMX(
					LinkProps{
						Href:  "/",
						Class: "flex h-full w-full select-none flex-col justify-end rounded-md bg-gradient-to-b from-muted/50 to-muted p-6 no-underline outline-none focus:shadow-md",
					},
					HTMXProps{},
					html.Div(html.Class("mb-2 mt-4 text-lg font-medium"),
						g.Text("shadcn/ui"),
					),
					html.P(html.Class("text-sm leading-tight text-muted-foreground"),
						g.Text("Beautifully designed components built with Radix UI and Tailwind CSS."),
					),
				),
			),
			ListItem("Introduction", "/docs", "Re-usable components built using Radix UI and Tailwind CSS."),
			ListItem("Installation", "/docs/installation", "How to install dependencies and structure your app."),
			ListItem("Typography", "/docs/primitives/typography", "Styles for headings, paragraphs, lists...etc"),
		),
	)
}

// RenderComponentsContent renders the components menu content (for server response)
func RenderComponentsContent() g.Node {
	return ContentHTMX(
		ContentProps{},
		"components",
		html.Ul(html.Class("grid w-[400px] gap-3 p-4 md:w-[500px] md:grid-cols-2 lg:w-[600px]"),
			ListItem("Alert Dialog", "/docs/primitives/alert-dialog", "A modal dialog that interrupts the user with important content and expects a response."),
			ListItem("Hover Card", "/docs/primitives/hover-card", "For sighted users to preview content available behind a link."),
			ListItem("Progress", "/docs/primitives/progress", "Displays an indicator showing the completion progress of a task, typically displayed as a progress bar."),
			ListItem("Scroll-area", "/docs/primitives/scroll-area", "Visually or semantically separates content."),
			ListItem("Tabs", "/docs/primitives/tabs", "A set of layered sections of content—known as tab panels—that are displayed one at a time."),
			ListItem("Tooltip", "/docs/primitives/tooltip", "A popup that displays information related to an element when the element receives keyboard focus or the mouse hovers over it."),
		),
	)
}

// MegaMenuHTMX creates a mega menu navigation example
func MegaMenuHTMX() g.Node {
	htmxProps := HTMXProps{
		ID:             "mega-nav",
		PreloadOnHover: true,
	}

	return NewHTMX(
		Props{},
		htmxProps,
		ListComponent(
			ListProps{},
			Item(
				ItemProps{Value: "products"},
				TriggerHTMX(
					TriggerProps{},
					"products",
					ItemHTMXProps{
						ContentPath: "/api/nav/products",
					},
					g.Text("Products"),
				),
				html.Div(html.ID("nav-content-products")),
			),
			Item(
				ItemProps{Value: "solutions"},
				TriggerHTMX(
					TriggerProps{},
					"solutions",
					ItemHTMXProps{
						ContentPath: "/api/nav/solutions",
					},
					g.Text("Solutions"),
				),
				html.Div(html.ID("nav-content-solutions")),
			),
			Item(
				ItemProps{Value: "resources"},
				TriggerHTMX(
					TriggerProps{},
					"resources",
					ItemHTMXProps{
						ContentPath: "/api/nav/resources",
					},
					g.Text("Resources"),
				),
				html.Div(html.ID("nav-content-resources")),
			),
			Item(
				ItemProps{},
				LinkHTMX(
					LinkProps{Href: "/pricing"},
					htmxProps,
					g.Text("Pricing"),
				),
			),
		),
	)
}

// RenderProductsMenu renders a products mega menu (for server response)
func RenderProductsMenu() g.Node {
	return ContentHTMX(
		ContentProps{},
		"products",
		ViewportHTMX(
			ViewportProps{},
			HTMXProps{},
			html.Div(html.Class("grid gap-3 p-6 md:w-[700px] md:grid-cols-3"),
				html.Div(html.Class("space-y-3"),
					html.H3(html.Class("font-medium leading-none mb-2"), g.Text("Analytics")),
					LinkHTMX(
						LinkProps{Href: "/products/analytics/overview"},
						HTMXProps{},
						html.Div(html.Class("text-sm font-medium leading-none"), g.Text("Overview")),
						html.P(html.Class("text-sm text-muted-foreground"), g.Text("Real-time analytics platform")),
					),
					LinkHTMX(
						LinkProps{Href: "/products/analytics/dashboards"},
						HTMXProps{},
						html.Div(html.Class("text-sm font-medium leading-none"), g.Text("Dashboards")),
						html.P(html.Class("text-sm text-muted-foreground"), g.Text("Customizable dashboards")),
					),
				),
				html.Div(html.Class("space-y-3"),
					html.H3(html.Class("font-medium leading-none mb-2"), g.Text("Automation")),
					LinkHTMX(
						LinkProps{Href: "/products/automation/workflows"},
						HTMXProps{},
						html.Div(html.Class("text-sm font-medium leading-none"), g.Text("Workflows")),
						html.P(html.Class("text-sm text-muted-foreground"), g.Text("Automate repetitive tasks")),
					),
					LinkHTMX(
						LinkProps{Href: "/products/automation/integrations"},
						HTMXProps{},
						html.Div(html.Class("text-sm font-medium leading-none"), g.Text("Integrations")),
						html.P(html.Class("text-sm text-muted-foreground"), g.Text("Connect with your tools")),
					),
				),
				html.Div(html.Class("space-y-3"),
					html.H3(html.Class("font-medium leading-none mb-2"), g.Text("Security")),
					LinkHTMX(
						LinkProps{Href: "/products/security/access"},
						HTMXProps{},
						html.Div(html.Class("text-sm font-medium leading-none"), g.Text("Access Control")),
						html.P(html.Class("text-sm text-muted-foreground"), g.Text("Manage user permissions")),
					),
					LinkHTMX(
						LinkProps{Href: "/products/security/audit"},
						HTMXProps{},
						html.Div(html.Class("text-sm font-medium leading-none"), g.Text("Audit Logs")),
						html.P(html.Class("text-sm text-muted-foreground"), g.Text("Track all activities")),
					),
				),
			),
		),
	)
}