package dialog

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
	"github.com/rizome-dev/shadcn-gomponents/lib/icons"
)

// Props defines the properties for the Dialog component
type Props struct {
	Open  bool   // Whether the dialog is open (for server-side rendering)
	Class string // Additional custom classes
}

// ContentProps defines the properties for the DialogContent
type ContentProps struct {
	Class            string
	ShowCloseButton  bool // Whether to show the close button
}

// HeaderProps defines the properties for the DialogHeader
type HeaderProps struct {
	Class string
}

// FooterProps defines the properties for the DialogFooter
type FooterProps struct {
	Class string
}

// TitleProps defines the properties for the DialogTitle
type TitleProps struct {
	Class string
}

// DescriptionProps defines the properties for the DialogDescription
type DescriptionProps struct {
	Class string
}

// TriggerProps defines the properties for the DialogTrigger
type TriggerProps struct {
	Class string
}

// CloseProps defines the properties for the DialogClose
type CloseProps struct {
	Class string
}

// New creates a new Dialog component
// Note: This is a static implementation for server-side rendering
// For full interactivity, JavaScript would be needed
func New(props Props, children ...g.Node) g.Node {
	if !props.Open {
		// If not open, don't render anything
		return g.Text("")
	}

	classes := lib.CN(
		"fixed inset-0 z-50",
		props.Class,
	)

	return html.Div(
		append([]g.Node{html.Class(classes)}, children...)...,
	)
}

// Overlay creates the Dialog overlay
func Overlay(class ...string) g.Node {
	classes := lib.CN(
		"fixed inset-0 z-50 bg-black/80",
		lib.CN(class...),
	)

	return html.Div(html.Class(classes))
}

// DialogContent creates the Dialog content container
func DialogContent(props ContentProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"fixed left-[50%] top-[50%] z-50 grid w-full max-w-lg translate-x-[-50%] translate-y-[-50%] gap-4 border bg-background p-6 shadow-lg sm:rounded-lg",
		props.Class,
	)

	contentChildren := children
	if props.ShowCloseButton {
		closeButton := html.Button(
			html.Type("button"),
			html.Class("absolute right-4 top-4 rounded-sm opacity-70 ring-offset-background transition-opacity hover:opacity-100 focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 disabled:pointer-events-none"),
			icons.X(html.Class("h-4 w-4")),
			html.Span(html.Class("sr-only"), g.Text("Close")),
		)
		contentChildren = append([]g.Node{closeButton}, children...)
	}

	return html.Div(
		html.Class(classes),
		g.Group(contentChildren),
	)
}

// DialogHeader creates the Dialog header
func DialogHeader(props HeaderProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"flex flex-col space-y-1.5 text-center sm:text-left",
		props.Class,
	)

	return html.Div(
		html.Class(classes),
		g.Group(children),
	)
}

// DialogFooter creates the Dialog footer
func DialogFooter(props FooterProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"flex flex-col-reverse sm:flex-row sm:justify-end sm:space-x-2",
		props.Class,
	)

	return html.Div(
		html.Class(classes),
		g.Group(children),
	)
}

// DialogTitle creates the Dialog title
func DialogTitle(props TitleProps, text string) g.Node {
	classes := lib.CN(
		"text-lg font-semibold leading-none tracking-tight",
		props.Class,
	)

	return html.H2(
		html.Class(classes),
		g.Text(text),
	)
}

// Description creates the Dialog description
func Description(props DescriptionProps, text string) g.Node {
	classes := lib.CN(
		"text-sm text-muted-foreground",
		props.Class,
	)

	return html.P(
		html.Class(classes),
		g.Text(text),
	)
}

// Trigger creates a trigger button for the dialog (requires JavaScript)
func Trigger(props TriggerProps, children ...g.Node) g.Node {
	classes := lib.CN(
		props.Class,
	)

	return html.Button(
		html.Type("button"),
		g.If(classes != "", html.Class(classes)),
		g.Group(children),
	)
}

// Close creates a close button
func Close(props CloseProps, children ...g.Node) g.Node {
	classes := lib.CN(
		props.Class,
	)

	return html.Button(
		html.Type("button"),
		g.If(classes != "", html.Class(classes)),
		g.Group(children),
	)
}

