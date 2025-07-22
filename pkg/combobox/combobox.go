package combobox

import (
	"fmt"
	"time"

	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
	"github.com/rizome-dev/shadcn-gomponents/lib/icons"
	"github.com/rizome-dev/shadcn-gomponents/pkg/button"
	"github.com/rizome-dev/shadcn-gomponents/pkg/popover"
)

// Option represents a single option in the combobox
type Option struct {
	Value    string // The value to be submitted
	Label    string // The display text
	Icon     g.Node // Optional icon
	Disabled bool   // Whether the option is disabled
}

// Props defines the properties for the Combobox component
type Props struct {
	ID          string   // ID for the input field
	Name        string   // Name for form submission
	Value       string   // Currently selected value
	Options     []Option // Available options
	Placeholder string   // Placeholder text when no value selected
	SearchPlaceholder string // Placeholder for search input
	EmptyText   string   // Text to show when no options match
	Open        bool     // Whether the popover is open
	Disabled    bool     // Whether the combobox is disabled
	Class       string   // Additional CSS classes
	Width       string   // Width of the combobox (e.g., "200px", "w-full")
	OnSelect    string   // JavaScript to run on selection
}

// New creates a new Combobox component
func New(props Props) g.Node {
	// Set defaults
	if props.Placeholder == "" {
		props.Placeholder = "Select option..."
	}
	if props.SearchPlaceholder == "" {
		props.SearchPlaceholder = "Search..."
	}
	if props.EmptyText == "" {
		props.EmptyText = "No results found."
	}
	if props.ID == "" {
		props.ID = fmt.Sprintf("combobox-%d", time.Now().UnixNano())
	}
	if props.Width == "" {
		props.Width = "w-[200px]"
	}

	// Find selected label
	selectedLabel := props.Placeholder
	for _, opt := range props.Options {
		if opt.Value == props.Value {
			selectedLabel = opt.Label
			break
		}
	}

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
						props.Width,
						"justify-between",
						func() string {
							if props.Value == "" {
								return "text-muted-foreground"
							}
							return ""
						}(),
					),
					Disabled: props.Disabled,
				},
				g.If(props.ID != "", html.ID(props.ID)),
				g.If(props.Name != "", html.Name(props.Name)),
				g.Attr("role", "combobox"),
				g.Attr("aria-expanded", func() string {
					if props.Open {
						return "true"
					}
					return "false"
				}()),
				g.Text(selectedLabel),
				icons.ChevronsUpDown(html.Class("ml-2 h-4 w-4 shrink-0 opacity-50")),
			),
		),
		// Popover content with options list
		popover.ContentComponent(
			popover.ContentProps{
				Side:  "bottom",
				Align: "start",
				Class: lib.CN("p-0", props.Width),
			},
			html.Div(
				html.Class("max-h-[300px] overflow-auto"),
				// Search input
				html.Div(
					html.Class("flex items-center border-b px-3"),
					icons.Search(html.Class("mr-2 h-4 w-4 shrink-0 opacity-50")),
					html.Input(
						html.Type("text"),
						html.Placeholder(props.SearchPlaceholder),
						html.Class("flex h-9 w-full rounded-md bg-transparent py-3 text-sm outline-none placeholder:text-muted-foreground disabled:cursor-not-allowed disabled:opacity-50"),
					),
				),
				// Options list
				html.Div(
					html.Class("py-1"),
					g.Attr("role", "listbox"),
					g.If(len(props.Options) == 0,
						html.Div(
							html.Class("py-6 text-center text-sm text-muted-foreground"),
							g.Text(props.EmptyText),
						),
					),
					g.Group(g.Map(props.Options, func(opt Option) g.Node {
						isSelected := opt.Value == props.Value
						return html.Div(
							g.Attr("role", "option"),
							g.Attr("aria-selected", func() string {
								if isSelected {
									return "true"
								}
								return "false"
							}()),
							html.Class(lib.CN(
								"relative flex cursor-default select-none items-center rounded-sm px-2 py-1.5 text-sm outline-none",
								"hover:bg-accent hover:text-accent-foreground",
								"data-[disabled=true]:pointer-events-none data-[disabled=true]:opacity-50",
								func() string {
									if isSelected {
										return "bg-accent text-accent-foreground"
									}
									return ""
								}(),
							)),
							g.If(opt.Disabled, g.Attr("data-disabled", "true")),
							g.If(props.OnSelect != "", g.Attr("onclick", props.OnSelect)),
							g.If(opt.Icon != nil, html.Span(
								html.Class("mr-2 h-4 w-4"),
								opt.Icon,
							)),
							g.Text(opt.Label),
							g.If(isSelected, icons.Check(
								html.Class("ml-auto h-4 w-4"),
							)),
						)
					})),
				),
			),
		),
	)
}

