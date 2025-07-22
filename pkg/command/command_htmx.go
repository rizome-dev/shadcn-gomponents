package command

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

// HTMXConfig contains HTMX-specific configuration
type HTMXConfig struct {
	SearchEndpoint   string // Endpoint for search requests
	SelectEndpoint   string // Endpoint for item selection
	FilterEndpoint   string // Endpoint for filtering
	DebounceMs       int    // Debounce time for search
	MinSearchLength  int    // Minimum characters before searching
}

// NewHTMX creates an HTMX-enhanced command menu
func NewHTMX(id string, groups []CommandGroup, htmxCfg HTMXConfig, opts ...Option) g.Node {
	cfg := &config{
		class:          "",
		placeholder:    "Type a command or searchtml...",
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

	if htmxCfg.DebounceMs == 0 {
		htmxCfg.DebounceMs = 300
	}
	if htmxCfg.MinSearchLength == 0 {
		htmxCfg.MinSearchLength = 1
	}

	classes := lib.CN(
		"relative overflow-hidden rounded-lg border bg-popover text-popover-foreground shadow-md",
		cfg.class,
	)

	return html.Div(
		html.ID(id),
		html.Class(classes),
		g.Attr("data-command", "true"),
		g.Attr("data-theme", cfg.theme),
		html.Style(fmt.Sprintf("width: %s", cfg.width)),
		g.Attr("role", "combobox"),
		g.Attr("aria-expanded", "false"),
		g.Attr("aria-haspopup", "listbox"),
		g.Attr("aria-label", "Command menu"),

		// Search input
		g.If(cfg.showSearch,
			SearchInputHTMX(id, cfg, htmxCfg),
		),

		// Command list
		CommandListHTMX(id, groups, cfg, htmxCfg),
	)
}

// SearchInputHTMX creates the HTMX-enhanced search input
func SearchInputHTMX(id string, cfg *config, htmxCfg HTMXConfig) g.Node {
	return html.Div(
		html.Class("flex items-center border-b px-3"),
		icons.Search(html.Class("mr-2 h-4 w-4 shrink-0 opacity-50")),

		// Input with HTMX
		html.Input(
			html.ID(id+"-input"),
			html.Type("text"),
			html.Class("flex h-11 w-full rounded-md bg-transparent py-3 text-sm outline-none placeholder:text-muted-foreground disabled:cursor-not-allowed disabled:opacity-50"),
			html.Placeholder(cfg.placeholder),
			g.Attr("role", "combobox"),
			g.Attr("aria-controls", id+"-list"),
			g.Attr("aria-autocomplete", "list"),
			g.Attr("autocomplete", "off"),
			g.Attr("data-command-input", "true"),
			// HTMX attributes
			hx.Get(htmxCfg.SearchEndpoint),
			hx.Trigger(fmt.Sprintf("keyup changed delay:%dms", htmxCfg.DebounceMs)),
			hx.Target("#"+id+"-list"),
			hx.Swap("innerHTML"),
			hx.Include("#"+id+"-input"),
			hx.Indicator("#"+id+"-loading"),
			g.Attr("name", "search"),
		),

		// Loading indicator
		html.Div(
			html.ID(id+"-loading"),
			html.Class("htmx-indicator"),
			icons.Loader(html.Class("h-4 w-4 animate-spin opacity-50")),
		),
	)
}

// CommandListHTMX creates the HTMX-enhanced command list
func CommandListHTMX(id string, groups []CommandGroup, cfg *config, htmxCfg HTMXConfig) g.Node {
	return html.Div(
		html.ID(id+"-list"),
		html.Class("max-h-[300px] overflow-y-auto overflow-x-hidden"),
		html.Style(fmt.Sprintf("max-height: %s", cfg.maxHeight)),
		g.Attr("role", "listbox"),
		g.Attr("aria-label", "Commands"),
		g.Attr("data-command-list", "true"),

		// Initial content
		CommandListContent(groups, cfg, htmxCfg),
	)
}

// CommandListContent renders the content of the command list
func CommandListContent(groups []CommandGroup, cfg *config, htmxCfg HTMXConfig) g.Node {
	if len(groups) == 0 {
		return html.Div(
			html.Class("py-6 text-center text-sm"),
			g.Attr("data-command-empty", "true"),
			g.Text(cfg.emptyMessage),
		)
	}

	return g.Group(renderGroupsHTMX(groups, cfg, htmxCfg))
}

// renderGroupsHTMX renders command groups with HTMX
func renderGroupsHTMX(groups []CommandGroup, cfg *config, htmxCfg HTMXConfig) []g.Node {
	nodes := make([]g.Node, 0, len(groups))

	for _, group := range groups {
		if len(group.Items) == 0 {
			continue
		}

		nodes = append(nodes, CommandGroupNodeHTMX(group, cfg, htmxCfg))
	}

	return nodes
}

// CommandGroupNodeHTMX creates an HTMX-enhanced command group
func CommandGroupNodeHTMX(group CommandGroup, cfg *config, htmxCfg HTMXConfig) g.Node {
	return html.Div(
		g.Attr("data-command-group", "true"),
		g.Attr("data-group-id", group.ID),
		g.If(group.ID != "", html.ID(group.ID)),

		// Group heading
		g.If(group.Label != "",
			html.Div(
				html.Class("px-2 py-1.5 text-xs font-medium text-muted-foreground"),
				g.Attr("data-command-group-heading", "true"),
				g.Text(group.Label),
			),
		),

		// Items
		html.Div(
			g.Attr("role", "group"),
			g.Group(renderItemsHTMX(group.Items, cfg, htmxCfg)),
		),
	)
}

// renderItemsHTMX renders command items with HTMX
func renderItemsHTMX(items []CommandItem, cfg *config, htmxCfg HTMXConfig) []g.Node {
	nodes := make([]g.Node, 0, len(items))

	for _, item := range items {
		nodes = append(nodes, CommandItemNodeHTMX(item, cfg, htmxCfg))
	}

	return nodes
}

// CommandItemNodeHTMX creates an HTMX-enhanced command item
func CommandItemNodeHTMX(item CommandItem, cfg *config, htmxCfg HTMXConfig) g.Node {
	classes := lib.CN(
		"relative flex cursor-default select-none items-center rounded-sm px-2 py-1.5 text-sm outline-none",
		lib.CNIf(item.Disabled,
			"opacity-50 cursor-not-allowed",
			"hover:bg-accent hover:text-accent-foreground data-[selected=true]:bg-accent data-[selected=true]:text-accent-foreground cursor-pointer",
		),
	)

	itemNode := html.Div(
		html.Class(classes),
		g.Attr("role", "option"),
		g.Attr("aria-selected", "false"),
		g.If(item.ID != "", g.Attr("id", item.ID)),
		g.Attr("data-command-item", "true"),
		g.Attr("data-value", item.Value),
		g.If(item.Category != "", g.Attr("data-category", item.Category)),
		g.If(item.Disabled, g.Attr("aria-disabled", "true")),

		// Icon
		g.If(item.Icon != nil,
			html.Span(html.Class("mr-2 h-4 w-4"), item.Icon),
		),

		// Label and description
		html.Div(
			html.Class("flex-1"),
			html.Span(g.Text(item.Label)),
			g.If(item.Description != "",
				html.Span(
					html.Class("ml-2 text-xs text-muted-foreground"),
					g.Text(item.Description),
				),
			),
		),

		// Shortcut
		g.If(cfg.showShortcuts && item.Shortcut != "",
			html.Span(
				html.Class("ml-auto text-xs tracking-widest text-muted-foreground"),
				g.Text(item.Shortcut),
			),
		),
	)

	// Add HTMX attributes if not disabled
	if !item.Disabled && htmxCfg.SelectEndpoint != "" {
		itemNode = g.Group([]g.Node{
			html.Form(
				hx.Post(htmxCfg.SelectEndpoint),
				hx.Trigger("click"),
				html.Input(html.Type("hidden"), html.Name("value"), html.Value(item.Value)),
				html.Input(html.Type("hidden"), html.Name("label"), html.Value(item.Label)),
				g.If(item.Category != "", html.Input(html.Type("hidden"), html.Name("category"), html.Value(item.Category))),
				itemNode,
			),
		})
	}

	return itemNode
}

// DialogHTMX creates an HTMX-enhanced command dialog
func DialogHTMX(id string, groups []CommandGroup, htmxCfg HTMXConfig, opts ...Option) g.Node {
	return html.Div(
		html.ID(id+"-dialog"),
		html.Class("fixed inset-0 z-50 hidden"),
		g.Attr("role", "dialog"),
		g.Attr("aria-modal", "true"),
		g.Attr("data-command-dialog", "true"),

		// Backdrop
		html.Div(
			html.Class("fixed inset-0 bg-background/80 backdrop-blur-sm"),
			g.Attr("aria-hidden", "true"),
			g.Attr("data-command-overlay", "true"),
			hx.On("click", fmt.Sprintf("document.getElementById('%s-dialog').classList.add('hidden')", id)),
		),

		// Dialog content
		html.Div(
			html.Class("fixed left-[50%] top-[50%] z-50 w-full max-w-lg translate-x-[-50%] translate-y-[-50%] gap-4 p-0"),
			hx.On("click", "event.stopPropagation()"),
			NewHTMX(id, groups, htmxCfg, opts...),
		),
	)
}

// RenderSearchResults renders search results for HTMX requests
func RenderSearchResults(query string, allGroups []CommandGroup, cfg *config, htmxCfg HTMXConfig) g.Node {
	if query == "" {
		return CommandListContent(allGroups, cfg, htmxCfg)
	}

	// Filter items based on query
	filteredGroups := FilterGroups(allGroups, query)

	if len(filteredGroups) == 0 {
		return html.Div(
			html.Class("py-6 text-center text-sm"),
			g.Attr("data-command-empty", "true"),
			g.Text(cfg.emptyMessage),
		)
	}

	return CommandListContent(filteredGroups, cfg, htmxCfg)
}

// FilterGroups filters command groups based on a search query
func FilterGroups(groups []CommandGroup, query string) []CommandGroup {
	query = strings.ToLower(strings.TrimSpace(query))
	if query == "" {
		return groups
	}

	var filtered []CommandGroup

	for _, group := range groups {
		var filteredItems []CommandItem

		for _, item := range group.Items {
			if matchesQuery(item, query) {
				filteredItems = append(filteredItems, item)
			}
		}

		if len(filteredItems) > 0 {
			filtered = append(filtered, CommandGroup{
				ID:       group.ID,
				Label:    group.Label,
				Items:    filteredItems,
				Expanded: group.Expanded,
			})
		}
	}

	return filtered
}

// matchesQuery checks if a command item matches the search query
func matchesQuery(item CommandItem, query string) bool {
	// Check label
	if strings.Contains(strings.ToLower(item.Label), query) {
		return true
	}

	// Check value
	if strings.Contains(strings.ToLower(item.Value), query) {
		return true
	}

	// Check description
	if strings.Contains(strings.ToLower(item.Description), query) {
		return true
	}

	// Check category
	if strings.Contains(strings.ToLower(item.Category), query) {
		return true
	}

	return false
}

// CommandHandlers creates HTTP handlers for command menu
func CommandHandlers(mux *http.ServeMux) {
	// Search handler
	mux.HandleFunc("/api/command/search", func(w http.ResponseWriter, r *http.Request) {
		query := r.FormValue("search")

		// Example command groups
		groups := []CommandGroup{
			{
				Label: "Suggestions",
				Items: []CommandItem{
					{Value: "calendar", Label: "Calendar", Icon: icons.Calendar(html.Class("h-4 w-4"))},
					{Value: "search-emoji", Label: "Search Emoji", Icon: icons.Dot(html.Class("h-4 w-4"))},
					{Value: "calculator", Label: "Calculator", Icon: icons.CircleIcon(html.Class("h-4 w-4"))},
				},
			},
			{
				Label: "Settings",
				Items: []CommandItem{
					{Value: "profile", Label: "Profile", Icon: icons.User(html.Class("h-4 w-4")), Shortcut: "⌘P"},
					{Value: "billing", Label: "Billing", Icon: icons.CreditCard(html.Class("h-4 w-4")), Shortcut: "⌘B"},
					{Value: "settings", Label: "Settings", Icon: icons.Settings(html.Class("h-4 w-4")), Shortcut: "⌘S"},
				},
			},
		}

		cfg := &config{
			emptyMessage:  "No results found.",
			showShortcuts: true,
		}

		htmxCfg := HTMXConfig{
			SelectEndpoint: "/api/command/select",
		}

		result := RenderSearchResults(query, groups, cfg, htmxCfg)
		result.Render(w)
	})

	// Select handler
	mux.HandleFunc("/api/command/select", func(w http.ResponseWriter, r *http.Request) {
		value := r.FormValue("value")
		label := r.FormValue("label")
		category := r.FormValue("category")

		// In a real app, handle the selection
		response := html.Div(
			html.Class("p-4 bg-green-50 border border-green-200 rounded-md"),
			html.Div(
				html.Class("text-green-800"),
				g.Textf("Selected: %s", label),
				g.If(category != "", g.Group([]g.Node{
					g.Text(" ("),
					html.Span(html.Class("text-green-600"), g.Text(category)),
					g.Text(")"),
				})),
			),
			html.Div(
				html.Class("text-sm text-green-600 mt-1"),
				g.Textf("Value: %s", value),
			),
		)

		response.Render(w)
	})
}

// CreateSampleGroups creates sample command groups for examples
func CreateSampleGroups() []CommandGroup {
	return []CommandGroup{
		{
			Label: "Suggestions",
			Items: []CommandItem{
				{
					Value:       "calendar",
					Label:       "Calendar",
					Description: "View your calendar",
					Icon:        icons.Calendar(html.Class("h-4 w-4")),
				},
				{
					Value:       "search",
					Label:       "Search",
					Description: "Search for anything",
					Icon:        icons.Search(html.Class("h-4 w-4")),
				},
				{
					Value:       "calculator",
					Label:       "Calculator",
					Description: "Open calculator",
					Icon:        icons.CircleIcon(html.Class("h-4 w-4")),
				},
			},
		},
		{
			Label: "Settings",
			Items: []CommandItem{
				{
					Value:    "profile",
					Label:    "Profile",
					Icon:     icons.User(html.Class("h-4 w-4")),
					Shortcut: "⌘P",
				},
				{
					Value:    "billing",
					Label:    "Billing",
					Icon:     icons.CreditCard(html.Class("h-4 w-4")),
					Shortcut: "⌘B",
				},
				{
					Value:    "settings",
					Label:    "Settings",
					Icon:     icons.Settings(html.Class("h-4 w-4")),
					Shortcut: "⌘S",
				},
			},
		},
		{
			Label: "Team",
			Items: []CommandItem{
				{
					Value:       "invite",
					Label:       "Invite users",
					Description: "Invite new team members",
					Icon:        icons.UserPlus(html.Class("h-4 w-4")),
				},
				{
					Value:       "team",
					Label:       "Team settings",
					Description: "Manage team settings",
					Icon:        icons.Users(html.Class("h-4 w-4")),
					Disabled:    true,
				},
			},
		},
	}
}