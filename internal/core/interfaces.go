package core

import "context"

// DataSourceConfig represents the configuration for a data source
type DataSourceConfig struct {
	ID     string                 `json:"id"`
	Type   string                 `json:"type"`
	Config map[string]interface{} `json:"config"`
}

// DataSource is an interface for accessing different types of data sources
type DataSource interface {
	// Initialize sets up the data source with the provided configuration
	Initialize(config map[string]interface{}) error
	
	// FetchData retrieves data from the source based on the provided query
	FetchData(ctx context.Context, query interface{}) (interface{}, error)
	
	// Close cleans up any resources used by the data source
	Close() error
	
	// Fetch retrieves data from the source (simplified interface for compatibility)
	Fetch() (interface{}, error)
}

// ReportGenerator is an interface for different report generation strategies
type ReportGenerator interface {
	// Initialize sets up the report generator with the provided configuration
	Initialize(config *Config) error
	
	// Generate creates the report using data from the provided data sources
	Generate(ctx context.Context, dataSources map[string]DataSource) error
}

// OutputFormatter is an interface for converting report data to different output formats
type OutputFormatter interface {
	// Format converts the report data to the desired output format and saves it
	Format(data interface{}, outputPath string) error
}

// Plugin is an interface for dynamically loadable components
type Plugin interface {
	// ID returns the unique identifier for the plugin
	ID() string
	
	// Type returns the type of the plugin (e.g., "report", "dataSource", "formatter")
	Type() string
	
	// Name returns the unique name of the plugin
	Name() string
	
	// Instance returns the actual implementation of the plugin
	Instance() interface{}
	
	// Initialize sets up the plugin with the provided configuration
	Initialize(config map[string]interface{}) error
}
