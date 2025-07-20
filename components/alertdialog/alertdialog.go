package alertdialog

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// Props defines the properties for the AlertDialog component
type Props struct {
	Open  bool   // Whether the dialog is open (for server-side rendering)
	Class string // Additional custom classes
}

// ContentProps defines the properties for the AlertDialogContent
type ContentProps struct {
	Class string
}

// HeaderProps defines the properties for the AlertDialogHeader
type HeaderProps struct {
	Class string
}

// FooterProps defines the properties for the AlertDialogFooter
type FooterProps struct {
	Class string
}

// TitleProps defines the properties for the AlertDialogTitle
type TitleProps struct {
	Class string
}

// DescriptionProps defines the properties for the AlertDialogDescription
type DescriptionProps struct {
	Class string
}

// ActionProps defines the properties for the AlertDialogAction
type ActionProps struct {
	Class string
	Href  string // Optional href to make it a link
}

// CancelProps defines the properties for the AlertDialogCancel
type CancelProps struct {
	Class string
}

// New creates a new AlertDialog component
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

// DialogOverlay creates the AlertDialog overlay
func DialogOverlay(class ...string) g.Node {
	classes := lib.CN(
		"fixed inset-0 z-50 bg-black/80",
		lib.CN(class...),
	)

	return html.Div(html.Class(classes))
}

// DialogContent creates the AlertDialog content container
func DialogContent(props ContentProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"fixed left-[50%] top-[50%] z-50 grid w-full max-w-lg translate-x-[-50%] translate-y-[-50%] gap-4 border bg-background p-6 shadow-lg sm:rounded-lg",
		props.Class,
	)

	return html.Div(
		html.Class(classes),
		g.Group(children),
	)
}

// DialogHeader creates the AlertDialog header
func DialogHeader(props HeaderProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"flex flex-col space-y-2 text-center sm:text-left",
		props.Class,
	)

	return html.Div(
		html.Class(classes),
		g.Group(children),
	)
}

// DialogFooter creates the AlertDialog footer
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

// DialogTitle creates the AlertDialog title
func DialogTitle(props TitleProps, text string) g.Node {
	classes := lib.CN(
		"text-lg font-semibold",
		props.Class,
	)

	return html.H3(
		html.Class(classes),
		g.Text(text),
	)
}

// DialogDescription creates the AlertDialog description
func DialogDescription(props DescriptionProps, text string) g.Node {
	classes := lib.CN(
		"text-sm text-muted-foreground",
		props.Class,
	)

	return html.P(
		html.Class(classes),
		g.Text(text),
	)
}

// DialogAction creates the AlertDialog action button
func DialogAction(props ActionProps, children ...g.Node) g.Node {
	// Use default button styles
	classes := lib.CN(
		"inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium transition-all disabled:pointer-events-none disabled:opacity-50 outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px]",
		"h-9 px-4 py-2",
		"bg-primary text-primary-foreground shadow-xs hover:bg-primary/90",
		props.Class,
	)

	if props.Href != "" {
		return html.A(
			html.Href(props.Href),
			html.Class(classes),
			g.Group(children),
		)
	}

	return html.Button(
		html.Type("button"),
		html.Class(classes),
		g.Group(children),
	)
}

// DialogCancel creates the AlertDialog cancel button
func DialogCancel(props CancelProps, children ...g.Node) g.Node {
	// Use outline button styles
	classes := lib.CN(
		"inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium transition-all disabled:pointer-events-none disabled:opacity-50 outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px]",
		"h-9 px-4 py-2",
		"border bg-background shadow-xs hover:bg-accent hover:text-accent-foreground",
		"mt-2 sm:mt-0",
		props.Class,
	)

	return html.Button(
		html.Type("button"),
		html.Class(classes),
		g.Group(children),
	)
}

// Example creates a complete alert dialog example
func Example() g.Node {
	return New(
		Props{Open: true},
		DialogOverlay(),
		DialogContent(
			ContentProps{},
			DialogHeader(
				HeaderProps{},
				DialogTitle(TitleProps{}, "Are you absolutely sure?"),
				DialogDescription(DescriptionProps{}, "This action cannot be undone. This will permanently delete your account and remove your data from our servers."),
			),
			DialogFooter(
				FooterProps{},
				DialogCancel(CancelProps{}, g.Text("Cancel")),
				DialogAction(ActionProps{}, g.Text("Continue")),
			),
		),
	)
}