package aspectratio

import (
	"bytes"
	"fmt"
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

func TestAspectRatio(t *testing.T) {
	tests := []struct {
		name     string
		props    Props
		children []g.Node
		contains []string
	}{
		{
			name: "16:9 aspect ratio",
			props: Props{
				Ratio: 16.0 / 9.0,
			},
			children: []g.Node{
				Img(Src("test.jpg"), Alt("Test image")),
			},
			contains: []string{
				`data-slot="aspect-ratio"`,
				`data-aspect-ratio="1.78"`,
				`padding-bottom: 56.2500%`,
				`relative overflow-hidden`,
				`absolute inset-0`,
				`<img src="test.jpg" alt="Test image">`,
			},
		},
		{
			name: "4:3 aspect ratio",
			props: Props{
				Ratio: 4.0 / 3.0,
			},
			children: []g.Node{
				Div(g.Text("Content")),
			},
			contains: []string{
				`data-aspect-ratio="1.33"`,
				`padding-bottom: 75.0000%`,
				"Content",
			},
		},
		{
			name: "1:1 aspect ratio (square)",
			props: Props{
				Ratio: 1,
			},
			children: []g.Node{
				Div(g.Text("Square content")),
			},
			contains: []string{
				`data-aspect-ratio="1.00"`,
				`padding-bottom: 100.0000%`,
				"Square content",
			},
		},
		{
			name: "default ratio (no ratio specified)",
			props: Props{},
			children: []g.Node{
				Div(g.Text("Default ratio")),
			},
			contains: []string{
				`data-aspect-ratio="1.00"`,
				`padding-bottom: 100.0000%`,
				"Default ratio",
			},
		},
		{
			name: "with custom class",
			props: Props{
				Ratio: 16.0 / 9.0,
				Class: "bg-muted rounded-lg",
			},
			children: []g.Node{
				Div(g.Text("Styled content")),
			},
			contains: []string{
				"bg-muted rounded-lg",
				"Styled content",
			},
		},
		{
			name: "cinema ratio (2.39:1)",
			props: Props{
				Ratio: 2.39,
			},
			children: []g.Node{
				Div(g.Text("Cinematic")),
			},
			contains: []string{
				`data-aspect-ratio="2.39"`,
				`padding-bottom: 41.8410%`,
				"Cinematic",
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

func TestHelperFunctions(t *testing.T) {
	tests := []struct {
		name      string
		component g.Node
		contains  []string
	}{
		{
			name:      "Square helper",
			component: Square(Div(g.Text("Square"))),
			contains: []string{
				`data-aspect-ratio="1.00"`,
				`padding-bottom: 100.0000%`,
				"Square",
			},
		},
		{
			name:      "Video helper (16:9)",
			component: Video16x9(Div(g.Text("Video"))),
			contains: []string{
				`data-aspect-ratio="1.78"`,
				`padding-bottom: 56.2500%`,
				"Video",
			},
		},
		{
			name:      "Portrait helper (4:5)",
			component: Portrait(Div(g.Text("Portrait"))),
			contains: []string{
				`data-aspect-ratio="0.80"`,
				`padding-bottom: 125.0000%`,
				"Portrait",
			},
		},
		{
			name:      "Landscape helper (3:2)",
			component: Landscape(Div(g.Text("Landscape"))),
			contains: []string{
				`data-aspect-ratio="1.50"`,
				`padding-bottom: 66.6667%`,
				"Landscape",
			},
		},
		{
			name:      "Cinema helper (2.39:1)",
			component: Cinema(Div(g.Text("Cinema"))),
			contains: []string{
				`data-aspect-ratio="2.39"`,
				`padding-bottom: 41.8410%`,
				"Cinema",
			},
		},
		{
			name:      "WithClass helper",
			component: WithClass(21.0/9.0, "border rounded", Div(g.Text("Ultra-wide"))),
			contains: []string{
				`data-aspect-ratio="2.33"`,
				"border rounded",
				"Ultra-wide",
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			html := renderToString(t, tt.component)
			
			for _, substring := range tt.contains {
				if !strings.Contains(html, substring) {
					t.Errorf("expected HTML to contain %q, but it didn't.\nHTML: %s", substring, html)
				}
			}
		})
	}
}

func TestPaddingCalculation(t *testing.T) {
	tests := []struct {
		ratio           float64
		expectedPadding string
	}{
		{16.0 / 9.0, "56.2500"},   // HD video
		{4.0 / 3.0, "75.0000"},    // Traditional TV
		{1.0, "100.0000"},         // Square
		{2.0, "50.0000"},          // 2:1
		{0.5, "200.0000"},         // 1:2 (tall)
		{21.0 / 9.0, "42.8571"},   // Ultra-wide
	}
	
	for _, tt := range tests {
		t.Run(fmt.Sprintf("ratio %.2f", tt.ratio), func(t *testing.T) {
			component := New(Props{Ratio: tt.ratio}, Div())
			html := renderToString(t, component)
			
			expectedStyle := fmt.Sprintf("padding-bottom: %s%%", tt.expectedPadding)
			if !strings.Contains(html, expectedStyle) {
				t.Errorf("expected padding-bottom to be %s%%, but HTML was: %s", tt.expectedPadding, html)
			}
		})
	}
}