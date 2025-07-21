package navigationmenu

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// Props defines the properties for the NavigationMenu component
type Props struct {
	Class       string // Additional custom classes
	Orientation string // "horizontal" | "vertical" (default: "horizontal")
}

// ListProps defines properties for the navigation menu list
type ListProps struct {
	Class string
}

// ItemProps defines properties for navigation menu items
type ItemProps struct {
	Class string
	Value string // Unique value for the item
}

// TriggerProps defines properties for navigation triggers
type TriggerProps struct {
	Class    string
	Disabled bool
}

// ContentProps defines properties for navigation content
type ContentProps struct {
	Class string
}

// LinkProps defines properties for navigation links
type LinkProps struct {
	Class    string
	Active   bool
	Disabled bool
	Href     string
}

// ViewportProps defines properties for the viewport
type ViewportProps struct {
	Class string
}

// IndicatorProps defines properties for the active indicator
type IndicatorProps struct {
	Class string
}

// New creates a new NavigationMenu component
func New(props Props, children ...g.Node) g.Node {
	// Set defaults
	if props.Orientation == "" {
		props.Orientation = "horizontal"
	}

	classes := lib.CN(
		"relative z-10 flex max-w-max flex-1 items-center justify-center",
		props.Class,
	)

	return html.Nav(
		html.Class(classes),
		html.Role("navigation"),
		g.Attr("data-orientation", props.Orientation),
		g.Group(children),
	)
}

// List creates a navigation menu list
func ListComponent(props ListProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"group flex flex-1 list-none items-center justify-center space-x-1",
		props.Class,
	)

	return html.Ul(
		html.Class(classes),
		html.Role("none"),
		g.Attr("data-orientation", "horizontal"),
		g.Group(children),
	)
}

// Item creates a navigation menu item
func Item(props ItemProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"relative",
		props.Class,
	)

	attrs := []g.Node{
		html.Class(classes),
		html.Role("none"),
	}

	if props.Value != "" {
		attrs = append(attrs, g.Attr("data-value", props.Value))
	}

	return html.Li(
		append(attrs, children...)...,
	)
}

// Trigger creates a navigation menu trigger
func Trigger(props TriggerProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"group inline-flex h-10 w-max items-center justify-center rounded-md bg-background px-4 py-2 text-sm font-medium transition-colors",
		"hover:bg-accent hover:text-accent-foreground",
		"focus:bg-accent focus:text-accent-foreground focus:outline-none",
		"disabled:pointer-events-none disabled:opacity-50",
		"data-[active]:bg-accent/50 data-[state=open]:bg-accent/50",
		props.Class,
	)

	attrs := []g.Node{
		html.Type("button"),
		html.Class(classes),
		html.Role("menuitem"),
		g.Attr("aria-haspopup", "menu"),
		g.Attr("aria-expanded", "false"),
		g.Attr("data-state", "closed"),
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

// Content creates the dropdown content for navigation
func ContentComponent(props ContentProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"left-0 top-0 w-full data-[motion^=from-]:animate-in data-[motion^=to-]:animate-out",
		"data-[motion^=from-]:fade-in data-[motion^=to-]:fade-out",
		"data-[motion=from-end]:slide-in-from-right-52",
		"data-[motion=from-start]:slide-in-from-left-52",
		"data-[motion=to-end]:slide-out-to-right-52",
		"data-[motion=to-start]:slide-out-to-left-52",
		"md:absolute md:w-auto",
		props.Class,
	)

	return html.Div(
		html.Class(classes),
		g.Attr("data-state", "closed"),
		html.Style("display: none;"), // Hidden by default
		g.Group(children),
	)
}

// Link creates a navigation link
func LinkComponent(props LinkProps, children ...g.Node) g.Node {
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

	return html.A(
		append(attrs, children...)...,
	)
}

// Viewport creates the viewport container for content
func Viewport(props ViewportProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"origin-top-center relative mt-1.5 h-[var(--radix-navigation-menu-viewport-height)] w-full overflow-hidden rounded-md border bg-popover text-popover-foreground shadow-lg",
		"data-[state=open]:animate-in data-[state=closed]:animate-out",
		"data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-90",
		"md:w-[var(--radix-navigation-menu-viewport-width)]",
		props.Class,
	)

	return html.Div(
		html.Class(classes),
		g.Attr("data-state", "closed"),
		html.Style("display: none;"), // Hidden by default
		g.Group(children),
	)
}

