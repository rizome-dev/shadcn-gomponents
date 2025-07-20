package sidebar

import (
	"fmt"

	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

const (
	SidebarWidth       = "16rem"
	SidebarWidthMobile = "18rem"
	SidebarWidthIcon   = "3rem"
)

// Props defines the properties for the sidebar component
type Props struct {
	Side        string // "left" | "right"
	Variant     string // "sidebar" | "floating" | "inset"
	Collapsible string // "offcanvas" | "icon" | "none"
	Class       string
	ID          string
}

// ProviderProps defines the properties for the sidebar provider
type ProviderProps struct {
	DefaultOpen bool
	Open        *bool
	Class       string
	Style       string
}

// Provider creates a sidebar provider wrapper
func Provider(props ProviderProps, children ...g.Node) g.Node {
	style := fmt.Sprintf("--sidebar-width: %s; --sidebar-width-icon: %s;", SidebarWidth, SidebarWidthIcon)
	if props.Style != "" {
		style = style + " " + props.Style
	}

	classes := lib.CN(
		"group/sidebar-wrapper has-[[data-variant=inset]]:bg-sidebar flex min-h-svh w-full",
		props.Class,
	)

	return html.Div(
		g.Attr("data-sidebar-wrapper", "true"),
		g.Attr("style", style),
		html.Class(classes),
		g.Group(children),
	)
}

// New creates a new sidebar component
func New(props Props, children ...g.Node) g.Node {
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
		g.Attr("data-variant", props.Variant),
		g.Attr("data-side", props.Side),
		g.Attr("data-collapsible", props.Collapsible),
		g.If(props.ID != "", g.Attr("id", props.ID)),
		
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

// Trigger creates a sidebar trigger button
func Trigger(props Props, children ...g.Node) g.Node {
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
		g.Group(buttonChildren),
	)
}

// Rail creates a sidebar rail (drag handle)
func Rail(props Props) g.Node {
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
	)
}

// Inset creates a sidebar inset (main content area)
func Inset(props Props, children ...g.Node) g.Node {
	return g.El("main",
		g.Attr("data-sidebar-inset", "true"),
		html.Class(lib.CN(
			"bg-background relative flex w-full flex-1 flex-col",
			"md:peer-data-[variant=inset]:m-2 md:peer-data-[variant=inset]:ml-0 md:peer-data-[variant=inset]:rounded-xl md:peer-data-[variant=inset]:shadow-sm md:peer-data-[variant=inset]:peer-data-[state=collapsed]:ml-2",
			props.Class,
		)),
		g.Group(children),
	)
}

// Header creates a sidebar header
func HeaderComponent(props Props, children ...g.Node) g.Node {
	return html.Div(
		g.Attr("data-sidebar-header", "true"),
		html.Class(lib.CN("flex flex-col gap-2 p-2", props.Class)),
		g.Group(children),
	)
}

// Content creates a sidebar content area
func ContentComponent(props Props, children ...g.Node) g.Node {
	return html.Div(
		g.Attr("data-sidebar-content", "true"),
		html.Class(lib.CN(
			"flex min-h-0 flex-1 flex-col gap-2 overflow-auto group-data-[collapsible=icon]:overflow-hidden",
			props.Class,
		)),
		g.Group(children),
	)
}

// Footer creates a sidebar footer
func FooterComponent(props Props, children ...g.Node) g.Node {
	return html.Div(
		g.Attr("data-sidebar-footer", "true"),
		html.Class(lib.CN("flex flex-col gap-2 p-2", props.Class)),
		g.Group(children),
	)
}

// Separator creates a sidebar separator
func Separator(props Props) g.Node {
	return html.Hr(
		g.Attr("data-sidebar-separator", "true"),
		html.Class(lib.CN("bg-sidebar-border mx-2 w-auto", props.Class)),
	)
}

// Input creates a sidebar input
func Input(props Props) g.Node {
	return g.El("input",
		g.Attr("data-sidebar-input", "true"),
		g.Attr("type", "text"),
		html.Class(lib.CN("bg-background h-8 w-full shadow-none", props.Class)),
	)
}

// Group creates a sidebar group
func Group(props Props, children ...g.Node) g.Node {
	return html.Div(
		g.Attr("data-sidebar-group", "true"),
		html.Class(lib.CN("relative flex w-full min-w-0 flex-col p-2", props.Class)),
		g.Group(children),
	)
}

// GroupLabel creates a sidebar group label
func GroupLabel(props Props, children ...g.Node) g.Node {
	return html.Div(
		g.Attr("data-sidebar-group-label", "true"),
		html.Class(lib.CN(
			"text-sidebar-foreground/70 ring-sidebar-ring flex h-8 shrink-0 items-center rounded-md px-2 text-xs font-medium outline-none transition-[margin,opacity] duration-200 ease-linear focus-visible:ring-2 [&>svg]:size-4 [&>svg]:shrink-0",
			"group-data-[collapsible=icon]:-mt-8 group-data-[collapsible=icon]:opacity-0",
			props.Class,
		)),
		g.Group(children),
	)
}

