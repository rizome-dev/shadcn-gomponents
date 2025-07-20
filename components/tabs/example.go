package tabs

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/components/button"
	"github.com/rizome-dev/shadcn-gomponents/components/card"
	"github.com/rizome-dev/shadcn-gomponents/components/input"
	"github.com/rizome-dev/shadcn-gomponents/components/label"
)

// Example demonstrates various tabs usage patterns
func Example() g.Node {
	return html.Div(
		html.Class("max-w-4xl mx-auto p-8 space-y-8"),
		html.H2(html.Class("text-2xl font-bold mb-6"), g.Text("Tabs Examples")),
		
		// Example 1: Basic tabs with account/password
		html.Section(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Account Settings Tabs")),
			html.Div(
				html.Class("flex w-full max-w-sm flex-col gap-6"),
				WithDefault("account",
					TabsList(ListProps{},
						Trigger(TriggerProps{Value: "account"}, g.Text("Account")),
						Trigger(TriggerProps{Value: "password"}, g.Text("Password")),
					),
					TabsContent(ContentProps{Value: "account"},
						card.Card(
							card.CardHeader(
								card.CardTitle(g.Text("Account")),
								card.CardDescription(
									g.Text("Make changes to your account here. Click save when you're done."),
								),
							),
							card.CardContent(
								html.Div(html.Class("grid gap-3"),
									label.New(label.Props{For: "tabs-demo-name"}, g.Text("Name")),
									input.New(input.Props{
										ID:    "tabs-demo-name",
										Value: "Pedro Duarte",
									}),
								),
								html.Div(html.Class("grid gap-3"),
									label.New(label.Props{For: "tabs-demo-username"}, g.Text("Username")),
									input.New(input.Props{
										ID:    "tabs-demo-username",
										Value: "@peduarte",
									}),
								),
							),
							card.CardFooter(
								button.Default(g.Text("Save changes")),
							),
						),
					),
					TabsContent(ContentProps{Value: "password"},
						card.Card(
							card.CardHeader(
								card.CardTitle(g.Text("Password")),
								card.CardDescription(
									g.Text("Change your password here. After saving, you'll be logged out."),
								),
							),
							card.CardContent(
								html.Div(html.Class("grid gap-3"),
									label.New(label.Props{For: "tabs-demo-current"}, g.Text("Current password")),
									input.New(input.Props{
										ID:   "tabs-demo-current",
										Type: "password",
									}),
								),
								html.Div(html.Class("grid gap-3"),
									label.New(label.Props{For: "tabs-demo-new"}, g.Text("New password")),
									input.New(input.Props{
										ID:   "tabs-demo-new",
										Type: "password",
									}),
								),
							),
							card.CardFooter(
								button.Default(g.Text("Save password")),
							),
						),
					),
				),
			),
		),
		
		// Example 2: Simple content tabs
		html.Section(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Simple Content Tabs")),
			New(Props{DefaultValue: "overview"},
				TabsList(ListProps{},
					Trigger(TriggerProps{Value: "overview"}, g.Text("Overview")),
					Trigger(TriggerProps{Value: "analytics"}, g.Text("Analytics")),
					Trigger(TriggerProps{Value: "reports"}, g.Text("Reports")),
					Trigger(TriggerProps{Value: "notifications", Disabled: true}, g.Text("Notifications")),
				),
				TabsContent(ContentProps{Value: "overview", Class: "p-4 bg-muted/50 rounded-lg mt-2"},
					html.H4(html.Class("font-semibold mb-2"), g.Text("Overview")),
					html.P(html.Class("text-muted-foreground"), g.Text("Welcome to your dashboard. Here you can see an overview of your account activity and recent updates.")),
				),
				TabsContent(ContentProps{Value: "analytics", Class: "p-4 bg-muted/50 rounded-lg mt-2"},
					html.H4(html.Class("font-semibold mb-2"), g.Text("Analytics")),
					html.P(html.Class("text-muted-foreground"), g.Text("View detailed analytics about your usage patterns, performance metrics, and growth trends.")),
				),
				TabsContent(ContentProps{Value: "reports", Class: "p-4 bg-muted/50 rounded-lg mt-2"},
					html.H4(html.Class("font-semibold mb-2"), g.Text("Reports")),
					html.P(html.Class("text-muted-foreground"), g.Text("Generate and download comprehensive reports for your records and analysis.")),
				),
			),
		),
		
		// Example 3: Code/Preview tabs pattern
		html.Section(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Code Preview Tabs")),
			New(Props{DefaultValue: "preview", Class: "w-full"},
				TabsList(ListProps{Class: "grid w-full grid-cols-2"},
					Trigger(TriggerProps{Value: "preview"}, g.Text("Preview")),
					Trigger(TriggerProps{Value: "code"}, g.Text("Code")),
				),
				TabsContent(ContentProps{Value: "preview", Class: "mt-2"},
					html.Div(html.Class("border rounded-lg p-6"),
						button.Default(g.Text("Click me")),
					),
				),
				TabsContent(ContentProps{Value: "code", Class: "mt-2"},
					html.Pre(html.Class("border rounded-lg p-4 bg-muted overflow-x-auto"),
						html.Code(g.Text(`<button class="px-4 py-2 bg-primary text-white rounded">
  Click me
</button>`)),
					),
				),
			),
		),
		
		// Example 4: Styled tabs
		html.Section(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Custom Styled Tabs")),
			New(Props{DefaultValue: "features", Class: "w-full"},
				TabsList(ListProps{Class: "bg-primary/10 border-primary/20"},
					Trigger(TriggerProps{Value: "features", Class: "data-[state=active]:bg-primary data-[state=active]:text-primary-foreground"}, 
						g.Raw(`<svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path></svg>`),
						g.Text("Features"),
					),
					Trigger(TriggerProps{Value: "pricing", Class: "data-[state=active]:bg-primary data-[state=active]:text-primary-foreground"}, 
						g.Raw(`<svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>`),
						g.Text("Pricing"),
					),
					Trigger(TriggerProps{Value: "team", Class: "data-[state=active]:bg-primary data-[state=active]:text-primary-foreground"}, 
						g.Raw(`<svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"></path></svg>`),
						g.Text("Team"),
					),
				),
				TabsContent(ContentProps{Value: "features", Class: "mt-4 p-6 bg-card rounded-lg border"},
					html.H4(html.Class("text-xl font-semibold mb-3"), g.Text("Powerful Features")),
					html.Ul(html.Class("space-y-2 text-muted-foreground"),
						html.Li(g.Text("✓ Real-time collaboration")),
						html.Li(g.Text("✓ Advanced analytics dashboard")),
						html.Li(g.Text("✓ Custom integrations")),
						html.Li(g.Text("✓ 24/7 priority support")),
					),
				),
				TabsContent(ContentProps{Value: "pricing", Class: "mt-4 p-6 bg-card rounded-lg border"},
					html.H4(html.Class("text-xl font-semibold mb-3"), g.Text("Simple Pricing")),
					html.Div(html.Class("grid gap-4"),
						html.Div(html.Class("p-4 border rounded-lg"),
							html.H5(html.Class("font-semibold"), g.Text("Starter - $9/month")),
							html.P(html.Class("text-sm text-muted-foreground"), g.Text("Perfect for individuals and small teams")),
						),
						html.Div(html.Class("p-4 border rounded-lg"),
							html.H5(html.Class("font-semibold"), g.Text("Pro - $29/month")),
							html.P(html.Class("text-sm text-muted-foreground"), g.Text("For growing businesses with advanced needs")),
						),
					),
				),
				TabsContent(ContentProps{Value: "team", Class: "mt-4 p-6 bg-card rounded-lg border"},
					html.H4(html.Class("text-xl font-semibold mb-3"), g.Text("Meet Our Team")),
					html.P(html.Class("text-muted-foreground mb-4"), g.Text("We're a diverse group of individuals passionate about creating amazing products.")),
					html.Div(html.Class("grid grid-cols-3 gap-4"),
						html.Div(html.Class("text-center"),
							html.Div(html.Class("w-16 h-16 bg-muted rounded-full mx-auto mb-2")),
							html.P(html.Class("text-sm font-medium"), g.Text("Alice Chen")),
							html.P(html.Class("text-xs text-muted-foreground"), g.Text("CEO")),
						),
						html.Div(html.Class("text-center"),
							html.Div(html.Class("w-16 h-16 bg-muted rounded-full mx-auto mb-2")),
							html.P(html.Class("text-sm font-medium"), g.Text("Bob Smith")),
							html.P(html.Class("text-xs text-muted-foreground"), g.Text("CTO")),
						),
						html.Div(html.Class("text-center"),
							html.Div(html.Class("w-16 h-16 bg-muted rounded-full mx-auto mb-2")),
							html.P(html.Class("text-sm font-medium"), g.Text("Carol White")),
							html.P(html.Class("text-xs text-muted-foreground"), g.Text("Designer")),
						),
					),
				),
			),
		),
	)
}