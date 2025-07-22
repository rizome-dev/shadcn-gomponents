package main

import (
	"fmt"
	"log"
	"net/http"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/pkg/accordion"
	"github.com/rizome-dev/shadcn-gomponents/pkg/alert"
	"github.com/rizome-dev/shadcn-gomponents/pkg/alertdialog"
	"github.com/rizome-dev/shadcn-gomponents/pkg/aspectratio"
	"github.com/rizome-dev/shadcn-gomponents/pkg/avatar"
	"github.com/rizome-dev/shadcn-gomponents/pkg/badge"
	"github.com/rizome-dev/shadcn-gomponents/pkg/breadcrumb"
	"github.com/rizome-dev/shadcn-gomponents/pkg/button"
	"github.com/rizome-dev/shadcn-gomponents/pkg/calendar"
	"github.com/rizome-dev/shadcn-gomponents/pkg/card"
	"github.com/rizome-dev/shadcn-gomponents/pkg/carousel"
	"github.com/rizome-dev/shadcn-gomponents/pkg/chart"
	"github.com/rizome-dev/shadcn-gomponents/pkg/checkbox"
	"github.com/rizome-dev/shadcn-gomponents/pkg/collapsible"
	"github.com/rizome-dev/shadcn-gomponents/pkg/command"
	"github.com/rizome-dev/shadcn-gomponents/pkg/contextmenu"
	"github.com/rizome-dev/shadcn-gomponents/pkg/dialog"
	"github.com/rizome-dev/shadcn-gomponents/pkg/drawer"
	"github.com/rizome-dev/shadcn-gomponents/pkg/dropdownmenu"
	"github.com/rizome-dev/shadcn-gomponents/pkg/form"
	"github.com/rizome-dev/shadcn-gomponents/pkg/hovercard"
	"github.com/rizome-dev/shadcn-gomponents/pkg/input"
	"github.com/rizome-dev/shadcn-gomponents/pkg/inputotp"
	"github.com/rizome-dev/shadcn-gomponents/pkg/label"
	"github.com/rizome-dev/shadcn-gomponents/pkg/menubar"
	"github.com/rizome-dev/shadcn-gomponents/pkg/navigationmenu"
	"github.com/rizome-dev/shadcn-gomponents/pkg/pagination"
	"github.com/rizome-dev/shadcn-gomponents/pkg/popover"
	"github.com/rizome-dev/shadcn-gomponents/pkg/progress"
	"github.com/rizome-dev/shadcn-gomponents/pkg/radio"
	"github.com/rizome-dev/shadcn-gomponents/pkg/resizable"
	"github.com/rizome-dev/shadcn-gomponents/pkg/scrollarea"
	"github.com/rizome-dev/shadcn-gomponents/pkg/selector"
	"github.com/rizome-dev/shadcn-gomponents/pkg/separator"
	"github.com/rizome-dev/shadcn-gomponents/pkg/sheet"
	"github.com/rizome-dev/shadcn-gomponents/pkg/sidebar"
	"github.com/rizome-dev/shadcn-gomponents/pkg/skeleton"
	"github.com/rizome-dev/shadcn-gomponents/pkg/slider"
	"github.com/rizome-dev/shadcn-gomponents/pkg/sonner"
	switchcomp "github.com/rizome-dev/shadcn-gomponents/pkg/switch"
	"github.com/rizome-dev/shadcn-gomponents/pkg/table"
	"github.com/rizome-dev/shadcn-gomponents/pkg/tabs"
	"github.com/rizome-dev/shadcn-gomponents/pkg/textarea"
	"github.com/rizome-dev/shadcn-gomponents/pkg/toast"
	"github.com/rizome-dev/shadcn-gomponents/pkg/toggle"
	"github.com/rizome-dev/shadcn-gomponents/pkg/togglegroup"
	"github.com/rizome-dev/shadcn-gomponents/pkg/tooltip"
)

