package accordion

import (
	"fmt"
	
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// Props defines the properties for the Accordion component
type Props struct {
	Type         string   // "single" | "multiple"
	Collapsible  bool     // Whether the accordion items can be collapsed
	DefaultValue string   // Default open item(s) for single mode
	DefaultOpen  []string // Default open items for multiple mode
	Class        string   // Additional custom classes
	ID           string   // Unique identifier for the accordion
}

// ItemProps defines the properties for an AccordionItem
type ItemProps struct {
	Value string // Unique value for the item
	Class string // Additional custom classes
}

// TriggerProps defines the properties for an AccordionTrigger
type TriggerProps struct {
	Class string // Additional custom classes
	Icon  g.Node // Custom icon (optional)
}

// ContentProps defines the properties for AccordionContent
type ContentProps struct {
	Class string // Additional custom classes
}

// New creates a new Accordion component
func New(props Props, children ...g.Node) g.Node {
	classes := lib.CN("w-full", props.Class)
	
	attrs := []g.Node{
		html.Class(classes),
		dataAttr("slot", "accordion"),
		dataAttr("accordion-type", props.Type),
	}
	
	if props.ID != "" {
		attrs = append(attrs, html.ID(props.ID))
	}
	
	if props.Collapsible {
		attrs = append(attrs, dataAttr("accordion-collapsible", "true"))
	}
	
	if props.DefaultValue != "" {
		attrs = append(attrs, dataAttr("accordion-default", props.DefaultValue))
	}
	
	// Add JavaScript for interactivity
	return g.Group([]g.Node{
		html.Div(append(attrs, children...)...),
		accordionScript(props.ID),
	})
}

// Item creates an AccordionItem component
func Item(props ItemProps, children ...g.Node) g.Node {
	classes := lib.CN("border-b last:border-b-0", props.Class)
	
	return html.Div(
		html.Class(classes),
		dataAttr("slot", "accordion-item"),
		dataAttr("accordion-value", props.Value),
		g.Group(children),
	)
}

// Trigger creates an AccordionTrigger component
func Trigger(props TriggerProps, children ...g.Node) g.Node {
	triggerClasses := lib.CN(
		"flex flex-1 items-start justify-between gap-4 rounded-md py-4 text-left text-sm font-medium transition-all outline-none hover:underline focus-visible:ring-[3px] focus-visible:ring-ring/50 focus-visible:border-ring disabled:pointer-events-none disabled:opacity-50 cursor-pointer",
		props.Class,
	)
	
	// Default chevron icon
	icon := props.Icon
	if icon == nil {
		icon = chevronIcon()
	}
	
	return html.Div(
		html.Class("flex"),
		html.Button(
			html.Type("button"),
			html.Class(triggerClasses),
			dataAttr("slot", "accordion-trigger"),
			AriaAttr("expanded", "false"),
			g.Group(children),
			icon,
		),
	)
}

// ItemContent creates an AccordionContent component
func ItemContent(props ContentProps, children ...g.Node) g.Node {
	contentClasses := lib.CN("overflow-hidden text-sm transition-all duration-200", props.Class)
	innerClasses := lib.CN("pt-0 pb-4", props.Class)
	
	return html.Div(
		html.Class(contentClasses),
		dataAttr("slot", "accordion-content"),
		dataAttr("state", "closed"),
		html.Style("max-height: 0;"),
		html.Div(
			html.Class(innerClasses),
			g.Group(children),
		),
	)
}

// Helper functions

// chevronIcon creates the default chevron icon
func chevronIcon() g.Node {
	return g.Raw(`<svg class="text-muted-foreground pointer-events-none size-4 shrink-0 translate-y-0.5 transition-transform duration-200" data-accordion-icon xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m6 9 6 6 6-6"/></svg>`)
}

// accordionScript generates the JavaScript for accordion functionality
func accordionScript(id string) g.Node {
	selector := ""
	if id != "" {
		selector = fmt.Sprintf("#%s", id)
	} else {
		selector = "[data-slot='accordion']"
	}
	
	return html.Script(g.Raw(fmt.Sprintf(`
	(function() {
		const accordion = document.querySelector('%s');
		if (!accordion) return;
		
		const type = accordion.dataset.accordionType || 'single';
		const collapsible = accordion.dataset.accordionCollapsible === 'true';
		const defaultValue = accordion.dataset.accordionDefault;
		
		const items = accordion.querySelectorAll('[data-slot="accordion-item"]');
		const openItems = new Set();
		
		// Initialize default open items
		if (defaultValue) {
			openItems.add(defaultValue);
			const item = accordion.querySelector('[data-accordion-value="' + defaultValue + '"]');
			if (item) {
				const trigger = item.querySelector('[data-slot="accordion-trigger"]');
				const content = item.querySelector('[data-slot="accordion-content"]');
				const icon = trigger.querySelector('[data-accordion-icon]');
				
				trigger.setAttribute('aria-expanded', 'true');
				content.dataset.state = 'open';
				content.style.maxHeight = content.scrollHeight + 'px';
				if (icon) {
					icon.style.transform = 'rotate(180deg)';
				}
			}
		}
		
		items.forEach(item => {
			const trigger = item.querySelector('[data-slot="accordion-trigger"]');
			const content = item.querySelector('[data-slot="accordion-content"]');
			const value = item.dataset.accordionValue;
			const icon = trigger.querySelector('[data-accordion-icon]');
			
			trigger.addEventListener('click', () => {
				const isOpen = openItems.has(value);
				
				if (type === 'single') {
					// Close all other items
					items.forEach(otherItem => {
						const otherValue = otherItem.dataset.accordionValue;
						if (otherValue !== value && openItems.has(otherValue)) {
							const otherTrigger = otherItem.querySelector('[data-slot="accordion-trigger"]');
							const otherContent = otherItem.querySelector('[data-slot="accordion-content"]');
							const otherIcon = otherTrigger.querySelector('[data-accordion-icon]');
							
							openItems.delete(otherValue);
							otherTrigger.setAttribute('aria-expanded', 'false');
							otherContent.dataset.state = 'closed';
							otherContent.style.maxHeight = '0';
							if (otherIcon) {
								otherIcon.style.transform = 'rotate(0deg)';
							}
						}
					});
				}
				
				if (isOpen && (!collapsible && type === 'single')) {
					// Can't close if not collapsible in single mode
					return;
				}
				
				if (isOpen) {
					openItems.delete(value);
					trigger.setAttribute('aria-expanded', 'false');
					content.dataset.state = 'closed';
					content.style.maxHeight = '0';
					if (icon) {
						icon.style.transform = 'rotate(0deg)';
					}
				} else {
					openItems.add(value);
					trigger.setAttribute('aria-expanded', 'true');
					content.dataset.state = 'open';
					content.style.maxHeight = content.scrollHeight + 'px';
					if (icon) {
						icon.style.transform = 'rotate(180deg)';
					}
				}
			});
		});
	})();
	`, selector)))
}

// Helper functions for common patterns

// Single creates a single-select accordion
func Single(collapsible bool, defaultValue string, children ...g.Node) g.Node {
	return New(Props{
		Type:         "single",
		Collapsible:  collapsible,
		DefaultValue: defaultValue,
	}, children...)
}

// MultipleAccordion creates a multi-select accordion
func MultipleAccordion(defaultOpen []string, children ...g.Node) g.Node {
	return New(Props{
		Type:        "multiple",
		Collapsible: true,
		DefaultOpen: defaultOpen,
	}, children...)
}

// dataAttr creates a data attribute
func dataAttr(name, value string) g.Node {
	return g.Attr("data-" + name, value)
}

// AriaAttr creates an aria attribute
func AriaAttr(name, value string) g.Node {
	return g.Attr("aria-" + name, value)
}