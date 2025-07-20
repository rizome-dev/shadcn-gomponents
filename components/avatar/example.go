package avatar

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

// Example demonstrates how to use the Avatar component
func Example() g.Node {
	return html.Div(
		html.Class("p-8 space-y-8"),
		
		// Basic avatars with different sizes
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Avatar Sizes")),
			html.Div(html.Class("flex items-center gap-4"),
				html.Small(
					Image(ImageProps{
						Src: "https://github.com/shadcn.png",
						Alt: "@shadcn",
					}),
				),
				Default(
					Image(ImageProps{
						Src: "https://github.com/shadcn.png",
						Alt: "@shadcn",
					}),
				),
				Large(
					Image(ImageProps{
						Src: "https://github.com/shadcn.png",
						Alt: "@shadcn",
					}),
				),
				New(
					Props{Size: "h-20 w-20"},
					Image(ImageProps{
						Src: "https://github.com/shadcn.png",
						Alt: "@shadcn",
					}),
				),
			),
		),
		
		// Avatars with fallback content
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("With Fallbacks")),
			html.Div(html.Class("flex items-center gap-4"),
				WithInitials("JD"),
				WithInitials("AB"),
				WithInitials("CN"),
				Default(
					Fallback(
						html.Span(html.Class("text-xs"), g.Text("USR")),
					),
				),
			),
		),
		
		// Avatars with icons
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("With Icons")),
			html.Div(html.Class("flex items-center gap-4"),
				WithIcon(
					// User icon
					g.El("svg",
						g.Attr("xmlns", "http://www.w3.org/2000/svg"),
						g.Attr("viewBox", "0 0 24 24"),
						g.Attr("fill", "none"),
						g.Attr("stroke", "currentColor"),
						g.Attr("stroke-width", "2"),
						g.Attr("stroke-linecap", "round"),
						g.Attr("stroke-linejoin", "round"),
						html.Class("h-4 w-4"),
						g.El("path", g.Attr("d", "M19 21v-2a4 4 0 0 0-4-4H9a4 4 0 0 0-4 4v2")),
						g.El("circle", g.Attr("cx", "12"), g.Attr("cy", "7"), g.Attr("r", "4")),
					),
				),
				New(
					Props{Class: "bg-primary text-primary-foreground"},
					Fallback(
						// Star icon
						g.El("svg",
							g.Attr("xmlns", "http://www.w3.org/2000/svg"),
							g.Attr("viewBox", "0 0 24 24"),
							g.Attr("fill", "currentColor"),
							html.Class("h-5 w-5"),
							g.El("polygon", g.Attr("points", "12 2 15.09 8.26 22 9.27 17 14.14 18.18 21.02 12 17.77 5.82 21.02 7 14.14 2 9.27 8.91 8.26 12 2")),
						),
					),
				),
			),
		),
		
		// Avatar group (overlapping avatars)
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Avatar Group")),
			Group(
				GroupItem(
					New(
						Props{Class: "border-2 border-background"},
						Image(ImageProps{
							Src: "https://i.pravatar.cc/150?img=1",
							Alt: "User 1",
						}),
					),
					0,
				),
				GroupItem(
					New(
						Props{Class: "border-2 border-background"},
						Image(ImageProps{
							Src: "https://i.pravatar.cc/150?img=2",
							Alt: "User 2",
						}),
					),
					1,
				),
				GroupItem(
					New(
						Props{Class: "border-2 border-background"},
						Image(ImageProps{
							Src: "https://i.pravatar.cc/150?img=3",
							Alt: "User 3",
						}),
					),
					2,
				),
				GroupItem(
					New(
						Props{Class: "border-2 border-background"},
						Fallback(
							html.Span(html.Class("text-xs"), g.Text("+5")),
						),
					),
					3,
				),
			),
		),
		
		// Avatars with custom styling
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Custom Styled Avatars")),
			html.Div(html.Class("flex items-center gap-4"),
				New(
					Props{Class: "ring-2 ring-primary ring-offset-2"},
					WithInitials("PR"),
				),
				New(
					Props{
						Size:  "h-12 w-12",
						Class: "rounded-lg", // Square with rounded corners
					},
					Image(ImageProps{
						Src: "https://github.com/vercel.png",
						Alt: "@vercel",
					}),
				),
				New(
					Props{Class: "bg-gradient-to-br from-purple-500 to-pink-500"},
					Fallback(
						html.Span(html.Class("text-white font-bold"), g.Text("AI")),
					),
				),
			),
		),
		
		// Avatar in context (user info card)
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("In Context")),
			html.Div(html.Class("flex items-center space-x-4 p-4 rounded-lg border"),
				Default(
					Image(ImageProps{
						Src: "https://github.com/shadcn.png",
						Alt: "@shadcn",
					}),
				),
				html.Div(
					html.H3(html.Class("font-semibold"), g.Text("shadcn")),
					html.P(html.Class("text-sm text-muted-foreground"), g.Text("Building beautiful UIs")),
				),
			),
		),
	)
}