package command

import (
	"bytes"
	"strings"
	"testing"

	h "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib/icons"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		id       string
		groups   []CommandGroup
		opts     []Option
		contains []string
		notContains []string
	}{
		{
			name: "basic command menu",
			id:   "test-command",
			groups: []CommandGroup{
				{
					Label: "Test Group",
					Items: []CommandItem{
						{Value: "item1", Label: "Item 1"},
						{Value: "item2", Label: "Item 2"},
					},
				},
			},
			contains: []string{
				`id="test-command"`,
				`data-command="true"`,
				`role="combobox"`,
				`Test Group`,
				`Item 1`,
				`Item 2`,
				`Type a command or search...`,
			},
		},
		{
			name: "without search",
			id:   "no-search",
			groups: []CommandGroup{
				{
					Items: []CommandItem{
						{Value: "item1", Label: "Item 1"},
					},
				},
			},
			opts: []Option{WithoutSearch()},
			contains: []string{
				`Item 1`,
			},
			notContains: []string{
				`Type a command or search...`,
				`data-command-input`,
			},
		},
		{
			name: "with custom placeholder",
			id:   "custom-placeholder",
			groups: []CommandGroup{},
			opts: []Option{WithPlaceholder("Search commands...")},
			contains: []string{
				`Search commands...`,
			},
		},
		{
			name:   "empty state",
			id:     "empty",
			groups: []CommandGroup{},
			opts:   []Option{WithEmptyMessage("No items found")},
			contains: []string{
				`data-command-empty`,
			},
		},
		{
			name: "with custom dimensions",
			id:   "custom-size",
			groups: []CommandGroup{},
			opts: []Option{
				WithWidth("600px"),
				WithMaxHeight("500px"),
			},
			contains: []string{
				`width: 600px`,
				`max-height: 500px`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			component := New(tt.id, tt.groups, tt.opts...)
			err := component.Render(&buf)
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

func TestCommandItem(t *testing.T) {
	tests := []struct {
		name     string
		item     CommandItem
		cfg      *config
		contains []string
		notContains []string
	}{
		{
			name: "basic item",
			item: CommandItem{
				Value: "test",
				Label: "Test Item",
			},
			cfg: &config{showShortcuts: true},
			contains: []string{
				`data-command-item="true"`,
				`data-value="test"`,
				`Test Item`,
				`role="option"`,
			},
		},
		{
			name: "item with icon",
			item: CommandItem{
				Value: "user",
				Label: "User Profile",
				Icon:  icons.User(h.Class("h-4 w-4")),
			},
			cfg: &config{showShortcuts: true},
			contains: []string{
				`User Profile`,
				`<svg`,
			},
		},
		{
			name: "item with shortcut",
			item: CommandItem{
				Value:    "save",
				Label:    "Save",
				Shortcut: "⌘S",
			},
			cfg: &config{showShortcuts: true},
			contains: []string{
				`Save`,
				`⌘S`,
				`text-xs tracking-widest`,
			},
		},
		{
			name: "item with description",
			item: CommandItem{
				Value:       "settings",
				Label:       "Settings",
				Description: "Manage your preferences",
			},
			cfg: &config{showShortcuts: true},
			contains: []string{
				`Settings`,
				`Manage your preferences`,
				`text-xs text-muted-foreground`,
			},
		},
		{
			name: "disabled item",
			item: CommandItem{
				Value:    "delete",
				Label:    "Delete",
				Disabled: true,
			},
			cfg: &config{showShortcuts: true},
			contains: []string{
				`Delete`,
				`opacity-50`,
				`cursor-not-allowed`,
				`aria-disabled="true"`,
			},
		},
		{
			name: "item with category",
			item: CommandItem{
				Value:    "copy",
				Label:    "Copy",
				Category: "Edit",
			},
			cfg: &config{showShortcuts: true},
			contains: []string{
				`data-category="Edit"`,
			},
		},
		{
			name: "item without shortcuts shown",
			item: CommandItem{
				Value:    "test",
				Label:    "Test",
				Shortcut: "⌘T",
			},
			cfg: &config{showShortcuts: false},
			notContains: []string{
				`⌘T`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			node := CommandItemNode(tt.item, tt.cfg)
			err := node.Render(&buf)
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

func TestDialog(t *testing.T) {
	groups := []CommandGroup{
		{
			Label: "Test",
			Items: []CommandItem{
				{Value: "item1", Label: "Item 1"},
			},
		},
	}

	var buf bytes.Buffer
	dialog := Dialog("test-dialog", groups)
	err := dialog.Render(&buf)
	if err != nil {
		t.Fatalf("Render() error = %v", err)
	}

	result := buf.String()
	expected := []string{
		`id="test-dialog-dialog"`,
		`role="dialog"`,
		`aria-modal="true"`,
		`data-command-dialog="true"`,
		`fixed inset-0 z-50 hidden`,
		`backdrop-blur-sm`,
		`data-command-overlay="true"`,
	}

	for _, exp := range expected {
		if !strings.Contains(result, exp) {
			t.Errorf("Expected dialog to contain %q", exp)
		}
	}
}

func TestFilterGroups(t *testing.T) {
	groups := []CommandGroup{
		{
			Label: "Actions",
			Items: []CommandItem{
				{Value: "copy", Label: "Copy", Description: "Copy selection"},
				{Value: "paste", Label: "Paste", Description: "Paste from clipboard"},
				{Value: "cut", Label: "Cut", Description: "Cut selection"},
			},
		},
		{
			Label: "Navigation",
			Items: []CommandItem{
				{Value: "home", Label: "Go Home", Category: "nav"},
				{Value: "back", Label: "Go Back", Category: "nav"},
			},
		},
	}

	tests := []struct {
		name          string
		query         string
		expectedCount int
		expectedItems []string
	}{
		{
			name:          "empty query returns all",
			query:         "",
			expectedCount: 2,
			expectedItems: []string{"Copy", "Paste", "Cut", "Go Home", "Go Back"},
		},
		{
			name:          "search by label",
			query:         "copy",
			expectedCount: 1,
			expectedItems: []string{"Copy"},
		},
		{
			name:          "search by description",
			query:         "clipboard",
			expectedCount: 1,
			expectedItems: []string{"Paste"},
		},
		{
			name:          "search by category",
			query:         "nav",
			expectedCount: 1,
			expectedItems: []string{"Go Home", "Go Back"},
		},
		{
			name:          "search partial match",
			query:         "go",
			expectedCount: 1,
			expectedItems: []string{"Go Home", "Go Back"},
		},
		{
			name:          "case insensitive search",
			query:         "COPY",
			expectedCount: 1,
			expectedItems: []string{"Copy"},
		},
		{
			name:          "no matches",
			query:         "xyz",
			expectedCount: 0,
			expectedItems: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filtered := FilterGroups(groups, tt.query)
			
			if len(filtered) != tt.expectedCount {
				t.Errorf("Expected %d groups, got %d", tt.expectedCount, len(filtered))
			}

			// Check that expected items are present
			var allItems []string
			for _, group := range filtered {
				for _, item := range group.Items {
					allItems = append(allItems, item.Label)
				}
			}

			for _, expected := range tt.expectedItems {
				found := false
				for _, item := range allItems {
					if item == expected {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Expected to find item %q in filtered results", expected)
				}
			}
		})
	}
}

func TestHTMX(t *testing.T) {
	htmxConfig := HTMXConfig{
		SearchEndpoint: "/api/search",
		SelectEndpoint: "/api/select",
		DebounceMs:     500,
	}

	groups := []CommandGroup{
		{
			Label: "Test",
			Items: []CommandItem{
				{Value: "item1", Label: "Item 1"},
			},
		},
	}

	var buf bytes.Buffer
	menu := NewHTMX("htmx-command", groups, htmxConfig)
	err := menu.Render(&buf)
	if err != nil {
		t.Fatalf("Render() error = %v", err)
	}

	result := buf.String()
	expected := []string{
		`hx-get="/api/search"`,
		`hx-trigger="keyup changed delay:500ms"`,
		`hx-target="#htmx-command-list"`,
		`hx-swap="innerHTML"`,
		`htmx-indicator`,
	}

	for _, exp := range expected {
		if !strings.Contains(result, exp) {
			t.Errorf("Expected HTMX command to contain %q", exp)
		}
	}
}

func TestSeparator(t *testing.T) {
	var buf bytes.Buffer
	sep := Separator()
	err := sep.Render(&buf)
	if err != nil {
		t.Fatalf("Render() error = %v", err)
	}

	result := buf.String()
	expected := []string{
		`role="separator"`,
		`aria-orientation="horizontal"`,
		`data-command-separator="true"`,
		`h-px bg-border`,
	}

	for _, exp := range expected {
		if !strings.Contains(result, exp) {
			t.Errorf("Expected separator to contain %q", exp)
		}
	}
}