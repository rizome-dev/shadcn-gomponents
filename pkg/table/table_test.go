package table

import (
	"bytes"
	"strings"
	"testing"

	g "maragu.dev/gomponents"
)

func TestTable(t *testing.T) {
	tests := []struct {
		name     string
		props    Props
		children []g.Node
		contains []string
	}{
		{
			name:  "basic table",
			props: Props{},
			children: []g.Node{
				HeaderComponent(Props{}, g.Text("Header")),
				Body(Props{}, g.Text("Body")),
			},
			contains: []string{
				`data-table-container="true"`,
				`data-table="true"`,
				`<table`,
				`class="w-full caption-bottom text-sm"`,
				"Header",
				"Body",
			},
		},
		{
			name: "table with ID and class",
			props: Props{
				ID:    "my-table",
				Class: "custom-table",
			},
			children: []g.Node{
				Body(Props{}, g.Text("Content")),
			},
			contains: []string{
				`id="my-table-container"`,
				`id="my-table"`,
				`custom-table`,
				"Content",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			component := TableComponent(tt.props, tt.children...)
			err := component.Render(&buf)
			if err != nil {
				t.Fatalf("Render() error = %v", err)
			}

			result := buf.String()
			for _, expected := range tt.contains {
				if !strings.Contains(result, expected) {
					t.Errorf("Expected output to contain %q, but it didn't. Got:\n%s", expected, result)
				}
			}
		})
	}
}

func TestTableComponents(t *testing.T) {
	// Test Header
	t.Run("header", func(t *testing.T) {
		var buf bytes.Buffer
		component := HeaderComponent(Props{Class: "custom-header"}, g.Text("Header Content"))
		err := component.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}

		result := buf.String()
		expected := []string{
			`<thead`,
			`data-table-header="true"`,
			`[&amp;_tr]:border-b`,
			`custom-header`,
			"Header Content",
		}

		for _, exp := range expected {
			if !strings.Contains(result, exp) {
				t.Errorf("Expected output to contain %q\nGot: %s", exp, result)
			}
		}
	})

	// Test Body
	t.Run("body", func(t *testing.T) {
		var buf bytes.Buffer
		component := Body(Props{}, g.Text("Body Content"))
		err := component.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}

		result := buf.String()
		expected := []string{
			`<tbody`,
			`data-table-body="true"`,
			`[&amp;_tr:last-child]:border-0`,
			"Body Content",
		}

		for _, exp := range expected {
			if !strings.Contains(result, exp) {
				t.Errorf("Expected output to contain %q\nGot: %s", exp, result)
			}
		}
	})

	// Test Footer
	t.Run("footer", func(t *testing.T) {
		var buf bytes.Buffer
		component := FooterComponent(Props{}, g.Text("Footer Content"))
		err := component.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}

		result := buf.String()
		expected := []string{
			`<tfoot`,
			`data-table-footer="true"`,
			`bg-muted/50`,
			"Footer Content",
		}

		for _, exp := range expected {
			if !strings.Contains(result, exp) {
				t.Errorf("Expected output to contain %q\nGot: %s", exp, result)
			}
		}
	})

	// Test Caption
	t.Run("caption", func(t *testing.T) {
		var buf bytes.Buffer
		component := Caption(Props{}, g.Text("Table Caption"))
		err := component.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}

		result := buf.String()
		expected := []string{
			`<caption`,
			`data-table-caption="true"`,
			`text-muted-foreground`,
			"Table Caption",
		}

		for _, exp := range expected {
			if !strings.Contains(result, exp) {
				t.Errorf("Expected output to contain %q\nGot: %s", exp, result)
			}
		}
	})
}

func TestRow(t *testing.T) {
	tests := []struct {
		name     string
		props    RowProps
		contains []string
	}{
		{
			name:  "basic row",
			props: RowProps{},
			contains: []string{
				`<tr`,
				`data-table-row="true"`,
				`hover:bg-muted/50`,
				`border-b`,
			},
		},
		{
			name: "selected row",
			props: RowProps{
				Selected: true,
				ID:       "row-1",
			},
			contains: []string{
				`data-state="selected"`,
				`id="row-1"`,
			},
		},
		{
			name: "row with click handler",
			props: RowProps{
				OnClick: `onclick="handleRowClick()"`,
			},
			contains: []string{
				`onclick="handleRowClick()"`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			component := Row(tt.props, g.Text("Row Content"))
			err := component.Render(&buf)
			if err != nil {
				t.Fatalf("Render() error = %v", err)
			}

			result := buf.String()
			for _, expected := range tt.contains {
				if !strings.Contains(result, expected) {
					t.Errorf("Expected output to contain %q", expected)
				}
			}
		})
	}
}