// GroupAction creates a sidebar group action button
func GroupAction(props Props, children ...g.Node) g.Node {
	return html.Button(
		g.Attr("data-sidebar-group-action", "true"),
		g.Attr("type", "button"),
		html.Class(lib.CN(
			"text-sidebar-foreground ring-sidebar-ring hover:bg-sidebar-accent hover:text-sidebar-accent-foreground absolute top-3.5 right-3 flex aspect-square w-5 items-center justify-center rounded-md p-0 outline-none transition-transform focus-visible:ring-2 [&>svg]:size-4 [&>svg]:shrink-0",
			"after:absolute after:-inset-2 md:after:hidden",
			"group-data-[collapsible=icon]:hidden",
			props.Class,
		)),
		g.Group(children),
	)
}

// GroupContent creates a sidebar group content area
func GroupContent(props Props, children ...g.Node) g.Node {
	return html.Div(
		g.Attr("data-sidebar-group-content", "true"),
		html.Class(lib.CN("w-full text-sm", props.Class)),
		g.Group(children),
	)
}

// Menu creates a sidebar menu list
func Menu(props Props, children ...g.Node) g.Node {
	return html.Ul(
		g.Attr("data-sidebar-menu", "true"),
		html.Class(lib.CN("flex w-full min-w-0 flex-col gap-1", props.Class)),
		g.Group(children),
	)
}

// MenuItem creates a sidebar menu item
func MenuItem(props Props, children ...g.Node) g.Node {
	return html.Li(
		g.Attr("data-sidebar-menu-item", "true"),
		html.Class(lib.CN("group/menu-item relative", props.Class)),
		g.Group(children),
	)
}

// MenuButtonProps defines properties for menu buttons
type MenuButtonProps struct {
	Variant  string // "default" | "outline"
	Size     string // "default" | "sm" | "lg"
	IsActive bool
	Tooltip  string
	Href     string
	Class    string
}

var menuButtonVariants = lib.VariantConfig{
	Base: "peer/menu-button flex w-full items-center gap-2 overflow-hidden rounded-md p-2 text-left text-sm outline-none ring-sidebar-ring transition-[width,height,padding] hover:bg-sidebar-accent hover:text-sidebar-accent-foreground focus-visible:ring-2 active:bg-sidebar-accent active:text-sidebar-accent-foreground disabled:pointer-events-none disabled:opacity-50 group-has-[[data-sidebar=menu-action]]/menu-item:pr-8 aria-disabled:pointer-events-none aria-disabled:opacity-50 data-[active=true]:bg-sidebar-accent data-[active=true]:font-medium data-[active=true]:text-sidebar-accent-foreground data-[state=open]:hover:bg-sidebar-accent data-[state=open]:hover:text-sidebar-accent-foreground group-data-[collapsible=icon]:!size-8 group-data-[collapsible=icon]:!p-2 [&>span:last-child]:truncate [&>svg]:size-4 [&>svg]:shrink-0",
	Variants: map[string]map[string]string{
		"variant": {
			"default": "hover:bg-sidebar-accent hover:text-sidebar-accent-foreground",
			"outline": "bg-background shadow-[0_0_0_1px_hsl(var(--sidebar-border))] hover:bg-sidebar-accent hover:text-sidebar-accent-foreground hover:shadow-[0_0_0_1px_hsl(var(--sidebar-accent))]",
		},
		"size": {
			"default": "h-8 text-sm",
			"sm":      "h-7 text-xs",
			"lg":      "h-12 text-sm group-data-[collapsible=icon]:!p-0",
		},
	},
	Defaults: map[string]string{
		"variant": "default",
		"size":    "default",
	},
}

// MenuButton creates a sidebar menu button
func MenuButton(props MenuButtonProps, children ...g.Node) g.Node {
	// Set defaults
	if props.Variant == "" {
		props.Variant = "default"
	}
	if props.Size == "" {
		props.Size = "default"
	}

	classes := menuButtonVariants.GetClasses(lib.VariantProps{
		Variant: props.Variant,
		Size:    props.Size,
		Class:   props.Class,
	})

	elem := "button"
	attrs := []g.Node{
		g.Attr("data-sidebar-menu-button", "true"),
		g.Attr("data-size", props.Size),
		html.Class(classes),
	}

	if props.Href != "" {
		elem = "a"
		attrs = append(attrs, html.Href(props.Href))
	} else {
		attrs = append(attrs, html.Type("button"))
	}

	if props.IsActive {
		attrs = append(attrs, g.Attr("data-active", "true"))
	}

	return g.El(elem,
		append(attrs, children...)...,
	)
}

