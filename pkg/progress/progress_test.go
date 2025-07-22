package progress_test

import (
	"bytes"
	"strings"
	"testing"
	g "maragu.dev/gomponents"
	"github.com/rizome-dev/shadcn-gomponents/pkg/progress"
)

func TestProgress(t *testing.T) {
	tests := []struct {
		name     string
		progress g.Node
		contains []string
	}{
		{
			name:     "default progress",
			progress: progress.Default(50),
			contains: []string{
				`role="progressbar"`,
				`aria-valuemin="0"`,
				`aria-valuemax="100"`,
				`aria-valuenow="50"`,
				`aria-label="Progress: 50%"`,
				`data-value="50"`,
				`width: 50%`,
				`class="relative w-full overflow-hidden rounded-full bg-secondary h-4"`,
			},
		},
		{
			name:     "zero progress",
			progress: progress.Default(0),
			contains: []string{
				`aria-valuenow="0"`,
				`width: 0%`,
			},
		},
		{
			name:     "complete progress",
			progress: progress.Default(100),
			contains: []string{
				`aria-valuenow="100"`,
				`width: 100%`,
			},
		},
		{
			name:     "small progress",
			progress: progress.SmallComponent(75),
			contains: []string{
				`h-2`,
				`width: 75%`,
			},
		},
		{
			name:     "large progress",
			progress: progress.Large(25),
			contains: []string{
				`h-6`,
				`width: 25%`,
			},
		},
		{
			name: "progress with custom max",
			progress: progress.New(progress.Props{
				Value: 25,
				Max:   50,
			}),
			contains: []string{
				`aria-valuemax="50"`,
				`aria-valuenow="25"`,
				`aria-label="Progress: 50%"`, // 25/50 = 50%
				`width: 50%`,
			},
		},
		{
			name:     "progress with label",
			progress: progress.WithLabel(60, "Upload progress"),
			contains: []string{
				`<span>Upload progress</span>`,
				`<span class="text-muted-foreground">60%</span>`,
				`aria-valuenow="60"`,
			},
		},
		{
			name:     "indeterminate progress",
			progress: progress.Indeterminate(),
			contains: []string{
				`aria-label="Loading..."`,
				`animate-pulse`,
				`width: 50%`,
			},
		},
		{
			name:     "striped progress",
			progress: progress.Striped(40),
			contains: []string{
				`background-image: linear-gradient`,
				`45deg`,
				`background-size: 1rem 1rem`,
				`width: 40%`,
			},
		},
		{
			name: "multi-segment progress",
			progress: progress.Multi([]progress.Segment{
				{Value: 30, Color: "bg-blue-500"},
				{Value: 20, Color: "bg-green-500"},
				{Value: 10, Color: "bg-yellow-500"},
			}),
			contains: []string{
				`aria-valuenow="60"`, // Total
				`aria-label="Total progress: 60%"`,
				`width: 30%`,
				`width: 20%`,
				`width: 10%`,
				`bg-blue-500`,
				`bg-green-500`,
				`bg-yellow-500`,
			},
		},
		{
			name:     "circular progress default",
			progress: progress.Circular(70, "default"),
			contains: []string{
				`<svg width="48" height="48"`,
				`viewBox="0 0 48 48"`,
				`aria-valuenow="70"`,
				`<circle cx="24" cy="24" r="22"`, // Background circle
				`stroke-dasharray`,
				`stroke-dashoffset`,
				`<span class="text-xs font-semibold">70%</span>`,
			},
		},
		{
			name:     "circular progress small",
			progress: progress.Circular(50, "sm"),
			contains: []string{
				`<svg width="32" height="32"`,
				`stroke-width="3"`,
				`<span class="text-xs font-semibold">50%</span>`,
			},
		},
		{
			name: "progress with custom class",
			progress: progress.New(progress.Props{
				Value: 80,
				Class: "w-1/2",
			}),
			contains: []string{
				`w-1/2`,
			},
		},
		{
			name: "progress exceeding max",
			progress: progress.New(progress.Props{
				Value: 150,
				Max:   100,
			}),
			contains: []string{
				`aria-valuenow="100"`, // Clamped to max
				`width: 100%`,
			},
		},
		{
			name: "negative progress",
			progress: progress.New(progress.Props{
				Value: -10,
			}),
			contains: []string{
				`aria-valuenow="0"`, // Clamped to 0
				`width: 0%`,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := test.progress.Render(&buf)
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