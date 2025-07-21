package command

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib/icons"
)

// Example creates a basic command menu example
func Example() g.Node {
	groups := []CommandGroup{
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
					Value:    "settings",
					Label:    "Settings",
					Icon:     icons.Settings(html.Class("h-4 w-4")),
					Shortcut: "⌘S",
				},
			},
		},
	}

	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Basic Command Menu")),
		New("command-basic", groups, WithWidth("500px")),
	)
}

// ExampleDialog creates a dialog command menu example
func ExampleDialog() g.Node {
	groups := CreateSampleGroups()

	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Command Dialog")),
		html.P(html.Class("text-sm text-muted-foreground mb-4"), g.Text("Click the button to open the command palette")),
		html.Button(
			html.Type("button"),
			html.Class("inline-flex items-center gap-2 rounded-md border px-3 py-2 text-sm hover:bg-accent"),
			g.Attr("onclick", "document.getElementById('command-dialog-dialog').classList.remove('hidden')"),
			g.Text("Open Command Palette"),
			html.Span(html.Class("text-xs text-muted-foreground"), g.Text("⌘K")),
		),
		Dialog("command-dialog", groups, WithWidth("600px")),
	)
}

// ExampleNoSearch creates a command menu without search
func ExampleNoSearch() g.Node {
	groups := []CommandGroup{
		{
			Items: []CommandItem{
				{Value: "new-file", Label: "New File", Icon: icons.Plus(html.Class("h-4 w-4")), Shortcut: "⌘N"},
				{Value: "new-folder", Label: "New Folder", Icon: icons.Plus(html.Class("h-4 w-4"))},
				{Value: "new-project", Label: "New Project", Icon: icons.Plus(html.Class("h-4 w-4"))},
			},
		},
	}

	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Command Menu Without Search")),
		New("command-no-search", groups, WithoutSearch(), WithWidth("300px")),
	)
}

// ExampleWithCategories creates a command menu with categories
func ExampleWithCategories() g.Node {
	groups := []CommandGroup{
		{
			Label: "Actions",
			Items: []CommandItem{
				{Value: "copy", Label: "Copy", Category: "Edit", Shortcut: "⌘C"},
				{Value: "paste", Label: "Paste", Category: "Edit", Shortcut: "⌘V"},
				{Value: "cut", Label: "Cut", Category: "Edit", Shortcut: "⌘X"},
			},
		},
		{
			Label: "Navigation",
			Items: []CommandItem{
				{Value: "home", Label: "Go to Home", Category: "Navigate"},
				{Value: "back", Label: "Go Back", Category: "Navigate", Shortcut: "⌘["},
				{Value: "forward", Label: "Go Forward", Category: "Navigate", Shortcut: "⌘]"},
			},
		},
	}

	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Command Menu with Categories")),
		New("command-categories", groups, WithWidth("400px")),
	)
}

// ExampleHTMX creates an HTMX-enhanced command menu example
func ExampleHTMX() g.Node {
	htmxConfig := HTMXConfig{
		SearchEndpoint: "/api/command/search",
		SelectEndpoint: "/api/command/select",
		DebounceMs:     300,
	}

	groups := CreateSampleGroups()

	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("HTMX Command Menu")),
		html.P(html.Class("text-sm text-muted-foreground mb-4"), g.Text("Type to search and click to select items")),
		NewHTMX("command-htmx", groups, htmxConfig, WithWidth("500px")),
		html.Div(html.ID("command-result"), html.Class("mt-4")),
	)
}

