package dialog_test

import (
	"bytes"
	"strings"
	"testing"

	g "maragu.dev/gomponents"
	"github.com/rizome-dev/shadcn-gomponents/pkg/dialog"
)

func TestDialog(t *testing.T) {
	t.Run("renders when open", func(t *testing.T) {
		d := dialog.New(
			dialog.Props{Open: true},
			dialog.DialogContent(
				dialog.ContentProps{},
				g.Text("Dialog content"),
			),
		)

		var buf bytes.Buffer
		err := d.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if output == "" {
			t.Error("expected dialog to render when open")
		}
		if !strings.Contains(output, "Dialog content") {
			t.Error("expected dialog to contain content")
		}
	})

	t.Run("doesn't render when closed", func(t *testing.T) {
		d := dialog.New(
			dialog.Props{Open: false},
			dialog.DialogContent(
				dialog.ContentProps{},
				g.Text("Dialog content"),
			),
		)

		var buf bytes.Buffer
		err := d.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if output != "" {
			t.Error("expected dialog not to render when closed")
		}
	})

	t.Run("renders with custom class", func(t *testing.T) {
		d := dialog.New(
			dialog.Props{Open: true, Class: "custom-dialog"},
			g.Text("Test"),
		)

		var buf bytes.Buffer
		err := d.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "custom-dialog") {
			t.Error("expected custom class")
		}
	})
}

func TestOverlay(t *testing.T) {
	t.Run("renders overlay with default styles", func(t *testing.T) {
		overlay := dialog.Overlay()

		var buf bytes.Buffer
		err := overlay.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "bg-black/80") {
			t.Error("expected background color")
		}
		if !strings.Contains(output, "fixed inset-0") {
			t.Error("expected fixed positioning")
		}
	})

	t.Run("renders with custom class", func(t *testing.T) {
		overlay := dialog.Overlay("bg-white/50")

		var buf bytes.Buffer
		err := overlay.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "bg-white/50") {
			t.Error("expected custom background")
		}
	})
}

func TestContent(t *testing.T) {
	t.Run("renders content with positioning", func(t *testing.T) {
		content := dialog.DialogContent(
			dialog.ContentProps{},
			g.Text("Test content"),
		)

		var buf bytes.Buffer
		err := content.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "translate-x-[-50%]") {
			t.Error("expected horizontal centering")
		}
		if !strings.Contains(output, "translate-y-[-50%]") {
			t.Error("expected vertical centering")
		}
		if !strings.Contains(output, "Test content") {
			t.Error("expected content text")
		}
	})

	t.Run("renders with close button when requested", func(t *testing.T) {
		content := dialog.DialogContent(
			dialog.ContentProps{ShowCloseButton: true},
			g.Text("Test content"),
		)

		var buf bytes.Buffer
		err := content.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "<button") {
			t.Error("expected close button")
		}
		if !strings.Contains(output, "absolute right-4 top-4") {
			t.Error("expected close button positioning")
		}
		if !strings.Contains(output, "<svg") {
			t.Error("expected close icon")
		}
		if !strings.Contains(output, "Close") {
			t.Error("expected screen reader text")
		}
	})

	t.Run("renders without close button by default", func(t *testing.T) {
		content := dialog.DialogContent(
			dialog.ContentProps{},
			g.Text("Test content"),
		)

		var buf bytes.Buffer
		err := content.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if strings.Contains(output, "absolute right-4 top-4") {
			t.Error("should not have close button by default")
		}
	})
}

func TestHeader(t *testing.T) {
	t.Run("renders header with styling", func(t *testing.T) {
		header := dialog.DialogHeader(
			dialog.HeaderProps{},
			g.Text("Header content"),
		)

		var buf bytes.Buffer
		err := header.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "flex-col") {
			t.Error("expected flex column")
		}
		if !strings.Contains(output, "space-y-1.5") {
			t.Error("expected spacing")
		}
		if !strings.Contains(output, "Header content") {
			t.Error("expected header content")
		}
	})
}

func TestFooter(t *testing.T) {
	t.Run("renders footer with responsive layout", func(t *testing.T) {
		footer := dialog.DialogFooter(
			dialog.FooterProps{},
			g.Text("Footer content"),
		)

		var buf bytes.Buffer
		err := footer.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "flex-col-reverse") {
			t.Error("expected reverse column on mobile")
		}
		if !strings.Contains(output, "sm:flex-row") {
			t.Error("expected row on desktop")
		}
		if !strings.Contains(output, "sm:justify-end") {
			t.Error("expected end justification")
		}
	})
}

