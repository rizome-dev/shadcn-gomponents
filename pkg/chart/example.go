package chart

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ExampleLineChart demonstrates a basic line chart
func ExampleLineChart() g.Node {
	data := ChartData{
		Labels: []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun"},
		Series: []SeriesData{
			{
				Name:  "Revenue",
				Data:  []float64{12, 19, 15, 25, 22, 30},
				Color: "#3b82f6",
			},
			{
				Name:  "Expenses",
				Data:  []float64{8, 12, 10, 18, 15, 20},
				Color: "#ef4444",
			},
		},
	}

	return h.Div(
		h.Class("w-full max-w-4xl mx-auto p-4"),
		h.H2(h.Class("text-2xl font-bold mb-4"), g.Text("Line Chart Example")),
		New("line-chart-example", data,
			WithType(ChartTypeLine),
			WithTitle("Monthly Revenue vs Expenses"),
			WithSubtitle("Financial overview for the first half of the year"),
			WithHeight("400px"),
		),
	)
}

// ExampleBarChart demonstrates a bar chart
func ExampleBarChart() g.Node {
	data := ChartData{
		Labels: []string{"Q1", "Q2", "Q3", "Q4"},
		Series: []SeriesData{
			{
				Name:  "Product A",
				Data:  []float64{45, 52, 38, 65},
				Color: "#10b981",
			},
			{
				Name:  "Product B",
				Data:  []float64{28, 34, 42, 40},
				Color: "#f59e0b",
			},
			{
				Name:  "Product C",
				Data:  []float64{15, 20, 25, 30},
				Color: "#8b5cf6",
			},
		},
	}

	return h.Div(
		h.Class("w-full max-w-4xl mx-auto p-4"),
		h.H2(h.Class("text-2xl font-bold mb-4"), g.Text("Bar Chart Example")),
		New("bar-chart-example", data,
			WithType(ChartTypeBar),
			WithTitle("Quarterly Sales by Product"),
			WithHeight("350px"),
		),
	)
}

// ExampleAreaChart demonstrates an area chart
func ExampleAreaChart() g.Node {
	data := ChartData{
		Labels: []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"},
		Series: []SeriesData{
			{
				Name:  "Page Views",
				Data:  []float64{120, 132, 101, 134, 90, 230, 210},
				Color: "#3b82f6",
			},
		},
	}

	return h.Div(
		h.Class("w-full max-w-4xl mx-auto p-4"),
		h.H2(h.Class("text-2xl font-bold mb-4"), g.Text("Area Chart Example")),
		New("area-chart-example", data,
			WithType(ChartTypeArea),
			WithTitle("Weekly Page Views"),
			WithSubtitle("Website traffic for the current week"),
			WithHeight("300px"),
			WithClass("bg-gray-50 rounded-lg p-4"),
		),
	)
}

// ExamplePieChart demonstrates a pie chart
func ExamplePieChart() g.Node {
	data := ChartData{
		Labels: []string{"Desktop", "Mobile", "Tablet", "Smart TV"},
		Series: []SeriesData{
			{
				Name: "Device Usage",
				Data: []float64{45, 35, 15, 5},
			},
		},
	}

	return h.Div(
		h.Class("w-full max-w-2xl mx-auto p-4"),
		h.H2(h.Class("text-2xl font-bold mb-4"), g.Text("Pie Chart Example")),
		New("pie-chart-example", data,
			WithType(ChartTypePie),
			WithTitle("Device Usage Distribution"),
			WithHeight("400px"),
		),
	)
}

// ExampleDonutChart demonstrates a donut chart
func ExampleDonutChart() g.Node {
	data := ChartData{
		Labels: []string{"Completed", "In Progress", "Pending", "Cancelled"},
		Series: []SeriesData{
			{
				Name: "Task Status",
				Data: []float64{120, 45, 30, 15},
			},
		},
	}

	return h.Div(
		h.Class("w-full max-w-2xl mx-auto p-4"),
		h.H2(h.Class("text-2xl font-bold mb-4"), g.Text("Donut Chart Example")),
		New("donut-chart-example", data,
			WithType(ChartTypeDonut),
			WithTitle("Task Status Overview"),
			WithHeight("400px"),
		),
	)
}

// ExampleHTMXChart demonstrates an HTMX-enabled chart
func ExampleHTMXChart() g.Node {
	data := ChartData{
		Labels: []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun"},
		Series: []SeriesData{
			{
				Name:  "Sales",
				Data:  []float64{65, 59, 80, 81, 56, 95},
				Color: "#10b981",
			},
			{
				Name:  "Profit",
				Data:  []float64{28, 48, 40, 52, 36, 62},
				Color: "#3b82f6",
			},
		},
	}

	return h.Div(
		h.Class("w-full max-w-4xl mx-auto p-4"),
		h.H2(h.Class("text-2xl font-bold mb-4"), g.Text("HTMX Chart with Live Updates")),
		h.P(h.Class("text-gray-600 mb-4"), g.Text("This chart refreshes every 5 seconds with new data from the server.")),
		NewHTMX("htmx-chart-example", data,
			WithHTMXType(ChartTypeBar),
			WithHTMXTitle("Real-time Sales Dashboard"),
			WithHTMXSubtitle("Live data updates via HTMX"),
			WithHTMXPolling(5),
			WithHTMXHeight("400px"),
		),
	)
}

