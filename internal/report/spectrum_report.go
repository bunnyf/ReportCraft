package report

import (
	"context"
	"fmt"
	"time"

	"github.com/genrep/internal/core"
	"github.com/genrep/internal/formatter"
)

// SpectrumReport generates a spectrum analysis report for vibration data
type SpectrumReport struct {
	config    *core.Config
	formatter core.OutputFormatter
	data      map[string]interface{}
}

// Initialize sets up the spectrum report generator with the provided configuration
func (r *SpectrumReport) Initialize(config *core.Config) error {
	r.config = config
	r.data = make(map[string]interface{})
	
	// Create the appropriate formatter based on the output format
	var err error
	r.formatter, err = formatter.CreateFormatter(config.OutputFormat)
	if err != nil {
		return fmt.Errorf("failed to create formatter: %w", err)
	}
	
	return nil
}

// Generate creates the spectrum report using data from the provided data sources
func (r *SpectrumReport) Generate(ctx context.Context, dataSources map[string]core.DataSource) error {
	// Collect data from all required data sources
	if err := r.collectData(ctx, dataSources); err != nil {
		return fmt.Errorf("failed to collect data: %w", err)
	}
	
	// Process the data for the report
	if err := r.processData(); err != nil {
		return fmt.Errorf("failed to process data: %w", err)
	}
	
	// Format and save the report
	if err := r.formatter.Format(r.data, r.config.OutputPath); err != nil {
		return fmt.Errorf("failed to format report: %w", err)
	}
	
	return nil
}

// collectData collects data from all required data sources
func (r *SpectrumReport) collectData(ctx context.Context, dataSources map[string]core.DataSource) error {
	// Extract basic report parameters
	r.data["title"] = r.config.Parameters["title"]
	r.data["reportDate"] = r.config.Parameters["reportDate"]
	r.data["description"] = r.config.Parameters["description"]
	
	// Add device information
	if deviceInfo, ok := r.config.Parameters["deviceInfo"].(map[string]interface{}); ok {
		r.data["deviceInfo"] = deviceInfo
	}
	
	// Add metadata
	if metadata, ok := r.config.Parameters["metadata"].(map[string]interface{}); ok {
		r.data["metadata"] = metadata
	}
	
	// Handle embedded data if present
	if embeddedData, ok := r.config.Parameters["embeddedData"].(map[string]interface{}); ok {
		// Extract waveform and spectrum data
		if waveformData, ok := embeddedData["waveformData"].([]interface{}); ok {
			r.data["waveformData"] = waveformData
		}
		
		if spectrumData, ok := embeddedData["spectrumData"].([]interface{}); ok {
			r.data["spectrumData"] = spectrumData
		}
		
		if waveformParams, ok := embeddedData["waveformParameters"].([]interface{}); ok {
			r.data["waveformParameters"] = waveformParams
		}
		
		if spectrumParams, ok := embeddedData["spectrumParameters"].([]interface{}); ok {
			r.data["spectrumParameters"] = spectrumParams
		}
	} else {
		// Use data from data sources
		for sourceID, source := range dataSources {
			if sourceID == "waveformData" {
				data, err := source.FetchData(ctx, nil)
				if err != nil {
					return fmt.Errorf("failed to fetch waveform data: %w", err)
				}
				r.data["waveformData"] = data
			} else if sourceID == "spectrumData" {
				data, err := source.FetchData(ctx, nil)
				if err != nil {
					return fmt.Errorf("failed to fetch spectrum data: %w", err)
				}
				r.data["spectrumData"] = data
			} else if sourceID == "waveformParameters" {
				data, err := source.FetchData(ctx, nil)
				if err != nil {
					return fmt.Errorf("failed to fetch waveform parameters: %w", err)
				}
				r.data["waveformParameters"] = data
			} else if sourceID == "spectrumParameters" {
				data, err := source.FetchData(ctx, nil)
				if err != nil {
					return fmt.Errorf("failed to fetch spectrum parameters: %w", err)
				}
				r.data["spectrumParameters"] = data
			}
		}
	}
	
	// Add chart configuration
	if waveformConfig, ok := r.config.Parameters["waveformConfig"].(map[string]interface{}); ok {
		r.data["waveformConfig"] = r.processChartConfig(waveformConfig)
	}
	
	if spectrumConfig, ok := r.config.Parameters["spectrumConfig"].(map[string]interface{}); ok {
		r.data["spectrumConfig"] = r.processChartConfig(spectrumConfig)
	}
	
	// Add tables configuration
	if tablesConfig, ok := r.config.Parameters["tablesConfig"].([]interface{}); ok {
		r.data["tablesConfig"] = tablesConfig
	}
	
	return nil
}

