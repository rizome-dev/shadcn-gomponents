package sonner

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

// HTMXToasterProps defines properties for HTMX-enhanced toaster
type HTMXToasterProps struct {
	ID         string
	AddPath    string
	RemovePath string
	UpdatePath string
}

// HTMXToaster creates an HTMX-enhanced toaster container
func HTMXToaster(props ToasterProps, htmxProps HTMXToasterProps) g.Node {
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
	if htmxProps.ID == "" {
		htmxProps.ID = "toaster"
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
		g.Attr("id", htmxProps.ID),
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
		// SSE connection for server-sent events
		hx.Ext("sse"),
		g.Attr("sse-connect", htmxProps.UpdatePath),
		// Toast list
	html.Ol(
			g.Attr("id", htmxProps.ID+"-list"),
			html.Class("flex flex-col gap-[var(--gap)]"),
			g.Attr("data-toaster-list", "true"),
			// Server pushes toast updates here
			hx.Get(htmxProps.UpdatePath+"/list"),
			hx.Trigger("sse:toastUpdate"),
			hx.Target("this"),
			hx.Swap("innerHTML"),
		),
	)
}

// HTMXToast creates an HTMX-enhanced toast
func HTMXToast(props ToastProps, htmxProps HTMXToasterProps) g.Node {
	// Generate ID if not provided
	if props.ID == "" {
		props.ID = fmt.Sprintf("toast-%d", time.Now().UnixNano())
	}

	// Base toast
	toast := Toast(props)

	// Add HTMX attributes for auto-dismiss
	if props.Duration > 0 {
		toast = g.Group([]g.Node{
			toast,
			html.Script(g.Raw(fmt.Sprintf(`
				setTimeout(function() {
					htmx.ajax('DELETE', '%s', {
						target: '#%s',
						swap: 'outerHTML swap:0.3s',
						values: { id: '%s' }
					});
				}, %d);
			`, htmxProps.RemovePath, props.ID, props.ID, props.Duration))),
		})
	}

	// Add close button handler
	if props.CloseButton {
		closeScript := html.Script(g.Raw(fmt.Sprintf(`
			document.getElementById('%s').querySelector('[data-toast-close]').addEventListener('click', function() {
				htmx.ajax('DELETE', '%s', {
					target: '#%s',
					swap: 'outerHTML swap:0.3s',
					values: { id: '%s' }
				});
			});
		`, props.ID, htmxProps.RemovePath, props.ID, props.ID)))
		toast = g.Group([]g.Node{toast, closeScript})
	}

	return toast
}

// ToastStore manages server-side toast state
type ToastStore struct {
	mu     sync.RWMutex
	toasts map[string][]ToastProps
	order  map[string][]string // Maintains toast order
}

var toastStore = &ToastStore{
	toasts: make(map[string][]ToastProps),
	order:  make(map[string][]string),
}

// AddToast adds a toast to the store
func (s *ToastStore) AddToast(toasterID string, toast ToastProps) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if toast.ID == "" {
		toast.ID = fmt.Sprintf("toast-%d", time.Now().UnixNano())
	}

	s.toasts[toasterID] = append(s.toasts[toasterID], toast)
	s.order[toasterID] = append(s.order[toasterID], toast.ID)
}

// RemoveToast removes a toast from the store
func (s *ToastStore) RemoveToast(toasterID, toastID string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	toasts := s.toasts[toasterID]
	newToasts := make([]ToastProps, 0, len(toasts))
	for _, t := range toasts {
		if t.ID != toastID {
			newToasts = append(newToasts, t)
		}
	}
	s.toasts[toasterID] = newToasts

	// Update order
	order := s.order[toasterID]
	newOrder := make([]string, 0, len(order))
	for _, id := range order {
		if id != toastID {
			newOrder = append(newOrder, id)
		}
	}
	s.order[toasterID] = newOrder
}

// GetToasts returns all toasts for a toaster
func (s *ToastStore) GetToasts(toasterID string) []ToastProps {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.toasts[toasterID]
}

