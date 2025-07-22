package typography

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// Props defines common properties for typography components
type Props struct {
	Class string      // Additional custom classes
	ID    string      // HTML id attribute
}

// H1 creates a large heading
func H1(props Props, children ...g.Node) g.Node {
	classes := lib.CN(
		"scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl",
		props.Class,
	)

	attrs := []g.Node{html.Class(classes)}
	if props.ID != "" {
		attrs = append(attrs, html.ID(props.ID))
	}

	return html.H1(append(attrs, children...)...)
}

// H2 creates a section heading
func H2(props Props, children ...g.Node) g.Node {
	classes := lib.CN(
		"scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight first:mt-0",
		props.Class,
	)

	attrs := []g.Node{html.Class(classes)}
	if props.ID != "" {
		attrs = append(attrs, html.ID(props.ID))
	}

	return html.H2(append(attrs, children...)...)
}

// H3 creates a sub-section heading
func H3(props Props, children ...g.Node) g.Node {
	classes := lib.CN(
		"scroll-m-20 text-2xl font-semibold tracking-tight",
		props.Class,
	)

	attrs := []g.Node{html.Class(classes)}
	if props.ID != "" {
		attrs = append(attrs, html.ID(props.ID))
	}

	return html.H3(append(attrs, children...)...)
}

// H4 creates a small heading
func H4(props Props, children ...g.Node) g.Node {
	classes := lib.CN(
		"scroll-m-20 text-xl font-semibold tracking-tight",
		props.Class,
	)

	attrs := []g.Node{html.Class(classes)}
	if props.ID != "" {
		attrs = append(attrs, html.ID(props.ID))
	}

	return html.H4(append(attrs, children...)...)
}

// P creates a paragraph
func P(props Props, children ...g.Node) g.Node {
	classes := lib.CN(
		"leading-7 [&:not(:first-child)]:mt-6",
		props.Class,
	)

	attrs := []g.Node{html.Class(classes)}
	if props.ID != "" {
		attrs = append(attrs, html.ID(props.ID))
	}

	return html.P(append(attrs, children...)...)
}

// Lead creates a lead paragraph (larger text)
func Lead(props Props, children ...g.Node) g.Node {
	classes := lib.CN(
		"text-xl text-muted-foreground",
		props.Class,
	)

	attrs := []g.Node{html.Class(classes)}
	if props.ID != "" {
		attrs = append(attrs, html.ID(props.ID))
	}

	return html.P(append(attrs, children...)...)
}

// Large creates large text
func Large(props Props, children ...g.Node) g.Node {
	classes := lib.CN(
		"text-lg font-semibold",
		props.Class,
	)

	attrs := []g.Node{html.Class(classes)}
	if props.ID != "" {
		attrs = append(attrs, html.ID(props.ID))
	}

	return html.Div(append(attrs, children...)...)
}

// Small creates small text
func Small(props Props, children ...g.Node) g.Node {
	classes := lib.CN(
		"text-sm font-medium leading-none",
		props.Class,
	)

	attrs := []g.Node{html.Class(classes)}
	if props.ID != "" {
		attrs = append(attrs, html.ID(props.ID))
	}

	return html.Small(append(attrs, children...)...)
}

// Muted creates muted text
func Muted(props Props, children ...g.Node) g.Node {
	classes := lib.CN(
		"text-sm text-muted-foreground",
		props.Class,
	)

	attrs := []g.Node{html.Class(classes)}
	if props.ID != "" {
		attrs = append(attrs, html.ID(props.ID))
	}

	return html.P(append(attrs, children...)...)
}

// Blockquote creates a blockquote
func Blockquote(props Props, children ...g.Node) g.Node {
	classes := lib.CN(
		"mt-6 border-l-2 pl-6 italic",
		props.Class,
	)

	attrs := []g.Node{html.Class(classes)}
	if props.ID != "" {
		attrs = append(attrs, html.ID(props.ID))
	}

	return g.El("blockquote", append(attrs, children...)...)
}

// List creates an unordered list
func List(props Props, children ...g.Node) g.Node {
	classes := lib.CN(
		"my-6 ml-6 list-disc [&>li]:mt-2",
		props.Class,
	)

	attrs := []g.Node{html.Class(classes)}
	if props.ID != "" {
		attrs = append(attrs, html.ID(props.ID))
	}

	return html.Ul(append(attrs, children...)...)
}

// OrderedList creates an ordered list
func OrderedList(props Props, children ...g.Node) g.Node {
	classes := lib.CN(
		"my-6 ml-6 list-decimal [&>li]:mt-2",
		props.Class,
	)

	attrs := []g.Node{html.Class(classes)}
	if props.ID != "" {
		attrs = append(attrs, html.ID(props.ID))
	}

	return html.Ol(append(attrs, children...)...)
}

// ListItem creates a list item
func ListItem(props Props, children ...g.Node) g.Node {
	attrs := []g.Node{}
	if props.Class != "" {
		attrs = append(attrs, html.Class(props.Class))
	}
	if props.ID != "" {
		attrs = append(attrs, html.ID(props.ID))
	}

	return html.Li(append(attrs, children...)...)
}

// InlineCode creates inline code
func InlineCode(props Props, children ...g.Node) g.Node {
	classes := lib.CN(
		"relative rounded bg-muted px-[0.3rem] py-[0.2rem] font-mono text-sm font-semibold",
		props.Class,
	)

	attrs := []g.Node{html.Class(classes)}
	if props.ID != "" {
		attrs = append(attrs, html.ID(props.ID))
	}

	return html.Code(append(attrs, children...)...)
}

