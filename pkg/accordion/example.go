package accordion

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

// Example demonstrates a basic accordion usage
func Example() g.Node {
	return html.Div(
		html.Class("max-w-2xl mx-auto p-8"),
		html.H2(html.Class("text-2xl font-bold mb-6"), g.Text("Accordion Examples")),
		
		// Example 1: Single accordion with collapsible
		html.Section(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Single Accordion (Collapsible)")),
			New(Props{
				Type:         "single",
				Collapsible:  true,
				DefaultValue: "item-1",
				Class:        "mb-8",
			},
				Item(ItemProps{Value: "item-1"},
					Trigger(TriggerProps{}, g.Text("Product Information")),
					ItemContent(ContentProps{Class: "flex flex-col gap-4 text-balance"},
						html.P(g.Text("Our flagship product combines cutting-edge technology with sleek design. Built with premium materials, it offers unparalleled performance and reliability.")),
						html.P(g.Text("Key features include advanced processing capabilities, and an intuitive user interface designed for both beginners and experts.")),
					),
				),
				Item(ItemProps{Value: "item-2"},
					Trigger(TriggerProps{}, g.Text("Shipping Details")),
					ItemContent(ContentProps{Class: "flex flex-col gap-4 text-balance"},
						html.P(g.Text("We offer worldwide shipping through trusted courier partners. Standard delivery takes 3-5 business days, while express shipping ensures delivery within 1-2 business days.")),
						html.P(g.Text("All orders are carefully packaged and fully insured. Track your shipment in real-time through our dedicated tracking portal.")),
					),
				),
				Item(ItemProps{Value: "item-3"},
					Trigger(TriggerProps{}, g.Text("Return Policy")),
					ItemContent(ContentProps{Class: "flex flex-col gap-4 text-balance"},
						html.P(g.Text("We stand behind our products with a comprehensive 30-day return policy. If you're not completely satisfied, simply return the item in its original condition.")),
						html.P(g.Text("Our hassle-free return process includes free return shipping and full refunds processed within 48 hours of receiving the returned item.")),
					),
				),
			),
		),
		
		// Example 2: Single accordion non-collapsible
		html.Section(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Single Accordion (Always One Open)")),
			Single(false, "faq-1",
				Item(ItemProps{Value: "faq-1"},
					Trigger(TriggerProps{}, g.Text("What payment methods do you accept?")),
					ItemContent(ContentProps{},
						html.P(g.Text("We accept all major credit cards (Visa, MasterCard, American Express), PayPal, Apple Pay, and Google Pay. All transactions are secured with industry-standard encryption.")),
					),
				),
				Item(ItemProps{Value: "faq-2"},
					Trigger(TriggerProps{}, g.Text("Do you offer international shipping?")),
					ItemContent(ContentProps{},
						html.P(g.Text("Yes, we ship to over 180 countries worldwide. International shipping rates and delivery times vary by destination. You can check the exact cost during checkout.")),
					),
				),
				Item(ItemProps{Value: "faq-3"},
					Trigger(TriggerProps{}, g.Text("How can I track my order?")),
					ItemContent(ContentProps{},
						html.P(g.Text("Once your order ships, you'll receive a confirmation email with a tracking number. Click the tracking link in the email or enter the number on our website to track your package.")),
					),
				),
			),
		),
		
		// Example 3: Multiple accordion
		html.Section(
			html.Class("mt-8"),
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Multiple Accordion (Can Open Many)")),
			MultipleAccordion([]string{},
				Item(ItemProps{Value: "feature-1"},
					Trigger(TriggerProps{}, g.Text("Advanced Analytics")),
					ItemContent(ContentProps{},
						html.P(g.Text("Get detailed insights into your data with our advanced analytics dashboard. Track metrics, visualize trends, and make data-driven decisions.")),
					),
				),
				Item(ItemProps{Value: "feature-2"},
					Trigger(TriggerProps{}, g.Text("Team Collaboration")),
					ItemContent(ContentProps{},
						html.P(g.Text("Work seamlessly with your team using our collaboration tools. Share projects, assign tasks, and communicate in real-time.")),
					),
				),
				Item(ItemProps{Value: "feature-3"},
					Trigger(TriggerProps{}, g.Text("Security & Compliance")),
					ItemContent(ContentProps{},
						html.P(g.Text("Enterprise-grade security with SOC 2 Type II certification. Your data is encrypted at rest and in transit, with regular security audits.")),
					),
				),
			),
		),
		
		// Example 4: Styled accordion
		html.Section(
			html.Class("mt-8"),
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Styled Accordion")),
			New(Props{
				Type:        "single",
				Collapsible: true,
				Class:       "bg-muted/50 rounded-lg p-4",
			},
				Item(ItemProps{Value: "styled-1", Class: "border-muted-foreground/20"},
					Trigger(TriggerProps{Class: "hover:no-underline hover:bg-muted rounded px-2 -mx-2"}, 
						html.Span(html.Class("font-bold text-primary"), g.Text("Premium Features"))),
					ItemContent(ContentProps{Class: "text-muted-foreground px-2"},
						html.P(g.Text("Access exclusive premium features including advanced customization options, priority support, and early access to new releases.")),
					),
				),
				Item(ItemProps{Value: "styled-2", Class: "border-muted-foreground/20"},
					Trigger(TriggerProps{Class: "hover:no-underline hover:bg-muted rounded px-2 -mx-2"}, 
						html.Span(html.Class("font-bold text-primary"), g.Text("Developer API"))),
					ItemContent(ContentProps{Class: "text-muted-foreground px-2"},
						html.P(g.Text("Build powerful integrations with our comprehensive REST API. Full documentation, SDKs, and code examples available.")),
					),
				),
			),
		),
	)
}