package alertdialog_test

import (
	"bytes"
	"strings"
	"testing"

	g "maragu.dev/gomponents"
	"github.com/rizome-dev/shadcn-gomponents/pkg/alertdialog"
)

func TestAlertDialog(t *testing.T) {
	t.Run("renders when open", func(t *testing.T) {
		dialog := alertdialog.New(
			alertdialog.Props{Open: true},
			alertdialog.DialogContent(
				alertdialog.ContentProps{},
				g.Text("Dialog content"),
			),
		)

		var buf bytes.Buffer
		err := dialog.Render(&buf)
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
		dialog := alertdialog.New(
			alertdialog.Props{Open: false},
			alertdialog.DialogContent(
				alertdialog.ContentProps{},
				g.Text("Dialog content"),
			),
		)

		var buf bytes.Buffer
		err := dialog.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()
		if output != "" {
			t.Error("expected dialog not to render when closed")
		}
	})

	t.Run("renders overlay", func(t *testing.T) {
		overlay := alertdialog.DialogOverlay()
		var buf bytes.Buffer
		err := overlay.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "bg-background/80") {
			t.Error("expected overlay to have background color")
		}
		if !strings.Contains(output, "fixed inset-0") {
			t.Error("expected overlay to be fixed and full screen")
		}
	})

	t.Run("renders content with proper styling", func(t *testing.T) {
		content := alertdialog.DialogContent(
			alertdialog.ContentProps{},
			g.Text("Test content"),
		)
		var buf bytes.Buffer
		err := content.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "translate-x-[-50%]") {
			t.Error("expected content to be centered horizontally")
		}
		if !strings.Contains(output, "translate-y-[-50%]") {
			t.Error("expected content to be centered vertically")
		}
		if !strings.Contains(output, "max-w-lg") {
			t.Error("expected content to have max width")
		}
	})

	t.Run("renders header", func(t *testing.T) {
		header := alertdialog.DialogHeader(
			alertdialog.HeaderProps{},
			g.Text("Header content"),
		)
		var buf bytes.Buffer
		err := header.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "flex-col") {
			t.Error("expected header to be flex column")
		}
		if !strings.Contains(output, "space-y-2") {
			t.Error("expected header to have spacing")
		}
	})

	t.Run("renders footer", func(t *testing.T) {
		footer := alertdialog.DialogFooter(
			alertdialog.FooterProps{},
			g.Text("Footer content"),
		)
		var buf bytes.Buffer
		err := footer.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "sm:justify-end") {
			t.Error("expected footer to justify end on small screens")
		}
		if !strings.Contains(output, "flex-col-reverse") {
			t.Error("expected footer to reverse column on mobile")
		}
	})

	t.Run("renders title", func(t *testing.T) {
		title := alertdialog.DialogTitle(alertdialog.TitleProps{}, "Test Title")
		var buf bytes.Buffer
		err := title.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "Test Title") {
			t.Error("expected title to contain text")
		}
		if !strings.Contains(output, "text-lg") {
			t.Error("expected title to have large text")
		}
		if !strings.Contains(output, "font-semibold") {
			t.Error("expected title to be semibold")
		}
	})

	t.Run("renders description", func(t *testing.T) {
		desc := alertdialog.DialogDescription(alertdialog.DescriptionProps{}, "Test description")
		var buf bytes.Buffer
		err := desc.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "Test description") {
			t.Error("expected description to contain text")
		}
		if !strings.Contains(output, "text-sm") {
			t.Error("expected description to have small text")
		}
		if !strings.Contains(output, "text-muted-foreground") {
			t.Error("expected description to have muted color")
		}
	})

	t.Run("renders action button", func(t *testing.T) {
		action := alertdialog.DialogAction(alertdialog.ActionProps{}, g.Text("Confirm"))
		var buf bytes.Buffer
		err := action.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "Confirm") {
			t.Error("expected action to contain text")
		}
		if !strings.Contains(output, "bg-primary") {
			t.Error("expected action to have primary background")
		}
		if !strings.Contains(output, "<button") {
			t.Error("expected action to be a button")
		}
	})

	t.Run("renders action as link when href provided", func(t *testing.T) {
		action := alertdialog.DialogAction(
			alertdialog.ActionProps{Href: "/confirm"},
			g.Text("Confirm"),
		)
		var buf bytes.Buffer
		err := action.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "<a") {
			t.Error("expected action to be a link")
		}
		if !strings.Contains(output, "href=\"/confirm\"") {
			t.Error("expected action to have href")
		}
	})

	t.Run("renders cancel button", func(t *testing.T) {
		cancel := alertdialog.DialogCancel(alertdialog.CancelProps{}, g.Text("Cancel"))
		var buf bytes.Buffer
		err := cancel.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "Cancel") {
			t.Error("expected cancel to contain text")
		}
		if !strings.Contains(output, "hover:bg-accent") {
			t.Error("expected cancel to have accent hover")
		}
		if !strings.Contains(output, "mt-2 sm:mt-0") {
			t.Error("expected cancel to have responsive margin")
		}
	})

	t.Run("renders complete example", func(t *testing.T) {
		example := alertdialog.Example()
		var buf bytes.Buffer
		err := example.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}
		output := buf.String()

		if !strings.Contains(output, "Are you absolutely sure?") {
			t.Error("expected example to contain title")
		}
		if !strings.Contains(output, "This action cannot be undone") {
			t.Error("expected example to contain description")
		}
		if !strings.Contains(output, "Cancel") {
			t.Error("expected example to contain cancel button")
		}
		if !strings.Contains(output, "Continue") {
			t.Error("expected example to contain action button")
		}
	})
}