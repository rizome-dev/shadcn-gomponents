package hovercard

import (
	"strings"
	"testing"

	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func renderToString(node g.Node) string {
	var buf strings.Builder
	node.Render(&buf)
	return buf.String()
}

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		props    Props
		children []g.Node
		want     string
	}{
		{
			name:  "basic hover card",
			props: Props{},
			want:  `data-hover-card="root" data-state="closed"`,
		},
		{
			name:  "open hover card",
			props: Props{Open: true},
			want:  `data-state="open"`,
		},
		{
			name:  "with custom class",
			props: Props{Class: "custom-hover"},
			want:  `class="relative inline-block custom-hover"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := renderToString(New(tt.props, tt.children...))
			if !strings.Contains(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrigger(t *testing.T) {
	tests := []struct {
		name  string
		props TriggerProps
		want  string
	}{
		{
			name:  "basic trigger",
			props: TriggerProps{},
			want:  `data-hover-card="trigger"`,
		},
		{
			name:  "with aria attributes",
			props: TriggerProps{},
			want:  `aria-haspopup="dialog" aria-expanded="false"`,
		},
		{
			name:  "with custom class",
			props: TriggerProps{Class: "custom-trigger"},
			want:  `class="cursor-pointer custom-trigger"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := renderToString(Trigger(tt.props, g.Text("Hover me")))
			if !strings.Contains(got, tt.want) {
				t.Errorf("Trigger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContent(t *testing.T) {
	tests := []struct {
		name  string
		props ContentProps
		want  []string
	}{
		{
			name:  "basic content",
			props: ContentProps{},
			want: []string{
				`data-hover-card="content"`,
				`data-state="closed"`,
				`data-side="bottom"`,
				`data-align="center"`,
				`role="dialog"`,
			},
		},
		{
			name: "with custom positioning",
			props: ContentProps{
				Side:  "top",
				Align: "start",
			},
			want: []string{
				`data-side="top"`,
				`data-align="start"`,
			},
		},
		{
			name:  "with custom class",
			props: ContentProps{Class: "custom-content"},
			want:  []string{`custom-content`},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := renderToString(Content(tt.props, g.Text("Content")))
			for _, want := range tt.want {
				if !strings.Contains(got, want) {
					t.Errorf("Content() missing: %v", want)
				}
			}
		})
	}
}

func TestProfileCard(t *testing.T) {
	card := ProfileCard("johndoe", "John Doe", "Software Developer", "")
	got := renderToString(card)
	
	wants := []string{
		`John Doe`,
		`@johndoe`,
		`Software Developer`,
		`Following`,
		`Followers`,
	}
	
	for _, want := range wants {
		if !strings.Contains(got, want) {
			t.Errorf("ProfileCard() missing: %v", want)
		}
	}
}

func TestLinkPreview(t *testing.T) {
	preview := LinkPreview(
		"https://example.com",
		"Example Site",
		"This is an example website",
		"",
	)
	got := renderToString(preview)
	
	wants := []string{
		`https://example.com`,
		`Example Site`,
		`This is an example website`,
		`w-80`, // wider content
	}
	
	for _, want := range wants {
		if !strings.Contains(got, want) {
			t.Errorf("LinkPreview() missing: %v", want)
		}
	}
}

func TestCalendar(t *testing.T) {
	tests := []struct {
		name   string
		date   string
		events []string
		wants  []string
	}{
		{
			name: "with events",
			date: "Today",
			events: []string{
				"9:00 AM - Meeting",
				"2:00 PM - Review",
			},
			wants: []string{
				`Today`,
				`9:00 AM - Meeting`,
				`2:00 PM - Review`,
			},
		},
		{
			name:   "no events",
			date:   "Tomorrow",
			events: []string{},
			wants: []string{
				`Tomorrow`,
				`No events scheduled`,
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := renderToString(Calendar(tt.date, tt.events))
			for _, want := range tt.wants {
				if !strings.Contains(got, want) {
					t.Errorf("Calendar() missing: %v", want)
				}
			}
		})
	}
}

func TestCompleteHoverCard(t *testing.T) {
	card := New(
		Props{},
		Trigger(
			TriggerProps{},
			Button(
				Type("button"),
				Class("rounded-md border px-3 py-2"),
				g.Text("Hover for info"),
			),
		),
		Content(
			ContentProps{Side: "top", Align: "start"},
			Div(Class("space-y-2"),
				H4(Class("font-semibold"), g.Text("Information")),
				P(Class("text-sm"), g.Text("This is helpful information that appears on hover.")),
				Div(Class("flex gap-2"),
					Button(
						Type("button"),
						Class("text-xs border rounded px-2 py-1"),
						g.Text("Action 1"),
					),
					Button(
						Type("button"),
						Class("text-xs border rounded px-2 py-1"),
						g.Text("Action 2"),
					),
				),
			),
		),
	)
	
	got := renderToString(card)
	
	wants := []string{
		`data-hover-card="root"`,
		`data-hover-card="trigger"`,
		`Hover for info`,
		`data-hover-card="content"`,
		`data-side="top"`,
		`data-align="start"`,
		`Information`,
		`This is helpful information`,
		`Action 1`,
		`Action 2`,
	}
	
	for _, want := range wants {
		if !strings.Contains(got, want) {
			t.Errorf("Complete hover card missing: %v", want)
		}
	}
}

func TestHTMXHoverCard(t *testing.T) {
	htmxProps := HTMXProps{
		ID:          "test-hover",
		ContentPath: "/api/hover/test",
		Delay:       500,
	}
	
	card := NewHTMX(
		Props{},
		htmxProps,
		TriggerHTMX(
			TriggerProps{},
			g.Text("Hover me"),
		),
	)
	
	got := renderToString(card)
	
	wants := []string{
		`id="test-hover"`,
		`test-hover-content`,
		`/api/hover/test`,
		`500`, // delay
		`mouseenter`,
		`mouseleave`,
	}
	
	for _, want := range wants {
		if !strings.Contains(got, want) {
			t.Errorf("HTMX hover card missing: %v", want)
		}
	}
}