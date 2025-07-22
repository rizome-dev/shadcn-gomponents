package toast

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	hx "maragu.dev/gomponents-htmx"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// HTMXProps defines properties for HTMX-enhanced toasts
type HTMXProps struct {
	ToasterID   string // ID of the toaster container
	ShowPath    string // Path to show a toast
	DismissPath string // Path to dismiss a toast
	ClearPath   string // Path to clear all toasts
}

// HTMXToast creates an HTMX-enhanced toast
func HTMXToast(props Props, htmxProps HTMXProps) g.Node {
	// Set defaults
	if props.Variant == "" {
		props.Variant = VariantDefault
	}
	if props.Duration == 0 && !props.Closable {
		props.Duration = 5 * time.Second
	}
	if props.ID == "" {
		props.ID = fmt.Sprintf("htmx-toast-%d", time.Now().UnixNano())
	}

	// Build toast attributes
	attrs := []g.Node{
		html.ID(props.ID),
		g.Attr("role", "alert"),
		g.Attr("aria-live", "polite"),
		g.Attr("data-toast", "true"),
		g.Attr("data-variant", string(props.Variant)),
		g.Attr("data-state", "open"),
		html.Class(lib.CN(
			"htmx-toast",
			"relative flex w-full items-center justify-between space-x-2 overflow-hidden rounded-md border p-4 pr-6 shadow-lg transition-all",
			"animate-in slide-in-from-top-full",
			getVariantClasses(props.Variant),
			props.Class,
		)),
	}
	
	// Auto-dismiss with HTMX
	if props.Duration > 0 {
		attrs = append(attrs,
			hx.Get(fmt.Sprintf("%s?id=%s", htmxProps.DismissPath, props.ID)),
			hx.Trigger(fmt.Sprintf("load delay:%dms", props.Duration.Milliseconds())),
			hx.Target("this"),
			hx.Swap("outerHTML swap:300ms"),
		)
	}

	// Build content
	contentNodes := []g.Node{}

	// Add icon
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

	// Add action button
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

	// Add close button
	if props.Closable {
		toastContent = append(toastContent, html.Button(
			html.Type("button"),
			html.Class("toast-close absolute right-1 top-1 rounded-md p-1 opacity-70 transition-opacity hover:opacity-100 focus:outline-none focus:ring-2"),
			g.Attr("aria-label", "Close"),
			hx.Delete(fmt.Sprintf("%s?id=%s", htmxProps.DismissPath, props.ID)),
			hx.Target("closest [data-toast]"),
			hx.Swap("outerHTML swap:300ms"),
			g.Raw(`<svg class="h-4 w-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="18" y1="6" x2="6" y2="18"></line><line x1="6" y1="6" x2="18" y2="18"></line></svg>`),
		))
	}

	// Add progress bar if requested
	if props.Progress && props.Duration > 0 {
		toastContent = append(toastContent, html.Div(
			html.Class("absolute bottom-0 left-0 h-1 w-full bg-current opacity-20"),
			html.Div(
				html.Class("h-full bg-current transition-all ease-linear"),
				html.Style(fmt.Sprintf("animation: htmx-toast-progress %dms linear", props.Duration.Milliseconds())),
			),
		))
	}

	return html.Div(
		append(attrs, toastContent...)...,
	)
}

// HTMXToaster creates an HTMX-enhanced toaster container
func HTMXToaster(props ToasterProps, htmxProps HTMXProps) g.Node {
	// Set defaults
	if props.Position == "" {
		props.Position = PositionBottomRight
	}
	if props.MaxToast == 0 {
		props.MaxToast = 3
	}
	if props.ID == "" {
		props.ID = "htmx-toaster"
	}

	return html.Div(
		html.ID(props.ID),
		g.Attr("data-toaster", "true"),
		g.Attr("data-position", string(props.Position)),
		g.Attr("data-max-toasts", fmt.Sprintf("%d", props.MaxToast)),
		html.Class(lib.CN(
			"htmx-toaster",
			"fixed z-50 flex flex-col gap-2 md:max-w-sm",
			getPositionClasses(props.Position),
			props.Class,
		)),
		// Listen for toast events
		hx.Get(htmxProps.ShowPath),
		hx.Trigger("toast-show from:body"),
		hx.Target("this"),
		hx.Swap("beforeend"),
		// Add CSS for animations
		html.Style(`
			@keyframes htmx-toast-progress {
				from { width: 0%; }
				to { width: 100%; }
			}
			.htmx-settling {
				opacity: 0;
			}
		`),
	)
}

// Toast storage
type ToastStore struct {
	mu     sync.RWMutex
	toasts map[string]*Props
	order  []string
}

var globalToastStore = &ToastStore{
	toasts: make(map[string]*Props),
	order:  []string{},
}

