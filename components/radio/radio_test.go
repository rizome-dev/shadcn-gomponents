package radio_test

import (
	"strings"
	"testing"
	g "maragu.dev/gomponents"
	"github.com/rizome-dev/shadcn-gomponents/components/radio"
)

func TestRadioGroup(t *testing.T) {
	tests := []struct {
		name     string
		radio    g.Node
		contains []string
	}{
		{
			name: "basic radio group",
			radio: radio.Group(
				radio.GroupProps{Name: "plan"},
				radio.WithLabel(
					radio.ItemProps{Value: "free"},
					"plan",
					"Free",
				),
				radio.WithLabel(
					radio.ItemProps{Value: "pro", Checked: true},
					"plan",
					"Pro",
				),
			),
			contains: []string{
				`role="radiogroup"`,
				`data-name="plan"`,
				`class="grid gap-2"`,
				`name="plan"`,
				`value="free"`,
				`value="pro"`,
				`checked`,
				`data-state="checked"`,
			},
		},
		{
			name: "radio item with ID",
			radio: radio.Item(
				radio.ItemProps{
					ID:    "r1",
					Value: "option1",
				},
				"group1",
			),
			contains: []string{
				`type="radio"`,
				`id="r1"`,
				`name="group1"`,
				`value="option1"`,
				`class="aspect-square h-4 w-4 rounded-full border`,
			},
		},
		{
			name: "checked radio with indicator",
			radio: radio.Item(
				radio.ItemProps{
					Value:   "selected",
					Checked: true,
				},
				"group2",
			),
			contains: []string{
				`checked`,
				`data-state="checked"`,
				`<svg viewBox="0 0 16 16"`,
				`<circle cx="8" cy="8" r="3"`,
			},
		},
		{
			name: "disabled radio",
			radio: radio.WithLabel(
				radio.ItemProps{
					Value:    "disabled-opt",
					Disabled: true,
				},
				"group3",
				"Disabled Option",
			),
			contains: []string{
				`disabled`,
				`disabled:cursor-not-allowed disabled:opacity-50`,
				`peer-disabled:cursor-not-allowed peer-disabled:opacity-70`,
			},
		},
		{
			name: "simple radio group",
			radio: radio.Simple(
				"color",
				[]radio.Option{
					{Value: "red", Label: "Red"},
					{Value: "blue", Label: "Blue"},
					{Value: "green", Label: "Green"},
				},
				"blue",
			),
			contains: []string{
				`data-name="color"`,
				`data-default-value="blue"`,
				`value="red"`,
				`value="blue"`,
				`value="green"`,
				`id="radio-blue"`,
				`<label for="radio-blue"`,
			},
		},
		{
			name: "horizontal radio group",
			radio: radio.Horizontal(
				"size",
				[]radio.Option{
					{Value: "sm", Label: "Small"},
					{Value: "md", Label: "Medium"},
					{Value: "lg", Label: "Large"},
				},
				"md",
			),
			contains: []string{
				`class="flex items-center gap-4"`,
				`value="sm"`,
				`value="md"`,
				`value="lg"`,
			},
		},
		{
			name: "radio form field",
			radio: radio.FormField(
				"subscription",
				"Subscription Plan",
				"Choose your subscription tier",
				[]radio.Option{
					{Value: "basic", Label: "Basic - $9/month"},
					{Value: "premium", Label: "Premium - $19/month"},
					{Value: "enterprise", Label: "Enterprise - Contact us"},
				},
				"basic",
			),
			contains: []string{
				`<label class="text-sm font-medium">Subscription Plan</label>`,
				`<p class="text-sm text-muted-foreground">Choose your subscription tier</p>`,
				`value="basic"`,
				`value="premium"`,
				`value="enterprise"`,
			},
		},
		{
			name: "card style radio",
			radio: radio.Card(
				radio.ItemProps{Value: "standard"},
				"shipping",
				"Standard Shipping",
				"5-7 business days",
			),
			contains: []string{
				`class="relative flex cursor-pointer rounded-lg border p-4`,
				`id="radio-card-standard"`,
				`<p class="font-medium leading-none">Standard Shipping</p>`,
				`<p class="text-sm text-muted-foreground">5-7 business days</p>`,
			},
		},
		{
			name: "selected card style radio",
			radio: radio.Card(
				radio.ItemProps{Value: "express", Checked: true},
				"shipping",
				"Express Shipping",
				"1-2 business days",
			),
			contains: []string{
				`border-primary bg-accent`,
				`checked`,
				`Express Shipping`,
			},
		},
		{
			name: "radio with custom class",
			radio: radio.Item(
				radio.ItemProps{
					Value: "custom",
					Class: "custom-radio",
				},
				"group4",
			),
			contains: []string{
				`custom-radio`,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.radio.String()
			for _, expected := range test.contains {
				if !strings.Contains(result, expected) {
					t.Errorf("expected result to contain %q, but it didn't.\nGot: %s", expected, result)
				}
			}
		})
	}
}