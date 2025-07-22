package datatable

import (
	"fmt"

	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
	"github.com/rizome-dev/shadcn-gomponents/lib/icons"
	"github.com/rizome-dev/shadcn-gomponents/pkg/button"
	"github.com/rizome-dev/shadcn-gomponents/pkg/checkbox"
	"github.com/rizome-dev/shadcn-gomponents/pkg/input"
	"github.com/rizome-dev/shadcn-gomponents/pkg/table"
)

// Column defines a column in the data table
type Column struct {
	ID         string      // Unique identifier for the column
	Header     string      // Header text
	Accessor   string      // Field accessor for data binding
	Cell       CellFunc    // Custom cell renderer
	Sortable   bool        // Whether column is sortable
	Filterable bool        // Whether column is filterable
	Width      string      // Column width (e.g., "100px", "20%")
	Align      string      // Text alignment: "left", "center", "right"
	Class      string      // Additional CSS classes
}

// CellFunc is a function that renders a cell
type CellFunc func(value interface{}, row interface{}) g.Node

// Props defines properties for the DataTable component
type Props struct {
	ID              string      // Table ID
	Columns         []Column    // Column definitions
	Data            []interface{} // Table data
	Caption         string      // Table caption
	EmptyMessage    string      // Message when no data
	Selectable      bool        // Enable row selection
	SelectedRows    []int       // Currently selected row indices
	Sortable        bool        // Enable sorting
	SortColumn      string      // Currently sorted column ID
	SortDirection   string      // "asc" or "desc"
	Filterable      bool        // Enable filtering
	FilterValue     string      // Current filter value
	Pagination      bool        // Enable pagination
	PageSize        int         // Rows per page
	CurrentPage     int         // Current page (0-indexed)
	TotalRows       int         // Total number of rows (for server-side pagination)
	Loading         bool        // Show loading state
	Striped         bool        // Striped rows
	Hoverable       bool        // Highlight rows on hover
	Dense           bool        // Compact table layout
	ShowHeader      bool        // Show/hide header
	StickyHeader    bool        // Make header sticky
	Class           string      // Additional CSS classes
	OnSort          string      // JavaScript to run on sort
	OnFilter        string      // JavaScript to run on filter
	OnPageChange    string      // JavaScript to run on page change
	OnRowSelect     string      // JavaScript to run on row selection
}

// New creates a new DataTable component
func New(props Props) g.Node {
	// Set defaults
	if props.EmptyMessage == "" {
		props.EmptyMessage = "No data available"
	}
	if props.PageSize == 0 {
		props.PageSize = 10
	}
	if props.ShowHeader == false && props.Caption == "" {
		props.ShowHeader = true
	}

	// Calculate pagination
	totalRows := props.TotalRows
	if totalRows == 0 {
		totalRows = len(props.Data)
	}
	totalPages := (totalRows + props.PageSize - 1) / props.PageSize
	
	// Paginate data if client-side pagination
	displayData := props.Data
	if props.Pagination && props.TotalRows == 0 {
		start := props.CurrentPage * props.PageSize
		end := start + props.PageSize
		if end > len(props.Data) {
			end = len(props.Data)
		}
		if start < len(props.Data) {
			displayData = props.Data[start:end]
		} else {
			displayData = []interface{}{}
		}
	}

	// Build table classes
	tableClasses := []string{}
	if props.Striped {
		tableClasses = append(tableClasses, "[&_tbody_tr:nth-child(even)]:bg-muted/50")
	}
	if props.Hoverable {
		tableClasses = append(tableClasses, "[&_tbody_tr]:hover:bg-muted/50")
	}
	if props.Dense {
		tableClasses = append(tableClasses, "[&_td]:py-2 [&_th]:py-2")
	}
	if props.StickyHeader {
		tableClasses = append(tableClasses, "[&_thead]:sticky [&_thead]:top-0 [&_thead]:bg-background [&_thead]:z-10")
	}

	// Build wrapper classes
	wrapperClasses := lib.CN(
		"space-y-4",
		props.Class,
	)

	return html.Div(
		html.Class(wrapperClasses),
		g.If(props.ID != "", html.ID(props.ID)),

		// Header with filter
		g.If(props.Filterable,
			renderTableHeader(props),
		),

		// Table
		table.TableComponent(
			table.Props{
				Class: lib.CN(tableClasses...),
			},
			// Caption
			g.If(props.Caption != "",
				table.Caption(table.Props{}, g.Text(props.Caption)),
			),
			// Header
			g.If(props.ShowHeader,
				renderTableHead(props, displayData),
			),
			// Body
			renderTableBody(props, displayData),
		),

		// Footer with pagination
		g.If(props.Pagination,
			renderTableFooter(props, totalPages, totalRows, len(displayData)),
		),
	)
}

