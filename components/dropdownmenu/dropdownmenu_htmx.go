package dropdownmenu

import (
	"fmt"
	"net/http"
	"strings"
	
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	hx "maragu.dev/gomponents-htmx"
	"github.com/rizome-dev/shadcn-gomponents/lib"
	"github.com/rizome-dev/shadcn-gomponents/lib/icons"
)

// HTMXProps defines HTMX-specific properties for the DropdownMenu
type HTMXProps struct {
	ID         string // Unique ID for the dropdown
	TogglePath string // Server path for toggle actions
	ItemPath   string // Server path for item actions
}

// NewHTMX creates an HTMX-enhanced DropdownMenu component
func NewHTMX(props Props, htmxProps HTMXProps, trigger g.Node, content g.Node) g.Node {
	classes := lib.CN(
		"relative inline-block text-left",
		props.Class,
	)

	return html.Div(
		html.ID(htmxProps.ID),
		html.Class(classes),
		trigger,
		g.If(props.Open, content),
	)
}

// TriggerHTMX creates an HTMX-enhanced trigger
func TriggerHTMX(props TriggerProps, htmxProps HTMXProps, isOpen bool, children ...g.Node) g.Node {
	classes := lib.CN(
		"inline-flex items-center justify-center",
		props.Class,
	)

	attrs := []g.Node{
		html.Type("button"),
		g.If(classes != "", html.Class(classes)),
		g.If(props.Disabled, html.Disabled()),
		g.Attr("aria-haspopup", "menu"),
		g.Attr("aria-expanded", fmt.Sprintf("%t", isOpen)),
		hx.Post(htmxProps.TogglePath),
		html.Target("#" + htmxProps.ID),
		hx.Swap("outerHTML"),
		hx.Vals(fmt.Sprintf(`{"open": "%t"}`, !isOpen)),
	}

	if props.AsChild && len(children) > 0 {
		// Assume the child element can handle HTMX attributes
		return g.Group(children)
	}

	return html.Button(
		append(attrs, g.Group(children))...,
	)
}

// DropdownContentHTMX creates HTMX-enhanced dropdown content
func DropdownContentHTMX(props ContentProps, htmxProps HTMXProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"absolute z-50 min-w-[8rem] overflow-hidden rounded-md border bg-popover p-1 text-popover-foreground shadow-md",
		"animate-in fade-in-0 zoom-in-95",
		props.Class,
	)

	// Determine position
	position := "top-full mt-2"
	if props.Side == "top" {
		position = "bottom-full mb-2"
	} else if props.Side == "left" {
		position = "right-full mr-2 top-0"
	} else if props.Side == "right" {
		position = "left-full ml-2 top-0"
	}

	// Determine alignment
	alignment := ""
	if props.Align == "start" {
		alignment = "left-0"
	} else if props.Align == "end" {
		alignment = "right-0"
	} else {
		alignment = "left-1/2 -translate-x-1/2"
	}

	return html.Div(
		html.Class(lib.CN(classes, position, alignment)),
		g.Attr("role", "menu"),
		g.Attr("aria-orientation", "vertical"),
		// Close on click outside
		hx.On("click", "event.stopPropagation()"),
		g.Group(children),
	)
}

// ItemHTMX creates an HTMX-enhanced menu item
func ItemHTMX(props ItemProps, htmxProps HTMXProps, action string, children ...g.Node) g.Node {
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

	if !props.Disabled && action != "" {
		attrs = append(attrs,
			hx.Post(action),
			html.Target("#" + htmxProps.ID),
			hx.Swap("outerHTML"),
		)
	}

	if props.Disabled {
		attrs = append(attrs, g.Attr("aria-disabled", "true"))
		attrs = append(attrs, g.Attr("data-disabled", ""))
	}

	return html.Div(
		append(attrs, g.Group(children))...,
	)
}

