package slider

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	hx "maragu.dev/gomponents-htmx"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// HTMXProps defines properties for HTMX-enhanced slider
type HTMXProps struct {
	ID          string
	UpdatePath  string
	DragPath    string
	InitPath    string
}

// HTMXSlider creates an HTMX-enhanced slider
func HTMXSlider(props Props, htmxProps HTMXProps) g.Node {
	// Set defaults
	if props.Max == 0 {
		props.Max = 100
	}
	if props.Step == 0 {
		props.Step = 1
	}
	if props.Orientation == "" {
		props.Orientation = "horizontal"
	}
	if len(props.Value) == 0 {
		props.Value = []int{props.Min}
	}

	// Calculate range percentage
	rangeStart := float64(props.Value[0]-props.Min) / float64(props.Max-props.Min) * 100
	rangeEnd := float64(100)
	if len(props.Value) > 1 {
		rangeEnd = float64(props.Value[len(props.Value)-1]-props.Min) / float64(props.Max-props.Min) * 100
	}

	attributes := []g.Node{
		g.Attr("id", htmxProps.ID),
		g.Attr("data-slider", "true"),
		g.Attr("data-orientation", props.Orientation),
		g.Attr("data-min", fmt.Sprintf("%d", props.Min)),
		g.Attr("data-max", fmt.Sprintf("%d", props.Max)),
		g.Attr("data-step", fmt.Sprintf("%d", props.Step)),
		html.Class(lib.CN(
			"relative flex w-full touch-none items-center select-none",
			"data-[disabled]:opacity-50",
			"data-[orientation=vertical]:h-full data-[orientation=vertical]:min-h-44",
			"data-[orientation=vertical]:w-auto data-[orientation=vertical]:flex-col",
			props.Class,
		)),
		hx.Get(htmxProps.InitPath),
		hx.Trigger("load"),
		hx.Target("this"),
		hx.Swap("outerHTML"),
	}

	if props.Disabled {
		attributes = append(attributes, g.Attr("data-disabled", "true"))
	}

	// Build slider structure
	sliderContent := []g.Node{
		// Track
		html.Div(
			g.Attr("data-slider-track", "true"),
			html.Class(lib.CN(
				"bg-muted relative grow overflow-hidden rounded-full",
				"data-[orientation=horizontal]:h-1.5 data-[orientation=horizontal]:w-full",
				"data-[orientation=vertical]:h-full data-[orientation=vertical]:w-1.5",
			)),
			// Click handler for track
			hx.Post(htmxProps.UpdatePath),
			hx.Target("#" + htmxProps.ID),
			hx.Swap("outerHTML"),
			hx.Vals(`js:{
				const rect = event.currentTarget.getBoundingClientRect();
				const orientation = document.getElementById('` + htmxProps.ID + `').dataset.orientation;
				let percent;
				if (orientation === 'vertical') {
					percent = 1 - ((event.clientY - rect.top) / rect.height);
				} else {
					percent = (event.clientX - rect.left) / rect.width;
				}
				return { percent: percent, index: -1 };
			}`),
			
			// Range (filled portion)
			html.Div(
				g.Attr("data-slider-range", "true"),
				html.Class(lib.CN(
					"bg-primary absolute pointer-events-none",
					"data-[orientation=horizontal]:h-full",
					"data-[orientation=vertical]:w-full",
				)),
				g.If(props.Orientation == "horizontal",
					g.Attr("style", fmt.Sprintf("left: %.2f%%; width: %.2f%%", rangeStart, rangeEnd-rangeStart)),
				),
				g.If(props.Orientation == "vertical",
					g.Attr("style", fmt.Sprintf("bottom: %.2f%%; height: %.2f%%", rangeStart, rangeEnd-rangeStart)),
				),
			),
		),
	}

	// Add thumbs
	for i, value := range props.Value {
		position := float64(value-props.Min) / float64(props.Max-props.Min) * 100
		
		thumb := html.Div(
			g.Attr("data-slider-thumb", "true"),
			g.Attr("data-index", fmt.Sprintf("%d", i)),
			g.Attr("data-value", fmt.Sprintf("%d", value)),
			g.Attr("tabindex", "0"),
			g.Attr("role", "slider"),
			g.Attr("aria-valuemin", fmt.Sprintf("%d", props.Min)),
			g.Attr("aria-valuemax", fmt.Sprintf("%d", props.Max)),
			g.Attr("aria-valuenow", fmt.Sprintf("%d", value)),
			g.Attr("aria-orientation", props.Orientation),
			html.Class(lib.CN(
				"border-primary bg-background ring-ring/50",
				"block size-4 shrink-0 rounded-full border shadow-sm",
				"transition-[color,box-shadow] hover:ring-4 focus-visible:ring-4",
				"focus-visible:outline-none disabled:pointer-events-none disabled:opacity-50",
				"absolute cursor-grab active:cursor-grabbing",
			)),
			g.If(props.Orientation == "horizontal",
				g.Attr("style", fmt.Sprintf("left: calc(%.2f%% - 8px)", position)),
			),
			g.If(props.Orientation == "vertical",
				g.Attr("style", fmt.Sprintf("bottom: calc(%.2f%% - 8px)", position)),
			),
			g.If(props.Disabled,
				g.Attr("disabled", "true"),
			),
			
			// Drag handling
			hx.Post(htmxProps.DragPath),
			hx.Target("#" + htmxProps.ID),
			hx.Swap("outerHTML"),
			hx.Trigger("mousedown"),
			hx.Vals(fmt.Sprintf(`{"index": %d, "action": "start"}`, i)),
			
			// Keyboard handling
			g.Attr("onkeydown", fmt.Sprintf(`
				const step = parseInt(document.getElementById('%s').dataset.step);
				const min = parseInt(document.getElementById('%s').dataset.min);
				const max = parseInt(document.getElementById('%s').dataset.max);
				let value = parseInt(this.dataset.value);
				let changed = false;
				
				if (event.key === 'ArrowRight' || event.key === 'ArrowUp') {
					value = Mathtml.min(max, value + step);
					changed = true;
				} else if (event.key === 'ArrowLeft' || event.key === 'ArrowDown') {
					value = Mathtml.max(min, value - step);
					changed = true;
				} else if (event.key === 'Home') {
					value = min;
					changed = true;
				} else if (event.key === 'End') {
					value = max;
					changed = true;
				}
				
				if (changed) {
					event.preventDefault();
					htmx.ajax('POST', '%s', {
						target: '#%s',
						swap: 'outerHTML',
						values: { index: %d, value: value }
					});
				}
			`, htmxProps.ID, htmxProps.ID, htmxProps.ID, htmxProps.UpdatePath, htmxProps.ID, i)),
		)

		// Hidden input for form submission
		if props.Name != "" {
			sliderContent = append(sliderContent,
				html.Input(
					html.Type("hidden"),
					html.Name(fmt.Sprintf("%s[%d]", props.Name, i)),
					html.Value(fmt.Sprintf("%d", value)),
					g.Attr("id", fmt.Sprintf("%s-input-%d", htmxProps.ID, i)),
				),
			)
		}

		sliderContent = append(sliderContent, thumb)
	}

	return html.Div(
		g.Group(append(attributes, sliderContent...)),
	)
}

