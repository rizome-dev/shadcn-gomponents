package pagination

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
	"fmt"
	"math"
)

// Props defines the properties for the Pagination component
type Props struct {
	CurrentPage  int    // Current active page (1-based)
	TotalPages   int    // Total number of pages
	ShowFirst    bool   // Show first page button
	ShowLast     bool   // Show last page button
	ShowPrevNext bool   // Show previous/next buttons (default: true)
	MaxVisible   int    // Maximum number of page buttons to show (default: 7)
	Class        string // Additional custom classes
}

// ContentProps defines properties for pagination content container
type ContentProps struct {
	Class string
}

// ItemProps defines properties for pagination items
type ItemProps struct {
	Class    string
	Active   bool // Whether this is the current page
	Disabled bool // Whether this item is disabled
}

// LinkProps defines properties for pagination links
type LinkProps struct {
	Href     string // URL for the page
	Page     int    // Page number
	Active   bool   // Whether this is the current page
	Disabled bool   // Whether this link is disabled
	Class    string // Additional custom classes
}

// EllipsisProps defines properties for ellipsis indicators
type EllipsisProps struct {
	Class string
}

// New creates a new Pagination component
func New(props Props, children ...g.Node) g.Node {
	// Set defaults
	if props.ShowPrevNext {
		props.ShowPrevNext = true
	}
	if props.MaxVisible == 0 {
		props.MaxVisible = 7
	}

	classes := lib.CN(
		"mx-auto flex w-full justify-center",
		props.Class,
	)

	return html.Nav(
		html.Role("navigation"),
		g.Attr("aria-label", "pagination"),
		html.Class(classes),
		g.Group(children),
	)
}

// Content creates the pagination content container
func ContentComponent(props ContentProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"flex flex-row items-center gap-1",
		props.Class,
	)

	return html.Ul(
		html.Class(classes),
		g.Group(children),
	)
}

// Item creates a pagination item
func Item(props ItemProps, children ...g.Node) g.Node {
	classes := lib.CN(
		props.Class,
	)

	attrs := []g.Node{
		html.Class(classes),
	}

	if props.Active {
		attrs = append(attrs, g.Attr("aria-current", "page"))
	}

	return html.Li(
		append(attrs, children...)...,
	)
}

// Link creates a pagination link
func LinkComponent(props LinkProps, children ...g.Node) g.Node {
	baseClasses := "inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50"
	
	variantClasses := ""
	if props.Active {
		variantClasses = "border border-input bg-background hover:bg-accent hover:text-accent-foreground"
	} else {
		variantClasses = "hover:bg-accent hover:text-accent-foreground"
	}

	sizeClasses := "h-10 w-10"

	classes := lib.CN(
		baseClasses,
		variantClasses,
		sizeClasses,
		props.Class,
	)

	attrs := []g.Node{
		html.Class(classes),
		html.Href(props.Href),
	}

	if props.Active {
		attrs = append(attrs, 
			g.Attr("aria-current", "page"),
			html.TabIndex("-1"), // Active page link shouldn't be tabbable
		)
	}

	if props.Disabled {
		attrs = append(attrs,
			g.Attr("aria-disabled", "true"),
			html.TabIndex("-1"),
		)
		// Replace href with # for disabled links
		attrs[1] = html.Href("#")
	}

	return html.A(
		append(attrs, children...)...,
	)
}

// PreviousButton creates a previous page button
func PreviousButton(href string, disabled bool) g.Node {
	return Item(
		ItemProps{},
	LinkComponent(
			LinkProps{
				Href:     href,
				Disabled: disabled,
				Class:    "gap-1 pl-2.5",
			},
			g.Raw(`<svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg" class="h-4 w-4">
				<path d="M8.84182 3.13514C9.04327 3.32401 9.05348 3.64042 8.86462 3.84188L5.43521 7.49991L8.86462 11.1579C9.05348 11.3594 9.04327 11.6758 8.84182 11.8647C8.64036 12.0535 8.32394 12.0433 8.13508 11.8419L4.38508 7.84188C4.20477 7.64955 4.20477 7.35027 4.38508 7.15794L8.13508 3.15794C8.32394 2.95648 8.64036 2.94628 8.84182 3.13514Z" fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"></path>
			</svg>`),
			html.Span(html.Class("sr-only"), g.Text("Go to previous page")),
			g.Text("Previous"),
		),
	)
}

