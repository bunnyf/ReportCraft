package formatter

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

// OutputFormatter formats data into a specific output format
type OutputFormatter interface {
	Format(data interface{}, outputPath string) error
}

// DocxFormatter formats data into a Word document
type DocxFormatter struct{}

// Format formats the data into a Word document
func (f *DocxFormatter) Format(data interface{}, outputPath string) error {
	// Implementation for Word document generation
	return nil
}

// ExcelFormatter formats data into an Excel spreadsheet
type ExcelFormatter struct{}

// Format formats the data into an Excel spreadsheet
func (f *ExcelFormatter) Format(data interface{}, outputPath string) error {
	// Implementation for Excel spreadsheet generation
	return nil
}

// PDFFormatter formats data into a PDF document
type PDFFormatter struct{}

// Format formats the data into a PDF document
func (f *PDFFormatter) Format(data interface{}, outputPath string) error {
	// Implementation for PDF document generation
	return nil
}

// HTMLFormatter formats data into an HTML document
type HTMLFormatter struct{}

// Format formats the data into an HTML document
func (f *HTMLFormatter) Format(data interface{}, outputPath string) error {
	// Check if template path is provided in the data
	dataMap, ok := data.(map[string]interface{})
	if !ok {
		return fmt.Errorf("data is not a map")
	}

	templatePath, ok := dataMap["_templatePath"].(string)
	if !ok {
		return fmt.Errorf("template path not found in data")
	}

	// Read the template file
	templateContent, err := os.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("failed to read template file: %w", err)
	}

	// Parse the template
	tmpl, err := template.New(filepath.Base(templatePath)).Parse(string(templateContent))
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	// Create output directory if it doesn't exist
	outputDir := filepath.Dir(outputPath)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Create output file
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outputFile.Close()

	// Apply the template with the provided data
	if err := tmpl.Execute(outputFile, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	fmt.Printf("Generated HTML document at %s\n", outputPath)
	return nil
}

// JSONFormatter formats data into a JSON file
type JSONFormatter struct{}

// Format formats the data into a JSON file
func (f *JSONFormatter) Format(data interface{}, outputPath string) error {
	// Implementation for JSON file generation
	return nil
}

// CreateFormatter creates a formatter based on the output format
func CreateFormatter(format string) (OutputFormatter, error) {
	switch strings.ToLower(format) {
	case "docx", ".docx":
		return &DocxFormatter{}, nil
	case "xlsx", ".xlsx":
		return &ExcelFormatter{}, nil
	case "pdf", ".pdf":
		return &PDFFormatter{}, nil
	case "html", ".html":
		return &HTMLFormatter{}, nil
	case "json", ".json":
		return &JSONFormatter{}, nil
	default:
		return nil, fmt.Errorf("unsupported format: %s", format)
	}
}

// GetFormatFromPath extracts the format from the file path
func GetFormatFromPath(path string) string {
	ext := filepath.Ext(path)
	if ext == "" {
		return ""
	}
	return strings.ToLower(ext[1:])
}
