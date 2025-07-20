package inputotp

import (
	"fmt"
	"strings"
	"testing"

	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func renderToString(node g.Node) string {
	var buf strings.Builder
	_ = node.Render(&buf)
	return buf.String()
}

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		props    Props
		children []g.Node
		want     []string
		notWant  []string
	}{
		{
			name:  "renders default OTP input",
			props: Props{},
			want: []string{
				`data-otp-container="true"`,
				`data-otp-length="6"`,
				`type="text"`,
				`maxlength="1"`,
				`pattern="[0-9]"`,
				`inputmode="numeric"`,
				`data-otp-input="0"`,
				`placeholder="○"`,
			},
		},
		{
			name: "renders with custom length",
			props: Props{
				Length: 4,
			},
			want: []string{
				`data-otp-length="4"`,
				`data-otp-input="3"`,
			},
			notWant: []string{
				`data-otp-input="4"`,
				`data-otp-input="5"`,
			},
		},
		{
			name: "renders alphanumeric type",
			props: Props{
				Type: "alphanumeric",
			},
			want: []string{
				`pattern="[a-zA-Z0-9]"`,
			},
			notWant: []string{
				`inputmode="numeric"`,
			},
		},
		{
			name: "renders with custom pattern",
			props: Props{
				Pattern: "[A-Z]",
			},
			want: []string{
				`pattern="[A-Z]"`,
			},
		},
		{
			name: "renders with name attribute",
			props: Props{
				Name: "verification",
			},
			want: []string{
				`name="verification[0]"`,
				`name="verification[1]"`,
				`name="verification[5]"`,
			},
		},
		{
			name: "renders with initial value",
			props: Props{
				Value: "123",
			},
			want: []string{
				`value="1"`,
				`value="2"`,
				`value="3"`,
			},
		},
		{
			name: "renders disabled state",
			props: Props{
				Disabled: true,
			},
			want: []string{
				`disabled`,
			},
		},
		{
			name: "renders with autofocus",
			props: Props{
				AutoFocus: true,
			},
			want: []string{
				`autofocus`,
			},
		},
		{
			name: "renders with ID",
			props: Props{
				ID: "otp-input",
			},
			want: []string{
				`id="otp-input"`,
			},
		},
		{
			name: "renders with completion handler",
			props: Props{
				OnComplete: "handleComplete",
			},
			want: []string{
				`data-otp-complete="handleComplete"`,
			},
		},
		{
			name: "renders with custom class",
			props: Props{
				Class: "custom-class",
			},
			want: []string{
				`custom-class`,
			},
		},
		{
			name: "renders with separator for 6 digits",
			props: Props{
				Length: 6,
			},
			want: []string{
				`<div class="mx-1">-</div>`,
			},
		},
		{
			name: "renders with custom placeholder",
			props: Props{
				Placeholder: "●",
			},
			want: []string{
				`placeholder="●"`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := renderToString(New(tt.props, tt.children...))

			for _, want := range tt.want {
				if !strings.Contains(result, want) {
					t.Errorf("New() = %v, want to contain %v", result, want)
				}
			}

			for _, notWant := range tt.notWant {
				if strings.Contains(result, notWant) {
					t.Errorf("New() = %v, don't want to contain %v", result, notWant)
				}
			}
		})
	}
}

func TestDefault(t *testing.T) {
	result := renderToString(Default())

	want := []string{
		`data-otp-length="6"`,
		`pattern="[0-9]"`,
		`inputmode="numeric"`,
	}

	for _, w := range want {
		if !strings.Contains(result, w) {
			t.Errorf("Default() = %v, want to contain %v", result, w)
		}
	}
}

func TestFourDigit(t *testing.T) {
	result := renderToString(FourDigit())

	if !strings.Contains(result, `data-otp-length="4"`) {
		t.Errorf("FourDigit() = %v, want to contain data-otp-length=\"4\"", result)
	}

	// Should not have separator
	if strings.Contains(result, `<div class="mx-1">-</div>`) {
		t.Errorf("FourDigit() = %v, should not contain separator", result)
	}
}

