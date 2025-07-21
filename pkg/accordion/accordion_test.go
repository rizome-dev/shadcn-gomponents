package accordion

import (
	"bytes"
	"strings"
	"testing"
	
	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// renderToString renders a node to a string for testing
func renderToString(t *testing.T, node g.Node) string {
	var buf bytes.Buffer
	if err := node.Render(&buf); err != nil {
		t.Fatal(err)
	}
	return buf.String()
}

func TestAccordion(t *testing.T) {
	tests := []struct {
		name     string
		props    Props
		children []g.Node
		contains []string
	}{
		{
			name: "single accordion",
			props: Props{
				Type:        "single",
				Collapsible: true,
			},
			children: []g.Node{
				Item(ItemProps{Value: "item-1"},
					Trigger(TriggerProps{}, g.Text("Item 1")),
					ItemContent(ContentProps{}, g.Text("Content 1")),
				),
			},
			contains: []string{
				`data-accordion-type="single"`,
				`data-accordion-collapsible="true"`,
				`data-accordion-value="item-1"`,
				"Item 1",
				"Content 1",
			},
		},
		{
			name: "accordion with default value",
			props: Props{
				Type:         "single",
				DefaultValue: "item-2",
			},
			children: []g.Node{
				Item(ItemProps{Value: "item-1"},
					Trigger(TriggerProps{}, g.Text("Item 1")),
					ItemContent(ContentProps{}, g.Text("Content 1")),
				),
				Item(ItemProps{Value: "item-2"},
					Trigger(TriggerProps{}, g.Text("Item 2")),
					ItemContent(ContentProps{}, g.Text("Content 2")),
				),
			},
			contains: []string{
				`data-accordion-default="item-2"`,
				`data-accordion-value="item-1"`,
				`data-accordion-value="item-2"`,
			},
		},
		{
			name: "multiple accordion",
			props: Props{
				Type:        "multiple",
				Collapsible: true,
			},
			children: []g.Node{
				Item(ItemProps{Value: "item-1"},
					Trigger(TriggerProps{}, g.Text("Item 1")),
					ItemContent(ContentProps{}, g.Text("Content 1")),
				),
				Item(ItemProps{Value: "item-2"},
					Trigger(TriggerProps{}, g.Text("Item 2")),
					ItemContent(ContentProps{}, g.Text("Content 2")),
				),
			},
			contains: []string{
				`data-accordion-type="multiple"`,
				`data-accordion-collapsible="true"`,
			},
		},
		{
			name: "accordion with custom classes",
			props: Props{
				Type:  "single",
				Class: "custom-accordion",
			},
			children: []g.Node{
				Item(ItemProps{Value: "item-1", Class: "custom-item"},
					Trigger(TriggerProps{Class: "custom-trigger"}, g.Text("Item 1")),
					ItemContent(ContentProps{Class: "custom-content"}, g.Text("Content 1")),
				),
			},
			contains: []string{
				"custom-accordion",
				"custom-item",
				"custom-trigger",
				"custom-content",
			},
		},
		{
			name: "accordion with ID",
			props: Props{
				Type: "single",
				ID:   "my-accordion",
			},
			children: []g.Node{
				Item(ItemProps{Value: "item-1"},
					Trigger(TriggerProps{}, g.Text("Item 1")),
					ItemContent(ContentProps{}, g.Text("Content 1")),
				),
			},
			contains: []string{
				`id="my-accordion"`,
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := New(tt.props, tt.children...)
			html := renderToString(t, component)
			
			for _, substring := range tt.contains {
				if !strings.Contains(html, substring) {
					t.Errorf("expected HTML to contain %q, but it didn't.\nHTML: %s", substring, html)
				}
			}
		})
	}
}

func TestAccordionItem(t *testing.T) {
	tests := []struct {
		name     string
		props    ItemProps
		children []g.Node
		contains []string
	}{
		{
			name:  "basic item",
			props: ItemProps{Value: "test-item"},
			children: []g.Node{
				Trigger(TriggerProps{}, g.Text("Trigger")),
				ItemContent(ContentProps{}, g.Text("Content")),
			},
			contains: []string{
				`data-slot="accordion-item"`,
				`data-accordion-value="test-item"`,
				"border-b",
			},
		},
		{
			name:  "item with custom class",
			props: ItemProps{Value: "test-item", Class: "custom-class"},
			children: []g.Node{
				Trigger(TriggerProps{}, g.Text("Trigger")),
			},
			contains: []string{
				"custom-class",
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := Item(tt.props, tt.children...)
			html := renderToString(t, component)
			
			for _, substring := range tt.contains {
				if !strings.Contains(html, substring) {
					t.Errorf("expected HTML to contain %q, but it didn't.\nHTML: %s", substring, html)
				}
			}
		})
	}
}

func TestAccordionTrigger(t *testing.T) {
	tests := []struct {
		name     string
		props    TriggerProps
		children []g.Node
		contains []string
	}{
		{
			name:     "basic trigger",
			props:    TriggerProps{},
			children: []g.Node{g.Text("Click me")},
			contains: []string{
				`data-slot="accordion-trigger"`,
				`aria-expanded="false"`,
				`type="button"`,
				"Click me",
				"<svg", // Default chevron icon
			},
		},
		{
			name:  "trigger with custom icon",
			props: TriggerProps{Icon: Span(g.Text("→"))},
			children: []g.Node{g.Text("Custom icon")},
			contains: []string{
				"Custom icon",
				"→",
			},
		},
		{
			name:     "trigger with custom class",
			props:    TriggerProps{Class: "custom-trigger"},
			children: []g.Node{g.Text("Styled trigger")},
			contains: []string{
				"custom-trigger",
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := Trigger(tt.props, tt.children...)
			html := renderToString(t, component)
			
			for _, substring := range tt.contains {
				if !strings.Contains(html, substring) {
					t.Errorf("expected HTML to contain %q, but it didn't.\nHTML: %s", substring, html)
				}
			}
		})
	}
}

func TestAccordionItemContent(t *testing.T) {
	tests := []struct {
		name     string
		props    ContentProps
		children []g.Node
		contains []string
	}{
		{
			name:     "basic content",
			props:    ContentProps{},
			children: []g.Node{g.Text("Content text")},
			contains: []string{
				`data-slot="accordion-content"`,
				`data-state="closed"`,
				`max-height: 0;`,
				"Content text",
				"overflow-hidden",
			},
		},
		{
			name:     "content with custom class",
			props:    ContentProps{Class: "custom-content"},
			children: []g.Node{g.Text("Styled content")},
			contains: []string{
				"custom-content",
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := ItemContent(tt.props, tt.children...)
			html := renderToString(t, component)
			
			for _, substring := range tt.contains {
				if !strings.Contains(html, substring) {
					t.Errorf("expected HTML to contain %q, but it didn't.\nHTML: %s", substring, html)
				}
			}
		})
	}
}

func TestHelperFunctions(t *testing.T) {
	t.Run("Single accordion", func(t *testing.T) {
		accordion := Single(true, "item-1",
			Item(ItemProps{Value: "item-1"},
				Trigger(TriggerProps{}, g.Text("Item 1")),
				ItemContent(ContentProps{}, g.Text("Content 1")),
			),
		)
		html := renderToString(t, accordion)
		expected := []string{
			`data-accordion-type="single"`,
			`data-accordion-collapsible="true"`,
			`data-accordion-default="item-1"`,
		}
		
		for _, substring := range expected {
			if !strings.Contains(html, substring) {
				t.Errorf("expected HTML to contain %q, but it didn't.\nHTML: %s", substring, html)
			}
		}
	})
	
	t.Run("Multiple accordion", func(t *testing.T) {
		accordion := MultipleAccordion([]string{"item-1", "item-2"},
			Item(ItemProps{Value: "item-1"},
				Trigger(TriggerProps{}, g.Text("Item 1")),
				ItemContent(ContentProps{}, g.Text("Content 1")),
			),
			Item(ItemProps{Value: "item-2"},
				Trigger(TriggerProps{}, g.Text("Item 2")),
				ItemContent(ContentProps{}, g.Text("Content 2")),
			),
		)
		html := renderToString(t, accordion)
		expected := []string{
			`data-accordion-type="multiple"`,
			`data-accordion-collapsible="true"`,
		}
		
		for _, substring := range expected {
			if !strings.Contains(html, substring) {
				t.Errorf("expected HTML to contain %q, but it didn't.\nHTML: %s", substring, html)
			}
		}
	})
}

// TestDataAndAriaAttributes is removed since dataAttr is now unexported
// The functionality is tested through the component tests above