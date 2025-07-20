package selector

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

// Example demonstrates how to use the Select component
func Example() g.Node {
	return html.Div(
		html.Class("p-8 space-y-8"),
		
		// Basic select
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Basic Select")),
			Simple(
				"fruit",
				[]OptionType{
					{Value: "apple", Label: "Apple"},
					{Value: "banana", Label: "Banana"},
					{Value: "orange", Label: "Orange"},
					{Value: "grape", Label: "Grape"},
				},
				"banana",
			),
		),
		
		// Select with placeholder
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("With Placeholder")),
			WithPlaceholder(
				"country",
				"Select a country",
				[]OptionType{
					{Value: "us", Label: "United States"},
					{Value: "uk", Label: "United Kingdom"},
					{Value: "ca", Label: "Canada"},
					{Value: "au", Label: "Australia"},
					{Value: "de", Label: "Germany"},
					{Value: "fr", Label: "France"},
				},
			),
		),
		
		// Grouped options
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Grouped Options")),
			WithGroups(
				"timezone",
				[]Group{
					{
						Label: "North America",
						Options: []OptionType{
							{Value: "pst", Label: "Pacific Standard Time (PST)"},
							{Value: "mst", Label: "Mountain Standard Time (MST)"},
							{Value: "cst", Label: "Central Standard Time (CST)"},
							{Value: "est", Label: "Eastern Standard Time (EST)"},
						},
					},
					{
						Label: "Europe",
						Options: []OptionType{
							{Value: "gmt", Label: "Greenwich Mean Time (GMT)"},
							{Value: "cet", Label: "Central European Time (CET)"},
							{Value: "eet", Label: "Eastern European Time (EET)"},
						},
					},
					{
						Label: "Asia",
						Options: []OptionType{
							{Value: "jst", Label: "Japan Standard Time (JST)"},
							{Value: "cst-china", Label: "China Standard Time (CST)"},
							{Value: "ist", Label: "India Standard Time (IST)"},
						},
					},
				},
				"est",
			),
		),
		
		// Form field with label and description
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Form Field")),
			FormField(
				Props{
					Name:     "department",
					Required: true,
					Options: []OptionType{
						{Value: "engineering", Label: "Engineering"},
						{Value: "design", Label: "Design"},
						{Value: "marketing", Label: "Marketing"},
						{Value: "sales", Label: "Sales"},
						{Value: "hr", Label: "Human Resources"},
						{Value: "finance", Label: "Finance"},
					},
				},
				"Department",
				"Select the department you work in",
			),
		),
		
		// Different sizes
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Different Sizes")),
			html.Div(html.Class("space-y-4"),
				html.Div(html.Class("space-y-2"),
					html.P(html.Class("text-sm text-muted-foreground"), g.Text("Small")),
					New(Props{
						Name: "size-small",
						Size: "sm",
						Options: []OptionType{
							{Value: "1", Label: "Option 1"},
							{Value: "2", Label: "Option 2"},
						},
					}),
				),
				html.Div(html.Class("space-y-2"),
					html.P(html.Class("text-sm text-muted-foreground"), g.Text("Default")),
					New(Props{
						Name: "size-default",
						Options: []OptionType{
							{Value: "1", Label: "Option 1"},
							{Value: "2", Label: "Option 2"},
						},
					}),
				),
				html.Div(html.Class("space-y-2"),
					html.P(html.Class("text-sm text-muted-foreground"), g.Text("Large")),
					New(Props{
						Name: "size-large",
						Size: "lg",
						Options: []OptionType{
							{Value: "1", Label: "Option 1"},
							{Value: "2", Label: "Option 2"},
						},
					}),
				),
			),
		),
		
		// Disabled states
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Disabled States")),
			html.Div(html.Class("space-y-4"),
				New(Props{
					Name:     "disabled-select",
					Disabled: true,
					Value:    "selected",
					Options: []OptionType{
						{Value: "selected", Label: "This select is disabled"},
					},
				}),
				New(Props{
					Name:        "partial-disabled",
					Placeholder: "Some options are disabled",
					Options: []OptionType{
						{Value: "enabled1", Label: "Enabled Option 1"},
						{Value: "disabled1", Label: "Disabled Option 1", Disabled: true},
						{Value: "enabled2", Label: "Enabled Option 2"},
						{Value: "disabled2", Label: "Disabled Option 2", Disabled: true},
					},
				}),
			),
		),
		
		// Multiple selection
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Multiple Selection")),
			New(Props{
				Name:     "skills",
				Multiple: true,
				Options: []OptionType{
					{Value: "html", Label: "HTML"},
					{Value: "css", Label: "CSS"},
					{Value: "js", Label: "JavaScript"},
					{Value: "go", Label: "Go"},
					{Value: "python", Label: "Python"},
					{Value: "rust", Label: "Rust"},
				},
			}),
		),
		
		// Custom styled select
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Custom Styled Select")),
			Custom(Props{
				Name:        "custom-select",
				Placeholder: "Choose your favorite framework",
				Value:       "react",
				Options: []OptionType{
					{Value: "react", Label: "React"},
					{Value: "vue", Label: "Vue"},
					{Value: "angular", Label: "Angular"},
					{Value: "svelte", Label: "Svelte"},
					{Value: "solid", Label: "Solid"},
				},
			}),
		),
		
		// In a form context
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("In Form Context")),
			html.Form(html.Class("max-w-md space-y-4 rounded-lg border bg-card p-6"),
				html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("User Preferences")),
				
				FormField(
					Props{
						Name:        "language",
						Placeholder: "Select language",
						Options: []OptionType{
							{Value: "en", Label: "English"},
							{Value: "es", Label: "Spanish"},
							{Value: "fr", Label: "French"},
							{Value: "de", Label: "German"},
							{Value: "it", Label: "Italian"},
						},
					},
					"Language",
					"",
				),
				
				FormField(
					Props{
						Name: "theme",
						Value: "system",
						Options: []OptionType{
							{Value: "light", Label: "Light"},
							{Value: "dark", Label: "Dark"},
							{Value: "system", Label: "System"},
						},
					},
					"Theme",
					"Select your preferred color theme",
				),
				
				FormField(
					Props{
						Name:     "notifications",
						Required: true,
						Options: []OptionType{
							{Value: "all", Label: "All notifications"},
							{Value: "important", Label: "Important only"},
							{Value: "none", Label: "None"},
						},
					},
					"Email Notifications",
					"How often would you like to receive emails?",
				),
				
				html.Div(html.Class("flex gap-2 pt-4"),
					html.Button(html.Type("submit"), html.Class("bg-primary text-primary-foreground"), g.Text("Save Preferences")),
					html.Button(html.Type("button"), html.Class("variant-outline"), g.Text("Cancel")),
				),
			),
		),
	)
}