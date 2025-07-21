package toggle

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

// Example demonstrates how to use the Toggle component
func Example() g.Node {
	// Helper function to create simple icons
	boldIcon := func() g.Node {
		return g.El("svg",
			g.Attr("xmlns", "http://www.w3.org/2000/svg"),
			g.Attr("viewBox", "0 0 24 24"),
			g.Attr("fill", "none"),
			g.Attr("stroke", "currentColor"),
			g.Attr("stroke-width", "2"),
			html.Class("h-4 w-4"),
			g.El("path", g.Attr("d", "M6 4h8a4 4 0 0 1 4 4 4 4 0 0 1-4 4H6z")),
			g.El("path", g.Attr("d", "M6 12h9a4 4 0 0 1 4 4 4 4 0 0 1-4 4H6z")),
		)
	}

	italicIcon := func() g.Node {
		return g.El("svg",
			g.Attr("xmlns", "http://www.w3.org/2000/svg"),
			g.Attr("viewBox", "0 0 24 24"),
			g.Attr("fill", "none"),
			g.Attr("stroke", "currentColor"),
			g.Attr("stroke-width", "2"),
			html.Class("h-4 w-4"),
			g.El("line", g.Attr("x1", "19"), g.Attr("y1", "4"), g.Attr("x2", "10"), g.Attr("y2", "4")),
			g.El("line", g.Attr("x1", "14"), g.Attr("y1", "20"), g.Attr("x2", "5"), g.Attr("y2", "20")),
			g.El("line", g.Attr("x1", "15"), g.Attr("y1", "4"), g.Attr("x2", "9"), g.Attr("y2", "20")),
		)
	}

	underlineIcon := func() g.Node {
		return g.El("svg",
			g.Attr("xmlns", "http://www.w3.org/2000/svg"),
			g.Attr("viewBox", "0 0 24 24"),
			g.Attr("fill", "none"),
			g.Attr("stroke", "currentColor"),
			g.Attr("stroke-width", "2"),
			html.Class("h-4 w-4"),
			g.El("path", g.Attr("d", "M6 3v7a6 6 0 0 0 6 6 6 6 0 0 0 6-6V3")),
			g.El("line", g.Attr("x1", "4"), g.Attr("y1", "21"), g.Attr("x2", "20"), g.Attr("y2", "21")),
		)
	}

	return html.Div(
		html.Class("p-8 space-y-8"),
		
		// Basic toggles
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Basic Toggles")),
			html.Div(html.Class("flex items-center gap-2"),
				Default(g.Text("Toggle")),
				Pressed(g.Text("Pressed")),
				New(Props{Disabled: true}, g.Text("Disabled")),
			),
		),
		
		// Variants
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Variants")),
			html.Div(html.Class("flex items-center gap-2"),
				Default(g.Text("Default")),
				Outline(g.Text("Outline")),
			),
		),
		
		// Sizes
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Sizes")),
			html.Div(html.Class("flex items-center gap-2"),
				SmallToggle("default", g.Text("Small")),
				Default(g.Text("Default")),
				Large("default", g.Text("Large")),
			),
		),
		
		// With icons
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("With Icons")),
			html.Div(html.Class("flex items-center gap-2"),
				Icon(boldIcon(), "Toggle bold"),
				WithIcon(italicIcon(), "Italic", false),
				WithIcon(underlineIcon(), "", true),
			),
		),
		
		// Text formatting toolbar
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Text Formatting Toolbar")),
			html.Div(html.Class("rounded-md border p-1"),
				Group(
					FormatButton("bold", boldIcon(), true),
					FormatButton("italic", italicIcon(), false),
					FormatButton("underline", underlineIcon(), false),
				),
			),
		),
		
		// View toggles
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("View Toggles")),
			html.Div(html.Class("flex items-center gap-1"),
				ViewToggle("grid", 
					g.El("svg",
						g.Attr("xmlns", "http://www.w3.org/2000/svg"),
						g.Attr("viewBox", "0 0 24 24"),
						g.Attr("fill", "none"),
						g.Attr("stroke", "currentColor"),
						g.Attr("stroke-width", "2"),
						html.Class("h-4 w-4"),
						g.El("rect", g.Attr("x", "3"), g.Attr("y", "3"), g.Attr("width", "7"), g.Attr("height", "7")),
						g.El("rect", g.Attr("x", "14"), g.Attr("y", "3"), g.Attr("width", "7"), g.Attr("height", "7")),
						g.El("rect", g.Attr("x", "14"), g.Attr("y", "14"), g.Attr("width", "7"), g.Attr("height", "7")),
						g.El("rect", g.Attr("x", "3"), g.Attr("y", "14"), g.Attr("width", "7"), g.Attr("height", "7")),
					),
					true,
				),
				ViewToggle("list",
					g.El("svg",
						g.Attr("xmlns", "http://www.w3.org/2000/svg"),
						g.Attr("viewBox", "0 0 24 24"),
						g.Attr("fill", "none"),
						g.Attr("stroke", "currentColor"),
						g.Attr("stroke-width", "2"),
						html.Class("h-4 w-4"),
						g.El("line", g.Attr("x1", "8"), g.Attr("y1", "6"), g.Attr("x2", "21"), g.Attr("y2", "6")),
						g.El("line", g.Attr("x1", "8"), g.Attr("y1", "12"), g.Attr("x2", "21"), g.Attr("y2", "12")),
						g.El("line", g.Attr("x1", "8"), g.Attr("y1", "18"), g.Attr("x2", "21"), g.Attr("y2", "18")),
						g.El("line", g.Attr("x1", "3"), g.Attr("y1", "6"), g.Attr("x2", "3.01"), g.Attr("y2", "6")),
						g.El("line", g.Attr("x1", "3"), g.Attr("y1", "12"), g.Attr("x2", "3.01"), g.Attr("y2", "12")),
						g.El("line", g.Attr("x1", "3"), g.Attr("y1", "18"), g.Attr("x2", "3.01"), g.Attr("y2", "18")),
					),
					false,
				),
			),
		),
		
		// Settings toggles
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Settings")),
			html.Div(html.Class("space-y-3"),
				html.Div(html.Class("flex items-center justify-between"),
					html.Span(html.Class("text-sm"), g.Text("Show notifications")),
					New(Props{
						Variant: "outline",
						Size:    "sm",
						Pressed: true,
					}, g.Text("On")),
				),
				html.Div(html.Class("flex items-center justify-between"),
					html.Span(html.Class("text-sm"), g.Text("Auto-save")),
					New(Props{
						Variant: "outline",
						Size:    "sm",
					}, g.Text("Off")),
				),
				html.Div(html.Class("flex items-center justify-between"),
					html.Span(html.Class("text-sm"), g.Text("Dark mode")),
					New(Props{
						Variant: "outline",
						Size:    "sm",
						Pressed: true,
					}, g.Text("On")),
				),
			),
		),
		
		// Multiple selection group
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Multi-Select Options")),
			html.P(html.Class("text-sm text-muted-foreground mb-2"), g.Text("Select your preferences:")),
			html.Div(html.Class("flex flex-wrap gap-2"),
				Outline(g.Text("JavaScript")),
				New(Props{Variant: "outline", Pressed: true}, g.Text("TypeScript")),
				Outline(g.Text("Python")),
				New(Props{Variant: "outline", Pressed: true}, g.Text("Go")),
				Outline(g.Text("Rust")),
			),
		),
		
		// Interactive example
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Interactive Example")),
			html.P(html.Class("text-sm text-muted-foreground mb-4"), g.Text("(Add JavaScript to make this functional)")),
			html.Div(html.Class("space-y-4"),
				html.Div(
					html.P(html.Class("text-sm mb-2"), g.Text("Toggle to show/hide content:")),
					New(Props{
						ID:      "content-toggle",
						OnClick: "document.getElementById('hidden-content').classList.toggle('hidden')",
					}, g.Text("Show Content")),
				),
				html.Div(
					html.ID("hidden-content"),
					html.Class("hidden p-4 rounded-md bg-muted"),
					html.P(html.Class("text-sm"), g.Text("This content is toggled by the button above!")),
				),
			),
		),
		
		// Disabled states
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Disabled States")),
			html.Div(html.Class("flex items-center gap-2"),
				New(Props{
					Disabled: true,
					Variant:  "default",
				}, g.Text("Disabled Off")),
				New(Props{
					Disabled: true,
					Variant:  "default",
					Pressed:  true,
				}, g.Text("Disabled On")),
				New(Props{
					Disabled: true,
					Variant:  "outline",
				}, g.Text("Disabled Outline")),
			),
		),
	)
}