func TestHead(t *testing.T) {
	tests := []struct {
		name     string
		props    HeadProps
		children []g.Node
		contains []string
	}{
		{
			name:     "basic header",
			props:    HeadProps{},
			children: []g.Node{g.Text("Name")},
			contains: []string{
				`<th`,
				`data-table-head="true"`,
				`text-left`,
				"Name",
			},
		},
		{
			name: "sortable header",
			props: HeadProps{
				Sortable: true,
				Sorted:   "asc",
			},
			children: []g.Node{g.Text("Date")},
			contains: []string{
				`data-sortable="true"`,
				`data-sorted="asc"`,
				`cursor-pointer`,
				"Date",
				`viewBox="0 0 15 15"`, // Sort icon
			},
		},
		{
			name: "center aligned with colspan",
			props: HeadProps{
				Align:   "center",
				ColSpan: 3,
			},
			children: []g.Node{g.Text("Actions")},
			contains: []string{
				`text-center`,
				`colspan="3"`,
				"Actions",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			component := Head(tt.props, tt.children...)
			err := component.Render(&buf)
			if err != nil {
				t.Fatalf("Render() error = %v", err)
			}

			result := buf.String()
			for _, expected := range tt.contains {
				if !strings.Contains(result, expected) {
					t.Errorf("Expected output to contain %q", expected)
				}
			}
		})
	}
}

func TestCell(t *testing.T) {
	tests := []struct {
		name     string
		props    CellProps
		children []g.Node
		contains []string
	}{
		{
			name:     "basic cell",
			props:    CellProps{},
			children: []g.Node{g.Text("Cell Content")},
			contains: []string{
				`<td`,
				`data-table-cell="true"`,
				`text-left`,
				"Cell Content",
			},
		},
		{
			name: "right aligned with rowspan",
			props: CellProps{
				Align:   "right",
				RowSpan: 2,
			},
			children: []g.Node{g.Text("123")},
			contains: []string{
				`text-right`,
				`rowspan="2"`,
				"123",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			component := Cell(tt.props, tt.children...)
			err := component.Render(&buf)
			if err != nil {
				t.Fatalf("Render() error = %v", err)
			}

			result := buf.String()
			for _, expected := range tt.contains {
				if !strings.Contains(result, expected) {
					t.Errorf("Expected output to contain %q", expected)
				}
			}
		})
	}
}

func TestHelperTables(t *testing.T) {
	// Test SimpleTable
	t.Run("simple table", func(t *testing.T) {
		headers := []string{"Name", "Email", "Role"}
		rows := [][]g.Node{
			{g.Text("John Doe"), g.Text("john@example.com"), g.Text("Admin")},
			{g.Text("Jane Smith"), g.Text("jane@example.com"), g.Text("User")},
		}

		var buf bytes.Buffer
		component := SimpleTable(headers, rows)
		err := component.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}

		result := buf.String()
		expected := []string{
			">Name</th>",
			">Email</th>",
			">Role</th>",
			">John Doe</td>",
			">john@example.com</td>",
			">Admin</td>",
			">Jane Smith</td>",
		}

		for _, exp := range expected {
			if !strings.Contains(result, exp) {
				t.Errorf("Expected output to contain %q\nGot: %s", exp, result)
			}
		}
	})

	// Test StripedTable
	t.Run("striped table", func(t *testing.T) {
		var buf bytes.Buffer
		component := StripedTable(Props{}, Body(Props{}, g.Text("Content")))
		err := component.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}

		result := buf.String()
		if !strings.Contains(result, "[&amp;_tbody_tr:nth-child(even)]:bg-muted/50") {
			t.Errorf("Expected striped table styling\nGot: %s", result)
		}
	})

	// Test BorderlessTable
	t.Run("borderless table", func(t *testing.T) {
		var buf bytes.Buffer
		component := BorderlessTable(Props{}, Body(Props{}, g.Text("Content")))
		err := component.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}

		result := buf.String()
		if !strings.Contains(result, "[&amp;_tr]:border-0") {
			t.Errorf("Expected borderless table styling\nGot: %s", result)
		}
	})

	// Test CompactTable
	t.Run("compact table", func(t *testing.T) {
		var buf bytes.Buffer
		component := CompactTable(Props{}, Body(Props{}, g.Text("Content")))
		err := component.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}

		result := buf.String()
		if !strings.Contains(result, "[&amp;_th]:px-1") || !strings.Contains(result, "[&amp;_td]:p-1") {
			t.Errorf("Expected compact table styling\nGot: %s", result)
		}
	})

	// Test ResponsiveTable
	t.Run("responsive table", func(t *testing.T) {
		var buf bytes.Buffer
		component := ResponsiveTable(Props{}, Body(Props{}, g.Text("Content")))
		err := component.Render(&buf)
		if err != nil {
			t.Fatalf("Render() error = %v", err)
		}

		result := buf.String()
		if !strings.Contains(result, "overflow-hidden rounded-md border") {
			t.Errorf("Expected responsive table wrapper")
		}
	})
}