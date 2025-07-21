package slider

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

// Examples demonstrates various slider configurations
func Examples() g.Node {
	return html.Div(
		html.Class("space-y-8 p-8"),
		
		// Basic Slider
		html.Div(
			html.Class("space-y-4"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Basic Slider")),
			html.Div(
				html.Class("space-y-4 max-w-md"),
				// Default slider
				html.Div(
					html.Class("space-y-2"),
					html.H4(html.Class("text-sm font-medium"), g.Text("Default")),
					New(Props{Value: []int{50}}),
				),
				
				// With custom range
				html.Div(
					html.Class("space-y-2"),
					html.H4(html.Class("text-sm font-medium"), g.Text("Custom Range (0-10)")),
					New(Props{
						Min:   0,
						Max:   10,
						Value: []int{5},
					}),
				),
				
				// With step
				html.Div(
					html.Class("space-y-2"),
					html.H4(html.Class("text-sm font-medium"), g.Text("With Step (10)")),
					New(Props{
						Min:   0,
						Max:   100,
						Step:  10,
						Value: []int{30},
					}),
				),
			),
		),

		// Range Slider
		html.Div(
			html.Class("space-y-4"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Range Slider")),
			html.Div(
				html.Class("space-y-4 max-w-md"),
				// Basic range
				html.Div(
					html.Class("space-y-2"),
					html.H4(html.Class("text-sm font-medium"), g.Text("Basic Range")),
					Range(Props{
						Value: []int{25, 75},
					}),
				),
				
				// Custom range
				html.Div(
					html.Class("space-y-2"),
					html.H4(html.Class("text-sm font-medium"), g.Text("Temperature Range (°C)")),
					Range(Props{
						Min:   -10,
						Max:   40,
						Step:  1,
						Value: []int{18, 24},
					}),
				),
			),
		),

		// Orientations
		html.Div(
			html.Class("space-y-4"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Orientations")),
			html.Div(
				html.Class("flex gap-8"),
				// Horizontal
				html.Div(
					html.Class("flex-1 space-y-2"),
					html.H4(html.Class("text-sm font-medium"), g.Text("Horizontal (Default)")),
					New(Props{
						Value: []int{60},
					}),
				),
				
				// Vertical
				html.Div(
					html.Class("space-y-2"),
					html.H4(html.Class("text-sm font-medium"), g.Text("Vertical")),
					html.Div(
						html.Class("h-48 flex justify-center"),
						Vertical(Props{
							Value: []int{60},
						}),
					),
				),
			),
		),

		// With Labels
		html.Div(
			html.Class("space-y-4"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("With Labels")),
			html.Div(
				html.Class("space-y-4 max-w-md"),
				// Min/Max labels
				html.Div(
					html.Class("space-y-2"),
					html.H4(html.Class("text-sm font-medium"), g.Text("Min/Max Labels")),
					WithLabels(Props{
						Min:   0,
						Max:   100,
						Value: []int{40},
					}),
				),
				
				// Custom labels
				html.Div(
					html.Class("space-y-2"),
					html.H4(html.Class("text-sm font-medium"), g.Text("Volume Control")),
					WithLabels(Props{
						Min:   0,
						Max:   100,
						Step:  5,
						Value: []int{65},
					}),
				),
			),
		),

		// With Value Display
		html.Div(
			html.Class("space-y-4"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("With Value Display")),
			html.Div(
				html.Class("space-y-4 max-w-md"),
				// Single value
				WithValue(Props{
					Value: []int{33},
					ID:    "value-single",
				}),
				
				// Range value
				WithValue(Props{
					Value: []int{20, 80},
					ID:    "value-range",
				}),
			),
		),

		// With Ticks
		html.Div(
			html.Class("space-y-4"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("With Tick Marks")),
			html.Div(
				html.Class("space-y-6 max-w-md"),
				// 5 ticks
				html.Div(
					html.Class("space-y-2 pb-4"),
					html.H4(html.Class("text-sm font-medium"), g.Text("5 Tick Marks")),
					WithTicks(Props{
						Value: []int{60},
					}, 5),
				),
				
				// 11 ticks (0-100 by 10s)
				html.Div(
					html.Class("space-y-2 pb-4"),
					html.H4(html.Class("text-sm font-medium"), g.Text("11 Tick Marks")),
					WithTicks(Props{
						Min:   0,
						Max:   100,
						Step:  10,
						Value: []int{70},
					}, 11),
				),
			),
		),

		// States
		html.Div(
			html.Class("space-y-4"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("States")),
			html.Div(
				html.Class("space-y-4 max-w-md"),
				// Disabled
				html.Div(
					html.Class("space-y-2"),
					html.H4(html.Class("text-sm font-medium"), g.Text("Disabled")),
					New(Props{
						Value:    []int{50},
						Disabled: true,
					}),
				),
				
				// Disabled range
				html.Div(
					html.Class("space-y-2"),
					html.H4(html.Class("text-sm font-medium"), g.Text("Disabled Range")),
					Range(Props{
						Value:    []int{30, 70},
						Disabled: true,
					}),
				),
			),
		),

		// Form Integration
		html.Div(
			html.Class("space-y-4"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Form Integration")),
			html.Form(
				html.Class("space-y-4 max-w-md"),
				html.Method("post"),
				html.Action("#"),
				
				// With name for form submission
				html.Div(
					html.Class("space-y-2"),
					html.Label(html.For("brightness"), g.Text("Brightness")),
					New(Props{
						ID:    "brightness",
						Name:  "brightness",
						Value: []int{75},
					}),
				),
				
				// Range with name
				html.Div(
					html.Class("space-y-2"),
					html.Label(g.Text("Price Range")),
					Range(Props{
						Name:  "price",
						Min:   0,
						Max:   1000,
						Step:  50,
						Value: []int{200, 600},
					}),
				),
				
				html.Button(
					html.Type("submit"),
					html.Class("bg-primary text-primary-foreground px-4 py-2 rounded"),
					g.Text("Submit"),
				),
			),
		),

		// Complex Examples
		html.Div(
			html.Class("space-y-4"),
			html.H3(html.Class("text-lg font-semibold"), g.Text("Complex Examples")),
			html.Div(
				html.Class("space-y-6 max-w-md"),
				// RGB Color Picker
				html.Div(
					html.Class("space-y-4"),
					html.H4(html.Class("text-sm font-medium"), g.Text("RGB Color Picker")),
					html.Div(
						html.Class("space-y-3"),
						// Red
						html.Div(
							html.Class("flex items-center gap-4"),
							html.Span(html.Class("w-8 text-sm"), g.Text("R")),
							html.Div(html.Class("flex-1"),
								New(Props{
									Max:   255,
									Value: []int{128},
									Class: "[&_[data-slider-range]]:bg-red-500",
								}),
							),
							html.Span(html.Class("w-12 text-sm text-right"), g.Text("128")),
						),
						// Green
						html.Div(
							html.Class("flex items-center gap-4"),
							html.Span(html.Class("w-8 text-sm"), g.Text("G")),
							html.Div(html.Class("flex-1"),
								New(Props{
									Max:   255,
									Value: []int{200},
									Class: "[&_[data-slider-range]]:bg-green-500",
								}),
							),
							html.Span(html.Class("w-12 text-sm text-right"), g.Text("200")),
						),
						// Blue
						html.Div(
							html.Class("flex items-center gap-4"),
							html.Span(html.Class("w-8 text-sm"), g.Text("B")),
							html.Div(html.Class("flex-1"),
								New(Props{
									Max:   255,
									Value: []int{64},
									Class: "[&_[data-slider-range]]:bg-blue-500",
								}),
							),
							html.Span(html.Class("w-12 text-sm text-right"), g.Text("64")),
						),
						// Color preview
						html.Div(
							html.Class("mt-4 h-16 rounded-md"),
							g.Attr("style", "background-color: rgb(128, 200, 64)"),
						),
					),
				),

				// Audio Mixer
				html.Div(
					html.Class("space-y-4"),
					html.H4(html.Class("text-sm font-medium"), g.Text("Audio Mixer")),
					html.Div(
						html.Class("flex gap-4"),
						// Channel sliders
						g.Group([]g.Node{
							html.Div(
								html.Class("text-center space-y-2"),
								html.Label(html.Class("text-xs"), g.Text("Master")),
								html.Div(html.Class("h-32 flex justify-center"),
									Vertical(Props{
										Max:   100,
										Value: []int{75},
									}),
								),
								html.Span(html.Class("text-xs text-muted-foreground"), g.Text("75%")),
							),
							html.Div(
								html.Class("text-center space-y-2"),
								html.Label(html.Class("text-xs"), g.Text("Bass")),
								html.Div(html.Class("h-32 flex justify-center"),
									Vertical(Props{
										Max:   100,
										Value: []int{60},
									}),
								),
								html.Span(html.Class("text-xs text-muted-foreground"), g.Text("60%")),
							),
							html.Div(
								html.Class("text-center space-y-2"),
								html.Label(html.Class("text-xs"), g.Text("Mid")),
								html.Div(html.Class("h-32 flex justify-center"),
									Vertical(Props{
										Max:   100,
										Value: []int{50},
									}),
								),
								html.Span(html.Class("text-xs text-muted-foreground"), g.Text("50%")),
							),
							html.Div(
								html.Class("text-center space-y-2"),
								html.Label(html.Class("text-xs"), g.Text("Treble")),
								html.Div(html.Class("h-32 flex justify-center"),
									Vertical(Props{
										Max:   100,
										Value: []int{65},
									}),
								),
								html.Span(html.Class("text-xs text-muted-foreground"), g.Text("65%")),
							),
						}),
					),
				),
			),
		),
	)
}