package checkbox

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

// Example demonstrates how to use the Checkbox component
func Example() g.Node {
	return html.Div(
		html.Class("p-8 space-y-8"),
		
		// Basic checkboxes
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Basic Checkboxes")),
			html.Div(html.Class("space-y-2"),
				Default(),
				html.Checked(),
				New(Props{Indeterminate: true}),
			),
		),
		
		// Checkboxes with labels
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("With Labels")),
			html.Div(html.Class("space-y-2"),
				WithLabel(Props{}, "Accept terms and conditions"),
				WithLabel(Props{Checked: true}, "Send me promotional emails"),
				WithLabel(Props{Disabled: true}, "This option is disabled"),
			),
		),
		
		// Form fields with descriptions
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Form Fields")),
			FormField(
				Props{Name: "mobile-notifications"},
				"Use different settings for my mobile devices",
				"You can manage your mobile notifications in the mobile settings page.",
			),
			html.Div(html.Class("mt-4"),
				FormField(
					Props{Name: "security-emails", Checked: true},
					"Email me about security updates",
					"We'll only email you when there are critical security updates.",
				),
			),
		),
		
		// Checkbox group
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Checkbox Group")),
			Group(
				"Select your interests",
				WithLabel(Props{Name: "interests", Value: "design"}, "Design"),
				WithLabel(Props{Name: "interests", Value: "development", Checked: true}, "Development"),
				WithLabel(Props{Name: "interests", Value: "business"}, "Business"),
				WithLabel(Props{Name: "interests", Value: "marketing"}, "Marketing"),
			),
		),
		
		// Different states
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Different States")),
			html.Div(html.Class("grid grid-cols-2 gap-4"),
				html.Div(html.Class("space-y-2"),
					html.P(html.Class("text-sm text-muted-foreground"), g.Text("Unchecked")),
					Default(),
				),
				html.Div(html.Class("space-y-2"),
					html.P(html.Class("text-sm text-muted-foreground"), g.Text("Checked")),
					html.Checked(),
				),
				html.Div(html.Class("space-y-2"),
					html.P(html.Class("text-sm text-muted-foreground"), g.Text("Indeterminate")),
					New(Props{Indeterminate: true}),
				),
				html.Div(html.Class("space-y-2"),
					html.P(html.Class("text-sm text-muted-foreground"), g.Text("Disabled")),
					Disabled(false),
				),
			),
		),
		
		// In a card/form context
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("In Context")),
			html.Div(html.Class("max-w-md rounded-lg border bg-card p-6"),
				html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Email Preferences")),
				html.Div(html.Class("space-y-4"),
					FormField(
						Props{Name: "marketing", Checked: true},
						"Marketing communications",
						"Receive updates about new features and products.",
					),
					FormField(
						Props{Name: "social"},
						"Social emails",
						"Receive notifications about friend requests and messages.",
					),
					FormField(
						Props{Name: "security", Checked: true, Required: true},
						"Security alerts",
						"Important notifications about your account security.",
					),
				),
			),
		),
		
		// With JavaScript interaction (example)
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Interactive Example")),
			html.P(html.Class("text-sm text-muted-foreground mb-2"), g.Text("(Add JavaScript to make this functional)")),
			WithLabel(
				Props{
					ID:       "toggle-all",
					OnChange: "document.querySelectorAll('[name=\"item\"]').forEach(cb => cb.checked = this.checked)",
				},
				"Select all items",
			),
			html.Div(html.Class("ml-6 mt-2 space-y-2"),
				WithLabel(Props{Name: "item", Value: "1"}, "Item 1"),
				WithLabel(Props{Name: "item", Value: "2"}, "Item 2"),
				WithLabel(Props{Name: "item", Value: "3"}, "Item 3"),
			),
		),
		
		// Indeterminate state example
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Table Selection Example")),
			html.Div(html.Class("rounded-md border"),
				html.Div(html.Class("p-4 border-b"),
					WithLabel(
						Props{Indeterminate: true},
						"Select all (2 of 4 selected)",
					),
				),
				html.Div(html.Class("divide-y"),
					html.Div(html.Class("p-4 flex items-center space-x-2"),
						New(Props{Checked: true}),
						html.Span(html.Class("text-sm"), g.Text("Row 1")),
					),
					html.Div(html.Class("p-4 flex items-center space-x-2"),
						New(Props{}),
						html.Span(html.Class("text-sm"), g.Text("Row 2")),
					),
					html.Div(html.Class("p-4 flex items-center space-x-2"),
						New(Props{Checked: true}),
						html.Span(html.Class("text-sm"), g.Text("Row 3")),
					),
					html.Div(html.Class("p-4 flex items-center space-x-2"),
						New(Props{}),
						html.Span(html.Class("text-sm"), g.Text("Row 4")),
					),
				),
			),
		),
	)
}