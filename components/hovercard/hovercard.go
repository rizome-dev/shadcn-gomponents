package hovercard

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// Props defines the properties for the HoverCard component
type Props struct {
	Open  bool
	Class string
}

// TriggerProps defines properties for the hover card trigger
type TriggerProps struct {
	Class string
}

// ContentProps defines properties for the hover card content
type ContentProps struct {
	Class    string
	Side     string // "top" | "bottom" | "left" | "right"
	Align    string // "start" | "center" | "end"
	SideOffset int  // Offset from the trigger
	AlignOffset int // Offset along the alignment axis
}

// New creates a new HoverCard component
func New(props Props, children ...g.Node) g.Node {
	classes := lib.CN("relative inline-block", props.Class)
	
	return html.Div(
		html.Class(classes),
		g.Attr("data-hover-card", "root"),
		g.If(props.Open, g.Attr("data-state", "open")),
		g.If(!props.Open, g.Attr("data-state", "closed")),
		g.Group(children),
	)
}

// Trigger creates a hover card trigger
func Trigger(props TriggerProps, children ...g.Node) g.Node {
	classes := lib.CN("cursor-pointer", props.Class)
	
	return html.Div(
		html.Class(classes),
		g.Attr("data-hover-card", "trigger"),
		g.Attr("aria-haspopup", "dialog"),
		g.Attr("aria-expanded", "false"),
		g.Group(children),
	)
}

// Content creates the hover card content
func ContentComponent(props ContentProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"z-50 w-64 rounded-md border bg-popover p-4 text-popover-foreground shadow-md outline-none",
		"data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95",
		"data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2",
		props.Class,
	)
	
	// Set default values
	side := props.Side
	if side == "" {
		side = "bottom"
	}
	
	align := props.Align
	if align == "" {
		align = "center"
	}
	
	sideOffset := props.SideOffset
	if sideOffset == 0 {
		sideOffset = 4
	}
	
	return html.Div(
		html.Class(classes),
		g.Attr("data-hover-card", "content"),
		g.Attr("data-state", "closed"),
		g.Attr("data-side", side),
		g.Attr("data-align", align),
		html.Style("position: absolute; display: none;"),
		g.Attr("role", "dialog"),
		g.Group(children),
	)
}

// BasicExample creates a basic hover card example
func BasicExample() g.Node {
	return New(
		Props{},
		Trigger(
			TriggerProps{},
			html.A(
				html.Href("#"),
				html.Class("text-sm font-medium underline-offset-4 hover:underline"),
				g.Text("@nextjs"),
			),
		),
	ContentComponent(
			ContentProps{},
			html.Div(html.Class("flex justify-between space-x-4"),
				html.Div(html.Class("space-y-1"),
					html.H4(html.Class("text-sm font-semibold"), g.Text("@nextjs")),
					html.P(html.Class("text-sm"), 
						g.Text("The React Framework – created and maintained by @vercel."),
					),
					html.Div(html.Class("flex items-center pt-2"),
						html.Span(html.Class("text-xs text-muted-foreground"),
							g.El("svg",
								g.Attr("viewBox", "0 0 24 24"),
								g.Attr("fill", "none"),
								g.Attr("stroke", "currentColor"),
								g.Attr("stroke-width", "2"),
								g.Attr("stroke-linecap", "round"),
								g.Attr("stroke-linejoin", "round"),
								html.Class("mr-2 h-4 w-4 opacity-70"),
								g.El("path", g.Attr("d", "M22 4s-.7 2.1-2 3.4c1.6 10-9.4 17.3-18 11.6 2.2.1 4.4-.6 6-2C3 15.5.5 9.6 3 5c2.2 2.6 5.6 4.1 9 4-.9-4.2 4-6.6 7-3.8 1.1 0 3-1.2 3-1.2z")),
							),
							g.Text("Joined December 2021"),
						),
					),
				),
			),
		),
	)
}

