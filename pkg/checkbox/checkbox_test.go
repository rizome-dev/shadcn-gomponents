package checkbox_test

import (
	"strings"
	"testing"
	g "maragu.dev/gomponents"
	"github.com/rizome-dev/shadcn-gomponents/pkg/checkbox"
)

func TestCheckbox(t *testing.T) {
	tests := []struct {
		name     string
		checkbox g.Node
		contains []string
		notContains []string
	}{
		{
			name:     "default checkbox",
			checkbox: checkbox.Default(),
			contains: []string{
				`type="checkbox"`,
				`class="peer h-4 w-4 shrink-0 rounded-sm border`,
				`data-state="unchecked"`,
				`role="checkbox"`,
				`aria-checked="false"`,
			},
			notContains: []string{
				`checked`,
				`<svg`,
			},
		},
		{
			name:     "checked checkbox",
			checkbox: checkbox.Checked(),
			contains: []string{
				`checked`,
				`data-state="checked"`,
				`aria-checked="true"`,
				`<svg`,
				`polyline`, // Check icon
			},
		},
		{
			name:     "disabled checkbox",
			checkbox: checkbox.Disabled(false),
			contains: []string{
				`disabled`,
				`aria-disabled="true"`,
				`disabled:cursor-not-allowed disabled:opacity-50`,
			},
		},
		{
			name: "checkbox with props",
			checkbox: checkbox.New(checkbox.Props{
				ID:       "terms",
				Name:     "terms",
				Value:    "agree",
				Checked:  true,
				Required: true,
			}),
			contains: []string{
				`id="terms"`,
				`name="terms"`,
				`value="agree"`,
				`checked`,
				`required`,
				`data-state="checked"`,
			},
		},
		{
			name: "indeterminate checkbox",
			checkbox: checkbox.New(checkbox.Props{
				Indeterminate: true,
			}),
			contains: []string{
				`data-indeterminate="true"`,
				`<svg`, // Should show minus icon
				`<line`, // Minus icon line
			},
		},
		{
			name: "checkbox with label",
			checkbox: checkbox.WithLabel(
				checkbox.Props{Name: "newsletter"},
				"Subscribe to newsletter",
			),
			contains: []string{
				`<label`,
				`for="checkbox-Subscribe to newsletter"`,
				`Subscribe to newsletter</label>`,
				`class="text-sm font-medium leading-none`,
			},
		},
		{
			name: "checkbox form field",
			checkbox: checkbox.FormField(
				checkbox.Props{
					ID:   "marketing",
					Name: "marketing",
				},
				"Marketing emails",
				"Receive emails about new products and features.",
			),
			contains: []string{
				`id="marketing"`,
				`<label for="marketing"`,
				`Marketing emails</label>`,
				`<p class="text-sm text-muted-foreground">Receive emails about new products`,
			},
		},
		{
			name: "checkbox with custom class",
			checkbox: checkbox.New(checkbox.Props{
				Class: "custom-checkbox",
			}),
			contains: []string{
				`custom-checkbox`,
			},
		},
		{
			name: "checkbox with onChange",
			checkbox: checkbox.New(checkbox.Props{
				OnChange: "handleCheckboxChange(this)",
			}),
			contains: []string{
				`onchange="handleCheckboxChange(this)"`,
			},
		},
		{
			name: "checkbox group",
			checkbox: checkbox.Group(
				"Notifications",
				checkbox.WithLabel(checkbox.Props{Name: "email"}, "Email"),
				checkbox.WithLabel(checkbox.Props{Name: "sms"}, "SMS"),
				checkbox.WithLabel(checkbox.Props{Name: "push"}, "Push"),
			),
			contains: []string{
				`<h3 class="text-sm font-medium">Notifications</h3>`,
				`name="email"`,
				`name="sms"`,
				`name="push"`,
				`class="space-y-3"`,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.checkbox.String()
			
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