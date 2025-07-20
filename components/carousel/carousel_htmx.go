package carousel

import (
	"fmt"
	"strings"

	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

// HTMXConfig provides HTMX-specific configuration
type HTMXConfig struct {
	config
	endpoint       string
	updateEndpoint string
	swapTarget     string
}

// SetOrientation sets the orientation for HTMXConfig
func (c *HTMXConfig) SetOrientation(orientation string) {
	c.orientation = orientation
}

// SetEndpoint sets the endpoint for HTMXConfig
func (c *HTMXConfig) SetEndpoint(endpoint string) {
	c.endpoint = endpoint
}

// HTMXOption is a functional option for HTMX carousel
type HTMXOption func(*HTMXConfig)

// NewHTMX creates an HTMX-enabled carousel
func NewHTMX(id string, slides []g.Node, opts ...HTMXOption) g.Node {
	cfg := &HTMXConfig{
		config: config{
			class:          "",
			orientation:    "horizontal",
			loop:           false,
			autoPlay:       false,
			autoPlayDelay:  3000,
			showIndicators: true,
			showControls:   true,
			align:          "start",
			slidesToScroll: 1,
		},
		endpoint:       "/carousel/" + id,
		updateEndpoint: "/carousel/" + id + "/slide",
		swapTarget:     "#" + id + "-viewport",
	}

	for _, opt := range opts {
		opt(cfg)
	}

	classes := strings.TrimSpace("relative " + cfg.class)
	
	return html.Div(
		html.ID(id),
		html.Class(classes),
		g.Attr("data-carousel", "htmx"),
		g.Attr("data-orientation", cfg.orientation),
		g.If(cfg.loop, g.Attr("data-loop", "true")),
		g.If(cfg.autoPlay, 
			g.Group([]g.Node{
				g.Attr("data-auto-play", fmt.Sprintf("%d", cfg.autoPlayDelay)),
				g.Attr("hx-trigger", fmt.Sprintf("load, every %dms", cfg.autoPlayDelay)),
				g.Attr("hx-get", cfg.endpoint+"/auto-next"),
				g.Attr("hx-target", cfg.swapTarget),
				g.Attr("hx-swap", "innerHTML"),
			}),
		),
		g.Attr("data-align", cfg.align),
		g.Attr("data-slides-to-scroll", fmt.Sprintf("%d", cfg.slidesToScroll)),
		g.Attr("data-current-slide", "0"),
		g.Attr("data-total-slides", fmt.Sprintf("%d", len(slides))),
		
		// Carousel content wrapper
		HTMXContent(id, slides, cfg),
		
		// Controls
		g.If(cfg.showControls,
			g.Group([]g.Node{
				HTMXPreviousButton(id, cfg),
				HTMXNextButton(id, cfg),
			}),
		),
		
		// Indicators
		g.If(cfg.showIndicators,
			HTMXIndicators(id, len(slides), cfg),
		),
	)
}

// HTMXContent creates the HTMX carousel content container
func HTMXContent(id string, slides []g.Node, cfg *HTMXConfig) g.Node {
	contentClasses := "overflow-hidden"
	viewportClasses := "flex transition-transform duration-300"
	
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
			html.ID(id+"-viewport"),
			html.Class(viewportClasses),
			g.Attr("data-carousel-viewport", "htmx"),
			g.Attr("style", "transform: translateX(0)"),
			g.Group(slides),
		),
	)
}

// HTMXPreviousButton creates an HTMX-enabled previous button
func HTMXPreviousButton(id string, cfg *HTMXConfig) g.Node {
	position := "absolute left-4"
	if cfg.orientation == "vertical" {
		position = "absolute top-4"
	}
	
	return html.Button(
		html.Type("button"),
		html.Class(position+" top-1/2 -translate-y-1/2 h-8 w-8 rounded-full bg-white/80 hover:bg-white shadow-md flex items-center justify-center disabled:opacity-50 disabled:cursor-not-allowed transition-opacity"),
		g.Attr("aria-label", "Previous slide"),
		g.Attr("data-carousel-prev", "htmx"),
		g.Attr("hx-get", cfg.endpoint+"/prev"),
		g.Attr("hx-target", cfg.swapTarget),
		g.Attr("hx-swap", "innerHTML"),
		g.Attr("hx-trigger", "click"),
		g.If(!cfg.loop, g.Attr("disabled", "disabled")),
		
		// Chevron left icon
		g.Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
			<path d="m15 18-6-6 6-6"/>
		</svg>`),
	)
}

// HTMXNextButton creates an HTMX-enabled next button
func HTMXNextButton(id string, cfg *HTMXConfig) g.Node {
	position := "absolute right-4"
	if cfg.orientation == "vertical" {
		position = "absolute bottom-4"
	}
	
	return html.Button(
		html.Type("button"),
		html.Class(position+" top-1/2 -translate-y-1/2 h-8 w-8 rounded-full bg-white/80 hover:bg-white shadow-md flex items-center justify-center disabled:opacity-50 disabled:cursor-not-allowed transition-opacity"),
		g.Attr("aria-label", "Next slide"),
		g.Attr("data-carousel-next", "htmx"),
		g.Attr("hx-get", cfg.endpoint+"/next"),
		g.Attr("hx-target", cfg.swapTarget),
		g.Attr("hx-swap", "innerHTML"),
		g.Attr("hx-trigger", "click"),
		
		// Chevron right icon
		g.Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
			<path d="m9 18 6-6-6-6"/>
		</svg>`),
	)
}

