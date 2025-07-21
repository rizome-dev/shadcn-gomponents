package avatar

import (
	"fmt"
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// Props defines the properties for the Avatar component
type Props struct {
	Size  string // "sm" | "default" | "lg" or custom size class
	Class string // Additional custom classes
}

// ImageProps defines properties for the AvatarImage
type ImageProps struct {
	Src   string
	Alt   string
	Class string
}

// New creates a new Avatar container
func New(props Props, children ...g.Node) g.Node {
	// Default size
	size := "h-10 w-10"
	
	// Handle predefined sizes
	switch props.Size {
	case "sm":
		size = "h-8 w-8"
	case "lg":
		size = "h-12 w-12"
	case "":
		// Keep default
	default:
		// Allow custom size classes
		size = props.Size
	}

	classes := lib.CN(
		"relative flex shrink-0 overflow-hidden rounded-full",
		size,
		props.Class,
	)

	return html.Div(
		html.Class(classes),
		g.Group(children),
	)
}

// Default creates an avatar with default size
func Default(children ...g.Node) g.Node {
	return New(Props{}, children...)
}

// Small creates a small avatar
func SmallComponent(children ...g.Node) g.Node {
	return New(Props{Size: "sm"}, children...)
}

// Large creates a large avatar
func Large(children ...g.Node) g.Node {
	return New(Props{Size: "lg"}, children...)
}

// Image creates an AvatarImage component
func Image(props ImageProps) g.Node {
	classes := lib.CN(
		"aspect-square h-full w-full",
		props.Class,
	)

	attrs := []g.Node{
		html.Src(props.Src),
		html.Class(classes),
	}
	
	if props.Alt != "" {
		attrs = append(attrs, html.Alt(props.Alt))
	}

	return html.Img(attrs...)
}

// Fallback creates an AvatarFallback component
func Fallback(children ...g.Node) g.Node {
	return html.Div(
		html.Class("flex h-full w-full items-center justify-center rounded-full bg-muted"),
		g.Group(children),
	)
}

// WithImage creates an avatar with an image and fallback
// Note: In Go, we can't automatically detect image load failures,
// so both image and fallback will be rendered (fallback hidden by image when loaded)
func WithImage(src, alt string, fallback g.Node) g.Node {
	return Default(
		Image(ImageProps{Src: src, Alt: alt}),
		// In a real implementation, you might want JavaScript to handle
		// showing/hiding fallback based on image load status
		g.If(fallback != nil, Fallback(fallback)),
	)
}

// WithInitials creates an avatar with initials as fallback
func WithInitials(initials string) g.Node {
	return Default(
		Fallback(
			html.Span(html.Class("text-sm font-medium"), g.Text(initials)),
		),
	)
}

// WithIcon creates an avatar with an icon
func WithIcon(icon g.Node) g.Node {
	return Default(
		Fallback(icon),
	)
}

// Group creates a group of overlapping avatars
func Group(avatars ...g.Node) g.Node {
	return html.Div(
		html.Class("flex -space-x-4"),
		g.Group(avatars),
	)
}

// GroupItem wraps an avatar for use in a group with proper z-index
func GroupItem(avatar g.Node, index int) g.Node {
	return html.Div(
		html.Class("relative inline-block"),
		html.Style(fmt.Sprintf("z-index: %d", 10-index)),
		avatar,
	)
}