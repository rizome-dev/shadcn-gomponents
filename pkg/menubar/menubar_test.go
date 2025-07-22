package menubar

import (
	"strings"
	"testing"

	g "maragu.dev/gomponents"
)

func renderToString(node g.Node) string {
	var buf strings.Builder
	_ = node.Render(&buf)
	return buf.String()
}

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		props    Props
		children []g.Node
		want     []string
	}{
		{
			name:  "renders basic menubar",
			props: Props{},
			want: []string{
				`role="menubar"`,
				`class="flex h-10 items-center space-x-1 rounded-md border bg-background p-1"`,
			},
		},
		{
			name: "renders with custom class",
			props: Props{
				Class: "custom-menubar",
			},
			want: []string{
				`custom-menubar`,
				`role="menubar"`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := renderToString(New(tt.props, tt.children...))

			for _, want := range tt.want {
				if !strings.Contains(result, want) {
					t.Errorf("New() = %v, want to contain %v", result, want)
				}
			}
		})
	}
}

func TestMenu(t *testing.T) {
	tests := []struct {
		name     string
		props    MenuProps
		children []g.Node
		want     []string
	}{
		{
			name:  "renders basic menu",
			props: MenuProps{},
			want: []string{
				`class="relative"`,
			},
		},
		{
			name: "renders with custom class",
			props: MenuProps{
				Class: "custom-menu",
			},
			want: []string{
				`custom-menu`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := renderToString(Menu(tt.props, tt.children...))

			for _, want := range tt.want {
				if !strings.Contains(result, want) {
					t.Errorf("Menu() = %v, want to contain %v", result, want)
				}
			}
		})
	}
}

func TestTrigger(t *testing.T) {
	tests := []struct {
		name     string
		props    TriggerProps
		children []g.Node
		want     []string
		notWant  []string
	}{
		{
			name:  "renders basic trigger",
			props: TriggerProps{},
			want: []string{
				`type="button"`,
				`role="menuitem"`,
				`aria-haspopup="menu"`,
				`data-state="closed"`,
			},
		},
		{
			name: "renders disabled trigger",
			props: TriggerProps{
				Disabled: true,
			},
			want: []string{
				`disabled`,
				`opacity-50 cursor-not-allowed`,
			},
		},
		{
			name: "renders with custom class",
			props: TriggerProps{
				Class: "custom-trigger",
			},
			want: []string{
				`custom-trigger`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := renderToString(Trigger(tt.props, tt.children...))

			for _, want := range tt.want {
				if !strings.Contains(result, want) {
					t.Errorf("Trigger() = %v, want to contain %v", result, want)
				}
			}

			for _, notWant := range tt.notWant {
				if strings.Contains(result, notWant) {
					t.Errorf("Trigger() = %v, don't want to contain %v", result, notWant)
				}
			}
		})
	}
}

func TestContent(t *testing.T) {
	tests := []struct {
		name     string
		props    ContentProps
		children []g.Node
		want     []string
	}{
		{
			name:  "renders basic content",
			props: ContentProps{},
			want: []string{
				`role="menu"`,
				`aria-orientation="vertical"`,
				`data-state="closed"`,
				`data-side="bottom"`,
				`style="display: none;"`,
				`left-0`, // default alignment
			},
		},
		{
			name: "renders with center alignment",
			props: ContentProps{
				Align: "center",
			},
			want: []string{
				`left-1/2 -translate-x-1/2`,
			},
		},
		{
			name: "renders with end alignment",
			props: ContentProps{
				Align: "end",
			},
			want: []string{
				`right-0`,
			},
		},
		{
			name: "renders with top side",
			props: ContentProps{
				Side: "top",
			},
			want: []string{
				`bottom-full mb-1`,
				`data-side="top"`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := renderToString(ContentComponent(tt.props, tt.children...))

			for _, want := range tt.want {
				if !strings.Contains(result, want) {
					t.Errorf("Content() = %v, want to contain %v", result, want)
				}
			}
		})
	}
}

func TestItem(t *testing.T) {
	tests := []struct {
		name     string
		props    ItemProps
		children []g.Node
		want     []string
	}{
		{
			name:  "renders basic item",
			props: ItemProps{},
			want: []string{
				`role="menuitem"`,
				`tabindex="-1"`,
			},
		},
		{
			name: "renders disabled item",
			props: ItemProps{
				Disabled: true,
			},
			want: []string{
				`data-disabled="true"`,
			},
		},
		{
			name: "renders inset item",
			props: ItemProps{
				Inset: true,
			},
			want: []string{
				`pl-8`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := renderToString(Item(tt.props, tt.children...))

			for _, want := range tt.want {
				if !strings.Contains(result, want) {
					t.Errorf("Item() = %v, want to contain %v", result, want)
				}
			}
		})
	}
}

func TestCheckboxItem(t *testing.T) {
	tests := []struct {
		name     string
		props    CheckboxItemProps
		children []g.Node
		want     []string
	}{
		{
			name:  "renders unchecked checkbox item",
			props: CheckboxItemProps{},
			want: []string{
				`role="menuitemcheckbox"`,
				`aria-checked="false"`,
				`tabindex="-1"`,
			},
		},
		{
			name: "renders checked checkbox item",
			props: CheckboxItemProps{
				Checked: true,
			},
			want: []string{
				`aria-checked="true"`,
				`<svg`, // checkmark icon
			},
		},
		{
			name: "renders disabled checkbox item",
			props: CheckboxItemProps{
				Disabled: true,
			},
			want: []string{
				`data-disabled="true"`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := renderToString(CheckboxItem(tt.props, tt.children...))

			for _, want := range tt.want {
				if !strings.Contains(result, want) {
					t.Errorf("CheckboxItem() = %v, want to contain %v", result, want)
				}
			}
		})
	}
}

