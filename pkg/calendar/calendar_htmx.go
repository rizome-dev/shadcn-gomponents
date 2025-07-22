package calendar

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	hx "maragu.dev/gomponents-htmx"
	"github.com/rizome-dev/shadcn-gomponents/lib"
	"github.com/rizome-dev/shadcn-gomponents/lib/icons"
)

// HTMXProps defines HTMX-specific properties for the Calendar
type HTMXProps struct {
	ID           string // Unique ID for the calendar
	NavigatePath string // Server path for month navigation
	SelectPath   string // Server path for date selection
	UpdatePath   string // Server path for calendar updates
}

// NewHTMX creates an HTMX-enhanced Calendar component
func NewHTMX(props Props, htmxProps HTMXProps, children ...g.Node) g.Node {
	// Set defaults
	if props.Month.IsZero() {
		props.Month = time.Now()
	}
	if !props.ShowDays {
		props.ShowDays = true
	}

	classes := lib.CN(
		"bg-background p-3 rounded-lg border",
		props.Class,
	)

	return html.Div(
		html.ID(htmxProps.ID),
		html.Class(classes),
		CalendarHeaderHTMX(HeaderProps{Month: props.Month}, htmxProps),
		CalendarGridHTMX(props, htmxProps),
		g.Group(children),
	)
}

// CalendarHeaderHTMX creates an HTMX-enhanced calendar header
func CalendarHeaderHTMX(props HeaderProps, htmxProps HTMXProps) g.Node {
	classes := lib.CN(
		"flex items-center justify-between mb-4",
		props.Class,
	)

	monthYear := props.Month.Format("January 2006")

	return html.Div(
		html.Class(classes),
		html.Button(
			html.Type("button"),
			html.Class("size-8 rounded-md hover:bg-accent hover:text-accent-foreground inline-flex items-center justify-center"),
			g.Attr("aria-label", "Previous month"),
			hx.Get(fmt.Sprintf("%s?month=%d&year=%d", htmxProps.NavigatePath, 
				getPrevMonth(props.Month).Month(), getPrevMonth(props.Month).Year())),
			hx.Target("#" + htmxProps.ID),
			hx.Swap("outerHTML"),
			icons.ChevronLeft(html.Class("h-4 w-4")),
		),
		html.H2(html.Class("text-sm font-medium"), g.Text(monthYear)),
		html.Button(
			html.Type("button"),
			html.Class("size-8 rounded-md hover:bg-accent hover:text-accent-foreground inline-flex items-center justify-center"),
			g.Attr("aria-label", "Next month"),
			hx.Get(fmt.Sprintf("%s?month=%d&year=%d", htmxProps.NavigatePath, 
				getNextMonth(props.Month).Month(), getNextMonth(props.Month).Year())),
			hx.Target("#" + htmxProps.ID),
			hx.Swap("outerHTML"),
			icons.ChevronRight(html.Class("h-4 w-4")),
		),
	)
}

// CalendarGridHTMX creates an HTMX-enhanced calendar grid
func CalendarGridHTMX(props Props, htmxProps HTMXProps) g.Node {
	// Get the first day of the month
	firstDay := time.Date(props.Month.Year(), props.Month.Month(), 1, 0, 0, 0, 0, props.Month.Location())
	
	// Get the last day of the month
	lastDay := firstDay.AddDate(0, 1, -1)
	
	// Find the start of the calendar (previous Sunday)
	startDate := firstDay
	for startDate.Weekday() != time.Sunday {
		startDate = startDate.AddDate(0, 0, -1)
	}
	
	// Find the end of the calendar (next Saturday)
	endDate := lastDay
	for endDate.Weekday() != time.Saturday {
		endDate = endDate.AddDate(0, 0, 1)
	}

	return html.Div(
		html.Class("w-full"),
		// Weekday headers
		g.If(props.ShowDays,
			html.Div(html.Class("grid grid-cols-7 mb-1"),
				g.Group(g.Map([]string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}, func(day string) g.Node {
					return html.Div(
						html.Class("text-center text-xs font-medium text-muted-foreground p-0"),
						g.Text(day[:2]),
					)
				})),
			),
		),
		// Calendar days
		CalendarDaysHTMX(props, htmxProps, startDate, endDate, firstDay, lastDay),
	)
}

