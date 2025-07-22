package popover

import (
	"bytes"
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
		want     string
	}{
		{
			name:  "default popover (closed)",
			props: Props{},
			children: []g.Node{
				g.Text("Trigger"),
				g.Text("Content"),
			},
			want: `<div class="relative inline-block" data-state="closed">Trigger</div>`,
		},
		{
			name: "open popover",
			props: Props{
				Open: true,
			},
			children: []g.Node{
				g.Text("Trigger"),
				g.Text("Content"),
			},
			want: `<div class="relative inline-block" data-state="open">TriggerContent</div>`,
		},
		{
			name: "popover with custom class",
			props: Props{
				Class: "custom-class",
			},
			children: []g.Node{
				g.Text("Trigger"),
				g.Text("Content"),
			},
			want: `<div class="relative inline-block custom-class" data-state="closed">Trigger</div>`,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.props, tt.children...)
			var buf bytes.Buffer
			err := got.Render(&buf)
			if err != nil {
				t.Fatalf("Render() error = %v", err)
			}
			gotStr := buf.String()
			
			// Normalize whitespace for comparison
			gotStr = strings.TrimSpace(gotStr)
			tt.want = strings.TrimSpace(tt.want)
			
			if gotStr != tt.want {
				t.Errorf("\ngot:  %s\nwant: %s", gotStr, tt.want)
			}
		})
	}
}

func TestTrigger(t *testing.T) {
	tests := []struct {
		name     string
		props    TriggerProps
		children []g.Node
		want     string
	}{
		{
			name:  "default trigger",
			props: TriggerProps{},
			children: []g.Node{
				g.Text("Click me"),
			},
			want: `<button type="button" aria-haspopup="dialog" aria-expanded="false">Click me</button>`,
		},
		{
			name: "trigger with custom class",
			props: TriggerProps{
				Class: "custom-trigger",
			},
			children: []g.Node{
				g.Text("Click me"),
			},
			want: `<button type="button" class="custom-trigger" aria-haspopup="dialog" aria-expanded="false">Click me</button>`,
		},
		{
			name: "trigger as child",
			props: TriggerProps{
				AsChild: true,
			},
			children: []g.Node{
				Div(Class("custom-element"), g.Text("Custom")),
			},
			want: `<div class="custom-element">Custom</div>`,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Trigger(tt.props, tt.children...)
			var buf bytes.Buffer
			err := got.Render(&buf)
			if err != nil {
				t.Fatalf("Render() error = %v", err)
			}
			gotStr := buf.String()
			
			// Normalize whitespace for comparison
			gotStr = strings.TrimSpace(gotStr)
			tt.want = strings.TrimSpace(tt.want)
			
			if gotStr != tt.want {
				t.Errorf("\ngot:  %s\nwant: %s", gotStr, tt.want)
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
			name:  "default content",
			props: ContentProps{},
			children: []g.Node{
				g.Text("Popover content"),
			},
			wantContains: []string{
				"role=\"dialog\"",
				"data-side=\"bottom\"",
				"data-align=\"center\"",
				"Popover content",
			},
		},
		{
			name: "content with side top",
			props: ContentProps{
				Side: "top",
			},
			children: []g.Node{
				g.Text("Top content"),
			},
			wantContains: []string{
				"data-side=\"top\"",
				"bottom-full",
				"mb-2",
			},
		},
		{
			name: "content with align start",
			props: ContentProps{
				Align: "start",
			},
			children: []g.Node{
				g.Text("Start aligned"),
			},
			wantContains: []string{
				"data-align=\"start\"",
				"left-0",
			},
		},
		{
			name: "content with custom class",
			props: ContentProps{
				Class: "custom-content",
			},
			children: []g.Node{
				g.Text("Custom content"),
			},
			wantContains: []string{
				"custom-content",
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContentComponent(tt.props, tt.children...)
			var buf bytes.Buffer
			err := got.Render(&buf)
			if err != nil {
				t.Fatalf("Render() error = %v", err)
			}
			gotStr := buf.String()
			
			for _, want := range tt.wantContains {
				if !strings.Contains(gotStr, want) {
					t.Errorf("expected output to contain %q, got: %s", want, gotStr)
				}
			}
		})
	}
}

func TestGetPositionClasses(t *testing.T) {
	tests := []struct {
		name  string
		side  string
		align string
		wantContains []string
	}{
		{
			name:  "top center",
			side:  "top",
			align: "center",
			wantContains: []string{"bottom-full", "mb-2", "left-1/2", "-translate-x-1/2"},
		},
		{
			name:  "right center",
			side:  "right",
			align: "center",
			wantContains: []string{"left-full", "ml-2", "top-1/2", "-translate-y-1/2"},
		},
		{
			name:  "bottom start",
			side:  "bottom",
			align: "start",
			wantContains: []string{"top-full", "mt-2", "left-0"},
		},
		{
			name:  "left end",
			side:  "left",
			align: "end",
			wantContains: []string{"right-full", "mr-2", "bottom-0"},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getPositionClasses(tt.side, tt.align)
			
			for _, want := range tt.wantContains {
				if !strings.Contains(got, want) {
					t.Errorf("expected position classes to contain %q, got: %s", want, got)
				}
			}
		})
	}
}

func TestGetAnimationClasses(t *testing.T) {
	tests := []struct {
		name string
		side string
		wantContains []string
	}{
		{
			name: "top animations",
			side: "top",
			wantContains: []string{"slide-out-to-bottom-2", "slide-in-from-bottom-2"},
		},
		{
			name: "right animations",
			side: "right",
			wantContains: []string{"slide-out-to-left-2", "slide-in-from-left-2"},
		},
		{
			name: "bottom animations",
			side: "bottom",
			wantContains: []string{"slide-out-to-top-2", "slide-in-from-top-2"},
		},
		{
			name: "left animations",
			side: "left",
			wantContains: []string{"slide-out-to-right-2", "slide-in-from-right-2"},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getAnimationClasses(tt.side)
			
			for _, want := range tt.wantContains {
				if !strings.Contains(got, want) {
					t.Errorf("expected animation classes to contain %q, got: %s", want, got)
				}
			}
		})
	}
}

func TestWithArrow(t *testing.T) {
	got := WithArrow(
		ContentProps{Side: "bottom"},
		g.Text("Content with arrow"),
	)
	var buf bytes.Buffer
	err := got.Render(&buf)
	if err != nil {
		t.Fatalf("Render() error = %v", err)
	}
	gotStr := buf.String()
	
	// Check for arrow element
	if !strings.Contains(gotStr, "rotate-45") {
		t.Error("expected output to contain arrow with rotate-45 class")
	}
	
	// Check for content
	if !strings.Contains(gotStr, "Content with arrow") {
		t.Error("expected output to contain content text")
	}
}

func TestClose(t *testing.T) {
	tests := []struct {
		name  string
		class []string
		wantContains []string
	}{
		{
			name:  "default close button",
			class: []string{},
			wantContains: []string{
				"type=\"button\"",
				"aria-label=\"Close\"",
				"absolute right-2 top-2",
			},
		},
		{
			name:  "close button with custom class",
			class: []string{"custom-close"},
			wantContains: []string{
				"custom-close",
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Close(tt.class...)
			var buf bytes.Buffer
			err := got.Render(&buf)
			if err != nil {
				t.Fatalf("Render() error = %v", err)
			}
			gotStr := buf.String()
			
			for _, want := range tt.wantContains {
				if !strings.Contains(gotStr, want) {
					t.Errorf("expected output to contain %q, got: %s", want, gotStr)
				}
			}
		})
	}
}