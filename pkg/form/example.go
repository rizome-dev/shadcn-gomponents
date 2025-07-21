package form

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

// Example demonstrates the Form component
func Example() g.Node {
	return html.Div(html.Class("space-y-8"),
		// Basic Form placeholder
		html.Div(
			html.H3(html.Class("text-lg font-semibold"), g.Text("Basic Form")),
			html.P(html.Class("text-muted-foreground"), g.Text("Basic form example would go here")),
		),

		// Complete Registration Form
		html.Div(
			html.H3(html.Class("text-lg font-semibold"), g.Text("Registration Form")),
			ExampleRegistration(),
		),

		// Settings Form with Sections
		html.Div(
			html.H3(html.Class("text-lg font-semibold"), g.Text("Settings Form")),
			ExampleSettings(),
		),

		// Form with Input Groups
		html.Div(
			html.H3(html.Class("text-lg font-semibold"), g.Text("Form with Input Groups")),
			ExampleInputGroups(),
		),

		// Inline Form
		html.Div(
			html.H3(html.Class("text-lg font-semibold"), g.Text("Inline Form")),
			ExampleInline(),
		),
	)
}

// ExampleRegistration demonstrates a complete registration form
func ExampleRegistration() g.Node {
	return New(
		Props{Method: "post", Action: "/register", Class: "max-w-2xl"},
		Section("Personal Information", "Please provide your personal details.",
			FormRow(
				FormItem(
					ItemProps{},
					FormLabel(LabelProps{For: "firstName", Required: true}, g.Text("First Name")),
					FormControl(
						ControlProps{},
						html.Input(
							html.Type("text"),
							html.ID("firstName"),
							html.Name("firstName"),
							html.Placeholder("John"),
							html.Class("w-full"),
							html.Required(),
						),
					),
				),
				FormItem(
					ItemProps{},
					FormLabel(LabelProps{For: "lastName", Required: true}, g.Text("Last Name")),
					FormControl(
						ControlProps{},
						html.Input(
							html.Type("text"),
							html.ID("lastName"),
							html.Name("lastName"),
							html.Placeholder("Doe"),
							html.Class("w-full"),
							html.Required(),
						),
					),
				),
			),
			FormItem(
				ItemProps{},
				FormLabel(LabelProps{For: "email", Required: true}, g.Text("Email Address")),
				FormControl(
					ControlProps{},
					html.Input(
						html.Type("email"),
						html.ID("email"),
						html.Name("email"),
						html.Placeholder("john.doe@example.com"),
						html.Class("w-full"),
						html.Required(),
					),
				),
				FormDescription(
					DescriptionProps{},
					g.Text("We'll use this for account-related communications."),
				),
			),
			FormItem(
				ItemProps{},
				FormLabel(LabelProps{For: "phone"}, g.Text("Phone Number")),
				FormControl(
					ControlProps{},
					html.Input(
						html.Type("tel"),
						html.ID("phone"),
						html.Name("phone"),
						html.Placeholder("+1 (555) 123-4567"),
						html.Class("w-full"),
					),
				),
			),
		),
		
		Section("Account Security", "Create a strong password for your account.",
			FormItem(
				ItemProps{},
				FormLabel(LabelProps{For: "password", Required: true}, g.Text("Password")),
				FormControl(
					ControlProps{},
					html.Input(
						html.Type("password"),
						html.ID("password"),
						html.Name("password"),
						html.Class("w-full"),
						html.Required(),
					),
				),
				FormDescription(
					DescriptionProps{},
					g.Text("Must be at least 8 characters with one uppercase letter and one number."),
				),
			),
			FormItem(
				ItemProps{},
				FormLabel(LabelProps{For: "confirmPassword", Required: true}, g.Text("Confirm Password")),
				FormControl(
					ControlProps{},
					html.Input(
						html.Type("password"),
						html.ID("confirmPassword"),
						html.Name("confirmPassword"),
						html.Class("w-full"),
						html.Required(),
					),
				),
			),
		),
		
		FormFieldset(
			FieldsetProps{},
			FormLegend(LegendProps{}, g.Text("Terms and Conditions")),
			FormItem(
				ItemProps{Class: "flex flex-row items-start space-x-3 space-y-0"},
				FormControl(
					ControlProps{},
					html.Input(
						html.Type("checkbox"),
						html.ID("terms"),
						html.Name("terms"),
						html.Class("mt-1"),
						html.Required(),
					),
				),
				html.Div(html.Class("space-y-1 leading-none"),
					FormLabel(
						LabelProps{For: "terms", Required: true},
						g.Text("I accept the terms and conditions"),
					),
					FormDescription(
						DescriptionProps{},
						g.Text("By checking this box, you agree to our "),
						html.A(html.Href("#"), html.Class("underline"), g.Text("Terms of Service")),
						g.Text(" and "),
						html.A(html.Href("#"), html.Class("underline"), g.Text("Privacy Policy")),
						g.Text("."),
					),
				),
			),
			FormItem(
				ItemProps{Class: "flex flex-row items-start space-x-3 space-y-0"},
				FormControl(
					ControlProps{},
					html.Input(
						html.Type("checkbox"),
						html.ID("marketing"),
						html.Name("marketing"),
						html.Class("mt-1"),
					),
				),
				html.Div(html.Class("space-y-1 leading-none"),
					FormLabel(
						LabelProps{For: "marketing"},
						g.Text("Send me marketing emails"),
					),
					FormDescription(
						DescriptionProps{},
						g.Text("Receive updates about new features and promotions."),
					),
				),
			),
		),
		
		FormActions(
			ItemProps{},
			html.Button(
				html.Type("button"),
				html.Class("border hover:bg-accent"),
				g.Text("Cancel"),
			),
			html.Button(
				html.Type("submit"),
				html.Class("bg-primary text-primary-foreground hover:bg-primary/90"),
				g.Text("Create Account"),
			),
		),
	)
}

