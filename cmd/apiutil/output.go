package apiutil

import (
	"encoding/json"
	"fmt"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// OutputMode represents the --output/-o flag value.
type OutputMode string

const (
	// OutputJSON is the default output format: pretty-printed JSON.
	OutputJSON OutputMode = "json"
	// OutputYAML serialises the response as YAML.
	OutputYAML OutputMode = "yaml"
	// OutputTable renders results in a tab-separated table.
	OutputTable OutputMode = "table"
)

// AddOutputFlag registers the --output/-o flag on cmd.
func AddOutputFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("output", "o", "json", "Output format: json, yaml, table")
}

// GetOutputMode reads the --output flag from cmd and returns the corresponding
// OutputMode. Defaults to OutputJSON for unknown values.
func GetOutputMode(cmd *cobra.Command) OutputMode {
	mode, _ := cmd.Flags().GetString("output")
	switch OutputMode(mode) {
	case OutputYAML:
		return OutputYAML
	case OutputTable:
		return OutputTable
	default:
		return OutputJSON
	}
}

// PrintOutput serialises v according to mode and writes it to cmd's output.
func PrintOutput(cmd *cobra.Command, v interface{}, mode OutputMode) error {
	switch mode {
	case OutputYAML:
		out, err := yaml.Marshal(v)
		if err != nil {
			return fmt.Errorf("failed to marshal as YAML: %w", err)
		}
		cmd.Print(string(out))
		return nil
	case OutputTable:
		return printTable(cmd, v)
	default:
		out, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal as JSON: %w", err)
		}
		cmd.Println(string(out))
		return nil
	}
}

// PrintRawOutput parses data as JSON then re-renders it according to mode.
func PrintRawOutput(cmd *cobra.Command, data []byte, mode OutputMode) error {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return fmt.Errorf("failed to parse response: %w", err)
	}
	return PrintOutput(cmd, v, mode)
}

// printTable renders v in a tab-separated table. When v contains a "results"
// key, each element of that slice becomes a row. For other objects the key/value
// pairs are printed as a two-column table. Falls back to YAML for complex nested
// structures that cannot be rendered as a flat table.
func printTable(cmd *cobra.Command, v interface{}) error {
	// Normalise to map via JSON round-trip.
	raw, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("failed to serialise for table: %w", err)
	}

	var top map[string]interface{}
	if err := json.Unmarshal(raw, &top); err != nil {
		// Not a JSON object — fall back to YAML.
		return PrintOutput(cmd, v, OutputYAML)
	}

	// If the response wraps a results array, render that.
	if results, ok := top["results"].([]interface{}); ok {
		return renderResultsTable(cmd, results)
	}

	// Otherwise render the object itself as key→value rows.
	return renderObjectTable(cmd, top)
}

// renderResultsTable prints a table where each row is one element from results.
func renderResultsTable(cmd *cobra.Command, results []interface{}) error {
	if len(results) == 0 {
		cmd.Println("(no results)")
		return nil
	}

	// Collect a stable column order from the first element.
	first, ok := results[0].(map[string]interface{})
	if !ok {
		return PrintOutput(cmd, results, OutputYAML)
	}
	cols := tableColumns(first)

	tw := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 2, ' ', 0)
	fmt.Fprintln(tw, strings.Join(cols, "\t"))
	for _, item := range results {
		row, ok := item.(map[string]interface{})
		if !ok {
			continue
		}
		vals := make([]string, len(cols))
		for i, col := range cols {
			vals[i] = fmt.Sprintf("%v", row[col])
		}
		fmt.Fprintln(tw, strings.Join(vals, "\t"))
	}
	return tw.Flush()
}

// renderObjectTable prints a two-column key/value table for a single object.
func renderObjectTable(cmd *cobra.Command, obj map[string]interface{}) error {
	tw := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 2, ' ', 0)
	fmt.Fprintln(tw, "KEY\tVALUE")
	for k, v := range obj {
		fmt.Fprintf(tw, "%s\t%v\n", k, v)
	}
	return tw.Flush()
}

// tableColumns returns a deterministic column order for a result row. The ID
// column always comes first when present; then mediaType; then title/name;
// then all remaining string/number fields in alphabetical order.
func tableColumns(row map[string]interface{}) []string {
	priority := []string{"id", "mediaType", "title", "name"}
	seen := map[string]bool{}
	var cols []string

	for _, key := range priority {
		if _, ok := row[key]; ok {
			cols = append(cols, key)
			seen[key] = true
		}
	}
	// Add remaining scalar fields (skip nested objects/arrays).
	for k, v := range row {
		if seen[k] {
			continue
		}
		switch v.(type) {
		case map[string]interface{}, []interface{}:
			continue
		}
		cols = append(cols, k)
	}
	return cols
}
