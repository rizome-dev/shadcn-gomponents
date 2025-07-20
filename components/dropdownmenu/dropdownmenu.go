package dropdownmenu

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
	"github.com/rizome-dev/shadcn-gomponents/lib/icons"
)

// Props defines the properties for the DropdownMenu component
type Props struct {
	Open  bool   // Whether the dropdown is open (for server-side rendering)
	Class string // Additional custom classes
}

// TriggerProps defines the properties for the Trigger
type TriggerProps struct {
	Class    string
	AsChild  bool // Whether to render as child element
	Disabled bool
}

// ContentProps defines the properties for the Content
type ContentProps struct {
	Class      string
	SideOffset int    // Offset from the trigger
	Align      string // Alignment: "start", "center", "end"
	Side       string // Side: "top", "right", "bottom", "left"
}

// ItemProps defines the properties for menu items
type ItemProps struct {
	Class    string
	Inset    bool // Whether to add inset padding
	Disabled bool
}

// CheckboxItemProps defines the properties for checkbox items
type CheckboxItemProps struct {
	Class    string
	Checked  bool
	Disabled bool
}

// RadioItemProps defines the properties for radio items
type RadioItemProps struct {
	Class    string
	Value    string
	Disabled bool
}

// LabelProps defines the properties for labels
type LabelProps struct {
	Class string
	Inset bool
}

// SeparatorProps defines the properties for separators
type SeparatorProps struct {
	Class string
}

// ShortcutProps defines the properties for shortcuts
type ShortcutProps struct {
	Class string
}

// GroupProps defines the properties for groups
type GroupProps struct {
	Class string
}

// SubProps defines the properties for submenus
type SubProps struct {
	Open  bool
	Class string
}

// SubTriggerProps defines the properties for submenu triggers
type SubTriggerProps struct {
	Class    string
	Inset    bool
	Disabled bool
}

// SubContentProps defines the properties for submenu content
type SubContentProps struct {
	Class string
}

// RadioGroupProps defines the properties for radio groups
type RadioGroupProps struct {
	Class string
	Value string // Selected value
}

// New creates a new DropdownMenu component
// Note: This is a static implementation for server-side rendering
// For full interactivity, JavaScript would be needed
func New(props Props, children ...g.Node) g.Node {
	classes := lib.CN(
		"relative inline-block text-left",
		props.Class,
	)

	return html.Div(
		append([]g.Node{html.Class(classes)}, children...)...,
	)
}

// Trigger creates the dropdown trigger element
func Trigger(props TriggerProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"inline-flex items-center justify-center",
		props.Class,
	)

	attrs := []g.Node{
		html.Type("button"),
		g.If(classes != "", html.Class(classes)),
		g.If(props.Disabled, html.Disabled()),
		g.Attr("aria-haspopup", "menu"),
		g.Attr("aria-expanded", "false"),
	}

	if props.AsChild && len(children) > 0 {
		// If AsChild is true, we assume the child is already a button-like element
		return g.Group(children)
	}

	return html.Button(
		append(attrs, g.Group(children))...,
	)
}

// DropdownContent creates the dropdown content container
func DropdownContent(props ContentProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"z-50 min-w-[8rem] overflow-hidden rounded-md border bg-popover p-1 text-popover-foreground shadow-md",
		props.Class,
	)

	// For server-side rendering, we'll just render the content as visible
	// In a real implementation, this would be controlled by JavaScript
	return html.Div(
		html.Class(classes),
		g.Attr("role", "menu"),
		g.Attr("aria-orientation", "vertical"),
		g.Group(children),
	)
}

// Item creates a menu item
func Item(props ItemProps, children ...g.Node) g.Node {
	baseClasses := "relative flex cursor-default select-none items-center gap-2 rounded-sm px-2 py-1.5 text-sm outline-none transition-colors focus:bg-accent focus:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50"
	if props.Inset {
		baseClasses += " pl-8"
	}
	classes := lib.CN(baseClasses, props.Class)

	attrs := []g.Node{
		html.Class(classes),
		g.Attr("role", "menuitem"),
		g.Attr("tabindex", "-1"),
	}

	if props.Disabled {
		attrs = append(attrs, g.Attr("aria-disabled", "true"))
		attrs = append(attrs, g.Attr("data-disabled", ""))
	}

	return html.Div(
		append(attrs, g.Group(children))...,
	)
}

// CheckboxItem creates a checkbox menu item
func CheckboxItem(props CheckboxItemProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"relative flex cursor-default select-none items-center rounded-sm py-1.5 pl-8 pr-2 text-sm outline-none transition-colors focus:bg-accent focus:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50",
		props.Class,
	)

	checkedValue := "false"
	if props.Checked {
		checkedValue = "true"
	}
	
	attrs := []g.Node{
		html.Class(classes),
		g.Attr("role", "menuitemcheckbox"),
		g.Attr("aria-checked", checkedValue),
		g.Attr("tabindex", "-1"),
	}

	if props.Disabled {
		attrs = append(attrs, g.Attr("aria-disabled", "true"))
		attrs = append(attrs, g.Attr("data-disabled", ""))
	}

	checkmark := html.Span(
		html.Class("absolute left-2 flex h-3.5 w-3.5 items-center justify-center"),
		g.If(props.Checked,
			icons.Check(g.Attr("class", "h-4 w-4")),
		),
	)

	return html.Div(
		append(attrs, checkmark, g.Group(children))...,
	)
}

