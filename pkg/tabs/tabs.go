package tabs

import (
	"fmt"
	
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// Props defines the properties for the Tabs component
type Props struct {
	DefaultValue string // Default active tab
	Class        string // Additional custom classes
	ID           string // Unique identifier for the tabs
}

// ListProps defines the properties for TabsList
type ListProps struct {
	Class string // Additional custom classes
}

// TriggerProps defines the properties for TabsTrigger
type TriggerProps struct {
	Value    string // Unique value for the tab
	Class    string // Additional custom classes
	Disabled bool   // Whether the tab is disabled
}

// ContentProps defines the properties for TabsContent
type ContentProps struct {
	Value string // Value that corresponds to a trigger
	Class string // Additional custom classes
}

// New creates a new Tabs component
func New(props Props, children ...g.Node) g.Node {
	classes := lib.CN("flex flex-col gap-2", props.Class)
	
	attrs := []g.Node{
		html.Class(classes),
		dataAttr("slot", "tabs"),
	}
	
	if props.ID != "" {
		attrs = append(attrs, html.ID(props.ID))
	}
	
	if props.DefaultValue != "" {
		attrs = append(attrs, dataAttr("tabs-default", props.DefaultValue))
	}
	
	// Add JavaScript for interactivity
	return g.Group([]g.Node{
		html.Div(append(attrs, children...)...),
		tabsScript(props.ID),
	})
}

// TabsList creates a TabsList component
func TabsList(props ListProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"bg-muted text-muted-foreground inline-flex h-9 w-fit items-center justify-center rounded-lg p-[3px]",
		props.Class,
	)
	
	return html.Div(
		html.Class(classes),
		dataAttr("slot", "tabs-list"),
		html.Role("tablist"),
		g.Group(children),
	)
}

// Trigger creates a TabsTrigger component
func Trigger(props TriggerProps, children ...g.Node) g.Node {
	classes := lib.CN(
		"inline-flex h-[calc(100%-1px)] flex-1 items-center justify-center gap-1.5 rounded-md border border-transparent px-2 py-1 text-sm font-medium whitespace-nowrap transition-[color,box-shadow] focus-visible:ring-[3px] focus-visible:outline-1 disabled:pointer-events-none disabled:opacity-50 [&_svg]:pointer-events-none [&_svg]:shrink-0 [&_svg:not([class*='size-'])]:size-4 cursor-pointer",
		props.Class,
	)
	
	attrs := []g.Node{
		html.Type("button"),
		html.Class(classes),
		dataAttr("slot", "tabs-trigger"),
		dataAttr("tabs-value", props.Value),
		dataAttr("state", "inactive"),
		html.Role("tab"),
		ariaAttr("selected", "false"),
		ariaAttr("controls", fmt.Sprintf("content-%s", props.Value)),
		html.TabIndex("-1"),
	}
	
	if props.Disabled {
		attrs = append(attrs, html.Disabled())
	}
	
	return html.Button(
		append(attrs, g.Group(children))...,
	)
}

// TabsContent creates a TabsContent component
func TabsContent(props ContentProps, children ...g.Node) g.Node {
	classes := lib.CN("flex-1 outline-none", props.Class)
	
	return html.Div(
		html.Class(classes),
		dataAttr("slot", "tabs-content"),
		dataAttr("tabs-value", props.Value),
		dataAttr("state", "inactive"),
		html.Role("tabpanel"),
		html.ID(fmt.Sprintf("content-%s", props.Value)),
		ariaAttr("labelledby", fmt.Sprintf("trigger-%s", props.Value)),
		html.TabIndex("0"),
		html.Style("display: none;"),
		g.Group(children),
	)
}

// Helper functions

// WithDefault creates tabs with a default active tab
func WithDefault(defaultValue string, children ...g.Node) g.Node {
	return New(Props{DefaultValue: defaultValue}, children...)
}

// dataAttr creates a data attribute
func dataAttr(name, value string) g.Node {
	return g.Attr("data-" + name, value)
}

// ariaAttr creates an aria attribute
func ariaAttr(name, value string) g.Node {
	return g.Attr("aria-" + name, value)
}

// tabsScript generates the JavaScript for tabs functionality
func tabsScript(id string) g.Node {
	selector := ""
	if id != "" {
		selector = fmt.Sprintf("#%s", id)
	} else {
		selector = "[data-slot='tabs']"
	}
	
	return html.Script(g.Raw(fmt.Sprintf(`
	(function() {
		const tabs = document.querySelector('%s');
		if (!tabs) return;
		
		const defaultValue = tabs.dataset.tabsDefault;
		const triggers = tabs.querySelectorAll('[data-slot="tabs-trigger"]');
		const contents = tabs.querySelectorAll('[data-slot="tabs-content"]');
		
		// Helper function to activate a tab
		function activateTab(value) {
			triggers.forEach(trigger => {
				const triggerValue = trigger.dataset.tabsValue;
				const isActive = triggerValue === value;
				
				trigger.dataset.state = isActive ? 'active' : 'inactive';
				trigger.setAttribute('aria-selected', isActive ? 'true' : 'false');
				trigger.setAttribute('tabindex', isActive ? '0' : '-1');
				
				// Update styling
				if (isActive) {
					trigger.classList.add('data-[state=active]:bg-background', 'dark:data-[state=active]:text-foreground', 'dark:data-[state=active]:border-input', 'dark:data-[state=active]:bg-input/30', 'text-foreground', 'data-[state=active]:shadow-sm');
					trigger.classList.remove('dark:text-muted-foreground');
				} else {
					trigger.classList.remove('data-[state=active]:bg-background', 'dark:data-[state=active]:text-foreground', 'dark:data-[state=active]:border-input', 'dark:data-[state=active]:bg-input/30', 'text-foreground', 'data-[state=active]:shadow-sm');
					trigger.classList.add('dark:text-muted-foreground');
				}
			});
			
			contents.forEach(content => {
				const contentValue = content.dataset.tabsValue;
				const isActive = contentValue === value;
				
				content.dataset.state = isActive ? 'active' : 'inactive';
				content.style.display = isActive ? 'block' : 'none';
			});
		}
		
		// Initialize default tab
		if (defaultValue) {
			activateTab(defaultValue);
		} else if (triggers.length > 0) {
			// If no default, activate the first tab
			const firstValue = triggers[0].dataset.tabsValue;
			activateTab(firstValue);
		}
		
		// Add click handlers to triggers
		triggers.forEach(trigger => {
			trigger.addEventListener('click', () => {
				if (!trigger.disabled) {
					const value = trigger.dataset.tabsValue;
					activateTab(value);
				}
			});
			
			// Add keyboard navigation
			trigger.addEventListener('keydown', (e) => {
				if (trigger.disabled) return;
				
				let newIndex = -1;
				const currentIndex = Array.from(triggers).indexOf(trigger);
				
				switch (e.key) {
					case 'ArrowLeft':
						e.preventDefault();
						newIndex = currentIndex - 1;
						if (newIndex < 0) newIndex = triggers.length - 1;
						break;
					case 'ArrowRight':
						e.preventDefault();
						newIndex = currentIndex + 1;
						if (newIndex >= triggers.length) newIndex = 0;
						break;
					case 'Home':
						e.preventDefault();
						newIndex = 0;
						break;
					case 'End':
						e.preventDefault();
						newIndex = triggers.length - 1;
						break;
				}
				
				if (newIndex !== -1) {
					const newTrigger = triggers[newIndex];
					if (!newTrigger.disabled) {
						newTrigger.click();
						newTrigger.focus();
					}
				}
			});
		});
	})();
	`, selector)))
}