// CalendarDaysHTMX creates an HTMX-enhanced grid of calendar days
func CalendarDaysHTMX(props Props, htmxProps HTMXProps, startDate, endDate, firstDay, lastDay time.Time) g.Node {
	var weeks []g.Node
	currentDate := startDate
	
	for currentDate.Before(endDate.AddDate(0, 0, 1)) {
		var weekDays []g.Node
		
		// Show week number if enabled
		if props.ShowWeeks {
			_, week := currentDate.ISOWeek()
			weekDays = append(weekDays, html.Div(
				html.Class("text-xs text-muted-foreground pr-2"),
				g.Text(fmt.Sprintf("%d", week)),
			))
		}
		
		// Add 7 days for this week
		for i := 0; i < 7; i++ {
			date := currentDate
			dayProps := DayProps{
				Date:     date,
				Selected: isSameDay(date, props.Value),
				Today:    isSameDay(date, time.Now()),
				Outside:  date.Before(firstDay) || date.After(lastDay),
				Disabled: (!props.MinDate.IsZero() && date.Before(props.MinDate)) || 
				         (!props.MaxDate.IsZero() && date.After(props.MaxDate)),
			}
			
			weekDays = append(weekDays, CalendarDayHTMX(dayProps, htmxProps))
			currentDate = currentDate.AddDate(0, 0, 1)
		}
		
		weeks = append(weeks, html.Div(
			html.Class("grid grid-cols-7 mt-2"),
			g.Group(weekDays),
		))
	}
	
	return g.Group(weeks)
}

// CalendarDayHTMX creates an HTMX-enhanced calendar day
func CalendarDayHTMX(props DayProps, htmxProps HTMXProps) g.Node {
	classes := lib.CN(
		"relative p-0 text-center text-sm",
		"inline-flex h-9 w-9 items-center justify-center rounded-md",
		"hover:bg-accent hover:text-accent-foreground",
		"focus:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2",
		lib.CNIf(props.Selected, "bg-primary text-primary-foreground hover:bg-primary hover:text-primary-foreground", ""),
		lib.CNIf(props.Today && !props.Selected, "bg-accent text-accent-foreground", ""),
		lib.CNIf(props.Outside, "text-muted-foreground opacity-50", ""),
		lib.CNIf(props.Disabled, "text-muted-foreground opacity-50 pointer-events-none", ""),
		props.Class,
	)

	dateStr := props.Date.Format("2006-01-02")

	return html.Button(
		html.Type("button"),
		html.Class(classes),
		g.Attr("aria-label", fmt.Sprintf("Select %s", props.Date.Format("January 2, 2006"))),
		g.Attr("aria-selected", fmt.Sprintf("%t", props.Selected)),
		g.If(props.Disabled, html.Disabled()),
		g.If(!props.Disabled, g.Group([]g.Node{
			hx.Post(fmt.Sprintf("%s?date=%s", htmxProps.SelectPath, dateStr)),
			hx.Target("#" + htmxProps.ID),
			hx.Swap("outerHTML"),
		})),
		g.Text(fmt.Sprintf("%d", props.Date.Day())),
	)
}

// getPrevMonth returns the previous month
func getPrevMonth(current time.Time) time.Time {
	return current.AddDate(0, -1, 0)
}

// getNextMonth returns the next month
func getNextMonth(current time.Time) time.Time {
	return current.AddDate(0, 1, 0)
}