// ExampleSettings demonstrates a settings form
func ExampleSettings() g.Node {
	return New(
		Props{Method: "post", Action: "/settings", Class: "max-w-4xl"},
		Section("Profile", "This is how others will see you on the site.",
			FormItem(
				ItemProps{},
				FormLabel(LabelProps{For: "username", Required: true}, g.Text("Username")),
				FormControl(
					ControlProps{},
					html.Input(
						html.Type("text"),
						html.ID("username"),
						html.Name("username"),
						html.Value("johndoe"),
						html.Class("w-full max-w-md"),
					),
				),
				FormDescription(
					DescriptionProps{},
					g.Text("This is your public display name. It can be your real name or a pseudonym."),
				),
			),
			FormItem(
				ItemProps{},
				FormLabel(LabelProps{For: "bio"}, g.Text("Bio")),
				FormControl(
					ControlProps{},
					html.Textarea(
						html.ID("bio"),
						html.Name("bio"),
						html.Rows("4"),
						html.Class("w-full max-w-md"),
						html.Placeholder("Tell us a little bit about yourself"),
						g.Text("I'm a software developer who loves building web applications."),
					),
				),
				FormDescription(
					DescriptionProps{},
					g.Text("You can @mention other users and organizations to link to them."),
				),
			),
			FormItem(
				ItemProps{},
				FormLabel(LabelProps{For: "urls"}, g.Text("URLs")),
				FormDescription(
					DescriptionProps{},
					g.Text("Add links to your website, blog, or social media profiles."),
				),
				html.Div(html.Class("space-y-2 mt-2"),
					html.Input(
						html.Type("url"),
						html.Name("urls[]"),
						html.Placeholder("https://example.com"),
						html.Class("w-full max-w-md"),
					),
					html.Input(
						html.Type("url"),
						html.Name("urls[]"),
						html.Placeholder("https://twitter.com/johndoe"),
						html.Class("w-full max-w-md"),
					),
					html.Button(
						html.Type("button"),
						html.Class("text-sm border"),
						g.Text("Add another URL"),
					),
				),
			),
		),
		
		Section("Notifications", "Configure how you receive notifications.",
			FormFieldset(
				FieldsetProps{},
				FormLegend(LegendProps{}, g.Text("Email Notifications")),
				FormItem(
					ItemProps{Class: "flex flex-row items-center justify-between rounded-lg border p-3 shadow-sm"},
					html.Div(html.Class("space-y-0.5"),
						FormLabel(LabelProps{For: "marketing-emails"}, g.Text("Marketing emails")),
						FormDescription(
							DescriptionProps{},
							g.Text("Receive emails about new products, features, and more."),
						),
					),
					FormControl(
						ControlProps{},
						html.Input(
							html.Type("checkbox"),
							html.ID("marketing-emails"),
							html.Name("marketing-emails"),
							html.Class("h-4 w-4"),
						),
					),
				),
				FormItem(
					ItemProps{Class: "flex flex-row items-center justify-between rounded-lg border p-3 shadow-sm"},
					html.Div(html.Class("space-y-0.5"),
						FormLabel(LabelProps{For: "social-emails"}, g.Text("Social emails")),
						FormDescription(
							DescriptionProps{},
							g.Text("Receive emails for friend requests, follows, and more."),
						),
					),
					FormControl(
						ControlProps{},
						html.Input(
							html.Type("checkbox"),
							html.ID("social-emails"),
							html.Name("social-emails"),
							html.Class("h-4 w-4"),
							html.Checked(),
						),
					),
				),
				FormItem(
					ItemProps{Class: "flex flex-row items-center justify-between rounded-lg border p-3 shadow-sm"},
					html.Div(html.Class("space-y-0.5"),
						FormLabel(LabelProps{For: "security-emails"}, g.Text("Security emails")),
						FormDescription(
							DescriptionProps{},
							g.Text("Receive emails about your account security."),
						),
					),
					FormControl(
						ControlProps{},
						html.Input(
							html.Type("checkbox"),
							html.ID("security-emails"),
							html.Name("security-emails"),
							html.Class("h-4 w-4"),
							html.Checked(),
							html.Disabled(),
						),
					),
				),
			),
		),
		
		FormActions(
			ItemProps{Class: "border-t pt-6"},
			html.Button(
				html.Type("submit"),
				html.Class("bg-primary text-primary-foreground hover:bg-primary/90"),
				g.Text("Update Settings"),
			),
		),
	)
}

