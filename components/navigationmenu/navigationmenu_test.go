package navigationmenu

import (
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
	}{
		{
			name:  "renders basic navigation menu",
			props: Props{},
			want: []string{
				`role="navigation"`,
				`data-orientation="horizontal"`,
				`relative z-10 flex max-w-max flex-1 items-center justify-center`,
			},
		},
		{
			name: "renders with vertical orientation",
			props: Props{
				Orientation: "vertical",
			},
			want: []string{
				`data-orientation="vertical"`,
			},
		},
		{
			name: "renders with custom class",
			props: Props{
				Class: "custom-nav",
			},
			want: []string{
				`custom-nav`,
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
		})
	}
}

func TestList(t *testing.T) {
	tests := []struct {
		name     string
		props    ListProps
		children []g.Node
		want     []string
	}{
		{
			name:  "renders basic list",
			props: ListProps{},
			want: []string{
				`role="none"`,
				`data-orientation="horizontal"`,
				`group flex flex-1 list-none items-center justify-center space-x-1`,
			},
		},
		{
			name: "renders with custom class",
			props: ListProps{
				Class: "custom-list",
			},
			want: []string{
				`custom-list`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := renderToString(List(tt.props, tt.children...))

			for _, want := range tt.want {
				if !strings.Contains(result, want) {
					t.Errorf("List() = %v, want to contain %v", result, want)
				}
			}
		})
	}
}

func TestItem(t *testing.T) {
	tests := []struct {
		name     string
		props    ItemProps
		children []g.Node
		want     []string
	}{
		{
			name:  "renders basic item",
			props: ItemProps{},
			want: []string{
				`role="none"`,
				`class="relative"`,
			},
		},
		{
			name: "renders with value",
			props: ItemProps{
				Value: "item-1",
			},
			want: []string{
				`data-value="item-1"`,
			},
		},
		{
			name: "renders with custom class",
			props: ItemProps{
				Class: "custom-item",
			},
			want: []string{
				`custom-item`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := renderToString(Item(tt.props, tt.children...))

			for _, want := range tt.want {
				if !strings.Contains(result, want) {
					t.Errorf("Item() = %v, want to contain %v", result, want)
				}
			}
		})
	}
}

func TestTrigger(t *testing.T) {
	tests := []struct {
		name     string
		props    TriggerProps
		children []g.Node
		want     []string
	}{
		{
			name:  "renders basic trigger",
			props: TriggerProps{},
			want: []string{
				`type="button"`,
				`role="menuitem"`,
				`aria-haspopup="menu"`,
				`aria-expanded="false"`,
				`data-state="closed"`,
				`<svg`, // chevron icon
			},
		},
		{
			name: "renders disabled trigger",
			props: TriggerProps{
				Disabled: true,
			},
			want: []string{
				`disabled`,
				`disabled:pointer-events-none disabled:opacity-50`,
			},
		},
		{
			name: "renders with custom class",
			props: TriggerProps{
				Class: "custom-trigger",
			},
			want: []string{
				`custom-trigger`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := renderToString(Trigger(tt.props, tt.children...))

			for _, want := range tt.want {
				if !strings.Contains(result, want) {
					t.Errorf("Trigger() = %v, want to contain %v", result, want)
				}
			}
		})
	}
}

func TestContent(t *testing.T) {
	tests := []struct {
		name     string
		props    ContentProps
		children []g.Node
		want     []string
	}{
		{
			name:  "renders basic content",
			props: ContentProps{},
			want: []string{
				`data-state="closed"`,
				`style="display: none;"`,
				`md:absolute md:w-auto`,
			},
		},
		{
			name: "renders with custom class",
			props: ContentProps{
				Class: "custom-content",
			},
			want: []string{
				`custom-content`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := renderToString(Content(tt.props, tt.children...))

			for _, want := range tt.want {
				if !strings.Contains(result, want) {
					t.Errorf("Content() = %v, want to contain %v", result, want)
				}
			}
		})
	}
}

func TestLink(t *testing.T) {
	tests := []struct {
		name     string
		props    LinkProps
		children []g.Node
		want     []string
		notWant  []string
	}{
		{
			name: "renders basic link",
			props: LinkProps{
				Href: "/home",
			},
			want: []string{
				`href="/home"`,
				`role="menuitem"`,
			},
			notWant: []string{
				`aria-current`,
				`aria-disabled`,
			},
		},
		{
			name: "renders active link",
			props: LinkProps{
				Href:   "/home",
				Active: true,
			},
			want: []string{
				`aria-current="page"`,
				`bg-accent text-accent-foreground`,
			},
		},
		{
			name: "renders disabled link",
			props: LinkProps{
				Href:     "/home",
				Disabled: true,
			},
			want: []string{
				`aria-disabled="true"`,
				`pointer-events-none opacity-50`,
			},
		},
		{
			name: "renders with custom class",
			props: LinkProps{
				Href:  "/home",
				Class: "custom-link",
			},
			want: []string{
				`custom-link`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := renderToString(Link(tt.props, tt.children...))

			for _, want := range tt.want {
				if !strings.Contains(result, want) {
					t.Errorf("Link() = %v, want to contain %v", result, want)
				}
			}

			for _, notWant := range tt.notWant {
				if strings.Contains(result, notWant) {
					t.Errorf("Link() = %v, don't want to contain %v", result, notWant)
				}
			}
		})
	}
}

