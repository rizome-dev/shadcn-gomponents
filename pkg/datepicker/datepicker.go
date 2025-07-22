package datepicker

import (
	"fmt"
	"time"

	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
	"github.com/rizome-dev/shadcn-gomponents/lib/icons"
	"github.com/rizome-dev/shadcn-gomponents/pkg/button"
	"github.com/rizome-dev/shadcn-gomponents/pkg/calendar"
	"github.com/rizome-dev/shadcn-gomponents/pkg/popover"
)

// Props defines the properties for a DatePicker component
type Props struct {
	ID          string    // ID for the input field
	Name        string    // Name for form submission
	Value       time.Time // Selected date
	Placeholder string    // Placeholder text
	Format      string    // Date format (e.g., "Jan 2, 2006")
	MinDate     time.Time // Minimum selectable date
	MaxDate     time.Time // Maximum selectable date
	Disabled    bool      // Whether the date picker is disabled
	Required    bool      // Whether the field is required
	Open        bool      // Whether the popover is open
	Class       string    // Additional CSS classes
	OnSelect    string    // JavaScript to run on date selection
}

// New creates a new DatePicker component
func New(props Props) g.Node {
	// Set defaults
	if props.Format == "" {
		props.Format = "Jan 2, 2006"
	}
	if props.Placeholder == "" {
		props.Placeholder = "Pick a date"
	}
	if props.ID == "" {
		props.ID = fmt.Sprintf("datepicker-%d", time.Now().UnixNano())
	}

	// Format the display value
	displayValue := props.Placeholder
	if !props.Value.IsZero() {
		displayValue = props.Value.Format(props.Format)
	}

	// Build the date picker using popover and calendar
	return popover.New(
		popover.Props{
			Open:  props.Open,
			Class: props.Class,
		},
		// Trigger button
		popover.Trigger(
			popover.TriggerProps{
				AsChild: true,
			},
			button.New(
				button.Props{
					Variant: "outline",
					Class: lib.CN(
						"w-[240px] justify-start text-left font-normal",
						func() string {
							if props.Value.IsZero() {
								return "text-muted-foreground"
							}
							return ""
						}(),
					),
					Disabled: props.Disabled,
				},
				g.If(props.ID != "", html.ID(props.ID)),
				g.If(props.Name != "", html.Name(props.Name)),
				g.If(props.Required, html.Required()),
				// Calendar icon
				icons.Calendar(html.Class("mr-2 h-4 w-4")),
				g.Text(displayValue),
			),
		),
		// Popover content with calendar
		popover.ContentComponent(
			popover.ContentProps{
				Side:  "bottom",
				Align: "start",
				Class: "w-auto p-0",
			},
			calendar.New(calendar.Props{
				Value:   props.Value,
				Month:   func() time.Time {
					if !props.Value.IsZero() {
						return props.Value
					}
					return time.Now()
				}(),
				MinDate: props.MinDate,
				MaxDate: props.MaxDate,
			}),
		),
	)
}

// WithRange creates a date picker for selecting a date range
func WithRange(props RangeProps) g.Node {
	// Set defaults
	if props.Format == "" {
		props.Format = "Jan 2, 2006"
	}
	if props.Placeholder == "" {
		props.Placeholder = "Pick a date range"
	}
	if props.ID == "" {
		props.ID = fmt.Sprintf("daterangepicker-%d", time.Now().UnixNano())
	}

	// Format the display value
	displayValue := props.Placeholder
	if !props.StartDate.IsZero() && !props.EndDate.IsZero() {
		displayValue = fmt.Sprintf("%s - %s",
			props.StartDate.Format(props.Format),
			props.EndDate.Format(props.Format),
		)
	} else if !props.StartDate.IsZero() {
		displayValue = fmt.Sprintf("%s - ...", props.StartDate.Format(props.Format))
	}

	// Build the date range picker
	return popover.New(
		popover.Props{
			Open:  props.Open,
			Class: props.Class,
		},
		// Trigger button
		popover.Trigger(
			popover.TriggerProps{
				AsChild: true,
			},
			button.New(
				button.Props{
					Variant: "outline",
					Class: lib.CN(
						"w-[300px] justify-start text-left font-normal",
						func() string {
							if props.StartDate.IsZero() && props.EndDate.IsZero() {
								return "text-muted-foreground"
							}
							return ""
						}(),
					),
					Disabled: props.Disabled,
				},
				g.If(props.ID != "", html.ID(props.ID)),
				// Calendar icon
				icons.Calendar(html.Class("mr-2 h-4 w-4")),
				g.Text(displayValue),
			),
		),
		// Popover content with calendar - showing two months for range
		popover.ContentComponent(
			popover.ContentProps{
				Side:  "bottom",
				Align: "start",
				Class: "w-auto p-0",
			},
			html.Div(
				html.Class("flex gap-4 p-3"),
				calendar.New(calendar.Props{
					Value:   props.StartDate,
					Month:   func() time.Time {
						if !props.StartDate.IsZero() {
							return props.StartDate
						}
						return time.Now()
					}(),
					MinDate: props.MinDate,
					MaxDate: props.MaxDate,
				}),
				calendar.New(calendar.Props{
					Value:   props.EndDate,
					Month:   func() time.Time {
						if !props.StartDate.IsZero() {
							return props.StartDate.AddDate(0, 1, 0)
						}
						return time.Now().AddDate(0, 1, 0)
					}(),
					MinDate: props.MinDate,
					MaxDate: props.MaxDate,
				}),
			),
		),
	)
}

