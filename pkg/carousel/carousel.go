package carousel

import (
	"fmt"
	"strings"

	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

// Option is a functional option for configuring a carousel
type Option func(*config)

type config struct {
	class           string
	orientation     string // horizontal or vertical
	loop            bool
	autoPlay        bool
	autoPlayDelay   int // milliseconds
	showIndicators  bool
	showControls    bool
	align           string // start, center, end
	slidesToScroll  int
}

// New creates a new carousel component
func New(slides []g.Node, opts ...Option) g.Node {
	cfg := &config{
		class:          "",
		orientation:    "horizontal",
		loop:           false,
		autoPlay:       false,
		autoPlayDelay:  3000,
		showIndicators: true,
		showControls:   true,
		align:          "start",
		slidesToScroll: 1,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	classes := strings.TrimSpace("relative " + cfg.class)
	
	return html.Div(
		html.Class(classes),
		g.Attr("data-carousel", "true"),
		g.Attr("data-orientation", cfg.orientation),
		g.If(cfg.loop, g.Attr("data-loop", "true")),
		g.If(cfg.autoPlay, g.Attr("data-auto-play", fmt.Sprintf("%d", cfg.autoPlayDelay))),
		g.Attr("data-align", cfg.align),
		g.Attr("data-slides-to-scroll", fmt.Sprintf("%d", cfg.slidesToScroll)),
		
		// Carousel content wrapper
		ContentComponent(slides, cfg),
		
		// Controls
		g.If(cfg.showControls,
			g.Group([]g.Node{
				PreviousButton(cfg),
				NextButton(cfg),
			}),
		),
		
		// Indicators
		g.If(cfg.showIndicators,
			Indicators(len(slides), cfg),
		),
	)
}

// Content creates the carousel content container
func ContentComponent(slides []g.Node, cfg *config) g.Node {
	contentClasses := "overflow-hidden"
	viewportClasses := "flex"
	
	if cfg.orientation == "vertical" {
		viewportClasses += " flex-col"
	} else {
		viewportClasses += " flex-row"
	}
	
	if cfg.align == "center" {
		viewportClasses += " items-center"
	} else if cfg.align == "end" {
		viewportClasses += " items-end"
	}
	
	return html.Div(
		html.Class(contentClasses),
		html.Div(
			html.Class(viewportClasses),
			g.Attr("data-carousel-viewport", "true"),
			g.Group(slides),
		),
	)
}

// Item creates a carousel slide item
func Item(content g.Node, opts ...Option) g.Node {
	cfg := &config{
		class: "",
	}
	
	for _, opt := range opts {
		opt(cfg)
	}
	
	classes := strings.TrimSpace("min-w-0 shrink-0 grow-0 basis-full " + cfg.class)
	
	return html.Div(
		html.Class(classes),
		g.Attr("role", "group"),
		g.Attr("aria-roledescription", "slide"),
		g.Attr("data-carousel-item", "true"),
		content,
	)
}

// PreviousButton creates the previous button
func PreviousButton(cfg *config) g.Node {
	position := "absolute left-4"
	if cfg.orientation == "vertical" {
		position = "absolute top-4"
	}
	
	return html.Button(
		html.Type("button"),
		html.Class(position+" top-1/2 -translate-y-1/2 h-8 w-8 rounded-full bg-white/80 hover:bg-white shadow-md flex items-center justify-center disabled:opacity-50 disabled:cursor-not-allowed"),
		g.Attr("aria-label", "Previous slide"),
		g.Attr("data-carousel-prev", "true"),
		g.If(!cfg.loop, g.Attr("disabled", "disabled")),
		
		// Chevron left icon
		g.Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
			<path d="m15 18-6-6 6-6"/>
		</svg>`),
	)
}

// NextButton creates the next button
func NextButton(cfg *config) g.Node {
	position := "absolute right-4"
	if cfg.orientation == "vertical" {
		position = "absolute bottom-4"
	}
	
	return html.Button(
		html.Type("button"),
		html.Class(position+" top-1/2 -translate-y-1/2 h-8 w-8 rounded-full bg-white/80 hover:bg-white shadow-md flex items-center justify-center disabled:opacity-50 disabled:cursor-not-allowed"),
		g.Attr("aria-label", "Next slide"),
		g.Attr("data-carousel-next", "true"),
		
		// Chevron right icon
		g.Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
			<path d="m9 18 6-6-6-6"/>
		</svg>`),
	)
}

// Indicators creates the carousel indicators
func Indicators(count int, cfg *config) g.Node {
	position := "absolute bottom-4 left-1/2 -translate-x-1/2"
	containerClasses := "flex gap-1"
	
	if cfg.orientation == "vertical" {
		position = "absolute right-4 top-1/2 -translate-y-1/2"
		containerClasses = "flex flex-col gap-1"
	}
	
	indicators := make([]g.Node, count)
	for i := 0; i < count; i++ {
		indicators[i] = Indicator(i, i == 0)
	}
	
	return html.Div(
		html.Class(position),
		html.Div(
			html.Class(containerClasses),
			g.Attr("role", "tablist"),
			g.Attr("data-carousel-indicators", "true"),
			g.Group(indicators),
		),
	)
}

// Indicator creates a single carousel indicator
func Indicator(index int, active bool) g.Node {
	classes := "w-2 h-2 rounded-full transition-all"
	if active {
		classes += " bg-primary w-6"
	} else {
		classes += " bg-primary/30"
	}
	
	return html.Button(
		html.Type("button"),
		html.Class(classes),
		g.Attr("role", "tab"),
		g.Attr("aria-label", fmt.Sprintf("Go to slide %d", index+1)),
		g.Attr("data-carousel-indicator", fmt.Sprintf("%d", index)),
		g.If(active, g.Attr("aria-selected", "true")),
	)
}

// Option functions

// WithClass adds custom CSS classes
func WithClass(class string) Option {
	return func(c *config) {
		c.class = class
	}
}

// WithOrientation sets the carousel orientation
func WithOrientation(orientation string) Option {
	return func(c *config) {
		c.orientation = orientation
	}
}

// WithLoop enables infinite loop scrolling
func WithLoop() Option {
	return func(c *config) {
		c.loop = true
	}
}

// WithAutoPlay enables auto-play with optional delay
func WithAutoPlay(delay ...int) Option {
	return func(c *config) {
		c.autoPlay = true
		if len(delay) > 0 {
			c.autoPlayDelay = delay[0]
		}
	}
}

// WithoutIndicators hides the indicators
func WithoutIndicators() Option {
	return func(c *config) {
		c.showIndicators = false
	}
}

// WithoutControls hides the navigation controls
func WithoutControls() Option {
	return func(c *config) {
		c.showControls = false
	}
}

// WithAlign sets the slide alignment
func WithAlign(align string) Option {
	return func(c *config) {
		c.align = align
	}
}

// WithSlidesToScroll sets how many slides to scroll at once
func WithSlidesToScroll(count int) Option {
	return func(c *config) {
		c.slidesToScroll = count
	}
}