package radio

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

// Example demonstrates how to use the Radio component
func Example() g.Node {
	return html.Div(
		html.Class("p-8 space-y-8"),
		
		// Basic radio group
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Basic Radio Group")),
			Group(
				GroupProps{Name: "plan"},
				WithLabel(ItemProps{Value: "free"}, "plan", "Free"),
				WithLabel(ItemProps{Value: "pro", Checked: true}, "plan", "Pro"),
				WithLabel(ItemProps{Value: "enterprise"}, "plan", "Enterprise"),
			),
		),
		
		// Horizontal radio group
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Horizontal Layout")),
			Horizontal(
				"size",
				[]Option{
					{Value: "xs", Label: "Extra Small"},
					{Value: "sm", Label: "Small"},
					{Value: "md", Label: "Medium"},
					{Value: "lg", Label: "Large"},
					{Value: "xl", Label: "Extra Large"},
				},
				"md",
			),
		),
		
		// Simple radio group with disabled option
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("With Disabled Option")),
			Simple(
				"notification",
				[]Option{
					{Value: "all", Label: "All notifications"},
					{Value: "important", Label: "Important only"},
					{Value: "none", Label: "No notifications"},
					{Value: "custom", Label: "Custom (Pro only)", Disabled: true},
				},
				"important",
			),
		),
		
		// Form field with description
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Form Field")),
			FormField(
				"privacy",
				"Privacy Settings",
				"Control who can see your profile and activity",
				[]Option{
					{Value: "public", Label: "Public - Anyone can view"},
					{Value: "friends", Label: "Friends only"},
					{Value: "private", Label: "Private - Only you"},
				},
				"friends",
			),
		),
		
		// Card style radio options
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Card Style Options")),
			Group(
				GroupProps{Name: "shipping", Class: "grid gap-3"},
				Card(
					ItemProps{Value: "standard"},
					"shipping",
					"Standard Shipping",
					"5-7 business days • Free",
				),
				Card(
					ItemProps{Value: "express", Checked: true},
					"shipping",
					"Express Shipping",
					"1-2 business days • $9.99",
				),
				Card(
					ItemProps{Value: "overnight"},
					"shipping",
					"Overnight Shipping",
					"Next business day • $24.99",
				),
				Card(
					ItemProps{Value: "international", Disabled: true},
					"shipping",
					"International",
					"Not available for this order",
				),
			),
		),
		
		// In a settings context
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Settings Example")),
			html.Div(html.Class("max-w-md rounded-lg border bg-card p-6 space-y-6"),
				html.Div(
					html.H3(html.Class("text-lg font-semibold"), g.Text("Appearance")),
					html.P(html.Class("text-sm text-muted-foreground mt-1"), g.Text("Customize how the app looks on your device")),
				),
				
				FormField(
					"theme",
					"Theme",
					"",
					[]Option{
						{Value: "light", Label: "Light"},
						{Value: "dark", Label: "Dark"},
						{Value: "system", Label: "System default"},
					},
					"system",
				),
				
				html.Div(html.Class("pt-4 border-t"),
					FormField(
						"language",
						"Language",
						"",
						[]Option{
							{Value: "en", Label: "English"},
							{Value: "es", Label: "Español"},
							{Value: "fr", Label: "Français"},
							{Value: "de", Label: "Deutsch"},
						},
						"en",
					),
				),
			),
		),
		
		// Complex form example
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Complex Form")),
			html.Form(html.Class("space-y-6 max-w-md"),
				FormField(
					"account-type",
					"Account Type",
					"Choose the type of account that best fits your needs",
					[]Option{
						{Value: "personal", Label: "Personal Account"},
						{Value: "business", Label: "Business Account"},
						{Value: "developer", Label: "Developer Account"},
					},
					"personal",
				),
				
				FormField(
					"billing",
					"Billing Cycle",
					"",
					[]Option{
						{Value: "monthly", Label: "Monthly billing"},
						{Value: "annual", Label: "Annual billing (save 20%)"},
					},
					"monthly",
				),
				
				html.Div(html.Class("flex gap-2 pt-4"),
					html.Button(html.Type("submit"), html.Class("bg-primary text-primary-foreground"), g.Text("Continue")),
					html.Button(html.Type("button"), html.Class("variant-outline"), g.Text("Cancel")),
				),
			),
		),
		
		// Custom styled radio group
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Custom Styles")),
			Group(
				GroupProps{Name: "color", Orientation: "horizontal"},
				html.Div(html.Class("flex items-center space-x-2"),
					Item(ItemProps{Value: "red", ID: "radio-red", Class: "border-red-500 text-red-500"}, "color"),
					html.Label(html.For("radio-red"), g.Text("Red")),
				),
				html.Div(html.Class("flex items-center space-x-2"),
					Item(ItemProps{Value: "green", ID: "radio-green", Class: "border-green-500 text-green-500"}, "color"),
					html.Label(html.For("radio-green"), g.Text("Green")),
				),
				html.Div(html.Class("flex items-center space-x-2"),
					Item(ItemProps{Value: "blue", ID: "radio-blue", Class: "border-blue-500 text-blue-500", Checked: true}, "color"),
					html.Label(html.For("radio-blue"), g.Text("Blue")),
				),
			),
		),
	)
}