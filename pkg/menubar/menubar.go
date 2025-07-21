package menubar

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// Props defines the properties for the Menubar component
type Props struct {
	Class string // Additional custom classes
}

// MenuProps defines properties for a menu within the menubar
type MenuProps struct {
	Class string
}

// TriggerProps defines properties for a menu trigger
type TriggerProps struct {
	Class    string
	Disabled bool
}

// ContentProps defines properties for menu content
type ContentProps struct {
	Class     string
	Align     string // "start" | "center" | "end"
	Side      string // "top" | "bottom"
	SideOffset int
	AlignOffset int
}

// ItemProps defines properties for menu items
type ItemProps struct {
	Class    string
	Disabled bool
	Inset    bool
}

// SeparatorProps defines properties for menu separators
type SeparatorProps struct {
	Class string
}

// LabelProps defines properties for menu labels
type LabelProps struct {
	Class string
	Inset bool
}

// CheckboxItemProps defines properties for checkbox menu items
type CheckboxItemProps struct {
	Class    string
	Checked  bool
	Disabled bool
	Name     string
	Value    string
}

// RadioGroupProps defines properties for radio group
type RadioGroupProps struct {
	Value string
	Name  string
}

// RadioItemProps defines properties for radio items
type RadioItemProps struct {
	Class    string
	Value    string
	Disabled bool
}

// SubMenuProps defines properties for submenus
type SubMenuProps struct {
	Class string
}

// SubTriggerProps defines properties for submenu triggers
type SubTriggerProps struct {
	Class    string
	Disabled bool
}

// SubContentProps defines properties for submenu content
type SubContentProps struct {
	Class string
}

// ShortcutProps defines properties for keyboard shortcuts
type ShortcutProps struct {
	Class string
}

// New creates a new Menubar component
func New(props Props, children ...g.Node) g.Node {
	classes := lib.CN(
		"flex h-10 items-center space-x-1 rounded-md border bg-background p-1",
		props.Class,
	)

	return html.Div(
		html.Class(classes),
		html.Role("menubar"),
		g.Group(children),
	)
}

// Menu creates a menu within the menubar
func Menu(props MenuProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"relative",
		props.Class,
	)

	return html.Div(
		html.Class(classes),
		g.Group(children),
	)
}

// Trigger creates a menu trigger button
func Trigger(props TriggerProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"flex cursor-default select-none items-center rounded-sm px-3 py-1.5 text-sm font-medium outline-none",
		"focus:bg-accent focus:text-accent-foreground",
		"data-[state=open]:bg-accent data-[state=open]:text-accent-foreground",
		lib.CNIf(props.Disabled,
			"opacity-50 cursor-not-allowed",
			"",
		),
		props.Class,
	)

	attrs := []g.Node{
		html.Type("button"),
		html.Class(classes),
		html.Role("menuitem"),
		g.Attr("aria-haspopup", "menu"),
		g.Attr("data-state", "closed"),
	}

	if props.Disabled {
		attrs = append(attrs, html.Disabled())
	}

	return html.Button(
		append(attrs, children...)...,
	)
}

// Content creates the dropdown content for a menu
func ContentComponent(props ContentProps, children ...g.Node) g.Node {
	// Set defaults
	if props.Align == "" {
		props.Align = "start"
	}
	if props.Side == "" {
		props.Side = "bottom"
	}

	classes := lib.CN(
		"z-50 min-w-[8rem] overflow-hidden rounded-md border bg-popover p-1 text-popover-foreground shadow-md",
		"data-[state=open]:animate-in data-[state=closed]:animate-out",
		"data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0",
		"data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95",
		"data-[side=bottom]:slide-in-from-top-2",
		"data-[side=left]:slide-in-from-right-2",
		"data-[side=right]:slide-in-from-left-2",
		"data-[side=top]:slide-in-from-bottom-2",
		props.Class,
	)

	// Position classes based on alignment and side
	positionClasses := "absolute "
	switch props.Side {
	case "top":
		positionClasses += "bottom-full mb-1 "
	default: // bottom
		positionClasses += "top-full mt-1 "
	}

	switch props.Align {
	case "center":
		positionClasses += "left-1/2 -translate-x-1/2"
	case "end":
		positionClasses += "right-0"
	default: // start
		positionClasses += "left-0"
	}

	return html.Div(
		html.Class(lib.CN(classes, positionClasses)),
		html.Role("menu"),
		g.Attr("aria-orientation", "vertical"),
		g.Attr("data-state", "closed"),
		g.Attr("data-side", props.Side),
		html.Style("display: none;"), // Hidden by default, shown via JavaScript
		g.Group(children),
	)
}

// Item creates a menu item
func Item(props ItemProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"relative flex cursor-default select-none items-center rounded-sm px-2 py-1.5 text-sm outline-none transition-colors",
		"focus:bg-accent focus:text-accent-foreground",
		"data-[disabled]:pointer-events-none data-[disabled]:opacity-50",
		lib.CNIf(props.Inset, "pl-8", ""),
		props.Class,
	)

	attrs := []g.Node{
		html.Class(classes),
		html.Role("menuitem"),
		html.TabIndex("-1"),
	}

	if props.Disabled {
		attrs = append(attrs, g.Attr("data-disabled", "true"))
	}

	return html.Div(
		append(attrs, children...)...,
	)
}

