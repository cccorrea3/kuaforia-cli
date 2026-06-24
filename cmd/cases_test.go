package cmd

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestCasesListCmd(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/cli/list_cases" {
			t.Errorf("expected /api/v1/cli/list_cases, got %s", r.URL.Path)
		}
		w.Write([]byte(`{"success":true,"data":[{"id":1,"title":"Test","status":"draft"}]}`))
	}))
	defer server.Close()

	t.Setenv("KUAFORIA_SERVER", server.URL)
	t.Setenv("KUAFORIA_API_KEY", "test-key")
	t.Setenv("KUAFORIA_TENANT", "test")
	t.Setenv("KUAFORIA_DEFAULT_WORKSPACE", "1")

	rootCmd.SetArgs([]string{"cases", "list"})
	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestAuthTestCmd(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/cli/health" {
			t.Errorf("expected /api/v1/cli/health, got %s", r.URL.Path)
		}
		w.Write([]byte(`{"success":true,"data":{"status":"ok"}}`))
	}))
	defer server.Close()

	t.Setenv("KUAFORIA_SERVER", server.URL)
	t.Setenv("KUAFORIA_API_KEY", "test-key")
	t.Setenv("KUAFORIA_TENANT", "test")

	rootCmd.SetArgs([]string{"auth", "test"})
	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestJSONOutputFormat(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"success":true,"data":[{"id":1,"title":"T","status":"draft"}]}`))
	}))
	defer server.Close()

	t.Setenv("KUAFORIA_SERVER", server.URL)
	t.Setenv("KUAFORIA_API_KEY", "test-key")
	t.Setenv("KUAFORIA_TENANT", "test")
	t.Setenv("KUAFORIA_DEFAULT_WORKSPACE", "1")

	r, w, _ := os.Pipe()
	stdout := os.Stdout
	os.Stdout = w

	rootCmd.SetArgs([]string{"cases", "list", "-o", "json"})
	rootCmd.Execute()

	w.Close()
	os.Stdout = stdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// JSON output may have a trailing message line; extract only the JSON part
	jsonEnd := strings.LastIndex(output, "Total:")
	if jsonEnd > 0 {
		output = output[:jsonEnd]
	}
	output = strings.TrimSpace(output)
	var parsed any
	if err := json.Unmarshal([]byte(output), &parsed); err != nil {
		t.Fatalf("invalid JSON from command: %v\noutput: %s", err, output)
	}
}