func TestViewport(t *testing.T) {
	tests := []struct {
		name     string
		props    ViewportProps
		children []g.Node
		want     []string
	}{
		{
			name:  "renders basic viewport",
			props: ViewportProps{},
			want: []string{
				`data-state="closed"`,
				`style="display: none;"`,
				`origin-top-center relative mt-1.5`,
				`overflow-hidden rounded-md border bg-popover`,
			},
		},
		{
			name: "renders with custom class",
			props: ViewportProps{
				Class: "custom-viewport",
			},
			want: []string{
				`custom-viewport`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := renderToString(Viewport(tt.props, tt.children...))

			for _, want := range tt.want {
				if !strings.Contains(result, want) {
					t.Errorf("Viewport() = %v, want to contain %v", result, want)
				}
			}
		})
	}
}

func TestIndicator(t *testing.T) {
	tests := []struct {
		name  string
		props IndicatorProps
		want  []string
	}{
		{
			name:  "renders basic indicator",
			props: IndicatorProps{},
			want: []string{
				`data-state="hidden"`,
				`top-full z-[1] flex h-1.5 items-end justify-center overflow-hidden`,
				`rotate-45 rounded-tl-sm bg-border shadow-md`,
			},
		},
		{
			name: "renders with custom class",
			props: IndicatorProps{
				Class: "custom-indicator",
			},
			want: []string{
				`custom-indicator`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := renderToString(Indicator(tt.props))

			for _, want := range tt.want {
				if !strings.Contains(result, want) {
					t.Errorf("Indicator() = %v, want to contain %v", result, want)
				}
			}
		})
	}
}

func TestListItem(t *testing.T) {
	result := renderToString(ListItem("Test Title", "/test", "Test description"))

	want := []string{
		`href="/test"`,
		`Test Title`,
		`Test description`,
		`text-sm font-medium leading-none`,
		`line-clamp-2 text-sm leading-snug text-muted-foreground`,
	}

	for _, w := range want {
		if !strings.Contains(result, w) {
			t.Errorf("ListItem() = %v, want to contain %v", result, w)
		}
	}
}

func TestDefault(t *testing.T) {
	result := renderToString(Default())

	if !strings.Contains(result, `role="navigation"`) {
		t.Errorf("Default() = %v, want to contain navigation role", result)
	}
}

func TestSimpleMenu(t *testing.T) {
	result := renderToString(SimpleMenu())

	want := []string{
		`href="/"`,
		`Home`,
		`href="/about"`,
		`About`,
		`href="/services"`,
		`Services`,
		`href="/contact"`,
		`Contact`,
	}

	for _, w := range want {
		if !strings.Contains(result, w) {
			t.Errorf("SimpleMenu() = %v, want to contain %v", result, w)
		}
	}
}

func TestWithViewport(t *testing.T) {
	result := renderToString(WithViewport(
		List(
			ListProps{},
			Item(ItemProps{}, Link(LinkProps{Href: "#"}, g.Text("Test"))),
		),
	))

	want := []string{
		`role="navigation"`,
		`data-state="closed"`,
		`Test`,
	}

	for _, w := range want {
		if !strings.Contains(result, w) {
			t.Errorf("WithViewport() = %v, want to contain %v", result, w)
		}
	}
}

func TestWithDropdowns(t *testing.T) {
	result := renderToString(WithDropdowns())

	want := []string{
		`Getting started`,
		`Components`,
		`Documentation`,
		`data-value="getting-started"`,
		`data-value="components"`,
		`Introduction`,
		`Installation`,
		`Alert Dialog`,
		`Hover Card`,
	}

	for _, w := range want {
		if !strings.Contains(result, w) {
			t.Errorf("WithDropdowns() = %v, want to contain %v", result, w)
		}
	}
}

func TestCompleteNavigation(t *testing.T) {
	nav := New(
		Props{},
		List(
			ListProps{},
			Item(
				ItemProps{Value: "products"},
				Trigger(TriggerProps{}, g.Text("Products")),
				Content(
					ContentProps{},
					Ul(Class("p-4"),
						ListItem("Product 1", "/product1", "Description 1"),
						ListItem("Product 2", "/product2", "Description 2"),
					),
				),
			),
			Item(
				ItemProps{},
				Link(LinkProps{Href: "/pricing", Active: true}, g.Text("Pricing")),
			),
			Item(
				ItemProps{},
				Link(LinkProps{Href: "/contact", Disabled: true}, g.Text("Contact")),
			),
		),
	)

	result := renderToString(nav)

	// Check structure
	expectedElements := []string{
		`role="navigation"`,
		`Products`,
		`Product 1`,
		`Product 2`,
		`Pricing`,
		`Contact`,
		`aria-current="page"`,
		`aria-disabled="true"`,
		`data-value="products"`,
	}

	for _, elem := range expectedElements {
		if !strings.Contains(result, elem) {
			t.Errorf("Complete navigation = %v, want to contain %v", result, elem)
		}
	}
}