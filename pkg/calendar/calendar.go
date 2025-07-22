package calendar

import (
	"fmt"
	"time"

	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
	"github.com/rizome-dev/shadcn-gomponents/lib/icons"
)

// Props defines the properties for the Calendar component
type Props struct {
	Value     time.Time // Currently selected date
	Month     time.Time // Month to display
	ShowDays  bool      // Whether to show weekday names
	ShowWeeks bool      // Whether to show week numbers
	MinDate   time.Time // Minimum selectable date
	MaxDate   time.Time // Maximum selectable date
	Class     string    // Additional custom classes
}

// HeaderProps defines the properties for the CalendarHeader
type HeaderProps struct {
	Month         time.Time
	ShowDropdowns bool   // Whether to show month/year dropdowns
	Class         string
}

// DayProps defines the properties for a calendar day
type DayProps struct {
	Date      time.Time
	Selected  bool
	Today     bool
	Outside   bool // Day is outside the current month
	Disabled  bool
	Class     string
}

// New creates a new Calendar component
func New(props Props, children ...g.Node) g.Node {
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
		html.Class(classes),
		CalendarHeader(HeaderProps{Month: props.Month}),
		CalendarGrid(props),
		g.Group(children),
	)
}

// CalendarHeader creates the calendar header with month/year
func CalendarHeader(props HeaderProps) g.Node {
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
			icons.ChevronLeft(html.Class("h-4 w-4")),
		),
		html.H2(html.Class("text-sm font-medium"), g.Text(monthYear)),
		html.Button(
			html.Type("button"),
			html.Class("size-8 rounded-md hover:bg-accent hover:text-accent-foreground inline-flex items-center justify-center"),
			g.Attr("aria-label", "Next month"),
			icons.ChevronRight(html.Class("h-4 w-4")),
		),
	)
}

// CalendarGrid creates the calendar grid with days
func CalendarGrid(props Props) g.Node {
	// Get the first day of the month
	firstDay := time.Date(props.Month.Year(), props.Month.Month(), 1, 0, 0, 0, 0, props.Month.Location())
	
	// Get the last day of the month
	lastDay := firstDay.AddDate(0, 1, -1)
	
	// Find the start of the calendar (previous Sunday or Monday depending on locale)
	// For now, we'll use Sunday as the first day of the week
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
			html.Div(
				html.Class("grid grid-cols-7 mb-1"),
				html.Div(html.Class("text-center text-xs font-medium text-muted-foreground p-0"), g.Text("Su")),
				html.Div(html.Class("text-center text-xs font-medium text-muted-foreground p-0"), g.Text("Mo")),
				html.Div(html.Class("text-center text-xs font-medium text-muted-foreground p-0"), g.Text("Tu")),
				html.Div(html.Class("text-center text-xs font-medium text-muted-foreground p-0"), g.Text("We")),
				html.Div(html.Class("text-center text-xs font-medium text-muted-foreground p-0"), g.Text("Th")),
				html.Div(html.Class("text-center text-xs font-medium text-muted-foreground p-0"), g.Text("Fr")),
				html.Div(html.Class("text-center text-xs font-medium text-muted-foreground p-0"), g.Text("Sa")),
			),
		),
		// Calendar days
		CalendarDays(props, startDate, endDate, firstDay, lastDay),
	)
}

// CalendarDays creates the grid of calendar days
func CalendarDays(props Props, startDate, endDate, firstDay, lastDay time.Time) g.Node {
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
			
			weekDays = append(weekDays, CalendarDay(dayProps))
			currentDate = currentDate.AddDate(0, 0, 1)
		}
		
		weeks = append(weeks, html.Div(
			html.Class("grid grid-cols-7 mt-2"),
			g.Group(weekDays),
		))
	}
	
	return g.Group(weeks)
}

// CalendarDay creates a single calendar day
func CalendarDay(props DayProps) g.Node {
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

	return html.Button(
		html.Type("button"),
		html.Class(classes),
		g.Attr("aria-label", fmt.Sprintf("Select %s", props.Date.Format("January 2, 2006"))),
		g.Attr("aria-selected", fmt.Sprintf("%t", props.Selected)),
		g.If(props.Disabled, html.Disabled()),
		g.Text(fmt.Sprintf("%d", props.Date.Day())),
	)
}

