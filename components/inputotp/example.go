package inputotp

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

// Examples demonstrates various InputOTP usage patterns
func Examples() g.Node {
	return html.Div(
		html.Class("space-y-8 p-8"),
		
		// Basic OTP Input
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Basic OTP Input")),
			html.Div(
				html.Class("space-y-4"),
				html.Label(g.Text("Enter verification code")),
				Default(),
			),
		),

		// Different Lengths
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Different Lengths")),
			html.Div(
				html.Class("space-y-4"),
				html.Div(
					html.Label(g.Text("4-digit code")),
					FourDigit(),
				),
				html.Div(
					html.Label(g.Text("6-digit code (default)")),
					Default(),
				),
				html.Div(
					html.Label(g.Text("8-digit code")),
					New(Props{Length: 8}),
				),
			),
		),

		// Input Types
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Input Types")),
			html.Div(
				html.Class("space-y-4"),
				html.Div(
					html.Label(g.Text("Numeric (default)")),
					Default(),
				),
				html.Div(
					html.Label(g.Text("Alphanumeric")),
					Alphanumeric(),
				),
				html.Div(
					html.Label(g.Text("Custom pattern (letters only)")),
					New(Props{
						Pattern: "[a-zA-Z]",
						Type:    "custom",
					}),
				),
			),
		),

		// With Value
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Pre-filled Value")),
			html.Div(
				html.Class("space-y-4"),
				html.Label(g.Text("Partially filled")),
				New(Props{
					Value: "123",
				}),
			),
		),

		// Disabled State
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Disabled State")),
			html.Div(
				html.Class("space-y-4"),
				html.Label(g.Text("Disabled input")),
				New(Props{
					Disabled: true,
					Value:    "123456",
				}),
			),
		),

		// Custom Styling
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Custom Styling")),
			html.Div(
				html.Class("space-y-4"),
				html.Div(
					html.Label(g.Text("Larger inputs")),
					New(Props{
						Class: "[&_input]:w-14 [&_input]:h-14 [&_input]:text-lg",
					}),
				),
				html.Div(
					html.Label(g.Text("Custom colors")),
					New(Props{
						Class: "[&_input]:border-primary [&_input]:focus:ring-primary",
					}),
				),
			),
		),

		// With Separator
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("With Custom Separator")),
			html.Div(
				html.Class("space-y-4"),
				html.Label(g.Text("Grouped inputs")),
				WithSeparator(),
			),
		),

		// Form Integration
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Form Integration")),
			html.Form(
				html.Class("space-y-4"),
				html.Method("POST"),
				html.Action("/verify-otp"),
				html.Div(
					html.Label(html.For("verification-code"), g.Text("Verification Code")),
					New(Props{
						ID:        "verification-code",
						Name:      "code",
						AutoFocus: true,
					}),
				),
				html.Button(
					html.Type("submit"),
					html.Class("w-full bg-primary text-primary-foreground hover:bg-primary/90 h-10 px-4 py-2 rounded-md"),
					g.Text("Verify"),
				),
			),
		),

		// With Helper Text
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("With Helper Text")),
			html.Div(
				html.Class("space-y-2"),
				html.Label(g.Text("Enter code")),
				Default(),
				html.P(html.Class("text-sm text-muted-foreground"),
					g.Text("We sent a verification code to your email"),
				),
			),
		),

		// Error State
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Error State")),
			html.Div(
				html.Class("space-y-2"),
				html.Label(g.Text("Verification code")),
				New(Props{
					Class: "[&_input]:border-destructive [&_input]:focus:ring-destructive",
					Value: "123456",
				}),
				html.P(html.Class("text-sm text-destructive"),
					g.Text("Invalid verification code. Please try again."),
				),
			),
		),

		// With Completion Handler
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("With Completion Handler")),
			html.Div(
				html.Class("space-y-4"),
				html.Label(g.Text("Auto-submit on completion")),
				New(Props{
					OnComplete: "handleOTPComplete",
				}),
				html.Script(g.Raw(`
					function handleOTPComplete(value) {
						console.log('OTP Complete:', value);
						// You can auto-submit a form or make an API call here
						alert('OTP entered: ' + value);
					}
				`)),
			),
		),

		// Mobile-Optimized
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Mobile-Optimized")),
			html.Div(
				html.Class("space-y-2"),
				html.P(html.Class("text-sm text-muted-foreground"),
					g.Text("This input uses inputmode='numeric' for better mobile experience"),
				),
				Default(),
			),
		),

		// Two-Factor Authentication Example
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Two-Factor Authentication")),
			html.Div(
				html.Class("max-w-md p-6 border rounded-lg space-y-4"),
				html.Div(
					html.Class("space-y-2"),
					html.H4(html.Class("text-xl font-semibold"), g.Text("Enter authentication code")),
					html.P(html.Class("text-sm text-muted-foreground"),
						g.Text("Enter the 6-digit code from your authenticator app"),
					),
				),
				New(Props{
					Name:      "2fa_code",
					AutoFocus: true,
				}),
				html.Div(
					html.Class("flex gap-2"),
					html.Button(
						html.Type("submit"),
						html.Class("flex-1 bg-primary text-primary-foreground hover:bg-primary/90 h-10 px-4 py-2 rounded-md"),
						g.Text("Verify"),
					),
					html.Button(
						html.Type("button"),
						html.Class("flex-1 border border-input bg-background hover:bg-accent hover:text-accent-foreground h-10 px-4 py-2 rounded-md"),
						g.Text("Use backup code"),
					),
				),
			),
		),

		// SMS Verification Example
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("SMS Verification")),
			html.Div(
				html.Class("max-w-md p-6 border rounded-lg space-y-4"),
				html.Div(
					html.Class("space-y-2"),
					html.H4(html.Class("text-xl font-semibold"), g.Text("Verify your phone number")),
					html.P(html.Class("text-sm text-muted-foreground"),
						g.Text("We've sent a code to +1 (555) 123-4567"),
					),
				),
				New(Props{
					Name:      "sms_code",
					AutoFocus: true,
				}),
				html.P(html.Class("text-sm text-muted-foreground"),
					g.Text("Didn't receive the code? "),
					html.A(
						html.Href("#"),
						html.Class("text-primary hover:underline"),
						g.Text("Resend"),
					),
				),
			),
		),
	)
}