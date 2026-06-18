package org

import (
	"github.com/invopop/gobl/cbc"
	"github.com/invopop/gobl/rules"
	"github.com/invopop/gobl/rules/is"
)

// Attribute describes a named feature or property of the parent object,
// such as the colour or size of an item.
type Attribute struct {
	// Machine-readable key that identifies the attribute or property.
	Key cbc.Key `json:"key" jsonschema:"title=Key"`
	// Human-readable label for internal use, not included in output documents.
	Label string `json:"label,omitempty" jsonschema:"title=Label"`
	// Value of the attribute or property.
	Value string `json:"value" jsonschema:"title=Value"`
}

func attributeRules() *rules.Set {
	return rules.For(new(Attribute),
		rules.Field("key",
			rules.Assert("01", "attribute key is required", is.Present),
		),
		rules.Field("value",
			rules.Assert("02", "attribute value is required", is.Present),
		),
	)
}

func normalizeAttribute(a *Attribute) {
	a.Label = cbc.NormalizeString(a.Label)
	a.Value = cbc.NormalizeString(a.Value)
}

// CleanAttributes removes any nil or empty attributes from the list,
// returning nil if none remain.
func CleanAttributes(attrs []*Attribute) []*Attribute {
	cleaned := make([]*Attribute, 0, len(attrs))
	for _, a := range attrs {
		if a == nil || (a.Key == "" && a.Label == "" && a.Value == "") {
			continue
		}
		cleaned = append(cleaned, a)
	}
	if len(cleaned) == 0 {
		return nil
	}
	return cleaned
}