func main() {
	// Setup routes
	mux := http.NewServeMux()

	// Main demo page
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		DemoPage().Render(w)
	})

	// Component demo pages
	mux.HandleFunc("/components/", func(w http.ResponseWriter, r *http.Request) {
		ComponentsListPage().Render(w)
	})

	// Individual component example pages
	mux.HandleFunc("/accordion", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Accordion", accordion.Example()).Render(w)
	})
	mux.HandleFunc("/alert", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Alert", alert.Example()).Render(w)
	})
	mux.HandleFunc("/alert-dialog", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Alert Dialog", alertdialog.Example()).Render(w)
	})
	mux.HandleFunc("/aspect-ratio", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Aspect Ratio", aspectratio.Example()).Render(w)
	})
	mux.HandleFunc("/avatar", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Avatar", avatar.Example()).Render(w)
	})
	mux.HandleFunc("/badge", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Badge", badge.Example()).Render(w)
	})
	mux.HandleFunc("/breadcrumb", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Breadcrumb", breadcrumb.Example()).Render(w)
	})
	mux.HandleFunc("/button", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Button", button.Example()).Render(w)
	})
	mux.HandleFunc("/calendar", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Calendar", calendar.Example()).Render(w)
	})
	mux.HandleFunc("/card", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Card", card.Examples()).Render(w)
	})
	mux.HandleFunc("/carousel", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Carousel", carousel.ExampleBasic()).Render(w)
	})
	mux.HandleFunc("/chart", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Chart", chart.ExampleLineChart()).Render(w)
	})
	mux.HandleFunc("/checkbox", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Checkbox", checkbox.Example()).Render(w)
	})
	mux.HandleFunc("/collapsible", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Collapsible", collapsible.Example()).Render(w)
	})
	mux.HandleFunc("/command", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Command", command.Example()).Render(w)
	})
	mux.HandleFunc("/context-menu", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Context Menu", contextmenu.Example()).Render(w)
	})
	mux.HandleFunc("/dialog", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Dialog", dialog.Example()).Render(w)
	})
	mux.HandleFunc("/drawer", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Drawer", drawer.Example()).Render(w)
	})
	mux.HandleFunc("/dropdown-menu", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Dropdown Menu", dropdownmenu.Example()).Render(w)
	})
	mux.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Form", form.Example()).Render(w)
	})
	mux.HandleFunc("/hover-card", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Hover Card", hovercard.Example()).Render(w)
	})
	mux.HandleFunc("/input", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Input", input.Examples()).Render(w)
	})
	mux.HandleFunc("/input-otp", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Input OTP", inputotp.Examples()).Render(w)
	})
	mux.HandleFunc("/label", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Label", label.Examples()).Render(w)
	})
	mux.HandleFunc("/menubar", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Menubar", menubar.Examples()).Render(w)
	})
	mux.HandleFunc("/navigation-menu", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Navigation Menu", navigationmenu.Examples()).Render(w)
	})
	mux.HandleFunc("/pagination", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Pagination", pagination.Examples()).Render(w)
	})
	mux.HandleFunc("/popover", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Popover", popover.Example()).Render(w)
	})
	mux.HandleFunc("/progress", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Progress", progress.Example()).Render(w)
	})
	mux.HandleFunc("/radio-group", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Radio Group", radio.Example()).Render(w)
	})
	mux.HandleFunc("/resizable", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Resizable", resizable.Example()).Render(w)
	})
	mux.HandleFunc("/scroll-area", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Scroll Area", scrollarea.Example()).Render(w)
	})
	mux.HandleFunc("/select", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Select", selector.Example()).Render(w)
	})
	mux.HandleFunc("/separator", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Separator", separator.Example()).Render(w)
	})
	mux.HandleFunc("/sheet", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Sheet", sheet.Example()).Render(w)
	})
	mux.HandleFunc("/sidebar", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Sidebar", sidebar.Examples()).Render(w)
	})
	mux.HandleFunc("/skeleton", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Skeleton", skeleton.Example()).Render(w)
	})
	mux.HandleFunc("/slider", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Slider", slider.Examples()).Render(w)
	})
	mux.HandleFunc("/sonner", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Sonner", sonner.Examples()).Render(w)
	})
	mux.HandleFunc("/switch", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Switch", switchcomp.Example()).Render(w)
	})
	mux.HandleFunc("/table", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Table", table.Examples()).Render(w)
	})
	mux.HandleFunc("/tabs", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Tabs", tabs.Example()).Render(w)
	})
	mux.HandleFunc("/textarea", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Textarea", textarea.Example()).Render(w)
	})
	mux.HandleFunc("/toast", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Toast", toast.Example()).Render(w)
	})
	mux.HandleFunc("/toggle", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Toggle", toggle.Example()).Render(w)
	})
	mux.HandleFunc("/toggle-group", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Toggle Group", togglegroup.Example()).Render(w)
	})
	mux.HandleFunc("/tooltip", func(w http.ResponseWriter, r *http.Request) {
		ComponentPage("Tooltip", tooltip.Example()).Render(w)
	})

	// Static file server
	fs := http.FileServer(http.Dir("../../public"))
	mux.Handle("/public/", http.StripPrefix("/public/", fs))

	// Register HTMX handlers for components
	registerHTMXHandlers(mux)

	// Start server
	fmt.Println("Demo app running at http://localhost:8080")
	fmt.Println("View all components at http://localhost:8080/components/")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

