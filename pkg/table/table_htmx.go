package table

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"

	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	hx "maragu.dev/gomponents-htmx"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// HTMXProps defines properties for HTMX-enhanced tables
type HTMXProps struct {
	ID            string
	SortPath      string
	SelectPath    string
	PaginatePath  string
	FilterPath    string
	LoadPath      string
}

// HTMXTable creates an HTMX-enhanced table
func HTMXTable(props Props, htmxProps HTMXProps, children ...g.Node) g.Node {
	return html.Div(
		g.Attr("id", htmxProps.ID+"-wrapper"),
		g.Attr("data-table-wrapper", "true"),
		hx.Get(htmxProps.LoadPath),
		hx.Trigger("load"),
		hx.Target("this"),
		hx.Swap("innerHTML"),
		TableComponent(props, children...),
	)
}

// HTMXSortableHead creates a sortable table header
func HTMXSortableHead(props HeadProps, htmxProps HTMXProps, column string, children ...g.Node) g.Node {
	props.Sortable = true
	
	attrs := []g.Node{
		g.Attr("data-table-head", "true"),
		g.Attr("data-column", column),
		html.Class(lib.CN(
			"text-foreground h-10 px-2 align-middle font-medium whitespace-nowrap cursor-pointer select-none",
			"[&:has([role=checkbox])]:pr-0 [&>[role=checkbox]]:translate-y-[2px]",
			props.Class,
		)),
		hx.Get(htmxProps.SortPath),
		hx.Target("#" + htmxProps.ID + "-wrapper"),
		hx.Swap("innerHTML"),
		hx.Vals(fmt.Sprintf(`{"column": "%s"}`, column)),
	}

	if props.Sorted != "" {
		attrs = append(attrs, g.Attr("data-sorted", props.Sorted))
	}

	// Add sort indicator
	var sortIcon g.Node
	if props.Sorted == "asc" {
		sortIcon = g.Raw(`<svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg" class="ml-2 h-4 w-4 inline"><path d="M7.14645 2.14645C7.34171 1.95118 7.65829 1.95118 7.85355 2.14645L11.8536 6.14645C12.0488 6.34171 12.0488 6.65829 11.8536 6.85355C11.6583 7.04882 11.3417 7.04882 11.1464 6.85355L7.5 3.20711L3.85355 6.85355C3.65829 7.04882 3.34171 7.04882 3.14645 6.85355C2.95118 6.65829 2.95118 6.34171 3.14645 6.14645L7.14645 2.14645Z" fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"></path></svg>`)
	} else if props.Sorted == "desc" {
		sortIcon = g.Raw(`<svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg" class="ml-2 h-4 w-4 inline"><path d="M7.85355 12.8536C7.65829 13.0488 7.34171 13.0488 7.14645 12.8536L3.14645 8.85355C2.95118 8.65829 2.95118 8.34171 3.14645 8.14645C3.34171 7.95118 3.65829 7.95118 3.85355 8.14645L7.5 11.7929L11.1464 8.14645C11.3417 7.95118 11.6583 7.95118 11.8536 8.14645C12.0488 8.34171 12.0488 8.65829 11.8536 8.85355L7.85355 12.8536Z" fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"></path></svg>`)
	} else {
		sortIcon = g.Raw(`<svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg" class="ml-2 h-4 w-4 inline opacity-50"><path d="M4.93179 5.43179C4.75605 5.60753 4.75605 5.89245 4.93179 6.06819C5.10753 6.24392 5.39245 6.24392 5.56819 6.06819L7.49999 4.13638L9.43179 6.06819C9.60753 6.24392 9.89245 6.24392 10.0682 6.06819C10.2439 5.89245 10.2439 5.60753 10.0682 5.43179L7.81819 3.18179C7.73379 3.0974 7.61933 3.04999 7.49999 3.04999C7.38064 3.04999 7.26618 3.0974 7.18179 3.18179L4.93179 5.43179ZM10.0682 9.56819C10.2439 9.39245 10.2439 9.10753 10.0682 8.93179C9.89245 8.75606 9.60753 8.75606 9.43179 8.93179L7.49999 10.8636L5.56819 8.93179C5.39245 8.75606 5.10753 8.75606 4.93179 8.93179C4.75605 9.10753 4.75605 9.39245 4.93179 9.56819L7.18179 11.8182C7.26618 11.9026 7.38064 11.95 7.49999 11.95C7.61933 11.95 7.73379 11.9026 7.81819 11.8182L10.0682 9.56819Z" fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"></path></svg>`)
	}

	return g.El("th",
		append(attrs, append(children, sortIcon)...)...,
	)
}