// renderTableHeader renders the filter header
func renderTableHeader(props Props) g.Node {
	return html.Div(
		html.Class("flex items-center justify-between gap-4"),
		// Filter input
		html.Div(
			html.Class("flex items-center gap-2"),
			icons.Search(html.Class("h-4 w-4 text-muted-foreground")),
			input.New(
				input.Props{
					Type:        "text",
					Placeholder: "Filter...",
					Value:       props.FilterValue,
					Class:       "max-w-sm",
					ID:          "datatable-filter",
				},
			),
		),
		// Additional toolbar space
		html.Div(
			html.Class("flex items-center gap-2"),
			// Space for additional controls
		),
	)
}

// renderTableHead renders the table header
func renderTableHead(props Props, data []interface{}) g.Node {
	return table.HeaderComponent(
		table.Props{},
		table.Row(
			table.RowProps{},
			// Selection checkbox column
			g.If(props.Selectable,
				table.Head(
					table.HeadProps{Class: "w-[40px]"},
					checkbox.New(
						checkbox.Props{
							ID:      "select-all",
							Checked: len(props.SelectedRows) == len(data) && len(data) > 0,
							OnChange: props.OnRowSelect,
						},
					),
				),
			),
			// Data columns
			g.Group(g.Map(props.Columns, func(col Column) g.Node {
				// Build header classes
				headerClasses := lib.CN(
					col.Class,
					func() string {
						if col.Align == "center" {
							return "text-center"
						}
						return ""
					}(),
					func() string {
						if col.Align == "right" {
							return "text-right"
						}
						return ""
					}(),
					func() string {
						if props.Sortable && col.Sortable {
							return "cursor-pointer select-none"
						}
						return ""
					}(),
				)

				return table.Head(
					table.HeadProps{
						Class: headerClasses,
					},
					g.If(col.Width != "", html.Style(fmt.Sprintf("width: %s", col.Width))),
					g.If(props.Sortable && col.Sortable,
						html.Div(
							html.Class(lib.CN(
								"flex items-center gap-1",
								func() string {
									if col.Align == "right" {
										return "justify-end"
									}
									if col.Align == "center" {
										return "justify-center"
									}
									return ""
								}(),
							)),
							g.Text(col.Header),
							renderSortIcon(props.SortColumn == col.ID, props.SortDirection),
						),
					),
					g.If(!props.Sortable || !col.Sortable,
						g.Text(col.Header),
					),
					g.If(props.OnSort != "" && col.Sortable, 
						g.Attr("onclick", props.OnSort),
					),
					g.If(props.OnSort != "" && col.Sortable,
						g.Attr("data-column", col.ID),
					),
				)
			})),
		),
	)
}

