package switchcomp_test

import (
	"strings"
	"testing"
	g "maragu.dev/gomponents"
	"github.com/rizome-dev/shadcn-gomponents/components/switch"
)

func TestSwitch(t *testing.T) {
	tests := []struct {
		name     string
		switch_  g.Node
		contains []string
	}{
		{
			name:     "default switch",
			switch_:  switchcomp.Default(),
			contains: []string{
				`type="checkbox"`,
				`class="sr-only peer"`,
				`class="peer inline-flex shrink-0 cursor-pointer`,
				`h-6 w-11`, // default track size
				`h-5 w-5`,  // default thumb size
				`role="switch"`,
				`aria-checked="false"`,
				`<style>`,
				`transform: translateX(0)`,
			},
		},
		{
			name:     "checked switch",
			switch_:  switchcomp.Checked(),
			contains: []string{
				`checked`,
				`aria-checked="true"`,
				`transform: translate-x-5`, // checked position
			},
		},
		{
			name:     "disabled switch",
			switch_:  switchcomp.Disabled(false),
			contains: []string{
				`disabled`,
				`aria-disabled="true"`,
				`disabled:cursor-not-allowed disabled:opacity-50`,
			},
		},
		{
			name: "switch with all props",
			switch_: switchcomp.New(switchcomp.Props{
				ID:       "notifications",
				Name:     "notifications",
				Value:    "on",
				Checked:  true,
				Required: true,
				OnChange: "handleToggle(this)",
			}),
			contains: []string{
				`id="notifications"`,
				`name="notifications"`,
				`value="on"`,
				`checked`,
				`required`,
				`onchange="handleToggle(this)"`,
			},
		},
		{
			name:     "small switch",
			switch_:  switchcomp.Small(),
			contains: []string{
				`h-5 w-9`, // small track size
				`h-4 w-4`, // small thumb size
				`transform: translate-x-4`,
			},
		},
		{
			name:     "large switch",
			switch_:  switchcomp.Large(),
			contains: []string{
				`h-7 w-14`, // large track size
				`h-6 w-6`,  // large thumb size
				`transform: translate-x-7`,
			},
		},
		{
			name: "switch with label",
			switch_: switchcomp.WithLabel(
				switchcomp.Props{Name: "airplane"},
				"Airplane Mode",
			),
			contains: []string{
				`<label for="switch-Airplane Mode"`,
				`Airplane Mode</label>`,
				`id="switch-Airplane Mode"`,
				`class="text-sm font-medium leading-none`,
			},
		},
		{
			name: "switch form field",
			switch_: switchcomp.FormField(
				switchcomp.Props{
					Name:    "marketing",
					Checked: true,
				},
				"Marketing emails",
				"Receive emails about new products, features, and more.",
			),
			contains: []string{
				`class="flex flex-row items-center justify-between rounded-lg border p-4"`,
				`<label for="switch-field"`,
				`Marketing emails</label>`,
				`<p class="text-sm text-muted-foreground">Receive emails about`,
			},
		},
		{
			name: "switch setting",
			switch_: switchcomp.Setting(
				"notifications",
				"Push Notifications",
				"Enable push notifications on this device",
				false,
			),
			contains: []string{
				`name="notifications"`,
				`Push Notifications`,
				`Enable push notifications`,
				`aria-checked="false"`,
			},
		},
		{
			name: "switch with custom class",
			switch_: switchcomp.New(switchcomp.Props{
				Class: "custom-switch",
			}),
			contains: []string{
				`custom-switch`,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.switch_.String()
			for _, expected := range test.contains {
				if !strings.Contains(result, expected) {
					t.Errorf("expected result to contain %q, but it didn't.\nGot: %s", expected, result)
				}
			}
		})
	}
}