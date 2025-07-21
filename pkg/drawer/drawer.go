package drawer

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// Props defines the properties for the Drawer component
type Props struct {
	Open  bool
	Side  string // "left" | "right" | "top" | "bottom"
	Class string
}

// TriggerProps defines properties for the drawer trigger
type TriggerProps struct {
	Class string
}

// ContentProps defines properties for the drawer content
type ContentProps struct {
	Class string
}

// HeaderProps defines properties for the drawer header
type HeaderProps struct {
	Class string
}

// TitleProps defines properties for the drawer title
type TitleProps struct {
	Class string
}

// DescriptionProps defines properties for the drawer description
type DescriptionProps struct {
	Class string
}

// FooterProps defines properties for the drawer footer
type FooterProps struct {
	Class string
}

// CloseProps defines properties for the drawer close button
type CloseProps struct {
	Class string
}

// OverlayProps defines properties for the drawer overlay
type OverlayProps struct {
	Class string
}

// drawerVariants defines the variant configuration for drawers
var drawerVariants = lib.VariantConfig{
	Base: "fixed z-50 gap-4 bg-background p-6 shadow-lg transition ease-in-out data-[state=open]:duration-500 data-[state=closed]:duration-300",
	Variants: map[string]map[string]string{
		"side": {
			"top":    "inset-x-0 top-0 border-b data-[state=closed]:-translate-y-full data-[state=open]:translate-y-0",
			"bottom": "inset-x-0 bottom-0 border-t data-[state=closed]:translate-y-full data-[state=open]:translate-y-0",
			"left":   "inset-y-0 left-0 h-full w-3/4 border-r data-[state=closed]:-translate-x-full data-[state=open]:translate-x-0 sm:max-w-sm",
			"right":  "inset-y-0 right-0 h-full w-3/4 border-l data-[state=closed]:translate-x-full data-[state=open]:translate-x-0 sm:max-w-sm",
		},
	},
	Defaults: map[string]string{
		"side": "right",
	},
}

// New creates a new Drawer component
func New(props Props, children ...g.Node) g.Node {
	if props.Open {
		classes := lib.CN("fixed inset-0 z-50", props.Class)
		return html.Div(
			html.Class(classes),
			g.Attr("data-state", "open"),
			g.Group(children),
		)
	}
	// Return empty div when closed
	return html.Div(g.Attr("data-state", "closed"))
}

// Trigger creates a drawer trigger
func Trigger(props TriggerProps, children ...g.Node) g.Node {
	classes := lib.CN(props.Class)
	
	return html.Button(
		html.Type("button"),
		g.If(classes != "", html.Class(classes)),
		g.Group(children),
	)
}

// Overlay creates a drawer overlay
func Overlay(props OverlayProps) g.Node {
	classes := lib.CN(
		"fixed inset-0 z-50 bg-black/80",
		"data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0",
		props.Class,
	)

	return html.Div(
		html.Class(classes),
		g.Attr("data-state", "open"),
	)
}

// Content creates the drawer content container
func ContentComponent(props ContentProps, side string, children ...g.Node) g.Node {
	// Get variant classes
	classes := drawerVariants.GetClasses(lib.VariantProps{
		Variant: side,
		Class:   props.Class,
	})

	return html.Div(
		html.Class(classes),
		g.Attr("data-state", "open"),
		g.Attr("data-side", side),
		g.Group(children),
	)
}

// DrawerHeader creates a drawer header section
func DrawerHeader(props HeaderProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"flex flex-col space-y-2",
		props.Class,
	)

	return html.Div(
		html.Class(classes),
		g.Group(children),
	)
}

// DrawerTitle creates a drawer title
func DrawerTitle(props TitleProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"text-lg font-semibold text-foreground",
		props.Class,
	)

	return html.H2(
		html.Class(classes),
		g.Group(children),
	)
}

// DrawerDescription creates a drawer description
func DrawerDescription(props DescriptionProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"text-sm text-muted-foreground",
		props.Class,
	)

	return html.P(
		html.Class(classes),
		g.Group(children),
	)
}

// DrawerFooter creates a drawer footer section
func DrawerFooter(props FooterProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"flex flex-col-reverse sm:flex-row sm:justify-end sm:space-x-2",
		props.Class,
	)

	return html.Div(
		html.Class(classes),
		g.Group(children),
	)
}

