package button

import (
	"bytes"
	"strings"
	"testing"

	g "maragu.dev/gomponents"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		props    Props
		children []g.Node
		contains []string
	}{
		{
			name:  "default button",
			props: Props{},
			children: []g.Node{
				g.Text("Click me"),
			},
			contains: []string{
				`<button`,
				`type="button"`,
				`Click me`,
				`bg-primary`,
				`text-primary-foreground`,
			},
		},
		{
			name: "destructive button",
			props: Props{
				Variant: "destructive",
			},
			children: []g.Node{
				g.Text("Delete"),
			},
			contains: []string{
				`bg-destructive`,
				`text-white`,
				`Delete`,
			},
		},
		{
			name: "disabled button",
			props: Props{
				Disabled: true,
			},
			contains: []string{
				`disabled`,
			},
		},
		{
			name: "submit button",
			props: Props{
				Type: "submit",
			},
			contains: []string{
				`type="submit"`,
			},
		},
		{
			name: "small button",
			props: Props{
				Size: "sm",
			},
			contains: []string{
				`h-8`,
				`gap-1.5`,
			},
		},
		{
			name: "icon button",
			props: Props{
				Size: "icon",
			},
			contains: []string{
				`size-9`,
			},
		},
		{
			name: "button with custom class",
			props: Props{
				Class: "ml-2",
			},
			contains: []string{
				`ml-2`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			btn := New(tt.props, tt.children...)
			err := btn.Render(&buf)
			if err != nil {
				t.Fatalf("Render() error = %v", err)
			}

			result := buf.String()
			for _, expected := range tt.contains {
				if !strings.Contains(result, expected) {
					t.Errorf("Expected output to contain %q, but it didn't.\nGot: %s", expected, result)
				}
			}
		})
	}
}

func TestVariantHelpers(t *testing.T) {
	tests := []struct {
		name     string
		fn       func(...g.Node) g.Node
		expected string
	}{
		{
			name:     "Default helper",
			fn:       Default,
			expected: "bg-primary",
		},
		{
			name:     "Destructive helper",
			fn:       Destructive,
			expected: "bg-destructive",
		},
		{
			name:     "Outline helper",
			fn:       Outline,
			expected: "border",
		},
		{
			name:     "Secondary helper",
			fn:       Secondary,
			expected: "bg-secondary",
		},
		{
			name:     "Ghost helper",
			fn:       Ghost,
			expected: "hover:bg-accent",
		},
		{
			name:     "LinkButton helper",
			fn:       LinkButton,
			expected: "text-primary",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			btn := tt.fn(g.Text("Test"))
			err := btn.Render(&buf)
			if err != nil {
				t.Fatalf("Render() error = %v", err)
			}

			result := buf.String()
			if !strings.Contains(result, tt.expected) {
				t.Errorf("Expected output to contain %q, but it didn't.\nGot: %s", tt.expected, result)
			}
		})
	}
}