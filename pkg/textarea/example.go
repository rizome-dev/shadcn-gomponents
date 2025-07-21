package textarea

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

// Example demonstrates how to use the Textarea component
func Example() g.Node {
	return html.Div(
		html.Class("p-8 space-y-8"),
		
		// Basic textarea
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Basic Textarea")),
			Default(),
		),
		
		// With placeholder
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("With Placeholder")),
			WithPlaceholder("Type your message here..."),
		),
		
		// Different sizes
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Different Sizes")),
			html.Div(html.Class("space-y-4"),
				html.Div(
					html.P(html.Class("text-sm text-muted-foreground mb-2"), g.Text("Small (3 rows)")),
					New(Props{
						Rows:        3,
						Placeholder: "Small textarea",
					}),
				),
				html.Div(
					html.P(html.Class("text-sm text-muted-foreground mb-2"), g.Text("Default (auto)")),
					New(Props{
						Placeholder: "Default height textarea",
					}),
				),
				html.Div(
					html.P(html.Class("text-sm text-muted-foreground mb-2"), g.Text("Large (8 rows)")),
					New(Props{
						Rows:        8,
						Placeholder: "Large textarea",
					}),
				),
			),
		),
		
		// Auto-resize textarea
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Auto-resize")),
			html.P(html.Class("text-sm text-muted-foreground mb-2"), g.Text("This textarea grows as you type (if browser supports field-sizing-content)")),
			AutoResize(),
		),
		
		// Resize options
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Resize Options")),
			html.Div(html.Class("grid grid-cols-2 gap-4"),
				html.Div(
					html.P(html.Class("text-sm text-muted-foreground mb-2"), g.Text("No resize")),
					NoResize(),
				),
				html.Div(
					html.P(html.Class("text-sm text-muted-foreground mb-2"), g.Text("Resize both")),
					New(Props{
						Resize:      "both",
						Placeholder: "Drag corner to resize",
					}),
				),
				html.Div(
					html.P(html.Class("text-sm text-muted-foreground mb-2"), g.Text("Horizontal only")),
					New(Props{
						Resize:      "horizontal",
						Placeholder: "Drag right edge",
					}),
				),
				html.Div(
					html.P(html.Class("text-sm text-muted-foreground mb-2"), g.Text("Vertical only")),
					New(Props{
						Resize:      "vertical",
						Placeholder: "Drag bottom edge",
					}),
				),
			),
		),
		
		// Form fields
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Form Fields")),
			html.Div(html.Class("space-y-4"),
				FormField(
					Props{
						Name:        "comment",
						Placeholder: "Leave a comment...",
						Rows:        3,
					},
					"Comment",
					"",
				),
				FormField(
					Props{
						Name:        "feedback",
						Required:    true,
						Placeholder: "We'd love to hear your thoughts",
						Rows:        4,
					},
					"Feedback",
					"Your feedback helps us improve our service",
				),
			),
		),
		
		// With character count
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("With Character Count")),
			html.Div(html.Class("space-y-4"),
				WithCharacterCount(
					Props{
						Name:        "tweet",
						MaxLength:   280,
						Placeholder: "What's happening?",
						Rows:        3,
					},
					"New Tweet",
				),
				Bio("profile-bio", 500),
			),
		),
		
		// Different states
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Different States")),
			html.Div(html.Class("space-y-4"),
				html.Div(
					html.P(html.Class("text-sm text-muted-foreground mb-2"), g.Text("Disabled")),
					New(Props{
						Disabled: true,
						Value:    "This textarea is disabled",
					}),
				),
				html.Div(
					html.P(html.Class("text-sm text-muted-foreground mb-2"), g.Text("Read-only")),
					New(Props{
						ReadOnly: true,
						Value:    "This textarea is read-only. You can select and copy text but not edit it.",
					}),
				),
			),
		),
		
		// In a form
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Contact Form Example")),
			html.Form(html.Class("max-w-md space-y-4 rounded-lg border bg-card p-6"),
				html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Contact Us")),
				
				html.Div(html.Class("grid grid-cols-2 gap-4"),
					html.Div(
						html.Label(html.For("name"), html.Class("text-sm font-medium"), g.Text("Name")),
						html.Input(html.Type("text"), html.ID("name"), html.Name("name"), html.Class("mt-1.5")),
					),
					html.Div(
						html.Label(html.For("email"), html.Class("text-sm font-medium"), g.Text("Email")),
						html.Input(html.Type("email"), html.ID("email"), html.Name("email"), html.Class("mt-1.5")),
					),
				),
				
				FormField(
					Props{
						Name:        "subject",
						Required:    true,
						Placeholder: "Brief description of your inquiry",
					},
					"Subject",
					"",
				),
				
				FormField(
					Props{
						Name:        "message",
						Required:    true,
						Rows:        5,
						MinLength:   20,
						Placeholder: "Please provide details about your inquiry...",
					},
					"Message",
					"Minimum 20 characters",
				),
				
				html.Div(html.Class("flex gap-2 pt-2"),
					html.Button(html.Type("submit"), html.Class("bg-primary text-primary-foreground"), g.Text("Send Message")),
					html.Button(html.Type("button"), html.Class("variant-outline"), g.Text("Cancel")),
				),
			),
		),
		
		// Advanced example with preview
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("With Live Preview")),
			html.P(html.Class("text-sm text-muted-foreground mb-4"), g.Text("(Add JavaScript to make preview functional)")),
			html.Div(html.Class("grid grid-cols-2 gap-4"),
				New(Props{
					ID:          "markdown-input",
					Name:        "markdown",
					Placeholder: "Type some markdown...",
					Rows:        8,
					OnInput:     "updatePreview(this.value)",
				}),
				html.Div(
					html.Class("rounded-md border bg-muted/50 p-4 min-h-[200px]"),
					html.ID("preview"),
					html.P(html.Class("text-muted-foreground"), g.Text("Preview will appear here...")),
				),
			),
		),
		
		// Message styles
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Common Use Cases")),
			html.Div(html.Class("space-y-4"),
				html.Div(
					html.P(html.Class("text-sm font-medium mb-2"), g.Text("Quick Message")),
					Message("quick-message", 3),
				),
				html.Div(
					html.P(html.Class("text-sm font-medium mb-2"), g.Text("Support Ticket")),
					New(Props{
						Name:        "ticket",
						Placeholder: "Describe your issue in detail. Include steps to reproduce if applicable.",
						Rows:        6,
						Required:    true,
					}),
				),
				html.Div(
					html.P(html.Class("text-sm font-medium mb-2"), g.Text("Code Input")),
					New(Props{
						Name:        "code",
						Placeholder: "Paste your code here...",
						Rows:        10,
						Class:       "font-mono text-xs",
					}),
				),
			),
		),
	)
}