package togglegroup_test

import (
	"strings"
	"testing"
	g "maragu.dev/gomponents"
	"github.com/rizome-dev/shadcn-gomponents/components/togglegroup"
)

func renderToString(node g.Node) string {
	var buf strings.Builder
	node.Render(&buf)
	return buf.String()
}

func TestToggleGroup(t *testing.T) {
	tests := []struct {
		name     string
		group    g.Node
		contains []string
	}{
		{
			name:  "default toggle group",
			group: togglegroup.New(togglegroup.Props{}, g.Text("Content")),
			contains: []string{
				`role="group"`,
				`data-toggle-group="true"`,
				`data-type="single"`,
				`data-variant="default"`,
				`data-size="default"`,
				`class="group/toggle-group flex w-fit items-center rounded-md"`,
				`>Content</div>`,
			},
		},
		{
			name: "single selection toggle group",
			group: togglegroup.Single(togglegroup.Props{
				ID:    "single-group",
				Value: []string{"option1"},
			}, g.Text("Options")),
			contains: []string{
				`id="single-group"`,
				`data-type="single"`,
				`data-value="option1"`,
			},
		},
		{
			name: "multiple selection toggle group",
			group: togglegroup.MultipleSelection(togglegroup.Props{
				Value: []string{"option1", "option2"},
			}, g.Text("Options")),
			contains: []string{
				`data-type="multiple"`,
				`data-value="option1,option2"`,
			},
		},
		{
			name: "outline variant toggle group",
			group: togglegroup.New(togglegroup.Props{
				Variant: "outline",
			}, g.Text("Outline")),
			contains: []string{
				`data-variant="outline"`,
				`shadow-sm`,
			},
		},
		{
			name: "disabled toggle group",
			group: togglegroup.New(togglegroup.Props{
				Disabled: true,
			}, g.Text("Disabled")),
			contains: []string{
				`data-disabled="true"`,
			},
		},
		{
			name: "toggle group with custom class",
			group: togglegroup.New(togglegroup.Props{
				Class: "custom-class",
			}, g.Text("Custom")),
			contains: []string{
				`custom-class"`,
			},
		},
		{
			name:  "default group helper",
			group: togglegroup.DefaultGroup(g.Text("Default")),
			contains: []string{
				`data-variant="default"`,
			},
		},
		{
			name:  "outline group helper",
			group: togglegroup.OutlineGroup(g.Text("Outline")),
			contains: []string{
				`data-variant="outline"`,
			},
		},
		{
			name:  "small group helper",
			group: togglegroup.SmallGroup(togglegroup.Props{}, g.Text("Small")),
			contains: []string{
				`data-size="sm"`,
			},
		},
		{
			name:  "large group helper",
			group: togglegroup.LargeGroup(togglegroup.Props{}, g.Text("Large")),
			contains: []string{
				`data-size="lg"`,
			},
		},
		{
			name:  "text formatting preset",
			group: togglegroup.TextFormatting(),
			contains: []string{
				`data-type="single"`,
				`data-variant="outline"`,
				`data-value="bold"`,
				`aria-label="Toggle bold"`,
				`aria-label="Toggle italic"`,
				`aria-label="Toggle underline"`,
			},
		},
		{
			name:  "alignment preset",
			group: togglegroup.Alignment(),
			contains: []string{
				`data-type="single"`,
				`data-variant="outline"`,
				`data-value="center"`,
				`aria-label="Align left"`,
				`aria-label="Align center"`,
				`aria-label="Align right"`,
				`aria-label="Justify"`,
			},
		},
		{
			name: "with items helper",
			group: togglegroup.WithItems(
				togglegroup.Props{Type: togglegroup.TypeMultiple},
				[]struct {
					Value     string
					Label     string
					Icon      g.Node
					Disabled  bool
					AriaLabel string
				}{
					{Value: "bold", Label: "B", AriaLabel: "Bold text"},
					{Value: "italic", Label: "I", Disabled: true},
					{Value: "underline", Label: "U"},
				},
			),
			contains: []string{
				`data-value="bold"`,
				`aria-label="Bold text"`,
				`>B</`,
				`data-value="italic"`,
				`disabled`,
				`>I</`,
				`data-value="underline"`,
				`>U</`,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := renderToString(test.group)
			for _, expected := range test.contains {
				if !strings.Contains(result, expected) {
					t.Errorf("expected result to contain %q, but it didn't.\nGot: %s", expected, result)
				}
			}
		})
	}
}

