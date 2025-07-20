package toast_test

import (
	"strings"
	"testing"
	"time"
	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/components/toast"
)

func renderToString(node g.Node) string {
	var buf strings.Builder
	node.Render(&buf)
	return buf.String()
}

func TestToast(t *testing.T) {
	tests := []struct {
		name     string
		toast    g.Node
		contains []string
	}{
		{
			name: "simple toast",
			toast: toast.Simple("Hello, World!"),
			contains: []string{
				`role="alert"`,
				`aria-live="polite"`,
				`data-toast="true"`,
				`data-variant="default"`,
				`Hello, World!`,
				`toast-close`,
				`aria-label="Close"`,
			},
		},
		{
			name: "toast with title and description",
			toast: toast.New(toast.Props{
				Title:       "Success",
				Description: "Your changes have been saved.",
				Variant:     toast.VariantSuccess,
			}),
			contains: []string{
				`data-variant="success"`,
				`text-sm font-semibold`,
				`Success`,
				`text-sm opacity-90`,
				`Your changes have been saved.`,
				`border-green-200 bg-green-50 text-green-900`,
			},
		},
		{
			name: "error toast",
			toast: toast.Error("Error", "Something went wrong!"),
			contains: []string{
				`data-variant="error"`,
				`Error`,
				`Something went wrong!`,
				`border-red-200 bg-red-50 text-red-900`,
				`data-duration="10000"`, // 10 seconds
			},
		},
		{
			name: "warning toast",
			toast: toast.Warning("Warning", "Please check your input"),
			contains: []string{
				`data-variant="warning"`,
				`Warning`,
				`Please check your input`,
				`border-yellow-200 bg-yellow-50 text-yellow-900`,
			},
		},
		{
			name: "info toast",
			toast: toast.Info("Info", "New update available"),
			contains: []string{
				`data-variant="info"`,
				`Info`,
				`New update available`,
				`border-blue-200 bg-blue-50 text-blue-900`,
			},
		},
		{
			name: "loading toast",
			toast: toast.LoadingToast("Processing your request..."),
			contains: []string{
				`Processing your request...`,
				`animate-spin`, // Loading spinner
			},
		},
		{
			name: "toast with action",
			toast: toast.WithAction(
				"File deleted",
				"The file has been permanently deleted.",
				"Undo",
				"undoDelete()",
			),
			contains: []string{
				`File deleted`,
				`The file has been permanently deleted.`,
				`onclick="undoDelete()"`,
				`Undo`,
				`bg-primary text-primary-foreground`,
			},
		},
		{
			name: "toast with custom duration",
			toast: toast.New(toast.Props{
				Description: "Custom duration",
				Duration:    3 * time.Second,
			}),
			contains: []string{
				`Custom duration`,
				`data-duration="3000"`,
			},
		},
		{
			name: "persistent toast",
			toast: toast.New(toast.Props{
				Description: "Persistent toast",
				Duration:    0,
				Closable:    true,
			}),
			contains: []string{
				`Persistent toast`,
				`toast-close`,
				// Should not contain data-duration
			},
		},
		{
			name: "toast with custom icon",
			toast: toast.New(toast.Props{
				Description: "Custom icon",
				Icon:        Span(Class("custom-icon"), g.Text("🎉")),
			}),
			contains: []string{
				`Custom icon`,
				`custom-icon`,
				`🎉`,
			},
		},
		{
			name: "toast with progress",
			toast: toast.New(toast.Props{
				Description: "Progress toast",
				Progress:    true,
				Duration:    5 * time.Second,
			}),
			contains: []string{
				`Progress toast`,
				`animation: toast-progress 5000ms linear`,
				`absolute bottom-0 left-0 h-1`,
			},
		},
		{
			name: "destructive variant",
			toast: toast.New(toast.Props{
				Title:   "Destructive",
				Variant: toast.VariantDestructive,
			}),
			contains: []string{
				`data-variant="destructive"`,
				`border-red-200 bg-red-50 text-red-900`,
			},
		},
		{
			name: "promise toast",
			toast: toast.Promise("promise-123", "Uploading file..."),
			contains: []string{
				`id="promise-123"`,
				`Uploading file...`,
				`animate-spin`, // Loading icon
			},
		},
		{
			name: "toast with custom class",
			toast: toast.New(toast.Props{
				Description: "Custom class",
				Class:       "custom-toast-class",
			}),
			contains: []string{
				`Custom class`,
				`custom-toast-class`,
			},
		},
		{
			name: "success toast with default icon",
			toast: toast.Success("Success", "Operation completed"),
			contains: []string{
				`data-variant="success"`,
				`<svg class="h-5 w-5 text-green-600"`, // Success icon
				`Operation completed`,
			},
		},
		{
			name: "error toast with default icon",
			toast: toast.Error("Error", "Operation failed"),
			contains: []string{
				`data-variant="error"`,
				`<svg class="h-5 w-5 text-red-600"`, // Error icon
				`Operation failed`,
			},
		},
		{
			name: "warning toast with default icon",
			toast: toast.Warning("Warning", "Be careful"),
			contains: []string{
				`data-variant="warning"`,
				`<svg class="h-5 w-5 text-yellow-600"`, // Warning icon
				`Be careful`,
			},
		},
		{
			name: "info toast with default icon",
			toast: toast.Info("Info", "For your information"),
			contains: []string{
				`data-variant="info"`,
				`<svg class="h-5 w-5 text-blue-600"`, // Info icon
				`For your information`,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := renderToString(test.toast)
			for _, expected := range test.contains {
				if !strings.Contains(result, expected) {
					t.Errorf("expected result to contain %q, but it didn't.\nGot: %s", expected, result)
				}
			}
		})
	}
}

