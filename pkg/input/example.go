package input

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/pkg/label"
)

// Examples demonstrates various input usage patterns
func Examples() g.Node {
	return html.Div(
		html.Class("space-y-8 p-8 max-w-2xl"),

		// Basic Inputs section
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Basic Inputs")),
			html.Div(
				html.Class("space-y-4"),
				Text("Enter your name"),
				Email("Enter your email"),
				Password("Enter your password"),
				Number("Enter your age"),
			),
		),

		// Input with Labels section
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Input with Labels")),
			html.Div(
				html.Class("space-y-4"),
				html.Div(
					html.Class("space-y-2"),
					label.ForInput("username", "Username"),
					New(Props{
						Type:        "text",
						ID:          "username",
						Placeholder: "johndoe",
						Name:        "username",
					}),
				),
				html.Div(
					html.Class("space-y-2"),
					label.ForInputRequired("email", "Email"),
					New(Props{
						Type:        "email",
						ID:          "email",
						Placeholder: "john@example.com",
						Name:        "email",
						Required:    true,
					}),
				),
			),
		),

		// Different Input Types section
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Different Input Types")),
			html.Div(
				html.Class("grid grid-cols-2 gap-4"),
				html.Div(
					html.Class("space-y-2"),
					label.ForInput("search", "Search"),
					Search("Searchtml..."),
				),
				html.Div(
					html.Class("space-y-2"),
					label.ForInput("tel", "Phone"),
					Tel("+1 (555) 000-0000"),
				),
				html.Div(
					html.Class("space-y-2"),
					label.ForInput("url", "Website"),
					URL("https://example.com"),
				),
				html.Div(
					html.Class("space-y-2"),
					label.ForInput("date", "Date"),
					New(Props{
						Type: "date",
						ID:   "date",
					}),
				),
			),
		),

		// States section
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("States")),
			html.Div(
				html.Class("space-y-4"),
				html.Div(
					html.Class("space-y-2"),
					label.Default("Disabled Input"),
					New(Props{
						Type:        "text",
						Placeholder: "Cannot type here",
						Disabled:    true,
					}),
				),
				html.Div(
					html.Class("space-y-2"),
					label.Default("Input with Error"),
					New(Props{
						Type:        "email",
						Placeholder: "Invalid email",
						Value:       "not-an-email",
						AriaInvalid: true,
					}),
					html.P(html.Class("text-sm text-destructive"), g.Text("Please enter a valid email address")),
				),
			),
		),

		// Form Example section
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Form Example")),
			html.Form(
				html.Class("space-y-4 border rounded-lg p-4"),
				html.Div(
					html.Class("grid grid-cols-2 gap-4"),
					html.Div(
						html.Class("space-y-2"),
						label.ForInputRequired("first-name", "First Name"),
						New(Props{
							Type:         "text",
							ID:           "first-name",
							Name:         "firstName",
							Required:     true,
							AutoComplete: "given-name",
						}),
					),
					html.Div(
						html.Class("space-y-2"),
						label.ForInputRequired("last-name", "Last Name"),
						New(Props{
							Type:         "text",
							ID:           "last-name",
							Name:         "lastName",
							Required:     true,
							AutoComplete: "family-name",
						}),
					),
				),
				html.Div(
					html.Class("space-y-2"),
					label.ForInputRequired("email-address", "Email Address"),
					New(Props{
						Type:         "email",
						ID:           "email-address",
						Name:         "email",
						Required:     true,
						AutoComplete: "email",
					}),
				),
				html.Div(
					html.Class("space-y-2"),
					label.ForInput("bio", "Bio"),
					New(Props{
						Type:        "text",
						ID:          "bio",
						Name:        "bio",
						Placeholder: "Tell us about yourself",
						Class:       "h-20",
					}),
				),
			),
		),

		// Custom Styling section
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Custom Styling")),
			html.Div(
				html.Class("space-y-4"),
				New(Props{
					Type:        "text",
					Placeholder: "Full width input",
					Class:       "w-full",
				}),
				New(Props{
					Type:        "text",
					Placeholder: "Half width input",
					Class:       "w-1/2",
				}),
				New(Props{
					Type:        "text",
					Placeholder: "Custom border color",
					Class:       "border-blue-500 focus-visible:border-blue-600",
				}),
			),
		),
	)
}