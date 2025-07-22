package contextmenu

import (
	"strings"
	"testing"

	g "maragu.dev/gomponents"
)

func renderToString(node g.Node) string {
	var buf strings.Builder
	err := node.Render(&buf)
	if err != nil {
		panic(err) // For tests, panic on render error
	}
	return buf.String()
}

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		props    Props
		children []g.Node
		want     string
	}{
		{
			name:  "basic context menu",
			props: Props{ID: "test-menu"},
			children: []g.Node{
				Trigger(TriggerProps{}, g.Text("Right click")),
			},
			want: `<div id="test-menu" data-context-menu="root"><div class="cursor-context-menu" data-context-menu="trigger" oncontextmenu="return false;">Right click</div></div>`,
		},
		{
			name:  "with custom class",
			props: Props{ID: "test", Class: "custom-class"},
			children: []g.Node{
				Trigger(TriggerProps{}, g.Text("Test")),
			},
			want: `<div id="test" class="custom-class" data-context-menu="root">`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := renderToString(New(tt.props, tt.children...))
			if !strings.Contains(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrigger(t *testing.T) {
	tests := []struct {
		name  string
		props TriggerProps
		want  string
	}{
		{
			name:  "basic trigger",
			props: TriggerProps{},
			want:  `<div class="cursor-context-menu" data-context-menu="trigger" oncontextmenu="return false;">`,
		},
		{
			name:  "with custom class",
			props: TriggerProps{Class: "custom-trigger"},
			want:  `<div class="cursor-context-menu custom-trigger" data-context-menu="trigger"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := renderToString(Trigger(tt.props, g.Text("Content")))
			if !strings.Contains(got, tt.want) {
				t.Errorf("Trigger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContent(t *testing.T) {
	tests := []struct {
		name  string
		props ContentProps
		want  string
	}{
		{
			name:  "basic content",
			props: ContentProps{},
			want:  `data-context-menu="content" data-state="closed"`,
		},
		{
			name:  "with position",
			props: ContentProps{Position: "bottom"},
			want:  `data-side="bottom"`,
		},
		{
			name:  "with align",
			props: ContentProps{Align: "start"},
			want:  `data-align="start"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := renderToString(ContentComponent(tt.props, g.Text("Menu content")))
			if !strings.Contains(got, tt.want) {
				t.Errorf("ContentComponent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItem(t *testing.T) {
	tests := []struct {
		name  string
		props ItemProps
		want  string
	}{
		{
			name:  "basic item",
			props: ItemProps{},
			want:  `data-context-menu="item" tabindex="-1"`,
		},
		{
			name:  "disabled item",
			props: ItemProps{Disabled: true},
			want:  `data-disabled="true"`,
		},
		{
			name:  "inset item",
			props: ItemProps{Inset: true},
			want:  `pl-8`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := renderToString(Item(tt.props, g.Text("Menu item")))
			if !strings.Contains(got, tt.want) {
				t.Errorf("Item() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckboxItem(t *testing.T) {
	tests := []struct {
		name  string
		props CheckboxItemProps
		want  string
	}{
		{
			name:  "unchecked",
			props: CheckboxItemProps{Checked: false},
			want:  `data-state="unchecked" aria-checked="false"`,
		},
		{
			name:  "checked",
			props: CheckboxItemProps{Checked: true},
			want:  `data-state="checked" aria-checked="true"`,
		},
		{
			name:  "disabled",
			props: CheckboxItemProps{Disabled: true},
			want:  `data-disabled="true"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := renderToString(CheckboxItem(tt.props, g.Text("Checkbox")))
			if !strings.Contains(got, tt.want) {
				t.Errorf("CheckboxItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRadioItem(t *testing.T) {
	tests := []struct {
		name  string
		props RadioItemProps
		want  string
	}{
		{
			name:  "basic radio",
			props: RadioItemProps{Value: "option1"},
			want:  `data-value="option1"`,
		},
		{
			name:  "disabled radio",
			props: RadioItemProps{Value: "option2", Disabled: true},
			want:  `data-disabled="true"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := renderToString(RadioItem(tt.props, g.Text("Radio option")))
			if !strings.Contains(got, tt.want) {
				t.Errorf("RadioItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLabel(t *testing.T) {
	tests := []struct {
		name  string
		props LabelProps
		want  string
	}{
		{
			name:  "basic label",
			props: LabelProps{},
			want:  `data-context-menu="label"`,
		},
		{
			name:  "inset label",
			props: LabelProps{Inset: true},
			want:  `pl-8`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := renderToString(LabelComponent(tt.props, g.Text("Label")))
			if !strings.Contains(got, tt.want) {
				t.Errorf("LabelComponent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSeparator(t *testing.T) {
	got := renderToString(Separator(SeparatorProps{}))
	want := `data-context-menu="separator" role="separator"`
	if !strings.Contains(got, want) {
		t.Errorf("Separator() = %v, want %v", got, want)
	}
}

func TestSubMenu(t *testing.T) {
	tests := []struct {
		name  string
		props SubProps
		want  string
	}{
		{
			name:  "closed submenu",
			props: SubProps{Open: false},
			want:  `data-state="closed"`,
		},
		{
			name:  "open submenu",
			props: SubProps{Open: true},
			want:  `data-state="open"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := renderToString(SubMenu(tt.props, g.Text("Submenu")))
			if !strings.Contains(got, tt.want) {
				t.Errorf("SubMenu() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubTrigger(t *testing.T) {
	tests := []struct {
		name  string
		props SubTriggerProps
		want  string
	}{
		{
			name:  "basic trigger",
			props: SubTriggerProps{},
			want:  `data-context-menu="sub-trigger"`,
		},
		{
			name:  "disabled trigger",
			props: SubTriggerProps{Disabled: true},
			want:  `data-disabled="true"`,
		},
		{
			name:  "inset trigger",
			props: SubTriggerProps{Inset: true},
			want:  `pl-8`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := renderToString(SubTrigger(tt.props, g.Text("More")))
			if !strings.Contains(got, tt.want) {
				t.Errorf("SubTrigger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShortcut(t *testing.T) {
	got := renderToString(Shortcut(ShortcutProps{}, g.Text("⌘K")))
	want := `data-context-menu="shortcut"`
	if !strings.Contains(got, want) {
		t.Errorf("Shortcut() = %v, want %v", got, want)
	}
}

func TestCompleteMenu(t *testing.T) {
	menu := New(
		Props{ID: "test-menu"},
		Trigger(
			TriggerProps{Class: "trigger-area"},
			g.Text("Right click me"),
		),
		ContentComponent(
			ContentProps{},
			Item(ItemProps{}, g.Text("Copy")),
			Item(ItemProps{}, g.Text("Paste")),
			Separator(SeparatorProps{}),
			CheckboxItem(CheckboxItemProps{Checked: true}, g.Text("Show hints")),
			Separator(SeparatorProps{}),
			RadioGroup(
				RadioGroupProps{Value: "medium"},
				RadioItem(RadioItemProps{Value: "small"}, g.Text("Small")),
				RadioItem(RadioItemProps{Value: "medium"}, g.Text("Medium")),
				RadioItem(RadioItemProps{Value: "large"}, g.Text("Large")),
			),
			Separator(SeparatorProps{}),
			SubMenu(
				SubProps{},
				SubTrigger(SubTriggerProps{}, g.Text("More options")),
				SubContent(
					SubContentProps{},
					Item(ItemProps{}, g.Text("Settings")),
					Item(ItemProps{}, g.Text("Help")),
				),
			),
		),
	)

	got := renderToString(menu)
	
	// Check for key elements
	wants := []string{
		`id="test-menu"`,
		`data-context-menu="root"`,
		`Right click me`,
		`Copy`,
		`Paste`,
		`Show hints`,
		`role="group"`,
		`Small`,
		`Medium`,
		`Large`,
		`More options`,
		`Settings`,
	}

	for _, want := range wants {
		if !strings.Contains(got, want) {
			t.Errorf("Complete menu missing: %v", want)
		}
	}
}