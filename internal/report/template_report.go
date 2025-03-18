package report

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/genrep/internal/core"
)

// TemplateReport is a base report generator that uses templates for report generation
type TemplateReport struct {
	config       *core.Config
	templatePath string
	templateData map[string]interface{}
	outputPath   string
	outputFormat string
}

// Initialize sets up the template report with the provided configuration
func (r *TemplateReport) Initialize(config *core.Config) error {
	r.config = config
	r.outputPath = config.OutputPath
	r.outputFormat = config.OutputFormat
	r.templateData = make(map[string]interface{})

	// Extract template path from parameters
	if templatePath, ok := config.Parameters["template_path"].(string); ok {
		r.templatePath = templatePath
	} else {
		return fmt.Errorf("template_path parameter is required")
	}

	// Verify template file exists
	if _, err := os.Stat(r.templatePath); os.IsNotExist(err) {
		return fmt.Errorf("template file not found: %s", r.templatePath)
	}

	fmt.Printf("Initialized template report with template: %s\n", r.templatePath)
	return nil
}

// Generate creates a report based on a template and data sources
func (r *TemplateReport) Generate(ctx context.Context, dataSources map[string]core.DataSource) error {
	// Collect data from all data sources
	for _, dsConfig := range r.config.DataSources {
		dataSource, ok := dataSources[dsConfig.ID]
		if !ok {
			return fmt.Errorf("data source not found: %s", dsConfig.ID)
		}

		// For simplicity, we'll use ID as the query for now
		// In a real implementation, the query would come from the configuration
		query := dsConfig.ID

		// Fetch data using the query
		data, err := dataSource.FetchData(ctx, query)
		if err != nil {
			return fmt.Errorf("failed to fetch data from %s: %w", dsConfig.ID, err)
		}

		// Store the data with the data source ID as key
		r.templateData[dsConfig.ID] = data
	}

	// Add parameters to template data
	for key, value := range r.config.Parameters {
		if key != "template_path" { // Skip template_path as it's already processed
			r.templateData[key] = value
		}
	}

	// Create output directory if it doesn't exist
	outputDir := filepath.Dir(r.outputPath)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Read the template file
	templateContent, err := os.ReadFile(r.templatePath)
	if err != nil {
		return fmt.Errorf("failed to read template file: %w", err)
	}

	// Parse the template
	tmpl, err := template.New(filepath.Base(r.templatePath)).Parse(string(templateContent))
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	// Apply the template with the collected data
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, r.templateData); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	// Write the output to a file
	if err := os.WriteFile(r.outputPath, buf.Bytes(), 0644); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	fmt.Printf("Generated report at %s\n", r.outputPath)
	return nil
}

// NewTemplateReport creates a new template-based report generator
func NewTemplateReport() *TemplateReport {
	return &TemplateReport{}
}