// RadioItem creates a radio menu item
func RadioItem(props RadioItemProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"relative flex cursor-default select-none items-center rounded-sm py-1.5 pl-8 pr-2 text-sm outline-none transition-colors focus:bg-accent focus:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50",
		props.Class,
	)

	attrs := []g.Node{
		html.Class(classes),
		g.Attr("role", "menuitemradio"),
		g.Attr("tabindex", "-1"),
		g.Attr("data-value", props.Value),
	}

	if props.Disabled {
		attrs = append(attrs, g.Attr("aria-disabled", "true"))
		attrs = append(attrs, g.Attr("data-disabled", ""))
	}

	// Radio indicator (will be controlled by parent RadioGroup)
	indicator := html.Span(
		html.Class("absolute left-2 flex h-3.5 w-3.5 items-center justify-center"),
		// The circle will be shown based on selection state
	)

	return html.Div(
		append(attrs, indicator, g.Group(children))...,
	)
}

// DropdownLabel creates a menu label
func DropdownLabel(props LabelProps, text string) g.Node {
	baseClasses := "px-2 py-1.5 text-sm font-semibold"
	if props.Inset {
		baseClasses += " pl-8"
	}
	classes := lib.CN(baseClasses, props.Class)

	return html.Div(
		html.Class(classes),
		g.Text(text),
	)
}

// Separator creates a menu separator
func Separator(props SeparatorProps) g.Node {
	classes := lib.CN(
		"-mx-1 my-1 h-px bg-muted",
		props.Class,
	)

	return html.Div(
		html.Class(classes),
		g.Attr("role", "separator"),
	)
}

// Shortcut creates a keyboard shortcut indicator
func Shortcut(props ShortcutProps, text string) g.Node {
	classes := lib.CN(
		"ml-auto text-xs tracking-widest opacity-60",
		props.Class,
	)

	return html.Span(
		html.Class(classes),
		g.Text(text),
	)
}

// Group creates a menu group
func Group(props GroupProps, children ...g.Node) g.Node {
	classes := lib.CN(props.Class)

	return html.Div(
		g.Attr("role", "group"),
		g.If(classes != "", html.Class(classes)),
		g.Group(children),
	)
}

// DropdownSub creates a submenu container
func DropdownSub(props SubProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"relative",
		props.Class,
	)

	return html.Div(
		html.Class(classes),
		g.Group(children),
	)
}

// SubTrigger creates a submenu trigger
func SubTrigger(props SubTriggerProps, children ...g.Node) g.Node {
	baseClasses := "flex cursor-default select-none items-center gap-2 rounded-sm px-2 py-1.5 text-sm outline-none focus:bg-accent data-[state=open]:bg-accent"
	if props.Inset {
		baseClasses += " pl-8"
	}
	classes := lib.CN(baseClasses, props.Class)

	attrs := []g.Node{
		html.Class(classes),
		g.Attr("role", "menuitem"),
		g.Attr("aria-haspopup", "menu"),
		g.Attr("aria-expanded", "false"),
		g.Attr("tabindex", "-1"),
	}

	if props.Disabled {
		attrs = append(attrs, g.Attr("aria-disabled", "true"))
		attrs = append(attrs, g.Attr("data-disabled", ""))
	}

	// Add chevron icon
	chevron := icons.ChevronRight(g.Attr("class", "ml-auto h-4 w-4"))

	return html.Div(
		append(attrs, g.Group(children), chevron)...,
	)
}

// SubContent creates submenu content
func SubContent(props SubContentProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"z-50 min-w-[8rem] overflow-hidden rounded-md border bg-popover p-1 text-popover-foreground shadow-lg",
		props.Class,
	)

	return html.Div(
		html.Class(classes),
		g.Attr("role", "menu"),
		g.Attr("aria-orientation", "vertical"),
		g.Group(children),
	)
}

// RadioGroup creates a radio group container
func RadioGroup(props RadioGroupProps, children ...g.Node) g.Node {
	classes := lib.CN(props.Class)

	// For server-side rendering, we need to mark the selected item
	// This would normally be handled by JavaScript
	processedChildren := make([]g.Node, 0, len(children))
	for _, child := range children {
		// In a real implementation, we'd check if this is a RadioItem
		// and set its selected state based on props.Value
		processedChildren = append(processedChildren, child)
	}

	return html.Div(
		g.Attr("role", "radiogroup"),
		g.If(classes != "", html.Class(classes)),
		g.Group(processedChildren),
	)
}