// processChartConfig enhances the chart configuration with the advanced style options
func (r *SpectrumReport) processChartConfig(config map[string]interface{}) map[string]interface{} {
	// Create a copy of the config to avoid modifying the original
	enhancedConfig := make(map[string]interface{})
	for k, v := range config {
		enhancedConfig[k] = v
	}
	
	// Process chart type if present
	if chartType, ok := config["chartType"].(string); ok {
		enhancedConfig["chartType"] = chartType
		
		// Add default settings based on chart type if not already specified
		if chartStyle, ok := config["chartStyle"].(map[string]interface{}); ok {
			enhancedStyle := make(map[string]interface{})
			
			// Copy basic style properties
			for k, v := range chartStyle {
				enhancedStyle[k] = v
			}
			
			// Apply type-specific default configurations if they don't exist
			switch chartType {
			case "heatmap":
				if _, exists := chartStyle["heatmap"]; !exists {
					enhancedStyle["heatmap"] = map[string]interface{}{
						"colorScale":  "viridis",
						"showColorBar": true,
						"interpolate": true,
					}
				}
			case "3dsurface":
				if _, exists := chartStyle["3dsurface"]; !exists {
					enhancedStyle["3dsurface"] = map[string]interface{}{
						"wireframe": true,
						"colorScale": "jet",
						"rotation": map[string]interface{}{
							"x": 45,
							"y": 30,
							"z": 0,
						},
					}
				}
			case "waterfall":
				if _, exists := chartStyle["waterfall"]; !exists {
					enhancedStyle["waterfall"] = map[string]interface{}{
						"baseColor": "#1E90FF",
						"colorGradient": true,
						"spacing": 0.1,
						"perspective": 30,
					}
				}
			}
			
			enhancedConfig["chartStyle"] = enhancedStyle
		}
	}
	
	// Process chart style if present
	if chartStyle, ok := config["chartStyle"].(map[string]interface{}); ok {
		enhancedStyle := make(map[string]interface{})
		
		// Copy basic style properties
		for k, v := range chartStyle {
			enhancedStyle[k] = v
		}
		
		// Process axis configuration
		if axis, ok := chartStyle["axis"].(map[string]interface{}); ok {
			enhancedStyle["axisConfig"] = axis
		}
		
		// Process markers configuration
		if markers, ok := chartStyle["markers"].(map[string]interface{}); ok {
			enhancedStyle["markersConfig"] = markers
		}
		
		// Process grid configuration
		if grid, ok := chartStyle["grid"].(map[string]interface{}); ok {
			enhancedStyle["gridConfig"] = grid
		}
		
		// Process highlight regions
		if highlight, ok := chartStyle["highlight"].(map[string]interface{}); ok {
			enhancedStyle["highlightConfig"] = highlight
		}
		
		// Process label styles
		if labels, ok := chartStyle["labels"].(map[string]interface{}); ok {
			enhancedStyle["labelsConfig"] = labels
		}
		
		// Process legend configuration
		if legend, ok := chartStyle["legend"].(map[string]interface{}); ok {
			enhancedStyle["legendConfig"] = legend
		}
		
		// Process heatmap specific configuration
		if heatmap, ok := chartStyle["heatmap"].(map[string]interface{}); ok {
			enhancedStyle["heatmapConfig"] = heatmap
		}
		
		// Process 3D surface specific configuration
		if surface3d, ok := chartStyle["3dsurface"].(map[string]interface{}); ok {
			enhancedStyle["surface3dConfig"] = surface3d
		}
		
		// Process waterfall specific configuration
		if waterfall, ok := chartStyle["waterfall"].(map[string]interface{}); ok {
			enhancedStyle["waterfallConfig"] = waterfall
		}
		
		enhancedConfig["chartStyle"] = enhancedStyle
	}
	
	return enhancedConfig
}

// processData processes the collected data for the report
func (r *SpectrumReport) processData() error {
	// Set timestamp if not provided
	if _, ok := r.data["reportDate"]; !ok {
		r.data["reportDate"] = time.Now().Format("2006-01-02")
	}
	
	// Set generated timestamp
	r.data["generatedTimestamp"] = time.Now().Format("2006-01-02 15:04:05")
	
	return nil
}

// Factory function for creating SpectrumReport instances
func NewSpectrumReport() *SpectrumReport {
	return &SpectrumReport{}
}