// MultiProps defines properties for a multi-select combobox
type MultiProps struct {
	ID          string   // ID for the input field
	Name        string   // Name for form submission
	Values      []string // Currently selected values
	Options     []Option // Available options
	Placeholder string   // Placeholder text when no values selected
	SearchPlaceholder string // Placeholder for search input
	EmptyText   string   // Text to show when no options match
	MaxItems    int      // Maximum number of items that can be selected (0 = unlimited)
	Open        bool     // Whether the popover is open
	Disabled    bool     // Whether the combobox is disabled
	Class       string   // Additional CSS classes
	Width       string   // Width of the combobox
	OnSelect    string   // JavaScript to run on selection
}

// Multi creates a multi-select combobox
func Multi(props MultiProps) g.Node {
	// Set defaults
	if props.Placeholder == "" {
		props.Placeholder = "Select options..."
	}
	if props.SearchPlaceholder == "" {
		props.SearchPlaceholder = "Search..."
	}
	if props.EmptyText == "" {
		props.EmptyText = "No results found."
	}
	if props.ID == "" {
		props.ID = fmt.Sprintf("multi-combobox-%d", time.Now().UnixNano())
	}
	if props.Width == "" {
		props.Width = "w-[280px]"
	}

	// Build display text
	displayText := props.Placeholder
	if len(props.Values) > 0 {
		if len(props.Values) == 1 {
			// Find label for single value
			for _, opt := range props.Options {
				if opt.Value == props.Values[0] {
					displayText = opt.Label
					break
				}
			}
		} else {
			displayText = fmt.Sprintf("%d selected", len(props.Values))
		}
	}

	// Check if value is selected
	isSelected := func(value string) bool {
		for _, v := range props.Values {
			if v == value {
				return true
			}
		}
		return false
	}

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
						props.Width,
						"justify-between",
						func() string {
							if len(props.Values) == 0 {
								return "text-muted-foreground"
							}
							return ""
						}(),
					),
					Disabled: props.Disabled,
				},
				g.If(props.ID != "", html.ID(props.ID)),
				g.If(props.Name != "", html.Name(props.Name)),
				g.Attr("role", "combobox"),
				g.Attr("aria-expanded", func() string {
					if props.Open {
						return "true"
					}
					return "false"
				}()),
				html.Span(html.Class("truncate"), g.Text(displayText)),
				icons.ChevronsUpDown(html.Class("ml-2 h-4 w-4 shrink-0 opacity-50")),
			),
		),
		// Popover content with options list
		popover.ContentComponent(
			popover.ContentProps{
				Side:  "bottom",
				Align: "start",
				Class: lib.CN("p-0", props.Width),
			},
			html.Div(
				html.Class("max-h-[300px] overflow-auto"),
				// Search input
				html.Div(
					html.Class("flex items-center border-b px-3"),
					icons.Search(html.Class("mr-2 h-4 w-4 shrink-0 opacity-50")),
					html.Input(
						html.Type("text"),
						html.Placeholder(props.SearchPlaceholder),
						html.Class("flex h-9 w-full rounded-md bg-transparent py-3 text-sm outline-none placeholder:text-muted-foreground disabled:cursor-not-allowed disabled:opacity-50"),
					),
				),
				// Options list
				html.Div(
					html.Class("py-1"),
					g.Attr("role", "listbox"),
					g.If(len(props.Options) == 0,
						html.Div(
							html.Class("py-6 text-center text-sm text-muted-foreground"),
							g.Text(props.EmptyText),
						),
					),
					g.Group(g.Map(props.Options, func(opt Option) g.Node {
						selected := isSelected(opt.Value)
						canSelect := props.MaxItems == 0 || len(props.Values) < props.MaxItems || selected
						
						return html.Div(
							g.Attr("role", "option"),
							g.Attr("aria-selected", func() string {
								if selected {
									return "true"
								}
								return "false"
							}()),
							html.Class(lib.CN(
								"relative flex cursor-default select-none items-center rounded-sm px-2 py-1.5 text-sm outline-none",
								"hover:bg-accent hover:text-accent-foreground",
								"data-[disabled=true]:pointer-events-none data-[disabled=true]:opacity-50",
								func() string {
									if selected {
										return "bg-accent text-accent-foreground"
									}
									return ""
								}(),
							)),
							g.If(opt.Disabled || (!selected && !canSelect), g.Attr("data-disabled", "true")),
							g.If(props.OnSelect != "", g.Attr("onclick", props.OnSelect)),
							g.If(opt.Icon != nil, html.Span(
								html.Class("mr-2 h-4 w-4"),
								opt.Icon,
							)),
							g.Text(opt.Label),
							g.If(selected, icons.Check(
								html.Class("ml-auto h-4 w-4"),
							)),
						)
					})),
				),
			),
		),
	)
}

