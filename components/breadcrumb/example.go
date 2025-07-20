package breadcrumb

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

// DemoBasic shows a basic breadcrumb
func DemoBasic() g.Node {
	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Basic Breadcrumb")),
		Example(),
	)
}

// DemoWithDropdown shows a breadcrumb with ellipsis for collapsed items
func DemoWithDropdown() g.Node {
	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Breadcrumb with Ellipsis")),
		ExampleWithDropdown(),
	)
}

// DemoCustomSeparator shows a breadcrumb with custom separators
func DemoCustomSeparator() g.Node {
	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Custom Separator")),
		ExampleCustomSeparator(),
	)
}

// DemoResponsive shows a responsive breadcrumb that adapts to screen size
func DemoResponsive() g.Node {
	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Responsive Breadcrumb")),
		New(
			Props{},
			BreadcrumbList(
				ListProps{},
				Item(
					ItemProps{Class: "hidden sm:inline-flex"},
					BreadcrumbLink(LinkProps{Href: "/"}, g.Text("Home")),
				),
				Separator(SeparatorProps{Class: "hidden sm:inline-flex"}),
				Item(
					ItemProps{Class: "hidden md:inline-flex"},
					BreadcrumbLink(LinkProps{Href: "/category"}, g.Text("Category")),
				),
				Separator(SeparatorProps{Class: "hidden md:inline-flex"}),
				Item(
					ItemProps{Class: "hidden lg:inline-flex"},
					BreadcrumbLink(LinkProps{Href: "/category/subcategory"}, g.Text("Subcategory")),
				),
				Separator(SeparatorProps{Class: "hidden lg:inline-flex"}),
				Item(
					ItemProps{},
					Page(PageProps{}, g.Text("Product")),
				),
			),
		),
	)
}

// DemoWithIcons shows a breadcrumb with icons
func DemoWithIcons() g.Node {
	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Breadcrumb with Icons")),
		New(
			Props{},
			BreadcrumbList(
				ListProps{},
				Item(
					ItemProps{},
					BreadcrumbLink(
						LinkProps{Href: "/"},
						// Using a simple home icon representation
						g.El("svg",
							g.Attr("class", "h-4 w-4 mr-1"),
							g.Attr("viewBox", "0 0 24 24"),
							g.Attr("fill", "none"),
							g.Attr("stroke", "currentColor"),
							g.Attr("stroke-width", "2"),
							g.El("path", g.Attr("d", "M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z")),
							g.El("polyline", g.Attr("points", "9 22 9 12 15 12 15 22")),
						),
						g.Text("Home"),
					),
				),
				Separator(SeparatorProps{}),
				Item(
					ItemProps{},
					BreadcrumbLink(
						LinkProps{Href: "/settings"},
						// Using a simple settings icon representation
						g.El("svg",
							g.Attr("class", "h-4 w-4 mr-1"),
							g.Attr("viewBox", "0 0 24 24"),
							g.Attr("fill", "none"),
							g.Attr("stroke", "currentColor"),
							g.Attr("stroke-width", "2"),
							g.El("circle", g.Attr("cx", "12"), g.Attr("cy", "12"), g.Attr("r", "3")),
							g.El("path", g.Attr("d", "M12 1v6m0 6v6m11-11h-6m-6 0H1")),
						),
						g.Text("Settings"),
					),
				),
				Separator(SeparatorProps{}),
				Item(
					ItemProps{},
					Page(PageProps{}, g.Text("Profile")),
				),
			),
		),
	)
}

// DemoLongPath shows how to handle long breadcrumb paths
func DemoLongPath() g.Node {
	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Long Path with Ellipsis")),
		New(
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
					Ellipsis(EllipsisProps{}),
				),
				Separator(SeparatorProps{}),
				Item(
					ItemProps{},
					BreadcrumbLink(LinkProps{Href: "/docs/components"}, g.Text("Components")),
				),
				Separator(SeparatorProps{}),
				Item(
					ItemProps{},
					Page(PageProps{}, g.Text("Breadcrumb")),
				),
			),
		),
	)
}

// DemoStyledVariants shows different styling variations
func DemoStyledVariants() g.Node {
	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Styled Variants")),
		html.Div(html.Class("space-y-4"),
			// Default style
			html.Div(
				html.P(html.Class("text-sm text-muted-foreground mb-2"), g.Text("Default")),
				Example(),
			),
			// Dark background
			html.Div(html.Class("bg-slate-900 p-4 rounded-lg"),
				html.P(html.Class("text-sm text-slate-400 mb-2"), g.Text("On Dark Background")),
				New(
					Props{Class: "text-slate-200"},
					BreadcrumbList(
						ListProps{Class: "text-slate-400"},
						Item(
							ItemProps{},
							BreadcrumbLink(LinkProps{Href: "/", Class: "hover:text-slate-100"}, g.Text("Home")),
						),
						Separator(SeparatorProps{Class: "text-slate-600"}),
						Item(
							ItemProps{},
							BreadcrumbLink(LinkProps{Href: "/products", Class: "hover:text-slate-100"}, g.Text("Products")),
						),
						Separator(SeparatorProps{Class: "text-slate-600"}),
						Item(
							ItemProps{},
							Page(PageProps{Class: "text-slate-100"}, g.Text("Details")),
						),
					),
				),
			),
			// With background
			html.Div(
				html.P(html.Class("text-sm text-muted-foreground mb-2"), g.Text("With Background")),
				html.Div(html.Class("bg-muted/50 p-3 rounded-md"),
					Example(),
				),
			),
		),
	)
}