// Pre creates a preformatted text block
func Pre(props Props, children ...g.Node) g.Node {
	classes := lib.CN(
		"overflow-x-auto rounded-lg bg-muted p-4",
		props.Class,
	)

	attrs := []g.Node{html.Class(classes)}
	if props.ID != "" {
		attrs = append(attrs, html.ID(props.ID))
	}

	return html.Pre(append(attrs, children...)...)
}

// Code creates a code block (pre with code inside)
func Code(props Props, children ...g.Node) g.Node {
	return Pre(Props{Class: props.Class, ID: props.ID},
		html.Code(
			html.Class("font-mono text-sm"),
			g.Group(children),
		),
	)
}

// TableProps defines properties for the Table component
type TableProps struct {
	Class   string
	ID      string
	Caption string // Optional table caption
}

// Table creates a styled table
func Table(props TableProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"w-full",
		props.Class,
	)

	attrs := []g.Node{html.Class(classes)}
	if props.ID != "" {
		attrs = append(attrs, html.ID(props.ID))
	}

	tableChildren := children
	if props.Caption != "" {
		tableChildren = append([]g.Node{
			html.Caption(html.Class("mt-4 text-sm text-muted-foreground"), g.Text(props.Caption)),
		}, children...)
	}

	return html.Div(
		html.Class("my-6 w-full overflow-y-auto"),
		html.Table(append(attrs, tableChildren...)...),
	)
}

// TableHeader creates a table header
func TableHeader(props Props, children ...g.Node) g.Node {
	attrs := []g.Node{}
	if props.Class != "" {
		attrs = append(attrs, html.Class(props.Class))
	}
	if props.ID != "" {
		attrs = append(attrs, html.ID(props.ID))
	}

	return g.El("thead", append(attrs, children...)...)
}

// TableBody creates a table body
func TableBody(props Props, children ...g.Node) g.Node {
	attrs := []g.Node{}
	if props.Class != "" {
		attrs = append(attrs, html.Class(props.Class))
	}
	if props.ID != "" {
		attrs = append(attrs, html.ID(props.ID))
	}

	return g.El("tbody", append(attrs, children...)...)
}

// TableRow creates a table row
func TableRow(props Props, children ...g.Node) g.Node {
	classes := lib.CN(
		"m-0 border-t p-0 even:bg-muted",
		props.Class,
	)

	attrs := []g.Node{html.Class(classes)}
	if props.ID != "" {
		attrs = append(attrs, html.ID(props.ID))
	}

	return html.Tr(append(attrs, children...)...)
}

// TableHead creates a table header cell
func TableHead(props Props, children ...g.Node) g.Node {
	classes := lib.CN(
		"border px-4 py-2 text-left font-bold [&[align=center]]:text-center [&[align=right]]:text-right",
		props.Class,
	)

	attrs := []g.Node{html.Class(classes)}
	if props.ID != "" {
		attrs = append(attrs, html.ID(props.ID))
	}

	return html.Th(append(attrs, children...)...)
}

// TableCell creates a table data cell
func TableCell(props Props, children ...g.Node) g.Node {
	classes := lib.CN(
		"border px-4 py-2 text-left [&[align=center]]:text-center [&[align=right]]:text-right",
		props.Class,
	)

	attrs := []g.Node{html.Class(classes)}
	if props.ID != "" {
		attrs = append(attrs, html.ID(props.ID))
	}

	return html.Td(append(attrs, children...)...)
}

// Link creates a styled link
func Link(href string, props Props, children ...g.Node) g.Node {
	classes := lib.CN(
		"font-medium text-primary underline underline-offset-4",
		props.Class,
	)

	attrs := []g.Node{
		html.Href(href),
		html.Class(classes),
	}
	if props.ID != "" {
		attrs = append(attrs, html.ID(props.ID))
	}

	return html.A(append(attrs, children...)...)
}

// ExternalLink creates a link that opens in a new tab
func ExternalLink(href string, props Props, children ...g.Node) g.Node {
	classes := lib.CN(
		"font-medium text-primary underline underline-offset-4",
		props.Class,
	)

	attrs := []g.Node{
		html.Href(href),
		html.Class(classes),
		html.Target("_blank"),
		html.Rel("noopener noreferrer"),
	}
	if props.ID != "" {
		attrs = append(attrs, html.ID(props.ID))
	}

	return html.A(append(attrs, children...)...)
}

// Mark creates highlighted text
func Mark(props Props, children ...g.Node) g.Node {
	classes := lib.CN(
		"bg-yellow-200 dark:bg-yellow-900",
		props.Class,
	)

	attrs := []g.Node{html.Class(classes)}
	if props.ID != "" {
		attrs = append(attrs, html.ID(props.ID))
	}

	return html.Mark(append(attrs, children...)...)
}

// Kbd creates keyboard input styling
func Kbd(props Props, children ...g.Node) g.Node {
	classes := lib.CN(
		"rounded bg-muted px-1 py-0.5 font-mono text-sm font-semibold",
		props.Class,
	)

	attrs := []g.Node{html.Class(classes)}
	if props.ID != "" {
		attrs = append(attrs, html.ID(props.ID))
	}

	return html.Kbd(append(attrs, children...)...)
}

// Hr creates a horizontal rule
func Hr(props Props) g.Node {
	classes := lib.CN(
		"my-8 border-t",
		props.Class,
	)

	attrs := []g.Node{html.Class(classes)}
	if props.ID != "" {
		attrs = append(attrs, html.ID(props.ID))
	}

	return html.Hr(attrs...)
}