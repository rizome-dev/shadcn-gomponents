package alert

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

// Example demonstrates how to use the Alert component
func Example() g.Node {
	return html.Div(
		html.Class("space-y-4 p-8"),
		// Default alert
		Default(
			Title(g.Text("Heads up!")),
			Description(g.Text("You can add components to your app using the cli.")),
		),
		
		// Destructive alert
		Destructive(
			Title(g.Text("Error")),
			Description(g.Text("Your session has expired. Please log in again.")),
		),
		
		// Alert with icon (using a simple SVG)
		WithIcon(
			// Info icon SVG
			g.El("svg",
				g.Attr("xmlns", "http://www.w3.org/2000/svg"),
				g.Attr("viewBox", "0 0 24 24"),
				g.Attr("fill", "none"),
				g.Attr("stroke", "currentColor"),
				g.Attr("stroke-width", "2"),
				g.Attr("stroke-linecap", "round"),
				g.Attr("stroke-linejoin", "round"),
				html.Class("h-4 w-4"),
				g.El("circle", g.Attr("cx", "12"), g.Attr("cy", "12"), g.Attr("r", "10")),
				g.El("line", g.Attr("x1", "12"), g.Attr("y1", "16"), g.Attr("x2", "12"), g.Attr("y2", "12")),
				g.El("line", g.Attr("x1", "12"), g.Attr("y1", "8"), g.Attr("x2", "12.01"), g.Attr("y2", "8")),
			),
			Props{},
			Title(g.Text("Information")),
			Description(g.Text("This is an informational alert with an icon.")),
		),
		
		// Custom styled alert
		New(
			Props{
				Class: "border-blue-500 bg-blue-50 text-blue-900",
			},
			Title(g.Text("Custom Alert")),
			Description(g.Text("This alert has custom styling applied.")),
		),
	)
}