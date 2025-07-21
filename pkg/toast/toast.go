package toast

import (
	"fmt"
	"time"

	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// Variant defines the visual style of the toast
type Variant string

const (
	VariantDefault     Variant = "default"
	VariantSuccess     Variant = "success"
	VariantError       Variant = "error"
	VariantWarning     Variant = "warning"
	VariantInfo        Variant = "info"
	VariantDestructive Variant = "destructive"
)

// Position defines where toasts appear on screen
type Position string

const (
	PositionTopLeft      Position = "top-left"
	PositionTopCenter    Position = "top-center"
	PositionTopRight     Position = "top-right"
	PositionBottomLeft   Position = "bottom-left"
	PositionBottomCenter Position = "bottom-center"
	PositionBottomRight  Position = "bottom-right"
)

// Props defines the properties for a Toast component
type Props struct {
	ID          string        // Unique identifier
	Title       string        // Toast title
	Description string        // Toast description/body text
	Variant     Variant       // Visual variant
	Duration    time.Duration // How long to show (0 = persistent)
	Closable    bool          // Show close button
	Action      *ActionProps  // Optional action button
	Icon        g.Node        // Optional icon
	Progress    bool          // Show progress bar
	Class       string        // Additional CSS classes
	OnClose     string        // JavaScript to run on close
}

// ActionProps defines properties for toast action buttons
type ActionProps struct {
	Label   string // Button text
	OnClick string // JavaScript click handler
	Class   string // Additional CSS classes
}

// ToasterProps defines properties for the Toaster container
type ToasterProps struct {
	ID       string   // Container ID
	Position Position // Where to show toasts
	MaxToast int      // Maximum visible toasts
	Class    string   // Additional CSS classes
}

// New creates a single toast notification
func New(props Props) g.Node {
	// Set defaults
	if props.Variant == "" {
		props.Variant = VariantDefault
	}
	if props.ID == "" {
		props.ID = fmt.Sprintf("toast-%d", time.Now().UnixNano())
	}

	// Build toast attributes
	attrs := []g.Node{
		html.ID(props.ID),
		g.Attr("role", "alert"),
		g.Attr("aria-live", "polite"),
		g.Attr("data-toast", "true"),
		g.Attr("data-variant", string(props.Variant)),
		html.Class(lib.CN(
			"toast",
			"relative flex w-full items-center justify-between space-x-2 overflow-hidden rounded-md border p-4 pr-6 shadow-lg transition-all",
			"data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-80",
			"data-[state=closed]:slide-out-to-right-full data-[state=open]:slide-in-from-top-full",
			getVariantClasses(props.Variant),
			props.Class,
		)),
	}

	// Add duration data attribute if set
	if props.Duration > 0 {
		attrs = append(attrs, g.Attr("data-duration", fmt.Sprintf("%d", props.Duration.Milliseconds())))
	}

	// Build content
	contentNodes := []g.Node{}

	// Add icon if provided
	if props.Icon != nil {
		contentNodes = append(contentNodes, html.Div(
			html.Class("flex-shrink-0"),
			props.Icon,
		))
	} else if defaultIcon := getDefaultIcon(props.Variant); defaultIcon != nil {
		contentNodes = append(contentNodes, html.Div(
			html.Class("flex-shrink-0"),
			defaultIcon,
		))
	}

	// Add text content
	textContent := []g.Node{}
	if props.Title != "" {
		textContent = append(textContent, html.Div(
			html.Class("text-sm font-semibold"),
			g.Text(props.Title),
		))
	}
	if props.Description != "" {
		textContent = append(textContent, html.Div(
			html.Class("text-sm opacity-90"),
			g.Text(props.Description),
		))
	}
	
	contentNodes = append(contentNodes, html.Div(
		html.Class("flex-1 space-y-1"),
		g.Group(textContent),
	))

	// Add action button if provided
	if props.Action != nil {
		contentNodes = append(contentNodes, html.Button(
			html.Type("button"),
			html.Class(lib.CN(
				"ml-auto flex-shrink-0 rounded-md px-3 py-1 text-sm font-medium",
				"hover:opacity-90 focus:outline-none focus:ring-2 focus:ring-offset-2",
				getActionButtonClasses(props.Variant),
				props.Action.Class,
			)),
			g.Attr("onclick", props.Action.OnClick),
			g.Text(props.Action.Label),
		))
	}

	// Build the toast
	toastContent := []g.Node{
		html.Div(
			html.Class("flex w-full items-center space-x-2"),
			g.Group(contentNodes),
		),
	}

	// Add close button if closable
	if props.Closable {
		closeHandler := `this.closest('[data-toast]').setAttribute('data-state', 'closed'); setTimeout(() => this.closest('[data-toast]').remove(), 300);`
		if props.OnClose != "" {
			closeHandler = props.OnClose + "; " + closeHandler
		}
		
		toastContent = append(toastContent, html.Button(
			html.Type("button"),
			html.Class("toast-close absolute right-1 top-1 rounded-md p-1 opacity-70 transition-opacity hover:opacity-100 focus:outline-none focus:ring-2"),
			g.Attr("aria-label", "Close"),
			g.Attr("onclick", closeHandler),
			g.Raw(`<svg class="h-4 w-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="18" y1="6" x2="6" y2="18"></line><line x1="6" y1="6" x2="18" y2="18"></line></svg>`),
		))
	}

	// Add progress bar if requested
	if props.Progress && props.Duration > 0 {
		toastContent = append(toastContent, html.Div(
			html.Class("absolute bottom-0 left-0 h-1 bg-current opacity-20"),
			html.Style(fmt.Sprintf("animation: toast-progress %dms linear", props.Duration.Milliseconds())),
			html.Div(
				html.Class("h-full bg-current"),
				html.Style("width: 0%; animation: inherit"),
			),
		))
	}

	return html.Div(
		append(attrs, toastContent...)...,
	)
}

// Toaster creates a container for displaying toasts
func Toaster(props ToasterProps) g.Node {
	// Set defaults
	if props.Position == "" {
		props.Position = PositionBottomRight
	}
	if props.MaxToast == 0 {
		props.MaxToast = 3
	}
	if props.ID == "" {
		props.ID = "toaster"
	}

	return html.Div(
		html.ID(props.ID),
		g.Attr("data-toaster", "true"),
		g.Attr("data-position", string(props.Position)),
		g.Attr("data-max-toasts", fmt.Sprintf("%d", props.MaxToast)),
		html.Class(lib.CN(
			"toaster",
			"fixed z-50 flex flex-col gap-2 md:max-w-sm",
			getPositionClasses(props.Position),
			props.Class,
		)),
		// Add CSS for animations
		html.Style(`
			@keyframes toast-progress {
				from { width: 100%; }
				to { width: 0%; }
			}
		`),
	)
}

// Helper functions for creating common toasts

// Success creates a success toast
func Success(title, description string) g.Node {
	return New(Props{
		Title:       title,
		Description: description,
		Variant:     VariantSuccess,
		Closable:    true,
	})
}

// Error creates an error toast
func Error(title, description string) g.Node {
	return New(Props{
		Title:       title,
		Description: description,
		Variant:     VariantError,
		Closable:    true,
		Duration:    10 * time.Second, // Errors stay longer
	})
}

// Warning creates a warning toast
func Warning(title, description string) g.Node {
	return New(Props{
		Title:       title,
		Description: description,
		Variant:     VariantWarning,
		Closable:    true,
	})
}

// Info creates an info toast
func Info(title, description string) g.Node {
	return New(Props{
		Title:       title,
		Description: description,
		Variant:     VariantInfo,
		Closable:    true,
	})
}

// Simple creates a simple toast with just a message
func Simple(message string) g.Node {
	return New(Props{
		Description: message,
		Closable:    true,
	})
}

// LoadingToast creates a loading toast
func LoadingToast(message string) g.Node {
	return New(Props{
		Description: message,
		Icon: g.Raw(`<svg class="h-4 w-4 animate-spin" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
			<circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
			<path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
		</svg>`),
		Duration: 0, // No auto-dismiss
		Closable: false,
	})
}


// WithAction creates a toast with an action button
func WithAction(title, description, actionLabel string, actionHandler string) g.Node {
	return New(Props{
		Title:       title,
		Description: description,
		Closable:    true,
		Action: &ActionProps{
			Label:   actionLabel,
			OnClick: actionHandler,
		},
	})
}

// Promise creates a promise-style toast (loading -> success/error)
func Promise(id, loadingMessage string) g.Node {
	return New(Props{
		ID:          id,
		Description: loadingMessage,
		Icon: g.Raw(`<svg class="h-5 w-5 animate-spin" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"><circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle><path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path></svg>`),
		Duration:    0,
		Closable:    false,
	})
}

// Helper to update a promise toast to success
func PromiseSuccess(id, message string) string {
	return fmt.Sprintf(`
		const toast = document.getElementById('%s');
		if (toast) {
			toast.querySelector('.text-sm').textContent = '%s';
			toast.setAttribute('data-variant', 'success');
			const icon = toast.querySelector('svg').parentElement;
			icon.innerHTML = '%s';
			setTimeout(() => {
				toast.setAttribute('data-state', 'closed');
				setTimeout(() => toast.remove(), 300);
			}, 3000);
		}
	`, id, message, getSuccessIconString())
}

// Helper to update a promise toast to error
func PromiseError(id, message string) string {
	return fmt.Sprintf(`
		const toast = document.getElementById('%s');
		if (toast) {
			toast.querySelector('.text-sm').textContent = '%s';
			toast.setAttribute('data-variant', 'error');
			const icon = toast.querySelector('svg').parentElement;
			icon.innerHTML = '%s';
			toast.innerHTML += '<button type="button" class="toast-close absolute right-1 top-1 rounded-md p-1 opacity-70 transition-opacity hover:opacity-100" onclick="this.closest(\\'[data-toast]\\').setAttribute(\\'data-state\\', \\'closed\\'); setTimeout(() => this.closest(\\'[data-toast]\\').remove(), 300);">%s</button>';
		}
	`, id, message, getErrorIconString(), getCloseIconString())
}

// Styling helper functions

func getVariantClasses(variant Variant) string {
	switch variant {
	case VariantSuccess:
		return "border-border bg-card text-card-foreground"
	case VariantError, VariantDestructive:
		return "border-destructive bg-destructive text-destructive-foreground"
	case VariantWarning:
		return "border-border bg-accent text-accent-foreground"
	case VariantInfo:
		return "border-border bg-muted text-muted-foreground"
	default:
		return "border-border bg-popover text-popover-foreground"
	}
}

func getActionButtonClasses(variant Variant) string {
	switch variant {
	case VariantSuccess:
		return "bg-primary text-primary-foreground hover:bg-primary/90 focus:ring-ring"
	case VariantError, VariantDestructive:
		return "bg-destructive text-destructive-foreground hover:bg-destructive/90 focus:ring-destructive"
	case VariantWarning:
		return "bg-accent text-accent-foreground hover:bg-accent/90 focus:ring-ring"
	case VariantInfo:
		return "bg-secondary text-secondary-foreground hover:bg-secondary/90 focus:ring-ring"
	default:
		return "bg-primary text-primary-foreground hover:bg-primary/90 focus:ring-ring"
	}
}

func getPositionClasses(position Position) string {
	switch position {
	case PositionTopLeft:
		return "top-4 left-4"
	case PositionTopCenter:
		return "top-4 left-1/2 -translate-x-1/2"
	case PositionTopRight:
		return "top-4 right-4"
	case PositionBottomLeft:
		return "bottom-4 left-4"
	case PositionBottomCenter:
		return "bottom-4 left-1/2 -translate-x-1/2"
	case PositionBottomRight:
		return "bottom-4 right-4"
	default:
		return "bottom-4 right-4"
	}
}

func getDefaultIcon(variant Variant) g.Node {
	switch variant {
	case VariantSuccess:
		return g.Raw(getSuccessIconString())
	case VariantError, VariantDestructive:
		return g.Raw(getErrorIconString())
	case VariantWarning:
		return g.Raw(getWarningIconString())
	case VariantInfo:
		return g.Raw(getInfoIconString())
	default:
		return nil
	}
}

// Icon string helpers (for JavaScript updates)
func getSuccessIconString() string {
	return `<svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor"><path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" /></svg>`
}

func getErrorIconString() string {
	return `<svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor"><path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clip-rule="evenodd" /></svg>`
}

func getWarningIconString() string {
	return `<svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor"><path fill-rule="evenodd" d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z" clip-rule="evenodd" /></svg>`
}

func getInfoIconString() string {
	return `<svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor"><path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z" clip-rule="evenodd" /></svg>`
}

func getLoadingIcon() g.Node {
	return g.Raw(`<svg class="h-5 w-5 animate-spin" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"><circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle><path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path></svg>`)
}

func getCloseIconString() string {
	return `<svg class="h-4 w-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="18" y1="6" x2="6" y2="18"></line><line x1="6" y1="6" x2="18" y2="18"></line></svg>`
}