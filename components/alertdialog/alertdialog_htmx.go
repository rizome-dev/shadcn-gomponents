package alertdialog

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	hx "maragu.dev/gomponents-htmx"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// HTMXProps defines HTMX-specific properties for the AlertDialog
type HTMXProps struct {
	ID          string // Unique ID for the dialog
	TriggerPath string // Server path for trigger actions
	ClosePath   string // Server path for close actions
}

// TriggerProps defines properties for trigger button
type TriggerProps struct {
	Class string // Additional CSS classes
}

// NewHTMX creates an HTMX-enhanced AlertDialog component
func NewHTMX(props Props, htmxProps HTMXProps, children ...g.Node) g.Node {
	if props.Open {
		classes := lib.CN("fixed inset-0 z-50", props.Class)
		return html.Div(
			html.ID(htmxProps.ID),
			html.Class(classes),
			g.Group(children),
		)
	}
	// Return empty div that can be replaced by HTMX
	return html.Div(html.ID(htmxProps.ID))
}

// TriggerHTMX creates an HTMX-enhanced trigger button
func TriggerHTMX(props TriggerProps, htmxProps HTMXProps, children ...g.Node) g.Node {
	classes := lib.CN(props.Class)
	
	return html.Button(
		html.Type("button"),
		g.If(classes != "", html.Class(classes)),
		hx.Get(htmxProps.TriggerPath),
		html.Target("#" + htmxProps.ID),
		hx.Swap("outerHTML"),
		g.Group(children),
	)
}

// DialogOverlayHTMX creates an HTMX-enhanced overlay with close functionality
func DialogOverlayHTMX(htmxProps HTMXProps, class ...string) g.Node {
	classes := lib.CN(
		"fixed inset-0 z-50 bg-black/80",
		"data-[state=open]:animate-in data-[state=closed]:animate-out",
		"data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0",
		lib.CN(class...),
	)

	return html.Div(
		html.Class(classes),
		hx.Get(htmxProps.ClosePath),
		html.Target("#" + htmxProps.ID),
		hx.Swap("outerHTML"),
		hx.Trigger("click"),
	)
}

// CancelHTMX creates an HTMX-enhanced cancel button
func CancelHTMX(props ActionProps, htmxProps HTMXProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors",
		"focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2",
		"disabled:pointer-events-none disabled:opacity-50",
		"border border-input bg-background hover:bg-accent hover:text-accent-foreground",
		"h-10 px-4 py-2",
		props.Class,
	)

	return html.Button(
		html.Type("button"),
		html.Class(classes),
		hx.Get(htmxProps.ClosePath),
		html.Target("#" + htmxProps.ID),
		hx.Swap("outerHTML"),
		g.Group(children),
	)
}

// ActionHTMX creates an HTMX-enhanced action button
func ActionHTMX(props ActionProps, htmxProps HTMXProps, actionPath string, children ...g.Node) g.Node {
	classes := lib.CN(
		"inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors",
		"focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2",
		"disabled:pointer-events-none disabled:opacity-50",
		"bg-primary text-primary-foreground hover:bg-primary/90",
		"h-10 px-4 py-2",
		props.Class,
	)

	return html.Button(
		html.Type("button"),
		html.Class(classes),
		hx.Post(actionPath),
		html.Target("#" + htmxProps.ID),
		hx.Swap("outerHTML"),
		g.Group(children),
	)
}

// ExampleHTMX creates an HTMX-enhanced alert dialog example
func ExampleHTMX() g.Node {
	htmxProps := HTMXProps{
		ID:          "alert-dialog-example",
		TriggerPath: "/api/alert-dialog/open",
		ClosePath:   "/api/alert-dialog/close",
	}
	
	return html.Div(
		// Trigger button
		TriggerHTMX(
			TriggerProps{},
			htmxProps,
			g.Text("Show Alert Dialog"),
		),
		// Dialog placeholder
		html.Div(html.ID(htmxProps.ID)),
	)
}

