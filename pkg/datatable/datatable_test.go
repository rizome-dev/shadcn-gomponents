package datatable

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

func TestNew(t *testing.T) {
	// Test data
	testData := []interface{}{
		map[string]interface{}{"id": 1, "name": "John Doe", "email": "john@example.com", "status": "active"},
		map[string]interface{}{"id": 2, "name": "Jane Smith", "email": "jane@example.com", "status": "inactive"},
		map[string]interface{}{"id": 3, "name": "Bob Johnson", "email": "bob@example.com", "status": "active"},
	}

	testColumns := []Column{
		{ID: "id", Header: "ID", Accessor: "id"},
		{ID: "name", Header: "Name", Accessor: "name"},
		{ID: "email", Header: "Email", Accessor: "email"},
		{ID: "status", Header: "Status", Accessor: "status"},
	}

	tests := []struct {
		name     string
		props    Props
		contains []string
	}{
		{
			name: "default data table",
			props: Props{
				Columns: testColumns,
				Data:    testData,
			},
			contains: []string{
				`<table`,
				`ID`,
				`Name`,
				`Email`,
				`Status`,
				`John Doe`,
				`jane@example.com`,
				`Bob Johnson`,
			},
		},
		{
			name: "empty data table",
			props: Props{
				Columns: testColumns,
				Data:    []interface{}{},
			},
			contains: []string{
				`No data available`,
			},
		},
		{
			name: "loading state",
			props: Props{
				Columns: testColumns,
				Data:    testData,
				Loading: true,
			},
			contains: []string{
				`Loading...`,
				`animate-spin`,
			},
		},
		{
			name: "with caption",
			props: Props{
				Columns: testColumns,
				Data:    testData,
				Caption: "User Data Table",
			},
			contains: []string{
				`User Data Table`,
			},
		},
		{
			name: "selectable table",
			props: Props{
				Columns:      testColumns,
				Data:         testData,
				Selectable:   true,
				SelectedRows: []int{1},
			},
			contains: []string{
				`type="checkbox"`,
				`id="select-all"`,
				`id="select-0"`,
				`bg-muted`, // Selected row styling
			},
		},
		{
			name: "sortable table",
			props: Props{
				Columns: []Column{
					{ID: "id", Header: "ID", Accessor: "id", Sortable: true},
					{ID: "name", Header: "Name", Accessor: "name", Sortable: true},
				},
				Data:          testData,
				Sortable:      true,
				SortColumn:    "name",
				SortDirection: "asc",
			},
			contains: []string{
				`cursor-pointer`,
				`h-4 w-4`, // Sort icon
			},
		},
		{
			name: "with pagination",
			props: Props{
				Columns:     testColumns,
				Data:        testData,
				Pagination:  true,
				PageSize:    2,
				CurrentPage: 0,
			},
			contains: []string{
				`Showing 1 to 2 of 3 rows`,
				`Page 1 of 2`,
				`Previous`,
				`Next`,
			},
		},
		{
			name: "filterable table",
			props: Props{
				Columns:     testColumns,
				Data:        testData,
				Filterable:  true,
				FilterValue: "john",
			},
			contains: []string{
				`Filter...`,
				`value="john"`,
			},
		},
		{
			name: "striped and hoverable",
			props: Props{
				Columns:   testColumns,
				Data:      testData,
				Striped:   true,
				Hoverable: true,
			},
			contains: []string{
				`[&amp;_tbody_tr:nth-child(even)]:bg-muted/50`,
				`[&amp;_tbody_tr]:hover:bg-muted/50`,
			},
		},
		{
			name: "dense layout",
			props: Props{
				Columns: testColumns,
				Data:    testData,
				Dense:   true,
			},
			contains: []string{
				`[&amp;_td]:py-2`,
				`[&amp;_th]:py-2`,
			},
		},
		{
			name: "sticky header",
			props: Props{
				Columns:      testColumns,
				Data:         testData,
				StickyHeader: true,
			},
			contains: []string{
				`[&amp;_thead]:sticky`,
				`[&amp;_thead]:top-0`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := New(tt.props)
			var buf bytes.Buffer
			err := component.Render(&buf)
			if err != nil {
				t.Fatalf("failed to render: %v", err)
			}

			html := buf.String()
			for _, want := range tt.contains {
				if !strings.Contains(html, want) {
					t.Errorf("expected HTML to contain %q, but got:\n%s", want, html)
				}
			}
		})
	}
}

