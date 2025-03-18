package report

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/genrep/internal/core"
)

func TestTemplateReport(t *testing.T) {
	// Create a temp directory for testing
	tempDir := t.TempDir()
	
	// Create a simple template file
	templatePath := filepath.Join(tempDir, "template.txt")
	templateContent := "Hello, {{.name}}!\n\nData from source: {{.test_source}}"
	err := os.WriteFile(templatePath, []byte(templateContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create template file: %v", err)
	}
	
	// Create output path
	outputPath := filepath.Join(tempDir, "output.txt")
	
	// Create test config
	config := &core.Config{
		ReportType:   "template",
		OutputPath:   outputPath,
		OutputFormat: "txt",
		Parameters: map[string]interface{}{
			"template_path": templatePath,
			"name":          "World",
		},
		DataSources: []core.DataSourceConfig{
			{
				ID:     "test_source",
				Type:   "mock",
				Config: map[string]interface{}{},
			},
		},
	}
	
	// Create a mock data source
	mockDataSource := &MockDataSource{
		data: "Test Data",
	}
	
	dataSources := map[string]core.DataSource{
		"test_source": mockDataSource,
	}
	
	// Create and initialize the report generator
	report := NewTemplateReport()
	err = report.Initialize(config)
	if err != nil {
		t.Fatalf("Failed to initialize template report: %v", err)
	}
	
	// Generate the report
	ctx := context.Background()
	err = report.Generate(ctx, dataSources)
	if err != nil {
		t.Fatalf("Failed to generate template report: %v", err)
	}
	
	// Verify the output file exists
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Errorf("Expected output file %s to exist", outputPath)
	}
	
	// Verify the content of the output file
	content, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}
	
	expectedContent := "Hello, World!\n\nData from source: Test Data"
	if string(content) != expectedContent {
		t.Errorf("Output content doesn't match.\nExpected: %s\nGot: %s", expectedContent, string(content))
	}
}

// MockDataSource is a mock implementation of DataSource for testing
type MockDataSource struct {
	data interface{}
}

func (ds *MockDataSource) Initialize(config map[string]interface{}) error {
	return nil
}

func (ds *MockDataSource) FetchData(ctx context.Context, query interface{}) (interface{}, error) {
	return ds.data, nil
}

func (ds *MockDataSource) Close() error {
	return nil
}
