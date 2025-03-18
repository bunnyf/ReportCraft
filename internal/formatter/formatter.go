package formatter

import (
	"fmt"
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
	// Implementation for HTML document generation
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
