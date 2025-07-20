package lib

// VariantProps represents the common properties for component variants
type VariantProps struct {
	Variant string
	Size    string
	Class   string // Additional custom classes
}

// VariantConfig defines the structure for component variant configurations
type VariantConfig struct {
	Base     string                       // Base classes always applied
	Variants map[string]map[string]string // variant type -> variant name -> classes
	Defaults map[string]string            // Default variant selections
}

// GetClasses returns the combined classes for the given variant configuration
func (vc *VariantConfig) GetClasses(props VariantProps) string {
	classes := []string{vc.Base}

	// Apply variant classes
	if props.Variant == "" && vc.Defaults["variant"] != "" {
		props.Variant = vc.Defaults["variant"]
	}
	if variantClasses, ok := vc.Variants["variant"]; ok {
		if class, ok := variantClasses[props.Variant]; ok {
			classes = append(classes, class)
		}
	}

	// Apply size classes
	if props.Size == "" && vc.Defaults["size"] != "" {
		props.Size = vc.Defaults["size"]
	}
	if sizeClasses, ok := vc.Variants["size"]; ok {
		if class, ok := sizeClasses[props.Size]; ok {
			classes = append(classes, class)
		}
	}

	// Add custom classes
	if props.Class != "" {
		classes = append(classes, props.Class)
	}

	return MergeClasses(classes...)
}