func TestRadioGroup(t *testing.T) {
	result := renderToString(RadioGroup(RadioGroupProps{Value: "test"}))

	want := []string{
		`role="group"`,
		`aria-orientation="vertical"`,
	}

	for _, w := range want {
		if !strings.Contains(result, w) {
			t.Errorf("RadioGroup() = %v, want to contain %v", result, w)
		}
	}
}

func TestRadioItem(t *testing.T) {
	tests := []struct {
		name     string
		props    RadioItemProps
		children []g.Node
		want     []string
	}{
		{
			name: "renders basic radio item",
			props: RadioItemProps{
				Value: "option1",
			},
			want: []string{
				`role="menuitemradio"`,
				`tabindex="-1"`,
			},
		},
		{
			name: "renders disabled radio item",
			props: RadioItemProps{
				Value:    "option1",
				Disabled: true,
			},
			want: []string{
				`data-disabled="true"`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := renderToString(RadioItem(tt.props, tt.children...))

			for _, want := range tt.want {
				if !strings.Contains(result, want) {
					t.Errorf("RadioItem() = %v, want to contain %v", result, want)
				}
			}
		})
	}
}

func TestLabel(t *testing.T) {
	tests := []struct {
		name  string
		props LabelProps
		want  []string
	}{
		{
			name:  "renders basic label",
			props: LabelProps{},
			want: []string{
				`px-2 py-1.5 text-sm font-semibold`,
			},
		},
		{
			name: "renders inset label",
			props: LabelProps{
				Inset: true,
			},
			want: []string{
				`pl-8`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := renderToString(LabelComponent(tt.props, g.Text("Test Label")))

			for _, want := range tt.want {
				if !strings.Contains(result, want) {
					t.Errorf("Label() = %v, want to contain %v", result, want)
				}
			}
		})
	}
}

func TestSeparator(t *testing.T) {
	result := renderToString(Separator(SeparatorProps{}))

	want := []string{
		`role="separator"`,
		`-mx-1 my-1 h-px bg-muted`,
	}

	for _, w := range want {
		if !strings.Contains(result, w) {
			t.Errorf("Separator() = %v, want to contain %v", result, w)
		}
	}
}

func TestSubMenu(t *testing.T) {
	result := renderToString(SubMenu(SubMenuProps{}, g.Text("submenu")))

	if !strings.Contains(result, `class="relative"`) {
		t.Errorf("SubMenu() = %v, want to contain relative class", result)
	}
}

func TestSubTrigger(t *testing.T) {
	tests := []struct {
		name  string
		props SubTriggerProps
		want  []string
	}{
		{
			name:  "renders basic sub trigger",
			props: SubTriggerProps{},
			want: []string{
				`role="menuitem"`,
				`aria-haspopup="menu"`,
				`aria-expanded="false"`,
				`<svg`, // chevron icon
			},
		},
		{
			name: "renders disabled sub trigger",
			props: SubTriggerProps{
				Disabled: true,
			},
			want: []string{
				`data-disabled="true"`,
				`opacity-50 cursor-not-allowed`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := renderToString(SubTrigger(tt.props, g.Text("Submenu")))

			for _, want := range tt.want {
				if !strings.Contains(result, want) {
					t.Errorf("SubTrigger() = %v, want to contain %v", result, want)
				}
			}
		})
	}
}

func TestSubContent(t *testing.T) {
	result := renderToString(SubContent(SubContentProps{}, g.Text("content")))

	want := []string{
		`role="menu"`,
		`aria-orientation="vertical"`,
		`data-state="closed"`,
		`style="display: none;"`,
		`absolute left-full top-0 ml-1`,
	}

	for _, w := range want {
		if !strings.Contains(result, w) {
			t.Errorf("SubContent() = %v, want to contain %v", result, w)
		}
	}
}

func TestShortcut(t *testing.T) {
	result := renderToString(Shortcut(ShortcutProps{}, g.Text("⌘K")))

	want := []string{
		`ml-auto text-xs tracking-widest text-muted-foreground`,
		`⌘K`,
	}

	for _, w := range want {
		if !strings.Contains(result, w) {
			t.Errorf("Shortcut() = %v, want to contain %v", result, w)
		}
	}
}

func TestDefault(t *testing.T) {
	result := renderToString(Default())

	if !strings.Contains(result, `role="menubar"`) {
		t.Errorf("Default() = %v, want to contain menubar role", result)
	}
}

func TestCompleteMenubar(t *testing.T) {
	menubar := New(
		Props{},
		Menu(
			MenuProps{},
			Trigger(TriggerProps{}, g.Text("File")),
			ContentComponent(
				ContentProps{},
				Item(ItemProps{}, g.Text("New")),
				Item(ItemProps{}, g.Text("Open")),
				Separator(SeparatorProps{}),
				CheckboxItem(CheckboxItemProps{Checked: true}, g.Text("Auto Save")),
				Separator(SeparatorProps{}),
				SubMenu(
					SubMenuProps{},
					SubTrigger(SubTriggerProps{}, g.Text("Recent")),
					SubContent(
						SubContentProps{},
						Item(ItemProps{}, g.Text("File 1")),
						Item(ItemProps{}, g.Text("File 2")),
					),
				),
			),
		),
	)

	result := renderToString(menubar)

	// Check structure
	expectedElements := []string{
		`role="menubar"`,
		`File`,
		`New`,
		`Open`,
		`Auto Save`,
		`Recent`,
		`File 1`,
		`File 2`,
		`role="separator"`,
		`aria-checked="true"`,
	}

	for _, elem := range expectedElements {
		if !strings.Contains(result, elem) {
			t.Errorf("Complete menubar = %v, want to contain %v", result, elem)
		}
	}
}