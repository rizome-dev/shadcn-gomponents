package inputotp

import (
	"fmt"
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	hx "maragu.dev/gomponents-htmx"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// HTMXProps defines HTMX-specific properties for the InputOTP
type HTMXProps struct {
	ID           string // Unique ID for the input group
	VerifyPath   string // Server path for OTP verification
	Target       string // HTMX target for responses
	Indicator    string // ID of loading indicator
	TriggerDelay string // Delay before triggering verification (default: "500ms")
	SwapOOB      bool   // Whether to use out-of-band swapping
}

// NewHTMX creates an HTMX-enhanced InputOTP component
func NewHTMX(props Props, htmxProps HTMXProps, children ...g.Node) g.Node {
	// Set defaults
	if props.Length == 0 {
		props.Length = 6
	}
	if props.Type == "" {
		props.Type = "numeric"
	}
	if props.Placeholder == "" {
		props.Placeholder = "○"
	}
	if htmxProps.TriggerDelay == "" {
		htmxProps.TriggerDelay = "500ms"
	}
	if htmxProps.Target == "" {
		htmxProps.Target = "#" + htmxProps.ID + "-feedback"
	}

	// Build container classes
	classes := lib.CN(
		"flex items-center gap-2",
		props.Class,
	)

	// Generate input pattern based on type
	pattern := props.Pattern
	inputType := "text"
	inputMode := ""
	
	if pattern == "" {
		switch props.Type {
		case "numeric":
			pattern = "[0-9]"
			inputMode = "numeric"
		case "alphanumeric":
			pattern = "[a-zA-Z0-9]"
		}
	}

	// Build container attributes
	containerAttrs := []g.Node{
		html.Class(classes),
		html.ID(htmxProps.ID),
		g.Attr("data-otp-container", "true"),
		g.Attr("data-otp-length", fmt.Sprintf("%d", props.Length)),
	}

	// Create input fields with HTMX
	var inputs []g.Node
	for i := 0; i < props.Length; i++ {
		inputAttrs := []g.Node{
			html.Type(inputType),
			html.Class("w-10 h-12 text-center text-sm font-medium border border-input bg-background rounded-md focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 transition-colors"),
			g.Attr("maxlength", "1"),
			g.If(pattern != "", g.Attr("pattern", pattern)),
			g.If(inputMode != "", g.Attr("inputmode", inputMode)),
			g.If(props.Name != "", html.Name(fmt.Sprintf("%s[%d]", props.Name, i))),
			g.If(props.Disabled, html.Disabled()),
			g.If(props.AutoFocus && i == 0, html.AutoFocus()),
			html.Placeholder(props.Placeholder),
			g.Attr("data-otp-input", fmt.Sprintf("%d", i)),
			g.Attr("autocomplete", "one-time-code"),
			
			// HTMX attributes for the last input
			g.If(i == props.Length-1 && htmxProps.VerifyPath != "",
				g.Group([]g.Node{
					hx.Post(htmxProps.VerifyPath),
					hx.Trigger(fmt.Sprintf("input changed delay:%s", htmxProps.TriggerDelay)),
					hx.Target(htmxProps.Target),
					hx.Include(fmt.Sprintf("#%s input", htmxProps.ID)),
					g.If(htmxProps.Indicator != "", hx.Indicator("#"+htmxProps.Indicator)),
					g.If(htmxProps.SwapOOB, hx.SwapOOB("true")),
				}),
			),
			
			// JavaScript for handling input and navigation with HTMX awareness
			g.Attr("oninput", fmt.Sprintf(`
				const input = event.target;
				const value = input.value;
				const container = input.closest('[data-otp-container]');
				const inputs = container.querySelectorAll('[data-otp-input]');
				const currentIndex = parseInt(input.dataset.otpInput);
				
				// Only allow valid characters
				if (value && !input.checkValidity()) {
					input.value = '';
					return;
				}
				
				// Move to next input if value entered
				if (value && currentIndex < inputs.length - 1) {
					inputs[currentIndex + 1].focus();
				}
				
				// Check if all inputs are filled for HTMX trigger
				const allFilled = Array.from(inputs).every(i => i.value);
				if (allFilled && currentIndex === inputs.length - 1) {
					// HTMX will handle the verification automatically
					container.dataset.otpComplete = 'true';
				}
			`)),
			
			// Handle backspace navigation
			g.Attr("onkeydown", `
				const input = event.target;
				const container = input.closest('[data-otp-container]');
				const inputs = container.querySelectorAll('[data-otp-input]');
				const currentIndex = parseInt(input.dataset.otpInput);
				
				if (event.key === 'Backspace' && !input.value && currentIndex > 0) {
					event.preventDefault();
					inputs[currentIndex - 1].focus();
					inputs[currentIndex - 1].value = '';
					container.dataset.otpComplete = 'false';
				} else if (event.key === 'ArrowLeft' && currentIndex > 0) {
					event.preventDefault();
					inputs[currentIndex - 1].focus();
				} else if (event.key === 'ArrowRight' && currentIndex < inputs.length - 1) {
					event.preventDefault();
					inputs[currentIndex + 1].focus();
				}
			`),
			
			// Handle paste
			g.Attr("onpaste", `
				event.preventDefault();
				const paste = (event.clipboardData || window.clipboardData).getData('text');
				const container = event.target.closest('[data-otp-container]');
				const inputs = container.querySelectorAll('[data-otp-input]');
				const currentIndex = parseInt(event.target.dataset.otpInput);
				
				// Fill inputs starting from current position
				for (let i = 0; i < paste.length && currentIndex + i < inputs.length; i++) {
					const char = paste[i];
					const targetInput = inputs[currentIndex + i];
					
					// Validate character against pattern
					targetInput.value = char;
					if (!targetInput.checkValidity()) {
						targetInput.value = '';
					}
				}
				
				// Focus last filled input or next empty one
				const lastFilledIndex = Array.from(inputs).findLastIndex(i => i.value);
				if (lastFilledIndex >= 0 && lastFilledIndex < inputs.length - 1) {
					inputs[Mathtml.min(lastFilledIndex + 1, inputs.length - 1)].focus();
				}
				
				// Trigger HTMX if all filled
				const allFilled = Array.from(inputs).every(i => i.value);
				if (allFilled) {
					container.dataset.otpComplete = 'true';
					// Trigger HTMX on the last input
					inputs[inputs.length - 1].dispatchEvent(new Event('input', { bubbles: true }));
				}
			`),
		}

		// Set initial value if provided
		if props.Value != "" && i < len(props.Value) {
			inputAttrs = append(inputAttrs, html.Value(string(props.Value[i])))
		}

		inputs = append(inputs, html.Input(inputAttrs...))

		// Add separator in the middle (for 6 digits, after 3rd input)
		if props.Length == 6 && i == 2 {
			inputs = append(inputs, Separator())
		}
	}

	// Add feedback container
	feedbackContainer := html.Div(
		html.ID(htmxProps.ID+"-feedback"),
		html.Class("mt-2"),
	)

	// Add loading indicator if specified
	var loadingIndicator g.Node
	if htmxProps.Indicator != "" {
		loadingIndicator = html.Div(
			html.ID(htmxProps.Indicator),
			html.Class("htmx-indicator inline-flex items-center gap-2 text-sm text-muted-foreground"),
			g.Raw(`<svg class="animate-spin h-4 w-4" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
				<circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
				<path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
			</svg>`),
			g.Text("Verifying..."),
		)
	}

	return html.Div(
		html.Div(
			append(containerAttrs, g.Group(inputs))...,
		),
		feedbackContainer,
		g.If(htmxProps.Indicator != "", loadingIndicator),
		g.Group(children),
	)
}

// VerificationFeedback creates a feedback message for OTP verification
func VerificationFeedback(success bool, message string, htmxProps HTMXProps) g.Node {
	classes := lib.CN(
		"text-sm font-medium",
		lib.CNIf(success,
			"text-green-600 dark:text-green-500",
			"text-destructive",
		),
	)

	attrs := []g.Node{
		html.Class(classes),
	}

	// Add swap out-of-band if needed
	if htmxProps.SwapOOB {
		attrs = append(attrs, 
			html.ID(htmxProps.ID+"-feedback"),
			hx.SwapOOB("true"),
		)
	}

	return html.Div(
		append(attrs, g.Text(message))...,
	)
}

// ResendButton creates a button to resend OTP code
func ResendButton(htmxProps HTMXProps, resendPath string, cooldownSeconds int) g.Node {
	return html.Button(
		html.Type("button"),
		html.Class("text-sm text-primary hover:underline disabled:opacity-50 disabled:cursor-not-allowed"),
		hx.Post(resendPath),
		hx.Target("#"+htmxProps.ID+"-feedback"),
		hx.Swap("innerHTML"),
		g.Attr("onclick", fmt.Sprintf(`
			const btn = this;
			btn.disabled = true;
			let seconds = %d;
			const originalText = btn.innerText;
			
			const timer = setInterval(() => {
				btn.innerText = 'Resend in ' + seconds + 's';
				seconds--;
				
				if (seconds < 0) {
					clearInterval(timer);
					btn.disabled = false;
					btn.innerText = originalText;
				}
			}, 1000);
		`, cooldownSeconds)),
		g.Text("Resend code"),
	)
}

// ExampleHTMX creates an HTMX-enhanced OTP input example
func ExampleHTMX() g.Node {
	htmxProps := HTMXProps{
		ID:         "otp-example",
		VerifyPath: "/api/otp/verify",
		Indicator:  "otp-loading",
	}

	return html.Div(
		html.Class("space-y-4"),
		html.Label(html.For(htmxProps.ID), g.Text("Enter verification code")),
		NewHTMX(
			Props{
				Name:      "otp",
				AutoFocus: true,
			},
			htmxProps,
		),
		html.Div(
			html.Class("flex items-center justify-between"),
			html.P(html.Class("text-sm text-muted-foreground"),
				g.Text("Didn't receive a code? "),
				ResendButton(htmxProps, "/api/otp/resend", 30),
			),
		),
	)
}

// TwoFactorExampleHTMX creates a two-factor authentication OTP example
func TwoFactorExampleHTMX() g.Node {
	htmxProps := HTMXProps{
		ID:           "2fa-otp",
		VerifyPath:   "/api/2fa/verify",
		Indicator:    "2fa-loading",
		TriggerDelay: "300ms",
	}

	return html.Div(
		html.Class("max-w-md mx-auto space-y-6"),
		html.Div(
			html.Class("text-center space-y-2"),
			html.H2(html.Class("text-2xl font-bold"), g.Text("Two-factor authentication")),
			html.P(html.Class("text-muted-foreground"),
				g.Text("Enter the 6-digit code from your authenticator app"),
			),
		),
		NewHTMX(
			Props{
				Name:      "twoFactorCode",
				AutoFocus: true,
			},
			htmxProps,
		),
		html.Div(
			html.Class("space-y-2"),
			html.Button(
				html.Type("button"),
				html.Class("w-full"),
				g.Text("Verify and sign in"),
			),
			html.Button(
				html.Type("button"),
				html.Class("w-full variant-outline"),
				g.Text("Use backup code instead"),
			),
		),
	)
}

// RenderVerificationSuccess renders a success response for HTMX
func RenderVerificationSuccess(htmxProps HTMXProps) g.Node {
	return html.Div(
		html.ID(htmxProps.ID+"-feedback"),
		hx.SwapOOB("true"),
		html.Class("flex items-center gap-2 text-sm text-green-600 dark:text-green-500"),
		g.Raw(`<svg class="h-4 w-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
			<path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"></path>
			<polyline points="22 4 12 14.01 9 11.01"></polyline>
		</svg>`),
		g.Text("Verification successful!"),
	)
}

// RenderVerificationError renders an error response for HTMX
func RenderVerificationError(htmxProps HTMXProps, errorMessage string) g.Node {
	return html.Div(
		html.ID(htmxProps.ID+"-feedback"),
		hx.SwapOOB("true"),
		html.Class("flex items-center gap-2 text-sm text-destructive"),
		g.Raw(`<svg class="h-4 w-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
			<circle cx="12" cy="12" r="10"></circle>
			<line x1="12" y1="8" x2="12" y2="12"></line>
			<line x1="12" y1="16" x2="12.01" y2="16"></line>
		</svg>`),
		g.Text(errorMessage),
	)
}