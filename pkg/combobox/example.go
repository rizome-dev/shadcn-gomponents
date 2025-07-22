package combobox

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib/icons"
)

// Example demonstrates various Combobox configurations
func Example() g.Node {
	return html.Div(
		html.Class("p-8 space-y-8"),

		// Basic combobox
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Basic Combobox")),
			html.Div(html.Class("space-y-4"),
				// Default combobox
				html.Div(
					html.Label(html.Class("text-sm text-muted-foreground"), g.Text("Default")),
					New(Props{
						Options: []Option{
							{Value: "next", Label: "Next.js"},
							{Value: "sveltekit", Label: "SvelteKit"},
							{Value: "nuxt", Label: "Nuxt.js"},
							{Value: "remix", Label: "Remix"},
							{Value: "astro", Label: "Astro"},
						},
					}),
				),

				// With selected value
				html.Div(
					html.Label(html.Class("text-sm text-muted-foreground"), g.Text("With Value")),
					New(Props{
						Value: "nuxt",
						Options: []Option{
							{Value: "next", Label: "Next.js"},
							{Value: "sveltekit", Label: "SvelteKit"},
							{Value: "nuxt", Label: "Nuxt.js"},
							{Value: "remix", Label: "Remix"},
							{Value: "astro", Label: "Astro"},
						},
					}),
				),

				// Disabled
				html.Div(
					html.Label(html.Class("text-sm text-muted-foreground"), g.Text("Disabled")),
					New(Props{
						Value:    "next",
						Disabled: true,
						Options: []Option{
							{Value: "next", Label: "Next.js"},
							{Value: "sveltekit", Label: "SvelteKit"},
						},
					}),
				),

				// Custom placeholder
				html.Div(
					html.Label(html.Class("text-sm text-muted-foreground"), g.Text("Custom Placeholder")),
					New(Props{
						Placeholder: "Choose a framework...",
						Options: []Option{
							{Value: "next", Label: "Next.js"},
							{Value: "sveltekit", Label: "SvelteKit"},
							{Value: "nuxt", Label: "Nuxt.js"},
						},
					}),
				),
			),
		),

		// Combobox with icons
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Combobox with Icons")),
			html.Div(html.Class("space-y-4"),
				html.Div(
					html.Label(html.Class("text-sm text-muted-foreground"), g.Text("Status Selection")),
					WithIcons("", []Option{
						{
							Value: "backlog",
							Label: "Backlog",
							Icon:  html.Span(html.Class("h-4 w-4 text-muted-foreground"), g.Text("○")),
						},
						{
							Value: "todo",
							Label: "Todo",
							Icon:  icons.CircleIcon(html.Class("h-4 w-4")),
						},
						{
							Value: "in-progress",
							Label: "In Progress",
							Icon:  icons.Loader(html.Class("h-4 w-4 text-yellow-500")),
						},
						{
							Value: "done",
							Label: "Done",
							Icon:  icons.Check(html.Class("h-4 w-4 text-green-500")),
						},
						{
							Value: "canceled",
							Label: "Canceled",
							Icon:  icons.X(html.Class("h-4 w-4 text-red-500")),
						},
					}, "Set status..."),
				),

				html.Div(
					html.Label(html.Class("text-sm text-muted-foreground"), g.Text("Priority")),
					WithIcons("high", []Option{
						{
							Value: "low",
							Label: "Low",
							Icon:  icons.ChevronDown(html.Class("h-4 w-4 text-blue-500")),
						},
						{
							Value: "medium",
							Label: "Medium",
							Icon:  icons.ArrowRight(html.Class("h-4 w-4 text-yellow-500")),
						},
						{
							Value: "high",
							Label: "High",
							Icon:  icons.ChevronUp(html.Class("h-4 w-4 text-orange-500")),
						},
						{
							Value: "urgent",
							Label: "Urgent",
							Icon:  html.Span(html.Class("h-4 w-4 text-red-500"), g.Text("!")),
						},
					}),
				),
			),
		),

		// Multi-select combobox
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Multi-Select Combobox")),
			html.Div(html.Class("space-y-4"),
				// Default multi-select
				html.Div(
					html.Label(html.Class("text-sm text-muted-foreground"), g.Text("Select Languages")),
					Multi(MultiProps{
						Options: []Option{
							{Value: "go", Label: "Go"},
							{Value: "rust", Label: "Rust"},
							{Value: "python", Label: "Python"},
							{Value: "javascript", Label: "JavaScript"},
							{Value: "typescript", Label: "TypeScript"},
							{Value: "java", Label: "Java"},
							{Value: "csharp", Label: "C#"},
						},
					}),
				),

				// With selected values
				html.Div(
					html.Label(html.Class("text-sm text-muted-foreground"), g.Text("Pre-selected")),
					Multi(MultiProps{
						Values: []string{"go", "rust", "typescript"},
						Options: []Option{
							{Value: "go", Label: "Go"},
							{Value: "rust", Label: "Rust"},
							{Value: "python", Label: "Python"},
							{Value: "javascript", Label: "JavaScript"},
							{Value: "typescript", Label: "TypeScript"},
						},
					}),
				),

				// With max items
				html.Div(
					html.Label(html.Class("text-sm text-muted-foreground"), g.Text("Max 3 Items")),
					Multi(MultiProps{
						Values:   []string{"red", "blue"},
						MaxItems: 3,
						Options: []Option{
							{Value: "red", Label: "Red"},
							{Value: "blue", Label: "Blue"},
							{Value: "green", Label: "Green"},
							{Value: "yellow", Label: "Yellow"},
							{Value: "purple", Label: "Purple"},
							{Value: "orange", Label: "Orange"},
						},
						Placeholder: "Select up to 3 colors...",
					}),
				),
			),
		),

		// Grouped combobox
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Grouped Combobox")),
			html.Div(html.Class("space-y-4"),
				html.Div(
					html.Label(html.Class("text-sm text-muted-foreground"), g.Text("Select a Food")),
					WithGroups(GroupedProps{
						Groups: []OptionGroup{
							{
								Label: "Fruits",
								Options: []Option{
									{Value: "apple", Label: "Apple"},
									{Value: "banana", Label: "Banana"},
									{Value: "orange", Label: "Orange"},
									{Value: "grape", Label: "Grape"},
									{Value: "strawberry", Label: "Strawberry"},
								},
							},
							{
								Label: "Vegetables",
								Options: []Option{
									{Value: "carrot", Label: "Carrot"},
									{Value: "broccoli", Label: "Broccoli"},
									{Value: "spinach", Label: "Spinach"},
									{Value: "tomato", Label: "Tomato"},
									{Value: "cucumber", Label: "Cucumber"},
								},
							},
							{
								Label: "Grains",
								Options: []Option{
									{Value: "rice", Label: "Rice"},
									{Value: "wheat", Label: "Wheat"},
									{Value: "oats", Label: "Oats"},
									{Value: "quinoa", Label: "Quinoa"},
								},
							},
						},
					}),
				),

				html.Div(
					html.Label(html.Class("text-sm text-muted-foreground"), g.Text("Team Members by Department")),
					WithGroups(GroupedProps{
						Value: "john",
						Groups: []OptionGroup{
							{
								Label: "Engineering",
								Options: []Option{
									{Value: "john", Label: "John Doe"},
									{Value: "jane", Label: "Jane Smith"},
									{Value: "bob", Label: "Bob Johnson"},
								},
							},
							{
								Label: "Design",
								Options: []Option{
									{Value: "alice", Label: "Alice Brown"},
									{Value: "charlie", Label: "Charlie Wilson"},
								},
							},
							{
								Label: "Marketing",
								Options: []Option{
									{Value: "eve", Label: "Eve Davis"},
									{Value: "frank", Label: "Frank Miller"},
								},
							},
						},
						Width: "w-[250px]",
					}),
				),
			),
		),

		// Searchable combobox
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Searchable Combobox")),
			html.Div(html.Class("space-y-4"),
				html.Div(
					html.Label(html.Class("text-sm text-muted-foreground"), g.Text("Large Dataset")),
					Searchable("", []Option{
						{Value: "us", Label: "United States"},
						{Value: "ca", Label: "Canada"},
						{Value: "mx", Label: "Mexico"},
						{Value: "uk", Label: "United Kingdom"},
						{Value: "de", Label: "Germany"},
						{Value: "fr", Label: "France"},
						{Value: "it", Label: "Italy"},
						{Value: "es", Label: "Spain"},
						{Value: "jp", Label: "Japan"},
						{Value: "cn", Label: "China"},
						{Value: "kr", Label: "South Korea"},
						{Value: "in", Label: "India"},
						{Value: "br", Label: "Brazil"},
						{Value: "ar", Label: "Argentina"},
						{Value: "au", Label: "Australia"},
						{Value: "nz", Label: "New Zealand"},
					}, "Search countries..."),
				),
			),
		),

		// Different sizes
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Different Sizes")),
			html.Div(html.Class("space-y-4"),
				html.Div(
					html.Label(html.Class("text-sm text-muted-foreground"), g.Text("Small (200px)")),
					New(Props{
						Width: "w-[200px]",
						Options: []Option{
							{Value: "1", Label: "Option 1"},
							{Value: "2", Label: "Option 2"},
						},
					}),
				),

				html.Div(
					html.Label(html.Class("text-sm text-muted-foreground"), g.Text("Medium (300px)")),
					New(Props{
						Width: "w-[300px]",
						Options: []Option{
							{Value: "1", Label: "Option 1"},
							{Value: "2", Label: "Option 2"},
						},
					}),
				),

				html.Div(
					html.Label(html.Class("text-sm text-muted-foreground"), g.Text("Full Width")),
					New(Props{
						Width: "w-full",
						Options: []Option{
							{Value: "1", Label: "Option 1"},
							{Value: "2", Label: "Option 2"},
						},
					}),
				),
			),
		),

		// Form integration
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Form Integration")),
			html.Form(
				html.Class("space-y-4 max-w-md"),
				html.Method("POST"),
				html.Action("/submit"),

				// Single select
				html.Div(
					html.Label(
						html.For("framework"),
						html.Class("text-sm font-medium"),
						g.Text("Framework"),
					),
					New(Props{
						ID:   "framework",
						Name: "framework",
						Options: []Option{
							{Value: "next", Label: "Next.js"},
							{Value: "sveltekit", Label: "SvelteKit"},
							{Value: "nuxt", Label: "Nuxt.js"},
							{Value: "remix", Label: "Remix"},
						},
						Width: "w-full",
					}),
				),

				// Multi-select
				html.Div(
					html.Label(
						html.For("features"),
						html.Class("text-sm font-medium"),
						g.Text("Features"),
					),
					Multi(MultiProps{
						ID:   "features",
						Name: "features[]",
						Options: []Option{
							{Value: "ssr", Label: "Server-Side Rendering"},
							{Value: "ssg", Label: "Static Site Generation"},
							{Value: "api", Label: "API Routes"},
							{Value: "auth", Label: "Authentication"},
							{Value: "db", Label: "Database Integration"},
						},
						Width: "w-full",
					}),
				),

				// Submit button
				html.Button(
					html.Type("submit"),
					html.Class("px-4 py-2 bg-primary text-primary-foreground rounded-md hover:bg-primary/90"),
					g.Text("Submit"),
				),
			),
		),

		// Usage notes
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Usage Notes")),
			html.Div(html.Class("rounded-lg border bg-muted/50 p-4"),
				html.Ul(html.Class("text-sm space-y-2 list-disc list-inside"),
					html.Li(g.Text("Combobox combines a button, popover, and command menu")),
					html.Li(g.Text("Supports single and multi-select modes")),
					html.Li(g.Text("Options can have icons and be disabled")),
					html.Li(g.Text("Searchable variant helps with large datasets")),
					html.Li(g.Text("Grouped options organize related items")),
					html.Li(g.Text("Selected items show a check mark")),
					html.Li(g.Text("For HTMX integration, use the HTMX variant (coming soon)")),
				),
			),
		),

		// Code example
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Code Example")),
			html.Pre(html.Class("text-xs bg-muted p-4 rounded-lg overflow-x-auto"),
				html.Code(g.Raw(`// Basic combobox
combobox.New(combobox.Props{
    Value: "selected-value",
    Options: []combobox.Option{
        {Value: "1", Label: "Option 1"},
        {Value: "2", Label: "Option 2"},
    },
})

// With icons
combobox.WithIcons("", []combobox.Option{
    {
        Value: "user",
        Label: "User",
        Icon:  icons.User(html.Class("h-4 w-4")),
    },
})

// Multi-select
combobox.Multi(combobox.MultiProps{
    Values: []string{"opt1", "opt2"},
    Options: options,
    MaxItems: 5,
})

// Grouped
combobox.WithGroups(combobox.GroupedProps{
    Groups: []combobox.OptionGroup{
        {
            Label: "Group 1",
            Options: options1,
        },
        {
            Label: "Group 2",
            Options: options2,
        },
    },
})`)),
			),
		),
	)
}