// isSameDay checks if two dates are the same day
func isSameDay(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

// MonthPicker creates a month picker component
func MonthPicker(currentMonth time.Time, class ...string) g.Node {
	classes := lib.CN(
		"grid grid-cols-3 gap-2 p-2",
		lib.CN(class...),
	)
	
	months := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	
	return html.Div(
		html.Class(classes),
		g.Group(g.Map(months, func(month string) g.Node {
			monthNum := indexOf(months, month) + 1
			isSelected := int(currentMonth.Month()) == monthNum
			
			return html.Button(
				html.Type("button"),
				html.Class(lib.CN(
					"text-sm rounded-md px-3 py-1.5 hover:bg-accent hover:text-accent-foreground",
					lib.CNIf(isSelected, "bg-primary text-primary-foreground", ""),
				)),
				g.Text(month),
			)
		})),
	)
}

// YearPicker creates a year picker component
func YearPicker(currentYear int, startYear, endYear int, class ...string) g.Node {
	classes := lib.CN(
		"grid grid-cols-3 gap-2 p-2 max-h-64 overflow-y-auto",
		lib.CN(class...),
	)
	
	if startYear == 0 {
		startYear = currentYear - 50
	}
	if endYear == 0 {
		endYear = currentYear + 50
	}
	
	var years []int
	for y := startYear; y <= endYear; y++ {
		years = append(years, y)
	}
	
	return html.Div(
		html.Class(classes),
		g.Group(g.Map(years, func(year int) g.Node {
			isSelected := year == currentYear
			
			return html.Button(
				html.Type("button"),
				html.Class(lib.CN(
					"text-sm rounded-md px-3 py-1.5 hover:bg-accent hover:text-accent-foreground",
					lib.CNIf(isSelected, "bg-primary text-primary-foreground", ""),
				)),
				g.Text(fmt.Sprintf("%d", year)),
			)
		})),
	)
}

// indexOf returns the index of a string in a slice
func indexOf(slice []string, item string) int {
	for i, v := range slice {
		if v == item {
			return i
		}
	}
	return -1
}

// DateRangeCalendar creates a calendar for selecting date ranges
func DateRangeCalendar(startDate, endDate time.Time, month time.Time, class ...string) g.Node {
	props := Props{
		Month: month,
		Class: lib.CN(class...),
	}
	
	// Custom day renderer for date ranges
	firstDay := time.Date(month.Year(), month.Month(), 1, 0, 0, 0, 0, month.Location())
	lastDay := firstDay.AddDate(0, 1, -1)
	
	startCalDate := firstDay
	for startCalDate.Weekday() != time.Sunday {
		startCalDate = startCalDate.AddDate(0, 0, -1)
	}
	
	endCalDate := lastDay
	for endCalDate.Weekday() != time.Saturday {
		endCalDate = endCalDate.AddDate(0, 0, 1)
	}
	
	return New(props, CustomDateRangeGrid(startDate, endDate, month, startCalDate, endCalDate, firstDay, lastDay))
}

// CustomDateRangeGrid creates a custom grid for date range selection
func CustomDateRangeGrid(startDate, endDate, month, startCalDate, endCalDate, firstDay, lastDay time.Time) g.Node {
	var weeks []g.Node
	currentDate := startCalDate
	
	// Weekday headers
	weeks = append(weeks, html.Div(
		html.Class("grid grid-cols-7 mb-1"),
		html.Div(html.Class("text-center text-xs font-medium text-muted-foreground p-0"), g.Text("Su")),
		html.Div(html.Class("text-center text-xs font-medium text-muted-foreground p-0"), g.Text("Mo")),
		html.Div(html.Class("text-center text-xs font-medium text-muted-foreground p-0"), g.Text("Tu")),
		html.Div(html.Class("text-center text-xs font-medium text-muted-foreground p-0"), g.Text("We")),
		html.Div(html.Class("text-center text-xs font-medium text-muted-foreground p-0"), g.Text("Th")),
		html.Div(html.Class("text-center text-xs font-medium text-muted-foreground p-0"), g.Text("Fr")),
		html.Div(html.Class("text-center text-xs font-medium text-muted-foreground p-0"), g.Text("Sa")),
	))
	
	for currentDate.Before(endCalDate.AddDate(0, 0, 1)) {
		var weekDays []g.Node
		
		for i := 0; i < 7; i++ {
			date := currentDate
			isStart := isSameDay(date, startDate)
			isEnd := isSameDay(date, endDate)
			isInRange := !startDate.IsZero() && !endDate.IsZero() && 
			            date.After(startDate) && date.Before(endDate)
			isOutside := date.Before(firstDay) || date.After(lastDay)
			
			classes := lib.CN(
				"relative p-0 text-center text-sm",
				"inline-flex h-9 w-9 items-center justify-center",
				"hover:bg-accent hover:text-accent-foreground",
				lib.CNIf(isStart || isEnd, "bg-primary text-primary-foreground hover:bg-primary hover:text-primary-foreground", ""),
				lib.CNIf(isStart, "rounded-l-md", ""),
				lib.CNIf(isEnd, "rounded-r-md", ""),
				lib.CNIf(isInRange, "bg-accent rounded-none", ""),
				lib.CNIf(isOutside, "text-muted-foreground opacity-50", ""),
				lib.CNIf(!isStart && !isEnd && !isInRange, "rounded-md", ""),
			)
			
			weekDays = append(weekDays, html.Button(
				html.Type("button"),
				html.Class(classes),
				g.Text(fmt.Sprintf("%d", date.Day())),
			))
			
			currentDate = currentDate.AddDate(0, 0, 1)
		}
		
		weeks = append(weeks, html.Div(
			html.Class("grid grid-cols-7 mt-2"),
			g.Group(weekDays),
		))
	}
	
	return g.Group(weeks)
}