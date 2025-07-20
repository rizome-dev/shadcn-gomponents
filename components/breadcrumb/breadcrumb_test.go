package breadcrumb_test

import (
	"bytes"
	"strings"
	"testing"

	g "maragu.dev/gomponents"
	"github.com/rizome-dev/shadcn-gomponents/components/breadcrumb"
)

func TestBreadcrumb(t *testing.T) {
	t.Run("renders nav with aria-label", func(t *testing.T) {
		b := breadcrumb.New(breadcrumb.Props{})
		
		var buf bytes.Buffer
		err := b.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "<nav") {
			t.Error("expected nav element")
		}
		if !strings.Contains(output, "aria-label=\"breadcrumb\"") {
			t.Error("expected aria-label")
		}
	})

	t.Run("renders with custom class", func(t *testing.T) {
		b := breadcrumb.New(breadcrumb.Props{Class: "custom-class"})
		
		var buf bytes.Buffer
		err := b.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "custom-class") {
			t.Error("expected custom class")
		}
	})
}

func TestBreadcrumbList(t *testing.T) {
	t.Run("renders ordered list with styling", func(t *testing.T) {
		list := breadcrumb.BreadcrumbList(breadcrumb.ListProps{}, g.Text("test"))
		
		var buf bytes.Buffer
		err := list.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "<ol") {
			t.Error("expected ordered list")
		}
		if !strings.Contains(output, "flex") {
			t.Error("expected flex display")
		}
		if !strings.Contains(output, "text-sm") {
			t.Error("expected small text")
		}
		if !strings.Contains(output, "text-muted-foreground") {
			t.Error("expected muted foreground color")
		}
	})
}

func TestBreadcrumbItem(t *testing.T) {
	t.Run("renders list item with styling", func(t *testing.T) {
		item := breadcrumb.Item(breadcrumb.ItemProps{}, g.Text("test"))
		
		var buf bytes.Buffer
		err := item.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "<li") {
			t.Error("expected list item")
		}
		if !strings.Contains(output, "inline-flex") {
			t.Error("expected inline-flex display")
		}
		if !strings.Contains(output, "items-center") {
			t.Error("expected items-center")
		}
	})
}

func TestBreadcrumbLink(t *testing.T) {
	t.Run("renders link with href and styling", func(t *testing.T) {
		link := breadcrumb.BreadcrumbLink(
			breadcrumb.LinkProps{Href: "/test"},
			g.Text("Test Link"),
		)
		
		var buf bytes.Buffer
		err := link.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "<a") {
			t.Error("expected anchor element")
		}
		if !strings.Contains(output, "href=\"/test\"") {
			t.Error("expected href attribute")
		}
		if !strings.Contains(output, "transition-colors") {
			t.Error("expected transition-colors")
		}
		if !strings.Contains(output, "hover:text-foreground") {
			t.Error("expected hover state")
		}
		if !strings.Contains(output, "Test Link") {
			t.Error("expected link text")
		}
	})
}

func TestBreadcrumbPage(t *testing.T) {
	t.Run("renders current page span with aria attributes", func(t *testing.T) {
		page := breadcrumb.Page(breadcrumb.PageProps{}, g.Text("Current"))
		
		var buf bytes.Buffer
		err := page.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "<span") {
			t.Error("expected span element")
		}
		if !strings.Contains(output, "role=\"link\"") {
			t.Error("expected role link")
		}
		if !strings.Contains(output, "aria-disabled=\"true\"") {
			t.Error("expected aria-disabled")
		}
		if !strings.Contains(output, "aria-current=\"page\"") {
			t.Error("expected aria-current")
		}
		if !strings.Contains(output, "font-normal") {
			t.Error("expected font-normal")
		}
		if !strings.Contains(output, "text-foreground") {
			t.Error("expected text-foreground")
		}
	})
}

