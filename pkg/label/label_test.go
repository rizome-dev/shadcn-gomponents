package label

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
			name:  "basic label",
			props: Props{},
			children: []g.Node{
				g.Text("Username"),
			},
			contains: []string{
				`<label`,
				`Username`,
				`text-sm`,
				`font-medium`,
			},
		},
		{
			name: "label with for attribute",
			props: Props{
				For: "username",
			},
			children: []g.Node{
				g.Text("Username"),
			},
			contains: []string{
				`for="username"`,
				`Username`,
			},
		},
		{
			name: "label with required indicator",
			props: Props{
				Required: true,
			},
			children: []g.Node{
				g.Text("Email"),
			},
			contains: []string{
				`Email`,
				`<span`,
				`*`,
				`text-destructive`,
			},
		},
		{
			name: "label with custom class",
			props: Props{
				Class: "mt-4",
			},
			children: []g.Node{
				g.Text("Password"),
			},
			contains: []string{
				`Password`,
				`mt-4`,
			},
		},
		{
			name: "label with all props",
			props: Props{
				For:      "email-field",
				Required: true,
				Class:    "text-red-500",
			},
			children: []g.Node{
				g.Text("Email Address"),
			},
			contains: []string{
				`for="email-field"`,
				`Email Address`,
				`*`,
				`text-red-500`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			label := New(tt.props, tt.children...)
			err := label.Render(&buf)
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

func TestHelpers(t *testing.T) {
	tests := []struct {
		name     string
		fn       func() g.Node
		expected []string
	}{
		{
			name: "Default helper",
			fn: func() g.Node {
				return Default("Simple Label")
			},
			expected: []string{`Simple Label`},
		},
		{
			name: "WithRequired helper - not required",
			fn: func() g.Node {
				return WithRequired("Optional Field", false)
			},
			expected: []string{`Optional Field`},
		},
		{
			name: "WithRequired helper - required",
			fn: func() g.Node {
				return WithRequired("Required Field", true)
			},
			expected: []string{`Required Field`, `*`},
		},
		{
			name: "ForInput helper",
			fn: func() g.Node {
				return ForInput("my-input", "My Input")
			},
			expected: []string{`for="my-input"`, `My Input`},
		},
		{
			name: "ForInputRequired helper",
			fn: func() g.Node {
				return ForInputRequired("req-input", "Required Input")
			},
			expected: []string{`for="req-input"`, `Required Input`, `*`},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			label := tt.fn()
			err := label.Render(&buf)
			if err != nil {
				t.Fatalf("Render() error = %v", err)
			}

			result := buf.String()
			for _, expected := range tt.expected {
				if !strings.Contains(result, expected) {
					t.Errorf("Expected output to contain %q, but it didn't.\nGot: %s", expected, result)
				}
			}
		})
	}
}