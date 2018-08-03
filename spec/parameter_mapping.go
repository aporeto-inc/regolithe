package spec

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"time"

	"github.com/araddon/dateparse"
	"gopkg.in/yaml.v2"
)

// ParameterType represents the various type for a parameter.
type ParameterType string

// Various values for ParameterType.
const (
	ParameterTypeString   ParameterType = "string"
	ParameterTypeInt      ParameterType = "integer"
	ParameterTypeFloat    ParameterType = "float"
	ParameterTypeBool     ParameterType = "boolean"
	ParameterTypeTime     ParameterType = "time"
	ParameterTypeEnum     ParameterType = "enum"
	ParameterTypeDuration ParameterType = "duration"
)

// A ParameterMapping is a list parameter mapping
type ParameterMapping map[string]*ParameterDefinition

// NewParameterMapping returns a new ParameterMapping.
func NewParameterMapping() ParameterMapping {
	return ParameterMapping{}
}

// LoadGlobalParameters loads the global parameters file.
func LoadGlobalParameters(path string) (ParameterMapping, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close() // nolint: errcheck

	pm := ParameterMapping{}

	if err = pm.Read(file, true); err != nil {
		return nil, err
	}

	return pm, nil
}

// Read loads a validation mapping from the given io.Reader
func (p ParameterMapping) Read(reader io.Reader, validate bool) (err error) {

	decoder := yaml.NewDecoder(reader)
	decoder.SetStrict(true)

	if err = decoder.Decode(&p); err != nil {
		return err
	}

	if validate {
		if errs := p.Validate(); len(errs) != 0 {
			return formatValidationErrors(errs)
		}
	}

	return nil
}

// Write dumps the specification into a []byte.
func (p ParameterMapping) Write(writer io.Writer) error {

	repr := yaml.MapSlice{}

	keys := make([]string, len(p))
	var i int
	for k := range p {
		keys[i] = k
		i++
	}

	sort.Strings(keys)

	for _, k := range keys {
		repr = append(repr, yaml.MapItem{
			Key:   k,
			Value: p[k],
		})
	}

	data, err := yaml.Marshal(repr)
	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	prfx1 := []byte("  - name: ")
	prfx2 := []byte("  entries:")
	lines := bytes.Split(data, []byte("\n"))
	var previousLine []byte

	for i, line := range lines {
		condFirstLine := i == 0

		if !condFirstLine &&
			(len(line) > 0 && line[0] != ' ') ||
			(bytes.HasPrefix(line, prfx1) && !bytes.HasPrefix(previousLine, prfx2)) {
			buf.WriteRune('\n')
		}

		buf.Write(line)

		if i+1 < len(lines) {
			buf.WriteRune('\n')
		}

		previousLine = line
	}

	_, err = writer.Write(buf.Bytes())
	return err
}

// Validate the ParameterMapping
func (p ParameterMapping) Validate() []error {

	var errs []error

	for _, v := range p {
		if err := v.Validate("_parameter.mapping"); err != nil {
			errs = append(errs, err...)
		}
	}

	return errs
}

// ParameterDefinition represents a parameter definition.
type ParameterDefinition struct {
	Required [][][]string `yaml:"required,omitempty"    json:"required,omitempty"`
	Entries  []*Parameter `yaml:"entries,omitempty"     json:"entries,omitempty"`
}

func (p *ParameterDefinition) extend(additional *ParameterDefinition, key string) error {

	if additional == nil {
		return fmt.Errorf("unable to find global parameter key '%s'", key)
	}

	p.Required = append(p.Required, additional.Required...)
	p.Entries = append(p.Entries, additional.Entries...)

	return nil
}

// Validate validates the parameter definition.
func (p *ParameterDefinition) Validate(relatedReSTName string) []error {

	var errs []error
	for _, p := range p.Entries {
		if err := p.Validate(relatedReSTName); err != nil {
			errs = append(errs, err...)
		}
	}

	return errs
}

// A Parameter represent one parameter that can be
// sent with a query
type Parameter struct {
	Name           string        `yaml:"name,omitempty"              json:"name,omitempty"`
	Description    string        `yaml:"description,omitempty"       json:"description,omitempty"`
	Type           ParameterType `yaml:"type,omitempty"              json:"type,omitempty"`
	Multiple       bool          `yaml:"multiple,omitempty"          json:"multiple,omitempty"`
	AllowedChoices []string      `yaml:"allowed_choices,omitempty"   json:"allowed_choices,omitempty"`
	DefaultValue   interface{}   `yaml:"default_value,omitempty"     json:"default_value,omitempty"`
	ExampleValue   interface{}   `yaml:"example_value,omitempty"     json:"example_value,omitempty"`
}

