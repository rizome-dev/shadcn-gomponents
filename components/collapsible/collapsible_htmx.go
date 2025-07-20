package collapsible

import (
	"fmt"
	"net/http"
	"strings"
	
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	hx "maragu.dev/gomponents-htmx"
	"github.com/rizome-dev/shadcn-gomponents/lib"
	"github.com/rizome-dev/shadcn-gomponents/lib/icons"
)

// HTMXProps defines HTMX-specific properties for the Collapsible
type HTMXProps struct {
	ID         string // Unique ID for the collapsible
	TogglePath string // Server path for toggle actions
}

// NewHTMX creates an HTMX-enhanced Collapsible component
func NewHTMX(props Props, htmxProps HTMXProps, trigger g.Node, content g.Node) g.Node {
	classes := lib.CN(props.Class)
	
	return html.Div(
		html.ID(htmxProps.ID),
		g.If(classes != "", html.Class(classes)),
		g.If(props.Open, g.Attr("data-state", "open")),
		g.If(!props.Open, g.Attr("data-state", "closed")),
		trigger,
		g.If(props.Open, content),
	)
}

// TriggerHTMX creates an HTMX-enhanced trigger
func TriggerHTMX(props TriggerProps, htmxProps HTMXProps, isOpen bool, children ...g.Node) g.Node {
	classes := lib.CN(
		"flex w-full items-center justify-between cursor-pointer",
		props.Class,
	)
	
	return html.Div(
		html.Class(classes),
		g.Attr("role", "button"),
		g.Attr("tabindex", "0"),
		hx.Post(htmxProps.TogglePath),
		html.Target("#" + htmxProps.ID),
		hx.Swap("outerHTML"),
		hx.Vals(fmt.Sprintf(`{"open": "%t"}`, !isOpen)),
		g.Group(children),
	)
}

// ContentHTMX creates HTMX-enhanced collapsible content
func ContentHTMX(props ContentProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"overflow-hidden transition-all",
		"data-[state=closed]:animate-accordion-up data-[state=open]:animate-accordion-down",
		props.Class,
	)
	
	return html.Div(
		html.Class(classes),
		g.Group(children),
	)
}

// TriggerButtonHTMX creates an HTMX-enhanced trigger button with icon
func TriggerButtonHTMX(props TriggerProps, htmxProps HTMXProps, isOpen bool) g.Node {
	return html.Button(
		html.Type("button"),
		html.Class(lib.CN(
			"flex items-center justify-between gap-4 px-4 py-2 font-medium transition-all hover:bg-accent",
			"[&[data-state=open]>svg]:rotate-180",
			props.Class,
		)),
		hx.Post(htmxProps.TogglePath),
		html.Target("#" + htmxProps.ID),
		hx.Swap("outerHTML"),
		hx.Vals(fmt.Sprintf(`{"open": "%t"}`, !isOpen)),
		html.Span(html.Class("sr-only"), g.Text("Toggle")),
		icons.ChevronDown(
			html.Class(lib.CN(
				"h-4 w-4 shrink-0 transition-transform duration-200",
				map[bool]string{true: "rotate-180", false: ""}[isOpen],
			)),
		),
	)
}

// ExampleHTMX creates an HTMX-enhanced collapsible example
func ExampleHTMX() g.Node {
	htmxProps := HTMXProps{
		ID:         "collapsible-example",
		TogglePath: "/api/collapsible/toggle",
	}
	
	return RenderCollapsible(htmxProps, true, 
		html.Div(html.Class("flex items-center justify-between space-x-4 px-4"),
			html.H4(html.Class("text-sm font-semibold"),
				g.Text("@peduarte starred 3 repositories"),
			),
			TriggerButtonHTMX(TriggerProps{}, htmxProps, true),
		),
		html.Div(html.Class("rounded-md border px-4 py-3 font-mono text-sm"),
			g.Text("@radix-ui/primitives"),
		),
		html.Div(html.Class("rounded-md border px-4 py-3 font-mono text-sm"),
			g.Text("@radix-ui/colors"),
		),
		html.Div(html.Class("rounded-md border px-4 py-3 font-mono text-sm"),
			g.Text("@stitches/react"),
		),
	)
}

// RenderCollapsible renders a collapsible component (for server response)
func RenderCollapsible(htmxProps HTMXProps, isOpen bool, trigger g.Node, items ...g.Node) g.Node {
	return NewHTMX(
		Props{Open: isOpen},
		htmxProps,
		trigger,
		ContentHTMX(
			ContentProps{},
			html.Div(html.Class("space-y-2"),
				g.Group(items),
			),
		),
	)
}

// AccordionHTMX creates an HTMX-enhanced accordion (multiple collapsibles)
func AccordionHTMX(id string, items []AccordionItem) g.Node {
	return html.Div(
		html.ID(id),
		html.Class("space-y-2"),
		g.Group(g.Map(items, func(item AccordionItem) g.Node {
			htmxProps := HTMXProps{
				ID:         id + "-" + item.ID,
				TogglePath: "/api/accordion/" + id + "/" + item.ID + "/toggle",
			}
			
			return RenderAccordionItem(htmxProps, item)
		})),
	)
}

// AccordionItem represents an item in an accordion
type AccordionItem struct {
	ID      string
	Title   string
	Content g.Node
	Open    bool
}

