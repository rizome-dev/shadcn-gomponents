package lib

import (
	"strings"
)

// CN combines multiple class strings, filtering out empty strings and duplicates.
// This is similar to the cn() function in shadcn/ui which uses clsx and tailwind-merge.
func CN(classes ...string) string {
	classMap := make(map[string]bool)
	var result []string

	for _, class := range classes {
		// Split by spaces to handle multiple classes in one string
		parts := strings.Fields(class)
		for _, part := range parts {
			if part != "" && !classMap[part] {
				classMap[part] = true
				result = append(result, part)
			}
		}
	}

	return strings.Join(result, " ")
}

// CNIf conditionally includes a class string based on a condition
func CNIf(condition bool, trueClass, falseClass string) string {
	if condition {
		return trueClass
	}
	return falseClass
}

// MergeClasses is a helper to merge Tailwind classes intelligently
// It handles conflicting classes by keeping the last one
func MergeClasses(classes ...string) string {
	// For now, this is a simple implementation
	// In a more complex version, we'd parse Tailwind classes and handle conflicts
	return CN(classes...)
}