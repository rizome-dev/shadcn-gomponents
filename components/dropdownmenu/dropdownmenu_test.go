package dropdownmenu_test

import (
	"bytes"
	"strings"
	"testing"

	g "maragu.dev/gomponents"
	"github.com/rizome-dev/shadcn-gomponents/components/dropdownmenu"
)

func TestDropdownMenu(t *testing.T) {
	t.Run("renders container", func(t *testing.T) {
		menu := dropdownmenu.New(
			dropdownmenu.Props{},
			g.Text("Test content"),
		)

		var buf bytes.Buffer
		err := menu.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "relative inline-block text-left") {
			t.Error("expected relative positioning")
		}
		if !strings.Contains(output, "Test content") {
			t.Error("expected content")
		}
	})

	t.Run("renders with custom class", func(t *testing.T) {
		menu := dropdownmenu.New(
			dropdownmenu.Props{Class: "custom-dropdown"},
			g.Text("Test"),
		)

		var buf bytes.Buffer
		err := menu.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "custom-dropdown") {
			t.Error("expected custom class")
		}
	})
}

func TestTrigger(t *testing.T) {
	t.Run("renders trigger button", func(t *testing.T) {
		trigger := dropdownmenu.Trigger(
			dropdownmenu.TriggerProps{},
			g.Text("Open Menu"),
		)

		var buf bytes.Buffer
		err := trigger.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "<button") {
			t.Error("expected button element")
		}
		if !strings.Contains(output, `aria-haspopup="menu"`) {
			t.Error("expected aria-haspopup")
		}
		if !strings.Contains(output, `aria-expanded="false"`) {
			t.Error("expected aria-expanded")
		}
		if !strings.Contains(output, "Open Menu") {
			t.Error("expected trigger text")
		}
	})

	t.Run("renders disabled trigger", func(t *testing.T) {
		trigger := dropdownmenu.Trigger(
			dropdownmenu.TriggerProps{Disabled: true},
			g.Text("Disabled"),
		)

		var buf bytes.Buffer
		err := trigger.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "disabled") {
			t.Error("expected disabled attribute")
		}
	})

	t.Run("renders as child element", func(t *testing.T) {
		trigger := dropdownmenu.Trigger(
			dropdownmenu.TriggerProps{AsChild: true},
			g.El("span", g.Text("Custom Element")),
		)

		var buf bytes.Buffer
		err := trigger.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if strings.Contains(output, "<button") {
			t.Error("should not have button wrapper when AsChild is true")
		}
		if !strings.Contains(output, "<span") {
			t.Error("expected span element")
		}
		if !strings.Contains(output, "Custom Element") {
			t.Error("expected custom element text")
		}
	})
}

func TestContent(t *testing.T) {
	t.Run("renders content container", func(t *testing.T) {
		content := dropdownmenu.DropdownContent(
			dropdownmenu.ContentProps{},
			g.Text("Menu content"),
		)

		var buf bytes.Buffer
		err := content.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, `role="menu"`) {
			t.Error("expected menu role")
		}
		if !strings.Contains(output, `aria-orientation="vertical"`) {
			t.Error("expected vertical orientation")
		}
		if !strings.Contains(output, "Menu content") {
			t.Error("expected content")
		}
		if !strings.Contains(output, "z-50") {
			t.Error("expected z-index")
		}
		if !strings.Contains(output, "min-w-[8rem]") {
			t.Error("expected minimum width")
		}
	})
}

func TestItem(t *testing.T) {
	t.Run("renders menu item", func(t *testing.T) {
		item := dropdownmenu.Item(
			dropdownmenu.ItemProps{},
			g.Text("Menu Item"),
		)

		var buf bytes.Buffer
		err := item.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, `role="menuitem"`) {
			t.Error("expected menuitem role")
		}
		if !strings.Contains(output, `tabindex="-1"`) {
			t.Error("expected tabindex")
		}
		if !strings.Contains(output, "Menu Item") {
			t.Error("expected item text")
		}
	})

	t.Run("renders disabled item", func(t *testing.T) {
		item := dropdownmenu.Item(
			dropdownmenu.ItemProps{Disabled: true},
			g.Text("Disabled Item"),
		)

		var buf bytes.Buffer
		err := item.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, `aria-disabled="true"`) {
			t.Error("expected aria-disabled")
		}
		if !strings.Contains(output, `data-disabled`) {
			t.Error("expected data-disabled")
		}
	})

	t.Run("renders inset item", func(t *testing.T) {
		item := dropdownmenu.Item(
			dropdownmenu.ItemProps{Inset: true},
			g.Text("Inset Item"),
		)

		var buf bytes.Buffer
		err := item.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "pl-8") {
			t.Error("expected inset padding")
		}
	})
}

