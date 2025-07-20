package dropdownmenu

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/components/button"
)

// DemoBasic shows a basic dropdown menu
func DemoBasic() g.Node {
	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Basic Dropdown Menu")),
		html.P(html.Class("text-sm text-muted-foreground mb-4"), 
			g.Text("A standard dropdown menu with various item types and shortcuts."),
		),
		Example(),
	)
}

// DemoCheckboxes shows a dropdown with checkbox items
func DemoCheckboxes() g.Node {
	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Dropdown with Checkboxes")),
		html.P(html.Class("text-sm text-muted-foreground mb-4"), 
			g.Text("Dropdown menu with checkbox items for selecting multiple options."),
		),
		ExampleCheckboxes(),
	)
}

// DemoRadioGroup shows a dropdown with radio items
func DemoRadioGroup() g.Node {
	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Dropdown with Radio Group")),
		html.P(html.Class("text-sm text-muted-foreground mb-4"), 
			g.Text("Dropdown menu with radio items for selecting a single option from a group."),
		),
		ExampleRadioGroup(),
	)
}

// DemoWithSubmenu shows a dropdown with submenu
func DemoWithSubmenu() g.Node {
	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Dropdown with Submenu")),
		html.P(html.Class("text-sm text-muted-foreground mb-4"), 
			g.Text("Dropdown menu with nested submenus. Note: Submenus require JavaScript for full functionality."),
		),
		ExampleWithSubmenu(),
	)
}

// DemoWithCustomTrigger shows a dropdown with custom trigger button
func DemoWithCustomTrigger() g.Node {
	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Custom Trigger")),
		html.P(html.Class("text-sm text-muted-foreground mb-4"), 
			g.Text("Dropdown menu with a custom styled trigger button."),
		),
		New(
			Props{Open: true},
			Trigger(
				TriggerProps{AsChild: true},
				button.New(
					button.Props{
						Variant: "secondary",
						Size:    "sm",
					},
					g.Text("Options"),
				),
			),
			DropdownContent(
				ContentProps{Class: "w-48"},
				Item(ItemProps{}, g.Text("Edit")),
				Item(ItemProps{}, g.Text("Duplicate")),
				Item(ItemProps{}, g.Text("Archive")),
				Separator(SeparatorProps{}),
				Item(ItemProps{Class: "text-destructive"}, g.Text("Delete")),
			),
		),
	)
}

// DemoAlignment shows dropdown menu alignment options
func DemoAlignment() g.Node {
	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Alignment Options")),
		html.P(html.Class("text-sm text-muted-foreground mb-4"), 
			g.Text("Dropdown menus can be aligned to different sides and positions."),
		),
		html.Div(html.Class("flex gap-4 flex-wrap"),
			// Start aligned
			html.Div(html.Class("space-y-2"),
				html.P(html.Class("text-sm font-medium"), g.Text("Align Start")),
				New(
					Props{Open: true},
					Trigger(TriggerProps{}, g.Text("Open")),
					DropdownContent(
						ContentProps{
							Class: "w-48",
							Align: "start",
						},
						DropdownLabel(LabelProps{}, "Start Aligned"),
						Item(ItemProps{}, g.Text("Item 1")),
						Item(ItemProps{}, g.Text("Item 2")),
					),
				),
			),
			// Center aligned
			html.Div(html.Class("space-y-2"),
				html.P(html.Class("text-sm font-medium"), g.Text("Align Center")),
				New(
					Props{Open: true},
					Trigger(TriggerProps{}, g.Text("Open")),
					DropdownContent(
						ContentProps{
							Class: "w-48",
							Align: "center",
						},
						DropdownLabel(LabelProps{}, "Center Aligned"),
						Item(ItemProps{}, g.Text("Item 1")),
						Item(ItemProps{}, g.Text("Item 2")),
					),
				),
			),
			// End aligned
			html.Div(html.Class("space-y-2"),
				html.P(html.Class("text-sm font-medium"), g.Text("Align End")),
				New(
					Props{Open: true},
					Trigger(TriggerProps{}, g.Text("Open")),
					DropdownContent(
						ContentProps{
							Class: "w-48",
							Align: "end",
						},
						DropdownLabel(LabelProps{}, "End Aligned"),
						Item(ItemProps{}, g.Text("Item 1")),
						Item(ItemProps{}, g.Text("Item 2")),
					),
				),
			),
		),
	)
}

