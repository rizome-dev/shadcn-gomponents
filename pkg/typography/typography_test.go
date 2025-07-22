package typography

import (
	"bytes"
	"strings"
	"testing"

	g "maragu.dev/gomponents"
)

func TestHeadings(t *testing.T) {
	tests := []struct {
		name     string
		fn       func(Props, ...g.Node) g.Node
		props    Props
		text     string
		contains []string
	}{
		{
			name:  "H1 default",
			fn:    H1,
			props: Props{},
			text:  "Main Heading",
			contains: []string{
				"<h1",
				"class=\"scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl\"",
				"Main Heading",
			},
		},
		{
			name:  "H1 with custom class",
			fn:    H1,
			props: Props{Class: "text-center"},
			text:  "Centered Heading",
			contains: []string{
				"text-center",
				"Centered Heading",
			},
		},
		{
			name:  "H2 default",
			fn:    H2,
			props: Props{},
			text:  "Section Heading",
			contains: []string{
				"<h2",
				"scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight first:mt-0",
				"Section Heading",
			},
		},
		{
			name:  "H3 default",
			fn:    H3,
			props: Props{},
			text:  "Subsection",
			contains: []string{
				"<h3",
				"scroll-m-20 text-2xl font-semibold tracking-tight",
				"Subsection",
			},
		},
		{
			name:  "H4 default",
			fn:    H4,
			props: Props{},
			text:  "Small Heading",
			contains: []string{
				"<h4",
				"scroll-m-20 text-xl font-semibold tracking-tight",
				"Small Heading",
			},
		},
		{
			name:  "heading with ID",
			fn:    H1,
			props: Props{ID: "main-title"},
			text:  "Title",
			contains: []string{
				`id="main-title"`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := tt.fn(tt.props, g.Text(tt.text))
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

func TestTextVariants(t *testing.T) {
	tests := []struct {
		name     string
		fn       func(Props, ...g.Node) g.Node
		props    Props
		text     string
		contains []string
	}{
		{
			name:  "Paragraph",
			fn:    P,
			props: Props{},
			text:  "Regular paragraph text",
			contains: []string{
				"<p",
				`leading-7 [&amp;:not(:first-child)]:mt-6`,
				"Regular paragraph text",
			},
		},
		{
			name:  "Lead",
			fn:    Lead,
			props: Props{},
			text:  "Lead text",
			contains: []string{
				"<p",
				"text-xl text-muted-foreground",
				"Lead text",
			},
		},
		{
			name:  "Large",
			fn:    Large,
			props: Props{},
			text:  "Large text",
			contains: []string{
				"<div",
				"text-lg font-semibold",
				"Large text",
			},
		},
		{
			name:  "Small",
			fn:    Small,
			props: Props{},
			text:  "Small text",
			contains: []string{
				"<small",
				"text-sm font-medium leading-none",
				"Small text",
			},
		},
		{
			name:  "Muted",
			fn:    Muted,
			props: Props{},
			text:  "Muted text",
			contains: []string{
				"<p",
				"text-sm text-muted-foreground",
				"Muted text",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := tt.fn(tt.props, g.Text(tt.text))
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

func TestBlockquote(t *testing.T) {
	component := Blockquote(Props{}, g.Text("Important quote"))
	var buf bytes.Buffer
	err := component.Render(&buf)
	if err != nil {
		t.Fatalf("failed to render: %v", err)
	}

	html := buf.String()
	expected := []string{
		"<blockquote",
		"mt-6 border-l-2 pl-6 italic",
		"Important quote",
	}

	for _, want := range expected {
		if !strings.Contains(html, want) {
			t.Errorf("expected HTML to contain %q, but got:\n%s", want, html)
		}
	}
}

func TestLists(t *testing.T) {
	t.Run("unordered list", func(t *testing.T) {
		component := List(Props{},
			ListItem(Props{}, g.Text("Item 1")),
			ListItem(Props{}, g.Text("Item 2")),
		)
		var buf bytes.Buffer
		err := component.Render(&buf)
		if err != nil {
			t.Fatalf("failed to render: %v", err)
		}

		html := buf.String()
		expected := []string{
			"<ul",
			`my-6 ml-6 list-disc [&amp;&gt;li]:mt-2`,
			"<li",
			"Item 1",
			"Item 2",
		}

		for _, want := range expected {
			if !strings.Contains(html, want) {
				t.Errorf("expected HTML to contain %q, but got:\n%s", want, html)
			}
		}
	})

	t.Run("ordered list", func(t *testing.T) {
		component := OrderedList(Props{},
			ListItem(Props{}, g.Text("Step 1")),
			ListItem(Props{}, g.Text("Step 2")),
		)
		var buf bytes.Buffer
		err := component.Render(&buf)
		if err != nil {
			t.Fatalf("failed to render: %v", err)
		}

		html := buf.String()
		expected := []string{
			"<ol",
			`my-6 ml-6 list-decimal [&amp;&gt;li]:mt-2`,
			"Step 1",
			"Step 2",
		}

		for _, want := range expected {
			if !strings.Contains(html, want) {
				t.Errorf("expected HTML to contain %q, but got:\n%s", want, html)
			}
		}
	})
}

func TestCode(t *testing.T) {
	t.Run("inline code", func(t *testing.T) {
		component := InlineCode(Props{}, g.Text("const x = 42"))
		var buf bytes.Buffer
		err := component.Render(&buf)
		if err != nil {
			t.Fatalf("failed to render: %v", err)
		}

		html := buf.String()
		expected := []string{
			"<code",
			"relative rounded bg-muted px-[0.3rem] py-[0.2rem] font-mono text-sm font-semibold",
			"const x = 42",
		}

		for _, want := range expected {
			if !strings.Contains(html, want) {
				t.Errorf("expected HTML to contain %q, but got:\n%s", want, html)
			}
		}
	})

	t.Run("code block", func(t *testing.T) {
		component := Code(Props{}, g.Text("func main() {}"))
		var buf bytes.Buffer
		err := component.Render(&buf)
		if err != nil {
			t.Fatalf("failed to render: %v", err)
		}

		html := buf.String()
		expected := []string{
			"<pre",
			"overflow-x-auto rounded-lg bg-muted p-4",
			"<code",
			"font-mono text-sm",
			"func main() {}",
		}

		for _, want := range expected {
			if !strings.Contains(html, want) {
				t.Errorf("expected HTML to contain %q, but got:\n%s", want, html)
			}
		}
	})
}

func TestTable(t *testing.T) {
	component := Table(TableProps{Caption: "Test table"},
		TableHeader(Props{},
			TableRow(Props{},
				TableHead(Props{}, g.Text("Column 1")),
				TableHead(Props{}, g.Text("Column 2")),
			),
		),
		TableBody(Props{},
			TableRow(Props{},
				TableCell(Props{}, g.Text("Cell 1")),
				TableCell(Props{}, g.Text("Cell 2")),
			),
		),
	)

	var buf bytes.Buffer
	err := component.Render(&buf)
	if err != nil {
		t.Fatalf("failed to render: %v", err)
	}

	html := buf.String()
	expected := []string{
		"<div",
		"my-6 w-full overflow-y-auto",
		"<table",
		"w-full",
		"<caption",
		"Test table",
		"<thead",
		"<tr",
		"m-0 border-t p-0 even:bg-muted",
		"<th",
		`border px-4 py-2 text-left font-bold [&amp;[align=center]]:text-center [&amp;[align=right]]:text-right`,
		"Column 1",
		"<tbody",
		"<td",
		`border px-4 py-2 text-left [&amp;[align=center]]:text-center [&amp;[align=right]]:text-right`,
		"Cell 1",
	}

	for _, want := range expected {
		if !strings.Contains(html, want) {
			t.Errorf("expected HTML to contain %q, but got:\n%s", want, html)
		}
	}
}

func TestLinks(t *testing.T) {
	t.Run("regular link", func(t *testing.T) {
		component := Link("/about", Props{}, g.Text("About Us"))
		var buf bytes.Buffer
		err := component.Render(&buf)
		if err != nil {
			t.Fatalf("failed to render: %v", err)
		}

		html := buf.String()
		expected := []string{
			"<a",
			`href="/about"`,
			"font-medium text-primary underline underline-offset-4",
			"About Us",
		}

		for _, want := range expected {
			if !strings.Contains(html, want) {
				t.Errorf("expected HTML to contain %q, but got:\n%s", want, html)
			}
		}
	})

	t.Run("external link", func(t *testing.T) {
		component := ExternalLink("https://example.com", Props{}, g.Text("Example"))
		var buf bytes.Buffer
		err := component.Render(&buf)
		if err != nil {
			t.Fatalf("failed to render: %v", err)
		}

		html := buf.String()
		expected := []string{
			"<a",
			`href="https://example.com"`,
			`target="_blank"`,
			`rel="noopener noreferrer"`,
			"Example",
		}

		for _, want := range expected {
			if !strings.Contains(html, want) {
				t.Errorf("expected HTML to contain %q, but got:\n%s", want, html)
			}
		}
	})
}

func TestSpecialText(t *testing.T) {
	t.Run("mark", func(t *testing.T) {
		component := Mark(Props{}, g.Text("Highlighted"))
		var buf bytes.Buffer
		err := component.Render(&buf)
		if err != nil {
			t.Fatalf("failed to render: %v", err)
		}

		html := buf.String()
		expected := []string{
			"<mark",
			"bg-yellow-200 dark:bg-yellow-900",
			"Highlighted",
		}

		for _, want := range expected {
			if !strings.Contains(html, want) {
				t.Errorf("expected HTML to contain %q, but got:\n%s", want, html)
			}
		}
	})

	t.Run("kbd", func(t *testing.T) {
		component := Kbd(Props{}, g.Text("Ctrl"))
		var buf bytes.Buffer
		err := component.Render(&buf)
		if err != nil {
			t.Fatalf("failed to render: %v", err)
		}

		html := buf.String()
		expected := []string{
			"<kbd",
			"rounded bg-muted px-1 py-0.5 font-mono text-sm font-semibold",
			"Ctrl",
		}

		for _, want := range expected {
			if !strings.Contains(html, want) {
				t.Errorf("expected HTML to contain %q, but got:\n%s", want, html)
			}
		}
	})
}

func TestHr(t *testing.T) {
	component := Hr(Props{})
	var buf bytes.Buffer
	err := component.Render(&buf)
	if err != nil {
		t.Fatalf("failed to render: %v", err)
	}

	html := buf.String()
	expected := []string{
		"<hr",
		"my-8 border-t",
	}

	for _, want := range expected {
		if !strings.Contains(html, want) {
			t.Errorf("expected HTML to contain %q, but got:\n%s", want, html)
		}
	}
}