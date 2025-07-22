package popover

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib/icons"
)

// Example demonstrates how to use the Popover component
func Example() g.Node {
	return html.Div(
		html.Class("space-y-8 p-8"),
		
		// Basic popover
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Basic Popover")),
			New(
				Props{},
				Trigger(TriggerProps{}, g.Text("Open popover")),
	ContentComponent(
					ContentProps{},
					html.Div(
						html.H4(html.Class("font-medium leading-none"), g.Text("Popover Content")),
						html.P(html.Class("text-sm text-muted-foreground mt-2"), 
							g.Text("This is a basic popover with some content."),
						),
					),
				),
			),
		),
		
		// Popover with different positions
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Popover Positions")),
			html.Div(html.Class("flex flex-wrap gap-4"),
				// Top
				New(
					Props{},
					Trigger(TriggerProps{}, g.Text("Top")),
	ContentComponent(
						ContentProps{Side: "top", Class: "w-40"},
						html.P(html.Class("text-sm"), g.Text("Popover on top")),
					),
				),
				
				// Right
				New(
					Props{},
					Trigger(TriggerProps{}, g.Text("Right")),
	ContentComponent(
						ContentProps{Side: "right", Class: "w-40"},
						html.P(html.Class("text-sm"), g.Text("Popover on right")),
					),
				),
				
				// Bottom
				New(
					Props{},
					Trigger(TriggerProps{}, g.Text("Bottom")),
	ContentComponent(
						ContentProps{Side: "bottom", Class: "w-40"},
						html.P(html.Class("text-sm"), g.Text("Popover on bottom")),
					),
				),
				
				// Left
				New(
					Props{},
					Trigger(TriggerProps{}, g.Text("Left")),
	ContentComponent(
						ContentProps{Side: "left", Class: "w-40"},
						html.P(html.Class("text-sm"), g.Text("Popover on left")),
					),
				),
			),
		),
		
		// Popover with alignments
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Popover Alignments")),
			html.Div(html.Class("flex flex-wrap gap-4"),
				// Start alignment
				New(
					Props{},
					Trigger(TriggerProps{}, g.Text("Align Start")),
	ContentComponent(
						ContentProps{Align: "start"},
						html.P(html.Class("text-sm"), g.Text("Aligned to start")),
					),
				),
				
				// Center alignment
				New(
					Props{},
					Trigger(TriggerProps{}, g.Text("Align Center")),
	ContentComponent(
						ContentProps{Align: "center"},
						html.P(html.Class("text-sm"), g.Text("Aligned to center")),
					),
				),
				
				// End alignment
				New(
					Props{},
					Trigger(TriggerProps{}, g.Text("Align End")),
	ContentComponent(
						ContentProps{Align: "end"},
						html.P(html.Class("text-sm"), g.Text("Aligned to end")),
					),
				),
			),
		),
		
		// Popover with arrow
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Popover with Arrow")),
			New(
				Props{},
				Trigger(TriggerProps{}, g.Text("With Arrow")),
				WithArrow(
					ContentProps{},
					html.H4(html.Class("font-medium"), g.Text("Popover with Arrow")),
					html.P(html.Class("text-sm text-muted-foreground mt-2"), 
						g.Text("This popover has an arrow pointing to the trigger."),
					),
				),
			),
		),
		
		// Popover with close button
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Popover with Close Button")),
			New(
				Props{},
				Trigger(TriggerProps{}, g.Text("With Close")),
	ContentComponent(
					ContentProps{},
					Close(),
					html.H4(html.Class("font-medium"), g.Text("Closable Popover")),
					html.P(html.Class("text-sm text-muted-foreground mt-2"), 
						g.Text("This popover has a close button."),
					),
				),
			),
		),
		
		// Complex popover content
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Complex Popover")),
			New(
				Props{},
				Trigger(
					TriggerProps{Class: "inline-flex items-center gap-2"},
					icons.Settings(html.Class("h-4 w-4")),
					g.Text("Settings"),
				),
	ContentComponent(
					ContentProps{Class: "w-80"},
					html.Div(html.Class("grid gap-4"),
						html.Div(html.Class("space-y-2"),
							html.H4(html.Class("font-medium leading-none"), g.Text("Dimensions")),
							html.P(html.Class("text-sm text-muted-foreground"), 
								g.Text("Set the dimensions for the layer."),
							),
						),
						html.Div(html.Class("grid gap-2"),
							html.Div(html.Class("grid grid-cols-3 items-center gap-4"),
								html.Label(html.For("width"), g.Text("Width")),
								html.Input(
									html.ID("width"),
									html.Value("100%"),
									html.Class("col-span-2 h-8"),
								),
							),
							html.Div(html.Class("grid grid-cols-3 items-center gap-4"),
								html.Label(html.For("maxWidth"), g.Text("Max. width")),
								html.Input(
									html.ID("maxWidth"),
									html.Value("300px"),
									html.Class("col-span-2 h-8"),
								),
							),
							html.Div(html.Class("grid grid-cols-3 items-center gap-4"),
								html.Label(html.For("height"), g.Text("Height")),
								html.Input(
									html.ID("height"),
									html.Value("25px"),
									html.Class("col-span-2 h-8"),
								),
							),
							html.Div(html.Class("grid grid-cols-3 items-center gap-4"),
								html.Label(html.For("maxHeight"), g.Text("Max. height")),
								html.Input(
									html.ID("maxHeight"),
									html.Value("none"),
									html.Class("col-span-2 h-8"),
								),
							),
						),
					),
				),
			),
		),
		
		// HTMX Example
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("HTMX-Enhanced Popover")),
			html.P(html.Class("text-sm text-muted-foreground mb-4"), 
				g.Text("These examples require HTMX and server endpoints to be set up."),
			),
			html.Div(html.Class("flex flex-wrap gap-4"),
				// Basic HTMX popover
				// ExampleHTMX() // TODO: Implement,
				
				// Menu popover
				html.Div(
					html.ID("popover-menu"),
					TriggerHTMX(
						TriggerProps{Class: "inline-flex items-center gap-2"},
						HTMXProps{
							ID:         "popover-menu",
							TogglePath: "/api/popover/menu/toggle",
							ClosePath:  "/api/popover/menu/close",
						},
						icons.MoreVertical(html.Class("h-4 w-4")),
						g.Text("Options"),
					),
				),
			),
		),
	)
}

