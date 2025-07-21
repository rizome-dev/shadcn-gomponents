package resizable

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

// Example demonstrates how to use the Resizable component
func Example() g.Node {
	return html.Div(
		html.Class("space-y-8 p-8"),
		
		// Basic horizontal resizable
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Basic Horizontal Resizable")),
			html.Div(
				html.Class("h-64 max-w-4xl rounded-lg border"),
				PanelGroup(
					Props{Direction: "horizontal"},
					Panel(
						PanelProps{DefaultSize: 50},
						html.Div(
							html.Class("flex h-full items-center justify-center p-6"),
							html.Span(html.Class("font-semibold"), g.Text("Panel 1")),
						),
					),
					Handle(HandleProps{}),
					Panel(
						PanelProps{DefaultSize: 50},
						html.Div(
							html.Class("flex h-full items-center justify-center p-6"),
							html.Span(html.Class("font-semibold"), g.Text("Panel 2")),
						),
					),
				),
			),
		),
		
		// Vertical resizable
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Vertical Resizable")),
			html.Div(
				html.Class("h-96 max-w-md rounded-lg border"),
				VerticalPanelGroup(
					Props{},
					Panel(
						PanelProps{DefaultSize: 50},
						html.Div(
							html.Class("flex h-full items-center justify-center p-6"),
							html.Span(html.Class("font-semibold"), g.Text("Top Panel")),
						),
					),
					Handle(HandleProps{}),
					Panel(
						PanelProps{DefaultSize: 50},
						html.Div(
							html.Class("flex h-full items-center justify-center p-6"),
							html.Span(html.Class("font-semibold"), g.Text("Bottom Panel")),
						),
					),
				),
			),
		),
		
		// Resizable with handle
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Resizable with Visual Handle")),
			html.Div(
				html.Class("h-64 max-w-4xl rounded-lg border"),
				HorizontalPanelGroup(
					Props{},
					Panel(
						PanelProps{DefaultSize: 30},
						html.Div(
							html.Class("flex h-full items-center justify-center p-6 bg-muted"),
							html.Span(html.Class("font-semibold"), g.Text("Sidebar")),
						),
					),
					Handle(HandleProps{WithHandle: true}),
					Panel(
						PanelProps{DefaultSize: 70},
						html.Div(
							html.Class("flex h-full items-center justify-center p-6"),
							html.Span(html.Class("font-semibold"), g.Text("Main Content")),
						),
					),
				),
			),
		),
		
		// Collapsible panels
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Collapsible Panels")),
			html.Div(
				html.Class("h-64 max-w-4xl rounded-lg border"),
				HorizontalPanelGroup(
					Props{},
					CollapsiblePanel(
						PanelProps{
							DefaultSize:   25,
							MinSize:       15,
							MaxSize:       40,
							CollapsedSize: 4,
						},
						html.Div(
							html.Class("flex h-full items-center justify-center p-6 bg-muted"),
							html.Span(html.Class("font-semibold"), g.Text("Collapsible Sidebar")),
						),
					),
					Handle(HandleProps{WithHandle: true}),
					Panel(
						PanelProps{DefaultSize: 75},
						html.Div(
							html.Class("flex h-full items-center justify-center p-6"),
							html.Span(html.Class("font-semibold"), g.Text("Main Content")),
						),
					),
				),
			),
		),
		
		// Three column layout
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Three Column Layout")),
			html.Div(
				html.Class("h-64 max-w-5xl rounded-lg border"),
				ThreeColumnLayout(
					html.Div(
						html.Class("flex h-full items-center justify-center p-6 bg-muted"),
						html.Span(html.Class("font-semibold"), g.Text("Left")),
					),
					html.Div(
						html.Class("flex h-full items-center justify-center p-6"),
						html.Span(html.Class("font-semibold"), g.Text("Center")),
					),
					html.Div(
						html.Class("flex h-full items-center justify-center p-6 bg-muted"),
						html.Span(html.Class("font-semibold"), g.Text("Right")),
					),
				),
			),
		),
		
		// IDE-like layout
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("IDE Layout")),
			html.Div(
				html.Class("h-96 max-w-6xl rounded-lg border overflow-hidden"),
				IDELayout(
					// Sidebar
					html.Div(
						html.Class("h-full bg-muted p-4"),
						html.H4(html.Class("font-semibold mb-2"), g.Text("Explorer")),
						html.Ul(html.Class("space-y-1 text-sm"),
							html.Li(g.Text("📁 src")),
							html.Li(g.Text("📁 components")),
							html.Li(g.Text("📄 main.go")),
							html.Li(g.Text("📄 go.mod")),
						),
					),
					// Editor
					html.Div(
						html.Class("h-full bg-background p-4"),
						html.H4(html.Class("font-semibold mb-2"), g.Text("main.go")),
	html.Pre(html.Class("text-xs font-mono"),
	html.Code(g.Text(`package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}`)),
						),
					),
					// Terminal
					html.Div(
						html.Class("h-full bg-black text-white p-4"),
						html.H4(html.Class("font-semibold mb-2"), g.Text("Terminal")),
	html.Pre(html.Class("text-xs font-mono"),
							g.Text(`$ go run main.go
Hello, World!`),
						),
					),
				),
			),
		),
		
		// Custom min/max sizes
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Custom Min/Max Sizes")),
			html.Div(
				html.Class("h-64 max-w-4xl rounded-lg border"),
				HorizontalPanelGroup(
					Props{},
					Panel(
						PanelProps{
							DefaultSize: 30,
							MinSize:     20,
							MaxSize:     40,
						},
						html.Div(
							html.Class("flex h-full items-center justify-center p-6 bg-muted"),
							html.Div(html.Class("text-center"),
								html.P(html.Class("font-semibold"), g.Text("Constrained Panel")),
								html.P(html.Class("text-sm text-muted-foreground mt-2"), g.Text("Min: 20%, html.Max: 40%")),
							),
						),
					),
					Handle(HandleProps{WithHandle: true}),
					Panel(
						PanelProps{DefaultSize: 70},
						html.Div(
							html.Class("flex h-full items-center justify-center p-6"),
							html.Span(html.Class("font-semibold"), g.Text("Flexible Panel")),
						),
					),
				),
			),
		),
		
		// Nested resizable
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Nested Resizable Panels")),
			html.Div(
				html.Class("h-96 max-w-5xl rounded-lg border"),
				HorizontalPanelGroup(
					Props{},
					Panel(
						PanelProps{DefaultSize: 50},
						html.Div(
							html.Class("flex h-full items-center justify-center p-6 bg-muted"),
							html.Span(html.Class("font-semibold"), g.Text("Left Panel")),
						),
					),
					Handle(HandleProps{WithHandle: true}),
					Panel(
						PanelProps{DefaultSize: 50},
						// Nested vertical panels
						VerticalPanelGroup(
							Props{},
							Panel(
								PanelProps{DefaultSize: 50},
								html.Div(
									html.Class("flex h-full items-center justify-center p-6"),
									html.Span(html.Class("font-semibold"), g.Text("Top Right")),
								),
							),
							Handle(HandleProps{}),
							Panel(
								PanelProps{DefaultSize: 50},
								html.Div(
									html.Class("flex h-full items-center justify-center p-6 bg-muted/50"),
									html.Span(html.Class("font-semibold"), g.Text("Bottom Right")),
								),
							),
						),
					),
				),
			),
		),
		
		// Persistent layout
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Persistent Layout")),
			html.P(html.Class("text-sm text-muted-foreground mb-4"), 
				g.Text("This layout persists its size in localStorage (requires JavaScript)."),
			),
			html.Div(
				html.Class("h-64 max-w-4xl rounded-lg border"),
				PanelGroup(
					Props{
						Direction:  "horizontal",
						Storage:    true,
						StorageKey: "example-layout",
					},
					Panel(
						PanelProps{DefaultSize: 33, ID: "panel-1"},
						html.Div(
							html.Class("flex h-full items-center justify-center p-6 bg-muted"),
							html.Span(html.Class("font-semibold"), g.Text("Panel 1")),
						),
					),
					Handle(HandleProps{WithHandle: true}),
					Panel(
						PanelProps{DefaultSize: 67, ID: "panel-2"},
						html.Div(
							html.Class("flex h-full items-center justify-center p-6"),
							html.Span(html.Class("font-semibold"), g.Text("Panel 2")),
						),
					),
				),
			),
		),
	)
}