// RangeProps defines properties for a date range picker
type RangeProps struct {
	ID          string    // ID for the input field
	StartDate   time.Time // Start date of the range
	EndDate     time.Time // End date of the range
	Placeholder string    // Placeholder text
	Format      string    // Date format
	MinDate     time.Time // Minimum selectable date
	MaxDate     time.Time // Maximum selectable date
	Disabled    bool      // Whether the picker is disabled
	Open        bool      // Whether the popover is open
	Class       string    // Additional CSS classes
	OnSelect    string    // JavaScript to run on selection
}

// WithPresets creates a date picker with preset date options
func WithPresets(props PresetsProps) g.Node {
	// Set defaults
	if props.Format == "" {
		props.Format = "Jan 2, 2006"
	}
	if props.Placeholder == "" {
		props.Placeholder = "Select a date"
	}
	if props.ID == "" {
		props.ID = fmt.Sprintf("datepicker-presets-%d", time.Now().UnixNano())
	}

	// Set default presets if none provided
	if len(props.Presets) == 0 {
		now := time.Now()
		props.Presets = []Preset{
			{Label: "Today", Date: now},
			{Label: "Tomorrow", Date: now.AddDate(0, 0, 1)},
			{Label: "In a week", Date: now.AddDate(0, 0, 7)},
			{Label: "In a month", Date: now.AddDate(0, 1, 0)},
		}
	}

	// Format the display value
	displayValue := props.Placeholder
	if !props.Value.IsZero() {
		displayValue = props.Value.Format(props.Format)
	}

	// Build the date picker with presets
	return popover.New(
		popover.Props{
			Open:  props.Open,
			Class: props.Class,
		},
		// Trigger button
		popover.Trigger(
			popover.TriggerProps{
				AsChild: true,
			},
			button.New(
				button.Props{
					Variant: "outline",
					Class: lib.CN(
						"w-[240px] justify-start text-left font-normal",
						func() string {
							if props.Value.IsZero() {
								return "text-muted-foreground"
							}
							return ""
						}(),
					),
					Disabled: props.Disabled,
				},
				g.If(props.ID != "", html.ID(props.ID)),
				// Calendar icon
				icons.Calendar(html.Class("mr-2 h-4 w-4")),
				g.Text(displayValue),
			),
		),
		// Popover content with presets and calendar
		popover.ContentComponent(
			popover.ContentProps{
				Side:  "bottom",
				Align: "start",
				Class: "flex w-auto flex-col space-y-2 p-2",
			},
			// Presets section
			html.Div(
				html.Class("flex flex-col space-y-1"),
				g.Group(g.Map(props.Presets, func(preset Preset) g.Node {
					isSelected := !props.Value.IsZero() && props.Value.Format("2006-01-02") == preset.Date.Format("2006-01-02")
					return button.New(
						button.Props{
							Variant: "ghost",
							Class: lib.CN(
								"justify-start text-left font-normal",
								func() string {
									if isSelected {
										return "bg-accent text-accent-foreground"
									}
									return ""
								}(),
							),
						},
						g.If(preset.OnClick != "", g.Attr("onclick", preset.OnClick)),
						g.Text(preset.Label),
					)
				})),
			),
			// Separator
			html.Div(html.Class("rounded-md border")),
			// Calendar
			calendar.New(calendar.Props{
				Value:   props.Value,
				Month:   func() time.Time {
					if !props.Value.IsZero() {
						return props.Value
					}
					return time.Now()
				}(),
				MinDate: props.MinDate,
				MaxDate: props.MaxDate,
			}),
		),
	)
}

// PresetsProps defines properties for a date picker with presets
type PresetsProps struct {
	ID          string    // ID for the input field
	Name        string    // Name for form submission
	Value       time.Time // Selected date
	Placeholder string    // Placeholder text
	Format      string    // Date format
	MinDate     time.Time // Minimum selectable date
	MaxDate     time.Time // Maximum selectable date
	Disabled    bool      // Whether the picker is disabled
	Open        bool      // Whether the popover is open
	Class       string    // Additional CSS classes
	OnSelect    string    // JavaScript to run on selection
	Presets     []Preset  // Preset date options
}

