package progress

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

// Example demonstrates how to use the Progress component
func Example() g.Node {
	return html.Div(
		html.Class("p-8 space-y-8"),
		
		// Basic progress bars
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Basic Progress")),
			html.Div(html.Class("space-y-4"),
				Default(0),
				Default(25),
				Default(50),
				Default(75),
				Default(100),
			),
		),
		
		// Different sizes
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Different Sizes")),
			html.Div(html.Class("space-y-4"),
				html.Div(html.Class("space-y-2"),
					html.P(html.Class("text-sm text-muted-foreground"), g.Text("Small")),
					SmallComponent(60),
				),
				html.Div(html.Class("space-y-2"),
					html.P(html.Class("text-sm text-muted-foreground"), g.Text("Default")),
					Default(60),
				),
				html.Div(html.Class("space-y-2"),
					html.P(html.Class("text-sm text-muted-foreground"), g.Text("Large")),
					Large(60),
				),
			),
		),
		
		// With labels
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("With Labels")),
			html.Div(html.Class("space-y-6"),
				WithLabel(33, "Uploading files"),
				WithLabel(67, "Processing"),
				WithLabel(100, "Complete"),
			),
		),
		
		// Indeterminate state
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Indeterminate")),
			html.P(html.Class("text-sm text-muted-foreground mb-4"), g.Text("Use when the duration cannot be determined")),
			Indeterminate(),
		),
		
		// Striped progress
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Striped Pattern")),
			html.Div(html.Class("space-y-4"),
				Striped(45),
				Striped(80),
			),
		),
		
		// Multi-segment progress
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Multi-segment Progress")),
			html.P(html.Class("text-sm text-muted-foreground mb-4"), g.Text("Useful for showing multiple categories")),
			html.Div(html.Class("space-y-4"),
				Multi([]Segment{
					{Value: 30, Color: "bg-blue-500"},
					{Value: 25, Color: "bg-green-500"},
					{Value: 20, Color: "bg-yellow-500"},
					{Value: 15, Color: "bg-red-500"},
				}),
				// With legend
				html.Div(html.Class("mt-4 flex flex-wrap gap-4 text-sm"),
					html.Div(html.Class("flex items-center gap-2"),
						html.Div(html.Class("h-3 w-3 rounded-full bg-blue-500")),
						html.Span(g.Text("Category A (30%)")),
					),
					html.Div(html.Class("flex items-center gap-2"),
						html.Div(html.Class("h-3 w-3 rounded-full bg-green-500")),
						html.Span(g.Text("Category B (25%)")),
					),
					html.Div(html.Class("flex items-center gap-2"),
						html.Div(html.Class("h-3 w-3 rounded-full bg-yellow-500")),
						html.Span(g.Text("Category C (20%)")),
					),
					html.Div(html.Class("flex items-center gap-2"),
						html.Div(html.Class("h-3 w-3 rounded-full bg-red-500")),
						html.Span(g.Text("Category D (15%)")),
					),
				),
			),
		),
		
		// Circular progress
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Circular Progress")),
			html.Div(html.Class("flex items-center gap-8"),
				html.Div(html.Class("text-center"),
					Circular(25, "sm"),
					html.P(html.Class("text-sm text-muted-foreground mt-2"), g.Text("Small")),
				),
				html.Div(html.Class("text-center"),
					Circular(50, "default"),
					html.P(html.Class("text-sm text-muted-foreground mt-2"), g.Text("Default")),
				),
				html.Div(html.Class("text-center"),
					Circular(75, "lg"),
					html.P(html.Class("text-sm text-muted-foreground mt-2"), g.Text("Large")),
				),
			),
		),
		
		// Custom styled progress
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Custom Styles")),
			html.Div(html.Class("space-y-4"),
				New(Props{
					Value: 60,
					Class: "w-1/2", // Half width
				}),
				html.Div(
					// Gradient background
					html.Class("relative h-4 w-full overflow-hidden rounded-full bg-gradient-to-r from-blue-100 to-blue-200"),
					html.Div(
						html.Class("h-full w-full flex-1 bg-gradient-to-r from-blue-500 to-blue-600 transition-all"),
						html.Style("transform: translateX(-40%)"),
					),
				),
			),
		),
		
		// Real-world examples
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Real-world Examples")),
			
			// File upload
			html.Div(html.Class("rounded-lg border bg-card p-4 space-y-3"),
				html.Div(html.Class("flex items-center justify-between"),
					html.Div(html.Class("flex items-center gap-3"),
						// File icon
						g.El("svg",
							g.Attr("xmlns", "http://www.w3.org/2000/svg"),
							g.Attr("viewBox", "0 0 24 24"),
							g.Attr("fill", "none"),
							g.Attr("stroke", "currentColor"),
							g.Attr("stroke-width", "2"),
							html.Class("h-8 w-8 text-muted-foreground"),
							g.El("path", g.Attr("d", "M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z")),
							g.El("polyline", g.Attr("points", "14 2 14 8 20 8")),
						),
						html.Div(
							html.H4(html.Class("font-medium"), g.Text("document.pdf")),
							html.P(html.Class("text-sm text-muted-foreground"), g.Text("2.4 MB")),
						),
					),
					html.Span(html.Class("text-sm font-medium"), g.Text("75%")),
				),
				SmallComponent(75),
			),
			
			// Storage usage
			html.Div(html.Class("mt-4 rounded-lg border bg-card p-4 space-y-3"),
				html.H4(html.Class("font-medium"), g.Text("Storage Usage")),
				Multi([]Segment{
					{Value: 40, Color: "bg-blue-500"},
					{Value: 30, Color: "bg-purple-500"},
					{Value: 20, Color: "bg-orange-500"},
				}),
				html.Div(html.Class("mt-3 space-y-1 text-sm"),
					html.Div(html.Class("flex items-center justify-between"),
						html.Span(html.Class("flex items-center gap-2"),
							html.Div(html.Class("h-2 w-2 rounded-full bg-blue-500")),
							g.Text("Documents"),
						),
						html.Span(html.Class("text-muted-foreground"), g.Text("4.2 GB")),
					),
					html.Div(html.Class("flex items-center justify-between"),
						html.Span(html.Class("flex items-center gap-2"),
							html.Div(html.Class("h-2 w-2 rounded-full bg-purple-500")),
							g.Text("Media"),
						),
						html.Span(html.Class("text-muted-foreground"), g.Text("3.1 GB")),
					),
					html.Div(html.Class("flex items-center justify-between"),
						html.Span(html.Class("flex items-center gap-2"),
							html.Div(html.Class("h-2 w-2 rounded-full bg-orange-500")),
							g.Text("Other"),
						),
						html.Span(html.Class("text-muted-foreground"), g.Text("2.1 GB")),
					),
				),
				html.P(html.Class("text-sm text-muted-foreground mt-2"), g.Text("9.4 GB of 10 GB used")),
			),
		),
	)
}