func TestBreadcrumbSeparator(t *testing.T) {
	t.Run("renders default chevron separator", func(t *testing.T) {
		sep := breadcrumb.Separator(breadcrumb.SeparatorProps{})
		
		var buf bytes.Buffer
		err := sep.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "<li") {
			t.Error("expected list item")
		}
		if !strings.Contains(output, "role=\"presentation\"") {
			t.Error("expected presentation role")
		}
		if !strings.Contains(output, "aria-hidden=\"true\"") {
			t.Error("expected aria-hidden")
		}
		if !strings.Contains(output, "<svg") {
			t.Error("expected svg icon")
		}
		// Check for the class (HTML entities are escaped)
		if !strings.Contains(output, "[&amp;&gt;svg]:w-3.5") {
			t.Errorf("expected svg width styling, got: %s", output)
		}
	})

	t.Run("renders custom separator", func(t *testing.T) {
		sep := breadcrumb.Separator(breadcrumb.SeparatorProps{}, g.Text("/"))
		
		var buf bytes.Buffer
		err := sep.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "/") {
			t.Error("expected custom separator text")
		}
		if strings.Contains(output, "<svg") {
			t.Error("should not have svg when custom content provided")
		}
	})
}

func TestBreadcrumbEllipsis(t *testing.T) {
	t.Run("renders ellipsis with icon and sr-only text", func(t *testing.T) {
		ellipsis := breadcrumb.Ellipsis(breadcrumb.EllipsisProps{})
		
		var buf bytes.Buffer
		err := ellipsis.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "<span") {
			t.Error("expected span element")
		}
		if !strings.Contains(output, "role=\"presentation\"") {
			t.Error("expected presentation role")
		}
		if !strings.Contains(output, "aria-hidden=\"true\"") {
			t.Error("expected aria-hidden")
		}
		if !strings.Contains(output, "<svg") {
			t.Error("expected svg icon")
		}
		if !strings.Contains(output, "h-4 w-4") {
			t.Error("expected icon sizing")
		}
		if !strings.Contains(output, "sr-only") {
			t.Error("expected screen reader only text")
		}
		if !strings.Contains(output, "More") {
			t.Error("expected 'More' text")
		}
	})
}

func TestBreadcrumbExample(t *testing.T) {
	t.Run("renders complete breadcrumb example", func(t *testing.T) {
		example := breadcrumb.Example()
		
		var buf bytes.Buffer
		err := example.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "Home") {
			t.Error("expected Home link")
		}
		if !strings.Contains(output, "Documentation") {
			t.Error("expected Documentation link")
		}
		if !strings.Contains(output, "Components") {
			t.Error("expected Components page")
		}
		if !strings.Contains(output, "href=\"/\"") {
			t.Error("expected home href")
		}
		if !strings.Contains(output, "href=\"/docs\"") {
			t.Error("expected docs href")
		}
		// Should have separators
		if strings.Count(output, "role=\"presentation\"") < 2 {
			t.Error("expected at least 2 separators")
		}
	})
}

func TestBreadcrumbExampleWithDropdown(t *testing.T) {
	t.Run("renders breadcrumb with ellipsis", func(t *testing.T) {
		example := breadcrumb.ExampleWithDropdown()
		
		var buf bytes.Buffer
		err := example.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "Home") {
			t.Error("expected Home link")
		}
		if !strings.Contains(output, "Current Page") {
			t.Error("expected Current Page")
		}
		// Should have ellipsis
		if !strings.Contains(output, "More") {
			t.Error("expected ellipsis with More text")
		}
		// Ellipsis icon should be present
		if strings.Count(output, "<circle") < 3 {
			t.Error("expected 3 circles for ellipsis icon")
		}
	})
}

func TestBreadcrumbExampleCustomSeparator(t *testing.T) {
	t.Run("renders breadcrumb with custom separator", func(t *testing.T) {
		example := breadcrumb.ExampleCustomSeparator()
		
		var buf bytes.Buffer
		err := example.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "Home") {
			t.Error("expected Home link")
		}
		if !strings.Contains(output, "Products") {
			t.Error("expected Products link")
		}
		if !strings.Contains(output, "Electronics") {
			t.Error("expected Electronics page")
		}
		// Should have / separators instead of chevrons
		// Count the separators (they appear inside <li> with role="presentation")
		separatorCount := strings.Count(output, "role=\"presentation\"")
		if separatorCount < 2 {
			t.Error("expected at least 2 separators")
		}
		// Should not have SVG icons
		if strings.Contains(output, "<svg") {
			t.Error("should not have svg icons with custom separator")
		}
	})
}