// ToasterHandlers creates HTTP handlers for toaster functionality
func ToasterHandlers(mux *http.ServeMux, baseProps ToasterProps, htmxProps HTMXToasterProps) {
	// Add toast handler
	mux.HandleFunc(htmxProps.AddPath, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parse toast data
		var toastData struct {
			Type        string `json:"type"`
			Title       string `json:"title"`
			Description string `json:"description"`
			Duration    int    `json:"duration"`
			Action      *struct {
				Label   string `json:"label"`
				OnClick string `json:"onClick"`
			} `json:"action"`
		}

		if err := json.NewDecoder(r.Body).Decode(&toastData); err != nil {
			// Fallback to form data
			r.ParseForm()
			toastData.Type = r.FormValue("type")
			toastData.Title = r.FormValue("title")
			toastData.Description = r.FormValue("description")
		}

		// Create toast
		toast := ToastProps{
			Type:        ToastType(toastData.Type),
			Title:       toastData.Title,
			Description: toastData.Description,
			Duration:    toastData.Duration,
			CloseButton: baseProps.CloseButton,
		}

		if toastData.Duration == 0 {
			toast.Duration = baseProps.Duration
		}

		if toastData.Action != nil {
			toast.Action = &ToastAction{
				Label:   toastData.Action.Label,
				OnClick: toastData.Action.OnClick,
			}
		}

		// Add to store
		toastStore.AddToast(htmxProps.ID, toast)

		// Trigger SSE update
		w.Header().Set("HX-Trigger", "toastUpdate")
		w.WriteHeader(http.StatusOK)
	})

	// Remove toast handler
	mux.HandleFunc(htmxProps.RemovePath, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		toastID := r.FormValue("id")
		if toastID == "" {
			http.Error(w, "Toast ID required", http.StatusBadRequest)
			return
		}

		toastStore.RemoveToast(htmxProps.ID, toastID)

		// Return empty response with fade out
		w.Header().Set("HX-Reswap", "outerHTML swap:0.3s")
		w.Write([]byte(""))
	})

	// Update handler (returns current toast list)
	mux.HandleFunc(htmxProps.UpdatePath+"/list", func(w http.ResponseWriter, r *http.Request) {
		toasts := toastStore.GetToasts(htmxProps.ID)
		
		// Render all toasts
		for _, toast := range toasts {
			HTMXToast(toast, htmxProps).Render(w)
		}
	})

	// SSE endpoint for real-time updates
	mux.HandleFunc(htmxProps.UpdatePath, func(w http.ResponseWriter, r *http.Request) {
		// Set headers for SSE
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		// Send initial connection message
		fmt.Fprintf(w, "event: connect\ndata: connected\n\n")
		w.(http.Flusher).Flush()

		// Keep connection alive
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				fmt.Fprintf(w, "event: ping\ndata: ping\n\n")
				w.(http.Flusher).Flush()
			case <-r.Context().Done():
				return
			}
		}
	})
}

// Helper functions for common toast operations

// ShowToast is a helper to trigger a toast via HTMX response header
func ShowToast(w http.ResponseWriter, toastType ToastType, title, description string) {
	toastData := map[string]interface{}{
		"type":        string(toastType),
		"title":       title,
		"description": description,
	}
	
	jsonData, _ := json.Marshal(toastData)
	w.Header().Set("HX-Trigger", fmt.Sprintf(`{"showToast": %s}`, string(jsonData)))
}

// ToastTriggerScript returns JavaScript to handle HX-Trigger showToast events
func ToastTriggerScript(htmxProps HTMXToasterProps) string {
	return fmt.Sprintf(`
		document.body.addEventListener('showToast', function(evt) {
			htmx.ajax('POST', '%s', {
				values: evt.detail
			});
		});
	`, htmxProps.AddPath)
}

// ExampleToastButtons creates example buttons to trigger different toast types
func ExampleToastButtons(htmxProps HTMXToasterProps) g.Node {
	return html.Div(
		html.Class("flex flex-wrap gap-2"),
		
		// Success
		html.Button(
			html.Type("button"),
			html.Class("px-4 py-2 rounded bg-green-500 text-white hover:bg-green-600"),
			hx.Post(htmxProps.AddPath),
			hx.Vals(`{
				"type": "success",
				"title": "Success!",
				"description": "Your action was completed successfully."
			}`),
			g.Text("Show Success"),
		),
		
		// Error
		html.Button(
			html.Type("button"),
			html.Class("px-4 py-2 rounded bg-red-500 text-white hover:bg-red-600"),
			hx.Post(htmxProps.AddPath),
			hx.Vals(`{
				"type": "error",
				"title": "Error",
				"description": "Something went wrong. Please try again."
			}`),
			g.Text("Show Error"),
		),
		
		// Warning
		html.Button(
			html.Type("button"),
			html.Class("px-4 py-2 rounded bg-yellow-500 text-white hover:bg-yellow-600"),
			hx.Post(htmxProps.AddPath),
			hx.Vals(`{
				"type": "warning",
				"title": "Warning",
				"description": "This action may have unintended consequences."
			}`),
			g.Text("Show Warning"),
		),
		
		// Info
		html.Button(
			html.Type("button"),
			html.Class("px-4 py-2 rounded bg-blue-500 text-white hover:bg-blue-600"),
			hx.Post(htmxProps.AddPath),
			hx.Vals(`{
				"type": "info",
				"title": "Information",
				"description": "Here's some helpful information for you."
			}`),
			g.Text("Show Info"),
		),
		
		// Default
		html.Button(
			html.Type("button"),
			html.Class("px-4 py-2 rounded bg-gray-500 text-white hover:bg-gray-600"),
			hx.Post(htmxProps.AddPath),
			hx.Vals(`{
				"type": "default",
				"description": "This is a simple message."
			}`),
			g.Text("Show Message"),
		),
		
		// With Action
		html.Button(
			html.Type("button"),
			html.Class("px-4 py-2 rounded bg-purple-500 text-white hover:bg-purple-600"),
			hx.Post(htmxProps.AddPath),
			hx.Vals(`{
				"type": "default",
				"title": "Update Available",
				"description": "A new version is available.",
				"action": {
					"label": "Update Now",
					"onClick": "onclick=\"alert('Updating...')\""
				}
			}`),
			g.Text("Show with Action"),
		),
	)
}