func TestToggleGroupItem(t *testing.T) {
	groupProps := togglegroup.Props{
		Type:    togglegroup.TypeSingle,
		Variant: "outline",
		Size:    "default",
		Value:   []string{"selected"},
	}

	tests := []struct {
		name     string
		item     g.Node
		contains []string
	}{
		{
			name: "basic item",
			item: togglegroup.Item(togglegroup.ItemProps{
				Value: "item1",
			}, groupProps, g.Text("Item 1")),
			contains: []string{
				`data-value="item1"`,
				`type="button"`,
				`role="button"`,
				`aria-pressed="false"`,
				`data-state="off"`,
				`min-w-0 flex-1 shrink-0 rounded-none shadow-none`,
				`first:rounded-l-md last:rounded-r-md`,
				`>Item 1</button>`,
			},
		},
		{
			name: "selected item",
			item: togglegroup.Item(togglegroup.ItemProps{
				Value: "selected",
			}, groupProps, g.Text("Selected")),
			contains: []string{
				`data-value="selected"`,
				`aria-pressed="true"`,
				`data-state="on"`,
			},
		},
		{
			name: "item with pressed prop",
			item: togglegroup.Item(togglegroup.ItemProps{
				Value:   "pressed",
				Pressed: true,
			}, groupProps, g.Text("Pressed")),
			contains: []string{
				`aria-pressed="true"`,
				`data-state="on"`,
			},
		},
		{
			name: "disabled item",
			item: togglegroup.Item(togglegroup.ItemProps{
				Value:    "disabled",
				Disabled: true,
			}, groupProps, g.Text("Disabled")),
			contains: []string{
				`disabled`,
			},
		},
		{
			name: "item with aria label",
			item: togglegroup.Item(togglegroup.ItemProps{
				Value:     "accessible",
				AriaLabel: "Custom label",
			}, groupProps, g.Text("A")),
			contains: []string{
				`aria-label="Custom label"`,
			},
		},
		{
			name: "item with click handler",
			item: togglegroup.Item(togglegroup.ItemProps{
				Value:   "clickable",
				OnClick: "handleClick()",
			}, groupProps, g.Text("Click")),
			contains: []string{
				`onclick="handleClick()"`,
			},
		},
		{
			name: "item with custom class",
			item: togglegroup.Item(togglegroup.ItemProps{
				Value: "custom",
				Class: "custom-item-class",
			}, groupProps, g.Text("Custom")),
			contains: []string{
				`custom-item-class`,
			},
		},
		{
			name: "item in outline group",
			item: togglegroup.Item(togglegroup.ItemProps{
				Value: "outline",
			}, groupProps, g.Text("Outline")),
			contains: []string{
				`border-l-0 first:border-l`,
			},
		},
		{
			name: "item in disabled group",
			item: togglegroup.Item(togglegroup.ItemProps{
				Value: "group-disabled",
			}, togglegroup.Props{
				Disabled: true,
				Variant:  "default",
			}, g.Text("Disabled by Group")),
			contains: []string{
				`disabled`,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := renderToString(test.item)
			for _, expected := range test.contains {
				if !strings.Contains(result, expected) {
					t.Errorf("expected result to contain %q, but it didn't.\nGot: %s", expected, result)
				}
			}
		})
	}
}