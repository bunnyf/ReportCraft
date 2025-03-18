package core

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// ReportEngine is the core component that orchestrates the report generation process
type ReportEngine struct {
	reportGenerators map[string]ReportGenerator
	dataSources      map[string]func() DataSource
	outputFormatters map[string]OutputFormatter
	pluginsMutex     sync.RWMutex
}

// NewReportEngine creates a new instance of the report engine
func NewReportEngine() *ReportEngine {
	engine := &ReportEngine{
		reportGenerators: make(map[string]ReportGenerator),
		dataSources:      make(map[string]func() DataSource),
		outputFormatters: make(map[string]OutputFormatter),
	}
	
	// Register built-in components
	engine.registerBuiltinComponents()
	
	return engine
}

// registerBuiltinComponents registers the built-in report generators, data sources, and formatters
func (e *ReportEngine) registerBuiltinComponents() {
	// This would be implemented to register all built-in components
	// For example:
	// e.RegisterDataSource("db", func() DataSource { return &DatabaseSource{} })
}

// RegisterReportGenerator registers a new report generator
func (e *ReportEngine) RegisterReportGenerator(name string, generator ReportGenerator) {
	e.pluginsMutex.Lock()
	defer e.pluginsMutex.Unlock()
	e.reportGenerators[name] = generator
}

// RegisterDataSource registers a factory function for creating data sources
func (e *ReportEngine) RegisterDataSource(name string, factory func() DataSource) {
	e.pluginsMutex.Lock()
	defer e.pluginsMutex.Unlock()
	e.dataSources[name] = factory
}

// RegisterOutputFormatter registers a new output formatter
func (e *ReportEngine) RegisterOutputFormatter(name string, formatter OutputFormatter) {
	e.pluginsMutex.Lock()
	defer e.pluginsMutex.Unlock()
	e.outputFormatters[name] = formatter
}

// LoadPlugins loads plugins from the specified directory
func (e *ReportEngine) LoadPlugins(pluginDir string) error {
	// This would be implemented to load external plugins
	// using Go's plugin package or a custom plugin system
	return nil
}

// GenerateReport generates a report based on the provided configuration file
func (e *ReportEngine) GenerateReport(configFile string) error {
	// Read and parse the configuration file
	config, err := readConfig(configFile)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}
	
	// Validate the configuration
	if err := validateConfig(config); err != nil {
		return fmt.Errorf("invalid configuration: %w", err)
	}
	
	// Get the report generator
	generator, ok := e.getReportGenerator(config.ReportType)
	if !ok {
		return fmt.Errorf("unsupported report type: %s", config.ReportType)
	}
	
	// Initialize the report generator
	if err := generator.Initialize(config); err != nil {
		return fmt.Errorf("failed to initialize report generator: %w", err)
	}
	
	// Create and initialize data sources
	dataSources, err := e.createDataSources(config.DataSources)
	if err != nil {
		return fmt.Errorf("failed to create data sources: %w", err)
	}
	defer closeDataSources(dataSources)
	
	// Generate the report
	ctx := context.Background()
	if err := generator.Generate(ctx, dataSources); err != nil {
		return fmt.Errorf("failed to generate report: %w", err)
	}
	
	return nil
}

// readConfig reads and parses the configuration file
func readConfig(configFile string) (*Config, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}
	
	return &config, nil
}

// validateConfig validates the configuration
func validateConfig(config *Config) error {
	if config.ReportType == "" {
		return fmt.Errorf("reportType is required")
	}
	if config.OutputPath == "" {
		return fmt.Errorf("outputPath is required")
	}
	if config.OutputFormat == "" {
		return fmt.Errorf("outputFormat is required")
	}
	
	// Ensure the output directory exists
	outputDir := filepath.Dir(config.OutputPath)
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}
	}
	
	return nil
}

// getReportGenerator returns the report generator for the specified type
func (e *ReportEngine) getReportGenerator(reportType string) (ReportGenerator, bool) {
	e.pluginsMutex.RLock()
	defer e.pluginsMutex.RUnlock()
	generator, ok := e.reportGenerators[reportType]
	return generator, ok
}

// createDataSources creates and initializes the data sources based on the configuration
func (e *ReportEngine) createDataSources(configs []DataSourceConfig) (map[string]DataSource, error) {
	dataSources := make(map[string]DataSource)
	
	for _, config := range configs {
		factory, ok := e.getDataSourceFactory(config.Type)
		if !ok {
			return nil, fmt.Errorf("unsupported data source type: %s", config.Type)
		}
		
		dataSource := factory()
		if err := dataSource.Initialize(config.Config); err != nil {
			return nil, fmt.Errorf("failed to initialize data source %s: %w", config.ID, err)
		}
		
		dataSources[config.ID] = dataSource
	}
	
	return dataSources, nil
}

// getDataSourceFactory returns the factory function for the specified data source type
func (e *ReportEngine) getDataSourceFactory(dataSourceType string) (func() DataSource, bool) {
	e.pluginsMutex.RLock()
	defer e.pluginsMutex.RUnlock()
	factory, ok := e.dataSources[dataSourceType]
	return factory, ok
}

// closeDataSources closes all the data sources
func closeDataSources(dataSources map[string]DataSource) {
	for _, dataSource := range dataSources {
		_ = dataSource.Close()
	}
}