// ExampleHTMXInteractiveChart demonstrates an interactive HTMX chart
func ExampleHTMXInteractiveChart() g.Node {
	data := ChartData{
		Labels: []string{"Week 1", "Week 2", "Week 3", "Week 4"},
		Series: []SeriesData{
			{
				Name:  "Active Users",
				Data:  []float64{1250, 1380, 1520, 1690},
				Color: "#8b5cf6",
			},
			{
				Name:  "New Signups",
				Data:  []float64{180, 220, 195, 250},
				Color: "#ec4899",
			},
		},
	}

	return h.Div(
		h.Class("w-full max-w-4xl mx-auto p-4"),
		h.H2(h.Class("text-2xl font-bold mb-4"), g.Text("Interactive HTMX Chart")),
		h.P(h.Class("text-gray-600 mb-4"), g.Text("Click the refresh button to update the chart with new data.")),
		NewHTMX("htmx-interactive-chart", data,
			WithHTMXType(ChartTypeLine),
			WithHTMXTitle("User Growth Metrics"),
			WithHTMXHeight("350px"),
		),
	)
}

// ExampleMinimalChart demonstrates a minimal chart without extras
func ExampleMinimalChart() g.Node {
	data := ChartData{
		Labels: []string{"A", "B", "C", "D", "E"},
		Series: []SeriesData{
			{
				Name: "Values",
				Data: []float64{10, 25, 15, 30, 20},
			},
		},
	}

	return h.Div(
		h.Class("w-full max-w-2xl mx-auto p-4"),
		h.H2(h.Class("text-2xl font-bold mb-4"), g.Text("Minimal Chart")),
		New("minimal-chart", data,
			WithType(ChartTypeBar),
			WithHeight("200px"),
			WithoutLegend(),
			WithoutGrid(),
			WithoutAnimations(),
		),
	)
}

// ExampleDarkThemeChart demonstrates a dark theme chart
func ExampleDarkThemeChart() g.Node {
	data := ChartData{
		Labels: []string{"Mon", "Tue", "Wed", "Thu", "Fri"},
		Series: []SeriesData{
			{
				Name:  "Temperature (°C)",
				Data:  []float64{22, 24, 21, 26, 23},
				Color: "#fbbf24",
			},
		},
	}

	return h.Div(
		h.Class("w-full max-w-4xl mx-auto p-4 bg-gray-900 rounded-lg"),
		h.H2(h.Class("text-2xl font-bold mb-4 text-white"), g.Text("Dark Theme Chart")),
		New("dark-theme-chart", data,
			WithType(ChartTypeLine),
			WithTitle("Weekly Temperature"),
			WithHeight("300px"),
			WithTheme("dark"),
			WithClass("text-white"),
		),
	)
}

// ExampleMultipleCharts demonstrates a dashboard with multiple charts
func ExampleMultipleCharts() g.Node {
	salesData := ChartData{
		Labels: []string{"Jan", "Feb", "Mar", "Apr"},
		Series: []SeriesData{
			{
				Name:  "2023",
				Data:  []float64{30, 40, 35, 50},
				Color: "#3b82f6",
			},
			{
				Name:  "2024",
				Data:  []float64{35, 45, 40, 60},
				Color: "#10b981",
			},
		},
	}

	categoryData := ChartData{
		Labels: []string{"Electronics", "Clothing", "Food", "Books", "Other"},
		Series: []SeriesData{
			{
				Name: "Sales by Category",
				Data: []float64{30, 25, 20, 15, 10},
			},
		},
	}

	return h.Div(
		h.Class("w-full max-w-6xl mx-auto p-4"),
		h.H2(h.Class("text-2xl font-bold mb-6"), g.Text("Sales Dashboard")),
		
		h.Div(
			h.Class("grid grid-cols-1 md:grid-cols-2 gap-6"),
			
			// Sales comparison chart
			h.Div(
				h.Class("bg-white rounded-lg shadow-lg p-4"),
				New("sales-comparison", salesData,
					WithType(ChartTypeBar),
					WithTitle("Sales Comparison"),
					WithSubtitle("Year over year"),
					WithHeight("300px"),
				),
			),
			
			// Category distribution chart
			h.Div(
				h.Class("bg-white rounded-lg shadow-lg p-4"),
				New("category-distribution", categoryData,
					WithType(ChartTypePie),
					WithTitle("Sales by Category"),
					WithHeight("300px"),
				),
			),
		),
	)
}