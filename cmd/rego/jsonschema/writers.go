package jsonschema

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"text/template"

	"go.aporeto.io/regolithe/spec"
)

var functions = template.FuncMap{
	"convertType":             convertType,
	"convertRegexp":           convertRegexp,
	"jsonStringify":           jsonStringify,
	"isNil":                   isNil,
	"stripFirstLevelBrackets": stripFirstLevelBrackets,
}

func writeGlobalResources(set spec.SpecificationSet, outFolder string) error {

	if err := os.MkdirAll(outFolder, 0755); err != nil && !os.IsExist(err) {
		return err
	}

	tmpl, err := makeTemplate("templates/json-schema-restname.gotpl")
	if err != nil {
		return err
	}

	var buf bytes.Buffer

	if err = tmpl.Execute(
		&buf,
		struct {
			Name string
			Set  spec.SpecificationSet
		}{
			Name: set.Configuration().Name,
			Set:  set,
		}); err != nil {
		return fmt.Errorf("Unable to generate global resource code: %s", err)
	}

	data := map[string]interface{}{}
	if err := json.Unmarshal(buf.Bytes(), &data); err != nil {
		return fmt.Errorf("Unable to unmarshal model code: %s", err)
	}

	out, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("Unable to marshal model code: %s", err)
	}

	return writeFile(path.Join(outFolder, "_models.json"), out)
}

func writeGlobalResourceLists(set spec.SpecificationSet, outFolder string) error {

	if err := os.MkdirAll(outFolder, 0755); err != nil && !os.IsExist(err) {
		return err
	}

	tmpl, err := makeTemplate("templates/json-schema-resourcename.gotpl")
	if err != nil {
		return err
	}

	var buf bytes.Buffer

	if err = tmpl.Execute(
		&buf,
		struct {
			Name string
			Set  spec.SpecificationSet
		}{
			Name: set.Configuration().Name,
			Set:  set,
		}); err != nil {
		return fmt.Errorf("Unable to generate global resource lists code: %s", err)
	}

	data := map[string]interface{}{}
	if err := json.Unmarshal(buf.Bytes(), &data); err != nil {
		return fmt.Errorf("Unable to unmarshal model code: %s", err)
	}

	out, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("Unable to marshal model code: %s", err)
	}

	return writeFile(path.Join(outFolder, "_lists.json"), out)
}

func writeModel(set spec.SpecificationSet, outFolder string) error {

	if err := os.MkdirAll(outFolder, 0755); err != nil && !os.IsExist(err) {
		return err
	}

	tmpl, err := makeTemplate("templates/json-schema.gotpl")
	if err != nil {
		return err
	}

	for _, s := range set.Specifications() {
		var buf bytes.Buffer

		if err = tmpl.Execute(
			&buf,
			struct {
				Name string
				Spec spec.Specification
			}{
				Name: set.Configuration().Name,
				Spec: s,
			}); err != nil {
			return fmt.Errorf("Unable to generate model code: %s", err)
		}

		data := map[string]interface{}{}
		if err := json.Unmarshal(buf.Bytes(), &data); err != nil {
			return fmt.Errorf("Unable to unmarshal model code: %s", err)
		}

		out, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			return fmt.Errorf("Unable to marshal model code: %s", err)
		}

		if err := writeFile(path.Join(outFolder, s.Model().RestName+".json"), out); err != nil {
			return err
		}
	}

	return nil
}
