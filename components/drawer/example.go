package drawer

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

// Example demonstrates the Drawer component
func Example() g.Node {
	return html.Div(html.Class("space-y-8"),
		// Basic Example
		html.Div(
			html.H3(html.Class("text-lg font-semibold"), g.Text("Basic Drawer")),
			BasicExample(),
		),

		// Different Sides
		html.Div(
			html.H3(html.Class("text-lg font-semibold"), g.Text("Drawer Sides")),
			ExampleSides(),
		),

		// HTMX Example
		html.Div(
			html.H3(html.Class("text-lg font-semibold"), g.Text("HTMX Drawer")),
			ExampleHTMX(),
			html.P(html.Class("text-sm text-muted-foreground mt-2"), 
				g.Text("This drawer uses HTMX for dynamic interactions."),
			),
		),

		// Navigation Drawer
		html.Div(
			html.H3(html.Class("text-lg font-semibold"), g.Text("Navigation Drawer")),
			NavigationDrawerExampleHTMX(),
			html.P(html.Class("text-sm text-muted-foreground mt-2"), 
				g.Text("A drawer used for navigation with HTMX."),
			),
		),
	)
}

// ExampleWithForm demonstrates a drawer with a form
func ExampleWithForm() g.Node {
	return html.Div(
		Trigger(
			TriggerProps{Class: "bg-primary text-primary-foreground hover:bg-primary/90 px-4 py-2 rounded-md"},
			g.Text("Open Form Drawer"),
		),
		New(
			Props{Open: false},
			Overlay(OverlayProps{}),
			ContentComponent(
				ContentProps{},
				"right",
				DrawerHeader(
					HeaderProps{},
					DrawerTitle(TitleProps{}, g.Text("Create Account")),
					DrawerDescription(
						DescriptionProps{},
						g.Text("Enter your details to create a new account."),
					),
				),
				html.Form(html.Class("space-y-4 py-4"),
					html.Div(html.Class("space-y-2"),
						html.Label(html.For("email"), g.Text("Email")),
						html.Input(
							html.Type("email"),
							html.ID("email"),
							html.Name("email"),
							html.Placeholder("m@example.com"),
							html.Required(),
						),
					),
					html.Div(html.Class("space-y-2"),
						html.Label(html.For("password"), g.Text("Password")),
						html.Input(
							html.Type("password"),
							html.ID("password"),
							html.Name("password"),
							html.Required(),
						),
					),
					html.Div(html.Class("space-y-2"),
						html.Label(html.For("confirm"), g.Text("Confirm Password")),
						html.Input(
							html.Type("password"),
							html.ID("confirm"),
							html.Name("confirm"),
							html.Required(),
						),
					),
				),
				DrawerFooter(
					FooterProps{},
					Close(
						CloseProps{Class: "border hover:bg-accent px-4 py-2 rounded-md"},
						g.Text("Cancel"),
					),
					html.Button(
						html.Type("submit"),
						html.Class("bg-primary text-primary-foreground hover:bg-primary/90 px-4 py-2 rounded-md"),
						g.Text("Create Account"),
					),
				),
			),
		),
	)
}

// ExampleScrollable demonstrates a drawer with scrollable content
func ExampleScrollable() g.Node {
	return html.Div(
		Trigger(
			TriggerProps{Class: "border px-4 py-2 rounded-md"},
			g.Text("Open Scrollable Drawer"),
		),
		New(
			Props{Open: false},
			Overlay(OverlayProps{}),
			ContentComponent(
				ContentProps{Class: "flex flex-col max-h-[90vh]"},
				"right",
				DrawerHeader(
					HeaderProps{Class: "flex-shrink-0"},
					DrawerTitle(TitleProps{}, g.Text("Terms of Service")),
					DrawerDescription(
						DescriptionProps{},
						g.Text("Please read our terms carefully."),
					),
				),
				html.Div(html.Class("flex-1 overflow-y-auto py-4"),
					g.Group(g.Map(make([]int, 20), func(i int) g.Node {
						return html.Div(html.Class("mb-4"),
							html.H4(html.Class("font-medium mb-2"), g.Textf("Section %d", i+1)),
							html.P(html.Class("text-sm text-muted-foreground"),
								g.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."),
							),
						)
					})),
				),
				DrawerFooter(
					FooterProps{Class: "flex-shrink-0 border-t pt-4"},
					Close(
						CloseProps{Class: "border hover:bg-accent px-4 py-2 rounded-md"},
						g.Text("Decline"),
					),
					html.Button(
						html.Type("button"),
						html.Class("bg-primary text-primary-foreground hover:bg-primary/90 px-4 py-2 rounded-md"),
						g.Text("Accept"),
					),
				),
			),
		),
	)
}

