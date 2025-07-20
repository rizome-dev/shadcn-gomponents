package card

import (
	"bytes"
	"strings"
	"testing"

	g "maragu.dev/gomponents"
)

func TestCard(t *testing.T) {
	tests := []struct {
		name     string
		card     g.Node
		contains []string
	}{
		{
			name: "basic card",
			card: Card(
				CardContent(g.Text("Hello, World!")),
			),
			contains: []string{
				`<div class="bg-card text-card-foreground flex flex-col gap-6 rounded-xl border py-6 shadow-sm"`,
				`Hello, World!`,
				`px-6`,
			},
		},
		{
			name: "card with header",
			card: Card(
				CardHeader(
					CardTitle(g.Text("Card Title")),
					CardDescription(g.Text("Card Description")),
				),
				CardContent(g.Text("Card content goes here")),
			),
			contains: []string{
				`<h3 class="leading-none font-semibold">Card Title</h3>`,
				`<p class="text-muted-foreground text-sm">Card Description</p>`,
				`Card content goes here`,
			},
		},
		{
			name: "card with footer",
			card: Card(
				CardHeader(CardTitle(g.Text("Title"))),
				CardContent(g.Text("Content")),
				CardFooter(g.Text("Footer")),
			),
			contains: []string{
				`Title`,
				`Content`,
				`Footer`,
				`flex items-center px-6`,
			},
		},
		{
			name: "card with action",
			card: Card(
				CardHeader(
					CardTitle(g.Text("Title")),
					CardAction(g.Text("Action")),
				),
			),
			contains: []string{
				`Title`,
				`Action`,
				`data-slot="card-action"`,
				`col-start-2 row-span-2`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := tt.card.Render(&buf)
			if err != nil {
				t.Fatalf("Render() error = %v", err)
			}

			result := buf.String()
			for _, expected := range tt.contains {
				if !strings.Contains(result, expected) {
					t.Errorf("Expected output to contain %q, but it didn't.\nGot: %s", expected, result)
				}
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		props    Props
		content  []g.Node
		contains []string
	}{
		{
			name: "card with props",
			props: Props{
				Title:       "My Card",
				Description: "This is a card",
			},
			content: []g.Node{g.Text("Card content")},
			contains: []string{
				`My Card`,
				`This is a card`,
				`Card content`,
			},
		},
		{
			name: "card with custom class",
			props: Props{
				Title: "Custom Card",
				Class: "w-96",
			},
			contains: []string{
				`Custom Card`,
				`w-96`,
			},
		},
		{
			name:  "card with only content",
			props: Props{},
			content: []g.Node{
				g.Text("Just content"),
			},
			contains: []string{
				`Just content`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			card := New(tt.props, tt.content...)
			err := card.Render(&buf)
			if err != nil {
				t.Fatalf("Render() error = %v", err)
			}

			result := buf.String()
			for _, expected := range tt.contains {
				if !strings.Contains(result, expected) {
					t.Errorf("Expected output to contain %q, but it didn't.\nGot: %s", expected, result)
				}
			}
		})
	}
}

func TestHelpers(t *testing.T) {
	tests := []struct {
		name     string
		fn       func() g.Node
		contains []string
	}{
		{
			name: "WithFooter helper",
			fn: func() g.Node {
				return WithFooter(
					"Card Title",
					"Card Description",
					[]g.Node{g.Text("Content")},
					[]g.Node{g.Text("Footer")},
				)
			},
			contains: []string{
				`Card Title`,
				`Card Description`,
				`Content`,
				`Footer`,
			},
		},
		{
			name: "Simple helper",
			fn: func() g.Node {
				return Simple(g.Text("Simple card content"))
			},
			contains: []string{
				`Simple card content`,
				`px-6`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			card := tt.fn()
			err := card.Render(&buf)
			if err != nil {
				t.Fatalf("Render() error = %v", err)
			}

			result := buf.String()
			for _, expected := range tt.contains {
				if !strings.Contains(result, expected) {
					t.Errorf("Expected output to contain %q, but it didn't.\nGot: %s", expected, result)
				}
			}
		})
	}
}