// Validate validates the parameter definition.
func (p *Parameter) Validate(relatedReSTName string) []error {

	var errs []error

	if p.Description == "" || p.Description[len(p.Description)-1] != '.' {
		errs = append(errs, fmt.Errorf("%s.spec: description of parameter '%s' must end with a period", relatedReSTName, p.Name))
	}

	if p.Type == "" {
		errs = append(errs, fmt.Errorf("%s.spec: type of parameter '%s' must be set", relatedReSTName, p.Name))
	} else {

		if p.Type != ParameterTypeString &&
			p.Type != ParameterTypeInt &&
			p.Type != ParameterTypeFloat &&
			p.Type != ParameterTypeBool &&
			p.Type != ParameterTypeTime &&
			p.Type != ParameterTypeDuration &&
			p.Type != ParameterTypeEnum {
			errs = append(errs, fmt.Errorf("%s.spec: type of parameter '%s' must be 'string', 'integer', 'float', 'boolean', 'enum', 'time' or 'duration'", relatedReSTName, p.Name))
		}
	}

	if p.Type == ParameterTypeEnum && len(p.AllowedChoices) == 0 {
		errs = append(errs, fmt.Errorf("%s.spec: enum parameter '%s' must define allowed_choices", relatedReSTName, p.Name))
	}

	if p.Type != ParameterTypeEnum && len(p.AllowedChoices) > 0 {
		errs = append(errs, fmt.Errorf("%s.spec: parameter '%s' is not an enum but defines allowed_choices", relatedReSTName, p.Name))
	}

	if p.DefaultValue == nil && p.ExampleValue == nil && p.Type == ParameterTypeString {
		errs = append(errs, fmt.Errorf("%s.spec: parameter '%s' must provide an example value as it doesn't have a default", relatedReSTName, p.Name))
	}

	if p.DefaultValue != nil {
		switch p.Type {
		case ParameterTypeString:
			if _, ok := p.DefaultValue.(string); !ok {
				errs = append(errs, fmt.Errorf("%s.spec: parameter '%s' is defined as an string, but the default value is not", relatedReSTName, p.Name))
			}
		case ParameterTypeEnum:
			if _, ok := p.DefaultValue.(string); !ok {
				errs = append(errs, fmt.Errorf("%s.spec: parameter '%s' is defined as an enum, but the default value is not", relatedReSTName, p.Name))
			}
		case ParameterTypeInt:
			if _, ok := p.DefaultValue.(int); !ok {
				errs = append(errs, fmt.Errorf("%s.spec: parameter '%s' is defined as an integer, but the default value is not", relatedReSTName, p.Name))
			}
		case ParameterTypeFloat:
			if _, ok := p.DefaultValue.(float64); !ok {
				errs = append(errs, fmt.Errorf("%s.spec: parameter '%s' is defined as an float, but the default value is not", relatedReSTName, p.Name))
			}
		case ParameterTypeBool:
			if _, ok := p.DefaultValue.(bool); !ok {
				errs = append(errs, fmt.Errorf("%s.spec: parameter '%s' is defined as an boolean, but the default value is not", relatedReSTName, p.Name))
			}
		case ParameterTypeDuration:
			if _, ok := p.DefaultValue.(string); !ok {
				errs = append(errs, fmt.Errorf("%s.spec: parameter '%s' is defined as an duration, but the default value is not", relatedReSTName, p.Name))
				break
			}
			if _, err := time.ParseDuration(p.DefaultValue.(string)); err != nil {
				errs = append(errs, fmt.Errorf("%s.spec: parameter '%s' is defined as an duration, but the default value is not", relatedReSTName, p.Name))
			}
		case ParameterTypeTime:
			if _, ok := p.DefaultValue.(string); !ok {
				errs = append(errs, fmt.Errorf("%s.spec: parameter '%s' is defined as an time, but the default value is not", relatedReSTName, p.Name))
				break
			}
			if _, err := dateparse.ParseAny(p.DefaultValue.(string)); err != nil {
				errs = append(errs, fmt.Errorf("%s.spec: parameter '%s' is defined as an time, but the default value is not", relatedReSTName, p.Name))
			}
		}
	}

	return errs
}