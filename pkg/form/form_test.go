package form

import (
	"strings"
	"testing"

	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func renderToString(node g.Node) string {
	var buf strings.Builder
	err := node.Render(&buf)
	if err != nil {
		panic(err) // For tests, panic on render error
	}
	return buf.String()
}

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		props    Props
		children []g.Node
		want     string
	}{
		{
			name:  "basic form",
			props: Props{},
			want:  `<form class="space-y-8">`,
		},
		{
			name:  "with method and action",
			props: Props{Method: "post", Action: "/submit"},
			want:  `<form class="space-y-8" method="post" action="/submit">`,
		},
		{
			name:  "with custom class",
			props: Props{Class: "custom-form"},
			want:  `<form class="space-y-8 custom-form"`,
		},
		{
			name:  "with onsubmit handler",
			props: Props{OnSubmit: "return validateForm();"},
			want:  `onsubmit="return validateForm();"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := renderToString(New(tt.props, tt.children...))
			if !strings.Contains(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormItem(t *testing.T) {
	tests := []struct {
		name  string
		props ItemProps
		want  string
	}{
		{
			name:  "basic item",
			props: ItemProps{},
			want:  `<div class="space-y-2">`,
		},
		{
			name:  "with custom class",
			props: ItemProps{Class: "custom-item"},
			want:  `<div class="space-y-2 custom-item">`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := renderToString(FormItem(tt.props, g.Text("Content")))
			if !strings.Contains(got, tt.want) {
				t.Errorf("FormItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormLabel(t *testing.T) {
	tests := []struct {
		name  string
		props LabelProps
		want  string
	}{
		{
			name:  "basic label",
			props: LabelProps{},
			want:  `<label class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70">`,
		},
		{
			name:  "with for attribute",
			props: LabelProps{For: "email"},
			want:  `for="email"`,
		},
		{
			name:  "required label",
			props: LabelProps{Required: true},
			want:  `<span class="text-destructive ml-1">*</span>`,
		},
		{
			name:  "with custom class",
			props: LabelProps{Class: "custom-label"},
			want:  `custom-label`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := renderToString(FormLabel(tt.props, g.Text("Label")))
			if !strings.Contains(got, tt.want) {
				t.Errorf("FormLabel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormControl(t *testing.T) {
	tests := []struct {
		name  string
		props ControlProps
		want  string
	}{
		{
			name:  "basic control",
			props: ControlProps{},
			want:  `<div>`,
		},
		{
			name:  "with custom class",
			props: ControlProps{Class: "custom-control"},
			want:  `<div class="custom-control">`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := renderToString(FormControl(tt.props, g.Text("Control")))
			if !strings.Contains(got, tt.want) {
				t.Errorf("FormControl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormDescription(t *testing.T) {
	tests := []struct {
		name  string
		props DescriptionProps
		want  string
	}{
		{
			name:  "basic description",
			props: DescriptionProps{},
			want:  `<p class="text-[0.8rem] text-muted-foreground">`,
		},
		{
			name:  "with custom class",
			props: DescriptionProps{Class: "custom-desc"},
			want:  `custom-desc`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := renderToString(FormDescription(tt.props, g.Text("Description")))
			if !strings.Contains(got, tt.want) {
				t.Errorf("FormDescription() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormMessage(t *testing.T) {
	tests := []struct {
		name  string
		props MessageProps
		want  string
	}{
		{
			name:  "basic message",
			props: MessageProps{},
			want:  `text-muted-foreground`,
		},
		{
			name:  "error message",
			props: MessageProps{Error: true},
			want:  `text-destructive`,
		},
		{
			name:  "with custom class",
			props: MessageProps{Class: "custom-message"},
			want:  `custom-message`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := renderToString(FormMessage(tt.props, g.Text("Message")))
			if !strings.Contains(got, tt.want) {
				t.Errorf("FormMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormFieldset(t *testing.T) {
	tests := []struct {
		name  string
		props FieldsetProps
		want  string
	}{
		{
			name:  "basic fieldset",
			props: FieldsetProps{},
			want:  `<fieldset class="space-y-4">`,
		},
		{
			name:  "disabled fieldset",
			props: FieldsetProps{Disabled: true},
			want:  `disabled`,
		},
		{
			name:  "with custom class",
			props: FieldsetProps{Class: "custom-fieldset"},
			want:  `custom-fieldset`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := renderToString(FormFieldset(tt.props, g.Text("Content")))
			if !strings.Contains(got, tt.want) {
				t.Errorf("FormFieldset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormLegend(t *testing.T) {
	tests := []struct {
		name  string
		props LegendProps
		want  string
	}{
		{
			name:  "basic legend",
			props: LegendProps{},
			want:  `<legend class="text-sm font-medium leading-none">`,
		},
		{
			name:  "with custom class",
			props: LegendProps{Class: "custom-legend"},
			want:  `custom-legend`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := renderToString(FormLegend(tt.props, g.Text("Legend")))
			if !strings.Contains(got, tt.want) {
				t.Errorf("FormLegend() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSection(t *testing.T) {
	got := renderToString(FormSection("Title", "Description", g.Text("Content")))
	wants := []string{
		`<h3 class="text-lg font-medium">Title</h3>`,
		`<p class="text-sm text-muted-foreground">Description</p>`,
		`Content`,
	}
	
	for _, want := range wants {
		if !strings.Contains(got, want) {
			t.Errorf("Section() missing: %v", want)
		}
	}
}

func TestInputGroup(t *testing.T) {
	input := Input(Type("text"), Class("flex-1"))
	
	tests := []struct {
		name   string
		prefix g.Node
		suffix g.Node
		want   []string
	}{
		{
			name:   "with prefix",
			prefix: g.Text("$"),
			want:   []string{`rounded-l-md`, "$"},
		},
		{
			name:   "with suffix",
			suffix: g.Text("USD"),
			want:   []string{`rounded-r-md`, "USD"},
		},
		{
			name:   "with both",
			prefix: g.Text("$"),
			suffix: g.Text("USD"),
			want:   []string{`rounded-l-md`, `rounded-r-md`, "$", "USD"},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := renderToString(InputGroup(tt.prefix, input, tt.suffix))
			for _, want := range tt.want {
				if !strings.Contains(got, want) {
					t.Errorf("InputGroup() missing: %v", want)
				}
			}
		})
	}
}

func TestFormRow(t *testing.T) {
	got := renderToString(FormRow(g.Text("Item1"), g.Text("Item2")))
	want := `grid gap-4 sm:grid-cols-2 lg:grid-cols-3`
	if !strings.Contains(got, want) {
		t.Errorf("FormRow() = %v, want %v", got, want)
	}
}

func TestFormActions(t *testing.T) {
	tests := []struct {
		name  string
		props ItemProps
		want  string
	}{
		{
			name:  "basic actions",
			props: ItemProps{},
			want:  `flex flex-col gap-2 sm:flex-row sm:justify-end`,
		},
		{
			name:  "with custom class",
			props: ItemProps{Class: "custom-actions"},
			want:  `custom-actions`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := renderToString(FormActions(tt.props, g.Text("Actions")))
			if !strings.Contains(got, tt.want) {
				t.Errorf("FormActions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompleteForm(t *testing.T) {
	form := New(
		Props{Method: "post", Action: "/submit"},
		FormItem(
			ItemProps{},
			FormLabel(LabelProps{For: "name", Required: true}, g.Text("Name")),
			FormControl(
				ControlProps{},
				Input(
					Type("text"),
					ID("name"),
					Name("name"),
					Class("w-full"),
				),
			),
			FormDescription(
				DescriptionProps{},
				g.Text("Enter your full name."),
			),
		),
		FormFieldset(
			FieldsetProps{},
			FormLegend(LegendProps{}, g.Text("Preferences")),
			FormItem(
				ItemProps{},
				FormControl(
					ControlProps{},
					Input(
						Type("checkbox"),
						ID("newsletter"),
						Name("newsletter"),
					),
				),
				FormLabel(LabelProps{For: "newsletter"}, g.Text("Subscribe to newsletter")),
			),
		),
		FormActions(
			ItemProps{},
			Button(Type("submit"), g.Text("Submit")),
		),
	)

	got := renderToString(form)
	
	// Check for key elements
	wants := []string{
		`method="post"`,
		`action="/submit"`,
		`Name`,
		`<span class="text-destructive ml-1">*</span>`,
		`Enter your full name.`,
		`<fieldset`,
		`<legend`,
		`Preferences`,
		`Subscribe to newsletter`,
		`Submit`,
	}

	for _, want := range wants {
		if !strings.Contains(got, want) {
			t.Errorf("Complete form missing: %v", want)
		}
	}
}