// NextButton creates a next page button
func NextButton(href string, disabled bool) g.Node {
	return Item(
		ItemProps{},
	LinkComponent(
			LinkProps{
				Href:     href,
				Disabled: disabled,
				Class:    "gap-1 pr-2.5",
			},
			g.Text("Next"),
			html.Span(html.Class("sr-only"), g.Text("Go to next page")),
			g.Raw(`<svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg" class="h-4 w-4">
				<path d="M6.1584 3.13508C6.35985 2.94621 6.67627 2.95642 6.86514 3.15788L10.6151 7.15788C10.7954 7.3502 10.7954 7.64949 10.6151 7.84182L6.86514 11.8418C6.67627 12.0433 6.35985 12.0535 6.1584 11.8646C5.95694 11.6757 5.94673 11.3593 6.1356 11.1579L9.565 7.49985L6.1356 3.84182C5.94673 3.64036 5.95694 3.32394 6.1584 3.13508Z" fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"></path>
			</svg>`),
		),
	)
}

// FirstButton creates a first page button
func FirstButton(href string, disabled bool) g.Node {
	return Item(
		ItemProps{},
	LinkComponent(
			LinkProps{
				Href:     href,
				Disabled: disabled,
			},
			g.Raw(`<svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg" class="h-4 w-4">
				<path d="M6.85355 3.14645C7.04882 3.34171 7.04882 3.65829 6.85355 3.85355L3.70711 7H12.5C12.7761 7 13 7.22386 13 7.5C13 7.77614 12.7761 8 12.5 8H3.70711L6.85355 11.1464C7.04882 11.3417 7.04882 11.6583 6.85355 11.8536C6.65829 12.0488 6.34171 12.0488 6.14645 11.8536L2.14645 7.85355C1.95118 7.65829 1.95118 7.34171 2.14645 7.14645L6.14645 3.14645C6.34171 2.95118 6.65829 2.95118 6.85355 3.14645Z" fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"></path>
			</svg>`),
			html.Span(html.Class("sr-only"), g.Text("Go to first page")),
		),
	)
}

// LastButton creates a last page button
func LastButton(href string, disabled bool) g.Node {
	return Item(
		ItemProps{},
	LinkComponent(
			LinkProps{
				Href:     href,
				Disabled: disabled,
			},
			g.Raw(`<svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg" class="h-4 w-4">
				<path d="M8.14645 3.14645C8.34171 2.95118 8.65829 2.95118 8.85355 3.14645L12.8536 7.14645C13.0488 7.34171 13.0488 7.65829 12.8536 7.85355L8.85355 11.8536C8.65829 12.0488 8.34171 12.0488 8.14645 11.8536C7.95118 11.6583 7.95118 11.3417 8.14645 11.1464L11.2929 8H2.5C2.22386 8 2 7.77614 2 7.5C2 7.22386 2.22386 7 2.5 7H11.2929L8.14645 3.85355C7.95118 3.65829 7.95118 3.34171 8.14645 3.14645Z" fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"></path>
			</svg>`),
			html.Span(html.Class("sr-only"), g.Text("Go to last page")),
		),
	)
}

// Ellipsis creates an ellipsis indicator
func Ellipsis(props ...EllipsisProps) g.Node {
	var p EllipsisProps
	if len(props) > 0 {
		p = props[0]
	}

	classes := lib.CN(
		"flex h-9 w-9 items-center justify-center",
		p.Class,
	)

	return html.Li(
		g.Attr("aria-hidden", "true"),
		html.Span(
			html.Class(classes),
			g.Raw(`<svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg" class="h-4 w-4">
				<path d="M3.625 7.5C3.625 8.12132 3.12132 8.625 2.5 8.625C1.87868 8.625 1.375 8.12132 1.375 7.5C1.375 6.87868 1.87868 6.375 2.5 6.375C3.12132 6.375 3.625 6.87868 3.625 7.5ZM8.625 7.5C8.625 8.12132 8.12132 8.625 7.5 8.625C6.87868 8.625 6.375 8.12132 6.375 7.5C6.375 6.87868 6.87868 6.375 7.5 6.375C8.12132 6.375 8.625 6.87868 8.625 7.5ZM12.5 8.625C13.1213 8.625 13.625 8.12132 13.625 7.5C13.625 6.87868 13.1213 6.375 12.5 6.375C11.8787 6.375 11.375 6.87868 11.375 7.5C11.375 8.12132 11.8787 8.625 12.5 8.625Z" fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"></path>
			</svg>`),
			html.Span(html.Class("sr-only"), g.Text("More pages")),
		),
	)
}

