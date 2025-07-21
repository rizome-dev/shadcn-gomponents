package calendar

import (
	"time"

	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
)

// Example creates a basic calendar example
func Example() g.Node {
	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Basic Calendar")),
		New(Props{
			Value: time.Now(),
			Month: time.Now(),
		}),
	)
}

// ExampleWithWeekNumbers creates a calendar with week numbers
func ExampleWithWeekNumbers() g.Node {
	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Calendar with Week Numbers")),
		New(Props{
			Value:     time.Now(),
			Month:     time.Now(),
			ShowWeeks: true,
		}),
	)
}

// ExampleDateRange creates a calendar showing a date range
func ExampleDateRange() g.Node {
	startDate := time.Now()
	endDate := startDate.AddDate(0, 0, 7) // 7 days later

	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Date Range Calendar")),
		DateRangeCalendar(startDate, endDate, time.Now()),
	)
}

// ExampleWithMinMax creates a calendar with min/max date constraints
func ExampleWithMinMax() g.Node {
	minDate := time.Now().AddDate(0, 0, -7)  // 7 days ago
	maxDate := time.Now().AddDate(0, 0, 14) // 14 days from now

	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Calendar with Date Constraints")),
		html.P(html.Class("text-sm text-muted-foreground mb-2"), g.Text("Only dates within ±7-14 days from today are selectable")),
		New(Props{
			Value:   time.Now(),
			Month:   time.Now(),
			MinDate: minDate,
			MaxDate: maxDate,
		}),
	)
}

// ExampleHTMX creates an HTMX-enhanced calendar example
func ExampleHTMX() g.Node {
	htmxProps := HTMXProps{
		ID:           "calendar-interactive",
		NavigatePath: "/api/calendar/navigate",
		SelectPath:   "/api/calendar/select",
		UpdatePath:   "/api/calendar/update",
	}

	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Interactive Calendar")),
		html.P(html.Class("text-sm text-muted-foreground mb-2"), g.Text("Click dates and navigation arrows to interact")),
		NewHTMX(Props{
			Value: time.Now(),
			Month: time.Now(),
		}, htmxProps),
	)
}

// ExampleDatePicker creates a date picker example
func ExampleDatePicker() g.Node {
	htmxProps := HTMXProps{
		ID:           "date-picker",
		NavigatePath: "/api/datepicker/navigate",
		SelectPath:   "/api/datepicker/select",
		UpdatePath:   "/api/datepicker/show",
	}

	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Date Picker")),
		html.P(html.Class("text-sm text-muted-foreground mb-2"), g.Text("Click the input or calendar icon to open the date picker")),
		DatePickerHTMX(htmxProps, "", "Select a date"),
	)
}

// ExampleMonthYearPicker creates a month/year picker example
func ExampleMonthYearPicker() g.Node {
	htmxProps := HTMXProps{
		ID:         "monthyear-picker",
		UpdatePath: "/api/monthyear/update",
	}

	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Month & Year Picker")),
		MonthYearPickerHTMX(htmxProps, time.Now()),
	)
}

// ExampleMultiMonth creates a multi-month calendar view
func ExampleMultiMonth() g.Node {
	currentMonth := time.Now()
	nextMonth := currentMonth.AddDate(0, 1, 0)

	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Multi-Month View")),
		html.Div(html.Class("grid grid-cols-1 md:grid-cols-2 gap-4"),
			New(Props{
				Value: time.Now(),
				Month: currentMonth,
			}),
			New(Props{
				Value: time.Now(),
				Month: nextMonth,
			}),
		),
	)
}

// ExampleCustomStyling creates a calendar with custom styling
func ExampleCustomStyling() g.Node {
	return html.Div(
		html.H3(html.Class("text-lg font-semibold mb-4"), g.Text("Custom Styled Calendar")),
		New(Props{
			Value: time.Now(),
			Month: time.Now(),
			Class: "bg-slate-900 text-slate-50 border-slate-800",
		}),
	)
}