// registerHTMXHandlers registers all HTMX endpoints for interactive components
func registerHTMXHandlers(mux *http.ServeMux) {
	// NOTE: Some components (alertdialog, carousel, chart, inputotp, menubar, navigationmenu)
	// don't have handler registration functions yet

	// Calendar handlers
	calendar.CalendarHandlers(mux)

	// Collapsible handlers
	collapsible.CollapsibleHandlers(mux)

	// Command handlers
	command.CommandHandlers(mux)

	// Context Menu handlers
	contextmenu.ContextMenuHandlers(mux)

	// Dialog handlers
	dialog.DialogHandlers(mux)

	// Drawer handlers
	drawer.DrawerHandlers(mux)

	// Dropdown Menu handlers
	dropdownmenu.DropdownMenuHandlers(mux)

	// Hover Card handlers
	hovercard.HoverCardHandlers(mux)

	// Popover handlers
	popover.PopoverHandlers(mux)

	// Sheet handlers
	sheet.SheetHandlers(mux)

	// Sidebar handlers
	sidebar.SidebarHandlers(mux, sidebar.Props{
		Side: "left",
		Open: true,
	}, sidebar.HTMXProps{
		ID:         "demo-sidebar",
		TogglePath: "/htmx/sidebar/toggle",
		StatePath:  "/htmx/sidebar/state",
	})

	// Slider handlers
	slider.SliderHandlers(mux, slider.Props{
		Min:  0,
		Max:  100,
		Step: 1,
		Value: []int{50},
	}, slider.HTMXProps{
		ID:         "demo-slider",
		UpdatePath: "/htmx/slider/update",
		DragPath:   "/htmx/slider/drag",
		InitPath:   "/htmx/slider/init",
	})

	// Sonner (toast) handlers
	sonner.ToasterHandlers(mux, sonner.ToasterProps{
		Position: sonner.PositionBottomRight,
	}, sonner.HTMXToasterProps{
		ID:         "demo-sonner",
		AddPath:    "/htmx/sonner/add",
		RemovePath: "/htmx/sonner/remove",
		UpdatePath: "/htmx/sonner/update",
	})

	// Table handlers
	table.TableHandlers(mux, table.HTMXProps{
		ID:           "demo-table",
		LoadPath:     "/htmx/table/load",
		SortPath:     "/htmx/table/sort",
		SelectPath:   "/htmx/table/select",
		FilterPath:   "/htmx/table/filter",
		PaginatePath: "/htmx/table/page",
	})

	// Toast handlers
	toast.ToastHandlers(mux, toast.HTMXProps{
		ToasterID:   "demo-toast",
		ShowPath:    "/htmx/toast/show",
		DismissPath: "/htmx/toast/dismiss",
	})

	// Toggle Group handlers
	togglegroup.ToggleGroupHandlers(mux, togglegroup.Props{
		Type: "single",
	}, togglegroup.HTMXProps{
		ID:         "demo-toggle-group",
		TogglePath: "/htmx/toggle-group/toggle",
		LoadPath:   "/htmx/toggle-group/load",
	})

	// Tooltip handlers
	tooltip.TooltipHandlers(mux, tooltip.Props{}, tooltip.HTMXProps{
		ID:       "demo-tooltip",
		ShowPath: "/htmx/tooltip/show",
	})
}

