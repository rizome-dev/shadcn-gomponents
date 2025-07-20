package button

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

// Example demonstrates various button usage patterns
func Example() g.Node {
	return html.Div(
		html.Class("space-y-8 p-8"),
		
		// Variants section
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Variants")),
			html.Div(
				html.Class("flex flex-wrap gap-2"),
				Default(g.Text("Default")),
				Destructive(g.Text("Destructive")),
				Outline(g.Text("Outline")),
				Secondary(g.Text("Secondary")),
				Ghost(g.Text("Ghost")),
				LinkButton(g.Text("Link")),
			),
		),

		// Sizes section
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Sizes")),
			html.Div(
				html.Class("flex items-center gap-2"),
				New(Props{Size: "lg"}, g.Text("Large")),
				New(Props{Size: "default"}, g.Text("Default")),
				New(Props{Size: "sm"}, g.Text("Small")),
				Icon("default", g.Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M5 12h14"/><path d="m12 5 7 7-7 7"/></svg>`)),
			),
		),

		// With Icons section
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("With Icons")),
			html.Div(
				html.Class("flex flex-wrap gap-2"),
				Default(
					g.Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M5 12h14"/><path d="m12 5 7 7-7 7"/></svg>`),
					g.Text("Next"),
				),
				Secondary(
					g.Text("Previous"),
					g.Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m19 12-7-7-7 7"/><path d="M5 12h14"/></svg>`),
				),
			),
		),

		// States section
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("States")),
			html.Div(
				html.Class("flex flex-wrap gap-2"),
				New(Props{Disabled: true}, g.Text("Disabled")),
				New(Props{Disabled: true}, 
					g.Raw(`<svg class="animate-spin" xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 12a9 9 0 1 1-6.219-8.56"/></svg>`),
					g.Text("Loading..."),
				),
			),
		),

		// Form Buttons section
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Form Buttons")),
			html.Form(
				html.Class("flex gap-2"),
				New(Props{Type: "submit", Variant: "default"}, g.Text("Submit")),
				New(Props{Type: "reset", Variant: "outline"}, g.Text("Reset")),
				New(Props{Type: "button", Variant: "ghost"}, g.Text("Cancel")),
			),
		),

		// Custom Styling section
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Custom Styling")),
			html.Div(
				html.Class("flex gap-2"),
				New(Props{Class: "w-full"}, g.Text("Full Width")),
				New(Props{Class: "rounded-full", Variant: "outline"}, g.Text("Rounded")),
			),
		),
	)
}