// HTMXIndicators creates HTMX-enabled carousel indicators
func HTMXIndicators(id string, count int, cfg *HTMXConfig) g.Node {
	position := "absolute bottom-4 left-1/2 -translate-x-1/2"
	containerClasses := "flex gap-1"
	
	if cfg.orientation == "vertical" {
		position = "absolute right-4 top-1/2 -translate-y-1/2"
		containerClasses = "flex flex-col gap-1"
	}
	
	indicators := make([]g.Node, count)
	for i := 0; i < count; i++ {
		indicators[i] = HTMXIndicator(id, i, i == 0, cfg)
	}
	
	return html.Div(
		html.Class(position),
		html.Div(
			html.ID(id+"-indicators"),
			html.Class(containerClasses),
			g.Attr("role", "tablist"),
			g.Attr("data-carousel-indicators", "htmx"),
			g.Group(indicators),
		),
	)
}

// HTMXIndicator creates a single HTMX-enabled carousel indicator
func HTMXIndicator(id string, index int, active bool, cfg *HTMXConfig) g.Node {
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
		g.Attr("hx-get", fmt.Sprintf("%s/goto/%d", cfg.endpoint, index)),
		g.Attr("hx-target", cfg.swapTarget),
		g.Attr("hx-swap", "innerHTML"),
		g.Attr("hx-trigger", "click"),
		g.If(active, g.Attr("aria-selected", "true")),
	)
}

// Slide response for HTMX updates
func SlideResponse(slides []g.Node, currentIndex int, cfg *HTMXConfig) g.Node {
	viewportClasses := "flex transition-transform duration-300"
	
	if cfg.orientation == "vertical" {
		viewportClasses += " flex-col"
		var translateValue string
		if currentIndex == 0 {
			translateValue = "translateY(0%)"
		} else {
			translateValue = fmt.Sprintf("translateY(-%d%%)", currentIndex*100)
		}
		return html.Div(
			html.Class(viewportClasses),
			g.Attr("data-carousel-viewport", "htmx"),
			g.Attr("style", fmt.Sprintf("transform: %s", translateValue)),
			g.Group(slides),
		)
	}
	
	var translateValue string
	if currentIndex == 0 {
		translateValue = "translateX(0%)"
	} else {
		translateValue = fmt.Sprintf("translateX(-%d%%)", currentIndex*100)
	}
	return html.Div(
		html.Class(viewportClasses),
		g.Attr("data-carousel-viewport", "htmx"),
		g.Attr("style", fmt.Sprintf("transform: %s", translateValue)),
		g.Group(slides),
	)
}

// IndicatorsResponse updates indicators via HTMX
func IndicatorsResponse(id string, count int, currentIndex int, cfg *HTMXConfig) g.Node {
	containerClasses := "flex gap-1"
	
	if cfg.orientation == "vertical" {
		containerClasses = "flex flex-col gap-1"
	}
	
	indicators := make([]g.Node, count)
	for i := 0; i < count; i++ {
		indicators[i] = HTMXIndicator(id, i, i == currentIndex, cfg)
	}
	
	return html.Div(
		html.Class(containerClasses),
		g.Attr("role", "tablist"),
		g.Attr("data-carousel-indicators", "htmx"),
		g.Group(indicators),
	)
}

// HTMX Option functions

// WithHTMXClass adds custom CSS classes
func WithHTMXClass(class string) HTMXOption {
	return func(c *HTMXConfig) {
		c.class = class
	}
}

// WithHTMXOrientation sets the carousel orientation
func WithHTMXOrientation(orientation string) HTMXOption {
	return func(c *HTMXConfig) {
		c.orientation = orientation
	}
}

// WithHTMXLoop enables infinite loop scrolling
func WithHTMXLoop() HTMXOption {
	return func(c *HTMXConfig) {
		c.loop = true
	}
}

// WithHTMXAutoPlay enables auto-play with optional delay
func WithHTMXAutoPlay(delay ...int) HTMXOption {
	return func(c *HTMXConfig) {
		c.autoPlay = true
		if len(delay) > 0 {
			c.autoPlayDelay = delay[0]
		}
	}
}

// WithHTMXEndpoint sets custom endpoints
func WithHTMXEndpoint(endpoint string) HTMXOption {
	return func(c *HTMXConfig) {
		c.endpoint = endpoint
		c.updateEndpoint = endpoint + "/slide"
	}
}

// WithHTMXSwapTarget sets custom swap target
func WithHTMXSwapTarget(target string) HTMXOption {
	return func(c *HTMXConfig) {
		c.swapTarget = target
	}
}

// WithoutHTMXIndicators hides the indicators
func WithoutHTMXIndicators() HTMXOption {
	return func(c *HTMXConfig) {
		c.showIndicators = false
	}
}

// WithoutHTMXControls hides the navigation controls
func WithoutHTMXControls() HTMXOption {
	return func(c *HTMXConfig) {
		c.showControls = false
	}
}