// BasePage creates the base HTML structure
func BasePage(title string, content Node) Node {
	return HTML(
		Lang("en"),
		Head(
			Meta(Charset("UTF-8")),
			Meta(Name("viewport"), Content("width=device-width, initial-scale=1.0")),
			TitleEl(Text(title + " - Shadcn Gomponents Demo")),
			Link(Rel("stylesheet"), Href("/public/styles/app.css")),
			Script(Src("/public/scripts/htmx.js")),
			Script(Src("/public/scripts/app.js")),
		),
		Body(
			Class("min-h-screen bg-background text-foreground"),
			content,
		),
	)
}

// ComponentPage creates a page for a single component demo
func ComponentPage(name string, example Node) Node {
	return BasePage(name,
		Div(Class("container mx-auto py-8"),
			// Header with navigation
			Header(Class("mb-8"),
				Nav(Class("flex items-center justify-between"),
					A(Href("/"), Class("text-2xl font-bold"), Text("Shadcn Gomponents")),
					Div(Class("flex gap-4"),
						A(Href("/"), Class("text-sm hover:underline"), Text("Home")),
						A(Href("/components/"), Class("text-sm hover:underline"), Text("All Components")),
					),
				),
			),
			
			// Component demo
			Main(
				H1(Class("text-4xl font-bold mb-2"), Text(name)),
				P(Class("text-muted-foreground mb-8"), Text("Examples and usage of the "+name+" component")),
				Div(Class("border rounded-lg bg-card"),
					example,
				),
			),
		),
	)
}

// DemoPage creates the main landing page
func DemoPage() Node {
	return BasePage("Home",
		Div(Class("min-h-screen"),
			// Hero section
			Section(Class("py-20 px-4 text-center bg-gradient-to-b from-background to-muted"),
				Div(Class("container mx-auto"),
					H1(Class("text-5xl font-bold mb-4"), Text("Shadcn Gomponents")),
					P(Class("text-xl text-muted-foreground mb-8"), 
						Text("Beautiful UI components for Go, built with Gomponents and HTMX")),
					Div(Class("flex gap-4 justify-center"),
						button.New(button.Props{Variant: "default", Size: "lg"}, 
							A(Href("/components/"), Text("View Components")),
						),
						button.New(button.Props{Variant: "outline", Size: "lg"},
							A(Href("https://github.com/rizome-dev/shadcn-gomponents"), 
								Target("_blank"),
								Text("GitHub"),
							),
						),
					),
				),
			),
			
			// Features section
			Section(Class("py-16 px-4"),
				Div(Class("container mx-auto"),
					H2(Class("text-3xl font-bold text-center mb-12"), Text("Features")),
					Div(Class("grid md:grid-cols-3 gap-8"),
						FeatureCard(
							"🚀 Fast & Lightweight",
							"Server-side rendered components with minimal JavaScript",
						),
						FeatureCard(
							"🎨 Beautifully Designed",
							"Based on shadcn/ui design system with Tailwind CSS",
						),
						FeatureCard(
							"⚡ HTMX Enhanced",
							"Interactive components without complex JavaScript frameworks",
						),
						FeatureCard(
							"📦 47 Components",
							"Complete set of UI components for modern web applications",
						),
						FeatureCard(
							"🔧 Customizable",
							"Easy to customize with Tailwind classes and Go code",
						),
						FeatureCard(
							"📱 Responsive",
							"Mobile-first design that works on all screen sizes",
						),
					),
				),
			),
			
			// Component showcase
			Section(Class("py-16 px-4 bg-muted/50"),
				Div(Class("container mx-auto"),
					H2(Class("text-3xl font-bold text-center mb-12"), Text("Component Showcase")),
					Div(Class("grid md:grid-cols-2 lg:grid-cols-3 gap-6"),
						ComponentCard("Button", "Buttons in various styles and sizes", "/button"),
						ComponentCard("Card", "Flexible container components", "/card"),
						ComponentCard("Dialog", "Modal dialogs and popups", "/dialog"),
						ComponentCard("Form", "Form inputs and validation", "/form"),
						ComponentCard("Table", "Data tables with sorting", "/table"),
						ComponentCard("Toast", "Toast notifications", "/toast"),
					),
					Div(Class("text-center mt-8"),
						button.Default(
							A(Href("/components/"), Text("View All Components →")),
						),
					),
				),
			),
			
			// Footer
			Footer(Class("py-8 px-4 border-t"),
				Div(Class("container mx-auto text-center text-sm text-muted-foreground"),
					P(Text("Built with ❤️ using Go, Gomponents, and HTMX")),
					P(Class("mt-2"),
						Text("© 2024 Shadcn Gomponents. Based on "),
						A(Href("https://ui.shadcn.com"), Target("_blank"), 
							Class("underline"), Text("shadcn/ui")),
					),
				),
			),
		),
	)
}

