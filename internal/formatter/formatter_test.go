package formatter

import (
	"testing"
)

func TestCreateFormatter(t *testing.T) {
	// Test creating a DOCX formatter
	docxFormatter, err := CreateFormatter("docx")
	if err != nil {
		t.Errorf("Expected no error for DOCX formatter, got %v", err)
	}
	if docxFormatter == nil {
		t.Error("Expected non-nil DOCX formatter")
	}

	// Test creating an Excel formatter
	excelFormatter, err := CreateFormatter("xlsx")
	if err != nil {
		t.Errorf("Expected no error for Excel formatter, got %v", err)
	}
	if excelFormatter == nil {
		t.Error("Expected non-nil Excel formatter")
	}

	// Test with unsupported format
	_, err = CreateFormatter("unknown")
	if err == nil {
		t.Error("Expected error for unsupported format, got nil")
	}
}

func TestDocxFormatter(t *testing.T) {
	formatter := &DocxFormatter{}
	data := map[string]interface{}{
		"title": "Test Report",
		"content": "Test Content",
	}
	
	// Since the actual implementation is a placeholder,
	// we just check that it doesn't panic
	err := formatter.Format(data, "test.docx")
	if err != nil {
		t.Errorf("Unexpected error from DocxFormatter.Format: %v", err)
	}
}

func TestExcelFormatter(t *testing.T) {
	formatter := &ExcelFormatter{}
	data := map[string]interface{}{
		"title": "Test Report",
		"content": "Test Content",
	}
	
	// Since the actual implementation is a placeholder,
	// we just check that it doesn't panic
	err := formatter.Format(data, "test.xlsx")
	if err != nil {
		t.Errorf("Unexpected error from ExcelFormatter.Format: %v", err)
	}
}
