package tooltip_test

import (
	"strings"
	"testing"
	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/pkg/tooltip"
)

func renderToString(node g.Node) string {
	var buf strings.Builder
	node.Render(&buf)
	return buf.String()
}

func TestTooltip(t *testing.T) {
	tests := []struct {
		name     string
		tooltip  g.Node
		contains []string
	}{
		{
			name: "simple tooltip",
			tooltip: tooltip.Simple("Hello tooltip", 
				Button(g.Text("Hover me"))),
			contains: []string{
				`data-tooltip-container="true"`,
				`role="tooltip"`,
				`data-state="closed"`,
				`data-side="top"`,
				`data-align="center"`,
				`Hello tooltip`,
				`>Hover me</button>`,
				`tooltip-content`,
				`bg-primary`,
				`text-primary-foreground`,
				`tooltip-arrow`,
			},
		},
		{
			name: "tooltip with custom side",
			tooltip: tooltip.New(tooltip.Props{
				Content: "Right side tooltip",
				Side:    tooltip.SideRight,
			}, 
				Span(g.Text("Trigger")),
				g.Text("Right side tooltip")),
			contains: []string{
				`data-side="right"`,
				`left-full ml-2`,
				`>Trigger</span>`,
			},
		},
		{
			name: "tooltip with custom align",
			tooltip: tooltip.New(tooltip.Props{
				Content: "Start aligned",
				Side:    tooltip.SideBottom,
				Align:   tooltip.AlignStart,
			}, 
				Span(g.Text("Trigger")),
				g.Text("Start aligned")),
			contains: []string{
				`data-side="bottom"`,
				`data-align="start"`,
				`top-full mt-2`,
			},
		},
		{
			name: "tooltip with delay",
			tooltip: tooltip.WithDelay("Delayed tooltip", 500,
				Button(g.Text("Wait for it"))),
			contains: []string{
				`transition-delay: 500ms`,
				`Delayed tooltip`,
			},
		},
		{
			name: "tooltip with custom classes",
			tooltip: tooltip.New(tooltip.Props{
				Content:      "Custom styled",
				Class:        "custom-tooltip",
				ContentClass: "bg-blue-500",
				ArrowClass:   "bg-blue-500",
			},
				Span(g.Text("Custom")),
				g.Text("Custom styled")),
			contains: []string{
				`custom-tooltip`,
				`bg-blue-500`,
				`Custom styled`,
			},
		},
		{
			name: "open tooltip",
			tooltip: tooltip.New(tooltip.Props{
				Content: "I'm open",
				Open:    true,
			},
				Span(g.Text("Always visible")),
				g.Text("I'm open")),
			contains: []string{
				`data-state="open"`,
				`opacity-100 visible`,
			},
		},
		{
			name: "tooltip with side offset",
			tooltip: tooltip.New(tooltip.Props{
				Content:    "Offset tooltip",
				SideOffset: 8,
			},
				Span(g.Text("Offset")),
				g.Text("Offset tooltip")),
			contains: []string{
				`Offset tooltip`,
			},
		},
		{
			name: "info tooltip preset",
			tooltip: tooltip.InfoTooltip("This is information",
				Button(g.Text("Info"))),
			contains: []string{
				`bg-blue-600 text-white`,
				`This is information`,
				`<svg`, // Info icon
			},
		},
		{
			name: "warning tooltip preset",
			tooltip: tooltip.WarningTooltip("This is a warning",
				Button(g.Text("Warning"))),
			contains: []string{
				`bg-yellow-600 text-white`,
				`This is a warning`,
				`<svg`, // Warning icon
			},
		},
		{
			name: "error tooltip preset",
			tooltip: tooltip.ErrorTooltip("This is an error",
				Button(g.Text("Error"))),
			contains: []string{
				`bg-red-600 text-white`,
				`This is an error`,
				`<svg`, // Error icon
			},
		},
		{
			name: "tooltip with provider",
			tooltip: tooltip.Provider(200,
				tooltip.Simple("Inside provider", Span(g.Text("Trigger")))),
			contains: []string{
				`data-tooltip-provider="true"`,
				`data-delay="200"`,
				`Inside provider`,
			},
		},
		{
			name: "tooltip trigger",
			tooltip: tooltip.Trigger(tooltip.TriggerProps{
				ID:      "my-trigger",
				Class:   "custom-trigger",
				OnHover: "showTooltip()",
				OnLeave: "hideTooltip()",
				OnFocus: "focusTooltip()",
				OnBlur:  "blurTooltip()",
			}, g.Text("Trigger content")),
			contains: []string{
				`id="my-trigger"`,
				`data-tooltip-trigger="true"`,
				`aria-describedby="my-trigger-content"`,
				`custom-trigger`,
				`onmouseenter="showTooltip()"`,
				`onmouseleave="hideTooltip()"`,
				`onfocus="focusTooltip()"`,
				`onblur="blurTooltip()"`,
				`>Trigger content</span>`,
			},
		},
		{
			name: "tooltip with all sides",
			tooltip: Div(
				tooltip.New(tooltip.Props{Side: tooltip.SideTop}, Span(g.Text("T")), g.Text("Top")),
				tooltip.New(tooltip.Props{Side: tooltip.SideRight}, Span(g.Text("R")), g.Text("Right")),
				tooltip.New(tooltip.Props{Side: tooltip.SideBottom}, Span(g.Text("B")), g.Text("Bottom")),
				tooltip.New(tooltip.Props{Side: tooltip.SideLeft}, Span(g.Text("L")), g.Text("Left")),
			),
			contains: []string{
				`data-side="top"`,
				`data-side="right"`,
				`data-side="bottom"`,
				`data-side="left"`,
				`bottom-full mb-2`, // Top
				`left-full ml-2`,   // Right
				`top-full mt-2`,    // Bottom
				`right-full mr-2`,  // Left
			},
		},
		{
			name: "tooltip with all alignments",
			tooltip: Div(
				tooltip.New(tooltip.Props{Align: tooltip.AlignStart}, Span(g.Text("S")), g.Text("Start")),
				tooltip.New(tooltip.Props{Align: tooltip.AlignCenter}, Span(g.Text("C")), g.Text("Center")),
				tooltip.New(tooltip.Props{Align: tooltip.AlignEnd}, Span(g.Text("E")), g.Text("End")),
			),
			contains: []string{
				`data-align="start"`,
				`data-align="center"`,
				`data-align="end"`,
			},
		},
		{
			name: "icon tooltip with custom icon",
			tooltip: tooltip.IconTooltip("Custom icon tooltip",
				Span(Class("text-blue-500"), g.Text("ℹ")),
				Button(g.Text("Icon"))),
			contains: []string{
				`text-blue-500`,
				`ℹ`,
				`Custom icon tooltip`,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := renderToString(test.tooltip)
			for _, expected := range test.contains {
				if !strings.Contains(result, expected) {
					t.Errorf("expected result to contain %q, but it didn't.\nGot: %s", expected, result)
				}
			}
		})
	}
}

