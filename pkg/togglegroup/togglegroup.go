package togglegroup

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/pkg/toggle"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// SelectionType defines the selection type for the toggle group
type SelectionType string

const (
	TypeSingle   SelectionType = "single"   // Only one item can be selected
	TypeMultiple SelectionType = "multiple" // Multiple items can be selected
)

// Props defines the properties for the ToggleGroup component
type Props struct {
	ID       string        // HTML id attribute
	Type     SelectionType // "single" | "multiple"
	Value    []string      // Currently selected values
	Variant  string        // "default" | "outline"
	Size     string        // "sm" | "default" | "lg"
	Disabled bool          // Whether the entire group is disabled
	Class    string        // Additional custom classes
}

// ItemProps defines the properties for a ToggleGroupItem
type ItemProps struct {
	Value       string // The value this item represents
	Pressed     bool   // Whether this item is pressed/selected
	Disabled    bool   // Whether this specific item is disabled
	AriaLabel   string // Accessibility label
	Class       string // Additional custom classes
	OnClick     string // JavaScript onClick handler
}

// New creates a new ToggleGroup component
func New(props Props, children ...g.Node) g.Node {
	// Set defaults
	if props.Type == "" {
		props.Type = TypeSingle
	}
	if props.Variant == "" {
		props.Variant = "default"
	}
	if props.Size == "" {
		props.Size = "default"
	}

	// Build attributes
	attrs := []g.Node{
		g.Attr("role", "group"),
		g.Attr("data-toggle-group", "true"),
		g.Attr("data-type", string(props.Type)),
		g.Attr("data-variant", props.Variant),
		g.Attr("data-size", props.Size),
		html.Class(lib.CN(
			"group/toggle-group flex w-fit items-center rounded-md",
			lib.CNIf(props.Variant == "outline", "shadow-sm", ""),
			props.Class,
		)),
	}

	if props.ID != "" {
		attrs = append(attrs, g.Attr("id", props.ID))
	}

	if props.Disabled {
		attrs = append(attrs, g.Attr("data-disabled", "true"))
	}

	// Store current values as data attribute for JavaScript access
	if len(props.Value) > 0 {
		valueStr := ""
		for i, v := range props.Value {
			if i > 0 {
				valueStr += ","
			}
			valueStr += v
		}
		attrs = append(attrs, g.Attr("data-value", valueStr))
	}

	return html.Div(
		append(attrs, children...)...,
	)
}

// Item creates a toggle group item
func Item(props ItemProps, groupProps Props, children ...g.Node) g.Node {
	// Determine if this item is pressed based on group value
	pressed := props.Pressed
	if !pressed && len(groupProps.Value) > 0 {
		for _, v := range groupProps.Value {
			if v == props.Value {
				pressed = true
				break
			}
		}
	}

	// Create toggle with group-specific styling
	toggleProps := toggle.Props{
		Pressed:   pressed,
		Disabled:  props.Disabled || groupProps.Disabled,
		AriaLabel: props.AriaLabel,
		Variant:   groupProps.Variant,
		Size:      groupProps.Size,
		Class: lib.CN(
			"min-w-0 flex-1 shrink-0 rounded-none shadow-none",
			"first:rounded-l-md last:rounded-r-md",
			"focus:z-10 focus-visible:z-10",
			lib.CNIf(groupProps.Variant == "outline",
				"border-l-0 first:border-l",
				"",
			),
			props.Class,
		),
		OnClick: props.OnClick,
		Attrs: []g.Node{
			g.Attr("data-value", props.Value),
		},
	}

	// Create and return the toggle
	return toggle.New(toggleProps, children...)
}

// Single creates a single-selection toggle group
func Single(props Props, children ...g.Node) g.Node {
	props.Type = TypeSingle
	return New(props, children...)
}

// MultipleSelection creates a multiple-selection toggle group
func MultipleSelection(props Props, children ...g.Node) g.Node {
	props.Type = TypeMultiple
	return New(props, children...)
}

// Helper functions for common configurations

// DefaultGroup creates a default variant toggle group
func DefaultGroup(children ...g.Node) g.Node {
	return New(Props{
		Variant: "default",
	}, children...)
}

// OutlineGroup creates an outline variant toggle group
func OutlineGroup(children ...g.Node) g.Node {
	return New(Props{
		Variant: "outline",
	}, children...)
}

// SmallGroup creates a small size toggle group
func SmallGroup(props Props, children ...g.Node) g.Node {
	props.Size = "sm"
	return New(props, children...)
}

