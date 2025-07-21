package inputotp

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
	"fmt"
)

// Props defines the properties for the InputOTP component
type Props struct {
	ID           string   // Unique ID for the input group
	Length       int      // Number of input fields (default: 6)
	Type         string   // "numeric" | "alphanumeric" (default: "numeric")
	Pattern      string   // Custom pattern for validation
	Name         string   // Name attribute for form submission
	Value        string   // Initial value
	Disabled     bool     // Whether the input is disabled
	AutoFocus    bool     // Whether to auto-focus the first input
	OnComplete   string   // JavaScript function to call when all inputs are filled
	Class        string   // Additional custom classes
	Placeholder  string   // Placeholder for each input (default: "○")
}

// GroupProps defines properties for the OTP input group container
type GroupProps struct {
	Class string
}

// SlotProps defines properties for individual OTP slots
type SlotProps struct {
	Index       int
	IsActive    bool
	HasValue    bool
	Class       string
}

// SeparatorProps defines properties for the separator between slots
type SeparatorProps struct {
	Class string
}

// New creates a new InputOTP component
func New(props Props, children ...g.Node) g.Node {
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

	// Build attributes
	attrs := []g.Node{
		html.Class(classes),
		g.If(props.ID != "", html.ID(props.ID)),
		g.Attr("data-otp-container", "true"),
		g.Attr("data-otp-length", fmt.Sprintf("%d", props.Length)),
	}

	// Create input fields
	var inputs []g.Node
	for i := 0; i < props.Length; i++ {
		inputAttrs := []g.Node{
			html.Type(inputType),
			html.Class("w-10 h-12 text-center text-sm font-medium border border-input bg-background rounded-md focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"),
			g.Attr("maxlength", "1"),
			g.If(pattern != "", g.Attr("pattern", pattern)),
			g.If(inputMode != "", g.Attr("inputmode", inputMode)),
			g.If(props.Name != "", html.Name(fmt.Sprintf("%s[%d]", props.Name, i))),
			g.If(props.Disabled, html.Disabled()),
			g.If(props.AutoFocus && i == 0, html.AutoFocus()),
			html.Placeholder(props.Placeholder),
			g.Attr("data-otp-input", fmt.Sprintf("%d", i)),
			g.Attr("autocomplete", "one-time-code"),
			
			// JavaScript for handling input and navigation
			g.Attr("oninput", `
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
				
				// Check if all inputs are filled
				const allFilled = Array.from(inputs).every(i => i.value);
				if (allFilled && container.dataset.otpComplete) {
					const completeValue = Array.from(inputs).map(i => i.value).join('');
					window[container.dataset.otpComplete](completeValue);
				}
			`),
			
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

	// Add onComplete handler if provided
	if props.OnComplete != "" {
		attrs = append(attrs, g.Attr("data-otp-complete", props.OnComplete))
	}

	return html.Div(
		append(attrs, g.Group(inputs), g.Group(children))...,
	)
}

// Group creates a container for OTP inputs with custom styling
func Group(props GroupProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"flex items-center",
		props.Class,
	)

	return html.Div(
		html.Class(classes),
		g.Group(children),
	)
}

// Slot creates a single OTP input slot with enhanced styling
func Slot(props SlotProps) g.Node {
	classes := lib.CN(
		"relative w-10 h-12 text-center text-sm transition-all",
		"border rounded-md",
		lib.CNIf(props.IsActive,
			"border-ring ring-2 ring-ring ring-offset-2",
			"border-input",
		),
		lib.CNIf(props.HasValue,
			"text-foreground",
			"text-muted-foreground",
		),
		props.Class,
	)

	return html.Div(
		html.Class(classes),
		html.Input(
			html.Type("text"),
			html.Class("absolute inset-0 w-full h-full text-center bg-transparent border-0 outline-none"),
			g.Attr("maxlength", "1"),
			g.Attr("data-otp-input", fmt.Sprintf("%d", props.Index)),
		),
	)
}

// Separator creates a visual separator between OTP input groups
func Separator(props ...SeparatorProps) g.Node {
	var p SeparatorProps
	if len(props) > 0 {
		p = props[0]
	}

	classes := lib.CN(
		"mx-1",
		p.Class,
	)

	return html.Div(
		html.Class(classes),
		g.Text("-"),
	)
}

// Default creates a default 6-digit numeric OTP input
func Default() g.Node {
	return New(Props{})
}

// FourDigit creates a 4-digit numeric OTP input
func FourDigit() g.Node {
	return New(Props{
		Length: 4,
	})
}

// Alphanumeric creates an alphanumeric OTP input
func Alphanumeric() g.Node {
	return New(Props{
		Type: "alphanumeric",
	})
}

// WithSeparator creates an OTP input with custom separator
func WithSeparator() g.Node {
	return html.Div(
		html.Class("flex items-center gap-2"),
		Group(
			GroupProps{},
			Slot(SlotProps{Index: 0}),
			Slot(SlotProps{Index: 1}),
			Slot(SlotProps{Index: 2}),
		),
		Separator(),
		Group(
			GroupProps{},
			Slot(SlotProps{Index: 3}),
			Slot(SlotProps{Index: 4}),
			Slot(SlotProps{Index: 5}),
		),
	)
}