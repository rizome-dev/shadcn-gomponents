package skeleton

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

// Example demonstrates how to use the Skeleton component
func Example() g.Node {
	return html.Div(
		html.Class("p-8 space-y-8"),
		
		// Basic skeleton shapes
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Basic Shapes")),
			html.Div(html.Class("space-y-4"),
				html.Div(
					html.P(html.Class("text-sm text-muted-foreground mb-2"), g.Text("Rectangle")),
					WithClass("h-20 w-full"),
				),
				html.Div(
					html.P(html.Class("text-sm text-muted-foreground mb-2"), g.Text("Square")),
					WithClass("h-20 w-20"),
				),
				html.Div(
					html.P(html.Class("text-sm text-muted-foreground mb-2"), g.Text("Circle")),
					WithClass("h-20 w-20 rounded-full"),
				),
			),
		),
		
		// Text content
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Text Content")),
			html.Div(html.Class("space-y-4"),
				html.Div(
					html.P(html.Class("text-sm text-muted-foreground mb-2"), g.Text("Single line")),
					Text(),
				),
				html.Div(
					html.P(html.Class("text-sm text-muted-foreground mb-2"), g.Text("Multiple lines")),
					TextLines(3),
				),
				html.Div(
					html.P(html.Class("text-sm text-muted-foreground mb-2"), g.Text("Paragraph")),
					TextLines(5),
				),
			),
		),
		
		// Common UI elements
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("UI Elements")),
			html.Div(html.Class("grid grid-cols-2 gap-4"),
				html.Div(
					html.P(html.Class("text-sm text-muted-foreground mb-2"), g.Text("Buttons")),
					html.Div(html.Class("flex gap-2"),
						Button("sm"),
						Button("default"),
						Button("lg"),
						Button("icon"),
					),
				),
				html.Div(
					html.P(html.Class("text-sm text-muted-foreground mb-2"), g.Text("Avatars")),
					html.Div(html.Class("flex gap-2 items-center"),
						Avatar("sm"),
						Avatar("default"),
						Avatar("lg"),
					),
				),
				html.Div(
					html.P(html.Class("text-sm text-muted-foreground mb-2"), g.Text("Input")),
					html.Input(),
				),
				html.Div(
					html.P(html.Class("text-sm text-muted-foreground mb-2"), g.Text("Select")),
					WithClass("h-10 w-full"),
				),
			),
		),
		
		// Card layouts
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Card Layouts")),
			html.Div(html.Class("grid grid-cols-2 gap-4"),
				// Simple card
				html.Div(html.Class("space-y-3 p-4 border rounded-lg"),
					Card(),
					TextLines(2),
					Button("default"),
				),
				// Post card
				html.Div(html.Class("space-y-3 p-4 border rounded-lg"),
					PostCard(),
				),
			),
		),
		
		// Profile layouts
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Profile Layouts")),
			html.Div(html.Class("space-y-4"),
				// Basic profile
				ProfileCard(),
				
				// Detailed profile
				html.Div(html.Class("flex gap-4"),
					Avatar("lg"),
					html.Div(html.Class("flex-1 space-y-3"),
						WithClass("h-6 w-48"), // Name
						TextLines(2),          // Bio
						html.Div(html.Class("flex gap-2"),
							Button("sm"),
							Button("sm"),
						),
					),
				),
			),
		),
		
		// List layouts
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("List Layouts")),
			html.Div(html.Class("grid grid-cols-2 gap-4"),
				html.Div(
					html.P(html.Class("text-sm text-muted-foreground mb-2"), g.Text("Simple list")),
					ListComponent(4),
				),
				html.Div(
					html.P(html.Class("text-sm text-muted-foreground mb-2"), g.Text("Media list")),
					html.Div(html.Class("space-y-4"),
						html.Div(html.Class("flex gap-3"),
							WithClass("h-12 w-12 rounded"),
							html.Div(html.Class("flex-1 space-y-2"),
								WithClass("h-4 w-3/4"),
								WithClass("h-3 w-1/2"),
							),
						),
						html.Div(html.Class("flex gap-3"),
							WithClass("h-12 w-12 rounded"),
							html.Div(html.Class("flex-1 space-y-2"),
								WithClass("h-4 w-2/3"),
								WithClass("h-3 w-1/3"),
							),
						),
					),
				),
			),
		),
		
		// Table
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Table")),
			TableComponent(5),
		),
		
		// Form
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Form")),
			html.Div(html.Class("max-w-md"),
				html.Form(),
			),
		),
		
		// Grid layouts
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Grid Layouts")),
			html.Div(html.Class("space-y-4"),
				html.Div(
					html.P(html.Class("text-sm text-muted-foreground mb-2"), g.Text("2 columns")),
					Grid(2, 4),
				),
				html.Div(
					html.P(html.Class("text-sm text-muted-foreground mb-2"), g.Text("3 columns")),
					Grid(3, 6),
				),
			),
		),
		
		// Real-world example: Blog loading state
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Blog Loading State")),
			html.Div(html.Class("max-w-2xl space-y-8"),
				// Header
				html.Div(html.Class("space-y-2 text-center"),
					WithClass("h-8 w-64 mx-auto"), // Title
					WithClass("h-4 w-96 mx-auto"), // Subtitle
				),
				
				// Article
				html.Article(html.Class("space-y-6"),
					// Meta
					html.Div(html.Class("flex items-center justify-between"),
						ProfileCard(),
						WithClass("h-4 w-24"), // Date
					),
					
					// Featured image
					Image("video"),
					
					// Content
					html.Div(html.Class("space-y-4"),
						TextLines(5),
						TextLines(4),
						TextLines(6),
					),
					
					// Tags
					html.Div(html.Class("flex gap-2"),
						WithClass("h-6 w-16 rounded-full"),
						WithClass("h-6 w-20 rounded-full"),
						WithClass("h-6 w-18 rounded-full"),
					),
				),
			),
		),
		
		// Dashboard loading state
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Dashboard Loading State")),
			html.Div(html.Class("space-y-4"),
				// Stats cards
				html.Div(html.Class("grid grid-cols-4 gap-4"),
					html.Div(html.Class("p-4 border rounded-lg space-y-2"),
						WithClass("h-4 w-20"),
						WithClass("h-8 w-32"),
					),
					html.Div(html.Class("p-4 border rounded-lg space-y-2"),
						WithClass("h-4 w-20"),
						WithClass("h-8 w-28"),
					),
					html.Div(html.Class("p-4 border rounded-lg space-y-2"),
						WithClass("h-4 w-20"),
						WithClass("h-8 w-36"),
					),
					html.Div(html.Class("p-4 border rounded-lg space-y-2"),
						WithClass("h-4 w-20"),
						WithClass("h-8 w-24"),
					),
				),
				
				// Chart area
				html.Div(html.Class("p-4 border rounded-lg"),
					WithClass("h-64 w-full"),
				),
			),
		),
	)
}