// Example creates a basic dialog example
func Example() g.Node {
	return ExampleWithHTMX()
}

// ExampleScrollable creates a scrollable dialog example
func ExampleScrollable() g.Node {
	return New(
		Props{Open: true},
		Overlay(),
		DialogContent(
			ContentProps{ShowCloseButton: true, Class: "max-h-[80vh]"},
			DialogHeader(
				HeaderProps{},
				DialogTitle(TitleProps{}, "Terms of Service"),
				Description(DescriptionProps{}, "Please read and accept our terms of service."),
			),
			html.Div(html.Class("overflow-y-auto max-h-[60vh] pr-2"),
				html.Div(html.Class("space-y-4"),
					html.H3(html.Class("font-medium"), g.Text("1. Introduction")),
					html.P(html.Class("text-sm text-muted-foreground"),
						g.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."),
					),
					html.H3(html.Class("font-medium"), g.Text("2. Terms of Use")),
					html.P(html.Class("text-sm text-muted-foreground"),
						g.Text("Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat."),
					),
					html.H3(html.Class("font-medium"), g.Text("3. Privacy Policy")),
					html.P(html.Class("text-sm text-muted-foreground"),
						g.Text("Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur."),
					),
					html.H3(html.Class("font-medium"), g.Text("4. User Responsibilities")),
					html.P(html.Class("text-sm text-muted-foreground"),
						g.Text("Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."),
					),
					html.H3(html.Class("font-medium"), g.Text("5. Disclaimers")),
					html.P(html.Class("text-sm text-muted-foreground"),
						g.Text("Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium."),
					),
				),
			),
			DialogFooter(
				FooterProps{},
				html.Button(html.Type("button"), html.Class("border hover:bg-accent"), g.Text("Decline")),
				html.Button(html.Type("button"), html.Class("bg-primary text-primary-foreground hover:bg-primary/90"), g.Text("Accept")),
			),
		),
	)
}

// ExampleCustom creates a custom styled dialog
func ExampleCustom() g.Node {
	return New(
		Props{Open: true},
		Overlay("bg-white/80 backdrop-blur-sm"),
		DialogContent(
			ContentProps{Class: "border-0 bg-slate-950 text-slate-50"},
			html.Div(html.Class("text-center space-y-6"),
				html.Div(html.Class("mx-auto w-12 h-12 bg-green-500/20 rounded-full flex items-center justify-center"),
					// Checkmark icon
					g.El("svg",
						g.Attr("class", "w-6 h-6 text-green-500"),
						g.Attr("fill", "none"),
						g.Attr("viewBox", "0 0 24 24"),
						g.Attr("stroke", "currentColor"),
						g.Attr("stroke-width", "3"),
						g.El("path",
							g.Attr("stroke-linecap", "round"),
							g.Attr("stroke-linejoin", "round"),
							g.Attr("d", "M5 13l4 4L19 7"),
						),
					),
				),
				DialogHeader(
					HeaderProps{},
					DialogTitle(TitleProps{Class: "text-slate-50"}, "Payment Successful"),
					Description(DescriptionProps{Class: "text-slate-400"}, "Your payment has been processed successfully."),
				),
				html.Div(html.Class("bg-slate-900/50 rounded-lg p-4 space-y-2"),
					html.Div(html.Class("flex justify-between text-sm"),
						html.Span(html.Class("text-slate-400"), g.Text("Amount")),
						html.Span(html.Class("font-medium"), g.Text("$99.00")),
					),
					html.Div(html.Class("flex justify-between text-sm"),
						html.Span(html.Class("text-slate-400"), g.Text("Reference")),
						html.Span(html.Class("font-mono text-xs"), g.Text("TXN-20240115-001")),
					),
				),
				DialogFooter(
					FooterProps{Class: "sm:justify-center"},
					html.Button(html.Type("button"), html.Class("bg-slate-50 text-slate-950 hover:bg-slate-200 w-full sm:w-auto"), g.Text("Done")),
				),
			),
		),
	)
}