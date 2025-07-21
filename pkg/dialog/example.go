package dialog

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/pkg/button"
)

// DemoBasic shows a basic dialog
func DemoBasic() g.Node {
	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Basic Dialog")),
		Example(),
	)
}

// DemoScrollable shows a scrollable dialog
func DemoScrollable() g.Node {
	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Scrollable Dialog")),
		ExampleScrollable(),
	)
}

// DemoCustom shows a custom styled dialog
func DemoCustom() g.Node {
	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Custom Styled Dialog")),
		ExampleCustom(),
	)
}

// DemoWithForm shows a dialog with a form
func DemoWithForm() g.Node {
	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Dialog with Form")),
		New(
			Props{Open: true},
			Overlay(),
			DialogContent(
				ContentProps{ShowCloseButton: true},
				html.Form(html.Class("space-y-6"),
					DialogHeader(
						HeaderProps{},
						DialogTitle(TitleProps{}, "Create Account"),
						Description(DescriptionProps{}, "Enter your information to create a new account."),
					),
					html.Div(html.Class("space-y-4"),
						html.Div(html.Class("space-y-2"),
							html.Label(html.For("email"), g.Text("Email")),
							html.Input(
								html.Type("email"),
								html.ID("email"),
								html.Placeholder("name@example.com"),
								html.Required(),
								html.Class("w-full"),
							),
						),
						html.Div(html.Class("space-y-2"),
							html.Label(html.For("password"), g.Text("Password")),
							html.Input(
								html.Type("password"),
								html.ID("password"),
								html.Required(),
								html.Class("w-full"),
							),
						),
						html.Div(html.Class("space-y-2"),
							html.Label(html.For("confirm"), g.Text("Confirm Password")),
							html.Input(
								html.Type("password"),
								html.ID("confirm"),
								html.Required(),
								html.Class("w-full"),
							),
						),
					),
					DialogFooter(
						FooterProps{},
						button.New(
							button.Props{Variant: "outline"},
							g.Text("Cancel"),
						),
						button.New(
							button.Props{Type: "submit"},
							g.Text("Create Account"),
						),
					),
				),
			),
		),
	)
}

// DemoSizes shows dialogs in different sizes
func DemoSizes() g.Node {
	return html.Div(html.Class("space-y-8"),
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Dialog Sizes")),
		
		// Small dialog
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-2"), g.Text("Small Dialog")),
			New(
				Props{Open: true},
				Overlay(),
				DialogContent(
					ContentProps{Class: "max-w-sm"},
					DialogHeader(
						HeaderProps{},
						DialogTitle(TitleProps{}, "Small Dialog"),
						Description(DescriptionProps{}, "This is a small dialog window."),
					),
					DialogFooter(
						FooterProps{},
						button.New(button.Props{}, g.Text("Close")),
					),
				),
			),
		),
		
		// Large dialog
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-2 mt-8"), g.Text("Large Dialog")),
			New(
				Props{Open: true},
				Overlay(),
				DialogContent(
					ContentProps{Class: "max-w-3xl"},
					DialogHeader(
						HeaderProps{},
						DialogTitle(TitleProps{}, "Large Dialog"),
						Description(DescriptionProps{}, "This is a large dialog window with more space for content."),
					),
					html.Div(html.Class("grid gap-4 py-4"),
						html.P(html.Class("text-sm text-muted-foreground"),
							g.Text("Large dialogs are useful for displaying more complex content, forms, or data tables."),
						),
					),
					DialogFooter(
						FooterProps{},
						button.New(button.Props{}, g.Text("Close")),
					),
				),
			),
		),
	)
}

// DemoNested shows nested dialogs (note: requires careful JavaScript handling)
func DemoNested() g.Node {
	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Nested Dialog Example")),
		New(
			Props{Open: true},
			Overlay(),
			DialogContent(
				ContentProps{ShowCloseButton: true},
				DialogHeader(
					HeaderProps{},
					DialogTitle(TitleProps{}, "First Dialog"),
					Description(DescriptionProps{}, "This dialog contains a button to open another dialog."),
				),
				html.Div(html.Class("py-6"),
					html.P(html.Class("text-sm text-muted-foreground mb-4"),
						g.Text("Click the button below to open a nested dialog. Note: Nested dialogs require careful JavaScript handling to manage z-index and focus."),
					),
					button.New(
						button.Props{Class: "w-full"},
						g.Text("Open Nested Dialog"),
					),
				),
				DialogFooter(
					FooterProps{},
					button.New(
						button.Props{Variant: "outline"},
						g.Text("Close"),
					),
				),
			),
		),
		// Second dialog (would be shown conditionally with JavaScript)
		html.Div(html.Class("hidden"), // Hidden by default
			New(
				Props{Open: true},
				Overlay("z-[60]"), // Higher z-index
				DialogContent(
					ContentProps{Class: "z-[60]"}, // Higher z-index
					DialogHeader(
						HeaderProps{},
						DialogTitle(TitleProps{}, "Nested Dialog"),
						Description(DescriptionProps{}, "This is a dialog opened from another dialog."),
					),
					DialogFooter(
						FooterProps{},
						button.New(button.Props{}, g.Text("Close")),
					),
				),
			),
		),
	)
}

// DemoConfirmation shows a confirmation dialog
func DemoConfirmation() g.Node {
	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Confirmation Dialog")),
		New(
			Props{Open: true},
			Overlay(),
			DialogContent(
				ContentProps{Class: "max-w-md"},
				html.Div(html.Class("flex gap-4"),
					html.Div(html.Class("flex-shrink-0"),
						html.Div(html.Class("w-10 h-10 rounded-full bg-destructive/10 flex items-center justify-center"),
							// Warning icon
							g.El("svg",
								g.Attr("class", "w-5 h-5 text-destructive"),
								g.Attr("viewBox", "0 0 24 24"),
								g.Attr("fill", "none"),
								g.Attr("stroke", "currentColor"),
								g.Attr("stroke-width", "2"),
								g.El("path",
									g.Attr("stroke-linecap", "round"),
									g.Attr("stroke-linejoin", "round"),
									g.Attr("d", "M12 9v2m0 4html.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"),
								),
							),
						),
					),
					html.Div(html.Class("flex-1"),
						DialogHeader(
							HeaderProps{},
							DialogTitle(TitleProps{}, "Delete Repository"),
							Description(DescriptionProps{}, "This action cannot be undone. This will permanently delete the repository and all of its contents."),
						),
					),
				),
				DialogFooter(
					FooterProps{Class: "mt-6"},
					button.New(
						button.Props{Variant: "outline"},
						g.Text("Cancel"),
					),
					button.New(
						button.Props{Variant: "destructive"},
						g.Text("Delete"),
					),
				),
			),
		),
	)
}