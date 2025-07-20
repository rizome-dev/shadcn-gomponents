package badge

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

// Example demonstrates how to use the Badge component
func Example() g.Node {
	return html.Div(
		html.Class("p-8 space-y-8"),
		
		// Basic variants
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Badge Variants")),
			html.Div(html.Class("flex flex-wrap gap-2"),
				Default(g.Text("Default")),
				Secondary(g.Text("Secondary")),
				Destructive(g.Text("Destructive")),
				Outline(g.Text("Outline")),
			),
		),
		
		// Badges with different content
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Different Content")),
			html.Div(html.Class("flex flex-wrap gap-2"),
				Default(g.Text("New")),
				Secondary(g.Text("v2.0.0")),
				Destructive(g.Text("Deprecated")),
				Outline(g.Text("Beta")),
				Default(g.Text("Coming Soon")),
			),
		),
		
		// Badges with icons
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("With Icons")),
			html.Div(html.Class("flex flex-wrap gap-2"),
				WithIcon(
					// Check icon
					g.El("svg",
						g.Attr("xmlns", "http://www.w3.org/2000/svg"),
						g.Attr("viewBox", "0 0 24 24"),
						g.Attr("fill", "none"),
						g.Attr("stroke", "currentColor"),
						g.Attr("stroke-width", "2"),
						g.El("polyline", g.Attr("points", "20 6 9 17 4 12")),
					),
					"default",
					"Verified",
				),
				WithIcon(
					// X icon
					g.El("svg",
						g.Attr("xmlns", "http://www.w3.org/2000/svg"),
						g.Attr("viewBox", "0 0 24 24"),
						g.Attr("fill", "none"),
						g.Attr("stroke", "currentColor"),
						g.Attr("stroke-width", "2"),
						g.El("line", g.Attr("x1", "18"), g.Attr("y1", "6"), g.Attr("x2", "6"), g.Attr("y2", "18")),
						g.El("line", g.Attr("x1", "6"), g.Attr("y1", "6"), g.Attr("x2", "18"), g.Attr("y2", "18")),
					),
					"destructive",
					"Failed",
				),
			),
		),
		
		// Badges as links
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("As Links")),
			html.Div(html.Class("flex flex-wrap gap-2"),
				LinkComponent("/docs", Props{}, g.Text("Documentation")),
				LinkComponent("/api", Props{Variant: "secondary"}, g.Text("API Reference")),
				LinkComponent("/changelog", Props{Variant: "outline"}, g.Text("Changelog")),
			),
		),
		
		// Badges in context
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("In Context")),
			
			// Card with badges
			html.Div(html.Class("rounded-lg border bg-card p-6 space-y-3"),
				html.Div(html.Class("flex items-start justify-between"),
					html.Div(
						html.H3(html.Class("text-lg font-semibold"), g.Text("Product Name")),
						html.P(html.Class("text-sm text-muted-foreground mt-1"), g.Text("A brief description of the product")),
					),
					html.Div(html.Class("flex gap-2"),
						Default(g.Text("New")),
						Secondary(g.Text("Popular")),
					),
				),
			),
			
			// List with badges
			html.Div(html.Class("mt-4 space-y-2"),
				html.Div(html.Class("flex items-center justify-between p-3 rounded-md border"),
					html.Span(html.Class("text-sm"), g.Text("Feature A")),
					Default(g.Text("Stable")),
				),
				html.Div(html.Class("flex items-center justify-between p-3 rounded-md border"),
					html.Span(html.Class("text-sm"), g.Text("Feature B")),
					Secondary(g.Text("Preview")),
				),
				html.Div(html.Class("flex items-center justify-between p-3 rounded-md border"),
					html.Span(html.Class("text-sm"), g.Text("Feature C")),
					Destructive(g.Text("Deprecated")),
				),
			),
		),
		
		// Custom styled badges
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Custom Styles")),
			html.Div(html.Class("flex flex-wrap gap-2"),
				New(Props{Class: "bg-purple-500 text-white border-purple-500"}, g.Text("Purple")),
				New(Props{Class: "bg-gradient-to-r from-blue-500 to-purple-500 text-white border-transparent"}, g.Text("Gradient")),
				New(Props{Class: "text-xs uppercase tracking-wider"}, g.Text("Uppercase")),
			),
		),
	)
}