// DatePickerHTMX creates a date picker with input and calendar dropdown
func DatePickerHTMX(htmxProps HTMXProps, value string, placeholder string, class ...string) g.Node {
	inputID := htmxProps.ID + "-input"
	dropdownID := htmxProps.ID + "-dropdown"

	return html.Div(
		html.Class("relative"),
		html.Div(
			html.Class("flex"),
			html.Input(
				html.ID(inputID),
				html.Type("text"),
				html.Value(value),
				html.Placeholder(placeholder),
				html.Class(lib.CN("pr-10", lib.CN(class...))),
				hx.Get(htmxProps.UpdatePath),
				html.Target("#" + dropdownID),
				hx.Swap("innerHTML"),
				hx.Trigger("focus"),
			),
			html.Button(
				html.Type("button"),
				html.Class("absolute right-0 top-0 h-full px-3 py-2 hover:bg-transparent"),
				hx.Get(htmxProps.UpdatePath),
				html.Target("#" + dropdownID),
				hx.Swap("innerHTML"),
				icons.Calendar(html.Class("h-4 w-4 opacity-50")),
			),
		),
		html.Div(
			html.ID(dropdownID),
			html.Class("absolute z-50 mt-1"),
			// Calendar will be loaded here
		),
	)
}

// RenderCalendarDropdown renders the calendar dropdown for date picker
func RenderCalendarDropdown(htmxProps HTMXProps, selectedDate time.Time) g.Node {
	calendarProps := Props{
		Value: selectedDate,
		Month: selectedDate,
	}
	
	if calendarProps.Month.IsZero() {
		calendarProps.Month = time.Now()
	}

	return html.Div(
		html.Class("bg-background border rounded-md shadow-lg"),
		hx.On("click", "event.stopPropagation()"),
		NewHTMX(calendarProps, htmxProps),
	)
}

// MonthYearPickerHTMX creates a month/year picker with HTMX
func MonthYearPickerHTMX(htmxProps HTMXProps, currentMonth time.Time, class ...string) g.Node {
	return html.Div(
		html.ID(htmxProps.ID),
		html.Class(lib.CN("bg-background border rounded-md p-4", lib.CN(class...))),
		html.Div(
			html.Class("grid grid-cols-2 gap-4"),
			// Month picker
			html.Div(
				html.H3(html.Class("text-sm font-medium mb-2"), g.Text("Month")),
				html.Div(
					html.Class("grid grid-cols-3 gap-1"),
					g.Group(g.Map([]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}, func(month string) g.Node {
						monthNum := indexOf([]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}, month) + 1
						isSelected := int(currentMonth.Month()) == monthNum
						
						return html.Button(
							html.Type("button"),
							html.Class(lib.CN(
								"text-xs rounded px-2 py-1 hover:bg-accent hover:text-accent-foreground",
								lib.CNIf(isSelected, "bg-primary text-primary-foreground", ""),
							)),
							hx.Post(fmt.Sprintf("%s?month=%d&year=%d", htmxProps.UpdatePath, monthNum, currentMonth.Year())),
							hx.Target("#" + htmxProps.ID),
							hx.Swap("outerHTML"),
							g.Text(month),
						)
					})),
				),
			),
			// Year picker
			html.Div(
				html.H3(html.Class("text-sm font-medium mb-2"), g.Text("Year")),
				html.Div(
					html.Class("flex items-center gap-2"),
					html.Button(
						html.Type("button"),
						html.Class("size-6 rounded hover:bg-accent hover:text-accent-foreground inline-flex items-center justify-center"),
						hx.Post(fmt.Sprintf("%s?month=%d&year=%d", htmxProps.UpdatePath, currentMonth.Month(), currentMonth.Year()-1)),
						hx.Target("#" + htmxProps.ID),
						hx.Swap("outerHTML"),
						icons.ChevronLeft(html.Class("h-3 w-3")),
					),
					html.Span(html.Class("text-sm font-medium w-12 text-center"), g.Text(fmt.Sprintf("%d", currentMonth.Year()))),
					html.Button(
						html.Type("button"),
						html.Class("size-6 rounded hover:bg-accent hover:text-accent-foreground inline-flex items-center justify-center"),
						hx.Post(fmt.Sprintf("%s?month=%d&year=%d", htmxProps.UpdatePath, currentMonth.Month(), currentMonth.Year()+1)),
						hx.Target("#" + htmxProps.ID),
						hx.Swap("outerHTML"),
						icons.ChevronRight(html.Class("h-3 w-3")),
					),
				),
			),
		),
	)
}

