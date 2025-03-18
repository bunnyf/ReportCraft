package core

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// MockDataSource is a data source implementation for testing
type MockDataSource struct {
	InitializeFunc func(config map[string]interface{}) error
	FetchDataFunc  func(ctx context.Context, query interface{}) (interface{}, error)
	FetchFunc      func() (interface{}, error)
	CloseFunc      func() error
}

func (ds *MockDataSource) Initialize(config map[string]interface{}) error {
	if ds.InitializeFunc != nil {
		return ds.InitializeFunc(config)
	}
	return nil
}

func (ds *MockDataSource) FetchData(ctx context.Context, query interface{}) (interface{}, error) {
	if ds.FetchDataFunc != nil {
		return ds.FetchDataFunc(ctx, query)
	}
	return "test-data", nil
}

func (ds *MockDataSource) Fetch() (interface{}, error) {
	if ds.FetchFunc != nil {
		return ds.FetchFunc()
	}
	return ds.FetchData(context.Background(), nil)
}

func (ds *MockDataSource) Close() error {
	if ds.CloseFunc != nil {
		return ds.CloseFunc()
	}
	return nil
}

// MockReportGenerator is a report generator implementation for testing
type MockReportGenerator struct {
	InitializeFunc func(config *Config) error
	GenerateFunc   func(ctx context.Context, dataSources map[string]DataSource) error
}

func (rg *MockReportGenerator) Initialize(config *Config) error {
	if rg.InitializeFunc != nil {
		return rg.InitializeFunc(config)
	}
	return nil
}

func (rg *MockReportGenerator) Generate(ctx context.Context, dataSources map[string]DataSource) error {
	if rg.GenerateFunc != nil {
		return rg.GenerateFunc(ctx, dataSources)
	}
	return nil
}

// TestNewReportEngine tests the creation of a new report engine
func TestNewReportEngine(t *testing.T) {
	engine := NewReportEngine()
	if engine == nil {
		t.Fatal("Expected a non-nil report engine")
	}
	if engine.reportGenerators == nil {
		t.Error("Expected report generators map to be initialized")
	}
	if engine.dataSources == nil {
		t.Error("Expected data sources map to be initialized")
	}
	if engine.outputFormatters == nil {
		t.Error("Expected output formatters map to be initialized")
	}
}

// TestRegisterReportGenerator tests the registration of a report generator
func TestRegisterReportGenerator(t *testing.T) {
	engine := NewReportEngine()
	generator := &MockReportGenerator{}
	engine.RegisterReportGenerator("test", generator)

	if len(engine.reportGenerators) != 1 {
		t.Errorf("Expected 1 report generator, got %d", len(engine.reportGenerators))
	}
	if _, ok := engine.reportGenerators["test"]; !ok {
		t.Error("Expected 'test' report generator to be registered")
	}
}

// TestRegisterDataSource tests the registration of a data source
func TestRegisterDataSource(t *testing.T) {
	engine := NewReportEngine()
	factory := func() DataSource { return &MockDataSource{} }
	engine.RegisterDataSource("test", factory)

	if len(engine.dataSources) != 1 {
		t.Errorf("Expected 1 data source, got %d", len(engine.dataSources))
	}
	if _, ok := engine.dataSources["test"]; !ok {
		t.Error("Expected 'test' data source to be registered")
	}
}

// TestReadConfig tests the reading and parsing of a configuration file
func TestReadConfig(t *testing.T) {
	// Create a temporary config file
	tempDir := t.TempDir()
	configFile := filepath.Join(tempDir, "config.json")

	config := Config{
		ReportType:   "test",
		OutputPath:   filepath.Join(tempDir, "output.txt"),
		OutputFormat: "text",
		Parameters:   map[string]interface{}{"param1": "value1"},
		DataSources: []DataSourceConfig{
			{
				ID:     "src1",
				Type:   "test",
				Config: map[string]interface{}{"key1": "value1"},
			},
		},
	}

	configJSON, err := json.Marshal(config)
	if err != nil {
		t.Fatalf("Failed to marshal config: %v", err)
	}

	if err := os.WriteFile(configFile, configJSON, 0644); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	// Read the config file
	readConfig, err := readConfig(configFile)
	if err != nil {
		t.Fatalf("Failed to read config file: %v", err)
	}

	// Verify the config contents
	if readConfig.ReportType != config.ReportType {
		t.Errorf("Expected report type %s, got %s", config.ReportType, readConfig.ReportType)
	}
	if readConfig.OutputPath != config.OutputPath {
		t.Errorf("Expected output path %s, got %s", config.OutputPath, readConfig.OutputPath)
	}
	if readConfig.OutputFormat != config.OutputFormat {
		t.Errorf("Expected output format %s, got %s", config.OutputFormat, readConfig.OutputFormat)
	}
	if len(readConfig.DataSources) != len(config.DataSources) {
		t.Errorf("Expected %d data sources, got %d", len(config.DataSources), len(readConfig.DataSources))
	}
}

// TestValidateConfig tests the validation of a configuration
func TestValidateConfig(t *testing.T) {
	// Valid config
	validConfig := &Config{
		ReportType:   "test",
		OutputPath:   "output.txt",
		OutputFormat: "text",
	}
	if err := validateConfig(validConfig); err != nil {
		t.Errorf("Expected no error for valid config, got %v", err)
	}

	// Invalid config - missing report type
	invalidConfig1 := &Config{
		OutputPath:   "output.txt",
		OutputFormat: "text",
	}
	if err := validateConfig(invalidConfig1); err == nil {
		t.Error("Expected error for missing report type, got nil")
	}

	// Invalid config - missing output path
	invalidConfig2 := &Config{
		ReportType:   "test",
		OutputFormat: "text",
	}
	if err := validateConfig(invalidConfig2); err == nil {
		t.Error("Expected error for missing output path, got nil")
	}

	// Invalid config - missing output format
	invalidConfig3 := &Config{
		ReportType: "test",
		OutputPath: "output.txt",
	}
	if err := validateConfig(invalidConfig3); err == nil {
		t.Error("Expected error for missing output format, got nil")
	}
}