func TestHelperFunctions(t *testing.T) {
	testData := []interface{}{
		map[string]interface{}{"id": 1, "name": "Item 1"},
		map[string]interface{}{"id": 2, "name": "Item 2"},
	}

	testColumns := []Column{
		{ID: "id", Header: "ID", Accessor: "id"},
		{ID: "name", Header: "Name", Accessor: "name"},
	}

	t.Run("Simple", func(t *testing.T) {
		component := Simple(testColumns, testData)
		var buf bytes.Buffer
		err := component.Render(&buf)
		if err != nil {
			t.Fatalf("failed to render: %v", err)
		}

		html := buf.String()
		if !strings.Contains(html, "Item 1") {
			t.Errorf("expected Simple data table to contain data")
		}
		if !strings.Contains(html, "[&amp;_tbody_tr:nth-child(even)]:bg-muted/50") {
			t.Errorf("expected Simple data table to be striped")
		}
	})

	t.Run("WithSelection", func(t *testing.T) {
		component := WithSelection(testColumns, testData, []int{0})
		var buf bytes.Buffer
		err := component.Render(&buf)
		if err != nil {
			t.Fatalf("failed to render: %v", err)
		}

		html := buf.String()
		if !strings.Contains(html, "checkbox") {
			t.Errorf("expected WithSelection data table to have checkboxes")
		}
	})

	t.Run("WithSorting", func(t *testing.T) {
		component := WithSorting(testColumns, testData, "name", "asc")
		var buf bytes.Buffer
		err := component.Render(&buf)
		if err != nil {
			t.Fatalf("failed to render: %v", err)
		}

		html := buf.String()
		if !strings.Contains(html, "cursor-pointer") {
			t.Errorf("expected WithSorting data table to have sortable columns")
		}
	})

	t.Run("WithPagination", func(t *testing.T) {
		component := WithPagination(testColumns, testData, 0, 1)
		var buf bytes.Buffer
		err := component.Render(&buf)
		if err != nil {
			t.Fatalf("failed to render: %v", err)
		}

		html := buf.String()
		if !strings.Contains(html, "Page 1 of 2") {
			t.Errorf("expected WithPagination data table to show pagination")
		}
	})

	t.Run("Full", func(t *testing.T) {
		component := Full(testColumns, testData)
		var buf bytes.Buffer
		err := component.Render(&buf)
		if err != nil {
			t.Fatalf("failed to render: %v", err)
		}

		html := buf.String()
		if !strings.Contains(html, "checkbox") {
			t.Errorf("expected Full data table to have selection")
		}
		if !strings.Contains(html, "Filter...") {
			t.Errorf("expected Full data table to have filter")
		}
		if !strings.Contains(html, "Page") {
			t.Errorf("expected Full data table to have pagination")
		}
		if !strings.Contains(html, "[&amp;_thead]:sticky") {
			t.Errorf("expected Full data table to have sticky header")
		}
	})
}

func TestColumnBuilder(t *testing.T) {
	t.Run("basic column", func(t *testing.T) {
		col := NewColumn("test", "Test Header").Build()
		if col.ID != "test" {
			t.Errorf("expected column ID to be 'test', got %s", col.ID)
		}
		if col.Header != "Test Header" {
			t.Errorf("expected column header to be 'Test Header', got %s", col.Header)
		}
		if col.Accessor != "test" {
			t.Errorf("expected column accessor to be 'test', got %s", col.Accessor)
		}
	})

	t.Run("column with options", func(t *testing.T) {
		col := NewColumn("test", "Test").
			WithAccessor("customField").
			Sortable().
			Filterable().
			WithWidth("100px").
			WithAlign("center").
			WithClass("custom-class").
			Build()

		if col.Accessor != "customField" {
			t.Errorf("expected accessor to be 'customField'")
		}
		if !col.Sortable {
			t.Errorf("expected column to be sortable")
		}
		if !col.Filterable {
			t.Errorf("expected column to be filterable")
		}
		if col.Width != "100px" {
			t.Errorf("expected width to be '100px'")
		}
		if col.Align != "center" {
			t.Errorf("expected align to be 'center'")
		}
		if col.Class != "custom-class" {
			t.Errorf("expected class to be 'custom-class'")
		}
	})

	t.Run("column with custom cell", func(t *testing.T) {
		customCell := func(value interface{}, row interface{}) g.Node {
			return html.Span(g.Text("custom"))
		}

		col := NewColumn("test", "Test").
			WithCell(customCell).
			Build()

		if col.Cell == nil {
			t.Errorf("expected column to have custom cell function")
		}
	})
}

