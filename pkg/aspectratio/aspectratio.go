package aspectratio

import (
	"fmt"
	
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// Props defines the properties for the AspectRatio component
type Props struct {
	Ratio float64 // The aspect ratio (e.g., 16/9, 4/3, 1/1)
	Class string  // Additional custom classes
}

// New creates a new AspectRatio component
// The AspectRatio component maintains a specific aspect ratio for its content
func New(props Props, children ...g.Node) g.Node {
	// Default ratio is 1:1 if not specified
	ratio := props.Ratio
	if ratio == 0 {
		ratio = 1
	}
	
	// Calculate padding percentage for the aspect ratio
	// paddingBottom = (1 / ratio) * 100
	paddingBottom := (1 / ratio) * 100
	
	classes := lib.CN("relative overflow-hidden", props.Class)
	
	return html.Div(
		html.Class(classes),
		dataAttr("slot", "aspect-ratio"),
		dataAttr("aspect-ratio", fmt.Sprintf("%.2f", ratio)),
		html.Style(fmt.Sprintf("padding-bottom: %.4f%%;", paddingBottom)),
		
		// Inner container for content
		html.Div(
			html.Class("absolute inset-0"),
			g.Group(children),
		),
	)
}

// Common aspect ratio helpers

// Square creates an aspect ratio container with 1:1 ratio
func Square(children ...g.Node) g.Node {
	return New(Props{Ratio: 1}, children...)
}

// Video16x9 creates an aspect ratio container with 16:9 ratio (standard video)
func Video16x9(children ...g.Node) g.Node {
	return New(Props{Ratio: 16.0 / 9.0}, children...)
}

// Portrait creates an aspect ratio container with 4:5 ratio
func Portrait(children ...g.Node) g.Node {
	return New(Props{Ratio: 4.0 / 5.0}, children...)
}

// Landscape creates an aspect ratio container with 3:2 ratio
func Landscape(children ...g.Node) g.Node {
	return New(Props{Ratio: 3.0 / 2.0}, children...)
}

// Cinema creates an aspect ratio container with 2.39:1 ratio (cinematic)
func Cinema(children ...g.Node) g.Node {
	return New(Props{Ratio: 2.39}, children...)
}

// WithClass creates an aspect ratio with a specific class
func WithClass(ratio float64, class string, children ...g.Node) g.Node {
	return New(Props{Ratio: ratio, Class: class}, children...)
}

// dataAttr creates a data attribute
func dataAttr(name, value string) g.Node {
	return g.Attr("data-" + name, value)
}