package card

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// Card creates a card container
func Card(children ...g.Node) g.Node {
	return html.Div(
		html.Class("bg-card text-card-foreground flex flex-col gap-6 rounded-xl border py-6 shadow-sm"),
		g.Group(children),
	)
}

// CardHeader creates a card header section
func CardHeader(children ...g.Node) g.Node {
	return html.Div(
		html.Class("@container/card-header grid auto-rows-min grid-rows-[auto_auto] items-start gap-1.5 px-6 has-data-[slot=card-action]:grid-cols-[1fr_auto] [.border-b]:pb-6"),
		g.Group(children),
	)
}

// CardTitle creates a card title
func CardTitle(children ...g.Node) g.Node {
	return html.H3(
		html.Class("leading-none font-semibold"),
		g.Group(children),
	)
}

// CardDescription creates a card description
func CardDescription(children ...g.Node) g.Node {
	return html.P(
		html.Class("text-muted-foreground text-sm"),
		g.Group(children),
	)
}

// CardContent creates a card content section
func CardContent(children ...g.Node) g.Node {
	return html.Div(
		html.Class("px-6"),
		g.Group(children),
	)
}

// CardFooter creates a card footer section
func CardFooter(children ...g.Node) g.Node {
	return html.Div(
		html.Class("flex items-center px-6 [.border-t]:pt-6"),
		g.Group(children),
	)
}

// CardAction creates a card action (for header)
func CardAction(children ...g.Node) g.Node {
	return html.Div(
		html.Class("col-start-2 row-span-2 row-start-1 self-start justify-self-end"),
		g.Attr("data-slot", "card-action"),
		g.Group(children),
	)
}

// Props defines properties for creating a complete card
type Props struct {
	Title       string
	Description string
	Class       string // Additional custom classes
}

// New creates a card with the given props and content
func New(props Props, content ...g.Node) g.Node {
	var headerContent []g.Node

	if props.Title != "" {
		headerContent = append(headerContent, CardTitle(g.Text(props.Title)))
	}
	if props.Description != "" {
		headerContent = append(headerContent, CardDescription(g.Text(props.Description)))
	}

	classes := lib.CN("bg-card text-card-foreground flex flex-col gap-6 rounded-xl border py-6 shadow-sm", props.Class)

	return html.Div(
		html.Class(classes),
		g.If(len(headerContent) > 0, CardHeader(headerContent...)),
		g.If(len(content) > 0, CardContent(content...)),
	)
}

// WithFooter creates a card with header, content, and footer sections
func WithFooter(title, description string, content, footerContent []g.Node) g.Node {
	return Card(
		CardHeader(
			CardTitle(g.Text(title)),
			g.If(description != "", CardDescription(g.Text(description))),
		),
		g.If(len(content) > 0, CardContent(g.Group(content))),
		g.If(len(footerContent) > 0, CardFooter(g.Group(footerContent))),
	)
}

// Simple creates a simple card with just content
func Simple(children ...g.Node) g.Node {
	return Card(
		CardContent(children...),
	)
}