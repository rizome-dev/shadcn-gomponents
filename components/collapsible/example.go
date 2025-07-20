package collapsible

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

// DemoBasic shows a basic collapsible using native details/summary
func DemoBasic() g.Node {
	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Basic Collapsible")),
		Example(),
	)
}

// DemoClosed shows a collapsible that starts closed
func DemoClosed() g.Node {
	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Closed by Default")),
		ExampleClosed(),
	)
}

// DemoStyled shows a styled collapsible with custom appearance
func DemoStyled() g.Node {
	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Styled Collapsible")),
		ExampleStyled(),
	)
}

// DemoMultiple shows multiple collapsibles (accordion-like behavior)
func DemoMultiple() g.Node {
	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Multiple Collapsibles")),
		html.Div(html.Class("space-y-2"),
			g.El("details",
				html.Class("rounded-lg border"),
				g.El("summary",
					html.Class("flex items-center justify-between p-4 cursor-pointer list-none hover:bg-muted/50 transition-colors"),
					html.Span(html.Class("font-medium"), g.Text("Is it accessible?")),
					ChevronRight(html.Class("h-4 w-4 transition-transform [details[open]>&]:rotate-90")),
				),
				html.Div(html.Class("px-4 pb-4 pt-0"),
					html.P(html.Class("text-muted-foreground"),
						g.Text("Yes. It uses native HTML details/summary elements which are accessible by default."),
					),
				),
			),
			g.El("details",
				html.Class("rounded-lg border"),
				g.El("summary",
					html.Class("flex items-center justify-between p-4 cursor-pointer list-none hover:bg-muted/50 transition-colors"),
					html.Span(html.Class("font-medium"), g.Text("Is it styled?")),
					ChevronRight(html.Class("h-4 w-4 transition-transform [details[open]>&]:rotate-90")),
				),
				html.Div(html.Class("px-4 pb-4 pt-0"),
					html.P(html.Class("text-muted-foreground"),
						g.Text("Yes. It comes with default styles that matches the other components' aesthetic."),
					),
				),
			),
			g.El("details",
				html.Class("rounded-lg border"),
				g.El("summary",
					html.Class("flex items-center justify-between p-4 cursor-pointer list-none hover:bg-muted/50 transition-colors"),
					html.Span(html.Class("font-medium"), g.Text("Is it animated?")),
					ChevronRight(html.Class("h-4 w-4 transition-transform [details[open]>&]:rotate-90")),
				),
				html.Div(html.Class("px-4 pb-4 pt-0"),
					html.P(html.Class("text-muted-foreground"),
						g.Text("The chevron icon rotates smoothly. For content animation, you would need to add CSS transitions."),
					),
				),
			),
		),
	)
}

// DemoNested shows nested collapsibles
func DemoNested() g.Node {
	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Nested Collapsibles")),
		g.El("details",
			html.Class("w-full max-w-md border rounded-lg"),
			g.Attr("open", ""),
			g.El("summary",
				html.Class("flex items-center justify-between p-4 cursor-pointer list-none"),
				html.Span(html.Class("font-medium"), g.Text("Main Section")),
				ChevronRight(html.Class("h-4 w-4 transition-transform [details[open]>&]:rotate-90")),
			),
			html.Div(html.Class("px-4 pb-4 space-y-4"),
				html.P(html.Class("text-sm text-muted-foreground"), g.Text("This section contains nested collapsibles.")),
				g.El("details",
					html.Class("border rounded-lg"),
					g.El("summary",
						html.Class("flex items-center justify-between p-3 cursor-pointer list-none hover:bg-muted/50"),
						html.Span(html.Class("text-sm font-medium"), g.Text("Subsection 1")),
						ChevronRight(html.Class("h-3 w-3 transition-transform [details[open]>&]:rotate-90")),
					),
					html.Div(html.Class("px-3 pb-3 pt-0"),
						html.P(html.Class("text-sm text-muted-foreground"), g.Text("Content of subsection 1")),
					),
				),
				g.El("details",
					html.Class("border rounded-lg"),
					g.El("summary",
						html.Class("flex items-center justify-between p-3 cursor-pointer list-none hover:bg-muted/50"),
						html.Span(html.Class("text-sm font-medium"), g.Text("Subsection 2")),
						ChevronRight(html.Class("h-3 w-3 transition-transform [details[open]>&]:rotate-90")),
					),
					html.Div(html.Class("px-3 pb-3 pt-0"),
						html.P(html.Class("text-sm text-muted-foreground"), g.Text("Content of subsection 2")),
					),
				),
			),
		),
	)
}

// DemoWithForm shows a collapsible containing a form
func DemoWithForm() g.Node {
	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Collapsible with Form")),
		g.El("details",
			html.Class("w-full max-w-md border rounded-lg"),
			g.El("summary",
				html.Class("flex items-center justify-between p-4 cursor-pointer list-none bg-muted/50"),
				html.Div(
					html.H4(html.Class("font-medium"), g.Text("Advanced Settings")),
					html.P(html.Class("text-sm text-muted-foreground"), g.Text("Configure advanced options")),
				),
				ChevronRight(html.Class("h-4 w-4 transition-transform [details[open]>&]:rotate-90")),
			),
			html.Form(html.Class("p-4 space-y-4"),
				html.Div(html.Class("space-y-2"),
					html.Label(html.For("notifications"), g.Text("Email Notifications")),
					html.Select(
						html.ID("notifications"),
						html.Name("notifications"),
						html.Class("w-full px-3 py-2 border rounded-md"),
						html.Option(html.Value("all"), g.Text("All notifications")),
						html.Option(html.Value("important"), html.Selected(), g.Text("Important only")),
						html.Option(html.Value("none"), g.Text("None")),
					),
				),
				html.Div(html.Class("space-y-2"),
					html.Label(html.For("theme"), g.Text("Theme Preference")),
					html.Div(html.Class("space-y-1"),
						html.Label(html.Class("flex items-center gap-2"),
							html.Input(html.Type("radio"), html.Name("theme"), html.Value("light"), html.Class("w-4 h-4")),
							html.Span(g.Text("Light")),
						),
						html.Label(html.Class("flex items-center gap-2"),
							html.Input(html.Type("radio"), html.Name("theme"), html.Value("dark"), html.Checked(), html.Class("w-4 h-4")),
							html.Span(g.Text("Dark")),
						),
						html.Label(html.Class("flex items-center gap-2"),
							html.Input(html.Type("radio"), html.Name("theme"), html.Value("system"), html.Class("w-4 h-4")),
							html.Span(g.Text("System")),
						),
					),
				),
				html.Div(html.Class("pt-2"),
					html.Button(
						html.Type("submit"),
						html.Class("w-full px-4 py-2 bg-primary text-primary-foreground rounded-md hover:bg-primary/90"),
						g.Text("Save Settings"),
					),
				),
			),
		),
	)
}

// Helper function for chevron icon (simplified)
func ChevronRight(attrs ...g.Node) g.Node {
	defaultAttrs := []g.Node{
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("width", "24"),
		g.Attr("height", "24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
	}
	
	allAttrs := append(defaultAttrs, attrs...)
	
	return g.El("svg",
		append(allAttrs,
			g.El("path", g.Attr("d", "m9 18 6-6-6-6")),
		)...,
	)
}