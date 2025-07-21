package carousel

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ExampleBasic demonstrates a basic carousel
func ExampleBasic() g.Node {
	slides := []g.Node{
		Item(
			h.Div(
				h.Class("h-64 bg-gradient-to-r from-blue-500 to-purple-500 rounded-lg flex items-center justify-center"),
				h.H3(h.Class("text-3xl font-bold text-white"), g.Text("Slide 1")),
			),
		),
		Item(
			h.Div(
				h.Class("h-64 bg-gradient-to-r from-green-500 to-teal-500 rounded-lg flex items-center justify-center"),
				h.H3(h.Class("text-3xl font-bold text-white"), g.Text("Slide 2")),
			),
		),
		Item(
			h.Div(
				h.Class("h-64 bg-gradient-to-r from-orange-500 to-red-500 rounded-lg flex items-center justify-center"),
				h.H3(h.Class("text-3xl font-bold text-white"), g.Text("Slide 3")),
			),
		),
	}

	return h.Div(
		h.Class("w-full max-w-4xl mx-auto p-4"),
		h.H2(h.Class("text-2xl font-bold mb-4"), g.Text("Basic Carousel")),
		New(slides),
	)
}

// ExampleWithLoop demonstrates a carousel with loop enabled
func ExampleWithLoop() g.Node {
	slides := []g.Node{
		Item(
			h.Div(
				h.Class("p-6"),
				h.Img(
					h.Src("https://via.placeholder.com/600x400/4F46E5/ffffff?text=Image+1"),
					h.Alt("Slide 1"),
					h.Class("w-full rounded-lg"),
				),
			),
		),
		Item(
			h.Div(
				h.Class("p-6"),
				h.Img(
					h.Src("https://via.placeholder.com/600x400/7C3AED/ffffff?text=Image+2"),
					h.Alt("Slide 2"),
					h.Class("w-full rounded-lg"),
				),
			),
		),
		Item(
			h.Div(
				h.Class("p-6"),
				h.Img(
					h.Src("https://via.placeholder.com/600x400/EC4899/ffffff?text=Image+3"),
					h.Alt("Slide 3"),
					h.Class("w-full rounded-lg"),
				),
			),
		),
	}

	return h.Div(
		h.Class("w-full max-w-4xl mx-auto p-4"),
		h.H2(h.Class("text-2xl font-bold mb-4"), g.Text("Carousel with Loop")),
		New(slides, WithLoop(), WithClass("shadow-lg rounded-lg bg-gray-100")),
	)
}

// ExampleAutoPlay demonstrates an auto-playing carousel
func ExampleAutoPlay() g.Node {
	slides := []g.Node{
		Item(
			h.Div(
				h.Class("relative h-96"),
				h.Img(
					h.Src("https://via.placeholder.com/800x600/0EA5E9/ffffff?text=Auto+Slide+1"),
					h.Alt("Auto Slide 1"),
					h.Class("w-full h-full object-cover rounded-lg"),
				),
				h.Div(
					h.Class("absolute bottom-0 left-0 right-0 p-6 bg-gradient-to-t from-black/70 to-transparent"),
					h.H3(h.Class("text-2xl font-bold text-white"), g.Text("Automatic Slideshow")),
					h.P(h.Class("text-white/90"), g.Text("This carousel auto-advances every 5 seconds")),
				),
			),
		),
		Item(
			h.Div(
				h.Class("relative h-96"),
				h.Img(
					h.Src("https://via.placeholder.com/800x600/10B981/ffffff?text=Auto+Slide+2"),
					h.Alt("Auto Slide 2"),
					h.Class("w-full h-full object-cover rounded-lg"),
				),
				h.Div(
					h.Class("absolute bottom-0 left-0 right-0 p-6 bg-gradient-to-t from-black/70 to-transparent"),
					h.H3(h.Class("text-2xl font-bold text-white"), g.Text("Seamless Transitions")),
					h.P(h.Class("text-white/90"), g.Text("Smooth animations between slides")),
				),
			),
		),
		Item(
			h.Div(
				h.Class("relative h-96"),
				h.Img(
					h.Src("https://via.placeholder.com/800x600/F59E0B/ffffff?text=Auto+Slide+3"),
					h.Alt("Auto Slide 3"),
					h.Class("w-full h-full object-cover rounded-lg"),
				),
				h.Div(
					h.Class("absolute bottom-0 left-0 right-0 p-6 bg-gradient-to-t from-black/70 to-transparent"),
					h.H3(h.Class("text-2xl font-bold text-white"), g.Text("User Control")),
					h.P(h.Class("text-white/90"), g.Text("Users can still navigate manually")),
				),
			),
		),
	}

	return h.Div(
		h.Class("w-full max-w-4xl mx-auto p-4"),
		h.H2(h.Class("text-2xl font-bold mb-4"), g.Text("Auto-Play Carousel")),
		New(slides, WithAutoPlay(5000), WithLoop()),
	)
}

