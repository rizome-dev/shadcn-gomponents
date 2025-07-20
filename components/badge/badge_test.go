package badge_test

import (
	"strings"
	"testing"
	g "maragu.dev/gomponents"
	"github.com/rizome-dev/shadcn-gomponents/components/badge"
)

func TestBadge(t *testing.T) {
	tests := []struct {
		name     string
		badge    g.Node
		contains []string
	}{
		{
			name:  "default badge",
			badge: badge.Default(g.Text("Badge")),
			contains: []string{
				`class="inline-flex items-center rounded-md border`,
				`border-transparent bg-primary text-primary-foreground shadow hover:bg-primary/80`,
				`px-2.5 py-0.5 text-xs font-semibold`,
				">Badge</div>",
			},
		},
		{
			name:  "secondary badge",
			badge: badge.Secondary(g.Text("Secondary")),
			contains: []string{
				`border-transparent bg-secondary text-secondary-foreground hover:bg-secondary/80`,
				">Secondary</div>",
			},
		},
		{
			name:  "destructive badge",
			badge: badge.Destructive(g.Text("Error")),
			contains: []string{
				`border-transparent bg-destructive text-destructive-foreground shadow hover:bg-destructive/80`,
				">Error</div>",
			},
		},
		{
			name:  "outline badge",
			badge: badge.Outline(g.Text("Outline")),
			contains: []string{
				`text-foreground`,
				`border px-2.5`,
				">Outline</div>",
			},
		},
		{
			name: "badge with custom class",
			badge: badge.New(
				badge.Props{
					Variant: "default",
					Class:   "ml-2",
				},
				g.Text("Custom"),
			),
			contains: []string{
				`ml-2`,
				">Custom</div>",
			},
		},
		{
			name: "badge with icon",
			badge: badge.WithIcon(
				g.El("svg", g.Attr("viewBox", "0 0 24 24")),
				"default",
				"Icon Badge",
			),
			contains: []string{
				`gap-1`,
				`<svg viewBox="0 0 24 24"`,
				`class="h-3 w-3"`,
				"Icon Badge",
			},
		},
		{
			name: "badge as link",
			badge: badge.Link(
				"/docs",
				badge.Props{Variant: "secondary"},
				g.Text("Documentation"),
			),
			contains: []string{
				`<a href="/docs"`,
				`hover:underline`,
				`bg-secondary`,
				">Documentation</a>",
			},
		},
		{
			name: "multiple badges together",
			badge: g.Group{
				badge.Default(g.Text("New")),
				badge.Secondary(g.Text("v2.0")),
				badge.Outline(g.Text("Beta")),
			},
			contains: []string{
				">New</div>",
				">v2.0</div>",
				">Beta</div>",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.badge.String()
			for _, expected := range test.contains {
				if !strings.Contains(result, expected) {
					t.Errorf("expected result to contain %q, but it didn't.\nGot: %s", expected, result)
				}
			}
		})
	}
}