// ExampleInputGroups demonstrates forms with input groups
func ExampleInputGroups() g.Node {
	return New(
		Props{Class: "max-w-lg"},
		FormItem(
			ItemProps{},
			FormLabel(LabelProps{For: "website"}, g.Text("Website")),
			FormControl(
				ControlProps{},
				InputGroup(
					g.Text("https://"),
					html.Input(
						html.Type("text"),
						html.ID("website"),
						html.Name("website"),
						html.Placeholder("example.com"),
						html.Class("flex-1 rounded-none"),
					),
					nil,
				),
			),
		),
		FormItem(
			ItemProps{},
			FormLabel(LabelProps{For: "price"}, g.Text("Price")),
			FormControl(
				ControlProps{},
				InputGroup(
					g.Text("$"),
					html.Input(
						html.Type("number"),
						html.ID("price"),
						html.Name("price"),
						html.Placeholder("0.00"),
						html.Step("0.01"),
						html.Class("flex-1 rounded-none"),
					),
					g.Text("USD"),
				),
			),
		),
		FormItem(
			ItemProps{},
			FormLabel(LabelProps{For: "username"}, g.Text("Username")),
			FormControl(
				ControlProps{},
				InputGroup(
					g.Text("@"),
					html.Input(
						html.Type("text"),
						html.ID("username"),
						html.Name("username"),
						html.Placeholder("username"),
						html.Class("flex-1 rounded-l-none"),
					),
					nil,
				),
			),
			FormDescription(
				DescriptionProps{},
				g.Text("Choose a unique username."),
			),
		),
		html.Button(
			html.Type("submit"),
			html.Class("w-full"),
			g.Text("Submit"),
		),
	)
}

// ExampleInline demonstrates an inline form
func ExampleInline() g.Node {
	return New(
		Props{Class: "flex flex-col sm:flex-row gap-4 items-end"},
		FormItem(
			ItemProps{Class: "flex-1 space-y-1"},
			FormLabel(LabelProps{For: "email-inline"}, g.Text("Email")),
			FormControl(
				ControlProps{},
				html.Input(
					html.Type("email"),
					html.ID("email-inline"),
					html.Name("email"),
					html.Placeholder("Enter your email"),
					html.Class("w-full"),
				),
			),
		),
		FormItem(
			ItemProps{Class: "flex-1 space-y-1"},
			FormLabel(LabelProps{For: "name-inline"}, g.Text("Name")),
			FormControl(
				ControlProps{},
				html.Input(
					html.Type("text"),
					html.ID("name-inline"),
					html.Name("name"),
					html.Placeholder("Enter your name"),
					html.Class("w-full"),
				),
			),
		),
		html.Button(
			html.Type("submit"),
			html.Class("bg-primary text-primary-foreground hover:bg-primary/90"),
			g.Text("Subscribe"),
		),
	)
}