// CheckboxItem creates a checkbox menu item
func CheckboxItem(props CheckboxItemProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"relative flex cursor-default select-none items-center rounded-sm py-1.5 pl-8 pr-2 text-sm outline-none transition-colors",
		"focus:bg-accent focus:text-accent-foreground",
		"data-[disabled]:pointer-events-none data-[disabled]:opacity-50",
		props.Class,
	)

	attrs := []g.Node{
		html.Class(classes),
		html.Role("menuitemcheckbox"),
		g.Attr("aria-checked", lib.CNIf(props.Checked, "true", "false")),
		html.TabIndex("-1"),
	}

	if props.Disabled {
		attrs = append(attrs, g.Attr("data-disabled", "true"))
	}

	checkmark := html.Span(
		html.Class("absolute left-2 flex h-3.5 w-3.5 items-center justify-center"),
		g.If(props.Checked,
			g.Raw(`<svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg">
				<path d="M11.4669 3.72684C11.7558 3.91574 11.8369 4.30308 11.648 4.59198L7.39799 11.092C7.29783 11.2452 7.13556 11.3467 6.95402 11.3699C6.77247 11.3931 6.58989 11.3355 6.45446 11.2124L3.70446 8.71241C3.44905 8.48022 3.43023 8.08494 3.66242 7.82953C3.89461 7.57412 4.28989 7.55529 4.5453 7.78749L6.75292 9.79441L10.6018 3.90792C10.7907 3.61902 11.178 3.53795 11.4669 3.72684Z" fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"></path>
			</svg>`),
		),
	)

	return html.Div(
		append(attrs, checkmark, g.Group(children))...,
	)
}

// RadioGroup creates a radio button group
func RadioGroup(props RadioGroupProps, children ...g.Node) g.Node {
	return html.Div(
		html.Role("group"),
		g.Attr("aria-orientation", "vertical"),
		g.Group(children),
	)
}

// RadioItem creates a radio menu item
func RadioItem(props RadioItemProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"relative flex cursor-default select-none items-center rounded-sm py-1.5 pl-8 pr-2 text-sm outline-none transition-colors",
		"focus:bg-accent focus:text-accent-foreground",
		"data-[disabled]:pointer-events-none data-[disabled]:opacity-50",
		props.Class,
	)

	attrs := []g.Node{
		html.Class(classes),
		html.Role("menuitemradio"),
		html.TabIndex("-1"),
	}

	if props.Disabled {
		attrs = append(attrs, g.Attr("data-disabled", "true"))
	}

	indicator := html.Span(
		html.Class("absolute left-2 flex h-3.5 w-3.5 items-center justify-center"),
		g.Raw(`<svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg">
			<path d="M7.5 9.5C8.60457 9.5 9.5 8.60457 9.5 7.5C9.5 6.39543 8.60457 5.5 7.5 5.5C6.39543 5.5 5.5 6.39543 5.5 7.5C5.5 8.60457 6.39543 9.5 7.5 9.5Z" fill="currentColor"></path>
		</svg>`),
	)

	return html.Div(
		append(attrs, indicator, g.Group(children))...,
	)
}

// Label creates a menu label
func LabelComponent(props LabelProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"px-2 py-1.5 text-sm font-semibold",
		lib.CNIf(props.Inset, "pl-8", ""),
		props.Class,
	)

	return html.Div(
		html.Class(classes),
		g.Group(children),
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
		html.Role("separator"),
	)
}

// SubMenu creates a submenu
func SubMenu(props SubMenuProps, children ...g.Node) g.Node {
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
	classes := lib.CN(
		"flex cursor-default select-none items-center rounded-sm px-2 py-1.5 text-sm outline-none",
		"focus:bg-accent focus:text-accent-foreground",
		"data-[state=open]:bg-accent data-[state=open]:text-accent-foreground",
		lib.CNIf(props.Disabled,
			"opacity-50 cursor-not-allowed",
			"",
		),
		props.Class,
	)

	attrs := []g.Node{
		html.Class(classes),
		html.Role("menuitem"),
		g.Attr("aria-haspopup", "menu"),
		g.Attr("aria-expanded", "false"),
		html.TabIndex("-1"),
	}

	if props.Disabled {
		attrs = append(attrs, g.Attr("data-disabled", "true"))
	}

	// Add chevron icon
	chevron := g.Raw(`<svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg" class="ml-auto h-4 w-4">
		<path d="M6.1584 3.13508C6.35985 2.94621 6.67627 2.95642 6.86514 3.15788L10.6151 7.15788C10.7954 7.3502 10.7954 7.64949 10.6151 7.84182L6.86514 11.8418C6.67627 12.0433 6.35985 12.0535 6.1584 11.8646C5.95694 11.6757 5.94673 11.3593 6.1356 11.1579L9.565 7.49985L6.1356 3.84182C5.94673 3.64036 5.95694 3.32394 6.1584 3.13508Z" fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"></path>
	</svg>`)

	return html.Div(
		append(attrs, g.Group(children), chevron)...,
	)
}

// SubContent creates submenu content
func SubContent(props SubContentProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"z-50 min-w-[8rem] overflow-hidden rounded-md border bg-popover p-1 text-popover-foreground shadow-lg",
		"data-[state=open]:animate-in data-[state=closed]:animate-out",
		"data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0",
		"data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95",
		"data-[side=bottom]:slide-in-from-top-2",
		"data-[side=left]:slide-in-from-right-2",
		"data-[side=right]:slide-in-from-left-2",
		"data-[side=top]:slide-in-from-bottom-2",
		"absolute left-full top-0 ml-1",
		props.Class,
	)

	return html.Div(
		html.Class(classes),
		html.Role("menu"),
		g.Attr("aria-orientation", "vertical"),
		g.Attr("data-state", "closed"),
		html.Style("display: none;"), // Hidden by default
		g.Group(children),
	)
}

// Shortcut creates a keyboard shortcut indicator
func Shortcut(props ShortcutProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"ml-auto text-xs tracking-widest text-muted-foreground",
		props.Class,
	)

	return html.Span(
		html.Class(classes),
		g.Group(children),
	)
}

// Default creates a default menubar
func Default() g.Node {
	return New(Props{})
}