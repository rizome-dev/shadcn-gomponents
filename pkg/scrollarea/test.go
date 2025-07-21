package scrollarea

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

func Test() g.Node {
	return html.Div(
		html.Class("test"),
		html.Div(
			html.H3(g.Text("Section 1")),
		),
		html.Div(
			html.H3(g.Text("Section 2")),
		),
	)
}