func TestCellRendering(t *testing.T) {
	tests := []struct {
		name     string
		column   Column
		rowData  interface{}
		contains string
	}{
		{
			name: "string value",
			column: Column{
				ID:       "name",
				Accessor: "name",
			},
			rowData:  map[string]interface{}{"name": "Test Name"},
			contains: "Test Name",
		},
		{
			name: "numeric value",
			column: Column{
				ID:       "count",
				Accessor: "count",
			},
			rowData:  map[string]interface{}{"count": 42},
			contains: "42",
		},
		{
			name: "boolean true",
			column: Column{
				ID:       "active",
				Accessor: "active",
			},
			rowData:  map[string]interface{}{"active": true},
			contains: "Yes",
		},
		{
			name: "boolean false",
			column: Column{
				ID:       "active",
				Accessor: "active",
			},
			rowData:  map[string]interface{}{"active": false},
			contains: "No",
		},
		{
			name: "nil value",
			column: Column{
				ID:       "missing",
				Accessor: "missing",
			},
			rowData:  map[string]interface{}{},
			contains: "-",
		},
		{
			name: "custom cell renderer",
			column: Column{
				ID:       "status",
				Accessor: "status",
				Cell: func(value interface{}, row interface{}) g.Node {
					if str, ok := value.(string); ok && str == "active" {
						return html.Span(html.Class("text-green-500"), g.Text("Active"))
					}
					return html.Span(html.Class("text-red-500"), g.Text("Inactive"))
				},
			},
			rowData:  map[string]interface{}{"status": "active"},
			contains: "text-green-500",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			content := renderCellContent(tt.column, tt.rowData)
			var buf bytes.Buffer
			err := content.Render(&buf)
			if err != nil {
				t.Fatalf("failed to render cell: %v", err)
			}

			html := buf.String()
			if !strings.Contains(html, tt.contains) {
				t.Errorf("expected cell to contain %q, but got: %s", tt.contains, html)
			}
		})
	}
}

func TestPagination(t *testing.T) {
	// Test data with 10 items
	var testData []interface{}
	for i := 1; i <= 10; i++ {
		testData = append(testData, map[string]interface{}{
			"id":   i,
			"name": fmt.Sprintf("Item %d", i),
		})
	}

	testColumns := []Column{
		{ID: "id", Header: "ID", Accessor: "id"},
		{ID: "name", Header: "Name", Accessor: "name"},
	}

	tests := []struct {
		name         string
		currentPage  int
		pageSize     int
		expectRows   []string
		expectFooter string
	}{
		{
			name:        "first page",
			currentPage: 0,
			pageSize:    3,
			expectRows:  []string{"Item 1", "Item 2", "Item 3"},
			expectFooter: "Showing 1 to 3 of 10 rows",
		},
		{
			name:        "middle page",
			currentPage: 1,
			pageSize:    3,
			expectRows:  []string{"Item 4", "Item 5", "Item 6"},
			expectFooter: "Showing 4 to 6 of 10 rows",
		},
		{
			name:        "last page",
			currentPage: 3,
			pageSize:    3,
			expectRows:  []string{"Item 10"},
			expectFooter: "Showing 10 to 10 of 10 rows",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := New(Props{
				Columns:     testColumns,
				Data:        testData,
				Pagination:  true,
				CurrentPage: tt.currentPage,
				PageSize:    tt.pageSize,
			})

			var buf bytes.Buffer
			err := component.Render(&buf)
			if err != nil {
				t.Fatalf("failed to render: %v", err)
			}

			html := buf.String()

			// Check expected rows are present
			for _, row := range tt.expectRows {
				if !strings.Contains(html, row) {
					t.Errorf("expected page to contain row %q", row)
				}
			}

			// Check footer text
			if !strings.Contains(html, tt.expectFooter) {
				t.Errorf("expected footer to contain %q", tt.expectFooter)
			}
		})
	}
}