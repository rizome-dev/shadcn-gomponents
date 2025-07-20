package contextmenu

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// Props defines the properties for the ContextMenu component
type Props struct {
	ID    string // Unique ID for the context menu
	Class string // Additional custom classes
}

// TriggerProps defines properties for the context menu trigger
type TriggerProps struct {
	Class string
}

// ContentProps defines properties for the context menu content
type ContentProps struct {
	Class    string
	Position string // Position relative to trigger
	Align    string // Alignment relative to trigger
}

// ItemProps defines properties for context menu items
type ItemProps struct {
	Class    string
	Disabled bool
	Inset    bool
}

// LabelProps defines properties for context menu labels
type LabelProps struct {
	Class string
	Inset bool
}

// SeparatorProps defines properties for context menu separators
type SeparatorProps struct {
	Class string
}

// CheckboxItemProps defines properties for checkbox items
type CheckboxItemProps struct {
	Class    string
	Checked  bool
	Disabled bool
}

// RadioGroupProps defines properties for radio groups
type RadioGroupProps struct {
	Value string
	Class string
}

// RadioItemProps defines properties for radio items
type RadioItemProps struct {
	Value    string
	Class    string
	Disabled bool
}

// SubProps defines properties for submenus
type SubProps struct {
	Open bool
}

// SubTriggerProps defines properties for submenu triggers
type SubTriggerProps struct {
	Class    string
	Disabled bool
	Inset    bool
}

// SubContentProps defines properties for submenu content
type SubContentProps struct {
	Class string
}

// ShortcutProps defines properties for keyboard shortcuts
type ShortcutProps struct {
	Class string
}

// New creates a new ContextMenu container
func New(props Props, children ...g.Node) g.Node {
	return html.Div(
		g.If(props.ID != "", html.ID(props.ID)),
		g.If(props.Class != "", html.Class(props.Class)),
		g.Attr("data-context-menu", "root"),
		g.Group(children),
	)
}

// Trigger creates a context menu trigger area
func Trigger(props TriggerProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"cursor-context-menu",
		props.Class,
	)

	return html.Div(
		html.Class(classes),
		g.Attr("data-context-menu", "trigger"),
		g.Attr("oncontextmenu", "return false;"), // Prevent default context menu
		g.Group(children),
	)
}

// Content creates the context menu content container
func ContentComponent(props ContentProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"z-50 min-w-[8rem] overflow-hidden rounded-md border bg-popover p-1 text-popover-foreground shadow-md",
		"data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2",
		props.Class,
	)

	return html.Div(
		html.Class(classes),
		g.Attr("data-context-menu", "content"),
		g.Attr("data-state", "closed"),
		g.If(props.Position != "", g.Attr("data-side", props.Position)),
		g.If(props.Align != "", g.Attr("data-align", props.Align)),
		html.Style("position: absolute; display: none;"),
		g.Group(children),
	)
}

// Item creates a context menu item
func Item(props ItemProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"relative flex cursor-default select-none items-center rounded-sm px-2 py-1.5 text-sm outline-none transition-colors focus:bg-accent focus:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50",
		lib.CNIf(props.Inset, "pl-8", ""),
		props.Class,
	)

	attrs := []g.Node{
		html.Class(classes),
		g.Attr("data-context-menu", "item"),
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
		"relative flex cursor-default select-none items-center rounded-sm py-1.5 pl-8 pr-2 text-sm outline-none transition-colors focus:bg-accent focus:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50",
		props.Class,
	)

	attrs := []g.Node{
		html.Class(classes),
		g.Attr("data-context-menu", "checkbox-item"),
		g.Attr("role", "menuitemcheckbox"),
		html.TabIndex("-1"),
	}

	if props.Checked {
		attrs = append(attrs, g.Attr("data-state", "checked"), g.Attr("aria-checked", "true"))
	} else {
		attrs = append(attrs, g.Attr("data-state", "unchecked"), g.Attr("aria-checked", "false"))
	}

	if props.Disabled {
		attrs = append(attrs, g.Attr("data-disabled", "true"))
	}

	itemChildren := []g.Node{
		html.Span(
			html.Class("absolute left-2 flex h-3.5 w-3.5 items-center justify-center"),
			g.If(props.Checked,
				g.El("svg",
					g.Attr("viewBox", "0 0 24 24"),
					g.Attr("fill", "none"),
					g.Attr("stroke", "currentColor"),
					g.Attr("stroke-width", "2"),
					g.Attr("stroke-linecap", "round"),
					g.Attr("stroke-linejoin", "round"),
					html.Class("h-4 w-4"),
					g.El("polyline", g.Attr("points", "20 6 9 17 4 12")),
				),
			),
		),
	}

	return html.Div(
		append(attrs, append(itemChildren, children...)...)...,
	)
}