// ToastHandlers creates HTTP handlers for toast functionality
func ToastHandlers(mux *http.ServeMux, htmxProps HTMXProps) {
	// Validate required paths
	if htmxProps.ShowPath == "" {
		panic("ToastHandlers: ShowPath is required")
	}
	if htmxProps.DismissPath == "" {
		panic("ToastHandlers: DismissPath is required")
	}
	
	// Show toast handler
	mux.HandleFunc(htmxProps.ShowPath, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet && r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var props Props

		// Parse toast data from request
		if r.Method == http.MethodPost {
			// Read from body
			if err := json.NewDecoder(r.Body).Decode(&props); err != nil {
				// Try form data
				props = Props{
					Title:       r.FormValue("title"),
					Description: r.FormValue("description"),
					Variant:     Variant(r.FormValue("variant")),
				}
			}
		} else {
			// Read from query params
			props = Props{
				Title:       r.URL.Query().Get("title"),
				Description: r.URL.Query().Get("description"),
				Variant:     Variant(r.URL.Query().Get("variant")),
			}
		}

		// Set defaults
		if props.Variant == "" {
			props.Variant = VariantDefault
		}
		if props.Duration == 0 {
			props.Duration = 5 * time.Second
		}
		props.Closable = true

		// Store toast
		globalToastStore.mu.Lock()
		props.ID = fmt.Sprintf("toast-%d", time.Now().UnixNano())
		globalToastStore.toasts[props.ID] = &props
		globalToastStore.order = append(globalToastStore.order, props.ID)
		
		// Limit number of toasts
		maxToasts := 3
		if len(globalToastStore.order) > maxToasts {
			// Remove oldest toasts
			toRemove := len(globalToastStore.order) - maxToasts
			for i := 0; i < toRemove; i++ {
				delete(globalToastStore.toasts, globalToastStore.order[i])
			}
			globalToastStore.order = globalToastStore.order[toRemove:]
		}
		globalToastStore.mu.Unlock()

		// Render the toast
		HTMXToast(props, htmxProps).Render(w)
	})

	// Dismiss toast handler
	mux.HandleFunc(htmxProps.DismissPath, func(w http.ResponseWriter, r *http.Request) {
		toastID := r.URL.Query().Get("id")
		if toastID == "" {
			http.Error(w, "Missing toast ID", http.StatusBadRequest)
			return
		}

		// Remove from store
		globalToastStore.mu.Lock()
		delete(globalToastStore.toasts, toastID)
		newOrder := []string{}
		for _, id := range globalToastStore.order {
			if id != toastID {
				newOrder = append(newOrder, id)
			}
		}
		globalToastStore.order = newOrder
		globalToastStore.mu.Unlock()

		// Return empty response (toast will remove itself)
		w.WriteHeader(http.StatusOK)
	})

	// Clear all toasts handler
	if htmxProps.ClearPath != "" {
		mux.HandleFunc(htmxProps.ClearPath, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost && r.Method != http.MethodDelete {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}

			// Clear store
			globalToastStore.mu.Lock()
			globalToastStore.toasts = make(map[string]*Props)
			globalToastStore.order = []string{}
			globalToastStore.mu.Unlock()

			// Return empty toaster
			html.Div(html.ID(htmxProps.ToasterID)).Render(w)
		})
	}
}

// ShowToastScript generates JavaScript to trigger a toast via HTMX
func ShowToastScript(title, description string, variant Variant) string {
	return fmt.Sprintf(`
		htmx.trigger(document.body, 'toast-show', {
			detail: {
				title: '%s',
				description: '%s',
				variant: '%s'
			}
		});
	`, title, description, variant)
}

// HTMXToastHelpers - Helper functions for common HTMX toast patterns

// FormSuccessToast creates a toast for form submission success
func FormSuccessToast(message string) g.Node {
	return html.Div(
		hx.Get("/api/toast/show"),
		hx.Vals(fmt.Sprintf(`{"description": "%s", "variant": "success"}`, message)),
		hx.Trigger("load"),
		hx.Target("#htmx-toaster"),
		hx.Swap("beforeend"),
	)
}

// AsyncToast creates a toast that updates based on async operation
func AsyncToast(taskID string, htmxProps HTMXProps) g.Node {
	props := Props{
		ID:          fmt.Sprintf("async-toast-%s", taskID),
		Description: "Processing...",
		Icon:        getLoadingIcon(),
		Duration:    0,
		Closable:    false,
	}

	// Create toast with polling for status updates
	attrs := []g.Node{
		html.ID(props.ID),
		g.Attr("role", "alert"),
		g.Attr("data-toast", "true"),
		html.Class("htmx-toast relative flex w-full items-center justify-between space-x-2 overflow-hidden rounded-md border p-4 shadow-lg"),
		hx.Get(fmt.Sprintf("/api/tasks/%s/status", taskID)),
		hx.Trigger("every 1s"),
		hx.Target("this"),
		hx.Swap("outerHTML"),
	}

	return html.Div(append(attrs,
		html.Div(html.Class("flex items-center space-x-2"),
			html.Div(html.Class("flex-shrink-0"), props.Icon),
			html.Div(html.Class("text-sm"), g.Text(props.Description)),
		))...,
	)
}

// NotificationToast creates a notification-style toast
func NotificationToast(props Props, htmxProps HTMXProps) g.Node {
	// Add notification-specific styling
	props.Class = lib.CN(props.Class, "min-w-[300px]")
	props.Closable = true
	
	return HTMXToast(props, htmxProps)
}

// ToastGroup creates a group of related toasts
func ToastGroup(groupID string, toasts []Props, htmxProps HTMXProps) g.Node {
	nodes := []g.Node{}
	
	for i, toast := range toasts {
		toast.ID = fmt.Sprintf("%s-%d", groupID, i)
		nodes = append(nodes, HTMXToast(toast, htmxProps))
	}
	
	return html.Div(
		html.Class("space-y-2"),
		g.Group(nodes),
	)
}

// PushToast sends a toast to connected clients via Server-Sent Events
func PushToast(w http.ResponseWriter, props Props) {
	data, _ := json.Marshal(props)
	fmt.Fprintf(w, "event: toast\ndata: %s\n\n", data)
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}
}

// ToastSSEHandler creates a Server-Sent Events handler for real-time toasts
func ToastSSEHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		// Send initial connection message
		fmt.Fprintf(w, "event: connected\ndata: {\"status\": \"connected\"}\n\n")
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}

		// Keep connection alive
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				fmt.Fprintf(w, "event: ping\ndata: {\"time\": \"%s\"}\n\n", time.Now().Format(time.RFC3339))
				if f, ok := w.(http.Flusher); ok {
					f.Flush()
				}
			case <-r.Context().Done():
				return
			}
		}
	}
}