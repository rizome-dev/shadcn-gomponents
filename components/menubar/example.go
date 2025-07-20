package menubar

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

// Examples demonstrates various Menubar usage patterns
func Examples() g.Node {
	return html.Div(
		html.Class("space-y-8 p-8"),
		
		// Basic Menubar
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Basic Menubar")),
			New(
				Props{},
				Menu(
					MenuProps{},
					Trigger(TriggerProps{}, g.Text("File")),
					ContentComponent(
						ContentProps{},
						Item(ItemProps{}, g.Text("New Tab"), Shortcut(ShortcutProps{}, g.Text("⌘T"))),
						Item(ItemProps{}, g.Text("New Window"), Shortcut(ShortcutProps{}, g.Text("⌘N"))),
						Item(ItemProps{}, g.Text("New Incognito Window"), Shortcut(ShortcutProps{}, g.Text("⇧⌘N"))),
						Separator(SeparatorProps{}),
						SubMenu(
							SubMenuProps{},
							SubTrigger(SubTriggerProps{}, g.Text("Share")),
							SubContent(
								SubContentProps{},
								Item(ItemProps{}, g.Text("Email link")),
								Item(ItemProps{}, g.Text("Messages")),
								Item(ItemProps{}, g.Text("Notes")),
							),
						),
						Separator(SeparatorProps{}),
						Item(ItemProps{}, g.Text("Print..."), Shortcut(ShortcutProps{}, g.Text("⌘P"))),
					),
				),
				Menu(
					MenuProps{},
					Trigger(TriggerProps{}, g.Text("Edit")),
					ContentComponent(
						ContentProps{},
						Item(ItemProps{}, g.Text("Undo"), Shortcut(ShortcutProps{}, g.Text("⌘Z"))),
						Item(ItemProps{}, g.Text("Redo"), Shortcut(ShortcutProps{}, g.Text("⇧⌘Z"))),
						Separator(SeparatorProps{}),
						SubMenu(
							SubMenuProps{},
							SubTrigger(SubTriggerProps{}, g.Text("Find")),
							SubContent(
								SubContentProps{},
								Item(ItemProps{}, g.Text("Search the web")),
								Separator(SeparatorProps{}),
								Item(ItemProps{}, g.Text("Find...")),
								Item(ItemProps{}, g.Text("Find Next")),
								Item(ItemProps{}, g.Text("Find Previous")),
							),
						),
						Separator(SeparatorProps{}),
						Item(ItemProps{}, g.Text("Cut")),
						Item(ItemProps{}, g.Text("Copy")),
						Item(ItemProps{}, g.Text("Paste")),
					),
				),
				Menu(
					MenuProps{},
					Trigger(TriggerProps{}, g.Text("View")),
					ContentComponent(
						ContentProps{},
						CheckboxItem(CheckboxItemProps{Checked: true}, g.Text("Always Show Bookmarks Bar")),
						CheckboxItem(CheckboxItemProps{}, g.Text("Always Show Full URLs")),
						Separator(SeparatorProps{}),
						Item(ItemProps{Inset: true}, g.Text("Reload"), Shortcut(ShortcutProps{}, g.Text("⌘R"))),
						Item(ItemProps{Inset: true, Disabled: true}, g.Text("Force Reload"), Shortcut(ShortcutProps{}, g.Text("⇧⌘R"))),
						Separator(SeparatorProps{}),
						Item(ItemProps{Inset: true}, g.Text("Toggle Fullscreen")),
						Separator(SeparatorProps{}),
						Item(ItemProps{Inset: true}, g.Text("Hide Sidebar")),
					),
				),
				Menu(
					MenuProps{},
					Trigger(TriggerProps{}, g.Text("Profiles")),
					ContentComponent(
						ContentProps{},
						RadioGroup(
							RadioGroupProps{Value: "benoit"},
							RadioItem(RadioItemProps{Value: "andy"}, g.Text("Andy")),
							RadioItem(RadioItemProps{Value: "benoit"}, g.Text("Benoit")),
							RadioItem(RadioItemProps{Value: "Luis"}, g.Text("Luis")),
						),
						Separator(SeparatorProps{}),
						Item(ItemProps{Inset: true}, g.Text("Edit...")),
						Separator(SeparatorProps{}),
						Item(ItemProps{Inset: true}, g.Text("Add Profile...")),
					),
				),
			),
		),

		// Application Menubar Example
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Application Menubar")),
			New(
				Props{},
				Menu(
					MenuProps{},
					Trigger(TriggerProps{}, g.Text("Application")),
					ContentComponent(
						ContentProps{},
						Item(ItemProps{}, g.Text("About Application")),
						Separator(SeparatorProps{}),
						Item(ItemProps{}, g.Text("Preferences..."), Shortcut(ShortcutProps{}, g.Text("⌘,"))),
						Separator(SeparatorProps{}),
						Item(ItemProps{}, g.Text("Services")),
						Separator(SeparatorProps{}),
						Item(ItemProps{}, g.Text("Hide Application"), Shortcut(ShortcutProps{}, g.Text("⌘H"))),
						Item(ItemProps{}, g.Text("Hide Others"), Shortcut(ShortcutProps{}, g.Text("⌘⌥H"))),
						Item(ItemProps{}, g.Text("Show All")),
						Separator(SeparatorProps{}),
						Item(ItemProps{}, g.Text("Quit Application"), Shortcut(ShortcutProps{}, g.Text("⌘Q"))),
					),
				),
				Menu(
					MenuProps{},
					Trigger(TriggerProps{}, g.Text("File")),
					ContentComponent(
						ContentProps{},
						Item(ItemProps{}, g.Text("New"), Shortcut(ShortcutProps{}, g.Text("⌘N"))),
						Item(ItemProps{}, g.Text("Open..."), Shortcut(ShortcutProps{}, g.Text("⌘O"))),
						SubMenu(
							SubMenuProps{},
							SubTrigger(SubTriggerProps{}, g.Text("Open Recent")),
							SubContent(
								SubContentProps{},
								Item(ItemProps{}, g.Text("Document 1.txt")),
								Item(ItemProps{}, g.Text("Document 2.txt")),
								Item(ItemProps{}, g.Text("Document 3.txt")),
								Separator(SeparatorProps{}),
								Item(ItemProps{}, g.Text("Clear Recent")),
							),
						),
						Separator(SeparatorProps{}),
						Item(ItemProps{}, g.Text("Save"), Shortcut(ShortcutProps{}, g.Text("⌘S"))),
						Item(ItemProps{}, g.Text("Save As..."), Shortcut(ShortcutProps{}, g.Text("⇧⌘S"))),
						Separator(SeparatorProps{}),
						Item(ItemProps{}, g.Text("Close"), Shortcut(ShortcutProps{}, g.Text("⌘W"))),
					),
				),
				Menu(
					MenuProps{},
					Trigger(TriggerProps{}, g.Text("Edit")),
					ContentComponent(
						ContentProps{},
						Item(ItemProps{Disabled: true}, g.Text("Undo"), Shortcut(ShortcutProps{}, g.Text("⌘Z"))),
						Item(ItemProps{Disabled: true}, g.Text("Redo"), Shortcut(ShortcutProps{}, g.Text("⇧⌘Z"))),
						Separator(SeparatorProps{}),
						Item(ItemProps{}, g.Text("Cut"), Shortcut(ShortcutProps{}, g.Text("⌘X"))),
						Item(ItemProps{}, g.Text("Copy"), Shortcut(ShortcutProps{}, g.Text("⌘C"))),
						Item(ItemProps{}, g.Text("Paste"), Shortcut(ShortcutProps{}, g.Text("⌘V"))),
						Item(ItemProps{}, g.Text("Paste and Match Style"), Shortcut(ShortcutProps{}, g.Text("⌥⇧⌘V"))),
						Separator(SeparatorProps{}),
						Item(ItemProps{}, g.Text("Select All"), Shortcut(ShortcutProps{}, g.Text("⌘A"))),
					),
				),
				Menu(
					MenuProps{},
					Trigger(TriggerProps{}, g.Text("View")),
					ContentComponent(
						ContentProps{},
						html.Label(g.Text("Appearance")),
						CheckboxItem(CheckboxItemProps{Checked: true}, g.Text("Show Toolbar")),
						CheckboxItem(CheckboxItemProps{Checked: true}, g.Text("Show Status Bar")),
						CheckboxItem(CheckboxItemProps{}, g.Text("Show Path Bar")),
						Separator(SeparatorProps{}),
						html.Label(g.Text("Layout")),
						RadioGroup(
							RadioGroupProps{Value: "icons"},
							RadioItem(RadioItemProps{Value: "icons"}, g.Text("as Icons")),
							RadioItem(RadioItemProps{Value: "list"}, g.Text("as List")),
							RadioItem(RadioItemProps{Value: "columns"}, g.Text("as Columns")),
							RadioItem(RadioItemProps{Value: "gallery"}, g.Text("as Gallery")),
						),
						Separator(SeparatorProps{}),
						Item(ItemProps{}, g.Text("Enter Full Screen"), Shortcut(ShortcutProps{}, g.Text("⌃⌘F"))),
					),
				),
				Menu(
					MenuProps{},
					Trigger(TriggerProps{}, g.Text("Window")),
					ContentComponent(
						ContentProps{},
						Item(ItemProps{}, g.Text("Minimize"), Shortcut(ShortcutProps{}, g.Text("⌘M"))),
						Item(ItemProps{}, g.Text("Zoom")),
						Separator(SeparatorProps{}),
						Item(ItemProps{}, g.Text("Bring All to Front")),
						Separator(SeparatorProps{}),
						CheckboxItem(CheckboxItemProps{Checked: true}, g.Text("Welcome")),
						CheckboxItem(CheckboxItemProps{}, g.Text("Document 1")),
						CheckboxItem(CheckboxItemProps{}, g.Text("Document 2")),
					),
				),
				Menu(
					MenuProps{},
					Trigger(TriggerProps{}, g.Text("Help")),
					ContentComponent(
						ContentProps{},
						Item(ItemProps{}, g.Text("Application Help")),
						Separator(SeparatorProps{}),
						Item(ItemProps{}, g.Text("Keyboard Shortcuts")),
						Item(ItemProps{}, g.Text("Report Issue...")),
						Separator(SeparatorProps{}),
						Item(ItemProps{}, g.Text("View License")),
						Item(ItemProps{}, g.Text("Privacy Policy")),
					),
				),
			),
		),

		// Disabled Menu Items
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Disabled States")),
			New(
				Props{},
				Menu(
					MenuProps{},
					Trigger(TriggerProps{}, g.Text("Actions")),
					ContentComponent(
						ContentProps{},
						Item(ItemProps{}, g.Text("Action 1")),
						Item(ItemProps{Disabled: true}, g.Text("Action 2 (Disabled)")),
						Item(ItemProps{}, g.Text("Action 3")),
						Separator(SeparatorProps{}),
						SubMenu(
							SubMenuProps{},
							SubTrigger(SubTriggerProps{Disabled: true}, g.Text("More Actions (Disabled)")),
							SubContent(
								SubContentProps{},
								Item(ItemProps{}, g.Text("Sub Action 1")),
								Item(ItemProps{}, g.Text("Sub Action 2")),
							),
						),
					),
				),
				Menu(
					MenuProps{},
					Trigger(TriggerProps{Disabled: true}, g.Text("Disabled Menu")),
					ContentComponent(
						ContentProps{},
						Item(ItemProps{}, g.Text("This menu is disabled")),
					),
				),
			),
		),

		// Complex Nested Menus
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Nested Menus")),
			New(
				Props{},
				Menu(
					MenuProps{},
					Trigger(TriggerProps{}, g.Text("Format")),
					ContentComponent(
						ContentProps{},
						SubMenu(
							SubMenuProps{},
							SubTrigger(SubTriggerProps{}, g.Text("Font")),
							SubContent(
								SubContentProps{},
								Item(ItemProps{}, g.Text("Bold"), Shortcut(ShortcutProps{}, g.Text("⌘B"))),
								Item(ItemProps{}, g.Text("Italic"), Shortcut(ShortcutProps{}, g.Text("⌘I"))),
								Item(ItemProps{}, g.Text("Underline"), Shortcut(ShortcutProps{}, g.Text("⌘U"))),
								Separator(SeparatorProps{}),
								SubMenu(
									SubMenuProps{},
									SubTrigger(SubTriggerProps{}, g.Text("More")),
									SubContent(
										SubContentProps{},
										Item(ItemProps{}, g.Text("Strikethrough")),
										Item(ItemProps{}, g.Text("Superscript")),
										Item(ItemProps{}, g.Text("Subscript")),
									),
								),
							),
						),
						SubMenu(
							SubMenuProps{},
							SubTrigger(SubTriggerProps{}, g.Text("Paragraph")),
							SubContent(
								SubContentProps{},
								RadioGroup(
									RadioGroupProps{Value: "left"},
									RadioItem(RadioItemProps{Value: "left"}, g.Text("Align Left")),
									RadioItem(RadioItemProps{Value: "center"}, g.Text("Center")),
									RadioItem(RadioItemProps{Value: "right"}, g.Text("Align Right")),
									RadioItem(RadioItemProps{Value: "justify"}, g.Text("Justify")),
								),
								Separator(SeparatorProps{}),
								CheckboxItem(CheckboxItemProps{}, g.Text("Add Indent")),
								CheckboxItem(CheckboxItemProps{}, g.Text("Remove Indent")),
							),
						),
					),
				),
			),
		),

		// With Icons
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("With Icons")),
			New(
				Props{},
				Menu(
					MenuProps{},
					Trigger(TriggerProps{}, g.Text("Tools")),
					ContentComponent(
						ContentProps{},
						Item(
							ItemProps{},
							html.Div(html.Class("flex items-center gap-2"),
								g.Raw(`<svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg">
									<path d="M3 2.5C3 2.22386 3.22386 2 3.5 2H11.5C11.7761 2 12 2.22386 12 2.5V12.5C12 12.7761 11.7761 13 11.5 13H3.5C3.22386 13 3 12.7761 3 12.5V2.5ZM4 3V12H11V3H4Z" fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"></path>
								</svg>`),
								g.Text("Copy"),
							),
						),
						Item(
							ItemProps{},
							html.Div(html.Class("flex items-center gap-2"),
								g.Raw(`<svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg">
									<path d="M1 2C0.447715 2 0 2.44772 0 3V12C0 12.5523 0.447715 13 1 13H14C14.5523 13 15 12.5523 15 12V3C15 2.44772 14.5523 2 14 2H1ZM1 3H14V12H1V3Z" fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"></path>
								</svg>`),
								g.Text("Paste"),
							),
						),
						Separator(SeparatorProps{}),
						Item(
							ItemProps{},
							html.Div(html.Class("flex items-center gap-2"),
								g.Raw(`<svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg">
									<path d="M5 2V1H10V2H5ZM4.5 0C4.22386 0 4 0.223858 4 0.5V2H2.5C2.22386 2 2 2.22386 2 2.5C2 2.77614 2.22386 3 2.5 3H3V12C3 12.5523 3.44772 13 4 13H11C11.5523 13 12 12.5523 12 12V3H12.5C12.7761 3 13 2.77614 13 2.5C13 2.22386 12.7761 2 12.5 2H11V0.5C11 0.223858 10.7761 0 10.5 0H4.5ZM5 3H10V12H5V3Z" fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"></path>
								</svg>`),
								g.Text("Delete"),
							),
						),
					),
				),
			),
		),

		// Custom Styling
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Custom Styling")),
			New(
				Props{Class: "bg-primary text-primary-foreground"},
				Menu(
					MenuProps{},
					Trigger(TriggerProps{Class: "text-primary-foreground hover:bg-primary/90"}, g.Text("Custom")),
					ContentComponent(
						ContentProps{},
						Item(ItemProps{}, g.Text("Option 1")),
						Item(ItemProps{}, g.Text("Option 2")),
						Item(ItemProps{}, g.Text("Option 3")),
					),
				),
			),
		),
	)
}