func TestToaster(t *testing.T) {
	tests := []struct {
		name     string
		toaster  g.Node
		contains []string
	}{
		{
			name: "default toaster",
			toaster: toast.Toaster(toast.ToasterProps{}),
			contains: []string{
				`id="toaster"`,
				`data-toaster="true"`,
				`data-position="bottom-right"`,
				`data-max-toasts="3"`,
				`bottom-4 right-4`,
				`fixed z-50 flex flex-col gap-2`,
			},
		},
		{
			name: "toaster with custom position",
			toaster: toast.Toaster(toast.ToasterProps{
				Position: toast.PositionTopCenter,
			}),
			contains: []string{
				`data-position="top-center"`,
				`top-4 left-1/2 -translate-x-1/2`,
			},
		},
		{
			name: "toaster with custom max toasts",
			toaster: toast.Toaster(toast.ToasterProps{
				MaxToast: 5,
			}),
			contains: []string{
				`data-max-toasts="5"`,
			},
		},
		{
			name: "toaster with custom ID",
			toaster: toast.Toaster(toast.ToasterProps{
				ID: "custom-toaster",
			}),
			contains: []string{
				`id="custom-toaster"`,
			},
		},
		{
			name: "toaster with custom class",
			toaster: toast.Toaster(toast.ToasterProps{
				Class: "custom-toaster-class",
			}),
			contains: []string{
				`custom-toaster-class`,
			},
		},
		{
			name: "toaster positions",
			toaster: Div(
				toast.Toaster(toast.ToasterProps{Position: toast.PositionTopLeft}),
				toast.Toaster(toast.ToasterProps{Position: toast.PositionTopCenter}),
				toast.Toaster(toast.ToasterProps{Position: toast.PositionTopRight}),
				toast.Toaster(toast.ToasterProps{Position: toast.PositionBottomLeft}),
				toast.Toaster(toast.ToasterProps{Position: toast.PositionBottomCenter}),
				toast.Toaster(toast.ToasterProps{Position: toast.PositionBottomRight}),
			),
			contains: []string{
				`top-4 left-4`,
				`top-4 left-1/2 -translate-x-1/2`,
				`top-4 right-4`,
				`bottom-4 left-4`,
				`bottom-4 left-1/2 -translate-x-1/2`,
				`bottom-4 right-4`,
			},
		},
		{
			name: "toaster with CSS",
			toaster: toast.Toaster(toast.ToasterProps{}),
			contains: []string{
				`@keyframes toast-progress`,
				`from { width: 100%; }`,
				`to { width: 0%; }`,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := renderToString(test.toaster)
			for _, expected := range test.contains {
				if !strings.Contains(result, expected) {
					t.Errorf("expected result to contain %q, but it didn't.\nGot: %s", expected, result)
				}
			}
		})
	}
}

func TestToastHelpers(t *testing.T) {
	t.Run("promise success script", func(t *testing.T) {
		script := toast.PromiseSuccess("test-id", "Upload complete!")
		expected := []string{
			`document.getElementById('test-id')`,
			`Upload complete!`,
			`data-variant', 'success'`,
			`setTimeout(() => toast.remove(), 300)`,
		}
		
		for _, exp := range expected {
			if !strings.Contains(script, exp) {
				t.Errorf("expected script to contain %q, but it didn't.\nGot: %s", exp, script)
			}
		}
	})

	t.Run("promise error script", func(t *testing.T) {
		script := toast.PromiseError("test-id", "Upload failed!")
		expected := []string{
			`document.getElementById('test-id')`,
			`Upload failed!`,
			`data-variant', 'error'`,
			`toast-close`,
		}
		
		for _, exp := range expected {
			if !strings.Contains(script, exp) {
				t.Errorf("expected script to contain %q, but it didn't.\nGot: %s", exp, script)
			}
		}
	})
}

func TestToastShouldNotContain(t *testing.T) {
	tests := []struct {
		name         string
		toast        g.Node
		shouldNotHave []string
	}{
		{
			name: "persistent toast should not have duration",
			toast: toast.New(toast.Props{
				Description: "Persistent",
				Duration:    0,
				Closable:    true,
			}),
			shouldNotHave: []string{
				`data-duration=`,
			},
		},
		{
			name: "non-closable toast should not have close button",
			toast: toast.New(toast.Props{
				Description: "Not closable",
				Closable:    false,
				Duration:    5 * time.Second,
			}),
			shouldNotHave: []string{
				`toast-close`,
				`aria-label="Close"`,
			},
		},
		{
			name: "toast without progress should not have progress bar",
			toast: toast.New(toast.Props{
				Description: "No progress",
				Progress:    false,
			}),
			shouldNotHave: []string{
				`animation: toast-progress`,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := renderToString(test.toast)
			for _, notExpected := range test.shouldNotHave {
				if strings.Contains(result, notExpected) {
					t.Errorf("expected result NOT to contain %q, but it did.\nGot: %s", notExpected, result)
				}
			}
		})
	}
}