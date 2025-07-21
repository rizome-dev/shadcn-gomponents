package switchcomp

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

// Example demonstrates how to use the Switch component
func Example() g.Node {
	return html.Div(
		html.Class("p-8 space-y-8"),
		
		// Basic switches
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Basic Switches")),
			html.Div(html.Class("flex items-center space-x-4"),
				Default(),
				Checked(),
				Disabled(false),
				Disabled(true),
			),
		),
		
		// Different sizes
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Different Sizes")),
			html.Div(html.Class("flex items-center space-x-6"),
				html.Div(html.Class("flex flex-col items-center space-y-2"),
					html.P(html.Class("text-sm text-muted-foreground"), g.Text("Small")),
					SmallComponent(),
				),
				html.Div(html.Class("flex flex-col items-center space-y-2"),
					html.P(html.Class("text-sm text-muted-foreground"), g.Text("Default")),
					Default(),
				),
				html.Div(html.Class("flex flex-col items-center space-y-2"),
					html.P(html.Class("text-sm text-muted-foreground"), g.Text("Large")),
					Large(),
				),
			),
		),
		
		// With labels
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("With Labels")),
			html.Div(html.Class("space-y-4"),
				WithLabel(Props{}, "Airplane Mode"),
				WithLabel(Props{Checked: true}, "Wi-Fi"),
				WithLabel(Props{}, "Bluetooth"),
				WithLabel(Props{Disabled: true}, "Location Services"),
			),
		),
		
		// Form fields
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Form Fields")),
			html.Div(html.Class("space-y-4"),
				FormField(
					Props{Name: "marketing", Checked: true},
					"Marketing emails",
					"Receive emails about new products, features, and more.",
				),
				FormField(
					Props{Name: "security"},
					"Security emails",
					"Receive emails about your account security.",
				),
				FormField(
					Props{Name: "updates", Checked: true},
					"Product updates",
					"Get notified when we release new features.",
				),
			),
		),
		
		// Settings panel
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Settings Panel")),
			html.Div(html.Class("max-w-md rounded-lg border bg-card"),
				html.Div(html.Class("p-6 space-y-6"),
					html.H3(html.Class("text-lg font-semibold"), g.Text("Privacy Settings")),
					
					Setting(
						"analytics",
						"Analytics",
						"Help us improve by sending anonymous usage data",
						true,
					),
					
					Setting(
						"personalization",
						"Personalization",
						"Enable personalized recommendations based on your usage",
						true,
					),
					
					Setting(
						"tracking",
						"Third-party tracking",
						"Allow third-party services to track your activity",
						false,
					),
					
					Setting(
						"cookies",
						"Functional cookies",
						"Enable cookies for enhanced functionality",
						true,
					),
				),
			),
		),
		
		// Notification preferences
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Notification Preferences")),
			html.Div(html.Class("max-w-md space-y-4"),
				html.Div(html.Class("rounded-lg border p-4"),
					html.H4(html.Class("font-medium mb-4"), g.Text("Email Notifications")),
					html.Div(html.Class("space-y-3"),
						WithLabel(
							Props{Name: "email_comments", Checked: true},
							"Comments and replies",
						),
						WithLabel(
							Props{Name: "email_mentions"},
							"Mentions",
						),
						WithLabel(
							Props{Name: "email_newsletter", Checked: true},
							"Weekly newsletter",
						),
					),
				),
				
				html.Div(html.Class("rounded-lg border p-4"),
					html.H4(html.Class("font-medium mb-4"), g.Text("Push Notifications")),
					html.Div(html.Class("space-y-3"),
						WithLabel(
							Props{Name: "push_messages", Checked: true},
							"Direct messages",
						),
						WithLabel(
							Props{Name: "push_updates", Checked: true},
							"Product updates",
						),
						WithLabel(
							Props{Name: "push_reminders"},
							"Task reminders",
						),
					),
				),
			),
		),
		
		// In a form
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("In a Form")),
			html.Form(html.Class("max-w-md space-y-6 rounded-lg border bg-card p-6"),
				html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Account Settings")),
				
				html.Div(html.Class("space-y-4"),
					FormField(
						Props{
							Name:     "two_factor",
							Required: true,
						},
						"Two-factor authentication",
						"Add an extra layer of security to your account",
					),
					
					FormField(
						Props{
							Name:    "session_timeout",
							Checked: true,
						},
						"Auto logout",
						"Automatically log out after 30 minutes of inactivity",
					),
					
					FormField(
						Props{
							Name: "api_access",
						},
						"API Access",
						"Enable programmatic access to your account",
					),
				),
				
				html.Div(html.Class("flex gap-2 pt-4"),
					html.Button(html.Type("submit"), html.Class("bg-primary text-primary-foreground"), g.Text("Save Changes")),
					html.Button(html.Type("button"), html.Class("variant-outline"), g.Text("Cancel")),
				),
			),
		),
		
		// Interactive example
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Interactive Example")),
			html.P(html.Class("text-sm text-muted-foreground mb-4"), g.Text("(Add JavaScript to make this functional)")),
			html.Div(html.Class("max-w-md rounded-lg border bg-card p-6"),
				FormField(
					Props{
						ID:       "master-switch",
						OnChange: "document.querySelectorAll('[data-child-switch]').forEach(s => s.disabled = !this.checked)",
					},
					"Enable All Features",
					"Toggle to enable/disable all features below",
				),
				html.Div(html.Class("ml-6 mt-4 space-y-3"),
					WithLabel(
						Props{
							Name:    "feature1",
							Checked: true,
							Class:   "data-[child-switch]:true",
						},
						"Feature 1",
					),
					WithLabel(
						Props{
							Name:    "feature2",
							Checked: true,
							Class:   "data-[child-switch]:true",
						},
						"Feature 2",
					),
					WithLabel(
						Props{
							Name:  "feature3",
							Class: "data-[child-switch]:true",
						},
						"Feature 3",
					),
				),
			),
		),
	)
}