// Preset defines a preset date option
type Preset struct {
	Label   string    // Display label
	Date    time.Time // The preset date
	OnClick string    // Optional JavaScript handler
}

// WithInput creates a date picker with a text input field
func WithInput(props InputProps) g.Node {
	// Set defaults
	if props.Format == "" {
		props.Format = "2006-01-02" // ISO format for input
	}
	if props.DisplayFormat == "" {
		props.DisplayFormat = "Jan 2, 2006"
	}
	if props.Placeholder == "" {
		props.Placeholder = "YYYY-MM-DD"
	}
	if props.ID == "" {
		props.ID = fmt.Sprintf("datepicker-input-%d", time.Now().UnixNano())
	}

	// Format the input value
	inputValue := ""
	if !props.Value.IsZero() {
		inputValue = props.Value.Format(props.Format)
	}

	return html.Div(
		html.Class(lib.CN("grid gap-2", props.Class)),
		// Label if provided
		g.If(props.Label != "", html.Label(
			html.For(props.ID),
			html.Class("text-sm font-medium"),
			g.Text(props.Label),
		)),
		// Date picker popover
		popover.New(
			popover.Props{
				Open: props.Open,
			},
			// Input trigger
			popover.Trigger(
				popover.TriggerProps{
					AsChild: true,
				},
				html.Div(
					html.Class("relative"),
					html.Input(
						html.Type("text"),
						html.ID(props.ID),
						g.If(props.Name != "", html.Name(props.Name)),
						html.Value(inputValue),
						html.Placeholder(props.Placeholder),
						html.Class(lib.CN(
							"flex h-9 w-full rounded-md border border-input bg-transparent px-3 py-1 text-sm shadow-sm transition-colors",
							"file:border-0 file:bg-transparent file:text-sm file:font-medium",
							"placeholder:text-muted-foreground",
							"focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring",
							"disabled:cursor-not-allowed disabled:opacity-50",
							"pr-10", // Space for icon
						)),
						g.If(props.Disabled, html.Disabled()),
						g.If(props.Required, html.Required()),
						g.If(props.OnChange != "", g.Attr("onchange", props.OnChange)),
					),
					// Calendar icon button
					html.Button(
						html.Type("button"),
						html.Class("absolute right-0 top-0 h-full px-3 py-1 hover:bg-transparent"),
						g.If(props.Disabled, html.Disabled()),
						icons.Calendar(html.Class("h-4 w-4 opacity-50")),
					),
				),
			),
			// Calendar popover
			popover.ContentComponent(
				popover.ContentProps{
					Side:  "bottom",
					Align: "start",
					Class: "w-auto p-0",
				},
				calendar.New(calendar.Props{
					Value:   props.Value,
					Month:   func() time.Time {
						if !props.Value.IsZero() {
							return props.Value
						}
						return time.Now()
					}(),
					MinDate: props.MinDate,
					MaxDate: props.MaxDate,
				}),
			),
		),
		// Helper text if provided
		g.If(props.HelperText != "", html.P(
			html.Class("text-sm text-muted-foreground"),
			g.Text(props.HelperText),
		)),
	)
}

// InputProps defines properties for a date picker with input field
type InputProps struct {
	ID            string    // ID for the input field
	Name          string    // Name for form submission
	Value         time.Time // Selected date
	Label         string    // Label text
	Placeholder   string    // Placeholder text
	Format        string    // Input value format
	DisplayFormat string    // Display format in calendar
	HelperText    string    // Helper text below input
	MinDate       time.Time // Minimum selectable date
	MaxDate       time.Time // Maximum selectable date
	Disabled      bool      // Whether the picker is disabled
	Required      bool      // Whether the field is required
	Open          bool      // Whether the popover is open
	Class         string    // Additional CSS classes
	OnSelect      string    // JavaScript to run on selection
	OnChange      string    // JavaScript to run on input change
}

// Simple creates a simple date picker button
func Simple(value time.Time, onSelect ...string) g.Node {
	var onSelectHandler string
	if len(onSelect) > 0 {
		onSelectHandler = onSelect[0]
	}

	return New(Props{
		Value:    value,
		OnSelect: onSelectHandler,
	})
}

// Disabled creates a disabled date picker
func Disabled(value time.Time) g.Node {
	return New(Props{
		Value:    value,
		Disabled: true,
	})
}

// WithMinMax creates a date picker with min/max constraints
func WithMinMax(value, min, max time.Time, onSelect ...string) g.Node {
	var onSelectHandler string
	if len(onSelect) > 0 {
		onSelectHandler = onSelect[0]
	}

	return New(Props{
		Value:    value,
		MinDate:  min,
		MaxDate:  max,
		OnSelect: onSelectHandler,
	})
}