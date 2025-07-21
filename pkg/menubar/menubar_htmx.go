package menubar

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	hx "maragu.dev/gomponents-htmx"
	"github.com/rizome-dev/shadcn-gomponents/lib"
	"fmt"
)

// HTMXProps defines HTMX-specific properties for the Menubar
type HTMXProps struct {
	ID           string // Unique ID for the menubar
	MenuPath     string // Server path for menu content
	TriggerEvent string // Event to trigger menu (default: "click")
	CloseOnClick bool   // Close menu on item click (default: true)
}

// MenuHTMXProps defines HTMX properties for individual menus
type MenuHTMXProps struct {
	ID          string // Unique ID for the menu
	ContentPath string // Server path to fetch menu content
	Target      string // HTMX target for content
}

// NewHTMX creates an HTMX-enhanced Menubar component
func NewHTMX(props Props, htmxProps HTMXProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"flex h-10 items-center space-x-1 rounded-md border bg-background p-1",
		props.Class,
	)

	attrs := []g.Node{
		html.Class(classes),
		html.Role("menubar"),
		g.If(htmxProps.ID != "", html.ID(htmxProps.ID)),
		g.Attr("data-menubar", "true"),
	}

	// Add global click handler to close menus
	if htmxProps.CloseOnClick {
		attrs = append(attrs, g.Attr("data-close-on-click", "true"))
	}

	return html.Div(
		append(attrs, children...)...,
	)
}

// MenuHTMX creates an HTMX-enhanced menu
func MenuHTMX(props MenuProps, htmxProps MenuHTMXProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"relative",
		props.Class,
	)

	return html.Div(
		html.ID(htmxProps.ID),
		html.Class(classes),
		g.Attr("data-menu", "true"),
		g.Group(children),
	)
}

// TriggerHTMX creates an HTMX-enhanced menu trigger
func TriggerHTMX(props TriggerProps, htmxProps MenuHTMXProps, children ...g.Node) g.Node {
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

	target := htmxProps.Target
	if target == "" {
		target = "#" + htmxProps.ID + "-content"
	}

	attrs := []g.Node{
		html.Type("button"),
		html.Class(classes),
		html.Role("menuitem"),
		g.Attr("aria-haspopup", "menu"),
		g.Attr("aria-expanded", "false"),
		g.Attr("data-state", "closed"),
		g.Attr("data-menu-trigger", htmxProps.ID),
		
		// HTMX attributes
		hx.Get(htmxProps.ContentPath),
		hx.Target(target),
		hx.Swap("innerHTML"),
		hx.Trigger("click"),
		
		// Toggle menu state
		g.Attr("onclick", fmt.Sprintf(`
			const trigger = this;
			const menu = trigger.closest('[data-menu]');
			const menubar = trigger.closest('[data-menubar]');
			const isOpen = trigger.getAttribute('data-state') === 'open';
			
			// Close all other menus
			if (menubar) {
				menubar.querySelectorAll('[data-menu-trigger]').forEach(t => {
					if (t !== trigger) {
						t.setAttribute('data-state', 'closed');
						t.setAttribute('aria-expanded', 'false');
						const contentId = t.getAttribute('hx-target');
						if (contentId) {
							const content = document.querySelector(contentId);
							if (content) content.style.display = 'none';
						}
					}
				});
			}
			
			// Toggle current menu
			if (!isOpen) {
				trigger.setAttribute('data-state', 'open');
				trigger.setAttribute('aria-expanded', 'true');
			} else {
				trigger.setAttribute('data-state', 'closed');
				trigger.setAttribute('aria-expanded', 'false');
				const content = document.querySelector('%s');
				if (content) content.style.display = 'none';
			}
		`, target)),
	}

	if props.Disabled {
		attrs = append(attrs, html.Disabled())
	}

	return html.Button(
		append(attrs, children...)...,
	)
}

