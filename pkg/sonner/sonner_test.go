package sonner

import (
	"bytes"
	"strings"
	"testing"

	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func TestToaster(t *testing.T) {
	tests := []struct {
		name     string
		props    ToasterProps
		contains []string
	}{
		{
			name:  "default toaster",
			props: ToasterProps{},
			contains: []string{
				`id="toaster"`,
				`data-toaster="true"`,
				`data-position="top-right"`,
				`data-duration="5000"`,
				`data-gap="16"`,
				`data-max-visible="3"`,
				`top-4 right-4`,
				`--gap: 16px`,
			},
		},
		{
			name: "bottom left toaster with custom settings",
			props: ToasterProps{
				Position:    PositionBottomLeft,
				Duration:    3000,
				Gap:         24,
				MaxVisible:  5,
				Expand:      true,
				RichColors:  true,
				CloseButton: true,
			},
			contains: []string{
				`data-position="bottom-left"`,
				`data-duration="3000"`,
				`data-gap="24"`,
				`data-max-visible="5"`,
				`data-expand="true"`,
				`data-rich-colors="true"`,
				`data-close-button="true"`,
				`bottom-4 left-4`,
				`--gap: 24px`,
			},
		},
		{
			name: "center positions",
			props: ToasterProps{
				Position: PositionTopCenter,
			},
			contains: []string{
				`data-position="top-center"`,
				`top-4 left-1/2 -translate-x-1/2`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			component := Toaster(tt.props)
			err := component.Render(&buf)
			if err != nil {
				t.Fatalf("Render() error = %v", err)
			}

			result := buf.String()
			for _, expected := range tt.contains {
				if !strings.Contains(result, expected) {
					t.Errorf("Expected output to contain %q, but it didn't. Got:\n%s", expected, result)
				}
			}
		})
	}
}

func TestToast(t *testing.T) {
	tests := []struct {
		name     string
		props    ToastProps
		contains []string
	}{
		{
			name: "default toast",
			props: ToastProps{
				ID:          "test-toast",
				Title:       "Hello",
				Description: "This is a toast",
			},
			contains: []string{
				`id="test-toast"`,
				`data-toast="true"`,
				`data-type="default"`,
				`>Hello</div>`,
				`>This is a toast</div>`,
				`bg-popover`,
			},
		},
		{
			name: "success toast with close button",
			props: ToastProps{
				Type:        ToastSuccess,
				Title:       "Success",
				Description: "Operation completed",
				CloseButton: true,
				Duration:    3000,
			},
			contains: []string{
				`data-type="success"`,
				`data-duration="3000"`,
				`>Success</div>`,
				`>Operation completed</div>`,
				`data-toast-close`,
				`bg-green-50`,
				// Should have success icon
				`viewBox="0 0 20 20"`,
			},
		},
		{
			name: "error toast",
			props: ToastProps{
				Type:        ToastError,
				Title:       "Error",
				Description: "Something went wrong",
			},
			contains: []string{
				`data-type="error"`,
				`>Error</div>`,
				`bg-red-50`,
			},
		},
		{
			name: "toast with action",
			props: ToastProps{
				Title:       "Update Available",
				Description: "A new version is ready",
				Action: &ToastAction{
					Label:   "Update Now",
					OnClick: `updateApp()`,
				},
			},
			contains: []string{
				`>Update Available</div>`,
				`>Update Now</button>`,
				`onclick="updateApp()"`,
			},
		},
		{
			name: "toast with custom class",
			props: ToastProps{
				Description: "Custom styled toast",
				Class:       "custom-toast-class",
			},
			contains: []string{
				`custom-toast-class`,
				`>Custom styled toast</div>`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			component := Toast(tt.props)
			err := component.Render(&buf)
			if err != nil {
				t.Fatalf("Render() error = %v", err)
			}

			result := buf.String()
			for _, expected := range tt.contains {
				if !strings.Contains(result, expected) {
					t.Errorf("Expected output to contain %q, but it didn't. Got:\n%s", expected, result)
				}
			}
		})
	}
}

func TestHelperFunctions(t *testing.T) {
	tests := []struct {
		name      string
		component g.Node
		contains  []string
	}{
		{
			name:      "success helper",
			component: Success("Great!", "Your changes have been saved"),
			contains: []string{
				`data-type="success"`,
				`>Great!</div>`,
				`>Your changes have been saved</div>`,
				`data-toast-close`,
			},
		},
		{
			name:      "error helper",
			component: Error("Oops!", "Something went wrong"),
			contains: []string{
				`data-type="error"`,
				`>Oops!</div>`,
				`>Something went wrong</div>`,
			},
		},
		{
			name:      "warning helper",
			component: Warning("Careful", "This action cannot be undone"),
			contains: []string{
				`data-type="warning"`,
				`>Careful</div>`,
				`>This action cannot be undone</div>`,
			},
		},
		{
			name:      "info helper",
			component: Info("FYI", "New features available"),
			contains: []string{
				`data-type="info"`,
				`>FYI</div>`,
				`>New features available</div>`,
			},
		},
		{
			name:      "message helper",
			component: Message("Quick notification"),
			contains: []string{
				`data-type="default"`,
				`>Quick notification</div>`,
			},
		},
		{
			name: "with action helper",
			component: WithAction(ToastProps{
				Title:       "Confirm",
				Description: "Are you sure?",
			}, "Yes", `confirm()`),
			contains: []string{
				`>Confirm</div>`,
				`>Are you sure?</div>`,
				`>Yes</button>`,
				`onclick="confirm()"`,
			},
		},
		{
			name:      "loading helper",
			component: LoadingToast("Processing..."),
			contains: []string{
				`data-type="default"`,
				`>Processing...</div>`,
				`data-duration="0"`, // Infinite duration
				`pr-12`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := tt.component.Render(&buf)
			if err != nil {
				t.Fatalf("Render() error = %v", err)
			}

			result := buf.String()
			for _, expected := range tt.contains {
				if !strings.Contains(result, expected) {
					t.Errorf("Expected output to contain %q, but it didn't. Got:\n%s", expected, result)
				}
			}
		})
	}
}

func TestPromiseToasts(t *testing.T) {
	promise := PromiseToast{
		ID:      "promise-1",
		Loading: "Saving changes...",
		Success: "Changes saved successfully!",
		Error:   "Failed to save changes",
	}

	tests := []struct {
		name      string
		component g.Node
		contains  []string
	}{
		{
			name:      "promise loading",
			component: PromiseLoading(promise),
			contains: []string{
				`id="promise-1"`,
				`data-type="default"`,
				`>Saving changes...</div>`,
				`data-duration="0"`,
			},
		},
		{
			name:      "promise success",
			component: PromiseSuccess(promise),
			contains: []string{
				`id="promise-1"`,
				`data-type="success"`,
				`>Success</div>`,
				`>Changes saved successfully!</div>`,
			},
		},
		{
			name:      "promise error",
			component: PromiseError(promise),
			contains: []string{
				`id="promise-1"`,
				`data-type="error"`,
				`>Error</div>`,
				`>Failed to save changes</div>`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := tt.component.Render(&buf)
			if err != nil {
				t.Fatalf("Render() error = %v", err)
			}

			result := buf.String()
			for _, expected := range tt.contains {
				if !strings.Contains(result, expected) {
					t.Errorf("Expected output to contain %q, but it didn't. Got:\n%s", expected, result)
				}
			}
		})
	}
}

func TestCustomToast(t *testing.T) {
	customContent := Div(
		Class("flex items-center gap-4 p-4"),
		Img(Src("/avatar.jpg"), Alt("User"), Class("w-10 h-10 rounded-full")),
		Div(
			P(Class("font-semibold"), g.Text("John Doe")),
			P(Class("text-sm text-muted-foreground"), g.Text("Just sent you a message")),
		),
	)

	var buf bytes.Buffer
	component := Custom(customContent)
	err := component.Render(&buf)
	if err != nil {
		t.Fatalf("Render() error = %v", err)
	}

	result := buf.String()
	expected := []string{
		`data-toast="true"`,
		`data-type="custom"`,
		`id="toast-`,
		`<img src="/avatar.jpg"`,
		`>John Doe</p>`,
		`>Just sent you a message</p>`,
	}

	for _, exp := range expected {
		if !strings.Contains(result, exp) {
			t.Errorf("Expected output to contain %q, but it didn't. Got:\n%s", exp, result)
		}
	}
}