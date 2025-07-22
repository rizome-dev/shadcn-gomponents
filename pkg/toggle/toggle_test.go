package toggle_test

import (
	"strings"
	"testing"
	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/pkg/toggle"
)

func renderToString(node g.Node) string {
	var buf strings.Builder
	err := node.Render(&buf)
	if err != nil {
		panic(err) // For tests, panic on render error
	}
	return buf.String()
}

func TestToggle(t *testing.T) {
	tests := []struct {
		name     string
		toggle   g.Node
		contains []string
	}{
		{
			name:   "default toggle",
			toggle: toggle.Default(g.Text("Toggle")),
			contains: []string{
				`type="button"`,
				`role="button"`,
				`data-state="off"`,
				`aria-pressed="false"`,
				`class="inline-flex items-center justify-center rounded-md`,
				`bg-transparent`,
				`>Toggle</button>`,
			},
		},
		{
			name:   "pressed toggle",
			toggle: toggle.Pressed(g.Text("Active")),
			contains: []string{
				`data-state="on"`,
				`aria-pressed="true"`,
				`data-[state=on]:bg-accent`,
				`>Active</button>`,
			},
		},
		{
			name:   "outline toggle",
			toggle: toggle.Outline(g.Text("Outline")),
			contains: []string{
				`border border-input`,
				`hover:bg-accent hover:text-accent-foreground`,
			},
		},
		{
			name: "toggle with all props",
			toggle: toggle.New(toggle.Props{
				ID:        "my-toggle",
				Pressed:   true,
				Disabled:  true,
				AriaLabel: "Toggle feature",
				Variant:   "outline",
				Size:      "lg",
				OnClick:   "handleToggle()",
			}, g.Text("Feature")),
			contains: []string{
				`id="my-toggle"`,
				`data-state="on"`,
				`disabled`,
				`aria-label="Toggle feature"`,
				`aria-pressed="true"`,
				`onclick="handleToggle()"`,
				`h-11 px-5 min-w-11`, // Large size
			},
		},
		{
			name:   "small toggle",
			toggle: toggle.SmallToggle("default", g.Text("Small")),
			contains: []string{
				`h-9 px-2.5 min-w-9`,
			},
		},
		{
			name:   "large toggle",
			toggle: toggle.Large("outline", g.Text("Large")),
			contains: []string{
				`h-11 px-5 min-w-11`,
				`border border-input`,
			},
		},
		{
			name: "toggle with icon only",
			toggle: toggle.Icon(
				g.El("svg", Class("h-4 w-4")),
				"Toggle bold",
			),
			contains: []string{
				`aria-label="Toggle bold"`,
				`<svg class="h-4 w-4"`,
				`[&amp;_svg]:size-4`,
			},
		},
		{
			name: "toggle with icon and text",
			toggle: toggle.WithIcon(
				g.El("svg", Class("h-4 w-4")),
				"Bold",
				false,
			),
			contains: []string{
				`<svg class="h-4 w-4"`,
				`Bold</button>`,
				`gap-2`,
			},
		},
		{
			name: "toggle group",
			toggle: toggle.Group(
				toggle.Default(g.Text("A")),
				toggle.Default(g.Text("B")),
				toggle.Default(g.Text("C")),
			),
			contains: []string{
				`class="flex items-center gap-1"`,
				`role="group"`,
				`>A</button>`,
				`>B</button>`,
				`>C</button>`,
			},
		},
		{
			name: "toolbar item",
			toggle: toggle.ToolbarItem(
				g.El("svg"),
				"Bold text",
				true,
			),
			contains: []string{
				`aria-label="Bold text"`,
				`data-state="on"`,
				`h-9 px-2.5 min-w-9`, // Small size
			},
		},
		{
			name: "format button",
			toggle: toggle.FormatButton(
				"bold",
				g.El("svg"),
				false,
			),
			contains: []string{
				`aria-label="Toggle bold"`,
				`onclick="document.execCommand(&#39;bold&#39;)"`,
			},
		},
		{
			name: "view toggle",
			toggle: toggle.ViewToggle(
				"grid",
				g.El("svg"),
				true,
			),
			contains: []string{
				`aria-label="grid view"`,
				`data-state="on"`,
				`border border-input`, // Outline variant
			},
		},
		{
			name: "custom data state",
			toggle: toggle.New(toggle.Props{
				Pressed:   false,
				DataState: "indeterminate",
			}, g.Text("Custom")),
			contains: []string{
				`data-state="indeterminate"`,
				`aria-pressed="false"`,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := renderToString(test.toggle)
			for _, expected := range test.contains {
				if !strings.Contains(result, expected) {
					t.Errorf("expected result to contain %q, but it didn't.\nGot: %s", expected, result)
				}
			}
		})
	}
}