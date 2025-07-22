package pagination

import (
	"fmt"
	"strings"
	"testing"

	g "maragu.dev/gomponents"
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
			name:  "renders basic pagination",
			props: Props{},
			want: []string{
				`role="navigation"`,
				`aria-label="pagination"`,
				`mx-auto flex w-full justify-center`,
			},
		},
		{
			name: "renders with custom class",
			props: Props{
				Class: "custom-pagination",
			},
			want: []string{
				`custom-pagination`,
			},
		},
		{
			name: "renders with properties",
			props: Props{
				CurrentPage: 5,
				TotalPages:  10,
			},
			want: []string{
				`role="navigation"`,
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

func TestContent(t *testing.T) {
	result := renderToString(ContentComponent(ContentProps{}, g.Text("content")))

	want := []string{
		`flex flex-row items-center gap-1`,
		"content",
	}

	for _, w := range want {
		if !strings.Contains(result, w) {
			t.Errorf("Content() = %v, want to contain %v", result, w)
		}
	}
}

func TestItem(t *testing.T) {
	tests := []struct {
		name  string
		props ItemProps
		want  []string
	}{
		{
			name:  "renders basic item",
			props: ItemProps{},
			want:  []string{`<li`},
		},
		{
			name: "renders active item",
			props: ItemProps{
				Active: true,
			},
			want: []string{
				`aria-current="page"`,
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
			result := renderToString(Item(tt.props, g.Text("test")))

			for _, want := range tt.want {
				if !strings.Contains(result, want) {
					t.Errorf("Item() = %v, want to contain %v", result, want)
				}
			}
		})
	}
}

func TestLink(t *testing.T) {
	tests := []struct {
		name  string
		props LinkProps
		want  []string
	}{
		{
			name: "renders basic link",
			props: LinkProps{
				Href: "/page/2",
				Page: 2,
			},
			want: []string{
				`href="/page/2"`,
				`h-10 w-10`,
			},
		},
		{
			name: "renders active link",
			props: LinkProps{
				Href:   "/page/5",
				Page:   5,
				Active: true,
			},
			want: []string{
				`aria-current="page"`,
				`tabindex="-1"`,
				`border border-input bg-background`,
			},
		},
		{
			name: "renders disabled link",
			props: LinkProps{
				Href:     "/page/1",
				Disabled: true,
			},
			want: []string{
				`aria-disabled="true"`,
				`href="#"`,
				`tabindex="-1"`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := renderToString(LinkComponent(tt.props, g.Text("test")))

			for _, want := range tt.want {
				if !strings.Contains(result, want) {
					t.Errorf("Link() = %v, want to contain %v", result, want)
				}
			}
		})
	}
}

func TestPreviousButton(t *testing.T) {
	tests := []struct {
		name     string
		href     string
		disabled bool
		want     []string
	}{
		{
			name: "renders enabled previous button",
			href: "/page/4",
			want: []string{
				`href="/page/4"`,
				`Previous`,
				`<svg`,
				`Go to previous page`,
			},
		},
		{
			name:     "renders disabled previous button",
			href:     "/page/0",
			disabled: true,
			want: []string{
				`aria-disabled="true"`,
				`href="#"`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := renderToString(PreviousButton(tt.href, tt.disabled))

			for _, want := range tt.want {
				if !strings.Contains(result, want) {
					t.Errorf("PreviousButton() = %v, want to contain %v", result, want)
				}
			}
		})
	}
}

func TestNextButton(t *testing.T) {
	tests := []struct {
		name     string
		href     string
		disabled bool
		want     []string
	}{
		{
			name: "renders enabled next button",
			href: "/page/6",
			want: []string{
				`href="/page/6"`,
				`Next`,
				`<svg`,
				`Go to next page`,
			},
		},
		{
			name:     "renders disabled next button",
			href:     "/page/11",
			disabled: true,
			want: []string{
				`aria-disabled="true"`,
				`href="#"`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := renderToString(NextButton(tt.href, tt.disabled))

			for _, want := range tt.want {
				if !strings.Contains(result, want) {
					t.Errorf("NextButton() = %v, want to contain %v", result, want)
				}
			}
		})
	}
}

func TestFirstButton(t *testing.T) {
	result := renderToString(FirstButton("/page/1", false))

	want := []string{
		`href="/page/1"`,
		`Go to first page`,
		`<svg`,
	}

	for _, w := range want {
		if !strings.Contains(result, w) {
			t.Errorf("FirstButton() = %v, want to contain %v", result, w)
		}
	}
}

func TestLastButton(t *testing.T) {
	result := renderToString(LastButton("/page/10", false))

	want := []string{
		`href="/page/10"`,
		`Go to last page`,
		`<svg`,
	}

	for _, w := range want {
		if !strings.Contains(result, w) {
			t.Errorf("LastButton() = %v, want to contain %v", result, w)
		}
	}
}

func TestEllipsis(t *testing.T) {
	result := renderToString(Ellipsis())

	want := []string{
		`aria-hidden="true"`,
		`More pages`,
		`<svg`,
	}

	for _, w := range want {
		if !strings.Contains(result, w) {
			t.Errorf("Ellipsis() = %v, want to contain %v", result, w)
		}
	}
}

func TestPageButton(t *testing.T) {
	tests := []struct {
		name   string
		page   int
		href   string
		active bool
		want   []string
	}{
		{
			name: "renders regular page button",
			page: 3,
			href: "/page/3",
			want: []string{
				`href="/page/3"`,
				`>3<`,
			},
		},
		{
			name:   "renders active page button",
			page:   5,
			href:   "/page/5",
			active: true,
			want: []string{
				`aria-current="page"`,
				`>5<`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := renderToString(PageButton(tt.page, tt.href, tt.active))

			for _, want := range tt.want {
				if !strings.Contains(result, want) {
					t.Errorf("PageButton() = %v, want to contain %v", result, want)
				}
			}
		})
	}
}

func TestSimple(t *testing.T) {
	getPageURL := func(page int) string {
		return fmt.Sprintf("/page/%d", page)
	}

	result := renderToString(Simple(5, 10, getPageURL))

	want := []string{
		`Previous`,
		`Next`,
		`href="/page/4"`,
		`href="/page/6"`,
	}

	for _, w := range want {
		if !strings.Contains(result, w) {
			t.Errorf("Simple() = %v, want to contain %v", result, w)
		}
	}
}

func TestDefault(t *testing.T) {
	getPageURL := func(page int) string {
		return fmt.Sprintf("/page/%d", page)
	}

	tests := []struct {
		name        string
		currentPage int
		totalPages  int
		want        []string
	}{
		{
			name:        "renders pagination for first page",
			currentPage: 1,
			totalPages:  10,
			want: []string{
				`>1<`,
				`>2<`,
				`>3<`,
				`More pages`,
				`>10<`,
			},
		},
		{
			name:        "renders pagination for middle page",
			currentPage: 5,
			totalPages:  10,
			want: []string{
				`>1<`,
				`>4<`,
				`>5<`,
				`>6<`,
				`>10<`,
			},
		},
		{
			name:        "renders pagination for last page",
			currentPage: 10,
			totalPages:  10,
			want: []string{
				`>1<`,
				`More pages`,
				`>8<`,
				`>9<`,
				`>10<`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := renderToString(Default(tt.currentPage, tt.totalPages, getPageURL))

			for _, want := range tt.want {
				if !strings.Contains(result, want) {
					t.Errorf("Default() = %v, want to contain %v", result, want)
				}
			}
		})
	}
}

func TestWithFirstLast(t *testing.T) {
	getPageURL := func(page int) string {
		return fmt.Sprintf("/page/%d", page)
	}

	result := renderToString(WithFirstLast(5, 20, getPageURL))

	want := []string{
		`Go to first page`,
		`Previous`,
		`Next`,
		`Go to last page`,
		`>5<`,
	}

	for _, w := range want {
		if !strings.Contains(result, w) {
			t.Errorf("WithFirstLast() = %v, want to contain %v", result, w)
		}
	}
}

func TestGeneratePageNumbers(t *testing.T) {
	tests := []struct {
		name       string
		current    int
		total      int
		maxVisible int
		want       []int
	}{
		{
			name:       "all pages when total <= maxVisible",
			current:    3,
			total:      5,
			maxVisible: 7,
			want:       []int{1, 2, 3, 4, 5},
		},
		{
			name:       "with ellipsis at end",
			current:    2,
			total:      10,
			maxVisible: 7,
			want:       []int{1, 2, 3, 4, 5, -1, 10},
		},
		{
			name:       "with ellipsis at start",
			current:    9,
			total:      10,
			maxVisible: 7,
			want:       []int{1, -1, 6, 7, 8, 9, 10},
		},
		{
			name:       "with ellipsis at both ends",
			current:    15,
			total:      30,
			maxVisible: 7,
			want:       []int{1, -1, 13, 14, 15, 16, 17, -1, 30},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generatePageNumbers(tt.current, tt.total, tt.maxVisible)

			if len(got) != len(tt.want) {
				t.Errorf("generatePageNumbers() = %v, want %v", got, tt.want)
				return
			}

			for i, v := range got {
				if v != tt.want[i] {
					t.Errorf("generatePageNumbers()[%d] = %v, want %v", i, v, tt.want[i])
				}
			}
		})
	}
}