func TestCheckboxItem(t *testing.T) {
	t.Run("renders unchecked checkbox item", func(t *testing.T) {
		item := dropdownmenu.CheckboxItem(
			dropdownmenu.CheckboxItemProps{Checked: false},
			g.Text("Checkbox Item"),
		)

		var buf bytes.Buffer
		err := item.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, `role="menuitemcheckbox"`) {
			t.Error("expected menuitemcheckbox role")
		}
		if !strings.Contains(output, `aria-checked="false"`) {
			t.Error("expected aria-checked false")
		}
		if !strings.Contains(output, "Checkbox Item") {
			t.Error("expected item text")
		}
	})

	t.Run("renders checked checkbox item", func(t *testing.T) {
		item := dropdownmenu.CheckboxItem(
			dropdownmenu.CheckboxItemProps{Checked: true},
			g.Text("Checked Item"),
		)

		var buf bytes.Buffer
		err := item.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, `aria-checked="true"`) {
			t.Error("expected aria-checked true")
		}
		// Should have check icon
		if !strings.Contains(output, "svg") || !strings.Contains(output, "h-4 w-4") {
			t.Error("expected check icon")
		}
	})

	t.Run("renders disabled checkbox item", func(t *testing.T) {
		item := dropdownmenu.CheckboxItem(
			dropdownmenu.CheckboxItemProps{Disabled: true},
			g.Text("Disabled"),
		)

		var buf bytes.Buffer
		err := item.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, `aria-disabled="true"`) {
			t.Error("expected aria-disabled")
		}
	})
}

func TestRadioItem(t *testing.T) {
	t.Run("renders radio item", func(t *testing.T) {
		item := dropdownmenu.RadioItem(
			dropdownmenu.RadioItemProps{Value: "option1"},
			g.Text("Option 1"),
		)

		var buf bytes.Buffer
		err := item.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, `role="menuitemradio"`) {
			t.Error("expected menuitemradio role")
		}
		if !strings.Contains(output, `data-value="option1"`) {
			t.Error("expected data-value")
		}
		if !strings.Contains(output, "Option 1") {
			t.Error("expected item text")
		}
	})

	t.Run("renders selected radio item", func(t *testing.T) {
		item := dropdownmenu.RadioItemWithSelection(
			dropdownmenu.RadioItemProps{Value: "selected"},
			true,
			g.Text("Selected"),
		)

		var buf bytes.Buffer
		err := item.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, `aria-checked="true"`) {
			t.Error("expected aria-checked true")
		}
		// Should have circle icon
		if !strings.Contains(output, "svg") {
			t.Error("expected circle icon")
		}
	})
}

func TestLabel(t *testing.T) {
	t.Run("renders label", func(t *testing.T) {
		label := dropdownmenu.DropdownLabel(
			dropdownmenu.LabelProps{},
			"Section Label",
		)

		var buf bytes.Buffer
		err := label.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "Section Label") {
			t.Error("expected label text")
		}
		if !strings.Contains(output, "font-semibold") {
			t.Error("expected font-semibold")
		}
		if !strings.Contains(output, "px-2 py-1.5") {
			t.Error("expected padding")
		}
	})

	t.Run("renders inset label", func(t *testing.T) {
		label := dropdownmenu.DropdownLabel(
			dropdownmenu.LabelProps{Inset: true},
			"Inset Label",
		)

		var buf bytes.Buffer
		err := label.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "pl-8") {
			t.Error("expected inset padding")
		}
	})
}

func TestSeparator(t *testing.T) {
	t.Run("renders separator", func(t *testing.T) {
		separator := dropdownmenu.Separator(dropdownmenu.SeparatorProps{})

		var buf bytes.Buffer
		err := separator.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, `role="separator"`) {
			t.Error("expected separator role")
		}
		if !strings.Contains(output, "h-px") {
			t.Error("expected height")
		}
		if !strings.Contains(output, "bg-muted") {
			t.Error("expected background color")
		}
		if !strings.Contains(output, "-mx-1 my-1") {
			t.Error("expected margin")
		}
	})
}

func TestShortcut(t *testing.T) {
	t.Run("renders shortcut", func(t *testing.T) {
		shortcut := dropdownmenu.Shortcut(
			dropdownmenu.ShortcutProps{},
			"⌘K",
		)

		var buf bytes.Buffer
		err := shortcut.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "⌘K") {
			t.Error("expected shortcut text")
		}
		if !strings.Contains(output, "ml-auto") {
			t.Error("expected margin-left auto")
		}
		if !strings.Contains(output, "text-xs") {
			t.Error("expected small text")
		}
		if !strings.Contains(output, "opacity-60") {
			t.Error("expected reduced opacity")
		}
	})
}

func TestGroup(t *testing.T) {
	t.Run("renders group", func(t *testing.T) {
		group := dropdownmenu.Group(
			dropdownmenu.GroupProps{},
			g.Text("Group content"),
		)

		var buf bytes.Buffer
		err := group.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, `role="group"`) {
			t.Error("expected group role")
		}
		if !strings.Contains(output, "Group content") {
			t.Error("expected group content")
		}
	})
}

