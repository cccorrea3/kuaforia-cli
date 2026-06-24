package schema

import "fmt"

var ValidPriorities = []string{"low", "medium", "high", "critical"}
var ValidStatuses = []string{"draft", "needs_review", "active", "archived", "deprecated"}
var RequiredFields = []string{"title", "description"}

type ValidationError struct {
	Field   string
	Message string
}

func Validate(data map[string]any) []ValidationError {
	var errs []ValidationError

	for _, f := range RequiredFields {
		v, ok := data[f]
		if !ok || v == nil || fmt.Sprint(v) == "" {
			errs = append(errs, ValidationError{Field: f, Message: "campo requerido"})
		}
	}

	if p, ok := data["priority"].(string); ok && p != "" {
		if !contains(ValidPriorities, p) {
			errs = append(errs, ValidationError{Field: "priority", Message: "debe ser uno de: " + join(ValidPriorities)})
		}
	}

	if s, ok := data["status"].(string); ok && s != "" {
		if !contains(ValidStatuses, s) {
			errs = append(errs, ValidationError{Field: "status", Message: "debe ser uno de: " + join(ValidStatuses)})
		}
	}

	if tags, ok := data["tags"].([]any); ok {
		for i, t := range tags {
			if _, ok := t.(string); !ok {
				errs = append(errs, ValidationError{Field: fmt.Sprintf("tags[%d]", i), Message: "debe ser string"})
			}
		}
	}

	if steps, ok := data["steps"].([]any); ok {
		for i, s := range steps {
			step, ok := s.(map[string]any)
			if !ok {
				errs = append(errs, ValidationError{Field: fmt.Sprintf("steps[%d]", i), Message: "debe ser objeto"})
				continue
			}
			if desc, ok := step["description"]; !ok || desc == nil || fmt.Sprint(desc) == "" {
				errs = append(errs, ValidationError{Field: fmt.Sprintf("steps[%d].description", i), Message: "campo requerido"})
			}
		}
	}

	if deps, ok := data["dependencies"].([]any); ok {
		for i, d := range deps {
			dep, ok := d.(map[string]any)
			if !ok {
				errs = append(errs, ValidationError{Field: fmt.Sprintf("dependencies[%d]", i), Message: "debe ser objeto"})
				continue
			}
			for _, f := range []string{"case", "type"} {
				if v, ok := dep[f]; !ok || v == nil || fmt.Sprint(v) == "" {
					errs = append(errs, ValidationError{Field: fmt.Sprintf("dependencies[%d].%s", i, f), Message: "campo requerido"})
				}
			}
		}
	}

	if tests, ok := data["tests"].([]any); ok {
		for i, t := range tests {
			test, ok := t.(map[string]any)
			if !ok {
				errs = append(errs, ValidationError{Field: fmt.Sprintf("tests[%d]", i), Message: "debe ser objeto"})
				continue
			}
			for _, f := range []string{"type", "definition"} {
				if v, ok := test[f]; !ok || v == nil || fmt.Sprint(v) == "" {
					errs = append(errs, ValidationError{Field: fmt.Sprintf("tests[%d].%s", i, f), Message: "campo requerido"})
				}
			}
		}
	}

	return errs
}

func contains(list []string, s string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}

func join(list []string) string {
	s := ""
	for i, v := range list {
		if i > 0 {
			s += ", "
		}
		s += v
	}
	return s
}