func TestTooltipPositioning(t *testing.T) {
	// Test specific positioning scenarios
	tests := []struct {
		name     string
		props    tooltip.Props
		contains []string
	}{
		{
			name: "top center positioning",
			props: tooltip.Props{
				Side:  tooltip.SideTop,
				Align: tooltip.AlignCenter,
			},
			contains: []string{
				`bottom-full mb-2`,
				`left: 50%; transform: translateX(-50%)`,
			},
		},
		{
			name: "right start positioning",
			props: tooltip.Props{
				Side:        tooltip.SideRight,
				Align:       tooltip.AlignStart,
				AlignOffset: 10,
			},
			contains: []string{
				`left-full ml-2`,
				`top: 10px`,
			},
		},
		{
			name: "bottom end positioning",
			props: tooltip.Props{
				Side:        tooltip.SideBottom,
				Align:       tooltip.AlignEnd,
				AlignOffset: 20,
			},
			contains: []string{
				`top-full mt-2`,
				`right: 20px`,
			},
		},
		{
			name: "left center positioning",
			props: tooltip.Props{
				Side:  tooltip.SideLeft,
				Align: tooltip.AlignCenter,
			},
			contains: []string{
				`right-full mr-2`,
				`top: 50%; transform: translateY(-50%)`,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tooltip := tooltip.New(test.props, 
				Span(g.Text("Trigger")),
				g.Text("Content"))
			result := renderToString(tooltip)
			
			for _, expected := range test.contains {
				if !strings.Contains(result, expected) {
					t.Errorf("expected result to contain %q, but it didn't.\nGot: %s", expected, result)
				}
			}
		})
	}
}

func TestTooltipArrowPositioning(t *testing.T) {
	// Test arrow positioning for different sides
	tests := []struct {
		side     tooltip.Side
		contains string
	}{
		{
			side:     tooltip.SideTop,
			contains: "bottom-0 left-1/2 -translate-x-1/2 translate-y-1/2",
		},
		{
			side:     tooltip.SideBottom,
			contains: "top-0 left-1/2 -translate-x-1/2 -translate-y-1/2",
		},
		{
			side:     tooltip.SideLeft,
			contains: "right-0 top-1/2 translate-x-1/2 -translate-y-1/2",
		},
		{
			side:     tooltip.SideRight,
			contains: "left-0 top-1/2 -translate-x-1/2 -translate-y-1/2",
		},
	}

	for _, test := range tests {
		t.Run(string(test.side)+" arrow", func(t *testing.T) {
			tooltip := tooltip.New(tooltip.Props{Side: test.side},
				Span(g.Text("T")),
				g.Text("Content"))
			result := renderToString(tooltip)
			
			if !strings.Contains(result, test.contains) {
				t.Errorf("expected arrow classes %q for side %s, but not found.\nGot: %s", 
					test.contains, test.side, result)
			}
		})
	}
}