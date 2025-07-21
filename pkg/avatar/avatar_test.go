package avatar_test

import (
	"strings"
	"testing"
	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/pkg/avatar"
)

func TestAvatar(t *testing.T) {
	tests := []struct {
		name     string
		avatar   g.Node
		contains []string
	}{
		{
			name: "default avatar with image",
			avatar: avatar.Default(
				avatar.Image(avatar.ImageProps{
					Src: "https://example.com/avatar.jpg",
					Alt: "User Avatar",
				}),
			),
			contains: []string{
				`class="relative flex shrink-0 overflow-hidden rounded-full h-10 w-10"`,
				`<img src="https://example.com/avatar.jpg"`,
				`alt="User Avatar"`,
				`class="aspect-square h-full w-full"`,
			},
		},
		{
			name: "avatar with fallback",
			avatar: avatar.Default(
				avatar.Fallback(Span(g.Text("JD"))),
			),
			contains: []string{
				`class="flex h-full w-full items-center justify-center rounded-full bg-muted"`,
				`<span>JD</span>`,
			},
		},
		{
			name:   "small avatar",
			avatar: avatar.Small(avatar.Fallback(g.Text("A"))),
			contains: []string{
				`h-8 w-8`,
			},
		},
		{
			name:   "large avatar",
			avatar: avatar.Large(avatar.Fallback(g.Text("B"))),
			contains: []string{
				`h-12 w-12`,
			},
		},
		{
			name: "avatar with custom size",
			avatar: avatar.New(
				avatar.Props{Size: "h-16 w-16"},
				avatar.Fallback(g.Text("C")),
			),
			contains: []string{
				`h-16 w-16`,
			},
		},
		{
			name: "avatar with custom class",
			avatar: avatar.New(
				avatar.Props{Class: "border-2 border-primary"},
				avatar.Fallback(g.Text("D")),
			),
			contains: []string{
				`border-2 border-primary`,
			},
		},
		{
			name:   "avatar with initials",
			avatar: avatar.WithInitials("AB"),
			contains: []string{
				`<span class="text-sm font-medium">AB</span>`,
			},
		},
		{
			name: "avatar with icon",
			avatar: avatar.WithIcon(
				g.El("svg", 
					g.Attr("class", "h-6 w-6"),
					g.Attr("viewBox", "0 0 24 24"),
				),
			),
			contains: []string{
				`<svg class="h-6 w-6" viewBox="0 0 24 24"`,
			},
		},
		{
			name: "avatar with image and fallback",
			avatar: avatar.WithImage(
				"/avatar.png",
				"Profile",
				g.Text("PF"),
			),
			contains: []string{
				`src="/avatar.png"`,
				`alt="Profile"`,
				`>PF<`,
			},
		},
		{
			name: "avatar group",
			avatar: avatar.Group(
				avatar.GroupItem(avatar.WithInitials("A"), 0),
				avatar.GroupItem(avatar.WithInitials("B"), 1),
				avatar.GroupItem(avatar.WithInitials("C"), 2),
			),
			contains: []string{
				`class="flex -space-x-4"`,
				`z-index: 10`,
				`z-index: 9`,
				`z-index: 8`,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.avatar.String()
			for _, expected := range test.contains {
				if !strings.Contains(result, expected) {
					t.Errorf("expected result to contain %q, but it didn't.\nGot: %s", expected, result)
				}
			}
		})
	}
}