// DemoStates shows different dropdown menu states
func DemoStates() g.Node {
	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Menu Item States")),
		html.P(html.Class("text-sm text-muted-foreground mb-4"), 
			g.Text("Different states for dropdown menu items."),
		),
		New(
			Props{Open: true},
			Trigger(TriggerProps{}, g.Text("View States")),
			DropdownContent(
				ContentProps{Class: "w-56"},
				DropdownLabel(LabelProps{}, "Item States"),
				Separator(SeparatorProps{}),
				Item(ItemProps{}, g.Text("Normal Item")),
				Item(ItemProps{Class: "bg-accent text-accent-foreground"}, g.Text("Focused/Hover Item")),
				Item(ItemProps{Disabled: true}, g.Text("Disabled Item")),
				Separator(SeparatorProps{}),
				CheckboxItem(CheckboxItemProps{Checked: false}, g.Text("Unchecked")),
				CheckboxItem(CheckboxItemProps{Checked: true}, g.Text("Checked")),
				CheckboxItem(CheckboxItemProps{Checked: false, Disabled: true}, g.Text("Disabled Unchecked")),
				CheckboxItem(CheckboxItemProps{Checked: true, Disabled: true}, g.Text("Disabled Checked")),
			),
		),
	)
}

// DemoComplex shows a complex dropdown menu
func DemoComplex() g.Node {
	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Complex Dropdown Menu")),
		html.P(html.Class("text-sm text-muted-foreground mb-4"), 
			g.Text("A comprehensive example with all dropdown menu features."),
		),
		New(
			Props{Open: true},
			Trigger(
				TriggerProps{AsChild: true},
				button.New(
					button.Props{Variant: "outline"},
					g.Text("Account Settings"),
				),
			),
			DropdownContent(
				ContentProps{Class: "w-64"},
				// User info section
				html.Div(html.Class("px-2 py-1.5"),
					html.P(html.Class("text-sm font-medium"), g.Text("john.doe@example.com")),
					html.P(html.Class("text-xs text-muted-foreground"), g.Text("Free Plan · 2GB used")),
				),
				Separator(SeparatorProps{}),
				
				// Main actions
				Group(
					GroupProps{},
					DropdownLabel(LabelProps{}, "Account"),
					Item(ItemProps{}, g.Text("Profile Settings")),
					Item(ItemProps{}, g.Text("Billing")),
					Item(ItemProps{}, g.Text("Notifications")),
				),
				Separator(SeparatorProps{}),
				
				// View options with checkboxes
				Group(
					GroupProps{},
					DropdownLabel(LabelProps{}, "View Options"),
					CheckboxItem(CheckboxItemProps{Checked: true}, g.Text("Show Sidebar")),
					CheckboxItem(CheckboxItemProps{Checked: true}, g.Text("Show Toolbar")),
					CheckboxItem(CheckboxItemProps{Checked: false}, g.Text("Compact Mode")),
				),
				Separator(SeparatorProps{}),
				
				// Theme selection with radio items
				Group(
					GroupProps{},
					DropdownLabel(LabelProps{}, "Theme"),
					RadioGroup(
						RadioGroupProps{Value: "system"},
						RadioItemWithSelection(RadioItemProps{Value: "light"}, false, g.Text("Light")),
						RadioItemWithSelection(RadioItemProps{Value: "dark"}, false, g.Text("Dark")),
						RadioItemWithSelection(RadioItemProps{Value: "system"}, true, g.Text("System")),
					),
				),
				Separator(SeparatorProps{}),
				
				// Danger zone
				Group(
					GroupProps{},
					DropdownLabel(LabelProps{Class: "text-destructive"}, "Danger Zone"),
					Item(ItemProps{Class: "text-destructive focus:bg-destructive/10 focus:text-destructive"}, 
						g.Text("Delete Account"),
					),
				),
			),
		),
	)
}