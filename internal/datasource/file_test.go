package datasource

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestFileDataSourceInitialize(t *testing.T) {
	ds := NewFileDataSource()

	// Test with valid config
	validConfig := map[string]interface{}{
		"path": "test.json",
	}
	if err := ds.Initialize(validConfig); err != nil {
		t.Errorf("Expected no error for valid config, got %v", err)
	}
	if ds.path != "test.json" {
		t.Errorf("Expected path to be 'test.json', got '%s'", ds.path)
	}
	if ds.fileType != "json" {
		t.Errorf("Expected file type to be 'json', got '%s'", ds.fileType)
	}

	// Test with missing path
	invalidConfig := map[string]interface{}{}
	if err := ds.Initialize(invalidConfig); err == nil {
		t.Error("Expected error for missing path, got nil")
	}

	// Test with invalid path type
	invalidConfig2 := map[string]interface{}{
		"path": 123,
	}
	if err := ds.Initialize(invalidConfig2); err == nil {
		t.Error("Expected error for invalid path type, got nil")
	}

	// Test with unsupported file type
	invalidConfig3 := map[string]interface{}{
		"path": "test.xyz",
	}
	if err := ds.Initialize(invalidConfig3); err == nil {
		t.Error("Expected error for unsupported file type, got nil")
	}
}

func TestFileDataSourceFetchData(t *testing.T) {
	ds := NewFileDataSource()
	ctx := context.Background()

	// Create a temporary directory for test files
	tempDir := t.TempDir()

	// Test with JSON file
	jsonFile := filepath.Join(tempDir, "test.json")
	jsonContent := `{"key": "value"}`
	if err := os.WriteFile(jsonFile, []byte(jsonContent), 0644); err != nil {
		t.Fatalf("Failed to write test JSON file: %v", err)
	}

	// Initialize with JSON file
	ds.Initialize(map[string]interface{}{"path": jsonFile})

	// Fetch data from the JSON file
	data, err := ds.FetchData(ctx, nil)
	if err != nil {
		t.Fatalf("Failed to fetch data from JSON file: %v", err)
	}

	// Verify the data
	jsonData, ok := data.(map[string]interface{})
	if !ok {
		t.Fatal("Expected a map from JSON data")
	}
	if value, ok := jsonData["key"]; !ok || value != "value" {
		t.Errorf("Expected 'key' to be 'value', got %v", value)
	}

	// Test with a non-existent file
	ds.Initialize(map[string]interface{}{"path": "nonexistent.json"})
	_, err = ds.FetchData(ctx, nil)
	if err == nil {
		t.Error("Expected error for non-existent file, got nil")
	}
}

func TestFileDataSourceClose(t *testing.T) {
	ds := NewFileDataSource()
	err := ds.Close()
	if err != nil {
		t.Errorf("Expected no error from Close, got %v", err)
	}
}