// Close creates a drawer close button
func Close(props CloseProps, children ...g.Node) g.Node {
	classes := lib.CN(props.Class)

	return html.Button(
		html.Type("button"),
		g.If(classes != "", html.Class(classes)),
		g.Group(children),
	)
}

// BasicExample creates a basic drawer example
func BasicExample() g.Node {
	return html.Div(
		Trigger(
			TriggerProps{Class: "bg-primary text-primary-foreground hover:bg-primary/90 px-4 py-2 rounded-md"},
			g.Text("Open Drawer"),
		),
		New(
			Props{Open: false},
			Overlay(OverlayProps{}),
			ContentComponent(
				ContentProps{},
				"right",
				DrawerHeader(
					HeaderProps{},
					DrawerTitle(TitleProps{}, g.Text("Edit Profile")),
					DrawerDescription(
						DescriptionProps{},
						g.Text("Make changes to your profile here. Click save when you're done."),
					),
				),
				html.Div(html.Class("grid gap-4 py-4"),
					html.Div(html.Class("grid grid-cols-4 items-center gap-4"),
						html.Label(html.For("name"), html.Class("text-right"), g.Text("Name")),
						html.Input(
							html.ID("name"),
							html.Value("Pedro Duarte"),
							html.Class("col-span-3"),
						),
					),
					html.Div(html.Class("grid grid-cols-4 items-center gap-4"),
						html.Label(html.For("username"), html.Class("text-right"), g.Text("Username")),
						html.Input(
							html.ID("username"),
							html.Value("@peduarte"),
							html.Class("col-span-3"),
						),
					),
				),
				DrawerFooter(
					FooterProps{},
					Close(
						CloseProps{Class: "border hover:bg-accent px-4 py-2 rounded-md"},
						g.Text("Cancel"),
					),
					html.Button(
						html.Type("submit"),
						html.Class("bg-primary text-primary-foreground hover:bg-primary/90 px-4 py-2 rounded-md"),
						g.Text("Save changes"),
					),
				),
			),
		),
	)
}

// ExampleSides demonstrates drawers from different sides
func ExampleSides() g.Node {
	return html.Div(html.Class("grid grid-cols-2 gap-4"),
		// Left drawer
		html.Div(
			Trigger(
				TriggerProps{Class: "border px-4 py-2 rounded-md w-full"},
				g.Text("Open Left Drawer"),
			),
			New(
				Props{Open: false},
				Overlay(OverlayProps{}),
				ContentComponent(
					ContentProps{},
					"left",
					DrawerHeader(
						HeaderProps{},
						DrawerTitle(TitleProps{}, g.Text("Left Drawer")),
						DrawerDescription(
							DescriptionProps{},
							g.Text("This drawer slides in from the left."),
						),
					),
				),
			),
		),
		
		// Right drawer
		html.Div(
			Trigger(
				TriggerProps{Class: "border px-4 py-2 rounded-md w-full"},
				g.Text("Open Right Drawer"),
			),
			New(
				Props{Open: false},
				Overlay(OverlayProps{}),
				ContentComponent(
					ContentProps{},
					"right",
					DrawerHeader(
						HeaderProps{},
						DrawerTitle(TitleProps{}, g.Text("Right Drawer")),
						DrawerDescription(
							DescriptionProps{},
							g.Text("This drawer slides in from the right."),
						),
					),
				),
			),
		),
		
		// Top drawer
		html.Div(
			Trigger(
				TriggerProps{Class: "border px-4 py-2 rounded-md w-full"},
				g.Text("Open Top Drawer"),
			),
			New(
				Props{Open: false},
				Overlay(OverlayProps{}),
				ContentComponent(
					ContentProps{},
					"top",
					DrawerHeader(
						HeaderProps{},
						DrawerTitle(TitleProps{}, g.Text("Top Drawer")),
						DrawerDescription(
							DescriptionProps{},
							g.Text("This drawer slides in from the top."),
						),
					),
				),
			),
		),
		
		// Bottom drawer
		html.Div(
			Trigger(
				TriggerProps{Class: "border px-4 py-2 rounded-md w-full"},
				g.Text("Open Bottom Drawer"),
			),
			New(
				Props{Open: false},
				Overlay(OverlayProps{}),
				ContentComponent(
					ContentProps{},
					"bottom",
					DrawerHeader(
						HeaderProps{},
						DrawerTitle(TitleProps{}, g.Text("Bottom Drawer")),
						DrawerDescription(
							DescriptionProps{},
							g.Text("This drawer slides in from the bottom."),
						),
					),
				),
			),
		),
	)
}