// CalendarHandlers creates HTTP handlers for calendar components
func CalendarHandlers(mux *http.ServeMux) {
	// Basic calendar navigation
	htmxProps := HTMXProps{
		ID:           "calendar-example",
		NavigatePath: "/api/calendar/navigate",
		SelectPath:   "/api/calendar/select",
		UpdatePath:   "/api/calendar/update",
	}

	mux.HandleFunc("/api/calendar/navigate", func(w http.ResponseWriter, r *http.Request) {
		monthStr := r.URL.Query().Get("month")
		yearStr := r.URL.Query().Get("year")
		
		month, _ := strconv.Atoi(monthStr)
		year, _ := strconv.Atoi(yearStr)
		
		if month < 1 || month > 12 {
			month = int(time.Now().Month())
		}
		if year < 1900 || year > 2100 {
			year = time.Now().Year()
		}
		
		selectedDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
		
		props := Props{
			Month: selectedDate,
		}
		
		node := NewHTMX(props, htmxProps)
		node.Render(w)
	})

	mux.HandleFunc("/api/calendar/select", func(w http.ResponseWriter, r *http.Request) {
		dateStr := r.URL.Query().Get("date")
		selectedDate, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			selectedDate = time.Now()
		}
		
		props := Props{
			Value: selectedDate,
			Month: selectedDate,
		}
		
		node := NewHTMX(props, htmxProps)
		node.Render(w)
	})

	// Date picker handlers
	datePickerProps := HTMXProps{
		ID:           "datepicker-example",
		NavigatePath: "/api/datepicker/navigate",
		SelectPath:   "/api/datepicker/select",
		UpdatePath:   "/api/datepicker/show",
	}

	mux.HandleFunc("/api/datepicker/show", func(w http.ResponseWriter, r *http.Request) {
		// Show the calendar dropdown
		node := RenderCalendarDropdown(datePickerProps, time.Now())
		node.Render(w)
	})

	mux.HandleFunc("/api/datepicker/select", func(w http.ResponseWriter, r *http.Request) {
		dateStr := r.URL.Query().Get("date")
		selectedDate, _ := time.Parse("2006-01-02", dateStr)
		
		// Return the selected date and close the dropdown
		node := html.Div(
			// Update the input value
			html.Script(g.Raw(fmt.Sprintf(`
				document.getElementById('%s-input').value = '%s';
				document.getElementById('%s-dropdown').innerHTML = '';
			`, datePickerProps.ID, selectedDate.Format("01/02/2006"), datePickerProps.ID))),
		)
		node.Render(w)
	})

	// Month/Year picker handlers
	monthYearProps := HTMXProps{
		ID:         "monthyear-picker",
		UpdatePath: "/api/monthyear/update",
	}

	mux.HandleFunc("/api/monthyear/update", func(w http.ResponseWriter, r *http.Request) {
		monthStr := r.URL.Query().Get("month")
		yearStr := r.URL.Query().Get("year")
		
		month, _ := strconv.Atoi(monthStr)
		year, _ := strconv.Atoi(yearStr)
		
		if month < 1 || month > 12 {
			month = int(time.Now().Month())
		}
		if year < 1900 || year > 2100 {
			year = time.Now().Year()
		}
		
		selectedDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
		
		node := MonthYearPickerHTMX(monthYearProps, selectedDate)
		node.Render(w)
	})
}