// SliderState represents the state of a slider
type SliderState struct {
	Values []int
	Min    int
	Max    int
	Step   int
}

// Server-side state management
var sliderStates = make(map[string]*SliderState)

// SliderHandlers creates HTTP handlers for slider functionality
func SliderHandlers(mux *http.ServeMux, baseProps Props, htmxProps HTMXProps) {
	// Initialize state
	if _, exists := sliderStates[htmxProps.ID]; !exists {
		sliderStates[htmxProps.ID] = &SliderState{
			Values: baseProps.Value,
			Min:    baseProps.Min,
			Max:    baseProps.Max,
			Step:   baseProps.Step,
		}
		if sliderStates[htmxProps.ID].Max == 0 {
			sliderStates[htmxProps.ID].Max = 100
		}
		if sliderStates[htmxProps.ID].Step == 0 {
			sliderStates[htmxProps.ID].Step = 1
		}
	}

	// Initialize handler
	mux.HandleFunc(htmxProps.InitPath, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		state := sliderStates[htmxProps.ID]
		props := baseProps
		props.Value = state.Values
		props.Min = state.Min
		props.Max = state.Max
		props.Step = state.Step

		HTMXSlider(props, htmxProps).Render(w)
	})

	// Update handler
	mux.HandleFunc(htmxProps.UpdatePath, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		state := sliderStates[htmxProps.ID]
		
		// Parse form values
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		// Handle different update types
		if percentStr := r.FormValue("percent"); percentStr != "" {
			// Track click - find nearest thumb or add new one
			percent, _ := strconv.ParseFloat(percentStr, 64)
			newValue := int(percent*float64(state.Max-state.Min) + float64(state.Min))
			
			// Round to nearest step
			newValue = ((newValue + state.Step/2) / state.Step) * state.Step
			newValue = max(state.Min, min(state.Max, newValue))

			// For single slider, just update the value
			if len(state.Values) == 1 {
				state.Values[0] = newValue
			} else {
				// For range slider, update the nearest thumb
				minDist := state.Max - state.Min
				nearestIdx := 0
				for i, v := range state.Values {
					dist := abs(v - newValue)
					if dist < minDist {
						minDist = dist
						nearestIdx = i
					}
				}
				state.Values[nearestIdx] = newValue
				
				// Ensure values are in order
				if len(state.Values) == 2 && state.Values[0] > state.Values[1] {
					state.Values[0], state.Values[1] = state.Values[1], state.Values[0]
				}
			}
		} else if indexStr := r.FormValue("index"); indexStr != "" {
			// Direct thumb update
			index, _ := strconv.Atoi(indexStr)
			if valueStr := r.FormValue("value"); valueStr != "" {
				value, _ := strconv.Atoi(valueStr)
				value = max(state.Min, min(state.Max, value))
				
				if index >= 0 && index < len(state.Values) {
					state.Values[index] = value
					
					// Ensure values are in order for range sliders
					if len(state.Values) == 2 {
						if index == 0 && state.Values[0] > state.Values[1] {
							state.Values[0] = state.Values[1]
						} else if index == 1 && state.Values[1] < state.Values[0] {
							state.Values[1] = state.Values[0]
						}
					}
				}
			}
		}

		// Return updated slider
		props := baseProps
		props.Value = state.Values
		props.Min = state.Min
		props.Max = state.Max
		props.Step = state.Step

		HTMXSlider(props, htmxProps).Render(w)
	})

	// Drag handler
	var dragState = make(map[string]struct {
		Index int
		Start int
	})

	mux.HandleFunc(htmxProps.DragPath, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		state := sliderStates[htmxProps.ID]
		sessionID := r.Header.Get("X-Session-ID")
		if sessionID == "" {
			sessionID = "default"
		}

		// Parse request
		var data map[string]interface{}
		json.NewDecoder(r.Body).Decode(&data)

		action, _ := data["action"].(string)
		index := int(data["index"].(float64))

		if action == "start" && index >= 0 && index < len(state.Values) {
			// Start dragging
			dragState[sessionID] = struct {
				Index int
				Start int
			}{
				Index: index,
				Start: state.Values[index],
			}

			// Add mouse move and up handlers
			w.Header().Set("HX-Trigger-After-Swap", "setupDrag")
			
			// Return current slider
			props := baseProps
			props.Value = state.Values
			props.Min = state.Min
			props.Max = state.Max
			props.Step = state.Step

			HTMXSlider(props, htmxProps).Render(w)
		}
	})
}