// HTMXSelectableRow creates a selectable table row
func HTMXSelectableRow(props RowProps, htmxProps HTMXProps, rowID string, children ...g.Node) g.Node {
	attrs := []g.Node{
		g.Attr("data-table-row", "true"),
		g.Attr("data-row-id", rowID),
		html.Class(lib.CN(
			"hover:bg-muted/50 data-[state=selected]:bg-muted border-b transition-colors cursor-pointer",
			props.Class,
		)),
		hx.Post(htmxProps.SelectPath),
		hx.Target("#" + htmxProps.ID + "-wrapper"),
		hx.Swap("innerHTML"),
		hx.Vals(fmt.Sprintf(`{"rowId": "%s"}`, rowID)),
	}

	if props.Selected {
		attrs = append(attrs, g.Attr("data-state", "selected"))
	}

	return g.El("tr",
		append(attrs, children...)...,
	)
}

// TableState represents the state of a table
type TableState struct {
	SortColumn string
	SortOrder  string // "asc" | "desc"
	SelectedRows map[string]bool
	CurrentPage int
	PageSize    int
	Filter      string
}

// TableData represents table data
type TableData struct {
	Headers []string
	Rows    []map[string]interface{}
}

var tableStates = make(map[string]*TableState)
var tableData = make(map[string]*TableData)

// InitializeTable sets up initial table data
func InitializeTable(id string, headers []string, rows []map[string]interface{}) {
	tableData[id] = &TableData{
		Headers: headers,
		Rows:    rows,
	}
	tableStates[id] = &TableState{
		SelectedRows: make(map[string]bool),
		CurrentPage:  1,
		PageSize:     10,
	}
}

// TableHandlers creates HTTP handlers for table functionality
func TableHandlers(mux *http.ServeMux, htmxProps HTMXProps) {
	// Load handler
	mux.HandleFunc(htmxProps.LoadPath, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		state := tableStates[htmxProps.ID]
		data := tableData[htmxProps.ID]
		if state == nil || data == nil {
			http.Error(w, "Table not initialized", http.StatusNotFound)
			return
		}

		// Apply sorting
		rows := make([]map[string]interface{}, len(data.Rows))
		copy(rows, data.Rows)
		
		if state.SortColumn != "" {
			sort.Slice(rows, func(i, j int) bool {
				valI := fmt.Sprintf("%v", rows[i][state.SortColumn])
				valJ := fmt.Sprintf("%v", rows[j][state.SortColumn])
				
				if state.SortOrder == "desc" {
					return valI > valJ
				}
				return valI < valJ
			})
		}

		// Apply filtering
		if state.Filter != "" {
			filtered := make([]map[string]interface{}, 0)
			for _, row := range rows {
				for _, v := range row {
					if strings.Contains(strings.ToLower(fmt.Sprintf("%v", v)), strings.ToLower(state.Filter)) {
						filtered = append(filtered, row)
						break
					}
				}
			}
			rows = filtered
		}

		// Apply pagination
		start := (state.CurrentPage - 1) * state.PageSize
		end := start + state.PageSize
		if end > len(rows) {
			end = len(rows)
		}
		pageRows := rows[start:end]

		// Render table
		renderTable(w, data.Headers, pageRows, state, htmxProps)
	})

	// Sort handler
	mux.HandleFunc(htmxProps.SortPath, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		column := r.URL.Query().Get("column")
		state := tableStates[htmxProps.ID]
		if state == nil {
			http.Error(w, "Table state not found", http.StatusNotFound)
			return
		}

		// Toggle sort order
		if state.SortColumn == column {
			if state.SortOrder == "asc" {
				state.SortOrder = "desc"
			} else if state.SortOrder == "desc" {
				state.SortColumn = ""
				state.SortOrder = ""
			} else {
				state.SortOrder = "asc"
			}
		} else {
			state.SortColumn = column
			state.SortOrder = "asc"
		}

		// Render updated table
		handleLoadPath(w, r, htmxProps)
	})

	// Select handler
	mux.HandleFunc(htmxProps.SelectPath, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		r.ParseForm()
		rowID := r.FormValue("rowId")
		
		state := tableStates[htmxProps.ID]
		if state == nil {
			http.Error(w, "Table state not found", http.StatusNotFound)
			return
		}

		// Toggle selection
		if state.SelectedRows[rowID] {
			delete(state.SelectedRows, rowID)
		} else {
			state.SelectedRows[rowID] = true
		}

		// Render updated table
		handleLoadPath(w, r, htmxProps)
	})

	// Pagination handler
	if htmxProps.PaginatePath != "" {
		mux.HandleFunc(htmxProps.PaginatePath, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}

			page, _ := strconv.Atoi(r.URL.Query().Get("page"))
			if page < 1 {
				page = 1
			}

			state := tableStates[htmxProps.ID]
			if state == nil {
				http.Error(w, "Table state not found", http.StatusNotFound)
				return
			}

			state.CurrentPage = page

			// Render updated table
			handleLoadPath(w, r, htmxProps)
		})
	}

	// Filter handler
	if htmxProps.FilterPath != "" {
		mux.HandleFunc(htmxProps.FilterPath, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}

			r.ParseForm()
			filter := r.FormValue("filter")

			state := tableStates[htmxProps.ID]
			if state == nil {
				http.Error(w, "Table state not found", http.StatusNotFound)
				return
			}

			state.Filter = filter
			state.CurrentPage = 1 // Reset to first page

			// Render updated table
			handleLoadPath(w, r, htmxProps)
		})
	}
}

