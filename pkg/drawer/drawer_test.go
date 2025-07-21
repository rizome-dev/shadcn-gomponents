package drawer

import (
	"strings"
	"testing"

	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func renderToString(node g.Node) string {
	var buf strings.Builder
	node.Render(&buf)
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
			name:  "closed drawer",
			props: Props{Open: false},
			want:  `<div data-state="closed"></div>`,
		},
		{
			name: "open drawer",
			props: Props{Open: true},
			children: []g.Node{
				Overlay(OverlayProps{}),
				Content(ContentProps{}, "right", g.Text("Content")),
			},
			want: `<div class="fixed inset-0 z-50" data-state="open">`,
		},
		{
			name:  "with custom class",
			props: Props{Open: true, Class: "custom-drawer"},
			want:  `<div class="fixed inset-0 z-50 custom-drawer" data-state="open">`,
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

func TestTrigger(t *testing.T) {
	tests := []struct {
		name  string
		props TriggerProps
		want  string
	}{
		{
			name:  "basic trigger",
			props: TriggerProps{},
			want:  `<button type="button">`,
		},
		{
			name:  "with custom class",
			props: TriggerProps{Class: "custom-trigger"},
			want:  `<button type="button" class="custom-trigger">`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := renderToString(Trigger(tt.props, g.Text("Open")))
			if !strings.Contains(got, tt.want) {
				t.Errorf("Trigger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOverlay(t *testing.T) {
	got := renderToString(Overlay(OverlayProps{}))
	want := `data-state="open"`
	if !strings.Contains(got, want) {
		t.Errorf("Overlay() = %v, want %v", got, want)
	}
	
	// Check for animation classes
	if !strings.Contains(got, "animate-in") || !strings.Contains(got, "fade-in-0") {
		t.Errorf("Overlay() missing animation classes")
	}
}

func TestContent(t *testing.T) {
	tests := []struct {
		name  string
		props ContentProps
		side  string
		want  string
	}{
		{
			name:  "right side",
			props: ContentProps{},
			side:  "right",
			want:  `data-side="right"`,
		},
		{
			name:  "left side",
			props: ContentProps{},
			side:  "left",
			want:  `data-side="left"`,
		},
		{
			name:  "top side",
			props: ContentProps{},
			side:  "top",
			want:  `data-side="top"`,
		},
		{
			name:  "bottom side",
			props: ContentProps{},
			side:  "bottom",
			want:  `data-side="bottom"`,
		},
		{
			name:  "with custom class",
			props: ContentProps{Class: "custom-content"},
			side:  "right",
			want:  `custom-content`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := renderToString(Content(tt.props, tt.side, g.Text("Content")))
			if !strings.Contains(got, tt.want) {
				t.Errorf("Content() = %v, want %v", got, tt.want)
			}
			if !strings.Contains(got, `data-state="open"`) {
				t.Errorf("Content() missing data-state")
			}
		})
	}
}

func TestDrawerHeader(t *testing.T) {
	tests := []struct {
		name  string
		props HeaderProps
		want  string
	}{
		{
			name:  "basic header",
			props: HeaderProps{},
			want:  `<div class="flex flex-col space-y-2">`,
		},
		{
			name:  "with custom class",
			props: HeaderProps{Class: "custom-header"},
			want:  `<div class="flex flex-col space-y-2 custom-header">`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := renderToString(DrawerHeader(tt.props, g.Text("Header")))
			if !strings.Contains(got, tt.want) {
				t.Errorf("DrawerHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDrawerTitle(t *testing.T) {
	tests := []struct {
		name  string
		props TitleProps
		want  string
	}{
		{
			name:  "basic title",
			props: TitleProps{},
			want:  `<h2 class="text-lg font-semibold text-foreground">`,
		},
		{
			name:  "with custom class",
			props: TitleProps{Class: "custom-title"},
			want:  `<h2 class="text-lg font-semibold text-foreground custom-title">`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := renderToString(DrawerTitle(tt.props, g.Text("Title")))
			if !strings.Contains(got, tt.want) {
				t.Errorf("DrawerTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDrawerDescription(t *testing.T) {
	tests := []struct {
		name  string
		props DescriptionProps
		want  string
	}{
		{
			name:  "basic description",
			props: DescriptionProps{},
			want:  `<p class="text-sm text-muted-foreground">`,
		},
		{
			name:  "with custom class",
			props: DescriptionProps{Class: "custom-desc"},
			want:  `<p class="text-sm text-muted-foreground custom-desc">`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := renderToString(DrawerDescription(tt.props, g.Text("Description")))
			if !strings.Contains(got, tt.want) {
				t.Errorf("DrawerDescription() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDrawerFooter(t *testing.T) {
	tests := []struct {
		name  string
		props FooterProps
		want  string
	}{
		{
			name:  "basic footer",
			props: FooterProps{},
			want:  `<div class="flex flex-col-reverse sm:flex-row sm:justify-end sm:space-x-2">`,
		},
		{
			name:  "with custom class",
			props: FooterProps{Class: "custom-footer"},
			want:  `<div class="flex flex-col-reverse sm:flex-row sm:justify-end sm:space-x-2 custom-footer">`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := renderToString(DrawerFooter(tt.props, g.Text("Footer")))
			if !strings.Contains(got, tt.want) {
				t.Errorf("DrawerFooter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClose(t *testing.T) {
	tests := []struct {
		name  string
		props CloseProps
		want  string
	}{
		{
			name:  "basic close",
			props: CloseProps{},
			want:  `<button type="button">`,
		},
		{
			name:  "with custom class",
			props: CloseProps{Class: "custom-close"},
			want:  `<button type="button" class="custom-close">`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := renderToString(Close(tt.props, g.Text("Close")))
			if !strings.Contains(got, tt.want) {
				t.Errorf("Close() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDrawerVariants(t *testing.T) {
	tests := []struct {
		name string
		side string
		want string
	}{
		{
			name: "right variant",
			side: "right",
			want: "inset-y-0 right-0",
		},
		{
			name: "left variant",
			side: "left",
			want: "inset-y-0 left-0",
		},
		{
			name: "top variant",
			side: "top",
			want: "inset-x-0 top-0",
		},
		{
			name: "bottom variant",
			side: "bottom",
			want: "inset-x-0 bottom-0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			classes := drawerVariants.GetClasses(lib.VariantProps{
				Variant: tt.side,
			})
			if !strings.Contains(classes, tt.want) {
				t.Errorf("drawerVariants for %s = %v, want %v", tt.side, classes, tt.want)
			}
		})
	}
}

func TestCompleteDrawer(t *testing.T) {
	drawer := New(
		Props{Open: true},
		Overlay(OverlayProps{}),
		Content(
			ContentProps{},
			"right",
			DrawerHeader(
				HeaderProps{},
				DrawerTitle(TitleProps{}, g.Text("Edit Profile")),
				DrawerDescription(
					DescriptionProps{},
					g.Text("Make changes to your profile here."),
				),
			),
			Div(Class("py-4"),
				g.Text("Form content here"),
			),
			DrawerFooter(
				FooterProps{},
				Close(CloseProps{}, g.Text("Cancel")),
				Button(Type("submit"), g.Text("Save")),
			),
		),
	)

	got := renderToString(drawer)
	
	// Check for key elements
	wants := []string{
		`data-state="open"`,
		`Edit Profile`,
		`Make changes to your profile here.`,
		`Form content here`,
		`Cancel`,
		`Save`,
		`data-side="right"`,
	}

	for _, want := range wants {
		if !strings.Contains(got, want) {
			t.Errorf("Complete drawer missing: %v", want)
		}
	}
}