// WithGroups creates a combobox with grouped options
func WithGroups(props GroupedProps) g.Node {
	// Set defaults
	if props.Placeholder == "" {
		props.Placeholder = "Select option..."
	}
	if props.SearchPlaceholder == "" {
		props.SearchPlaceholder = "Search..."
	}
	if props.EmptyText == "" {
		props.EmptyText = "No results found."
	}
	if props.ID == "" {
		props.ID = fmt.Sprintf("grouped-combobox-%d", time.Now().UnixNano())
	}
	if props.Width == "" {
		props.Width = "w-[200px]"
	}

	// Find selected label
	selectedLabel := props.Placeholder
	for _, group := range props.Groups {
		for _, opt := range group.Options {
			if opt.Value == props.Value {
				selectedLabel = opt.Label
				break
			}
		}
		if selectedLabel != props.Placeholder {
			break
		}
	}

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
						props.Width,
						"justify-between",
						func() string {
							if props.Value == "" {
								return "text-muted-foreground"
							}
							return ""
						}(),
					),
					Disabled: props.Disabled,
				},
				g.If(props.ID != "", html.ID(props.ID)),
				g.If(props.Name != "", html.Name(props.Name)),
				g.Attr("role", "combobox"),
				g.Attr("aria-expanded", func() string {
					if props.Open {
						return "true"
					}
					return "false"
				}()),
				g.Text(selectedLabel),
				icons.ChevronsUpDown(html.Class("ml-2 h-4 w-4 shrink-0 opacity-50")),
			),
		),
		// Popover content with grouped options
		popover.ContentComponent(
			popover.ContentProps{
				Side:  "bottom",
				Align: "start",
				Class: lib.CN("p-0", props.Width),
			},
			html.Div(
				html.Class("max-h-[300px] overflow-auto"),
				// Search input
				html.Div(
					html.Class("flex items-center border-b px-3"),
					icons.Search(html.Class("mr-2 h-4 w-4 shrink-0 opacity-50")),
					html.Input(
						html.Type("text"),
						html.Placeholder(props.SearchPlaceholder),
						html.Class("flex h-9 w-full rounded-md bg-transparent py-3 text-sm outline-none placeholder:text-muted-foreground disabled:cursor-not-allowed disabled:opacity-50"),
					),
				),
				// Grouped options list
				html.Div(
					html.Class("py-1"),
					g.Attr("role", "listbox"),
					g.If(len(props.Groups) == 0,
						html.Div(
							html.Class("py-6 text-center text-sm text-muted-foreground"),
							g.Text(props.EmptyText),
						),
					),
					g.Group(g.Map(props.Groups, func(group OptionGroup) g.Node {
						return html.Div(
							// Group label
							html.Div(
								html.Class("px-2 py-1.5 text-xs font-semibold text-muted-foreground"),
								g.Text(group.Label),
							),
							// Group items
							g.Group(g.Map(group.Options, func(opt Option) g.Node {
								isSelected := opt.Value == props.Value
								return html.Div(
									g.Attr("role", "option"),
									g.Attr("aria-selected", func() string {
										if isSelected {
											return "true"
										}
										return "false"
									}()),
									html.Class(lib.CN(
										"relative flex cursor-default select-none items-center rounded-sm px-2 py-1.5 text-sm outline-none",
										"hover:bg-accent hover:text-accent-foreground",
										"data-[disabled=true]:pointer-events-none data-[disabled=true]:opacity-50",
										func() string {
											if isSelected {
												return "bg-accent text-accent-foreground"
											}
											return ""
										}(),
									)),
									g.If(opt.Disabled, g.Attr("data-disabled", "true")),
									g.If(props.OnSelect != "", g.Attr("onclick", props.OnSelect)),
									g.If(opt.Icon != nil, html.Span(
										html.Class("mr-2 h-4 w-4"),
										opt.Icon,
									)),
									g.Text(opt.Label),
									g.If(isSelected, icons.Check(
										html.Class("ml-auto h-4 w-4"),
									)),
								)
							})),
						)
					})),
				),
			),
		),
	)
}

