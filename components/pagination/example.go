package pagination

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"fmt"
)

// Examples demonstrates various Pagination usage patterns
func Examples() g.Node {
	// Helper function to generate page URLs
	getPageURL := func(page int) string {
		return fmt.Sprintf("/page/%d", page)
	}

	return html.Div(
		html.Class("space-y-8 p-8"),
		
		// Basic Pagination
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Basic Pagination")),
			Default(1, 10, getPageURL),
		),

		// Simple Pagination (Previous/Next only)
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Simple Pagination")),
			Simple(5, 10, getPageURL),
		),

		// With First/Last Buttons
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("With First/Last Buttons")),
			WithFirstLast(5, 20, getPageURL),
		),

		// Different Page States
		html.Div(
			html.Class("space-y-4"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Different States")),
			
			html.Div(
				html.Class("space-y-2"),
				html.P(html.Class("text-sm text-muted-foreground"), g.Text("First page active:")),
				Default(1, 10, getPageURL),
			),
			
			html.Div(
				html.Class("space-y-2"),
				html.P(html.Class("text-sm text-muted-foreground"), g.Text("Middle page active:")),
				Default(5, 10, getPageURL),
			),
			
			html.Div(
				html.Class("space-y-2"),
				html.P(html.Class("text-sm text-muted-foreground"), g.Text("Last page active:")),
				Default(10, 10, getPageURL),
			),
		),

		// Large Number of Pages
		html.Div(
			html.Class("space-y-4"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Many Pages")),
			
			html.Div(
				html.Class("space-y-2"),
				html.P(html.Class("text-sm text-muted-foreground"), g.Text("100 pages, current page 1:")),
				Default(1, 100, getPageURL),
			),
			
			html.Div(
				html.Class("space-y-2"),
				html.P(html.Class("text-sm text-muted-foreground"), g.Text("100 pages, current page 50:")),
				Default(50, 100, getPageURL),
			),
			
			html.Div(
				html.Class("space-y-2"),
				html.P(html.Class("text-sm text-muted-foreground"), g.Text("100 pages, current page 100:")),
				Default(100, 100, getPageURL),
			),
		),

		// Custom Implementation
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Custom Implementation")),
			New(
				Props{},
	ContentComponent(
					ContentProps{},
					// Custom previous button with text
					Item(
						ItemProps{},
	LinkComponent(
							LinkProps{
								Href:     "/page/4",
								Disabled: false,
								Class:    "w-auto px-3",
							},
							g.Text("← Previous"),
						),
					),
					// Page numbers
					PageButton(3, "/page/3", false),
					PageButton(4, "/page/4", false),
					PageButton(5, "/page/5", true),
					PageButton(6, "/page/6", false),
					PageButton(7, "/page/7", false),
					// Custom next button with text
					Item(
						ItemProps{},
	LinkComponent(
							LinkProps{
								Href:  "/page/6",
								Class: "w-auto px-3",
							},
							g.Text("Next →"),
						),
					),
				),
			),
		),

		// With Page Info
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("With Page Info")),
			html.Div(
				html.Class("flex flex-col items-center gap-4"),
				Default(3, 10, getPageURL),
				html.P(html.Class("text-sm text-muted-foreground"),
					g.Text("Page 3 of 10"),
				),
			),
		),

		// Minimal Pagination
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Minimal")),
			New(
				Props{},
	ContentComponent(
					ContentProps{Class: "gap-4"},
					Item(
						ItemProps{},
	LinkComponent(
							LinkProps{
								Href: "/page/2",
								Class: "gap-1 pl-0 hover:bg-transparent hover:text-primary",
							},
							g.Raw(`<svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg" class="h-4 w-4">
								<path d="M8.84182 3.13514C9.04327 3.32401 9.05348 3.64042 8.86462 3.84188L5.43521 7.49991L8.86462 11.1579C9.05348 11.3594 9.04327 11.6758 8.84182 11.8647C8.64036 12.0535 8.32394 12.0433 8.13508 11.8419L4.38508 7.84188C4.20477 7.64955 4.20477 7.35027 4.38508 7.15794L8.13508 3.15794C8.32394 2.95648 8.64036 2.94628 8.84182 3.13514Z" fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"></path>
							</svg>`),
							g.Text("Previous"),
						),
					),
					html.Li(html.Class("text-sm"), g.Text("Page 3 of 10")),
					Item(
						ItemProps{},
	LinkComponent(
							LinkProps{
								Href: "/page/4",
								Class: "gap-1 pr-0 hover:bg-transparent hover:text-primary",
							},
							g.Text("Next"),
							g.Raw(`<svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg" class="h-4 w-4">
								<path d="M6.1584 3.13508C6.35985 2.94621 6.67627 2.95642 6.86514 3.15788L10.6151 7.15788C10.7954 7.3502 10.7954 7.64949 10.6151 7.84182L6.86514 11.8418C6.67627 12.0433 6.35985 12.0535 6.1584 11.8646C5.95694 11.6757 5.94673 11.3593 6.1356 11.1579L9.565 7.49985L6.1356 3.84182C5.94673 3.64036 5.95694 3.32394 6.1584 3.13508Z" fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"></path>
							</svg>`),
						),
					),
				),
			),
		),

		// Small Pagination (for tight spaces)
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Small Size")),
			New(
				Props{Class: "text-xs"},
	ContentComponent(
					ContentProps{},
					Item(
						ItemProps{},
	LinkComponent(
							LinkProps{
								Href:     "/page/1",
								Disabled: true,
								Class:    "h-8 w-8",
							},
							g.Text("←"),
						),
					),
					Item(
						ItemProps{},
	LinkComponent(
							LinkProps{
								Href:   "/page/1",
								Active: true,
								Class:  "h-8 w-8 text-xs",
							},
							g.Text("1"),
						),
					),
					Item(
						ItemProps{},
	LinkComponent(
							LinkProps{
								Href:  "/page/2",
								Class: "h-8 w-8 text-xs",
							},
							g.Text("2"),
						),
					),
					Item(
						ItemProps{},
	LinkComponent(
							LinkProps{
								Href:  "/page/3",
								Class: "h-8 w-8 text-xs",
							},
							g.Text("3"),
						),
					),
					Item(
						ItemProps{},
	LinkComponent(
							LinkProps{
								Href:  "/page/2",
								Class: "h-8 w-8",
							},
							g.Text("→"),
						),
					),
				),
			),
		),

		// With Results Info
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("With Results Info")),
			html.Div(
				html.Class("space-y-4"),
				html.Div(
					html.Class("flex items-center justify-between"),
					html.P(html.Class("text-sm text-muted-foreground"),
						g.Text("Showing 21-30 of 97 results"),
					),
					Simple(3, 10, getPageURL),
				),
			),
		),

		// Centered with Border
		html.Div(
			html.Class("space-y-2"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Styled Variations")),
			html.Div(
				html.Class("rounded-lg border p-4"),
				Default(5, 10, getPageURL),
			),
		),
	)
}