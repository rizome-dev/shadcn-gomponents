package icons

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

// ChevronRight creates a chevron-right icon
func ChevronRight(attrs ...g.Node) g.Node {
	defaultAttrs := []g.Node{
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("width", "24"),
		g.Attr("height", "24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
	}
	
	// Append custom attributes
	allAttrs := append(defaultAttrs, attrs...)
	
	return g.El("svg",
		append(allAttrs,
			g.El("path", g.Attr("d", "m9 18 6-6-6-6")),
		)...,
	)
}

// MoreHorizontal creates a more-horizontal (ellipsis) icon
func MoreHorizontal(attrs ...g.Node) g.Node {
	defaultAttrs := []g.Node{
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("width", "24"),
		g.Attr("height", "24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
	}
	
	// Append custom attributes
	allAttrs := append(defaultAttrs, attrs...)
	
	return g.El("svg",
		append(allAttrs,
			g.El("circle", g.Attr("cx", "12"), g.Attr("cy", "12"), g.Attr("r", "1")),
			g.El("circle", g.Attr("cx", "19"), g.Attr("cy", "12"), g.Attr("r", "1")),
			g.El("circle", g.Attr("cx", "5"), g.Attr("cy", "12"), g.Attr("r", "1")),
		)...,
	)
}

// Plus creates a plus icon
func Plus(attrs ...g.Node) g.Node {
	defaultAttrs := []g.Node{
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("width", "24"),
		g.Attr("height", "24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
	}
	
	// Append custom attributes
	allAttrs := append(defaultAttrs, attrs...)
	
	return g.El("svg",
		append(allAttrs,
			g.El("line", g.Attr("x1", "12"), g.Attr("y1", "5"), g.Attr("x2", "12"), g.Attr("y2", "19")),
			g.El("line", g.Attr("x1", "5"), g.Attr("y1", "12"), g.Attr("x2", "19"), g.Attr("y2", "12")),
		)...,
	)
}

// X creates an X (close) icon
func X(attrs ...g.Node) g.Node {
	defaultAttrs := []g.Node{
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("width", "24"),
		g.Attr("height", "24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
	}
	
	// Append custom attributes
	allAttrs := append(defaultAttrs, attrs...)
	
	return g.El("svg",
		append(allAttrs,
			g.El("line", g.Attr("x1", "18"), g.Attr("y1", "6"), g.Attr("x2", "6"), g.Attr("y2", "18")),
			g.El("line", g.Attr("x1", "6"), g.Attr("y1", "6"), g.Attr("x2", "18"), g.Attr("y2", "18")),
		)...,
	)
}

// MenuIcon creates a menu (hamburger) icon
func MenuIcon(attrs ...g.Node) g.Node {
	defaultAttrs := []g.Node{
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("width", "24"),
		g.Attr("height", "24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
	}
	
	// Append custom attributes
	allAttrs := append(defaultAttrs, attrs...)
	
	return g.El("svg",
		append(allAttrs,
			g.El("line", g.Attr("x1", "3"), g.Attr("y1", "12"), g.Attr("x2", "21"), g.Attr("y2", "12")),
			g.El("line", g.Attr("x1", "3"), g.Attr("y1", "6"), g.Attr("x2", "21"), g.Attr("y2", "6")),
			g.El("line", g.Attr("x1", "3"), g.Attr("y1", "18"), g.Attr("x2", "21"), g.Attr("y2", "18")),
		)...,
	)
}

// Check creates a check icon
func Check(attrs ...g.Node) g.Node {
	defaultAttrs := []g.Node{
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("width", "24"),
		g.Attr("height", "24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
	}
	
	// Append custom attributes
	allAttrs := append(defaultAttrs, attrs...)
	
	return g.El("svg",
		append(allAttrs,
			g.El("polyline", g.Attr("points", "20 6 9 17 4 12")),
		)...,
	)
}

// ChevronDown creates a chevron-down icon
func ChevronDown(attrs ...g.Node) g.Node {
	defaultAttrs := []g.Node{
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("width", "24"),
		g.Attr("height", "24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
	}
	
	// Append custom attributes
	allAttrs := append(defaultAttrs, attrs...)
	
	return g.El("svg",
		append(allAttrs,
			g.El("path", g.Attr("d", "m6 9 6 6 6-6")),
		)...,
	)
}

// ChevronUp creates a chevron-up icon
func ChevronUp(attrs ...g.Node) g.Node {
	defaultAttrs := []g.Node{
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("width", "24"),
		g.Attr("height", "24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
	}
	
	// Append custom attributes
	allAttrs := append(defaultAttrs, attrs...)
	
	return g.El("svg",
		append(allAttrs,
			g.El("path", g.Attr("d", "m18 15-6-6-6 6")),
		)...,
	)
}

// ChevronLeft creates a chevron-left icon
func ChevronLeft(attrs ...g.Node) g.Node {
	defaultAttrs := []g.Node{
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("width", "24"),
		g.Attr("height", "24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
	}
	
	// Append custom attributes
	allAttrs := append(defaultAttrs, attrs...)
	
	return g.El("svg",
		append(allAttrs,
			g.El("path", g.Attr("d", "m15 18-6-6 6-6")),
		)...,
	)
}

// ArrowRight creates an arrow-right icon
func ArrowRight(attrs ...g.Node) g.Node {
	defaultAttrs := []g.Node{
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("width", "24"),
		g.Attr("height", "24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
	}
	
	// Append custom attributes
	allAttrs := append(defaultAttrs, attrs...)
	
	return g.El("svg",
		append(allAttrs,
			g.El("line", g.Attr("x1", "5"), g.Attr("y1", "12"), g.Attr("x2", "19"), g.Attr("y2", "12")),
			g.El("polyline", g.Attr("points", "12 5 19 12 12 19")),
		)...,
	)
}

// ArrowLeft creates an arrow-left icon
func ArrowLeft(attrs ...g.Node) g.Node {
	defaultAttrs := []g.Node{
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("width", "24"),
		g.Attr("height", "24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
	}
	
	// Append custom attributes
	allAttrs := append(defaultAttrs, attrs...)
	
	return g.El("svg",
		append(allAttrs,
			g.El("line", g.Attr("x1", "19"), g.Attr("y1", "12"), g.Attr("x2", "5"), g.Attr("y2", "12")),
			g.El("polyline", g.Attr("points", "12 19 5 12 12 5")),
		)...,
	)
}

// CircleIcon creates a circle icon
func CircleIcon(attrs ...g.Node) g.Node {
	defaultAttrs := []g.Node{
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("width", "24"),
		g.Attr("height", "24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
	}
	
	// Append custom attributes
	allAttrs := append(defaultAttrs, attrs...)
	
	return g.El("svg",
		append(allAttrs,
			g.El("circle", g.Attr("cx", "12"), g.Attr("cy", "12"), g.Attr("r", "10")),
		)...,
	)
}

// Dot creates a dot icon (filled circle)
func Dot(attrs ...g.Node) g.Node {
	defaultAttrs := []g.Node{
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("width", "24"),
		g.Attr("height", "24"),
		g.Attr("fill", "currentColor"),
		g.Attr("stroke", "none"),
	}
	
	// Append custom attributes
	allAttrs := append(defaultAttrs, attrs...)
	
	return g.El("svg",
		append(allAttrs,
			g.El("circle", g.Attr("cx", "12"), g.Attr("cy", "12"), g.Attr("r", "3")),
		)...,
	)
}

// Search creates a search icon
func Search(attrs ...g.Node) g.Node {
	defaultAttrs := []g.Node{
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("width", "24"),
		g.Attr("height", "24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
	}
	
	// Append custom attributes
	allAttrs := append(defaultAttrs, attrs...)
	
	return g.El("svg",
		append(allAttrs,
			g.El("circle", g.Attr("cx", "11"), g.Attr("cy", "11"), g.Attr("r", "8")),
			g.El("path", g.Attr("d", "m21 21-4.35-4.35")),
		)...,
	)
}

// Loader creates a loader/spinner icon
func Loader(attrs ...g.Node) g.Node {
	defaultAttrs := []g.Node{
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("width", "24"),
		g.Attr("height", "24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
		html.Class("animate-spin"),
	}
	
	// Append custom attributes
	allAttrs := append(defaultAttrs, attrs...)
	
	return g.El("svg",
		append(allAttrs,
			g.El("line", g.Attr("x1", "12"), g.Attr("y1", "2"), g.Attr("x2", "12"), g.Attr("y2", "6")),
			g.El("line", g.Attr("x1", "12"), g.Attr("y1", "18"), g.Attr("x2", "12"), g.Attr("y2", "22")),
			g.El("line", g.Attr("x1", "4.93"), g.Attr("y1", "4.93"), g.Attr("x2", "7.76"), g.Attr("y2", "7.76")),
			g.El("line", g.Attr("x1", "16.24"), g.Attr("y1", "16.24"), g.Attr("x2", "19.07"), g.Attr("y2", "19.07")),
			g.El("line", g.Attr("x1", "2"), g.Attr("y1", "12"), g.Attr("x2", "6"), g.Attr("y2", "12")),
			g.El("line", g.Attr("x1", "18"), g.Attr("y1", "12"), g.Attr("x2", "22"), g.Attr("y2", "12")),
			g.El("line", g.Attr("x1", "4.93"), g.Attr("y1", "19.07"), g.Attr("x2", "7.76"), g.Attr("y2", "16.24")),
			g.El("line", g.Attr("x1", "16.24"), g.Attr("y1", "7.76"), g.Attr("x2", "19.07"), g.Attr("y2", "4.93")),
		)...,
	)
}

// ChevronsUpDown creates a chevrons-up-down icon
func ChevronsUpDown(attrs ...g.Node) g.Node {
	defaultAttrs := []g.Node{
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("width", "24"),
		g.Attr("height", "24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
	}
	
	// Append custom attributes
	allAttrs := append(defaultAttrs, attrs...)
	
	return g.El("svg",
		append(allAttrs,
			g.El("path", g.Attr("d", "m7 15 5 5 5-5")),
			g.El("path", g.Attr("d", "m7 9 5-5 5 5")),
		)...,
	)
}

// User creates an SVG user icon
func User(attrs ...g.Node) g.Node {
	defaultAttrs := []g.Node{
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("width", "24"),
		g.Attr("height", "24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
	}
	allAttrs := append(defaultAttrs, attrs...)
	return g.El("svg",
		append(allAttrs,
			g.El("path", g.Attr("d", "M19 21v-2a4 4 0 0 0-4-4H9a4 4 0 0 0-4 4v2")),
			g.El("circle", g.Attr("cx", "12"), g.Attr("cy", "7"), g.Attr("r", "4")),
		)...,
	)
}

// CreditCard creates an SVG credit card icon
func CreditCard(attrs ...g.Node) g.Node {
	defaultAttrs := []g.Node{
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("width", "24"),
		g.Attr("height", "24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
	}
	allAttrs := append(defaultAttrs, attrs...)
	return g.El("svg",
		append(allAttrs,
			g.El("rect", g.Attr("width", "20"), g.Attr("height", "14"), g.Attr("x", "2"), g.Attr("y", "5"), g.Attr("rx", "2")),
			g.El("line", g.Attr("x1", "2"), g.Attr("x2", "22"), g.Attr("y1", "10"), g.Attr("y2", "10")),
		)...,
	)
}

// Settings creates an SVG settings icon
func Settings(attrs ...g.Node) g.Node {
	defaultAttrs := []g.Node{
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("width", "24"),
		g.Attr("height", "24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
	}
	allAttrs := append(defaultAttrs, attrs...)
	return g.El("svg",
		append(allAttrs,
			g.El("path", g.Attr("d", "M12.22 2h-.44a2 2 0 0 0-2 2v.18a2 2 0 0 1-1 1.73l-.43.25a2 2 0 0 1-2 0l-.15-.08a2 2 0 0 0-2.73.73l-.22.38a2 2 0 0 0 .73 2.73l.15.1a2 2 0 0 1 1 1.72v.51a2 2 0 0 1-1 1.74l-.15.09a2 2 0 0 0-.73 2.73l.22.38a2 2 0 0 0 2.73.73l.15-.08a2 2 0 0 1 2 0l.43.25a2 2 0 0 1 1 1.73V20a2 2 0 0 0 2 2h.44a2 2 0 0 0 2-2v-.18a2 2 0 0 1 1-1.73l.43-.25a2 2 0 0 1 2 0l.15.08a2 2 0 0 0 2.73-.73l.22-.39a2 2 0 0 0-.73-2.73l-.15-.08a2 2 0 0 1-1-1.74v-.5a2 2 0 0 1 1-1.74l.15-.09a2 2 0 0 0 .73-2.73l-.22-.38a2 2 0 0 0-2.73-.73l-.15.08a2 2 0 0 1-2 0l-.43-.25a2 2 0 0 1-1-1.73V4a2 2 0 0 0-2-2z")),
			g.El("circle", g.Attr("cx", "12"), g.Attr("cy", "12"), g.Attr("r", "3")),
		)...,
	)
}

// Cloud creates an SVG cloud icon
func Cloud(attrs ...g.Node) g.Node {
	defaultAttrs := []g.Node{
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("width", "24"),
		g.Attr("height", "24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
	}
	allAttrs := append(defaultAttrs, attrs...)
	return g.El("svg",
		append(allAttrs,
			g.El("path", g.Attr("d", "M17.5 19H9a7 7 0 1 1 6.71-9h1.79a4.5 4.5 0 1 1 0 9Z")),
		)...,
	)
}

// LogOut creates an SVG log out icon
func LogOut(attrs ...g.Node) g.Node {
	defaultAttrs := []g.Node{
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("width", "24"),
		g.Attr("height", "24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
	}
	allAttrs := append(defaultAttrs, attrs...)
	return g.El("svg",
		append(allAttrs,
			g.El("path", g.Attr("d", "M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4")),
			g.El("polyline", g.Attr("points", "16 17 21 12 16 7")),
			g.El("line", g.Attr("x1", "21"), g.Attr("x2", "9"), g.Attr("y1", "12"), g.Attr("y2", "12")),
		)...,
	)
}

// Users creates an SVG users icon
func Users(attrs ...g.Node) g.Node {
	defaultAttrs := []g.Node{
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("width", "24"),
		g.Attr("height", "24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
	}
	allAttrs := append(defaultAttrs, attrs...)
	return g.El("svg",
		append(allAttrs,
			g.El("path", g.Attr("d", "M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2")),
			g.El("circle", g.Attr("cx", "9"), g.Attr("cy", "7"), g.Attr("r", "4")),
			g.El("path", g.Attr("d", "M22 21v-2a4 4 0 0 0-3-3.87")),
			g.El("path", g.Attr("d", "M16 3.13a4 4 0 0 1 0 7.75")),
		)...,
	)
}

// UserPlus creates an SVG user plus icon
func UserPlus(attrs ...g.Node) g.Node {
	defaultAttrs := []g.Node{
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("width", "24"),
		g.Attr("height", "24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
	}
	allAttrs := append(defaultAttrs, attrs...)
	return g.El("svg",
		append(allAttrs,
			g.El("path", g.Attr("d", "M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2")),
			g.El("circle", g.Attr("cx", "9"), g.Attr("cy", "7"), g.Attr("r", "4")),
			g.El("line", g.Attr("x1", "19"), g.Attr("x2", "19"), g.Attr("y1", "8"), g.Attr("y2", "14")),
			g.El("line", g.Attr("x1", "22"), g.Attr("x2", "16"), g.Attr("y1", "11"), g.Attr("y2", "11")),
		)...,
	)
}

// Calendar creates an SVG calendar icon
func Calendar(attrs ...g.Node) g.Node {
	defaultAttrs := []g.Node{
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("width", "24"),
		g.Attr("height", "24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
	}
	allAttrs := append(defaultAttrs, attrs...)
	return g.El("svg",
		append(allAttrs,
			g.El("rect", g.Attr("width", "18"), g.Attr("height", "18"), g.Attr("x", "3"), g.Attr("y", "4"), g.Attr("rx", "2"), g.Attr("ry", "2")),
			g.El("line", g.Attr("x1", "16"), g.Attr("x2", "16"), g.Attr("y1", "2"), g.Attr("y2", "6")),
			g.El("line", g.Attr("x1", "8"), g.Attr("x2", "8"), g.Attr("y1", "2"), g.Attr("y2", "6")),
			g.El("line", g.Attr("x1", "3"), g.Attr("x2", "21"), g.Attr("y1", "10"), g.Attr("y2", "10")),
		)...,
	)
}

// Menu creates a menu (hamburger) icon - alias for MenuIcon
func Menu(attrs ...g.Node) g.Node {
	return MenuIcon(attrs...)
}

// Home creates a home icon
func Home(attrs ...g.Node) g.Node {
	defaultAttrs := []g.Node{
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("width", "24"),
		g.Attr("height", "24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
	}
	allAttrs := append(defaultAttrs, attrs...)
	return g.El("svg",
		append(allAttrs,
			g.El("path", g.Attr("d", "m3 9 9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z")),
			g.El("polyline", g.Attr("points", "9 22 9 12 15 12 15 22")),
		)...,
	)
}

// Package creates a package icon
func Package(attrs ...g.Node) g.Node {
	defaultAttrs := []g.Node{
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("width", "24"),
		g.Attr("height", "24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
	}
	allAttrs := append(defaultAttrs, attrs...)
	return g.El("svg",
		append(allAttrs,
			g.El("path", g.Attr("d", "M16.5 9.4 7.55 4.24")),
			g.El("path", g.Attr("d", "M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z")),
			g.El("polyline", g.Attr("points", "3.29 7 12 12 20.71 7")),
			g.El("line", g.Attr("x1", "12"), g.Attr("x2", "12"), g.Attr("y1", "22"), g.Attr("y2", "12")),
		)...,
	)
}

// MoreVertical creates a more-vertical (ellipsis vertical) icon
func MoreVertical(attrs ...g.Node) g.Node {
	defaultAttrs := []g.Node{
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("width", "24"),
		g.Attr("height", "24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
	}
	
	// Append custom attributes
	allAttrs := append(defaultAttrs, attrs...)
	
	return g.El("svg",
		append(allAttrs,
			g.El("circle", g.Attr("cx", "12"), g.Attr("cy", "12"), g.Attr("r", "1")),
			g.El("circle", g.Attr("cx", "12"), g.Attr("cy", "5"), g.Attr("r", "1")),
			g.El("circle", g.Attr("cx", "12"), g.Attr("cy", "19"), g.Attr("r", "1")),
		)...,
	)
}

// Edit creates an edit icon
func Edit(attrs ...g.Node) g.Node {
	defaultAttrs := []g.Node{
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("width", "24"),
		g.Attr("height", "24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
	}
	allAttrs := append(defaultAttrs, attrs...)
	return g.El("svg",
		append(allAttrs,
			g.El("path", g.Attr("d", "M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7")),
			g.El("path", g.Attr("d", "M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z")),
		)...,
	)
}

// Copy creates a copy icon
func Copy(attrs ...g.Node) g.Node {
	defaultAttrs := []g.Node{
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("width", "24"),
		g.Attr("height", "24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
	}
	allAttrs := append(defaultAttrs, attrs...)
	return g.El("svg",
		append(allAttrs,
			g.El("rect", g.Attr("width", "14"), g.Attr("height", "14"), g.Attr("x", "8"), g.Attr("y", "8"), g.Attr("rx", "2"), g.Attr("ry", "2")),
			g.El("path", g.Attr("d", "M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2")),
		)...,
	)
}

// Archive creates an archive icon
func Archive(attrs ...g.Node) g.Node {
	defaultAttrs := []g.Node{
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("width", "24"),
		g.Attr("height", "24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
	}
	allAttrs := append(defaultAttrs, attrs...)
	return g.El("svg",
		append(allAttrs,
			g.El("rect", g.Attr("width", "20"), g.Attr("height", "5"), g.Attr("x", "2"), g.Attr("y", "3"), g.Attr("rx", "1")),
			g.El("path", g.Attr("d", "M4 8v11a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8")),
			g.El("line", g.Attr("x1", "10"), g.Attr("x2", "14"), g.Attr("y1", "12"), g.Attr("y2", "12")),
		)...,
	)
}

// Trash creates a trash icon
func Trash(attrs ...g.Node) g.Node {
	defaultAttrs := []g.Node{
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("width", "24"),
		g.Attr("height", "24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
	}
	allAttrs := append(defaultAttrs, attrs...)
	return g.El("svg",
		append(allAttrs,
			g.El("path", g.Attr("d", "M3 6h18")),
			g.El("path", g.Attr("d", "M19 6v14c0 1-1 2-2 2H7c-1 0-2-1-2-2V6")),
			g.El("path", g.Attr("d", "M8 6V4c0-1 1-2 2-2h4c1 0 2 1 2 2v2")),
		)...,
	)
}

// Cut creates a scissors icon
func Cut(attrs ...g.Node) g.Node {
	defaultAttrs := []g.Node{
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("width", "24"),
		g.Attr("height", "24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
	}
	
	// Append custom attributes
	allAttrs := append(defaultAttrs, attrs...)
	
	return g.El("svg",
		append(allAttrs,
			g.El("circle", g.Attr("cx", "6"), g.Attr("cy", "6"), g.Attr("r", "3")),
			g.El("circle", g.Attr("cx", "6"), g.Attr("cy", "18"), g.Attr("r", "3")),
			g.El("line", g.Attr("x1", "20"), g.Attr("y1", "4"), g.Attr("x2", "8.12"), g.Attr("y2", "15.88")),
			g.El("line", g.Attr("x1", "14.47"), g.Attr("y1", "14.48"), g.Attr("x2", "20"), g.Attr("y2", "20")),
			g.El("line", g.Attr("x1", "8.12"), g.Attr("y1", "8.12"), g.Attr("x2", "12"), g.Attr("y2", "12")),
		)...,
	)
}

// Paste creates a clipboard paste icon
func Paste(attrs ...g.Node) g.Node {
	defaultAttrs := []g.Node{
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("width", "24"),
		g.Attr("height", "24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
	}
	
	// Append custom attributes
	allAttrs := append(defaultAttrs, attrs...)
	
	return g.El("svg",
		append(allAttrs,
			g.El("path", g.Attr("d", "M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2")),
			g.El("rect", g.Attr("x", "8"), g.Attr("y", "2"), g.Attr("width", "8"), g.Attr("height", "4"), g.Attr("rx", "1"), g.Attr("ry", "1")),
		)...,
	)
}

// SelectAll creates a select all icon
func SelectAll(attrs ...g.Node) g.Node {
	defaultAttrs := []g.Node{
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("width", "24"),
		g.Attr("height", "24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
	}
	
	// Append custom attributes
	allAttrs := append(defaultAttrs, attrs...)
	
	return g.El("svg",
		append(allAttrs,
			g.El("rect", g.Attr("x", "3"), g.Attr("y", "3"), g.Attr("width", "18"), g.Attr("height", "18"), g.Attr("rx", "2"), g.Attr("ry", "2")),
			g.El("polyline", g.Attr("points", "9 11 12 14 22 4")),
		)...,
	)
}

// Undo creates an undo icon (curved arrow left)
func Undo(attrs ...g.Node) g.Node {
	defaultAttrs := []g.Node{
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("width", "24"),
		g.Attr("height", "24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
	}
	
	// Append custom attributes
	allAttrs := append(defaultAttrs, attrs...)
	
	return g.El("svg",
		append(allAttrs,
			g.El("path", g.Attr("d", "M3 7v6h6")),
			g.El("path", g.Attr("d", "M21 17a9 9 0 00-9-9 9 9 0 00-6 2.3L3 13")),
		)...,
	)
}

// Redo creates a redo icon (curved arrow right)
func Redo(attrs ...g.Node) g.Node {
	defaultAttrs := []g.Node{
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("width", "24"),
		g.Attr("height", "24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
	}
	
	// Append custom attributes
	allAttrs := append(defaultAttrs, attrs...)
	
	return g.El("svg",
		append(allAttrs,
			g.El("path", g.Attr("d", "M21 7v6h-6")),
			g.El("path", g.Attr("d", "M3 17a9 9 0 019-9 9 9 0 016 2.3l3 2.7")),
		)...,
	)
}