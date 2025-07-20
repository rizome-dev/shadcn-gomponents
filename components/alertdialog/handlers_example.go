package alertdialog

import (
	"net/http"
	"strings"
	
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	. "maragu.dev/gomponents/http"
)

// ExampleHandlers shows how to set up HTTP handlers for HTMX alert dialogs
func ExampleHandlers() {
	mux := http.NewServeMux()
	
	// Basic alert dialog handlers
	htmxProps := HTMXProps{
		ID:          "alert-dialog-example",
		TriggerPath: "/api/alert-dialog/open",
		ClosePath:   "/api/alert-dialog/close",
	}
	
	// Open dialog handler
	mux.HandleFunc("/api/alert-dialog/open", Adapt(func(w http.ResponseWriter, r *http.Request) (g.Node, error) {
		return RenderOpenDialog(htmxProps), nil
	}))
	
	// Close dialog handler
	mux.HandleFunc("/api/alert-dialog/close", Adapt(func(w http.ResponseWriter, r *http.Request) (g.Node, error) {
		return RenderClosedDialog(htmxProps), nil
	}))
	
	// Confirm action handler
	mux.HandleFunc("/api/alert-dialog/confirm", Adapt(func(w http.ResponseWriter, r *http.Request) (g.Node, error) {
		// Handle the confirmation action
		// In a real app, you would perform the action here
		
		// Return a success message or redirect
		return html.Div(
			html.ID(htmxProps.ID),
			html.Class("fixed bottom-4 right-4 bg-green-500 text-white px-4 py-2 rounded-md shadow-lg"),
			g.Text("Action confirmed!"),
			// Auto-hide after 3 seconds
			g.Attr("x-data", "{}"),
			g.Attr("x-init", "setTimeout(() => $el.remove(), 3000)"),
		), nil
	}))
	
	// Delete account dialog handlers
	deleteHTMXProps := HTMXProps{
		ID:          "delete-account-dialog",
		TriggerPath: "/api/alert-dialog/delete-account/open",
		ClosePath:   "/api/alert-dialog/delete-account/close",
	}
	
	// Open delete account dialog
	mux.HandleFunc("/api/alert-dialog/delete-account/open", Adapt(func(w http.ResponseWriter, r *http.Request) (g.Node, error) {
		return RenderDeleteAccountDialog(deleteHTMXProps), nil
	}))
	
	// Close delete account dialog
	mux.HandleFunc("/api/alert-dialog/delete-account/close", Adapt(func(w http.ResponseWriter, r *http.Request) (g.Node, error) {
		return RenderClosedDialog(deleteHTMXProps), nil
	}))
	
	// Validate delete confirmation input
	mux.HandleFunc("/api/alert-dialog/delete-account/validate", Adapt(func(w http.ResponseWriter, r *http.Request) (g.Node, error) {
		r.ParseForm()
		value := strings.TrimSpace(r.FormValue("delete-confirmation"))
		isValid := strings.ToUpper(value) == "DELETE"
		
		return RenderDeleteButton(deleteHTMXProps, isValid), nil
	}))
	
	// Confirm account deletion
	mux.HandleFunc("/api/alert-dialog/delete-account/confirm", Adapt(func(w http.ResponseWriter, r *http.Request) (g.Node, error) {
		// In a real app, you would delete the account here
		
		// Return a message
		return html.Div(
			html.ID(deleteHTMXProps.ID),
			html.Class("fixed inset-0 z-50 flex items-center justify-center bg-black/50"),
			html.Div(
				html.Class("bg-white rounded-lg p-6 max-w-sm mx-4"),
				html.H3(html.Class("text-lg font-semibold mb-2"), g.Text("Account Deleted")),
				html.P(html.Class("text-muted-foreground mb-4"), g.Text("Your account has been permanently deleted.")),
				html.Button(
					html.Type("button"),
					html.Class("w-full bg-primary text-primary-foreground hover:bg-primary/90 rounded-md px-4 py-2"),
					g.Text("OK"),
					// In a real app, you would redirect to login page
					g.Attr("onclick", "window.location.href = '/'"),
				),
			),
		), nil
	}))
}

// ConfirmationDialogHandler creates a reusable confirmation dialog handler
func ConfirmationDialogHandler(
	dialogID string,
	title string,
	description string,
	confirmText string,
	onConfirm func() error,
) http.HandlerFunc {
	htmxProps := HTMXProps{
		ID:          dialogID,
		TriggerPath: "/api/dialog/" + dialogID + "/open",
		ClosePath:   "/api/dialog/" + dialogID + "/close",
	}
	
	return Adapt(func(w http.ResponseWriter, r *http.Request) (g.Node, error) {
		switch r.URL.Path {
		case htmxProps.TriggerPath:
			// Open dialog
			return NewHTMX(
				Props{Open: true},
				htmxProps,
				DialogOverlayHTMX(htmxProps),
				DialogContent(
					ContentProps{},
					DialogHeader(
						HeaderProps{},
						DialogTitle(TitleProps{}, title),
						DialogDescription(DescriptionProps{}, description),
					),
					DialogFooter(
						FooterProps{},
						CancelHTMX(ActionProps{}, htmxProps, g.Text("Cancel")),
						ActionHTMX(
							ActionProps{},
							htmxProps,
							"/api/dialog/" + dialogID + "/confirm",
							g.Text(confirmText),
						),
					),
				),
			), nil
			
		case htmxProps.ClosePath:
			// Close dialog
			return RenderClosedDialog(htmxProps), nil
			
		case "/api/dialog/" + dialogID + "/confirm":
			// Handle confirmation
			if err := onConfirm(); err != nil {
				// Return error message
				return html.Div(
					html.ID(htmxProps.ID),
					html.Class("fixed bottom-4 right-4 bg-red-500 text-white px-4 py-2 rounded-md shadow-lg"),
					g.Text("Error: " + err.Error()),
				), nil
			}
			
			// Return success message
			return html.Div(
				html.ID(htmxProps.ID),
				html.Class("fixed bottom-4 right-4 bg-green-500 text-white px-4 py-2 rounded-md shadow-lg"),
				g.Text("Success!"),
			), nil
			
		default:
			return nil, nil
		}
	})
}