// renderTableBody renders the table body
func renderTableBody(props Props, data []interface{}) g.Node {
	// Loading state
	if props.Loading {
		return table.Body(
			table.Props{},
			table.Row(
				table.RowProps{},
				table.Cell(
					table.CellProps{
						ColSpan: len(props.Columns) + func() int {
							if props.Selectable {
								return 1
							}
							return 0
						}(),
						Class: "h-24 text-center",
					},
					html.Div(
						html.Class("flex items-center justify-center gap-2"),
						icons.Loader(html.Class("h-4 w-4 animate-spin")),
						g.Text("Loading..."),
					),
				),
			),
		)
	}

	// Empty state
	if len(data) == 0 {
		return table.Body(
			table.Props{},
			table.Row(
				table.RowProps{},
				table.Cell(
					table.CellProps{
						ColSpan: len(props.Columns) + func() int {
							if props.Selectable {
								return 1
							}
							return 0
						}(),
						Class: "h-24 text-center text-muted-foreground",
					},
					g.Text(props.EmptyMessage),
				),
			),
		)
	}

	// Data rows
	var rows []g.Node
	for index, row := range data {
		isSelected := false
		for _, selectedIndex := range props.SelectedRows {
			if selectedIndex == index {
				isSelected = true
				break
			}
		}

		rowNode := table.Row(
			table.RowProps{
				Class: lib.CN(
					func() string {
						if isSelected {
							return "bg-muted"
						}
						return ""
					}(),
				),
			},
				// Selection checkbox
				g.If(props.Selectable,
					table.Cell(
						table.CellProps{},
						checkbox.New(
							checkbox.Props{
								ID:       fmt.Sprintf("select-%d", index),
								Checked:  isSelected,
								Value:    fmt.Sprintf("%d", index),
								OnChange: props.OnRowSelect,
							},
						),
					),
				),
				// Data cells
				g.Group(g.Map(props.Columns, func(col Column) g.Node {
					cellClasses := lib.CN(
						func() string {
							if col.Align == "center" {
								return "text-center"
							}
							return ""
						}(),
						func() string {
							if col.Align == "right" {
								return "text-right"
							}
							return ""
						}(),
					)

					return table.Cell(
						table.CellProps{
							Class: cellClasses,
						},
						renderCellContent(col, row),
					)
				})),
		)
		rows = append(rows, rowNode)
	}

	return table.Body(
		table.Props{},
		g.Group(rows),
	)
}

// renderCellContent renders the content of a cell
func renderCellContent(col Column, row interface{}) g.Node {
	// If custom cell renderer is provided
	if col.Cell != nil {
		value := getFieldValue(row, col.Accessor)
		return col.Cell(value, row)
	}

	// Default rendering
	value := getFieldValue(row, col.Accessor)
	if value == nil {
		return g.Text("-")
	}

	// Convert value to string
	switch v := value.(type) {
	case string:
		return g.Text(v)
	case int, int32, int64, float32, float64:
		return g.Text(fmt.Sprintf("%v", v))
	case bool:
		if v {
			return g.Text("Yes")
		}
		return g.Text("No")
	default:
		return g.Text(fmt.Sprintf("%v", v))
	}
}

// getFieldValue extracts a field value from a row using the accessor
func getFieldValue(row interface{}, accessor string) interface{} {
	// This is a simplified version. In a real implementation,
	// you would use reflection or a proper data access method
	if m, ok := row.(map[string]interface{}); ok {
		return m[accessor]
	}
	return nil
}

// renderSortIcon renders the sort indicator
func renderSortIcon(isActive bool, direction string) g.Node {
	if !isActive {
		return icons.ChevronsUpDown(html.Class("h-4 w-4 opacity-50"))
	}

	if direction == "asc" {
		return icons.ChevronUp(html.Class("h-4 w-4"))
	}
	return icons.ChevronDown(html.Class("h-4 w-4"))
}

// renderTableFooter renders the pagination footer
func renderTableFooter(props Props, totalPages, totalRows, displayDataLen int) g.Node {
	startRow := props.CurrentPage*props.PageSize + 1
	endRow := startRow + displayDataLen - 1
	if endRow > totalRows {
		endRow = totalRows
	}

	return html.Div(
		html.Class("flex items-center justify-between"),
		// Row count
		html.Div(
			html.Class("text-sm text-muted-foreground"),
			g.Text(fmt.Sprintf("Showing %d to %d of %d rows", startRow, endRow, totalRows)),
		),
		// Pagination controls
		html.Div(
			html.Class("flex items-center gap-2"),
			// Previous button
			func() g.Node {
				buttonChildren := []g.Node{
					icons.ChevronLeft(html.Class("h-4 w-4")),
					g.Text("Previous"),
				}
				if props.OnPageChange != "" {
					buttonChildren = append([]g.Node{
						g.Attr("onclick", props.OnPageChange),
						g.Attr("data-page", fmt.Sprintf("%d", props.CurrentPage-1)),
					}, buttonChildren...)
				}
				return button.New(
					button.Props{
						Variant:  "outline",
						Size:     "sm",
						Disabled: props.CurrentPage == 0,
					},
					buttonChildren...,
				)
			}(),
			// Page info
			html.Span(
				html.Class("text-sm"),
				g.Text(fmt.Sprintf("Page %d of %d", props.CurrentPage+1, totalPages)),
			),
			// Next button
			func() g.Node {
				buttonChildren := []g.Node{
					g.Text("Next"),
					icons.ChevronRight(html.Class("h-4 w-4")),
				}
				if props.OnPageChange != "" {
					buttonChildren = append([]g.Node{
						g.Attr("onclick", props.OnPageChange),
						g.Attr("data-page", fmt.Sprintf("%d", props.CurrentPage+1)),
					}, buttonChildren...)
				}
				return button.New(
					button.Props{
						Variant:  "outline",
						Size:     "sm",
						Disabled: props.CurrentPage >= totalPages-1,
					},
					buttonChildren...,
				)
			}(),
		),
	)
}