// ComponentsListPage shows all available components
func ComponentsListPage() Node {
	components := []struct {
		Name        string
		Description string
		Path        string
		Category    string
	}{
		// Layout
		{"Aspect Ratio", "Displays content within a desired ratio", "/aspect-ratio", "Layout"},
		{"Card", "Displays content in a card container", "/card", "Layout"},
		{"Resizable", "Resizable panel groups", "/resizable", "Layout"},
		{"Scroll Area", "Augments scrolling functionality", "/scroll-area", "Layout"},
		{"Separator", "Visually separates content", "/separator", "Layout"},
		
		// Forms
		{"Button", "Interactive button component", "/button", "Forms"},
		{"Checkbox", "Checkbox input component", "/checkbox", "Forms"},
		{"Form", "Form layout and validation", "/form", "Forms"},
		{"Input", "Text input fields", "/input", "Forms"},
		{"Input OTP", "One-time password input", "/input-otp", "Forms"},
		{"Label", "Accessible label for inputs", "/label", "Forms"},
		{"Radio Group", "Radio button groups", "/radio-group", "Forms"},
		{"Select", "Select dropdown component", "/select", "Forms"},
		{"Slider", "Input slider for ranges", "/slider", "Forms"},
		{"Switch", "Toggle switch component", "/switch", "Forms"},
		{"Textarea", "Multiline text input", "/textarea", "Forms"},
		{"Toggle", "Toggle button component", "/toggle", "Forms"},
		{"Toggle Group", "Group of toggle buttons", "/toggle-group", "Forms"},
		
		// Data Display
		{"Avatar", "User avatar display", "/avatar", "Data Display"},
		{"Badge", "Small count and labeling", "/badge", "Data Display"},
		{"Calendar", "Date picker calendar", "/calendar", "Data Display"},
		{"Chart", "Data visualization charts", "/chart", "Data Display"},
		{"Progress", "Progress indicators", "/progress", "Data Display"},
		{"Skeleton", "Loading placeholder", "/skeleton", "Data Display"},
		{"Table", "Data table component", "/table", "Data Display"},
		
		// Feedback
		{"Alert", "Alert messages", "/alert", "Feedback"},
		{"Alert Dialog", "Modal alert dialogs", "/alert-dialog", "Feedback"},
		{"Dialog", "Modal dialog windows", "/dialog", "Feedback"},
		{"Drawer", "Sliding panel overlay", "/drawer", "Feedback"},
		{"Sheet", "Sliding sheet overlay", "/sheet", "Feedback"},
		{"Sonner", "Toast notifications (Sonner)", "/sonner", "Feedback"},
		{"Toast", "Toast notifications", "/toast", "Feedback"},
		{"Tooltip", "Informative tooltips", "/tooltip", "Feedback"},
		
		// Navigation
		{"Breadcrumb", "Navigation breadcrumbs", "/breadcrumb", "Navigation"},
		{"Command", "Command palette", "/command", "Navigation"},
		{"Context Menu", "Right-click context menus", "/context-menu", "Navigation"},
		{"Dropdown Menu", "Dropdown menu component", "/dropdown-menu", "Navigation"},
		{"Hover Card", "Card shown on hover", "/hover-card", "Navigation"},
		{"Menubar", "Application menubar", "/menubar", "Navigation"},
		{"Navigation Menu", "Navigation menu component", "/navigation-menu", "Navigation"},
		{"Pagination", "Pagination controls", "/pagination", "Navigation"},
		{"Popover", "Popover component", "/popover", "Navigation"},
		{"Sidebar", "Sidebar navigation", "/sidebar", "Navigation"},
		{"Tabs", "Tabbed interface", "/tabs", "Navigation"},
		
		// Display
		{"Accordion", "Collapsible content panels", "/accordion", "Display"},
		{"Carousel", "Image/content carousel", "/carousel", "Display"},
		{"Collapsible", "Collapsible content section", "/collapsible", "Display"},
	}

	// Group components by category
	categories := make(map[string][]struct {
		Name        string
		Description string
		Path        string
		Category    string
	})
	
	for _, comp := range components {
		categories[comp.Category] = append(categories[comp.Category], comp)
	}

	return BasePage("All Components",
		Div(Class("container mx-auto py-8"),
			// Header
			Header(Class("mb-8"),
				Nav(Class("flex items-center justify-between"),
					A(Href("/"), Class("text-2xl font-bold"), Text("Shadcn Gomponents")),
					A(Href("/"), Class("text-sm hover:underline"), Text("← Back to Home")),
				),
			),
			
			// Page title
			Div(Class("mb-12"),
				H1(Class("text-4xl font-bold mb-4"), Text("All Components")),
				P(Class("text-muted-foreground"), 
					Text(fmt.Sprintf("Explore all %d components available in Shadcn Gomponents", len(components))),
				),
			),
			
			// Components by category
			Main(
				RenderCategories(categories),
			),
			
			// Add toaster for demo
			toast.Toaster(toast.ToasterProps{
				Position: toast.PositionBottomRight,
			}),
		),
	)
}

