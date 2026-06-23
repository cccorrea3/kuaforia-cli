package format

import (
	"encoding/json"
	"os"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"gopkg.in/yaml.v3"
)

type FormatType string

const (
	FormatTable FormatType = "table"
	FormatJSON  FormatType = "json"
	FormatYAML  FormatType = "yaml"
)

type Formatter interface {
	Table(headers []string, rows [][]string)
	JSON(data any)
	YAML(data any)
	Error(msg string)
	Success(msg string)
	Warning(msg string)
}

type formatter struct {
	format FormatType
}

func New(f FormatType) Formatter {
	return &formatter{format: f}
}

func (f *formatter) Table(headers []string, rows [][]string) {
	switch f.format {
	case FormatJSON:
		data := make([]map[string]string, len(rows))
		for i, row := range rows {
			m := make(map[string]string)
			for j, h := range headers {
				if j < len(row) {
					m[h] = row[j]
				}
			}
			data[i] = m
		}
		f.JSON(data)
	case FormatYAML:
		data := make([]map[string]string, len(rows))
		for i, row := range rows {
			m := make(map[string]string)
			for j, h := range headers {
				if j < len(row) {
					m[h] = row[j]
				}
			}
			data[i] = m
		}
		f.YAML(data)
	default:
		table := tablewriter.NewWriter(os.Stdout)
		table.Header(headers)
		table.Bulk(rows)
		table.Render()
	}
}

func (f *formatter) JSON(data any) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(data)
}

func (f *formatter) YAML(data any) {
	enc := yaml.NewEncoder(os.Stdout)
	enc.SetIndent(2)
	enc.Encode(data)
}

func (f *formatter) Error(msg string) {
	color.Red("[ERROR] %s", msg)
}

func (f *formatter) Success(msg string) {
	color.Green("[OK] %s", msg)
}

func (f *formatter) Warning(msg string) {
	color.Yellow("[WARN] %s", msg)
}
