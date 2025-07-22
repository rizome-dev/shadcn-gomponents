package datepicker

import (
	"time"

	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

// Example demonstrates various DatePicker configurations
func Example() g.Node {
	now := time.Now()
	selectedDate := time.Date(2024, 6, 15, 0, 0, 0, 0, time.UTC)

	return html.Div(
		html.Class("p-8 space-y-8"),

		// Basic date picker
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Basic Date Picker")),
			html.Div(html.Class("space-y-4"),
				// Default date picker
				html.Div(
					html.Label(html.Class("text-sm text-muted-foreground"), g.Text("Default")),
					New(Props{}),
				),

				// Date picker with selected value
				html.Div(
					html.Label(html.Class("text-sm text-muted-foreground"), g.Text("With Value")),
					New(Props{
						Value: selectedDate,
					}),
				),

				// Disabled date picker
				html.Div(
					html.Label(html.Class("text-sm text-muted-foreground"), g.Text("Disabled")),
					Disabled(selectedDate),
				),

				// Required date picker
				html.Div(
					html.Label(html.Class("text-sm text-muted-foreground"), g.Text("Required")),
					New(Props{
						Required:    true,
						Placeholder: "Required date",
					}),
				),
			),
		),

		// Date range picker
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Date Range Picker")),
			html.Div(html.Class("space-y-4"),
				// Default range picker
				html.Div(
					html.Label(html.Class("text-sm text-muted-foreground"), g.Text("Default Range")),
					WithRange(RangeProps{}),
				),

				// Range with selected dates
				html.Div(
					html.Label(html.Class("text-sm text-muted-foreground"), g.Text("With Selected Range")),
					WithRange(RangeProps{
						StartDate: time.Date(2024, 6, 1, 0, 0, 0, 0, time.UTC),
						EndDate:   time.Date(2024, 6, 30, 0, 0, 0, 0, time.UTC),
					}),
				),

				// Range with only start date
				html.Div(
					html.Label(html.Class("text-sm text-muted-foreground"), g.Text("Partial Range")),
					WithRange(RangeProps{
						StartDate: time.Date(2024, 6, 15, 0, 0, 0, 0, time.UTC),
					}),
				),
			),
		),

		// Date picker with presets
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Date Picker with Presets")),
			html.Div(html.Class("space-y-4"),
				// Default presets
				html.Div(
					html.Label(html.Class("text-sm text-muted-foreground"), g.Text("Default Presets")),
					WithPresets(PresetsProps{
						Value: now,
					}),
				),

				// Custom presets
				html.Div(
					html.Label(html.Class("text-sm text-muted-foreground"), g.Text("Custom Presets")),
					WithPresets(PresetsProps{
						Presets: []Preset{
							{Label: "Yesterday", Date: now.AddDate(0, 0, -1)},
							{Label: "Last Week", Date: now.AddDate(0, 0, -7)},
							{Label: "Last Month", Date: now.AddDate(0, -1, 0)},
							{Label: "Last Year", Date: now.AddDate(-1, 0, 0)},
						},
					}),
				),
			),
		),

		// Date picker with input
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Date Picker with Input")),
			html.Div(html.Class("space-y-4"),
				// Basic input
				WithInput(InputProps{
					Label: "Date of Birth",
					Name:  "dob",
				}),

				// Input with value and helper text
				WithInput(InputProps{
					Label:      "Event Date",
					Name:       "eventDate",
					Value:      selectedDate,
					HelperText: "Select the date for your event",
				}),

				// Required input with constraints
				WithInput(InputProps{
					Label:      "Appointment Date",
					Name:       "appointment",
					Required:   true,
					MinDate:    now,
					MaxDate:    now.AddDate(0, 3, 0), // 3 months from now
					HelperText: "Available dates are within the next 3 months",
				}),
			),
		),

		// Date picker with constraints
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Date Picker with Constraints")),
			html.Div(html.Class("space-y-4"),
				// Min date only
				html.Div(
					html.Label(html.Class("text-sm text-muted-foreground"), g.Text("Future Dates Only")),
					New(Props{
						Placeholder: "Select a future date",
						MinDate:     now,
					}),
				),

				// Max date only
				html.Div(
					html.Label(html.Class("text-sm text-muted-foreground"), g.Text("Past Dates Only")),
					New(Props{
						Placeholder: "Select a past date",
						MaxDate:     now,
					}),
				),

				// Min and max date
				html.Div(
					html.Label(html.Class("text-sm text-muted-foreground"), g.Text("Date Range Constraint")),
					WithMinMax(
						time.Time{},
						now.AddDate(0, -1, 0), // 1 month ago
						now.AddDate(0, 1, 0),  // 1 month from now
					),
				),
			),
		),

		// Different formats
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Different Date Formats")),
			html.Div(html.Class("space-y-4"),
				// ISO format
				html.Div(
					html.Label(html.Class("text-sm text-muted-foreground"), g.Text("ISO Format")),
					New(Props{
						Value:  selectedDate,
						Format: "2006-01-02",
					}),
				),

				// Long format
				html.Div(
					html.Label(html.Class("text-sm text-muted-foreground"), g.Text("Long Format")),
					New(Props{
						Value:  selectedDate,
						Format: "Monday, January 2, 2006",
					}),
				),

				// Short format
				html.Div(
					html.Label(html.Class("text-sm text-muted-foreground"), g.Text("Short Format")),
					New(Props{
						Value:  selectedDate,
						Format: "01/02/06",
					}),
				),

				// Custom format
				html.Div(
					html.Label(html.Class("text-sm text-muted-foreground"), g.Text("Custom Format")),
					New(Props{
						Value:  selectedDate,
						Format: "2 Jan 2006",
					}),
				),
			),
		),

		// Form integration example
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Form Integration")),
			html.Form(
				html.Class("space-y-4 max-w-md"),
				html.Method("POST"),
				html.Action("/submit"),

				// Start date
				html.Div(
					html.Label(
						html.For("start-date"),
						html.Class("text-sm font-medium"),
						g.Text("Start Date"),
					),
					New(Props{
						ID:       "start-date",
						Name:     "startDate",
						Required: true,
					}),
				),

				// End date
				html.Div(
					html.Label(
						html.For("end-date"),
						html.Class("text-sm font-medium"),
						g.Text("End Date"),
					),
					New(Props{
						ID:   "end-date",
						Name: "endDate",
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
					html.Li(g.Text("Date picker combines Calendar and Popover components")),
					html.Li(g.Text("Supports single date and date range selection")),
					html.Li(g.Text("Can be integrated with form inputs for submission")),
					html.Li(g.Text("Supports min/max date constraints")),
					html.Li(g.Text("Customizable date formats using Go's time format")),
					html.Li(g.Text("Preset options for quick date selection")),
					html.Li(g.Text("For HTMX integration, use the HTMX variant (coming soon)")),
				),
			),
		),

		// Code example
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Code Example")),
			html.Pre(html.Class("text-xs bg-muted p-4 rounded-lg overflow-x-auto"),
				html.Code(g.Raw(`// Basic date picker
datepicker.New(datepicker.Props{
    Value:       time.Now(),
    Placeholder: "Select date",
    Format:      "Jan 2, 2006",
})

// Date range picker
datepicker.WithRange(datepicker.RangeProps{
    StartDate: startDate,
    EndDate:   endDate,
})

// With presets
datepicker.WithPresets(datepicker.PresetsProps{
    Value: selectedDate,
    Presets: []datepicker.Preset{
        {Label: "Today", Date: time.Now()},
        {Label: "Tomorrow", Date: time.Now().AddDate(0, 0, 1)},
    },
})

// With input field
datepicker.WithInput(datepicker.InputProps{
    Label:      "Birth Date",
    Name:       "birthDate",
    Required:   true,
    HelperText: "Enter your date of birth",
})`)),
			),
		),
	)
}