package report

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/genrep/internal/core"
)

func TestHelloWordReport(t *testing.T) {
	// Create a temp directory for output
	tempDir := t.TempDir()
	outputPath := filepath.Join(tempDir, "hello.docx")

	// Create test config
	config := &core.Config{
		ReportType:   "hello-word",
		OutputPath:   outputPath,
		OutputFormat: "docx",
		Parameters:   map[string]interface{}{},
		DataSources:  []core.DataSourceConfig{},
	}

	// Create and initialize the report generator
	report := NewHelloWordReport()
	err := report.Initialize(config)
	if err != nil {
		t.Fatalf("Failed to initialize report: %v", err)
	}

	// Generate the report
	ctx := context.Background()
	dataSources := make(map[string]core.DataSource)
	err = report.Generate(ctx, dataSources)
	if err != nil {
		t.Fatalf("Failed to generate report: %v", err)
	}

	// Verify the report was created
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Errorf("Expected report file %s to exist", outputPath)
	}

	// Verify the file has content
	content, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to read report file: %v", err)
	}
	if len(content) == 0 {
		t.Error("Expected non-empty report file")
	}
}