func TestAlphanumeric(t *testing.T) {
	result := renderToString(Alphanumeric())

	if !strings.Contains(result, `pattern="[a-zA-Z0-9]"`) {
		t.Errorf("Alphanumeric() = %v, want to contain alphanumeric pattern", result)
	}

	if strings.Contains(result, `inputmode="numeric"`) {
		t.Errorf("Alphanumeric() = %v, should not contain numeric inputmode", result)
	}
}

func TestGroup(t *testing.T) {
	tests := []struct {
		name     string
		props    GroupProps
		children []g.Node
		want     []string
	}{
		{
			name:  "renders basic group",
			props: GroupProps{},
			want: []string{
				`<div class="flex items-center">`,
			},
		},
		{
			name: "renders with custom class",
			props: GroupProps{
				Class: "gap-4",
			},
			want: []string{
				`class="flex items-center gap-4"`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := renderToString(Group(tt.props, tt.children...))

			for _, want := range tt.want {
				if !strings.Contains(result, want) {
					t.Errorf("Group() = %v, want to contain %v", result, want)
				}
			}
		})
	}
}

func TestSlot(t *testing.T) {
	tests := []struct {
		name  string
		props SlotProps
		want  []string
	}{
		{
			name: "renders basic slot",
			props: SlotProps{
				Index: 0,
			},
			want: []string{
				`data-otp-input="0"`,
				`border-input`,
			},
		},
		{
			name: "renders active slot",
			props: SlotProps{
				Index:    1,
				IsActive: true,
			},
			want: []string{
				`data-otp-input="1"`,
				`border-ring ring-2`,
			},
		},
		{
			name: "renders slot with value",
			props: SlotProps{
				Index:    2,
				HasValue: true,
			},
			want: []string{
				`data-otp-input="2"`,
				`text-foreground`,
			},
		},
		{
			name: "renders with custom class",
			props: SlotProps{
				Index: 3,
				Class: "custom-slot",
			},
			want: []string{
				`custom-slot`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := renderToString(Slot(tt.props))

			for _, want := range tt.want {
				if !strings.Contains(result, want) {
					t.Errorf("Slot() = %v, want to contain %v", result, want)
				}
			}
		})
	}
}

func TestSeparator(t *testing.T) {
	tests := []struct {
		name  string
		props []SeparatorProps
		want  string
	}{
		{
			name:  "renders default separator",
			props: []SeparatorProps{},
			want:  `<div class="mx-1">-</div>`,
		},
		{
			name:  "renders with custom class",
			props: []SeparatorProps{{Class: "mx-2 text-muted"}},
			want:  `<div class="mx-1 mx-2 text-muted">-</div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := renderToString(Separator(tt.props...))

			if !strings.Contains(result, tt.want) {
				t.Errorf("Separator() = %v, want to contain %v", result, tt.want)
			}
		})
	}
}

func TestWithSeparator(t *testing.T) {
	result := renderToString(WithSeparator())

	// Should have 6 slots
	for i := 0; i < 6; i++ {
		want := fmt.Sprintf(`data-otp-input="%d"`, i)
		if !strings.Contains(result, want) {
			t.Errorf("WithSeparator() = %v, want to contain %v", result, want)
		}
	}

	// Should have separator
	if !strings.Contains(result, `<div class="mx-1">-</div>`) {
		t.Errorf("WithSeparator() = %v, want to contain separator", result)
	}

	// Should have two groups
	if strings.Count(result, `<div class="flex items-center">`) != 2 {
		t.Errorf("WithSeparator() should have exactly 2 groups")
	}
}

func TestJavaScriptBehavior(t *testing.T) {
	result := renderToString(New(Props{Length: 4}))

	// Test that JavaScript handlers are present
	jsHandlers := []string{
		"oninput=",
		"onkeydown=",
		"onpaste=",
		"checkValidity()",
		"Backspace",
		"ArrowLeft",
		"ArrowRight",
		"clipboardData",
	}

	for _, handler := range jsHandlers {
		if !strings.Contains(result, handler) {
			t.Errorf("New() = %v, want to contain JavaScript handler: %v", result, handler)
		}
	}
}

func TestAutocompleteAttribute(t *testing.T) {
	result := renderToString(Default())

	if !strings.Contains(result, `autocomplete="one-time-code"`) {
		t.Errorf("Default() = %v, want to contain autocomplete attribute", result)
	}
}