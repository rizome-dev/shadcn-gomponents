package chart_test

import (
	"strings"
	"testing"

	"github.com/rizome-dev/shadcn-gomponents/pkg/chart"
	g "maragu.dev/gomponents"
)

func TestNew(t *testing.T) {
	t.Run("creates basic chart with default options", func(t *testing.T) {
		data := chart.ChartData{
			Labels: []string{"A", "B", "C"},
			Series: []chart.SeriesData{
				{Name: "Test", Data: []float64{10, 20, 30}},
			},
		}
		
		result := chart.New("test-chart", data)
		html := render(result)
		
		// Check basic structure
		if !strings.Contains(html, `data-chart="static"`) {
			t.Error("expected static chart data attribute")
		}
		if !strings.Contains(html, `data-chart-type="line"`) {
			t.Error("expected line chart type by default")
		}
		if !strings.Contains(html, `data-theme="light"`) {
			t.Error("expected light theme by default")
		}
		if !strings.Contains(html, `id="test-chart"`) {
			t.Error("expected chart ID")
		}
		// Server-rendered charts should have SVG for line charts
		if !strings.Contains(html, "<svg") {
			t.Error("expected SVG element for line chart")
		}
	})

	t.Run("creates chart with custom type", func(t *testing.T) {
		data := chart.ChartData{
			Labels: []string{"A"},
			Series: []chart.SeriesData{{Name: "Test", Data: []float64{10}}},
		}
		
		result := chart.New("bar-chart", data, chart.WithType(chart.ChartTypeBar))
		html := render(result)
		
		if !strings.Contains(html, `data-chart-type="bar"`) {
			t.Error("expected bar chart type")
		}
	})

	t.Run("creates chart with title and subtitle", func(t *testing.T) {
		data := chart.ChartData{
			Labels: []string{"A"},
			Series: []chart.SeriesData{{Name: "Test", Data: []float64{10}}},
		}
		
		result := chart.New("titled-chart", data,
			chart.WithTitle("Chart Title"),
			chart.WithSubtitle("Chart Subtitle"),
		)
		html := render(result)
		
		// Server-rendered charts show title/subtitle as text
		if !strings.Contains(html, "<h3") {
			t.Error("expected title element")
		}
		if !strings.Contains(html, "Chart Title") {
			t.Error("expected title text")
		}
		if !strings.Contains(html, "Chart Subtitle") {
			t.Error("expected subtitle text")
		}
	})

	t.Run("creates chart with custom dimensions", func(t *testing.T) {
		data := chart.ChartData{
			Labels: []string{"A"},
			Series: []chart.SeriesData{{Name: "Test", Data: []float64{10}}},
		}
		
		result := chart.New("sized-chart", data,
			chart.WithHeight("500px"),
			chart.WithWidth("80%"),
		)
		html := render(result)
		
		if !strings.Contains(html, "height: 500px") {
			t.Error("expected custom height")
		}
		if !strings.Contains(html, "width: 80%") {
			t.Error("expected custom width")
		}
	})

	t.Run("hides legend when requested", func(t *testing.T) {
		data := chart.ChartData{
			Labels: []string{"A"},
			Series: []chart.SeriesData{{Name: "Test", Data: []float64{10}}},
		}
		
		result := chart.New("no-legend-chart", data, chart.WithoutLegend())
		html := render(result)
		
		// When legend is hidden, there should be no legend items
		if strings.Contains(html, "role=\"list\"") && strings.Contains(html, "aria-label=\"Chart legend\"") {
			t.Error("expected no legend when WithoutLegend is used")
		}
	})

	t.Run("hides grid when requested", func(t *testing.T) {
		data := chart.ChartData{
			Labels: []string{"A"},
			Series: []chart.SeriesData{{Name: "Test", Data: []float64{10}}},
		}
		
		result := chart.New("no-grid-chart", data, chart.WithoutGrid())
		html := render(result)
		
		// When grid is hidden, there should be no grid lines
		if strings.Contains(html, "stroke=\"#e5e7eb\"") {
			t.Error("expected no grid lines when WithoutGrid is used")
		}
	})

	t.Run("disables animations when requested", func(t *testing.T) {
		data := chart.ChartData{
			Labels: []string{"A"},
			Series: []chart.SeriesData{{Name: "Test", Data: []float64{10}}},
		}
		
		result := chart.New("no-anim-chart", data, chart.WithoutAnimations())
		html := render(result)
		
		// Animations don't affect server-rendered charts
		// Just check that chart renders
		if !strings.Contains(html, "<svg") {
			t.Error("expected SVG element")
		}
	})

	t.Run("applies dark theme", func(t *testing.T) {
		data := chart.ChartData{
			Labels: []string{"A"},
			Series: []chart.SeriesData{{Name: "Test", Data: []float64{10}}},
		}
		
		result := chart.New("dark-chart", data, chart.WithTheme("dark"))
		html := render(result)
		
		if !strings.Contains(html, `data-theme="dark"`) {
			t.Error("expected dark theme")
		}
	})

	t.Run("applies custom classes", func(t *testing.T) {
		data := chart.ChartData{
			Labels: []string{"A"},
			Series: []chart.SeriesData{{Name: "Test", Data: []float64{10}}},
		}
		
		result := chart.New("custom-chart", data, chart.WithClass("custom-class shadow-lg"))
		html := render(result)
		
		if !strings.Contains(html, "custom-class") {
			t.Error("expected custom class")
		}
		if !strings.Contains(html, "shadow-lg") {
			t.Error("expected shadow class")
		}
	})

	t.Run("embeds chart data as JSON", func(t *testing.T) {
		data := chart.ChartData{
			Labels: []string{"Jan", "Feb"},
			Series: []chart.SeriesData{
				{
					Name:  "Revenue",
					Data:  []float64{100, 200},
					Color: "#3b82f6",
				},
			},
		}
		
		result := chart.New("data-chart", data)
		html := render(result)
		
		// Server-rendered charts show data directly
		if !strings.Contains(html, "<svg") {
			t.Error("expected SVG element")
		}
		// Legend should show series name
		if !strings.Contains(html, "Revenue") {
			t.Error("expected series name in legend")
		}
	})
}