// ContentHTMX creates HTMX-aware menu content container
func ContentHTMX(props ContentProps, htmxProps MenuHTMXProps, children ...g.Node) g.Node {
	// Set defaults
	if props.Align == "" {
		props.Align = "start"
	}
	if props.Side == "" {
		props.Side = "bottom"
	}

	classes := lib.CN(
		"z-50 min-w-[8rem] overflow-hidden rounded-md border bg-popover p-1 text-popover-foreground shadow-md",
		"animate-in fade-in-0 zoom-in-95",
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

	target := htmxProps.Target
	if target == "" {
		target = htmxProps.ID + "-content"
	}

	return html.Div(
		html.ID(target),
		html.Class(lib.CN(classes, positionClasses)),
		html.Role("menu"),
		g.Attr("aria-orientation", "vertical"),
		g.Attr("data-state", "open"),
		g.Attr("data-side", props.Side),
		g.Attr("data-menu-content", htmxProps.ID),
		
		// Click outside to close
		g.Attr("hx-on:click.outside", fmt.Sprintf(`
			const trigger = document.querySelector('[data-menu-trigger="%s"]');
			if (trigger) {
				trigger.setAttribute('data-state', 'closed');
				trigger.setAttribute('aria-expanded', 'false');
			}
			this.style.display = 'none';
		`, htmxProps.ID)),
		
		// Keyboard navigation
		g.Attr("onkeydown", `
			const content = event.currentTarget;
			const items = Array.from(content.querySelectorAll('[role="menuitem"]:not([data-disabled="true"])'));
			const currentIndex = items.findIndex(item => item === document.activeElement);
			
			switch(event.key) {
				case 'ArrowDown':
					event.preventDefault();
					const nextIndex = currentIndex < items.length - 1 ? currentIndex + 1 : 0;
					items[nextIndex]?.focus();
					break;
				case 'ArrowUp':
					event.preventDefault();
					const prevIndex = currentIndex > 0 ? currentIndex - 1 : items.length - 1;
					items[prevIndex]?.focus();
					break;
				case 'Escape':
					const trigger = document.querySelector('[data-menu-trigger="${htmxProps.ID}"]');
					if (trigger) {
						trigger.setAttribute('data-state', 'closed');
						trigger.setAttribute('aria-expanded', 'false');
						trigger.focus();
					}
					content.style.display = 'none';
					break;
			}
		`),
		
		g.Group(children),
	)
}

// ItemHTMX creates an HTMX-enhanced menu item
func ItemHTMX(props ItemProps, actionPath string, children ...g.Node) g.Node {
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
		
		// HTMX action if provided
		g.If(actionPath != "",
			g.Group([]g.Node{
				hx.Post(actionPath),
				hx.Trigger("click"),
			}),
		),
		
		// Close menu on click
		g.Attr("onclick", `
			const menuContent = this.closest('[data-menu-content]');
			const menuId = menuContent?.dataset.menuContent;
			if (menuId) {
				const trigger = document.querySelector('[data-menu-trigger="' + menuId + '"]');
				if (trigger) {
					trigger.setAttribute('data-state', 'closed');
					trigger.setAttribute('aria-expanded', 'false');
				}
				menuContent.style.display = 'none';
			}
		`),
	}

	if props.Disabled {
		attrs = append(attrs, g.Attr("data-disabled", "true"))
	}

	return html.Div(
		append(attrs, children...)...,
	)
}