// MenuAction creates a sidebar menu action button
func MenuAction(props Props, children ...g.Node) g.Node {
	return html.Button(
		g.Attr("data-sidebar-menu-action", "true"),
		g.Attr("type", "button"),
		html.Class(lib.CN(
			"text-sidebar-foreground ring-sidebar-ring hover:bg-sidebar-accent hover:text-sidebar-accent-foreground peer-hover/menu-button:text-sidebar-accent-foreground absolute top-1.5 right-1 flex aspect-square w-5 items-center justify-center rounded-md p-0 outline-none transition-transform focus-visible:ring-2 [&>svg]:size-4 [&>svg]:shrink-0",
			"after:absolute after:-inset-2 md:after:hidden",
			"peer-data-[size=sm]/menu-button:top-1",
			"peer-data-[size=default]/menu-button:top-1.5",
			"peer-data-[size=lg]/menu-button:top-2.5",
			"group-data-[collapsible=icon]:hidden",
			props.Class,
		)),
		g.Group(children),
	)
}

// MenuBadge creates a sidebar menu badge
func MenuBadge(props Props, children ...g.Node) g.Node {
	return html.Div(
		g.Attr("data-sidebar-menu-badge", "true"),
		html.Class(lib.CN(
			"text-sidebar-foreground pointer-events-none absolute right-1 flex h-5 min-w-5 items-center justify-center rounded-md px-1 text-xs font-medium tabular-nums select-none",
			"peer-hover/menu-button:text-sidebar-accent-foreground peer-data-[active=true]/menu-button:text-sidebar-accent-foreground",
			"peer-data-[size=sm]/menu-button:top-1",
			"peer-data-[size=default]/menu-button:top-1.5",
			"peer-data-[size=lg]/menu-button:top-2.5",
			"group-data-[collapsible=icon]:hidden",
			props.Class,
		)),
		g.Group(children),
	)
}

// MenuSkeleton creates a sidebar menu skeleton loader
func MenuSkeleton(props Props, showIcon bool) g.Node {
	skeletonItems := []g.Node{}
	
	if showIcon {
		skeletonItems = append(skeletonItems, 
			html.Div(
				html.Class("size-4 rounded-md bg-muted animate-pulse"),
				g.Attr("data-sidebar-menu-skeleton-icon", "true"),
			),
		)
	}
	
	skeletonItems = append(skeletonItems,
		html.Div(
			html.Class("h-4 w-[60%] bg-muted animate-pulse rounded"),
			g.Attr("data-sidebar-menu-skeleton-text", "true"),
		),
	)

	return html.Div(
		g.Attr("data-sidebar-menu-skeleton", "true"),
		html.Class(lib.CN("flex h-8 items-center gap-2 rounded-md px-2", props.Class)),
		g.Group(skeletonItems),
	)
}

// MenuSub creates a sidebar submenu
func MenuSub(props Props, children ...g.Node) g.Node {
	return html.Ul(
		g.Attr("data-sidebar-menu-sub", "true"),
		html.Class(lib.CN(
			"border-sidebar-border mx-3.5 flex min-w-0 translate-x-px flex-col gap-1 border-l px-2.5 py-0.5",
			"group-data-[collapsible=icon]:hidden",
			props.Class,
		)),
		g.Group(children),
	)
}

// MenuSubItem creates a sidebar submenu item
func MenuSubItem(props Props, children ...g.Node) g.Node {
	return html.Li(
		g.Attr("data-sidebar-menu-sub-item", "true"),
		html.Class(lib.CN("group/menu-sub-item relative", props.Class)),
		g.Group(children),
	)
}

// MenuSubButtonProps defines properties for submenu buttons
type MenuSubButtonProps struct {
	Size     string // "sm" | "md"
	IsActive bool
	Href     string
	Class    string
}

// MenuSubButton creates a sidebar submenu button
func MenuSubButton(props MenuSubButtonProps, children ...g.Node) g.Node {
	// Set defaults
	if props.Size == "" {
		props.Size = "md"
	}

	classes := lib.CN(
		"text-sidebar-foreground ring-sidebar-ring hover:bg-sidebar-accent hover:text-sidebar-accent-foreground active:bg-sidebar-accent active:text-sidebar-accent-foreground [&>svg]:text-sidebar-accent-foreground flex h-7 min-w-0 -translate-x-px items-center gap-2 overflow-hidden rounded-md px-2 outline-none focus-visible:ring-2 disabled:pointer-events-none disabled:opacity-50 aria-disabled:pointer-events-none aria-disabled:opacity-50 [&>span:last-child]:truncate [&>svg]:size-4 [&>svg]:shrink-0",
		"data-[active=true]:bg-sidebar-accent data-[active=true]:text-sidebar-accent-foreground",
		lib.CNIf(props.Size == "sm", "text-xs", "text-sm"),
		"group-data-[collapsible=icon]:hidden",
		props.Class,
	)

	elem := "a"
	attrs := []g.Node{
		g.Attr("data-sidebar-menu-sub-button", "true"),
		g.Attr("data-size", props.Size),
		html.Class(classes),
	}

	if props.Href != "" {
		attrs = append(attrs, html.Href(props.Href))
	}

	if props.IsActive {
		attrs = append(attrs, g.Attr("data-active", "true"))
	}

	return g.El(elem,
		append(attrs, children...)...,
	)
}