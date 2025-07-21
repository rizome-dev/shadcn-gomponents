package resizable

import (
	"bytes"
	"strings"
	"testing"
	
	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func TestPanelGroup(t *testing.T) {
	tests := []struct {
		name     string
		props    Props
		children []g.Node
		wantContains []string
	}{
		{
			name:  "default horizontal panel group",
			props: Props{},
			children: []g.Node{
				g.Text("Content"),
			},
			wantContains: []string{
				"data-panel-group",
				"data-panel-group-direction=\"horizontal\"",
				"flex h-full w-full",
				"Content",
			},
		},
		{
			name: "vertical panel group",
			props: Props{
				Direction: "vertical",
			},
			children: []g.Node{
				g.Text("Content"),
			},
			wantContains: []string{
				"data-panel-group-direction=\"vertical\"",
				"flex-col",
			},
		},
		{
			name: "panel group with storage",
			props: Props{
				Storage:    true,
				StorageKey: "my-layout",
			},
			wantContains: []string{
				"data-panel-group-storage=\"my-layout\"",
			},
		},
		{
			name: "panel group with resize handlers",
			props: Props{
				OnResizeStart: "handleStart()",
				OnResizeEnd:   "handleEnd()",
			},
			wantContains: []string{
				"data-onresizestart=\"handleStart()\"",
				"data-onresizeend=\"handleEnd()\"",
			},
		},
		{
			name: "panel group with custom class",
			props: Props{
				Class: "custom-class",
			},
			wantContains: []string{
				"custom-class",
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PanelGroup(tt.props, tt.children...)
			var buf bytes.Buffer
			err := got.Render(&buf)
			if err != nil {
				t.Fatalf("failed to render: %v", err)
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

func TestPanel(t *testing.T) {
	tests := []struct {
		name     string
		props    PanelProps
		children []g.Node
		wantContains []string
	}{
		{
			name:  "default panel",
			props: PanelProps{},
			children: []g.Node{
				g.Text("Panel content"),
			},
			wantContains: []string{
				"data-panel",
				"data-panel-size=\"50\"",
				"data-panel-min-size=\"10\"",
				"data-panel-max-size=\"90\"",
				"style=\"flex: 50 50 0%\"",
				"Panel content",
			},
		},
		{
			name: "panel with custom sizes",
			props: PanelProps{
				DefaultSize: 30,
				MinSize:     20,
				MaxSize:     40,
			},
			wantContains: []string{
				"data-panel-size=\"30\"",
				"data-panel-min-size=\"20\"",
				"data-panel-max-size=\"40\"",
				"style=\"flex: 30 30 0%\"",
			},
		},
		{
			name: "collapsible panel",
			props: PanelProps{
				Collapsible:   true,
				CollapsedSize: 4,
			},
			wantContains: []string{
				"data-panel-collapsible=\"true\"",
				"data-panel-collapsed-size=\"4\"",
			},
		},
		{
			name: "panel with ID and order",
			props: PanelProps{
				ID:    "sidebar",
				Order: 1,
			},
			wantContains: []string{
				"id=\"sidebar\"",
				"data-panel-id=\"sidebar\"",
				"data-panel-order=\"1\"",
			},
		},
		{
			name: "panel with custom class",
			props: PanelProps{
				Class: "custom-panel",
			},
			wantContains: []string{
				"custom-panel",
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Panel(tt.props, tt.children...)
			var buf bytes.Buffer
			err := got.Render(&buf)
			if err != nil {
				t.Fatalf("failed to render: %v", err)
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

func TestHandle(t *testing.T) {
	tests := []struct {
		name     string
		props    HandleProps
		wantContains []string
	}{
		{
			name:  "default handle",
			props: HandleProps{},
			wantContains: []string{
				"role=\"separator\"",
				"aria-valuenow=\"50\"",
				"aria-valuemin=\"0\"",
				"aria-valuemax=\"100\"",
				"tabindex=\"0\"",
				"data-panel-resize-handle",
				"bg-border",
				"focus-visible:ring-ring",
			},
		},
		{
			name: "handle with visual indicator",
			props: HandleProps{
				WithHandle: true,
			},
			wantContains: []string{
				"data-panel-resize-handle",
				"<svg",
				"viewBox=\"0 0 24 24\"",
			},
		},
		{
			name: "disabled handle",
			props: HandleProps{
				Disabled: true,
			},
			wantContains: []string{
				"data-disabled",
				"cursor-not-allowed",
				"opacity-50",
			},
		},
		{
			name: "handle with custom class",
			props: HandleProps{
				Class: "custom-handle",
			},
			wantContains: []string{
				"custom-handle",
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Handle(tt.props)
			var buf bytes.Buffer
			err := got.Render(&buf)
			if err != nil {
				t.Fatalf("failed to render: %v", err)
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

func TestHorizontalPanelGroup(t *testing.T) {
	got := HorizontalPanelGroup(
		Props{},
		g.Text("Content"),
	)
	var buf bytes.Buffer
	err := got.Render(&buf)
	if err != nil {
		t.Fatalf("failed to render: %v", err)
	}
	gotStr := buf.String()
	
	if !strings.Contains(gotStr, "data-panel-group-direction=\"horizontal\"") {
		t.Error("expected horizontal panel group to have horizontal direction")
	}
}

func TestVerticalPanelGroup(t *testing.T) {
	got := VerticalPanelGroup(
		Props{},
		g.Text("Content"),
	)
	var buf bytes.Buffer
	err := got.Render(&buf)
	if err != nil {
		t.Fatalf("failed to render: %v", err)
	}
	gotStr := buf.String()
	
	if !strings.Contains(gotStr, "data-panel-group-direction=\"vertical\"") {
		t.Error("expected vertical panel group to have vertical direction")
	}
	
	if !strings.Contains(gotStr, "flex-col") {
		t.Error("expected vertical panel group to have flex-col class")
	}
}

func TestCollapsiblePanel(t *testing.T) {
	got := CollapsiblePanel(
		PanelProps{},
		g.Text("Collapsible content"),
	)
	var buf bytes.Buffer
	err := got.Render(&buf)
	if err != nil {
		t.Fatalf("failed to render: %v", err)
	}
	gotStr := buf.String()
	
	wantContains := []string{
		"data-panel-collapsible=\"true\"",
		"data-panel-collapsed-size=\"4\"",
		"Collapsible content",
	}
	
	for _, want := range wantContains {
		if !strings.Contains(gotStr, want) {
			t.Errorf("expected output to contain %q, got: %s", want, gotStr)
		}
	}
}

func TestResizableLayout(t *testing.T) {
	got := ResizableLayout(
		Div(g.Text("Sidebar")),
		Div(g.Text("Main")),
	)
	var buf bytes.Buffer
	err := got.Render(&buf)
	if err != nil {
		t.Fatalf("failed to render: %v", err)
	}
	gotStr := buf.String()
	
	wantContains := []string{
		"data-panel-group-storage=\"layout\"",
		"id=\"sidebar\"",
		"id=\"main\"",
		"data-panel-collapsible=\"true\"",
		"Sidebar",
		"Main",
	}
	
	for _, want := range wantContains {
		if !strings.Contains(gotStr, want) {
			t.Errorf("expected output to contain %q, got: %s", want, gotStr)
		}
	}
}

func TestTwoColumnLayout(t *testing.T) {
	tests := []struct {
		name      string
		leftSize  int
		rightSize int
		wantLeft  string
		wantRight string
	}{
		{
			name:      "default sizes",
			leftSize:  0,
			rightSize: 0,
			wantLeft:  "data-panel-size=\"50\"",
			wantRight: "data-panel-size=\"50\"",
		},
		{
			name:      "custom sizes",
			leftSize:  30,
			rightSize: 70,
			wantLeft:  "data-panel-size=\"30\"",
			wantRight: "data-panel-size=\"70\"",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoColumnLayout(
				Div(g.Text("Left")),
				Div(g.Text("Right")),
				tt.leftSize,
				tt.rightSize,
			)
			var buf bytes.Buffer
			err := got.Render(&buf)
			if err != nil {
				t.Fatalf("failed to render: %v", err)
			}
			gotStr := buf.String()
			
			if !strings.Contains(gotStr, tt.wantLeft) {
				t.Errorf("expected left panel to have %s", tt.wantLeft)
			}
			if !strings.Contains(gotStr, tt.wantRight) {
				t.Errorf("expected right panel to have %s", tt.wantRight)
			}
		})
	}
}

func TestThreeColumnLayout(t *testing.T) {
	got := ThreeColumnLayout(
		Div(g.Text("Left")),
		Div(g.Text("Center")),
		Div(g.Text("Right")),
	)
	var buf bytes.Buffer
	err := got.Render(&buf)
	if err != nil {
		t.Fatalf("failed to render: %v", err)
	}
	gotStr := buf.String()
	
	// Check that we have 3 panels
	panelCount := strings.Count(gotStr, "data-panel=\"\"")
	if panelCount != 3 {
		t.Errorf("expected 3 panels, got %d", panelCount)
	}
	
	// Check that we have 2 handles
	handleCount := strings.Count(gotStr, "data-panel-resize-handle")
	if handleCount != 2 {
		t.Errorf("expected 2 handles, got %d", handleCount)
	}
	
	// Check content
	wantContains := []string{"Left", "Center", "Right"}
	for _, want := range wantContains {
		if !strings.Contains(gotStr, want) {
			t.Errorf("expected output to contain %q", want)
		}
	}
}

func TestIDELayout(t *testing.T) {
	got := IDELayout(
		Div(g.Text("Sidebar")),
		Div(g.Text("Editor")),
		Div(g.Text("Terminal")),
	)
	var buf bytes.Buffer
	err := got.Render(&buf)
	if err != nil {
		t.Fatalf("failed to render: %v", err)
	}
	gotStr := buf.String()
	
	wantContains := []string{
		"data-panel-group-direction=\"vertical\"",
		"h-screen",
		"id=\"sidebar\"",
		"id=\"editor\"",
		"id=\"terminal\"",
		"Sidebar",
		"Editor",
		"Terminal",
	}
	
	for _, want := range wantContains {
		if !strings.Contains(gotStr, want) {
			t.Errorf("expected output to contain %q", want)
		}
	}
	
	// Check that we have nested panel groups
	groupCount := strings.Count(gotStr, "data-panel-group=\"\"")
	if groupCount != 2 {
		t.Errorf("expected 2 panel groups (nested), got %d", groupCount)
	}
}