package hovercard

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

// Example demonstrates the HoverCard component
func Example() g.Node {
	return html.Div(html.Class("space-y-8"),
		// Basic Example
		html.Div(
			html.H3(html.Class("text-lg font-semibold"), g.Text("Basic Hover Card")),
			html.P(html.Class("text-sm text-muted-foreground mb-4"), 
				g.Text("Hover over the link to see the card."),
			),
			BasicExample(),
		),

		// Profile Cards
		html.Div(
			html.H3(html.Class("text-lg font-semibold"), g.Text("Profile Cards")),
			html.P(html.Class("text-sm text-muted-foreground mb-4"), 
				g.Text("Hover over usernames to see profile information."),
			),
			html.Div(html.Class("flex gap-4 flex-wrap"),
				ProfileCard("vercel", "Vercel", "Develop. Preview. Ship.", ""),
				ProfileCard("shadcn", "shadcn", "Design engineer. Building UI things.", ""),
				ProfileCard("radix", "Radix UI", "Low-level UI primitives for React.", ""),
			),
		),

		// Link Previews
		html.Div(
			html.H3(html.Class("text-lg font-semibold"), g.Text("Link Previews")),
			html.P(html.Class("text-sm text-muted-foreground mb-4"), 
				g.Text("Hover over links to see previews."),
			),
			html.P(html.Class("space-x-2"),
				g.Text("Check out "),
				LinkPreview(
					"https://nextjs.org",
					"Next.js by Vercel",
					"The React Framework for Production - Next.js gives you the best developer experience.",
					"",
				),
				g.Text(" and "),
				LinkPreview(
					"https://tailwindcss.com",
					"Tailwind CSS",
					"A utility-first CSS framework for rapidly building modern websites.",
					"",
				),
				g.Text("."),
			),
		),

		// Calendar Hover Card
		html.Div(
			html.H3(html.Class("text-lg font-semibold"), g.Text("Calendar Hover Card")),
			html.P(html.Class("text-sm text-muted-foreground mb-4"), 
				g.Text("Hover over the date to see events."),
			),
			html.Div(html.Class("flex gap-4"),
				Calendar("Today", []string{
					"9:00 AM - Team standup",
					"2:00 PM - Client meeting",
					"4:00 PM - Code review",
				}),
				Calendar("Tomorrow", []string{
					"10:00 AM - Sprint planning",
					"3:00 PM - Design review",
				}),
				Calendar("Friday", []string{}),
			),
		),

		// HTMX Examples
		html.Div(
			html.H3(html.Class("text-lg font-semibold"), g.Text("HTMX Hover Cards")),
			html.P(html.Class("text-sm text-muted-foreground mb-4"), 
				g.Text("Dynamic hover cards loaded with HTMX."),
			),
			html.Div(html.Class("space-y-2"),
				html.P(
					g.Text("Follow "),
					ExampleHTMX(),
					g.Text(" for the latest updates."),
				),
				html.P(
					g.Text("Contributors: "),
					ProfileCardHTMX("john"),
					g.Text(", "),
					ProfileCardHTMX("jane"),
					g.Text(" and "),
					ProfileCardHTMX("alex"),
					g.Text("."),
				),
				html.P(
					g.Text("Learn more at "),
					LinkPreviewHTMX("https://nextjs.org"),
					g.Text("."),
				),
			),
		),
	)
}

