package hovercard

import (
	"fmt"
	"net/http"
	"time"

	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	hx "maragu.dev/gomponents-htmx"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// HTMXProps defines HTMX-specific properties for the HoverCard
type HTMXProps struct {
	ID          string // Unique ID for the hover card
	ContentPath string // Server path for loading content
	Delay       int    // Delay in ms before showing (default 200)
}

// NewHTMX creates an HTMX-enhanced HoverCard component
func NewHTMX(props Props, htmxProps HTMXProps, children ...g.Node) g.Node {
	classes := lib.CN("relative inline-block", props.Class)
	
	delay := htmxProps.Delay
	if delay == 0 {
		delay = 200
	}
	
	return html.Div(
		html.ID(htmxProps.ID),
		html.Class(classes),
		g.Attr("data-hover-card", "root"),
		hx.On("mouseenter", fmt.Sprintf(`
			clearTimeout(window.hoverCardTimeout_%s);
			window.hoverCardTimeout_%s = setTimeout(() => {
				htmx.ajax('GET', '%s', {
					target: '#%s-content',
					swap: 'innerHTML'
				}).then(() => {
					const content = document.querySelector('#%s-content');
					if (content) {
						content.style.display = 'block';
						// Position the content
						const trigger = event.currentTarget.querySelector('[data-hover-card="trigger"]');
						const rect = trigger.getBoundingClientRect();
						content.style.left = rect.left + 'px';
						content.style.top = (rect.bottom + 4) + 'px';
					}
				});
			}, %d);
		`, htmxProps.ID, htmxProps.ID, htmxProps.ContentPath, htmxProps.ID, htmxProps.ID, delay)),
		hx.On("mouseleave", fmt.Sprintf(`
			clearTimeout(window.hoverCardTimeout_%s);
			setTimeout(() => {
				const content = document.querySelector('#%s-content');
				if (content && !content.matches(':hover')) {
					content.style.display = 'none';
					content.innerHTML = '';
				}
			}, 100);
		`, htmxProps.ID, htmxProps.ID)),
		g.Group(children),
		// Content container
		html.Div(
			html.ID(htmxProps.ID+"-content"),
			html.Class("fixed z-50 w-64 rounded-md border bg-popover p-4 text-popover-foreground shadow-md outline-none"),
			html.Style("display: none;"),
			hx.On("mouseleave", fmt.Sprintf(`
				setTimeout(() => {
					if (!document.querySelector('#%s').matches(':hover')) {
						this.style.display = 'none';
						this.innerHTML = '';
					}
				}, 100);
			`, htmxProps.ID)),
		),
	)
}

// TriggerHTMX creates an HTMX-enhanced trigger
func TriggerHTMX(props TriggerProps, children ...g.Node) g.Node {
	classes := lib.CN("cursor-pointer", props.Class)
	
	return html.Div(
		html.Class(classes),
		g.Attr("data-hover-card", "trigger"),
		g.Group(children),
	)
}

// ExampleHTMX creates an HTMX hover card example
func ExampleHTMX() g.Node {
	htmxProps := HTMXProps{
		ID:          "user-hovercard",
		ContentPath: "/api/hovercard/user/nextjs",
		Delay:       300,
	}
	
	return NewHTMX(
		Props{},
		htmxProps,
		TriggerHTMX(
			TriggerProps{},
			html.A(
				html.Href("#"),
				html.Class("text-sm font-medium underline-offset-4 hover:underline"),
				g.Text("@nextjs"),
			),
		),
	)
}

// ProfileCardHTMX creates an HTMX profile hover card
func ProfileCardHTMX(username string) g.Node {
	htmxProps := HTMXProps{
		ID:          fmt.Sprintf("profile-%s", username),
		ContentPath: fmt.Sprintf("/api/hovercard/profile/%s", username),
	}
	
	return NewHTMX(
		Props{},
		htmxProps,
		TriggerHTMX(
			TriggerProps{},
			html.A(
				html.Href(fmt.Sprintf("/users/%s", username)),
				html.Class("font-medium hover:underline"),
				g.Text("@"+username),
			),
		),
	)
}

// LinkPreviewHTMX creates an HTMX link preview hover card
func LinkPreviewHTMX(url string) g.Node {
	htmxProps := HTMXProps{
		ID:          fmt.Sprintf("link-%d", time.Now().UnixNano()),
		ContentPath: fmt.Sprintf("/api/hovercard/link?url=%s", url),
		Delay:       500,
	}
	
	return NewHTMX(
		Props{},
		htmxProps,
		TriggerHTMX(
			TriggerProps{},
			html.A(
				html.Href(url),
				html.Class("text-blue-600 hover:underline"),
				html.Target("_blank"),
				html.Rel("noopener noreferrer"),
				g.Text(url),
			),
		),
	)
}

// RenderUserContent renders user hover card content (for server response)
func RenderUserContent(username string, name string, bio string, followers int, following int) g.Node {
	return html.Div(html.Class("space-y-2"),
		html.Div(html.Class("flex items-center space-x-4"),
			html.Div(html.Class("h-12 w-12 rounded-full bg-gradient-to-br from-blue-500 to-purple-600")),
			html.Div(html.Class("space-y-1"),
				html.H4(html.Class("text-sm font-semibold"), g.Text(name)),
				html.P(html.Class("text-sm text-muted-foreground"), g.Text("@"+username)),
			),
		),
		html.P(html.Class("text-sm"), g.Text(bio)),
		html.Div(html.Class("flex gap-4 text-xs text-muted-foreground"),
			html.Div(
				html.Span(html.Class("font-semibold text-foreground"), g.Textf("%d", following)),
				g.Text(" Following"),
			),
			html.Div(
				html.Span(html.Class("font-semibold text-foreground"), 
					g.If(followers >= 1000, g.Textf("%.1fk", float64(followers)/1000)),
					g.If(followers < 1000, g.Textf("%d", followers)),
				),
				g.Text(" Followers"),
			),
		),
		html.Button(
			html.Type("button"),
			html.Class("w-full text-sm bg-primary text-primary-foreground hover:bg-primary/90"),
			g.Text("Follow"),
		),
	)
}

// RenderLinkContent renders link preview content (for server response)
func RenderLinkContent(url string, title string, description string, imageUrl string) g.Node {
	return html.Div(html.Class("space-y-2 w-80"),
		g.If(imageUrl != "",
	html.Img(
				html.Src(imageUrl),
				html.Alt(title),
				html.Class("h-40 w-full rounded-md object-cover"),
			),
		),
		html.Div(html.Class("space-y-1"),
			html.H4(html.Class("text-sm font-semibold line-clamp-1"), g.Text(title)),
			html.P(html.Class("text-sm text-muted-foreground line-clamp-2"), 
				g.Text(description),
			),
			html.P(html.Class("text-xs text-muted-foreground truncate"), g.Text(url)),
		),
	)
}

// HoverCardHandlers creates HTTP handlers for hover card components
func HoverCardHandlers(mux *http.ServeMux) {
	// User hover card handler
	mux.HandleFunc("/api/hovercard/user/", func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Path[len("/api/hovercard/user/"):]
		
		// Mock user data
		users := map[string]struct {
			Name      string
			Bio       string
			Followers int
			Following int
		}{
			"nextjs": {
				Name:      "Next.js",
				Bio:       "The React Framework – created and maintained by @vercel.",
				Followers: 155300,
				Following: 256,
			},
			"vercel": {
				Name:      "Vercel",
				Bio:       "Develop. Preview. Ship. Creators of Next.js.",
				Followers: 88900,
				Following: 142,
			},
			"shadcn": {
				Name:      "shadcn",
				Bio:       "Design engineer. Building UI things.",
				Followers: 42100,
				Following: 89,
			},
		}
		
		if user, ok := users[username]; ok {
			node := RenderUserContent(username, user.Name, user.Bio, user.Followers, user.Following)
			node.Render(w)
		} else {
			// Unknown user
			node := RenderUserContent(username, username, "User not found", 0, 0)
			node.Render(w)
		}
	})
	
	// Profile hover card handler
	mux.HandleFunc("/api/hovercard/profile/", func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Path[len("/api/hovercard/profile/"):]
		
		// Mock profile data
		profiles := map[string]struct {
			Name      string
			Bio       string
			Followers int
			Following int
		}{
			"john": {
				Name:      "John Doe",
				Bio:       "Software Engineer | Open Source Enthusiast | Coffee Lover ☕",
				Followers: 1234,
				Following: 567,
			},
			"jane": {
				Name:      "Jane Smith",
				Bio:       "Product Designer | UI/UX | Making the web beautiful ✨",
				Followers: 2345,
				Following: 432,
			},
		}
		
		if profile, ok := profiles[username]; ok {
			node := RenderUserContent(username, profile.Name, profile.Bio, profile.Followers, profile.Following)
			node.Render(w)
		} else {
			node := RenderUserContent(username, username, "Profile not found", 0, 0)
			node.Render(w)
		}
	})
	
	// Link preview handler
	mux.HandleFunc("/api/hovercard/link", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")
		
		// Mock link previews
		previews := map[string]struct {
			Title       string
			Description string
			ImageUrl    string
		}{
			"https://nextjs.org": {
				Title:       "Next.js by Vercel - The React Framework",
				Description: "Next.js gives you the best developer experience with all the features you need for production.",
				ImageUrl:    "/api/placeholder/640/320",
			},
			"https://tailwindcss.com": {
				Title:       "Tailwind CSS - Rapidly build modern websites",
				Description: "Tailwind CSS is a utility-first CSS framework for rapidly building modern websites without leaving your HTML.",
				ImageUrl:    "/api/placeholder/640/320",
			},
		}
		
		if preview, ok := previews[url]; ok {
			node := RenderLinkContent(url, preview.Title, preview.Description, preview.ImageUrl)
			node.Render(w)
		} else {
			// Generic preview for unknown URLs
			node := RenderLinkContent(url, url, "Preview not available", "")
			node.Render(w)
		}
	})
	
	// Placeholder image handler
	mux.HandleFunc("/api/placeholder/", func(w http.ResponseWriter, r *http.Request) {
		// Return a simple SVG placeholder
		w.Header().Set("Content-Type", "image/svg+xml")
		fmt.Fprintf(w, `<svg width="640" height="320" xmlns="http://www.w3.org/2000/svg">
			<rect width="640" height="320" fill="#e5e7eb"/>
			<text x="320" y="160" text-anchor="middle" fill="#9ca3af" font-family="sans-serif" font-size="20">Image Placeholder</text>
		</svg>`)
	})
}