// Helper functions

func FeatureCard(title, description string) Node {
	return card.Card(
		card.CardHeader(
			card.CardTitle(Text(title)),
		),
		card.CardContent(
			P(Class("text-muted-foreground"), Text(description)),
		),
	)
}

func ComponentCard(name, description, path string) Node {
	return A(Href(path), Class("block transition-colors hover:bg-muted/50"),
		card.Card(
			card.CardHeader(
				card.CardTitle(Text(name)),
				card.CardDescription(Text(description)),
			),
		),
	)
}

func RenderCategories(categories map[string][]struct {
	Name        string
	Description string
	Path        string
	Category    string
}) Node {
	categoryOrder := []string{"Layout", "Forms", "Data Display", "Feedback", "Navigation", "Display"}
	nodes := []Node{}
	
	for _, cat := range categoryOrder {
		if comps, ok := categories[cat]; ok {
			nodes = append(nodes,
				Div(Class("mb-12"),
					H2(Class("text-2xl font-bold mb-6"), Text(cat)),
					Div(Class("grid md:grid-cols-2 lg:grid-cols-3 gap-4"),
						Group(RenderComponentCards(comps)),
					),
				),
			)
		}
	}
	
	return Group(nodes)
}

func RenderComponentCards(components []struct {
	Name        string
	Description string
	Path        string
	Category    string
}) []Node {
	nodes := []Node{}
	for _, comp := range components {
		nodes = append(nodes, ComponentCard(comp.Name, comp.Description, comp.Path))
	}
	return nodes
}