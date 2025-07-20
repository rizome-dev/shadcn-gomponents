package collapsible_test

import (
	"bytes"
	"strings"
	"testing"

	g "maragu.dev/gomponents"
	"github.com/rizome-dev/shadcn-gomponents/components/collapsible"
)

func TestCollapsible(t *testing.T) {
	t.Run("renders details element", func(t *testing.T) {
		c := collapsible.New(
			collapsible.Props{},
			g.Text("Summary"),
			collapsible.ContentMarker,
			g.Text("Content"),
		)

		var buf bytes.Buffer
		err := c.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "<details") {
			t.Error("expected details element")
		}
		if !strings.Contains(output, "<summary") {
			t.Error("expected summary element")
		}
		if !strings.Contains(output, "Summary") {
			t.Error("expected summary text")
		}
		if !strings.Contains(output, "Content") {
			t.Error("expected content text")
		}
	})

	t.Run("renders open when open prop is true", func(t *testing.T) {
		c := collapsible.New(
			collapsible.Props{Open: true},
			g.Text("Summary"),
			collapsible.ContentMarker,
			g.Text("Content"),
		)

		var buf bytes.Buffer
		err := c.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, `open=""`) && !strings.Contains(output, `open`) {
			t.Error("expected open attribute")
		}
	})

	t.Run("renders closed when open prop is false", func(t *testing.T) {
		c := collapsible.New(
			collapsible.Props{Open: false},
			g.Text("Summary"),
			collapsible.ContentMarker,
			g.Text("Content"),
		)

		var buf bytes.Buffer
		err := c.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if strings.Contains(output, "open") {
			t.Error("should not have open attribute when closed")
		}
	})

	t.Run("renders with custom class", func(t *testing.T) {
		c := collapsible.New(
			collapsible.Props{Class: "custom-class"},
			g.Text("Summary"),
			collapsible.ContentMarker,
			g.Text("Content"),
		)

		var buf bytes.Buffer
		err := c.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "custom-class") {
			t.Error("expected custom class")
		}
	})

	t.Run("properly separates summary and content", func(t *testing.T) {
		c := collapsible.New(
			collapsible.Props{},
			g.Text("Part 1"),
			g.Text("Part 2"),
			collapsible.ContentMarker,
			g.Text("Content 1"),
			g.Text("Content 2"),
		)

		var buf bytes.Buffer
		err := c.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		// Check that summary contains Part 1 and Part 2
		summaryStart := strings.Index(output, "<summary")
		summaryEnd := strings.Index(output, "</summary>")
		if summaryStart == -1 || summaryEnd == -1 {
			t.Fatal("could not find summary tags")
		}
		summaryContent := output[summaryStart:summaryEnd]
		if !strings.Contains(summaryContent, "Part 1") || !strings.Contains(summaryContent, "Part 2") {
			t.Error("expected summary to contain Part 1 and Part 2")
		}

		// Check that content is after summary
		contentAfterSummary := output[summaryEnd:]
		if !strings.Contains(contentAfterSummary, "Content 1") || !strings.Contains(contentAfterSummary, "Content 2") {
			t.Error("expected content after summary")
		}
	})
}

func TestDivCollapsible(t *testing.T) {
	t.Run("renders div with data-state", func(t *testing.T) {
		c := collapsible.DivCollapsible(
			collapsible.Props{Open: true},
			g.Text("Test content"),
		)

		var buf bytes.Buffer
		err := c.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "<div") {
			t.Error("expected div element")
		}
		if !strings.Contains(output, `data-state="open"`) {
			t.Error("expected data-state open")
		}
	})

	t.Run("renders closed state", func(t *testing.T) {
		c := collapsible.DivCollapsible(
			collapsible.Props{Open: false},
			g.Text("Test content"),
		)

		var buf bytes.Buffer
		err := c.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, `data-state="closed"`) {
			t.Error("expected data-state closed")
		}
	})
}

