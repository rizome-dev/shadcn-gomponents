package toast

import (
	"time"

	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/pkg/button"
)

// Example demonstrates how to use the Toast component
func Example() g.Node {
	// Note: In a real app, you would need JavaScript to handle showing toasts dynamically
	// This example shows the static structure and appearance of toasts
	
	return html.Div(
		html.Class("p-8 space-y-8"),
		
		// Toast demos container
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4"), g.Text("Toast Examples")),
			html.P(html.Class("text-sm text-muted-foreground mb-6"), 
				g.Text("Note: These are static examples. In a real app, toasts would appear dynamically.")),
			
			// Basic toasts
			html.Div(
				html.H5(html.Class("text-xs font-medium uppercase text-muted-foreground mb-4"), g.Text("Basic Toasts")),
				html.Div(html.Class("space-y-2"),
					Simple("Your message has been sent."),
					New(Props{
						Title:       "Event Created",
						Description: "Your event has been added to the calendar.",
						Closable:    true,
					}),
					New(Props{
						Title:       "Scheduled",
						Description: "Your meeting is scheduled for tomorrow at 10 AM.",
						Closable:    true,
						Icon:        g.Raw(`<svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor"><path fill-rule="evenodd" d="M6 2a1 1 0 00-1 1v1H4a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V6a2 2 0 00-2-2h-1V3a1 1 0 10-2 0v1H7V3a1 1 0 00-1-1zM4 8h12v8H4V8z" clip-rule="evenodd" /></svg>`),
					}),
				),
			),
			
			// Toast variants
			html.Div(
				html.H5(html.Class("text-xs font-medium uppercase text-muted-foreground mb-4 mt-8"), g.Text("Toast Variants")),
				html.Div(html.Class("space-y-2"),
					Success("Success", "Your changes have been saved successfully."),
					Error("Error", "There was a problem with your request."),
					Warning("Warning", "Your session will expire in 5 minutes."),
					Info("Information", "A new version is available for download."),
				),
			),
			
			// Toasts with actions
			html.Div(
				html.H5(html.Class("text-xs font-medium uppercase text-muted-foreground mb-4 mt-8"), g.Text("With Actions")),
				html.Div(html.Class("space-y-2"),
					WithAction(
						"File Deleted",
						"important-file.pdf has been permanently deleted.",
						"Undo",
						"alert('Undo clicked')",
					),
					WithAction(
						"Email Sent",
						"Your email has been sent to 3 recipients.",
						"View",
						"alert('View clicked')",
					),
					New(Props{
						Title:       "Invitation Sent",
						Description: "An invitation has been sent to user@example.com",
						Variant:     VariantSuccess,
						Closable:    true,
						Action: &ActionProps{
							Label:   "Resend",
							OnClick: "alert('Resending invitation...')",
						},
					}),
				),
			),
			
			// Loading and promise toasts
			html.Div(
				html.H5(html.Class("text-xs font-medium uppercase text-muted-foreground mb-4 mt-8"), g.Text("Loading States")),
				html.Div(html.Class("space-y-2"),
					LoadingToast("Uploading your file..."),
					Promise("upload-123", "Processing document..."),
					New(Props{
						Description: "Generating report...",
						Icon:        getLoadingIcon(),
						Progress:    true,
						Duration:    10 * time.Second,
					}),
				),
			),
			
			// Different durations
			html.Div(
				html.H5(html.Class("text-xs font-medium uppercase text-muted-foreground mb-4 mt-8"), g.Text("Different Durations")),
				html.Div(html.Class("space-y-2"),
					New(Props{
						Description: "Quick notification (3s)",
						Duration:    3 * time.Second,
						Closable:    true,
					}),
					New(Props{
						Description: "Standard notification (5s)",
						Duration:    5 * time.Second,
						Closable:    true,
					}),
					New(Props{
						Description: "Long notification (10s)",
						Duration:    10 * time.Second,
						Closable:    true,
					}),
					New(Props{
						Description: "Persistent notification (no auto-dismiss)",
						Duration:    0,
						Closable:    true,
					}),
				),
			),
			
			// Custom styled toasts
			html.Div(
				html.H5(html.Class("text-xs font-medium uppercase text-muted-foreground mb-4 mt-8"), g.Text("Custom Styles")),
				html.Div(html.Class("space-y-2"),
					New(Props{
						Title:       "🎉 Congratulations!",
						Description: "You've unlocked a new achievement.",
						Class:       "bg-gradient-to-r from-purple-400 to-pink-400 text-white border-0",
						Closable:    true,
					}),
					New(Props{
						Title:       "System Update",
						Description: "A system update is available. Restart required.",
						Icon:        g.Raw(`<svg class="h-5 w-5 text-orange-500" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor"><path fill-rule="evenodd" d="M11.3 1.046A1 1 0 0112 2v5h4a1 1 0 01.82 1.573l-7 10A1 1 0 018 18v-5H4a1 1 0 01-.82-1.573l7-10a1 1 0 011.12-.38z" clip-rule="evenodd" /></svg>`),
						Class:       "border-orange-200 bg-orange-50 text-orange-900",
						Closable:    true,
					}),
				),
			),
		),
		
		// Toaster positions demo
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4 mt-12"), g.Text("Toaster Positions")),
			html.P(html.Class("text-sm text-muted-foreground mb-6"), 
				g.Text("The toaster container can be positioned in different corners of the screen.")),
			
			html.Div(html.Class("grid grid-cols-3 gap-4 max-w-2xl mx-auto"),
				// Top positions
				html.Div(html.Class("text-center p-4 border rounded-md"),
					html.P(html.Class("text-xs font-medium mb-2"), g.Text("Top Left")),
					html.Div(html.Class("relative h-24 bg-muted rounded"),
						html.Div(html.Class("absolute top-2 left-2 w-20 h-8 bg-primary/20 rounded")),
					),
				),
				html.Div(html.Class("text-center p-4 border rounded-md"),
					html.P(html.Class("text-xs font-medium mb-2"), g.Text("Top Center")),
					html.Div(html.Class("relative h-24 bg-muted rounded"),
						html.Div(html.Class("absolute top-2 left-1/2 -translate-x-1/2 w-20 h-8 bg-primary/20 rounded")),
					),
				),
				html.Div(html.Class("text-center p-4 border rounded-md"),
					html.P(html.Class("text-xs font-medium mb-2"), g.Text("Top Right")),
					html.Div(html.Class("relative h-24 bg-muted rounded"),
						html.Div(html.Class("absolute top-2 right-2 w-20 h-8 bg-primary/20 rounded")),
					),
				),
				// Bottom positions
				html.Div(html.Class("text-center p-4 border rounded-md"),
					html.P(html.Class("text-xs font-medium mb-2"), g.Text("Bottom Left")),
					html.Div(html.Class("relative h-24 bg-muted rounded"),
						html.Div(html.Class("absolute bottom-2 left-2 w-20 h-8 bg-primary/20 rounded")),
					),
				),
				html.Div(html.Class("text-center p-4 border rounded-md"),
					html.P(html.Class("text-xs font-medium mb-2"), g.Text("Bottom Center")),
					html.Div(html.Class("relative h-24 bg-muted rounded"),
						html.Div(html.Class("absolute bottom-2 left-1/2 -translate-x-1/2 w-20 h-8 bg-primary/20 rounded")),
					),
				),
				html.Div(html.Class("text-center p-4 border rounded-md bg-primary/5"),
					html.P(html.Class("text-xs font-medium mb-2"), g.Text("Bottom Right (Default)")),
					html.Div(html.Class("relative h-24 bg-muted rounded"),
						html.Div(html.Class("absolute bottom-2 right-2 w-20 h-8 bg-primary/40 rounded")),
					),
				),
			),
		),
		
		// Interactive demo buttons
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4 mt-12"), g.Text("Interactive Demo")),
			html.P(html.Class("text-sm text-muted-foreground mb-6"), 
				g.Text("Click these buttons to show toasts (requires JavaScript implementation):")),
			
			html.Div(html.Class("flex flex-wrap gap-2"),
				button.Default(
					g.Attr("onclick", `
						const toast = document.createElement('div');
						toast.innerHTML = 'Event has been created';
						toast.className = 'toast animate-in slide-in-from-top-full';
						// Add to toaster container
						alert('Toast would appear: Event has been created');
					`),
					g.Text("Show Toast"),
				),
				button.Secondary(
					g.Attr("onclick", `alert('Success toast would appear')`),
					g.Text("Success Toast"),
				),
				button.Secondary(
					g.Attr("onclick", `alert('Error toast would appear')`),
					g.Text("Error Toast"),
				),
				button.Secondary(
					g.Attr("onclick", `alert('Loading toast would appear')`),
					g.Text("Loading Toast"),
				),
				button.Secondary(
					g.Attr("onclick", `alert('Toast with action would appear')`),
					g.Text("With Action"),
				),
			),
		),
		
		// Implementation notes
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4 mt-12"), g.Text("Implementation Notes")),
			html.Div(html.Class("rounded-lg border bg-muted/50 p-4"),
				html.Ul(html.Class("text-sm space-y-2 list-disc list-inside"),
					html.Li(g.Text("Add a Toaster component to your app root to display toasts")),
					html.Li(g.Text("Use JavaScript to dynamically add toast elements to the toaster")),
					html.Li(g.Text("Toasts auto-dismiss based on their duration setting")),
					html.Li(g.Text("Users can manually dismiss toasts with the close button")),
					html.Li(g.Text("Multiple toasts stack vertically with proper spacing")),
					html.Li(g.Text("The toaster limits the number of visible toasts (default: 3)")),
					html.Li(g.Text("For HTMX integration, use HTMXToast and HTMXToaster components")),
				),
			),
		),
		
		// Code example
		html.Div(
			html.H4(html.Class("text-sm font-medium mb-4 mt-12"), g.Text("Usage Example")),
	html.Pre(html.Class("text-xs bg-muted p-4 rounded-lg overflow-x-auto"),
	html.Code(g.Raw(`// Add toaster to your app
toast.Toaster(toast.ToasterProps{
    Position: toast.PositionBottomRight,
    MaxToast: 3,
})

// Show different toast types
toast.Success("Success", "Your changes have been saved.")
toast.Error("Error", "Something went wrong!")
toast.Warning("Warning", "Please check your input.")
toast.Info("Info", "New update available.")

// Toast with action
toast.WithAction(
    "File deleted",
    "document.pdf has been deleted.",
    "Undo",
    "undoDelete()",
)

// Loading toast
toast.LoadingToast("Processing your request...")

// Custom toast
toast.New(toast.Props{
    Title:       "Custom Toast",
    Description: "With custom styling",
    Variant:     toast.VariantDefault,
    Duration:    7 * time.Second,
    Closable:    true,
    Progress:    true,
})`)),
			),
		),
	)
}