package slider

import (
	"bytes"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		props    Props
		contains []string
	}{
		{
			name:  "default slider",
			props: Props{},
			contains: []string{
				`data-slider="true"`,
				`data-orientation="horizontal"`,
				`data-slider-track="true"`,
				`data-slider-range="true"`,
				`data-slider-thumb="true"`,
				`aria-valuemin="0"`,
				`aria-valuemax="100"`,
				`aria-valuenow="0"`,
			},
		},
		{
			name: "slider with custom range",
			props: Props{
				Min:   10,
				Max:   50,
				Value: []int{25},
			},
			contains: []string{
				`aria-valuemin="10"`,
				`aria-valuemax="50"`,
				`aria-valuenow="25"`,
			},
		},
		{
			name: "range slider",
			props: Props{
				Value: []int{20, 80},
			},
			contains: []string{
				`data-index="0"`,
				`data-index="1"`,
				`aria-valuenow="20"`,
				`aria-valuenow="80"`,
			},
		},
		{
			name: "vertical slider",
			props: Props{
				Orientation: "vertical",
				Value:       []int{50},
			},
			contains: []string{
				`data-orientation="vertical"`,
				`aria-orientation="vertical"`,
				`bottom: calc(50.00%`,
			},
		},
		{
			name: "disabled slider",
			props: Props{
				Disabled: true,
			},
			contains: []string{
				`data-disabled="true"`,
				`disabled="true"`,
			},
		},
		{
			name: "slider with step",
			props: Props{
				Step:  5,
				Value: []int{15},
			},
			contains: []string{
				`aria-valuenow="15"`,
			},
		},
		{
			name: "slider with hidden inputs",
			props: Props{
				Name:  "volume",
				Value: []int{75},
				ID:    "volume-slider",
			},
			contains: []string{
				`<input type="hidden"`,
				`name="volume[0]"`,
				`value="75"`,
				`id="volume-slider-input-0"`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			component := New(tt.props)
			err := component.Render(&buf)
			if err != nil {
				t.Fatalf("Render() error = %v", err)
			}

			result := buf.String()
			for _, expected := range tt.contains {
				if !strings.Contains(result, expected) {
					t.Errorf("Expected output to contain %q, but it didn't. Got:\n%s", expected, result)
				}
			}
		})
	}
}

func TestSingle(t *testing.T) {
	// Test that Single ensures only one value
	props := Props{
		Value: []int{20, 40, 60}, // Multiple values provided
	}
	
	var buf bytes.Buffer
	component := Single(props)
	err := component.Render(&buf)
	if err != nil {
		t.Fatalf("Render() error = %v", err)
	}

	result := buf.String()
	
	// Should only have one thumb
	thumbCount := strings.Count(result, `data-slider-thumb="true"`)
	if thumbCount != 1 {
		t.Errorf("Expected exactly 1 thumb, but got %d", thumbCount)
	}
	
	// Should have the first value
	if !strings.Contains(result, `aria-valuenow="20"`) {
		t.Errorf("Expected to contain first value (20)")
	}
}

func TestRange(t *testing.T) {
	tests := []struct {
		name     string
		props    Props
		expected []string
	}{
		{
			name: "range with no values",
			props: Props{
				Min: 0,
				Max: 100,
			},
			expected: []string{
				`aria-valuenow="0"`,
				`aria-valuenow="100"`,
			},
		},
		{
			name: "range with one value",
			props: Props{
				Value: []int{30},
			},
			expected: []string{
				`aria-valuenow="30"`,
				`aria-valuenow="100"`,
			},
		},
		{
			name: "range with reversed values",
			props: Props{
				Value: []int{80, 20}, // Should be swapped
			},
			expected: []string{
				`aria-valuenow="20"`,
				`aria-valuenow="80"`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			component := Range(tt.props)
			err := component.Render(&buf)
			if err != nil {
				t.Fatalf("Render() error = %v", err)
			}

			result := buf.String()
			
			// Should have exactly 2 thumbs
			thumbCount := strings.Count(result, `data-slider-thumb="true"`)
			if thumbCount != 2 {
				t.Errorf("Expected exactly 2 thumbs, but got %d", thumbCount)
			}
			
			for _, expected := range tt.expected {
				if !strings.Contains(result, expected) {
					t.Errorf("Expected output to contain %q, but it didn't", expected)
				}
			}
		})
	}
}

func TestVertical(t *testing.T) {
	props := Props{
		Value: []int{50},
	}
	
	var buf bytes.Buffer
	component := Vertical(props)
	err := component.Render(&buf)
	if err != nil {
		t.Fatalf("Render() error = %v", err)
	}

	result := buf.String()
	
	expected := []string{
		`data-orientation="vertical"`,
		`aria-orientation="vertical"`,
		`data-[orientation=vertical]:h-full`,
		`bottom: calc(50.00%`,
	}
	
	for _, exp := range expected {
		if !strings.Contains(result, exp) {
			t.Errorf("Expected output to contain %q, but it didn't", exp)
		}
	}
}

func TestWithLabels(t *testing.T) {
	props := Props{
		Min: 0,
		Max: 100,
	}
	
	var buf bytes.Buffer
	component := WithLabels(props)
	err := component.Render(&buf)
	if err != nil {
		t.Fatalf("Render() error = %v", err)
	}

	result := buf.String()
	
	expected := []string{
		`>0</span>`,
		`>100</span>`,
		`text-muted-foreground`,
		`data-slider="true"`,
	}
	
	for _, exp := range expected {
		if !strings.Contains(result, exp) {
			t.Errorf("Expected output to contain %q, but it didn't", exp)
		}
	}
}

func TestWithValue(t *testing.T) {
	tests := []struct {
		name     string
		props    Props
		expected string
	}{
		{
			name: "single value display",
			props: Props{
				Value: []int{42},
				ID:    "test-slider",
			},
			expected: `>42</span>`,
		},
		{
			name: "range value display",
			props: Props{
				Value: []int{10, 90},
			},
			expected: `>10 - 90</span>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			component := WithValue(tt.props)
			err := component.Render(&buf)
			if err != nil {
				t.Fatalf("Render() error = %v", err)
			}

			result := buf.String()
			
			if !strings.Contains(result, tt.expected) {
				t.Errorf("Expected output to contain %q, but it didn't", tt.expected)
			}
			
			if !strings.Contains(result, `>Value</label>`) {
				t.Errorf("Expected to contain value label")
			}
			
			if tt.props.ID != "" && !strings.Contains(result, `id="test-slider-value"`) {
				t.Errorf("Expected value display to have correct ID")
			}
		})
	}
}

func TestWithTicks(t *testing.T) {
	props := Props{
		Value: []int{50},
	}
	
	var buf bytes.Buffer
	component := WithTicks(props, 5)
	err := component.Render(&buf)
	if err != nil {
		t.Fatalf("Render() error = %v", err)
	}

	result := buf.String()
	
	// Should have 5 tick marks
	tickCount := strings.Count(result, `bg-border`)
	if tickCount != 5 {
		t.Errorf("Expected 5 tick marks, but got %d", tickCount)
	}
	
	// Check for tick positioning
	expected := []string{
		`left: 0.00%`,
		`left: 25.00%`,
		`left: 50.00%`,
		`left: 75.00%`,
		`left: 100.00%`,
	}
	
	for _, exp := range expected {
		if !strings.Contains(result, exp) {
			t.Errorf("Expected output to contain tick at %q, but it didn't", exp)
		}
	}
}