// PageButton creates a numbered page button
func PageButton(page int, href string, active bool) g.Node {
	return Item(
		ItemProps{Active: active},
	LinkComponent(
			LinkProps{
				Href:   href,
				Page:   page,
				Active: active,
			},
			g.Text(fmt.Sprintf("%d", page)),
		),
	)
}

// Simple creates a simple pagination with just previous/next
func Simple(currentPage, totalPages int, getPageURL func(int) string) g.Node {
	return New(
		Props{},
	ContentComponent(
			ContentProps{},
			PreviousButton(
				getPageURL(currentPage-1),
				currentPage <= 1,
			),
			NextButton(
				getPageURL(currentPage+1),
				currentPage >= totalPages,
			),
		),
	)
}

// Default creates a default pagination with page numbers
func Default(currentPage, totalPages int, getPageURL func(int) string) g.Node {
	items := []g.Node{}

	// Previous button
	items = append(items, PreviousButton(
		getPageURL(currentPage-1),
		currentPage <= 1,
	))

	// Generate page numbers with ellipsis
	pages := generatePageNumbers(currentPage, totalPages, 7)
	
	for i, page := range pages {
		if page == -1 {
			items = append(items, Ellipsis())
		} else {
			items = append(items, PageButton(page, getPageURL(page), page == currentPage))
		}
		_ = i
	}

	// Next button
	items = append(items, NextButton(
		getPageURL(currentPage+1),
		currentPage >= totalPages,
	))

	return New(
		Props{
			CurrentPage: currentPage,
			TotalPages:  totalPages,
		},
	ContentComponent(
			ContentProps{},
			g.Group(items),
		),
	)
}

// WithFirstLast creates pagination with first/last buttons
func WithFirstLast(currentPage, totalPages int, getPageURL func(int) string) g.Node {
	items := []g.Node{}

	// First button
	items = append(items, FirstButton(
		getPageURL(1),
		currentPage <= 1,
	))

	// Previous button
	items = append(items, PreviousButton(
		getPageURL(currentPage-1),
		currentPage <= 1,
	))

	// Generate page numbers with ellipsis
	pages := generatePageNumbers(currentPage, totalPages, 5)
	
	for _, page := range pages {
		if page == -1 {
			items = append(items, Ellipsis())
		} else {
			items = append(items, PageButton(page, getPageURL(page), page == currentPage))
		}
	}

	// Next button
	items = append(items, NextButton(
		getPageURL(currentPage+1),
		currentPage >= totalPages,
	))

	// Last button
	items = append(items, LastButton(
		getPageURL(totalPages),
		currentPage >= totalPages,
	))

	return New(
		Props{
			CurrentPage: currentPage,
			TotalPages:  totalPages,
			ShowFirst:   true,
			ShowLast:    true,
		},
	ContentComponent(
			ContentProps{},
			g.Group(items),
		),
	)
}

// generatePageNumbers generates the page numbers to display with ellipsis
func generatePageNumbers(current, total, maxVisible int) []int {
	if total <= maxVisible {
		// Show all pages
		pages := make([]int, total)
		for i := 0; i < total; i++ {
			pages[i] = i + 1
		}
		return pages
	}

	// Calculate the range around current page
	halfVisible := maxVisible / 2
	pages := []int{}

	// Always show first page
	pages = append(pages, 1)

	// Calculate start and end of visible range
	start := int(math.Max(2, float64(current-halfVisible+1)))
	end := int(math.Min(float64(total-1), float64(current+halfVisible-1)))

	// Adjust range if at the edges
	if current <= halfVisible {
		end = maxVisible - 2
	} else if current > total-halfVisible {
		start = total - maxVisible + 3
	}

	// Add ellipsis if needed before range
	if start > 2 {
		pages = append(pages, -1) // -1 represents ellipsis
	}

	// Add visible range
	for i := start; i <= end; i++ {
		pages = append(pages, i)
	}

	// Add ellipsis if needed after range
	if end < total-1 {
		pages = append(pages, -1) // -1 represents ellipsis
	}

	// Always show last page
	if total > 1 {
		pages = append(pages, total)
	}

	return pages
}