// RenderAccordionItem renders a single accordion item
func RenderAccordionItem(htmxProps HTMXProps, item AccordionItem) g.Node {
	return html.Div(
		html.ID(htmxProps.ID),
		html.Class("border rounded-lg"),
		g.If(item.Open, g.Attr("data-state", "open")),
		g.If(!item.Open, g.Attr("data-state", "closed")),
		html.Button(
			html.Type("button"),
			html.Class("flex w-full items-center justify-between px-4 py-4 font-medium transition-all hover:bg-accent [&[data-state=open]>svg]:rotate-180"),
			hx.Post(htmxProps.TogglePath),
			html.Target("#" + htmxProps.ID),
			hx.Swap("outerHTML"),
			hx.Vals(fmt.Sprintf(`{"open": "%t"}`, !item.Open)),
			g.Text(item.Title),
			icons.ChevronDown(
				html.Class(lib.CN(
					"h-4 w-4 shrink-0 transition-transform duration-200",
					map[bool]string{true: "rotate-180", false: ""}[item.Open],
				)),
			),
		),
		g.If(item.Open,
			html.Div(
				html.Class("px-4 pb-4 pt-0"),
				item.Content,
			),
		),
	)
}

// FAQExampleHTMX creates an FAQ section using HTMX collapsibles
func FAQExampleHTMX() g.Node {
	items := []AccordionItem{
		{
			ID:    "item-1",
			Title: "Is it accessible?",
			Content: html.P(html.Class("text-muted-foreground"),
				g.Text("Yes. It adheres to the WAI-ARIA design pattern."),
			),
			Open: false,
		},
		{
			ID:    "item-2",
			Title: "Is it styled?",
			Content: html.P(html.Class("text-muted-foreground"),
				g.Text("Yes. It comes with default styles that matches the other components' aesthetic."),
			),
			Open: false,
		},
		{
			ID:    "item-3",
			Title: "Is it animated?",
			Content: html.P(html.Class("text-muted-foreground"),
				g.Text("Yes. It's animated by default, but you can disable it if you prefer."),
			),
			Open: false,
		},
	}
	
	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Frequently Asked Questions")),
		AccordionHTMX("faq-accordion", items),
	)
}

// CollapsibleHandlers creates HTTP handlers for collapsible components
func CollapsibleHandlers(mux *http.ServeMux) {
	// Basic collapsible toggle handler
	mux.HandleFunc("/api/collapsible/toggle", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		isOpen := r.FormValue("open") == "true"
		
		htmxProps := HTMXProps{
			ID:         "collapsible-example",
			TogglePath: "/api/collapsible/toggle",
		}
		
		// Render the collapsible with new state
		node := RenderCollapsible(htmxProps, isOpen,
			html.Div(html.Class("flex items-center justify-between space-x-4 px-4"),
				html.H4(html.Class("text-sm font-semibold"),
					g.Text("@peduarte starred 3 repositories"),
				),
				TriggerButtonHTMX(TriggerProps{}, htmxProps, isOpen),
			),
			html.Div(html.Class("rounded-md border px-4 py-3 font-mono text-sm"),
				g.Text("@radix-ui/primitives"),
			),
			html.Div(html.Class("rounded-md border px-4 py-3 font-mono text-sm"),
				g.Text("@radix-ui/colors"),
			),
			html.Div(html.Class("rounded-md border px-4 py-3 font-mono text-sm"),
				g.Text("@stitches/react"),
			),
		)
		
		node.Render(w)
	})
	
	// Accordion handlers
	mux.HandleFunc("/api/accordion/", func(w http.ResponseWriter, r *http.Request) {
		// Parse accordion ID and item ID from path
		// Format: /api/accordion/{accordionID}/{itemID}/toggle
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) < 6 {
			http.Error(w, "Invalid path", http.StatusBadRequest)
			return
		}
		
		accordionID := parts[3]
		itemID := parts[4]
		
		r.ParseForm()
		isOpen := r.FormValue("open") == "true"
		
		htmxProps := HTMXProps{
			ID:         accordionID + "-" + itemID,
			TogglePath: r.URL.Path,
		}
		
		// In a real app, you would fetch the item data from a database
		// For this example, we'll use hardcoded data
		var item AccordionItem
		switch itemID {
		case "item-1":
			item = AccordionItem{
				ID:    "item-1",
				Title: "Is it accessible?",
				Content: html.P(html.Class("text-muted-foreground"),
					g.Text("Yes. It adheres to the WAI-ARIA design pattern."),
				),
				Open: isOpen,
			}
		case "item-2":
			item = AccordionItem{
				ID:    "item-2",
				Title: "Is it styled?",
				Content: html.P(html.Class("text-muted-foreground"),
					g.Text("Yes. It comes with default styles that matches the other components' aesthetic."),
				),
				Open: isOpen,
			}
		case "item-3":
			item = AccordionItem{
				ID:    "item-3",
				Title: "Is it animated?",
				Content: html.P(html.Class("text-muted-foreground"),
					g.Text("Yes. It's animated by default, but you can disable it if you prefer."),
				),
				Open: isOpen,
			}
		}
		
		node := RenderAccordionItem(htmxProps, item)
		node.Render(w)
	})
}