package alertdialog

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

// ExampleBasic shows a basic alert dialog
func ExampleBasic() g.Node {
	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Basic Alert Dialog")),
		New(
			Props{Open: true},
			DialogOverlay(),
			DialogContent(
				ContentProps{},
				DialogHeader(
					HeaderProps{},
					DialogTitle(TitleProps{}, "Confirm Action"),
					DialogDescription(DescriptionProps{}, "Are you sure you want to continue?"),
				),
				DialogFooter(
					FooterProps{},
					DialogCancel(CancelProps{}, g.Text("Cancel")),
					DialogAction(ActionProps{}, g.Text("Confirm")),
				),
			),
		),
	)
}

// ExampleDestructive shows a destructive alert dialog
func ExampleDestructive() g.Node {
	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Destructive Alert Dialog")),
		New(
			Props{Open: true},
			DialogOverlay(),
			DialogContent(
				ContentProps{},
				DialogHeader(
					HeaderProps{},
					DialogTitle(TitleProps{}, "Delete Account"),
					DialogDescription(DescriptionProps{}, "This action cannot be undone. This will permanently delete your account and remove your data from our servers."),
				),
				DialogFooter(
					FooterProps{},
					DialogCancel(CancelProps{}, g.Text("Cancel")),
					DialogAction(ActionProps{Class: "bg-destructive hover:bg-destructive/90"}, g.Text("Delete Account")),
				),
			),
		),
	)
}

// ExampleWithLink shows an alert dialog with a link action
func ExampleWithLink() g.Node {
	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Alert Dialog with Link")),
		New(
			Props{Open: true},
			DialogOverlay(),
			DialogContent(
				ContentProps{},
				DialogHeader(
					HeaderProps{},
					DialogTitle(TitleProps{}, "Terms Updated"),
					DialogDescription(DescriptionProps{}, "Our terms of service have been updated. Please review the changes before continuing."),
				),
				DialogFooter(
					FooterProps{},
					DialogCancel(CancelProps{}, g.Text("Later")),
					DialogAction(ActionProps{Href: "/terms"}, g.Text("Review Terms")),
				),
			),
		),
	)
}

// ExampleCustomStyling shows an alert dialog with custom styling
func ExampleCustomStyling() g.Node {
	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Custom Styled Alert Dialog")),
		New(
			Props{Open: true, Class: "backdrop-blur-sm"},
			DialogOverlay("bg-blue-900/20"),
			DialogContent(
				ContentProps{Class: "border-blue-500"},
				DialogHeader(
					HeaderProps{Class: "text-blue-900"},
					DialogTitle(TitleProps{Class: "text-blue-900"}, "Information"),
					DialogDescription(DescriptionProps{}, "This is an informational alert with custom styling."),
				),
				DialogFooter(
					FooterProps{},
					DialogAction(ActionProps{Class: "bg-blue-600 hover:bg-blue-700"}, g.Text("Got it")),
				),
			),
		),
	)
}