// Helper function to handle load path logic
func handleLoadPath(w http.ResponseWriter, r *http.Request, htmxProps HTMXProps) {
	// TODO: Implement load path logic
	TableHandlers(nil, htmxProps)
}

// Helper function to render table
func renderTable(w http.ResponseWriter, headers []string, rows []map[string]interface{}, state *TableState, htmxProps HTMXProps) {
	// Create header cells
	headerCells := make([]g.Node, len(headers))
	for i, h := range headers {
		sorted := ""
		if state.SortColumn == h {
			sorted = state.SortOrder
		}
		headerCells[i] = HTMXSortableHead(HeadProps{Sorted: sorted}, htmxProps, h, g.Text(h))
	}

	// Create body rows
	bodyRows := make([]g.Node, len(rows))
	for i, row := range rows {
		rowID := fmt.Sprintf("row-%d", i)
		cells := make([]g.Node, len(headers))
		for j, h := range headers {
			cells[j] = Cell(CellProps{}, g.Text(fmt.Sprintf("%v", row[h])))
		}
		
		selected := state.SelectedRows[rowID]
		bodyRows[i] = HTMXSelectableRow(RowProps{Selected: selected}, htmxProps, rowID, g.Group(cells))
	}

	// Render table
	table := TableComponent(Props{ID: htmxProps.ID},
		HeaderComponent(Props{}, Row(RowProps{}, g.Group(headerCells))),
		Body(Props{}, g.Group(bodyRows)),
	)

	table.Render(w)
}

// PaginationControls creates pagination controls for a table
func PaginationControls(htmxProps HTMXProps, currentPage, totalPages int) g.Node {
	return html.Div(
		html.Class("flex items-center justify-between px-2 py-4"),
		
		// Previous button
		html.Button(
			html.Type("button"),
			html.Class(lib.CN(
				"px-3 py-1 text-sm border rounded",
				lib.CNIf(currentPage == 1, "opacity-50 cursor-not-allowed", "hover:bg-muted"),
			)),
			g.If(currentPage > 1, hx.Get(htmxProps.PaginatePath)),
			g.If(currentPage > 1, hx.Target("#" + htmxProps.ID + "-wrapper")),
			g.If(currentPage > 1, hx.Vals(fmt.Sprintf(`{"page": %d}`, currentPage-1))),
			g.Text("Previous"),
		),
		
		// Page info
		html.Span(html.Class("text-sm text-muted-foreground"),
			g.Text(fmt.Sprintf("Page %d of %d", currentPage, totalPages)),
		),
		
		// Next button
		html.Button(
			html.Type("button"),
			html.Class(lib.CN(
				"px-3 py-1 text-sm border rounded",
				lib.CNIf(currentPage == totalPages, "opacity-50 cursor-not-allowed", "hover:bg-muted"),
			)),
			g.If(currentPage < totalPages, hx.Get(htmxProps.PaginatePath)),
			g.If(currentPage < totalPages, hx.Target("#" + htmxProps.ID + "-wrapper")),
			g.If(currentPage < totalPages, hx.Vals(fmt.Sprintf(`{"page": %d}`, currentPage+1))),
			g.Text("Next"),
		),
	)
}

// FilterInput creates a filter input for a table
func FilterInput(htmxProps HTMXProps, placeholder string) g.Node {
	if placeholder == "" {
		placeholder = "Filter..."
	}
	
	return html.Div(
		html.Class("p-4"),
		html.Input(
			html.Type("text"),
			html.Class("w-full px-3 py-2 border rounded-md text-sm"),
			g.Attr("placeholder", placeholder),
			hx.Post(htmxProps.FilterPath),
			hx.Target("#" + htmxProps.ID + "-wrapper"),
			hx.Trigger("keyup changed delay:500ms"),
			html.Name("filter"),
		),
	)
}