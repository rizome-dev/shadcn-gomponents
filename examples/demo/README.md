# Shadcn Gomponents Demo App

This demo application showcases all 47 components available in the Shadcn Gomponents library.

## Running the Demo

1. Navigate to this directory:
   ```bash
   cd examples/demo
   ```

2. Run the demo app:
   ```bash
   go run main.go
   ```

3. Open your browser to http://localhost:8080

## Features

- **Component Gallery**: Browse all 47 components organized by category
- **Live Examples**: See each component in action with various configurations
- **Interactive Demos**: Components with HTMX integration for dynamic behavior
- **Responsive Design**: All components work seamlessly on mobile and desktop

## Available Routes

- `/` - Home page with feature overview
- `/components/` - Complete component gallery
- `/[component-name]` - Individual component demo pages

## Components Included

### Layout (5)
- Aspect Ratio
- Card
- Resizable
- Scroll Area  
- Separator

### Forms (13)
- Button
- Checkbox
- Form
- Input
- Input OTP
- Label
- Radio Group
- Select
- Slider
- Switch
- Textarea
- Toggle
- Toggle Group

### Data Display (7)
- Avatar
- Badge
- Calendar
- Chart
- Progress
- Skeleton
- Table

### Feedback (8)
- Alert
- Alert Dialog
- Dialog
- Drawer
- Sheet
- Sonner
- Toast
- Tooltip

### Navigation (11)
- Breadcrumb
- Command
- Context Menu
- Dropdown Menu
- Hover Card
- Menubar
- Navigation Menu
- Pagination
- Popover
- Sidebar
- Tabs

### Display (3)
- Accordion
- Carousel
- Collapsible

## Development

The demo app uses:
- **Gomponents** for HTML generation
- **HTMX** for dynamic interactions
- **Tailwind CSS** for styling
- **Go standard library** for HTTP server

## Contributing

To add a new component to the demo:

1. Import the component package
2. Add a route handler in `main.go`
3. Add the component to the `ComponentsListPage` function
4. Ensure the component has an `Example()` function

## License

This demo app is part of the Shadcn Gomponents project and follows the same license.