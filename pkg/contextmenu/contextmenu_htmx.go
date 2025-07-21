package contextmenu

import (
	"fmt"
	"net/http"
	"strings"

	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	hx "maragu.dev/gomponents-htmx"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// HTMXProps defines HTMX-specific properties for the ContextMenu
type HTMXProps struct {
	ID          string // Unique ID for the context menu
	MenuPath    string // Server path for menu actions
	ItemPath    string // Server path for item clicks
}

// NewHTMX creates an HTMX-enhanced ContextMenu component
func NewHTMX(props Props, htmxProps HTMXProps, children ...g.Node) g.Node {
	return html.Div(
		html.ID(htmxProps.ID),
		g.If(props.Class != "", html.Class(props.Class)),
		g.Attr("data-context-menu", "root"),
		hx.On("contextmenu", fmt.Sprintf(`
			event.preventDefault();
			htmx.ajax('POST', '%s', {
				target: '#%s-menu',
				swap: 'innerHTML',
				values: {
					x: event.pageX,
					y: event.pageY
				}
			});
		`, htmxProps.MenuPath, htmxProps.ID)),
		// Click outside to close
		hx.On("click", fmt.Sprintf(`
			const menu = document.querySelector('#%s-menu');
			if (menu && menu.children.length > 0) {
				menu.innerHTML = '';
			}
		`, htmxProps.ID)),
		g.Group(children),
		// Menu container
		html.Div(html.ID(htmxProps.ID+"-menu"), html.Style("position: relative;")),
	)
}

// TriggerHTMX creates an HTMX-enhanced context menu trigger
func TriggerHTMX(props TriggerProps, htmxProps HTMXProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"cursor-context-menu",
		props.Class,
	)

	return html.Div(
		html.Class(classes),
		g.Attr("data-context-menu", "trigger"),
		g.Group(children),
	)
}

// ContentHTMX creates HTMX-enhanced context menu content
func ContentHTMX(props ContentProps, x, y int, children ...g.Node) g.Node {
	classes := lib.CN(
		"z-50 min-w-[8rem] overflow-hidden rounded-md border bg-popover p-1 text-popover-foreground shadow-md",
		"animate-in fade-in-0 zoom-in-95",
		props.Class,
	)

	// Calculate position to ensure menu stays within viewport
	style := fmt.Sprintf("position: fixed; left: %dpx; top: %dpx;", x, y)

	return html.Div(
		html.Class(classes),
		g.Attr("data-context-menu", "content"),
		g.Attr("data-state", "open"),
		html.Style(style),
		// Prevent click propagation to avoid closing menu when clicking inside
		hx.On("click", "event.stopPropagation()"),
		g.Group(children),
	)
}

// ItemHTMX creates an HTMX-enhanced context menu item
func ItemHTMX(props ItemProps, htmxProps HTMXProps, action string, children ...g.Node) g.Node {
	classes := lib.CN(
		"relative flex cursor-default select-none items-center rounded-sm px-2 py-1.5 text-sm outline-none transition-colors hover:bg-accent hover:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50",
		lib.CNIf(props.Inset, "pl-8", ""),
		props.Class,
	)

	attrs := []g.Node{
		html.Class(classes),
		g.Attr("data-context-menu", "item"),
		html.TabIndex("-1"),
	}

	if !props.Disabled {
		attrs = append(attrs,
			hx.Post(htmxProps.ItemPath),
			hx.Vals(fmt.Sprintf(`{"action": "%s"}`, action)),
			html.Target("#"+htmxProps.ID+"-menu"),
			hx.Swap("innerHTML"),
		)
	} else {
		attrs = append(attrs, g.Attr("data-disabled", "true"))
	}

	return html.Div(
		append(attrs, children...)...,
	)
}