// CheckboxItemHTMX creates an HTMX-enhanced checkbox menu item
func CheckboxItemHTMX(props CheckboxItemProps, htmxProps HTMXProps, name string, children ...g.Node) g.Node {
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

	if !props.Disabled {
		attrs = append(attrs,
			hx.Post(htmxProps.ItemPath + "/checkbox"),
			html.Target("#" + htmxProps.ID),
			hx.Swap("outerHTML"),
			hx.Vals(fmt.Sprintf(`{"name": "%s", "checked": "%t"}`, name, !props.Checked)),
		)
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

// RadioItemHTMX creates an HTMX-enhanced radio menu item
func RadioItemHTMX(props RadioItemProps, htmxProps HTMXProps, groupName string, selected bool, children ...g.Node) g.Node {
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

	if !props.Disabled {
		attrs = append(attrs,
			hx.Post(htmxProps.ItemPath + "/radio"),
			html.Target("#" + htmxProps.ID),
			hx.Swap("outerHTML"),
			hx.Vals(`{"group": "` + groupName + `", "value": "` + props.Value + `"}`),
		)
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

// ExampleHTMX creates an HTMX-enhanced dropdown menu example
func ExampleHTMX() g.Node {
	htmxProps := HTMXProps{
		ID:         "dropdown-example",
		TogglePath: "/api/dropdown/toggle",
		ItemPath:   "/api/dropdown/item",
	}

	return RenderDropdownMenu(htmxProps, false, nil)
}

// DropdownMenuState represents the state of a dropdown menu
type DropdownMenuState struct {
	Open       bool              `json:"open"`
	Checkboxes map[string]bool   `json:"checkboxes"`
	RadioGroups map[string]string `json:"radioGroups"`
}

// RenderDropdownMenu renders a dropdown menu with current state
func RenderDropdownMenu(htmxProps HTMXProps, isOpen bool, state *DropdownMenuState) g.Node {
	if state == nil {
		state = &DropdownMenuState{
			Open:        isOpen,
			Checkboxes:  make(map[string]bool),
			RadioGroups: make(map[string]string),
		}
		// Default values
		state.Checkboxes["statusBar"] = true
		state.Checkboxes["activityBar"] = false
		state.Checkboxes["panel"] = false
		state.RadioGroups["position"] = "bottom"
	}

	return NewHTMX(
		Props{Open: isOpen},
		htmxProps,
		TriggerHTMX(
			TriggerProps{},
			htmxProps,
			isOpen,
			g.Text("Open"),
		),
		DropdownContentHTMX(
			ContentProps{Class: "w-56"},
			htmxProps,
			DropdownLabel(LabelProps{}, "My Account"),
			Separator(SeparatorProps{}),
			Group(
				GroupProps{},
				ItemHTMX(
					ItemProps{},
					htmxProps,
					htmxProps.ItemPath + "/profile",
					icons.User(g.Attr("class", "h-4 w-4")),
					html.Span(g.Text("Profile")),
					Shortcut(ShortcutProps{}, "⇧⌘P"),
				),
				ItemHTMX(
					ItemProps{},
					htmxProps,
					htmxProps.ItemPath + "/billing",
					icons.CreditCard(g.Attr("class", "h-4 w-4")),
					html.Span(g.Text("Billing")),
					Shortcut(ShortcutProps{}, "⌘B"),
				),
				ItemHTMX(
					ItemProps{},
					htmxProps,
					htmxProps.ItemPath + "/settings",
					icons.Settings(g.Attr("class", "h-4 w-4")),
					html.Span(g.Text("Settings")),
					Shortcut(ShortcutProps{}, "⌘S"),
				),
			),
			Separator(SeparatorProps{}),
			DropdownLabel(LabelProps{}, "Appearance"),
			CheckboxItemHTMX(
				CheckboxItemProps{Checked: state.Checkboxes["statusBar"]},
				htmxProps,
				"statusBar",
				g.Text("Status Bar"),
			),
			CheckboxItemHTMX(
				CheckboxItemProps{Checked: state.Checkboxes["activityBar"], Disabled: true},
				htmxProps,
				"activityBar",
				g.Text("Activity Bar"),
			),
			CheckboxItemHTMX(
				CheckboxItemProps{Checked: state.Checkboxes["panel"]},
				htmxProps,
				"panel",
				g.Text("Panel"),
			),
			Separator(SeparatorProps{}),
			DropdownLabel(LabelProps{}, "Panel Position"),
			RadioGroup(
				RadioGroupProps{Value: state.RadioGroups["position"]},
				RadioItemHTMX(
					RadioItemProps{Value: "top"},
					htmxProps,
					"position",
					state.RadioGroups["position"] == "top",
					g.Text("Top"),
				),
				RadioItemHTMX(
					RadioItemProps{Value: "bottom"},
					htmxProps,
					"position",
					state.RadioGroups["position"] == "bottom",
					g.Text("Bottom"),
				),
				RadioItemHTMX(
					RadioItemProps{Value: "right"},
					htmxProps,
					"position",
					state.RadioGroups["position"] == "right",
					g.Text("Right"),
				),
			),
			Separator(SeparatorProps{}),
			ItemHTMX(
				ItemProps{},
				htmxProps,
				htmxProps.ItemPath + "/logout",
				icons.LogOut(g.Attr("class", "h-4 w-4")),
				html.Span(g.Text("Log out")),
				Shortcut(ShortcutProps{}, "⇧⌘Q"),
			),
		),
	)
}

// CommandPaletteExampleHTMX creates a command palette dropdown with search
func CommandPaletteExampleHTMX() g.Node {
	htmxProps := HTMXProps{
		ID:         "command-dropdown",
		TogglePath: "/api/dropdown/command/toggle",
		ItemPath:   "/api/dropdown/command/item",
	}

	return html.Div(
		html.Button(
			html.Type("button"),
			html.Class("inline-flex items-center gap-2 border rounded-md px-3 py-2 text-sm"),
			hx.Post(htmxProps.TogglePath),
			html.Target("#" + htmxProps.ID),
			hx.Swap("outerHTML"),
			hx.Vals(`{"open": "true"}`),
			g.Text("Commands"),
	html.Kbd(html.Class("ml-2 text-xs"), g.Text("⌘K")),
		),
		html.Div(html.ID(htmxProps.ID)),
	)
}

// RenderCommandDropdown renders a command palette dropdown
func RenderCommandDropdown(htmxProps HTMXProps, isOpen bool, query string) g.Node {
	if !isOpen {
		return html.Div(html.ID(htmxProps.ID))
	}

	// Mock commands
	commands := []struct {
		Name     string
		Icon     g.Node
		Shortcut string
		Action   string
	}{
		{"Copy", icons.Copy(html.Class("h-4 w-4")), "⌘C", "/copy"},
		{"Cut", icons.Cut(html.Class("h-4 w-4")), "⌘X", "/cut"},
		{"Paste", icons.Paste(html.Class("h-4 w-4")), "⌘V", "/paste"},
		{"Select All", icons.SelectAll(html.Class("h-4 w-4")), "⌘A", "/select-all"},
		{"Undo", icons.Undo(html.Class("h-4 w-4")), "⌘Z", "/undo"},
		{"Redo", icons.Redo(html.Class("h-4 w-4")), "⇧⌘Z", "/redo"},
	}

	// Filter commands based on query
	var filtered []struct {
		Name     string
		Icon     g.Node
		Shortcut string
		Action   string
	}

	if query == "" {
		filtered = commands
	} else {
		for _, cmd := range commands {
			if strings.Contains(strings.ToLower(cmd.Name), strings.ToLower(query)) {
				filtered = append(filtered, cmd)
			}
		}
	}

	return NewHTMX(
		Props{Open: true},
		htmxProps,
		TriggerHTMX(
			TriggerProps{AsChild: true},
			htmxProps,
			true,
			html.Button(
				html.Type("button"),
				html.Class("inline-flex items-center gap-2 border rounded-md px-3 py-2 text-sm"),
				g.Text("Commands"),
	html.Kbd(html.Class("ml-2 text-xs"), g.Text("⌘K")),
			),
		),
		DropdownContentHTMX(
			ContentProps{Class: "w-72"},
			htmxProps,
			html.Div(html.Class("flex items-center border-b px-3"),
				icons.Search(html.Class("mr-2 h-4 w-4 shrink-0 opacity-50")),
				html.Input(
					html.Type("text"),
					html.Class("flex h-11 w-full rounded-md bg-transparent py-3 text-sm outline-none placeholder:text-muted-foreground"),
					html.Placeholder("Search commands..."),
					html.Value(query),
					hx.Post(htmxProps.ItemPath + "/search"),
					hx.Trigger("keyup changed delay:300ms"),
					html.Target("#" + htmxProps.ID),
					hx.Swap("outerHTML"),
					hx.Include("#" + htmxProps.ID),
				),
			),
			html.Div(html.Class("max-h-[300px] overflow-y-auto p-1"),
				g.If(len(filtered) == 0,
					html.P(html.Class("py-6 text-center text-sm text-muted-foreground"),
						g.Text("No commands found."),
					),
				),
				g.Group(g.Map(filtered, func(cmd struct {
					Name     string
					Icon     g.Node
					Shortcut string
					Action   string
				}) g.Node {
					return ItemHTMX(
						ItemProps{},
						htmxProps,
						htmxProps.ItemPath + cmd.Action,
						cmd.Icon,
						html.Span(g.Text(cmd.Name)),
						Shortcut(ShortcutProps{}, cmd.Shortcut),
					)
				})),
			),
		),
	)
}

// DropdownMenuHandlers creates HTTP handlers for dropdown menu components
func DropdownMenuHandlers(mux *http.ServeMux) {
	// State storage (in production, use a proper session store)
	menuStates := make(map[string]*DropdownMenuState)

	// Basic dropdown handlers
	htmxProps := HTMXProps{
		ID:         "dropdown-example",
		TogglePath: "/api/dropdown/toggle",
		ItemPath:   "/api/dropdown/item",
	}

	mux.HandleFunc("/api/dropdown/toggle", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		open := r.FormValue("open") == "true"

		// Get or create state
		state, ok := menuStates[htmxProps.ID]
		if !ok {
			state = &DropdownMenuState{
				Checkboxes:  make(map[string]bool),
				RadioGroups: make(map[string]string),
			}
			// Set defaults
			state.Checkboxes["statusBar"] = true
			state.Checkboxes["activityBar"] = false
			state.Checkboxes["panel"] = false
			state.RadioGroups["position"] = "bottom"
			menuStates[htmxProps.ID] = state
		}
		state.Open = open

		node := RenderDropdownMenu(htmxProps, open, state)
		node.Render(w)
	})

	mux.HandleFunc("/api/dropdown/item/checkbox", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		name := r.FormValue("name")
		checked := r.FormValue("checked") == "true"

		// Update state
		state := menuStates[htmxProps.ID]
		if state == nil {
			state = &DropdownMenuState{
				Checkboxes:  make(map[string]bool),
				RadioGroups: make(map[string]string),
			}
			menuStates[htmxProps.ID] = state
		}
		state.Checkboxes[name] = checked

		// Re-render menu
		node := RenderDropdownMenu(htmxProps, true, state)
		node.Render(w)
	})

	mux.HandleFunc("/api/dropdown/item/radio", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		group := r.FormValue("group")
		value := r.FormValue("value")

		// Update state
		state := menuStates[htmxProps.ID]
		if state == nil {
			state = &DropdownMenuState{
				Checkboxes:  make(map[string]bool),
				RadioGroups: make(map[string]string),
			}
			menuStates[htmxProps.ID] = state
		}
		state.RadioGroups[group] = value

		// Re-render menu
		node := RenderDropdownMenu(htmxProps, true, state)
		node.Render(w)
	})

	// Handle menu item actions
	mux.HandleFunc("/api/dropdown/item/", func(w http.ResponseWriter, r *http.Request) {
		action := strings.TrimPrefix(r.URL.Path, "/api/dropdown/item/")

		// Close menu and show notification
		var message string
		switch action {
		case "profile":
			message = "Opening profile..."
		case "billing":
			message = "Opening billing..."
		case "settings":
			message = "Opening settings..."
		case "logout":
			message = "Logging out..."
		default:
			message = fmt.Sprintf("Action: %s", action)
		}

		node := html.Div(
			html.ID(htmxProps.ID),
			html.Div(
				html.Class("fixed bottom-4 right-4 bg-green-500 text-white px-4 py-2 rounded-md shadow-lg"),
				g.Text(message),
				// Auto-hide after 3 seconds
				g.Attr("x-data", "{}"),
				g.Attr("x-init", "setTimeout(() => $el.remove(), 3000)"),
			),
		)
		node.Render(w)
	})

	// Command palette handlers
	cmdProps := HTMXProps{
		ID:         "command-dropdown",
		TogglePath: "/api/dropdown/command/toggle",
		ItemPath:   "/api/dropdown/command/item",
	}

	mux.HandleFunc("/api/dropdown/command/toggle", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		open := r.FormValue("open") == "true"

		node := RenderCommandDropdown(cmdProps, open, "")
		node.Render(w)
	})

	mux.HandleFunc("/api/dropdown/command/item/search", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		query := r.FormValue("value")

		node := RenderCommandDropdown(cmdProps, true, query)
		node.Render(w)
	})

	// Command actions
	mux.HandleFunc("/api/dropdown/command/item/", func(w http.ResponseWriter, r *http.Request) {
		action := strings.TrimPrefix(r.URL.Path, "/api/dropdown/command/item/")

		// Close menu and show action
		node := html.Div(
			html.ID(cmdProps.ID),
			html.Div(
				html.Class("fixed bottom-4 right-4 bg-blue-500 text-white px-4 py-2 rounded-md shadow-lg"),
				g.Textf("Executed: %s", action),
				// Auto-hide after 2 seconds
				g.Attr("x-data", "{}"),
				g.Attr("x-init", "setTimeout(() => $el.remove(), 2000)"),
			),
		)
		node.Render(w)
	})
}

