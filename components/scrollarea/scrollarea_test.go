package scrollarea

import (
	"strings"
	"testing"
	
	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		props    Props
		children []g.Node
		wantContains []string
	}{
		{
			name:  "default scroll area",
			props: Props{},
			children: []g.Node{
				g.Text("Content"),
			},
			wantContains: []string{
				"data-scroll-area",
				"data-orientation=\"vertical\"",
				"data-type=\"hover\"",
				"relative overflow-hidden",
				"Content",
			},
		},
		{
			name: "horizontal scroll area",
			props: Props{
				Orientation: "horizontal",
			},
			wantContains: []string{
				"data-orientation=\"horizontal\"",
			},
		},
		{
			name: "scroll area with custom type",
			props: Props{
				Type: "always",
			},
			wantContains: []string{
				"data-type=\"always\"",
			},
		},
		{
			name: "scroll area with hide delay",
			props: Props{
				ScrollHideDelay: 1000,
			},
			wantContains: []string{
				"data-scroll-hide-delay=\"1000\"",
			},
		},
		{
			name: "rtl scroll area",
			props: Props{
				Dir: "rtl",
			},
			wantContains: []string{
				"dir=\"rtl\"",
			},
		},
		{
			name: "scroll area with custom class",
			props: Props{
				Class: "custom-scroll",
			},
			wantContains: []string{
				"custom-scroll",
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.props, tt.children...)
			gotStr := got.Render()
			
			for _, want := range tt.wantContains {
				if !strings.Contains(gotStr, want) {
					t.Errorf("expected output to contain %q, got: %s", want, gotStr)
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
		wantContains []string
	}{
		{
			name:  "default viewport",
			props: ViewportProps{},
			children: []g.Node{
				g.Text("Scrollable content"),
			},
			wantContains: []string{
				"data-scroll-area-viewport",
				"h-full w-full rounded-[inherit]",
				"overflow: scroll",
				"-ms-overflow-style: none",
				"scrollbar-width: none",
				"Scrollable content",
			},
		},
		{
			name: "viewport with custom class",
			props: ViewportProps{
				Class: "custom-viewport",
			},
			wantContains: []string{
				"custom-viewport",
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Viewport(tt.props, tt.children...)
			gotStr := got.Render()
			
			for _, want := range tt.wantContains {
				if !strings.Contains(gotStr, want) {
					t.Errorf("expected output to contain %q, got: %s", want, gotStr)
				}
			}
		})
	}
}

func TestScrollbar(t *testing.T) {
	tests := []struct {
		name     string
		props    ScrollbarProps
		children []g.Node
		wantContains []string
	}{
		{
			name:  "default vertical scrollbar",
			props: ScrollbarProps{},
			wantContains: []string{
				"data-scroll-area-scrollbar",
				"data-orientation=\"vertical\"",
				"h-full w-2.5",
				"border-l border-l-transparent",
			},
		},
		{
			name: "horizontal scrollbar",
			props: ScrollbarProps{
				Orientation: "horizontal",
			},
			wantContains: []string{
				"data-orientation=\"horizontal\"",
				"h-2.5 w-full",
				"border-t border-t-transparent",
			},
		},
		{
			name: "force mounted scrollbar",
			props: ScrollbarProps{
				ForceMount: true,
			},
			wantContains: []string{
				"data-state=\"visible\"",
			},
		},
		{
			name: "scrollbar with custom class",
			props: ScrollbarProps{
				Class: "custom-scrollbar",
			},
			wantContains: []string{
				"custom-scrollbar",
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Scrollbar(tt.props, tt.children...)
			gotStr := got.Render()
			
			for _, want := range tt.wantContains {
				if !strings.Contains(gotStr, want) {
					t.Errorf("expected output to contain %q, got: %s", want, gotStr)
				}
			}
		})
	}
}

func TestThumb(t *testing.T) {
	tests := []struct {
		name     string
		props    ThumbProps
		wantContains []string
	}{
		{
			name:  "default thumb",
			props: ThumbProps{},
			wantContains: []string{
				"data-scroll-area-thumb",
				"relative flex-1 rounded-full bg-border",
				"position: relative",
			},
		},
		{
			name: "thumb with custom class",
			props: ThumbProps{
				Class: "bg-primary",
			},
			wantContains: []string{
				"bg-primary",
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Thumb(tt.props)
			gotStr := got.Render()
			
			for _, want := range tt.wantContains {
				if !strings.Contains(gotStr, want) {
					t.Errorf("expected output to contain %q, got: %s", want, gotStr)
				}
			}
		})
	}
}

func TestCorner(t *testing.T) {
	tests := []struct {
		name     string
		props    CornerProps
		wantContains []string
	}{
		{
			name:  "default corner",
			props: CornerProps{},
			wantContains: []string{
				"data-scroll-area-corner",
				"absolute right-0 bottom-0",
			},
		},
		{
			name: "corner with custom class",
			props: CornerProps{
				Class: "bg-muted",
			},
			wantContains: []string{
				"bg-muted",
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Corner(tt.props)
			gotStr := got.Render()
			
			for _, want := range tt.wantContains {
				if !strings.Contains(gotStr, want) {
					t.Errorf("expected output to contain %q, got: %s", want, gotStr)
				}
			}
		})
	}
}

func TestScrollAreaWithBar(t *testing.T) {
	got := ScrollAreaWithBar(
		Props{},
		g.Text("Content"),
	)
	gotStr := got.Render()
	
	wantContains := []string{
		"data-type=\"always\"",
		"data-scroll-area-viewport",
		"data-scroll-area-scrollbar",
		"data-scroll-area-thumb",
		"Content",
	}
	
	for _, want := range wantContains {
		if !strings.Contains(gotStr, want) {
			t.Errorf("expected output to contain %q", want)
		}
	}
}

func TestHorizontalScrollArea(t *testing.T) {
	got := HorizontalScrollArea(
		Props{},
		g.Text("Horizontal content"),
	)
	gotStr := got.Render()
	
	wantContains := []string{
		"data-orientation=\"horizontal\"",
		"data-scroll-area-viewport",
		"data-orientation=\"horizontal\"", // scrollbar should also be horizontal
		"Horizontal content",
	}
	
	for _, want := range wantContains {
		if !strings.Contains(gotStr, want) {
			t.Errorf("expected output to contain %q", want)
		}
	}
}

func TestScrollAreaAuto(t *testing.T) {
	got := ScrollAreaAuto(
		Props{},
		g.Text("Auto content"),
	)
	gotStr := got.Render()
	
	if !strings.Contains(gotStr, "data-type=\"auto\"") {
		t.Error("expected scroll area to have auto type")
	}
}

func TestScrollAreaHover(t *testing.T) {
	got := ScrollAreaHover(
		Props{},
		g.Text("Hover content"),
	)
	gotStr := got.Render()
	
	wantContains := []string{
		"data-type=\"hover\"",
		"data-scroll-hide-delay=\"600\"",
		"Hover content",
	}
	
	for _, want := range wantContains {
		if !strings.Contains(gotStr, want) {
			t.Errorf("expected output to contain %q", want)
		}
	}
}

func TestCodeScrollArea(t *testing.T) {
	got := CodeScrollArea(
		"custom-code",
		g.Text("const code = 'hello';"),
	)
	gotStr := got.Render()
	
	wantContains := []string{
		"data-orientation=\"both\"",
		"rounded-md border",
		"custom-code",
		"<pre",
		"<code",
		"const code = 'hello';",
	}
	
	for _, want := range wantContains {
		if !strings.Contains(gotStr, want) {
			t.Errorf("expected output to contain %q", want)
		}
	}
}

func TestListScrollArea(t *testing.T) {
	got := ListScrollArea(
		"200px",
		Li(g.Text("Item 1")),
		Li(g.Text("Item 2")),
	)
	gotStr := got.Render()
	
	wantContains := []string{
		"data-type=\"auto\"",
		"max-h-[200px]",
		"Item 1",
		"Item 2",
	}
	
	for _, want := range wantContains {
		if !strings.Contains(gotStr, want) {
			t.Errorf("expected output to contain %q", want)
		}
	}
}

func TestChatScrollArea(t *testing.T) {
	got := ChatScrollArea(
		Div(g.Text("Message 1")),
		Div(g.Text("Message 2")),
	)
	gotStr := got.Render()
	
	wantContains := []string{
		"data-type=\"auto\"",
		"h-full",
		"flex flex-col gap-4",
		"Message 1",
		"Message 2",
	}
	
	for _, want := range wantContains {
		if !strings.Contains(gotStr, want) {
			t.Errorf("expected output to contain %q", want)
		}
	}
}

func TestImageGalleryScrollArea(t *testing.T) {
	got := ImageGalleryScrollArea(
		Img(Src("1.jpg")),
		Img(Src("2.jpg")),
	)
	gotStr := got.Render()
	
	wantContains := []string{
		"data-orientation=\"horizontal\"",
		"w-full whitespace-nowrap",
		"flex w-max space-x-4",
		"1.jpg",
		"2.jpg",
	}
	
	for _, want := range wantContains {
		if !strings.Contains(gotStr, want) {
			t.Errorf("expected output to contain %q", want)
		}
	}
}

func TestTableScrollArea(t *testing.T) {
	got := TableScrollArea(
		Table(
			Tr(Td(g.Text("Cell"))),
		),
	)
	gotStr := got.Render()
	
	wantContains := []string{
		"data-orientation=\"both\"",
		"<table>",
		"Cell",
	}
	
	for _, want := range wantContains {
		if !strings.Contains(gotStr, want) {
			t.Errorf("expected output to contain %q", want)
		}
	}
}