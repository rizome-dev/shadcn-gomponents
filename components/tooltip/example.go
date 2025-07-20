package tooltip

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/components/button"
)

// Example demonstrates how to use the Tooltip component
func Example() g.Node {
	return html.Div(
		html.Class("p-8 space-y-8"),
		
		// Basic tooltips
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Basic Tooltips")),
			html.Div(html.Class("flex items-center gap-4"),
				Simple("Add to library",
					button.Default(
						g.Raw(`<svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" viewBox="0 0 20 20" fill="currentColor"><path d="M10 3a1 1 0 011 1v5h5a1 1 0 110 2h-5v5a1 1 0 11-2 0v-5H4a1 1 0 110-2h5V4a1 1 0 011-1z" /></svg>`),
						g.Text(" Add"),
					),
				),
				Simple("You have 3 unread messages",
					button.Outline(g.Text("Inbox (3)")),
				),
				Simple("Click to view profile",
					html.Span(html.Class("inline-flex items-center justify-center h-10 w-10 rounded-full bg-muted"),
						g.Text("JD"),
					),
				),
			),
		),
		
		// Different sides
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Tooltip Positions")),
			html.Div(html.Class("flex items-center gap-8 justify-center p-8"),
				WithSide("I'm on top", SideTop,
					button.Outline(g.Text("Top")),
				),
				WithSide("I'm on the right", SideRight,
					button.Outline(g.Text("Right")),
				),
				WithSide("I'm on the bottom", SideBottom,
					button.Outline(g.Text("Bottom")),
				),
				WithSide("I'm on the left", SideLeft,
					button.Outline(g.Text("Left")),
				),
			),
		),
		
		// Different alignments
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Tooltip Alignments")),
			html.Div(html.Class("flex flex-col gap-4"),
				New(Props{
					Content: "Aligned to start",
					Side:    SideRight,
					Align:   AlignStart,
				},
					button.Default(html.Class("w-32"), g.Text("Start")),
					g.Text("Aligned to start"),
				),
				New(Props{
					Content: "Aligned to center",
					Side:    SideRight,
					Align:   AlignCenter,
				},
					button.Default(html.Class("w-32"), g.Text("Center")),
					g.Text("Aligned to center"),
				),
				New(Props{
					Content: "Aligned to end",
					Side:    SideRight,
					Align:   AlignEnd,
				},
					button.Default(html.Class("w-32"), g.Text("End")),
					g.Text("Aligned to end"),
				),
			),
		),
		
		// With delays
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Tooltip Delays")),
			html.Div(html.Class("flex items-center gap-4"),
				WithDelay("Instant tooltip", 0,
					button.Secondary(g.Text("No delay")),
				),
				WithDelay("Quick tooltip", 200,
					button.Secondary(g.Text("200ms delay")),
				),
				WithDelay("Slower tooltip", 500,
					button.Secondary(g.Text("500ms delay")),
				),
				WithDelay("Very slow tooltip", 1000,
					button.Secondary(g.Text("1s delay")),
				),
			),
		),
		
		// Styled tooltips
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Styled Tooltips")),
			html.Div(html.Class("flex items-center gap-4"),
				InfoTooltip("This action cannot be undone",
					button.Default(g.Text("Delete")),
				),
				WarningTooltip("This will affect all users",
					button.Default(g.Text("Update Settings")),
				),
				ErrorTooltip("You don't have permission",
					button.Default(html.Disabled(), g.Text("Admin Only")),
				),
			),
		),
		
		// Interactive elements
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Interactive Elements")),
			html.Div(html.Class("space-y-4"),
				// Form field with help tooltip
				html.Div(html.Class("flex items-center gap-2"),
					html.Label(html.For("email"), g.Text("Email")),
					Simple("Enter your email address",
						g.Raw(`<svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 text-muted-foreground" viewBox="0 0 20 20" fill="currentColor"><path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z" clip-rule="evenodd" /></svg>`),
					),
				),
				
				// Status indicators
				html.Div(html.Class("flex items-center gap-4"),
					New(Props{
						Content: "Server is running normally",
						ContentClass: "bg-green-600 text-white",
						ArrowClass: "bg-green-600",
					},
						html.Span(html.Class("flex items-center gap-2"),
							html.Span(html.Class("h-2 w-2 bg-green-500 rounded-full")),
							g.Text("Online"),
						),
						g.Text("Server is running normally"),
					),
					New(Props{
						Content: "Scheduled maintenance at 2 AM",
						ContentClass: "bg-yellow-600 text-white",
						ArrowClass: "bg-yellow-600",
					},
						html.Span(html.Class("flex items-center gap-2"),
							html.Span(html.Class("h-2 w-2 bg-yellow-500 rounded-full")),
							g.Text("Maintenance"),
						),
						g.Text("Scheduled maintenance at 2 AM"),
					),
				),
				
				// Icon buttons
				html.Div(html.Class("flex items-center gap-2"),
					Simple("Bold (⌘B)",
						button.New(button.Props{Variant: "ghost", Size: "sm"},
							g.Raw(`<svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg" class="h-4 w-4"><path d="M5.10505 12C4.70805 12 4.4236 11.912 4.25171 11.736C4.0839 11.5559 4 11.2715 4 10.8827V4.11733C4 3.72033 4.08595 3.43588 4.25784 3.26398C4.43383 3.08799 4.71623 3 5.10505 3H8.2C8.94904 3 9.53636 3.21164 9.96196 3.63491C10.3917 4.05819 10.6065 4.63592 10.6065 5.36811C10.6065 5.92517 10.4763 6.39344 10.2159 6.77292C9.95958 7.14831 9.59099 7.42264 9.11014 7.59592C9.662 7.72607 10.1014 7.99438 10.4283 8.40085C10.7593 8.80322 10.9248 9.32428 10.9248 9.96401C10.9248 10.7383 10.6679 11.3679 10.1541 11.8527C9.64045 12.3375 8.96325 12.58 8.12222 12.58H5.10505V12ZM6.16134 6.91681H7.825C8.17632 6.91681 8.45268 6.81371 8.65409 6.6075C8.85958 6.3972 8.96233 6.11786 8.96233 5.76948C8.96233 5.4211 8.85958 5.14176 8.65409 4.93145C8.45268 4.72115 8.17632 4.616 7.825 4.616H6.16134V6.91681ZM6.16134 10.9641H8.0318C8.42062 10.9641 8.73039 10.8528 8.96111 10.6302C9.19591 10.4035 9.31331 10.1014 9.31331 9.72409C9.31331 9.34679 9.19591 9.04469 8.96111 8.81779C8.73039 8.59089 8.42062 8.47744 8.0318 8.47744H6.16134V10.9641Z" fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"></path></svg>`),
						),
					),
					Simple("Italic (⌘I)",
						button.New(button.Props{Variant: "ghost", Size: "sm"},
							g.Raw(`<svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg" class="h-4 w-4"><path d="M5.67494 3.50017C5.67494 3.25164 5.87641 3.05017 6.12494 3.05017H10.6249C10.8735 3.05017 11.0749 3.25164 11.0749 3.50017C11.0749 3.7487 10.8735 3.95017 10.6249 3.95017H9.00587L7.2309 11.05H8.87493C9.12345 11.05 9.32493 11.2515 9.32493 11.5C9.32493 11.7486 9.12345 11.95 8.87493 11.95H4.37493C4.1264 11.95 3.92493 11.7486 3.92493 11.5C3.92493 11.2515 4.1264 11.05 4.37493 11.05H5.99397L7.76894 3.95017H6.12494C5.87641 3.95017 5.67494 3.7487 5.67494 3.50017Z" fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"></path></svg>`),
						),
					),
					Simple("Underline (⌘U)",
						button.New(button.Props{Variant: "ghost", Size: "sm"},
							g.Raw(`<svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg" class="h-4 w-4"><path d="M5.00001 2.75C5.00001 2.47386 4.77615 2.25 4.50001 2.25C4.22387 2.25 4.00001 2.47386 4.00001 2.75V8.05C4.00001 9.983 5.56702 11.55 7.50001 11.55C9.43301 11.55 11 9.983 11 8.05V2.75C11 2.47386 10.7762 2.25 10.5 2.25C10.2239 2.25 10 2.47386 10 2.75V8.05C10 9.43071 8.88072 10.55 7.50001 10.55C6.1193 10.55 5.00001 9.43071 5.00001 8.05V2.75ZM3.49998 13.1001C3.27906 13.1001 3.09998 13.2791 3.09998 13.5001C3.09998 13.721 3.27906 13.9001 3.49998 13.9001H11.5C11.7209 13.9001 11.9 13.721 11.9 13.5001C11.9 13.2791 11.7209 13.1001 11.5 13.1001H3.49998Z" fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"></path></svg>`),
						),
					),
				),
			),
		),
		
		// Links with tooltips
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Links with Tooltips")),
			html.P(html.Class("text-sm text-muted-foreground"),
				g.Text("Hover over the "),
				Simple("This link opens in a new tab",
					html.A(html.Href("#"), html.Class("text-primary underline underline-offset-4"),
						g.Text("documentation"),
					),
				),
				g.Text(" to learn more. You can also check the "),
				Simple("View our API reference",
					html.A(html.Href("#"), html.Class("text-primary underline underline-offset-4"),
						g.Text("API"),
					),
				),
				g.Text(" for detailed information."),
			),
		),
		
		// Table with tooltips
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Table with Tooltips")),
			html.Div(html.Class("overflow-x-auto"),
				html.Table(html.Class("w-full"),
	html.THead(
						html.Tr(
							html.Th(html.Class("text-left p-2"), g.Text("Feature")),
							html.Th(html.Class("text-left p-2"), 
								html.Span(html.Class("flex items-center gap-1"),
									g.Text("Status"),
									Simple("Shows if the feature is enabled",
										g.Raw(`<svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 text-muted-foreground" viewBox="0 0 20 20" fill="currentColor"><path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z" clip-rule="evenodd" /></svg>`),
									),
								),
							),
							html.Th(html.Class("text-left p-2"), g.Text("Action")),
						),
					),
	html.TBody(
						html.Tr(
							html.Td(html.Class("p-2"), g.Text("Auto-save")),
							html.Td(html.Class("p-2"),
								New(Props{
									Content: "Automatically saves your work every 30 seconds",
									ContentClass: "max-w-xs",
								},
									html.Span(html.Class("text-green-600"), g.Text("Enabled")),
									g.Text("Automatically saves your work every 30 seconds"),
								),
							),
							html.Td(html.Class("p-2"),
								Simple("Turn off auto-save",
									button.New(button.Props{Variant: "ghost", Size: "sm"}, g.Text("Disable")),
								),
							),
						),
						html.Tr(
							html.Td(html.Class("p-2"), g.Text("Dark mode")),
							html.Td(html.Class("p-2"),
								New(Props{
									Content: "Dark mode is currently disabled",
									ContentClass: "max-w-xs",
								},
									html.Span(html.Class("text-muted-foreground"), g.Text("Disabled")),
									g.Text("Dark mode is currently disabled"),
								),
							),
							html.Td(html.Class("p-2"),
								Simple("Enable dark mode",
									button.New(button.Props{Variant: "ghost", Size: "sm"}, g.Text("Enable")),
								),
							),
						),
					),
				),
			),
		),
		
		// Always visible tooltip
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Always Visible")),
			html.P(html.Class("text-sm text-muted-foreground mb-4"), g.Text("This tooltip is always shown")),
			New(Props{
				Content: "I'm always visible!",
				Open:    true,
				Side:    SideRight,
			},
				button.Default(g.Text("No need to hover")),
				g.Text("I'm always visible!"),
			),
		),
		
		// Custom styled tooltips
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Custom Styles")),
			html.Div(html.Class("flex items-center gap-4"),
				New(Props{
					Content:      "Gradient tooltip",
					ContentClass: "bg-gradient-to-r from-purple-500 to-pink-500 text-white",
					ArrowClass:   "bg-pink-500",
				},
					button.Default(g.Text("Gradient")),
					g.Text("Gradient tooltip"),
				),
				New(Props{
					Content:      "Large tooltip with padding",
					ContentClass: "p-4 text-base",
				},
					button.Default(g.Text("Large")),
					g.Text("Large tooltip with padding"),
				),
				New(Props{
					Content:      "Bordered tooltip",
					ContentClass: "bg-background text-foreground border",
					ArrowClass:   "bg-background border-l border-t",
				},
					button.Default(g.Text("Bordered")),
					g.Text("Bordered tooltip"),
				),
			),
		),
	)
}