// RenderOpenDialog renders an open alert dialog (for server response)
func RenderOpenDialog(htmxProps HTMXProps) g.Node {
	return NewHTMX(
		Props{Open: true},
		htmxProps,
		DialogOverlayHTMX(htmxProps),
		DialogContent(
			ContentProps{},
			DialogHeader(
				HeaderProps{},
				DialogTitle(TitleProps{}, "Are you absolutely sure?"),
				DialogDescription(
					DescriptionProps{},
					"This action cannot be undone. This will permanently delete your account and remove your data from our servers.",
				),
			),
			DialogFooter(
				FooterProps{},
				CancelHTMX(ActionProps{}, htmxProps, g.Text("Cancel")),
				ActionHTMX(
					ActionProps{}, 
					htmxProps,
					"/api/alert-dialog/confirm",
					g.Text("Continue"),
				),
			),
		),
	)
}

// RenderClosedDialog renders a closed alert dialog (for server response)
func RenderClosedDialog(htmxProps HTMXProps) g.Node {
	return html.Div(html.ID(htmxProps.ID))
}

// DeleteAccountExampleHTMX creates a delete account alert dialog with HTMX
func DeleteAccountExampleHTMX() g.Node {
	htmxProps := HTMXProps{
		ID:          "delete-account-dialog",
		TriggerPath: "/api/alert-dialog/delete-account/open",
		ClosePath:   "/api/alert-dialog/delete-account/close",
	}
	
	return html.Div(
		TriggerHTMX(
			TriggerProps{Class: "bg-destructive text-destructive-foreground hover:bg-destructive/90"},
			htmxProps,
			g.Text("Delete Account"),
		),
		html.Div(html.ID(htmxProps.ID)),
	)
}

// RenderDeleteAccountDialog renders the delete account dialog (for server response)
func RenderDeleteAccountDialog(htmxProps HTMXProps) g.Node {
	return NewHTMX(
		Props{Open: true},
		htmxProps,
		DialogOverlayHTMX(htmxProps),
		DialogContent(
			ContentProps{Class: "sm:max-w-[425px]"},
			DialogHeader(
				HeaderProps{},
				DialogTitle(TitleProps{Class: "text-destructive"}, "Delete Account"),
				DialogDescription(
					DescriptionProps{},
					"This action cannot be undone. This will permanently delete your account and remove all of your data from our servers.",
				),
			),
			html.Div(html.Class("py-4"),
				html.P(html.Class("text-sm text-muted-foreground"),
					g.Text("All your projects, files, and settings will be permanently deleted. You will not be able to recover your account."),
				),
				html.Div(html.Class("mt-4 p-4 rounded-md bg-destructive/10 border border-destructive/20"),
					html.P(html.Class("text-sm font-medium text-destructive"),
						g.Text("Type 'DELETE' to confirm:"),
					),
					html.Input(
						html.Type("text"),
						html.Class("mt-2 w-full"),
						html.Placeholder("DELETE"),
						html.ID("delete-confirmation"),
						hx.Post("/api/alert-dialog/delete-account/validate"),
						hx.Trigger("keyup changed delay:500ms"),
						html.Target("#delete-button"),
						hx.Swap("outerHTML"),
					),
				),
			),
			DialogFooter(
				FooterProps{},
				CancelHTMX(ActionProps{}, htmxProps, g.Text("Cancel")),
				html.Div(html.ID("delete-button"),
					html.Button(
						html.Type("button"),
						html.Class("inline-flex items-center justify-center rounded-md text-sm font-medium h-10 px-4 py-2 bg-gray-300 text-gray-500 cursor-not-allowed"),
						html.Disabled(),
						g.Text("Delete Account"),
					),
				),
			),
		),
	)
}

// RenderDeleteButton renders the delete button based on validation (for server response)
func RenderDeleteButton(htmxProps HTMXProps, isValid bool) g.Node {
	if isValid {
		return ActionHTMX(
			ActionProps{Class: "bg-destructive text-destructive-foreground hover:bg-destructive/90"},
			htmxProps,
			"/api/alert-dialog/delete-account/confirm",
			g.Text("Delete Account"),
		)
	}
	
	return html.Button(
		html.Type("button"),
		html.Class("inline-flex items-center justify-center rounded-md text-sm font-medium h-10 px-4 py-2 bg-gray-300 text-gray-500 cursor-not-allowed"),
		html.Disabled(),
		g.Text("Delete Account"),
	)
}