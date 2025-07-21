package separator_test

import (
	"strings"
	"testing"
	g "maragu.dev/gomponents"
	"github.com/rizome-dev/shadcn-gomponents/pkg/separator"
)

func TestSeparator(t *testing.T) {
	tests := []struct {
		name      string
		separator g.Node
		contains  []string
		notContains []string
	}{
		{
			name:      "horizontal separator",
			separator: separator.Horizontal(),
			contains: []string{
				`class="shrink-0 bg-border h-[1px] w-full"`,
				`data-orientation="horizontal"`,
				`role="none"`,
			},
			notContains: []string{
				`aria-orientation`,
			},
		},
		{
			name:      "vertical separator",
			separator: separator.Vertical(),
			contains: []string{
				`class="shrink-0 bg-border h-full w-[1px]"`,
				`data-orientation="vertical"`,
				`role="none"`,
			},
		},
		{
			name:      "semantic horizontal separator",
			separator: separator.Semantic(),
			contains: []string{
				`role="separator"`,
				`aria-orientation="horizontal"`,
				`data-orientation="horizontal"`,
			},
		},
		{
			name:      "semantic vertical separator",
			separator: separator.SemanticVertical(),
			contains: []string{
				`role="separator"`,
				`aria-orientation="vertical"`,
				`data-orientation="vertical"`,
				`h-full w-[1px]`,
			},
		},
		{
			name:      "separator with custom class",
			separator: separator.WithClass("my-custom-class"),
			contains: []string{
				`my-custom-class`,
				`shrink-0 bg-border`,
			},
		},
		{
			name: "custom separator with all props",
			separator: separator.New(separator.Props{
				Orientation: "vertical",
				Decorative:  false,
				Class:       "custom-separator",
			}),
			contains: []string{
				`data-orientation="vertical"`,
				`role="separator"`,
				`aria-orientation="vertical"`,
				`custom-separator`,
				`h-full w-[1px]`,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.separator.String()
			
			// Check for expected content
			for _, expected := range test.contains {
				if !strings.Contains(result, expected) {
					t.Errorf("expected result to contain %q, but it didn't.\nGot: %s", expected, result)
				}
			}
			
			// Check for content that should not be present
			for _, notExpected := range test.notContains {
				if strings.Contains(result, notExpected) {
					t.Errorf("expected result to NOT contain %q, but it did.\nGot: %s", notExpected, result)
				}
			}
		})
	}
}