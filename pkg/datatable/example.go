package datatable

import (
	"fmt"

	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/pkg/badge"
)

// Example demonstrates various DataTable configurations
func Example() g.Node {
	// Sample data
	users := []interface{}{
		map[string]interface{}{
			"id":       1,
			"name":     "John Doe",
			"email":    "john.doe@example.com",
			"role":     "Admin",
			"status":   "active",
			"joinDate": "2023-01-15",
		},
		map[string]interface{}{
			"id":       2,
			"name":     "Jane Smith",
			"email":    "jane.smith@example.com",
			"role":     "User",
			"status":   "active",
			"joinDate": "2023-02-20",
		},
		map[string]interface{}{
			"id":       3,
			"name":     "Bob Johnson",
			"email":    "bob.johnson@example.com",
			"role":     "User",
			"status":   "inactive",
			"joinDate": "2023-03-10",
		},
		map[string]interface{}{
			"id":       4,
			"name":     "Alice Brown",
			"email":    "alice.brown@example.com",
			"role":     "Moderator",
			"status":   "active",
			"joinDate": "2023-04-05",
		},
		map[string]interface{}{
			"id":       5,
			"name":     "Charlie Wilson",
			"email":    "charlie.wilson@example.com",
			"role":     "User",
			"status":   "pending",
			"joinDate": "2023-05-12",
		},
	}

	// Larger dataset for pagination demo
	var products []interface{}
	for i := 1; i <= 50; i++ {
		products = append(products, map[string]interface{}{
			"id":       i,
			"name":     fmt.Sprintf("Product %d", i),
			"category": []string{"Electronics", "Clothing", "Home", "Books", "Sports"}[i%5],
			"price":    float64(10 + i*5),
			"stock":    i * 10,
			"status":   []string{"in_stock", "low_stock", "out_of_stock"}[i%3],
		})
	}

	return html.Div(
		html.Class("p-8 space-y-12"),

		// Basic data table
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Basic Data Table")),
			Simple(
				[]Column{
					{ID: "id", Header: "ID", Accessor: "id", Width: "60px"},
					{ID: "name", Header: "Name", Accessor: "name"},
					{ID: "email", Header: "Email", Accessor: "email"},
					{ID: "role", Header: "Role", Accessor: "role"},
				},
				users[:3],
			),
		),

		// Data table with custom cell renderers
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Custom Cell Renderers")),
			New(Props{
				Columns: []Column{
					{ID: "id", Header: "ID", Accessor: "id", Width: "60px", Align: "center"},
					{ID: "name", Header: "Name", Accessor: "name"},
					{ID: "email", Header: "Email", Accessor: "email"},
					{
						ID:       "status",
						Header:   "Status",
						Accessor: "status",
						Cell: func(value interface{}, row interface{}) g.Node {
							status, _ := value.(string)
							variant := "default"
							switch status {
							case "active":
								variant = "success"
							case "inactive":
								variant = "secondary"
							case "pending":
								variant = "warning"
							}
							return badge.New(badge.Props{
								Variant: variant,
							}, g.Text(status))
						},
					},
					{
						ID:       "actions",
						Header:   "Actions",
						Accessor: "",
						Align:    "center",
						Cell: func(value interface{}, row interface{}) g.Node {
							return html.Div(
								html.Class("flex gap-2 justify-center"),
								html.Button(
									html.Type("button"),
									html.Class("text-xs px-2 py-1 rounded hover:bg-muted"),
									g.Text("Edit"),
								),
								html.Button(
									html.Type("button"),
									html.Class("text-xs px-2 py-1 rounded hover:bg-muted text-destructive"),
									g.Text("Delete"),
								),
							)
						},
					},
				},
				Data:      users,
				Striped:   true,
				Hoverable: true,
			}),
		),

		// Selectable data table
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Row Selection")),
			WithSelection(
				[]Column{
					{ID: "id", Header: "ID", Accessor: "id", Width: "60px"},
					{ID: "name", Header: "Name", Accessor: "name"},
					{ID: "email", Header: "Email", Accessor: "email"},
					{ID: "role", Header: "Role", Accessor: "role"},
				},
				users,
				[]int{1, 3}, // Pre-selected rows
			),
		),

		// Sortable data table
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Sortable Columns")),
			html.P(html.Class("text-sm text-muted-foreground mb-4"), 
				g.Text("Click column headers to sort")),
			WithSorting(
				[]Column{
					{ID: "id", Header: "ID", Accessor: "id", Width: "60px", Align: "center"},
					{ID: "name", Header: "Name", Accessor: "name"},
					{ID: "email", Header: "Email", Accessor: "email"},
					{ID: "joinDate", Header: "Join Date", Accessor: "joinDate"},
				},
				users,
				"name", // Sort by name
				"asc",  // Ascending order
			),
		),

		// Paginated data table
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Pagination")),
			WithPagination(
				[]Column{
					{ID: "id", Header: "ID", Accessor: "id", Width: "60px"},
					{ID: "name", Header: "Product", Accessor: "name"},
					{ID: "category", Header: "Category", Accessor: "category"},
					{
						ID:       "price",
						Header:   "Price",
						Accessor: "price",
						Align:    "right",
						Cell: func(value interface{}, row interface{}) g.Node {
							price, _ := value.(float64)
							return g.Text(fmt.Sprintf("$%.2f", price))
						},
					},
					{
						ID:       "stock",
						Header:   "Stock",
						Accessor: "stock",
						Align:    "right",
					},
				},
				products,
				0,  // Current page
				10, // Items per page
			),
		),

		// Fully-featured data table
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Full-Featured Data Table")),
			html.P(html.Class("text-sm text-muted-foreground mb-4"), 
				g.Text("With selection, sorting, filtering, and pagination")),
			Full(
				[]Column{
					NewColumn("id", "ID").
						WithWidth("60px").
						WithAlign("center").
						Build(),
					NewColumn("name", "Product").
						Build(),
					NewColumn("category", "Category").
						Build(),
					NewColumn("price", "Price").
						WithAlign("right").
						WithCell(func(value interface{}, row interface{}) g.Node {
							price, _ := value.(float64)
							return g.Text(fmt.Sprintf("$%.2f", price))
						}).
						Build(),
					NewColumn("stock", "Stock").
						WithAlign("right").
						WithCell(func(value interface{}, row interface{}) g.Node {
							stock, _ := value.(int)
							textColor := "text-green-600"
							if stock < 50 {
								textColor = "text-yellow-600"
							}
							if stock < 10 {
								textColor = "text-red-600"
							}
							return html.Span(
								html.Class(textColor+" font-medium"),
								g.Text(fmt.Sprintf("%d", stock)),
							)
						}).
						Build(),
					NewColumn("status", "Status").
						WithCell(func(value interface{}, row interface{}) g.Node {
							status, _ := value.(string)
							variant := "default"
							label := "Unknown"
							switch status {
							case "in_stock":
								variant = "success"
								label = "In Stock"
							case "low_stock":
								variant = "warning"
								label = "Low Stock"
							case "out_of_stock":
								variant = "destructive"
								label = "Out of Stock"
							}
							return badge.New(badge.Props{
								Variant: variant,
							}, g.Text(label))
						}).
						Build(),
				},
				products[:20], // Use subset of products
			),
		),

		// Different styles
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Table Styles")),
			html.Div(html.Class("space-y-8"),
				// Dense table
				html.Div(
					html.H4(html.Class("text-sm font-medium mb-2"), g.Text("Dense Layout")),
					New(Props{
						Columns: []Column{
							{ID: "id", Header: "ID", Accessor: "id"},
							{ID: "name", Header: "Name", Accessor: "name"},
							{ID: "email", Header: "Email", Accessor: "email"},
						},
						Data:  users[:3],
						Dense: true,
					}),
				),

				// Table without striping
				html.Div(
					html.H4(html.Class("text-sm font-medium mb-2"), g.Text("No Striping")),
					New(Props{
						Columns: []Column{
							{ID: "id", Header: "ID", Accessor: "id"},
							{ID: "name", Header: "Name", Accessor: "name"},
							{ID: "email", Header: "Email", Accessor: "email"},
						},
						Data:      users[:3],
						Hoverable: true,
					}),
				),
			),
		),

		// Loading and empty states
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("States")),
			html.Div(html.Class("space-y-8"),
				// Loading state
				html.Div(
					html.H4(html.Class("text-sm font-medium mb-2"), g.Text("Loading State")),
					New(Props{
						Columns: []Column{
							{ID: "id", Header: "ID", Accessor: "id"},
							{ID: "name", Header: "Name", Accessor: "name"},
							{ID: "email", Header: "Email", Accessor: "email"},
						},
						Data:    []interface{}{},
						Loading: true,
					}),
				),

				// Empty state
				html.Div(
					html.H4(html.Class("text-sm font-medium mb-2"), g.Text("Empty State")),
					New(Props{
						Columns: []Column{
							{ID: "id", Header: "ID", Accessor: "id"},
							{ID: "name", Header: "Name", Accessor: "name"},
							{ID: "email", Header: "Email", Accessor: "email"},
						},
						Data:         []interface{}{},
						EmptyMessage: "No users found. Try adjusting your filters.",
					}),
				),
			),
		),

		// Usage notes
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Usage Notes")),
			html.Div(html.Class("rounded-lg border bg-muted/50 p-4"),
				html.Ul(html.Class("text-sm space-y-2 list-disc list-inside"),
					html.Li(g.Text("DataTable is built on top of the Table component")),
					html.Li(g.Text("Supports custom cell renderers for complex content")),
					html.Li(g.Text("Row selection with checkbox support")),
					html.Li(g.Text("Column sorting (requires server-side implementation)")),
					html.Li(g.Text("Filtering with search input")),
					html.Li(g.Text("Client and server-side pagination")),
					html.Li(g.Text("Loading and empty states")),
					html.Li(g.Text("Responsive with horizontal scroll")),
					html.Li(g.Text("Use ColumnBuilder for fluent column configuration")),
					html.Li(g.Text("For HTMX integration, use event handlers")),
				),
			),
		),

		// Code example
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Code Example")),
			html.Pre(html.Class("text-xs bg-muted p-4 rounded-lg overflow-x-auto"),
				html.Code(g.Raw(`// Define columns
columns := []datatable.Column{
    datatable.NewColumn("id", "ID").
        WithWidth("60px").
        WithAlign("center").
        Sortable().
        Build(),
    datatable.NewColumn("name", "Name").
        Sortable().
        Filterable().
        Build(),
    datatable.NewColumn("status", "Status").
        WithCell(func(value interface{}, row interface{}) g.Node {
            // Custom cell renderer
            return badge.New(badge.Props{
                Variant: "success",
                Text:    value.(string),
            })
        }).
        Build(),
}

// Create data table
datatable.New(datatable.Props{
    Columns:      columns,
    Data:         data,
    Selectable:   true,
    Sortable:     true,
    Filterable:   true,
    Pagination:   true,
    PageSize:     20,
    CurrentPage:  0,
    Striped:      true,
    Hoverable:    true,
    StickyHeader: true,
})`)),
			),
		),
	)
}