// GroupedProps defines properties for a combobox with grouped options
type GroupedProps struct {
	ID          string        // ID for the input field
	Name        string        // Name for form submission
	Value       string        // Currently selected value
	Groups      []OptionGroup // Grouped options
	Placeholder string        // Placeholder text
	SearchPlaceholder string  // Placeholder for search input
	EmptyText   string        // Text to show when no options match
	Open        bool          // Whether the popover is open
	Disabled    bool          // Whether the combobox is disabled
	Class       string        // Additional CSS classes
	Width       string        // Width of the combobox
	OnSelect    string        // JavaScript to run on selection
}

// OptionGroup represents a group of options
type OptionGroup struct {
	Label   string   // Group label
	Options []Option // Options in this group
}

// Simple creates a simple combobox with string options
func Simple(value string, options []string, placeholder ...string) g.Node {
	ph := "Select option..."
	if len(placeholder) > 0 {
		ph = placeholder[0]
	}

	opts := make([]Option, len(options))
	for i, opt := range options {
		opts[i] = Option{
			Value: opt,
			Label: opt,
		}
	}

	return New(Props{
		Value:       value,
		Options:     opts,
		Placeholder: ph,
	})
}

// WithIcons creates a combobox where each option has an icon
func WithIcons(value string, options []Option, placeholder ...string) g.Node {
	ph := "Select option..."
	if len(placeholder) > 0 {
		ph = placeholder[0]
	}

	return New(Props{
		Value:       value,
		Options:     options,
		Placeholder: ph,
		Width:       "w-[250px]",
	})
}

// Searchable creates a combobox optimized for searching
func Searchable(value string, options []Option, searchPlaceholder ...string) g.Node {
	sp := "Type to search..."
	if len(searchPlaceholder) > 0 {
		sp = searchPlaceholder[0]
	}

	return New(Props{
		Value:             value,
		Options:           options,
		SearchPlaceholder: sp,
		Width:             "w-[300px]",
	})
}