// ExampleDisabledItems creates a command menu with disabled items
func ExampleDisabledItems() g.Node {
	groups := []CommandGroup{
		{
			Label: "Account",
			Items: []CommandItem{
				{Value: "profile", Label: "View Profile", Icon: icons.User(html.Class("h-4 w-4"))},
				{Value: "edit-profile", Label: "Edit Profile", Icon: icons.Settings(html.Class("h-4 w-4"))},
				{Value: "delete-account", Label: "Delete Account", Icon: icons.X(html.Class("h-4 w-4")), Disabled: true},
			},
		},
		{
			Label: "Subscription",
			Items: []CommandItem{
				{Value: "upgrade", Label: "Upgrade Plan", Icon: icons.ArrowRight(html.Class("h-4 w-4"))},
				{Value: "cancel", Label: "Cancel Subscription", Icon: icons.X(html.Class("h-4 w-4")), Disabled: true},
			},
		},
	}

	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Command Menu with Disabled Items")),
		html.P(html.Class("text-sm text-muted-foreground mb-4"), g.Text("Some items are disabled and cannot be selected")),
		New("command-disabled", groups, WithWidth("400px")),
	)
}

// ExampleCustomStyling creates a command menu with custom styling
func ExampleCustomStyling() g.Node {
	groups := []CommandGroup{
		{
			Label: "Quick Actions",
			Items: []CommandItem{
				{Value: "lock", Label: "Lock Screen", Icon: icons.CircleIcon(html.Class("h-4 w-4")), Shortcut: "⌘L"},
				{Value: "sleep", Label: "Sleep", Icon: icons.CircleIcon(html.Class("h-4 w-4"))},
				{Value: "restart", Label: "Restart", Icon: icons.CircleIcon(html.Class("h-4 w-4"))},
				{Value: "shutdown", Label: "Shut Down", Icon: icons.CircleIcon(html.Class("h-4 w-4"))},
			},
		},
	}

	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Custom Styled Command Menu")),
		New("command-custom", groups,
			WithClass("bg-slate-900 text-slate-100 border-slate-800"),
			WithWidth("350px"),
			WithMaxHeight("250px"),
		),
	)
}

// ExampleEmpty creates an empty command menu example
func ExampleEmpty() g.Node {
	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Empty Command Menu")),
		html.P(html.Class("text-sm text-muted-foreground mb-4"), g.Text("Shows the empty state message")),
		New("command-empty", []CommandGroup{},
			WithEmptyMessage("No commands available at this time."),
			WithWidth("400px"),
		),
	)
}

// ExampleSeparators creates a command menu with separators
func ExampleSeparators() g.Node {
	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Command Menu with Separators")),
		html.Div(
			html.ID("command-separators"),
			html.Class("relative overflow-hidden rounded-lg border bg-popover text-popover-foreground shadow-md w-[400px]"),
			
			// Search input
			SearchInput("command-sep", &config{placeholder: "Type a command..."}),
			
			// Command list with manual separator
			html.Div(
				html.Class("max-h-[300px] overflow-y-auto overflow-x-hidden"),
				
				// First group
				CommandGroupNode(CommandGroup{
					Items: []CommandItem{
						{Value: "new", Label: "New Document", Icon: icons.Plus(html.Class("h-4 w-4"))},
						{Value: "open", Label: "Open", Icon: icons.Plus(html.Class("h-4 w-4"))},
					},
				}, &config{showShortcuts: true}),
				
				// Separator
				Separator(),
				
				// Second group
				CommandGroupNode(CommandGroup{
					Items: []CommandItem{
						{Value: "save", Label: "Save", Icon: icons.Check(html.Class("h-4 w-4")), Shortcut: "⌘S"},
						{Value: "save-as", Label: "Save As...", Icon: icons.Check(html.Class("h-4 w-4")), Shortcut: "⇧⌘S"},
					},
				}, &config{showShortcuts: true}),
				
				// Separator
				Separator(),
				
				// Third group
				CommandGroupNode(CommandGroup{
					Items: []CommandItem{
						{Value: "quit", Label: "Quit", Icon: icons.X(html.Class("h-4 w-4")), Shortcut: "⌘Q"},
					},
				}, &config{showShortcuts: true}),
			),
		),
	)
}