func TestTrigger(t *testing.T) {
	t.Run("renders trigger with role and tabindex", func(t *testing.T) {
		trigger := collapsible.Trigger(
			collapsible.TriggerProps{},
			g.Text("Click me"),
		)

		var buf bytes.Buffer
		err := trigger.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, `role="button"`) {
			t.Error("expected role button")
		}
		if !strings.Contains(output, `tabindex="0"`) {
			t.Error("expected tabindex")
		}
		if !strings.Contains(output, "cursor-pointer") {
			t.Error("expected cursor-pointer class")
		}
	})
}

func TestContent(t *testing.T) {
	t.Run("renders content with animation classes", func(t *testing.T) {
		content := collapsible.CollapsibleContent(
			collapsible.ContentProps{},
			g.Text("Hidden content"),
		)

		var buf bytes.Buffer
		err := content.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "overflow-hidden") {
			t.Error("expected overflow-hidden")
		}
		if !strings.Contains(output, "transition-all") {
			t.Error("expected transition-all")
		}
		if !strings.Contains(output, "data-[state=closed]:animate-accordion-up") {
			t.Error("expected close animation class")
		}
		if !strings.Contains(output, "data-[state=open]:animate-accordion-down") {
			t.Error("expected open animation class")
		}
	})
}

func TestTriggerButton(t *testing.T) {
	t.Run("renders button with icon", func(t *testing.T) {
		button := collapsible.TriggerButton(
			collapsible.TriggerProps{},
			false,
		)

		var buf bytes.Buffer
		err := button.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "<button") {
			t.Error("expected button element")
		}
		if !strings.Contains(output, "<svg") {
			t.Error("expected svg icon")
		}
		if !strings.Contains(output, "sr-only") {
			t.Error("expected screen reader text")
		}
		if !strings.Contains(output, "Toggle") {
			t.Error("expected toggle text")
		}
	})

	t.Run("rotates icon when open", func(t *testing.T) {
		button := collapsible.TriggerButton(
			collapsible.TriggerProps{},
			true,
		)

		var buf bytes.Buffer
		err := button.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "rotate-180") {
			t.Error("expected rotate-180 class when open")
		}
	})
}

func TestExample(t *testing.T) {
	t.Run("renders complete example", func(t *testing.T) {
		example := collapsible.Example()

		var buf bytes.Buffer
		err := example.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "@peduarte starred 3 repositories") {
			t.Error("expected header text")
		}
		if !strings.Contains(output, "@radix-ui/primitives") {
			t.Error("expected first item")
		}
		if !strings.Contains(output, "@radix-ui/colors") {
			t.Error("expected second item")
		}
		if !strings.Contains(output, "@stitches/react") {
			t.Error("expected third item")
		}
		if !strings.Contains(output, "open") {
			t.Error("expected to be open")
		}
	})
}

func TestExampleClosed(t *testing.T) {
	t.Run("renders closed example", func(t *testing.T) {
		example := collapsible.ExampleClosed()

		var buf bytes.Buffer
		err := example.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "Click to expand") {
			t.Error("expected header text")
		}
		if !strings.Contains(output, "Hidden content 1") {
			t.Error("expected hidden content")
		}
		// Should not have open attribute
		if strings.Contains(output, `open=""`) || strings.Contains(output, `open=`) {
			t.Error("should not have open attribute")
		}
	})
}

func TestExampleStyled(t *testing.T) {
	t.Run("renders styled example", func(t *testing.T) {
		example := collapsible.ExampleStyled()

		var buf bytes.Buffer
		err := example.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "Package Details") {
			t.Error("expected package details header")
		}
		if !strings.Contains(output, "View package information") {
			t.Error("expected description")
		}
		if !strings.Contains(output, "Version") {
			t.Error("expected version label")
		}
		if !strings.Contains(output, "1.0.0") {
			t.Error("expected version value")
		}
		if !strings.Contains(output, "bg-muted") {
			t.Error("expected styled background")
		}
		// Check for emoji
		if !strings.Contains(output, "📦") {
			t.Error("expected package emoji")
		}
	})
}