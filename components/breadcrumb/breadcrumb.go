package breadcrumb

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
	"github.com/rizome-dev/shadcn-gomponents/lib/icons"
)

// Props defines the properties for the Breadcrumb component
type Props struct {
	Class string // Additional custom classes
}

// ListProps defines the properties for the BreadcrumbList
type ListProps struct {
	Class string
}

// ItemProps defines the properties for the BreadcrumbItem
type ItemProps struct {
	Class string
}

// LinkProps defines the properties for the BreadcrumbLink
type LinkProps struct {
	Href  string
	Class string
}

// PageProps defines the properties for the BreadcrumbPage
type PageProps struct {
	Class string
}

// SeparatorProps defines the properties for the BreadcrumbSeparator
type SeparatorProps struct {
	Class string
}

// EllipsisProps defines the properties for the BreadcrumbEllipsis
type EllipsisProps struct {
	Class string
}

// New creates a new Breadcrumb navigation component
func New(props Props, children ...g.Node) g.Node {
	classes := lib.CN(
		props.Class,
	)

	return html.Nav(
		append([]g.Node{
			g.Attr("aria-label", "breadcrumb"),
			g.If(classes != "", html.Class(classes)),
		}, children...)...,
	)
}

// BreadcrumbList creates a BreadcrumbList component
func BreadcrumbList(props ListProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"flex flex-wrap items-center gap-1.5 break-words text-sm text-muted-foreground sm:gap-2.5",
		props.Class,
	)

	return html.Ol(
		append([]g.Node{html.Class(classes)}, children...)...,
	)
}

// Item creates a BreadcrumbItem component
func Item(props ItemProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"inline-flex items-center gap-1.5",
		props.Class,
	)

	return html.Li(
		append([]g.Node{html.Class(classes)}, children...)...,
	)
}

// BreadcrumbLink creates a BreadcrumbLink component
func BreadcrumbLink(props LinkProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"transition-colors hover:text-foreground",
		props.Class,
	)

	return html.A(
		append([]g.Node{
			html.Href(props.Href),
			html.Class(classes),
		}, children...)...,
	)
}

// Page creates a BreadcrumbPage component (current page, non-clickable)
func Page(props PageProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"font-normal text-foreground",
		props.Class,
	)

	return html.Span(
		append([]g.Node{
			html.Role("link"),
			g.Attr("aria-disabled", "true"),
			g.Attr("aria-current", "page"),
			html.Class(classes),
		}, children...)...,
	)
}

// Separator creates a BreadcrumbSeparator component
func Separator(props SeparatorProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"[&>svg]:w-3.5 [&>svg]:h-3.5",
		props.Class,
	)

	// Default to ChevronRight icon if no children provided
	var content []g.Node
	if len(children) == 0 {
		content = []g.Node{icons.ChevronRight()}
	} else {
		content = children
	}

	return html.Li(
		append([]g.Node{
			html.Role("presentation"),
			g.Attr("aria-hidden", "true"),
			html.Class(classes),
		}, content...)...,
	)
}

// Ellipsis creates a BreadcrumbEllipsis component
func Ellipsis(props EllipsisProps) g.Node {
	classes := lib.CN(
		"flex h-9 w-9 items-center justify-center",
		props.Class,
	)

	return html.Span(
		html.Role("presentation"),
		g.Attr("aria-hidden", "true"),
		html.Class(classes),
		icons.MoreHorizontal(html.Class("h-4 w-4")),
		html.Span(html.Class("sr-only"), g.Text("More")),
	)
}

// Example creates a basic breadcrumb example
func Example() g.Node {
	return New(
		Props{},
		BreadcrumbList(
			ListProps{},
			Item(
				ItemProps{},
				BreadcrumbLink(LinkProps{Href: "/"}, g.Text("Home")),
			),
			Separator(SeparatorProps{}),
			Item(
				ItemProps{},
				BreadcrumbLink(LinkProps{Href: "/docs"}, g.Text("Documentation")),
			),
			Separator(SeparatorProps{}),
			Item(
				ItemProps{},
				Page(PageProps{}, g.Text("Components")),
			),
		),
	)
}

// ExampleWithDropdown creates a breadcrumb with dropdown example
func ExampleWithDropdown() g.Node {
	return New(
		Props{},
		BreadcrumbList(
			ListProps{},
			Item(
				ItemProps{},
				BreadcrumbLink(LinkProps{Href: "/"}, g.Text("Home")),
			),
			Separator(SeparatorProps{}),
			Item(
				ItemProps{},
				Ellipsis(EllipsisProps{}),
			),
			Separator(SeparatorProps{}),
			Item(
				ItemProps{},
				Page(PageProps{}, g.Text("Current Page")),
			),
		),
	)
}

// ExampleCustomSeparator creates a breadcrumb with custom separator
func ExampleCustomSeparator() g.Node {
	return New(
		Props{},
		BreadcrumbList(
			ListProps{},
			Item(
				ItemProps{},
				BreadcrumbLink(LinkProps{Href: "/"}, g.Text("Home")),
			),
			Separator(SeparatorProps{}, g.Text("/")),
			Item(
				ItemProps{},
				BreadcrumbLink(LinkProps{Href: "/products"}, g.Text("Products")),
			),
			Separator(SeparatorProps{}, g.Text("/")),
			Item(
				ItemProps{},
				Page(PageProps{}, g.Text("Electronics")),
			),
		),
	)
}