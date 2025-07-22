package combobox

import (
	"bytes"
	"strings"
	"testing"

	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		props    Props
		contains []string
	}{
		{
			name: "default combobox",
			props: Props{
				Options: []Option{
					{Value: "1", Label: "Option 1"},
					{Value: "2", Label: "Option 2"},
				},
			},
			contains: []string{
				`role="combobox"`,
				`aria-expanded="false"`,
				`Select option...`,
				`class="ml-2 h-4 w-4 shrink-0 opacity-50"`, // ChevronsUpDown icon
			},
		},
		{
			name: "combobox with selected value",
			props: Props{
				Value: "2",
				Options: []Option{
					{Value: "1", Label: "Option 1"},
					{Value: "2", Label: "Option 2"},
				},
			},
			contains: []string{
				`Option 2`, // Selected label shown
			},
		},
		{
			name: "open combobox",
			props: Props{
				Open: true,
				Options: []Option{
					{Value: "1", Label: "Option 1"},
					{Value: "2", Label: "Option 2"},
				},
			},
			contains: []string{
				`aria-expanded="true"`,
				`data-state="open"`,
			},
		},
		{
			name: "combobox with custom placeholder",
			props: Props{
				Placeholder: "Choose an option",
				Options: []Option{
					{Value: "1", Label: "Option 1"},
				},
			},
			contains: []string{
				`Choose an option`,
			},
		},
		{
			name: "disabled combobox",
			props: Props{
				Disabled: true,
				Options: []Option{
					{Value: "1", Label: "Option 1"},
				},
			},
			contains: []string{
				`disabled`,
			},
		},
		{
			name: "combobox with custom width",
			props: Props{
				Width: "w-[300px]",
				Options: []Option{
					{Value: "1", Label: "Option 1"},
				},
			},
			contains: []string{
				`w-[300px]`,
			},
		},
		{
			name: "combobox with ID and name",
			props: Props{
				ID:   "my-combo",
				Name: "myCombo",
				Options: []Option{
					{Value: "1", Label: "Option 1"},
				},
			},
			contains: []string{
				`id="my-combo"`,
				`name="myCombo"`,
			},
		},
		{
			name: "combobox with empty text",
			props: Props{
				Open:      true,
				EmptyText: "No items found",
				Options:   []Option{},
			},
			contains: []string{
				`No items found`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := New(tt.props)
			var buf bytes.Buffer
			err := component.Render(&buf)
			if err != nil {
				t.Fatalf("failed to render: %v", err)
			}

			html := buf.String()
			for _, want := range tt.contains {
				if !strings.Contains(html, want) {
					t.Errorf("expected HTML to contain %q, but got:\n%s", want, html)
				}
			}
		})
	}
}

func TestMulti(t *testing.T) {
	tests := []struct {
		name     string
		props    MultiProps
		contains []string
	}{
		{
			name: "default multi-select",
			props: MultiProps{
				Options: []Option{
					{Value: "1", Label: "Option 1"},
					{Value: "2", Label: "Option 2"},
				},
			},
			contains: []string{
				`Select options...`,
				`w-[280px]`, // Default width for multi
			},
		},
		{
			name: "multi-select with single value",
			props: MultiProps{
				Values: []string{"1"},
				Options: []Option{
					{Value: "1", Label: "Option 1"},
					{Value: "2", Label: "Option 2"},
				},
			},
			contains: []string{
				`Option 1`, // Single selected label
			},
		},
		{
			name: "multi-select with multiple values",
			props: MultiProps{
				Values: []string{"1", "2"},
				Options: []Option{
					{Value: "1", Label: "Option 1"},
					{Value: "2", Label: "Option 2"},
					{Value: "3", Label: "Option 3"},
				},
			},
			contains: []string{
				`2 selected`,
			},
		},
		{
			name: "multi-select with max items",
			props: MultiProps{
				Values:   []string{"1", "2"},
				MaxItems: 2,
				Options: []Option{
					{Value: "1", Label: "Option 1"},
					{Value: "2", Label: "Option 2"},
					{Value: "3", Label: "Option 3"},
				},
			},
			contains: []string{
				`2 selected`,
			},
		},
		{
			name: "disabled multi-select",
			props: MultiProps{
				Disabled: true,
				Options: []Option{
					{Value: "1", Label: "Option 1"},
				},
			},
			contains: []string{
				`disabled`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := Multi(tt.props)
			var buf bytes.Buffer
			err := component.Render(&buf)
			if err != nil {
				t.Fatalf("failed to render: %v", err)
			}

			html := buf.String()
			for _, want := range tt.contains {
				if !strings.Contains(html, want) {
					t.Errorf("expected HTML to contain %q, but got:\n%s", want, html)
				}
			}
		})
	}
}