// Define missing icons for the command example
var (
	Copy = func(attrs ...g.Node) g.Node {
		defaultAttrs := []g.Node{
			g.Attr("viewBox", "0 0 24 24"),
			g.Attr("width", "24"),
			g.Attr("height", "24"),
			g.Attr("fill", "none"),
			g.Attr("stroke", "currentColor"),
			g.Attr("stroke-width", "2"),
			g.Attr("stroke-linecap", "round"),
			g.Attr("stroke-linejoin", "round"),
		}
		allAttrs := append(defaultAttrs, attrs...)
		return g.El("svg",
			append(allAttrs,
				g.El("rect", g.Attr("width", "14"), g.Attr("height", "14"), g.Attr("x", "8"), g.Attr("y", "8"), g.Attr("rx", "2"), g.Attr("ry", "2")),
				g.El("path", g.Attr("d", "M4 16c-1.11 0-2-.9-2-2V4c0-1.11.89-2 2-2h10c1.11 0 2 .89 2 2")),
			)...,
		)
	}

	Cut = func(attrs ...g.Node) g.Node {
		defaultAttrs := []g.Node{
			g.Attr("viewBox", "0 0 24 24"),
			g.Attr("width", "24"),
			g.Attr("height", "24"),
			g.Attr("fill", "none"),
			g.Attr("stroke", "currentColor"),
			g.Attr("stroke-width", "2"),
			g.Attr("stroke-linecap", "round"),
			g.Attr("stroke-linejoin", "round"),
		}
		allAttrs := append(defaultAttrs, attrs...)
		return g.El("svg",
			append(allAttrs,
				g.El("circle", g.Attr("cx", "6"), g.Attr("cy", "6"), g.Attr("r", "3")),
				g.El("circle", g.Attr("cx", "6"), g.Attr("cy", "18"), g.Attr("r", "3")),
				g.El("line", g.Attr("x1", "20"), g.Attr("y1", "4"), g.Attr("x2", "8.12"), g.Attr("y2", "15.88")),
				g.El("line", g.Attr("x1", "14.47"), g.Attr("y1", "14.48"), g.Attr("x2", "20"), g.Attr("y2", "20")),
				g.El("line", g.Attr("x1", "8.12"), g.Attr("y1", "8.12"), g.Attr("x2", "12"), g.Attr("y2", "12")),
			)...,
		)
	}

	Paste = func(attrs ...g.Node) g.Node {
		defaultAttrs := []g.Node{
			g.Attr("viewBox", "0 0 24 24"),
			g.Attr("width", "24"),
			g.Attr("height", "24"),
			g.Attr("fill", "none"),
			g.Attr("stroke", "currentColor"),
			g.Attr("stroke-width", "2"),
			g.Attr("stroke-linecap", "round"),
			g.Attr("stroke-linejoin", "round"),
		}
		allAttrs := append(defaultAttrs, attrs...)
		return g.El("svg",
			append(allAttrs,
				g.El("path", g.Attr("d", "M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2")),
				g.El("rect", g.Attr("width", "8"), g.Attr("height", "4"), g.Attr("x", "8"), g.Attr("y", "2"), g.Attr("rx", "1"), g.Attr("ry", "1")),
			)...,
		)
	}

	SelectAll = func(attrs ...g.Node) g.Node {
		defaultAttrs := []g.Node{
			g.Attr("viewBox", "0 0 24 24"),
			g.Attr("width", "24"),
			g.Attr("height", "24"),
			g.Attr("fill", "none"),
			g.Attr("stroke", "currentColor"),
			g.Attr("stroke-width", "2"),
			g.Attr("stroke-linecap", "round"),
			g.Attr("stroke-linejoin", "round"),
		}
		allAttrs := append(defaultAttrs, attrs...)
		return g.El("svg",
			append(allAttrs,
				g.El("path", g.Attr("d", "M3 5h2")),
				g.El("path", g.Attr("d", "M3 11h2")),
				g.El("path", g.Attr("d", "M3 17h2")),
				g.El("path", g.Attr("d", "M9 5h12")),
				g.El("path", g.Attr("d", "M9 11h12")),
				g.El("path", g.Attr("d", "M9 17h12")),
			)...,
		)
	}

	Undo = func(attrs ...g.Node) g.Node {
		defaultAttrs := []g.Node{
			g.Attr("viewBox", "0 0 24 24"),
			g.Attr("width", "24"),
			g.Attr("height", "24"),
			g.Attr("fill", "none"),
			g.Attr("stroke", "currentColor"),
			g.Attr("stroke-width", "2"),
			g.Attr("stroke-linecap", "round"),
			g.Attr("stroke-linejoin", "round"),
		}
		allAttrs := append(defaultAttrs, attrs...)
		return g.El("svg",
			append(allAttrs,
				g.El("path", g.Attr("d", "M3 7v6h6")),
				g.El("path", g.Attr("d", "M21 17a9 9 0 00-9-9 9 9 0 00-6 2.3L3 13")),
			)...,
		)
	}

	Redo = func(attrs ...g.Node) g.Node {
		defaultAttrs := []g.Node{
			g.Attr("viewBox", "0 0 24 24"),
			g.Attr("width", "24"),
			g.Attr("height", "24"),
			g.Attr("fill", "none"),
			g.Attr("stroke", "currentColor"),
			g.Attr("stroke-width", "2"),
			g.Attr("stroke-linecap", "round"),
			g.Attr("stroke-linejoin", "round"),
		}
		allAttrs := append(defaultAttrs, attrs...)
		return g.El("svg",
			append(allAttrs,
				g.El("path", g.Attr("d", "M21 7v6h-6")),
				g.El("path", g.Attr("d", "M3 17a9 9 0 019-9 9 9 0 016 2.3l3 2.7")),
			)...,
		)
	}
)