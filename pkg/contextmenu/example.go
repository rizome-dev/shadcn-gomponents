package contextmenu

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

// Example demonstrates the ContextMenu component
func Example() g.Node {
	return html.Div(html.Class("space-y-8"),
		// Basic Example
		html.Div(
			html.H3(html.Class("text-lg font-semibold"), g.Text("Basic Context Menu")),
			New(
				Props{ID: "basic-menu"},
				Trigger(
					TriggerProps{Class: "flex h-[150px] w-[300px] items-center justify-center rounded-md border border-dashed text-sm"},
					g.Text("Right click here"),
				),
				ContentComponent(
					ContentProps{},
					Item(ItemProps{}, g.Text("Back")),
					Item(ItemProps{Disabled: true}, g.Text("Forward")),
					Item(ItemProps{}, g.Text("Reload")),
					Item(ItemProps{}, g.Text("More Tools")),
					Separator(SeparatorProps{}),
					Item(ItemProps{}, 
						g.Text("Show Bookmarks Bar"),
						Shortcut(ShortcutProps{}, g.Text("⌘⇧B")),
					),
					Item(ItemProps{}, g.Text("Show Full URLs")),
					Separator(SeparatorProps{}),
					LabelComponent(LabelProps{Inset: true}, g.Text("People")),
					Item(ItemProps{}, g.Text("Pedro Duarte")),
					Item(ItemProps{}, g.Text("Colm Tuite")),
				),
			),
		),

		// With Checkboxes and Radio Items
		html.Div(
			html.H3(html.Class("text-lg font-semibold"), g.Text("With Checkboxes and Radio Items")),
			New(
				Props{ID: "checkbox-menu"},
				Trigger(
					TriggerProps{Class: "flex h-[150px] w-[300px] items-center justify-center rounded-md border border-dashed text-sm"},
					g.Text("Right click for options"),
				),
				ContentComponent(
					ContentProps{},
					Item(ItemProps{}, g.Text("Undo")),
					Item(ItemProps{}, g.Text("Redo")),
					Separator(SeparatorProps{}),
					CheckboxItem(CheckboxItemProps{Checked: true}, g.Text("Show Status Bar")),
					CheckboxItem(CheckboxItemProps{Checked: false, Disabled: true}, g.Text("Show Activity Bar")),
					CheckboxItem(CheckboxItemProps{Checked: true}, g.Text("Show Panel")),
					Separator(SeparatorProps{}),
					RadioGroup(
						RadioGroupProps{Value: "bottom"},
						LabelComponent(LabelProps{Inset: true}, g.Text("Panel Position")),
						RadioItem(RadioItemProps{Value: "top"}, g.Text("Top")),
						RadioItem(RadioItemProps{Value: "bottom"}, g.Text("Bottom")),
						RadioItem(RadioItemProps{Value: "right"}, g.Text("Right")),
					),
				),
			),
		),

		// With Submenus
		html.Div(
			html.H3(html.Class("text-lg font-semibold"), g.Text("With Submenus")),
			New(
				Props{ID: "submenu-menu"},
				Trigger(
					TriggerProps{Class: "flex h-[150px] w-[300px] items-center justify-center rounded-md border border-dashed text-sm"},
					g.Text("Right click for nested menus"),
				),
				ContentComponent(
					ContentProps{},
					Item(ItemProps{}, g.Text("New Tab")),
					Item(ItemProps{}, g.Text("New Window")),
					Item(ItemProps{}, g.Text("New Private Window")),
					Separator(SeparatorProps{}),
					SubMenu(
						SubProps{},
						SubTrigger(SubTriggerProps{}, g.Text("Share")),
						SubContent(
							SubContentProps{},
							Item(ItemProps{}, g.Text("Email link")),
							Item(ItemProps{}, g.Text("Messages")),
							Item(ItemProps{}, g.Text("Notes")),
						),
					),
					SubMenu(
						SubProps{},
						SubTrigger(SubTriggerProps{}, g.Text("More Tools")),
						SubContent(
							SubContentProps{},
							Item(ItemProps{}, 
								g.Text("Save Page As..."),
								Shortcut(ShortcutProps{}, g.Text("⇧⌘S")),
							),
							Item(ItemProps{}, g.Text("Create Shortcut...")),
							Item(ItemProps{}, g.Text("Name Window...")),
							Separator(SeparatorProps{}),
							Item(ItemProps{}, g.Text("Developer Tools")),
						),
					),
				),
			),
		),

		// HTMX Example
		html.Div(
			html.H3(html.Class("text-lg font-semibold"), g.Text("HTMX Context Menu")),
			ExampleHTMX(),
			html.P(html.Class("text-sm text-muted-foreground mt-2"), 
				g.Text("This context menu uses HTMX for dynamic interactions."),
			),
		),
	)
}

// ExampleWithIcons demonstrates context menu with icons
func ExampleWithIcons() g.Node {
	// Helper function for icon
	icon := func(path string) g.Node {
		return g.El("svg",
			g.Attr("viewBox", "0 0 24 24"),
			g.Attr("fill", "none"),
			g.Attr("stroke", "currentColor"),
			g.Attr("stroke-width", "2"),
			g.Attr("stroke-linecap", "round"),
			g.Attr("stroke-linejoin", "round"),
			html.Class("mr-2 h-4 w-4"),
			g.El("path", g.Attr("d", path)),
		)
	}

	return New(
		Props{ID: "icon-menu"},
		Trigger(
			TriggerProps{Class: "flex h-[150px] w-[300px] items-center justify-center rounded-md border border-dashed text-sm"},
			g.Text("Right click for icon menu"),
		),
		ContentComponent(
			ContentProps{},
			Item(ItemProps{}, 
				icon("M12 2v20M17 5H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6"),
				g.Text("Cut"),
			),
			Item(ItemProps{}, 
				icon("M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2"),
				g.Text("Copy"),
			),
			Item(ItemProps{}, 
				icon("M9 2h6v6l-3 3z"),
				g.Text("Paste"),
			),
			Separator(SeparatorProps{}),
			Item(ItemProps{}, 
				icon("M3 6h18M8 6V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6"),
				g.Text("Delete"),
			),
		),
	)
}

// ExampleCustomStyling demonstrates custom styling
func ExampleCustomStyling() g.Node {
	return New(
		Props{ID: "custom-menu", Class: "custom-context-menu"},
		Trigger(
			TriggerProps{Class: "flex h-[150px] w-[300px] items-center justify-center rounded-lg bg-gradient-to-r from-purple-500 to-pink-500 text-white font-medium"},
			g.Text("Right click for custom menu"),
		),
		ContentComponent(
			ContentProps{Class: "bg-gradient-to-br from-purple-50 to-pink-50 dark:from-purple-900/20 dark:to-pink-900/20"},
			Item(ItemProps{Class: "hover:bg-purple-100 dark:hover:bg-purple-800/30"}, g.Text("Purple Action")),
			Item(ItemProps{Class: "hover:bg-pink-100 dark:hover:bg-pink-800/30"}, g.Text("Pink Action")),
			Separator(SeparatorProps{Class: "bg-purple-200 dark:bg-purple-700"}),
			Item(ItemProps{Class: "hover:bg-gradient-to-r hover:from-purple-100 hover:to-pink-100"}, 
				g.Text("Gradient Action"),
			),
		),
	)
}