// RadioItemWithSelection creates a radio item with selection state
func RadioItemWithSelection(props RadioItemProps, selected bool, children ...g.Node) g.Node {
	classes := lib.CN(
		"relative flex cursor-default select-none items-center rounded-sm py-1.5 pl-8 pr-2 text-sm outline-none transition-colors focus:bg-accent focus:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50",
		props.Class,
	)

	checkedValue := "false"
	if selected {
		checkedValue = "true"
	}
	
	attrs := []g.Node{
		html.Class(classes),
		g.Attr("role", "menuitemradio"),
		g.Attr("aria-checked", checkedValue),
		g.Attr("tabindex", "-1"),
		g.Attr("data-value", props.Value),
	}

	if props.Disabled {
		attrs = append(attrs, g.Attr("aria-disabled", "true"))
		attrs = append(attrs, g.Attr("data-disabled", ""))
	}

	// Radio indicator
	indicator := html.Span(
		html.Class("absolute left-2 flex h-3.5 w-3.5 items-center justify-center"),
		g.If(selected,
			icons.CircleIcon(g.Attr("class", "h-2 w-2 fill-current")),
		),
	)

	return html.Div(
		append(attrs, indicator, g.Group(children))...,
	)
}

// Example creates a basic dropdown menu example
func Example() g.Node {
	return New(
		Props{Open: true},
		Trigger(
			TriggerProps{},
			g.Text("Open"),
		),
		DropdownContent(
			ContentProps{Class: "w-56"},
			DropdownLabel(LabelProps{}, "My Account"),
			Separator(SeparatorProps{}),
			Group(
				GroupProps{},
				Item(
					ItemProps{},
					icons.User(g.Attr("class", "h-4 w-4")),
					html.Span(g.Text("Profile")),
					Shortcut(ShortcutProps{}, "⇧⌘P"),
				),
				Item(
					ItemProps{},
					icons.CreditCard(g.Attr("class", "h-4 w-4")),
					html.Span(g.Text("Billing")),
					Shortcut(ShortcutProps{}, "⌘B"),
				),
				Item(
					ItemProps{},
					icons.Settings(g.Attr("class", "h-4 w-4")),
					html.Span(g.Text("Settings")),
					Shortcut(ShortcutProps{}, "⌘S"),
				),
			),
			Separator(SeparatorProps{}),
			Item(
				ItemProps{Disabled: true},
				icons.Cloud(g.Attr("class", "h-4 w-4")),
				html.Span(g.Text("API")),
			),
			Separator(SeparatorProps{}),
			Item(
				ItemProps{},
				icons.LogOut(g.Attr("class", "h-4 w-4")),
				html.Span(g.Text("Log out")),
				Shortcut(ShortcutProps{}, "⇧⌘Q"),
			),
		),
	)
}

// ExampleCheckboxes creates a dropdown with checkbox items
func ExampleCheckboxes() g.Node {
	return New(
		Props{Open: true},
		Trigger(
			TriggerProps{},
			g.Text("Open"),
		),
		DropdownContent(
			ContentProps{Class: "w-56"},
			DropdownLabel(LabelProps{}, "Appearance"),
			Separator(SeparatorProps{}),
			CheckboxItem(
				CheckboxItemProps{Checked: true},
				g.Text("Status Bar"),
			),
			CheckboxItem(
				CheckboxItemProps{Checked: false, Disabled: true},
				g.Text("Activity Bar"),
			),
			CheckboxItem(
				CheckboxItemProps{Checked: false},
				g.Text("Panel"),
			),
		),
	)
}

// ExampleRadioGroup creates a dropdown with radio items
func ExampleRadioGroup() g.Node {
	return New(
		Props{Open: true},
		Trigger(
			TriggerProps{},
			g.Text("Open"),
		),
		DropdownContent(
			ContentProps{Class: "w-56"},
			DropdownLabel(LabelProps{}, "Panel Position"),
			Separator(SeparatorProps{}),
			RadioGroup(
				RadioGroupProps{Value: "bottom"},
				RadioItemWithSelection(
					RadioItemProps{Value: "top"},
					false,
					g.Text("Top"),
				),
				RadioItemWithSelection(
					RadioItemProps{Value: "bottom"},
					true,
					g.Text("Bottom"),
				),
				RadioItemWithSelection(
					RadioItemProps{Value: "right"},
					false,
					g.Text("Right"),
				),
			),
		),
	)
}

// ExampleWithSubmenu creates a dropdown with a submenu
func ExampleWithSubmenu() g.Node {
	return New(
		Props{Open: true},
		Trigger(
			TriggerProps{},
			g.Text("Open"),
		),
		DropdownContent(
			ContentProps{Class: "w-56"},
			DropdownLabel(LabelProps{}, "Team"),
			Separator(SeparatorProps{}),
			Item(
				ItemProps{},
				icons.Users(g.Attr("class", "h-4 w-4")),
				html.Span(g.Text("Team")),
			),
			DropdownSub(
				SubProps{},
				SubTrigger(
					SubTriggerProps{},
					icons.UserPlus(g.Attr("class", "h-4 w-4")),
					html.Span(g.Text("Invite users")),
				),
				// Note: In a static implementation, submenus are challenging
				// This would typically be shown on hover/click with JavaScript
			),
			Item(
				ItemProps{},
				icons.Plus(g.Attr("class", "h-4 w-4")),
				html.Span(g.Text("New Team")),
				Shortcut(ShortcutProps{}, "⌘+T"),
			),
		),
	)
}