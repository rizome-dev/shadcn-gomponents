package separator

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

// Example demonstrates how to use the Separator component
func Example() g.Node {
	return html.Div(
		html.Class("p-8"),
		
		// Horizontal separator example
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Horizontal Separator")),
			html.Div(html.Class("space-y-1"),
				html.P(html.Class("text-sm text-muted-foreground"), g.Text("Content above separator")),
				Horizontal(),
				html.P(html.Class("text-sm text-muted-foreground"), g.Text("Content below separator")),
			),
		),
		
		// Vertical separator example
		html.Div(html.Class("mt-8"),
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Vertical Separator")),
			html.Div(html.Class("flex h-5 items-center space-x-4 text-sm"),
				html.Span(g.Text("Blog")),
				Vertical(),
				html.Span(g.Text("Docs")),
				Vertical(),
				html.Span(g.Text("Source")),
			),
		),
		
		// Semantic separator example
		html.Div(html.Class("mt-8"),
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Semantic Separator (for screen readers)")),
	html.Article(
				html.H3(html.Class("text-lg font-semibold"), g.Text("Article Title")),
				html.P(html.Class("text-sm text-muted-foreground mt-2"), g.Text("This is the article content.")),
				Semantic(), // This separator will be announced by screen readers
				html.H3(html.Class("text-lg font-semibold mt-4"), g.Text("Related Articles")),
				html.P(html.Class("text-sm text-muted-foreground mt-2"), g.Text("Links to related content would go here.")),
			),
		),
		
		// Custom styled separator
		html.Div(html.Class("mt-8"),
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Custom Styled Separator")),
			WithClass("bg-blue-500 h-1"),
			html.P(html.Class("text-sm text-muted-foreground mt-4"), g.Text("This separator has custom styling")),
		),
		
		// Card-like layout with separators
		html.Div(html.Class("mt-8 rounded-lg border bg-card p-6"),
			html.Div(html.Class("flex items-center justify-between"),
				html.H3(html.Class("text-lg font-semibold"), g.Text("Account Settings")),
				html.Button(html.Class("text-sm"), g.Text("Save")),
			),
			New(Props{Class: "my-4"}),
			html.Div(html.Class("space-y-4"),
				html.Div(
					html.Label(html.Class("text-sm font-medium"), g.Text("Username")),
					html.P(html.Class("text-sm text-muted-foreground"), g.Text("@johndoe")),
				),
				Horizontal(),
				html.Div(
					html.Label(html.Class("text-sm font-medium"), g.Text("Email")),
					html.P(html.Class("text-sm text-muted-foreground"), g.Text("john@example.com")),
				),
			),
		),
	)
}