// ExampleWithCustomContent demonstrates custom hover card content
func ExampleWithCustomContent() g.Node {
	return html.Div(html.Class("space-y-4"),
		// Stats Card
		New(
			Props{},
			Trigger(
				TriggerProps{},
				html.Button(
					html.Type("button"),
					html.Class("inline-flex items-center gap-2 rounded-md border px-3 py-2 text-sm hover:bg-accent"),
					g.El("svg",
						g.Attr("viewBox", "0 0 24 24"),
						g.Attr("fill", "none"),
						g.Attr("stroke", "currentColor"),
						g.Attr("stroke-width", "2"),
						g.Attr("stroke-linecap", "round"),
						g.Attr("stroke-linejoin", "round"),
						html.Class("h-4 w-4"),
						g.El("path", g.Attr("d", "M3 3v18h18")),
						g.El("path", g.Attr("d", "m19 9-5 5-4-4-3 3")),
					),
					g.Text("View Stats"),
				),
			),
	ContentComponent(
				ContentProps{Class: "w-80"},
				html.Div(html.Class("space-y-3"),
					html.H4(html.Class("font-medium"), g.Text("Performance Overview")),
					html.Div(html.Class("grid grid-cols-2 gap-4"),
						html.Div(
							html.P(html.Class("text-sm text-muted-foreground"), g.Text("Total Views")),
							html.P(html.Class("text-2xl font-bold"), g.Text("1,234")),
							html.P(html.Class("text-xs text-green-600"), g.Text("↑ 12% from last week")),
						),
						html.Div(
							html.P(html.Class("text-sm text-muted-foreground"), g.Text("Engagement")),
							html.P(html.Class("text-2xl font-bold"), g.Text("89%")),
							html.P(html.Class("text-xs text-green-600"), g.Text("↑ 3% from last week")),
						),
						html.Div(
							html.P(html.Class("text-sm text-muted-foreground"), g.Text("Conversions")),
							html.P(html.Class("text-2xl font-bold"), g.Text("43")),
							html.P(html.Class("text-xs text-red-600"), g.Text("↓ 2% from last week")),
						),
						html.Div(
							html.P(html.Class("text-sm text-muted-foreground"), g.Text("Bounce Rate")),
							html.P(html.Class("text-2xl font-bold"), g.Text("24%")),
							html.P(html.Class("text-xs text-green-600"), g.Text("↓ 5% from last week")),
						),
					),
				),
			),
		),

		// Color Picker Card
		New(
			Props{},
			Trigger(
				TriggerProps{},
				html.Div(html.Class("inline-flex items-center gap-2 rounded-md border px-3 py-2 text-sm cursor-pointer hover:bg-accent"),
					html.Div(html.Class("h-4 w-4 rounded bg-blue-500")),
					g.Text("Primary Color"),
				),
			),
	ContentComponent(
				ContentProps{},
				html.Div(html.Class("space-y-3"),
					html.H4(html.Class("font-medium"), g.Text("Color Palette")),
					html.Div(html.Class("grid grid-cols-5 gap-2"),
						g.Group(g.Map([]string{
							"bg-red-500", "bg-orange-500", "bg-yellow-500", "bg-green-500", "bg-blue-500",
							"bg-indigo-500", "bg-purple-500", "bg-pink-500", "bg-gray-500", "bg-black",
						}, func(color string) g.Node {
							return html.Button(
								html.Type("button"),
								html.Class("h-8 w-8 rounded "+color+" hover:scale-110 transition-transform"),
								g.Attr("aria-label", color),
							)
						})),
					),
					html.Input(
						html.Type("text"),
						html.Value("#3B82F6"),
						html.Class("w-full text-sm"),
						html.Placeholder("Enter hex color"),
					),
				),
			),
		),

		// Command Palette Preview
		New(
			Props{},
			Trigger(
				TriggerProps{},
	html.Kbd(html.Class("inline-flex items-center gap-1 rounded border bg-muted px-1.5 font-mono text-xs cursor-pointer"),
					g.Text("⌘K"),
				),
			),
	ContentComponent(
				ContentProps{Class: "w-96"},
				html.Div(html.Class("space-y-3"),
					html.H4(html.Class("font-medium"), g.Text("Command Palette")),
					html.P(html.Class("text-sm text-muted-foreground"), 
						g.Text("Quick access to all commands and navigation."),
					),
					html.Div(html.Class("space-y-1 rounded-md border p-2"),
						g.Group(g.Map([]struct{ Icon, Label, Shortcut string }{
							{"🔍", "Searchtml...", "⌘K"},
							{"📄", "New Document", "⌘N"},
							{"📁", "Open File", "⌘O"},
							{"⚙️", "Settings", "⌘,"},
							{"🎨", "Change Theme", "⌘T"},
						}, func(cmd struct{ Icon, Label, Shortcut string }) g.Node {
							return html.Div(html.Class("flex items-center justify-between rounded px-2 py-1 text-sm hover:bg-accent"),
								html.Div(html.Class("flex items-center gap-2"),
									html.Span(g.Text(cmd.Icon)),
									g.Text(cmd.Label),
								),
	html.Kbd(html.Class("text-xs"), g.Text(cmd.Shortcut)),
							)
						})),
					),
				),
			),
		),
	)
}

// ExampleWithPositioning demonstrates different positioning options
func ExampleWithPositioning() g.Node {
	positions := []struct {
		Side  string
		Align string
		Label string
	}{
		{"top", "start", "Top Start"},
		{"top", "center", "Top Center"},
		{"top", "end", "Top End"},
		{"bottom", "start", "Bottom Start"},
		{"bottom", "center", "Bottom Center"},
		{"bottom", "end", "Bottom End"},
		{"left", "start", "Left Start"},
		{"left", "center", "Left Center"},
		{"left", "end", "Left End"},
		{"right", "start", "Right Start"},
		{"right", "center", "Right Center"},
		{"right", "end", "Right End"},
	}

	return html.Div(html.Class("grid grid-cols-3 gap-8 p-8"),
		g.Group(g.Map(positions, func(pos struct {
			Side  string
			Align string
			Label string
		}) g.Node {
			return New(
				Props{},
				Trigger(
					TriggerProps{},
					html.Button(
						html.Type("button"),
						html.Class("w-full rounded-md border px-3 py-2 text-sm hover:bg-accent"),
						g.Text(pos.Label),
					),
				),
	ContentComponent(
					ContentProps{
						Side:  pos.Side,
						Align: pos.Align,
					},
					html.Div(html.Class("text-center"),
						html.P(html.Class("font-medium"), g.Text(pos.Label)),
						html.P(html.Class("text-sm text-muted-foreground"), 
							g.Textf("Side: %s, Align: %s", pos.Side, pos.Align),
						),
					),
				),
			)
		})),
	)
}