// Indicator creates an active indicator
func Indicator(props IndicatorProps) g.Node {
	classes := lib.CN(
		"top-full z-[1] flex h-1.5 items-end justify-center overflow-hidden",
		"data-[state=visible]:animate-in data-[state=hidden]:animate-out",
		"data-[state=hidden]:fade-out data-[state=visible]:fade-in",
		props.Class,
	)

	return html.Div(
		html.Class(classes),
		g.Attr("data-state", "hidden"),
		html.Div(
			html.Class("relative top-[60%] h-2 w-2 rotate-45 rounded-tl-sm bg-border shadow-md"),
		),
	)
}

// ListItem creates a styled list item for content
func ListItem(title string, href string, description string) g.Node {
	return html.Li(
	LinkComponent(
			LinkProps{Href: href},
			html.Div(html.Class("text-sm font-medium leading-none"), g.Text(title)),
			html.P(html.Class("line-clamp-2 text-sm leading-snug text-muted-foreground"),
				g.Text(description),
			),
		),
	)
}

// Default creates a default navigation menu
func Default() g.Node {
	return New(Props{})
}

// WithViewport creates a navigation menu with viewport
func WithViewport(children ...g.Node) g.Node {
	return html.Div(
		html.Class("relative"),
		New(
			Props{},
			g.Group(children),
		),
		Viewport(ViewportProps{}),
	)
}

// SimpleMenu creates a simple navigation menu with links
func SimpleMenu() g.Node {
	return New(
		Props{},
		ListComponent(
			ListProps{},
			Item(
				ItemProps{},
	LinkComponent(LinkProps{Href: "/"}, g.Text("Home")),
			),
			Item(
				ItemProps{},
	LinkComponent(LinkProps{Href: "/about"}, g.Text("About")),
			),
			Item(
				ItemProps{},
	LinkComponent(LinkProps{Href: "/services"}, g.Text("Services")),
			),
			Item(
				ItemProps{},
	LinkComponent(LinkProps{Href: "/contact"}, g.Text("Contact")),
			),
		),
	)
}

// WithDropdowns creates a navigation menu with dropdown menus
func WithDropdowns() g.Node {
	return New(
		Props{},
		ListComponent(
			ListProps{},
			Item(
				ItemProps{Value: "getting-started"},
				Trigger(TriggerProps{}, g.Text("Getting started")),
	ContentComponent(
					ContentProps{},
					html.Ul(html.Class("grid gap-3 p-4 md:w-[400px] lg:w-[500px] lg:grid-cols-[.75fr_1fr]"),
						html.Li(html.Class("row-span-3"),
	LinkComponent(
								LinkProps{
									Href:  "/",
									Class: "flex h-full w-full select-none flex-col justify-end rounded-md bg-gradient-to-b from-muted/50 to-muted p-6 no-underline outline-none focus:shadow-md",
								},
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
				),
			),
			Item(
				ItemProps{Value: "components"},
				Trigger(TriggerProps{}, g.Text("Components")),
	ContentComponent(
					ContentProps{},
					html.Ul(html.Class("grid w-[400px] gap-3 p-4 md:w-[500px] md:grid-cols-2 lg:w-[600px]"),
						ListItem("Alert Dialog", "/docs/primitives/alert-dialog", "A modal dialog that interrupts the user with important content and expects a response."),
						ListItem("Hover Card", "/docs/primitives/hover-card", "For sighted users to preview content available behind a link."),
						ListItem("Progress", "/docs/primitives/progress", "Displays an indicator showing the completion progress of a task, typically displayed as a progress bar."),
						ListItem("Scroll-area", "/docs/primitives/scroll-area", "Visually or semantically separates content."),
						ListItem("Tabs", "/docs/primitives/tabs", "A set of layered sections of content—known as tab panels—that are displayed one at a time."),
						ListItem("Tooltip", "/docs/primitives/tooltip", "A popup that displays information related to an element when the element receives keyboard focus or the mouse hovers over it."),
					),
				),
			),
			Item(
				ItemProps{},
	LinkComponent(LinkProps{Href: "/docs"}, g.Text("Documentation")),
			),
		),
	)
}