// LargeGroup creates a large size toggle group
func LargeGroup(props Props, children ...g.Node) g.Node {
	props.Size = "lg"
	return New(props, children...)
}

// WithItems creates a toggle group with multiple items from a slice
func WithItems(props Props, items []struct {
	Value     string
	Label     string
	Icon      g.Node
	Disabled  bool
	AriaLabel string
}) g.Node {
	itemNodes := make([]g.Node, len(items))
	for i, item := range items {
		content := []g.Node{}
		if item.Icon != nil {
			content = append(content, item.Icon)
		}
		if item.Label != "" {
			content = append(content, g.Text(item.Label))
		}

		itemNodes[i] = Item(ItemProps{
			Value:     item.Value,
			Disabled:  item.Disabled,
			AriaLabel: item.AriaLabel,
		}, props, g.Group(content))
	}

	return New(props, g.Group(itemNodes))
}

// TextFormatting creates a text formatting toggle group
func TextFormatting() g.Node {
	return Single(Props{
		Variant: "outline",
		Value:   []string{"bold"},
	},
		Item(ItemProps{Value: "bold", AriaLabel: "Toggle bold"}, Props{Variant: "outline"},
			g.Raw(`<svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg" class="h-4 w-4"><path d="M5.10505 12C4.70805 12 4.4236 11.912 4.25171 11.736C4.0839 11.5559 4 11.2715 4 10.8827V4.11733C4 3.72033 4.08595 3.43588 4.25784 3.26398C4.43383 3.08799 4.71623 3 5.10505 3H8.2C8.94904 3 9.53636 3.21164 9.96196 3.63491C10.3917 4.05819 10.6065 4.63592 10.6065 5.36811C10.6065 5.92517 10.4763 6.39344 10.2159 6.77292C9.95958 7.14831 9.59099 7.42264 9.11014 7.59592C9.662 7.72607 10.1014 7.99438 10.4283 8.40085C10.7593 8.80322 10.9248 9.32428 10.9248 9.96401C10.9248 10.7383 10.6679 11.3679 10.1541 11.8527C9.64045 12.3375 8.96325 12.58 8.12222 12.58H5.10505V12ZM6.16134 6.91681H7.825C8.17632 6.91681 8.45268 6.81371 8.65409 6.6075C8.85958 6.3972 8.96233 6.11786 8.96233 5.76948C8.96233 5.4211 8.85958 5.14176 8.65409 4.93145C8.45268 4.72115 8.17632 4.616 7.825 4.616H6.16134V6.91681ZM6.16134 10.9641H8.0318C8.42062 10.9641 8.73039 10.8528 8.96111 10.6302C9.19591 10.4035 9.31331 10.1014 9.31331 9.72409C9.31331 9.34679 9.19591 9.04469 8.96111 8.81779C8.73039 8.59089 8.42062 8.47744 8.0318 8.47744H6.16134V10.9641Z" fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"></path></svg>`),
		),
		Item(ItemProps{Value: "italic", AriaLabel: "Toggle italic"}, Props{Variant: "outline"},
			g.Raw(`<svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg" class="h-4 w-4"><path d="M5.67494 3.50017C5.67494 3.25164 5.87641 3.05017 6.12494 3.05017H10.6249C10.8735 3.05017 11.0749 3.25164 11.0749 3.50017C11.0749 3.7487 10.8735 3.95017 10.6249 3.95017H9.00587L7.2309 11.05H8.87493C9.12345 11.05 9.32493 11.2515 9.32493 11.5C9.32493 11.7486 9.12345 11.95 8.87493 11.95H4.37493C4.1264 11.95 3.92493 11.7486 3.92493 11.5C3.92493 11.2515 4.1264 11.05 4.37493 11.05H5.99397L7.76894 3.95017H6.12494C5.87641 3.95017 5.67494 3.7487 5.67494 3.50017Z" fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"></path></svg>`),
		),
		Item(ItemProps{Value: "underline", AriaLabel: "Toggle underline"}, Props{Variant: "outline"},
			g.Raw(`<svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg" class="h-4 w-4"><path d="M5.00001 2.75C5.00001 2.47386 4.77615 2.25 4.50001 2.25C4.22387 2.25 4.00001 2.47386 4.00001 2.75V8.05C4.00001 9.983 5.56702 11.55 7.50001 11.55C9.43301 11.55 11 9.983 11 8.05V2.75C11 2.47386 10.7762 2.25 10.5 2.25C10.2239 2.25 10 2.47386 10 2.75V8.05C10 9.43071 8.88072 10.55 7.50001 10.55C6.1193 10.55 5.00001 9.43071 5.00001 8.05V2.75ZM3.49998 13.1001C3.27906 13.1001 3.09998 13.2791 3.09998 13.5001C3.09998 13.721 3.27906 13.9001 3.49998 13.9001H11.5C11.7209 13.9001 11.9 13.721 11.9 13.5001C11.9 13.2791 11.7209 13.1001 11.5 13.1001H3.49998Z" fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"></path></svg>`),
		),
	)
}

