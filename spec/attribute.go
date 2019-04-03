package spec

import (
	"fmt"
)

// AttributeType represents the various type for an attribute.
type AttributeType string

// Various values for AttributeType.
const (
	AttributeTypeString  AttributeType = "string"
	AttributeTypeInt     AttributeType = "integer"
	AttributeTypeFloat   AttributeType = "float"
	AttributeTypeBool    AttributeType = "boolean"
	AttributeTypeEnum    AttributeType = "enum"
	AttributeTypeList    AttributeType = "list"
	AttributeTypeObject  AttributeType = "object"
	AttributeTypeTime    AttributeType = "time"
	AttributeTypeExt     AttributeType = "external"
	AttributeTypeRef     AttributeType = "ref"
	AttributeTypeRefList AttributeType = "refList"
	AttributeTypeRefMap  AttributeType = "refMap"
)

// An Attribute represents a regolithe specification attribute.
type Attribute struct {

	// NOTE: Order of attributes matters!
	// The YAML will be dumped respecting this order.

	Name                string                 `yaml:"name,omitempty"                   json:"name,omitempty"`
	Description         string                 `yaml:"description,omitempty"            json:"description,omitempty"`
	Type                AttributeType          `yaml:"type,omitempty"                   json:"type,omitempty"`
	Exposed             bool                   `yaml:"exposed,omitempty"                json:"exposed,omitempty"`
	SubType             string                 `yaml:"subtype,omitempty"                json:"subtype,omitempty"`
	Stored              bool                   `yaml:"stored,omitempty"                 json:"stored,omitempty"`
	Required            bool                   `yaml:"required,omitempty"               json:"required,omitempty"`
	ReadOnly            bool                   `yaml:"read_only,omitempty"              json:"read_only,omitempty"`
	CreationOnly        bool                   `yaml:"creation_only,omitempty"          json:"creation_only,omitempty"`
	AllowedChars        string                 `yaml:"allowed_chars,omitempty"          json:"allowed_chars,omitempty"`
	AllowedCharsMessage string                 `yaml:"allowed_chars_message,omitempty"  json:"allowed_chars_message,omitempty"`
	AllowedChoices      []string               `yaml:"allowed_choices,omitempty"        json:"allowed_choices,omitempty"`
	Autogenerated       bool                   `yaml:"autogenerated,omitempty"          json:"autogenerated,omitempty"`
	DefaultOrder        bool                   `yaml:"default_order,omitempty"          json:"default_order,omitempty"`
	DefaultValue        interface{}            `yaml:"default_value,omitempty"          json:"default_value,omitempty"`
	Deprecated          bool                   `yaml:"deprecated,omitempty"             json:"deprecated,omitempty"`
	ExampleValue        interface{}            `yaml:"example_value,omitempty"          json:"example_value,omitempty"`
	Filterable          bool                   `yaml:"filterable,omitempty"             json:"filterable,omitempty"`
	ForeignKey          bool                   `yaml:"foreign_key,omitempty"            json:"foreign_key,omitempty"`
	Getter              bool                   `yaml:"getter,omitempty"                 json:"getter,omitempty"`
	Setter              bool                   `yaml:"setter,omitempty"                 json:"setter,omitempty"`
	Identifier          bool                   `yaml:"identifier,omitempty"             json:"identifier,omitempty"`
	MaxLength           uint16                 `yaml:"max_length,omitempty"             json:"max_length,omitempty"`
	MinLength           uint16                 `yaml:"min_length,omitempty"             json:"min_length,omitempty"`
	MaxValue            float64                `yaml:"max_value,omitempty"              json:"max_value,omitempty"`
	MinValue            float64                `yaml:"min_value,omitempty"              json:"min_value,omitempty"`
	Orderable           bool                   `yaml:"orderable,omitempty"              json:"orderable,omitempty"`
	PrimaryKey          bool                   `yaml:"primary_key,omitempty"            json:"primary_key,omitempty"`
	Secret              bool                   `yaml:"secret,omitempty"                 json:"secret,omitempty"`
	Transient           bool                   `yaml:"transient,omitempty"              json:"transient,omitempty"`
	OmitEmpty           bool                   `yaml:"omit_empty,omitempty"             json:"omit_empty,omitempty"`
	Validations         []string               `yaml:"validations,omitempty"            json:"validations,omitempty"`
	Extensions          map[string]interface{} `yaml:"extensions,omitempty"             json:"extensions,omitempty"`

	ConvertedName       string                    `yaml:"-" json:"-"`
	ConvertedType       string                    `yaml:"-" json:"-"`
	TypeProvider        string                    `yaml:"-" json:"-"`
	Initializer         string                    `yaml:"-" json:"-"`
	ValidationProviders map[string]*ValidationMap `yaml:"-" json:"-"`

	linkedSpecification Specification
}

// Validate validates the attribute definition.
func (a *Attribute) Validate() []error {

	var errs []error

	if a.Required && a.DefaultValue == nil && a.ExampleValue == nil {
		errs = append(errs, fmt.Errorf("%s.spec: '%s' is required but has no default_value or example_value", a.linkedSpecification.Model().RestName, a.Name))
	}

	if a.Description != "" && a.Description[len(a.Description)-1] != '.' && a.linkedSpecification != nil && a.linkedSpecification.Model() != nil {
		errs = append(errs, fmt.Errorf("%s.spec: description of attribute '%s' must end with a period", a.linkedSpecification.Model().RestName, a.Name))
	}

	if a.Type == AttributeTypeEnum && len(a.AllowedChoices) == 0 {
		errs = append(errs, fmt.Errorf("%s.spec: enum attribute '%s' must define allowed_choices", a.linkedSpecification.Model().RestName, a.Name))
	}

	if a.AllowedChars != "" && a.AllowedCharsMessage == "" && a.linkedSpecification != nil && a.linkedSpecification.Model() != nil {
		errs = append(errs, fmt.Errorf("%s.spec: attribute '%s' must define allowed_chars_message", a.linkedSpecification.Model().RestName, a.Name))
	}

	return errs
}