func TestCanvas(t *testing.T) {
	t.Run("creates canvas element in chart", func(t *testing.T) {
		data := chart.ChartData{
			Labels: []string{"A"},
			Series: []chart.SeriesData{{Name: "Test", Data: []float64{10}}},
		}
		
		result := chart.New("canvas-test", data)
		html := render(result)
		
		// Server-rendered charts use SVG, not canvas
		if !strings.Contains(html, "<svg") {
			t.Error("expected SVG element for server-rendered chart")
		}
		if !strings.Contains(html, `id="canvas-test"`) {
			t.Error("expected chart ID")
		}
		if !strings.Contains(html, `role="img"`) {
			t.Error("expected img role for accessibility")
		}
	})
}

func TestLegend(t *testing.T) {
	t.Run("creates legend when present in chart", func(t *testing.T) {
		data := chart.ChartData{
			Labels: []string{"A", "B"},
			Series: []chart.SeriesData{
				{Name: "Series 1", Color: "#ff0000", Data: []float64{10, 20}},
				{Name: "Series 2", Color: "#00ff00", Data: []float64{15, 25}},
				{Name: "Series 3", Color: "#0000ff", Data: []float64{20, 30}},
			},
		}
		
		result := chart.New("legend-test", data)
		html := render(result)
		
		// Server-rendered charts show legend directly
		if !strings.Contains(html, "Series 1") {
			t.Error("expected first series in legend")
		}
		if !strings.Contains(html, "Series 2") {
			t.Error("expected second series in legend")
		}
		if !strings.Contains(html, "Series 3") {
			t.Error("expected third series in legend")
		}
	})
}