// Alignment creates an alignment toggle group
func Alignment() g.Node {
	return Single(Props{
		Variant: "outline",
		Value:   []string{"center"},
	},
		Item(ItemProps{Value: "left", AriaLabel: "Align left"}, Props{Variant: "outline"},
			g.Raw(`<svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg" class="h-4 w-4"><path d="M2 4.5C2 4.22386 2.22386 4 2.5 4H10.5C10.7761 4 11 4.22386 11 4.5C11 4.77614 10.7761 5 10.5 5H2.5C2.22386 5 2 4.77614 2 4.5ZM2 7.5C2 7.22386 2.22386 7 2.5 7H7.5C7.77614 7 8 7.22386 8 7.5C8 7.77614 7.77614 8 7.5 8H2.5C2.22386 8 2 7.77614 2 7.5ZM2 10.5C2 10.2239 2.22386 10 2.5 10H10.5C10.7761 10 11 10.2239 11 10.5C11 10.7761 10.7761 11 10.5 11H2.5C2.22386 11 2 10.7761 2 10.5Z" fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"></path></svg>`),
		),
		Item(ItemProps{Value: "center", AriaLabel: "Align center"}, Props{Variant: "outline"},
			g.Raw(`<svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg" class="h-4 w-4"><path d="M2 4.5C2 4.22386 2.22386 4 2.5 4H12.5C12.7761 4 13 4.22386 13 4.5C13 4.77614 12.7761 5 12.5 5H2.5C2.22386 5 2 4.77614 2 4.5ZM4 7.5C4 7.22386 4.22386 7 4.5 7H10.5C10.7761 7 11 7.22386 11 7.5C11 7.77614 10.7761 8 10.5 8H4.5C4.22386 8 4 7.77614 4 7.5ZM2 10.5C2 10.2239 2.22386 10 2.5 10H12.5C12.7761 10 13 10.2239 13 10.5C13 10.7761 12.7761 11 12.5 11H2.5C2.22386 11 2 10.7761 2 10.5Z" fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"></path></svg>`),
		),
		Item(ItemProps{Value: "right", AriaLabel: "Align right"}, Props{Variant: "outline"},
			g.Raw(`<svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg" class="h-4 w-4"><path d="M4 4.5C4 4.22386 4.22386 4 4.5 4H12.5C12.7761 4 13 4.22386 13 4.5C13 4.77614 12.7761 5 12.5 5H4.5C4.22386 5 4 4.77614 4 4.5ZM7 7.5C7 7.22386 7.22386 7 7.5 7H12.5C12.7761 7 13 7.22386 13 7.5C13 7.77614 12.7761 8 12.5 8H7.5C7.22386 8 7 7.77614 7 7.5ZM4 10.5C4 10.2239 4.22386 10 4.5 10H12.5C12.7761 10 13 10.2239 13 10.5C13 10.7761 12.7761 11 12.5 11H4.5C4.22386 11 4 10.7761 4 10.5Z" fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"></path></svg>`),
		),
		Item(ItemProps{Value: "justify", AriaLabel: "Justify"}, Props{Variant: "outline"},
			g.Raw(`<svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg" class="h-4 w-4"><path d="M2 4.5C2 4.22386 2.22386 4 2.5 4H12.5C12.7761 4 13 4.22386 13 4.5C13 4.77614 12.7761 5 12.5 5H2.5C2.22386 5 2 4.77614 2 4.5ZM2 7.5C2 7.22386 2.22386 7 2.5 7H12.5C12.7761 7 13 7.22386 13 7.5C13 7.77614 12.7761 8 12.5 8H2.5C2.22386 8 2 7.77614 2 7.5ZM2 10.5C2 10.2239 2.22386 10 2.5 10H12.5C12.7761 10 13 10.2239 13 10.5C13 10.7761 12.7761 11 12.5 11H2.5C2.22386 11 2 10.7761 2 10.5Z" fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"></path></svg>`),
		),
	)
}