// CheckboxItemHTMX creates an HTMX-enhanced checkbox menu item
func CheckboxItemHTMX(props CheckboxItemProps, togglePath string, children ...g.Node) g.Node {
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
		g.Attr("data-checked", lib.CNIf(props.Checked, "true", "false")),
		
		// HTMX toggle
		hx.Post(togglePath),
		hx.Vals(fmt.Sprintf(`{"name": "%s", "value": "%s", "checked": %v}`,
			props.Name, props.Value, !props.Checked)),
		hx.Trigger("click"),
		hx.Target("this"),
		hx.Swap("outerHTML"),
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

// RadioItemHTMX creates an HTMX-enhanced radio menu item
func RadioItemHTMX(props RadioItemProps, groupName string, selectPath string, isSelected bool, children ...g.Node) g.Node {
	classes := lib.CN(
		"relative flex cursor-default select-none items-center rounded-sm py-1.5 pl-8 pr-2 text-sm outline-none transition-colors",
		"focus:bg-accent focus:text-accent-foreground",
		"data-[disabled]:pointer-events-none data-[disabled]:opacity-50",
		props.Class,
	)

	attrs := []g.Node{
		html.Class(classes),
		html.Role("menuitemradio"),
		g.Attr("aria-checked", lib.CNIf(isSelected, "true", "false")),
		html.TabIndex("-1"),
		
		// HTMX selection
		hx.Post(selectPath),
		hx.Vals(fmt.Sprintf(`{"group": "%s", "value": "%s"}`, groupName, props.Value)),
		hx.Trigger("click"),
		hx.Target(fmt.Sprintf("[data-radio-group='%s']", groupName)),
		hx.Swap("innerHTML"),
	}

	if props.Disabled {
		attrs = append(attrs, g.Attr("data-disabled", "true"))
	}

	indicator := html.Span(
		html.Class("absolute left-2 flex h-3.5 w-3.5 items-center justify-center"),
		g.If(isSelected,
			g.Raw(`<svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg">
				<path d="M7.5 9.5C8.60457 9.5 9.5 8.60457 9.5 7.5C9.5 6.39543 8.60457 5.5 7.5 5.5C6.39543 5.5 5.5 6.39543 5.5 7.5C5.5 8.60457 6.39543 9.5 7.5 9.5Z" fill="currentColor"></path>
			</svg>`),
		),
	)

	return html.Div(
		append(attrs, indicator, g.Group(children))...,
	)
}

// ExampleHTMX creates an HTMX-enhanced menubar example
func ExampleHTMX() g.Node {
	return NewHTMX(
		Props{},
		HTMXProps{
			ID:           "main-menubar",
			CloseOnClick: true,
		},
		
		// File menu
		MenuHTMX(
			MenuProps{},
			MenuHTMXProps{
				ID:          "file-menu",
				ContentPath: "/api/menubar/file",
			},
			TriggerHTMX(
				TriggerProps{},
				MenuHTMXProps{ID: "file-menu", ContentPath: "/api/menubar/file"},
				g.Text("File"),
			),
			html.Div(html.ID("file-menu-content")), // Content placeholder
		),
		
		// Edit menu
		MenuHTMX(
			MenuProps{},
			MenuHTMXProps{
				ID:          "edit-menu",
				ContentPath: "/api/menubar/edit",
			},
			TriggerHTMX(
				TriggerProps{},
				MenuHTMXProps{ID: "edit-menu", ContentPath: "/api/menubar/edit"},
				g.Text("Edit"),
			),
			html.Div(html.ID("edit-menu-content")), // Content placeholder
		),
		
		// View menu
		MenuHTMX(
			MenuProps{},
			MenuHTMXProps{
				ID:          "view-menu",
				ContentPath: "/api/menubar/view",
			},
			TriggerHTMX(
				TriggerProps{},
				MenuHTMXProps{ID: "view-menu", ContentPath: "/api/menubar/view"},
				g.Text("View"),
			),
			html.Div(html.ID("view-menu-content")), // Content placeholder
		),
	)
}

// RenderFileMenu renders the file menu content (for server response)
func RenderFileMenu() g.Node {
	return ContentHTMX(
		ContentProps{},
		MenuHTMXProps{ID: "file-menu"},
		ItemHTMX(ItemProps{}, "/api/file/new", 
			g.Text("New File"),
			Shortcut(ShortcutProps{}, g.Text("⌘N")),
		),
		ItemHTMX(ItemProps{}, "/api/file/open",
			g.Text("Open..."),
			Shortcut(ShortcutProps{}, g.Text("⌘O")),
		),
		ItemHTMX(ItemProps{}, "/api/file/save",
			g.Text("Save"),
			Shortcut(ShortcutProps{}, g.Text("⌘S")),
		),
		Separator(SeparatorProps{}),
		SubMenu(
			SubMenuProps{},
			SubTrigger(SubTriggerProps{}, g.Text("Recent Files")),
			// Submenu would be handled similarly
		),
		Separator(SeparatorProps{}),
		ItemHTMX(ItemProps{}, "/api/file/exit",
			g.Text("Exit"),
			Shortcut(ShortcutProps{}, g.Text("⌘Q")),
		),
	)
}

// RenderEditMenu renders the edit menu content (for server response)
func RenderEditMenu() g.Node {
	return ContentHTMX(
		ContentProps{},
		MenuHTMXProps{ID: "edit-menu"},
		ItemHTMX(ItemProps{Disabled: true}, "",
			g.Text("Undo"),
			Shortcut(ShortcutProps{}, g.Text("⌘Z")),
		),
		ItemHTMX(ItemProps{Disabled: true}, "",
			g.Text("Redo"),
			Shortcut(ShortcutProps{}, g.Text("⇧⌘Z")),
		),
		Separator(SeparatorProps{}),
		ItemHTMX(ItemProps{}, "/api/edit/cut",
			g.Text("Cut"),
			Shortcut(ShortcutProps{}, g.Text("⌘X")),
		),
		ItemHTMX(ItemProps{}, "/api/edit/copy",
			g.Text("Copy"),
			Shortcut(ShortcutProps{}, g.Text("⌘C")),
		),
		ItemHTMX(ItemProps{}, "/api/edit/paste",
			g.Text("Paste"),
			Shortcut(ShortcutProps{}, g.Text("⌘V")),
		),
		Separator(SeparatorProps{}),
		ItemHTMX(ItemProps{}, "/api/edit/select-all",
			g.Text("Select All"),
			Shortcut(ShortcutProps{}, g.Text("⌘A")),
		),
	)
}

// RenderViewMenu renders the view menu content with checkboxes (for server response)
func RenderViewMenu(showStatusBar, showSidebar bool) g.Node {
	return ContentHTMX(
		ContentProps{},
		MenuHTMXProps{ID: "view-menu"},
		CheckboxItemHTMX(
			CheckboxItemProps{
				Checked: showStatusBar,
				Name:    "view",
				Value:   "statusbar",
			},
			"/api/view/toggle",
			g.Text("Status Bar"),
		),
		CheckboxItemHTMX(
			CheckboxItemProps{
				Checked: showSidebar,
				Name:    "view",
				Value:   "sidebar",
			},
			"/api/view/toggle",
			g.Text("Sidebar"),
		),
		Separator(SeparatorProps{}),
		html.Label( g.Text("Panel Position")),
		html.Div(
			g.Attr("data-radio-group", "panel-position"),
			RadioItemHTMX(
				RadioItemProps{Value: "top"},
				"panel-position",
				"/api/view/panel-position",
				true,
				g.Text("Top"),
			),
			RadioItemHTMX(
				RadioItemProps{Value: "bottom"},
				"panel-position",
				"/api/view/panel-position",
				false,
				g.Text("Bottom"),
			),
			RadioItemHTMX(
				RadioItemProps{Value: "right"},
				"panel-position",
				"/api/view/panel-position",
				false,
				g.Text("Right"),
			),
		),
	)
}