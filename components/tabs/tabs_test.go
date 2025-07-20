package tabs

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

func TestTabs(t *testing.T) {
	tests := []struct {
		name     string
		props    Props
		children []g.Node
		contains []string
	}{
		{
			name: "basic tabs",
			props: Props{
				DefaultValue: "tab1",
			},
			children: []g.Node{
				TabsList(ListProps{},
					Trigger(TriggerProps{Value: "tab1"}, g.Text("Tab 1")),
					Trigger(TriggerProps{Value: "tab2"}, g.Text("Tab 2")),
				),
				TabsContent(ContentProps{Value: "tab1"}, g.Text("Content 1")),
				TabsContent(ContentProps{Value: "tab2"}, g.Text("Content 2")),
			},
			contains: []string{
				`data-slot="tabs"`,
				`data-tabs-default="tab1"`,
				`flex flex-col gap-2`,
				"Tab 1",
				"Tab 2",
				"Content 1",
				"Content 2",
			},
		},
		{
			name: "tabs with custom class",
			props: Props{
				Class: "custom-tabs",
			},
			children: []g.Node{
				TabsList(ListProps{},
					Trigger(TriggerProps{Value: "tab1"}, g.Text("Tab 1")),
				),
				TabsContent(ContentProps{Value: "tab1"}, g.Text("Content 1")),
			},
			contains: []string{
				"custom-tabs",
			},
		},
		{
			name: "tabs with ID",
			props: Props{
				ID:           "my-tabs",
				DefaultValue: "settings",
			},
			children: []g.Node{
				TabsList(ListProps{},
					Trigger(TriggerProps{Value: "settings"}, g.Text("Settings")),
				),
				TabsContent(ContentProps{Value: "settings"}, g.Text("Settings content")),
			},
			contains: []string{
				`id="my-tabs"`,
				`data-tabs-default="settings"`,
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

func TestTabsTabsList(t *testing.T) {
	tests := []struct {
		name     string
		props    ListProps
		children []g.Node
		contains []string
	}{
		{
			name: "basic list",
			props: ListProps{},
			children: []g.Node{
				Trigger(TriggerProps{Value: "tab1"}, g.Text("Tab 1")),
				Trigger(TriggerProps{Value: "tab2"}, g.Text("Tab 2")),
			},
			contains: []string{
				`data-slot="tabs-list"`,
				`role="tablist"`,
				"bg-muted text-muted-foreground",
				"rounded-lg p-[3px]",
			},
		},
		{
			name: "list with custom class",
			props: ListProps{
				Class: "custom-list",
			},
			children: []g.Node{
				Trigger(TriggerProps{Value: "tab1"}, g.Text("Tab 1")),
			},
			contains: []string{
				"custom-list",
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := TabsList(tt.props, tt.children...)
			html := renderToString(t, component)
			
			for _, substring := range tt.contains {
				if !strings.Contains(html, substring) {
					t.Errorf("expected HTML to contain %q, but it didn't.\nHTML: %s", substring, html)
				}
			}
		})
	}
}

func TestTabsTrigger(t *testing.T) {
	tests := []struct {
		name     string
		props    TriggerProps
		children []g.Node
		contains []string
	}{
		{
			name: "basic trigger",
			props: TriggerProps{
				Value: "tab1",
			},
			children: []g.Node{g.Text("Tab 1")},
			contains: []string{
				`data-slot="tabs-trigger"`,
				`data-tabs-value="tab1"`,
				`data-state="inactive"`,
				`role="tab"`,
				`aria-selected="false"`,
				`aria-controls="content-tab1"`,
				`tabindex="-1"`,
				`type="button"`,
				"Tab 1",
			},
		},
		{
			name: "disabled trigger",
			props: TriggerProps{
				Value:    "tab2",
				Disabled: true,
			},
			children: []g.Node{g.Text("Disabled Tab")},
			contains: []string{
				`disabled`,
				"Disabled Tab",
			},
		},
		{
			name: "trigger with custom class",
			props: TriggerProps{
				Value: "tab3",
				Class: "custom-trigger",
			},
			children: []g.Node{g.Text("Custom Tab")},
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

func TestTabsTabsContent(t *testing.T) {
	tests := []struct {
		name     string
		props    ContentProps
		children []g.Node
		contains []string
	}{
		{
			name: "basic content",
			props: ContentProps{
				Value: "tab1",
			},
			children: []g.Node{g.Text("Tab content")},
			contains: []string{
				`data-slot="tabs-content"`,
				`data-tabs-value="tab1"`,
				`data-state="inactive"`,
				`role="tabpanel"`,
				`id="content-tab1"`,
				`aria-labelledby="trigger-tab1"`,
				`tabindex="0"`,
				`style="display: none;"`,
				"Tab content",
			},
		},
		{
			name: "content with custom class",
			props: ContentProps{
				Value: "tab2",
				Class: "custom-content",
			},
			children: []g.Node{g.Text("Custom content")},
			contains: []string{
				"custom-content",
				"Custom content",
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := TabsContent(tt.props, tt.children...)
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
	t.Run("WithDefault", func(t *testing.T) {
		tabs := WithDefault("profile",
			TabsList(ListProps{},
				Trigger(TriggerProps{Value: "account"}, g.Text("Account")),
				Trigger(TriggerProps{Value: "profile"}, g.Text("Profile")),
			),
			TabsContent(ContentProps{Value: "account"}, g.Text("Account content")),
			TabsContent(ContentProps{Value: "profile"}, g.Text("Profile content")),
		)
		
		html := renderToString(t, tabs)
		expected := []string{
			`data-tabs-default="profile"`,
			"Account",
			"Profile",
			"Account content",
			"Profile content",
		}
		
		for _, substring := range expected {
			if !strings.Contains(html, substring) {
				t.Errorf("expected HTML to contain %q, but it didn't.\nHTML: %s", substring, html)
			}
		}
	})
}

func TestDataAndAriaAttributes(t *testing.T) {
	t.Run("dataAttr", func(t *testing.T) {
		attr := dataAttr("test", "value")
		html := renderToString(t, Div(attr))
		
		if !strings.Contains(html, `data-test="value"`) {
			t.Errorf("expected data attribute to be rendered correctly, got: %s", html)
		}
	})
	
	t.Run("ariaAttr", func(t *testing.T) {
		attr := ariaAttr("label", "Test label")
		html := renderToString(t, Div(attr))
		
		if !strings.Contains(html, `aria-label="Test label"`) {
			t.Errorf("expected aria attribute to be rendered correctly, got: %s", html)
		}
	})
}