func TestTitle(t *testing.T) {
	t.Run("renders title with styling", func(t *testing.T) {
		title := dialog.DialogTitle(dialog.TitleProps{}, "Test Title")

		var buf bytes.Buffer
		err := title.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "<h2") {
			t.Error("expected h2 element")
		}
		if !strings.Contains(output, "Test Title") {
			t.Error("expected title text")
		}
		if !strings.Contains(output, "text-lg") {
			t.Error("expected large text")
		}
		if !strings.Contains(output, "font-semibold") {
			t.Error("expected semibold font")
		}
		if !strings.Contains(output, "tracking-tight") {
			t.Error("expected tight tracking")
		}
	})
}

func TestDescription(t *testing.T) {
	t.Run("renders description with styling", func(t *testing.T) {
		desc := dialog.Description(dialog.DescriptionProps{}, "Test description")

		var buf bytes.Buffer
		err := desc.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "<p") {
			t.Error("expected paragraph element")
		}
		if !strings.Contains(output, "Test description") {
			t.Error("expected description text")
		}
		if !strings.Contains(output, "text-sm") {
			t.Error("expected small text")
		}
		if !strings.Contains(output, "text-muted-foreground") {
			t.Error("expected muted color")
		}
	})
}

func TestTrigger(t *testing.T) {
	t.Run("renders trigger button", func(t *testing.T) {
		trigger := dialog.Trigger(
			dialog.TriggerProps{Class: "custom-trigger"},
			g.Text("Open Dialog"),
		)

		var buf bytes.Buffer
		err := trigger.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "<button") {
			t.Error("expected button element")
		}
		if !strings.Contains(output, "Open Dialog") {
			t.Error("expected button text")
		}
		if !strings.Contains(output, "custom-trigger") {
			t.Error("expected custom class")
		}
	})
}

func TestClose(t *testing.T) {
	t.Run("renders close button", func(t *testing.T) {
		close := dialog.Close(
			dialog.CloseProps{Class: "custom-close"},
			g.Text("Cancel"),
		)

		var buf bytes.Buffer
		err := close.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "<button") {
			t.Error("expected button element")
		}
		if !strings.Contains(output, "Cancel") {
			t.Error("expected button text")
		}
		if !strings.Contains(output, "custom-close") {
			t.Error("expected custom class")
		}
	})
}

func TestExample(t *testing.T) {
	t.Run("renders complete example", func(t *testing.T) {
		example := dialog.Example()

		var buf bytes.Buffer
		err := example.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "Edit Profile") {
			t.Error("expected title")
		}
		if !strings.Contains(output, "Make changes to your profile") {
			t.Error("expected description")
		}
		if !strings.Contains(output, "Name") {
			t.Error("expected name label")
		}
		if !strings.Contains(output, "Pedro Duarte") {
			t.Error("expected name value")
		}
		if !strings.Contains(output, "Username") {
			t.Error("expected username label")
		}
		if !strings.Contains(output, "@peduarte") {
			t.Error("expected username value")
		}
		if !strings.Contains(output, "Save changes") {
			t.Error("expected save button")
		}
		// Should have close button
		if !strings.Contains(output, "absolute right-4 top-4") {
			t.Error("expected close button")
		}
	})
}

func TestExampleScrollable(t *testing.T) {
	t.Run("renders scrollable example", func(t *testing.T) {
		example := dialog.ExampleScrollable()

		var buf bytes.Buffer
		err := example.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "Terms of Service") {
			t.Error("expected title")
		}
		if !strings.Contains(output, "overflow-y-auto") {
			t.Error("expected scrollable container")
		}
		if !strings.Contains(output, "max-h-[60vh]") {
			t.Error("expected max height")
		}
		if !strings.Contains(output, "1. Introduction") {
			t.Error("expected section 1")
		}
		if !strings.Contains(output, "5. Disclaimers") {
			t.Error("expected section 5")
		}
		if !strings.Contains(output, "Accept") {
			t.Error("expected accept button")
		}
		if !strings.Contains(output, "Decline") {
			t.Error("expected decline button")
		}
	})
}

func TestExampleCustom(t *testing.T) {
	t.Run("renders custom styled example", func(t *testing.T) {
		example := dialog.ExampleCustom()

		var buf bytes.Buffer
		err := example.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "Payment Successful") {
			t.Error("expected title")
		}
		if !strings.Contains(output, "bg-slate-950") {
			t.Error("expected dark background")
		}
		if !strings.Contains(output, "text-slate-50") {
			t.Error("expected light text")
		}
		if !strings.Contains(output, "$99.00") {
			t.Error("expected amount")
		}
		if !strings.Contains(output, "TXN-20240115-001") {
			t.Error("expected reference")
		}
		// Check for checkmark icon
		if !strings.Contains(output, "M5 13l4 4L19 7") {
			t.Error("expected checkmark icon path")
		}
	})
}