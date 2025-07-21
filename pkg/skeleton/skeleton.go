package skeleton

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// Props defines the properties for the Skeleton component
type Props struct {
	Class string // Additional custom classes
}

// New creates a new Skeleton component
func New(props Props) g.Node {
	classes := lib.CN(
		"animate-pulse rounded-md bg-muted",
		props.Class,
	)

	return html.Div(
		html.Class(classes),
		g.Attr("aria-hidden", "true"), // Hide from screen readers as it's decorative
	)
}

// Default creates a default skeleton
func Default() g.Node {
	return New(Props{})
}

// WithClass creates a skeleton with custom classes
func WithClass(class string) g.Node {
	return New(Props{Class: class})
}

// Text creates a text line skeleton
func Text() g.Node {
	return New(Props{
		Class: "h-4 w-full",
	})
}

// TextShort creates a short text line skeleton
func TextShort() g.Node {
	return New(Props{
		Class: "h-4 w-3/4",
	})
}

// TextLines creates multiple text line skeletons
func TextLines(count int) g.Node {
	var lines []g.Node
	for i := 0; i < count; i++ {
		// Make last line shorter for more realistic look
		if i == count-1 {
			lines = append(lines, TextShort())
		} else {
			lines = append(lines, Text())
		}
	}

	return html.Div(
		html.Class("space-y-2"),
		g.Group(lines),
	)
}

// Avatar creates an avatar skeleton
func Avatar(size string) g.Node {
	sizeClass := "h-10 w-10" // default
	switch size {
	case "sm":
		sizeClass = "h-8 w-8"
	case "lg":
		sizeClass = "h-12 w-12"
	}

	return New(Props{
		Class: lib.CN(sizeClass, "rounded-full"),
	})
}

// Button creates a button skeleton
func Button(size string) g.Node {
	sizeClass := "h-10 w-24" // default
	switch size {
	case "sm":
		sizeClass = "h-8 w-20"
	case "lg":
		sizeClass = "h-12 w-28"
	case "icon":
		sizeClass = "h-10 w-10"
	}

	return New(Props{
		Class: sizeClass,
	})
}

// Input creates an input field skeleton
func Input() g.Node {
	return New(Props{
		Class: "h-10 w-full",
	})
}

// Card creates a card skeleton
func Card() g.Node {
	return New(Props{
		Class: "h-[125px] w-full rounded-xl",
	})
}

// Image creates an image skeleton
func Image(aspectRatio string) g.Node {
	aspectClass := "aspect-video" // default 16:9
	switch aspectRatio {
	case "square":
		aspectClass = "aspect-square"
	case "portrait":
		aspectClass = "aspect-[3/4]"
	}

	return New(Props{
		Class: lib.CN("w-full", aspectClass),
	})
}

// Table creates a table skeleton with rows
func TableComponent(rows int) g.Node {
	var tableRows []g.Node

	// Header row
	tableRows = append(tableRows, 
		html.Div(html.Class("flex gap-4 p-4 border-b"),
			New(Props{Class: "h-4 w-24"}),
			New(Props{Class: "h-4 w-32"}),
			New(Props{Class: "h-4 w-20"}),
			New(Props{Class: "h-4 w-16"}),
		),
	)

	// Data rows
	for i := 0; i < rows; i++ {
		tableRows = append(tableRows,
			html.Div(html.Class("flex gap-4 p-4 border-b"),
				New(Props{Class: "h-4 w-20"}),
				New(Props{Class: "h-4 w-40"}),
				New(Props{Class: "h-4 w-16"}),
				New(Props{Class: "h-4 w-12"}),
			),
		)
	}

	return html.Div(
		html.Class("rounded-md border"),
		g.Group(tableRows),
	)
}

// ProfileCard creates a profile card skeleton
func ProfileCard() g.Node {
	return html.Div(
		html.Class("flex items-center space-x-4"),
		Avatar("default"),
		html.Div(html.Class("space-y-2"),
			New(Props{Class: "h-4 w-[200px]"}),
			New(Props{Class: "h-4 w-[150px]"}),
		),
	)
}

// PostCard creates a blog post card skeleton
func PostCard() g.Node {
	return html.Div(
		html.Class("space-y-3"),
		Image("video"),
		html.Div(html.Class("space-y-2"),
			New(Props{Class: "h-4 w-3/4"}),
			TextLines(2),
		),
	)
}

// List creates a list skeleton
func ListComponent(items int) g.Node {
	var listItems []g.Node
	
	for i := 0; i < items; i++ {
		listItems = append(listItems,
			html.Div(html.Class("flex items-center space-x-3"),
				New(Props{Class: "h-4 w-4 rounded-full"}), // Bullet
				New(Props{Class: "h-4 flex-1"}),            // Text
			),
		)
	}

	return html.Div(
		html.Class("space-y-3"),
		g.Group(listItems),
	)
}

// Form creates a form skeleton
func Form() g.Node {
	return html.Div(
		html.Class("space-y-6"),
		// Field 1
		html.Div(html.Class("space-y-2"),
			New(Props{Class: "h-4 w-24"}), // Label
			Input(),                        // Input
		),
		// Field 2
		html.Div(html.Class("space-y-2"),
			New(Props{Class: "h-4 w-32"}), // Label
			Input(),                        // Input
		),
		// Buttons
		html.Div(html.Class("flex gap-3"),
			Button("default"),
			Button("default"),
		),
	)
}

// Grid creates a grid of skeleton cards
func Grid(cols, items int) g.Node {
	var gridItems []g.Node
	
	for i := 0; i < items; i++ {
		gridItems = append(gridItems, Card())
	}

	gridClass := "grid gap-4"
	switch cols {
	case 2:
		gridClass += " grid-cols-2"
	case 3:
		gridClass += " grid-cols-3"
	case 4:
		gridClass += " grid-cols-4"
	default:
		gridClass += " grid-cols-1"
	}

	return html.Div(
		html.Class(gridClass),
		g.Group(gridItems),
	)
}