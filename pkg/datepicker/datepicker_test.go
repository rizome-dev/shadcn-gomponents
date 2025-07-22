package datepicker

import (
	"bytes"
	"strings"
	"testing"
	"time"

	g "maragu.dev/gomponents"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		props    Props
		contains []string
	}{
		{
			name:  "default date picker",
			props: Props{},
			contains: []string{
				`type="button"`,
				`border bg-background`, // outline variant styling
				`Pick a date`,
				`class="mr-2 h-4 w-4"`, // Calendar icon
			},
		},
		{
			name: "date picker with value",
			props: Props{
				Value: time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC),
			},
			contains: []string{
				`Jan 15, 2024`,
			},
		},
		{
			name: "date picker with custom format",
			props: Props{
				Value:  time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC),
				Format: "2006-01-02",
			},
			contains: []string{
				`2024-01-15`,
			},
		},
		{
			name: "date picker with custom placeholder",
			props: Props{
				Placeholder: "Select a date",
			},
			contains: []string{
				`Select a date`,
			},
		},
		{
			name: "disabled date picker",
			props: Props{
				Disabled: true,
			},
			contains: []string{
				`disabled`,
			},
		},
		{
			name: "required date picker",
			props: Props{
				Required: true,
			},
			contains: []string{
				`required`,
			},
		},
		{
			name: "date picker with ID and name",
			props: Props{
				ID:   "birth-date",
				Name: "birthDate",
			},
			contains: []string{
				`id="birth-date"`,
				`name="birthDate"`,
			},
		},
		{
			name: "date picker with custom class",
			props: Props{
				Class: "custom-datepicker",
			},
			contains: []string{
				`custom-datepicker`,
			},
		},
		{
			name: "open date picker",
			props: Props{
				Open: true,
			},
			contains: []string{
				`data-state="open"`,
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

func TestWithRange(t *testing.T) {
	tests := []struct {
		name     string
		props    RangeProps
		contains []string
	}{
		{
			name:  "default range picker",
			props: RangeProps{},
			contains: []string{
				`Pick a date range`,
				`w-[300px]`,
			},
		},
		{
			name: "range picker with dates",
			props: RangeProps{
				StartDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC),
			},
			contains: []string{
				`Jan 1, 2024 - Jan 31, 2024`,
			},
		},
		{
			name: "range picker with start date only",
			props: RangeProps{
				StartDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			contains: []string{
				`Jan 1, 2024 - ...`,
			},
		},
		{
			name: "disabled range picker",
			props: RangeProps{
				Disabled: true,
			},
			contains: []string{
				`disabled`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := WithRange(tt.props)
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

func TestWithPresets(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name     string
		props    PresetsProps
		contains []string
	}{
		{
			name: "default presets",
			props: PresetsProps{
				Open: true, // Open to show preset content
			},
			contains: []string{
				`Today`,
				`Tomorrow`,
				`In a week`,
				`In a month`,
			},
		},
		{
			name: "custom presets",
			props: PresetsProps{
				Open: true, // Open to show preset content
				Presets: []Preset{
					{Label: "Yesterday", Date: now.AddDate(0, 0, -1)},
					{Label: "Last week", Date: now.AddDate(0, 0, -7)},
				},
			},
			contains: []string{
				`Yesterday`,
				`Last week`,
			},
		},
		{
			name: "preset with selected value",
			props: PresetsProps{
				Open:  true, // Open to show preset content
				Value: now,
				Presets: []Preset{
					{Label: "Today", Date: now},
				},
			},
			contains: []string{
				`bg-accent text-accent-foreground`, // Selected preset styling
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := WithPresets(tt.props)
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

func TestWithInput(t *testing.T) {
	tests := []struct {
		name     string
		props    InputProps
		contains []string
	}{
		{
			name:  "default input date picker",
			props: InputProps{},
			contains: []string{
				`type="text"`,
				`placeholder="YYYY-MM-DD"`,
				`<button type="button"`, // Calendar icon button
			},
		},
		{
			name: "input with label",
			props: InputProps{
				Label: "Birth Date",
			},
			contains: []string{
				`<label`,
				`Birth Date`,
			},
		},
		{
			name: "input with value",
			props: InputProps{
				Value: time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC),
			},
			contains: []string{
				`value="2024-01-15"`,
			},
		},
		{
			name: "input with helper text",
			props: InputProps{
				HelperText: "Please select your birth date",
			},
			contains: []string{
				`text-sm text-muted-foreground`,
				`Please select your birth date`,
			},
		},
		{
			name: "disabled input",
			props: InputProps{
				Disabled: true,
			},
			contains: []string{
				`disabled`,
			},
		},
		{
			name: "required input",
			props: InputProps{
				Required: true,
			},
			contains: []string{
				`required`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := WithInput(tt.props)
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
	testDate := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)

	t.Run("Simple", func(t *testing.T) {
		component := Simple(testDate)
		var buf bytes.Buffer
		err := component.Render(&buf)
		if err != nil {
			t.Fatalf("failed to render: %v", err)
		}

		html := buf.String()
		if !strings.Contains(html, "Jan 15, 2024") {
			t.Errorf("expected Simple date picker to contain formatted date")
		}
	})

	t.Run("Disabled", func(t *testing.T) {
		component := Disabled(testDate)
		var buf bytes.Buffer
		err := component.Render(&buf)
		if err != nil {
			t.Fatalf("failed to render: %v", err)
		}

		html := buf.String()
		if !strings.Contains(html, "disabled") {
			t.Errorf("expected Disabled date picker to be disabled")
		}
	})

	t.Run("WithMinMax", func(t *testing.T) {
		minDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
		maxDate := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)
		component := WithMinMax(testDate, minDate, maxDate)
		var buf bytes.Buffer
		err := component.Render(&buf)
		if err != nil {
			t.Fatalf("failed to render: %v", err)
		}

		// Just ensure it renders without error
		html := buf.String()
		if !strings.Contains(html, "Jan 15, 2024") {
			t.Errorf("expected WithMinMax date picker to contain formatted date")
		}
	})
}

func TestDatePickerNotContains(t *testing.T) {
	tests := []struct {
		name        string
		props       Props
		notContains []string
	}{
		{
			name: "closed date picker should not show calendar",
			props: Props{
				Open: false,
			},
			notContains: []string{
				`role="grid"`, // Calendar grid
			},
		},
		{
			name: "date picker without value should not show date",
			props: Props{
				Value: time.Time{},
			},
			notContains: []string{
				`2024`, // Should not contain any year
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

// renderToString is a helper function to render a component to string
func renderToString(component g.Node) (string, error) {
	var buf bytes.Buffer
	err := component.Render(&buf)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}