// CheckboxItemHTMX creates an HTMX-enhanced checkbox menu item
func CheckboxItemHTMX(props CheckboxItemProps, htmxProps HTMXProps, name string, children ...g.Node) g.Node {
	classes := lib.CN(
		"relative flex cursor-default select-none items-center rounded-sm py-1.5 pl-8 pr-2 text-sm outline-none transition-colors hover:bg-accent hover:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50",
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

	if !props.Disabled {
		attrs = append(attrs,
			hx.Post(htmxProps.ItemPath),
			hx.Vals(fmt.Sprintf(`{"action": "toggle-%s", "checked": %v}`, name, !props.Checked)),
			html.Target("#"+htmxProps.ID+"-menu"),
			hx.Swap("innerHTML"),
		)
	} else {
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

// RadioItemHTMX creates an HTMX-enhanced radio menu item
func RadioItemHTMX(props RadioItemProps, htmxProps HTMXProps, groupName string, selected bool, children ...g.Node) g.Node {
	classes := lib.CN(
		"relative flex cursor-default select-none items-center rounded-sm py-1.5 pl-8 pr-2 text-sm outline-none transition-colors hover:bg-accent hover:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50",
		props.Class,
	)

	attrs := []g.Node{
		html.Class(classes),
		g.Attr("data-context-menu", "radio-item"),
		g.Attr("role", "menuitemradio"),
		g.Attr("data-value", props.Value),
		html.TabIndex("-1"),
	}

	if selected {
		attrs = append(attrs, g.Attr("aria-checked", "true"))
	} else {
		attrs = append(attrs, g.Attr("aria-checked", "false"))
	}

	if !props.Disabled {
		attrs = append(attrs,
			hx.Post(htmxProps.ItemPath),
			hx.Vals(fmt.Sprintf(`{"action": "select-%s", "value": "%s"}`, groupName, props.Value)),
			html.Target("#"+htmxProps.ID+"-menu"),
			hx.Swap("innerHTML"),
		)
	} else {
		attrs = append(attrs, g.Attr("data-disabled", "true"))
	}

	itemChildren := []g.Node{
		html.Span(
			html.Class("absolute left-2 flex h-3.5 w-3.5 items-center justify-center"),
			g.If(selected,
				g.El("svg",
					g.Attr("viewBox", "0 0 24 24"),
					g.Attr("fill", "currentColor"),
					html.Class("h-2 w-2 fill-current"),
					g.El("circle", g.Attr("cx", "12"), g.Attr("cy", "12"), g.Attr("r", "12")),
				),
			),
		),
	}

	return html.Div(
		append(attrs, append(itemChildren, children...)...)...,
	)
}

// SubHTMX creates an HTMX-enhanced submenu
func SubHTMX(props SubProps, htmxProps HTMXProps, subID string, children ...g.Node) g.Node {
	return html.Div(
		g.Attr("data-context-menu", "sub"),
		g.Attr("data-sub-id", subID),
		g.If(props.Open, g.Attr("data-state", "open")),
		g.If(!props.Open, g.Attr("data-state", "closed")),
		g.Group(children),
	)
}

// SubTriggerHTMX creates an HTMX-enhanced submenu trigger
func SubTriggerHTMX(props SubTriggerProps, htmxProps HTMXProps, subID string, children ...g.Node) g.Node {
	classes := lib.CN(
		"flex cursor-default select-none items-center rounded-sm px-2 py-1.5 text-sm outline-none hover:bg-accent data-[state=open]:bg-accent",
		lib.CNIf(props.Inset, "pl-8", ""),
		props.Class,
	)

	attrs := []g.Node{
		html.Class(classes),
		g.Attr("data-context-menu", "sub-trigger"),
		html.TabIndex("-1"),
	}

	if !props.Disabled {
		attrs = append(attrs,
			hx.On("mouseenter", fmt.Sprintf(`
				const submenu = this.nextElementSibling;
				if (submenu && submenu.getAttribute('data-context-menu') === 'sub-content') {
					submenu.style.display = 'block';
					const rect = this.getBoundingClientRect();
					submenu.style.left = rect.right + 'px';
					submenu.style.top = rect.top + 'px';
				}
			`)),
			hx.On("mouseleave", fmt.Sprintf(`
				setTimeout(() => {
					const submenu = this.nextElementSibling;
					if (submenu && !submenu.matches(':hover')) {
						submenu.style.display = 'none';
					}
				}, 100);
			`)),
		)
	} else {
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

// ExampleHTMX creates an HTMX-enhanced context menu example
func ExampleHTMX() g.Node {
	htmxProps := HTMXProps{
		ID:       "context-menu-example",
		MenuPath: "/api/context-menu/show",
		ItemPath: "/api/context-menu/action",
	}

	return NewHTMX(
		Props{},
		htmxProps,
		TriggerHTMX(
			TriggerProps{Class: "flex h-[150px] w-[300px] items-center justify-center rounded-md border border-dashed text-sm"},
			htmxProps,
			g.Text("Right click here"),
		),
	)
}

// ContextMenuHandlers creates HTTP handlers for context menu
func ContextMenuHandlers(mux *http.ServeMux) {
	// Track state
	var bookmarksChecked = true
	var fullURLsChecked = false
	var selectedPerson = "pedro"

	htmxProps := HTMXProps{
		ID:       "context-menu-example",
		MenuPath: "/api/context-menu/show",
		ItemPath: "/api/context-menu/action",
	}

	mux.HandleFunc("/api/context-menu/show", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		x := r.FormValue("x")
		y := r.FormValue("y")

		// Convert string to int
		var xPos, yPos int
		fmt.Sscanf(x, "%d", &xPos)
		fmt.Sscanf(y, "%d", &yPos)

		node := ContentHTMX(
			ContentProps{},
			xPos, yPos,
			ItemHTMX(ItemProps{}, htmxProps, "back", g.Text("Back")),
			ItemHTMX(ItemProps{Disabled: true}, htmxProps, "forward", g.Text("Forward")),
			ItemHTMX(ItemProps{}, htmxProps, "reload", g.Text("Reload")),
			Sub(SubProps{},
				SubTriggerHTMX(SubTriggerProps{}, htmxProps, "more-tools", g.Text("More Tools")),
				SubContent(SubContentProps{},
					ItemHTMX(ItemProps{}, htmxProps, "save-page", 
						g.Text("Save Page As..."),
						Shortcut(ShortcutProps{}, g.Text("⇧⌘S")),
					),
					ItemHTMX(ItemProps{}, htmxProps, "create-shortcut", g.Text("Create Shortcut...")),
					ItemHTMX(ItemProps{}, htmxProps, "name-window", g.Text("Name Window...")),
					Separator(SeparatorProps{}),
					ItemHTMX(ItemProps{}, htmxProps, "developer-tools", g.Text("Developer Tools")),
				),
			),
			Separator(SeparatorProps{}),
			CheckboxItemHTMX(
				CheckboxItemProps{Checked: bookmarksChecked},
				htmxProps, "bookmarks",
				g.Text("Show Bookmarks Bar"),
				Shortcut(ShortcutProps{}, g.Text("⌘⇧B")),
			),
			CheckboxItemHTMX(
				CheckboxItemProps{Checked: fullURLsChecked},
				htmxProps, "urls",
				g.Text("Show Full URLs"),
			),
			Separator(SeparatorProps{}),
			LabelComponent(LabelProps{Inset: true}, g.Text("People")),
			RadioGroup(RadioGroupProps{},
				RadioItemHTMX(RadioItemProps{Value: "pedro"}, htmxProps, "person", selectedPerson == "pedro", g.Text("Pedro Duarte")),
				RadioItemHTMX(RadioItemProps{Value: "colm"}, htmxProps, "person", selectedPerson == "colm", g.Text("Colm Tuite")),
			),
		)

		node.Render(w)
	})

	mux.HandleFunc("/api/context-menu/action", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		action := r.FormValue("action")

		// Handle actions
		switch {
		case action == "toggle-bookmarks":
			bookmarksChecked = !bookmarksChecked
		case action == "toggle-urls":
			fullURLsChecked = !fullURLsChecked
		case strings.HasPrefix(action, "select-person"):
			value := r.FormValue("value")
			selectedPerson = value
		}

		// Return empty to close menu
		w.Write([]byte(""))
	})
}