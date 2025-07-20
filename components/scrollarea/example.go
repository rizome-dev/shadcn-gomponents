package scrollarea

import (
	"fmt"
	
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

// Example demonstrates how to use the ScrollArea component
func Example() g.Node {
	return html.Div(
		html.Class("space-y-8 p-8"),
		
		// Basic vertical scroll area
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Basic Vertical ScrollArea")),
			ScrollAreaWithBar(
				Props{Class: "h-72 w-96 rounded-md border"},
				Viewport(
					ViewportProps{},
					html.Div(
						html.Class("p-4"),
						html.H4(html.Class("mb-4 text-sm font-medium leading-none"), g.Text("Tags")),
						g.Group(generateTags()),
					),
				),
			),
		),
		
		// Horizontal scroll area
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Horizontal ScrollArea")),
			HorizontalScrollArea(
				Props{Class: "w-96 rounded-md border"},
				html.Div(
					html.Class("flex w-max space-x-4 p-4"),
					g.Group(generateHorizontalItems()),
				),
			),
		),
		
		// Auto-hiding scrollbar
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Auto-hiding Scrollbar")),
			ScrollAreaAuto(
				Props{Class: "h-72 w-96 rounded-md border"},
				Viewport(
					ViewportProps{},
					html.Div(
						html.Class("p-4"),
						html.P(html.Class("text-sm"), g.Text("This scroll area automatically hides the scrollbar when not scrolling.")),
						html.Div(html.Class("mt-4"), g.Group(generateParagraphs())),
					),
				),
			),
		),
		
		// Hover scrollbar
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Hover Scrollbar")),
			ScrollAreaHover(
				Props{Class: "h-72 w-96 rounded-md border"},
				Viewport(
					ViewportProps{},
					html.Div(
						html.Class("p-4"),
						html.P(html.Class("text-sm mb-4"), g.Text("The scrollbar appears on hover and disappears after a delay.")),
						html.Div(html.Class("space-y-4"), g.Group(generateCards())),
					),
				),
			),
		),
		
		// Both scrollbars
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Both Scrollbars")),
			ScrollAreaWithBar(
				Props{
					Orientation: "both",
					Class:       "h-72 w-96 rounded-md border",
				},
				Viewport(
					ViewportProps{},
					html.Div(
						html.Class("p-4"),
						html.Table(
							html.Class("w-[700px]"),
							generateLargeTable(),
						),
					),
				),
			),
		),
		
		// Code block with scroll
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Code ScrollArea")),
			CodeScrollArea(
				"h-72 w-full max-w-2xl",
				g.Text(`package main

import (
    "fmt"
    "net/http"
    "log"
)

func main() {
    // Create a new HTTP server
    mux := http.NewServeMux()
    
    // Register handlers
    mux.HandleFunc("/", homeHandler)
    mux.HandleFunc("/api/users", usersHandler)
    mux.HandleFunc("/api/posts", postsHandler)
    
    // Start server
    port := ":8080"
    fmt.Printf("Server starting on port %s\n", port)
    if err := http.ListenAndServe(port, mux); err != nil {
        log.Fatal(err)
    }
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    fmt.Fprintf(w, "<h1>Welcome to the API</h1>")
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
    // Handle user requests
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, "{\"users\": [{\"id\": 1, \"name\": \"John\"}, {\"id\": 2, \"name\": \"Jane\"}]}")
}

func postsHandler(w http.ResponseWriter, r *http.Request) {
    // Handle post requests
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, "{\"posts\": [{\"id\": 1, \"title\": \"Hello World\"}, {\"id\": 2, \"title\": \"Second Post\"}]}")
}`),
			),
		),
		
		// List scroll area
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("List ScrollArea")),
			ListScrollArea(
				"300px",
				generateListItems()...,
			),
		),
		
		// Chat scroll area
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Chat ScrollArea")),
			html.Div(
				html.Class("h-96 w-full max-w-md rounded-lg border bg-background"),
				ChatScrollArea(
					generateChatMessages()...,
				),
			),
		),
		
		// Image gallery scroll area
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Image Gallery ScrollArea")),
			ImageGalleryScrollArea(
				generateImagePlaceholders()...,
			),
		),
		
		// Table scroll area
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Table ScrollArea")),
			TableScrollArea(
				html.Table(
					html.Class("w-full"),
					html.THead(
						html.Tr(
							html.Th(html.Class("sticky left-0 bg-background"), g.Text("Name")),
							html.Th(g.Text("Status")),
							html.Th(g.Text("Email")),
							html.Th(g.Text("Department")),
							html.Th(g.Text("Role")),
							html.Th(g.Text("Start Date")),
							html.Th(g.Text("Salary")),
						),
					),
					html.TBody(
						g.Group(generateTableRows()),
					),
				),
			),
		),
		
		// Custom styled scrollbar
		html.Div(
			html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Custom Styled Scrollbar")),
			New(
				Props{Class: "h-72 w-96 rounded-md border bg-slate-50 dark:bg-slate-900"},
				Viewport(
					ViewportProps{},
					html.Div(
						html.Class("p-4"),
						html.H4(html.Class("mb-4 text-sm font-medium"), g.Text("Custom Colors")),
						html.P(html.Class("text-sm text-muted-foreground"), 
							g.Text("This scroll area has custom styling applied to match the theme."),
						),
						html.Div(html.Class("mt-4 space-y-2"), g.Group(generateColoredItems())),
					),
				),
				Scrollbar(
					ScrollbarProps{
						Orientation: "vertical",
						Class:       "bg-transparent",
					},
					Thumb(ThumbProps{Class: "bg-primary/20 hover:bg-primary/30"}),
				),
			),
		),
	)
}

