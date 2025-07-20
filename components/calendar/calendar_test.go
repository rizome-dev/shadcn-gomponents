package calendar

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

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
			name: "default calendar",
			props: Props{
				Month: time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC),
			},
			contains: []string{
				`class="bg-background p-3 rounded-lg border"`,
				`January 2024`,
				`<h2`,
				`Su`, `Mo`, `Tu`, `We`, `Th`, `Fr`, `Sa`,
			},
		},
		{
			name: "calendar with selected date",
			props: Props{
				Value: time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC),
				Month: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			contains: []string{
				`15`,
				`aria-selected="true"`,
				`bg-primary text-primary-foreground`,
			},
		},
		{
			name: "calendar without weekday headers",
			props: Props{
				Month:    time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				ShowDays: false,
			},
			contains: []string{
				`January 2024`,
			},
		},
		{
			name: "calendar with custom class",
			props: Props{
				Month: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				Class: "custom-calendar",
			},
			contains: []string{
				`custom-calendar`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			calendar := New(tt.props, tt.children...)
			err := calendar.Render(&buf)
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

func TestCalendarDay(t *testing.T) {
	tests := []struct {
		name     string
		props    DayProps
		contains []string
		notContains []string
	}{
		{
			name: "regular day",
			props: DayProps{
				Date: time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC),
			},
			contains: []string{
				`15`,
				`type="button"`,
				`aria-label="Select January 15, 2024"`,
			},
		},
		{
			name: "selected day",
			props: DayProps{
				Date:     time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC),
				Selected: true,
			},
			contains: []string{
				`15`,
				`bg-primary text-primary-foreground`,
				`aria-selected="true"`,
			},
		},
		{
			name: "today",
			props: DayProps{
				Date:  time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC),
				Today: true,
			},
			contains: []string{
				`15`,
				`bg-accent text-accent-foreground`,
			},
		},
		{
			name: "disabled day",
			props: DayProps{
				Date:     time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC),
				Disabled: true,
			},
			contains: []string{
				`15`,
				`disabled`,
				`opacity-50 pointer-events-none`,
			},
		},
		{
			name: "outside month",
			props: DayProps{
				Date:    time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC),
				Outside: true,
			},
			contains: []string{
				`15`,
				`text-muted-foreground opacity-50`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			day := CalendarDay(tt.props)
			err := day.Render(&buf)
			if err != nil {
				t.Fatalf("Render() error = %v", err)
			}

			result := buf.String()
			for _, expected := range tt.contains {
				if !strings.Contains(result, expected) {
					t.Errorf("Expected output to contain %q, but it didn't.\nGot: %s", expected, result)
				}
			}
			for _, notExpected := range tt.notContains {
				if strings.Contains(result, notExpected) {
					t.Errorf("Expected output NOT to contain %q, but it did.\nGot: %s", notExpected, result)
				}
			}
		})
	}
}

func TestMonthPicker(t *testing.T) {
	currentMonth := time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC)
	
	var buf bytes.Buffer
	picker := MonthPicker(currentMonth)
	err := picker.Render(&buf)
	if err != nil {
		t.Fatalf("Render() error = %v", err)
	}

	result := buf.String()
	
	// Check that all months are present
	months := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	for _, month := range months {
		if !strings.Contains(result, month) {
			t.Errorf("Expected month %q to be present, but it wasn't", month)
		}
	}
	
	// Check that March is selected
	if !strings.Contains(result, `bg-primary text-primary-foreground`) {
		t.Error("Expected selected month to have primary styling")
	}
}

func TestYearPicker(t *testing.T) {
	currentYear := 2024
	
	var buf bytes.Buffer
	picker := YearPicker(currentYear, 2020, 2025)
	err := picker.Render(&buf)
	if err != nil {
		t.Fatalf("Render() error = %v", err)
	}

	result := buf.String()
	
	// Check that years are present
	for year := 2020; year <= 2025; year++ {
		yearStr := strings.TrimSpace(fmt.Sprintf("%d", year))
		if !strings.Contains(result, yearStr) {
			t.Errorf("Expected year %d to be present, but it wasn't", year)
		}
	}
	
	// Check that 2024 is selected
	if !strings.Contains(result, `bg-primary text-primary-foreground`) {
		t.Error("Expected selected year to have primary styling")
	}
}

func TestDateRangeCalendar(t *testing.T) {
	startDate := time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 1, 20, 0, 0, 0, 0, time.UTC)
	month := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	
	var buf bytes.Buffer
	calendar := DateRangeCalendar(startDate, endDate, month)
	err := calendar.Render(&buf)
	if err != nil {
		t.Fatalf("Render() error = %v", err)
	}

	result := buf.String()
	
	// Check for range styling
	if !strings.Contains(result, `rounded-l-md`) {
		t.Error("Expected start date to have rounded-l-md class")
	}
	if !strings.Contains(result, `rounded-r-md`) {
		t.Error("Expected end date to have rounded-r-md class")
	}
	if !strings.Contains(result, `bg-accent rounded-none`) {
		t.Error("Expected dates in range to have accent background")
	}
}

func TestHTMXCalendar(t *testing.T) {
	htmxProps := HTMXProps{
		ID:           "test-calendar",
		NavigatePath: "/api/calendar/navigate",
		SelectPath:   "/api/calendar/select",
		UpdatePath:   "/api/calendar/update",
	}
	
	props := Props{
		Month: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	
	var buf bytes.Buffer
	calendar := NewHTMX(props, htmxProps)
	err := calendar.Render(&buf)
	if err != nil {
		t.Fatalf("Render() error = %v", err)
	}

	result := buf.String()
	
	// Check for HTMX attributes
	if !strings.Contains(result, `id="test-calendar"`) {
		t.Error("Expected calendar to have the specified ID")
	}
	if !strings.Contains(result, `hx-get="/api/calendar/navigate`) {
		t.Error("Expected navigation buttons to have hx-get attribute")
	}
	if !strings.Contains(result, `hx-target="#test-calendar"`) {
		t.Error("Expected HTMX target to be set")
	}
	if !strings.Contains(result, `hx-swap="outerHTML"`) {
		t.Error("Expected HTMX swap to be outerHTML")
	}
}

func TestIsSameDay(t *testing.T) {
	tests := []struct {
		name     string
		date1    time.Time
		date2    time.Time
		expected bool
	}{
		{
			name:     "same day",
			date1:    time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC),
			date2:    time.Date(2024, 1, 15, 14, 45, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "different day",
			date1:    time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC),
			date2:    time.Date(2024, 1, 16, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "different month",
			date1:    time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC),
			date2:    time.Date(2024, 2, 15, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "different year",
			date1:    time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC),
			date2:    time.Date(2025, 1, 15, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isSameDay(tt.date1, tt.date2)
			if result != tt.expected {
				t.Errorf("isSameDay(%v, %v) = %v, want %v", tt.date1, tt.date2, result, tt.expected)
			}
		})
	}
}