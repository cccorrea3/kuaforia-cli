package format

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestJSONOutput(t *testing.T) {
	f := New(FormatJSON)

	data := map[string]any{"id": 1, "title": "test", "status": "draft"}

	result := captureOutput(func() {
		f.JSON(data)
	})

	var parsed map[string]any
	if err := json.Unmarshal([]byte(result), &parsed); err != nil {
		t.Fatalf("invalid JSON output: %v", err)
	}
	if parsed["title"] != "test" {
		t.Errorf("expected title=test, got %v", parsed["title"])
	}
}

func TestYAMLOutput(t *testing.T) {
	f := New(FormatYAML)

	data := map[string]any{"id": 1, "title": "test", "status": "draft"}

	result := captureOutput(func() {
		f.YAML(data)
	})

	var parsed map[string]any
	if err := yaml.Unmarshal([]byte(result), &parsed); err != nil {
		t.Fatalf("invalid YAML output: %v", err)
	}
	if parsed["title"] != "test" {
		t.Errorf("expected title=test, got %v", parsed["title"])
	}
}

func TestTableOutput(t *testing.T) {
	f := New(FormatTable)

	f.Table([]string{"id", "title"}, [][]string{{"1", "test"}})
}

func captureOutput(fn func()) string {
	r, w, _ := os.Pipe()
	stdout := os.Stdout
	os.Stdout = w

	fn()

	w.Close()
	os.Stdout = stdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
