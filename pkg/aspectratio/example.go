package aspectratio

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

// Example demonstrates various aspect ratio usage patterns
func Example() g.Node {
	return html.Div(
		html.Class("max-w-6xl mx-auto p-8 space-y-8"),
		html.H2(html.Class("text-2xl font-bold mb-6"), g.Text("Aspect Ratio Examples")),
		
		// Example 1: Image gallery with different ratios
		html.Section(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Image Gallery with Different Ratios")),
			html.Div(html.Class("grid grid-cols-1 md:grid-cols-3 gap-4"),
				// Square (1:1)
				html.Div(
					html.H4(html.Class("text-sm font-medium mb-2"), g.Text("Square (1:1)")),
					Square(
						html.Img(
							html.Class("h-full w-full object-cover rounded-lg"),
							html.Src("https://via.placeholder.com/400x400"),
							html.Alt("Square image"),
						),
					),
				),
				
				// Video (16:9)
				html.Div(
					html.H4(html.Class("text-sm font-medium mb-2"), g.Text("Video (16:9)")),
					Video16x9(
						html.Img(
							html.Class("h-full w-full object-cover rounded-lg"),
							html.Src("https://via.placeholder.com/1920x1080"),
							html.Alt("16:9 video ratio"),
						),
					),
				),
				
				// Portrait (4:5)
				html.Div(
					html.H4(html.Class("text-sm font-medium mb-2"), g.Text("Portrait (4:5)")),
					Portrait(
						html.Img(
							html.Class("h-full w-full object-cover rounded-lg"),
							html.Src("https://via.placeholder.com/800x1000"),
							html.Alt("Portrait image"),
						),
					),
				),
			),
		),
		
		// Example 2: Video player placeholder
		html.Section(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Video Player Placeholder")),
			html.Div(html.Class("max-w-2xl"),
				WithClass(16.0/9.0, "bg-muted rounded-lg",
					html.Div(
						html.Class("flex items-center justify-center h-full"),
						html.Button(
							html.Class("bg-primary text-primary-foreground rounded-full p-4"),
							g.Raw(`<svg class="w-8 h-8" fill="currentColor" viewBox="0 0 20 20"><path d="M10 18a8 8 0 100-16 8 8 0 000 16zM9.555 7.168A1 1 0 008 8v4a1 1 0 001.555.832l3-2a1 1 0 000-1.664l-3-2z"/></svg>`),
						),
					),
				),
			),
		),
		
		// Example 3: Content cards with consistent ratios
		html.Section(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Content Cards")),
			html.Div(html.Class("grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6"),
				// Card 1
				html.Div(html.Class("border rounded-lg overflow-hidden"),
					Landscape(
						html.Div(html.Class("bg-gradient-to-br from-purple-500 to-pink-500")),
					),
					html.Div(html.Class("p-4"),
						html.H4(html.Class("font-semibold"), g.Text("Landscape Card")),
						html.P(html.Class("text-sm text-muted-foreground mt-1"), g.Text("3:2 aspect ratio for landscape content")),
					),
				),
				
				// Card 2
				html.Div(html.Class("border rounded-lg overflow-hidden"),
					Landscape(
						html.Div(html.Class("bg-gradient-to-br from-blue-500 to-teal-500")),
					),
					html.Div(html.Class("p-4"),
						html.H4(html.Class("font-semibold"), g.Text("Another Landscape")),
						html.P(html.Class("text-sm text-muted-foreground mt-1"), g.Text("Consistent aspect ratios")),
					),
				),
				
				// Card 3
				html.Div(html.Class("border rounded-lg overflow-hidden"),
					Landscape(
						html.Div(html.Class("bg-gradient-to-br from-orange-500 to-red-500")),
					),
					html.Div(html.Class("p-4"),
						html.H4(html.Class("font-semibold"), g.Text("Third Card")),
						html.P(html.Class("text-sm text-muted-foreground mt-1"), g.Text("All cards align perfectly")),
					),
				),
			),
		),
		
		// Example 4: Custom ratios
		html.Section(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Custom Aspect Ratios")),
			html.Div(html.Class("grid grid-cols-2 md:grid-cols-4 gap-4"),
				// Ultra-wide
				html.Div(
					html.H4(html.Class("text-sm font-medium mb-2"), g.Text("21:9 Ultra-wide")),
					WithClass(21.0/9.0, "bg-muted rounded",
						html.Div(html.Class("flex items-center justify-center h-full text-xs text-muted-foreground"),
							g.Text("21:9"),
						),
					),
				),
				
				// Cinema
				html.Div(
					html.H4(html.Class("text-sm font-medium mb-2"), g.Text("2.39:1 Cinema")),
					Cinema(
						html.Div(html.Class("bg-muted rounded flex items-center justify-center h-full text-xs text-muted-foreground"),
							g.Text("2.39:1"),
						),
					),
				),
				
				// Golden ratio
				html.Div(
					html.H4(html.Class("text-sm font-medium mb-2"), g.Text("Golden Ratio")),
					WithClass(1.618, "bg-muted rounded",
						html.Div(html.Class("flex items-center justify-center h-full text-xs text-muted-foreground"),
							g.Text("1.618:1"),
						),
					),
				),
				
				// Vertical
				html.Div(
					html.H4(html.Class("text-sm font-medium mb-2"), g.Text("9:16 Vertical")),
					WithClass(9.0/16.0, "bg-muted rounded",
						html.Div(html.Class("flex items-center justify-center h-full text-xs text-muted-foreground"),
							g.Text("9:16"),
						),
					),
				),
			),
		),
		
		// Example 5: Responsive aspect ratio container
		html.Section(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Responsive Container")),
			html.P(html.Class("text-muted-foreground mb-4"), g.Text("The aspect ratio container maintains its ratio regardless of width")),
			html.Div(html.Class("space-y-4"),
				// Full width
				html.Div(
					html.P(html.Class("text-sm text-muted-foreground mb-2"), g.Text("Full width")),
					Video16x9(
						html.Div(html.Class("bg-primary/10 border-2 border-dashed border-primary/30 rounded-lg flex items-center justify-center"),
							g.Text("16:9 Content"),
						),
					),
				),
				
				// Half width
				html.Div(html.Class("w-1/2"),
					html.P(html.Class("text-sm text-muted-foreground mb-2"), g.Text("Half width")),
					Video16x9(
						html.Div(html.Class("bg-primary/10 border-2 border-dashed border-primary/30 rounded-lg flex items-center justify-center"),
							g.Text("16:9 Content"),
						),
					),
				),
				
				// Quarter width
				html.Div(html.Class("w-1/4"),
					html.P(html.Class("text-sm text-muted-foreground mb-2"), g.Text("Quarter width")),
					Video16x9(
						html.Div(html.Class("bg-primary/10 border-2 border-dashed border-primary/30 rounded-lg flex items-center justify-center text-sm"),
							g.Text("16:9"),
						),
					),
				),
			),
		),
	)
}