func TestWithGroups(t *testing.T) {
	tests := []struct {
		name     string
		props    GroupedProps
		contains []string
	}{
		{
			name: "grouped combobox",
			props: GroupedProps{
				Groups: []OptionGroup{
					{
						Label: "Fruits",
						Options: []Option{
							{Value: "apple", Label: "Apple"},
							{Value: "banana", Label: "Banana"},
						},
					},
					{
						Label: "Vegetables",
						Options: []Option{
							{Value: "carrot", Label: "Carrot"},
							{Value: "broccoli", Label: "Broccoli"},
						},
					},
				},
			},
			contains: []string{
				`Select option...`,
			},
		},
		{
			name: "grouped combobox with selection",
			props: GroupedProps{
				Value: "banana",
				Groups: []OptionGroup{
					{
						Label: "Fruits",
						Options: []Option{
							{Value: "apple", Label: "Apple"},
							{Value: "banana", Label: "Banana"},
						},
					},
				},
			},
			contains: []string{
				`Banana`, // Selected item shown
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := WithGroups(tt.props)
			var buf bytes.Buffer
			err := component.Render(&buf)
			if err != nil {
				t.Fatalf("failed to render: %v", err)
			}

			html := buf.String()
			for _, want := range tt.contains {
				if !strings.Contains(html, want) {
					t.Errorf("expected HTML to contain %q, but got:\n%s", want, html)
				}
			}
		})
	}
}

func TestHelperFunctions(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		component := Simple("opt2", []string{"opt1", "opt2", "opt3"})
		var buf bytes.Buffer
		err := component.Render(&buf)
		if err != nil {
			t.Fatalf("failed to render: %v", err)
		}

		html := buf.String()
		if !strings.Contains(html, "opt2") {
			t.Errorf("expected Simple combobox to show selected value")
		}
	})

	t.Run("WithIcons", func(t *testing.T) {
		component := WithIcons("1", []Option{
			{Value: "1", Label: "User", Icon: html.Span(g.Text("👤"))},
			{Value: "2", Label: "Team", Icon: html.Span(g.Text("👥"))},
		})
		var buf bytes.Buffer
		err := component.Render(&buf)
		if err != nil {
			t.Fatalf("failed to render: %v", err)
		}

		html := buf.String()
		if !strings.Contains(html, "User") {
			t.Errorf("expected WithIcons combobox to show selected label")
		}
		if !strings.Contains(html, "w-[250px]") {
			t.Errorf("expected WithIcons combobox to have wider width")
		}
	})

	t.Run("Searchable", func(t *testing.T) {
		component := Searchable("", []Option{
			{Value: "1", Label: "Option 1"},
			{Value: "2", Label: "Option 2"},
		}, "Search items...")
		var buf bytes.Buffer
		err := component.Render(&buf)
		if err != nil {
			t.Fatalf("failed to render: %v", err)
		}

		html := buf.String()
		if !strings.Contains(html, "w-[300px]") {
			t.Errorf("expected Searchable combobox to have wide width")
		}
	})
}

func TestComboboxWithOptions(t *testing.T) {
	tests := []struct {
		name     string
		props    Props
		contains []string
	}{
		{
			name: "options with icons",
			props: Props{
				Open: true,
				Options: []Option{
					{
						Value: "user",
						Label: "User",
						Icon:  html.Span(g.Text("👤")),
					},
					{
						Value: "team",
						Label: "Team",
						Icon:  html.Span(g.Text("👥")),
					},
				},
			},
			contains: []string{
				`class="mr-2 h-4 w-4"`, // Icon wrapper
			},
		},
		{
			name: "disabled options",
			props: Props{
				Open: true,
				Options: []Option{
					{Value: "1", Label: "Enabled"},
					{Value: "2", Label: "Disabled", Disabled: true},
				},
			},
			contains: []string{
				`Enabled`,
				`Disabled`,
			},
		},
		{
			name: "selected option with check",
			props: Props{
				Open:  true,
				Value: "1",
				Options: []Option{
					{Value: "1", Label: "Selected"},
					{Value: "2", Label: "Not Selected"},
				},
			},
			contains: []string{
				`bg-accent text-accent-foreground`, // Selected styling
				`class="ml-auto h-4 w-4"`,          // Check icon
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := New(tt.props)
			var buf bytes.Buffer
			err := component.Render(&buf)
			if err != nil {
				t.Fatalf("failed to render: %v", err)
			}

			html := buf.String()
			for _, want := range tt.contains {
				if !strings.Contains(html, want) {
					t.Errorf("expected HTML to contain %q, but got:\n%s", want, html)
				}
			}
		})
	}
}

func TestComboboxNotContains(t *testing.T) {
	tests := []struct {
		name        string
		props       Props
		notContains []string
	}{
		{
			name: "closed combobox should not show options",
			props: Props{
				Open: false,
				Options: []Option{
					{Value: "1", Label: "Option 1"},
				},
			},
			notContains: []string{
				`Option 1`, // Options should not be visible
			},
		},
		{
			name: "combobox without value should not show check",
			props: Props{
				Open:  true,
				Value: "",
				Options: []Option{
					{Value: "1", Label: "Option 1"},
				},
			},
			notContains: []string{
				`ml-auto h-4 w-4`, // Check icon should not appear
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := New(tt.props)
			var buf bytes.Buffer
			err := component.Render(&buf)
			if err != nil {
				t.Fatalf("failed to render: %v", err)
			}

			html := buf.String()
			for _, notWant := range tt.notContains {
				if strings.Contains(html, notWant) {
					t.Errorf("expected HTML to NOT contain %q, but got:\n%s", notWant, html)
				}
			}
		})
	}
}