// ProfileCard creates a profile hover card
func ProfileCard(username string, name string, bio string, avatarUrl string) g.Node {
	return New(
		Props{},
		Trigger(
			TriggerProps{},
			html.Button(
				html.Type("button"),
				html.Class("flex items-center space-x-2 rounded-md p-2 hover:bg-accent"),
				g.If(avatarUrl != "",
	html.Img(
						html.Src(avatarUrl),
						html.Alt(name),
						html.Class("h-8 w-8 rounded-full"),
					),
				),
				g.If(avatarUrl == "",
					html.Div(html.Class("h-8 w-8 rounded-full bg-muted")),
				),
				html.Span(html.Class("text-sm font-medium"), g.Text(name)),
			),
		),
	ContentComponent(
			ContentProps{Side: "bottom", Align: "start"},
			html.Div(html.Class("space-y-2"),
				html.Div(html.Class("flex items-center space-x-4"),
					g.If(avatarUrl != "",
	html.Img(
							html.Src(avatarUrl),
							html.Alt(name),
							html.Class("h-12 w-12 rounded-full"),
						),
					),
					g.If(avatarUrl == "",
						html.Div(html.Class("h-12 w-12 rounded-full bg-muted")),
					),
					html.Div(html.Class("space-y-1"),
						html.H4(html.Class("text-sm font-semibold"), g.Text(name)),
						html.P(html.Class("text-sm text-muted-foreground"), g.Text("@"+username)),
					),
				),
				g.If(bio != "",
					html.P(html.Class("text-sm"), g.Text(bio)),
				),
				html.Div(html.Class("flex gap-4 text-xs text-muted-foreground"),
					html.Div(
						html.Span(html.Class("font-semibold text-foreground"), g.Text("256")),
						g.Text(" Following"),
					),
					html.Div(
						html.Span(html.Class("font-semibold text-foreground"), g.Text("2.1k")),
						g.Text(" Followers"),
					),
				),
			),
		),
	)
}

// LinkPreview creates a link preview hover card
func LinkPreview(url string, title string, description string, imageUrl string) g.Node {
	return New(
		Props{},
		Trigger(
			TriggerProps{},
			html.A(
				html.Href(url),
				html.Class("text-blue-600 hover:underline"),
				g.Text(url),
			),
		),
	ContentComponent(
			ContentProps{Class: "w-80"},
			html.Div(html.Class("space-y-2"),
				g.If(imageUrl != "",
	html.Img(
						html.Src(imageUrl),
						html.Alt(title),
						html.Class("h-40 w-full rounded-md object-cover"),
					),
				),
				html.Div(html.Class("space-y-1"),
					html.H4(html.Class("text-sm font-semibold"), g.Text(title)),
					html.P(html.Class("text-sm text-muted-foreground line-clamp-2"), 
						g.Text(description),
					),
					html.P(html.Class("text-xs text-muted-foreground"), g.Text(url)),
				),
			),
		),
	)
}

// Calendar creates a calendar hover card
func Calendar(date string, events []string) g.Node {
	return New(
		Props{},
		Trigger(
			TriggerProps{},
			html.Button(
				html.Type("button"),
				html.Class("inline-flex items-center rounded-md px-3 py-1 text-sm border hover:bg-accent"),
				g.El("svg",
					g.Attr("viewBox", "0 0 24 24"),
					g.Attr("fill", "none"),
					g.Attr("stroke", "currentColor"),
					g.Attr("stroke-width", "2"),
					g.Attr("stroke-linecap", "round"),
					g.Attr("stroke-linejoin", "round"),
					html.Class("mr-2 h-4 w-4"),
					g.El("rect", g.Attr("width", "18"), g.Attr("height", "18"), g.Attr("x", "3"), g.Attr("y", "4"), g.Attr("rx", "2"), g.Attr("ry", "2")),
					g.El("line", g.Attr("x1", "16"), g.Attr("x2", "16"), g.Attr("y1", "2"), g.Attr("y2", "6")),
					g.El("line", g.Attr("x1", "8"), g.Attr("x2", "8"), g.Attr("y1", "2"), g.Attr("y2", "6")),
					g.El("line", g.Attr("x1", "3"), g.Attr("x2", "21"), g.Attr("y1", "10"), g.Attr("y2", "10")),
				),
				g.Text(date),
			),
		),
	ContentComponent(
			ContentProps{},
			html.Div(html.Class("space-y-2"),
				html.H4(html.Class("font-medium"), g.Text(date)),
				g.If(len(events) > 0,
					html.Div(html.Class("space-y-1"),
						g.Group(g.Map(events, func(event string) g.Node {
							return html.Div(html.Class("flex items-start gap-2 text-sm"),
								html.Div(html.Class("h-2 w-2 rounded-full bg-primary mt-1")),
								g.Text(event),
							)
						})),
					),
				),
				g.If(len(events) == 0,
					html.P(html.Class("text-sm text-muted-foreground"), g.Text("No events scheduled")),
				),
			),
		),
	)
}