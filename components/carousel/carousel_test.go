package carousel_test

import (
	"strings"
	"testing"

	"github.com/rizome-dev/shadcn-gomponents/components/carousel"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestNew(t *testing.T) {
	t.Run("creates basic carousel with default options", func(t *testing.T) {
		slides := []g.Node{
			carousel.Item(h.Div(g.Text("Slide 1"))),
			carousel.Item(h.Div(g.Text("Slide 2"))),
		}
		
		result := carousel.New(slides)
		html := render(result)
		
		// Check basic structure
		if !strings.Contains(html, `data-carousel="true"`) {
			t.Error("expected carousel data attribute")
		}
		if !strings.Contains(html, `data-orientation="horizontal"`) {
			t.Error("expected horizontal orientation by default")
		}
		if !strings.Contains(html, "Slide 1") {
			t.Error("expected first slide content")
		}
		if !strings.Contains(html, "Slide 2") {
			t.Error("expected second slide content")
		}
		if !strings.Contains(html, `data-carousel-prev="true"`) {
			t.Error("expected previous button")
		}
		if !strings.Contains(html, `data-carousel-next="true"`) {
			t.Error("expected next button")
		}
	})

	t.Run("creates carousel with loop enabled", func(t *testing.T) {
		slides := []g.Node{
			carousel.Item(h.Div(g.Text("Slide"))),
		}
		
		result := carousel.New(slides, carousel.WithLoop())
		html := render(result)
		
		if !strings.Contains(html, `data-loop="true"`) {
			t.Error("expected loop attribute")
		}
		// Previous button should not be disabled when loop is enabled
		if strings.Contains(html, `data-carousel-prev="true" disabled`) {
			t.Error("previous button should not be disabled with loop")
		}
	})

	t.Run("creates carousel with auto-play", func(t *testing.T) {
		slides := []g.Node{
			carousel.Item(h.Div(g.Text("Slide"))),
		}
		
		result := carousel.New(slides, carousel.WithAutoPlay(5000))
		html := render(result)
		
		if !strings.Contains(html, `data-auto-play="5000"`) {
			t.Error("expected auto-play attribute with delay")
		}
	})

	t.Run("creates vertical carousel", func(t *testing.T) {
		slides := []g.Node{
			carousel.Item(h.Div(g.Text("Slide"))),
		}
		
		result := carousel.New(slides, carousel.WithOrientation("vertical"))
		html := render(result)
		
		if !strings.Contains(html, `data-orientation="vertical"`) {
			t.Error("expected vertical orientation")
		}
		if !strings.Contains(html, "flex-col") {
			t.Error("expected flex-col class for vertical orientation")
		}
	})

	t.Run("hides indicators when requested", func(t *testing.T) {
		slides := []g.Node{
			carousel.Item(h.Div(g.Text("Slide"))),
		}
		
		result := carousel.New(slides, carousel.WithoutIndicators())
		html := render(result)
		
		if strings.Contains(html, `data-carousel-indicators`) {
			t.Error("indicators should be hidden")
		}
	})

	t.Run("hides controls when requested", func(t *testing.T) {
		slides := []g.Node{
			carousel.Item(h.Div(g.Text("Slide"))),
		}
		
		result := carousel.New(slides, carousel.WithoutControls())
		html := render(result)
		
		if strings.Contains(html, `data-carousel-prev`) {
			t.Error("previous button should be hidden")
		}
		if strings.Contains(html, `data-carousel-next`) {
			t.Error("next button should be hidden")
		}
	})

	t.Run("applies custom classes", func(t *testing.T) {
		slides := []g.Node{
			carousel.Item(h.Div(g.Text("Slide"))),
		}
		
		result := carousel.New(slides, carousel.WithClass("custom-class shadow-lg"))
		html := render(result)
		
		if !strings.Contains(html, "custom-class") {
			t.Error("expected custom class")
		}
		if !strings.Contains(html, "shadow-lg") {
			t.Error("expected shadow-lg class")
		}
	})

	t.Run("sets slides to scroll", func(t *testing.T) {
		slides := []g.Node{
			carousel.Item(h.Div(g.Text("Slide"))),
		}
		
		result := carousel.New(slides, carousel.WithSlidesToScroll(3))
		html := render(result)
		
		if !strings.Contains(html, `data-slides-to-scroll="3"`) {
			t.Error("expected slides-to-scroll attribute")
		}
	})

	t.Run("sets alignment", func(t *testing.T) {
		slides := []g.Node{
			carousel.Item(h.Div(g.Text("Slide"))),
		}
		
		result := carousel.New(slides, carousel.WithAlign("center"))
		html := render(result)
		
		if !strings.Contains(html, `data-align="center"`) {
			t.Error("expected align attribute")
		}
		if !strings.Contains(html, "items-center") {
			t.Error("expected items-center class")
		}
	})
}

func TestItem(t *testing.T) {
	t.Run("creates carousel item with default classes", func(t *testing.T) {
		result := carousel.Item(h.Div(g.Text("Content")))
		html := render(result)
		
		if !strings.Contains(html, `data-carousel-item="true"`) {
			t.Error("expected carousel item data attribute")
		}
		if !strings.Contains(html, `role="group"`) {
			t.Error("expected group role")
		}
		if !strings.Contains(html, `aria-roledescription="slide"`) {
			t.Error("expected slide role description")
		}
		if !strings.Contains(html, "basis-full") {
			t.Error("expected basis-full class")
		}
		if !strings.Contains(html, "Content") {
			t.Error("expected item content")
		}
	})

	t.Run("applies custom classes to item", func(t *testing.T) {
		result := carousel.Item(
			h.Div(g.Text("Content")),
			carousel.WithClass("custom-item-class"),
		)
		html := render(result)
		
		if !strings.Contains(html, "custom-item-class") {
			t.Error("expected custom item class")
		}
	})
}

func TestIndicators(t *testing.T) {
	t.Run("creates correct number of indicators", func(t *testing.T) {
		// Test indicators through the full carousel component
		slides := make([]g.Node, 3)
		for i := 0; i < 3; i++ {
			slides[i] = carousel.Item(h.Div(g.Text("Slide")))
		}
		
		result := carousel.New(slides)
		html := render(result)
		
		indicatorCount := strings.Count(html, `data-carousel-indicator=`)
		if indicatorCount != 3 {
			t.Errorf("expected 3 indicators, got %d", indicatorCount)
		}
	})

	t.Run("marks first indicator as active", func(t *testing.T) {
		slides := make([]g.Node, 2)
		for i := 0; i < 2; i++ {
			slides[i] = carousel.Item(h.Div(g.Text("Slide")))
		}
		
		result := carousel.New(slides)
		html := render(result)
		
		if !strings.Contains(html, `aria-selected="true"`) {
			t.Error("expected first indicator to be selected")
		}
		if !strings.Contains(html, "bg-primary w-6") {
			t.Error("expected active indicator styling")
		}
	})

	t.Run("positions indicators for vertical carousel", func(t *testing.T) {
		slides := make([]g.Node, 2)
		for i := 0; i < 2; i++ {
			slides[i] = carousel.Item(h.Div(g.Text("Slide")))
		}
		
		result := carousel.New(slides, carousel.WithOrientation("vertical"))
		html := render(result)
		
		if !strings.Contains(html, "flex-col") {
			t.Error("expected flex-col for vertical indicators")
		}
		if !strings.Contains(html, "right-4") {
			t.Error("expected right positioning for vertical carousel")
		}
	})
}

func TestHTMXCarousel(t *testing.T) {
	t.Run("creates HTMX-enabled carousel", func(t *testing.T) {
		slides := []g.Node{
			carousel.Item(h.Div(g.Text("Slide 1"))),
			carousel.Item(h.Div(g.Text("Slide 2"))),
		}
		
		result := carousel.NewHTMX("test-carousel", slides)
		html := render(result)
		
		if !strings.Contains(html, `id="test-carousel"`) {
			t.Error("expected carousel ID")
		}
		if !strings.Contains(html, `data-carousel="htmx"`) {
			t.Error("expected HTMX carousel data attribute")
		}
		if !strings.Contains(html, `hx-get="/carousel/test-carousel/prev"`) {
			t.Error("expected HTMX previous endpoint")
		}
		if !strings.Contains(html, `hx-get="/carousel/test-carousel/next"`) {
			t.Error("expected HTMX next endpoint")
		}
		if !strings.Contains(html, `hx-target="#test-carousel-viewport"`) {
			t.Error("expected HTMX target")
		}
		if !strings.Contains(html, `data-total-slides="2"`) {
			t.Error("expected total slides attribute")
		}
	})

	t.Run("creates HTMX carousel with auto-play", func(t *testing.T) {
		slides := []g.Node{
			carousel.Item(h.Div(g.Text("Slide"))),
		}
		
		result := carousel.NewHTMX("auto-carousel", slides, carousel.WithHTMXAutoPlay(2000))
		html := render(result)
		
		if !strings.Contains(html, `hx-trigger="load, every 2000ms"`) {
			t.Error("expected HTMX auto-play trigger")
		}
		if !strings.Contains(html, `hx-get="/carousel/auto-carousel/auto-next"`) {
			t.Error("expected auto-next endpoint")
		}
	})

	t.Run("creates HTMX carousel with custom endpoint", func(t *testing.T) {
		slides := []g.Node{
			carousel.Item(h.Div(g.Text("Slide"))),
		}
		
		result := carousel.NewHTMX("custom-carousel", slides, carousel.WithHTMXEndpoint("/api/carousel"))
		html := render(result)
		
		if !strings.Contains(html, `hx-get="/api/carousel/prev"`) {
			t.Error("expected custom endpoint for prev")
		}
		if !strings.Contains(html, `hx-get="/api/carousel/next"`) {
			t.Error("expected custom endpoint for next")
		}
	})

	t.Run("creates HTMX indicators with goto endpoints", func(t *testing.T) {
		slides := []g.Node{
			carousel.Item(h.Div(g.Text("Slide 1"))),
			carousel.Item(h.Div(g.Text("Slide 2"))),
			carousel.Item(h.Div(g.Text("Slide 3"))),
		}
		
		result := carousel.NewHTMX("indicator-carousel", slides)
		html := render(result)
		
		if !strings.Contains(html, `hx-get="/carousel/indicator-carousel/goto/0"`) {
			t.Error("expected goto endpoint for first indicator")
		}
		if !strings.Contains(html, `hx-get="/carousel/indicator-carousel/goto/1"`) {
			t.Error("expected goto endpoint for second indicator")
		}
		if !strings.Contains(html, `hx-get="/carousel/indicator-carousel/goto/2"`) {
			t.Error("expected goto endpoint for third indicator")
		}
	})

	t.Run("creates HTMX carousel without indicators", func(t *testing.T) {
		slides := []g.Node{
			carousel.Item(h.Div(g.Text("Slide"))),
		}
		
		result := carousel.NewHTMX("no-indicators", slides, carousel.WithoutHTMXIndicators())
		html := render(result)
		
		if strings.Contains(html, "data-carousel-indicators") {
			t.Error("indicators should not be present")
		}
	})

	t.Run("creates HTMX carousel without controls", func(t *testing.T) {
		slides := []g.Node{
			carousel.Item(h.Div(g.Text("Slide"))),
		}
		
		result := carousel.NewHTMX("no-controls", slides, carousel.WithoutHTMXControls())
		html := render(result)
		
		if strings.Contains(html, "data-carousel-prev") {
			t.Error("previous button should not be present")
		}
		if strings.Contains(html, "data-carousel-next") {
			t.Error("next button should not be present")
		}
	})

	t.Run("creates HTMX carousel with custom swap target", func(t *testing.T) {
		slides := []g.Node{
			carousel.Item(h.Div(g.Text("Slide"))),
		}
		
		result := carousel.NewHTMX("custom-target", slides, carousel.WithHTMXSwapTarget("#custom-viewport"))
		html := render(result)
		
		if !strings.Contains(html, `hx-target="#custom-viewport"`) {
			t.Error("expected custom swap target")
		}
	})
}

func TestHTMXResponses(t *testing.T) {
	t.Run("SlideResponse generates correct viewport", func(t *testing.T) {
		slides := []g.Node{
			carousel.Item(h.Div(g.Text("Slide 1"))),
			carousel.Item(h.Div(g.Text("Slide 2"))),
		}
		
		// Test initial state (index 0)
		result := carousel.SlideResponse(slides, 0, &carousel.HTMXConfig{})
		html := render(result)
		
		if !strings.Contains(html, `transform: translateX(0%)`) {
			t.Errorf("expected translateX(0%%) for first slide, got: %s", html)
		}
		
		// Test second slide (index 1)
		result = carousel.SlideResponse(slides, 1, &carousel.HTMXConfig{})
		html = render(result)
		
		if !strings.Contains(html, `transform: translateX(-100%)`) {
			t.Error("expected -100% transform for second slide")
		}
	})

	t.Run("SlideResponse handles vertical orientation", func(t *testing.T) {
		slides := []g.Node{
			carousel.Item(h.Div(g.Text("Slide 1"))),
			carousel.Item(h.Div(g.Text("Slide 2"))),
		}
		
		cfg := &carousel.HTMXConfig{}
		cfg.SetOrientation("vertical")
		
		result := carousel.SlideResponse(slides, 1, cfg)
		html := render(result)
		
		if !strings.Contains(html, `transform: translateY(-100%)`) {
			t.Error("expected translateY for vertical carousel")
		}
		if !strings.Contains(html, "flex-col") {
			t.Error("expected flex-col for vertical orientation")
		}
	})

	t.Run("IndicatorsResponse updates active state", func(t *testing.T) {
		cfg := &carousel.HTMXConfig{}
		cfg.SetEndpoint("/test")
		
		result := carousel.IndicatorsResponse("test", 3, 1, cfg)
		html := render(result)
		
		// Should have 3 indicators
		indicatorCount := strings.Count(html, `data-carousel-indicator="`)
		if indicatorCount != 3 {
			t.Errorf("expected 3 indicators, got %d in html: %s", indicatorCount, html)
		}
		
		// Second indicator should be active
		parts := strings.Split(html, `data-carousel-indicator="1"`)
		if len(parts) < 2 {
			t.Error("expected second indicator")
		}
		// Look for aria-selected in the portion after the second indicator
		if !strings.Contains(parts[1], `aria-selected="true"`) {
			t.Error("expected second indicator to be selected")
		}
	})
}

// Helper function to render components
func render(node g.Node) string {
	var sb strings.Builder
	_ = node.Render(&sb)
	return sb.String()
}