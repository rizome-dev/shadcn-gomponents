package sonner

import (
	"fmt"
	"time"

	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// ToastType defines the type of toast
type ToastType string

const (
	ToastDefault ToastType = "default"
	ToastSuccess ToastType = "success"
	ToastError   ToastType = "error"
	ToastWarning ToastType = "warning"
	ToastInfo    ToastType = "info"
)

// Position defines where toasts appear
type Position string

const (
	PositionTopRight    Position = "top-right"
	PositionTopCenter   Position = "top-center"
	PositionTopLeft     Position = "top-left"
	PositionBottomRight Position = "bottom-right"
	PositionBottomCenter Position = "bottom-center"
	PositionBottomLeft  Position = "bottom-left"
)

// ToasterProps defines properties for the toaster container
type ToasterProps struct {
	Position      Position
	Expand        bool
	RichColors    bool
	CloseButton   bool
	Duration      int // milliseconds
	Gap           int // pixels between toasts
	MaxVisible    int
	Class         string
	ID            string
}

// ToastProps defines properties for individual toasts
type ToastProps struct {
	ID          string
	Type        ToastType
	Title       string
	Description string
	Action      *ToastAction
	CloseButton bool
	Duration    int // milliseconds, 0 = infinite
	Class       string
}

// ToastAction defines an action button for toasts
type ToastAction struct {
	Label   string
	OnClick string // JavaScript or HTMX attributes
}

// Toaster creates a toast container
func Toaster(props ToasterProps) g.Node {
	// Set defaults
	if props.Position == "" {
		props.Position = PositionTopRight
	}
	if props.Duration == 0 {
		props.Duration = 5000
	}
	if props.Gap == 0 {
		props.Gap = 16
	}
	if props.MaxVisible == 0 {
		props.MaxVisible = 3
	}
	if props.ID == "" {
		props.ID = "toaster"
	}

	// Position classes
	positionClasses := map[Position]string{
		PositionTopRight:     "top-4 right-4",
		PositionTopCenter:    "top-4 left-1/2 -translate-x-1/2",
		PositionTopLeft:      "top-4 left-4",
		PositionBottomRight:  "bottom-4 right-4",
		PositionBottomCenter: "bottom-4 left-1/2 -translate-x-1/2",
		PositionBottomLeft:   "bottom-4 left-4",
	}

	return html.Div(
		g.Attr("id", props.ID),
		g.Attr("data-toaster", "true"),
		g.Attr("data-position", string(props.Position)),
		g.Attr("data-duration", fmt.Sprintf("%d", props.Duration)),
		g.Attr("data-gap", fmt.Sprintf("%d", props.Gap)),
		g.Attr("data-max-visible", fmt.Sprintf("%d", props.MaxVisible)),
		g.If(props.Expand, g.Attr("data-expand", "true")),
		g.If(props.RichColors, g.Attr("data-rich-colors", "true")),
		g.If(props.CloseButton, g.Attr("data-close-button", "true")),
		html.Class(lib.CN(
			"toaster group fixed z-[100] max-h-screen w-full max-w-[420px] pointer-events-none",
			positionClasses[props.Position],
			props.Class,
		)),
		g.Attr("style", fmt.Sprintf("--gap: %dpx", props.Gap)),
		// Toast items will be appended here dynamically
	html.Ol(
			html.Class("flex flex-col gap-[var(--gap)]"),
			g.Attr("data-toaster-list", "true"),
			// Toasts are inserted here
		),
	)
}

// Toast creates an individual toast notification
func Toast(props ToastProps) g.Node {
	// Set defaults
	if props.Type == "" {
		props.Type = ToastDefault
	}

	// Type-based styling
	typeClasses := map[ToastType]string{
		ToastDefault: "bg-popover text-popover-foreground border-border",
		ToastSuccess: "bg-green-50 dark:bg-green-900/20 text-green-900 dark:text-green-100 border-green-200 dark:border-green-800",
		ToastError:   "bg-red-50 dark:bg-red-900/20 text-red-900 dark:text-red-100 border-red-200 dark:border-red-800",
		ToastWarning: "bg-yellow-50 dark:bg-yellow-900/20 text-yellow-900 dark:text-yellow-100 border-yellow-200 dark:border-yellow-800",
		ToastInfo:    "bg-blue-50 dark:bg-blue-900/20 text-blue-900 dark:text-blue-100 border-blue-200 dark:border-blue-800",
	}

	// Icon based on type
	icons := map[ToastType]g.Node{
		ToastSuccess: g.Raw(`<svg width="20" height="20" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg" class="text-green-600 dark:text-green-400"><path fill-rule="evenodd" clip-rule="evenodd" d="M10 18C14.4183 18 18 14.4183 18 10C18 5.58172 14.4183 2 10 2C5.58172 2 2 5.58172 2 10C2 14.4183 5.58172 18 10 18ZM13.7071 8.70711C14.0976 8.31658 14.0976 7.68342 13.7071 7.29289C13.3166 6.90237 12.6834 6.90237 12.2929 7.29289L9 10.5858L7.70711 9.29289C7.31658 8.90237 6.68342 8.90237 6.29289 9.29289C5.90237 9.68342 5.90237 10.3166 6.29289 10.7071L8.29289 12.7071C8.68342 13.0976 9.31658 13.0976 9.70711 12.7071L13.7071 8.70711Z" fill="currentColor"/></svg>`),
		ToastError:   g.Raw(`<svg width="20" height="20" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg" class="text-red-600 dark:text-red-400"><path fill-rule="evenodd" clip-rule="evenodd" d="M10 18C14.4183 18 18 14.4183 18 10C18 5.58172 14.4183 2 10 2C5.58172 2 2 5.58172 2 10C2 14.4183 5.58172 18 10 18ZM8.28033 7.21967C7.98744 6.92678 7.51256 6.92678 7.21967 7.21967C6.92678 7.51256 6.92678 7.98744 7.21967 8.28033L8.93934 10L7.21967 11.7197C6.92678 12.0126 6.92678 12.4874 7.21967 12.7803C7.51256 13.0732 7.98744 13.0732 8.28033 12.7803L10 11.0607L11.7197 12.7803C12.0126 13.0732 12.4874 13.0732 12.7803 12.7803C13.0732 12.4874 13.0732 12.0126 12.7803 11.7197L11.0607 10L12.7803 8.28033C13.0732 7.98744 13.0732 7.51256 12.7803 7.21967C12.4874 6.92678 12.0126 6.92678 11.7197 7.21967L10 8.93934L8.28033 7.21967Z" fill="currentColor"/></svg>`),
		ToastWarning: g.Raw(`<svg width="20" height="20" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg" class="text-yellow-600 dark:text-yellow-400"><path fill-rule="evenodd" clip-rule="evenodd" d="M8.48528 3.44632C9.0843 2.24737 10.9157 2.24737 11.5147 3.44632L17.3229 14.8018C17.9003 15.9568 17.0141 17.3333 15.8082 17.3333H4.19179C2.98587 17.3333 2.09969 15.9568 2.67712 14.8018L8.48528 3.44632ZM10 7C10.5523 7 11 7.44772 11 8V11C11 11.5523 10.5523 12 10 12C9.44772 12 9 11.5523 9 11V8C9 7.44772 9.44772 7 10 7ZM10 15C10.5523 15 11 14.5523 11 14C11 13.4477 10.5523 13 10 13C9.44772 13 9 13.4477 9 14C9 14.5523 9.44772 15 10 15Z" fill="currentColor"/></svg>`),
		ToastInfo:    g.Raw(`<svg width="20" height="20" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg" class="text-blue-600 dark:text-blue-400"><path fill-rule="evenodd" clip-rule="evenodd" d="M18 10C18 14.4183 14.4183 18 10 18C5.58172 18 2 14.4183 2 10C2 5.58172 5.58172 2 10 2C14.4183 2 18 5.58172 18 10ZM11 6C11 6.55228 10.5523 7 10 7C9.44772 7 9 6.55228 9 6C9 5.44772 9.44772 5 10 5C10.5523 5 11 5.44772 11 6ZM9 9C9 8.44772 9.44772 8 10 8C10.5523 8 11 8.44772 11 9V14C11 14.5523 10.5523 15 10 15C9.44772 15 9 14.5523 9 14V9Z" fill="currentColor"/></svg>`),
	}

	// Build toast content
	toastContent := []g.Node{}

	// Icon
	if icon, hasIcon := icons[props.Type]; hasIcon && props.Type != ToastDefault {
		toastContent = append(toastContent, html.Div(html.Class("shrink-0"), icon))
	}

	// Content
	contentItems := []g.Node{}
	if props.Title != "" {
		contentItems = append(contentItems,
			html.Div(html.Class("font-semibold"), g.Text(props.Title)),
		)
	}
	if props.Description != "" {
		contentItems = append(contentItems,
			html.Div(html.Class("text-sm opacity-90"), g.Text(props.Description)),
		)
	}
	toastContent = append(toastContent, html.Div(html.Class("flex-1 space-y-1"), g.Group(contentItems)))

	// Action button
	if props.Action != nil {
		toastContent = append(toastContent,
			html.Button(
				html.Type("button"),
				html.Class("shrink-0 rounded-md px-3 py-1.5 text-sm font-medium ring-offset-background transition-opacity hover:opacity-90 focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 disabled:pointer-events-none disabled:opacity-50"),
				g.If(props.Type == ToastDefault, html.Class("bg-primary text-primary-foreground")),
				g.If(props.Type != ToastDefault, html.Class("bg-current/10 text-current")),
				g.Attr("onclick", props.Action.OnClick),
				g.Text(props.Action.Label),
			),
		)
	}

	// Close button
	if props.CloseButton {
		toastContent = append(toastContent,
			html.Button(
				html.Type("button"),
				html.Class("shrink-0 rounded-md p-1 transition-opacity hover:opacity-60 focus:opacity-100 focus:outline-none focus:ring-2 focus:ring-ring"),
				g.Attr("data-toast-close", props.ID),
				g.Raw(`<svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg" class="h-4 w-4"><path d="M11.7816 4.03157C12.0062 3.80702 12.0062 3.44295 11.7816 3.2184C11.5571 2.99385 11.193 2.99385 10.9685 3.2184L7.50005 6.68682L4.03164 3.2184C3.80708 2.99385 3.44301 2.99385 3.21846 3.2184C2.99391 3.44295 2.99391 3.80702 3.21846 4.03157L6.68688 7.49999L3.21846 10.9684C2.99391 11.193 2.99391 11.557 3.21846 11.7816C3.44301 12.0061 3.80708 12.0061 4.03164 11.7816L7.50005 8.31316L10.9685 11.7816C11.193 12.0061 11.5571 12.0061 11.7816 11.7816C12.0062 11.557 12.0062 11.193 11.7816 10.9684L8.31322 7.49999L11.7816 4.03157Z" fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"></path></svg>`),
			),
		)
	}

	return html.Li(
		g.Attr("id", props.ID),
		g.Attr("data-toast", "true"),
		g.Attr("data-type", string(props.Type)),
		g.If(props.Duration >= 0, g.Attr("data-duration", fmt.Sprintf("%d", props.Duration))),
		html.Class(lib.CN(
			"pointer-events-auto relative flex w-full items-center gap-3 overflow-hidden rounded-lg border p-4 pr-6 shadow-lg transition-all",
			"data-[swipe=cancel]:translate-x-0 data-[swipe=end]:translate-x-[var(--radix-toast-swipe-end-x)] data-[swipe=move]:translate-x-[var(--radix-toast-swipe-move-x)] data-[swipe=move]:transition-none",
			"data-[state=open]:animate-in data-[state=closed]:animate-out data-[swipe=end]:animate-out data-[state=closed]:fade-out-80 data-[state=closed]:slide-out-to-right-full data-[state=open]:slide-in-from-top-full data-[state=open]:sm:slide-in-from-bottom-full",
			typeClasses[props.Type],
			props.Class,
		)),
		g.Attr("role", "status"),
		g.Attr("aria-live", "polite"),
		g.Group(toastContent),
	)
}

// Helper functions to create toasts of specific types

// Success creates a success toast
func Success(title, description string) g.Node {
	return Toast(ToastProps{
		Type:        ToastSuccess,
		Title:       title,
		Description: description,
		CloseButton: true,
	})
}

// Error creates an error toast
func Error(title, description string) g.Node {
	return Toast(ToastProps{
		Type:        ToastError,
		Title:       title,
		Description: description,
		CloseButton: true,
	})
}

// Warning creates a warning toast
func Warning(title, description string) g.Node {
	return Toast(ToastProps{
		Type:        ToastWarning,
		Title:       title,
		Description: description,
		CloseButton: true,
	})
}

// Info creates an info toast
func Info(title, description string) g.Node {
	return Toast(ToastProps{
		Type:        ToastInfo,
		Title:       title,
		Description: description,
		CloseButton: true,
	})
}

// Message creates a default toast with just a message
func Message(message string) g.Node {
	return Toast(ToastProps{
		Type:        ToastDefault,
		Description: message,
		CloseButton: true,
	})
}

// WithAction creates a toast with an action button
func WithAction(props ToastProps, label string, onClick string) g.Node {
	props.Action = &ToastAction{
		Label:   label,
		OnClick: onClick,
	}
	return Toast(props)
}

// LoadingToast creates a loading toast
func LoadingToast(message string) g.Node {
	return Toast(ToastProps{
		Type:        ToastDefault,
		Description: message,
		Duration:    0, // Infinite duration
		Class:       "pr-12",
	})
}

// Promise creates a promise-style toast (loading -> success/error)
type PromiseToast struct {
	ID          string
	Loading     string
	Success     string
	Error       string
}

// PromiseLoading creates the loading state of a promise toast
func PromiseLoading(p PromiseToast) g.Node {
	return Toast(ToastProps{
		ID:          p.ID,
		Type:        ToastDefault,
		Description: p.Loading,
		Duration:    0,
		Class:       "pr-12",
	})
}

// PromiseSuccess creates the success state of a promise toast
func PromiseSuccess(p PromiseToast) g.Node {
	return Toast(ToastProps{
		ID:          p.ID,
		Type:        ToastSuccess,
		Title:       "Success",
		Description: p.Success,
		CloseButton: true,
	})
}

// PromiseError creates the error state of a promise toast
func PromiseError(p PromiseToast) g.Node {
	return Toast(ToastProps{
		ID:          p.ID,
		Type:        ToastError,
		Title:       "Error",
		Description: p.Error,
		CloseButton: true,
	})
}

// Custom creates a custom toast with full control
func Custom(content g.Node) g.Node {
	id := fmt.Sprintf("toast-%d", time.Now().UnixNano())
	return html.Li(
		g.Attr("id", id),
		g.Attr("data-toast", "true"),
		g.Attr("data-type", "custom"),
		html.Class(lib.CN(
			"pointer-events-auto relative flex w-full items-center overflow-hidden rounded-lg shadow-lg transition-all",
			"data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-80",
		)),
		g.Attr("role", "status"),
		g.Attr("aria-live", "polite"),
		content,
	)
}