func TestSubMenu(t *testing.T) {
	t.Run("renders submenu container", func(t *testing.T) {
		sub := dropdownmenu.DropdownSub(
			dropdownmenu.SubProps{},
			g.Text("Submenu"),
		)

		var buf bytes.Buffer
		err := sub.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "relative") {
			t.Error("expected relative positioning")
		}
		if !strings.Contains(output, "Submenu") {
			t.Error("expected submenu content")
		}
	})

	t.Run("renders submenu trigger", func(t *testing.T) {
		trigger := dropdownmenu.SubTrigger(
			dropdownmenu.SubTriggerProps{},
			g.Text("More options"),
		)

		var buf bytes.Buffer
		err := trigger.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, `aria-haspopup="menu"`) {
			t.Error("expected aria-haspopup")
		}
		if !strings.Contains(output, `aria-expanded="false"`) {
			t.Error("expected aria-expanded")
		}
		if !strings.Contains(output, "More options") {
			t.Error("expected trigger text")
		}
		// Should have chevron icon
		if !strings.Contains(output, "svg") || !strings.Contains(output, "ml-auto") {
			t.Error("expected chevron icon")
		}
	})

	t.Run("renders submenu content", func(t *testing.T) {
		content := dropdownmenu.SubContent(
			dropdownmenu.SubContentProps{},
			g.Text("Submenu items"),
		)

		var buf bytes.Buffer
		err := content.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, `role="menu"`) {
			t.Error("expected menu role")
		}
		if !strings.Contains(output, "z-50") {
			t.Error("expected z-index")
		}
		if !strings.Contains(output, "Submenu items") {
			t.Error("expected submenu content")
		}
	})
}

func TestRadioGroup(t *testing.T) {
	t.Run("renders radio group", func(t *testing.T) {
		group := dropdownmenu.RadioGroup(
			dropdownmenu.RadioGroupProps{Value: "option1"},
			g.Text("Radio items"),
		)

		var buf bytes.Buffer
		err := group.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, `role="radiogroup"`) {
			t.Error("expected radiogroup role")
		}
		if !strings.Contains(output, "Radio items") {
			t.Error("expected radio items")
		}
	})
}

func TestExample(t *testing.T) {
	t.Run("renders basic example", func(t *testing.T) {
		example := dropdownmenu.Example()

		var buf bytes.Buffer
		err := example.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "My Account") {
			t.Error("expected account label")
		}
		if !strings.Contains(output, "Profile") {
			t.Error("expected profile item")
		}
		if !strings.Contains(output, "Settings") {
			t.Error("expected settings item")
		}
		if !strings.Contains(output, "Log out") {
			t.Error("expected log out item")
		}
		if !strings.Contains(output, "⇧⌘P") {
			t.Error("expected keyboard shortcut")
		}
		// Check for disabled API item
		if !strings.Contains(output, "API") || !strings.Contains(output, `aria-disabled="true"`) {
			t.Error("expected disabled API item")
		}
	})
}

func TestExampleCheckboxes(t *testing.T) {
	t.Run("renders checkboxes example", func(t *testing.T) {
		example := dropdownmenu.ExampleCheckboxes()

		var buf bytes.Buffer
		err := example.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "Appearance") {
			t.Error("expected appearance label")
		}
		if !strings.Contains(output, "Status Bar") {
			t.Error("expected status bar item")
		}
		if !strings.Contains(output, "Activity Bar") {
			t.Error("expected activity bar item")
		}
		if !strings.Contains(output, "Panel") {
			t.Error("expected panel item")
		}
		// Status bar should be checked
		if !strings.Contains(output, `aria-checked="true"`) {
			t.Error("expected checked item")
		}
	})
}

func TestExampleRadioGroup(t *testing.T) {
	t.Run("renders radio group example", func(t *testing.T) {
		example := dropdownmenu.ExampleRadioGroup()

		var buf bytes.Buffer
		err := example.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "Panel Position") {
			t.Error("expected panel position label")
		}
		if !strings.Contains(output, "Top") {
			t.Error("expected top option")
		}
		if !strings.Contains(output, "Bottom") {
			t.Error("expected bottom option")
		}
		if !strings.Contains(output, "Right") {
			t.Error("expected right option")
		}
		// Bottom should be selected
		if !strings.Contains(output, `data-value="bottom"`) {
			t.Error("expected bottom value")
		}
	})
}

func TestExampleWithSubmenu(t *testing.T) {
	t.Run("renders submenu example", func(t *testing.T) {
		example := dropdownmenu.ExampleWithSubmenu()

		var buf bytes.Buffer
		err := example.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "Team") {
			t.Error("expected team label")
		}
		if !strings.Contains(output, "Invite users") {
			t.Error("expected invite users submenu trigger")
		}
		if !strings.Contains(output, "New Team") {
			t.Error("expected new team item")
		}
		if !strings.Contains(output, "⌘+T") {
			t.Error("expected new team shortcut")
		}
	})
}