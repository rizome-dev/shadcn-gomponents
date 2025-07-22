package form

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// Props defines the properties for the Form component
type Props struct {
	Method  string // "get" | "post"
	Action  string
	Class   string
	OnSubmit string // JavaScript onsubmit handler
}

// ItemProps defines properties for form items
type ItemProps struct {
	Class string
}

// LabelProps defines properties for form labels
type LabelProps struct {
	For      string
	Required bool
	Class    string
}

// ControlProps defines properties for form controls
type ControlProps struct {
	Class string
}

// DescriptionProps defines properties for form descriptions
type DescriptionProps struct {
	Class string
}

// MessageProps defines properties for form messages
type MessageProps struct {
	Class string
	Error bool // If true, styles as error message
}

// FieldsetProps defines properties for fieldsets
type FieldsetProps struct {
	Class    string
	Disabled bool
}

// LegendProps defines properties for legends
type LegendProps struct {
	Class string
}

// New creates a new Form component
func New(props Props, children ...g.Node) g.Node {
	classes := lib.CN("space-y-8", props.Class)
	
	attrs := []g.Node{
		html.Class(classes),
	}
	
	if props.Method != "" {
		attrs = append(attrs, html.Method(props.Method))
	}
	
	if props.Action != "" {
		attrs = append(attrs, html.Action(props.Action))
	}
	
	if props.OnSubmit != "" {
		attrs = append(attrs, g.Attr("onsubmit", props.OnSubmit))
	}
	
	return html.Form(append(attrs, children...)...)
}

// FormItem creates a form item container
func FormItem(props ItemProps, children ...g.Node) g.Node {
	classes := lib.CN("space-y-2", props.Class)
	
	return html.Div(
		html.Class(classes),
		g.Group(children),
	)
}

// FormLabel creates a form label
func FormLabel(props LabelProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70",
		props.Class,
	)
	
	attrs := []g.Node{
		html.Class(classes),
	}
	
	if props.For != "" {
		attrs = append(attrs, html.For(props.For))
	}
	
	labelChildren := children
	if props.Required {
		labelChildren = append(labelChildren, 
			html.Span(html.Class("text-destructive ml-1"), g.Text("*")),
		)
	}
	
	return html.Label(append(attrs, labelChildren...)...)
}

// FormControl creates a form control wrapper
func FormControl(props ControlProps, children ...g.Node) g.Node {
	classes := lib.CN(props.Class)
	
	return html.Div(
		g.If(classes != "", html.Class(classes)),
		g.Group(children),
	)
}

// FormDescription creates a form field description
func FormDescription(props DescriptionProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"text-[0.8rem] text-muted-foreground",
		props.Class,
	)
	
	return html.P(
		html.Class(classes),
		g.Group(children),
	)
}

// FormMessage creates a form validation message
func FormMessage(props MessageProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"text-[0.8rem] font-medium",
		lib.CNIf(props.Error, "text-destructive", "text-muted-foreground"),
		props.Class,
	)
	
	return html.P(
		html.Class(classes),
		g.Group(children),
	)
}

// FormFieldset creates a fieldset
func FormFieldset(props FieldsetProps, children ...g.Node) g.Node {
	classes := lib.CN("space-y-4", props.Class)
	
	attrs := []g.Node{
		html.Class(classes),
	}
	
	if props.Disabled {
		attrs = append(attrs, html.Disabled())
	}
	
	return html.FieldSet(append(attrs, children...)...)
}

// FormLegend creates a legend for a fieldset
func FormLegend(props LegendProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"text-sm font-medium leading-none",
		props.Class,
	)
	
	return html.Legend(
		html.Class(classes),
		g.Group(children),
	)
}


// FormSection creates a form section with heading
func FormSection(title string, description string, children ...g.Node) g.Node {
	return html.Div(html.Class("space-y-6"),
		html.Div(
			html.H3(html.Class("text-lg font-medium"), g.Text(title)),
			g.If(description != "", 
				html.P(html.Class("text-sm text-muted-foreground"), g.Text(description)),
			),
		),
		g.Group(children),
	)
}

// InputGroup creates a grouped input with prefix/suffix
func InputGroup(prefix g.Node, input g.Node, suffix g.Node) g.Node {
	return html.Div(html.Class("flex"),
		g.If(prefix != nil,
			html.Span(html.Class("inline-flex items-center rounded-l-md border border-r-0 border-input bg-muted px-3 text-sm text-muted-foreground"),
				prefix,
			),
		),
		input,
		g.If(suffix != nil,
			html.Span(html.Class("inline-flex items-center rounded-r-md border border-l-0 border-input bg-muted px-3 text-sm text-muted-foreground"),
				suffix,
			),
		),
	)
}

// FormRow creates a horizontal form row
func FormRow(children ...g.Node) g.Node {
	return html.Div(
		html.Class("grid gap-4 sm:grid-cols-2 lg:grid-cols-3"),
		g.Group(children),
	)
}

// FormActions creates a form actions container
func FormActions(props ItemProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"flex flex-col gap-2 sm:flex-row sm:justify-end",
		props.Class,
	)
	
	return html.Div(
		html.Class(classes),
		g.Group(children),
	)
}