// Helper functions to generate content

func generateTags() []g.Node {
	tags := []string{
		"React", "Vue", "Angular", "Svelte", "Next.js", "Nuxt", "Gatsby",
		"Remix", "Astro", "SolidJS", "Qwik", "Alpine.js", "Lit", "Ember",
		"Backbone", "Meteor", "Polymer", "Aurelia", "Knockout", "Mithril",
		"Stimulus", "Stencil", "Hyperapp", "Preact", "Inferno", "Cycle.js",
		"Riot", "Dojo", "Marko", "Elm",
	}
	
	nodes := make([]g.Node, len(tags))
	for i, tag := range tags {
		nodes[i] = html.Div(
			html.Class("text-sm mb-2"),
			g.Text(tag),
		)
	}
	return nodes
}

func generateHorizontalItems() []g.Node {
	items := make([]g.Node, 20)
	for i := 0; i < 20; i++ {
		items[i] = html.Div(
			html.Class("w-32 h-32 rounded-lg bg-muted flex items-center justify-center shrink-0"),
			html.Span(html.Class("text-2xl font-semibold"), g.Textf("Item %d", i+1)),
		)
	}
	return items
}

func generateParagraphs() []g.Node {
	return []g.Node{
		html.P(html.Class("mb-4"), g.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec a diam lectus. Sed sit amet ipsum mauris. Maecenas congue ligula ac quam viverra nec consectetur ante hendrerit.")),
		html.P(html.Class("mb-4"), g.Text("Donec et mollis dolor. Praesent et diam eget libero egestas mattis sit amet vitae augue. Nam tincidunt congue enim, ut porta lorem lacinia consectetur.")),
		html.P(html.Class("mb-4"), g.Text("Donec ut libero sed arcu vehicula ultricies a non tortor. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean ut gravida lorem.")),
		html.P(html.Class("mb-4"), g.Text("Ut turpis felis, pulvinar a semper sed, adipiscing id dolor. Pellentesque auctor nisi id magna consequat sagittis. Curabitur dapibus enim sit amet elit pharetra tincidunt feugiat nisl imperdiet.")),
		html.P(html.Class("mb-4"), g.Text("Ut convallis libero in urna ultrices accumsan. Donec sed odio eros. Donec viverra mi quis quam pulvinar at malesuada arcu rhoncus.")),
		html.P(g.Text("Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. In rutrum accumsan ultricies. Mauris vitae nisi at sem facilisis semper ac in est.")),
	}
}

func generateCards() []g.Node {
	cards := make([]g.Node, 10)
	for i := 0; i < 10; i++ {
		cards[i] = html.Div(
			html.Class("rounded-lg border bg-card p-4"),
			html.H4(html.Class("font-semibold"), g.Textf("Card %d", i+1)),
			html.P(html.Class("text-sm text-muted-foreground mt-2"), 
				g.Text("This is a card with some content that demonstrates the scroll area."),
			),
		)
	}
	return cards
}

func generateLargeTable() g.Node {
	rows := make([]g.Node, 50)
	for i := 0; i < 50; i++ {
		cells := make([]g.Node, 10)
		for j := 0; j < 10; j++ {
			cells[j] = html.Td(html.Class("p-2 border"), g.Textf("Cell %d-%d", i+1, j+1))
		}
		rows[i] = html.Tr(cells...)
	}
	return html.TBody(rows...)
}

func generateListItems() []g.Node {
	items := make([]g.Node, 50)
	for i := 0; i < 50; i++ {
		items[i] = html.Div(
			html.Class("flex items-center space-x-2 mb-2"),
			html.Div(html.Class("w-2 h-2 rounded-full bg-primary")),
			html.Span(html.Class("text-sm"), g.Textf("List item number %d", i+1)),
		)
	}
	return items
}

func generateChatMessages() []g.Node {
	messages := []struct {
		sender  string
		message string
		isMe    bool
	}{
		{"John", "Hey, how's it going?", false},
		{"Me", "Pretty good! Just working on some components.", true},
		{"John", "Nice! What are you building?", false},
		{"Me", "A scroll area component for a chat interface.", true},
		{"John", "That sounds useful. Is it responsive?", false},
		{"Me", "Yes, it adapts to different screen sizes.", true},
		{"John", "Can you share the code when you're done?", false},
		{"Me", "Of course! I'll send it over once it's complete.", true},
		{"John", "Thanks! Looking forward to it.", false},
		{"Me", "No problem! Happy to help.", true},
	}
	
	nodes := make([]g.Node, len(messages))
	for i, msg := range messages {
		alignClass := "justify-start"
		bgClass := "bg-muted"
		if msg.isMe {
			alignClass = "justify-end"
			bgClass = "bg-primary text-primary-foreground"
		}
		
		nodes[i] = html.Div(
			html.Class(fmt.Sprintf("flex %s", alignClass)),
			html.Div(
				html.Class(fmt.Sprintf("rounded-lg px-3 py-2 max-w-[70%%] %s", bgClass)),
				html.P(html.Class("text-sm"), g.Text(msg.message)),
			),
		)
	}
	return nodes
}

func generateImagePlaceholders() []g.Node {
	images := make([]g.Node, 10)
	for i := 0; i < 10; i++ {
		images[i] = html.Div(
			html.Class("relative aspect-video w-64 shrink-0 overflow-hidden rounded-lg bg-muted"),
			html.Div(
				html.Class("flex h-full items-center justify-center"),
				html.Span(html.Class("text-lg font-semibold"), g.Textf("Image %d", i+1)),
			),
		)
	}
	return images
}

func generateTableRows() []g.Node {
	data := []struct {
		name       string
		status     string
		email      string
		department string
		role       string
		startDate  string
		salary     string
	}{
		{"Alice Johnson", "Active", "alice@example.com", "Engineering", "Senior Developer", "2020-03-15", "$120,000"},
		{"Bob Smith", "Active", "bob@example.com", "Design", "UI/UX Designer", "2021-06-01", "$95,000"},
		{"Carol Williams", "On Leave", "carol@example.com", "Marketing", "Marketing Manager", "2019-11-20", "$105,000"},
		{"David Brown", "Active", "david@example.com", "Sales", "Sales Representative", "2022-01-10", "$75,000"},
		{"Eve Davis", "Active", "eve@example.com", "Engineering", "DevOps Engineer", "2020-08-05", "$115,000"},
		{"Frank Miller", "Active", "frank@example.com", "HR", "HR Specialist", "2021-04-12", "$80,000"},
		{"Grace Wilson", "Active", "grace@example.com", "Finance", "Financial Analyst", "2020-12-01", "$90,000"},
		{"Henry Moore", "Active", "henry@example.com", "Engineering", "Frontend Developer", "2022-03-28", "$100,000"},
		{"Iris Taylor", "Active", "iris@example.com", "Product", "Product Manager", "2019-09-15", "$130,000"},
		{"Jack Anderson", "Active", "jack@example.com", "Engineering", "Backend Developer", "2021-07-20", "$110,000"},
	}
	
	rows := make([]g.Node, len(data))
	for i, row := range data {
		statusClass := "text-green-600"
		if row.status == "On Leave" {
			statusClass = "text-yellow-600"
		}
		
		rows[i] = html.Tr(
			html.Td(html.Class("font-medium sticky left-0 bg-background"), g.Text(row.name)),
			html.Td(html.Span(html.Class(statusClass), g.Text(row.status))),
			html.Td(g.Text(row.email)),
			html.Td(g.Text(row.department)),
			html.Td(g.Text(row.role)),
			html.Td(g.Text(row.startDate)),
			html.Td(g.Text(row.salary)),
		)
	}
	return rows
}

func generateColoredItems() []g.Node {
	colors := []struct {
		name  string
		class string
	}{
		{"Primary", "bg-primary"},
		{"Secondary", "bg-secondary"},
		{"Destructive", "bg-destructive"},
		{"Muted", "bg-muted"},
		{"Accent", "bg-accent"},
		{"Card", "bg-card"},
		{"Popover", "bg-popover"},
	}
	
	items := make([]g.Node, 20)
	for i := 0; i < 20; i++ {
		color := colors[i%len(colors)]
		items[i] = html.Div(
			html.Class(fmt.Sprintf("rounded-md p-3 %s", color.class)),
			html.Span(html.Class("text-sm font-medium"), g.Textf("%s Item %d", color.name, i+1)),
		)
	}
	return items
}