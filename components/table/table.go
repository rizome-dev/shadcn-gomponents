package table

import (
	"fmt"

	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// Props defines common properties for table components
type Props struct {
	Class string
	ID    string
}

// Table creates a table container with horizontal scroll
func TableComponent(props Props, children ...g.Node) g.Node {
	return html.Div(
		g.Attr("data-table-container", "true"),
		html.Class("relative w-full overflow-x-auto"),
		g.If(props.ID != "", g.Attr("id", props.ID+"-container")),
		g.El("table",
			g.Attr("data-table", "true"),
			html.Class(lib.CN("w-full caption-bottom text-sm", props.Class)),
			g.If(props.ID != "", g.Attr("id", props.ID)),
			g.Group(children),
		),
	)
}

// Header creates a table header
func HeaderComponent(props Props, children ...g.Node) g.Node {
	return g.El("thead",
		g.Attr("data-table-header", "true"),
		html.Class(lib.CN("[&_tr]:border-b", props.Class)),
		g.Group(children),
	)
}

// Body creates a table body
func Body(props Props, children ...g.Node) g.Node {
	return g.El("tbody",
		g.Attr("data-table-body", "true"),
		html.Class(lib.CN("[&_tr:last-child]:border-0", props.Class)),
		g.Group(children),
	)
}

// Footer creates a table footer
func FooterComponent(props Props, children ...g.Node) g.Node {
	return g.El("tfoot",
		g.Attr("data-table-footer", "true"),
		html.Class(lib.CN(
			"bg-muted/50 border-t font-medium [&>tr]:last:border-b-0",
			props.Class,
		)),
		g.Group(children),
	)
}

// RowProps defines properties for table rows
type RowProps struct {
	Class    string
	ID       string
	Selected bool
	OnClick  string
}

// Row creates a table row
func Row(props RowProps, children ...g.Node) g.Node {
	attrs := []g.Node{
		g.Attr("data-table-row", "true"),
		html.Class(lib.CN(
			"hover:bg-muted/50 data-[state=selected]:bg-muted border-b transition-colors",
			props.Class,
		)),
	}

	if props.ID != "" {
		attrs = append(attrs, g.Attr("id", props.ID))
	}

	if props.Selected {
		attrs = append(attrs, g.Attr("data-state", "selected"))
	}

	if props.OnClick != "" {
		attrs = append(attrs, g.Raw(props.OnClick))
	}

	return g.El("tr",
		g.Group(append(attrs, children...)),
	)
}

// HeadProps defines properties for table headers
type HeadProps struct {
	Class     string
	Sortable  bool
	Sorted    string // "asc" | "desc" | ""
	Align     string // "left" | "center" | "right"
	ColSpan   int
	RowSpan   int
}

// Head creates a table header cell
func Head(props HeadProps, children ...g.Node) g.Node {
	// Set default alignment
	if props.Align == "" {
		props.Align = "left"
	}

	alignClass := map[string]string{
		"left":   "text-left",
		"center": "text-center",
		"right":  "text-right",
	}[props.Align]

	attrs := []g.Node{
		g.Attr("data-table-head", "true"),
		html.Class(lib.CN(
			"text-foreground h-10 px-2 align-middle font-medium whitespace-nowrap",
			alignClass,
			"[&:has([role=checkbox])]:pr-0 [&>[role=checkbox]]:translate-y-[2px]",
			lib.CNIf(props.Sortable, "cursor-pointer select-none", ""),
			props.Class,
		)),
	}

	if props.ColSpan > 0 {
		attrs = append(attrs, g.Attr("colspan", fmt.Sprintf("%d", props.ColSpan)))
	}

	if props.RowSpan > 0 {
		attrs = append(attrs, g.Attr("rowspan", fmt.Sprintf("%d", props.RowSpan)))
	}

	if props.Sortable {
		attrs = append(attrs, g.Attr("data-sortable", "true"))
		if props.Sorted != "" {
			attrs = append(attrs, g.Attr("data-sorted", props.Sorted))
		}
	}

	// Add sort indicator for sortable columns
	content := children
	if props.Sortable {
		var sortIcon g.Node
		if props.Sorted == "asc" {
			sortIcon = g.Raw(`<svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg" class="ml-2 h-4 w-4 inline"><path d="M7.14645 2.14645C7.34171 1.95118 7.65829 1.95118 7.85355 2.14645L11.8536 6.14645C12.0488 6.34171 12.0488 6.65829 11.8536 6.85355C11.6583 7.04882 11.3417 7.04882 11.1464 6.85355L7.5 3.20711L3.85355 6.85355C3.65829 7.04882 3.34171 7.04882 3.14645 6.85355C2.95118 6.65829 2.95118 6.34171 3.14645 6.14645L7.14645 2.14645Z" fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"></path></svg>`)
		} else if props.Sorted == "desc" {
			sortIcon = g.Raw(`<svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg" class="ml-2 h-4 w-4 inline"><path d="M7.85355 12.8536C7.65829 13.0488 7.34171 13.0488 7.14645 12.8536L3.14645 8.85355C2.95118 8.65829 2.95118 8.34171 3.14645 8.14645C3.34171 7.95118 3.65829 7.95118 3.85355 8.14645L7.5 11.7929L11.1464 8.14645C11.3417 7.95118 11.6583 7.95118 11.8536 8.14645C12.0488 8.34171 12.0488 8.65829 11.8536 8.85355L7.85355 12.8536Z" fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"></path></svg>`)
		} else {
			sortIcon = g.Raw(`<svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg" class="ml-2 h-4 w-4 inline opacity-50"><path d="M4.93179 5.43179C4.75605 5.60753 4.75605 5.89245 4.93179 6.06819C5.10753 6.24392 5.39245 6.24392 5.56819 6.06819L7.49999 4.13638L9.43179 6.06819C9.60753 6.24392 9.89245 6.24392 10.0682 6.06819C10.2439 5.89245 10.2439 5.60753 10.0682 5.43179L7.81819 3.18179C7.73379 3.0974 7.61933 3.04999 7.49999 3.04999C7.38064 3.04999 7.26618 3.0974 7.18179 3.18179L4.93179 5.43179ZM10.0682 9.56819C10.2439 9.39245 10.2439 9.10753 10.0682 8.93179C9.89245 8.75606 9.60753 8.75606 9.43179 8.93179L7.49999 10.8636L5.56819 8.93179C5.39245 8.75606 5.10753 8.75606 4.93179 8.93179C4.75605 9.10753 4.75605 9.39245 4.93179 9.56819L7.18179 11.8182C7.26618 11.9026 7.38064 11.95 7.49999 11.95C7.61933 11.95 7.73379 11.9026 7.81819 11.8182L10.0682 9.56819Z" fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"></path></svg>`)
		}
		
		content = append(children, sortIcon)
	}

	return g.El("th",
		g.Group(append(attrs, content...)),
	)
}

// CellProps defines properties for table cells
type CellProps struct {
	Class   string
	Align   string // "left" | "center" | "right"
	ColSpan int
	RowSpan int
}

// Cell creates a table cell
func Cell(props CellProps, children ...g.Node) g.Node {
	// Set default alignment
	if props.Align == "" {
		props.Align = "left"
	}

	alignClass := map[string]string{
		"left":   "text-left",
		"center": "text-center",
		"right":  "text-right",
	}[props.Align]

	attrs := []g.Node{
		g.Attr("data-table-cell", "true"),
		html.Class(lib.CN(
			"p-2 align-middle whitespace-nowrap",
			alignClass,
			"[&:has([role=checkbox])]:pr-0 [&>[role=checkbox]]:translate-y-[2px]",
			props.Class,
		)),
	}

	if props.ColSpan > 0 {
		attrs = append(attrs, g.Attr("colspan", fmt.Sprintf("%d", props.ColSpan)))
	}

	if props.RowSpan > 0 {
		attrs = append(attrs, g.Attr("rowspan", fmt.Sprintf("%d", props.RowSpan)))
	}

	return g.El("td",
		g.Group(append(attrs, children...)),
	)
}

// Caption creates a table caption
func Caption(props Props, children ...g.Node) g.Node {
	return g.El("caption",
		g.Attr("data-table-caption", "true"),
		html.Class(lib.CN("text-muted-foreground mt-4 text-sm", props.Class)),
		g.Group(children),
	)
}

// Helper functions for common table patterns

// SimpleTable creates a basic table structure
func SimpleTable(headers []string, rows [][]g.Node) g.Node {
	headerCells := make([]g.Node, len(headers))
	for i, h := range headers {
		headerCells[i] = Head(HeadProps{}, g.Text(h))
	}

	bodyRows := make([]g.Node, len(rows))
	for i, row := range rows {
		cells := make([]g.Node, len(row))
		for j, cell := range row {
			cells[j] = Cell(CellProps{}, cell)
		}
		bodyRows[i] = Row(RowProps{}, g.Group(cells))
	}

	return TableComponent(Props{},
		HeaderComponent(Props{}, Row(RowProps{}, g.Group(headerCells))),
		Body(Props{}, g.Group(bodyRows)),
	)
}

// StripedTable creates a table with striped rows
func StripedTable(props Props, children ...g.Node) g.Node {
	props.Class = lib.CN(props.Class, "[&_tbody_tr:nth-child(even)]:bg-muted/50")
	return TableComponent(props, children...)
}

// BorderlessTable creates a table without borders
func BorderlessTable(props Props, children ...g.Node) g.Node {
	props.Class = lib.CN(props.Class, "[&_tr]:border-0")
	return TableComponent(props, children...)
}

// CompactTable creates a table with reduced padding
func CompactTable(props Props, children ...g.Node) g.Node {
	props.Class = lib.CN(props.Class, "[&_th]:px-1 [&_th]:h-8 [&_td]:p-1")
	return TableComponent(props, children...)
}

// ResponsiveTable creates a table optimized for mobile
func ResponsiveTable(props Props, children ...g.Node) g.Node {
	return html.Div(
		html.Class("w-full overflow-hidden rounded-md border"),
		TableComponent(props, children...),
	)
}

// StickyHeaderTable creates a table with a sticky header
func StickyHeaderTable(props Props, maxHeight string, children ...g.Node) g.Node {
	if maxHeight == "" {
		maxHeight = "400px"
	}
	
	return html.Div(
		html.Class("relative overflow-auto rounded-md border"),
		g.Attr("style", fmt.Sprintf("max-height: %s", maxHeight)),
		g.El("table",
			g.Attr("data-table", "true"),
			html.Class(lib.CN("w-full caption-bottom text-sm", props.Class)),
			g.If(props.ID != "", g.Attr("id", props.ID)),
			// Apply sticky positioning to thead
			g.Raw(`<style>
				#` + props.ID + ` thead {
					position: sticky;
					top: 0;
					z-index: 10;
					background: hsl(var(--background));
				}
			</style>`),
			g.Group(children),
		),
	)
}