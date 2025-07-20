package label

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

// Examples demonstrates various label usage patterns
func Examples() g.Node {
	return html.Div(
		html.Class("space-y-8 p-8 max-w-2xl"),

		// Basic Labels section
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Basic Labels")),
			html.Div(
				html.Class("space-y-2"),
				Default("Simple Label"),
				Default("Another Label"),
			),
		),

		// Labels with Required Indicator section
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Labels with Required Indicator")),
			html.Div(
				html.Class("space-y-2"),
				WithRequired("Optional Field", false),
				WithRequired("Required Field", true),
				New(Props{Required: true}, g.Text("Custom Required Label")),
			),
		),

		// Labels for Inputs section
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Labels for Inputs")),
			html.Div(
				html.Class("space-y-4"),
				html.Div(
					html.Class("space-y-2"),
					ForInput("username-field", "Username"),
					html.Input(
						html.Type("text"),
						html.ID("username-field"),
						html.Class("flex h-9 w-full rounded-md border"),
					),
				),
				html.Div(
					html.Class("space-y-2"),
					ForInputRequired("email-field", "Email Address"),
					html.Input(
						html.Type("email"),
						html.ID("email-field"),
						html.Class("flex h-9 w-full rounded-md border"),
					),
				),
			),
		),

		// Labels with Custom Content section
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Labels with Custom Content")),
			html.Div(
				html.Class("space-y-2"),
				New(Props{}, 
					g.Text("Label with icon "),
					g.Raw(`<svg class="inline w-4 h-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 16v-4"/><path d="M12 8html.01"/></svg>`),
				),
				New(Props{Class: "text-blue-600"}, 
					g.Text("Colored Label"),
				),
				New(Props{}, 
					g.Text("Label with "),
					html.Code(g.Text("code")),
					g.Text(" inside"),
				),
			),
		),

		// Disabled State section
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Disabled State")),
			html.Div(
				html.Class("space-y-4"),
				html.Div(
					html.Class("space-y-2"),
					html.P(html.Class("text-sm text-muted-foreground"), g.Text("Label with disabled input (peer-disabled styles apply)")),
					Default("This label will be dimmed"),
					html.Input(
						html.Type("text"),
						html.Disabled(),
						html.Class("peer flex h-9 w-full rounded-md border"),
					),
				),
				html.Div(
					html.Class("group space-y-2"),
					g.Attr("data-disabled", "true"),
					html.P(html.Class("text-sm text-muted-foreground"), g.Text("Label in disabled group")),
					Default("This label is in a disabled group"),
				),
			),
		),

		// Form Layout Examples section
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Form Layout Examples")),
			html.Div(
				html.Class("space-y-4"),
				// Vertical layout
				html.Div(
					html.Class("space-y-4 border rounded-lg p-4"),
					html.H4(html.Class("font-medium"), g.Text("Vertical Layout")),
					html.Div(
						html.Class("space-y-2"),
						ForInputRequired("name", "Full Name"),
						html.Input(html.Type("text"), html.ID("name"), html.Class("flex h-9 w-full rounded-md border")),
					),
					html.Div(
						html.Class("space-y-2"),
						ForInput("company", "Company"),
						html.Input(html.Type("text"), html.ID("company"), html.Class("flex h-9 w-full rounded-md border")),
					),
				),
				// Horizontal layout
				html.Div(
					html.Class("space-y-4 border rounded-lg p-4"),
					html.H4(html.Class("font-medium"), g.Text("Horizontal Layout")),
					html.Div(
						html.Class("grid grid-cols-3 gap-2 items-center"),
						ForInput("horizontal-name", "Name"),
						html.Input(html.Type("text"), html.ID("horizontal-name"), html.Class("col-span-2 flex h-9 w-full rounded-md border")),
					),
					html.Div(
						html.Class("grid grid-cols-3 gap-2 items-center"),
						ForInputRequired("horizontal-email", "Email"),
						html.Input(html.Type("email"), html.ID("horizontal-email"), html.Class("col-span-2 flex h-9 w-full rounded-md border")),
					),
				),
			),
		),

		// Custom Styling section
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Custom Styling")),
			html.Div(
				html.Class("space-y-2"),
				New(Props{Class: "text-lg font-bold"}, g.Text("Large Bold Label")),
				New(Props{Class: "text-xs text-gray-500"}, g.Text("Small Gray Label")),
				New(Props{Class: "uppercase tracking-wide"}, g.Text("Uppercase Label")),
				New(Props{Class: "italic"}, g.Text("Italic Label")),
			),
		),
	)
}