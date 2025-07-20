package lib

import "testing"

func TestCN(t *testing.T) {
	tests := []struct {
		name     string
		classes  []string
		expected string
	}{
		{
			name:     "single class",
			classes:  []string{"foo"},
			expected: "foo",
		},
		{
			name:     "multiple classes",
			classes:  []string{"foo", "bar"},
			expected: "foo bar",
		},
		{
			name:     "duplicate classes",
			classes:  []string{"foo", "bar", "foo"},
			expected: "foo bar",
		},
		{
			name:     "empty strings",
			classes:  []string{"foo", "", "bar"},
			expected: "foo bar",
		},
		{
			name:     "classes with spaces",
			classes:  []string{"foo bar", "baz"},
			expected: "foo bar baz",
		},
		{
			name:     "all empty",
			classes:  []string{"", "", ""},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CN(tt.classes...)
			if result != tt.expected {
				t.Errorf("CN() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestCNIf(t *testing.T) {
	tests := []struct {
		name       string
		condition  bool
		trueClass  string
		falseClass string
		expected   string
	}{
		{
			name:       "condition true",
			condition:  true,
			trueClass:  "active",
			falseClass: "inactive",
			expected:   "active",
		},
		{
			name:       "condition false",
			condition:  false,
			trueClass:  "active",
			falseClass: "inactive",
			expected:   "inactive",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CNIf(tt.condition, tt.trueClass, tt.falseClass)
			if result != tt.expected {
				t.Errorf("CNIf() = %v, want %v", result, tt.expected)
			}
		})
	}
}