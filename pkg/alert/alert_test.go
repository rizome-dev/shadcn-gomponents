package alert_test

import (
	"bytes"
	"strings"
	"testing"
	g "maragu.dev/gomponents"
	"github.com/rizome-dev/shadcn-gomponents/pkg/alert"
)

func TestAlert(t *testing.T) {
	tests := []struct {
		name     string
		alert    g.Node
		contains []string
	}{
		{
			name:  "default alert",
			alert: alert.Default(alert.Title(g.Text("Default Alert")), alert.Description(g.Text("This is a default alert"))),
			contains: []string{
				`role="alert"`,
				`class="relative w-full rounded-lg border p-4`,
				`bg-card text-card-foreground`,
				`<h5 class="mb-1 font-medium leading-none tracking-tight">Default Alert</h5>`,
				`<div class="text-sm text-muted-foreground [&amp;_p]:leading-relaxed">This is a default alert</div>`,
			},
		},
		{
			name:  "destructive alert",
			alert: alert.Destructive(alert.Title(g.Text("Error")), alert.Description(g.Text("Something went wrong"))),
			contains: []string{
				`role="alert"`,
				`border-destructive/50 text-destructive`,
				`<h5 class="mb-1 font-medium leading-none tracking-tight">Error</h5>`,
				`<div class="text-sm text-muted-foreground [&amp;_p]:leading-relaxed">Something went wrong</div>`,
			},
		},
		{
			name: "alert with custom class",
			alert: alert.New(
				alert.Props{Class: "mt-4"},
				alert.Title(g.Text("Custom Alert")),
			),
			contains: []string{
				`mt-4`,
			},
		},
		{
			name: "alert with icon",
			alert: alert.WithIcon(
				g.El("svg", g.Attr("class", "h-4 w-4")),
				alert.Props{},
				alert.Title(g.Text("Info")),
				alert.Description(g.Text("This alert has an icon")),
			),
			contains: []string{
				`<svg class="h-4 w-4"></svg>`,
				`[&amp;&gt;svg]:absolute`,
				`[&amp;&gt;svg+div]:pl-7`,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := test.alert.Render(&buf)
			if err != nil {
				t.Fatalf("failed to render: %v", err)
			}
			result := buf.String()
			for _, expected := range test.contains {
				if !strings.Contains(result, expected) {
					t.Errorf("expected result to contain %q, but it didn't.\nGot: %s", expected, result)
				}
			}
		})
	}
}