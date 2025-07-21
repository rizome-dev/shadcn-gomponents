package sidebar

import (
	"bytes"
	"strings"
	"testing"

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
			name:  "default sidebar",
			props: Props{},
			children: []g.Node{
				Header(Props{}, g.Text("Header")),
				Content(Props{}, g.Text("Content")),
				Footer(Props{}, g.Text("Footer")),
			},
			contains: []string{
				`data-sidebar="true"`,
				`data-variant="sidebar"`,
				`data-side="left"`,
				`data-collapsible="offcanvas"`,
				"Header",
				"Content",
				"Footer",
			},
		},
		{
			name: "right side floating sidebar",
			props: Props{
				Side:    "right",
				Variant: "floating",
			},
			children: []g.Node{
				Content(Props{}, g.Text("Floating Content")),
			},
			contains: []string{
				`data-side="right"`,
				`data-variant="floating"`,
				"Floating Content",
			},
		},
		{
			name: "non-collapsible sidebar",
			props: Props{
				Collapsible: "none",
			},
			children: []g.Node{
				Content(Props{}, g.Text("Static Content")),
			},
			contains: []string{
				`data-sidebar="true"`,
				"bg-sidebar",
				"Static Content",
			},
		},
		{
			name: "icon collapsible sidebar",
			props: Props{
				Collapsible: "icon",
			},
			children: []g.Node{
				Content(Props{}, g.Text("Icon Collapsible")),
			},
			contains: []string{
				`data-collapsible="icon"`,
				"Icon Collapsible",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			component := New(tt.props, tt.children...)
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

func TestProvider(t *testing.T) {
	tests := []struct {
		name     string
		props    ProviderProps
		children []g.Node
		contains []string
	}{
		{
			name: "default provider",
			props: ProviderProps{
				DefaultOpen: true,
			},
			children: []g.Node{
				g.Text("Provider Content"),
			},
			contains: []string{
				`data-sidebar-wrapper="true"`,
				`--sidebar-width:`,
				`--sidebar-width-icon:`,
				"Provider Content",
			},
		},
		{
			name: "provider with custom style",
			props: ProviderProps{
				Style: "custom-style: value",
			},
			children: []g.Node{
				g.Text("Styled Provider"),
			},
			contains: []string{
				"custom-style: value",
				"Styled Provider",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			component := Provider(tt.props, tt.children...)
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

func TestTrigger(t *testing.T) {
	tests := []struct {
		name     string
		props    Props
		children []g.Node
		contains []string
	}{
		{
			name:     "default trigger",
			props:    Props{},
			children: []g.Node{},
			contains: []string{
				`data-sidebar-trigger="true"`,
				`type="button"`,
				"Toggle Sidebar",
				"size-7",
			},
		},
		{
			name:  "trigger with custom content",
			props: Props{},
			children: []g.Node{
				g.Text("Custom Trigger"),
			},
			contains: []string{
				"Custom Trigger",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			component := Trigger(tt.props, tt.children...)
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

func TestMenuButton(t *testing.T) {
	tests := []struct {
		name     string
		props    MenuButtonProps
		children []g.Node
		contains []string
	}{
		{
			name:  "default menu button",
			props: MenuButtonProps{},
			children: []g.Node{
				g.Text("Menu Item"),
			},
			contains: []string{
				`data-sidebar-menu-button="true"`,
				`type="button"`,
				"Menu Item",
				"hover:bg-sidebar-accent",
			},
		},
		{
			name: "active menu button with href",
			props: MenuButtonProps{
				IsActive: true,
				Href:     "/dashboard",
			},
			children: []g.Node{
				g.Text("Dashboard"),
			},
			contains: []string{
				`data-active="true"`,
				`href="/dashboard"`,
				"Dashboard",
			},
		},
		{
			name: "outline variant large size",
			props: MenuButtonProps{
				Variant: "outline",
				Size:    "lg",
			},
			children: []g.Node{
				g.Text("Large Outline"),
			},
			contains: []string{
				`data-size="lg"`,
				"shadow-[0_0_0_1px",
				"Large Outline",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			component := MenuButton(tt.props, tt.children...)
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

func TestMenuComponents(t *testing.T) {
	// Test menu structure
	menu := Menu(Props{},
		MenuItem(Props{},
			MenuButton(MenuButtonProps{}, g.Text("Item 1")),
			MenuAction(Props{}, g.Text("Action")),
			MenuBadge(Props{}, g.Text("3")),
		),
		MenuItem(Props{},
			MenuButton(MenuButtonProps{IsActive: true}, g.Text("Item 2")),
		),
	)

	var buf bytes.Buffer
	err := menu.Render(&buf)
	if err != nil {
		t.Fatalf("Render() error = %v", err)
	}

	result := buf.String()
	expected := []string{
		`data-sidebar-menu="true"`,
		`data-sidebar-menu-item="true"`,
		`data-sidebar-menu-button="true"`,
		`data-sidebar-menu-action="true"`,
		`data-sidebar-menu-badge="true"`,
		"Item 1",
		"Item 2",
		"Action",
		"3",
	}

	for _, exp := range expected {
		if !strings.Contains(result, exp) {
			t.Errorf("Expected output to contain %q, but it didn't. Got:\n%s", exp, result)
		}
	}
}

func TestSubMenu(t *testing.T) {
	submenu := MenuSub(Props{},
		MenuSubItem(Props{},
			MenuSubButton(MenuSubButtonProps{
				Href: "/sub1",
			}, g.Text("Sub Item 1")),
		),
		MenuSubItem(Props{},
			MenuSubButton(MenuSubButtonProps{
				IsActive: true,
				Size:     "sm",
			}, g.Text("Sub Item 2")),
		),
	)

	var buf bytes.Buffer
	err := submenu.Render(&buf)
	if err != nil {
		t.Fatalf("Render() error = %v", err)
	}

	result := buf.String()
	expected := []string{
		`data-sidebar-menu-sub="true"`,
		`data-sidebar-menu-sub-item="true"`,
		`data-sidebar-menu-sub-button="true"`,
		`href="/sub1"`,
		`data-active="true"`,
		`data-size="sm"`,
		"Sub Item 1",
		"Sub Item 2",
	}

	for _, exp := range expected {
		if !strings.Contains(result, exp) {
			t.Errorf("Expected output to contain %q, but it didn't. Got:\n%s", exp, result)
		}
	}
}