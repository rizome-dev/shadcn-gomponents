package sheet

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib/icons"
)

// Example demonstrates how to use the Sheet component
func Example() g.Node {
	return html.Div(
		html.Class("space-y-8 p-8"),
		
		// Basic sheet from right
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Basic Sheet (Right)")),
			Trigger(TriggerProps{}, g.Text("Open Sheet")),
			RightSheet(
				Props{Open: false},
				HeaderComponent(
					HeaderProps{},
					TitleComponent(TitleProps{}, g.Text("Edit Profile")),
					Description(DescriptionProps{}, g.Text("Make changes to your profile here. Click save when you're done.")),
				),
				html.Div(html.Class("grid gap-4 py-4"),
					html.Div(html.Class("grid grid-cols-4 items-center gap-4"),
						html.Label(html.For("name"), html.Class("text-right"), g.Text("Name")),
						html.Input(
							html.ID("name"),
							html.Value("Pedro Duarte"),
							html.Class("col-span-3"),
						),
					),
					html.Div(html.Class("grid grid-cols-4 items-center gap-4"),
						html.Label(html.For("username"), html.Class("text-right"), g.Text("Username")),
						html.Input(
							html.ID("username"),
							html.Value("@peduarte"),
							html.Class("col-span-3"),
						),
					),
				),
				FooterComponent(
					FooterProps{},
					Close(CloseProps{Class: "border hover:bg-accent"}, g.Text("Cancel")),
					html.Button(html.Type("submit"), html.Class("bg-primary text-primary-foreground hover:bg-primary/90"), g.Text("Save changes")),
				),
			),
		),
		
		// Sheet from different sides
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Sheet Positions")),
			html.Div(html.Class("flex flex-wrap gap-4"),
				// Left sheet
				html.Div(
					Trigger(TriggerProps{}, g.Text("From Left")),
					LeftSheet(
						Props{Open: false},
						HeaderComponent(
							HeaderProps{},
							TitleComponent(TitleProps{}, g.Text("Navigation")),
						),
						html.Nav(html.Class("grid gap-4 py-4"),
							html.A(html.Href("#"), html.Class("flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground hover:text-primary"),
								icons.Home(html.Class("h-4 w-4")),
								g.Text("Dashboard"),
							),
							html.A(html.Href("#"), html.Class("flex items-center gap-3 rounded-lg bg-muted px-3 py-2 text-primary"),
								icons.Package(html.Class("h-4 w-4")),
								g.Text("Products"),
							),
							html.A(html.Href("#"), html.Class("flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground hover:text-primary"),
								icons.Users(html.Class("h-4 w-4")),
								g.Text("Customers"),
							),
							html.A(html.Href("#"), html.Class("flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground hover:text-primary"),
								icons.Settings(html.Class("h-4 w-4")),
								g.Text("Settings"),
							),
						),
					),
				),
				
				// Top sheet
				html.Div(
					Trigger(TriggerProps{}, g.Text("From Top")),
					TopSheet(
						Props{Open: false},
						html.Div(html.Class("mx-auto max-w-2xl"),
							HeaderComponent(
								HeaderProps{},
								TitleComponent(TitleProps{}, g.Text("Notifications")),
								Description(DescriptionProps{}, g.Text("You have 3 unread messages.")),
							),
							html.Div(html.Class("grid gap-4 py-4"),
								html.Div(html.Class("flex items-start gap-4 rounded-lg border p-4"),
									html.Div(html.Class("h-2 w-2 translate-y-1.5 rounded-full bg-blue-500")),
									html.Div(html.Class("grid gap-1"),
										html.P(html.Class("text-sm font-medium"), g.Text("New message from Alice")),
										html.P(html.Class("text-sm text-muted-foreground"), g.Text("Hey, are you available for a call?")),
										html.P(html.Class("text-xs text-muted-foreground"), g.Text("5 minutes ago")),
									),
								),
								html.Div(html.Class("flex items-start gap-4 rounded-lg border p-4"),
									html.Div(html.Class("h-2 w-2 translate-y-1.5 rounded-full bg-green-500")),
									html.Div(html.Class("grid gap-1"),
										html.P(html.Class("text-sm font-medium"), g.Text("Your order has been shipped")),
										html.P(html.Class("text-sm text-muted-foreground"), g.Text("Track your package with order #12345")),
										html.P(html.Class("text-xs text-muted-foreground"), g.Text("2 hours ago")),
									),
								),
							),
						),
					),
				),
				
				// Bottom sheet
				html.Div(
					Trigger(TriggerProps{}, g.Text("From Bottom")),
					BottomSheet(
						Props{Open: false},
						html.Div(html.Class("mx-auto max-w-2xl"),
							HeaderComponent(
								HeaderProps{Class: "text-center"},
								TitleComponent(TitleProps{}, g.Text("Share")),
								Description(DescriptionProps{}, g.Text("Share this link with others")),
							),
							html.Div(html.Class("grid gap-4 py-4"),
								html.Div(html.Class("flex items-center gap-2"),
									html.Input(
										html.Value("https://example.com/share/abc123"),
										html.ReadOnly(),
										html.Class("flex-1"),
									),
									html.Button(
										html.Class("shrink-0"),
										icons.Copy(html.Class("h-4 w-4")),
									),
								),
								html.Div(html.Class("grid grid-cols-4 gap-4"),
									html.Button(html.Class("flex flex-col items-center gap-2 border p-4"),
										icons.User(html.Class("h-8 w-8")),
										html.Span(html.Class("text-xs"), g.Text("Email")),
									),
									html.Button(html.Class("flex flex-col items-center gap-2 border p-4"),
										icons.User(html.Class("h-8 w-8")),
										html.Span(html.Class("text-xs"), g.Text("Twitter")),
									),
									html.Button(html.Class("flex flex-col items-center gap-2 border p-4"),
										icons.User(html.Class("h-8 w-8")),
										html.Span(html.Class("text-xs"), g.Text("Facebook")),
									),
									html.Button(html.Class("flex flex-col items-center gap-2 border p-4"),
										icons.User(html.Class("h-8 w-8")),
										html.Span(html.Class("text-xs"), g.Text("LinkedIn")),
									),
								),
							),
						),
					),
				),
			),
		),
		
		// Mobile-optimized sheet
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Mobile Sheet")),
			html.P(html.Class("text-sm text-muted-foreground mb-4"), 
				g.Text("This sheet appears from bottom on mobile and right on desktop."),
			),
			Trigger(TriggerProps{}, g.Text("Open Mobile Sheet")),
			MobileSheet(
				Props{Open: false},
				HeaderComponent(
					HeaderProps{},
					TitleComponent(TitleProps{}, g.Text("Mobile Optimized")),
					Description(DescriptionProps{}, g.Text("This sheet adapts to different screen sizes.")),
				),
				html.Div(html.Class("py-4"),
					html.P(html.Class("text-sm"), g.Text("On mobile devices, this sheet slides up from the bottom. On desktop, it slides in from the right side.")),
				),
			),
		),
		
		// Sheet with form
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Sheet with Form")),
			Trigger(TriggerProps{}, g.Text("Create New Item")),
			WithForm(
				Props{Open: false},
				ContentProps{Side: "right", ShowCloseButton: true},
				"/api/items",
				HeaderComponent(
					HeaderProps{},
					TitleComponent(TitleProps{}, g.Text("Create New Item")),
					Description(DescriptionProps{}, g.Text("Fill in the details below to create a new item.")),
				),
				html.Div(html.Class("grid gap-4 py-4"),
					html.Div(html.Class("grid gap-2"),
						html.Label(html.For("item-name"), g.Text("Name")),
						html.Input(
							html.ID("item-name"),
							html.Name("name"),
							html.Placeholder("Enter item name"),
							html.Required(),
						),
					),
					html.Div(html.Class("grid gap-2"),
						html.Label(html.For("item-category"), g.Text("Category")),
						html.Select(
							html.ID("item-category"),
							html.Name("category"),
							html.Option(html.Value(""), g.Text("Select a category")),
							html.Option(html.Value("electronics"), g.Text("Electronics")),
							html.Option(html.Value("clothing"), g.Text("Clothing")),
							html.Option(html.Value("food"), g.Text("Food")),
							html.Option(html.Value("other"), g.Text("Other")),
						),
					),
					html.Div(html.Class("grid gap-2"),
						html.Label(html.For("item-description"), g.Text("Description")),
						html.Textarea(
							html.ID("item-description"),
							html.Name("description"),
							html.Placeholder("Enter item description"),
							html.Class("min-h-[100px]"),
						),
					),
				),
				FooterComponent(
					FooterProps{},
					Close(CloseProps{Class: "border hover:bg-accent"}, g.Text("Cancel")),
					html.Button(html.Type("submit"), html.Class("bg-primary text-primary-foreground hover:bg-primary/90"), g.Text("Create")),
				),
			),
		),
		
		// Complex sheet content
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Complex Sheet Content")),
			Trigger(TriggerProps{}, g.Text("View Details")),
			RightSheet(
				Props{Open: false},
				html.Div(html.Class("flex flex-col h-full"),
					HeaderComponent(
						HeaderProps{},
						TitleComponent(TitleProps{}, g.Text("Order Details")),
						Description(DescriptionProps{}, g.Text("Order #12345 • Placed on Jan 15, 2024")),
					),
					
					// Scrollable content
					html.Div(html.Class("flex-1 overflow-y-auto py-4 -mx-6 px-6"),
						// Order status
						html.Div(html.Class("mb-6"),
							html.H3(html.Class("text-sm font-medium mb-3"), g.Text("Status")),
							html.Div(html.Class("flex items-center gap-2"),
								html.Div(html.Class("h-2 w-2 rounded-full bg-green-500")),
								html.Span(html.Class("text-sm"), g.Text("Delivered")),
							),
						),
						
						// Customer info
						html.Div(html.Class("mb-6"),
							html.H3(html.Class("text-sm font-medium mb-3"), g.Text("Customer")),
							html.Div(html.Class("space-y-1 text-sm"),
								html.P(g.Text("John Doe")),
								html.P(html.Class("text-muted-foreground"), g.Text("john.doe@example.com")),
								html.P(html.Class("text-muted-foreground"), g.Text("+1 (555) 123-4567")),
							),
						),
						
						// Shipping address
						html.Div(html.Class("mb-6"),
							html.H3(html.Class("text-sm font-medium mb-3"), g.Text("Shipping Address")),
							html.Div(html.Class("space-y-1 text-sm text-muted-foreground"),
								html.P(g.Text("123 Main Street")),
								html.P(g.Text("Apt 4B")),
								html.P(g.Text("New York, NY 10001")),
								html.P(g.Text("United States")),
							),
						),
						
						// Order items
						html.Div(html.Class("mb-6"),
							html.H3(html.Class("text-sm font-medium mb-3"), g.Text("Items")),
							html.Div(html.Class("space-y-3"),
								html.Div(html.Class("flex items-center gap-3"),
									html.Div(html.Class("h-12 w-12 rounded bg-muted")),
									html.Div(html.Class("flex-1"),
										html.P(html.Class("text-sm font-medium"), g.Text("Product Name")),
										html.P(html.Class("text-sm text-muted-foreground"), g.Text("Qty: 2 × $29.99")),
									),
									html.P(html.Class("text-sm font-medium"), g.Text("$59.98")),
								),
								html.Div(html.Class("flex items-center gap-3"),
									html.Div(html.Class("h-12 w-12 rounded bg-muted")),
									html.Div(html.Class("flex-1"),
										html.P(html.Class("text-sm font-medium"), g.Text("Another Product")),
										html.P(html.Class("text-sm text-muted-foreground"), g.Text("Qty: 1 × $49.99")),
									),
									html.P(html.Class("text-sm font-medium"), g.Text("$49.99")),
								),
							),
						),
						
						// Order summary
						html.Div(html.Class("border-t pt-4"),
							html.Div(html.Class("space-y-2"),
								html.Div(html.Class("flex justify-between text-sm"),
									html.Span(g.Text("Subtotal")),
									html.Span(g.Text("$109.97")),
								),
								html.Div(html.Class("flex justify-between text-sm"),
									html.Span(g.Text("Shipping")),
									html.Span(g.Text("$5.00")),
								),
								html.Div(html.Class("flex justify-between text-sm"),
									html.Span(g.Text("Tax")),
									html.Span(g.Text("$9.20")),
								),
								html.Div(html.Class("flex justify-between font-medium"),
									html.Span(g.Text("Total")),
									html.Span(g.Text("$124.17")),
								),
							),
						),
					),
					
					// Footer actions
					FooterComponent(
						FooterProps{Class: "border-t pt-4"},
						Close(CloseProps{Class: "border hover:bg-accent"}, g.Text("Close")),
						html.Button(html.Class("bg-primary text-primary-foreground hover:bg-primary/90"), g.Text("Print Invoice")),
					),
				),
			),
		),
		
		// HTMX Examples
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("HTMX-Enhanced Sheets")),
			html.P(html.Class("text-sm text-muted-foreground mb-4"), 
				g.Text("These examples require HTMX and server endpoints to be set up."),
			),
			html.Div(html.Class("flex flex-wrap gap-4"),
				ExampleHTMX(),
				MobileMenuExampleHTMX(),
				FilterSheetExampleHTMX(),
				SettingsSheetExampleHTMX(),
			),
		),
	)
}