// RadioGroup creates a radio group container
func RadioGroup(props RadioGroupProps, children ...g.Node) g.Node {
	return html.Div(
		g.Attr("role", "group"),
		g.Attr("data-context-menu", "radio-group"),
		g.If(props.Value != "", g.Attr("data-value", props.Value)),
		g.If(props.Class != "", html.Class(props.Class)),
		g.Group(children),
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
		g.Attr("data-context-menu", "radio-item"),
		g.Attr("role", "menuitemradio"),
		g.Attr("data-value", props.Value),
		html.TabIndex("-1"),
	}

	if props.Disabled {
		attrs = append(attrs, g.Attr("data-disabled", "true"))
	}

	itemChildren := []g.Node{
		html.Span(
			html.Class("absolute left-2 flex h-3.5 w-3.5 items-center justify-center"),
			g.El("svg",
				g.Attr("viewBox", "0 0 24 24"),
				g.Attr("fill", "currentColor"),
				html.Class("h-2 w-2 fill-current"),
				html.Style("display: none;"),
				g.Attr("data-state", "unchecked"),
				g.El("circle", g.Attr("cx", "12"), g.Attr("cy", "12"), g.Attr("r", "12")),
			),
		),
	}

	return html.Div(
		append(attrs, append(itemChildren, children...)...)...,
	)
}

// Label creates a context menu label
func LabelComponent(props LabelProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"px-2 py-1.5 text-sm font-semibold text-foreground",
		lib.CNIf(props.Inset, "pl-8", ""),
		props.Class,
	)

	return html.Div(
		html.Class(classes),
		g.Attr("data-context-menu", "label"),
		g.Group(children),
	)
}

// Separator creates a context menu separator
func Separator(props SeparatorProps) g.Node {
	classes := lib.CN(
		"-mx-1 my-1 h-px bg-muted",
		props.Class,
	)

	return html.Div(
		html.Class(classes),
		g.Attr("data-context-menu", "separator"),
		g.Attr("role", "separator"),
	)
}

// Sub creates a submenu container
func Sub(props SubProps, children ...g.Node) g.Node {
	return html.Div(
		g.Attr("data-context-menu", "sub"),
		g.If(props.Open, g.Attr("data-state", "open")),
		g.If(!props.Open, g.Attr("data-state", "closed")),
		g.Group(children),
	)
}

// SubTrigger creates a submenu trigger item
func SubTrigger(props SubTriggerProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"flex cursor-default select-none items-center rounded-sm px-2 py-1.5 text-sm outline-none focus:bg-accent data-[state=open]:bg-accent",
		lib.CNIf(props.Inset, "pl-8", ""),
		props.Class,
	)

	attrs := []g.Node{
		html.Class(classes),
		g.Attr("data-context-menu", "sub-trigger"),
		html.TabIndex("-1"),
	}

	if props.Disabled {
		attrs = append(attrs, g.Attr("data-disabled", "true"))
	}

	// Add chevron icon
	childrenWithIcon := append(children,
		g.El("svg",
			g.Attr("viewBox", "0 0 24 24"),
			g.Attr("fill", "none"),
			g.Attr("stroke", "currentColor"),
			g.Attr("stroke-width", "2"),
			g.Attr("stroke-linecap", "round"),
			g.Attr("stroke-linejoin", "round"),
			html.Class("ml-auto h-4 w-4"),
			g.El("polyline", g.Attr("points", "9 18 15 12 9 6")),
		),
	)

	return html.Div(
		append(attrs, childrenWithIcon...)...,
	)
}

// SubContent creates submenu content
func SubContent(props SubContentProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"z-50 min-w-[8rem] overflow-hidden rounded-md border bg-popover p-1 text-popover-foreground shadow-lg",
		"data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2",
		props.Class,
	)

	return html.Div(
		html.Class(classes),
		g.Attr("data-context-menu", "sub-content"),
		html.Style("position: absolute; display: none;"),
		g.Group(children),
	)
}

// Shortcut creates a keyboard shortcut display
func Shortcut(props ShortcutProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"ml-auto text-xs tracking-widest text-muted-foreground",
		props.Class,
	)

	return html.Span(
		html.Class(classes),
		g.Attr("data-context-menu", "shortcut"),
		g.Group(children),
	)
}