// ExampleWithMenu demonstrates a popover with menu items
func ExampleWithMenu() g.Node {
	return New(
		Props{},
		Trigger(
			TriggerProps{Class: "inline-flex items-center gap-2"},
			icons.MoreHorizontal(html.Class("h-4 w-4")),
			g.Text("Actions"),
		),
	ContentComponent(
			ContentProps{Side: "bottom", Align: "end", Class: "w-56 p-1"},
			html.Button(
				html.Type("button"),
				html.Class("flex w-full items-center rounded-sm px-2 py-1.5 text-sm hover:bg-accent hover:text-accent-foreground"),
				icons.Edit(html.Class("mr-2 h-4 w-4")),
				g.Text("Edit"),
			),
			html.Button(
				html.Type("button"),
				html.Class("flex w-full items-center rounded-sm px-2 py-1.5 text-sm hover:bg-accent hover:text-accent-foreground"),
				icons.Copy(html.Class("mr-2 h-4 w-4")),
				g.Text("Duplicate"),
			),
			html.Div(html.Class("h-px bg-border my-1")),
			html.Button(
				html.Type("button"),
				html.Class("flex w-full items-center rounded-sm px-2 py-1.5 text-sm hover:bg-accent hover:text-accent-foreground"),
				icons.Archive(html.Class("mr-2 h-4 w-4")),
				g.Text("Archive"),
			),
			html.Button(
				html.Type("button"),
				html.Class("flex w-full items-center rounded-sm px-2 py-1.5 text-sm hover:bg-accent hover:text-accent-foreground text-destructive"),
				icons.Trash(html.Class("mr-2 h-4 w-4")),
				g.Text("Delete"),
			),
		),
	)
}

// ExampleAsChild demonstrates using AsChild for custom triggers
func ExampleAsChild() g.Node {
	return New(
		Props{},
		Trigger(
			TriggerProps{AsChild: true},
			html.Div(
				html.Class("inline-flex items-center justify-center rounded-md bg-primary px-4 py-2 text-primary-foreground hover:bg-primary/90"),
				g.Text("Custom Trigger"),
			),
		),
	ContentComponent(
			ContentProps{},
			html.P(html.Class("text-sm"), g.Text("This popover uses a custom trigger element.")),
		),
	)
}