// ExampleVertical demonstrates a vertical carousel
func ExampleVertical() g.Node {
	slides := []g.Node{
		Item(
			h.Div(
				h.Class("p-6 bg-blue-50 rounded-lg"),
				h.H3(h.Class("text-xl font-semibold mb-2"), g.Text("Vertical Slide 1")),
				h.P(h.Class("text-gray-600"), g.Text("This carousel scrolls vertically instead of horizontally.")),
			),
		),
		Item(
			h.Div(
				h.Class("p-6 bg-green-50 rounded-lg"),
				h.H3(h.Class("text-xl font-semibold mb-2"), g.Text("Vertical Slide 2")),
				h.P(h.Class("text-gray-600"), g.Text("Perfect for testimonials or feature lists.")),
			),
		),
		Item(
			h.Div(
				h.Class("p-6 bg-purple-50 rounded-lg"),
				h.H3(h.Class("text-xl font-semibold mb-2"), g.Text("Vertical Slide 3")),
				h.P(h.Class("text-gray-600"), g.Text("Indicators and controls adapt to vertical orientation.")),
			),
		),
	}

	return h.Div(
		h.Class("w-full max-w-2xl mx-auto p-4"),
		h.H2(h.Class("text-2xl font-bold mb-4"), g.Text("Vertical Carousel")),
		h.Div(
			h.Class("h-64"),
			New(slides, WithOrientation("vertical"), WithClass("h-full")),
		),
	)
}

// ExampleHTMX demonstrates an HTMX-enabled carousel
func ExampleHTMX() g.Node {
	slides := []g.Node{
		Item(
			h.Div(
				h.Class("p-8 bg-gradient-to-br from-indigo-500 to-purple-600 rounded-lg"),
				h.H3(h.Class("text-2xl font-bold text-white mb-4"), g.Text("HTMX Carousel")),
				h.P(h.Class("text-white/90"), g.Text("This carousel uses HTMX for smooth server-side updates.")),
			),
		),
		Item(
			h.Div(
				h.Class("p-8 bg-gradient-to-br from-pink-500 to-rose-600 rounded-lg"),
				h.H3(h.Class("text-2xl font-bold text-white mb-4"), g.Text("Server-Side Navigation")),
				h.P(h.Class("text-white/90"), g.Text("Each slide transition is handled by the server.")),
			),
		),
		Item(
			h.Div(
				h.Class("p-8 bg-gradient-to-br from-cyan-500 to-blue-600 rounded-lg"),
				h.H3(h.Class("text-2xl font-bold text-white mb-4"), g.Text("Progressive Enhancement")),
				h.P(h.Class("text-white/90"), g.Text("Works without JavaScript, enhanced with HTMX.")),
			),
		),
	}

	return h.Div(
		h.Class("w-full max-w-4xl mx-auto p-4"),
		h.H2(h.Class("text-2xl font-bold mb-4"), g.Text("HTMX Carousel")),
		NewHTMX("htmx-carousel-demo", slides, WithHTMXLoop(), WithHTMXAutoPlay(4000)),
	)
}

