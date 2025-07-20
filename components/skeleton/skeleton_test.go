package skeleton_test

import (
	"strings"
	"testing"
	g "maragu.dev/gomponents"
	"github.com/rizome-dev/shadcn-gomponents/components/skeleton"
)

func TestSkeleton(t *testing.T) {
	tests := []struct {
		name     string
		skeleton g.Node
		contains []string
	}{
		{
			name:     "default skeleton",
			skeleton: skeleton.Default(),
			contains: []string{
				`class="animate-pulse rounded-md bg-muted"`,
				`aria-hidden="true"`,
			},
		},
		{
			name:     "skeleton with custom class",
			skeleton: skeleton.WithClass("h-20 w-20"),
			contains: []string{
				`animate-pulse rounded-md bg-muted h-20 w-20`,
			},
		},
		{
			name:     "text skeleton",
			skeleton: skeleton.Text(),
			contains: []string{
				`h-4 w-full`,
			},
		},
		{
			name:     "short text skeleton",
			skeleton: skeleton.TextShort(),
			contains: []string{
				`h-4 w-3/4`,
			},
		},
		{
			name:     "text lines skeleton",
			skeleton: skeleton.TextLines(3),
			contains: []string{
				`class="space-y-2"`,
				`h-4 w-full`, // First two lines
				`h-4 w-3/4`,  // Last line is shorter
			},
		},
		{
			name:     "avatar skeleton default",
			skeleton: skeleton.Avatar("default"),
			contains: []string{
				`h-10 w-10 rounded-full`,
			},
		},
		{
			name:     "avatar skeleton small",
			skeleton: skeleton.Avatar("sm"),
			contains: []string{
				`h-8 w-8 rounded-full`,
			},
		},
		{
			name:     "avatar skeleton large",
			skeleton: skeleton.Avatar("lg"),
			contains: []string{
				`h-12 w-12 rounded-full`,
			},
		},
		{
			name:     "button skeleton default",
			skeleton: skeleton.Button("default"),
			contains: []string{
				`h-10 w-24`,
			},
		},
		{
			name:     "button skeleton icon",
			skeleton: skeleton.Button("icon"),
			contains: []string{
				`h-10 w-10`,
			},
		},
		{
			name:     "input skeleton",
			skeleton: skeleton.Input(),
			contains: []string{
				`h-10 w-full`,
			},
		},
		{
			name:     "card skeleton",
			skeleton: skeleton.Card(),
			contains: []string{
				`h-[125px] w-full rounded-xl`,
			},
		},
		{
			name:     "image skeleton default",
			skeleton: skeleton.Image("video"),
			contains: []string{
				`w-full aspect-video`,
			},
		},
		{
			name:     "image skeleton square",
			skeleton: skeleton.Image("square"),
			contains: []string{
				`w-full aspect-square`,
			},
		},
		{
			name:     "table skeleton",
			skeleton: skeleton.Table(2),
			contains: []string{
				`class="rounded-md border"`,
				`class="flex gap-4 p-4 border-b"`,
				`h-4 w-24`, // Header columns
				`h-4 w-32`,
				`h-4 w-20`, // Data columns
				`h-4 w-40`,
			},
		},
		{
			name:     "profile card skeleton",
			skeleton: skeleton.ProfileCard(),
			contains: []string{
				`class="flex items-center space-x-4"`,
				`h-10 w-10 rounded-full`, // Avatar
				`h-4 w-[200px]`,          // Name
				`h-4 w-[150px]`,          // Description
			},
		},
		{
			name:     "post card skeleton",
			skeleton: skeleton.PostCard(),
			contains: []string{
				`class="space-y-3"`,
				`aspect-video`, // Image
				`h-4 w-3/4`,    // Title
			},
		},
		{
			name:     "list skeleton",
			skeleton: skeleton.List(3),
			contains: []string{
				`class="flex items-center space-x-3"`,
				`h-4 w-4 rounded-full`, // Bullet
				`h-4 flex-1`,           // Text
			},
		},
		{
			name:     "form skeleton",
			skeleton: skeleton.Form(),
			contains: []string{
				`class="space-y-6"`,
				`h-4 w-24`,  // Label 1
				`h-4 w-32`,  // Label 2
				`h-10 w-24`, // Buttons
			},
		},
		{
			name:     "grid skeleton 2 cols",
			skeleton: skeleton.Grid(2, 4),
			contains: []string{
				`class="grid gap-4 grid-cols-2"`,
				`h-[125px] w-full rounded-xl`, // Card
			},
		},
		{
			name: "complex skeleton composition",
			skeleton: g.Group{
				skeleton.Avatar("default"),
				skeleton.TextLines(2),
				skeleton.Button("default"),
			},
			contains: []string{
				`rounded-full`,
				`h-4 w-full`,
				`h-10 w-24`,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.skeleton.String()
			for _, expected := range test.contains {
				if !strings.Contains(result, expected) {
					t.Errorf("expected result to contain %q, but it didn't.\nGot: %s", expected, result)
				}
			}
		})
	}
}