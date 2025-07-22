package sheet

import (
	"strings"
	"testing"
	
	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func renderToString(node g.Node) string {
	var buf strings.Builder
	err := node.Render(&buf)
	if err != nil {
		panic(err) // For tests, panic on render error
	}
	return buf.String()
}

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		props    Props
		children []g.Node
		wantContains []string
	}{
		{
			name:  "closed sheet",
			props: Props{Open: false},
			children: []g.Node{
				g.Text("Content"),
			},
			wantContains: []string{
				"data-state=\"closed\"",
			},
		},
		{
			name: "open sheet",
			props: Props{
				Open: true,
			},
			children: []g.Node{
				g.Text("Content"),
			},
			wantContains: []string{
				"fixed inset-0 z-50",
				"data-state=\"open\"",
				"Content",
			},
		},
		{
			name: "sheet with custom class",
			props: Props{
				Open:  true,
				Class: "custom-sheet",
			},
			wantContains: []string{
				"custom-sheet",
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.props, tt.children...)
			gotStr := renderToString(got)
			
			for _, want := range tt.wantContains {
				if !strings.Contains(gotStr, want) {
					t.Errorf("expected output to contain %q, got: %s", want, gotStr)
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
		wantContains []string
	}{
		{
			name:  "default trigger",
			props: TriggerProps{},
			children: []g.Node{
				g.Text("Open Sheet"),
			},
			wantContains: []string{
				"type=\"button\"",
				"aria-haspopup=\"dialog\"",
				"Open Sheet",
			},
		},
		{
			name: "trigger with custom class",
			props: TriggerProps{
				Class: "custom-trigger",
			},
			wantContains: []string{
				"class=\"custom-trigger\"",
			},
		},
		{
			name: "trigger as child",
			props: TriggerProps{
				AsChild: true,
			},
			children: []g.Node{
				Div(Class("custom-element"), g.Text("Custom")),
			},
			wantContains: []string{
				"<div class=\"custom-element\">Custom</div>",
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Trigger(tt.props, tt.children...)
			gotStr := renderToString(got)
			
			for _, want := range tt.wantContains {
				if !strings.Contains(gotStr, want) {
					t.Errorf("expected output to contain %q, got: %s", want, gotStr)
				}
			}
		})
	}
}

func TestOverlay(t *testing.T) {
	tests := []struct {
		name     string
		props    OverlayProps
		wantContains []string
	}{
		{
			name:  "default overlay",
			props: OverlayProps{},
			wantContains: []string{
				"fixed inset-0 z-50 bg-black/80",
				"data-sheet-overlay",
				"data-state=\"open\"",
				"animate-in",
				"fade-in-0",
			},
		},
		{
			name: "overlay with custom class",
			props: OverlayProps{
				Class: "custom-overlay",
			},
			wantContains: []string{
				"custom-overlay",
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Overlay(tt.props)
			gotStr := renderToString(got)
			
			for _, want := range tt.wantContains {
				if !strings.Contains(gotStr, want) {
					t.Errorf("expected output to contain %q, got: %s", want, gotStr)
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
		wantContains []string
	}{
		{
			name:  "default content (right)",
			props: ContentProps{},
			children: []g.Node{
				g.Text("Sheet content"),
			},
			wantContains: []string{
				"role=\"dialog\"",
				"aria-modal=\"true\"",
				"data-state=\"open\"",
				"data-sheet-content",
				"data-side=\"right\"",
				"inset-y-0 right-0",
				"Sheet content",
			},
		},
		{
			name: "left side content",
			props: ContentProps{
				Side: "left",
			},
			wantContains: []string{
				"data-side=\"left\"",
				"inset-y-0 left-0",
				"slide-in-from-left",
			},
		},
		{
			name: "top side content",
			props: ContentProps{
				Side: "top",
			},
			wantContains: []string{
				"data-side=\"top\"",
				"inset-x-0 top-0",
				"slide-in-from-top",
			},
		},
		{
			name: "bottom side content",
			props: ContentProps{
				Side: "bottom",
			},
			wantContains: []string{
				"data-side=\"bottom\"",
				"inset-x-0 bottom-0",
				"slide-in-from-bottom",
			},
		},
		{
			name: "content with close button",
			props: ContentProps{
				ShowCloseButton: true,
			},
			wantContains: []string{
				"aria-label=\"Close\"",
				"<svg",
				"sr-only",
			},
		},
		{
			name: "content with custom class",
			props: ContentProps{
				Class: "custom-content",
			},
			wantContains: []string{
				"custom-content",
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContentComponent(tt.props, tt.children...)
			gotStr := renderToString(got)
			
			for _, want := range tt.wantContains {
				if !strings.Contains(gotStr, want) {
					t.Errorf("expected output to contain %q, got: %s", want, gotStr)
				}
			}
		})
	}
}

func TestHeader(t *testing.T) {
	got := HeaderComponent(
		HeaderProps{},
		g.Text("Header content"),
	)
	gotStr := renderToString(got)
	
	wantContains := []string{
		"flex flex-col space-y-2",
		"Header content",
	}
	
	for _, want := range wantContains {
		if !strings.Contains(gotStr, want) {
			t.Errorf("expected output to contain %q", want)
		}
	}
}

func TestFooter(t *testing.T) {
	got := FooterComponent(
		FooterProps{},
		g.Text("Footer content"),
	)
	gotStr := renderToString(got)
	
	wantContains := []string{
		"flex flex-col-reverse sm:flex-row sm:justify-end sm:space-x-2",
		"Footer content",
	}
	
	for _, want := range wantContains {
		if !strings.Contains(gotStr, want) {
			t.Errorf("expected output to contain %q", want)
		}
	}
}

func TestTitle(t *testing.T) {
	got := TitleComponent(
		TitleProps{},
		g.Text("Sheet Title"),
	)
	gotStr := renderToString(got)
	
	wantContains := []string{
		"<h2",
		"text-lg font-semibold",
		"Sheet Title",
	}
	
	for _, want := range wantContains {
		if !strings.Contains(gotStr, want) {
			t.Errorf("expected output to contain %q", want)
		}
	}
}

func TestDescription(t *testing.T) {
	got := Description(
		DescriptionProps{},
		g.Text("Sheet description"),
	)
	gotStr := renderToString(got)
	
	wantContains := []string{
		"<p",
		"text-sm text-muted-foreground",
		"Sheet description",
	}
	
	for _, want := range wantContains {
		if !strings.Contains(gotStr, want) {
			t.Errorf("expected output to contain %q", want)
		}
	}
}

func TestClose(t *testing.T) {
	got := Close(
		CloseProps{Class: "custom-close"},
		g.Text("Cancel"),
	)
	gotStr := renderToString(got)
	
	wantContains := []string{
		"type=\"button\"",
		"custom-close",
		"Cancel",
	}
	
	for _, want := range wantContains {
		if !strings.Contains(gotStr, want) {
			t.Errorf("expected output to contain %q", want)
		}
	}
}

func TestWithForm(t *testing.T) {
	got := WithForm(
		Props{Open: true},
		ContentProps{Side: "right"},
		"/api/submit",
		g.Text("Form content"),
	)
	gotStr := renderToString(got)
	
	wantContains := []string{
		"<form",
		"action=\"/api/submit\"",
		"method=\"POST\"",
		"Form content",
		"data-sheet-overlay",
		"data-sheet-content",
	}
	
	for _, want := range wantContains {
		if !strings.Contains(gotStr, want) {
			t.Errorf("expected output to contain %q", want)
		}
	}
}

func TestRightSheet(t *testing.T) {
	got := RightSheet(
		Props{Open: true},
		g.Text("Right sheet content"),
	)
	gotStr := renderToString(got)
	
	wantContains := []string{
		"data-side=\"right\"",
		"Right sheet content",
		"aria-label=\"Close\"", // Should have close button
	}
	
	for _, want := range wantContains {
		if !strings.Contains(gotStr, want) {
			t.Errorf("expected output to contain %q", want)
		}
	}
}

func TestLeftSheet(t *testing.T) {
	got := LeftSheet(
		Props{Open: true},
		g.Text("Left sheet content"),
	)
	gotStr := renderToString(got)
	
	wantContains := []string{
		"data-side=\"left\"",
		"Left sheet content",
	}
	
	for _, want := range wantContains {
		if !strings.Contains(gotStr, want) {
			t.Errorf("expected output to contain %q", want)
		}
	}
}

func TestTopSheet(t *testing.T) {
	got := TopSheet(
		Props{Open: true},
		g.Text("Top sheet content"),
	)
	gotStr := renderToString(got)
	
	wantContains := []string{
		"data-side=\"top\"",
		"Top sheet content",
	}
	
	for _, want := range wantContains {
		if !strings.Contains(gotStr, want) {
			t.Errorf("expected output to contain %q", want)
		}
	}
}

func TestBottomSheet(t *testing.T) {
	got := BottomSheet(
		Props{Open: true},
		g.Text("Bottom sheet content"),
	)
	gotStr := renderToString(got)
	
	wantContains := []string{
		"data-side=\"bottom\"",
		"Bottom sheet content",
	}
	
	for _, want := range wantContains {
		if !strings.Contains(gotStr, want) {
			t.Errorf("expected output to contain %q", want)
		}
	}
}

func TestMobileSheet(t *testing.T) {
	got := MobileSheet(
		Props{Open: true},
		g.Text("Mobile sheet content"),
	)
	gotStr := renderToString(got)
	
	wantContains := []string{
		"data-side=\"bottom\"",
		"h-[90vh]",
		"sm:h-full",
		"sm:inset-y-0",
		"sm:right-0",
		"Mobile sheet content",
	}
	
	for _, want := range wantContains {
		if !strings.Contains(gotStr, want) {
			t.Errorf("expected output to contain %q", want)
		}
	}
}