// ExampleWithoutOverlay demonstrates a drawer without overlay
func ExampleWithoutOverlay() g.Node {
	return html.Div(
		Trigger(
			TriggerProps{Class: "border px-4 py-2 rounded-md"},
			g.Text("Open Drawer (No Overlay)"),
		),
		New(
			Props{Open: false},
			// No overlay component
			ContentComponent(
				ContentProps{},
				"left",
				DrawerHeader(
					HeaderProps{},
					DrawerTitle(TitleProps{}, g.Text("Settings")),
					DrawerDescription(
						DescriptionProps{},
						g.Text("Manage your preferences."),
					),
				),
				html.Div(html.Class("py-4 space-y-4"),
					html.Div(html.Class("flex items-center justify-between"),
						html.Label(html.For("notifications"), g.Text("Notifications")),
						html.Input(
							html.Type("checkbox"),
							html.ID("notifications"),
							html.Class("h-4 w-4"),
						),
					),
					html.Div(html.Class("flex items-center justify-between"),
						html.Label(html.For("marketing"), g.Text("Marketing emails")),
						html.Input(
							html.Type("checkbox"),
							html.ID("marketing"),
							html.Class("h-4 w-4"),
						),
					),
					html.Div(html.Class("flex items-center justify-between"),
						html.Label(html.For("social"), g.Text("Social emails")),
						html.Input(
							html.Type("checkbox"),
							html.ID("social"),
							html.Class("h-4 w-4"),
							html.Checked(),
						),
					),
					html.Div(html.Class("flex items-center justify-between"),
						html.Label(html.For("security"), g.Text("Security emails")),
						html.Input(
							html.Type("checkbox"),
							html.ID("security"),
							html.Class("h-4 w-4"),
							html.Checked(),
							html.Disabled(),
						),
					),
				),
				DrawerFooter(
					FooterProps{},
					Close(
						CloseProps{Class: "w-full border hover:bg-accent px-4 py-2 rounded-md"},
						g.Text("Close"),
					),
				),
			),
		),
	)
}

// ExampleCustomStyling demonstrates custom styling
func ExampleCustomStyling() g.Node {
	return html.Div(
		Trigger(
			TriggerProps{Class: "bg-gradient-to-r from-purple-500 to-pink-500 text-white px-4 py-2 rounded-md"},
			g.Text("Open Styled Drawer"),
		),
		New(
			Props{Open: false},
			Overlay(OverlayProps{Class: "bg-black/60 backdrop-blur-sm"}),
			ContentComponent(
				ContentProps{Class: "bg-gradient-to-br from-purple-50 to-pink-50 dark:from-purple-900/20 dark:to-pink-900/20"},
				"bottom",
				DrawerHeader(
					HeaderProps{},
					DrawerTitle(TitleProps{Class: "text-transparent bg-clip-text bg-gradient-to-r from-purple-600 to-pink-600"}, 
						g.Text("Gradient Drawer"),
					),
					DrawerDescription(
						DescriptionProps{Class: "text-purple-700 dark:text-purple-300"},
						g.Text("A beautifully styled drawer with gradients."),
					),
				),
				html.Div(html.Class("py-6"),
					html.P(html.Class("text-center text-purple-600 dark:text-purple-400"),
						g.Text("This drawer slides up from the bottom with custom styling."),
					),
				),
				DrawerFooter(
					FooterProps{},
					Close(
						CloseProps{Class: "bg-gradient-to-r from-purple-500 to-pink-500 text-white px-6 py-2 rounded-md hover:opacity-90"},
						g.Text("Got it!"),
					),
				),
			),
		),
	)
}