package card

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/components/button"
	"github.com/rizome-dev/shadcn-gomponents/components/input"
	"github.com/rizome-dev/shadcn-gomponents/components/label"
)

// Examples demonstrates various card usage patterns
func Examples() g.Node {
	return html.Div(
		html.Class("space-y-8 p-8"),

		// Basic Card section
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Basic Card")),
			Simple(
				html.P(g.Text("This is a simple card with just content inside.")),
			),
		),

		// Card with Header section
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Card with Header")),
			Card(
				CardHeader(
					CardTitle(g.Text("Card Title")),
					CardDescription(g.Text("Card description goes here")),
				),
				CardContent(
					html.P(g.Text("This is the main content of the card. You can put any content here.")),
				),
			),
		),

		// Card with Header and Footer section
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Card with Header and Footer")),
			WithFooter(
				"Featured Article",
				"Published on January 1, 2024",
				[]g.Node{
					html.P(g.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.")),
				},
				[]g.Node{
					button.Outline(g.Text("Read More")),
				},
			),
		),

		// Card with Action section
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Card with Action")),
			Card(
				CardHeader(
					CardTitle(g.Text("Settings")),
					CardDescription(g.Text("Manage your account settings")),
					CardAction(
						button.Icon("ghost", 
							g.Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="1"/><circle cx="12" cy="5" r="1"/><circle cx="12" cy="19" r="1"/></svg>`),
						),
					),
				),
				CardContent(
					html.P(g.Text("Configure your preferences and account details.")),
				),
			),
		),

		// Form Card section
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Form Card")),
			New(Props{
				Title:       "Create Account",
				Description: "Enter your details to get started",
				Class:       "w-full max-w-md",
			},
				html.Form(
					html.Class("space-y-4"),
					html.Div(
						html.Class("space-y-2"),
						label.ForInputRequired("name", "Name"),
						input.New(input.Props{
							ID:       "name",
							Name:     "name",
							Required: true,
						}),
					),
					html.Div(
						html.Class("space-y-2"),
						label.ForInputRequired("email", "Email"),
						input.New(input.Props{
							Type:     "email",
							ID:       "email",
							Name:     "email",
							Required: true,
						}),
					),
					html.Div(
						html.Class("space-y-2"),
						label.ForInputRequired("password", "Password"),
						input.New(input.Props{
							Type:     "password",
							ID:       "password",
							Name:     "password",
							Required: true,
						}),
					),
				),
			),
			Card(
				CardFooter(
					button.New(button.Props{
						Type:  "submit",
						Class: "w-full",
					}, g.Text("Create Account")),
				),
			),
		),

		// Multiple Cards Grid section
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Cards Grid")),
			html.Div(
				html.Class("grid grid-cols-1 md:grid-cols-3 gap-4"),
				g.Map([]struct {
					Title       string
					Description string
					Content     string
				}{
					{
						Title:       "Performance",
						Description: "Monitor your app performance",
						Content:     "Track response times, error rates, and throughput in real-time.",
					},
					{
						Title:       "Security",
						Description: "Keep your data safe",
						Content:     "Advanced encryption and authentication to protect your information.",
					},
					{
						Title:       "Analytics",
						Description: "Understand your users",
						Content:     "Detailed insights into user behavior and engagement patterns.",
					},
				}, func(item struct {
					Title       string
					Description string
					Content     string
				}) g.Node {
					return Card(
						CardHeader(
							CardTitle(g.Text(item.Title)),
							CardDescription(g.Text(item.Description)),
						),
						CardContent(
							html.P(g.Text(item.Content)),
						),
						CardFooter(
							button.LinkButton(g.Text("Learn more →")),
						),
					)
				}),
			),
		),

		// Custom Styled Card section
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Custom Styled Card")),
			New(Props{
				Title:       "Premium Plan",
				Description: "Best for growing teams",
				Class:       "border-2 border-blue-500 shadow-lg",
			},
				html.Div(
					html.Class("space-y-4"),
					html.Div(
						html.Class("text-3xl font-bold"),
						g.Text("$99"),
						html.Span(html.Class("text-sm font-normal text-muted-foreground"), g.Text("/month")),
					),
					html.Ul(
						html.Class("space-y-2 text-sm"),
						html.Li(g.Text("✓ Unlimited projects")),
						html.Li(g.Text("✓ Advanced analytics")),
						html.Li(g.Text("✓ Priority support")),
						html.Li(g.Text("✓ Custom integrations")),
					),
				),
			),
		),

		// Notification Card section
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Notification Card")),
			Card(
				CardHeader(
					html.Div(
						html.Class("flex items-center gap-2"),
						g.Raw(`<svg class="w-5 h-5 text-green-500" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"/><polyline points="22 4 12 14.01 9 11.01"/></svg>`),
						CardTitle(g.Text("Success!")),
					),
				),
				CardContent(
					html.P(g.Text("Your changes have been saved successfully.")),
				),
				CardFooter(
					html.Div(
						html.Class("flex gap-2"),
						button.Secondary(g.Text("Dismiss")),
						button.Default(g.Text("View Changes")),
					),
				),
			),
		),
	)
}