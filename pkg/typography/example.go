package typography

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

// Example demonstrates all typography components
func Example() g.Node {
	return html.Div(
		html.Class("p-8 space-y-12 max-w-4xl mx-auto"),

		// Headings section
		html.Div(
			html.Class("space-y-4"),
			H1(Props{}, g.Text("Typography Components")),
			Lead(Props{}, g.Text("Beautiful, consistent text styling for your Go applications")),
		),

		// Headings examples
		html.Div(
			html.Class("space-y-4"),
			H2(Props{}, g.Text("Headings")),
			P(Props{}, g.Text("Use heading components to create a clear content hierarchy.")),
			html.Div(
				html.Class("space-y-3 border rounded-lg p-6"),
				H1(Props{}, g.Text("H1: The quick brown fox")),
				H2(Props{}, g.Text("H2: Jumps over the lazy dog")),
				H3(Props{}, g.Text("H3: Pack my box with five dozen")),
				H4(Props{}, g.Text("H4: Liquor jugs for the trip")),
			),
		),

		// Text variants
		html.Div(
			html.Class("space-y-4"),
			H2(Props{}, g.Text("Text Variants")),
			P(Props{}, g.Text("Different text styles for various use cases.")),
			html.Div(
				html.Class("space-y-3 border rounded-lg p-6"),
				Lead(Props{}, g.Text("Lead text for introductions and important content")),
				P(Props{}, g.Text("Regular paragraph text forms the foundation of your content. It provides comfortable reading with appropriate line height and spacing.")),
				Large(Props{}, g.Text("Large text for emphasis")),
				Small(Props{}, g.Text("Small text for supporting information")),
				Muted(Props{}, g.Text("Muted text for secondary content")),
			),
		),

		// Lists
		html.Div(
			html.Class("space-y-4"),
			H2(Props{}, g.Text("Lists")),
			P(Props{}, g.Text("Organized content with styled lists.")),
			html.Div(
				html.Class("grid md:grid-cols-2 gap-6"),
				html.Div(
					H3(Props{}, g.Text("Unordered List")),
					List(Props{},
						ListItem(Props{}, g.Text("First item in the list")),
						ListItem(Props{}, g.Text("Second item with more content")),
						ListItem(Props{}, g.Text("Third item to complete the set")),
						ListItem(Props{}, 
							g.Text("Nested list example"),
							List(Props{},
								ListItem(Props{}, g.Text("Nested item one")),
								ListItem(Props{}, g.Text("Nested item two")),
							),
						),
					),
				),
				html.Div(
					H3(Props{}, g.Text("Ordered List")),
					OrderedList(Props{},
						ListItem(Props{}, g.Text("First step in the process")),
						ListItem(Props{}, g.Text("Second step to follow")),
						ListItem(Props{}, g.Text("Third step to complete")),
						ListItem(Props{}, g.Text("Final step for success")),
					),
				),
			),
		),

		// Blockquote
		html.Div(
			html.Class("space-y-4"),
			H2(Props{}, g.Text("Blockquote")),
			Blockquote(Props{},
				g.Text("\"The only way to do great work is to love what you do. If you haven't found it yet, keep looking. Don't settle.\""),
				html.Footer(
					html.Class("mt-2 text-sm"),
					g.Text("— Steve Jobs"),
				),
			),
		),

		// Code examples
		html.Div(
			html.Class("space-y-4"),
			H2(Props{}, g.Text("Code")),
			P(Props{}, 
				g.Text("Use "),
				InlineCode(Props{}, g.Text("InlineCode")),
				g.Text(" for inline code snippets and "),
				InlineCode(Props{}, g.Text("Code")),
				g.Text(" for code blocks."),
			),
			Code(Props{}, g.Text(`func Hello(name string) string {
    return fmt.Sprintf("Hello, %s!", name)
}`)),
		),

		// Table
		html.Div(
			html.Class("space-y-4"),
			H2(Props{}, g.Text("Table")),
			P(Props{}, g.Text("Display structured data in tables.")),
			Table(TableProps{Caption: "Example pricing table"},
				TableHeader(Props{},
					TableRow(Props{},
						TableHead(Props{}, g.Text("Plan")),
						TableHead(Props{}, g.Text("Price")),
						TableHead(Props{}, g.Text("Features")),
						TableHead(Props{Class: "text-right"}, g.Text("Action")),
					),
				),
				TableBody(Props{},
					TableRow(Props{},
						TableCell(Props{Class: "font-medium"}, g.Text("Free")),
						TableCell(Props{}, g.Text("$0/month")),
						TableCell(Props{}, g.Text("Basic features")),
						TableCell(Props{Class: "text-right"}, 
							Link("#", Props{}, g.Text("Get Started")),
						),
					),
					TableRow(Props{},
						TableCell(Props{Class: "font-medium"}, g.Text("Pro")),
						TableCell(Props{}, g.Text("$29/month")),
						TableCell(Props{}, g.Text("Advanced features")),
						TableCell(Props{Class: "text-right"}, 
							Link("#", Props{}, g.Text("Get Started")),
						),
					),
					TableRow(Props{},
						TableCell(Props{Class: "font-medium"}, g.Text("Enterprise")),
						TableCell(Props{}, g.Text("Custom")),
						TableCell(Props{}, g.Text("All features + support")),
						TableCell(Props{Class: "text-right"}, 
							Link("#", Props{}, g.Text("Contact Sales")),
						),
					),
				),
			),
		),

		// Links and special text
		html.Div(
			html.Class("space-y-4"),
			H2(Props{}, g.Text("Links and Special Text")),
			P(Props{}, 
				g.Text("Visit "),
				Link("#", Props{}, g.Text("our documentation")),
				g.Text(" or check out "),
				ExternalLink("https://github.com", Props{}, g.Text("our GitHub")),
				g.Text(" for more information."),
			),
			P(Props{},
				g.Text("You can "),
				Mark(Props{}, g.Text("highlight important text")),
				g.Text(" or show keyboard shortcuts like "),
				Kbd(Props{}, g.Text("⌘")),
				g.Text(" + "),
				Kbd(Props{}, g.Text("K")),
				g.Text("."),
			),
		),

		// Separator
		Hr(Props{}),

		// Combined example
		html.Div(
			html.Class("space-y-4"),
			H2(Props{}, g.Text("Real-World Example")),
			html.Article(
				html.Class("prose dark:prose-invert max-w-none"),
				H3(Props{}, g.Text("Getting Started with Go")),
				Lead(Props{}, g.Text("A comprehensive guide to starting your Go journey")),
				P(Props{}, g.Text("Go is a statically typed, compiled programming language designed at Google. It's known for its simplicity, efficiency, and excellent support for concurrent programming.")),
				
				H4(Props{}, g.Text("Installation")),
				P(Props{}, g.Text("To get started with Go, you'll need to install it on your system:")),
				OrderedList(Props{},
					ListItem(Props{}, 
						g.Text("Visit "),
						ExternalLink("https://golang.org/dl/", Props{}, g.Text("the official Go downloads page")),
					),
					ListItem(Props{}, g.Text("Download the installer for your operating system")),
					ListItem(Props{}, g.Text("Run the installer and follow the instructions")),
					ListItem(Props{}, 
						g.Text("Verify the installation by running "),
						InlineCode(Props{}, g.Text("go version")),
						g.Text(" in your terminal"),
					),
				),

				H4(Props{}, g.Text("Your First Program")),
				P(Props{}, g.Text("Create a new file called hello.go with the following content:")),
				Code(Props{}, g.Text(`package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}`)),
				P(Props{}, 
					g.Text("Run your program with "),
					InlineCode(Props{}, g.Text("go run hello.go")),
					g.Text("."),
				),

				Blockquote(Props{}, g.Text("\"Go will be the server language of the future.\" — Tobias Lütke, Shopify")),
			),
		),

		// Usage notes
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Usage Notes")),
			html.Div(html.Class("rounded-lg border bg-muted/50 p-4"),
				html.Ul(html.Class("text-sm space-y-2 list-disc list-inside"),
					html.Li(g.Text("Typography components provide consistent text styling")),
					html.Li(g.Text("All components support custom classes via the Class prop")),
					html.Li(g.Text("Components are designed to work well together")),
					html.Li(g.Text("Based on Tailwind CSS typography best practices")),
					html.Li(g.Text("Supports both light and dark modes")),
					html.Li(g.Text("Responsive sizing for different screen sizes")),
				),
			),
		),
	)
}