package command

import (
	"fmt"
	"strings"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// CommandItem represents an item in the command menu
type CommandItem struct {
	ID          string
	Label       string
	Value       string
	Icon        g.Node
	Shortcut    string
	Category    string
	Description string
	Disabled    bool
	OnSelect    string // JavaScript function or HTMX endpoint
}

// CommandGroup represents a group of command items
type CommandGroup struct {
	ID       string
	Label    string
	Items    []CommandItem
	Expanded bool
}

// Option is a functional option for configuring a command menu
type Option func(*config)

type config struct {
	class          string
	placeholder    string
	emptyMessage   string
	loadingMessage string
	showSearch     bool
	showShortcuts  bool
	showCategories bool
	maxHeight      string
	width          string
	theme          string // light or dark
}

// New creates a new command menu component
func New(id string, groups []CommandGroup, opts ...Option) g.Node {
	cfg := &config{
		class:          "",
		placeholder:    "Type a command or search...",
		emptyMessage:   "No results found.",
		loadingMessage: "Loading...",
		showSearch:     true,
		showShortcuts:  true,
		showCategories: true,
		maxHeight:      "400px",
		width:          "100%",
		theme:          "light",
	}

	for _, opt := range opts {
		opt(cfg)
	}

	classes := strings.TrimSpace("relative overflow-hidden rounded-lg border bg-popover text-popover-foreground shadow-md " + cfg.class)

	return h.Div(
		h.ID(id),
		h.Class(classes),
		g.Attr("data-command", "true"),
		g.Attr("data-theme", cfg.theme),
		h.Style(fmt.Sprintf("width: %s", cfg.width)),
		g.Attr("role", "combobox"),
		g.Attr("aria-expanded", "false"),
		g.Attr("aria-haspopup", "listbox"),
		g.Attr("aria-label", "Command menu"),

		// Search input
		g.If(cfg.showSearch,
			SearchInput(id, cfg),
		),

		// Command list
		CommandList(id, groups, cfg),
	)
}

// SearchInput creates the search input for the command menu
func SearchInput(id string, cfg *config) g.Node {
	return h.Div(
		h.Class("flex items-center border-b px-3"),
		// Search icon
		g.Raw(`<svg class="mr-2 h-4 w-4 shrink-0 opacity-50" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
			<circle cx="11" cy="11" r="8"></circle>
			<path d="m21 21-4.35-4.35"></path>
		</svg>`),

		// Input
		h.Input(
			h.ID(id+"-input"),
			h.Type("text"),
			h.Class("flex h-11 w-full rounded-md bg-transparent py-3 text-sm outline-none placeholder:text-muted-foreground disabled:cursor-not-allowed disabled:opacity-50"),
			h.Placeholder(cfg.placeholder),
			g.Attr("role", "combobox"),
			g.Attr("aria-controls", id+"-list"),
			g.Attr("aria-autocomplete", "list"),
			g.Attr("autocomplete", "off"),
			g.Attr("data-command-input", "true"),
		),
	)
}

// CommandList creates the scrollable list of command items
func CommandList(id string, groups []CommandGroup, cfg *config) g.Node {
	return h.Div(
		h.ID(id+"-list"),
		h.Class("max-h-[300px] overflow-y-auto overflow-x-hidden"),
		h.Style(fmt.Sprintf("max-height: %s", cfg.maxHeight)),
		g.Attr("role", "listbox"),
		g.Attr("aria-label", "Commands"),
		g.Attr("data-command-list", "true"),

		// Empty state
		h.Div(
			h.Class("py-6 text-center text-sm hidden"),
			g.Attr("data-command-empty", "true"),
			g.Text(cfg.emptyMessage),
		),

		// Loading state
		h.Div(
			h.Class("py-6 text-center text-sm hidden"),
			g.Attr("data-command-loading", "true"),
			g.Text(cfg.loadingMessage),
		),

		// Groups
		g.Group(renderGroups(groups, cfg)),
	)
}

// renderGroups renders all command groups
func renderGroups(groups []CommandGroup, cfg *config) []g.Node {
	nodes := make([]g.Node, 0, len(groups))

	for _, group := range groups {
		if len(group.Items) == 0 {
			continue
		}

		nodes = append(nodes, CommandGroupNode(group, cfg))
	}

	return nodes
}

// CommandGroupNode creates a single command group
func CommandGroupNode(group CommandGroup, cfg *config) g.Node {
	return h.Div(
		g.Attr("data-command-group", "true"),
		g.Attr("data-group-id", group.ID),
		g.If(group.ID != "", h.ID(group.ID)),

		// Group heading
		g.If(group.Label != "",
			h.Div(
				h.Class("px-2 py-1.5 text-xs font-medium text-muted-foreground"),
				g.Attr("data-command-group-heading", "true"),
				g.Text(group.Label),
			),
		),

		// Items
		h.Div(
			g.Attr("role", "group"),
			g.Group(renderItems(group.Items, cfg)),
		),
	)
}

// renderItems renders command items
func renderItems(items []CommandItem, cfg *config) []g.Node {
	nodes := make([]g.Node, 0, len(items))

	for _, item := range items {
		nodes = append(nodes, CommandItemNode(item, cfg))
	}

	return nodes
}

// CommandItemNode creates a single command item
func CommandItemNode(item CommandItem, cfg *config) g.Node {
	classes := "relative flex cursor-default select-none items-center rounded-sm px-2 py-1.5 text-sm outline-none"
	if item.Disabled {
		classes += " opacity-50 cursor-not-allowed"
	} else {
		classes += " hover:bg-accent hover:text-accent-foreground data-[selected=true]:bg-accent data-[selected=true]:text-accent-foreground"
	}

	return h.Div(
		h.Class(classes),
		g.Attr("role", "option"),
		g.Attr("aria-selected", "false"),
		g.If(item.ID != "", g.Attr("id", item.ID)),
		g.Attr("data-command-item", "true"),
		g.Attr("data-value", item.Value),
		g.If(item.Category != "", g.Attr("data-category", item.Category)),
		g.If(item.Disabled, g.Attr("aria-disabled", "true")),
		g.If(item.OnSelect != "", g.Attr("data-onselect", item.OnSelect)),

		// Icon
		g.If(item.Icon != nil,
			h.Span(h.Class("mr-2 h-4 w-4"), item.Icon),
		),

		// Label and description
		h.Div(
			h.Class("flex-1"),
			h.Span(g.Text(item.Label)),
			g.If(item.Description != "",
				h.Span(
					h.Class("ml-2 text-xs text-muted-foreground"),
					g.Text(item.Description),
				),
			),
		),

		// Shortcut
		g.If(cfg.showShortcuts && item.Shortcut != "",
			h.Span(
				h.Class("ml-auto text-xs tracking-widest text-muted-foreground"),
				g.Text(item.Shortcut),
			),
		),
	)
}

// Dialog creates a command dialog (modal command menu)
func Dialog(id string, groups []CommandGroup, opts ...Option) g.Node {
	return h.Div(
		h.ID(id+"-dialog"),
		h.Class("fixed inset-0 z-50 hidden"),
		g.Attr("role", "dialog"),
		g.Attr("aria-modal", "true"),
		g.Attr("data-command-dialog", "true"),

		// Backdrop
		h.Div(
			h.Class("fixed inset-0 bg-background/80 backdrop-blur-sm"),
			g.Attr("aria-hidden", "true"),
			g.Attr("data-command-overlay", "true"),
		),

		// Dialog content
		h.Div(
			h.Class("fixed left-[50%] top-[50%] z-50 w-full max-w-lg translate-x-[-50%] translate-y-[-50%] gap-4 p-0"),
			New(id, groups, opts...),
		),
	)
}

// Separator creates a separator between command items
func Separator() g.Node {
	return h.Div(
		h.Class("-mx-1 h-px bg-border"),
		g.Attr("role", "separator"),
		g.Attr("aria-orientation", "horizontal"),
		g.Attr("data-command-separator", "true"),
	)
}

// Option functions

// WithClass adds custom CSS classes
func WithClass(class string) Option {
	return func(c *config) {
		c.class = class
	}
}

// WithPlaceholder sets the search input placeholder
func WithPlaceholder(placeholder string) Option {
	return func(c *config) {
		c.placeholder = placeholder
	}
}

// WithEmptyMessage sets the message shown when no results are found
func WithEmptyMessage(message string) Option {
	return func(c *config) {
		c.emptyMessage = message
	}
}

// WithLoadingMessage sets the loading message
func WithLoadingMessage(message string) Option {
	return func(c *config) {
		c.loadingMessage = message
	}
}

// WithoutSearch hides the search input
func WithoutSearch() Option {
	return func(c *config) {
		c.showSearch = false
	}
}

// WithoutShortcuts hides keyboard shortcuts
func WithoutShortcuts() Option {
	return func(c *config) {
		c.showShortcuts = false
	}
}

// WithMaxHeight sets the maximum height of the command list
func WithMaxHeight(height string) Option {
	return func(c *config) {
		c.maxHeight = height
	}
}

// WithWidth sets the width of the command menu
func WithWidth(width string) Option {
	return func(c *config) {
		c.width = width
	}
}

// WithTheme sets the theme
func WithTheme(theme string) Option {
	return func(c *config) {
		c.theme = theme
	}
}