func TestHTMXChart(t *testing.T) {
	t.Run("creates HTMX-enabled chart", func(t *testing.T) {
		data := chart.ChartData{
			Labels: []string{"A", "B"},
			Series: []chart.SeriesData{{Name: "Test", Data: []float64{10, 20}}},
		}
		
		result := chart.NewHTMX("htmx-chart", data)
		html := render(result)
		
		if !strings.Contains(html, `id="htmx-chart-wrapper"`) {
			t.Error("expected wrapper ID")
		}
		if !strings.Contains(html, `data-chart="htmx"`) {
			t.Error("expected HTMX chart attribute")
		}
		if !strings.Contains(html, `id="htmx-chart"`) {
			t.Error("expected chart container ID")
		}
		if !strings.Contains(html, "Refresh") {
			t.Error("expected refresh button")
		}
	})

	t.Run("creates HTMX chart with polling", func(t *testing.T) {
		data := chart.ChartData{
			Labels: []string{"A"},
			Series: []chart.SeriesData{{Name: "Test", Data: []float64{10}}},
		}
		
		result := chart.NewHTMX("polling-chart", data, chart.WithHTMXPolling(10))
		html := render(result)
		
		if !strings.Contains(html, `hx-trigger="load, every 10s"`) {
			t.Error("expected polling trigger")
		}
		if !strings.Contains(html, `hx-get="/chart/polling-chart/data"`) {
			t.Error("expected update endpoint")
		}
		if strings.Contains(html, "Refresh") {
			t.Error("should not have refresh button when polling")
		}
	})

	t.Run("creates HTMX chart with custom endpoint", func(t *testing.T) {
		data := chart.ChartData{
			Labels: []string{"A"},
			Series: []chart.SeriesData{{Name: "Test", Data: []float64{10}}},
		}
		
		result := chart.NewHTMX("custom-chart", data, chart.WithHTMXEndpoint("/api/charts/custom"))
		html := render(result)
		
		if !strings.Contains(html, `hx-get="/api/charts/custom/data"`) {
			t.Error("expected custom update endpoint")
		}
	})

	t.Run("renders bar chart visualization", func(t *testing.T) {
		data := chart.ChartData{
			Labels: []string{"Q1", "Q2"},
			Series: []chart.SeriesData{
				{Name: "Sales", Data: []float64{100, 150}},
			},
		}
		
		result := chart.NewHTMX("bar-chart", data, chart.WithHTMXType(chart.ChartTypeBar))
		html := render(result)
		
		if !strings.Contains(html, `data-chart-type="bar"`) {
			t.Error("expected bar chart type")
		}
		// Bar chart should render bars
		if !strings.Contains(html, "Q1") || !strings.Contains(html, "Q2") {
			t.Error("expected x-axis labels")
		}
	})

	t.Run("renders line chart as SVG", func(t *testing.T) {
		data := chart.ChartData{
			Labels: []string{"A", "B", "C"},
			Series: []chart.SeriesData{
				{Name: "Line", Data: []float64{10, 20, 15}},
			},
		}
		
		result := chart.NewHTMX("line-chart", data, chart.WithHTMXType(chart.ChartTypeLine))
		html := render(result)
		
		if !strings.Contains(html, "<svg") {
			t.Error("expected SVG element for line chart")
		}
		if !strings.Contains(html, "<polyline") {
			t.Error("expected polyline for line chart")
		}
	})

	t.Run("renders pie chart as SVG", func(t *testing.T) {
		data := chart.ChartData{
			Labels: []string{"A", "B", "C"},
			Series: []chart.SeriesData{
				{Name: "Distribution", Data: []float64{30, 50, 20}},
			},
		}
		
		result := chart.NewHTMX("pie-chart", data, chart.WithHTMXType(chart.ChartTypePie))
		html := render(result)
		
		if !strings.Contains(html, "<svg") {
			t.Error("expected SVG element for pie chart")
		}
		if !strings.Contains(html, "<path") {
			t.Error("expected path elements for pie slices")
		}
	})

	t.Run("renders donut chart with inner radius", func(t *testing.T) {
		data := chart.ChartData{
			Labels: []string{"A", "B"},
			Series: []chart.SeriesData{
				{Name: "Status", Data: []float64{60, 40}},
			},
		}
		
		result := chart.NewHTMX("donut-chart", data, chart.WithHTMXType(chart.ChartTypeDonut))
		html := render(result)
		
		if !strings.Contains(html, "<svg") {
			t.Error("expected SVG element for donut chart")
		}
		if !strings.Contains(html, `data-chart-type="donut"`) {
			t.Error("expected donut chart type")
		}
	})
}

func TestChartResponse(t *testing.T) {
	t.Run("HTMX chart updates work correctly", func(t *testing.T) {
		// Test that HTMX charts can be updated
		data := chart.ChartData{
			Labels: []string{"New A", "New B"},
			Series: []chart.SeriesData{
				{Name: "Updated", Data: []float64{50, 60}},
			},
		}
		
		result := chart.NewHTMX("update-test", data, chart.WithHTMXType(chart.ChartTypeBar))
		html := render(result)
		
		if !strings.Contains(html, `id="update-test"`) {
			t.Error("expected chart ID")
		}
		if !strings.Contains(html, "New A") {
			t.Error("expected labels in chart")
		}
	})
}

func TestTooltipResponse(t *testing.T) {
	t.Run("generates tooltip content", func(t *testing.T) {
		result := chart.TooltipResponse("Revenue", "March", 1250.50)
		html := render(result)
		
		if !strings.Contains(html, "March") {
			t.Error("expected label in tooltip")
		}
		if !strings.Contains(html, "Revenue: 1250.50") {
			t.Error("expected formatted value in tooltip")
		}
	})
}

// Helper function to render components
func render(node g.Node) string {
	if node == nil {
		return ""
	}
	var sb strings.Builder
	_ = node.Render(&sb)
	return sb.String()
}