// Helper functions
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// HTMXSliderWithValue creates a slider with live value display
func HTMXSliderWithValue(props Props, htmxProps HTMXProps) g.Node {
	valueText := fmt.Sprintf("%d", props.Value[0])
	if len(props.Value) > 1 {
		valueText = fmt.Sprintf("%d - %d", props.Value[0], props.Value[1])
	}
	
	return html.Div(
		html.Class("space-y-2"),
		// Value display
		html.Div(
			html.Class("flex justify-between items-center"),
			html.Label(g.Text("Value")),
			html.Span(
				html.Class("text-sm font-medium"),
				g.Attr("id", htmxProps.ID+"-value"),
				g.Text(valueText),
				// Update value display when slider changes
				hx.Get(htmxProps.UpdatePath + "/value"),
				hx.Trigger(fmt.Sprintf("slider-update from:#%s", htmxProps.ID)),
				hx.Target("this"),
				hx.Swap("innerHTML"),
			),
		),
		// Slider
		HTMXSlider(props, htmxProps),
	)
}

// Additional handler for value display updates
func SliderValueHandler(mux *http.ServeMux, htmxProps HTMXProps) {
	mux.HandleFunc(htmxProps.UpdatePath+"/value", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		state := sliderStates[htmxProps.ID]
		if state == nil {
			http.Error(w, "Slider not found", http.StatusNotFound)
			return
		}

		valueText := fmt.Sprintf("%d", state.Values[0])
		if len(state.Values) > 1 {
			valueText = fmt.Sprintf("%d - %d", state.Values[0], state.Values[1])
		}

		w.Write([]byte(valueText))
	})
}

// DragScript returns JavaScript for handling drag operations
func DragScript(htmxProps HTMXProps) string {
	return fmt.Sprintf(`
		document.addEventListener('setupDrag', function(e) {
			const slider = document.getElementById('%s');
			const activeThumb = slider.querySelector('[data-slider-thumb][data-active="true"]');
			if (!activeThumb) return;
			
			const index = parseInt(activeThumb.dataset.index);
			const orientation = slider.dataset.orientation;
			const min = parseInt(slider.dataset.min);
			const max = parseInt(slider.dataset.max);
			const step = parseInt(slider.dataset.step);
			const track = slider.querySelector('[data-slider-track]');
			
			function handleMove(e) {
				e.preventDefault();
				const rect = track.getBoundingClientRect();
				let percent;
				
				if (orientation === 'vertical') {
					percent = 1 - ((e.clientY - rect.top) / rect.height);
				} else {
					percent = (e.clientX - rect.left) / rect.width;
				}
				
				percent = Mathtml.max(0, Mathtml.min(1, percent));
				let value = Mathtml.round(percent * (max - min) + min);
				value = Mathtml.round(value / step) * step;
				
				htmx.ajax('POST', '%s', {
					target: '#%s',
					swap: 'outerHTML',
					values: { index: index, value: value }
				});
			}
			
			function handleUp(e) {
				document.removeEventListener('mousemove', handleMove);
				document.removeEventListener('mouseup', handleUp);
				activeThumb.removeAttribute('data-active');
			}
			
			activeThumb.setAttribute('data-active', 'true');
			document.addEventListener('mousemove', handleMove);
			document.addEventListener('mouseup', handleUp);
		});
	`, htmxProps.ID, htmxProps.UpdatePath, htmxProps.ID)
}