// ExampleMultipleItems demonstrates showing multiple items at once
func ExampleMultipleItems() g.Node {
	// For multiple items, we adjust the item width
	slides := []g.Node{
		h.Div(
			h.Class("basis-1/3 px-2"),
			h.Div(
				h.Class("bg-gray-100 rounded-lg p-4 h-48 flex items-center justify-center"),
				h.Span(h.Class("text-lg font-semibold"), g.Text("Item 1")),
			),
		),
		h.Div(
			h.Class("basis-1/3 px-2"),
			h.Div(
				h.Class("bg-gray-100 rounded-lg p-4 h-48 flex items-center justify-center"),
				h.Span(h.Class("text-lg font-semibold"), g.Text("Item 2")),
			),
		),
		h.Div(
			h.Class("basis-1/3 px-2"),
			h.Div(
				h.Class("bg-gray-100 rounded-lg p-4 h-48 flex items-center justify-center"),
				h.Span(h.Class("text-lg font-semibold"), g.Text("Item 3")),
			),
		),
		h.Div(
			h.Class("basis-1/3 px-2"),
			h.Div(
				h.Class("bg-gray-100 rounded-lg p-4 h-48 flex items-center justify-center"),
				h.Span(h.Class("text-lg font-semibold"), g.Text("Item 4")),
			),
		),
		h.Div(
			h.Class("basis-1/3 px-2"),
			h.Div(
				h.Class("bg-gray-100 rounded-lg p-4 h-48 flex items-center justify-center"),
				h.Span(h.Class("text-lg font-semibold"), g.Text("Item 5")),
			),
		),
		h.Div(
			h.Class("basis-1/3 px-2"),
			h.Div(
				h.Class("bg-gray-100 rounded-lg p-4 h-48 flex items-center justify-center"),
				h.Span(h.Class("text-lg font-semibold"), g.Text("Item 6")),
			),
		),
	}

	return h.Div(
		h.Class("w-full max-w-6xl mx-auto p-4"),
		h.H2(h.Class("text-2xl font-bold mb-4"), g.Text("Multiple Items Carousel")),
		h.P(h.Class("text-gray-600 mb-4"), g.Text("Shows 3 items at once, scrolls by 3")),
		New(slides, WithSlidesToScroll(3), WithLoop()),
	)
}

// ExampleCustomStyling demonstrates custom styled carousel
func ExampleCustomStyling() g.Node {
	slides := []g.Node{
		Item(
			h.Div(
				h.Class("relative h-80"),
				h.Div(
					h.Class("absolute inset-0 bg-gradient-to-r from-violet-600 to-indigo-600 rounded-2xl"),
				),
				h.Div(
					h.Class("relative h-full flex flex-col justify-center items-center text-white p-8"),
					h.Div(h.Class("text-6xl mb-4"), g.Text("✨")),
					h.H3(h.Class("text-3xl font-bold mb-2"), g.Text("Premium Features")),
					h.P(h.Class("text-lg text-center max-w-md"), g.Text("Experience the best with our advanced carousel component")),
				),
			),
			WithClass("px-2"),
		),
		Item(
			h.Div(
				h.Class("relative h-80"),
				h.Div(
					h.Class("absolute inset-0 bg-gradient-to-r from-pink-600 to-rose-600 rounded-2xl"),
				),
				h.Div(
					h.Class("relative h-full flex flex-col justify-center items-center text-white p-8"),
					h.Div(h.Class("text-6xl mb-4"), g.Text("🚀")),
					h.H3(h.Class("text-3xl font-bold mb-2"), g.Text("Lightning Fast")),
					h.P(h.Class("text-lg text-center max-w-md"), g.Text("Optimized performance for the best user experience")),
				),
			),
			WithClass("px-2"),
		),
		Item(
			h.Div(
				h.Class("relative h-80"),
				h.Div(
					h.Class("absolute inset-0 bg-gradient-to-r from-amber-600 to-orange-600 rounded-2xl"),
				),
				h.Div(
					h.Class("relative h-full flex flex-col justify-center items-center text-white p-8"),
					h.Div(h.Class("text-6xl mb-4"), g.Text("💎")),
					h.H3(h.Class("text-3xl font-bold mb-2"), g.Text("Beautiful Design")),
					h.P(h.Class("text-lg text-center max-w-md"), g.Text("Stunning visuals that captivate your audience")),
				),
			),
			WithClass("px-2"),
		),
	}

	return h.Div(
		h.Class("w-full max-w-4xl mx-auto p-4"),
		h.H2(h.Class("text-2xl font-bold mb-4"), g.Text("Custom Styled Carousel")),
		New(slides, 
			WithLoop(),
			WithAutoPlay(6000),
			WithClass("bg-gradient-to-br from-gray-50 to-gray-100 rounded-3xl shadow-2xl p-4"),
		),
	)
}