package sonner

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

// Examples demonstrates various toast configurations
func Examples() g.Node {
	return html.Div(
		html.Class("space-y-8 p-8"),
		
		// Toast Types
		html.Div(
			html.Class("space-y-4"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Toast Types")),
			html.Div(
				html.Class("space-y-2"),
				// Success
				Success("Success!", "Your changes have been saved successfully."),
				
				// Error
				Error("Error", "Failed to save changes. Please try again."),
				
				// Warning
				Warning("Warning", "This action cannot be undone."),
				
				// Info
				Info("Information", "New updates are available for download."),
				
				// Default message
				Message("This is a simple message toast."),
			),
		),

		// Toast Variations
		html.Div(
			html.Class("space-y-4"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Toast Variations")),
			html.Div(
				html.Class("space-y-2"),
				// With title only
				Toast(ToastProps{
					Type:  ToastSuccess,
					Title: "Account created",
				}),
				
				// With description only
				Toast(ToastProps{
					Type:        ToastInfo,
					Description: "Your session will expire in 5 minutes",
				}),
				
				// With action button
				WithAction(ToastProps{
					Type:        ToastDefault,
					Title:       "Update Available",
					Description: "A new version of the app is available.",
				}, "Update Now", `onclick="alert('Updating...')"`),
				
				// Without close button
				Toast(ToastProps{
					Type:        ToastWarning,
					Title:       "Limited Time Offer",
					Description: "50% off all items today only!",
					CloseButton: false,
				}),
			),
		),

		// Loading States
		html.Div(
			html.Class("space-y-4"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Loading States")),
			html.Div(
				html.Class("space-y-2"),
				// Simple loading
				Loading("Processing your request..."),
				
				// Loading with custom styling
				Toast(ToastProps{
					Type:        ToastDefault,
					Description: "Uploading files...",
					Duration:    0,
					Class:       "pr-12",
				}),
			),
		),

		// Promise Toasts
		html.Div(
			html.Class("space-y-4"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Promise Toast States")),
			html.P(html.Class("text-sm text-muted-foreground"), 
				g.Text("These would transition from loading to success/error in a real application"),
			),
			html.Div(
				html.Class("space-y-2"),
				// Loading state
				PromiseLoading(PromiseToast{
					ID:      "save-1",
					Loading: "Saving your preferences...",
					Success: "Preferences saved successfully!",
					Error:   "Failed to save preferences",
				}),
				
				// Success state
				PromiseSuccess(PromiseToast{
					ID:      "save-2",
					Loading: "Saving your preferences...",
					Success: "Preferences saved successfully!",
					Error:   "Failed to save preferences",
				}),
				
				// Error state
				PromiseError(PromiseToast{
					ID:      "save-3",
					Loading: "Saving your preferences...",
					Success: "Preferences saved successfully!",
					Error:   "Failed to save preferences",
				}),
			),
		),

		// Custom Toasts
		html.Div(
			html.Class("space-y-4"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Custom Toasts")),
			html.Div(
				html.Class("space-y-2"),
				// User notification
				Custom(
					html.Div(
						html.Class("flex items-center gap-4 p-4"),
						html.Div(
							html.Class("w-10 h-10 rounded-full bg-blue-100 flex items-center justify-center"),
							html.Span(html.Class("text-blue-600 font-semibold"), g.Text("JD")),
						),
						html.Div(
							html.P(html.Class("font-semibold"), g.Text("John Doe")),
							html.P(html.Class("text-sm text-muted-foreground"), g.Text("Just sent you a message")),
						),
						html.Button(
							html.Type("button"),
							html.Class("ml-auto text-sm text-blue-600 hover:text-blue-700"),
							g.Text("View"),
						),
					),
				),
				
				// Achievement toast
				Custom(
					html.Div(
						html.Class("flex items-center gap-4 p-4 bg-gradient-to-r from-purple-500 to-pink-500 text-white"),
						g.Raw(`<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" class="text-yellow-300"><path d="M12 2L15.09 8.26L22 9.27L17 14.14L18.18 21.02L12 17.77L5.82 21.02L7 14.14L2 9.27L8.91 8.26L12 2Z" fill="currentColor" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/></svg>`),
						html.Div(
							html.P(html.Class("font-semibold"), g.Text("Achievement Unlocked!")),
							html.P(html.Class("text-sm opacity-90"), g.Text("You've completed 100 tasks")),
						),
					),
				),
				
				// Progress toast
				Custom(
					html.Div(
						html.Class("p-4"),
						html.Div(
							html.Class("flex justify-between mb-2"),
							html.Span(html.Class("text-sm font-medium"), g.Text("Uploading document.pdf")),
							html.Span(html.Class("text-sm text-muted-foreground"), g.Text("67%")),
						),
						html.Div(
							html.Class("w-full bg-gray-200 rounded-full h-2"),
							html.Div(
								html.Class("bg-blue-600 h-2 rounded-full transition-all duration-300"),
								g.Attr("style", "width: 67%"),
							),
						),
					),
				),
			),
		),

		// Toaster Positions
		html.Div(
			html.Class("space-y-4"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Toaster Positions")),
			html.P(html.Class("text-sm text-muted-foreground"), 
				g.Text("The toaster container can be positioned in different areas of the screen"),
			),
			html.Div(
				html.Class("grid grid-cols-3 gap-4 max-w-md mx-auto"),
				// Top positions
				html.Div(
					html.Class("text-center p-4 border rounded cursor-pointer hover:bg-muted"),
					g.Text("top-left"),
				),
				html.Div(
					html.Class("text-center p-4 border rounded cursor-pointer hover:bg-muted"),
					g.Text("top-center"),
				),
				html.Div(
					html.Class("text-center p-4 border rounded cursor-pointer hover:bg-muted"),
					g.Text("top-right"),
				),
				// Middle (placeholder)
				html.Div(html.Class("col-span-3 h-20")),
				// Bottom positions
				html.Div(
					html.Class("text-center p-4 border rounded cursor-pointer hover:bg-muted"),
					g.Text("bottom-left"),
				),
				html.Div(
					html.Class("text-center p-4 border rounded cursor-pointer hover:bg-muted"),
					g.Text("bottom-center"),
				),
				html.Div(
					html.Class("text-center p-4 border rounded cursor-pointer hover:bg-muted"),
					g.Text("bottom-right"),
				),
			),
		),

		// Toaster Configuration
		html.Div(
			html.Class("space-y-4"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Toaster Container Examples")),
			html.Div(
				html.Class("space-y-4"),
				// Default toaster
				html.Div(
					html.Class("space-y-2"),
					html.H4(html.Class("text-sm font-medium"), g.Text("Default Configuration")),
					html.Div(
						html.Class("relative h-32 border rounded bg-muted/20"),
						Toaster(ToasterProps{}),
					),
				),
				
				// Rich colors toaster
				html.Div(
					html.Class("space-y-2"),
					html.H4(html.Class("text-sm font-medium"), g.Text("Rich Colors")),
					html.Div(
						html.Class("relative h-32 border rounded bg-muted/20"),
						Toaster(ToasterProps{
							RichColors: true,
							Position:   PositionTopLeft,
						}),
					),
				),
				
				// Custom configuration
				html.Div(
					html.Class("space-y-2"),
					html.H4(html.Class("text-sm font-medium"), g.Text("Custom Configuration")),
					html.Div(
						html.Class("relative h-32 border rounded bg-muted/20"),
						Toaster(ToasterProps{
							Position:    PositionBottomCenter,
							Duration:    3000,
							Gap:         24,
							MaxVisible:  5,
							CloseButton: true,
							Expand:      true,
						}),
					),
				),
			),
		),

		// Complex Examples
		html.Div(
			html.Class("space-y-4"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Complex Examples")),
			html.Div(
				html.Class("space-y-2"),
				// Multi-line toast
				Toast(ToastProps{
					Type:  ToastInfo,
					Title: "System Maintenance",
					Description: "We'll be performing scheduled maintenance on Sunday, December 10th from 2:00 AM to 6:00 AM UTC. Some services may be temporarily unavailable.",
					Action: &ToastAction{
						Label:   "Learn More",
						OnClick: `onclick="window.open('/maintenance', '_blank')"`,
					},
					CloseButton: true,
				}),
				
				// Error with details
				Toast(ToastProps{
					Type:  ToastError,
					Title: "Upload Failed",
					Description: "File size exceeds the maximum limit of 10MB. Please compress your file and try again.",
					Action: &ToastAction{
						Label:   "Retry",
						OnClick: `onclick="retryUpload()"`,
					},
					CloseButton: true,
				}),
				
				// Success with next steps
				Toast(ToastProps{
					Type:  ToastSuccess,
					Title: "Welcome to Pro!",
					Description: "Your account has been upgraded. Explore your new features in the dashboard.",
					Action: &ToastAction{
						Label:   "Go to Dashboard",
						OnClick: `onclick="location.href='/dashboard'"`,
					},
					CloseButton: true,
				}),
			),
		),
	)
}