// Simple creates a simple data table
func Simple(columns []Column, data []interface{}) g.Node {
	return New(Props{
		Columns:   columns,
		Data:      data,
		Striped:   true,
		Hoverable: true,
	})
}

// WithSelection creates a data table with row selection
func WithSelection(columns []Column, data []interface{}, selectedRows []int) g.Node {
	return New(Props{
		Columns:      columns,
		Data:         data,
		Selectable:   true,
		SelectedRows: selectedRows,
		Striped:      true,
		Hoverable:    true,
	})
}

// WithSorting creates a sortable data table
func WithSorting(columns []Column, data []interface{}, sortColumn, sortDirection string) g.Node {
	// Mark sortable columns
	for i := range columns {
		columns[i].Sortable = true
	}

	return New(Props{
		Columns:       columns,
		Data:          data,
		Sortable:      true,
		SortColumn:    sortColumn,
		SortDirection: sortDirection,
		Striped:       true,
		Hoverable:     true,
	})
}

// WithPagination creates a paginated data table
func WithPagination(columns []Column, data []interface{}, currentPage, pageSize int) g.Node {
	return New(Props{
		Columns:     columns,
		Data:        data,
		Pagination:  true,
		CurrentPage: currentPage,
		PageSize:    pageSize,
		Striped:     true,
		Hoverable:   true,
	})
}

// Full creates a fully-featured data table
func Full(columns []Column, data []interface{}) g.Node {
	// Mark all columns as sortable and filterable
	for i := range columns {
		columns[i].Sortable = true
		columns[i].Filterable = true
	}

	return New(Props{
		Columns:      columns,
		Data:         data,
		Selectable:   true,
		Sortable:     true,
		Filterable:   true,
		Pagination:   true,
		PageSize:     10,
		CurrentPage:  0,
		Striped:      true,
		Hoverable:    true,
		StickyHeader: true,
	})
}

// ColumnBuilder provides a fluent API for building columns
type ColumnBuilder struct {
	column Column
}

// NewColumn creates a new column builder
func NewColumn(id, header string) *ColumnBuilder {
	return &ColumnBuilder{
		column: Column{
			ID:       id,
			Header:   header,
			Accessor: id,
		},
	}
}

// WithAccessor sets the data accessor
func (b *ColumnBuilder) WithAccessor(accessor string) *ColumnBuilder {
	b.column.Accessor = accessor
	return b
}

// WithCell sets a custom cell renderer
func (b *ColumnBuilder) WithCell(cell CellFunc) *ColumnBuilder {
	b.column.Cell = cell
	return b
}

// Sortable makes the column sortable
func (b *ColumnBuilder) Sortable() *ColumnBuilder {
	b.column.Sortable = true
	return b
}

// Filterable makes the column filterable
func (b *ColumnBuilder) Filterable() *ColumnBuilder {
	b.column.Filterable = true
	return b
}

// WithWidth sets the column width
func (b *ColumnBuilder) WithWidth(width string) *ColumnBuilder {
	b.column.Width = width
	return b
}

// WithAlign sets the text alignment
func (b *ColumnBuilder) WithAlign(align string) *ColumnBuilder {
	b.column.Align = align
	return b
}

// WithClass adds CSS classes
func (b *ColumnBuilder) WithClass(class string) *ColumnBuilder {
	b.column.Class = class
	return b
}

// Build returns the column
func (b *ColumnBuilder) Build() Column {
	return b.column
}