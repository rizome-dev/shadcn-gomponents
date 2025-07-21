package input

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
		contains []string
	}{
		{
			name:  "default text input",
			props: Props{},
			contains: []string{
				`<input`,
				`type="text"`,
				`class="`,
				`h-9`,
				`w-full`,
				`rounded-md`,
			},
		},
		{
			name: "email input with placeholder",
			props: Props{
				Type:        "email",
				Placeholder: "Enter your email",
			},
			contains: []string{
				`type="email"`,
				`placeholder="Enter your email"`,
			},
		},
		{
			name: "disabled input",
			props: Props{
				Disabled: true,
			},
			contains: []string{
				`disabled`,
			},
		},
		{
			name: "input with all props",
			props: Props{
				Type:         "password",
				Placeholder:  "Password",
				Name:         "password",
				ID:           "password-field",
				Value:        "secret",
				Required:     true,
				AriaInvalid:  true,
				AutoComplete: "current-password",
				Class:        "mt-2",
			},
			contains: []string{
				`type="password"`,
				`placeholder="Password"`,
				`name="password"`,
				`id="password-field"`,
				`value="secret"`,
				`required`,
				`aria-invalid="true"`,
				`autocomplete="current-password"`,
				`mt-2`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			input := New(tt.props)
			err := input.Render(&buf)
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
		name        string
		fn          func(string) g.Node
		placeholder string
		expected    []string
	}{
		{
			name:        "Text helper",
			fn:          func(p string) g.Node { return Text(p) },
			placeholder: "Enter text",
			expected:    []string{`type="text"`, `placeholder="Enter text"`},
		},
		{
			name:        "Email helper",
			fn:          func(p string) g.Node { return Email(p) },
			placeholder: "Email address",
			expected:    []string{`type="email"`, `placeholder="Email address"`, `autocomplete="email"`},
		},
		{
			name:        "Password helper",
			fn:          func(p string) g.Node { return Password(p) },
			placeholder: "Password",
			expected:    []string{`type="password"`, `placeholder="Password"`, `autocomplete="current-password"`},
		},
		{
			name:        "Number helper",
			fn:          func(p string) g.Node { return Number(p) },
			placeholder: "Age",
			expected:    []string{`type="number"`, `placeholder="Age"`},
		},
		{
			name:        "Search helper",
			fn:          func(p string) g.Node { return Search(p) },
			placeholder: "Search...",
			expected:    []string{`type="search"`, `placeholder="Search..."`},
		},
		{
			name:        "Tel helper",
			fn:          func(p string) g.Node { return Tel(p) },
			placeholder: "Phone",
			expected:    []string{`type="tel"`, `placeholder="Phone"`, `autocomplete="tel"`},
		},
		{
			name:        "URL helper",
			fn:          func(p string) g.Node { return URL(p) },
			placeholder: "Website",
			expected:    []string{`type="url"`, `placeholder="Website"`, `autocomplete="url"`},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			node := tt.fn(tt.placeholder)
			err := node.Render(&buf)
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