package selector_test

import (
	"strings"
	"testing"
	g "maragu.dev/gomponents"
	"github.com/rizome-dev/shadcn-gomponents/pkg/selector"
)

func renderToString(node g.Node) string {
	var buf strings.Builder
	err := node.Render(&buf)
	if err != nil {
		panic(err) // For tests, panic on render error
	}
	return buf.String()
}

func TestSelect(t *testing.T) {
	tests := []struct {
		name     string
		sel      g.Node
		contains []string
	}{
		{
			name: "basic select",
			sel: selector.New(selector.Props{
				Name: "country",
				Options: []selector.OptionType{
					{Value: "us", Label: "United States"},
					{Value: "uk", Label: "United Kingdom"},
					{Value: "ca", Label: "Canada"},
				},
			}),
			contains: []string{
				`<select`,
				`name="country"`,
				`class="flex w-full rounded-md border`,
				`<option value="us">United States</option>`,
				`<option value="uk">United Kingdom</option>`,
				`<option value="ca">Canada</option>`,
			},
		},
		{
			name: "select with value",
			sel: selector.Simple(
				"size",
				[]selector.OptionType{
					{Value: "sm", Label: "Small"},
					{Value: "md", Label: "Medium"},
					{Value: "lg", Label: "Large"},
				},
				"md",
			),
			contains: []string{
				`value="md"`,
				`<option value="md" selected>Medium</option>`,
			},
		},
		{
			name: "select with placeholder",
			sel: selector.WithPlaceholder(
				"category",
				"Choose a category",
				[]selector.OptionType{
					{Value: "electronics", Label: "Electronics"},
					{Value: "clothing", Label: "Clothing"},
				},
			),
			contains: []string{
				`<option value="" selected disabled hidden="">Choose a category</option>`,
				`value="electronics"`,
				`value="clothing"`,
			},
		},
		{
			name: "select with groups",
			sel: selector.WithGroups(
				"timezone",
				[]selector.Group{
					{
						Label: "North America",
						Options: []selector.OptionType{
							{Value: "pst", Label: "Pacific Time"},
							{Value: "est", Label: "Eastern Time"},
						},
					},
					{
						Label: "Europe",
						Options: []selector.OptionType{
							{Value: "gmt", Label: "GMT"},
							{Value: "cet", Label: "Central European Time"},
						},
					},
				},
				"pst",
			),
			contains: []string{
				`<optgroup label="North America"`,
				`<option value="pst" selected>Pacific Time</option>`,
				`<optgroup label="Europe"`,
				`<option value="gmt">GMT</option>`,
			},
		},
		{
			name: "disabled select",
			sel: selector.New(selector.Props{
				Name:     "disabled-select",
				Disabled: true,
				Options: []selector.OptionType{
					{Value: "opt1", Label: "Option 1"},
				},
			}),
			contains: []string{
				`disabled`,
				`disabled:cursor-not-allowed disabled:opacity-50`,
			},
		},
		{
			name: "select with all props",
			sel: selector.New(selector.Props{
				ID:       "my-select",
				Name:     "selection",
				Value:    "b",
				Required: true,
				OnChange: "handleChange(this)",
				Options: []selector.OptionType{
					{Value: "a", Label: "Option A"},
					{Value: "b", Label: "Option B"},
					{Value: "c", Label: "Option C", Disabled: true},
				},
			}),
			contains: []string{
				`id="my-select"`,
				`name="selection"`,
				`value="b"`,
				`required`,
				`onchange="handleChange(this)"`,
				`<option value="b" selected>Option B</option>`,
				`<option value="c" disabled>Option C</option>`,
			},
		},
		{
			name: "select form field",
			sel: selector.FormField(
				selector.Props{
					Name: "department",
					Options: []selector.OptionType{
						{Value: "eng", Label: "Engineering"},
						{Value: "sales", Label: "Sales"},
						{Value: "hr", Label: "Human Resources"},
					},
				},
				"Department",
				"Select your department",
			),
			contains: []string{
				`<label for="select-department"`,
				`Department</label>`,
				`<p class="text-sm text-muted-foreground">Select your department</p>`,
				`name="department"`,
			},
		},
		{
			name: "multiple select",
			sel: selector.New(selector.Props{
				Name:     "skills",
				Multiple: true,
				Options: []selector.OptionType{
					{Value: "js", Label: "JavaScript"},
					{Value: "go", Label: "Go"},
					{Value: "py", Label: "Python"},
				},
			}),
			contains: []string{
				`multiple`,
			},
		},
		{
			name: "small size select",
			sel: selector.New(selector.Props{
				Name: "small",
				Size: "sm",
				Options: []selector.OptionType{
					{Value: "1", Label: "One"},
				},
			}),
			contains: []string{
				`h-9 px-3 py-1.5 text-sm`,
			},
		},
		{
			name: "large size select",
			sel: selector.New(selector.Props{
				Name: "large",
				Size: "lg",
				Options: []selector.OptionType{
					{Value: "1", Label: "One"},
				},
			}),
			contains: []string{
				`h-11 px-4 py-2.5`,
			},
		},
		{
			name: "custom styled select",
			sel: selector.Custom(selector.Props{
				Name:        "custom",
				Placeholder: "Select an option",
				Value:       "opt2",
				Options: []selector.OptionType{
					{Value: "opt1", Label: "First Option"},
					{Value: "opt2", Label: "Second Option"},
				},
			}),
			contains: []string{
				`<span class="line-clamp-1">Second Option</span>`,
				`<svg`, // Chevron icon
				`polyline`,
				`opacity-0 cursor-pointer`, // Hidden select
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := renderToString(test.sel)
			for _, expected := range test.contains {
				if !strings.Contains(result, expected) {
					t.Errorf("expected result to contain %q, but it didn't.\nGot: %s", expected, result)
				}
			}
		})
	}
}