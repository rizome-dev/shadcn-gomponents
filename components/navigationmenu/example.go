package navigationmenu

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

// Examples demonstrates various NavigationMenu usage patterns
func Examples() g.Node {
	return html.Div(
		html.Class("space-y-8 p-8"),
		
		// Simple Navigation Menu
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Simple Navigation")),
			SimpleMenu(),
		),

		// Navigation with Dropdowns
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("With Dropdowns")),
			WithDropdowns(),
		),

		// Mega Menu Example
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Mega Menu")),
			New(
				Props{},
	ListComponent(
					ListProps{},
					Item(
						ItemProps{Value: "products"},
						Trigger(TriggerProps{}, g.Text("Products")),
						ContentComponent(
							ContentProps{},
							html.Div(html.Class("grid gap-3 p-6 md:w-[700px] md:grid-cols-3"),
								html.Div(html.Class("space-y-3"),
									html.H4(html.Class("text-sm font-medium leading-none mb-2"), g.Text("Analytics")),
	LinkComponent(
										LinkProps{Href: "#"},
										html.Div(html.Class("text-sm font-medium leading-none"), g.Text("Real-time Analytics")),
										html.P(html.Class("text-sm text-muted-foreground"), g.Text("Monitor your data in real-time")),
									),
	LinkComponent(
										LinkProps{Href: "#"},
										html.Div(html.Class("text-sm font-medium leading-none"), g.Text("Custom Dashboards")),
										html.P(html.Class("text-sm text-muted-foreground"), g.Text("Build personalized dashboards")),
									),
	LinkComponent(
										LinkProps{Href: "#"},
										html.Div(html.Class("text-sm font-medium leading-none"), g.Text("Reports")),
										html.P(html.Class("text-sm text-muted-foreground"), g.Text("Generate detailed reports")),
									),
								),
								html.Div(html.Class("space-y-3"),
									html.H4(html.Class("text-sm font-medium leading-none mb-2"), g.Text("Automation")),
	LinkComponent(
										LinkProps{Href: "#"},
										html.Div(html.Class("text-sm font-medium leading-none"), g.Text("Workflow Builder")),
										html.P(html.Class("text-sm text-muted-foreground"), g.Text("Create automated workflows")),
									),
	LinkComponent(
										LinkProps{Href: "#"},
										html.Div(html.Class("text-sm font-medium leading-none"), g.Text("API Integration")),
										html.P(html.Class("text-sm text-muted-foreground"), g.Text("Connect with any API")),
									),
	LinkComponent(
										LinkProps{Href: "#"},
										html.Div(html.Class("text-sm font-medium leading-none"), g.Text("Scheduled Tasks")),
										html.P(html.Class("text-sm text-muted-foreground"), g.Text("Schedule recurring tasks")),
									),
								),
								html.Div(html.Class("space-y-3"),
									html.H4(html.Class("text-sm font-medium leading-none mb-2"), g.Text("Security")),
	LinkComponent(
										LinkProps{Href: "#"},
										html.Div(html.Class("text-sm font-medium leading-none"), g.Text("Access Control")),
										html.P(html.Class("text-sm text-muted-foreground"), g.Text("Manage user permissions")),
									),
	LinkComponent(
										LinkProps{Href: "#"},
										html.Div(html.Class("text-sm font-medium leading-none"), g.Text("Audit Logs")),
										html.P(html.Class("text-sm text-muted-foreground"), g.Text("Track all activities")),
									),
	LinkComponent(
										LinkProps{Href: "#"},
										html.Div(html.Class("text-sm font-medium leading-none"), g.Text("2FA")),
										html.P(html.Class("text-sm text-muted-foreground"), g.Text("Two-factor authentication")),
									),
								),
							),
						),
					),
					Item(
						ItemProps{Value: "solutions"},
						Trigger(TriggerProps{}, g.Text("Solutions")),
						ContentComponent(
							ContentProps{},
							html.Div(html.Class("grid gap-3 p-6 md:w-[600px] md:grid-cols-2"),
								html.Div(html.Class("space-y-3"),
									html.H4(html.Class("text-sm font-medium leading-none mb-2"), g.Text("By Industry")),
									LinkComponent(LinkProps{Href: "#"}, g.Text("E-commerce")),
									LinkComponent(LinkProps{Href: "#"}, g.Text("Healthcare")),
									LinkComponent(LinkProps{Href: "#"}, g.Text("Finance")),
									LinkComponent(LinkProps{Href: "#"}, g.Text("Education")),
									LinkComponent(LinkProps{Href: "#"}, g.Text("Real Estate")),
								),
								html.Div(html.Class("space-y-3"),
									html.H4(html.Class("text-sm font-medium leading-none mb-2"), g.Text("By Use Case")),
									LinkComponent(LinkProps{Href: "#"}, g.Text("Data Analytics")),
									LinkComponent(LinkProps{Href: "#"}, g.Text("Marketing Automation")),
									LinkComponent(LinkProps{Href: "#"}, g.Text("Customer Support")),
									LinkComponent(LinkProps{Href: "#"}, g.Text("Sales Management")),
									LinkComponent(LinkProps{Href: "#"}, g.Text("Project Management")),
								),
							),
						),
					),
					Item(
						ItemProps{},
						LinkComponent(LinkProps{Href: "#pricing"}, g.Text("Pricing")),
					),
					Item(
						ItemProps{},
						LinkComponent(LinkProps{Href: "#about"}, g.Text("About")),
					),
				),
			),
		),

		// With Active States
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("With Active States")),
			New(
				Props{},
	ListComponent(
					ListProps{},
					Item(
						ItemProps{},
						LinkComponent(LinkProps{Href: "#", Active: true}, g.Text("Home")),
					),
					Item(
						ItemProps{},
						LinkComponent(LinkProps{Href: "#"}, g.Text("Features")),
					),
					Item(
						ItemProps{},
						LinkComponent(LinkProps{Href: "#"}, g.Text("Pricing")),
					),
					Item(
						ItemProps{},
						LinkComponent(LinkProps{Href: "#", Disabled: true}, g.Text("Coming Soon")),
					),
				),
			),
		),

		// Content Grid Layout
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Content Grid Layout")),
			New(
				Props{},
	ListComponent(
					ListProps{},
					Item(
						ItemProps{Value: "features"},
						Trigger(TriggerProps{}, g.Text("Features")),
						ContentComponent(
							ContentProps{},
							html.Div(html.Class("w-[600px] p-4"),
								html.H3(html.Class("mb-4 text-lg font-medium"), g.Text("Platform Features")),
								html.Div(html.Class("grid grid-cols-3 gap-4"),
									html.Div(html.Class("space-y-2"),
										html.H4(html.Class("text-sm font-medium"), g.Text("Core")),
										html.Ul(html.Class("space-y-1 text-sm"),
											html.Li(LinkComponent(LinkProps{Href: "#"}, g.Text("Dashboard"))),
											html.Li(LinkComponent(LinkProps{Href: "#"}, g.Text("Analytics"))),
											html.Li(LinkComponent(LinkProps{Href: "#"}, g.Text("Reports"))),
										),
									),
									html.Div(html.Class("space-y-2"),
										html.H4(html.Class("text-sm font-medium"), g.Text("Advanced")),
										html.Ul(html.Class("space-y-1 text-sm"),
											html.Li(LinkComponent(LinkProps{Href: "#"}, g.Text("Automations"))),
											html.Li(LinkComponent(LinkProps{Href: "#"}, g.Text("Integrations"))),
											html.Li(LinkComponent(LinkProps{Href: "#"}, g.Text("API Access"))),
										),
									),
									html.Div(html.Class("space-y-2"),
										html.H4(html.Class("text-sm font-medium"), g.Text("Support")),
										html.Ul(html.Class("space-y-1 text-sm"),
											html.Li(LinkComponent(LinkProps{Href: "#"}, g.Text("Documentation"))),
											html.Li(LinkComponent(LinkProps{Href: "#"}, g.Text("Tutorials"))),
											html.Li(LinkComponent(LinkProps{Href: "#"}, g.Text("Contact"))),
										),
									),
								),
							),
						),
					),
				),
			),
		),

		// With Icons
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("With Icons")),
			New(
				Props{},
	ListComponent(
					ListProps{},
					Item(
						ItemProps{},
	LinkComponent(
							LinkProps{Href: "#"},
							html.Div(html.Class("flex items-center gap-2"),
								g.Raw(`<svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg">
									<path d="M7.5 0.875C3.99187 0.875 1.125 3.74187 1.125 7.25C1.125 10.7581 3.99187 13.625 7.5 13.625C11.0081 13.625 13.875 10.7581 13.875 7.25C13.875 3.74187 11.0081 0.875 7.5 0.875ZM2.625 7.25C2.625 4.57031 4.82031 2.375 7.5 2.375C10.1797 2.375 12.375 4.57031 12.375 7.25C12.375 9.92969 10.1797 12.125 7.5 12.125C4.82031 12.125 2.625 9.92969 2.625 7.25Z" fill="currentColor"></path>
									<path d="M7.5 4.5C7.08579 4.5 6.75 4.83579 6.75 5.25V7.75C6.75 8.16421 7.08579 8.5 7.5 8.5C7.91421 8.5 8.25 8.16421 8.25 7.75V5.25C8.25 4.83579 7.91421 4.5 7.5 4.5Z" fill="currentColor"></path>
									<path d="M7.5 9.5C7.22386 9.5 7 9.72386 7 10C7 10.2761 7.22386 10.5 7.5 10.5C7.77614 10.5 8 10.2761 8 10C8 9.72386 7.77614 9.5 7.5 9.5Z" fill="currentColor"></path>
								</svg>`),
								g.Text("Dashboard"),
							),
						),
					),
					Item(
						ItemProps{Value: "settings"},
						Trigger(
							TriggerProps{},
							html.Div(html.Class("flex items-center gap-2"),
								g.Raw(`<svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg">
									<path d="M7.5 0.5C7.77614 0.5 8 0.723858 8 1V2.0738C8.91096 2.20173 9.74767 2.56092 10.4446 3.08679L11.2082 2.32318C11.4015 2.12985 11.7163 2.12985 11.9097 2.32318L12.6768 3.0903C12.8701 3.28363 12.8701 3.59849 12.6768 3.79182L11.9132 4.55544C12.4391 5.25233 12.7983 6.08904 12.9262 7H14C14.2761 7 14.5 7.22386 14.5 7.5C14.5 7.77614 14.2761 8 14 8H12.9262C12.7983 8.91096 12.4391 9.74767 11.9132 10.4446L12.6768 11.2082C12.8701 11.4015 12.8701 11.7164 12.6768 11.9097L11.9097 12.6768C11.7163 12.8701 11.4015 12.8701 11.2082 12.6768L10.4446 11.9132C9.74767 12.4391 8.91096 12.7983 8 12.9262V14C8 14.2761 7.77614 14.5 7.5 14.5C7.22386 14.5 7 14.2761 7 14V12.9262C6.08904 12.7983 5.25233 12.4391 4.55544 11.9132L3.79182 12.6768C3.59849 12.8701 3.28363 12.8701 3.0903 12.6768L2.32318 11.9097C2.12985 11.7164 2.12985 11.4015 2.32318 11.2082L3.08679 10.4446C2.56092 9.74767 2.20173 8.91096 2.0738 8H1C0.723858 8 0.5 7.77614 0.5 7.5C0.5 7.22386 0.723858 7 1 7H2.0738C2.20173 6.08904 2.56092 5.25233 3.08679 4.55544L2.32318 3.79182C2.12985 3.59849 2.12985 3.28363 2.32318 3.0903L3.0903 2.32318C3.28363 2.12985 3.59849 2.12985 3.79182 2.32318L4.55544 3.08679C5.25233 2.56092 6.08904 2.20173 7 2.0738V1C7 0.723858 7.22386 0.5 7.5 0.5ZM7.5 5C6.11929 5 5 6.11929 5 7.5C5 8.88071 6.11929 10 7.5 10C8.88071 10 10 8.88071 10 7.5C10 6.11929 8.88071 5 7.5 5Z" fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"></path>
								</svg>`),
								g.Text("Settings"),
							),
						),
						ContentComponent(
							ContentProps{},
							html.Ul(html.Class("grid gap-3 p-4 md:w-[300px]"),
								ListItem("Profile", "#", "Manage your profile settings"),
								ListItem("Account", "#", "Configure account preferences"),
								ListItem("Security", "#", "Update security settings"),
								ListItem("Notifications", "#", "Manage notification preferences"),
							),
						),
					),
				),
			),
		),

		// Compact Navigation
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Compact Navigation")),
			New(
				Props{Class: "h-8"},
	ListComponent(
					ListProps{},
					Item(
						ItemProps{},
						LinkComponent(LinkProps{Href: "#", Class: "h-8 px-3 py-1.5 text-xs"}, g.Text("Home")),
					),
					Item(
						ItemProps{},
						LinkComponent(LinkProps{Href: "#", Class: "h-8 px-3 py-1.5 text-xs"}, g.Text("About")),
					),
					Item(
						ItemProps{},
						LinkComponent(LinkProps{Href: "#", Class: "h-8 px-3 py-1.5 text-xs"}, g.Text("Services")),
					),
					Item(
						ItemProps{},
						LinkComponent(LinkProps{Href: "#", Class: "h-8 px-3 py-1.5 text-xs"}, g.Text("Contact")),
					),
				),
			),
		),

		// Centered Navigation
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Centered Navigation")),
			html.Div(
				html.Class("flex justify-center"),
				New(
					Props{},
	ListComponent(
						ListProps{},
						Item(ItemProps{}, LinkComponent(LinkProps{Href: "#"}, g.Text("Home"))),
						Item(ItemProps{}, LinkComponent(LinkProps{Href: "#"}, g.Text("Features"))),
						Item(ItemProps{}, LinkComponent(LinkProps{Href: "#"}, g.Text("Pricing"))),
						Item(ItemProps{}, LinkComponent(LinkProps{Href: "#"}, g.Text("About"))),
						Item(ItemProps{}, LinkComponent(LinkProps{Href: "#"}, g.Text("Contact"))),
					),
				),
			),
		),

		// With Badges
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("With Badges")),
			New(
				Props{},
	ListComponent(
					ListProps{},
					Item(
						ItemProps{},
	LinkComponent(
							LinkProps{Href: "#"},
							html.Div(html.Class("flex items-center gap-2"),
								g.Text("Messages"),
								html.Span(html.Class("flex h-5 w-5 items-center justify-center rounded-full bg-primary text-[10px] font-medium text-primary-foreground"),
									g.Text("5"),
								),
							),
						),
					),
					Item(
						ItemProps{},
	LinkComponent(
							LinkProps{Href: "#"},
							html.Div(html.Class("flex items-center gap-2"),
								g.Text("Notifications"),
								html.Span(html.Class("flex h-2 w-2 rounded-full bg-destructive")),
							),
						),
					),
					Item(
						ItemProps{},
	LinkComponent(
							LinkProps{Href: "#"},
							html.Div(html.Class("flex items-center gap-2"),
								g.Text("Updates"),
								html.Span(html.Class("rounded bg-green-50 px-1.5 py-0.5 text-[10px] font-medium text-green-700 dark:bg-green-900/30 dark:text-green-400"),
									g.Text("NEW"),
								),
							),
						),
					),
				),
			),
		),
	)
}